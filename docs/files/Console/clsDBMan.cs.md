# `Console/clsDBMan.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Database manager: multiple named databases, backup/restore/switch between them.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmDBMan.cs` (calls ×18)
  - `Console/console.cs` (calls ×2)
  - `Console/database.cs` (calls ×1)
  - `Console/setup.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/database.cs` (calls ×16)
  - `Console/InputBox.cs` (calls ×8)
  - `Console/common.cs` (calls ×5)
  - `Console/clsHardwareSpecific.cs` (calls ×4)
  - `Console/enums.cs` (references ×2)
  - `Console/frmDBMan.Designer.cs` (references ×1)
  - `Console/keyboard.cs` (calls ×1)
- Most-referenced symbols from other files: `.SelectedAvailable()` (×2), `.TakeBackup()` (×2), `.ShowDBMan()` (×1), `.Shutdown()` (×1), `.DBWritten()` (×1), `.MakeActiveDB()` (×1), `.NewDB()` (×1), `.BackupOnStartUpToggle()` (×1)

## Outline

### Types

#### `DBMan` (type, L56)

- `.ShowDBMan()` — L172
- `.LoadDB()` — L229
- `.updateFileExists()` — L428
- `.renameUpdatedb()` — L437
- `.checkVersion()` — L463
- `.moveToBroken()` — L523
- `.Shutdown()` — L541
- `.DBWritten()` — L559
- `.createNewDB()` — L604
- `.getActiveDB()` — L700
- `.makeDBActive()` — L721
- `.createNewDBFolder()` — L753
- `.getAvailableDBs()` — L770
- `.getAllActiveDBGUIDs()` — L906
- `.calculateFolderSize()` — L937
- `.MakeActiveDB()` — L962
- `.NewDB()` — L990
- `.BackupOnStartUpToggle()` — L1002
- `.BackupOnShutDownToggle()` — L1025
- `.RemoveDB()` — L1048
- `.DuplicateDB()` — L1114
- `.copyFolder()` — L1167
- `.SelectedAvailable()` — L1198
- `.TakeBackup()` — L1218
- `.getBackups()` — L1282
- `.getOrderedBackupFiles()` — L1288
- `.createUniqueFilename()` — L1333
- `.RemoveBackupDB()` — L1352
- `.Import()` — L1386
- `.ImportAsAvailable()` — L1454
- `.Rename()` — L1626
- `.RenameBackup()` — L1653
- `.OpenFolder()` — L1709
- `.Export()` — L1726
- `.ExportBackup()` — L1778
- `.MakeBackupAvailable()` — L1818
- `.getWeekOfYear()` — L1919
- `.pruneForGFS()` — L1929

#### `DBSettings` (type, L58)

_No extracted members._

#### `DatabaseInfo` (type, L68)

_No extracted members._

#### `DatabaseInfoDefaultStringEnumConverter` (type, L87)

- `.ReadJson()` — L96

#### `BackupFileInfo` (type, L128)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsDBMan.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
