# `wdsp/iqc.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** PureSignal calibration calculation and the I/Q correction applied to TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/calcc.c` (calls ×10)
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetTXAiqcStart()` (×2), `SetTXAiqcSwap()` (×2), `create_iqc()` (×1), `destroy_iqc()` (×1), `flush_iqc()` (×1), `xiqc()` (×1), `setSamplerate_iqc()` (×1), `setBuffers_iqc()` (×1)

## Outline

### Functions

- `size_iqc()` — L29
- `desize_iqc()` — L46
- `calc_iqc()` — L59
- `decalc_iqc()` — L80
- `create_iqc()` — L87
- `destroy_iqc()` — L102
- `flush_iqc()` — L108
- `xiqc()` — L122
- `setBuffers_iqc()` — L205
- `setSamplerate_iqc()` — L211
- `setSize_iqc()` — L218
- `GetTXAiqcValues()` — L229
- `SetTXAiqcValues()` — L241
- `SetTXAiqcSwap()` — L255
- `SetTXAiqcStart()` — L271
- `SetTXAiqcEnd()` — L288
- `GetTXAiqcDogCount()` — L301
- `SetTXAiqcDogCount()` — L309

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/iqc.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
