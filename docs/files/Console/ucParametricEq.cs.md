# `Console/ucParametricEq.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX/TX graphic and parametric equalizer forms (backed by wdsp `eq.c`).

## How this file is used

- Used by (incoming references from other files):
  - `Console/eqform.cs` (references ×1)
  - `Console/frmCFCConfig.Designer.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `ucParametricEq` (type, L52)

- `.getYAxisStepDb()` — L1007
- `.chooseDbStep()` — L1013
- `.raisePointSelected()` — L1025
- `.raisePointUnselected()` — L1033
- `.ResetPoints()` — L1041
- `.DrawBarChart()` — L1048
- `.GetDefaults()` — L1107
- `.SetPointHz()` — L1134
- `.GetIndexFromBandId()` — L1142
- `.findPointByBandId()` — L1152
- `.setPointHzInternal()` — L1162
- `.GetPointData()` — L1246
- `.SetPointData()` — L1260
- `.GetPointsData()` — L1290
- `.SetPointsData()` — L1311
- `.SaveToJsonFromPoints()` — L1353
- `.PointsFromJson()` — L1392
- `.SaveToJson()` — L1460
- `.LoadFromJson()` — L1488
- `.OnPaint()` — L1575
- `.Dispose()` — L1611
- `.OnMouseDown()` — L1625
- `.OnMouseMove()` — L1677
- `.OnMouseUp()` — L1803
- `.OnMouseWheel()` — L1846
- `.isDraggingNow()` — L1997
- `.getAxisLabelMaxWidth()` — L2002
- `.getComputedPlotMarginLeft()` — L2015
- `.getComputedPlotMarginRight()` — L2030
- `.getPlotRect()` — L2042
- `.drawGrid()` — L2061
- `.drawBandShading()` — L2192
- `.drawCurve()` — L2344
- `.drawGlobalGainHandle()` — L2372
- `.drawPoints()` — L2395
- `.drawDotReading()` — L2428
- `.createRoundedRectPath()` — L2480
- `.drawAxisScales()` — L2505
- `.formatDbTick()` — L2601
- `.formatHzTick()` — L2616
- `.chooseFrequencyStep()` — L2632
- `.drawBorder()` — L2645
- `.drawReadout()` — L2653
- `.responseDbAtFrequency()` — L2694
- `.barChartPeakTimer_Tick()` — L2751
- `.applyBarChartPeakDecay()` — L2769
- `.syncBarChartPeaksToData()` — L2817
- `.updateBarChartPeakTimerState()` — L2845
- `.getBandBaseColor()` — L2864
- `.getPointDisplayColor()` — L2873
- `.formatHz()` — L2887
- `.formatDotReadingHz()` — L2893
- `.formatDb()` — L2898
- `.formatDotReadingDb()` — L2904
- `.hitTestPoint()` — L2910
- `.hitTestGlobalGainHandle()` — L2936
- `.xFromFreq()` — L2951
- `.freqFromX()` — L2957
- `.yFromDb()` — L2965
- `.dbFromY()` — L2973
- `.clamp()` — L2983
- `.getLogFrequencyCentreHz()` — L2990
- `.getNormalizedFrequencyPosition()` — L3002
- `.frequencyFromNormalizedPosition()` — L3035
- `.getLogFrequencyShape()` — L3069
- `.getLogFrequencyTicks()` — L3080
- `.addLogFrequencyTickCandidate()` — L3156
- `.resetPointsDefault()` — L3163
- `.rescaleFrequencies()` — L3199
- `.enforceOrdering()` — L3223
- `.clampAllGains()` — L3314
- `.clampAllQ()` — L3325
- `.raisePointsChanged()` — L3334
- `.raiseGlobalGainChanged()` — L3340
- `.raiseSelectedIndexChanged()` — L3346
- `.raisePointDataChangedForPoint()` — L3354
- `.getComputedPlotMarginBottom()` — L3365
- `.isFrequencyLockedIndex()` — L3384
- `.getLockedFrequencyForIndex()` — L3389

#### `EqPoint` (type, L54)

_No extracted members._

#### `EqDraggingEventArgs` (type, L107)

_No extracted members._

#### `EqPointDataChangedEventArgs` (type, L122)

_No extracted members._

#### `EqPointSelectionChangedEventArgs` (type, L177)

_No extracted members._

#### `EqJsonState` (type, L220)

_No extracted members._

#### `EqJsonPoint` (type, L242)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucParametricEq.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
