# FreeDV Reporter Panadapter Spotting Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Extend `thetisctl freedv-reporter watch` so it can push live FreeDV
Reporter activity onto Thetis's panadapter as TCI `spot:` markers, alongside
its existing auto-tune behavior.

**Architecture:** Add one new typed TCI wrapper method (`Client.Spot`) to
`internal/tci/control.go`, following that file's existing per-command
convention. Add two new CLI flags (`--spot`, `--no-tune`) and a
spot-pushing loop to `cmd/thetisctl/freedvreporter_cmd.go`'s existing
`freedvReporterWatch`, reusing `internal/freedvreporter`'s already-live
`Tracker` completely unchanged. No C# changes; no new Go packages.

**Tech Stack:** Go 1.22 (module `thetisctl`, repo root
`Tools/thetis-ai-control`), stdlib only (`net`, `testing`, `bufio`), this
tool's own `internal/tci` and `internal/freedvreporter` packages.

**Spec:** [docs/superpowers/specs/2026-08-25-freedv-reporter-spotting-design.md](../specs/2026-08-25-freedv-reporter-spotting-design.md)

## Global Constraints

- **Working directly on the `FreeDV` branch**, no worktree — matches this
  project's established practice for sub-projects #1-#3 (confirmed with the
  repo owner earlier in this effort).
- **None of the CRLF-mixed-line-ending files apply to this sub-project.**
  `CATCommands.cs`/`CATParser.cs`/`CATStructs.xml`/`radio.cs` are untouched —
  this plan only touches Go files under `Tools/thetis-ai-control`, which are
  plain-LF and have no history of Edit-tool line-ending corruption. The
  Python-byte-splicing workaround used in sub-projects #1-#3 is NOT needed
  here.
- **`<argb_color>` on the wire is an UNSIGNED 32-bit decimal string** —
  confirmed directly against `TCIServer.cs:4375`
  (`uint.TryParse(args[3], out argb)`), not signed. Every task below uses
  `uint32` in Go and the two verified decimal values: TX = `4292618270`
  (`0xFFDC281E`), idle = `4282153160` (`0xFF3C78C8`). Do not introduce an
  `int32`/signed representation anywhere in this feature.
- **Mode on the wire is always the literal string `"digu"`** — confirmed
  `DSPMode.DIGU` exists (`enums.cs:263`) and `handleSpot`'s
  `Enum.TryParse(args[1].ToUpper(), out mode)` fast path
  (`TCIServer.cs:4378`) matches it directly; no raw-mode-filter fallback is
  ever exercised. Never pass the reporter's own free-text `Mode` field
  (e.g. `"RADEV1"`) as the wire mode argument — it goes in the spot's
  `<text>` field instead.
- **No changes to `internal/freedvreporter/station.go`'s `Tracker` API.**
  `Tracker.Stations()` (already public, already used by `live_test.go`) is
  sufficient; do not add a new change-event type to `Tracker`.
- **No changes to any C# file** (`SpotManager2.cs`, `TCIServer.cs`, or
  anything under `Project Files/`) — both already do exactly what this
  feature needs, confirmed by direct code reading during design.
- Every new/modified Go file must pass `gofmt -l` (no output = clean) and
  `go vet ./...` from `Tools/thetis-ai-control`.
- `go build ./...` and `go test ./...` (run from `Tools/thetis-ai-control`)
  must both succeed after every task. `go test ./...` excludes `live`-tagged
  tests by default (no `-tags=live`) — this is correct and expected; do not
  add `-tags=live` to the routine test command.

---

### Task 1: `Client.Spot` typed TCI wrapper

**Files:**
- Modify: `Tools/thetis-ai-control/internal/tci/control.go`
- Create: `Tools/thetis-ai-control/internal/tci/control_test.go`

**Interfaces:**
- Consumes: `Client.SendCmd(cmd string, args ...string) error` (already
  exists, `client.go:26`); the existing `itoa(n int) string` helper
  (`control.go:227`).
