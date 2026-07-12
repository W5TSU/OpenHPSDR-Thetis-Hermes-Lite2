# `Console/display.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** The spectrum/waterfall renderer — SharpDX (Direct2D/D3D11) drawing of panadapter, waterfall, band edges, notches, cursors, and TX filter overlays.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×11)
  - `Console/setup.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×14)
  - `Console/radio.cs` (references ×1, calls ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Channel.cs` (references ×1)
  - `Console/SpotManager2.cs` (calls ×1)
  - `Console/clsBandStackManager.cs` (references ×1)
  - `Console/hiperftimer.cs` (references ×1)
- Most-referenced symbols from other files: `.DX2Adaptors()` (×2), `.PurgeBuffers()` (×1), `.RenderDX2D()` (×1), `.SetPendingWaterfallPixelRef()` (×1), `.ResetWaterfallTimers()` (×1), `.ShutdownDX2D()` (×1), `.SetDX2BackgoundImage()` (×1), `.SetupDelegates()` (×1)

## Outline

### Functions

- `letItSnow()` — L11750
- `plotSanta()` — L11789
- `SetSantaGif()` — L11852
- `santaCleanUp()` — L11885

### Types

#### `WaterfallTimePosition` (type, L82)

_No extracted members._

#### `WaterfallTimeMode` (type, L91)

_No extracted members._

#### `Display` (type, L99)

