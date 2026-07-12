# `wdsp/apfshadow.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Matched CW filtering and audio peaking filter support.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/iir.c` (calls ×8)
  - `wdsp/doublepole.c` (calls ×7)
  - `wdsp/gaussian.c` (calls ×7)
  - `wdsp/matchedCW.c` (calls ×7)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_apfshadow()` (×1), `destroy_apfshadow()` (×1)

## Outline

### Functions

- `create_apfshadow()` — L28
- `destroy_apfshadow()` — L39
- `SetRXASPCWSelection()` — L44
- `SetRXASPCWRun()` — L92
- `SetRXASPCWFreq()` — L116
- `SetRXASPCWBandwidth()` — L140
- `SetRXASPCWGain()` — L164

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/apfshadow.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
