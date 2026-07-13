# `wdsp/delay.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/eer.c` (calls ×17)
  - `wdsp/calcc.c` (calls ×8)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_delay()` (×5), `destroy_delay()` (×5), `SetDelayValue()` (×5), `SetDelayBuffs()` (×4), `flush_delay()` (×2), `xdelay()` (×2), `SetDelayRun()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_delay()`** — L29 — `DELAY create_delay (int run, int size, double* in, double* out, int rate, double tdelta, double tdelay)`
  Constructor for the `delay` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_calcc()` (`wdsp/calcc.c`), `SetPSFeedbackRate()` (`wdsp/calcc.c`), `create_eer()` (`wdsp/eer.c`), `SetEERSamplerate()` (`wdsp/eer.c`), `pSetEERSamplerate()` (`wdsp/eer.c`)
- **`destroy_delay()`** — L57 — `void destroy_delay (DELAY a)`
  Destroys the `delay` block, freeing its allocated buffers.
  Called by: `destroy_calcc()` (`wdsp/calcc.c`), `SetPSFeedbackRate()` (`wdsp/calcc.c`), `destroy_eer()` (`wdsp/eer.c`), `SetEERSamplerate()` (`wdsp/eer.c`), `pSetEERSamplerate()` (`wdsp/eer.c`)
- **`flush_delay()`** — L65 — `void flush_delay (DELAY a)`
  Flushes (zeroes) the `delay` block’s internal buffers/state.
  Called by: `flush_calcc()` (`wdsp/calcc.c`), `flush_eer()` (`wdsp/eer.c`)
- **`xdelay()`** — L71 — `void xdelay (DELAY a)`
  Runs the `delay` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `pscc()` (`wdsp/calcc.c`), `xeer()` (`wdsp/eer.c`)
- **`SetDelayRun()`** — L107 — `void SetDelayRun (DELAY a, int run)`
  Sets delay run — API setter, typically called from the console via P/Invoke.
  Called by: `SetEERRunDelays()` (`wdsp/eer.c`), `pSetEERRunDelays()` (`wdsp/eer.c`)
- **`SetDelayValue()`** — L114 — `double SetDelayValue (DELAY a, double tdelay)`
  Sets delay value — API setter, typically called from the console via P/Invoke.
  Called by: `SetPSTXDelay()` (`wdsp/calcc.c`), `SetEERMdelay()` (`wdsp/eer.c`), `SetEERPdelay()` (`wdsp/eer.c`), `pSetEERMdelay()` (`wdsp/eer.c`), `pSetEERPdelay()` (`wdsp/eer.c`)
- **`SetDelayBuffs()`** — L128 — `void SetDelayBuffs (DELAY a, int size, double* in, double* out)`
  Sets delay buffs — API setter, typically called from the console via P/Invoke.
  Called by: `pscc()` (`wdsp/calcc.c`), `SetEERSize()` (`wdsp/eer.c`), `pSetEERSize()` (`wdsp/eer.c`), `xeerEXTF()` (`wdsp/eer.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/delay.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
