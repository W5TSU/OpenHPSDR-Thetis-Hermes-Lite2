# `wdsp/compress.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** TX speech compressor and continuous frequency compressor.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/TXA.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_compressor()` (×1), `destroy_compressor()` (×1), `flush_compressor()` (×1), `xcompressor()` (×1), `setSamplerate_compressor()` (×1), `setBuffers_compressor()` (×1), `setSize_compressor()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_compressor()`** — L32 — `COMPRESSOR create_compressor ( int run, int buffsize, double* inbuff, double* outbuff,`
  Constructor for the `compressor` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_compressor()`** — L49 — `void destroy_compressor (COMPRESSOR a)`
  Destroys the `compressor` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_compressor()`** — L54 — `void flush_compressor (COMPRESSOR a)`
  Flushes (zeroes) the `compressor` block’s internal buffers/state.
  Called by: `flush_txa()` (`wdsp/TXA.c`)
- **`xcompressor()`** — L59 — `void xcompressor (COMPRESSOR a)`
  Runs the `compressor` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_compressor()`** — L77 — `void setBuffers_compressor (COMPRESSOR a, double* in, double* out)`
  Re-points the `compressor` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_compressor()`** — L83 — `void setSamplerate_compressor (COMPRESSOR a, int rate)`
  Reconfigures the `compressor` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_compressor()`** — L88 — `void setSize_compressor (COMPRESSOR a, int size)`
  Reconfigures the `compressor` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXACompressorRun()`** — L99 — `PORT void SetTXACompressorRun (int channel, int run)`
  Sets txacompressor run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACompressorGain()`** — L111 — `PORT void SetTXACompressorGain (int channel, double gain)`
  Sets txacompressor gain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/compress.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
