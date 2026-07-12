# `wdsp/doublepole.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** IIR biquad sections (notches, peaking filters) and double-pole building blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/apfshadow.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×12)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/cmath.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
- Most-referenced symbols from other files: `SetRXADoublepoleFreqs()` (×3), `SetRXADoublepoleGain()` (×2), `SetRXADoublepoleRun()` (×2), `create_doublepole()` (×1), `destroy_doublepole()` (×1), `flush_doublepole()` (×1), `xdoublepole()` (×1), `setSamplerate_doublepole()` (×1)

## Outline

### Functions

- `calc_dpole_nc()` — L30
- `build_doublepole_1sided()` — L56
- `build_doublepole_2sided()` — L106
- `create_doublepole()` — L146
- `destroy_doublepole()` — L168
- `flush_doublepole()` — L174
- `xdoublepole()` — L179
- `setBuffers_doublepole()` — L200
- `setSamplerate_doublepole()` — L207
- `setSize_doublepole()` — L220
- `setGain_doublepole()` — L232
- `CalcDoublepoleFilter()` — L242
- `SetRXADoublepoleRun()` — L267
- `SetRXADoublepoleFreqs()` — L276
- `SetRXADoublepoleGain()` — L285

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/doublepole.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
