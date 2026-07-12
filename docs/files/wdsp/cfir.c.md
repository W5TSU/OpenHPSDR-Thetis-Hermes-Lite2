# `wdsp/cfir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×9)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
- Most-referenced symbols from other files: `create_cfir()` (×1), `destroy_cfir()` (×1), `flush_cfir()` (×1), `xcfir()` (×1), `setOutRate_cfir()` (×1), `setSamplerate_cfir()` (×1), `setBuffers_cfir()` (×1), `setSize_cfir()` (×1)

## Outline

### Functions

- `calc_cfir()` — L29
- `decalc_cfir()` — L38
- `create_cfir()` — L43
- `destroy_cfir()` — L79
- `flush_cfir()` — L85
- `xcfir()` — L90
- `setBuffers_cfir()` — L98
- `setSamplerate_cfir()` — L106
- `setSize_cfir()` — L113
- `setOutRate_cfir()` — L120
- `cfir_impulse()` — L127
- `SetTXACFIRRun()` — L232
- `SetTXACFIRNC()` — L240

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cfir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