- `.SetupDelegates()` — L767
- `.RemoveDelegates()` — L786
- `.OnMinRXNotchWidthChanged()` — L800
- `.OnMinTXNotchWidthChanged()` — L805
- `.OnCTUNChanged()` — L811
- `.OnPowerChangeHander()` — L818
- `.OnBandChangeHandler()` — L825
- `.processBlobsActivePeakDisplayDelay()` — L847
- `.delayBlobsActivePeakDisplay()` — L859
- `.resetPeaksAndNoise()` — L879
- `.OnAttenuatorDataChanged()` — L893
- `.OnPreampModeChanged()` — L900
- `.OnCentreFrequencyChanged()` — L907
- `.localMox()` — L2731
- `.getWaterfallRxIndex()` — L2784
- `.setCurrentWaterfallBand()` — L2789
- `.getWaterfallContextBandValue()` — L2796
- `.getWaterfallCachedPreviousMinOrFloor()` — L2810
- `.updateWaterfallAgcCache()` — L2822
- `.clearBuffers()` — L2835
- `.initDisplayArrays()` — L2878
- `.changeAlpha()` — L2938
- `.dBToPixel()` — L2943
- `.dBToRX2Pixel()` — L2955
- `.updateSharePointsArray()` — L2980
- `.ResetWaterfallBmp()` — L3073
- `.ResetWaterfallBmp2()` — L3131
- `.ShutdownDX2D()` — L3230
- `.DX2Adaptors()` — L3326
- `.getGPUNameInUse()` — L3386
- `.initDX2D()` — L3408
- `.DXVersion()` — L3615
- `.ResetDX2DModeDescription()` — L3641
- `.resizeDX2D()` — L3668
- `.setupAliasing()` — L3781
- `.pauseDisplay()` — L3831
- `.RenderDX2D()` — L3882
- `.showFPSProfile()` — L4303
- `.calcFps()` — L4365
- `.isOccupied()` — L4435
- `.processMaximums()` — L4462
- `.ResetSpectrumPeaks()` — L4527
- `.ResetBlobMaximums()` — L4542
- `.getFilterXPositions()` — L4564
- `.isRxDuplex()` — L4618
- `.modifyDataForNotches()` — L4733
- `.DrawPanadapterDX2D()` — L4976
- `.findImd()` — L5731
- `.processNoiseFloor()` — L5872
- `.resetWaterfallTimeOverlay()` — L5922
- `.resizeWaterfallTimeOverlay()` — L5936
- `.getWaterfallLineIntervalMs()` — L5976
- `.recordWaterfallAdvance()` — L5989
- `.toDisplayTicks()` — L6060
- `.chooseWaterfallLabelIntervalMs()` — L6068
- `.drawWaterfallTimeOverlay()` — L6091
- `.SetPendingWaterfallPixelRef()` — L6167
- `.resetWaterfallBitmapAlignment()` — L6174
- `.getWaterfallSpanHz()` — L6184
- `.isWaterfallNoiseFloorCompensationEnabled()` — L6189
- `.useWaterfallNoiseFloorCompensation()` — L6196
- `.getWaterfallNoiseFloorCompensationTarget()` — L6204
- `.prepareWaterfallBitmapShift()` — L6212
- `.clearWaterfallBitmapRegion()` — L6265
- `.ResetWaterfallTimers()` — L6288
- `.OnWaterfallRXGradientChanged()` — L6308
- `.OnWaterfallTXGradientChanged()` — L6340
- `.DrawWaterfallDX2D()` — L6357
- `.convertColour()` — L7717
- `.convertBrush()` — L7721
- `.SetDX2BackgoundImage()` — L7726
- `.SDXBitmapFromSysBitmap()` — L7759
- `.buildLinearGradientBrush()` — L7945
- `.buildLinearGradientBrushTX()` — L8032
- `.releaseDX2Resources()` — L8111
- `.buildDX2Resources()` — L8270
- `.releaseFonts()` — L8374
- `.buildFontsDX2D()` — L8404
- `.clearBackgroundDX2D()` — L8427
- `.drawLineDX2D()` — L8457
- `.drawFillRectangleDX2D()` — L8467
- `.drawRectangleDX2D()` — L8473
- `.drawElipseDX2D()` — L8479
- `.drawFillElipseDX2D()` — L8485
- `.drawStringDX2D()` — L8503
- `.drawFilterOverlayDX2D()` — L8509
- `.drawChannelBarDX2D()` — L8520
- `.measureStringDX2D()` — L8542
- `.getCWSideToneShift()` — L8601
- `.handleNotches()` — L8644
- `.fastPow10Shifted()` — L8772
- `.fastPow10Raw()` — L8794
- `.drawPanadapterAndWaterfallGridDX2D()` — L8802
- `.AppendCursorInfoLines()` — L10014
- `.DrawCursorInfoPanel()` — L10027
- `.DrawCursorInfo()` — L10072
- `.DrawSpectrumDX2D()` — L10088
- `.DrawSpectrumGridDX2D()` — L10233
- `.DrawScopeDX2D()` — L10613
- `.DrawScope2DX2D()` — L10702
- `.lerp()` — L10784
- `.lerpPointF()` — L10788
- `.DrawPhaseDX2D()` — L10796
- `.DrawPhaseGridDX2D()` — L10894
- `.DrawPhase2DX2D()` — L10906
- `.DrawHistogramDX2D()` — L10963
- `.getSpotLayer()` — L11159
- `.updateLayer()` — L11192
- `.getCallsignString()` — L11212
- `.drawSpots()` — L11230
- `.clearAllDynamicBrushes()` — L11553
- `.getDXBrushForColour()` — L11570
- `.PurgeBuffers()` — L11630
- `.EnumDisplaySettings()` — L11935
- `.GetCurrentMonitorRefreshRate()` — L11938
- `.getSpotTagSize()` — L11952
- `.drawSpotTagContent()` — L11967
- `.getSpotFlagBitmap()` — L12000
- `.getSpotFlagRenderSize()` — L12039
- `.clearSpotFlagBitmapCache()` — L12051

#### `BandEdgeRegionCacheDX2D` (type, L103)

- `.Update()` — L153
- `.Contains()` — L203

#### `WaterfallAgcCacheKey` (type, L2752)

- `.Equals()` — L2763
- `.GetHashCode()` — L2773

#### `WaterfallAgcCacheEntry` (type, L2779)

_No extracted members._

#### `AdaptorInfo` (type, L3313)

_No extracted members._

#### `Maximums` (type, L4427)

_No extracted members._

#### `clsNotchCoords` (type, L8627)

_No extracted members._

#### `SnowFlake` (type, L11662)

- `.Update()` — L11685

#### `DEVMODE` (type, L11903)

_No extracted members._

#### `ReferenceEqualityComparer` (type, L12085)

- `.Equals()` — L12089
- `.GetHashCode()` — L12094

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/display.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
