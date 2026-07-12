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

### Functions

- `calc_varsamp()` — L29
- `decalc_varsamp()` — L72
- `create_varsamp()` — L79
- `destroy_varsamp()` — L100
- `flush_varsamp()` — L106
- `hshift()` — L114
- `xvarsamp()` — L126
- `setBuffers_varsamp()` — L183
- `setSize_varsamp()` — L189
- `setInRate_varsamp()` — L195
- `setOutRate_varsamp()` — L202
- `setFCLow_varsamp()` — L209
- `setBandwidth_varsamp()` — L219
- `create_varsampV()` — L232
- `xvarsampV()` — L238
- `destroy_varsampV()` — L248

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/varsamp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
