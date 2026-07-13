# `Console/frmBandwidth.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Variable-bandwidth adjustment popup and its graphical bandwidth view.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmBandwidth` (type, L46)

- **`.timerReadBandwidth_Tick()`** — L57 — `private void timerReadBandwidth_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `timerReadBandwidth` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.RecoverShow()`** — L64 — `public void RecoverShow()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmBandwidth_FormClosing()`** — L71 — `private void frmBandwidth_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmBandwidth` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radUnits_CheckedChanged()`** — L84 — `private void radUnits_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUnits` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOnTop_CheckedChanged()`** — L89 — `private void chkOnTop_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOnTop` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmBandwidth.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
