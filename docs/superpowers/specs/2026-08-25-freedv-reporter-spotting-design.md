# FreeDV Reporter Panadapter Spotting — Design Spec

**Date**: 2026-08-25
**Status**: Approved, ready for implementation planning
**Sub-project**: 4 of 5 (see [2026-08-24-rade-setup-panel-design.md](2026-08-24-rade-setup-panel-design.md)'s
"Larger context" section)

## Goal

Show live FreeDV Reporter (qso.freedv.org) activity as markers on Thetis's
own panadapter — who's active, where, and whether they're currently
transmitting — without requiring a human to have the reporter's web page
open and watching it.

## What already exists (why this is small)

Two pieces of prior work in this codebase make this sub-project far smaller
than "build a FreeDV Reporter integration" sounds:

1. **A live, proven Socket.IO v4 client for the reporter's feed already
   exists**: `Tools/thetis-ai-control/internal/freedvreporter` (built and
   verified 2026-08-09, `ee88a402`; see `Documentation/FreeDV-Plan.md`'s
   "Stage D"). It's a hand-rolled client (the reporter's feed is Socket.IO,
   not REST — no maintained Go library was reused, matching this tool's
   `internal/tci` package's own precedent of a hand-rolled client where no
   good third-party option exists). `internal/freedvreporter/station.go`'s
   `Tracker` already parses `new_connection`/`freq_change`/`tx_report`/
   `remove_connection` events into a live `map[string]*Station` (keyed by
   the reporter's own session ID), each with `Callsign`, `GridSquare`,
   `FreqHz`, `Mode` (a free-text string like "RADEV1"), `Transmitting`,
   `RXOnly`. Currently this drives `thetisctl freedv-reporter watch`'s
   auto-tune-Thetis'-VFO-to-live-activity behavior — this sub-project adds
   a second, independent action on the same already-tracked data.
2. **Thetis already has a panadapter spot-overlay renderer with nothing left
   to build**: `Console/SpotManager2.cs`, already wired to accept spots
   pushed over TCI (`TCIServer.cs::handleSpot`, dispatched on wire command
   `spot`). Its `AddSpot(callsign, mode, frequencyHz, colour, additionalText,
   jsonSpotData=null)` method (verified by reading it directly, not assumed)
   already provides, for free, everything a live spot feed needs:
   - **Dedup + refresh**: re-adding the same callsign within 5kHz of its
     existing frequency (`SpotManager2.cs:598-599`) removes and replaces the
     old entry rather than duplicating it, resetting its age.
   - **Automatic expiry**: `SpotManager2.cs:251` — any non-SWL spot older
     than `_lifeTime` minutes is removed automatically on every render pass.
     No explicit "remove this station" wire command exists or is needed —
     a station that stops sending updates just ages out on its own.
   - **A max-spot cap** with oldest-first eviction (`SpotManager2.cs:626-656`),
     protecting against unbounded growth regardless of how many stations the
     reporter ever reports at once.
   - **Automatic country-flag lookup from the callsign prefix**
     (`SpotManager2.cs:512-524`, `getFlagImageFromCallsign`) even without
     the richer `[json]{...}` additional-text tag format — so the plain wire
     format is sufficient; no need to build or send the JSON variant.

The actual remaining gap, confirmed by reading both sides directly: nothing
in `internal/freedvreporter` currently sends a `spot:...;` TCI command for
any station update. That's the entire scope of this sub-project.

## Design

### Scope decision: extend the existing tool, not a native C# client

sv1eia's fork has a genuinely native FreeDV Reporter client
(`FreeDVReporter*.cs`, ~1,200 lines — request-QSY button, row selection,
error logging, its own Setup UI). Building an equivalent here would mean a
new C# Socket.IO v4 client from scratch (or awkward interop with the
existing Go one) plus real UI work — a large, mostly-duplicate undertaking.
Per this sub-project's own scoping decision: extend the already-live-tested
external tool instead. `internal/freedvreporter`'s Socket.IO client and
`Tracker` are reused entirely unchanged; only a new consumer of
`Tracker.Stations()`/`TxStarted` is added.

### Data mapping: `Station` → `spot:` command

**Trigger mechanism, precisely**: `Tracker.Apply()`'s own return value only
surfaces `TxStarted` transitions — it does *not* return a generic
"something changed" event, even though it updates its internal station map
on every `new_connection`/`freq_change`/`tx_report`. Rather than extend
`Tracker`'s API to expose more granular change events (new surface, more to
test), `--spot` calls `Tracker.Stations()` — the existing full-snapshot
method — after every `Apply()` call in the watch loop, and sends a
`spot:...;` command for every currently-tracked station with both a valid
`Callsign` and a nonzero `FreqHz`, not just the one that triggered the
event. This means unchanged stations get harmlessly re-sent on every loop
iteration too — `SpotManager2`'s dedup-by-callsign-and-frequency (see
above) treats that as a no-op refresh, not a duplicate, so this costs
nothing but a redundant TCI write. No changes to `Tracker`'s own API.

