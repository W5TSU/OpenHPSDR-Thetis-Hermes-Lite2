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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ColorButton` (type, L58)

- **`.OnChanged()`** — L94 — `protected virtual void OnChanged(EventArgs e)`
  Handles/raises the changed event.
  Called by: `.OnClick()` (same file)
- **`.Dispose()`** — L105 — `protected override void Dispose(bool disposing)`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L117 — `private void InitializeComponent()`
  Designer-generated UI construction (creates and lays out the form’s controls).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L123 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseDown()`** — L168 — `protected override void OnMouseDown(MouseEventArgs e)`
  Handles/raises the mouse down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseUp()`** — L174 — `protected override void OnMouseUp(MouseEventArgs e)`
  Handles/raises the mouse up event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnClick()`** — L180 — `protected override void OnClick(EventArgs e)`
  Handles/raises the click event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `ColorPanel` (type, L191)

- **`.OnClosed()`** — L254 — `protected override void OnClosed(EventArgs e)`
  Handles/raises the closed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L262 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  Called by: `.OnPaint()` (same file)
- **`.OnKeyDown()`** — L316 — `protected override void OnKeyDown(KeyEventArgs e)`
  Handles/raises the key down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MoveIndex()`** — L335 — `private void MoveIndex(int delta)`
  Called by: `.OnKeyDown()` (same file)
- **`.OnMouseDown()`** — L367 — `protected override void OnMouseDown(MouseEventArgs e)`
  Handles/raises the mouse down event.
  Called by: `.OnMouseDown()` (same file)
- **`.OnMouseMove()`** — L375 — `protected override void OnMouseMove(MouseEventArgs e)`
  Handles/raises the mouse move event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnClick()`** — L420 — `protected override void OnClick(EventArgs e)`
  Handles/raises the click event.
  Called by: `.OnKeyDown()` (same file)
- **`.DrawButton()`** — L448 — `protected void DrawButton(PaintEventArgs e, int x, int y, string text, int index, bool selected)`
  Draws button.
  Called by: `.OnPaint()` (same file)
- **`.SetColorIndex()`** — L479 — `protected bool SetColorIndex(Rectangle rc, Point pt, int index)`
  Sets color index.
  Called by: `.OnMouseMove()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ColorButton.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
