# FreeDV Tab: Reporter Enable UI + Self-Reporting Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Rename the Digital Voice tab to "FreeDV," add an in-Thetis UI that
launches/manages `thetisctl freedv-reporter watch` for both spotting and a
new self-reporting capability, and build that self-reporting capability
(publishing this station's own presence to qso.freedv.org) into `thetisctl`.

**Architecture:** Go side (`Tools/thetis-ai-control`): a new report-role
connection type in `internal/freedvreporter`, and a `--self-report` mode on
`cmd/thetisctl`'s existing `watch` subcommand that bridges Thetis's own live
TCI state (VFO/mode/TX) into reporter `Emit` calls via its own dedicated TCI
connection. C# side (`Project Files/Source/Console`): the tab rename, a new
UI group in Setup, and a `Process.Start`-based helper-process manager
following this codebase's existing `Dumpcap.cs` pattern.

**Tech Stack:** Go 1.22 (module `thetisctl`), Socket.IO v4 over TLS
WebSocket (hand-rolled, matching `internal/freedvreporter`'s existing
convention), C# / WinForms, `System.Diagnostics.Process`.

**Spec:** [docs/superpowers/specs/2026-08-26-freedv-tab-and-self-reporting-design.md](../specs/2026-08-26-freedv-tab-and-self-reporting-design.md)

## Global Constraints

- **CRLF-only files, plain-text Edit tool risk**: `Project Files/Source/Console/setup.cs`,
  `setup.designer.cs`, `console.cs`, and `Thetis.csproj` are all confirmed
  100% CRLF (verified this session — `console.cs`'s and `Thetis.csproj`'s
  status confirmed fresh while writing this plan, the other two unchanged
  since sub-project #5). Every edit to these four files in Task 3 below
  must be done via Python byte-level splicing
  (`data = open(path,'rb').read(); assert data.count(old_bytes) == 1; data
  = data.replace(old_bytes, new_bytes); open(path,'wb').write(data)`),
  verified with `git diff --stat` after every edit. **Go files
  (`Tools/thetis-ai-control/**`) are plain LF and are NOT subject to this —
  use the normal Edit tool there.**
- **Persistence direction, the exact lesson from sub-project #5's final
  review**: any code that syncs a checkbox's restored `.Checked` value with
  a backing config/state value at Setup-panel-init time must assign FROM
  the checkbox INTO the backing value (`SomeConfig.Value =
  chkSomething.Checked;`), never the reverse. The reverse direction
  silently overwrites whatever Setup's generic persistence mechanism just
  restored from the database with a hardcoded compile-time default,
  breaking "does this setting survive a restart" — exactly the bug caught
  in sub-project #5's final review. `setup.cs:37154`'s existing
  `Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;` is the
  correct pattern to copy.
- **No new native exports, no CAT protocol dependency**: this sub-project
  uses only Thetis's existing TCI wire commands (`vfo`, `modulation`,
  `trx`) queried in their existing GET form — no new C# TCIServer.cs
  changes, no CAT command additions.
- **`tx_report`'s `mode` field is a known v1 simplification**: it carries
  Thetis's raw DSPMode string (e.g. `"USB"`, `"DIGU"`) from a `modulation:`
  TCI query, not the actual active FreeDV codec name (`"RADEV1"`,
  `"700E"`) that a real freedv-gui report would show. Getting the real
  codec name would require also querying Thetis over CAT (the `ZZEX`
  family from sub-project #2), a new protocol dependency this plan
  deliberately does not add. Document this plainly in code comments and
  the final report; do not silently "fix" it into something more complex
  than this plan specifies.
- **Self-report gets its own dedicated TCI connection**, never sharing the
  `--spot`/auto-tune feature's existing `tciClient` variable. That
  variable is written to and nilled out by `--spot`/tune's own
  single-goroutine error-handling logic; adding a second (self-report)
  goroutine that reads from the same variable without synchronization
  would be a data race. Two independent TCP connections to the same local
  Thetis instance is cheap and simple; sharing one across goroutines is
  not.

---

### Task 1: `internal/freedvreporter` — report-role connection + `Emit`

**Files:**
- Modify: `Tools/thetis-ai-control/internal/freedvreporter/client.go`
- Create: `Tools/thetis-ai-control/internal/freedvreporter/report_test.go`

**Interfaces:**
- Consumes: `wsConn` (unexported, `ws.go`), `dialWS(host, path string,
  timeout time.Duration) (*wsConn, error)` (already exists), `(*wsConn)
  ReadFrame()`/`WriteText(s string) error` (already exist).
- Produces: `func DialReport(callsign, gridSquare string, rxOnly, writeOnly
  bool, timeout time.Duration) (*Client, error)` and `func (c *Client)
  Emit(event string, payload any) error` — Task 2 calls both.

- [ ] **Step 1: Refactor `Dial`'s handshake into a shared helper**

Read `internal/freedvreporter/client.go` first. Replace the entire current
body of `Dial` (everything from `func Dial(timeout time.Duration) (*Client,
error) {` through its matching closing `}`) with a version that delegates
to a new private helper, so `DialReport` (Step 2) can reuse the exact same
handshake without duplicating it:

Old code (the current, unmodified `Dial` function in full):
```go
func Dial(timeout time.Duration) (*Client, error) {
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

	// Socket.IO CONNECT to the default "/" namespace. The site's own client
	// (index.js) sends { role: "view", protocol_version: 2 } as the `auth`
	// option — matched exactly here since an unrecognised auth shape may be
	// rejected server-side.
	auth, err := json.Marshal(struct {
		Role            string `json:"role"`
		ProtocolVersion int    `json:"protocol_version"`
	}{Role: "view", ProtocolVersion: 2})
	if err != nil {
		ws.Close()
		return nil, fmt.Errorf("freedvreporter: marshal auth: %w", err)
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
```

New code (replaces the above exactly):
```go
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
```

- [ ] **Step 2: Add `Emit`**

In the same file, add this immediately after the `ReadEvent` method (i.e.
right after `ReadEvent`'s closing `}`, before `truncate`'s own doc
comment/definition):

```go
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
```

- [ ] **Step 3: Add the `runtime` import**

In the same file's import block, add `"runtime"` (alphabetically among the
existing imports):

```go
import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)
```

- [ ] **Step 4: Write the `Emit` wire-format test**

Create `Tools/thetis-ai-control/internal/freedvreporter/report_test.go`:

```go
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
```

- [ ] **Step 5: Run the new tests**

Run: `cd Tools/thetis-ai-control && go test ./internal/freedvreporter/... -run TestEmit -v`
Expected: both `TestEmitWireFormat` and `TestEmitWireFormatTxReport` PASS.

- [ ] **Step 6: Run the full non-live test suite and vet/format checks**

From `Tools/thetis-ai-control`:
```
gofmt -l internal/freedvreporter/client.go internal/freedvreporter/report_test.go
go vet ./...
go build ./...
go test ./...
```
Expected: `gofmt -l` prints nothing; the other three succeed (existing
`live`-tagged tests are skipped automatically).

- [ ] **Step 7: Commit**

```bash
cd Tools/thetis-ai-control
git add internal/freedvreporter/client.go internal/freedvreporter/report_test.go
git commit -m "feat(thetisctl): add report-role FreeDV Reporter connection + Emit

DialReport connects with role \"report\"/\"report_wo\" instead of
Dial's read-only \"view\" role, confirmed against freedv-gui's own
FreeDVReporter.cpp (fetched from its freedv_backend dependency), not
the read-only web viewer which never emits station data. Emit writes
the Socket.IO v4 EVENT frame this role needs to publish state -- the
exact inverse of ReadEvent's own wire-format parsing.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
```

---

### Task 2: `cmd/thetisctl` — `--self-report` bridge

**Files:**
- Modify: `Tools/thetis-ai-control/cmd/thetisctl/freedvreporter_cmd.go`
- Modify: `Tools/thetis-ai-control/cmd/thetisctl/args.go`

**Interfaces:**
- Consumes: `freedvreporter.DialReport(callsign, gridSquare string, rxOnly,
  writeOnly bool, timeout time.Duration) (*freedvreporter.Client, error)`
  and `(*freedvreporter.Client) Emit(event string, payload any) error`
  (Task 1); `tci.Client.SendCmd(cmd string, args ...string) error` and
  `tci.Client.RecvCmd() (cmd string, args []string, err error)` (already
  exist, `internal/tci/client.go`); `tci.Dial(addr string, timeout
  time.Duration) (*tci.Conn, error)` and `tci.NewClient(conn *tci.Conn)
  *tci.Client` (already exist).
- Produces: nothing further downstream — this is the outermost CLI layer.

- [ ] **Step 1: Register the new bool flags**

Read `Tools/thetis-ai-control/cmd/thetisctl/args.go` first. Change:

```go
var boolFlagNames = map[string]bool{
	"audio":   true,
	"spot":    true,
	"no-tune": true,
}
```

to:

```go
var boolFlagNames = map[string]bool{
	"audio":        true,
	"spot":         true,
	"no-tune":      true,
	"self-report":  true,
	"rx-only":      true,
}
```

- [ ] **Step 2: Run `gofmt` on `args.go` immediately**

`args.go` is plain LF Go source (no CRLF constraint), but map-literal
column alignment is easy to get wrong by hand. Run:
```
cd Tools/thetis-ai-control && gofmt -w cmd/thetisctl/args.go
```
This reformats the map's `:` alignment automatically if needed — safe to
run unconditionally on a Go file.

- [ ] **Step 3: Add self-report flags, validation, and the bridge**

Read `Tools/thetis-ai-control/cmd/thetisctl/freedvreporter_cmd.go` first.
Change:

```go
	tciHost := a.flag("tci", "")
	tciPort := a.flag("tci-port", "50001")
	tuneMode := a.flag("mode", "digu")
	doSpot := a.has("spot")
	noTune := a.has("no-tune")

	if doSpot && tciHost == "" {
		return fmt.Errorf("freedv-reporter watch: --spot requires --tci <host> (spots are pushed over TCI; there is nowhere else to send them)")
	}
```

to:

```go
	tciHost := a.flag("tci", "")
	tciPort := a.flag("tci-port", "50001")
	tuneMode := a.flag("mode", "digu")
	doSpot := a.has("spot")
	noTune := a.has("no-tune")
	doSelfReport := a.has("self-report")
	selfReportCallsign := a.flag("callsign", "")
	selfReportGrid := a.flag("grid", "")
	selfReportRxOnly := a.has("rx-only")

	if doSpot && tciHost == "" {
		return fmt.Errorf("freedv-reporter watch: --spot requires --tci <host> (spots are pushed over TCI; there is nowhere else to send them)")
	}
	if doSelfReport && tciHost == "" {
		return fmt.Errorf("freedv-reporter watch: --self-report requires --tci <host> (self-reporting reads Thetis's live VFO/mode/TX state over TCI)")
	}
	if doSelfReport && selfReportCallsign == "" {
		return fmt.Errorf("freedv-reporter watch: --self-report requires --callsign <call>")
	}
```

Then change the banner block:

```go
	if doSpot {
		fmt.Printf(" — pushing spots to %s panadapter", net.JoinHostPort(tciHost, tciPort))
	}
	fmt.Println(". Ctrl-C to stop.")
```

to:

```go
	if doSpot {
		fmt.Printf(" — pushing spots to %s panadapter", net.JoinHostPort(tciHost, tciPort))
	}
	if doSelfReport {
		fmt.Printf(" — reporting %s to FreeDV Reporter", selfReportCallsign)
	}
	fmt.Println(". Ctrl-C to stop.")
```

Then, right after the existing Ctrl-C-handling goroutine (immediately
before `tracker := freedvreporter.NewTracker()`), add the self-report
bridge's own goroutine launch. Change:

```go
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		fmt.Println("\nfreedv-reporter watch: stopping...")
		client.Close()
	}()

	tracker := freedvreporter.NewTracker()
```

to:

```go
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		fmt.Println("\nfreedv-reporter watch: stopping...")
		client.Close()
	}()

	if doSelfReport {
		go runSelfReport(tciHost, tciPort, selfReportCallsign, selfReportGrid, selfReportRxOnly)
	}

	tracker := freedvreporter.NewTracker()
```

(No explicit shutdown coordination is needed for this goroutine: when the
main loop's `client.ReadEvent()` errors out after Ctrl-C closes `client`
and `runFreeDVReporter` returns, the whole process exits and this
goroutine's own loop dies with it — matching how Go processes terminate
all goroutines on `main` returning.)

Finally, add the bridge's own implementation and a small helper. Append
this at the end of the file, after `parseInt64`'s closing `}`:

```go
// runSelfReport bridges Thetis's own live TCI state (VFO frequency, demod
// mode, TX/MOX) into FreeDV Reporter's write-side protocol, publishing
// this station's presence. It uses its OWN dedicated TCI connection,
// deliberately never sharing freedvReporterWatch's tciClient variable —
// that variable is written to and nilled out by the auto-tune/spot
// features' own single-goroutine error handling; a second goroutine
// reading it without synchronization would be a data race. Runs until the
// process exits (see the call site's own comment for why no explicit
// shutdown signal is threaded in here), reconnecting both sides on any
// error after a short pause.
//
// KNOWN SIMPLIFICATION: the tx_report "mode" field below carries Thetis's
// raw DSPMode string (e.g. "USB", "DIGU") from a plain "modulation:" TCI
// query, not the actual active FreeDV codec name ("RADEV1", "700E") a real
// freedv-gui report would show. Getting the real codec name needs a CAT
// query (the ZZEX family) in addition to TCI — a second protocol
// dependency deliberately out of scope for this pass.
func runSelfReport(tciHost, tciPort, callsign, gridSquare string, rxOnly bool) {
	for {
		addr := net.JoinHostPort(tciHost, tciPort)
		stateConn, derr := tci.Dial(addr, 5*time.Second)
		if derr != nil {
			fmt.Printf("[self-report] connect to Thetis TCI at %s failed: %v; retrying in 5s\n", addr, derr)
			time.Sleep(5 * time.Second)
			continue
		}
		stateClient := tci.NewClient(stateConn)

		reportClient, rerr := freedvreporter.DialReport(callsign, gridSquare, rxOnly, true, 10*time.Second)
		if rerr != nil {
			fmt.Printf("[self-report] connect to FreeDV Reporter failed: %v; retrying in 5s\n", rerr)
			stateClient.Close()
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Printf("[self-report] reporting %s to FreeDV Reporter\n", callsign)

		selfReportQueryInitialState(stateClient, reportClient)
		selfReportListenLoop(stateClient, reportClient)

		reportClient.Close()
		stateClient.Close()
		fmt.Println("[self-report] connection lost; reconnecting in 5s")
		time.Sleep(5 * time.Second)
	}
}

// selfReportQueryInitialState actively queries Thetis's current VFO A
// frequency, demod mode, and MOX state right after connecting (rather
// than only waiting passively for the next change), so the very first
// report reflects reality instead of a stale/empty value. Query wire
// forms confirmed directly against TCIServer.cs: "vfo:0,0;" (2 args) is
// the GET form replying "vfo:0,0,<hz>;" (TCIServer.cs:3859-3969,
// sendVFO/TCIServer.cs:2099-2132); "modulation:0;" (1 arg) is the GET form
// replying "modulation:0,<MODE>;" uppercase (TCIServer.cs:3972-4074,
// sendMode/TCIServer.cs:2174-2195); "trx:0;" (1 arg) is the GET form
// replying "trx:0,<true|false>[,tci];" (TCIServer.cs:3594-3694,
// sendMOX/TCIServer.cs:2159-2169). Errors here are logged and swallowed —
// a failed initial query just means the first report is skipped; the
// listen loop below will emit as soon as anything actually changes.
func selfReportQueryInitialState(stateClient *tci.Client, reportClient *freedvreporter.Client) {
	if err := stateClient.SendCmd("vfo", "0", "0"); err == nil {
		if cmd, args, rerr := stateClient.RecvCmd(); rerr == nil && cmd == "vfo" && len(args) >= 3 {
			if hz, perr := strconv.ParseInt(args[2], 10, 64); perr == nil {
				emitFreqChange(reportClient, hz)
			}
		}
	}
	if err := stateClient.SendCmd("modulation", "0"); err == nil {
		if cmd, args, rerr := stateClient.RecvCmd(); rerr == nil && cmd == "modulation" && len(args) >= 2 {
			lastKnownMode = args[1]
		}
	}
	if err := stateClient.SendCmd("trx", "0"); err == nil {
		if cmd, args, rerr := stateClient.RecvCmd(); rerr == nil && cmd == "trx" && len(args) >= 2 {
			if tx, perr := strconv.ParseBool(args[1]); perr == nil {
				emitTxReport(reportClient, lastKnownMode, tx)
			}
		}
	}
}

// lastKnownMode caches the most recently seen modulation string so a
// standalone "trx:" broadcast (which carries no mode of its own) can
// still emit a complete tx_report — mirrors freedv-gui's own
// FreeDVReporter::transmitImpl_, which similarly remembers mode_
// alongside tx_ state.
var lastKnownMode string

// selfReportListenLoop passively listens for Thetis's own unsolicited
// vfo:/modulation:/trx: broadcasts (sent whenever that state changes) and
// translates each into the matching Emit call. Returns (without erroring
// further) as soon as RecvCmd fails, letting the caller reconnect both
// connections.
func selfReportListenLoop(stateClient *tci.Client, reportClient *freedvreporter.Client) {
	for {
		cmd, args, err := stateClient.RecvCmd()
		if err != nil {
			fmt.Printf("[self-report] TCI connection lost: %v\n", err)
			return
		}
		switch cmd {
		case "vfo":
			if len(args) >= 3 {
				if hz, perr := strconv.ParseInt(args[2], 10, 64); perr == nil {
					emitFreqChange(reportClient, hz)
				}
			}
		case "modulation":
			if len(args) >= 2 {
				lastKnownMode = args[1]
			}
		case "trx":
			if len(args) >= 2 {
				if tx, perr := strconv.ParseBool(args[1]); perr == nil {
					emitTxReport(reportClient, lastKnownMode, tx)
				}
			}
		}
	}
}

func emitFreqChange(reportClient *freedvreporter.Client, hz int64) {
	type freqChangePayload struct {
		Freq int64 `json:"freq"`
	}
	if err := reportClient.Emit("freq_change", freqChangePayload{Freq: hz}); err != nil {
		fmt.Printf("[self-report] emit freq_change failed: %v\n", err)
	}
}

func emitTxReport(reportClient *freedvreporter.Client, mode string, transmitting bool) {
	type txReportPayload struct {
		Mode         string `json:"mode"`
		Transmitting bool   `json:"transmitting"`
	}
	if err := reportClient.Emit("tx_report", txReportPayload{Mode: mode, Transmitting: transmitting}); err != nil {
		fmt.Printf("[self-report] emit tx_report failed: %v\n", err)
	}
}
```

- [ ] **Step 4: Add the `strconv` import**

`freedvreporter_cmd.go`'s import block currently has no `strconv`. Change:

```go
import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"thetisctl/internal/freedvreporter"
	"thetisctl/internal/tci"
)
```

to:

```go
import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"thetisctl/internal/freedvreporter"
	"thetisctl/internal/tci"
)
```

- [ ] **Step 5: Verify it builds and vets cleanly**

From `Tools/thetis-ai-control`:
```
gofmt -l cmd/thetisctl/args.go cmd/thetisctl/freedvreporter_cmd.go
go vet ./...
go build ./...
go test ./...
```
Expected: `gofmt -l` prints nothing; the other three succeed.

- [ ] **Step 6: Manually verify the new flags parse correctly**

Run: `cd Tools/thetis-ai-control && go run ./cmd/thetisctl freedv-reporter watch --self-report`
Expected: fails immediately with exactly `freedv-reporter watch:
--self-report requires --tci <host> (self-reporting reads Thetis's live
VFO/mode/TX state over TCI)`.

Run: `go run ./cmd/thetisctl freedv-reporter watch --self-report --tci 127.0.0.1`
Expected: fails immediately with exactly `freedv-reporter watch:
--self-report requires --callsign <call>`.

- [ ] **Step 7: Commit**

```bash
cd Tools/thetis-ai-control
git add cmd/thetisctl/args.go cmd/thetisctl/freedvreporter_cmd.go
git commit -m "feat(thetisctl): add --self-report to freedv-reporter watch

Bridges Thetis's own live TCI state (VFO/mode/TX) into FreeDV
Reporter's write-side protocol via a dedicated TCI connection,
independent of --spot/auto-tune's own tciClient to avoid a
cross-goroutine data race. Actively queries current state on connect
so the first report isn't stale, then listens for Thetis's own
unsolicited state-change broadcasts. tx_report's mode field is a
known v1 simplification (Thetis's raw DSPMode, not the actual FreeDV
codec name) -- see the code comment for why.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
```

- [ ] **Step 8: Live verification against the real reporter**

This is a real deliverable, not optional — per the spec's own testing
section, the most convincing proof is the station actually appearing on
`https://qso.freedv.org`.

1. Determine `<thetis-host>` for this session first — the test box
   (`hl2winbox`/`hermes-pc`, same physical machine, two SSH aliases in
   `~/.ssh/config`) is reachable at different addresses depending on
   whether this session is on the local LAN or remote over VPN: LAN =
   `192.168.2.12`, VPN = `100.117.67.160`. Confirm which applies right now
   with `go run ./cmd/thetisctl cat --host 192.168.2.12 version` first; if
   that times out, try `go run ./cmd/thetisctl cat --host 100.117.67.160
   version` instead. Whichever responds is `<thetis-host>` for every step
   below. If Thetis isn't running at all (neither address responds even
   though `ssh hl2winbox`/`ssh hermes-pc` succeeds), relaunch via the
   scheduled-task interactive-session trick from sub-project #5 — `ssh
   hl2winbox` (or `hermes-pc`), then `schtasks /create /tn <name> /tr
   "'C:\Program Files\OpenHPSDR\Thetis-Test\Thetis.exe'" /sc onstart /ru
   mark /it /f`, then `schtasks /run /tn <name>`, then `schtasks /delete
   /tn <name> /f` to clean up.
2. Run (using a real or clearly-test callsign — check with whoever's
   running this session before using their real callsign on a public
   service, since this will be visible to other operators):
   ```
   timeout -s INT 60 go run ./cmd/thetisctl freedv-reporter watch \
     --tci <thetis-host> --self-report --callsign <CALLSIGN> --grid <GRID> --no-tune \
     > /tmp/self-report-live-check.log 2>&1
   echo "exit code: $?"
   cat /tmp/self-report-live-check.log
   ```
   (`--no-tune` isolates self-reporting from the pre-existing auto-tune
   behavior for this check; `timeout -s INT` sends the same signal Ctrl-C
   would, exercising the existing graceful-shutdown path.)
3. Confirm the log shows `reporting <CALLSIGN> to FreeDV Reporter` and no
   `emit ... failed` lines.
4. While it's running, change VFO A's frequency in Thetis (via CAT, e.g.
   `go run ./cmd/thetisctl cat --host <thetis-host> freq set a <hz>`, or by
   hand if you have visual/interactive access) and confirm the log or a
   direct look at `https://qso.freedv.org` shows the new frequency within
   a few seconds.
