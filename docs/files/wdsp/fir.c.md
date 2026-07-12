# `wdsp/fir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/bandpass.c` (calls ×12)
  - `wdsp/fmmod.c` (calls ×6)
  - `wdsp/fmd.c` (calls ×5)
  - `wdsp/eq.c` (calls ×3)
  - `wdsp/fcurve.c` (calls ×3)
  - `wdsp/firmin.c` (calls ×3)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/nbp.c` (calls ×2)
  - `wdsp/resample.c` (calls ×2)
  - `wdsp/cfir.c` (calls ×1)
  - `wdsp/delay.c` (calls ×1)
  - `wdsp/dexp.c` (calls ×1)
  - …and 3 more files
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×9)
  - `wdsp/impulse_cache.c` (calls ×4)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `fir_bandpass()` (×33), `fir_fsamp()` (×4), `fftcv_mults()` (×3), `fir_fsamp_odd()` (×2), `mp_imp()` (×2), `analytic()` (×1)

## Outline

### Functions

- `fftcv_mults()` — L29
- `get_fsamp_window()` — L44
- `fir_fsamp_odd()` — L83
- `fir_fsamp()` — L127
- `fir_bandpass()` — L187
- `fir_read()` — L288
- `analytic()` — L330
- `mp_imp()` — L357
- `zff_impulse()` — L442

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
