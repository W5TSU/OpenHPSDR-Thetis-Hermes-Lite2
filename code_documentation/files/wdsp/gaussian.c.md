# `wdsp/gaussian.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/apfshadow.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×13)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRXAGaussianFreqs()` (×3), `SetRXAGaussianGain()` (×2), `SetRXAGaussianRun()` (×2), `create_gaussian()` (×1), `destroy_gaussian()` (×1), `flush_gaussian()` (×1), `xgaussian()` (×1), `setSamplerate_gaussian()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_nc()`** — L30 — `static int calc_nc (double sigma, double nsigma, double rate)`
  Called by: `build_gaussian()` (same file)
- **`build_gaussian()`** — L45 — `double* build_gaussian(int* pnc, double rate, double f, double fwhm, double scale, double nsigma)`
  Called by: `create_gaussian()` (same file), `setSamplerate_gaussian()` (same file), `setSize_gaussian()` (same file), `setGain_gaussian()` (same file), `CalcGaussianFilter()` (same file), `SetRXAGaussianNC()` (same file)
- **`create_gaussian()`** — L105 — `GAUSSIAN create_gaussian (int run, int position, int size, int nc, double* in, double* out, double f_center, double bandwidth, int samplerate, double gain, double nsigma, int mode)`
  Constructor for the `gaussian` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_gaussian()`** — L132 — `void destroy_gaussian (GAUSSIAN a)`
  Destroys the `gaussian` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_gaussian()`** — L138 — `void flush_gaussian (GAUSSIAN a)`
  Flushes (zeroes) the `gaussian` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xgaussian()`** — L143 — `void xgaussian (GAUSSIAN a, int pos)`
  Runs the `gaussian` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_gaussian()`** — L164 — `void setBuffers_gaussian (GAUSSIAN a, double* in, double* out)`
  Re-points the `gaussian` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_gaussian()`** — L171 — `void setSamplerate_gaussian (GAUSSIAN a, int rate)`
  Reconfigures the `gaussian` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_gaussian()`** — L186 — `void setSize_gaussian (GAUSSIAN a, int size)`
  Reconfigures the `gaussian` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setGain_gaussian()`** — L198 — `void setGain_gaussian (GAUSSIAN a, double gain)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CalcGaussianFilter()`** — L208 — `void CalcGaussianFilter (GAUSSIAN a, double f_center, double bandwidth, double gain)`
  Called by: `SetRXAGaussianFreqs()` (same file), `SetRXAGaussianGain()` (same file)
- **`SetRXAGaussianRun()`** — L234 — `PORT void SetRXAGaussianRun (int channel, int run)`
  Sets rxagaussian run — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWRun()` (`wdsp/apfshadow.c`)
- **`SetRXAGaussianFreqs()`** — L243 — `PORT void SetRXAGaussianFreqs (int channel, double f_center, double bandwidth)`
  Sets rxagaussian freqs — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWFreq()` (`wdsp/apfshadow.c`), `SetRXASPCWBandwidth()` (`wdsp/apfshadow.c`)
- **`SetRXAGaussianGain()`** — L252 — `PORT void SetRXAGaussianGain (int channel, double gain)`
  Sets rxagaussian gain — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWGain()` (`wdsp/apfshadow.c`)
- **`SetRXAGaussianNC()`** — L261 — `PORT void SetRXAGaussianNC (int channel, int nc)`
  Sets rxagaussian nc — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/gaussian.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
