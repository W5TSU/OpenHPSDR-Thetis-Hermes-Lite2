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

### Functions

- `create_anr()` — L29
- `destroy_anr()` — L77
- `xanr()` — L82
- `flush_anr()` — L136
- `setBuffers_anr()` — L143
- `setSamplerate_anr()` — L149
- `setSize_anr()` — L154
- `SetRXAANRRun()` — L166
- `SetRXAANRVals()` — L184
- `SetRXAANRTaps()` — L196
- `SetRXAANRDelay()` — L205
- `SetRXAANRGain()` — L214
- `SetRXAANRLeakage()` — L223
- `SetRXAANRPosition()` — L232

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/anr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
