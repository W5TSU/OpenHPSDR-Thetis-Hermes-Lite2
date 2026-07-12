# `Console/frmInfoBarPopup.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** The info bar (status/warning strip) and its popup.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Invoke/checkboxts.cs` (references ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmInfoBarPopup` (type, L53)

- **`.SetStates()`** — L75 — `public void SetStates(Dictionary<ucInfoBar.ActionTypes, ucInfoBar.ActionState> states, ucInfoBar.ActionState b1, ucInfoBar.ActionState b2)`
  Sets states.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getCheckboxesDictionary()`** — L120 — `private Dictionary<string, CheckBoxTS> getCheckboxesDictionary()`
  Returns checkboxes dictionary.
  Called by: `.SetStates()` (same file), `.GetPopupButton()` (same file)
- **`.chkButton1_MouseUp()`** — L133 — `private void chkButton1_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkButton1` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.GetPopupButton()`** — L155 — `public CheckBoxTS GetPopupButton(int index)`
  Returns popup button.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `PopupActionSelected` (type, L55)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmInfoBarPopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
