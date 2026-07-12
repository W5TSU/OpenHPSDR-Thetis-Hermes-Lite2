# `wdsp/cfir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×9)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
- Most-referenced symbols from other files: `create_cfir()` (×1), `destroy_cfir()` (×1), `flush_cfir()` (×1), `xcfir()` (×1), `setOutRate_cfir()` (×1), `setSamplerate_cfir()` (×1), `setBuffers_cfir()` (×1), `setSize_cfir()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_cfir()`** — L29 — `void calc_cfir (CFIR a)`
  Called by: `create_cfir()` (same file), `setBuffers_cfir()` (same file), `setSamplerate_cfir()` (same file), `setSize_cfir()` (same file), `setOutRate_cfir()` (same file), `SetTXACFIRNC()` (same file)
- **`decalc_cfir()`** — L38 — `void decalc_cfir (CFIR a)`
  Called by: `destroy_cfir()` (same file), `setBuffers_cfir()` (same file), `setSamplerate_cfir()` (same file), `setSize_cfir()` (same file), `setOutRate_cfir()` (same file), `SetTXACFIRNC()` (same file)
- **`create_cfir()`** — L43 — `CFIR create_cfir (int run, int size, int nc, int mp, double* in, double* out, int runrate, int cicrate, int DD, int R, int Pairs, double cutoff, int xtype, double xbw, int wintype)`
  Constructor for the `cfir` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_cfir()`** — L79 — `void destroy_cfir (CFIR a)`
  Destroys the `cfir` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_cfir()`** — L85 — `void flush_cfir (CFIR a)`
  Flushes (zeroes) the `cfir` block’s internal buffers/state.
  Called by: `flush_txa()` (`wdsp/TXA.c`)
- **`xcfir()`** — L90 — `void xcfir (CFIR a)`
  Runs the `cfir` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_cfir()`** — L98 — `void setBuffers_cfir (CFIR a, double* in, double* out)`
  Re-points the `cfir` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_cfir()`** — L106 — `void setSamplerate_cfir (CFIR a, int rate)`
  Reconfigures the `cfir` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_cfir()`** — L113 — `void setSize_cfir (CFIR a, int size)`
  Reconfigures the `cfir` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setOutRate_cfir()`** — L120 — `void setOutRate_cfir (CFIR a, int rate)`
  Called by: `setOutputSamplerate_txa()` (`wdsp/TXA.c`)
- **`cfir_impulse()`** — L127 — `double* cfir_impulse (int N, int DD, int R, int Pairs, double runrate, double cicrate, double cutoff, int xtype, double xbw, int rtype, double scale, int wintype)`
  Called by: `calc_cfir()` (same file)
- **`SetTXACFIRRun()`** — L232 — `PORT void SetTXACFIRRun (int channel, int run)`
  Sets txacfirrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACFIRNC()`** — L240 — `PORT void SetTXACFIRNC(int channel, int nc)`
  Sets txacfirnc — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetNC()` (`wdsp/TXA.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cfir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
