# RADE Core Setup Panel — Design Spec

**Date**: 2026-08-24
**Status**: Implemented — UI, CAT, and backend wiring all built, deployed, and
verified live on `hl2winbox` (`git:510d90d9`). See "Testing" (As executed) and
"Bug found during implementation" below.
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
| RX level gain | `SetRadaeRxScale`/`SetRadaeRxDialScale` | ✅ | `ZZEO` | ✅ `udRadeRxLevel` |
| Mic level gain | `SetRadaeMicScale` | ✅ | `ZZEC` | ✅ `udRadeMicLevel` |
| Mic RNNoise enable | `SetRadaeMicRNNoiseEnabled` | ✅ | `ZZED` | ✅ `chkRadeMicRNNoise` |
| Mic AGC enable + target | `SetRadaeMicAGCEnabled`/`...AGCTargetLufs` | ✅ | `ZZEE`/`ZZEH` | ✅ `chkRadeMicAGC`/`udRadeMicAGCTarget` |
| Mic EQ enable + 3-band + vol | `SetRadaeMicEQEnabled`/`...Bass`/`...Mid`/`...Treble`/`...Vol` | ✅ | `ZZEI`/`ZZEJ`/`ZZEK`/`ZZEL`/`ZZEN` | ✅ `chkRadeMicEQ` + 8 band spinners |
| Protocol V1/V2 | `SetRadaeProtocolV2`/`Get...` | ✅ | `ZZEP` | ✅ `cmbRadeProtocol` |
| Diagnostics bypass ×5 | `SetRadaeBypassMicDsp`/`...EncoderCore`/`...Rmatch`/`...Encoder`/`...All` | ✅ | `ZZEQ`/`ZZES`/`ZZEU`/`ZZEV`/`ZZEW` | ✅ 5 bypass checkboxes |

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

16 new commands, all in the `ZZEx` family (next free codes after `ZZEA`/`ZZEB`/
`ZZEF`/`ZZEG`/`ZZEM`/`ZZER`/`ZZET`, which were already taken). Boolean commands
follow the existing `ZZDW`/`ZZDK`-style 1-char `"0"`/`"1"` convention; signed
numeric commands use a 3-char `<sign><2-digit>` field; the three EQ-band
commands pack freq(+gain[+Q]) into one combined field since the backend takes
them as a single call (`SetRadaeMicEQBass(freq, gain)` etc.) — `ZZEJ`/`ZZEK`/
`ZZEL` carry the gain sign mid-suffix (not leading), so they were added to
`CATParser`'s `FindSuffix` regex-exception list alongside the existing `ZZDJ`.

| Code | Setting | Format | nSet |
|---|---|---|---|
| `ZZEC` | Mic level (dB) | `<sign><2-digit>`, e.g. `-15` | 3 |
| `ZZED` | Mic RNNoise enable | `0`/`1` | 1 |
| `ZZEE` | Mic AGC enable | `0`/`1` | 1 |
| `ZZEH` | Mic AGC target (LUFS) | `<sign><2-digit>` | 3 |
| `ZZEI` | Mic EQ enable | `0`/`1` | 1 |
| `ZZEJ` | Mic EQ bass (freq+gain) | `<freq 4-digit><sign><gain 2-digit>` | 7 |
| `ZZEK` | Mic EQ mid (freq+gain+Q) | `<freq 4-digit><sign><gain 2-digit><Q×100 3-digit>` | 10 |
| `ZZEL` | Mic EQ treble (freq+gain) | `<freq 5-digit><sign><gain 2-digit>` | 8 |
| `ZZEN` | Mic EQ output/makeup gain (dB) | `<sign><2-digit>` | 3 |
| `ZZEO` | RX1 decoder-input level (dB) | `<sign><2-digit>` | 3 |
| `ZZEP` | Protocol (0=V1, 1=V2) | `0`/`1` | 1 |
| `ZZEQ` | Diagnostics: bypass Mic DSP | `0`/`1` | 1 |
| `ZZES` | Diagnostics: bypass Encoder Core | `0`/`1` | 1 |
| `ZZEU` | Diagnostics: bypass Rate Match | `0`/`1` | 1 |
| `ZZEV` | Diagnostics: bypass Entire Encoder | `0`/`1` | 1 |
| `ZZEW` | Diagnostics: bypass ALL | `0`/`1` | 1 |

RX1 enable/loopback already had CAT commands from the earlier prototype work
(`ZZDW`/`ZZDL`) and were left as-is — no new codes needed for those two.

### Testing

No local Windows environment exists to render/exercise WinForms UI directly —
`hl2winbox` is the only verification path. Per control: CI build compiles
clean → deploy → confirm CAT get/set round-trips via `thetisctl` → confirm the
UI control reflects and drives the same backend state → confirm the tab loads
without exceptions and reads correct initial values on open. No new C-level DSP
logic is being written, so no signal-quality testing is needed — the
RNNoise/AGC/EQ/protocol-version code already runs in production, it has simply
never had a switch.

**As executed**: a real screenshot of the deployed panel (via a scheduled-task/
interactive-session capture, since a plain SSH-invoked script can't touch the
real desktop session) confirmed all four groups render correctly with no
overlapping controls and no exceptions. All 16 CAT commands were then
round-trip tested live via a raw socket script mirroring `thetisctl`'s wire
protocol (`thetisctl` itself only exposes get-only `query`, not arbitrary
sets) — every GET matched the UI's on-screen values exactly, and every SET
followed by GET reflected the change, including the mid-string-sign combined
EQ-band fields. One real bug surfaced during this pass — see below — found,
fixed, redeployed, and reverified before values were restored to baseline.

## Bug found during implementation: Mic/RX Level was linear scale, not dB

`SetRadaeMicScale`/`SetRadaeRxScale` (`radae.c`) take a *linear* gain factor,
clamped there to `(0, 100]` and defaulting to unity (`1.0`) for any value
`<= 0`. The `udRadeMicLevel`/`udRadeRxLevel` spinners and their `ZZEC`/`ZZEO`
CAT commands were built passing the raw `-40..+40` spinner value straight
through as if it *were* the linear scale — so any negative value silently
collapsed to 0dB (unity gain), killing the entire bottom half of each
control's range.

Found live during the CAT round-trip pass: `ZZEO` set to `-20` read back `+01`
unchanged, while the same test on `ZZEC` with a positive value passed first,
masking the bug until the negative case was tried. The `-40..+40dB` range
maps exactly onto the `0.01..100.0` linear clamp (`20·log₁₀`), confirming dB
was the intended design all along — the log/exp conversion at the boundary
was just never written. Fixed in both directions: `udRadeMicLevel_ValueChanged`
/`udRadeRxLevel_ValueChanged` and `InitRadePanelFromBackend`'s reverse sync
(`setup.cs`), and `ZZEC`/`ZZEO` (`CATCommands.cs`) to match. Redeployed and
reverified live — `-40`, `-20`, `0`, and `+40` all round-trip exactly now.

## Explicitly out of scope

- 700E integration into this menu (sub-project #2)
- RX2/dual-channel support (sub-project #3, needs new backend)
- FreeDV Reporter network integration (sub-project #4)
- Panadapter overlays (sub-project #5)
- Settings persistence across restart

**Resolved during implementation** (was listed as out of scope, deferred a
decision rather than left undone): the RX level spinner drives `RxScale`
only. `RxDialScale` remains unwired — confirmed zero callers anywhere in the
codebase before leaving it untouched, same as the original design note
anticipated.
