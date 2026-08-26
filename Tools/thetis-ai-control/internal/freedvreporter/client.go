package freedvreporter

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

// ReporterHost is qso.freedv.org's fixed hostname. Exported as a const
// (rather than hardcoded in cmd/) since it's a property of the protocol
// this package speaks, not something a caller should need to override.
const ReporterHost = "qso.freedv.org"

const socketIOPath = "/socket.io/?EIO=4&transport=websocket"

// Client is a connected, authenticated ("view" role) Socket.IO v4 session
// against FreeDV Reporter's live feed.
type Client struct {
	ws *wsConn
}

// Dial connects to FreeDV Reporter and completes both the Engine.IO and
// Socket.IO v4 handshakes. Protocol confirmed by direct probing 2026-08-09:
// the server accepts a direct WebSocket connection (no polling transport
// needed first) at ReporterHost's "/socket.io/?EIO=4&transport=websocket",
// immediately sends an Engine.IO OPEN packet ("0{...}"), and expects a
// Socket.IO CONNECT packet ("40" + JSON auth) in reply before it will start
// pushing station/activity events.
func Dial(timeout time.Duration) (*Client, error) {
	// The site's own client (index.js) sends { role: "view", protocol_version: 2 }
	// as the `auth` option — matched exactly here since an unrecognised
	// auth shape may be rejected server-side.
	auth, err := json.Marshal(struct {
		Role            string `json:"role"`
		ProtocolVersion int    `json:"protocol_version"`
	}{Role: "view", ProtocolVersion: 2})
	if err != nil {
		return nil, fmt.Errorf("freedvreporter: marshal auth: %w", err)
	}
	return dialAndHandshake(auth, timeout)
}

// DialReport connects to FreeDV Reporter in a self-reporting role instead
// of Dial's read-only "view" role, publishing this station's own presence
// rather than watching others'. Auth payload shape and role names
// confirmed directly against freedv-gui's own FreeDVReporter.cpp (fetched
// from its freedv_backend dependency) — role/callsign/grid_square/version/
// rx_only/os/protocol_version fields, NOT inferred from the read-only web
// viewer, which never emits station data at all. writeOnly selects
// "report_wo" over "report": the server then skips streaming the full
// station table back to this connection, which a Client obtained this way
// (Emit-only — see below) has no use for.
func DialReport(callsign, gridSquare string, rxOnly, writeOnly bool, timeout time.Duration) (*Client, error) {
	role := "report"
	if writeOnly {
		role = "report_wo"
	}
	auth, err := json.Marshal(struct {
		Role            string `json:"role"`
		Callsign        string `json:"callsign"`
		GridSquare      string `json:"grid_square"`
		Version         string `json:"version"`
		RXOnly          bool   `json:"rx_only"`
		OS              string `json:"os"`
		ProtocolVersion int    `json:"protocol_version"`
	}{
		Role:            role,
		Callsign:        callsign,
		GridSquare:      gridSquare,
		Version:         "Thetis (thetisctl)",
		RXOnly:          rxOnly,
		OS:              runtime.GOOS,
		ProtocolVersion: 2,
	})
	if err != nil {
		return nil, fmt.Errorf("freedvreporter: marshal report auth: %w", err)
	}
	return dialAndHandshake(auth, timeout)
}

