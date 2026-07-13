# `Console/CAT/CATTester.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Serial receive event plumbing and an interactive CAT test window.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmMacroButtonConfig.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Invoke/textboxts.cs` (calls ×2, references ×1)
  - `Console/CAT/CATParser.cs` (references ×1, calls ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CATTester` (type, L13)

- **`.Dispose()`** — L47 — `protected override void Dispose( bool disposing )`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Setup()`** — L59 — `private void Setup()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L82 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnExit_Click()`** — L179 — `private void btnExit_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnExit` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtInput_KeyUp()`** — L184 — `private void txtInput_KeyUp(object sender, System.Windows.Forms.KeyEventArgs e)`
  WinForms event handler: runs when `txtInput` receives a key-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ExecuteCommand()`** — L192 — `private void ExecuteCommand()`
  Called by: `.CheckText()` (same file)
- **`.btnExecute_Click()`** — L200 — `private void btnExecute_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnExecute` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CheckText()`** — L205 — `private void CheckText()`
  Checks text.
  Called by: `.txtInput_KeyUp()` (same file), `.btnExecute_Click()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/CATTester.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
