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

### Functions

- `create_compressor()` — L32
- `destroy_compressor()` — L49
- `flush_compressor()` — L54
- `xcompressor()` — L59
- `setBuffers_compressor()` — L77
- `setSamplerate_compressor()` — L83
- `setSize_compressor()` — L88
- `SetTXACompressorRun()` — L99
- `SetTXACompressorGain()` — L111

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/compress.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
