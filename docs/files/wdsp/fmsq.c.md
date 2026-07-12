# `wdsp/fmsq.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM squelch, FM squelch, and syllabic (voice-detecting) squelch.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×7)
  - `wdsp/eq.c` (calls ×2)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `setBuffers_fmsq()` (×2), `create_fmsq()` (×1), `destroy_fmsq()` (×1), `flush_fmsq()` (×1), `xfmsq()` (×1), `setSamplerate_fmsq()` (×1), `setSize_fmsq()` (×1), `SetRXAFMSQNC()` (×1)

## Outline

### Functions

- `calc_fmsq()` — L29
- `decalc_fmsq()` — L80
- `create_fmsq()` — L88
- `destroy_fmsq()` — L116
- `flush_fmsq()` — L122
- `xfmsq()` — L141
- `setBuffers_fmsq()` — L207
- `setSamplerate_fmsq()` — L215
- `setSize_fmsq()` — L222
- `SetRXAFMSQRun()` — L235
- `SetRXAFMSQThreshold()` — L243
- `SetRXAFMSQNC()` — L252
- `SetRXAFMSQMP()` — L269

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fmsq.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
