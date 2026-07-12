# `Console/Andromeda/AndromedaEditForm.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Editor for assigning functions to Andromeda encoders and buttons.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/Andromeda.cs` (calls ×2, references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.SetEncoderNumber()` (×1), `.SetPushbuttonNumber()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `AndromedaEditForm` (type, L36)

- **`.Dispose()`** — L77 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L93 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BtnClose_Click()`** — L335 — `private void BtnClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.AndromedaEditForm_FormClosing()`** — L341 — `private void AndromedaEditForm_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `AndromedaEditForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.AndromedaEditForm_Activated()`** — L349 — `private void AndromedaEditForm_Activated(object sender, EventArgs e)`
  Called by: `.BtnReset_Click()` (same file), `.btnG2Reset_Click()` (same file)
- **`.DisplayRowNumbers()`** — L496 — `private void DisplayRowNumbers(DataGridView dgv)`
  this displays the row numbers in the row headers to show encoder/button/indicator number
  Called by: `.AndromedaEditForm_Activated()` (same file), `.TabControl1_SelectedIndexChanged()` (same file)
- **`.DisplayMenuNumbers()`** — L505 — `private void DisplayMenuNumbers(DataGridView dgv)`
  this displays the menu numbers in the row headers for a menu
  Called by: `.AndromedaEditForm_Activated()` (same file), `.TabControl1_SelectedIndexChanged()` (same file), `.btnDelete_Click()` (same file), `.BtnInsert_Click()` (same file)
- **`.SetEncoderNumber()`** — L516 — `public void SetEncoderNumber(int Encoder)`
  function to select an encoder row. Used to allow an encoder turn to highlight its data. parameter is encoder number (0-19)
  Called by: `.HandleFrontPanelEncoderStep()` (`Console/Andromeda/Andromeda.cs`)
- **`.SetPushbuttonNumber()`** — L525 — `public void SetPushbuttonNumber(int Button)`
  function to select a pushbutton row. Used to allow a button press to highlight its data. parameter is button number (0-49)
  Called by: `.HandleFrontPanelButtonPress()` (`Console/Andromeda/Andromeda.cs`)
- **`.BtnSave_Click()`** — L532 — `private void BtnSave_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnSave` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BtnReset_Click()`** — L538 — `private void BtnReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TabControl1_SelectedIndexChanged()`** — L546 — `private void TabControl1_SelectedIndexChanged(object sender, EventArgs e)`
  I don't know why but we have to set the row numbers when the tab is selected
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDelete_Click()`** — L579 — `private void btnDelete_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDelete` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MenuDataGridView_CellValueChanged()`** — L612 — `private void MenuDataGridView_CellValueChanged(object sender, DataGridViewCellEventArgs e)`
  this is called when a cell value is changed in the menu window if it is the combo box column, also retrieve text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BtnInsert_Click()`** — L636 — `private void BtnInsert_Click(object sender, EventArgs e)`
  if inserting at the end, simply add new rows; otherwise they will need to be inserted
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.AndromedaEditForm_Load()`** — L691 — `private void AndromedaEditForm_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `AndromedaEditForm` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnG2Reset_Click()`** — L696 — `private void btnG2Reset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnG2Reset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/AndromedaEditForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
