# `Console/ColorButton.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Custom-drawn slider and color-picker button used across forms.

## How this file is used

- Used by (incoming references from other files):
  - `Console/rxaControls.Designer.cs` (references ×1)
  - `Console/setup.designer.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/clsSpectrumProcessor.cs` (calls ×2)
  - `Console/Invoke/buttonts.cs` (inherits ×1)

## Outline

### Types

#### `ColorButton` (type, L58)

- `.OnChanged()` — L94
- `.Dispose()` — L105
- `.InitializeComponent()` — L117
- `.OnPaint()` — L123
- `.OnMouseDown()` — L168
- `.OnMouseUp()` — L174
- `.OnClick()` — L180

#### `ColorPanel` (type, L191)

- `.OnClosed()` — L254
- `.OnPaint()` — L262
- `.OnKeyDown()` — L316
- `.MoveIndex()` — L335
- `.OnMouseDown()` — L367
- `.OnMouseMove()` — L375
- `.OnClick()` — L420
- `.DrawButton()` — L448
- `.SetColorIndex()` — L479

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ColorButton.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
