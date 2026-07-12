# `Console/Andromeda/displaysettingsform.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/comboboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `DisplaySettingsForm` (type, L40)

- **`.Dispose()`** — L78 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L96 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RepopulateForm()`** — L324 — `public void RepopulateForm()`
  copy settings from relevant console controls
  Called by: `.DisplaySettingsForm_Activated()` (same file)
- **`.BtnClose_Click()`** — L373 — `private void BtnClose_Click(object sender, EventArgs e)`
  close form button: simply tell the form to close
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.DisplaySettingsForm_FormClosing()`** — L378 — `private void DisplaySettingsForm_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `DisplaySettingsForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.DisplaySettingsForm_Activated()`** — L386 — `private void DisplaySettingsForm_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ComboRX1Meter_SelectedIndexChanged()`** — L392 — `private void ComboRX1Meter_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboRX1Meter` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboRX2Meter_SelectedIndexChanged()`** — L397 — `private void ComboRX2Meter_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboRX2Meter` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboTXMeter_SelectedIndexChanged()`** — L402 — `private void ComboTXMeter_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboTXMeter` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboRX1Display_SelectedIndexChanged()`** — L407 — `private void ComboRX1Display_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboRX1Display` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboRX2Display_SelectedIndexChanged()`** — L412 — `private void ComboRX2Display_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboRX2Display` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkRX1Avg_CheckedChanged()`** — L417 — `private void ChkRX1Avg_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkRX1Avg` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkRX2Avg_CheckedChanged()`** — L425 — `private void ChkRX2Avg_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkRX2Avg` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkRX1Peak_CheckedChanged()`** — L433 — `private void ChkRX1Peak_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkRX1Peak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkRX2Peak_CheckedChanged()`** — L441 — `private void ChkRX2Peak_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkRX2Peak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/displaysettingsform.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
