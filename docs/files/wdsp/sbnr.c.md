# `wdsp/sbnr.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** libspecbleach spectral noise reduction "NR4".

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×6)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_sbnr()` (×1), `destroy_sbnr()` (×1), `xsbnr()` (×1), `setSamplerate_sbnr()` (×1), `setBuffers_sbnr()` (×1), `setSize_sbnr()` (×1)

## Outline

### Functions

- `setSize_sbnr()` — L52
- `setBuffers_sbnr()` — L61
- `create_sbnr()` — L67
- `setSamplerate_sbnr()` — L90
- `xsbnr()` — L97
- `destroy_sbnr()` — L136
- `SetRXASBNRRun()` — L144
- `SetRXASBNRreductionAmount()` — L162
- `SetRXASBNRsmoothingFactor()` — L176
- `SetRXASBNRwhiteningFactor()` — L190
- `SetRXASBNRnoiseRescale()` — L205
- `SetRXASBNRpostFilterThreshold()` — L218
- `SetRXASBNRnoiseScalingType()` — L233
- `SetRXASBNRPosition()` — L243

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/sbnr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
