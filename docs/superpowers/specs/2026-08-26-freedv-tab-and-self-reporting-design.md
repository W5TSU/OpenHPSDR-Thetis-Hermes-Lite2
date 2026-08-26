# FreeDV Tab: Reporter Enable UI + Self-Reporting — Design Spec

**Date**: 2026-08-26
**Status**: Approved, ready for implementation planning
**Sub-project**: 6 of 6 — the final piece of the Digital Voice effort (see
[2026-08-24-rade-setup-panel-design.md](2026-08-24-rade-setup-panel-design.md)'s
"Larger context" section for #1-#4; #5's own
["Larger context" update](2026-08-25-panadapter-sync-snr-overlay-design.md)
recorded that this sixth piece was added during its own brainstorming)

## Goal

Two additions to Thetis's Setup form, both centered on the Digital Voice tab
(renamed "FreeDV" as part of this work):

1. An in-Thetis UI to enable/manage sub-project #4's `thetisctl freedv-reporter
   watch --spot` (currently a manually-run external command) — a checkbox
   that launches and supervises the external process instead of requiring
   an operator to run it themselves in a terminal.
2. **Self-reporting**: publish *this station's own* presence (callsign, grid
   square, live frequency, mode, TX state) to qso.freedv.org, so other
   operators see it on the reporter's live map — the write-side of the
   protocol, the opposite direction from #4's read-only `--spot`, and a
   genuinely new capability not built anywhere in this codebase or in
   `thetisctl` today.

## Context that shapes this design

- **`thetisctl` (`Tools/thetis-ai-control`) is being extracted into its own
  project**, [thetis-ai-skill](https://github.com/W5TSU/thetis-ai-skill),
  outside this repo's control going forward. Thetis's enable-UI therefore
  does not bundle or build that binary itself — it launches whatever binary
  is configured (a path field) or found on `PATH`, the same way an operator
  would run it by hand today. The Go-side protocol work in this spec still
  lands in `Tools/thetis-ai-control` for now (this session can build and
  live-test it there) and simply migrates unchanged whenever that
  extraction happens — nothing in this spec's Go code depends on staying in
  this specific repo.
- **The self-reporting protocol is fully known, not guessed.** qso.freedv.org's
  own web viewer (`https://qso.freedv.org`) is read-only — direct inspection
  of its JS bundle showed it only ever emits `chat_login`/`chat_message`,
  never station data. The actual reporting client is `freedv-gui`'s own
  `FreeDVReporter` class (fetched from its `freedv_backend` dependency,
  read directly from a local checkout during this design's brainstorming:
  `src/reporting/FreeDVReporter.cpp`/`.h`). It connects to the same
  Socket.IO v4 server as the view-only feed, with a different `role` in the
  auth payload:
  ```
  { role: "report" | "report_wo", callsign, grid_square, version, rx_only, os, protocol_version: 2 }
  ```
  `"report_wo"` ("write-only") skips the server streaming the full station
  table back — the right choice here, since Thetis's self-report connection
  has no need to view other stations (that's `--spot`'s separate, existing
  view-role connection). After connecting, the client emits update events
  as its own state changes: `freq_change` → `{freq: <Hz>}`, `tx_report` →
  `{mode, transmitting}`. (`hide_self`/`show_self`/`message_update`/
  `rx_report` also exist in the real protocol but are out of scope here —
  see "Explicitly out of scope.")
- **This codebase already has an operator callsign setting** —
  `SetupForm.TCIOwnCallsign` (`setup.cs:12430-12433`, backed by
  `txtOwnCallsign.Text` in the existing TCI Server group) — reused directly
  for the report payload's `callsign` field; no new callsign UI needed.
  There is no existing grid-square setting anywhere in this codebase — one
  new text field is needed for that.
- **Thetis's TCI server already broadcasts unsolicited state-change frames**
  (`vfo:`/`modulation:`/`trx:`) that existing code in `internal/tci` already
  parses the wire shape of (via `RecvCmd`, currently exercised only in
  `live_test.go` to await specific replies) but no production `thetisctl`
  code currently *listens* for them proactively — every existing call site
  in `cmd/thetisctl` only *writes* (`SendCmd`-based setters). This sub-project's
  self-report bridge is the first consumer that needs to *read* Thetis's own
  live state continuously.

## Design

### Tab rename