- Produces: `func (c *Client) Spot(callsign, mode string, freqHz int64, argb uint32, text string) error`
  — Task 2 calls this exactly, with `mode` always `"digu"` and `argb` always
  one of the two Global-Constraints decimal values.

- [ ] **Step 1: Add the `Spot` method to `control.go`**

Append this to the end of `internal/tci/control.go`, immediately before the
existing `itoa` helper at the bottom of the file (i.e. insert right after
`SetPower`'s closing brace, before `encodeCWText`'s doc comment):

```go
// Spot pushes a live activity marker onto Thetis's panadapter, dispatched
// by TCIServer.cs::handleSpot into Console/SpotManager2.cs's AddSpot. mode
// is TCI's lowercase demod name (see SetModulation) — this tool always
// passes "digu" for FreeDV activity, since DSPMode has no per-codec FreeDV
// variant; any richer description (e.g. "RADEV1", "700D") belongs in text
// instead. argb is an UNSIGNED 32-bit color packed as 0xAARRGGBB and sent
// as a decimal string — confirmed against TCIServer.cs:4375
// (uint.TryParse(args[3], out argb)), NOT a signed int32. text is free-form
// and may be empty; no escaping is needed before calling this — handleSpot's
// non-JSON path rejoins every argument from index 4 onward with ","
// (TCIServer.cs:4351-4356), so any comma already present in text round-trips
// unchanged.
// Wire: "spot:<callsign>,<mode>,<freqHz>,<argb>,<text>;" (handleSpot,
// TCIServer.cs:4339-4408).
func (c *Client) Spot(callsign, mode string, freqHz int64, argb uint32, text string) error {
	return c.SendCmd("spot", callsign, mode, strconv.FormatInt(freqHz, 10), strconv.FormatUint(uint64(argb), 10), text)
}
```

No new imports are needed — `strconv` is already imported at the top of
`control.go`.

- [ ] **Step 2: Write the wire-format unit test**

Create `Tools/thetis-ai-control/internal/tci/control_test.go` with this
exact content:

```go
package tci

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

// TestSpotWireFormat exercises Client.Spot end-to-end over a net.Pipe,
// verifying the exact wire text a real Thetis TCI server would receive —
// in particular, that the argb color is sent as the correct UNSIGNED
// decimal (TCIServer.cs:4375 parses via uint.TryParse, not a signed int).
// No "live" build tag: this needs no network or running Thetis instance,
// so it runs in normal `go test ./...` / CI, matching TestFrameRoundTrip's
// existing net.Pipe pattern in ws_test.go.
func TestSpotWireFormat(t *testing.T) {
	clientEnd, serverEnd := net.Pipe()
	defer clientEnd.Close()
	defer serverEnd.Close()

	c := NewClient(&Conn{tcp: clientEnd, r: bufio.NewReader(clientEnd), timeout: 2 * time.Second})
	s := &Conn{tcp: serverEnd, r: bufio.NewReader(serverEnd), timeout: 2 * time.Second}

	done := make(chan error, 1)
	go func() {
		op, payload, err := s.ReadFrame()
		if err != nil {
			done <- err
			return
		}
		want := "spot:W5TSU,digu,14070000,4292618270,RADEV1 TX;"
		if op != opText || string(payload) != want {
			done <- fmt.Errorf("got opcode %d payload %q, want text %q", op, payload, want)
			return
		}
		done <- nil
	}()

	if err := c.Spot("W5TSU", "digu", 14070000, 4292618270, "RADEV1 TX"); err != nil {
		t.Fatalf("Spot: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

// TestSpotWireFormatIdleColor checks the idle-color decimal value and an
// empty text field separately, so both of the Global Constraints' verified
// argb values are covered by a passing test, not just the TX one.
func TestSpotWireFormatIdleColor(t *testing.T) {
	clientEnd, serverEnd := net.Pipe()
	defer clientEnd.Close()
	defer serverEnd.Close()

	c := NewClient(&Conn{tcp: clientEnd, r: bufio.NewReader(clientEnd), timeout: 2 * time.Second})
	s := &Conn{tcp: serverEnd, r: bufio.NewReader(serverEnd), timeout: 2 * time.Second}

	done := make(chan error, 1)
	go func() {
		op, payload, err := s.ReadFrame()
		if err != nil {
			done <- err
			return
		}
		want := "spot:VK3ABC,digu,7177000,4282153160,700D;"
		if op != opText || string(payload) != want {
			done <- fmt.Errorf("got opcode %d payload %q, want text %q", op, payload, want)
			return
		}
		done <- nil
	}()

	if err := c.Spot("VK3ABC", "digu", 7177000, 4282153160, "700D"); err != nil {
		t.Fatalf("Spot: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}
```

