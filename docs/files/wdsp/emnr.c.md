# `wdsp/emnr.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Spectral noise reduction "NR2" (MMSE-based).

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/FDnoiseIQ.h` (imports ×1)
  - `wdsp/calculus.h` (imports ×1)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/zetaHat.h` (imports ×1)
  - `wdsp/meterlog10.c` (calls ×1)
- Most-referenced symbols from other files: `create_emnr()` (×1), `destroy_emnr()` (×1), `flush_emnr()` (×1), `xemnr()` (×1), `setSamplerate_emnr()` (×1), `setBuffers_emnr()` (×1), `setSize_emnr()` (×1)

## Outline

### Functions

- `bessI0()` — L44
- `bessI1()` — L82
- `e1xb()` — L128
- `calc_window()` — L168
- `interpM()` — L189
- `readZetaHat()` — L207
- `CwriteZetaHat()` — L247
- `calc_emnr()` — L300
- `decalc_emnr()` — L576
- `create_emnr()` — L637
- `flush_emnr()` — L659
- `destroy_emnr()` — L674
- `LambdaD()` — L680
- `LambdaDs()` — L805
- `LambdaDl()` — L820
- `aepf()` — L849
- `post2_calc_w()` — L898
- `post2()` — L909
- `SetRXAEMNRpost2Run()` — L951
- `SetRXAEMNRpost2Factor()` — L959
- `SetRXAEMNRpost2Nlevel()` — L967
- `SetRXAEMNRpost2Taper()` — L975
- `SetRXAEMNRpost2Rate()` — L985
- `getKey()` — L1001
- `getZeta()` — L1047
- `calc_gain()` — L1069
- `xemnr()` — L1200
- `setBuffers_emnr()` — L1256
- `setSamplerate_emnr()` — L1262
- `setSize_emnr()` — L1269
- `SetRXAEMNRRun()` — L1282
- `SetRXAEMNRgainMethod()` — L1299
- `SetRXAEMNRnpeMethod()` — L1307
- `SetRXAEMNRaeRun()` — L1315
- `SetRXAEMNRPosition()` — L1323
- `SetRXAEMNRaeZetaThresh()` — L1335
- `SetRXAEMNRaePsi()` — L1343
- `SetRXAEMNRtrainZetaThresh()` — L1351
- `SetRXAEMNRtrainT2()` — L1359

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/emnr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
