# `wdsp/varsamp.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Fixed and variable-ratio resamplers, and the adaptive rate-matcher that reconciles independent sample clocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/rmatch.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
- Most-referenced symbols from other files: `create_varsamp()` (×1), `destroy_varsamp()` (×1), `xvarsamp()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_varsamp()`** — L29 — `void calc_varsamp (VARSAMP a)`
  Called by: `create_varsamp()` (same file), `setInRate_varsamp()` (same file), `setOutRate_varsamp()` (same file), `setFCLow_varsamp()` (same file), `setBandwidth_varsamp()` (same file)
- **`decalc_varsamp()`** — L72 — `void decalc_varsamp (VARSAMP a)`
  Called by: `destroy_varsamp()` (same file), `setInRate_varsamp()` (same file), `setOutRate_varsamp()` (same file), `setFCLow_varsamp()` (same file), `setBandwidth_varsamp()` (same file)
- **`create_varsamp()`** — L79 — `VARSAMP create_varsamp ( int run, int size, double* in, double* out, int in_rate, int out_rate, double fc, double fc_low, int R, double gain, double var, int varmode)`
  Constructor for the `varsamp` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_varsampV()` (same file), `calc_rmatch()` (`wdsp/rmatch.c`)
- **`destroy_varsamp()`** — L100 — `void destroy_varsamp (VARSAMP a)`
  Destroys the `varsamp` block, freeing its allocated buffers.
  Called by: `destroy_varsampV()` (same file), `decalc_rmatch()` (`wdsp/rmatch.c`)
- **`flush_varsamp()`** — L106 — `void flush_varsamp (VARSAMP a)`
  Flushes (zeroes) the `varsamp` block’s internal buffers/state.
  Called by: `setSize_varsamp()` (same file)
- **`hshift()`** — L114 — `void hshift (VARSAMP a)`
  Called by: `xvarsamp()` (same file)
- **`xvarsamp()`** — L126 — `int xvarsamp (VARSAMP a, double var)`
  Runs the `varsamp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xvarsampV()` (same file), `xrmatchIN()` (`wdsp/rmatch.c`)
- **`setBuffers_varsamp()`** — L183 — `void setBuffers_varsamp (VARSAMP a, double* in, double* out)`
  Re-points the `varsamp` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_varsamp()`** — L189 — `void setSize_varsamp (VARSAMP a, int size)`
  Reconfigures the `varsamp` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setInRate_varsamp()`** — L195 — `void setInRate_varsamp (VARSAMP a, int rate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setOutRate_varsamp()`** — L202 — `void setOutRate_varsamp (VARSAMP a, int rate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setFCLow_varsamp()`** — L209 — `void setFCLow_varsamp (VARSAMP a, double fc_low)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBandwidth_varsamp()`** — L219 — `void setBandwidth_varsamp (VARSAMP a, double fc_low, double fc_high)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_varsampV()`** — L232 — `PORT void* create_varsampV (int in_rate, int out_rate, int R)`
  Constructor for the `varsampV` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xvarsampV()`** — L238 — `PORT void xvarsampV (double* input, double* output, int numsamps, double var, int* outsamps, void* ptr)`
  Runs the `varsampV` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_varsampV()`** — L248 — `PORT void destroy_varsampV (void* ptr)`
  Destroys the `varsampV` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/varsamp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
