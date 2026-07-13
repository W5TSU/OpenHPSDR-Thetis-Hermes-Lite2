# `Console/frmReleaseNotes.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Version identification, release notes, and About box.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmReleaseNotes` (type, L55)

- **`.btnClose_Click()`** — L65 — `private void btnClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.InitPath()`** — L70 — `public void InitPath(string directoryPath)`
  Inits path.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowReleaseNotes()`** — L75 — `public void ShowReleaseNotes()`
  Shows release notes.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WebBrowser1_Navigating()`** — L98 — `private void WebBrowser1_Navigating(object sender, WebBrowserNavigatingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmReleaseNotes_FormClosing()`** — L106 — `private void frmReleaseNotes_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmReleaseNotes` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmReleaseNotes.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
