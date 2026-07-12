# `wdsp/gaussian.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/apfshadow.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×13)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRXAGaussianFreqs()` (×3), `SetRXAGaussianGain()` (×2), `SetRXAGaussianRun()` (×2), `create_gaussian()` (×1), `destroy_gaussian()` (×1), `flush_gaussian()` (×1), `xgaussian()` (×1), `setSamplerate_gaussian()` (×1)

## Outline

### Functions

- `calc_nc()` — L30
- `build_gaussian()` — L45
- `create_gaussian()` — L105
- `destroy_gaussian()` — L132
- `flush_gaussian()` — L138
- `xgaussian()` — L143
- `setBuffers_gaussian()` — L164
- `setSamplerate_gaussian()` — L171
- `setSize_gaussian()` — L186
- `setGain_gaussian()` — L198
- `CalcGaussianFilter()` — L208
- `SetRXAGaussianRun()` — L234
- `SetRXAGaussianFreqs()` — L243
- `SetRXAGaussianGain()` — L252
- `SetRXAGaussianNC()` — L261

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/gaussian.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
