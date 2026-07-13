# `Console/frmMeterDisplay.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** The meter user control and the floating multi-meter display window.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/ucMeter.Designer.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmMeterDisplay` (type, L49)

- **`.OnRX2Enabled()`** — L90 — `private void OnRX2Enabled(bool enabled)`
  Handles/raises the rx2 enabled event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnWindowStateChanged()`** — L114 — `private void OnWindowStateChanged(FormWindowState state)`
  Handles/raises the window state changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTitle()`** — L140 — `private void setTitle()`
  Sets title.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmMeterDisplay_FormClosing()`** — L158 — `private void frmMeterDisplay_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmMeterDisplay` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TakeOwner()`** — L168 — `public void TakeOwner(ucMeter m)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetWindowPos()`** — L180 — `[DllImport("user32.dll")] private static extern bool SetWindowPos(IntPtr hWnd, IntPtr hWndInsertAfter, int X, int Y, int cx, int cy, uint uFlags)`
  Sets window pos.
  Called by: `.OnLoad()` (same file)
- **`.OnLoad()`** — L184 — `protected override void OnLoad(EventArgs e)`
  Handles/raises the load event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmMeterDisplay.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
