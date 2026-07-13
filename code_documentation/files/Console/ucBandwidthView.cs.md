# `Console/ucBandwidthView.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Variable-bandwidth adjustment popup and its graphical bandwidth view.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucBandwidthView` (type, L48)

- **`.Reset()`** — L188 — `public void Reset()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PushSample()`** — L211 — `public void PushSample(double inbound_bps, double outbound_bps)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L261 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.drawGrid()`** — L301 — `private void drawGrid(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.drawAxisLeft()`** — L327 — `private void drawAxisLeft(Graphics g, Rectangle plot, double max_display)`
  Called by: `.OnPaint()` (same file)
- **`.drawLine()`** — L366 — `private void drawLine(Graphics g, Rectangle plot, double[] buf, int count, int head, double max_display, Color color, float width)`
  Called by: `.OnPaint()` (same file)
- **`.drawOverlay()`** — L410 — `private void drawOverlay(Graphics g, Rectangle plot)`
  Called by: `.OnPaint()` (same file)
- **`.formatOverlayLine()`** — L440 — `private string formatOverlayLine(string prefix, double value_display, string unit)`
  Called by: `.drawOverlay()` (same file)
- **`.formatAxisValue()`** — L452 — `private string formatAxisValue(double value_display)`
  Called by: `.drawAxisLeft()` (same file)
- **`.toDisplayUnits()`** — L464 — `private double toDisplayUnits(double bytes_per_second)`
  Called by: `.drawLine()` (same file), `.drawOverlay()` (same file), `.updateScale()` (same file)
- **`.updateScale()`** — L472 — `private void updateScale()`
  Called by: `.Reset()` (same file), `.PushSample()` (same file)
- **`.resizeBuffers()`** — L526 — `private void resizeBuffers(int seconds)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.resetSmoothing()`** — L545 — `private void resetSmoothing()`
  Called by: `.Reset()` (same file), `.resizeBuffers()` (same file)

#### `BandwidthUnits` (type, L50)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucBandwidthView.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
