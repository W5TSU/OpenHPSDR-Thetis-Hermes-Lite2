# `wdsp/cfcomp.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** TX speech compressor and continuous frequency compressor.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×4)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/meterlog10.c` (calls ×1)
- Most-referenced symbols from other files: `create_cfcomp()` (×1), `destroy_cfcomp()` (×1), `flush_cfcomp()` (×1), `xcfcomp()` (×1), `setSamplerate_cfcomp()` (×1), `setBuffers_cfcomp()` (×1), `setSize_cfcomp()` (×1)

## Outline

### Functions

- `calc_cfcwindow()` — L51
- `fCOMPcompare()` — L98
- `calc_comp()` — L108
- `calc_cfcomp()` — L358
- `decalc_cfcomp()` — L420
- `create_cfcomp()` — L449
- `flush_cfcomp()` — L482
- `destroy_cfcomp()` — L499
- `calc_mask()` — L511
- `xcfcomp()` — L554
- `setBuffers_cfcomp()` — L607
- `setSamplerate_cfcomp()` — L613
- `setSize_cfcomp()` — L620
- `SetTXACFCOMPRun()` — L633
- `SetTXACFCOMPPosition()` — L645
- `SetTXACFCOMPprofile()` — L657
- `SetTXACFCOMPPrecomp()` — L702
- `SetTXACFCOMPPeqRun()` — L719
- `SetTXACFCOMPPrePeq()` — L731
- `GetTXACFCOMPDisplayCompression()` — L741

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cfcomp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
