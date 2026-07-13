# `Console/frmLog.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Diagnostic/status logging windows.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmLog` (type, L46)

- **`.btnClear_Click()`** — L60 — `private void btnClear_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClear` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLog_CheckedChanged()`** — L69 — `private void chkLog_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLog` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Log()`** — L82 — `public void Log(bool bIn, string sMessage)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowWithTitle()`** — L111 — `public void ShowWithTitle(string title)`
  Shows with title.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmLog_FormClosing()`** — L117 — `private void frmLog_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmLog` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmLog.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
