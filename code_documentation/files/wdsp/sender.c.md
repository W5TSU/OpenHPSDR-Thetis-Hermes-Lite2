# `wdsp/sender.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sends DSP data (spectrum, audio taps) back toward the console.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/analyzer.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_sender()` (×1), `destroy_sender()` (×1), `flush_sender()` (×1), `xsender()` (×1), `setSamplerate_sender()` (×1), `setBuffers_sender()` (×1), `setSize_sender()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_sender()`** — L29 — `void calc_sender (SENDER a)`
  Called by: `create_sender()` (same file), `setSize_sender()` (same file)
- **`decalc_sender()`** — L34 — `void decalc_sender (SENDER a)`
  Called by: `destroy_sender()` (same file), `setSize_sender()` (same file)
- **`create_sender()`** — L39 — `SENDER create_sender (int run, int flag, int mode, int size, double* in, int arg0, int arg1, int arg2, int arg3)`
  Constructor for the `sender` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_sender()`** — L55 — `void destroy_sender (SENDER a)`
  Destroys the `sender` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_sender()`** — L61 — `void flush_sender (SENDER a)`
  Flushes (zeroes) the `sender` block’s internal buffers/state.
  Called by: `setSamplerate_sender()` (same file), `flush_rxa()` (`wdsp/RXA.c`)
- **`xsender()`** — L66 — `void xsender (SENDER a)`
  Runs the `sender` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_sender()`** — L88 — `void setBuffers_sender (SENDER a, double* in)`
  Re-points the `sender` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_sender()`** — L93 — `void setSamplerate_sender (SENDER a, int rate)`
  Reconfigures the `sender` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_sender()`** — L98 — `void setSize_sender (SENDER a, int size)`
  Reconfigures the `sender` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXASpectrum()`** — L111 — `PORT void SetRXASpectrum (int channel, int flag, int disp, int ss, int LO)`
  Sets rxaspectrum — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/sender.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
