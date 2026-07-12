# `Console/FilterForm.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX filter preset model per mode, the filter-edit form, and the filter-set manager.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/filter.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `FilterForm` (type, L52)

- **`.Dispose()`** — L116 — `protected override void Dispose( bool disposing )`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L135 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetFilterInfo()`** — L646 — `private void GetFilterInfo()`
  Returns filter info.
  Called by: `.txtName_LostFocus()` (same file)
- **`.HzToPixel()`** — L658 — `private int HzToPixel(float freq)`
  Called by: `.picDisplay_Paint()` (same file), `.picDisplay_MouseMove()` (same file), `.picDisplay_MouseDown()` (same file)
- **`.PixelToHz()`** — L666 — `private float PixelToHz(float x)`
  Called by: `.picDisplay_MouseMove()` (same file)
- **`.UpdateFilter()`** — L675 — `private void UpdateFilter(int low, int high)`
  Updates filter.
  Called by: `.GetFilterInfo()` (same file), `.udLow_ValueChanged()` (same file), `.udHigh_ValueChanged()` (same file), `.udWidth_ValueChanged()` (same file)
- **`.radFilter_CheckedChanged()`** — L705 — `private void radFilter_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radFilter` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboDSPMode_SelectedIndexChanged()`** — L722 — `private void comboDSPMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPMode` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtName_LostFocus()`** — L727 — `private void txtName_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtName` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udLow_ValueChanged()`** — L784 — `private void udLow_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udLow` value changes.
  Called by: `.udLow_LostFocus()` (same file)
- **`.udHigh_ValueChanged()`** — L806 — `private void udHigh_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udHigh` value changes.
  Called by: `.udHigh_LostFocus()` (same file)
- **`.udLow_LostFocus()`** — L828 — `private void udLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udHigh_LostFocus()`** — L833 — `private void udHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picDisplay_Paint()`** — L838 — `private void picDisplay_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picDisplay` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picDisplay_MouseMove()`** — L855 — `private void picDisplay_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `picDisplay` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picDisplay_MouseDown()`** — L884 — `private void picDisplay_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `picDisplay` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picDisplay_MouseUp()`** — L905 — `private void picDisplay_MouseUp(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `picDisplay` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udWidth_ValueChanged()`** — L918 — `private void udWidth_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udWidth` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.FilterForm_FormClosing()`** — L963 — `private void FilterForm_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `FilterForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/FilterForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
