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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `DBMan` (type, L56)

- **`.ShowDBMan()`** — L172 — `public static void ShowDBMan()`
  Shows dbman.
  Called by: `.databaseManagerToolStripMenuItem_Click()` (`Console/console.cs`)
- **`.LoadDB()`** — L229 — `public static bool LoadDB(string[] args, out string broken_folder)`
  Loads db.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateFileExists()`** — L428 — `private static bool updateFileExists()`
  Called by: `.LoadDB()` (same file)
- **`.renameUpdatedb()`** — L437 — `private static bool renameUpdatedb()`
  Called by: `.checkVersion()` (same file)
- **`.checkVersion()`** — L463 — `private static void checkVersion(bool made_new, bool force_upgrade = false, bool force_upgrade_via_file = false)`
  Called by: `.LoadDB()` (same file)
- **`.moveToBroken()`** — L523 — `private static void moveToBroken(Guid guid)`
  Called by: `.LoadDB()` (same file), `.getAvailableDBs()` (same file)
- **`.Shutdown()`** — L541 — `public static void Shutdown()`
  Called by: `.ExitConsole()` (`Console/console.cs`)
- **`.DBWritten()`** — L559 — `public static void DBWritten()`
  Called by: `.checkVersion()` (same file), `.WriteDB()` (`Console/database.cs`)
- **`.createNewDB()`** — L604 — `private static bool createNewDB(bool check_for_old_db, bool make_active, out bool old_db_found, string description = "")`
  Called by: `.LoadDB()` (same file), `.checkVersion()` (same file), `.NewDB()` (same file)
- **`.getActiveDB()`** — L700 — `private static DBSettings getActiveDB()`
  Returns active db.
  Called by: `.LoadDB()` (same file), `.createNewDB()` (same file)
- **`.makeDBActive()`** — L721 — `private static bool makeDBActive(Guid guid)`
  Called by: `.createNewDB()` (same file), `.MakeActiveDB()` (same file)
- **`.createNewDBFolder()`** — L753 — `private static string createNewDBFolder(out Guid guid)`
  Called by: `.createNewDB()` (same file), `.ImportAsAvailable()` (same file)
- **`.getAvailableDBs()`** — L770 — `private static Dictionary<Guid, DatabaseInfo> getAvailableDBs()`
  Returns available dbs.
  Called by: `.ShowDBMan()` (same file), `.LoadDB()` (same file), `.NewDB()` (same file), `.BackupOnStartUpToggle()` (same file), `.BackupOnShutDownToggle()` (same file), `.RemoveDB()` (same file) — and 4 more
- **`.getAllActiveDBGUIDs()`** — L906 — `private static List<Guid> getAllActiveDBGUIDs(string path)`
  Returns all active dbguids.
  Called by: `.getAvailableDBs()` (same file)
- **`.calculateFolderSize()`** — L937 — `private static long calculateFolderSize(DirectoryInfo directoryInfo)`
  Called by: `.DBWritten()` (same file), `.getAvailableDBs()` (same file), `.ImportAsAvailable()` (same file), `.MakeBackupAvailable()` (same file)
- **`.MakeActiveDB()`** — L962 — `public static void MakeActiveDB(Guid guid)`
  Called by: `.btnMakeActive_Click()` (`Console/frmDBMan.cs`)
- **`.NewDB()`** — L990 — `public static void NewDB()`
  Called by: `.btnNewDB_Click()` (`Console/frmDBMan.cs`)
- **`.BackupOnStartUpToggle()`** — L1002 — `public static void BackupOnStartUpToggle(Guid guid)`
  Called by: `.btnBackupOnStart_Click()` (`Console/frmDBMan.cs`)
- **`.BackupOnShutDownToggle()`** — L1025 — `public static void BackupOnShutDownToggle(Guid guid)`
  Called by: `.btnBackupOnShutdown_Click()` (`Console/frmDBMan.cs`)
- **`.RemoveDB()`** — L1048 — `public static void RemoveDB(Guid guid, bool force = false)`
  Removes db.
  Called by: `.btnRemoveDB_Click()` (`Console/frmDBMan.cs`)
- **`.DuplicateDB()`** — L1114 — `public static void DuplicateDB(Guid guid)`
  Called by: `.btnDuplicateBD_Click()` (`Console/frmDBMan.cs`)
- **`.copyFolder()`** — L1167 — `public static bool copyFolder(string sourceFolder, string destinationFolder)`
  Called by: `.DuplicateDB()` (same file)
- **`.SelectedAvailable()`** — L1198 — `public static void SelectedAvailable(Guid guid)`
  Called by: `.lstActiveDBs_SelectedIndexChanged()` (`Console/frmDBMan.cs`), `.btnRemoveBackup_Click()` (`Console/frmDBMan.cs`)
- **`.TakeBackup()`** — L1218 — `public static bool TakeBackup(Guid highlighted, string description = "", bool auto = false)`
  Called by: `.LoadDB()` (same file), `.Shutdown()` (same file), `.btnTakeBackupNow_Click()` (`Console/frmDBMan.cs`), `.btnContainer_load_Click()` (`Console/setup.cs`)
- **`.getBackups()`** — L1282 — `private static void getBackups(Guid guid)`
  Returns backups.
  Called by: `.TakeBackup()` (same file)
- **`.getOrderedBackupFiles()`** — L1288 — `private static List<BackupFileInfo> getOrderedBackupFiles(string backupFolderPath)`
  Returns ordered backup files.
  Called by: `.SelectedAvailable()` (same file), `.getBackups()` (same file), `.RenameBackup()` (same file)
- **`.createUniqueFilename()`** — L1333 — `private static string createUniqueFilename(string directoryPath)`
  Called by: `.TakeBackup()` (same file)
- **`.RemoveBackupDB()`** — L1352 — `public static void RemoveBackupDB(List<string> file_paths)`
  Removes backup db.
  Called by: `.btnRemoveBackup_Click()` (`Console/frmDBMan.cs`)
- **`.Import()`** — L1386 — `public static void Import()`
  Called by: `.btnImport_Click()` (`Console/frmDBMan.cs`)
- **`.ImportAsAvailable()`** — L1454 — `public static void ImportAsAvailable(Guid selected)`
  Called by: `.btnImport_to_available_list_Click()` (`Console/frmDBMan.cs`)
- **`.Rename()`** — L1626 — `public static void Rename(Guid guid)`
  Called by: `.btnRename_Click()` (`Console/frmDBMan.cs`)
- **`.RenameBackup()`** — L1653 — `public static void RenameBackup(Guid guid, string file_path)`
  Called by: `.btnRenameBackup_Click()` (`Console/frmDBMan.cs`)
- **`.OpenFolder()`** — L1709 — `public static void OpenFolder(Guid guid)`
  Opens folder.
  Called by: `.btnOpenFolder_Click()` (`Console/frmDBMan.cs`)
- **`.Export()`** — L1726 — `public static void Export(Guid guid)`
  Called by: `.btnExport_Click()` (`Console/frmDBMan.cs`)
- **`.ExportBackup()`** — L1778 — `public static void ExportBackup(string desc, string path)`
  Called by: `.btnExportBackup_Click()` (`Console/frmDBMan.cs`)
- **`.MakeBackupAvailable()`** — L1818 — `public static void MakeBackupAvailable(string file_path)`
  Called by: `.btnMakeBackupAvailable_Click()` (`Console/frmDBMan.cs`)
- **`.getWeekOfYear()`** — L1919 — `private static int getWeekOfYear(DateTime date)`
  Returns week of year.
  Called by: `.pruneForGFS()` (same file)
- **`.pruneForGFS()`** — L1929 — `private static void pruneForGFS(string backup_folder_path)`
  Called by: `.LoadDB()` (same file), `.TakeBackup()` (same file)

#### `DBSettings` (type, L58)

_No extracted members._

#### `DatabaseInfo` (type, L68)

_No extracted members._

#### `DatabaseInfoDefaultStringEnumConverter` (type, L87)

- **`.ReadJson()`** — L96 — `public override object ReadJson(JsonReader reader, Type objectType, object existingValue, JsonSerializer serializer)`
  Reads json.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `BackupFileInfo` (type, L128)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsDBMan.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