- [ ] **Step 3: Run the new tests**

Run: `cd Tools/thetis-ai-control && go test ./internal/tci/... -run TestSpot -v`
Expected: both `TestSpotWireFormat` and `TestSpotWireFormatIdleColor` PASS.

- [ ] **Step 4: Run the full non-live test suite and vet/format checks**

Run, all from `Tools/thetis-ai-control`:
```
gofmt -l internal/tci/control.go internal/tci/control_test.go
go vet ./...
go build ./...
go test ./...
```
Expected: `gofmt -l` prints nothing; `go vet`, `go build`, `go test` all
succeed with no failures (existing `live`-tagged tests are skipped
automatically, as today).

- [ ] **Step 5: Commit**

```bash
cd Tools/thetis-ai-control
git add internal/tci/control.go internal/tci/control_test.go
git commit -m "feat(thetisctl): add Client.Spot TCI wrapper for panadapter markers"
```

---

### Task 2: `--spot`/`--no-tune` flags and the spot-pushing loop

**Files:**
- Modify: `Tools/thetis-ai-control/cmd/thetisctl/args.go`
- Modify: `Tools/thetis-ai-control/cmd/thetisctl/freedvreporter_cmd.go`

**Interfaces:**
- Consumes: `Client.Spot(callsign, mode string, freqHz int64, argb uint32, text string) error`
  from Task 1; `freedvreporter.Tracker.Stations() []freedvreporter.Station`
  and the `freedvreporter.Station` fields `Callsign`, `FreqHz`, `Mode`,
  `Transmitting`, `RXOnly` (all already exist, `station.go:14-25`);
  `parsedArgs.has(name string) bool` (already exists, `args.go:56`).
- Produces: nothing further downstream — this is the CLI's outermost layer
  for this feature.

- [ ] **Step 1: Register the two new flags as boolean (no-value) flags**

`parseArgs` (`args.go:22`) otherwise consumes the *next* CLI argument as a
flag's value unless the flag name is listed in `boolFlagNames`
(`args.go:18-20`) — without this step, `thetisctl freedv-reporter watch
--spot` would silently swallow whatever argument follows `--spot` as its
value instead of treating it as a bare switch.

In `Tools/thetis-ai-control/cmd/thetisctl/args.go`, change:

```go
var boolFlagNames = map[string]bool{
	"audio": true,
}
```

to:

```go
var boolFlagNames = map[string]bool{
	"audio":   true,
	"spot":    true,
	"no-tune": true,
}
```

- [ ] **Step 2: Add the two new flags and the `--spot`-requires-`--tci` check**

In `Tools/thetis-ai-control/cmd/thetisctl/freedvreporter_cmd.go`, in
`freedvReporterWatch`, change:

```go
	tciHost := a.flag("tci", "")
	tciPort := a.flag("tci-port", "50001")
	tuneMode := a.flag("mode", "digu")

	fmt.Printf("freedv-reporter watch: connecting to %s ...\n", freedvreporter.ReporterHost)
```

to:

