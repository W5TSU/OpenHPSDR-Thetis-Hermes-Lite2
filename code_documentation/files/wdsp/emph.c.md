# `wdsp/emph.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FM pre-/de-emphasis.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×11)
  - `wdsp/fcurve.c` (calls ×6)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_emphp()` (×1), `destroy_emphp()` (×1), `flush_emphp()` (×1), `xemphp()` (×1), `setSamplerate_emphp()` (×1), `setBuffers_emphp()` (×1), `setSize_emphp()` (×1), `SetTXAFMEmphNC()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_emphp()`** — L35 — `EMPHP create_emphp (int run, int position, int size, int nc, int mp, double* in, double* out, int rate, int ctype, double f_low, double f_high)`
  Constructor for the `emphp` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_emphp()`** — L56 — `void destroy_emphp (EMPHP a)`
  Destroys the `emphp` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_emphp()`** — L62 — `void flush_emphp (EMPHP a)`
  Flushes (zeroes) the `emphp` block’s internal buffers/state.
  Called by: `flush_txa()` (`wdsp/TXA.c`)
- **`xemphp()`** — L67 — `void xemphp (EMPHP a, int position)`
  Runs the `emphp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_emphp()`** — L75 — `void setBuffers_emphp (EMPHP a, double* in, double* out)`
  Re-points the `emphp` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_emphp()`** — L82 — `void setSamplerate_emphp (EMPHP a, int rate)`
  Reconfigures the `emphp` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_emphp()`** — L91 — `void setSize_emphp (EMPHP a, int size)`
  Reconfigures the `emphp` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXAFMEmphPosition()`** — L107 — `PORT void SetTXAFMEmphPosition (int channel, int position)`
  Sets txafmemph position — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAFMEmphMP()`** — L115 — `PORT void SetTXAFMEmphMP (int channel, int mp)`
  Sets txafmemph mp — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetMP()` (`wdsp/TXA.c`)
- **`SetTXAFMEmphNC()`** — L127 — `PORT void SetTXAFMEmphNC (int channel, int nc)`
  Sets txafmemph nc — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetNC()` (`wdsp/TXA.c`)
- **`SetTXAFMPreEmphFreqs()`** — L144 — `PORT void SetTXAFMPreEmphFreqs (int channel, double low, double high)`
  Sets txafmpre emph freqs — API setter, typically called from the console via P/Invoke.
  Called by: `SetTXAFMAFFilter()` (`wdsp/TXA.c`)
- **`calc_emph()`** — L168 — `void calc_emph (EMPH a)`
  Called by: `create_emph()` (same file), `setBuffers_emph()` (same file), `setSamplerate_emph()` (same file), `setSize_emph()` (same file)
- **`decalc_emph()`** — L177 — `void decalc_emph (EMPH a)`
  Called by: `destroy_emph()` (same file), `setBuffers_emph()` (same file), `setSamplerate_emph()` (same file), `setSize_emph()` (same file)
- **`create_emph()`** — L186 — `EMPH create_emph (int run, int position, int size, double* in, double* out, int rate, int ctype, double f_low, double f_high)`
  Constructor for the `emph` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_emph()`** — L202 — `void destroy_emph (EMPH a)`
  Destroys the `emph` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_emph()`** — L208 — `void flush_emph (EMPH a)`
  Flushes (zeroes) the `emph` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xemph()`** — L213 — `void xemph (EMPH a, int position)`
  Runs the `emph` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_emph()`** — L235 — `void setBuffers_emph (EMPH a, double* in, double* out)`
  Re-points the `emph` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_emph()`** — L243 — `void setSamplerate_emph (EMPH a, int rate)`
  Reconfigures the `emph` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_emph()`** — L250 — `void setSize_emph (EMPH a, int size)`
  Reconfigures the `emph` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/emph.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
