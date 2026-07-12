# `Console/HPSDR/specHPSDR.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** Configures the wdsp spectrum analyzer instances for HPSDR data streams.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×26, references ×1)
  - `Console/clsSpectrumProcessor.cs` (calls ×9, references ×1)
  - `Console/MeterManager.cs` (calls ×5, references ×1)
  - `Console/wbDisplay.cs` (calls ×5)
  - `Console/PanDisplay.cs` (calls ×2)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetSpecRX()` (×22), `.resetPixelBuffers()` (×8), `.GetPixels()` (×5), `.initAnalyzer()` (×2), `.SetAnalyzer()` (×2), `.SnapSpectrum()` (×2), `.ZoomToBandwidth()` (×1), `.GetFrequencyExtents()` (×1)

## Outline

### Types

#### `SpecRX` (type, L29)

- `.GetSpecRX()` — L46

#### `SpecHPSDR` (type, L52)

- `.updateNormalizePan()` — L288
- `.resetPixelBuffers()` — L490
- `.initAnalyzer()` — L504
- `.ZoomToBandwidth()` — L645
- `.GetFrequencyExtents()` — L672
- `.CalcSpectrum()` — L738

#### `SpecHPSDRDLL` (type, L809)

- `.SetAnalyzer()` — L812
- `.XCreateAnalyzer()` — L835
- `.ResetPixelBuffers()` — L839
- `.DestroyAnalyzer()` — L842
- `.SetPixelRef()` — L845
- `.GetPixelsNative()` — L848
- `.GetPixels()` — L851
- `.Spectrum()` — L862
- `.SetCalibration()` — L865
- `.SnapSpectrum()` — L868
- `.SnapSpectrumTimeout()` — L871
- `.SetDisplayDetectorMode()` — L874
- `.SetDisplayAverageMode()` — L877
- `.SetDisplayNumAverage()` — L880
- `.SetDisplayAvBackmult()` — L883
- `.SetDisplayNormOneHz()` — L886
- `.GetDisplayENB()` — L889
- `.SetDisplaySampleRate()` — L892
- `.create_nobEXT()` — L895
- `.destroy_nobEXT()` — L909
- `.xnobEXTF()` — L912
- `.SetEXTNOBBuffsize()` — L915
- `.SetEXTNOBSamplerate()` — L918
- `.SetEXTNOBTau()` — L921
- `.SetEXTNOBHangtime()` — L924
- `.SetEXTNOBAdvtime()` — L927
- `.SetEXTNOBBacktau()` — L930
- `.SetEXTNOBThreshold()` — L933
- `.SetEXTNOBMode()` — L936
- `.create_anbEXT()` — L939
- `.destroy_anbEXT()` — L952
- `.xanbEXTF()` — L955
- `.SetEXTANBBuffsize()` — L958
- `.SetEXTANBSamplerate()` — L961
- `.SetEXTANBTau()` — L964
- `.SetEXTANBHangtime()` — L967
- `.SetEXTANBAdvtime()` — L970
- `.SetEXTANBBacktau()` — L973
- `.SetEXTANBThreshold()` — L976

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/specHPSDR.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
