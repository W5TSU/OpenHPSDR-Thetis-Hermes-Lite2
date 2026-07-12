# `Console/frmAbout.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Version identification, release notes, and About box.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×6)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Types

#### `frmAbout` (type, L50)

- `.InitVersions()` — L97
- `.btnOK_Click()` — L148
- `.btnCopyContributors_Click()` — L153
- `.btnSysInfo_Click()` — L168
- `.btnDXDiag_Click()` — L173
- `.lnkLicence_LinkClicked()` — L178
- `.btnVisit_Click()` — L183
- `.lstLinks_SelectedIndexChanged()` — L221
- `.handleVersionInfo()` — L225
- `.cancelFetchJsonTask()` — L260
- `.fetchJsonAsync()` — L269
- `.btnUpdatedRelease_Click()` — L322
- `.btnReleaseNotes_Click()` — L328

#### `ThetisVersionInfo` (type, L54)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmAbout.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
