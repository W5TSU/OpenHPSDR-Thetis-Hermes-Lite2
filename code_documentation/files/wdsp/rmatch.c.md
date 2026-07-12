# `wdsp/rmatch.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Fixed and variable-ratio resamplers, and the adaptive rate-matcher that reconciles independent sample clocks.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/ivac.c` (calls ×19)
  - `ChannelMaster/cmasio.c` (calls ×9)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×5)
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/combase.h` (calls ×3)
  - `wdsp/varsamp.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `xrmatchIN()` (×5), `xrmatchOUT()` (×4), `forceRMatchVar()` (×3), `create_rmatchV()` (×2), `destroy_rmatchV()` (×2), `getRMatchDiags()` (×2), `resetRMatchDiags()` (×2), `setRMatchFeedbackGain()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_mav()`** — L29 — `MAV create_mav (int ringmin, int ringmax, double nom_value)`
  Constructor for the `mav` block: allocates its state/buffers and computes initial coefficients.
  Called by: `calc_rmatch()` (same file)
- **`destroy_mav()`** — L43 — `void destroy_mav (MAV a)`
  Destroys the `mav` block, freeing its allocated buffers.
  Called by: `decalc_rmatch()` (same file)
- **`flush_mav()`** — L49 — `void flush_mav (MAV a)`
  Flushes (zeroes) the `mav` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xmav()`** — L56 — `void xmav (MAV a, int input, double* output)`
  Runs the `mav` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `control()` (same file)
- **`create_aamav()`** — L71 — `AAMAV create_aamav (int ringmin, int ringmax, double nom_ratio)`
  Constructor for the `aamav` block: allocates its state/buffers and computes initial coefficients.
  Called by: `calc_rmatch()` (same file)
- **`destroy_aamav()`** — L86 — `void destroy_aamav (AAMAV a)`
  Destroys the `aamav` block, freeing its allocated buffers.
  Called by: `decalc_rmatch()` (same file)
- **`flush_aamav()`** — L92 — `void flush_aamav (AAMAV a)`
  Flushes (zeroes) the `aamav` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xaamav()`** — L101 — `void xaamav (AAMAV a, int input, double* output)`
  Runs the `aamav` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `control()` (same file)
- **`calc_rmatch()`** — L128 — `void calc_rmatch (RMATCH a)`
  Called by: `create_rmatch()` (same file), `reset_rmatch()` (same file), `setRMatchInsize()` (same file), `setRMatchOutsize()` (same file), `setRMatchNomInrate()` (same file), `setRMatchNomOutrate()` (same file) — and 6 more
- **`decalc_rmatch()`** — L174 — `void decalc_rmatch (RMATCH a)`
  Called by: `destroy_rmatch()` (same file), `reset_rmatch()` (same file), `setRMatchInsize()` (same file), `setRMatchOutsize()` (same file), `setRMatchNomInrate()` (same file), `setRMatchNomOutrate()` (same file) — and 6 more
- **`create_rmatch()`** — L187 — `RMATCH create_rmatch ( int run, double* in, double* out, int insize,`
  Constructor for the `rmatch` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rmatchV()` (same file), `create_rmatchLegacyV()` (same file)
- **`destroy_rmatch()`** — L241 — `void destroy_rmatch (RMATCH a)`
  Destroys the `rmatch` block, freeing its allocated buffers.
  Called by: `destroy_rmatchV()` (same file)
- **`reset_rmatch()`** — L247 — `void reset_rmatch (RMATCH a)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`control()`** — L256 — `void control (RMATCH a, int change)`
  Called by: `xrmatchIN()` (same file), `xrmatchOUT()` (same file)
- **`blend()`** — L275 — `void blend (RMATCH a)`
  Called by: `xrmatchIN()` (same file)
- **`upslew()`** — L285 — `void upslew (RMATCH a, int newsamps)`
  Called by: `xrmatchIN()` (same file)
- **`xrmatchIN()`** — L300 — `PORT void xrmatchIN (void* b, double* in)`
  Runs the `rmatchIN` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `asioOUT()` (`ChannelMaster/cmasio.c`), `CallbackASIO()` (`ChannelMaster/cmasio.c`), `xvacOUT()` (`ChannelMaster/ivac.c`), `xvac_out()` (`ChannelMaster/ivac.c`), `CallbackIVAC()` (`ChannelMaster/ivac.c`)
- **`dslew()`** — L364 — `void dslew (RMATCH a)`
  Called by: `xrmatchOUT()` (same file)
