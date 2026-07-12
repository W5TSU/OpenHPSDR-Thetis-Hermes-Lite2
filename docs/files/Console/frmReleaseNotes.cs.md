# `Console/frmReleaseNotes.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Version identification, release notes, and About box.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)

## Outline

### Types

#### `frmReleaseNotes` (type, L55)

- `.btnClose_Click()` — L65
- `.InitPath()` — L70
- `.ShowReleaseNotes()` — L75
- `.WebBrowser1_Navigating()` — L98
- `.frmReleaseNotes_FormClosing()` — L106

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmReleaseNotes.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
