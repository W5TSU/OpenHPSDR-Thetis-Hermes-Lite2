# `wdsp/anf.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Legacy LMS adaptive noise reduction (NR) and automatic notch filter.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_anf()` (×1), `destroy_anf()` (×1), `flush_anf()` (×1), `xanf()` (×1), `setSamplerate_anf()` (×1), `setBuffers_anf()` (×1), `setSize_anf()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_anf()`** — L29 — `ANF create_anf ( int run, int position, int buff_size, double *in_buff,`
  Constructor for the `anf` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_anf()`** — L77 — `void destroy_anf (ANF a)`
  Destroys the `anf` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`xanf()`** — L82 — `void xanf(ANF a, int position)`
  Runs the `anf` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`flush_anf()`** — L136 — `void flush_anf (ANF a)`
  Flushes (zeroes) the `anf` block’s internal buffers/state.
  Called by: `setSamplerate_anf()` (same file), `setSize_anf()` (same file), `SetRXAANFRun()` (same file), `SetRXAANFVals()` (same file), `SetRXAANFTaps()` (same file), `SetRXAANFDelay()` (same file) — and 4 more
- **`setBuffers_anf()`** — L143 — `void setBuffers_anf (ANF a, double* in, double* out)`
  Re-points the `anf` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_anf()`** — L149 — `void setSamplerate_anf (ANF a, int rate)`
  Reconfigures the `anf` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_anf()`** — L154 — `void setSize_anf (ANF a, int size)`
  Reconfigures the `anf` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAANFRun()`** — L166 — `PORT void SetRXAANFRun (int channel, int run)`
  Sets rxaanfrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANFVals()`** — L185 — `PORT void SetRXAANFVals (int channel, int taps, int delay, double gain, double leakage)`
  Sets rxaanfvals — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANFTaps()`** — L197 — `PORT void SetRXAANFTaps (int channel, int taps)`
  Sets rxaanftaps — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANFDelay()`** — L206 — `PORT void SetRXAANFDelay (int channel, int delay)`
  Sets rxaanfdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANFGain()`** — L215 — `PORT void SetRXAANFGain (int channel, double gain)`
  Sets rxaanfgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANFLeakage()`** — L224 — `PORT void SetRXAANFLeakage (int channel, double leakage)`
  Sets rxaanfleakage — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAANFPosition()`** — L233 — `PORT void SetRXAANFPosition (int channel, int position)`
  Sets rxaanfposition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/anf.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
