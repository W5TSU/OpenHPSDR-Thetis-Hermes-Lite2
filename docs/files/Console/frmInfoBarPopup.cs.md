# `Console/frmInfoBarPopup.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** The info bar (status/warning strip) and its popup.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Invoke/checkboxts.cs` (references ×2)

## Outline

### Types

#### `frmInfoBarPopup` (type, L53)

- `.SetStates()` — L75
- `.getCheckboxesDictionary()` — L120
- `.chkButton1_MouseUp()` — L133
- `.GetPopupButton()` — L155

#### `PopupActionSelected` (type, L55)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmInfoBarPopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
