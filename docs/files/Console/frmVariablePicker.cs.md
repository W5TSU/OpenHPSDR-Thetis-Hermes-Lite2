# `Console/frmVariablePicker.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Small shared picker dialogs.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/MeterManager.cs` (calls ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmVariablePicker` (type, L48)

- **`.colour_for_type()`** — L118 — `private Color colour_for_type(clsVariableListItems.VariableListItemType t)`
  Called by: `.list_box_DrawItem()` (same file)
- **`.list_box_DrawItem()`** — L130 — `private void list_box_DrawItem(object sender, DrawItemEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnSelect_Click()`** — L154 — `private void btnSelect_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSelect` is clicked.
  Called by: `.lstVariables_MouseDoubleClick()` (same file)
- **`.btnCancel_Click()`** — L168 — `private void btnCancel_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCancel` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Init()`** — L174 — `public void Init(int variable, Guid g, string current, bool textoverlay_led_picker = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnDefault_Click()`** — L233 — `private void btnDefault_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDefault` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstVariables_MouseDoubleClick()`** — L239 — `private void lstVariables_MouseDoubleClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsVariableListItems` (type, L64)

- **`.ToString()`** — L99 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `VariableListItemType` (type, L66)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmVariablePicker.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
