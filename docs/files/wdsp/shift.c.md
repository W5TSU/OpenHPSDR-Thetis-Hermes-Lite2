# `wdsp/shift.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×11)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `setBuffers_shift()` (×3), `setSize_shift()` (×3), `create_shift()` (×1), `destroy_shift()` (×1), `flush_shift()` (×1), `xshift()` (×1), `setSamplerate_shift()` (×1)

## Outline

### Functions

- `calc_shift()` — L29
- `create_shift()` — L36
- `destroy_shift()` — L50
- `flush_shift()` — L55
- `xshift()` — L60
- `setBuffers_shift()` — L87
- `setSamplerate_shift()` — L93
- `setSize_shift()` — L100
- `SetRXAShiftRun()` — L112
- `SetRXAShiftFreq()` — L120

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/shift.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
