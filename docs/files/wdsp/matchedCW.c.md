# `wdsp/matchedCW.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Matched CW filtering and audio peaking filter support.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/apfshadow.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×12)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRXAMatchedFreqs()` (×3), `SetRXAMatchedGain()` (×2), `SetRXAMatchedRun()` (×2), `create_matched()` (×1), `destroy_matched()` (×1), `flush_matched()` (×1), `xmatched()` (×1), `setSamplerate_matched()` (×1)

## Outline

### Functions

- `calc_size()` — L30
- `build_matched()` — L44
- `create_matched()` — L123
- `destroy_matched()` — L145
- `flush_matched()` — L151
- `xmatched()` — L156
- `setBuffers_matched()` — L177
- `setSamplerate_matched()` — L184
- `setSize_matched()` — L197
- `setGain_matched()` — L209
- `CalcMatchedFilter()` — L219
- `SetRXAMatchedRun()` — L244
- `SetRXAMatchedFreqs()` — L253
- `SetRXAMatchedGain()` — L262

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/matchedCW.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
