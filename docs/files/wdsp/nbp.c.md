# `wdsp/nbp.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Main bandpass filter and the notched-bandpass (auto/manual notch database) filter.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×12)
  - `wdsp/snb.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×17)
  - `wdsp/utilities.c` (calls ×4)
  - `wdsp/snb.c` (calls ×3)
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/fir.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_nbp()` (×2), `destroy_nbp()` (×2), `flush_nbp()` (×2), `xnbp()` (×2), `create_notchdb()` (×1), `destroy_notchdb()` (×1), `setSamplerate_nbp()` (×1), `setBuffers_nbp()` (×1)

## Outline

### Functions

- `create_notchdb()` — L35
- `destroy_notchdb()` — L49
- `fir_mbandpass()` — L64
- `min_notch_width()` — L82
- `make_nbp()` — L100
- `calc_nbp_lightweight()` — L184
- `calc_nbp_impulse()` — L217
- `create_nbp()` — L244
- `destroy_nbp()` — L273
- `flush_nbp()` — L281
- `xnbp()` — L286
- `setBuffers_nbp()` — L294
- `setSamplerate_nbp()` — L301
- `setSize_nbp()` — L309
- `setNc_nbp()` — L319
- `setMp_nbp()` — L326
- `UpdateNBPFiltersLightWeight()` — L339
- `UpdateNBPFilters()` — L345
- `RXANBPAddNotch()` — L361
- `RXANBPGetNotch()` — L392
- `RXANBPDeleteNotch()` — L417
- `RXANBPEditNotch()` — L443
- `RXANBPGetNumNotches()` — L464
- `RXANBPSetTuneFrequency()` — L474
- `RXANBPSetShiftFrequency()` — L486
- `RXANBPSetNotchesRun()` — L498
- `RXANBPSetRun()` — L520
- `RXANBPSetFreqs()` — L530
- `RXANBPSetWindow()` — L545
- `RXANBPSetNC()` — L566
- `RXANBPSetMP()` — L581
- `RXANBPGetMinNotchWidth()` — L593
- `RXANBPSetAutoIncrease()` — L603

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/nbp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