- **`xrmatchOUT()`** — L427 — `PORT void xrmatchOUT (void* b, double* out)`
  Runs the `rmatchOUT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `asioIN()` (`ChannelMaster/cmasio.c`), `CallbackASIO()` (`ChannelMaster/cmasio.c`), `xvacIN()` (`ChannelMaster/ivac.c`), `CallbackIVAC()` (`ChannelMaster/ivac.c`)
- **`getRMatchDiags()`** — L469 — `PORT void getRMatchDiags (void* b, int* underflows, int* overflows, double* var, int* ringsize, int* nring)`
  Called by: `getCMAevents()` (`ChannelMaster/cmasio.c`), `getIVACdiags()` (`ChannelMaster/ivac.c`)
- **`resetRMatchDiags()`** — L482 — `PORT void resetRMatchDiags (void* b)`
  Called by: `resetCMAevents()` (`ChannelMaster/cmasio.c`), `resetIVACdiags()` (`ChannelMaster/ivac.c`)
- **`forceRMatchVar()`** — L490 — `PORT void forceRMatchVar (void* b, int force, double fvar)`
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`), `create_resamps()` (`ChannelMaster/ivac.c`), `forceIVACvar()` (`ChannelMaster/ivac.c`)
- **`create_rmatchV()`** — L500 — `PORT void* create_rmatchV(int in_size, int out_size, int nom_inrate, int nom_outrate, int ringsize, double var)`
  Constructor for the `rmatchV` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`), `create_resamps()` (`ChannelMaster/ivac.c`)
- **`destroy_rmatchV()`** — L529 — `PORT void destroy_rmatchV (void* ptr)`
  Destroys the `rmatchV` block, freeing its allocated buffers.
  Called by: `destroy_cmasio()` (`ChannelMaster/cmasio.c`), `destroy_resamps()` (`ChannelMaster/ivac.c`)
- **`setRMatchInsize()`** — L536 — `PORT void setRMatchInsize (void* ptr, int insize)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setRMatchOutsize()`** — L548 — `PORT void setRMatchOutsize (void* ptr, int outsize)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setRMatchNomInrate()`** — L560 — `PORT void setRMatchNomInrate (void* ptr, int nom_inrate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setRMatchNomOutrate()`** — L572 — `PORT void setRMatchNomOutrate (void* ptr, int nom_outrate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setRMatchRingsize()`** — L584 — `PORT void setRMatchRingsize (void* ptr, int ringsize)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setRMatchFeedbackGain()`** — L596 — `PORT void setRMatchFeedbackGain (void* b, double feedback_gain)`
  Called by: `SetIVACFeedbackGain()` (`ChannelMaster/ivac.c`)
- **`setRMatchSlewTime()`** — L606 — `PORT void setRMatchSlewTime (void* b, double slew_time)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setRMatchSlewTime1()`** — L618 — `PORT void setRMatchSlewTime1(void* b, double slew_time)`
  Called by: `SetIVACSlewTime()` (`ChannelMaster/ivac.c`)
- **`setRMatchPropRingMin()`** — L641 — `PORT void setRMatchPropRingMin(void* ptr, int prop_min)`
  Called by: `SetIVACPropRingMin()` (`ChannelMaster/ivac.c`)
- **`setRMatchPropRingMax()`** — L653 — `PORT void setRMatchPropRingMax(void* ptr, int prop_max)`
  Called by: `SetIVACPropRingMax()` (`ChannelMaster/ivac.c`)
- **`setRMatchFFRingMin()`** — L665 — `PORT void setRMatchFFRingMin(void* ptr, int ff_ringmin)`
  Called by: `SetIVACFFRingMin()` (`ChannelMaster/ivac.c`)
- **`setRMatchFFRingMax()`** — L677 — `PORT void setRMatchFFRingMax(void* ptr, int ff_ringmax)`
  Called by: `SetIVACFFRingMax()` (`ChannelMaster/ivac.c`)
- **`setRMatchFFAlpha()`** — L689 — `PORT void setRMatchFFAlpha(void* ptr, double ff_alpha)`
  Called by: `SetIVACFFAlpha()` (`ChannelMaster/ivac.c`)
- **`getControlFlag()`** — L699 — `PORT void getControlFlag(void* ptr, int* control_flag)`
  Called by: `GetIVACControlFlag()` (`ChannelMaster/ivac.c`)
- **`create_rmatchLegacyV()`** — L710 — `PORT void* create_rmatchLegacyV(int in_size, int out_size, int nom_inrate, int nom_outrate, int ringsize)`
  the following function is DEPRECATED it is intended for Legacy PowerSDR use only
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/rmatch.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
