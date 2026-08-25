# RADE Core Setup Panel — Design Spec

**Date**: 2026-08-24
**Status**: Approved, ready for implementation planning
**Sub-project**: 1 of 5 (see "Larger context" below)

## Goal

Duplicate sv1eia's `Setup → DSP → RADE` menu into this Hermes-Lite 2 fork, scoped
to the slice of it that maps onto backend functionality already implemented in
`ChannelMaster/radae.c` on this branch — no new DSP/C code, UI + wiring only.

## Larger context (not part of this spec)

The full ask ("duplicate everything, then fold 700E in as an option") spans five
independent pieces. This spec covers only #1. The rest are deferred to their own
future design → spec → plan cycles, in this order:

1. **RADE core + mic/TX conditioning UI panel** *(this spec)*
2. 700E folded into the same menu as a mode option — no sv1eia precedent exists
   for this; genuinely open design question, not decided here.
3. RX2/dual-channel RADE support — needs new backend work, not just UI.
4. FreeDV Reporter network integration (qso.freedv.org spotting/viewing) — a
   separate networking subsystem; this project's own docs already treat it as a
   distinct future stage.
5. Panadapter SNR/sync overlay display — touches `display.cs`, independent of
   the Setup-form work entirely.

## Source reference

sv1eia's actual menu (`Project Files/Source/Console/setup.designer.cs` on their
`main` branch, accessed via `git show sv1eia/main:"..."` from the peer session
working `trial-merge-upstream` — no screenshot was available, full control
inventory pulled from designer source instead) is a single flat tab page
(`tpDSPRADE`), not sub-tabbed, loosely grouped by function. Two orthogonal
controls, not an N-way mode picker: an enable checkbox per RX (RADE encode/
decode armed independently per receiver) and a separate `{"V1","V2"}` combo box
selecting which protocol version that RX's decoder uses.

