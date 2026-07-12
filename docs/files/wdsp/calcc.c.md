# `wdsp/calcc.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** PureSignal calibration calculation and the I/Q correction applied to TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×2)
  - `ChannelMaster/sync.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/iqc.c` (calls ×10)
  - `wdsp/delay.c` (calls ×8)
  - `wdsp/lmath.c` (calls ×3)
  - `wdsp/utilities.c` (calls ×2)
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/combase.h` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `pscc()` (×1), `create_calcc()` (×1), `destroy_calcc()` (×1)

## Outline

### Functions

- `size_calcc()` — L30
- `desize_calcc()` — L83
- `create_calcc()` — L119
- `destroy_calcc()` — L205
- `flush_calcc()` — L241
- `scheck()` — L247
- `rxscheck()` — L294
- `calc()` — L324
- `doPSCalcCorrection()` — L485
- `doPSTurnoff()` — L509
- `PSSaveCorrection()` — L539
- `PSRestoreCorrection()` — L572
- `pscc()` — L616
- `psccF()` — L839
- `PSSaveCorr()` — L859
- `PSRestoreCorr()` — L871
- `SetPSRunCal()` — L890
- `SetPSMox()` — L900
- `GetPSInfo()` — L913
- `SetPSReset()` — L923
- `SetPSMancal()` — L933
- `SetPSAutomode()` — L941
- `SetPSTurnon()` — L949
- `SetPSControl()` — L957
- `SetPSLoopDelay()` — L970
- `SetPSMoxDelay()` — L981
- `SetPSTXDelay()` — L992
- `SetPSHWPeak()` — L1015
- `GetPSHWPeak()` — L1025
- `GetPSMaxTX()` — L1033
- `SetPSPtol()` — L1041
- `GetPSDisp()` — L1049
- `SetPSFeedbackRate()` — L1064
- `SetPSPinMode()` — L1093
- `SetPSMapMode()` — L1101
- `SetPSStabilize()` — L1109
- `ForceShutDown()` — L1117
- `SetPSIntsAndSpi()` — L1131

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/calcc.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
