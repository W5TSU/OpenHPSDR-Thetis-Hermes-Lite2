# `Console/frmAbout.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Version identification, release notes, and About box.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×6)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmAbout` (type, L50)

- **`.InitVersions()`** — L97 — `public void InitVersions(string version, string build, string db_version, string dx_version, string radio_model, string firmware_version, string protocol, string supported_protocol`
  Inits versions.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnOK_Click()`** — L148 — `private void btnOK_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnOK` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCopyContributors_Click()`** — L153 — `private void btnCopyContributors_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCopyContributors` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSysInfo_Click()`** — L168 — `private void btnSysInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSysInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDXDiag_Click()`** — L173 — `private void btnDXDiag_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDXDiag` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lnkLicence_LinkClicked()`** — L178 — `private void lnkLicence_LinkClicked(object sender, LinkLabelLinkClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVisit_Click()`** — L183 — `private void btnVisit_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnVisit` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstLinks_SelectedIndexChanged()`** — L221 — `private void lstLinks_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstLinks` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.handleVersionInfo()`** — L225 — `private void handleVersionInfo()`
  Called by: `.fetchJsonAsync()` (same file)
- **`.cancelFetchJsonTask()`** — L260 — `private void cancelFetchJsonTask()`
  Called by: `.InitVersions()` (same file)
- **`.fetchJsonAsync()`** — L269 — `private async Task fetchJsonAsync(CancellationToken cancellationToken)`
  Called by: `.InitVersions()` (same file)
- **`.btnUpdatedRelease_Click()`** — L322 — `private void btnUpdatedRelease_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnUpdatedRelease` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnReleaseNotes_Click()`** — L328 — `private void btnReleaseNotes_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnReleaseNotes` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `ThetisVersionInfo` (type, L54)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmAbout.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
