# `wdsp/icfir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)

## Outline

### Functions

- `calc_icfir()` — L29
- `decalc_icfir()` — L38
- `create_icfir()` — L43
- `destroy_icfir()` — L79
- `flush_icfir()` — L85
- `xicfir()` — L90
- `setBuffers_icfir()` — L98
- `setSamplerate_icfir()` — L106
- `setSize_icfir()` — L113
- `setOutRate_icfir()` — L120
- `icfir_impulse()` — L127

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/icfir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
