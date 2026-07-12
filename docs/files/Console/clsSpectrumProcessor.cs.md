# `Console/clsSpectrumProcessor.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Pulls pixel-ready FFT data from the wdsp analyzer and post-processes it (averaging, peak/blend detection) for the display.

## How this file is used

- Used by (incoming references from other files):
  - `Console/ColorButton.cs` (calls ×2)
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/HPSDR/specHPSDR.cs` (calls ×9, references ×1)
  - `Console/Andromeda/Andromeda.cs` (imports ×1, references ×1)
  - `Console/enums.cs` (references ×1)
- Most-referenced symbols from other files: `.Refresh()` (×2)

## Outline

### Types

#### `clsSpectrumProcessor` (type, L51)

- `.AddReceiver()` — L101
- `.AddTransmitter()` — L111
- `.AddSource()` — L121
- `.RemoveReceiver()` — L148
- `.RemoveTransmitter()` — L153
- `.RemoveSource()` — L158
- `.Clear()` — L168
- `.ContainsSource()` — L183
- `.SetReceiverPixelResolution()` — L193
- `.SetTransmitterPixelResolution()` — L198
- `.SetPixelResolution()` — L203
- `.SetReceiverFrameRate()` — L212
- `.SetTransmitterFrameRate()` — L217
- `.SetFrameRate()` — L222
- `.SetReceiverFFTSize()` — L231
- `.SetTransmitterFFTSize()` — L236
- `.SetFFTSize()` — L241
- `.SetReceiverSampleRate()` — L250
- `.SetTransmitterSampleRate()` — L255
- `.SetSampleRate()` — L260
- `.SetReceiverZoomSlider()` — L269
- `.SetTransmitterZoomSlider()` — L274
- `.SetZoomSlider()` — L279
- `.SetReceiverPanSlider()` — L288
- `.SetTransmitterPanSlider()` — L293
- `.SetPanSlider()` — L298
- `.CentreReceiverPan()` — L307
- `.CentreTransmitterPan()` — L312
- `.CentrePan()` — L317
- `.SetReceiverEnabled()` — L322
- `.SetTransmitterEnabled()` — L327
- `.SetEnabled()` — L332
- `.ResetReceiverBuffers()` — L341
- `.ResetTransmitterBuffers()` — L346
- `.ResetBuffers()` — L351
- `.TryGetReceiverPixels()` — L360
- `.TryGetTransmitterPixels()` — L365
- `.TryCopyReceiverPixels()` — L370
- `.TryCopyTransmitterPixels()` — L375
- `.TryGetLatestPixels()` — L380
- `.TryCopyLatestPixels()` — L393
- `.TryGetViewport()` — L406
- `.TryGetSampleRate()` — L419
- `.TryGetFFTSize()` — L431
- `.ShowReceiverTestForm()` — L443
- `.ShowTransmitterTestForm()` — L448
- `.ShowTestForm()` — L453
- `.Dispose()` — L466
- `.RunWorker()` — L480
- `.SubscribeDelegates()` — L498
- `.UnsubscribeDelegates()` — L505
- `.OnCentreFrequencyChanged()` — L512
- `.OnHWSampleRateChanged()` — L525
- `.OnPowerChanged()` — L537
- `.GetEndpointArray()` — L546
- `.TryGetEndpoint()` — L551
- `.TryDetachEndpoint()` — L562
- `.UpdateEndpointSnapshotLocked()` — L579
- `.ThrowIfDisposed()` — L592
- `.MakeSourceKey()` — L597
- `.ValidateSourceId()` — L602
- `.ClampInt()` — L607
- `.ValidateSampleRate()` — L614
- `.ClampZoomSlider()` — L622
- `.ClampPanSlider()` — L630
- `.NormalizeFftSize()` — L638

#### `SpectrumSourceType` (type, L53)

_No extracted members._

#### `SpectrumEndpoint` (type, L657)

- `.MatchesReceiverEvent()` — L730
- `.SetPixels()` — L735
- `.SetFrameRate()` — L751
- `.SetFFTSize()` — L765
- `.SetSampleRate()` — L776
- `.SetZoomSlider()` — L788
- `.SetPanSlider()` — L803
- `.SetEnabled()` — L818
- `.SyncFrequencies()` — L830
- `.Refresh()` — L843
- `.ResetBuffers()` — L853
- `.ClearData()` — L866
- `.ProcessFrame()` — L876
- `.TryGetLatestPixels()` — L913
- `.TryCopyLatestPixels()` — L924
- `.TryGetViewport()` — L941
- `.TryGetSampleRate()` — L951
- `.TryGetFFTSize()` — L960
- `.Shutdown()` — L969
- `.RefreshLocked()` — L985
- `.ResolveDefaultBlockSize()` — L1014
- `.GetPixelOffset()` — L1020
- `.EnsurePixelBuffersLocked()` — L1033
- `.ClearPixelBuffersLocked()` — L1048
- `.RentPixelBuffer()` — L1055
- `.ReturnPixelBuffer()` — L1060
- `.FillPixelBuffer()` — L1068

#### `SpectrumTestForm` (type, L1075)

- `.BuildTitle()` — L1116
- `.RefreshTimer_Tick()` — L1128
- `.SpectrumTestForm_FormClosed()` — L1133

#### `SpectrumGraphPanel` (type, L1140)

- `.Dispose()` — L1175
- `.OnPaint()` — L1188
- `.TryCopyPixels()` — L1305
- `.TryCopyPixelsWithBuffer()` — L1321
- `.EnsurePointBuffer()` — L1328
- `.DrawGrid()` — L1334

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsSpectrumProcessor.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
