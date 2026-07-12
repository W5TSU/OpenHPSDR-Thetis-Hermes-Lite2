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

### Types

#### `frmDBMan` (type, L48)

- `.Restore()` — L64
- `.localDateTimeFormat()` — L70
- `.formatTimeSpanWithYears()` — L88
- `.InitBackups()` — L118
- `.InitAvailableDBs()` — L144
- `.getReadableFileSize()` — L192
- `.lstActiveDBs_ItemCheck()` — L208
- `.frmDBMan_Shown()` — L216
- `.frmDBMan_FormClosing()` — L221
- `.btnMakeActive_Click()` — L231
- `.btnNewDB_Click()` — L242
- `.btnDuplicateBD_Click()` — L247
- `.btnRemoveDB_Click()` — L258
- `.lstActiveDBs_SelectedIndexChanged()` — L268
- `.btnTakeBackupNow_Click()` — L311
- `.lstBackups_SelectedIndexChanged()` — L323
- `.btnMakeBackupAvailable_Click()` — L332
- `.btnRemoveBackup_Click()` — L343
- `.lstActiveDBs_ColumnWidthChanging()` — L365
- `.lstBackups_ColumnWidthChanging()` — L371
- `.btnBackupOnStart_Click()` — L378
- `.btnBackupOnShutdown_Click()` — L389
- `.btnImport_Click()` — L400
- `.btnExport_Click()` — L405
- `.btnRename_Click()` — L416
- `.btnExportBackup_Click()` — L427
- `.btnOpenFolder_Click()` — L436
- `.btnImport_to_available_list_Click()` — L447
- `.btnRenameBackup_Click()` — L459
- `.chkPruneBackups_CheckedChanged()` — L475

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmDBMan.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
