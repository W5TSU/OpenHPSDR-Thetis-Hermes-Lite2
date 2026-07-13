# `wdsp/dexp.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Downward expander / noise gate with VOX tie-in.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×5)
  - `wdsp/firmin.c` (calls ×4)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
- Most-referenced symbols from other files: `create_dexp()` (×1), `destroy_dexp()` (×1), `xdexp()` (×1), `SetDEXPRate()` (×1), `SetDEXPSize()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_delring()`** — L31 — `DELRING calc_delring (int rsize, int size, int delay, double* in, double* out)`
  Called by: `calc_dexp()` (same file), `calc_filter()` (same file)
- **`decalc_delring()`** — L45 — `void decalc_delring (DELRING a)`
  Called by: `decalc_dexp()` (same file), `decalc_filter()` (same file)
- **`flush_delring()`** — L51 — `void flush_delring (DELRING a)`
  Flushes (zeroes) the `delring` block’s internal buffers/state.
  Called by: `flush_dexp()` (same file)
- **`xdelring()`** — L58 — `void xdelring (DELRING a)`
  Runs the `delring` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xdexp()` (same file)
- **`calc_slews()`** — L91 — `void calc_slews (DEXP a)`
  Called by: `calc_dexp()` (same file)
- **`calc_buffs()`** — L111 — `void calc_buffs (DEXP a)`
  Called by: `create_dexp()` (same file), `SetDEXPSize()` (same file)
- **`decalc_buffs()`** — L118 — `void decalc_buffs (DEXP a)`
  Called by: `destroy_dexp()` (same file), `SetDEXPSize()` (same file)
- **`calc_dexp()`** — L125 — `void calc_dexp (DEXP a)`
  Called by: `create_dexp()` (same file), `SetDEXPSize()` (same file), `SetDEXPIOBuffers()` (same file), `SetDEXPRate()` (same file), `SetDEXPDetectorTau()` (same file), `SetDEXPAttackTime()` (same file) — and 6 more
- **`decalc_dexp()`** — L149 — `void decalc_dexp (DEXP a)`
  Called by: `destroy_dexp()` (same file), `SetDEXPSize()` (same file), `SetDEXPIOBuffers()` (same file), `SetDEXPRate()` (same file), `SetDEXPDetectorTau()` (same file), `SetDEXPAttackTime()` (same file) — and 6 more
- **`calc_filter()`** — L156 — `void calc_filter (DEXP a)`
  Called by: `create_dexp()` (same file), `SetDEXPSize()` (same file), `SetDEXPIOBuffers()` (same file), `SetDEXPRate()` (same file), `SetDEXPFilterTaps()` (same file), `SetDEXPWindowType()` (same file) — and 2 more
- **`decalc_filter()`** — L168 — `void decalc_filter (DEXP a)`
  Called by: `destroy_dexp()` (same file), `SetDEXPSize()` (same file), `SetDEXPIOBuffers()` (same file), `SetDEXPRate()` (same file), `SetDEXPFilterTaps()` (same file), `SetDEXPWindowType()` (same file) — and 2 more
- **`calc_antivox()`** — L174 — `void calc_antivox(DEXP a)`
  Called by: `create_dexp()` (same file), `SetAntiVOXSize()` (same file), `SetAntiVOXRate()` (same file), `SetAntiVOXDetectorTau()` (same file)
- **`decalc_antivox()`** — L181 — `void decalc_antivox(DEXP a)`
  Called by: `destroy_dexp()` (same file), `SetAntiVOXSize()` (same file), `SetAntiVOXRate()` (same file), `SetAntiVOXDetectorTau()` (same file)
- **`create_dexp()`** — L186 — `PORT void create_dexp (int id, int run_dexp, int size, double* in, double* out, int rate, double dettau, double tattack, double tdecay, double thold, double exp_ratio, double hyst_`
  Constructor for the `dexp` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_xmtr()` (`ChannelMaster/cmaster.c`)
- **`destroy_dexp()`** — L229 — `PORT void destroy_dexp (int id)`
  Destroys the `dexp` block, freeing its allocated buffers.
  Called by: `destroy_xmtr()` (`ChannelMaster/cmaster.c`)
- **`flush_dexp()`** — L241 — `PORT void flush_dexp (int id)`
  Flushes (zeroes) the `dexp` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xdexp()`** — L265 — `PORT void xdexp (int id)`
  Runs the `dexp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`SendCBPushDexpVox()`** — L398 — `PORT void SendCBPushDexpVox (int id, void (__stdcall *pushvox)(int id, int active))`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetDEXPRun()`** — L406 — `PORT void SetDEXPRun (int id, int run)`
  Sets dexprun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPSize()`** — L416 — `PORT void SetDEXPSize (int id, int size)`
  Sets dexpsize — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetDEXPIOBuffers()`** — L435 — `PORT void SetDEXPIOBuffers (int id, double* in, double* out)`
  Sets dexpiobuffers — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetDEXPRate()`** — L450 — `PORT void SetDEXPRate (int id, double rate)`
  Sets dexprate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetDEXPDetectorTau()`** — L465 — `PORT void SetDEXPDetectorTau (int id, double tau)`
  Sets dexpdetector tau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPAttackTime()`** — L478 — `PORT void SetDEXPAttackTime (int id, double time)`
  Sets dexpattack time — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPReleaseTime()`** — L491 — `PORT void SetDEXPReleaseTime (int id, double time)`
  Sets dexprelease time — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPHoldTime()`** — L504 — `PORT void SetDEXPHoldTime (int id, double time)`
  Sets dexphold time — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPExpansionRatio()`** — L517 — `PORT void SetDEXPExpansionRatio (int id, double ratio)`
  Sets dexpexpansion ratio — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPHysteresisRatio()`** — L530 — `PORT void SetDEXPHysteresisRatio (int id, double ratio)`
  Sets dexphysteresis ratio — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPAttackThreshold()`** — L543 — `PORT void SetDEXPAttackThreshold (int id, double thresh)`
  Sets dexpattack threshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPFilterTaps()`** — L555 — `PORT void SetDEXPFilterTaps (int id, int taps)`
  Sets dexpfilter taps — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetDEXPWindowType()`** — L567 — `PORT void SetDEXPWindowType (int id, int type)`
  Sets dexpwindow type — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetDEXPLowCut()`** — L581 — `PORT void SetDEXPLowCut (int id, double lowcut)`
  Sets dexplow cut — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPHighCut()`** — L593 — `PORT void SetDEXPHighCut (int id, double highcut)`
  Sets dexphigh cut — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPRunSideChannelFilter()`** — L605 — `PORT void SetDEXPRunSideChannelFilter (int id, int run)`
  Sets dexprun side channel filter — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPRunVox()`** — L615 — `PORT void SetDEXPRunVox (int id, int run)`
  Sets dexprun vox — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPRunAudioDelay()`** — L625 — `PORT void SetDEXPRunAudioDelay (int id, int run)`
  Sets dexprun audio delay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetDEXPAudioDelay()`** — L635 — `PORT void SetDEXPAudioDelay (int id, double delay)`
  Sets dexpaudio delay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`GetDEXPPeakSignal()`** — L647 — `PORT void GetDEXPPeakSignal (int id, double* peak)`
  Returns dexppeak signal — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetAntiVOXRun()`** — L656 — `PORT void SetAntiVOXRun (int id, int run)`
  Sets anti voxrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetAntiVOXSize()`** — L665 — `PORT void SetAntiVOXSize (int id, int size)`
  Sets anti voxsize — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAntiVOXRate()`** — L676 — `PORT void SetAntiVOXRate (int id, double rate)`
  Sets anti voxrate — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAntiVOXGain()`** — L687 — `PORT void SetAntiVOXGain (int id, double gain)`
  Sets anti voxgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetAntiVOXDetectorTau()`** — L696 — `PORT void SetAntiVOXDetectorTau (int id, double tau)`
  Sets anti voxdetector tau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SendAntiVOXData()`** — L707 — `PORT void SendAntiVOXData (int id, int nsamples, double* data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/dexp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
