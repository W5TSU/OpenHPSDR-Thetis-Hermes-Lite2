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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `PrettyTrackBar` (type, L52)

- **`.UpdateHeadRectPos()`** — L85 — `private void UpdateHeadRectPos()`
  Updates head rect pos.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateLimitBar()`** — L127 — `private void UpdateLimitBar()`
  Updates limit bar.
  Called by: `.OnMouseDown()` (same file), `.OnMouseMove()` (same file)
- **`.ConstrainAValue()`** — L267 — `public int ConstrainAValue(int value)`
  Called by: `.SetPowerUsingTargetDBM()` (`Console/console.cs`)
- **`.OnEnabledChanged()`** — L330 — `protected override void OnEnabledChanged(EventArgs e)`
  Handles/raises the enabled changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseDown()`** — L336 — `protected override void OnMouseDown(MouseEventArgs e)`
  Handles/raises the mouse down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseMove()`** — L446 — `protected override void OnMouseMove(MouseEventArgs e)`
  Handles/raises the mouse move event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseUp()`** — L575 — `protected override void OnMouseUp(MouseEventArgs e)`
  Handles/raises the mouse up event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L594 — `protected override void OnPaint(PaintEventArgs pe)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnScroll()`** — L706 — `protected virtual void OnScroll(object sender, EventArgs e)`
  Handles/raises the scroll event.
  Called by: `.OnMouseDown()` (same file), `.OnMouseMove()` (same file), `.OnMouseWheel()` (same file)
- **`.OnMouseWheel()`** — L715 — `protected override void OnMouseWheel(MouseEventArgs e)`
  Handles/raises the mouse wheel event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `LimitConstraint` (type, L64)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/PrettyTrackBar.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