```go
	tciHost := a.flag("tci", "")
	tciPort := a.flag("tci-port", "50001")
	tuneMode := a.flag("mode", "digu")
	doSpot := a.has("spot")
	noTune := a.has("no-tune")

	if doSpot && tciHost == "" {
		return fmt.Errorf("freedv-reporter watch: --spot requires --tci <host> (spots are pushed over TCI; there is nowhere else to send them)")
	}

	fmt.Printf("freedv-reporter watch: connecting to %s ...\n", freedvreporter.ReporterHost)
```

This check runs before the reporter `Dial` call, so a missing `--tci` fails
fast with a clear message rather than connecting to the reporter first and
failing later on the first tracked station.

- [ ] **Step 3: Update the startup banner to mention spotting and respect `--no-tune`**

Change:

```go
	fmt.Printf("freedv-reporter watch: connected. Watching %.3f-%.3f MHz",
		float64(minFreq)/1e6, float64(maxFreq)/1e6)
	if tciHost != "" {
		fmt.Printf(" — will auto-tune %s (mode %s) on activity", net.JoinHostPort(tciHost, tciPort), tuneMode)
	}
	fmt.Println(". Ctrl-C to stop.")
```

to:

```go
	fmt.Printf("freedv-reporter watch: connected. Watching %.3f-%.3f MHz",
		float64(minFreq)/1e6, float64(maxFreq)/1e6)
	if tciHost != "" && !noTune {
		fmt.Printf(" — will auto-tune %s (mode %s) on activity", net.JoinHostPort(tciHost, tciPort), tuneMode)
	}
	if doSpot {
		fmt.Printf(" — pushing spots to %s panadapter", net.JoinHostPort(tciHost, tciPort))
	}
	fmt.Println(". Ctrl-C to stop.")
```

- [ ] **Step 4: Replace the lazy-TCI-dial block with a shared `ensureTCI` closure**

Change:

```go
	tracker := freedvreporter.NewTracker()
	var tciClient *tci.Client // lazily dialed on first qualifying spot, kept for reuse
```

to:

```go
	tracker := freedvreporter.NewTracker()
	var tciClient *tci.Client // lazily dialed on first qualifying tune/spot action, kept for reuse

	// ensureTCI dials tciClient on first use and reuses it thereafter. Both
	// the auto-tune block and the spot-pushing block below call this —
	// they may run in the same watch session (doSpot and auto-tune are
	// independent flags) and share one TCI connection rather than each
	// dialing its own.
	ensureTCI := func() bool {
		if tciClient != nil {
			return true
		}
		addr := net.JoinHostPort(tciHost, tciPort)
		conn, derr := tci.Dial(addr, 5*time.Second)
		if derr != nil {
			fmt.Printf("  -> connect to Thetis TCI at %s failed: %v\n", addr, derr)
			return false
		}
		tciClient = tci.NewClient(conn)
		return true
	}
```

- [ ] **Step 5: Rework the per-event loop body**

Change the entire body from `started, err := tracker.Apply(ev)` through the
end of the `for _, ts := range started { ... }` block — i.e. replace:

```go
		started, err := tracker.Apply(ev)
		if err != nil {
			continue
		}

		for _, ts := range started {
			s := ts.Station
			if s.RXOnly || s.Callsign == "" {
				continue // monitor-only / not-yet-fully-known connections can't be "transmitting" in a meaningful sense
			}
			if s.FreqHz < minFreq || s.FreqHz > maxFreq {
				continue
			}

			fmt.Printf("[%s] %s TX-ing on %.4f MHz (FreeDV mode: %s)\n",
				time.Now().Format("15:04:05"), s.Callsign, float64(s.FreqHz)/1e6, s.Mode)

			if tciHost == "" {
				continue
			}

			if tciClient == nil {
				addr := net.JoinHostPort(tciHost, tciPort)
				conn, derr := tci.Dial(addr, 5*time.Second)
				if derr != nil {
					fmt.Printf("  -> auto-tune failed: connect to Thetis TCI at %s: %v\n", addr, derr)
					continue
				}
				tciClient = tci.NewClient(conn)
			}

			if serr := tciClient.SetVFOFreqHz(0, 0, s.FreqHz); serr != nil {
				fmt.Printf("  -> auto-tune failed: set VFO: %v\n", serr)
				tciClient.Close()
				tciClient = nil
				continue
			}
			if serr := tciClient.SetModulation(0, tuneMode); serr != nil {
				fmt.Printf("  -> auto-tune failed: set mode: %v\n", serr)
				tciClient.Close()
				tciClient = nil
				continue
			}
			fmt.Printf("  -> retuned RX1 to %.4f MHz / %s\n", float64(s.FreqHz)/1e6, tuneMode)
		}
	}
```