// dialAndHandshake performs the connection plus Engine.IO/Socket.IO v4
// handshake shared by every role this package supports (Dial's own
// original doc comment described this exact sequence, confirmed by direct
// protocol probing 2026-08-09): the server accepts a direct WebSocket
// connection (no polling transport needed first) at ReporterHost's
// "/socket.io/?EIO=4&transport=websocket", immediately sends an Engine.IO
// OPEN packet ("0{...}"), and expects a Socket.IO CONNECT packet ("40" +
// JSON auth) in reply before it will start pushing (view role) or
// accepting (report role) events. auth is the already-marshaled JSON
// value for the "auth" option — its shape depends on the caller's role.
func dialAndHandshake(auth []byte, timeout time.Duration) (*Client, error) {
	ws, err := dialWS(ReporterHost, socketIOPath, timeout)
	if err != nil {
		return nil, err
	}

	// Engine.IO OPEN packet: "0{"sid":"...","pingInterval":...,...}". Not
	// otherwise needed (this client answers pings reactively rather than
	// tracking the advertised interval), but must be consumed before the
	// connection is usable.
	op, payload, err := ws.ReadFrame()
	if err != nil {
		ws.Close()
		return nil, fmt.Errorf("freedvreporter: read engine.io open packet: %w", err)
	}
	if op != opText || len(payload) == 0 || payload[0] != '0' {
		ws.Close()
		return nil, fmt.Errorf("freedvreporter: expected engine.io OPEN packet, got %q", truncate(payload, 80))
	}

	if err := ws.WriteText("40" + string(auth)); err != nil {
		ws.Close()
		return nil, fmt.Errorf("freedvreporter: send socket.io connect: %w", err)
	}

	// Socket.IO CONNECT ack: "40{"sid":"..."}" (a different sid than the
	// engine.io one above — this is the socket.io-level session).
	op, payload, err = ws.ReadFrame()
	if err != nil {
		ws.Close()
		return nil, fmt.Errorf("freedvreporter: read socket.io connect ack: %w", err)
	}
	if op != opText || len(payload) < 2 || string(payload[:2]) != "40" {
		ws.Close()
		return nil, fmt.Errorf("freedvreporter: expected socket.io CONNECT ack, got %q", truncate(payload, 80))
	}

	return &Client{ws: ws}, nil
}

// Close closes the underlying connection.
func (c *Client) Close() error {
	return c.ws.Close()
}

// Event is one (event name, raw JSON payload) pair as sent by the server.
// For "bulk_update", Payload is a JSON array of [name, data] pairs rather
// than a single event's data — see Tracker.Apply, which knows how to
// unwrap it. Every other event name's Payload is that event's own data
// object directly (matching the array's second element in the wire form
// `42["event_name", {...}]`).
type Event struct {
	Name    string
	Payload json.RawMessage
}

// ReadEvent blocks for the next Socket.IO event, transparently handling
// Engine.IO PING/PONG keepalive (the reporter's advertised pingInterval is
// an aggressive 5s — a caller that doesn't drain events promptly risks the
// server timing the connection out, since PONGs can only be sent from
// inside this read loop).
func (c *Client) ReadEvent() (Event, error) {
	for {
		op, payload, err := c.ws.ReadFrame()
		if err != nil {
			return Event{}, err
		}
		if op == opClose {
			return Event{}, fmt.Errorf("freedvreporter: connection closed by server")
		}
		if op != opText || len(payload) == 0 {
			continue
		}

		switch payload[0] {
		case '2': // Engine.IO PING -> reply PONG, keep waiting for a real event
			if werr := c.ws.WriteText("3"); werr != nil {
				return Event{}, fmt.Errorf("freedvreporter: send pong: %w", werr)
			}
			continue
		case '3': // Engine.IO PONG (shouldn't normally arrive; this client never pings) - ignore
			continue
		}

		if len(payload) < 2 || string(payload[:2]) != "42" {
			continue // some other Engine.IO/Socket.IO packet type - ignore
		}

		var arr []json.RawMessage
		if err := json.Unmarshal(payload[2:], &arr); err != nil || len(arr) == 0 {
			continue // malformed event frame - skip rather than fail the whole session
		}

		var name string
		if err := json.Unmarshal(arr[0], &name); err != nil {
			continue
		}

		var data json.RawMessage
		if len(arr) > 1 {
			data = arr[1]
		} else {
			data = json.RawMessage("null")
		}

		return Event{Name: name, Payload: data}, nil
	}
}

// Emit writes a Socket.IO v4 EVENT frame ("42" + a JSON array [event,
// payload]) — the exact inverse of ReadEvent's own "42"-prefix parsing
// above, so both directions of this package share one understanding of
// the wire format. Confirmed against freedv-gui's own FreeDVReporter.cpp
// emit call shapes (e.g. its freqChangeImpl_ builds {"freq": <Hz>} and
// calls sioClient_->emit("freq_change", ...)). Only meaningful on a
// Client obtained from DialReport — a view-role Client from Dial has
// nothing to emit and the server will not be listening for events from it.
func (c *Client) Emit(event string, payload any) error {
	b, err := json.Marshal([]any{event, payload})
	if err != nil {
		return fmt.Errorf("freedvreporter: marshal emit payload for %q: %w", event, err)
	}
	return c.ws.WriteText("42" + string(b))
}

func truncate(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n]) + "..."
}