5. Confirm stopping the process (the `timeout -s INT` above, or Ctrl-C in
   an interactive run) removes the station from the reporter's map shortly
   after (the server's own disconnect-handling — no code here to verify
   beyond "did the process exit cleanly").

If reaching `https://qso.freedv.org` visually isn't possible this session,
the terminal-level confirmation (step 3, no emit failures, clean process
exit) is the part that must not be skipped — matching this project's
established precedent for when full visual confirmation isn't reachable
(sub-projects #4 and #5 both handled this the same way).

---

### Task 3: FreeDV tab rename + reporter enable UI + process management

**Files:**
- Modify: `Project Files/Source/Console/setup.designer.cs`
- Modify: `Project Files/Source/Console/setup.cs`
- Modify: `Project Files/Source/Console/console.cs:29015-29016`
- Modify: `Project Files/Source/Console/Thetis.csproj`
- Create: `Project Files/Source/Console/FreeDVReporterHelper.cs`

**Interfaces:**
- Consumes: nothing from Tasks 1-2 directly (Thetis never imports Go code)
  — it launches the CLI flags Tasks 1-2 define
  (`--spot`/`--self-report --callsign <c> --grid <g> [--rx-only] --tci
  <addr> [--no-tune]`) as command-line arguments to an external process.
  Also consumes `SetupForm.TCIOwnCallsign` (already exists, `setup.cs:12430-12433`).
- Produces: nothing further downstream — this is the final sub-project.

- [ ] **Step 1: Create the process-management helper class**

Following `Dumpcap.cs`'s existing exact pattern (static class,
`Process.Start` with `CreateNoWindow`, PID-tracked kill, no stdout/stderr
redirection — this codebase's only existing precedent for managing a
long-running external helper process). Create
`Project Files/Source/Console/FreeDVReporterHelper.cs`:

