# `wdsp/fmd.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM/SAM (synchronous) and FM demodulators.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×11)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×13)
  - `wdsp/wcpAGC.c` (calls ×7)
  - `wdsp/iir.c` (calls ×6)
  - `wdsp/fcurve.c` (calls ×5)
  - `wdsp/fir.c` (calls ×5)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_fmd()` (×1), `destroy_fmd()` (×1), `flush_fmd()` (×1), `xfmd()` (×1), `setSamplerate_fmd()` (×1), `setBuffers_fmd()` (×1), `setSize_fmd()` (×1), `SetRXAFMNCaud()` (×1)

## Outline

### Functions

- `calc_fmd()` — L29
- `decalc_fmd()` — L75
- `create_fmd()` — L81
- `destroy_fmd()` — L122
- `flush_fmd()` — L131
- `xfmd()` — L144
- `setBuffers_fmd()` — L190
- `setSamplerate_fmd()` — L201
- `setSize_fmd()` — L218
- `SetRXAFMDeviation()` — L245
- `SetRXACTCSSFreq()` — L256
- `SetRXACTCSSRun()` — L267
- `SetRXAFMNCde()` — L278
- `SetRXAFMMPde()` — L295
- `SetRXAFMNCaud()` — L307
- `SetRXAFMMPaud()` — L324
- `SetRXAFMLimRun()` — L336
- `SetRXAFMLimGain()` — L349
- `SetRXAFMAFFilter()` — L364

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fmd.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
