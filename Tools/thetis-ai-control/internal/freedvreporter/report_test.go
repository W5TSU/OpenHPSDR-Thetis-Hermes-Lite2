package freedvreporter

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

// TestEmitWireFormat exercises Client.Emit end-to-end over a net.Pipe,
// verifying the exact wire text a real reporter server would receive —
// mirrors internal/tci/control_test.go's TestSpotWireFormat pattern (same
// net.Pipe + ReadFrame technique, applied to this package's own wsConn).
func TestEmitWireFormat(t *testing.T) {
	clientEnd, serverEnd := net.Pipe()
	defer clientEnd.Close()
	defer serverEnd.Close()

	c := &Client{ws: &wsConn{tcp: clientEnd, r: bufio.NewReader(clientEnd), timeout: 2 * time.Second}}
	s := &wsConn{tcp: serverEnd, r: bufio.NewReader(serverEnd), timeout: 2 * time.Second}

	done := make(chan error, 1)
	go func() {
		op, payload, err := s.ReadFrame()
		if err != nil {
			done <- err
			return
		}
		want := `42["freq_change",{"freq":14236000}]`
		if op != opText || string(payload) != want {
			done <- fmt.Errorf("got opcode %d payload %q, want text %q", op, payload, want)
			return
		}
		done <- nil
	}()

	type freqChangePayload struct {
		Freq int64 `json:"freq"`
	}
	if err := c.Emit("freq_change", freqChangePayload{Freq: 14236000}); err != nil {
		t.Fatalf("Emit: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

// TestEmitWireFormatTxReport covers a second event shape (two fields, one
// of them a bool) so both of this feature's actual call shapes are tested,
// not just the single-field freq_change case.
func TestEmitWireFormatTxReport(t *testing.T) {
	clientEnd, serverEnd := net.Pipe()
	defer clientEnd.Close()
	defer serverEnd.Close()

	c := &Client{ws: &wsConn{tcp: clientEnd, r: bufio.NewReader(clientEnd), timeout: 2 * time.Second}}
	s := &wsConn{tcp: serverEnd, r: bufio.NewReader(serverEnd), timeout: 2 * time.Second}

	done := make(chan error, 1)
	go func() {
		op, payload, err := s.ReadFrame()
		if err != nil {
			done <- err
			return
		}
		want := `42["tx_report",{"mode":"USB","transmitting":true}]`
		if op != opText || string(payload) != want {
			done <- fmt.Errorf("got opcode %d payload %q, want text %q", op, payload, want)
			return
		}
		done <- nil
	}()

	type txReportPayload struct {
		Mode         string `json:"mode"`
		Transmitting bool   `json:"transmitting"`
	}
	if err := c.Emit("tx_report", txReportPayload{Mode: "USB", Transmitting: true}); err != nil {
		t.Fatalf("Emit: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}
