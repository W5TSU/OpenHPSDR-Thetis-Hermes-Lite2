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

### Functions

- `create_emphp()` — L35
- `destroy_emphp()` — L56
- `flush_emphp()` — L62
- `xemphp()` — L67
- `setBuffers_emphp()` — L75
- `setSamplerate_emphp()` — L82
- `setSize_emphp()` — L91
- `SetTXAFMEmphPosition()` — L107
- `SetTXAFMEmphMP()` — L115
- `SetTXAFMEmphNC()` — L127
- `SetTXAFMPreEmphFreqs()` — L144
- `calc_emph()` — L168
- `decalc_emph()` — L177
- `create_emph()` — L186
- `destroy_emph()` — L202
- `flush_emph()` — L208
- `xemph()` — L213
- `setBuffers_emph()` — L235
- `setSamplerate_emph()` — L243
- `setSize_emph()` — L250

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/emph.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