```csharp
using System;
using System.Diagnostics;

namespace Thetis
{
    // W5TSU: launches and supervises the external FreeDV Reporter helper
    // process (thetisctl or its thetis-ai-skill successor -- see
    // docs/superpowers/specs/2026-08-26-freedv-tab-and-self-reporting-design.md).
    // Follows Dumpcap.cs's existing exact pattern in this codebase: a
    // static class, Process.Start with CreateNoWindow, PID-tracked kill,
    // no stdout/stderr redirection (matching this project's only existing
    // precedent for supervising a long-running external process).
    public static class FreeDVReporterHelper
    {
        private static int m_nProcessID = -1;
        private const string PROCESS_NAME_NO_EXT = "thetisctl";

        public static bool IsRunning
        {
            get
            {
                if (m_nProcessID == -1) return false;

                bool bRet = false;
                Process[] proc = Process.GetProcessesByName(PROCESS_NAME_NO_EXT);
                foreach (Process p in proc)
                {
                    if (p.Id == m_nProcessID)
                    {
                        bRet = true;
                        break;
                    }
                }
                return bRet;
            }
        }

        // Starts (or restarts, if already running with different
        // arguments) the helper process. helperPath empty means resolve
        // "thetisctl" via PATH. Returns true if a process is running
        // afterward (either newly started or an unchanged prior instance
        // — callers pass the same arguments they'd want running now, so a
        // restart-on-every-call keeps this simple).
        public static bool EnsureRunning(string helperPath, string arguments)
        {
            Stop();

            try
            {
                string fileName = string.IsNullOrEmpty(helperPath) ? PROCESS_NAME_NO_EXT : helperPath;

                using (Process myProcess = new Process())
                {
                    myProcess.StartInfo.UseShellExecute = false;
                    myProcess.StartInfo.FileName = fileName;
                    myProcess.StartInfo.Arguments = arguments;
                    myProcess.StartInfo.CreateNoWindow = true;
                    myProcess.Start();
                    m_nProcessID = myProcess.Id;
                }
                return true;
            }
            catch
            {
                m_nProcessID = -1;
                return false;
            }
        }

        public static void Stop()
        {
            if (!IsRunning) return;

            Process[] proc = Process.GetProcessesByName(PROCESS_NAME_NO_EXT);
            foreach (Process p in proc)
            {
                if (p.Id == m_nProcessID)
                {
                    try
                    {
                        p.Kill();
                        m_nProcessID = -1;
                    }
                    catch { }
                    break;
                }
            }
        }
    }
}
```

