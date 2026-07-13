# `Console/Andromeda/FilterButtonsPopup.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** On-screen popups for band/filter/mode selection from the panel.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×5, references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.RepopulateForm()` (×5)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `FilterButtonsPopup` (type, L35)

- **`.Dispose()`** — L71 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L88 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RepopulateForm()`** — L361 — `public void RepopulateForm()`
  reloads the current button settings
  Called by: `.FilterButtonsPopup_Activated()` (same file), `.repopulateForms()` (`Console/console.cs`), `.radModeButton_CheckedChanged()` (`Console/console.cs`), `.radRX2Filter_CheckedChanged()` (`Console/console.cs`), `.radFilter_CheckedChanged()` (`Console/console.cs`), `.radRX2ModeButton_CheckedChanged()` (`Console/console.cs`)
- **`.btnClose_Click()`** — L461 — `private void btnClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.FilterButtonsPopup_FormClosing()`** — L466 — `private void FilterButtonsPopup_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `FilterButtonsPopup` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.FilterButtonsPopup_Activated()`** — L473 — `private void FilterButtonsPopup_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radBtn1_Click()`** — L478 — `private void radBtn1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn2_Click()`** — L486 — `private void radBtn2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn3_Click()`** — L494 — `private void radBtn3_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn3` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn4_Click()`** — L502 — `private void radBtn4_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn4` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn5_Click()`** — L510 — `private void radBtn5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn6_Click()`** — L518 — `private void radBtn6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn6` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn7_Click()`** — L526 — `private void radBtn7_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn7` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn8_Click()`** — L534 — `private void radBtn8_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn8` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn9_Click()`** — L542 — `private void radBtn9_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn9` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn10_Click()`** — L550 — `private void radBtn10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn10` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn11_Click()`** — L558 — `private void radBtn11_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn11` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn12_Click()`** — L566 — `private void radBtn12_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn12` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/FilterButtonsPopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
