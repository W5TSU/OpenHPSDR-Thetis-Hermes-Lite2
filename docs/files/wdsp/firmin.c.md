# `wdsp/firmin.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/eq.c` (calls ×22)
  - `wdsp/bandpass.c` (calls ×19)
  - `wdsp/nbp.c` (calls ×17)
  - `wdsp/fmd.c` (calls ×13)
  - `wdsp/gaussian.c` (calls ×13)
  - `wdsp/doublepole.c` (calls ×12)
  - `wdsp/fmmod.c` (calls ×12)
  - `wdsp/matchedCW.c` (calls ×12)
  - `wdsp/emph.c` (calls ×11)
  - `wdsp/fmsq.c` (calls ×7)
  - `wdsp/cfir.c` (calls ×4)
  - `wdsp/dexp.c` (calls ×4)
  - …and 3 more files
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×7)
  - `wdsp/fir.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `setImpulse_fircore()` (×49), `setNc_fircore()` (×17), `create_fircore()` (×14), `destroy_fircore()` (×14), `xfircore()` (×13), `flush_fircore()` (×12), `setBuffers_fircore()` (×10), `setMp_fircore()` (×10)

## Outline

### Functions

- `calc_firmin()` — L35
- `create_firmin()` — L44
- `destroy_firmin()` — L63
- `flush_firmin()` — L70
- `xfirmin()` — L76
- `setBuffers_firmin()` — L101
- `setSamplerate_firmin()` — L107
- `setSize_firmin()` — L113
- `setFreqs_firmin()` — L118
- `plan_firopt()` — L131
- `calc_firopt()` — L155
- `create_firopt()` — L172
- `deplan_firopt()` — L192
- `destroy_firopt()` — L212
- `flush_firopt()` — L218
- `xfiropt()` — L227
- `setBuffers_firopt()` — L253
- `setSamplerate_firopt()` — L262
- `setSize_firopt()` — L268
- `setFreqs_firopt()` — L276
- `plan_fircore()` — L290
- `calc_fircore()` — L322
- `create_fircore()` — L348
- `deplan_fircore()` — L365
- `destroy_fircore()` — L391
- `flush_fircore()` — L400
- `xfircore()` — L409
- `setBuffers_fircore()` — L441
- `setSize_fircore()` — L450
- `setImpulse_fircore()` — L458
- `setNc_fircore()` — L464
- `setMp_fircore()` — L478
- `setUpdate_fircore()` — L484

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/firmin.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
