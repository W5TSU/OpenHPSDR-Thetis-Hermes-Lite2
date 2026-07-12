# `wdsp/cblock.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sample buffering between the audio callback world and DSP blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/ssql.c` (calls ×4)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_cbl()` (×2), `destroy_cbl()` (×2), `flush_cbl()` (×2), `xcbl()` (×2), `setSamplerate_cbl()` (×1), `setBuffers_cbl()` (×1), `setSize_cbl()` (×1)

## Outline

### Functions

- `calc_cbl()` — L29
- `create_cbl()` — L38
- `destroy_cbl()` — L63
- `flush_cbl()` — L68
- `xcbl()` — L76
- `setBuffers_cbl()` — L99
- `setSamplerate_cbl()` — L105
- `setSize_cbl()` — L111
- `SetRXACBLRun()` — L123
- `SetRXACBLPosition()` — L131

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cblock.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
