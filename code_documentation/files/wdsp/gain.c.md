# `wdsp/gain.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_gain()`** — L29 — `PORT GAIN create_gain (int run, int* prun, int size, double* in, double* out, double Igain, double Qgain)`
  Constructor for the `gain` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_gain()`** — L44 — `PORT void destroy_gain (GAIN a)`
  Destroys the `gain` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_gain()`** — L51 — `PORT void flush_gain (GAIN a)`
  Flushes (zeroes) the `gain` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xgain()`** — L57 — `PORT void xgain (GAIN a)`
  Runs the `gain` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_gain()`** — L80 — `void setBuffers_gain (GAIN a, double* in, double* out)`
  Re-points the `gain` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_gain()`** — L86 — `void setSamplerate_gain (GAIN a, int rate)`
  Reconfigures the `gain` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_gain()`** — L91 — `void setSize_gain (GAIN a, int size)`
  Reconfigures the `gain` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`pSetTXOutputLevel()`** — L102 — `PORT void pSetTXOutputLevel (GAIN a, double level)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`pSetTXOutputLevelRun()`** — L111 — `PORT void pSetTXOutputLevelRun (GAIN a, int run)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`pSetTXOutputLevelSize()`** — L119 — `PORT void pSetTXOutputLevelSize (GAIN a, int size)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/gain.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
