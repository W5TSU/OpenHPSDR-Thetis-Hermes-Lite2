# `Console/InputBox.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Simple text-input dialog; searchable "find a setting" helper.

## How this file is used

- Used by (incoming references from other files):
  - `Console/clsDBMan.cs` (calls ×8)
  - `Console/setup.cs` (calls ×5)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Show()` (×13)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `InputBox` (type, L47)

- **`.Dispose()`** — L76 — `protected override void Dispose( bool disposing )`
  Clean up any resources being used.
  Called by: `.Show()` (same file)
- **`.InitializeComponent()`** — L95 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Show()`** — L158 — `public static string Show(string title, string label, string textbox, bool to_top = false)`
  Called by: `.NewDB()` (`Console/clsDBMan.cs`), `.RemoveDB()` (`Console/clsDBMan.cs`), `.DuplicateDB()` (`Console/clsDBMan.cs`), `.TakeBackup()` (`Console/clsDBMan.cs`), `.ImportAsAvailable()` (`Console/clsDBMan.cs`), `.Rename()` (`Console/clsDBMan.cs`) — and 7 more
- **`.btnOK_Click()`** — L176 — `private void btnOK_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnOK` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCancel_Click()`** — L182 — `private void btnCancel_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnCancel` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/InputBox.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
