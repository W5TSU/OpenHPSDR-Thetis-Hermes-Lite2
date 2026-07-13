# `Console/Andromeda/ModeButtonsPopup.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** On-screen popups for band/filter/mode selection from the panel.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×3, references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.RepopulateForm()` (×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ModeButtonsPopup` (type, L35)

- **`.Dispose()`** — L71 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L88 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RepopulateForm()`** — L362 — `public void RepopulateForm()`
  reloads the current button settings
  Called by: `.ModeButtonsPopup_Activated()` (same file), `.repopulateForms()` (`Console/console.cs`), `.radModeButton_CheckedChanged()` (`Console/console.cs`), `.radRX2ModeButton_CheckedChanged()` (`Console/console.cs`)
- **`.btnClose_Click()`** — L423 — `private void btnClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ModeButtonsPopup_FormClosing()`** — L428 — `private void ModeButtonsPopup_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `ModeButtonsPopup` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ModeButtonsPopup_Activated()`** — L435 — `private void ModeButtonsPopup_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radBtn1_Click()`** — L440 — `private void radBtn1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn2_Click()`** — L448 — `private void radBtn2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn3_Click()`** — L456 — `private void radBtn3_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn3` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn4_Click()`** — L464 — `private void radBtn4_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn4` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn5_Click()`** — L472 — `private void radBtn5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn6_Click()`** — L480 — `private void radBtn6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn6` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn7_Click()`** — L488 — `private void radBtn7_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn7` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn8_Click()`** — L496 — `private void radBtn8_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn8` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn9_Click()`** — L504 — `private void radBtn9_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn9` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn10_Click()`** — L512 — `private void radBtn10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn10` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn11_Click()`** — L520 — `private void radBtn11_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn11` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn12_Click()`** — L528 — `private void radBtn12_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn12` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/ModeButtonsPopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
