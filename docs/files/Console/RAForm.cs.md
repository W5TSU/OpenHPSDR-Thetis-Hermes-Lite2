# `Console/RAForm.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Radio-astronomy data collection utility (niche feature retained from upstream).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `RAForm` (type, L55)

- **`.Dispose()`** — L79 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RArecordCheckBox_CheckedChanged()`** — L125 — `private void RArecordCheckBox_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `RArecordCheckBox` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.construct_header()`** — L190 — `private void construct_header()`
  Called by: `.RArecordCheckBox_CheckedChanged()` (same file)
- **`.write_line_to_file()`** — L204 — `private void write_line_to_file(string s)`
  Called by: `.construct_header()` (same file)
- **`.numericUpDownTS1_ValueChanged()`** — L212 — `private void numericUpDownTS1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `numericUpDownTS1` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.RA_timer_Tick()`** — L218 — `private void RA_timer_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `RA_timer` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.numericUpDownTS2_ValueChanged()`** — L282 — `private void numericUpDownTS2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `numericUpDownTS2` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picRAGraph_Paint()`** — L288 — `private void picRAGraph_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  RA graphics plot
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.button_dBm_CheckedChanged()`** — L648 — `private void button_dBm_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `button_dBm` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.button_linear_CheckedChanged()`** — L670 — `private void button_linear_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `button_linear` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picRAGraph_MouseMove()`** — L676 — `private void picRAGraph_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `picRAGraph` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.RAForm_MouseMove()`** — L687 — `private void RAForm_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `RAForm` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.button_readFile_Click()`** — L693 — `private void button_readFile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `button_readFile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MyTryParse()`** — L765 — `private bool MyTryParse(string inValue, out float result)`
  Called by: `.button_readFile_Click()` (same file)
- **`.button_writeFile_Click()`** — L779 — `private void button_writeFile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `button_writeFile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.manual_ymax_ValueChanged()`** — L817 — `private void manual_ymax_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `manual_ymax` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.manual_ymin_ValueChanged()`** — L824 — `private void manual_ymin_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `manual_ymin` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.manual_xmax_ValueChanged()`** — L832 — `private void manual_xmax_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `manual_xmax` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.manual_xmin_ValueChanged()`** — L839 — `private void manual_xmin_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `manual_xmin` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.StrToByteArray()`** — L846 — `private byte[] StrToByteArray(string str)`
  Called by: `.write_line_to_file()` (same file), `.RA_timer_Tick()` (same file)
- **`.RAForm_Load()`** — L852 — `private void RAForm_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `RAForm` loads.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `Prompt` (type, L858)

- **`.ShowDialog()`** — L860 — `public static string ShowDialog(string text, string caption)`
  Shows dialog.
  Called by: `.button_writeFile_Click()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/RAForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