700E does not exist anywhere in sv1eia's design — confirmed via a full-text
search of their designer source. Every "FreeDV" string in their tooltips
explicitly means RADE V1. This fork's 700E support is entirely our own; there is
no precedent to reuse for how it should eventually fold into a shared menu
(sub-project #2's problem, not this one's).

## This fork's current state (baseline)

`Setup → DSP → FreeDV` tab (`tpDSPFreeDV`) holds two small side-by-side group
boxes, 235×76px each: `grpFreeDV` (one checkbox, "Decode FreeDV 700E (RX1)",
`chkFreeDVDecode` → `RXAFDVRun`) and `grpRADE` (one checkbox, "Decode RADE V1
(RX1)", `chkRADEDecode` → `RXRadaeEnabled`). No TX controls, no meters, no mic
conditioning, no protocol selection — everything else in `radae.c`'s public
surface is either CAT-only or has zero C#/CAT exposure at all.

**Backend audit** (verified directly against `radae.h`/`dsp.cs`/`CATCommands.cs`,
not assumed):

| Setting | C function | C# (`dsp.cs`) | CAT | UI |
|---|---|---|---|---|
| RX1 decode enable | `SetRadaeRxEnabled`/`Get...` | ✅ (via `radio.cs`) | `ZZDW` | ✅ `chkRADEDecode` |
| TX encode enable | `SetRadaeTxEnabled`/`Get...` | ✅ | `ZZDK` | ❌ |
| RX1 loopback | `SetRadaeLoopbackEnabled`/`Get...` | ✅ | `ZZDL` | ❌ |
| RX sync/SNR | `GetRadaeSync`/`GetRadaeSnrDb` | ✅ | `ZZDZ` | crude "off" label only |
| RX level/clip meter | `GetRadaeRxLevelDb`/`GetRadaeClip` | ✅ | `ZZDT` | ❌ |
| EOO callsign | `SetRadaeEooCallsign`/`Get...` | ✅ | `ZZDJ` | ❌ |
| RX level gain | `SetRadaeRxScale`/`SetRadaeRxDialScale` | ❌ | ❌ | ❌ (dead code, zero callers today) |
| Mic level gain | `SetRadaeMicScale` | ❌ | ❌ | ❌ |
| Mic RNNoise enable | `SetRadaeMicRNNoiseEnabled` | ❌ | ❌ | ❌ |
| Mic AGC enable + target | `SetRadaeMicAGCEnabled`/`...AGCTargetLufs` | ❌ | ❌ | ❌ |
| Mic EQ enable + 3-band + vol | `SetRadaeMicEQEnabled`/`...Bass`/`...Mid`/`...Treble`/`...Vol` | ❌ | ❌ | ❌ |
| Protocol V1/V2 | `SetRadaeProtocolV2`/`Get...` | ❌ | ❌ | ❌ |
| Diagnostics bypass ×5 | `SetRadaeBypassMicDsp`/`...EncoderCore`/`...Rmatch`/`...Encoder`/`...All` | ❌ | ❌ | ❌ |

Everything in the bottom seven rows has real, already-working DSP logic behind
it (RNNoise/AGC/EQ/protocol-version code runs today, just has no way to be
switched) — this is pure wiring work, not new signal processing.

## Design

### Layout

New tab page `tpDSPRADE`, labeled "RADE", inserted next to the existing
"FreeDV" tab. Four grouped sections on one flat page (matching sv1eia's own
flat-page structure):

- **RX1 Core**: enable checkbox (wires to the *existing* `RXRadaeEnabled`
  backend — same one `chkRADEDecode` already drives), loopback-test checkbox
  (wires to existing `ZZDL`/`SetRadaeLoopbackEnabled`), RX level spinner (new).
- **Mic/TX Conditioning**: mic level spinner (new), RNNoise enable (new), AGC
  enable + LUFS target spinner (new, default −23, range −30..0, matching
  sv1eia's own default), EQ enable + 3-band (bass/mid/treble — freq+gain each,
  mid also gets Q) + master vol spinner (new).
- **Protocol**: V1/V2 combo box (new).
- **Diagnostics** group box, "boots OFF" (matching sv1eia's own framing): 5
  bypass checkboxes (new).

The old `grpRADE` box (just the one enable checkbox) is removed from the
"FreeDV" tab — redundant once the new tab's RX1-core enable drives the exact
same backend property. `grpFreeDV` (700E) is left completely untouched;
reorganizing it is sub-project #2's job, not this one's.

### Wiring pattern

Every setting here is a plain `radae.c` (ChannelMaster) global with no wdsp
channel and no "survive a DSP rebuild" lifecycle — the same shape as
`SetRadaeTxEnabled`/`SetRadaeLoopbackEnabled`/`SetRadaeEooCallsign` (all built
tonight, all called directly from `CATCommands.cs` with **no** `radio.cs`
property indirection, unlike `RXRadaeEnabled` which genuinely needs that
indirection to survive channel rebuilds).

Every new control here follows that same direct-call pattern: each Setup
control's event handler calls `WDSP.SetXxx` directly; the tab's `Load`/`Show`
reads current state via the matching `GetXxx` so the panel reflects reality
even if something was changed via CAT while Setup was closed.

### Persistence — explicitly deferred

This fork has a known, documented, open bug: FreeDV/RADE checkbox state does
not survive a Thetis restart. Every setting in this panel inherits that same
gap; **building persistence is out of scope for this sub-project**, matching
current behavior exactly rather than fixing (or extending) it here. A
project-wide persistence fix, if wanted, is its own separate, well-scoped task.

### CAT commands

~16 new commands (RX-level gain, mic-level gain, RNNoise enable, AGC enable,
AGC LUFS target, EQ enable, EQ bass, EQ mid, EQ treble, EQ vol, protocol V1/V2,
+5 diagnostics bypass flags), using the established `ZZDx`/`ZZEx` free-code
convention (same family as tonight's `ZZDJ`/`ZZDK`/`ZZDL`/`ZZEF`/`ZZEG`). Exact
code letters get picked during implementation.

### Testing

No local Windows environment exists to render/exercise WinForms UI directly —
`hl2winbox` is the only verification path. Per control: CI build compiles
clean → deploy → confirm CAT get/set round-trips via `thetisctl` → confirm the
UI control reflects and drives the same backend state → confirm the tab loads
without exceptions and reads correct initial values on open. No new C-level DSP
logic is being written, so no signal-quality testing is needed — the
RNNoise/AGC/EQ/protocol-version code already runs in production, it has simply
never had a switch.

## Explicitly out of scope

- 700E integration into this menu (sub-project #2)
- RX2/dual-channel support (sub-project #3, needs new backend)
- FreeDV Reporter network integration (sub-project #4)
- Panadapter overlays (sub-project #5)
- Settings persistence across restart
- Exact `RxScale` vs `RxDialScale` semantic mapping for the single "RX level"
  spinner sv1eia shows — both C functions exist under related names; which one
  (or both) the UI control should drive gets resolved during implementation by
  reading `radae.c`'s own use of each, not guessed here.
