# `Console/CAT/CATTester.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Serial receive event plumbing and an interactive CAT test window.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmMacroButtonConfig.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Invoke/textboxts.cs` (calls ×2, references ×1)
  - `Console/CAT/CATParser.cs` (references ×1, calls ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)

## Outline

### Types

#### `CATTester` (type, L13)

- `.Dispose()` — L47
- `.Setup()` — L59
- `.InitializeComponent()` — L82
- `.btnExit_Click()` — L179
- `.txtInput_KeyUp()` — L184
- `.ExecuteCommand()` — L192
- `.btnExecute_Click()` — L200
- `.CheckText()` — L205

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/CATTester.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
