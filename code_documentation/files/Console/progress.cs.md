# `Console/progress.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Startup splash screen and progress reporting during initialization.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×4, calls ×3)
  - `Console/setup.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Invoke/buttonts.cs` (references ×1)
- Most-referenced symbols from other files: `.SetPercent()` (×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Progress` (type, L51)

- **`.Dispose()`** — L73 — `protected override void Dispose( bool disposing )`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L89 — `private void InitializeComponent()`
  Designer-generated UI construction (creates and lays out the form’s controls).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPercent()`** — L135 — `public void SetPercent(float f)`
  Sets percent.
  Called by: `.CalibrateLevel()` (`Console/console.cs`), `.CalibratePAGain()` (`Console/console.cs`), `.LowPowerPASweep()` (`Console/console.cs`)
- **`.btnAbort_Click()`** — L145 — `private void btnAbort_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnAbort` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panel1_Paint()`** — L170 — `private void panel1_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `panel1` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Progress_Closing()`** — L200 — `private void Progress_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `Progress` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/progress.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
