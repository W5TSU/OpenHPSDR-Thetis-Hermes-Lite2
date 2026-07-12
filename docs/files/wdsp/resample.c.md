# `wdsp/resample.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Fixed and variable-ratio resamplers, and the adaptive rate-matcher that reconciles independent sample clocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×15)
  - `wdsp/TXA.c` (calls ×15)
  - `ChannelMaster/aamix.c` (calls ×8)
  - `wdsp/snb.c` (calls ×6)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/fir.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `setBuffers_resample()` (×8), `create_resample()` (×6), `destroy_resample()` (×6), `setSize_resample()` (×6), `xresample()` (×4), `flush_resample()` (×4), `setInRate_resample()` (×4), `setOutRate_resample()` (×4)

## Outline

### Functions

- `calc_resample()` — L35
- `decalc_resample()` — L80
- `create_resample()` — L86
- `destroy_resample()` — L105
- `flush_resample()` — L112
- `xresample()` — L120
- `setBuffers_resample()` — L166
- `setSize_resample()` — L172
- `setInRate_resample()` — L178
- `setOutRate_resample()` — L185
- `setFCLow_resample()` — L192
- `setBandwidth_resample()` — L202
- `create_resampleV()` — L215
- `xresampleV()` — L221
- `destroy_resampleV()` — L231
- `create_resampleF()` — L243
- `destroy_resampleF()` — L289
- `flush_resampleF()` — L296
- `xresampleF()` — L303
- `create_resampleFV()` — L341
- `xresampleFV()` — L347
- `destroy_resampleFV()` — L357

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/resample.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
