# `Console/AmpView.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** PureSignal TX linearization control panel and the amplifier gain/phase view (backed by wdsp `calcc.c`/`iqc.c`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)
  - `Console/PSForm.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `AmpView` (type, L58)

- **`.AmpView_Load()`** — L84 — `private void AmpView_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `AmpView` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.disp_setup()`** — L107 — `private void disp_setup()`
  Called by: `.timer1_Tick()` (same file)
- **`.init_data()`** — L123 — `private void init_data(int ints, int spi)`
  MW0LGE [2.9.0.8] re-factored to use fixed set of chart points, which get adjusted, these poins are re-init under certain conditions
  Called by: `.timer1_Tick()` (same file)
- **`.disp_data_Update()`** — L156 — `private void disp_data_Update(int ints, int spi)`
  Called by: `.timer1_Tick()` (same file)
- **`.chkStayOnTop_CheckedChanged()`** — L329 — `private void chkStayOnTop_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkStayOnTop` checked state changes.
  Called by: `.AmpView_Load()` (same file)
- **`.CloseDown()`** — L335 — `public void CloseDown()`
  Closes down.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.timer1_Tick()`** — L355 — `private void timer1_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `timer1` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAVShowGain_CheckedChanged()`** — L435 — `private void chkAVShowGain_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAVShowGain` checked state changes.
  Called by: `.AmpView_Load()` (same file)
- **`.chkAVLowRes_CheckedChanged()`** — L457 — `private void chkAVLowRes_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAVLowRes` checked state changes.
  Called by: `.AmpView_Load()` (same file)
- **`.AmpView_FormClosing()`** — L465 — `private void AmpView_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `AmpView` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAVPhaseZoom_CheckedChanged()`** — L470 — `private void chkAVPhaseZoom_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAVPhaseZoom` checked state changes.
  Called by: `.AmpView_Load()` (same file)
- **`.AmpView_FormClosed()`** — L484 — `private void AmpView_FormClosed(object sender, FormClosedEventArgs e)`
  WinForms event handler: runs when `AmpView` has closed.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetWindowPos()`** — L490 — `[DllImport("user32.dll")] private static extern bool SetWindowPos(IntPtr hWnd, IntPtr hWndInsertAfter, int x, int y, int cx, int cy, uint uFlags)`
  Sets window pos.
  Called by: `.FixOnTop()` (same file)
- **`.FixOnTop()`** — L501 — `public void FixOnTop()`
  Called by: `.chkStayOnTop_CheckedChanged()` (same file), `.OnShown()` (same file)
- **`.OnShown()`** — L521 — `protected override void OnShown(EventArgs e)`
  Handles/raises the shown event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/AmpView.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
