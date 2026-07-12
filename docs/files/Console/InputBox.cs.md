# `Console/InputBox.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Simple text-input dialog; searchable "find a setting" helper.

## How this file is used

- Used by (incoming references from other files):
  - `Console/clsDBMan.cs` (calls ×8)
  - `Console/setup.cs` (calls ×5)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Show()` (×13)

## Outline

### Types

#### `InputBox` (type, L47)

- `.Dispose()` — L76
- `.InitializeComponent()` — L95
- `.Show()` — L158
- `.btnOK_Click()` — L176
- `.btnCancel_Click()` — L182

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/InputBox.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
