# `wdsp/ammod.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM and FM modulators for TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_ammod()` (×1), `destroy_ammod()` (×1), `flush_ammod()` (×1), `xammod()` (×1), `setSamplerate_ammod()` (×1), `setBuffers_ammod()` (×1), `setSize_ammod()` (×1)

## Outline

### Functions

- `create_ammod()` — L29
- `destroy_ammod()` — L43
- `flush_ammod()` — L48
- `xammod()` — L53
- `setBuffers_ammod()` — L81
- `setSamplerate_ammod()` — L87
- `setSize_ammod()` — L92
- `SetTXAAMCarrierLevel()` — L103

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/ammod.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
