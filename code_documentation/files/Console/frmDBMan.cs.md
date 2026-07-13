# `Console/frmDBMan.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Database manager: multiple named databases, backup/restore/switch between them.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/clsDBMan.cs` (calls ×18)
  - `Console/common.cs` (calls ×2)
  - `Console/clsHardwareSpecific.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmDBMan` (type, L48)

- **`.Restore()`** — L64 — `public void Restore()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.localDateTimeFormat()`** — L70 — `private string localDateTimeFormat(DateTime dateTime)`
  Called by: `.InitBackups()` (same file), `.InitAvailableDBs()` (same file)
- **`.formatTimeSpanWithYears()`** — L88 — `private string formatTimeSpanWithYears(TimeSpan difference)`
  Called by: `.InitBackups()` (same file), `.InitAvailableDBs()` (same file)
- **`.InitBackups()`** — L118 — `internal void InitBackups(List<DBMan.BackupFileInfo> backups)`
  Inits backups.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitAvailableDBs()`** — L144 — `internal void InitAvailableDBs(Dictionary<Guid, DBMan.DatabaseInfo> dbs, Guid active_guid, Guid reselect_guid)`
  Inits available dbs.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getReadableFileSize()`** — L192 — `private string getReadableFileSize(long byteSize)`
  Returns readable file size.
  Called by: `.InitAvailableDBs()` (same file)
- **`.lstActiveDBs_ItemCheck()`** — L208 — `private void lstActiveDBs_ItemCheck(object sender, ItemCheckEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmDBMan_Shown()`** — L216 — `private void frmDBMan_Shown(object sender, EventArgs e)`
  WinForms event handler: runs when `frmDBMan` is shown.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.frmDBMan_FormClosing()`** — L221 — `private void frmDBMan_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmDBMan` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMakeActive_Click()`** — L231 — `private void btnMakeActive_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMakeActive` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnNewDB_Click()`** — L242 — `private void btnNewDB_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnNewDB` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDuplicateBD_Click()`** — L247 — `private void btnDuplicateBD_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDuplicateBD` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRemoveDB_Click()`** — L258 — `private void btnRemoveDB_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRemoveDB` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstActiveDBs_SelectedIndexChanged()`** — L268 — `private void lstActiveDBs_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstActiveDBs` selection changes.
  Called by: `.InitAvailableDBs()` (same file)
- **`.btnTakeBackupNow_Click()`** — L311 — `private void btnTakeBackupNow_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTakeBackupNow` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstBackups_SelectedIndexChanged()`** — L323 — `private void lstBackups_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstBackups` selection changes.
  Called by: `.InitBackups()` (same file)
- **`.btnMakeBackupAvailable_Click()`** — L332 — `private void btnMakeBackupAvailable_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMakeBackupAvailable` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRemoveBackup_Click()`** — L343 — `private void btnRemoveBackup_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRemoveBackup` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstActiveDBs_ColumnWidthChanging()`** — L365 — `private void lstActiveDBs_ColumnWidthChanging(object sender, ColumnWidthChangingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lstBackups_ColumnWidthChanging()`** — L371 — `private void lstBackups_ColumnWidthChanging(object sender, ColumnWidthChangingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnBackupOnStart_Click()`** — L378 — `private void btnBackupOnStart_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnBackupOnStart` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnBackupOnShutdown_Click()`** — L389 — `private void btnBackupOnShutdown_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnBackupOnShutdown` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnImport_Click()`** — L400 — `private void btnImport_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnImport` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnExport_Click()`** — L405 — `private void btnExport_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnExport` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRename_Click()`** — L416 — `private void btnRename_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRename` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnExportBackup_Click()`** — L427 — `private void btnExportBackup_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnExportBackup` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnOpenFolder_Click()`** — L436 — `private void btnOpenFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnOpenFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnImport_to_available_list_Click()`** — L447 — `private void btnImport_to_available_list_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnImport_to_available_list` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRenameBackup_Click()`** — L459 — `private void btnRenameBackup_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRenameBackup` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPruneBackups_CheckedChanged()`** — L475 — `private void chkPruneBackups_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPruneBackups` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmDBMan.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
