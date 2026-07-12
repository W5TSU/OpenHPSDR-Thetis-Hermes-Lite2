# `Console/frmQuickRecallPopupList.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Quick recall (recent frequencies) list.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmQuickRecallPopupList.Designer.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `frmQuickRecallPopupList` (type, L47)

- `.AddItem()` — L63
- `.ClearItems()` — L67
- `.lstboxFrequencies_MouseClick()` — L74

#### `QuickRecallListBox` (type, L86)

- `.ClearItems()` — L102
- `.AddItem()` — L109
- `.OnDrawItem()` — L120
- `.OnFontChanged()` — L155
- `.OnSelectedIndexChanged()` — L165
- `.OnMouseEnter()` — L171
- `.OnMouseLeave()` — L175
- `.OnMouseMove()` — L182
- `.OnPaint()` — L199

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmQuickRecallPopupList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
