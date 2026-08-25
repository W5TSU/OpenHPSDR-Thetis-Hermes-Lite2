# Digital Voice Mode Selector — Design Spec

**Date**: 2026-08-25
**Status**: Approved, ready for implementation planning
**Sub-project**: 2 of 5 (see [2026-08-24-rade-setup-panel-design.md](2026-08-24-rade-setup-panel-design.md)'s
"Larger context" section)

## Goal

Fold FreeDV 700E into the RADE Setup panel (built in sub-project #1) as a peer
mode alongside RADE V1/V2, replacing three independent, uninterlocked enable
paths (700E RX, 700E TX, RADE RX+TX+protocol) with one exclusive Mode
selector. No sv1eia precedent exists for this — confirmed during #1's
research that 700E doesn't appear anywhere in their design at all — so the
interaction model below is this fork's own design, decided through the
questions in this doc's revision history, not copied from an upstream.

## Problem statement

700E (`wdsp/fdv.c`, hardcoded to `FREEDV_MODE_700E`, no other rate
configurable) and RADE V1/V2 (`ChannelMaster/radae.c`, a neural codec) are
structurally unrelated subsystems that happen to serve the same purpose:
decoding/encoding digital voice on RX1. Today they're controlled by entirely
separate, uninterlocked paths:

| | RX enable | TX enable | Backend layer |
|---|---|---|---|
| 700E | `RXAFDVRun` (radio.cs `DSPRX` property → `WDSP.SetRXAFDVRun`) | `TXAFDVRun` (radio.cs `DSPTX` property → `WDSP.SetTXAFDVRun`) | wdsp RXA/TXA channel |
| RADE | `RXRadaeEnabled` (radio.cs `DSPRX` property → `WDSP.SetRadaeRxEnabled`) | `WDSP.SetRadaeTxEnabled` (no radio.cs property, direct call) | ChannelMaster tap by rx index |

Nothing stops both being armed simultaneously — confirmed by reading both
setters, neither references the other. That's meaningless for a single RX1
antenna feed and is the actual problem this sub-project closes.

## Design

### RX1 Core group — the unified Mode selector

The existing "RADE Protocol: V1/V2" combo (`cmbRadeProtocol`, sub-project #1)
is repurposed into a 4-item selector — **Off / 700E / RADE V1 / RADE V2** —
and moves into the RX1 Core group box (the standalone "Protocol" group box is
removed; RX1 Core becomes the group that owns mode selection, matching that
it already owns the loopback and level controls). Control name is unchanged
(`cmbRadeMode` after rename — see "Renamed/removed controls" below); tab
label changes from "RADE" to **"Digital Voice"**.

Selecting a mode arms that subsystem's RX1 decode **and** TX encode together
(still no MOX/PTT arbiter — arming TX encode remains inert, same as today,
just now reachable through one control instead of needing raw CAT). Selecting
Off disarms everything. The existing separate "RX1 RADE Enable" checkbox is
removed — enable is now implicit in the mode selection.

"RX1 RADE Loopback Test" (`chkRadeRX1Loopback`) is kept as one checkbox, but
its handler becomes mode-aware: calls `WDSP.SetRadaeLoopbackEnabled` when the
current mode is RADE V1/V2, `WDSP.SetFDVLoopbackEnabled` when it's 700E,
and is itself disabled (grayed) when mode is Off (nothing to loop back).

On tab load/reopen, `InitRadePanelFromBackend` (from #1) is extended to derive
`cmbRadeMode`'s initial `SelectedIndex` by reading the interlocked state back:
`RXRadaeEnabled == 1` → index 2 or 3 depending on `GetRadaeProtocolV2`;
else `RXAFDVRun == 1` → index 1; else index 0 (Off). Since the low-level
interlock guarantees at most one of `RXAFDVRun`/`RXRadaeEnabled` is ever 1,
this ordering is unambiguous — there's no state where both checks could
disagree.

The two existing status labels/timers (`lblRadeRX1Status`+`radeStatusTimer_Tick`
from #1, `lblFreeDVStatus`+`freedvStatusTimer_Tick` from the old FreeDV tab)
collapse into one label (`lblRadeRX1Status` kept, `lblFreeDVStatus` removed)
and one timer, reading `WDSP.GetRadaeSync`/`GetRadaeSnrDb` or
`WDSP.GetRXAFDVSync`/`GetRXAFDVSnr` depending on the active mode. Same text/
color convention both already use: `"SYNC   SNR {0:F1} dB"` (green) / `"no
sync"` / `"off"` / `"radio off"` (default color).

### Interlock — enforced low, not just in the new control

Mutual exclusion is enforced in the property setters themselves, not just in
the Mode combo's handler, so it holds no matter which entry point is used
(new Mode combo, old CAT commands, or any future caller):

- `RXAFDVRun`'s setter (radio.cs `DSPRX`) sets `RXRadaeEnabled = 0` on the
  same instance when arming (`value == 1`).
- `RXRadaeEnabled`'s setter (same class) sets `RXAFDVRun = 0` when arming.
- `TXAFDVRun`'s setter (radio.cs `DSPTX`) calls `WDSP.SetRadaeTxEnabled(0)`
  when arming.
- RADE TX-arm call sites (the Mode combo's handler, and CAT `ZZDK`) call
  `console.radio.GetDSPTX(0).TXAFDVRun = 0` when arming RADE TX. RADE TX
  stays a plain-WDSP-global with no new radio.cs property — matching #1's
  established wiring pattern — the interlock is just written as an explicit
  cross-call at both arm sites instead of a symmetric property pair.

Net effect: the Mode combo's handler only needs to *arm* the newly selected
subsystem (or explicitly disarm everything for "Off") — arming a subsystem's
RX/TX already clears the other side as a side effect of the interlocked
setters, so the combo doesn't need its own disarm-the-old-mode logic.

### Gray-out behavior

The Mode combo's `SelectedIndexChanged` handler toggles `.Enabled` on: the RX
Level spinner, the whole Mic/TX Conditioning group box, and the Diagnostics
bypass group box — enabled only when mode is RADE V1 or RADE V2, disabled
(grayed, not hidden) for Off and 700E, since none of those controls do
anything for 700E's DSP chain. Same rule applies to the Loopback checkbox as
noted above (disabled for Off specifically, not for 700E — 700E's loopback
bridge is a real, meaningful control when 700E is the active mode).

### Old tab removal

`tpDSPFreeDV` currently holds only `grpFreeDV` (its other child, `grpRADE`,
was already removed in #1). Once `grpFreeDV`/`chkFreeDVDecode`/
`lblFreeDVStatus` are removed as part of folding 700E into the Mode selector,
`tpDSPFreeDV` becomes empty and is removed entirely — same call #1 made for
`grpRADE`, applied consistently here.

**Designer-file mechanics** (from #1's experience, expected to repeat
identically): WinForms designer code touches a removed/renamed control in
(at least) 5 places — instantiation, `SuspendLayout`/`BeginInit`, the
properties/`Controls.Add` block, `ResumeLayout`/`EndInit`, and the private
field declaration — miss any one and it's either a compile error or (for
`NumericUpDownTS` specifically) a missing `ISupportInitialize` call. Edit via
byte-precise Python splicing preserving the CRLF cluster (`setup.designer.cs`
is 100% CRLF; `CATCommands.cs` is a genuine CRLF/LF mix — the plain-text Edit
tool flattened an entire ~10,000-line file to LF once during #1's CAT-command
work and had to be reverted; Python byte-splicing is the only proven-safe
method for these two files on this branch).

### Renamed/removed controls

| Control | Change |
|---|---|
| `cmbRadeProtocol` | Renamed `cmbRadeMode`; items become `{"Off","700E","RADE V1","RADE V2"}`; moves from the standalone Protocol group into RX1 Core |
| `grpRadeProtocol` | Removed (its one child relocates, see above) |
| `chkRadeRX1Enable` | Removed (enable now implicit in `cmbRadeMode`) |
| `chkRadeRX1Loopback` | Kept, handler rewritten mode-aware |
| `lblRadeRX1Status` | Kept, becomes the single unified status label |
| `grpFreeDV`, `chkFreeDVDecode`, `lblFreeDVStatus` | Removed |
| `tpDSPFreeDV` | Removed |
| Tab `tpDSPRADE`'s `Text` | Changed from `"RADE"` to `"Digital Voice"` (control `Name` unchanged) |

### CAT commands

One new command, `ZZEX` (next free `ZZEx` code): sets/gets the mode index —
`0`=Off, `1`=700E, `2`=RADE V1, `3`=RADE V2. Same 1-char convention as the
existing boolean commands' family, `nSet=1`/`nGet=0`/`nAns=1`.

The existing per-subsystem commands (`ZZDV`/`ZZEF`/`ZZEG` for 700E, `ZZDW`/
`ZZDK`/`ZZDL` for RADE) are **not removed or deprecated** — they keep working
exactly as documented in #1's spec, and now simply participate in the same
low-level interlock as everything else (e.g. `ZZDV1` while RADE is active
silently disarms RADE too, same as flipping `cmbRadeMode` would). No new
commands are needed for the loopback/status unification since neither was
separately CAT-addressable before this either.

### Testing

Same approach as #1: CI build compiles clean → deploy to `hl2winbox` →
confirm `ZZEX` get/set round-trips via a raw CAT socket script (same one
built for #1's testing, since `thetisctl`'s `query` subcommand is get-only) →
confirm selecting each of the 4 modes correctly arms/disarms both subsystems
(verified via the *existing* low-level CAT commands as an independent check —
e.g. after `ZZEX2` (RADE V1), confirm `ZZDW` reads `1` and `ZZDV` reads `0`)
→ confirm the gray-out behavior and unified status label visually via
screenshot → confirm the old FreeDV tab is gone and Setup still opens without
exceptions.

## Explicitly out of scope

- MOX/PTT arbiter for either subsystem's TX encode — both remain inert
  (armed but unable to key real audio) exactly as they are today; this
  sub-project only unifies how they're armed, not whether arming them does
  anything on TX.
- RX2/dual-channel support (sub-project #3).
- FreeDV Reporter network integration (sub-project #4).
- Panadapter overlays (sub-project #5).
- Settings persistence across restart (same known, documented, pre-existing
  gap noted in #1's spec — every control here inherits it unchanged).
- Adding any new FreeDV codec2 rate beyond 700E — `wdsp/fdv.c` hardcodes
  `FREEDV_MODE_700E`; making that configurable would be new DSP work, not a
  fit for a menu-reorganization sub-project.