`tpDSPRADE.Text` changes from `"Digital Voice"` to `"FreeDV"`
(`setup.designer.cs:43931`) — a one-line change reflecting what the tab
has actually become across sub-projects #1-#6: RADE/700E controls plus
reporter integration, not just a mode selector.

### UI: new "FreeDV Reporter" group in the tab

A new group box in the renamed tab (exact pixel layout — including whether
the tab needs to grow or existing controls need to reflow to fit it —
is a planning-time decision; the tab is already fairly full, see
sub-project #5's layout map in its own plan) containing:

- **Helper path** (text field): path to the external helper binary; empty
  means "resolve via `PATH`." No validation beyond "does this file exist
  when we try to launch it" — a clear status-label error covers the
  not-found case, no need for eager validation as the operator types.
- **"Show other stations on panadapter"** (checkbox): wraps sub-project #4's
  `--spot` (and its `--min-freq`/`--max-freq`/`--no-tune` behavior, which
  keep their existing defaults — no new UI for those, YAGNI until an
  operator actually asks).
- **"Report my station to FreeDV Reporter"** (checkbox, new): enables
  self-reporting.
- **Grid square** (text field, new): free text, sent as-is as the report's
  `grid_square` field — no client-side Maidenhead-locator format validation
  in this pass (a malformed value is the operator's own data-entry mistake
  to notice on the reporter's map, same as any other reporting tool).
- **"I don't transmit"** (checkbox, new): maps to the report payload's
  `rx_only` field.
- **Status label**: shows not-running / running / last error from the
  helper process's stderr, refreshed same as sub-project #5's own status
  labels (a `System.Windows.Forms.Timer` polling process state — this
  process is short-lived output, not a 500ms DSP-status poll, so a slower
  interval, e.g. 2s, is appropriate).

Callsign is not a new field — it reads `SetupForm.TCIOwnCallsign` directly
at launch time.

### Process management

Thetis manages **one** external helper process for both checkboxes
combined, not two separate processes — both `--spot` and `--self-report`
point at the same underlying TCI connection inside the helper (see
"Go-side protocol implementation" below), so one process serves both
features whenever either or both are enabled:

- Neither checkbox on: no process running.
- Either or both checked: process running with the corresponding flags
  (`--spot`, `--self-report --callsign <TCIOwnCallsign> --grid <gridSquare>
  [--rx-only]`) plus the existing required `--tci <thetis-own-address>`
  (localhost, since the helper runs alongside Thetis — exact loopback
  address a planning-time detail, matching how sub-project #5's own
  live-testing used the box's real LAN address rather than `localhost`
  when testing cross-machine; when the helper runs on the *same* machine
  as Thetis, `127.0.0.1` is correct and simpler).
- Toggling either checkbox, or editing grid square / rx-only while
  running, restarts the process with the new flag set — simplest correct
  behavior, no in-process reconfiguration protocol needed for a helper
  this cheap to restart.
- The process is launched via `System.Diagnostics.Process`, stdout/stderr
  redirected into the status label (last line / last error), and killed on
  Setup close or Thetis exit (a `Process.Kill()` in the existing
  Setup-teardown path, not a new lifecycle hook).

### Go-side protocol implementation

**`internal/freedvreporter`** gains a report-role connection type,
following `Dial`'s existing handshake exactly (Engine.IO OPEN → Socket.IO
CONNECT+auth → CONNECT-ack, verified byte-for-byte against the existing,
already-live-tested `Dial` function) with a different auth payload:

```go
func DialReport(callsign, gridSquare string, rxOnly, writeOnly bool, timeout time.Duration) (*Client, error)
```

sending `{role: writeOnly ? "report_wo" : "report", callsign, grid_square:
gridSquare, version: "Thetis (thetisctl)", rx_only: rxOnly, os: runtime.GOOS,
protocol_version: 2}` — field names and the `report`/`report_wo` split
confirmed directly against `freedv-gui`'s own `FreeDVReporter.cpp`, not
inferred from the (read-only) web client.

A new method mirrors `ReadEvent`'s parsing in reverse:

```go
func (c *Client) Emit(event string, payload any) error
```

writing `"42" + json.Marshal([]any{event, payload})` — the exact Socket.IO
v4 EVENT wire form `ReadEvent` already parses on the receive side
(`client.go`'s existing `"42"`-prefix check), so both directions of this
package now share one understanding of the wire format.

**`cmd/thetisctl freedv-reporter watch`** gains:
- `--self-report` (bool, default false)
- `--callsign` (string) — required if `--self-report` is set
- `--grid` (string) — sent as-is, may be empty
- `--rx-only` (bool, default false)
- `--write-only` (bool, default true — matches the "this connection doesn't
  need to view the map" reasoning above; exposed as a flag rather than
  hardcoded in case a future caller wants the map data too)

`--self-report` requires `--tci`, exactly like `--spot` already does (this
plan's own fail-fast-before-dialing-anything pattern, established in
sub-project #4). When both `--spot` and `--self-report` are given, they
share the single `--tci` connection already established for spotting —
no second TCI dial.

When `--self-report` is set, on top of everything `watch` already does,
the command:
1. Dials the report-role connection via `DialReport`.
2. **Queries Thetis's current VFO A frequency, mode, and MOX state
   immediately** (rather than only waiting passively for the next change)
   so the very first report reflects reality, not a stale/empty value —
   the exact TCI query mechanism (a `get`-form wire command, per this
   package's existing `live_test.go` precedent of sending a command with
   fewer arguments than its "set" form to elicit the current value) is a
   plan-level detail to pin down against the real wire behavior, not
   guessed here.
3. Emits an initial `freq_change`/`tx_report` from that query's result.
4. Spawns a goroutine that loops on the shared `tciClient.RecvCmd()` —
   currently unused for reading anywhere in `cmd/thetisctl` — watching for
   subsequent unsolicited `vfo`/`modulation`/`trx` frames and translating
   each into the matching `Emit` call. Reading and writing happen on
   separate goroutines against the same connection, which is safe in Go as
   long as each direction has exactly one goroutine (true here: the
   existing code's writes — auto-tune's `SetVFOFreqHz`/`SetModulation`,
   sub-project #4's `Spot` — all originate from the existing single-goroutine
   main loop; only reading moves to its own goroutine).

### Testing

- **Unit test** for `Emit`'s wire format, mirroring sub-project #4's
  `TestSpotWireFormat` `net.Pipe` pattern exactly — assert the exact bytes
  `Emit("freq_change", ...)` produces.
- **Live verification**: run `--self-report` against the real
  `qso.freedv.org` with a real callsign, and confirm the station actually
  appears on the reporter's own web page (`https://qso.freedv.org`) with
  the right callsign/grid/frequency — the most convincing possible proof,
  since that page is exactly what other operators would see. Also confirm
  changing VFO frequency/mode/TX state in Thetis updates the reporter's
  map within a few seconds, and that stopping the helper process removes
  the station from the map (the reporter's own server-side disconnect
  handling, not new code here).
- Live verification of the C# side (checkbox → process launch → status
  label) follows this project's established hardware-in-the-loop practice
  from sub-projects #1-#5, deploying to `hl2winbox` and confirming via the
  interactive-session-launch technique established during sub-project #5
  (documented there; no remote UI automation needed for this feature's own
  own checks beyond confirming the process starts/stops correctly).

## Explicitly out of scope

- `hide_self`/`show_self` (temporarily hiding from the map without
  disconnecting) and `message_update` (a free-text status message shown
  next to your station) — real features of the protocol, no clear Thetis
  UI hook for them yet; natural follow-ons if wanted later.
- `rx_report` (reporting SNR of *other* stations you're hearing) — a
  separate capability from self-presence reporting, would need new
  RX-quality data Thetis doesn't currently expose to `thetisctl` at all.
- `qsy_request` (asking another station to change frequency) — an
  interactive feature with no operator-facing UI trigger in this design.
- Any change to `internal/freedvreporter`'s existing view-role `Dial`/
  `ReadEvent`/`Tracker` — the report-role connection is fully additive,
  sub-project #4's spotting code is untouched.
- Grid-square format validation, helper-path existence checking before
  launch, or any other input-hardening beyond a status label showing
  whatever error actually occurs — YAGNI until an operator hits a real
  confusing failure mode.
- Persisting the helper's running/stopped intent across a Thetis restart
  beyond what the two checkboxes' own Setup-persisted `.Checked` state
  already provides (sub-project #5's final review found and fixed exactly
  this class of bug for its own toggle — the same
  restore-then-launch-on-startup logic applies here and must get the same
  scrutiny in this sub-project's own review).
