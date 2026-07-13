# `wdsp/matchedCW.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Matched CW filtering and audio peaking filter support.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/apfshadow.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×12)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRXAMatchedFreqs()` (×3), `SetRXAMatchedGain()` (×2), `SetRXAMatchedRun()` (×2), `create_matched()` (×1), `destroy_matched()` (×1), `flush_matched()` (×1), `xmatched()` (×1), `setSamplerate_matched()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_size()`** — L30 — `static int calc_size (int nc)`
  Called by: `build_matched()` (same file)
- **`build_matched()`** — L44 — `double* build_matched(int* imp_size, double rate, double f, double fwhm, double scale, int imp_pos)`
  Called by: `create_matched()` (same file), `setSamplerate_matched()` (same file), `setSize_matched()` (same file), `setGain_matched()` (same file), `CalcMatchedFilter()` (same file)
- **`create_matched()`** — L123 — `MATCHED create_matched (int run, int position, int size, double* in, double* out, double f_center, double bandwidth, int samplerate, double gain, int mode)`
  Constructor for the `matched` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_matched()`** — L145 — `void destroy_matched (MATCHED a)`
  Destroys the `matched` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_matched()`** — L151 — `void flush_matched (MATCHED a)`
  Flushes (zeroes) the `matched` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xmatched()`** — L156 — `void xmatched (MATCHED a, int pos)`
  Runs the `matched` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_matched()`** — L177 — `void setBuffers_matched (MATCHED a, double* in, double* out)`
  Re-points the `matched` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_matched()`** — L184 — `void setSamplerate_matched (MATCHED a, int rate)`
  Reconfigures the `matched` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_matched()`** — L197 — `void setSize_matched (MATCHED a, int size)`
  Reconfigures the `matched` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setGain_matched()`** — L209 — `void setGain_matched (MATCHED a, double gain)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CalcMatchedFilter()`** — L219 — `void CalcMatchedFilter (MATCHED a, double f_center, double bandwidth, double gain)`
  Called by: `SetRXAMatchedFreqs()` (same file), `SetRXAMatchedGain()` (same file)
- **`SetRXAMatchedRun()`** — L244 — `PORT void SetRXAMatchedRun (int channel, int run)`
  Sets rxamatched run — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWRun()` (`wdsp/apfshadow.c`)
- **`SetRXAMatchedFreqs()`** — L253 — `PORT void SetRXAMatchedFreqs (int channel, double f_center, double bandwidth)`
  Sets rxamatched freqs — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWFreq()` (`wdsp/apfshadow.c`), `SetRXASPCWBandwidth()` (`wdsp/apfshadow.c`)
- **`SetRXAMatchedGain()`** — L262 — `PORT void SetRXAMatchedGain (int channel, double gain)`
  Sets rxamatched gain — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWGain()` (`wdsp/apfshadow.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/matchedCW.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
