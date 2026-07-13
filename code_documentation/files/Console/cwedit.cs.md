# `Console/cwedit.cs`

**Functional area:** [11. CW keying](../../CODE_OUTLINE.md#11-cw-keying)

**Role:** Editor for CWX stored messages.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `cwedit` (type, L50)

- **`.Dispose()`** — L82 — `protected override void Dispose( bool disposing )`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L100 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.extract_fields()`** — L252 — `private void extract_fields()`
  Called by: `.cwedit_Load()` (same file)
- **`.make_current()`** — L260 — `private void make_current()`
  Called by: `.cwedit_Load()` (same file), `.txtComments_Leave()` (same file), `.txtElements_Leave()` (same file)
- **`.cwedit_Load()`** — L266 — `private void cwedit_Load(object sender, System.EventArgs e)`
  WinForms event handler: runs when `cwedit` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.saveButton_Click()`** — L279 — `private void saveButton_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `saveButton` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.cancelButton_Click()`** — L285 — `private void cancelButton_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `cancelButton` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.slen()`** — L297 — `private string slen(string s,int len)`
  Called by: `.txtComments_Leave()` (same file), `.txtElements_Leave()` (same file)
- **`.txtComments_Leave()`** — L304 — `private void txtComments_Leave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtComments` is left.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtElements_Leave()`** — L317 — `private void txtElements_Leave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtElements` is left.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/cwedit.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
