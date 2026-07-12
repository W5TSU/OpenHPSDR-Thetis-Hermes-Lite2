# `wdsp/fmmod.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM and FM modulators for TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×12)
  - `wdsp/fir.c` (calls ×6)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_fmmod()` (×1), `destroy_fmmod()` (×1), `flush_fmmod()` (×1), `xfmmod()` (×1), `setSamplerate_fmmod()` (×1), `setBuffers_fmmod()` (×1), `setSize_fmmod()` (×1), `SetTXAFMNC()` (×1)

## Outline

### Functions

- `calc_fmmod()` — L29
- `create_fmmod()` — L42
- `destroy_fmmod()` — L68
- `flush_fmmod()` — L74
- `xfmmod()` — L80
- `setBuffers_fmmod()` — L112
- `setSamplerate_fmmod()` — L120
- `setSize_fmmod()` — L130
- `SetTXAFMDeviation()` — L147
- `SetTXACTCSSFreq()` — L166
- `SetTXACTCSSRun()` — L178
- `SetTXAFMNC()` — L186
- `SetTXAFMMP()` — L203
- `SetTXAFMAFFreqs()` — L215

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fmmod.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