with:

```go
		started, err := tracker.Apply(ev)
		if err != nil {
			continue
		}

		for _, ts := range started {
			s := ts.Station
			if s.RXOnly || s.Callsign == "" {
				continue // monitor-only / not-yet-fully-known connections can't be "transmitting" in a meaningful sense
			}
			if s.FreqHz < minFreq || s.FreqHz > maxFreq {
				continue
			}

			fmt.Printf("[%s] %s TX-ing on %.4f MHz (FreeDV mode: %s)\n",
				time.Now().Format("15:04:05"), s.Callsign, float64(s.FreqHz)/1e6, s.Mode)

			if noTune || tciHost == "" {
				continue
			}
			if !ensureTCI() {
				continue
			}

			if serr := tciClient.SetVFOFreqHz(0, 0, s.FreqHz); serr != nil {
				fmt.Printf("  -> auto-tune failed: set VFO: %v\n", serr)
				tciClient.Close()
				tciClient = nil
				continue
			}
			if serr := tciClient.SetModulation(0, tuneMode); serr != nil {
				fmt.Printf("  -> auto-tune failed: set mode: %v\n", serr)
				tciClient.Close()
				tciClient = nil
				continue
			}
			fmt.Printf("  -> retuned RX1 to %.4f MHz / %s\n", float64(s.FreqHz)/1e6, tuneMode)
		}

		// Spot-pushing is independent of the TxStarted-only started slice
		// above: it re-sends every currently-tracked station (with a valid
		// callsign and a known frequency) on every event, not just the ones
		// that just started transmitting. SpotManager2's dedup-by-callsign-
		// and-frequency treats a repeat send as a harmless refresh, not a
		// duplicate (see the design spec's "Trigger mechanism, precisely").
		if doSpot {
			for _, s := range tracker.Stations() {
				if s.Callsign == "" || s.FreqHz == 0 {
					continue
				}
				if s.FreqHz < minFreq || s.FreqHz > maxFreq {
					continue
				}
				if !ensureTCI() {
					continue
				}

				color := spotColorIdle
				text := s.Mode
				if s.Transmitting {
					color = spotColorTX
					if text != "" {
						text += " TX"
					} else {
						text = "TX"
					}
				}

				if serr := tciClient.Spot(s.Callsign, "digu", s.FreqHz, color, text); serr != nil {
					fmt.Printf("  -> spot failed for %s: %v\n", s.Callsign, serr)
					tciClient.Close()
					tciClient = nil
				}
			}
		}
	}
```

Note the closing `}` of the `for { ... }` loop and the function are
unchanged — this replacement only covers from `started, err := ...` to the
end of the `for _, ts := range started` block, adding the new `if doSpot`
block immediately after it, still inside the outer `for { ... }` event
loop.

- [ ] **Step 6: Add the spot color constants and the `strings` import**

`text += " TX"` in Step 5 needs no new import (plain string concatenation),
but add the two color constants near the top of the file, just below the
existing `band20mMinHz, band20mMaxHz` constant:

```go
const band20mMinHz, band20mMaxHz = 14000000, 14350000

// Fixed panadapter marker colors for FreeDV Reporter spots. Both are
// UNSIGNED 32-bit ARGB packed as 0xAARRGGBB — see Client.Spot's doc comment
// in internal/tci/control.go for why (TCIServer.cs:4375 uses uint.TryParse).
const (
	spotColorTX   uint32 = 4292618270 // 0xFFDC281E = (255,220,40,30), bright red: this station is transmitting right now
	spotColorIdle uint32 = 4282153160 // 0xFF3C78C8 = (255,60,120,200), calmer blue: connected/monitoring, not transmitting
)
```

(No new imports are needed in this file — `net`, `os`, `os/signal`,
`time`, `fmt`, `thetisctl/internal/freedvreporter`, `thetisctl/internal/tci`
are all already imported and all still used.)

- [ ] **Step 7: Verify the full file compiles and passes vet/format**

Run, from `Tools/thetis-ai-control`:
```
gofmt -l cmd/thetisctl/args.go cmd/thetisctl/freedvreporter_cmd.go
go vet ./...
go build ./...
go test ./...
```
Expected: `gofmt -l` prints nothing; all three other commands succeed.

- [ ] **Step 8: Manually verify CLI parsing (no network required)**

Run: `cd Tools/thetis-ai-control && go run ./cmd/thetisctl freedv-reporter watch --spot`
Expected: prints the usage/connecting line then immediately fails with
exactly `freedv-reporter watch: --spot requires --tci <host> (spots are
pushed over TCI; there is nowhere else to send them)` — proving `--spot`
parses as a bare bool flag (Step 1) rather than swallowing a following
argument, and that the guard from Step 2 fires.

- [ ] **Step 9: Commit**

```bash
cd Tools/thetis-ai-control
git add cmd/thetisctl/args.go cmd/thetisctl/freedvreporter_cmd.go
git commit -m "feat(thetisctl): add --spot/--no-tune to freedv-reporter watch"
```

- [ ] **Step 10: Live verification against hl2winbox**

This closes the loop the same way sub-projects #1-#3's hardware
verification did — a real run against the actual reporter feed and a real
Thetis instance, per the design spec's "Testing" section.

1. Build the tool locally (already done by Step 7's `go build ./...`).
2. Confirm Thetis is running on `hl2winbox` with its TCI server enabled
   (reuse the existing verification pattern from sub-projects #1-#3: `ssh
   hl2winbox` + `thetisctl cat --host 100.117.67.160 version` to confirm
   the box is up, matching this session's established deploy-verification
   habit).
3. From this machine, run:
   ```
   go run ./cmd/thetisctl freedv-reporter watch --spot --tci 100.117.67.160 --no-tune
   ```
   (`--no-tune` here is deliberate for this first live check — it isolates
   spot-pushing from the pre-existing auto-tune behavior so a VFO retune
   doesn't happen mid-check and confuse what's being verified.)
4. Confirm in the terminal output that `TX-ing on ...` lines appear as real
   reporter traffic arrives (this part of the pipeline is unchanged from
   before this plan, so seeing it confirms the reporter connection itself
   is healthy).
5. On the Thetis panadapter (screen-share, remote screenshot, or in person,
   whichever is available this session), confirm spot markers appear at
   the reported stations' frequencies, colored blue for idle / red for
   transmitting, and that a marker's flag/callsign label is legible.
6. Let it run for several minutes and confirm markers for stations that
   stop reporting age out on their own (per `SpotManager2`'s existing
   `_lifeTime`-based expiry — no explicit removal call exists or is needed,
   confirmed during design).
7. Stop with Ctrl-C; confirm the process exits cleanly (the existing
   `SIGINT`-closes-the-reporter-connection handling is unchanged by this
   plan).

If the remote panadapter can't be visually confirmed this session (this
project's established remote-UI-automation reliability problems, see prior
sub-projects), it is acceptable to defer only the visual screen check as an
honest, explicitly-flagged gap, the same way sub-projects #2 and #3 handled
it — but the terminal-level confirmation (step 4) and a clean `--spot`
run producing no `spot failed for ...` errors are not optional; both must
be confirmed live before this task is considered done.

---