For each station, send:

```
spot:<callsign>,digu,<freqHz>,<argb_color>,<text>;
```

- **`<callsign>`**: `Station.Callsign`, as reported (SpotManager2 upper-cases
  and trims it itself).
- **Mode is always `digu`**: Thetis's `DSPMode` enum has no per-codec FreeDV
  variants; `digu` is the existing convention this project's own CAT/TCI
  work already uses for FreeDV-family spots (matches
  `thetisctl freedv-reporter watch`'s own `--mode digu` default). The
  reporter's own free-text `Mode` field (e.g. "RADEV1", "700D") goes in
  `<text>` instead, not the mode field — it's informational, not a value
  `DSPMode` can represent.
- **`<freqHz>`**: `Station.FreqHz` directly (TCI's spot command takes Hz;
  `Tracker` already stores Hz).
- **`<argb_color>`**: two fixed colors, chosen by `Station.Transmitting`. The
  wire value is an **unsigned** 32-bit decimal string — confirmed by
  reading `TCIServer.cs:4375` directly (`uint.TryParse(args[3], out argb)`),
  correcting an earlier "signed" paraphrase in this project's own prior
  notes (`Documentation/FreeDV-Plan.md`'s Stage D section) that didn't
  match the actual parse call:
  - Transmitting: a bright, attention-getting red, ARGB `(255,220,40,30)` =
    `4292618270` (`0xFFDC281E`).
  - Not transmitting (idle or RX-only monitor connection): a calmer blue,
    ARGB `(255,60,120,200)` = `4282153160` (`0xFF3C78C8`).
  `Station.RXOnly` gets no special treatment beyond this — an RX-only
  connection is real, relevant situational-awareness data (someone is
  listening on that frequency), it just can't ever go into the
  "transmitting" (red) state, so it always renders in the idle color.
- **`<text>`**: the reporter's own `Mode` string, plus `"TX"` appended when
  `Transmitting` is true (e.g. `"RADEV1 TX"` vs. just `"700D"`). No SNR —
  see "Explicitly out of scope" below.

The plain (non-JSON) wire format is used throughout — no `[json]{...}` tag,
since `AddSpot` already derives country/flag from the callsign prefix
without it (confirmed by reading `AddSpot`'s own fallback path).

### CLI shape

Two new flags on the existing `thetisctl freedv-reporter watch` subcommand,
both backward compatible (default off / default matches current behavior):

- **`--spot`** (bool, default `false`): when set, sends a `spot:...;` TCI
  command for every station update, per the mapping above. Independent of
  auto-tune — can be used alone or together with it.
- **`--no-tune`** (bool, default `false`): when set, disables the existing
  auto-tune-on-TX-start behavior. Exists so `--spot` can be used standalone
  (a pure situational-awareness panadapter view) without retuning the
  operator's VFO — today, `watch` always tunes; this is the first time that
  becomes optional.

`--min-freq`/`--max-freq` (already existing, default 20m band) are reused
as-is for both behaviors — spotting is filtered to the same band-edge
bounds as tuning, rather than introducing a second, separate set of bounds
flags. A user wanting a wider spotting view than tuning view is a plausible
future refinement, not built now (YAGNI — the existing flags already do
useful work here with zero new code).

### Testing

Same convention as the existing `internal/freedvreporter/live_test.go`
(`go:build live`, excluded from `go test ./...`/CI, run manually against the
real service): a real session with `--spot` enabled, confirming spot
commands are actually sent (either by capturing them at the TCI mock/test
level, or, per the existing live test's own precedent, by observing real
station data flow end-to-end) and that repeated updates for the same
station don't produce runaway duplicate spot commands. A manual live check
against a real running Thetis instance (send `--spot`'s output at
`hl2winbox`, confirm markers actually render on the panadapter and expire
correctly after activity stops) closes the loop the same way sub-projects
#1-#3's hardware verification did.

## Explicitly out of scope

- A native in-Thetis C# FreeDV Reporter client (sv1eia-style) — this
  sub-project's own scoping decision; may be revisited later if the
  external-tool approach proves insufficient.
- SNR/`rx_report` data in spot text — `Tracker` currently doesn't parse
  `rx_report` events at all (deliberately, per `station.go`'s own comment);
  adding it is new event-parsing work beyond "act on data already tracked,"
  a reasonable future enhancement but not required for spotting to be
  useful (seeing who's active where is the core value; signal-quality
  detail is a refinement).
- Request-QSY / any interactive click-to-respond feature (sv1eia has this;
  it requires two-way TCI control this tool doesn't currently exercise for
  spotting, and there's no operator-facing UI element to click here beyond
  the panadapter marker itself).
- A separate min/max-frequency bound for spotting distinct from tuning's
  existing bounds — reuses the existing flags, see "CLI shape" above.
- Any change to `SpotManager2.cs`/`TCIServer.cs` — both already do exactly
  what's needed, confirmed by reading them directly; this sub-project is
  purely a new consumer in `internal/freedvreporter`/`thetisctl`.
