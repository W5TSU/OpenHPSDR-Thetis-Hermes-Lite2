# `wdsp/slew.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/TXA.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_uslew()` (×1), `destroy_uslew()` (×1), `flush_uslew()` (×1), `xuslew()` (×1), `setSamplerate_uslew()` (×1), `setBuffers_uslew()` (×1), `setSize_uslew()` (×1)

## Outline

### Functions

- `calc_uslew()` — L37
- `decalc_uslew()` — L57
- `create_uslew()` — L62
- `destroy_uslew()` — L77
- `flush_uslew()` — L83
- `xuslew()` — L90
- `setBuffers_uslew()` — L157
- `setSamplerate_uslew()` — L163
- `setSize_uslew()` — L170
- `SetTXAuSlewTime()` — L182

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/slew.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
