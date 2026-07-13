# `Console/Andromeda/vfosettingspopup.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/textboxts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `VFOSettingsPopup` (type, L36)

- **`.Dispose()`** — L60 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L77 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOSettingsPopup_FormClosing()`** — L170 — `private void VFOSettingsPopup_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `VFOSettingsPopup` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TextBoxTuneStep_MouseDown()`** — L178 — `private void TextBoxTuneStep_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `TextBoxTuneStep` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ButtonMinus_Click()`** — L185 — `private void ButtonMinus_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ButtonMinus` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ButtonPlus_Click()`** — L196 — `private void ButtonPlus_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ButtonPlus` is clicked.
  Called by: `.TextBoxTuneStep_MouseDown()` (same file)
- **`.VFOSettingsPopup_Load()`** — L206 — `private void VFOSettingsPopup_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `VFOSettingsPopup` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ButtonClose_Click()`** — L214 — `private void ButtonClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ButtonClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/vfosettingspopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
