# `Console/Andromeda/AndromedaEditForm.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Editor for assigning functions to Andromeda encoders and buttons.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/Andromeda.cs` (calls ×2, references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.SetEncoderNumber()` (×1), `.SetPushbuttonNumber()` (×1)

## Outline

### Types

#### `AndromedaEditForm` (type, L36)

- `.Dispose()` — L77
- `.InitializeComponent()` — L93
- `.BtnClose_Click()` — L335
- `.AndromedaEditForm_FormClosing()` — L341
- `.AndromedaEditForm_Activated()` — L349
- `.DisplayRowNumbers()` — L496
- `.DisplayMenuNumbers()` — L505
- `.SetEncoderNumber()` — L516
- `.SetPushbuttonNumber()` — L525
- `.BtnSave_Click()` — L532
- `.BtnReset_Click()` — L538
- `.TabControl1_SelectedIndexChanged()` — L546
- `.btnDelete_Click()` — L579
- `.MenuDataGridView_CellValueChanged()` — L612
- `.BtnInsert_Click()` — L636
- `.AndromedaEditForm_Load()` — L691
- `.btnG2Reset_Click()` — L696

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/AndromedaEditForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
