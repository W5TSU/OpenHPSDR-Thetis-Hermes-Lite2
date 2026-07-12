# `Console/ucBandwidthView.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Variable-bandwidth adjustment popup and its graphical bandwidth view.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `ucBandwidthView` (type, L48)

- `.Reset()` — L188
- `.PushSample()` — L211
- `.OnPaint()` — L261
- `.drawGrid()` — L301
- `.drawAxisLeft()` — L327
- `.drawLine()` — L366
- `.drawOverlay()` — L410
- `.formatOverlayLine()` — L440
- `.formatAxisValue()` — L452
- `.toDisplayUnits()` — L464
- `.updateScale()` — L472
- `.resizeBuffers()` — L526
- `.resetSmoothing()` — L545

#### `BandwidthUnits` (type, L50)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucBandwidthView.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