- [ ] **Step 2: Add this new file to the project file**

The project file is `Project Files/Source/Console/Thetis.csproj` (verified
this session — not `Console.csproj`, despite the project's display name),
also 100% CRLF (verified: 1362 CRLF line endings, 0 LF-only), so this
edit uses the same Python byte-splicing discipline as the rest of this
task. Read the file first. Locate this exact, unique anchor:

Old bytes:
```
    <Compile Include="Dumpcap.cs" />
```

New bytes:
```
    <Compile Include="Dumpcap.cs" />
    <Compile Include="FreeDVReporterHelper.cs" />
```

```python
path = "Project Files/Source/Console/Thetis.csproj"
data = open(path, "rb").read()
old = b'    <Compile Include="Dumpcap.cs" />\r\n'
assert data.count(old) == 1, "anchor not found or not unique"
new = old + b'    <Compile Include="FreeDVReporterHelper.cs" />\r\n'
data = data.replace(old, new)
open(path, "wb").write(data)
```

Verify: `git diff --stat -- "Project Files/Source/Console/Thetis.csproj"` —
expect exactly +1/-0.

- [ ] **Step 3: Rename the tab and enable scrolling**

Read `setup.designer.cs` first. Using Python byte-splicing, locate this
exact, unique anchor:

Old bytes:
```
            this.tpDSPRADE.Size = new System.Drawing.Size(724, 414);
```

New bytes:
```
            this.tpDSPRADE.Size = new System.Drawing.Size(724, 414);
            this.tpDSPRADE.AutoScroll = true;
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()
old = b"            this.tpDSPRADE.Size = new System.Drawing.Size(724, 414);\r\n"
assert data.count(old) == 1, "anchor not found or not unique"
new = old + b"            this.tpDSPRADE.AutoScroll = true;\r\n"
data = data.replace(old, new)
open(path, "wb").write(data)
```

Then, in a second splice, locate:

Old bytes:
```
            this.tpDSPRADE.Text = "Digital Voice";
```

New bytes:
```
            this.tpDSPRADE.Text = "FreeDV";
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()
old = b'            this.tpDSPRADE.Text = "Digital Voice";\r\n'
assert data.count(old) == 1, "anchor not found or not unique"
new = b'            this.tpDSPRADE.Text = "FreeDV";\r\n'
data = data.replace(old, new)
open(path, "wb").write(data)
```

Verify: `git diff --stat -- "Project Files/Source/Console/setup.designer.cs"`
— expect roughly +2/-1 for these two splices combined.

- [ ] **Step 4: Add the new controls' object instantiation**

Read `setup.designer.cs` again. Locate the RX2 Core controls'
instantiation block (the same anchor sub-project #5 inserted next to):

Old bytes:
```
            this.chkShowRadeSyncOverlay = new System.Windows.Forms.CheckBoxTS();
```

New bytes:
```
            this.chkShowRadeSyncOverlay = new System.Windows.Forms.CheckBoxTS();
            this.grpFreeDVReporter = new System.Windows.Forms.GroupBoxTS();
            this.lblFreeDVReporterHelperPath = new System.Windows.Forms.LabelTS();
            this.txtFreeDVReporterHelperPath = new System.Windows.Forms.TextBoxTS();
            this.chkFreeDVReporterSpot = new System.Windows.Forms.CheckBoxTS();
            this.chkFreeDVReporterSelfReport = new System.Windows.Forms.CheckBoxTS();
            this.lblFreeDVReporterGrid = new System.Windows.Forms.LabelTS();
            this.txtFreeDVReporterGrid = new System.Windows.Forms.TextBoxTS();
            this.chkFreeDVReporterRxOnly = new System.Windows.Forms.CheckBoxTS();
            this.lblFreeDVReporterStatus = new System.Windows.Forms.LabelTS();
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()
old = b"            this.chkShowRadeSyncOverlay = new System.Windows.Forms.CheckBoxTS();\r\n"
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"            this.grpFreeDVReporter = new System.Windows.Forms.GroupBoxTS();\r\n"
    b"            this.lblFreeDVReporterHelperPath = new System.Windows.Forms.LabelTS();\r\n"
    b"            this.txtFreeDVReporterHelperPath = new System.Windows.Forms.TextBoxTS();\r\n"
    b"            this.chkFreeDVReporterSpot = new System.Windows.Forms.CheckBoxTS();\r\n"
    b"            this.chkFreeDVReporterSelfReport = new System.Windows.Forms.CheckBoxTS();\r\n"
    b"            this.lblFreeDVReporterGrid = new System.Windows.Forms.LabelTS();\r\n"
    b"            this.txtFreeDVReporterGrid = new System.Windows.Forms.TextBoxTS();\r\n"
    b"            this.chkFreeDVReporterRxOnly = new System.Windows.Forms.CheckBoxTS();\r\n"
    b"            this.lblFreeDVReporterStatus = new System.Windows.Forms.LabelTS();\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

Verify: `git diff --stat -- "Project Files/Source/Console/setup.designer.cs"`
grows by ~9 lines for this splice.

- [ ] **Step 5: Add the new controls' properties**

Read `setup.designer.cs` again. Locate the end of `chkShowRadeSyncOverlay`'s
own properties block (its `CheckedChanged +=` line is the last line of
that block):

Old bytes:
```
            this.chkShowRadeSyncOverlay.CheckedChanged += new System.EventHandler(this.chkShowRadeSyncOverlay_CheckedChanged);
```

New bytes (old bytes + the new group and all its children's property
blocks appended immediately after — group placed at (16, 356), size
(320, 190), below `chkShowRadeSyncOverlay` at (16, 328) which ends around
y=345; `tpDSPRADE.AutoScroll = true` from Step 3 lets this extend past the
tab's nominal 414px height):

```
            this.chkShowRadeSyncOverlay.CheckedChanged += new System.EventHandler(this.chkShowRadeSyncOverlay_CheckedChanged);
            //
            // grpFreeDVReporter
            //
            this.grpFreeDVReporter.Controls.Add(this.lblFreeDVReporterHelperPath);
            this.grpFreeDVReporter.Controls.Add(this.txtFreeDVReporterHelperPath);
            this.grpFreeDVReporter.Controls.Add(this.chkFreeDVReporterSpot);
            this.grpFreeDVReporter.Controls.Add(this.chkFreeDVReporterSelfReport);
            this.grpFreeDVReporter.Controls.Add(this.lblFreeDVReporterGrid);
            this.grpFreeDVReporter.Controls.Add(this.txtFreeDVReporterGrid);
            this.grpFreeDVReporter.Controls.Add(this.chkFreeDVReporterRxOnly);
            this.grpFreeDVReporter.Controls.Add(this.lblFreeDVReporterStatus);
            this.grpFreeDVReporter.Location = new System.Drawing.Point(16, 356);
            this.grpFreeDVReporter.Name = "grpFreeDVReporter";
            this.grpFreeDVReporter.Size = new System.Drawing.Size(320, 190);
            this.grpFreeDVReporter.TabIndex = 52;
            this.grpFreeDVReporter.TabStop = false;
            this.grpFreeDVReporter.Text = "FreeDV Reporter";
            //
            // lblFreeDVReporterHelperPath
            //
            this.lblFreeDVReporterHelperPath.AutoSize = true;
            this.lblFreeDVReporterHelperPath.Location = new System.Drawing.Point(12, 20);
            this.lblFreeDVReporterHelperPath.Name = "lblFreeDVReporterHelperPath";
            this.lblFreeDVReporterHelperPath.Size = new System.Drawing.Size(140, 13);
            this.lblFreeDVReporterHelperPath.TabIndex = 0;
            this.lblFreeDVReporterHelperPath.Text = "Helper (blank = use PATH):";
            //
            // txtFreeDVReporterHelperPath
            //
            this.txtFreeDVReporterHelperPath.Location = new System.Drawing.Point(12, 36);
            this.txtFreeDVReporterHelperPath.Name = "txtFreeDVReporterHelperPath";
            this.txtFreeDVReporterHelperPath.Size = new System.Drawing.Size(296, 20);
            this.txtFreeDVReporterHelperPath.TabIndex = 1;
            //
            // chkFreeDVReporterSpot
            //
            this.chkFreeDVReporterSpot.AutoSize = true;
            this.chkFreeDVReporterSpot.Image = null;
            this.chkFreeDVReporterSpot.Location = new System.Drawing.Point(12, 62);
            this.chkFreeDVReporterSpot.Name = "chkFreeDVReporterSpot";
            this.chkFreeDVReporterSpot.Size = new System.Drawing.Size(180, 17);
            this.chkFreeDVReporterSpot.TabIndex = 2;
            this.chkFreeDVReporterSpot.Text = "Show other stations on panadapter";
            this.toolTip1.SetToolTip(this.chkFreeDVReporterSpot, "Launches the FreeDV Reporter helper to push other active stations onto the panadapter as spot markers (sub-project 4).");
            this.chkFreeDVReporterSpot.UseVisualStyleBackColor = true;
            this.chkFreeDVReporterSpot.CheckedChanged += new System.EventHandler(this.chkFreeDVReporter_SettingChanged);
            //
            // chkFreeDVReporterSelfReport
            //
            this.chkFreeDVReporterSelfReport.AutoSize = true;
            this.chkFreeDVReporterSelfReport.Image = null;
            this.chkFreeDVReporterSelfReport.Location = new System.Drawing.Point(12, 84);
            this.chkFreeDVReporterSelfReport.Name = "chkFreeDVReporterSelfReport";
            this.chkFreeDVReporterSelfReport.Size = new System.Drawing.Size(200, 17);
            this.chkFreeDVReporterSelfReport.TabIndex = 3;
            this.chkFreeDVReporterSelfReport.Text = "Report my station to FreeDV Reporter";
            this.toolTip1.SetToolTip(this.chkFreeDVReporterSelfReport, "Publishes this station's callsign, grid square, frequency, mode, and TX state to qso.freedv.org so other operators see it on the reporter's live map.");
            this.chkFreeDVReporterSelfReport.UseVisualStyleBackColor = true;
            this.chkFreeDVReporterSelfReport.CheckedChanged += new System.EventHandler(this.chkFreeDVReporter_SettingChanged);
            //
            // lblFreeDVReporterGrid
            //
            this.lblFreeDVReporterGrid.AutoSize = true;
            this.lblFreeDVReporterGrid.Location = new System.Drawing.Point(12, 110);
            this.lblFreeDVReporterGrid.Name = "lblFreeDVReporterGrid";
            this.lblFreeDVReporterGrid.Size = new System.Drawing.Size(64, 13);
            this.lblFreeDVReporterGrid.TabIndex = 4;
            this.lblFreeDVReporterGrid.Text = "Grid square:";
            //
            // txtFreeDVReporterGrid
            //
            this.txtFreeDVReporterGrid.Location = new System.Drawing.Point(100, 106);
            this.txtFreeDVReporterGrid.Name = "txtFreeDVReporterGrid";
            this.txtFreeDVReporterGrid.Size = new System.Drawing.Size(80, 20);
            this.txtFreeDVReporterGrid.TabIndex = 5;
            //
            // chkFreeDVReporterRxOnly
            //
            this.chkFreeDVReporterRxOnly.AutoSize = true;
            this.chkFreeDVReporterRxOnly.Image = null;
            this.chkFreeDVReporterRxOnly.Location = new System.Drawing.Point(190, 108);
            this.chkFreeDVReporterRxOnly.Name = "chkFreeDVReporterRxOnly";
            this.chkFreeDVReporterRxOnly.Size = new System.Drawing.Size(105, 17);
            this.chkFreeDVReporterRxOnly.TabIndex = 6;
            this.chkFreeDVReporterRxOnly.Text = "I don\'t transmit";
            this.chkFreeDVReporterRxOnly.UseVisualStyleBackColor = true;
            this.chkFreeDVReporterRxOnly.CheckedChanged += new System.EventHandler(this.chkFreeDVReporter_SettingChanged);
            //
            // lblFreeDVReporterStatus
            //
            this.lblFreeDVReporterStatus.AutoSize = true;
            this.lblFreeDVReporterStatus.Location = new System.Drawing.Point(12, 136);
            this.lblFreeDVReporterStatus.Name = "lblFreeDVReporterStatus";
            this.lblFreeDVReporterStatus.Size = new System.Drawing.Size(100, 13);
            this.lblFreeDVReporterStatus.TabIndex = 7;
            this.lblFreeDVReporterStatus.Text = "Status: not running";
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()
old = b"            this.chkShowRadeSyncOverlay.CheckedChanged += new System.EventHandler(this.chkShowRadeSyncOverlay_CheckedChanged);\r\n"
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"            //\r\n"
    b"            // grpFreeDVReporter\r\n"
    b"            //\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.lblFreeDVReporterHelperPath);\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.txtFreeDVReporterHelperPath);\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.chkFreeDVReporterSpot);\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.chkFreeDVReporterSelfReport);\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.lblFreeDVReporterGrid);\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.txtFreeDVReporterGrid);\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.chkFreeDVReporterRxOnly);\r\n"
    b"            this.grpFreeDVReporter.Controls.Add(this.lblFreeDVReporterStatus);\r\n"
    b"            this.grpFreeDVReporter.Location = new System.Drawing.Point(16, 356);\r\n"
    b'            this.grpFreeDVReporter.Name = "grpFreeDVReporter";\r\n'
    b"            this.grpFreeDVReporter.Size = new System.Drawing.Size(320, 190);\r\n"
    b"            this.grpFreeDVReporter.TabIndex = 52;\r\n"
    b"            this.grpFreeDVReporter.TabStop = false;\r\n"
    b'            this.grpFreeDVReporter.Text = "FreeDV Reporter";\r\n'
    b"            //\r\n"
    b"            // lblFreeDVReporterHelperPath\r\n"
    b"            //\r\n"
    b"            this.lblFreeDVReporterHelperPath.AutoSize = true;\r\n"
    b"            this.lblFreeDVReporterHelperPath.Location = new System.Drawing.Point(12, 20);\r\n"
    b'            this.lblFreeDVReporterHelperPath.Name = "lblFreeDVReporterHelperPath";\r\n'
    b"            this.lblFreeDVReporterHelperPath.Size = new System.Drawing.Size(140, 13);\r\n"
    b"            this.lblFreeDVReporterHelperPath.TabIndex = 0;\r\n"
    b'            this.lblFreeDVReporterHelperPath.Text = "Helper (blank = use PATH):";\r\n'
    b"            //\r\n"
    b"            // txtFreeDVReporterHelperPath\r\n"
    b"            //\r\n"
    b"            this.txtFreeDVReporterHelperPath.Location = new System.Drawing.Point(12, 36);\r\n"
    b'            this.txtFreeDVReporterHelperPath.Name = "txtFreeDVReporterHelperPath";\r\n'
    b"            this.txtFreeDVReporterHelperPath.Size = new System.Drawing.Size(296, 20);\r\n"
    b"            this.txtFreeDVReporterHelperPath.TabIndex = 1;\r\n"
    b"            //\r\n"
    b"            // chkFreeDVReporterSpot\r\n"
    b"            //\r\n"
    b"            this.chkFreeDVReporterSpot.AutoSize = true;\r\n"
    b"            this.chkFreeDVReporterSpot.Image = null;\r\n"
    b"            this.chkFreeDVReporterSpot.Location = new System.Drawing.Point(12, 62);\r\n"
    b'            this.chkFreeDVReporterSpot.Name = "chkFreeDVReporterSpot";\r\n'
    b"            this.chkFreeDVReporterSpot.Size = new System.Drawing.Size(180, 17);\r\n"
    b"            this.chkFreeDVReporterSpot.TabIndex = 2;\r\n"
    b'            this.chkFreeDVReporterSpot.Text = "Show other stations on panadapter";\r\n'
    b'            this.toolTip1.SetToolTip(this.chkFreeDVReporterSpot, "Launches the FreeDV Reporter helper to push other active stations onto the panadapter as spot markers (sub-project 4).");\r\n'
    b"            this.chkFreeDVReporterSpot.UseVisualStyleBackColor = true;\r\n"
    b"            this.chkFreeDVReporterSpot.CheckedChanged += new System.EventHandler(this.chkFreeDVReporter_SettingChanged);\r\n"
    b"            //\r\n"
    b"            // chkFreeDVReporterSelfReport\r\n"
    b"            //\r\n"
    b"            this.chkFreeDVReporterSelfReport.AutoSize = true;\r\n"
    b"            this.chkFreeDVReporterSelfReport.Image = null;\r\n"
    b"            this.chkFreeDVReporterSelfReport.Location = new System.Drawing.Point(12, 84);\r\n"
    b'            this.chkFreeDVReporterSelfReport.Name = "chkFreeDVReporterSelfReport";\r\n'
    b"            this.chkFreeDVReporterSelfReport.Size = new System.Drawing.Size(200, 17);\r\n"
    b"            this.chkFreeDVReporterSelfReport.TabIndex = 3;\r\n"
    b'            this.chkFreeDVReporterSelfReport.Text = "Report my station to FreeDV Reporter";\r\n'
    b'            this.toolTip1.SetToolTip(this.chkFreeDVReporterSelfReport, "Publishes this station\'s callsign, grid square, frequency, mode, and TX state to qso.freedv.org so other operators see it on the reporter\'s live map.");\r\n'
    b"            this.chkFreeDVReporterSelfReport.UseVisualStyleBackColor = true;\r\n"
    b"            this.chkFreeDVReporterSelfReport.CheckedChanged += new System.EventHandler(this.chkFreeDVReporter_SettingChanged);\r\n"
    b"            //\r\n"
    b"            // lblFreeDVReporterGrid\r\n"
    b"            //\r\n"
    b"            this.lblFreeDVReporterGrid.AutoSize = true;\r\n"
    b"            this.lblFreeDVReporterGrid.Location = new System.Drawing.Point(12, 110);\r\n"
    b'            this.lblFreeDVReporterGrid.Name = "lblFreeDVReporterGrid";\r\n'
    b"            this.lblFreeDVReporterGrid.Size = new System.Drawing.Size(64, 13);\r\n"
    b"            this.lblFreeDVReporterGrid.TabIndex = 4;\r\n"
    b'            this.lblFreeDVReporterGrid.Text = "Grid square:";\r\n'
    b"            //\r\n"
    b"            // txtFreeDVReporterGrid\r\n"
    b"            //\r\n"
    b"            this.txtFreeDVReporterGrid.Location = new System.Drawing.Point(100, 106);\r\n"
    b'            this.txtFreeDVReporterGrid.Name = "txtFreeDVReporterGrid";\r\n'
    b"            this.txtFreeDVReporterGrid.Size = new System.Drawing.Size(80, 20);\r\n"
    b"            this.txtFreeDVReporterGrid.TabIndex = 5;\r\n"
    b"            //\r\n"
    b"            // chkFreeDVReporterRxOnly\r\n"
    b"            //\r\n"
    b"            this.chkFreeDVReporterRxOnly.AutoSize = true;\r\n"
    b"            this.chkFreeDVReporterRxOnly.Image = null;\r\n"
    b"            this.chkFreeDVReporterRxOnly.Location = new System.Drawing.Point(190, 108);\r\n"
    b'            this.chkFreeDVReporterRxOnly.Name = "chkFreeDVReporterRxOnly";\r\n'
    b"            this.chkFreeDVReporterRxOnly.Size = new System.Drawing.Size(105, 17);\r\n"
    b"            this.chkFreeDVReporterRxOnly.TabIndex = 6;\r\n"
    b'            this.chkFreeDVReporterRxOnly.Text = "I don\'t transmit";\r\n'
    b"            this.chkFreeDVReporterRxOnly.UseVisualStyleBackColor = true;\r\n"
    b"            this.chkFreeDVReporterRxOnly.CheckedChanged += new System.EventHandler(this.chkFreeDVReporter_SettingChanged);\r\n"
    b"            //\r\n"
    b"            // lblFreeDVReporterStatus\r\n"
    b"            //\r\n"
    b"            this.lblFreeDVReporterStatus.AutoSize = true;\r\n"
    b"            this.lblFreeDVReporterStatus.Location = new System.Drawing.Point(12, 136);\r\n"
    b'            this.lblFreeDVReporterStatus.Name = "lblFreeDVReporterStatus";\r\n'
    b"            this.lblFreeDVReporterStatus.Size = new System.Drawing.Size(100, 13);\r\n"
    b"            this.lblFreeDVReporterStatus.TabIndex = 7;\r\n"
    b'            this.lblFreeDVReporterStatus.Text = "Status: not running";\r\n'
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

- [ ] **Step 6: Add the group to the tab and declare the fields**

Read `setup.designer.cs` again. Locate the tab's `Controls.Add` list:

Old bytes:
```
            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX2Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
            this.tpDSPRADE.Controls.Add(this.chkShowRadeSyncOverlay);
```

New bytes:
```
            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX2Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
            this.tpDSPRADE.Controls.Add(this.chkShowRadeSyncOverlay);
            this.tpDSPRADE.Controls.Add(this.grpFreeDVReporter);
```

Then locate the field declarations block:

Old bytes:
```
        private CheckBoxTS chkShowRadeSyncOverlay;
```

New bytes:
```
        private CheckBoxTS chkShowRadeSyncOverlay;
        private GroupBoxTS grpFreeDVReporter;
        private LabelTS lblFreeDVReporterHelperPath;
        private TextBoxTS txtFreeDVReporterHelperPath;
        private CheckBoxTS chkFreeDVReporterSpot;
        private CheckBoxTS chkFreeDVReporterSelfReport;
        private LabelTS lblFreeDVReporterGrid;
        private TextBoxTS txtFreeDVReporterGrid;
        private CheckBoxTS chkFreeDVReporterRxOnly;
        private LabelTS lblFreeDVReporterStatus;
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()

old1 = (
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);\r\n"
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);\r\n"
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeRX2Core);\r\n"
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);\r\n"
    b"            this.tpDSPRADE.Controls.Add(this.chkShowRadeSyncOverlay);\r\n"
)
assert data.count(old1) == 1, "anchor 1 not found or not unique"
new1 = old1 + b"            this.tpDSPRADE.Controls.Add(this.grpFreeDVReporter);\r\n"
data = data.replace(old1, new1)

old2 = b"        private CheckBoxTS chkShowRadeSyncOverlay;\r\n"
assert data.count(old2) == 1, "anchor 2 not found or not unique"
new2 = old2 + (
    b"        private GroupBoxTS grpFreeDVReporter;\r\n"
    b"        private LabelTS lblFreeDVReporterHelperPath;\r\n"
    b"        private TextBoxTS txtFreeDVReporterHelperPath;\r\n"
    b"        private CheckBoxTS chkFreeDVReporterSpot;\r\n"
    b"        private CheckBoxTS chkFreeDVReporterSelfReport;\r\n"
    b"        private LabelTS lblFreeDVReporterGrid;\r\n"
    b"        private TextBoxTS txtFreeDVReporterGrid;\r\n"
    b"        private CheckBoxTS chkFreeDVReporterRxOnly;\r\n"
    b"        private LabelTS lblFreeDVReporterStatus;\r\n"
)
data = data.replace(old2, new2)

open(path, "wb").write(data)
```

- [ ] **Step 7: Verify the designer changes with git diff**

`git diff --stat -- "Project Files/Source/Console/setup.designer.cs"` —
cumulative across Steps 3-6: expect on the order of +100 to +110 lines
total (9 instantiations + ~85 property lines + 1 Controls.Add + 9 field
declarations + 2 rename/autoscroll lines). Recover via `git checkout` +
redo if the diff is far larger (thousands of lines = flattened line
endings).

- [ ] **Step 8: Add the behavior in `setup.cs`**

Read `setup.cs` first. Locate the persistence-sync line sub-project #5
added (the exact, correct-direction pattern to copy):

Old bytes:
```
            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;
```

New bytes (old bytes + the new group's own restore-sync, in the SAME
correct direction — checkbox/text-field state INTO whatever backs it;
since this feature has no backing config object of its own yet beyond the
controls' own Setup-persisted values, this sync is about kicking off
`EnsureRunning`/`Stop` to match whatever Setup's generic mechanism just
restored into the checkboxes, not about a separate `Display`-style flag):

```
            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;

            chkFreeDVReporterSpot.CheckedChanged -= chkFreeDVReporter_SettingChanged;
            chkFreeDVReporterSelfReport.CheckedChanged -= chkFreeDVReporter_SettingChanged;
            chkFreeDVReporterRxOnly.CheckedChanged -= chkFreeDVReporter_SettingChanged;
            ApplyFreeDVReporterSettings();
            chkFreeDVReporterSpot.CheckedChanged += chkFreeDVReporter_SettingChanged;
            chkFreeDVReporterSelfReport.CheckedChanged += chkFreeDVReporter_SettingChanged;
            chkFreeDVReporterRxOnly.CheckedChanged += chkFreeDVReporter_SettingChanged;
```

```python
path = "Project Files/Source/Console/setup.cs"
data = open(path, "rb").read()
old = b"            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;\r\n"
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"\r\n"
    b"            chkFreeDVReporterSpot.CheckedChanged -= chkFreeDVReporter_SettingChanged;\r\n"
    b"            chkFreeDVReporterSelfReport.CheckedChanged -= chkFreeDVReporter_SettingChanged;\r\n"
    b"            chkFreeDVReporterRxOnly.CheckedChanged -= chkFreeDVReporter_SettingChanged;\r\n"
    b"            ApplyFreeDVReporterSettings();\r\n"
    b"            chkFreeDVReporterSpot.CheckedChanged += chkFreeDVReporter_SettingChanged;\r\n"
    b"            chkFreeDVReporterSelfReport.CheckedChanged += chkFreeDVReporter_SettingChanged;\r\n"
    b"            chkFreeDVReporterRxOnly.CheckedChanged += chkFreeDVReporter_SettingChanged;\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

Note this reads the checkboxes' own already-Setup-restored `.Checked`
values (via `ApplyFreeDVReporterSettings`, defined in Step 9) rather than
writing INTO them from some separate backing value — exactly the
correct direction per this plan's Global Constraint, since (unlike
`ShowRadeSyncOverlay`) there is no separate `Display`-style static flag
behind these three checkboxes; the checkboxes' own Setup-persisted
`.Checked` state IS the source of truth, and `InitRadePanelFromBackend`
must read from it, never overwrite it.

- [ ] **Step 9: Add the settings-changed handler and status timer**

Read `setup.cs` again. Locate `chkShowRadeSyncOverlay_CheckedChanged`'s
full existing body (added in sub-project #5):

Old bytes:
```
        // W5TSU: panadapter sync/SNR overlay toggle (sub-project 5 of 6) --
        // one-line passthrough to Display's static flag, guarded like every
        // other Digital Voice tab handler against firing during Setup's own
        // startup sequence.
        private void chkShowRadeSyncOverlay_CheckedChanged(object sender, EventArgs e)
        {
            if (initializing) return;
            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;
        }
```

New bytes (old bytes + the new handler, status timer, and
argument-building helper appended immediately after):

```
        // W5TSU: panadapter sync/SNR overlay toggle (sub-project 5 of 6) --
        // one-line passthrough to Display's static flag, guarded like every
        // other Digital Voice tab handler against firing during Setup's own
        // startup sequence.
        private void chkShowRadeSyncOverlay_CheckedChanged(object sender, EventArgs e)
        {
            if (initializing) return;
            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;
        }

        // W5TSU: FreeDV Reporter enable UI (sub-project 6 of 6, see
        // docs/superpowers/specs/2026-08-26-freedv-tab-and-self-reporting-design.md).
        // One shared handler for all three checkboxes -- any of them
        // changing means the helper process needs restarting with the
        // current combined argument set (or stopped, if none are checked
        // anymore). Grid square/RX-only text-field edits are picked up the
        // next time any checkbox changes, matching the spec's "changing a
        // setting while running restarts it" behavior without needing a
        // separate TextChanged-triggered restart for every keystroke.
        private void chkFreeDVReporter_SettingChanged(object sender, EventArgs e)
        {
            if (initializing) return;
            ApplyFreeDVReporterSettings();
        }

        private System.Windows.Forms.Timer _freedv_reporter_status_timer = null;

        // W5TSU: starts/stops/restarts FreeDVReporterHelper to match the
        // three checkboxes' current state, and (re)arms the status-label
        // poll timer. Called both from InitRadePanelFromBackend (Setup
        // startup, reading whatever Setup's generic persistence mechanism
        // already restored into the checkboxes) and from
        // chkFreeDVReporter_SettingChanged (a live operator toggle) --
        // both cases want exactly the same "make the running process match
        // what the checkboxes say" behavior.
        private void ApplyFreeDVReporterSettings()
        {
            bool doSpot = chkFreeDVReporterSpot.Checked;
            bool doSelfReport = chkFreeDVReporterSelfReport.Checked;

            if (!doSpot && !doSelfReport)
            {
                FreeDVReporterHelper.Stop();
                lblFreeDVReporterStatus.Text = "Status: not running";
                if (_freedv_reporter_status_timer != null)
                    _freedv_reporter_status_timer.Enabled = false;
                return;
            }

            // --no-tune is always passed: this UI never exposes the
            // pre-existing "auto-tune my own radio to other stations'
            // activity" behavior watch already has baked in whenever
            // --tci is given at all, and enabling that as an undeclared
            // side effect of checking either of THESE checkboxes would be
            // a surprise -- an operator who only wants self-reporting
            // does not also want their own VFO silently retuning to
            // someone else's transmission.
            string args = "freedv-reporter watch --tci 127.0.0.1 --no-tune";
            if (doSpot)
                args += " --spot";
            if (doSelfReport)
            {
                string callsign = SetupForm.TCIOwnCallsign;
                string grid = txtFreeDVReporterGrid.Text;
                args += " --self-report --callsign " + callsign + " --grid " + grid;
                if (chkFreeDVReporterRxOnly.Checked)
                    args += " --rx-only";
            }

            bool started = FreeDVReporterHelper.EnsureRunning(txtFreeDVReporterHelperPath.Text, args);
            lblFreeDVReporterStatus.Text = started ? "Status: running" : "Status: failed to start";

            if (_freedv_reporter_status_timer == null)
            {
                _freedv_reporter_status_timer = new System.Windows.Forms.Timer();
                _freedv_reporter_status_timer.Interval = 2000;
                _freedv_reporter_status_timer.Tick += freedvReporterStatusTimer_Tick;
            }
            _freedv_reporter_status_timer.Enabled = true;
        }

        private void freedvReporterStatusTimer_Tick(object sender, EventArgs e)
        {
            lblFreeDVReporterStatus.Text = FreeDVReporterHelper.IsRunning ? "Status: running" : "Status: not running";
        }
```

```python
path = "Project Files/Source/Console/setup.cs"
data = open(path, "rb").read()
old = (
    b"        // W5TSU: panadapter sync/SNR overlay toggle (sub-project 5 of 6) --\r\n"
    b"        // one-line passthrough to Display's static flag, guarded like every\r\n"
    b"        // other Digital Voice tab handler against firing during Setup's own\r\n"
    b"        // startup sequence.\r\n"
    b"        private void chkShowRadeSyncOverlay_CheckedChanged(object sender, EventArgs e)\r\n"
    b"        {\r\n"
    b"            if (initializing) return;\r\n"
    b"            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;\r\n"
    b"        }\r\n"
)
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"\r\n"
    b"        // W5TSU: FreeDV Reporter enable UI (sub-project 6 of 6, see\r\n"
    b"        // docs/superpowers/specs/2026-08-26-freedv-tab-and-self-reporting-design.md).\r\n"
    b"        // One shared handler for all three checkboxes -- any of them\r\n"
    b"        // changing means the helper process needs restarting with the\r\n"
    b"        // current combined argument set (or stopped, if none are checked\r\n"
    b"        // anymore). Grid square/RX-only text-field edits are picked up the\r\n"
    b"        // next time any checkbox changes, matching the spec's \"changing a\r\n"
    b"        // setting while running restarts it\" behavior without needing a\r\n"
    b"        // separate TextChanged-triggered restart for every keystroke.\r\n"
    b"        private void chkFreeDVReporter_SettingChanged(object sender, EventArgs e)\r\n"
    b"        {\r\n"
    b"            if (initializing) return;\r\n"
    b"            ApplyFreeDVReporterSettings();\r\n"
    b"        }\r\n"
    b"\r\n"
    b"        private System.Windows.Forms.Timer _freedv_reporter_status_timer = null;\r\n"
    b"\r\n"
    b"        // W5TSU: starts/stops/restarts FreeDVReporterHelper to match the\r\n"
    b"        // three checkboxes' current state, and (re)arms the status-label\r\n"
    b"        // poll timer. Called both from InitRadePanelFromBackend (Setup\r\n"
    b"        // startup, reading whatever Setup's generic persistence mechanism\r\n"
    b"        // already restored into the checkboxes) and from\r\n"
    b"        // chkFreeDVReporter_SettingChanged (a live operator toggle) --\r\n"
    b"        // both cases want exactly the same \"make the running process match\r\n"
    b"        // what the checkboxes say\" behavior.\r\n"
    b"        private void ApplyFreeDVReporterSettings()\r\n"
    b"        {\r\n"
    b"            bool doSpot = chkFreeDVReporterSpot.Checked;\r\n"
    b"            bool doSelfReport = chkFreeDVReporterSelfReport.Checked;\r\n"
    b"\r\n"
    b"            if (!doSpot && !doSelfReport)\r\n"
    b"            {\r\n"
    b"                FreeDVReporterHelper.Stop();\r\n"
    b'                lblFreeDVReporterStatus.Text = "Status: not running";\r\n'
    b"                if (_freedv_reporter_status_timer != null)\r\n"
    b"                    _freedv_reporter_status_timer.Enabled = false;\r\n"
    b"                return;\r\n"
    b"            }\r\n"
    b"\r\n"
    b"            // --no-tune is always passed: this UI never exposes the\r\n"
    b"            // pre-existing \"auto-tune my own radio to other stations'\r\n"
    b'            // activity" behavior watch already has baked in whenever\r\n'
    b"            // --tci is given at all, and enabling that as an undeclared\r\n"
    b"            // side effect of checking either of THESE checkboxes would be\r\n"
    b"            // a surprise -- an operator who only wants self-reporting\r\n"
    b"            // does not also want their own VFO silently retuning to\r\n"
    b"            // someone else's transmission.\r\n"
    b'            string args = "freedv-reporter watch --tci 127.0.0.1 --no-tune";\r\n'
    b"            if (doSpot)\r\n"
    b'                args += " --spot";\r\n'
    b"            if (doSelfReport)\r\n"
    b"            {\r\n"
    b"                string callsign = SetupForm.TCIOwnCallsign;\r\n"
    b"                string grid = txtFreeDVReporterGrid.Text;\r\n"
    b'                args += " --self-report --callsign " + callsign + " --grid " + grid;\r\n'
    b"                if (chkFreeDVReporterRxOnly.Checked)\r\n"
    b'                    args += " --rx-only";\r\n'
    b"            }\r\n"
    b"\r\n"
    b"            bool started = FreeDVReporterHelper.EnsureRunning(txtFreeDVReporterHelperPath.Text, args);\r\n"
    b'            lblFreeDVReporterStatus.Text = started ? "Status: running" : "Status: failed to start";\r\n'
    b"\r\n"
    b"            if (_freedv_reporter_status_timer == null)\r\n"
    b"            {\r\n"
    b"                _freedv_reporter_status_timer = new System.Windows.Forms.Timer();\r\n"
    b"                _freedv_reporter_status_timer.Interval = 2000;\r\n"
    b"                _freedv_reporter_status_timer.Tick += freedvReporterStatusTimer_Tick;\r\n"
    b"            }\r\n"
    b"            _freedv_reporter_status_timer.Enabled = true;\r\n"
    b"        }\r\n"
    b"\r\n"
    b"        private void freedvReporterStatusTimer_Tick(object sender, EventArgs e)\r\n"
    b"        {\r\n"
    b'            lblFreeDVReporterStatus.Text = FreeDVReporterHelper.IsRunning ? "Status: running" : "Status: not running";\r\n'
    b"        }\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

- [ ] **Step 10: Hook process teardown into the main console's close path**

**Not Setup's own close path** — `Setup_Closing` (`setup.cs:11759-11777`)
always sets `e.Cancel = true` and calls `this.Hide()`; the Setup *window*
is a singleton that's hidden, never actually closed, for the life of the
Thetis process (confirmed by reading it directly) — the helper process
must keep running while an operator merely closes the Setup dialog, so
this is the wrong hook.

The correct hook, confirmed directly: `console.cs`'s `Console_Closing`
(the *main* window's close handler) already tears down another
long-running external-process helper, `Dumpcap`, in exactly the shape
needed here:

Old bytes:
```
            shutdownLogStringToPath("Before DumpCap.StopDumpcap()");
            DumpCap.StopDumpcap();
```

New bytes:
```
            shutdownLogStringToPath("Before DumpCap.StopDumpcap()");
            DumpCap.StopDumpcap();

            shutdownLogStringToPath("Before FreeDVReporterHelper.Stop()");
            FreeDVReporterHelper.Stop();
```

```python
path = "Project Files/Source/Console/console.cs"
data = open(path, "rb").read()
old = (
    b'            shutdownLogStringToPath("Before DumpCap.StopDumpcap()");\r\n'
    b"            DumpCap.StopDumpcap();\r\n"
)
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"\r\n"
    b'            shutdownLogStringToPath("Before FreeDVReporterHelper.Stop()");\r\n'
    b"            FreeDVReporterHelper.Stop();\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

Verify: `git diff --stat -- "Project Files/Source/Console/console.cs"` —
expect exactly +3/-0.

- [ ] **Step 11: Verify with git diff**

`git diff --stat -- "Project Files/Source/Console/setup.cs"` — cumulative
across Steps 8-9: expect on the order of +65-75 lines (Step 8's ~8 lines,
Step 9's ~55 lines). `git diff --stat -- "Project Files/Source/Console/console.cs"`
— expect +3/-0 for Step 10. Recover via `git checkout` + redo if either is
far larger than expected.

- [ ] **Step 12: Build, deploy, and verify live on hl2winbox**

This is a real deliverable, not optional:

1. `git add` and commit all three files' changes (designer, setup.cs,
   the new `FreeDVReporterHelper.cs`, and the `.csproj` edit) — see commit
   message below.
2. `git push`, then `gh workflow run build.yml --ref FreeDV -R
   W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`, poll with `gh run list
   --workflow=build.yml -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 --limit 3`
   / `gh run watch <run-id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2` until
   it completes.
3. Download the `Thetis-HL2-installer` artifact, `scp` it to
   `hl2winbox:Downloads/`, admin-extract via `msiexec /a <msi> /qn
   TARGETDIR=<dir>`, `robocopy <extracted>\OpenHPSDR\Thetis-Test
   "C:\Program Files\OpenHPSDR\Thetis-Test" /MIR` (this exact sequence,
   confirmed working, was used repeatedly in sub-project #5 — see that
   plan's own Task 2 Step 8 for the full command forms).
4. Relaunch via the scheduled-task interactive-session trick established
   in sub-project #5 (`schtasks /create /tn <name> /tr "'C:\Program
   Files\OpenHPSDR\Thetis-Test\Thetis.exe'" /sc onstart /ru mark /it /f`,
   then `/run`, then `/delete` to clean up) — a plain SSH-launched
   `Start-Process` lands in a non-interactive session and crashes on this
   app's own early `MessageBox.Show` call, confirmed in sub-project #5.
5. Verify the new build is live: `cd Tools/thetis-ai-control && go run
   ./cmd/thetisctl cat --host <thetis-host> version` should report a
   `git:` short SHA matching this task's own commit — determine
   `<thetis-host>` the same way as Task 2 Step 8.1 (LAN `192.168.2.12` vs.
   VPN `100.117.67.160`, whichever responds; this can change between
   sessions depending on how this box is reached at the time).
6. Confirm visually (screenshot via the same interactive-session trick, or
   in-person/RDP access if available) that the tab now reads "FreeDV" and
   the new "FreeDV Reporter" group appears with its two checkboxes, grid
   field, RX-only checkbox, and status label.
7. Check the "Show other stations on panadapter" checkbox with a helper
   binary available (either build `thetisctl.exe` for Windows separately
   and place it on the box's PATH, or point `txtFreeDVReporterHelperPath`
   at a path where you've placed it) and confirm the status label shows
   "Status: running" and `Process.GetProcessesByName("thetisctl")` (check
   over SSH) shows it actually running.
8. Uncheck it; confirm the status label reverts to "Status: not running"
   and the process is gone.
9. Restart Thetis (same relaunch trick) with the checkbox left checked
   before the restart; confirm it comes back checked and the helper
   process is running again after restart — this is the exact
   persistence-direction check this plan's Global Constraint exists for,
   the highest-value thing to verify live given sub-project #5's own
   history with this exact class of bug.

If a Windows build of `thetisctl.exe` isn't readily available this
session to fully exercise step 7 end-to-end, it's acceptable to defer that
one specific check as an explicitly-flagged gap (the checkbox wiring,
process-launch-attempt, and status-label logic can still be confirmed via
steps 6, 8's "unchecked = not running" path, and 9's persistence check,
which don't require the helper to actually succeed at connecting to
anything) — but do not skip steps 6, 8, and 9.

- [ ] **Step 13: Commit**

```bash
git add "Project Files/Source/Console/setup.designer.cs" \
        "Project Files/Source/Console/setup.cs" \
        "Project Files/Source/Console/console.cs" \
        "Project Files/Source/Console/FreeDVReporterHelper.cs" \
        "Project Files/Source/Console/Thetis.csproj"
git commit -m "feat(setup): rename Digital Voice tab to FreeDV, add reporter enable UI

New FreeDVReporterHelper static class (following Dumpcap.cs's exact
existing process-supervision pattern) launches/manages an external
thetisctl (or thetis-ai-skill) process for both --spot and --self-report,
combined into one process per the design spec. Persistence wiring
follows sub-project #5's own corrected pattern: InitRadePanelFromBackend
reads the checkboxes' restored state, never overwrites it. Process
teardown hooks into console.cs's Console_Closing (the main window's own
close handler, which already tears down Dumpcap the same way) rather
than Setup's own close handler, which only hides that singleton window
and never actually closes it.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
```

---
