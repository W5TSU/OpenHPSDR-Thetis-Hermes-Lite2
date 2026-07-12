# `wdsp/analyzer.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** The multi-instance FFT spectrum analyzer behind every panadapter/waterfall.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
  - `ChannelMaster/analyzers.c` (calls ×2)
  - `ChannelMaster/network.c` (calls ×1)
  - `ChannelMaster/pipe.c` (calls ×1)
  - `wdsp/sender.c` (calls ×1)
  - `wdsp/siphon.c` (calls ×1)
- Uses (outgoing references to other files):
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/combase.h` (calls ×3)
  - `wdsp/meterlog10.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `XCreateAnalyzer()` (×3), `DestroyAnalyzer()` (×3), `Spectrum0()` (×3), `Spectrum()` (×1), `Spectrum2()` (×1)

## Outline

### Functions

- `bessi0()` — L33
- `new_window()` — L52
- `eliminate()` — L180
- `Celiminate()` — L215
- `detector()` — L283
- `avenger()` — L464
- `stitch()` — L556
- `spectra()` — L612
- `Init_DetectMaxBin()` — L688
- `Destroy_DetectMaxBin()` — L707
- `calc_dmb()` — L714
- `SetupDetectMaxBin()` — L774
- `DetectMaxBin()` — L795
- `GetDetectMaxBin()` — L829
- `Cspectra()` — L846
- `interpolate()` — L928
- `build_interpolants()` — L982
- `sendbuf()` — L1066
- `CalcBandwidthNormalization()` — L1101
- `ResetPixelBuffers()` — L1108
- `SetAnalyzer()` — L1188
- `XCreateAnalyzer()` — L1339
- `DestroyAnalyzer()` — L1448
- `SetPixelRef()` — L1529
- `GetPixels()` — L1545
- `SnapSpectrum()` — L1588
- `SnapSpectrumTimeout()` — L1600
- `calcompare()` — L1621
- `SetCalibration()` — L1631
- `OpenBuffer()` — L1663
- `CloseBuffer()` — L1673
- `Spectrum()` — L1702
- `Spectrum2()` — L1741
- `Spectrum0()` — L1787
- `SetDisplayDetectorMode()` — L1833
- `SetDisplayAverageMode()` — L1845
- `SetDisplayNumAverage()` — L1877
- `SetDisplayAvBackmult()` — L1892
- `SetDisplaySampleRate()` — L1904
- `SetDisplayNormOneHz()` — L1917
- `GetDisplayENB()` — L1929

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/analyzer.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
