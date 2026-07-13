# `Console/ucParametricEq.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX/TX graphic and parametric equalizer forms (backed by wdsp `eq.c`).

## How this file is used

- Used by (incoming references from other files):
  - `Console/eqform.cs` (references ×1)
  - `Console/frmCFCConfig.Designer.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucParametricEq` (type, L52)

- **`.getYAxisStepDb()`** — L1007 — `private double getYAxisStepDb()`
  Returns yaxis step db.
  Called by: `.getAxisLabelMaxWidth()` (same file), `.drawGrid()` (same file), `.drawAxisScales()` (same file)
- **`.chooseDbStep()`** — L1013 — `private double chooseDbStep(double span)`
  Called by: `.getYAxisStepDb()` (same file)
- **`.raisePointSelected()`** — L1025 — `private void raisePointSelected(int index, EqPoint p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.raisePointUnselected()`** — L1033 — `private void raisePointUnselected(int index, EqPoint p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetPoints()`** — L1041 — `public void ResetPoints()`
  Resets points.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawBarChart()`** — L1048 — `public void DrawBarChart(double[] data)`
  Draws bar chart.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDefaults()`** — L1107 — `public void GetDefaults(out double[] F, out double[] G, out double[] Q, out double global_preamp_db, out double min_hz, out double max_hz, out bool parametric_eq, out int band_coun`
  Returns defaults.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPointHz()`** — L1134 — `public bool SetPointHz(int band_id, double frequency_hz, bool is_dragging = false)`
  Sets point hz.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetIndexFromBandId()`** — L1142 — `public int GetIndexFromBandId(int band_id)`
  Returns index from band id.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.findPointByBandId()`** — L1152 — `private EqPoint findPointByBandId(int band_id)`
  Called by: `.SetPointHz()` (same file)
- **`.setPointHzInternal()`** — L1162 — `private bool setPointHzInternal(EqPoint p, double frequency_hz, bool is_dragging)`
  Sets point hz internal.
  Called by: `.SetPointHz()` (same file)
- **`.GetPointData()`** — L1246 — `public void GetPointData(int index, out double frequency_hz, out double gain_db, out double q)`
  Returns point data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPointData()`** — L1260 — `public bool SetPointData(int index, double frequency_hz, double gain_db, double q)`
  Sets point data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPointsData()`** — L1290 — `public void GetPointsData(out double[] frequency_hz, out double[] gain_db, out double[] q)`
  Returns points data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPointsData()`** — L1311 — `public bool SetPointsData(double[] frequency_hz, double[] gain_db, double[] q)`
  Sets points data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SaveToJsonFromPoints()`** — L1353 — `public string SaveToJsonFromPoints(double[] F, double[] G, double[] Q, double global_gain_db, double frequency_min_hz, double frequency_max_hz, bool parametric_eq)`
  Saves to json from points.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointsFromJson()`** — L1392 — `public bool PointsFromJson(string json, out double[] F, out double[] G, out double[] Q, out double global_gain_db, out double frequency_min_hz, out double frequency_max_hz, out boo`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SaveToJson()`** — L1460 — `public string SaveToJson()`
  Saves to json.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LoadFromJson()`** — L1488 — `public bool LoadFromJson(string json)`
  Loads from json.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L1575 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L1611 — `protected override void Dispose(bool disposing)`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseDown()`** — L1625 — `protected override void OnMouseDown(MouseEventArgs e)`
  Handles/raises the mouse down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseMove()`** — L1677 — `protected override void OnMouseMove(MouseEventArgs e)`
  Handles/raises the mouse move event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseUp()`** — L1803 — `protected override void OnMouseUp(MouseEventArgs e)`
  Handles/raises the mouse up event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseWheel()`** — L1846 — `protected override void OnMouseWheel(MouseEventArgs e)`
  Handles/raises the mouse wheel event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDraggingNow()`** — L1997 — `private bool isDraggingNow()`
  Called by: `.OnMouseWheel()` (same file), `.enforceOrdering()` (same file)
- **`.getAxisLabelMaxWidth()`** — L2002 — `private int getAxisLabelMaxWidth()`
  Returns axis label max width.
  Called by: `.getComputedPlotMarginLeft()` (same file)
- **`.getComputedPlotMarginLeft()`** — L2015 — `private int getComputedPlotMarginLeft()`
  Returns computed plot margin left.
  Called by: `.getPlotRect()` (same file)
- **`.getComputedPlotMarginRight()`** — L2030 — `private int getComputedPlotMarginRight()`
  Returns computed plot margin right.
  Called by: `.getPlotRect()` (same file)
- **`.getPlotRect()`** — L2042 — `private Rectangle getPlotRect()`
  Returns plot rect.
  Called by: `.OnPaint()` (same file), `.OnMouseDown()` (same file), `.OnMouseMove()` (same file), `.OnMouseWheel()` (same file)
- **`.drawGrid()`** — L2061 — `private void drawGrid(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.drawBandShading()`** — L2192 — `private void drawBandShading(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.drawCurve()`** — L2344 — `private void drawCurve(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.drawGlobalGainHandle()`** — L2372 — `private void drawGlobalGainHandle(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.drawPoints()`** — L2395 — `private void drawPoints(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.drawDotReading()`** — L2428 — `private void drawDotReading(Graphics g, Rectangle plot, EqPoint p, float dot_x, float dot_y, float dot_radius)`
  Called by: `.drawPoints()` (same file)
- **`.createRoundedRectPath()`** — L2480 — `private GraphicsPath createRoundedRectPath(RectangleF rect, float radius)`
  Called by: `.drawDotReading()` (same file)
- **`.drawAxisScales()`** — L2505 — `private void drawAxisScales(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.formatDbTick()`** — L2601 — `private string formatDbTick(double db, double step_db)`
  Called by: `.getAxisLabelMaxWidth()` (same file), `.drawAxisScales()` (same file)
- **`.formatHzTick()`** — L2616 — `private string formatHzTick(double hz)`
  Called by: `.drawAxisScales()` (same file)
- **`.chooseFrequencyStep()`** — L2632 — `private double chooseFrequencyStep(double span)`
  Called by: `.OnMouseWheel()` (same file), `.drawAxisScales()` (same file)
- **`.drawBorder()`** — L2645 — `private void drawBorder(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.drawReadout()`** — L2653 — `private void drawReadout(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.responseDbAtFrequency()`** — L2694 — `private double responseDbAtFrequency(double frequency_hz)`
  Called by: `.drawCurve()` (same file)
- **`.barChartPeakTimer_Tick()`** — L2751 — `private void barChartPeakTimer_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `barChartPeakTimer` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.applyBarChartPeakDecay()`** — L2769 — `private void applyBarChartPeakDecay(DateTime now_utc)`
  Called by: `.DrawBarChart()` (same file), `.barChartPeakTimer_Tick()` (same file)
- **`.syncBarChartPeaksToData()`** — L2817 — `private void syncBarChartPeaksToData()`
  Called by: `.applyBarChartPeakDecay()` (same file)
- **`.updateBarChartPeakTimerState()`** — L2845 — `private void updateBarChartPeakTimerState()`
  Called by: `.DrawBarChart()` (same file), `.barChartPeakTimer_Tick()` (same file)
- **`.getBandBaseColor()`** — L2864 — `private Color getBandBaseColor(int index)`
  Returns band base color.
  Called by: `.drawBandShading()` (same file), `.getPointDisplayColor()` (same file), `.resetPointsDefault()` (same file)
- **`.getPointDisplayColor()`** — L2873 — `private Color getPointDisplayColor(int index)`
  Returns point display color.
  Called by: `.drawPoints()` (same file)
- **`.formatHz()`** — L2887 — `private string formatHz(double hz)`
  Called by: `.drawReadout()` (same file)
- **`.formatDotReadingHz()`** — L2893 — `private string formatDotReadingHz(double hz)`
  Called by: `.drawDotReading()` (same file)
- **`.formatDb()`** — L2898 — `private string formatDb(double db)`
  Called by: `.drawReadout()` (same file)
- **`.formatDotReadingDb()`** — L2904 — `private string formatDotReadingDb(double db)`
  Called by: `.drawDotReading()` (same file)
- **`.hitTestPoint()`** — L2910 — `private int hitTestPoint(Rectangle plot, Point pt)`
  Called by: `.OnMouseDown()` (same file), `.OnMouseMove()` (same file)
- **`.hitTestGlobalGainHandle()`** — L2936 — `private bool hitTestGlobalGainHandle(Rectangle plot, Point pt)`
  Called by: `.OnMouseDown()` (same file), `.OnMouseMove()` (same file), `.OnMouseWheel()` (same file)
- **`.xFromFreq()`** — L2951 — `private float xFromFreq(Rectangle plot, double frequency_hz)`
  Called by: `.DrawBarChart()` (same file), `.drawGrid()` (same file), `.drawBandShading()` (same file), `.drawCurve()` (same file), `.drawPoints()` (same file), `.drawAxisScales()` (same file) — and 2 more
- **`.freqFromX()`** — L2957 — `private double freqFromX(Rectangle plot, int x)`
  Called by: `.OnMouseMove()` (same file)
- **`.yFromDb()`** — L2965 — `private float yFromDb(Rectangle plot, double db)`
  Called by: `.DrawBarChart()` (same file), `.drawGrid()` (same file), `.drawBandShading()` (same file), `.drawCurve()` (same file), `.drawGlobalGainHandle()` (same file), `.drawPoints()` (same file) — and 3 more
- **`.dbFromY()`** — L2973 — `private double dbFromY(Rectangle plot, int y)`
  Called by: `.OnMouseMove()` (same file)
- **`.clamp()`** — L2983 — `private double clamp(double v, double min, double max)`
  Called by: `.DrawBarChart()` (same file), `.setPointHzInternal()` (same file), `.SetPointData()` (same file), `.SetPointsData()` (same file), `.SaveToJsonFromPoints()` (same file), `.PointsFromJson()` (same file) — and 11 more
- **`.getLogFrequencyCentreHz()`** — L2990 — `private double getLogFrequencyCentreHz()`
  Returns log frequency centre hz.
  Called by: `.getNormalizedFrequencyPosition()` (same file), `.frequencyFromNormalizedPosition()` (same file), `.getLogFrequencyTicks()` (same file)
- **`.getNormalizedFrequencyPosition()`** — L3002 — `private double getNormalizedFrequencyPosition(double frequency_hz)`
  Returns normalized frequency position.
  Called by: `.xFromFreq()` (same file)
- **`.frequencyFromNormalizedPosition()`** — L3035 — `private double frequencyFromNormalizedPosition(double t)`
  Called by: `.DrawBarChart()` (same file), `.drawCurve()` (same file), `.freqFromX()` (same file)
- **`.getLogFrequencyShape()`** — L3069 — `private double getLogFrequencyShape(double centre_ratio)`
  Returns log frequency shape.
  Called by: `.getNormalizedFrequencyPosition()` (same file), `.frequencyFromNormalizedPosition()` (same file)
- **`.getLogFrequencyTicks()`** — L3080 — `private List<double> getLogFrequencyTicks(Rectangle plot)`
  Returns log frequency ticks.
  Called by: `.drawGrid()` (same file), `.drawAxisScales()` (same file)
- **`.addLogFrequencyTickCandidate()`** — L3156 — `private void addLogFrequencyTickCandidate(List<double> ticks, double frequency_hz)`
  Called by: `.getLogFrequencyTicks()` (same file)
- **`.resetPointsDefault()`** — L3163 — `private void resetPointsDefault()`
  Called by: `.ResetPoints()` (same file), `.LoadFromJson()` (same file)
- **`.rescaleFrequencies()`** — L3199 — `private void rescaleFrequencies(double old_min, double old_max, double new_min, double new_max)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.enforceOrdering()`** — L3223 — `private void enforceOrdering(bool enforce_spacing_all)`
  Called by: `.setPointHzInternal()` (same file), `.SetPointData()` (same file), `.SetPointsData()` (same file), `.LoadFromJson()` (same file), `.OnMouseMove()` (same file), `.OnMouseWheel()` (same file) — and 1 more
- **`.clampAllGains()`** — L3314 — `private void clampAllGains()`
  Called by: `.resetPointsDefault()` (same file)
- **`.clampAllQ()`** — L3325 — `private void clampAllQ()`
  Called by: `.resetPointsDefault()` (same file)
- **`.raisePointsChanged()`** — L3334 — `private void raisePointsChanged(bool is_dragging)`
  Called by: `.ResetPoints()` (same file), `.setPointHzInternal()` (same file), `.SetPointData()` (same file), `.SetPointsData()` (same file), `.LoadFromJson()` (same file), `.OnMouseMove()` (same file) — and 2 more
- **`.raiseGlobalGainChanged()`** — L3340 — `private void raiseGlobalGainChanged(bool is_dragging)`
  Called by: `.LoadFromJson()` (same file), `.OnMouseUp()` (same file)
- **`.raiseSelectedIndexChanged()`** — L3346 — `private void raiseSelectedIndexChanged(bool is_dragging)`
  Called by: `.OnMouseUp()` (same file), `.enforceOrdering()` (same file)
- **`.raisePointDataChangedForPoint()`** — L3354 — `private void raisePointDataChangedForPoint(EqPoint p, bool is_dragging)`
  Called by: `.setPointHzInternal()` (same file), `.SetPointData()` (same file), `.OnMouseMove()` (same file), `.OnMouseUp()` (same file), `.OnMouseWheel()` (same file)
- **`.getComputedPlotMarginBottom()`** — L3365 — `private int getComputedPlotMarginBottom()`
  Returns computed plot margin bottom.
  Called by: `.getPlotRect()` (same file)
- **`.isFrequencyLockedIndex()`** — L3384 — `private bool isFrequencyLockedIndex(int index)`
  Called by: `.setPointHzInternal()` (same file), `.SetPointData()` (same file), `.SetPointsData()` (same file), `.LoadFromJson()` (same file), `.OnMouseMove()` (same file), `.OnMouseWheel()` (same file)
- **`.getLockedFrequencyForIndex()`** — L3389 — `private double getLockedFrequencyForIndex(int index)`
  Returns locked frequency for index.
  Called by: `.setPointHzInternal()` (same file), `.SetPointData()` (same file), `.SetPointsData()` (same file), `.LoadFromJson()` (same file)

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
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucParametricEq.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
