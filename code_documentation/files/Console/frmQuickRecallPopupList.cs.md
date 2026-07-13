# `Console/frmQuickRecallPopupList.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Quick recall (recent frequencies) list.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmQuickRecallPopupList.Designer.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmQuickRecallPopupList` (type, L47)

- **`.AddItem()`** — L63 — `public int AddItem(double dFreq)`
  Adds item.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearItems()`** — L67 — `public void ClearItems()`
  Clears items.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lstboxFrequencies_MouseClick()`** — L74 — `private void lstboxFrequencies_MouseClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `QuickRecallListBox` (type, L86)

- **`.ClearItems()`** — L102 — `public void ClearItems()`
  Clears items.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddItem()`** — L109 — `public int AddItem(double dFreq)`
  Adds item.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDrawItem()`** — L120 — `protected override void OnDrawItem(DrawItemEventArgs e)`
  Handles/raises the draw item event.
  Called by: `.OnPaint()` (same file)
- **`.OnFontChanged()`** — L155 — `protected override void OnFontChanged(EventArgs e)`
  Handles/raises the font changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSelectedIndexChanged()`** — L165 — `protected override void OnSelectedIndexChanged(EventArgs e)`
  Handles/raises the selected index changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseEnter()`** — L171 — `protected override void OnMouseEnter(EventArgs e)`
  Handles/raises the mouse enter event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseLeave()`** — L175 — `protected override void OnMouseLeave(EventArgs e)`
  Handles/raises the mouse leave event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseMove()`** — L182 — `protected override void OnMouseMove(MouseEventArgs e)`
  Handles/raises the mouse move event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L199 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmQuickRecallPopupList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
