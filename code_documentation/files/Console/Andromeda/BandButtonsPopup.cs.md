# `Console/Andromeda/BandButtonsPopup.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** On-screen popups for band/filter/mode selection from the panel.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1, calls ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.RepopulateForm()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `BandButtonsPopup` (type, L36)

- **`.Dispose()`** — L77 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L94 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RepopulateForm()`** — L422 — `public void RepopulateForm()`
  relabel the controls according to the current set of bands RX2: always HF; else set by "Visible" property on the 3 band panels
  Called by: `.BandButtonsPopup_Activated()` (same file), `.radBtn13_Click()` (same file), `.radBtn15_Click()` (same file), `.repopulateForms()` (`Console/console.cs`)
- **`.btnClose_Click()`** — L651 — `private void btnClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BandButtonsPopup_FormClosing()`** — L656 — `private void BandButtonsPopup_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `BandButtonsPopup` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BandButtonsPopup_Activated()`** — L663 — `private void BandButtonsPopup_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radBtn1_Click()`** — L669 — `private void radBtn1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn2_Click()`** — L680 — `private void radBtn2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn3_Click()`** — L691 — `private void radBtn3_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn3` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn4_Click()`** — L702 — `private void radBtn4_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn4` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn5_Click()`** — L713 — `private void radBtn5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn6_Click()`** — L724 — `private void radBtn6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn6` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn7_Click()`** — L735 — `private void radBtn7_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn7` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn8_Click()`** — L746 — `private void radBtn8_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn8` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn9_Click()`** — L757 — `private void radBtn9_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn9` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn10_Click()`** — L768 — `private void radBtn10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn10` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn11_Click()`** — L779 — `private void radBtn11_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn11` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn12_Click()`** — L790 — `private void radBtn12_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn12` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn13_Click()`** — L800 — `private void radBtn13_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn13` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn14_Click()`** — L814 — `private void radBtn14_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn14` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBtn15_Click()`** — L825 — `private void radBtn15_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBtn15` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/BandButtonsPopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
