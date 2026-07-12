# `Console/frmVariablePicker.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Small shared picker dialogs.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/MeterManager.cs` (calls ×2)

## Outline

### Types

#### `frmVariablePicker` (type, L48)

- `.colour_for_type()` — L118
- `.list_box_DrawItem()` — L130
- `.btnSelect_Click()` — L154
- `.btnCancel_Click()` — L168
- `.Init()` — L174
- `.btnDefault_Click()` — L233
- `.lstVariables_MouseDoubleClick()` — L239

#### `clsVariableListItems` (type, L64)

- `.ToString()` — L99

#### `VariableListItemType` (type, L66)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmVariablePicker.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
