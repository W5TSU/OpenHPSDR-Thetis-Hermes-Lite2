# `wdsp/amd.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM/SAM (synchronous) and FM demodulators.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_amd()` (×1), `destroy_amd()` (×1), `flush_amd()` (×1), `xamd()` (×1), `setSamplerate_amd()` (×1), `setBuffers_amd()` (×1), `setSize_amd()` (×1)

## Outline

### Functions

- `create_amd()` — L29
- `destroy_amd()` — L67
- `init_amd()` — L72
- `flush_amd()` — L109
- `xamd()` — L115
- `setBuffers_amd()` — L241
- `setSamplerate_amd()` — L247
- `setSize_amd()` — L253
- `SetRXAAMDRun()` — L264
- `SetRXAAMDSBMode()` — L281
- `SetRXAAMDFadeLevel()` — L289

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/amd.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
