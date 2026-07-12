# `wdsp/amsq.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM squelch, FM squelch, and syllabic (voice-detecting) squelch.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×8)
  - `wdsp/TXA.c` (calls ×8)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_amsq()` (×2), `destroy_amsq()` (×2), `flush_amsq()` (×2), `xamsq()` (×2), `xamsqcap()` (×2), `setSamplerate_amsq()` (×2), `setBuffers_amsq()` (×2), `setSize_amsq()` (×2)

## Outline

### Functions

- `compute_slews()` — L29
- `calc_amsq()` — L49
- `decalc_amsq()` — L66
- `create_amsq()` — L73
- `destroy_amsq()` — L95
- `flush_amsq()` — L101
- `xamsq()` — L117
- `xamsqcap()` — L178
- `setBuffers_amsq()` — L183
- `setSamplerate_amsq()` — L190
- `setSize_amsq()` — L197
- `SetRXAAMSQRun()` — L210
- `SetRXAAMSQThreshold()` — L218
- `SetRXAAMSQMaxTail()` — L228
- `SetTXAAMSQRun()` — L245
- `SetTXAAMSQMutedGain()` — L253
- `SetTXAAMSQThreshold()` — L264

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/amsq.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
