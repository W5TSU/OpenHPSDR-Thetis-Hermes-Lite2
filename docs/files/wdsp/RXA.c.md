# `wdsp/RXA.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Define the complete receive and transmit DSP graphs — every block below is instantiated and ordered here.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/main.c` (calls ×8)
  - `wdsp/snb.c` (calls ×4)
  - `wdsp/amd.c` (calls ×2)
  - `wdsp/anf.c` (calls ×2)
  - `wdsp/anr.c` (calls ×2)
  - `wdsp/emnr.c` (calls ×2)
  - `wdsp/rnnr.c` (calls ×2)
  - `wdsp/sbnr.c` (calls ×2)
  - `wdsp/nbp.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/snb.c` (calls ×19)
  - `wdsp/resample.c` (calls ×15)
  - `wdsp/iir.c` (calls ×14)
  - `wdsp/bandpass.c` (calls ×12)
  - `wdsp/nbp.c` (calls ×12)
  - `wdsp/fmd.c` (calls ×11)
  - `wdsp/shift.c` (calls ×11)
  - `wdsp/fmsq.c` (calls ×10)
  - `wdsp/eq.c` (calls ×9)
  - `wdsp/amsq.c` (calls ×8)
  - `wdsp/amd.c` (calls ×7)
  - `wdsp/anf.c` (calls ×7)
  - …and 20 more files
- Most-referenced symbols from other files: `RXAbp1Check()` (×7), `RXAbp1Set()` (×7), `RXAbpsnbaCheck()` (×2), `RXAbpsnbaSet()` (×2), `create_rxa()` (×1), `destroy_rxa()` (×1), `flush_rxa()` (×1), `xrxa()` (×1)

## Outline

### Functions

- `create_rxa()` — L31
- `destroy_rxa()` — L560
- `flush_rxa()` — L601
- `xrxa()` — L638
- `setInputSamplerate_rxa()` — L685
- `setOutputSamplerate_rxa()` — L701
- `setDSPSamplerate_rxa()` — L712
- `setDSPBuffsize_rxa()` — L763
- `SetRXAMode()` — L848
- `RXAResCheck()` — L892
- `RXAbp1Check()` — L903
- `RXAbp1Set()` — L921
- `RXAbpsnbaCheck()` — L937
- `RXAbpsnbaSet()` — L991
- `RXASetPassband()` — L1034
- `RXASetNC()` — L1042
- `RXASetMP()` — L1056

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/RXA.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
