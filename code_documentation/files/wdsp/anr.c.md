# `wdsp/anr.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Legacy LMS adaptive noise reduction (NR) and automatic notch filter.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_anr()` (×1), `destroy_anr()` (×1), `flush_anr()` (×1), `xanr()` (×1), `setSamplerate_anr()` (×1), `setBuffers_anr()` (×1), `setSize_anr()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_anr()`** — L29 — `ANR create_anr ( int run, int position, int buff_size, double *in_buff,`
  Constructor for the `anr` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_anr()`** — L77 — `void destroy_anr (ANR a)`
  Destroys the `anr` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`xanr()`** — L82 — `void xanr (ANR a, int position)`
  Runs the `anr` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`flush_anr()`** — L136 — `void flush_anr (ANR a)`
  Flushes (zeroes) the `anr` block’s internal buffers/state.
  Called by: `setSamplerate_anr()` (same file), `setSize_anr()` (same file), `SetRXAANRRun()` (same file), `SetRXAANRVals()` (same file), `SetRXAANRTaps()` (same file), `SetRXAANRDelay()` (same file) — and 4 more
- **`setBuffers_anr()`** — L143 — `void setBuffers_anr (ANR a, double* in, double* out)`
  Re-points the `anr` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_anr()`** — L149 — `void setSamplerate_anr (ANR a, int rate)`
  Reconfigures the `anr` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_anr()`** — L154 — `void setSize_anr (ANR a, int size)`
  Reconfigures the `anr` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAANRRun()`** — L166 — `PORT void SetRXAANRRun (int channel, int run)`
  Sets rxaanrrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANRVals()`** — L184 — `PORT void SetRXAANRVals (int channel, int taps, int delay, double gain, double leakage)`
  Sets rxaanrvals — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANRTaps()`** — L196 — `PORT void SetRXAANRTaps (int channel, int taps)`
  Sets rxaanrtaps — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANRDelay()`** — L205 — `PORT void SetRXAANRDelay (int channel, int delay)`
  Sets rxaanrdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANRGain()`** — L214 — `PORT void SetRXAANRGain (int channel, double gain)`
  Sets rxaanrgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANRLeakage()`** — L223 — `PORT void SetRXAANRLeakage (int channel, double leakage)`
  Sets rxaanrleakage — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANRPosition()`** — L232 — `PORT void SetRXAANRPosition (int channel, int position)`
  Sets rxaanrposition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/anr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
