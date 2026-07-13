# `wdsp/icfir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_icfir()`** — L29 — `void calc_icfir (ICFIR a)`
  Called by: `create_icfir()` (same file), `setBuffers_icfir()` (same file), `setSamplerate_icfir()` (same file), `setSize_icfir()` (same file), `setOutRate_icfir()` (same file)
- **`decalc_icfir()`** — L38 — `void decalc_icfir (ICFIR a)`
  Called by: `destroy_icfir()` (same file), `setBuffers_icfir()` (same file), `setSamplerate_icfir()` (same file), `setSize_icfir()` (same file), `setOutRate_icfir()` (same file)
- **`create_icfir()`** — L43 — `ICFIR create_icfir (int run, int size, int nc, int mp, double* in, double* out, int runrate, int cicrate, int DD, int R, int Pairs, double cutoff, int xtype, double xbw, int wintyp`
  Constructor for the `icfir` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_icfir()`** — L79 — `void destroy_icfir (ICFIR a)`
  Destroys the `icfir` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_icfir()`** — L85 — `void flush_icfir (ICFIR a)`
  Flushes (zeroes) the `icfir` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xicfir()`** — L90 — `void xicfir (ICFIR a)`
  Runs the `icfir` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_icfir()`** — L98 — `void setBuffers_icfir (ICFIR a, double* in, double* out)`
  Re-points the `icfir` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_icfir()`** — L106 — `void setSamplerate_icfir (ICFIR a, int rate)`
  Reconfigures the `icfir` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_icfir()`** — L113 — `void setSize_icfir (ICFIR a, int size)`
  Reconfigures the `icfir` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setOutRate_icfir()`** — L120 — `void setOutRate_icfir (ICFIR a, int rate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`icfir_impulse()`** — L127 — `double* icfir_impulse (int N, int DD, int R, int Pairs, double runrate, double cicrate, double cutoff, int xtype, double xbw, int rtype, double scale, int wintype)`
  Called by: `calc_icfir()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/icfir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
