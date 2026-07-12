# `Console/PrettyTrackBar.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Custom-drawn slider and color-picker button used across forms.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Skin.cs` (references ×3)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/console.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.ConstrainAValue()` (×1)

## Outline

### Types

#### `PrettyTrackBar` (type, L52)

- `.UpdateHeadRectPos()` — L85
- `.UpdateLimitBar()` — L127
- `.ConstrainAValue()` — L267
- `.OnEnabledChanged()` — L330
- `.OnMouseDown()` — L336
- `.OnMouseMove()` — L446
- `.OnMouseUp()` — L575
- `.OnPaint()` — L594
- `.OnScroll()` — L706
- `.OnMouseWheel()` — L715

#### `LimitConstraint` (type, L64)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/PrettyTrackBar.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
