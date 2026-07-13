# `Console/Andromeda/ModeDependentSettingsForm.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ModeDependentSettingsForm` (type, L36)

- **`.Dispose()`** — L63 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L81 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BtnClose_Click()`** — L117 — `private void BtnClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ModeDependentSettingsForm_FormClosing()`** — L122 — `private void ModeDependentSettingsForm_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `ModeDependentSettingsForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/ModeDependentSettingsForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
