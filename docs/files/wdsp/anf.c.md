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

### Functions

- `create_anf()` — L29
- `destroy_anf()` — L77
- `xanf()` — L82
- `flush_anf()` — L136
- `setBuffers_anf()` — L143
- `setSamplerate_anf()` — L149
- `setSize_anf()` — L154
- `SetRXAANFRun()` — L166
- `SetRXAANFVals()` — L185
- `SetRXAANFTaps()` — L197
- `SetRXAANFDelay()` — L206
- `SetRXAANFGain()` — L215
- `SetRXAANFLeakage()` — L224
- `SetRXAANFPosition()` — L233

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/anf.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
