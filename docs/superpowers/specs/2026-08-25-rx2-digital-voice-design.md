# RX2 Digital Voice Support — Design Spec

**Date**: 2026-08-25
**Status**: Approved, ready for implementation planning
**Sub-project**: 3 of 5 (see [2026-08-24-rade-setup-panel-design.md](2026-08-24-rade-setup-panel-design.md)'s
"Larger context" section)

## Goal

Extend the Digital Voice Mode selector (sub-project #2) to RX2, giving it the
same Off/700E/RADE V1/RADE V2 choice RX1 already has — Mode combo, Loopback
test, RX Level, and status readout — plus matching CAT commands.

## Finding that reframed this sub-project's scope

Both #1's and #2's specs listed RX2/dual-channel support as needing "new
backend work, not just UI." That assumption doesn't hold up under a closer
read of `ChannelMaster/radae.c` and `wdsp/fdv.c`, done at the start of this
sub-project's brainstorming:

- `radae.c` has a header comment reading *"Dual-RX support. RADE decode runs
  concurrently on RX1 and RX2"* — not aspirational. `g_rade[RADAE_NRX]` holds
  two fully independent codec handles (own model instance, FIFOs, buffers,
  lock each); `pipe.c`'s `xpipe()` already calls `xradae_rx(rx, ...)` on
  *both* RX1's and RX2's real audio streams today, unconditionally at the
  ChannelMaster level (gated only by the existing per-rx `g_radae_rx_enabled[rx]`
  flag, which nothing in the C#/CAT/UI layer has ever set to anything but 0
  for `rx=1`). Every relevant getter/setter (`SetRadaeRxEnabled`,
  `GetRadaeSync`, `SetRadaeLoopbackEnabled`, `SetRadaeRxScale`, etc.) already
  takes an `rx` parameter.
- 700E's `wdsp/fdv.c` state (`rxa[channel].fdv.p`) is indexed by wdsp's own
  generic per-channel mechanism (`WDSP.id(thread, subrx)`) — the same one
  every other RXA feature (filters, AGC, NR, ...) already relies on for RX2
  to work at all. No 700E-specific dual-RX gap exists either.
- TX encode already supports targeting either RX's handle
  (`SetRadaeTxRx(rx)`, "0 = RX1, 1 = RX2, ... set at the MOX edge").

**The actual gap**: sub-projects #1 and #2 both deliberately scoped their
UI *and* CAT layer to RX1 only (`GetDSPRX(0, 0)`, `rx=0` hardcoded
throughout the C# layer) — the right call for those sub-projects, but it
left RX2 with full backend support and zero UI/CAT exposure. Plain
(non-RADE) RX2 reception already works in this app today (the existing
"RX2" toggle in the main console, `console.RX2Enabled`).

This sub-project is therefore **UI/CAT generalization to a second
receiver**, not new DSP engineering — a materially smaller shape than #1's
spec assumed. It also means RX2's RADE decode path will be **exercised
end-to-end for the first time** during this sub-project's testing, despite
having been structurally present since #1.

## Design

### Scope: RX2 Core only, not a full RX1-panel clone

Two of the Digital Voice tab's four groups are TX-side, process-wide
concerns, not per-receiver:

- **Mic/TX Conditioning** (`radae_micdsp.c`: `SetRadaeMicScale`,
  `SetRadaeMicRNNoiseEnabled`, `SetRadaeMicAGCEnabled`, `SetRadaeMicEQ*`) —
  none of these functions take an `rx` parameter. They feed the single
  shared TX encoder, not a specific receiver.
- **Diagnostics** (bypass ladder) — same: `SetRadaeBypassMicDsp` and
  siblings are global, not per-rx.

So RX2 needs only an "RX2 Core" equivalent to RX1's — Mode, Loopback, RX
Level, Status — not a duplicate of all four groups. Confirmed by reading
`radae_micdsp.c`'s function signatures directly, not assumed.

### TX semantics: RX2's Mode selector never touches TX

RX1 and RX2 can each independently be in any mode (e.g. RX1 = RADE V1, RX2
= 700E simultaneously) — decode is genuinely concurrent per RX. TX encode
is a single shared resource that only ever follows whichever RX is actually
keyed (`SetRadaeTxRx`). Per the decision made in this sub-project's
brainstorming: **RX2's Mode combo arms/disarms RX2's own RX decode only —
no TX side effects.** TX arming stays exactly as sub-project #2 built it:
RX1's Mode combo remains the only Mode selector that also arms TX. This
avoids two Mode combos racing to control one shared TX encoder, at the cost
of a deliberate asymmetry between the two combos' behavior (RX1's does more
than RX2's) — an accepted tradeoff, not an oversight.

RX2's own decode interlock (700E vs. RADE V1/V2, within RX2 only) follows
the exact same pattern sub-project #2 built for RX1: `RXAFDVRun`/
`RXRadaeEnabled` on `console.radio.GetDSPRX(1, 0)` interlock each other via
the same radio.cs property-setter mechanism already in place (that
mechanism is per-instance, so it applies automatically to the `(1, 0)`
instance with no code changes — see #2's spec, "Interlock" section, for why
this was built at the property level rather than only in the UI).

### RX2 Core group — compact 2-row layout

The Digital Voice tab (724×414) has no room for a same-sized fourth group:
Mic/TX Conditioning fills the left column (320×300, content-bound — its 9
rows can't shrink); RX1 Core (340×136) and Diagnostics (340×160) stack in
the right column with only a 76px gap between them. Per the layout decision
made in this sub-project's brainstorming, RX2 Core fills that gap
(`Location(352,152)`, `Size(340,76)`) with a compact 2-row layout instead
of RX1 Core's spacious 4-row one:

- **Row 1**: "Mode:" label + `cmbRadeRX2Mode` (same 4 items: Off/700E/RADE
  V1/RADE V2) + a "Loopback" checkbox, side by side.
- **Row 2**: RX Level spinner + status label, side by side.

Exact pixel placement within that 340×76 box (label widths, combo width,
checkbox text) is finalized during implementation — same deferral #1's spec
used for exact CAT-code letters, since it depends on how the real controls
actually measure once laid out, not something worth guessing precisely here.

### Gray-out: tied to `console.RX2Enabled`

The whole `grpRadeRX2Core` group is disabled (grayed) whenever
`console.RX2Enabled` is `false` — a RADE/700E mode on a receiver that isn't
even running is meaningless. This needs a way to react to RX2 being
toggled on/off *while Setup is open* (the main console's RX2 button isn't
inside Setup) — implementation should check whether `console.RX2Enabled`
already fires an event/notification this codebase can hook, or fall back to
re-checking it on the same 500ms status-timer tick RX2's status readout
already needs, whichever is cleaner given what's actually available; this
is an implementation detail, not a design fork worth a separate decision.

### CAT commands

Four new commands, next free codes in the `ZZFx` family (the `ZZEx` family
used for #2's 16 commands plus this decision's own is down to its last two
free letters — `ZZFx` has plenty of room):

| Code | Setting | Mirrors |
|---|---|---|
| `ZZFC` | RX2 mode (0=Off/1=700E/2=RADE V1/3=RADE V2) | `ZZEX` |
| `ZZFE` | RX2 loopback test enable | `ZZDL` |
| `ZZFG` | RX2 sync/SNR status (get-only) | `ZZDZ`/`ZZDS` |
| `ZZFK` | RX2 decoder-input level (dB) | `ZZEO` |

Same wire formats and conventions as their RX1 counterparts (see #1's and
#2's specs for the exact byte-level formats), with `rx=1` passed to the
underlying `radae.c`/`fdv.c` calls instead of `rx=0`. `ZZFC`'s set logic
mirrors `ZZEX`'s exactly except it never touches TX (per the TX-semantics
decision above) — arms/disarms only `console.radio.GetDSPRX(1, 0).RXAFDVRun`/
`.RXRadaeEnabled`.

Existing RX1 commands (`ZZEX`, `ZZDL`, `ZZDZ`, `ZZDS`, `ZZEO`, and all the
rest) are untouched.

### Testing

Same approach as #1/#2: CI build compiles clean → deploy to `hl2winbox` →
confirm CAT get/set round-trips for all four new commands via a raw CAT
socket script → confirm RX2's decode/loopback interlock behaves exactly
like RX1's did in #2's testing (arming 700E disarms RADE on RX2, and vice
versa, independent of whatever RX1 is doing) → confirm RX2 Core visibly
grays out when RX2 is powered off and re-enables when powered back on →
screenshot verification of the compact 2-row layout.

**This is also the first real end-to-end test of RX2 RADE decode ever
performed** — unlike #1/#2, which wired already-designed-for-RX1 code
paths, this sub-project's hardware test is the first time anyone confirms
`xradae_rx(1, ...)`'s dual-RX design actually decodes real off-air (or
loopback) audio on RX2, not just that it compiles and the CAT layer reports
sane values. If the loopback round-trip doesn't produce audible decoded
audio on RX2 the way #1's did on RX1, that's a genuine finding to surface,
not a testing-approach failure — see #1's spec for how the equivalent RX1
gap (missing TX-encode UI, discovered via live testing) was handled.

## Explicitly out of scope

- 700E TX / RADE TX per-RX targeting UI (`SetRadaeTxRx`) — TX remains a
  single shared resource controlled only via RX1's Mode selector and the
  existing MOX-edge mechanism; no new UI for choosing which RX transmits.
- Mic/TX Conditioning and Diagnostics duplication for RX2 — confirmed
  unnecessary, see "Scope" above.
- FreeDV Reporter network integration (sub-project #4).
- Panadapter overlays (sub-project #5).
- Settings persistence across restart — same known, documented,
  pre-existing gap noted in #1's and #2's specs; RX2's new controls inherit
  it unchanged, same as RX1's did. Per #2's final review, the *actual*
  persistence behavior for `ComboBoxTS`/`CheckBoxTS` controls is more
  subtle than "none": this codebase's generic Setup persistence restores
  every named control from disk and fires its change event for real during
  construction. Implementation should apply the same `initializing` guard
  convention #2's fix wave added to `cmbRadeMode` to `cmbRadeRX2Mode` (and
  `chkRadeRX2Loopback`, whichever name it ends up with) from the start,
  rather than rediscovering the same gap a second time.
