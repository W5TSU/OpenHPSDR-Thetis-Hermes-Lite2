# `Console/database.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Settings persistence — reads/writes all option and control state to the database (XML/SQLite), including import/merge on upgrade.

## How this file is used

- Used by (incoming references from other files):
  - `Console/clsDBMan.cs` (calls ×16)
  - `Console/clsBandStackManager.cs` (calls ×7)
  - `Console/setup.cs` (calls ×7)
  - `Console/console.cs` (calls ×6)
  - `Console/common.cs` (calls ×2)
  - `Console/wideband.cs` (calls ×2)
  - `Console/MeterManager.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/clsBandStackManager.cs` (references ×5, calls ×3)
  - `Console/common.cs` (calls ×2)
  - `Console/clsDBMan.cs` (calls ×1)
  - `Console/titlebar.cs` (calls ×1)
  - `Console/radio.cs` (calls ×1)
  - `Console/enums.cs` (references ×1)
- Most-referenced symbols from other files: `.GetVarsDictionary()` (×8), `.Init()` (×4), `.Exit()` (×4), `.BandText()` (×3), `.SaveVars()` (×3), `.GetVars()` (×3), `.AddBandStack2Entry()` (×2), `.ImportAndMergeDatabase()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `DB` (type, L59)

- **`.VerifyTables()`** — L199 — `private static void VerifyTables()`
  _table_snapshots.Remove(table); }
  Called by: `.Init()` (same file)
- **`.AddBandStack2HiddenEntriesTable()`** — L253 — `private static void AddBandStack2HiddenEntriesTable(string sTableName)`
  BandStack2
  Called by: `.VerifyTables()` (same file)
- **`.AddBandStack2FilterBandsTable()`** — L261 — `private static void AddBandStack2FilterBandsTable(string sTableName)`
  Adds band stack2 filter bands table.
  Called by: `.VerifyTables()` (same file)
- **`.AddBandStack2FilterSubModesTable()`** — L269 — `private static void AddBandStack2FilterSubModesTable(string sTableName)`
  Adds band stack2 filter sub modes table.
  Called by: `.VerifyTables()` (same file)
- **`.AddBandStack2FilterModesTable()`** — L277 — `private static void AddBandStack2FilterModesTable(string sTableName)`
  Adds band stack2 filter modes table.
  Called by: `.VerifyTables()` (same file)
- **`.AddBandStack2FilterFrequenciesTable()`** — L285 — `private static void AddBandStack2FilterFrequenciesTable(string sTableName)`
  Adds band stack2 filter frequencies table.
  Called by: `.VerifyTables()` (same file)
- **`.AddBandStack2FiltersTable()`** — L303 — `private static void AddBandStack2FiltersTable(string sTableName)`
  Adds band stack2 filters table.
  Called by: `.VerifyTables()` (same file)
- **`.AddBandStack2EntriesTable()`** — L341 — `private static void AddBandStack2EntriesTable(string sTableName)`
  Adds band stack2 entries table.
  Called by: `.VerifyTables()` (same file)
- **`.SaveBandStack2Filter()`** — L366 — `public static void SaveBandStack2Filter(BandStackFilter bsf)`
  Saves band stack2 filter.
  Called by: `.SaveToDB()` (`Console/clsBandStackManager.cs`)
- **`.RemoveBandStack2Entry()`** — L493 — `public static void RemoveBandStack2Entry(BandStackEntry bse)`
  Removes band stack2 entry.
  Called by: `.AddBandStack2Entry()` (same file), `.DeleteEntry()` (`Console/clsBandStackManager.cs`)
- **`.removeReferencesToThisBSEGuid()`** — L505 — `private static void removeReferencesToThisBSEGuid(string sGUID)`
  Called by: `.RemoveBandStack2Entry()` (same file), `.RemoveAllBandStack2Entries()` (same file)
- **`.RemoveAllBandStack2Entries()`** — L544 — `public static void RemoveAllBandStack2Entries()`
  Removes all band stack2 entries.
  Called by: `.RegionReset()` (`Console/clsBandStackManager.cs`)
- **`.AddBandStack2Entry()`** — L578 — `public static void AddBandStack2Entry(BandStackEntry bse)`
  Adds band stack2 entry.
  Called by: `.SaveToDB()` (`Console/clsBandStackManager.cs`), `.AddEntry()` (`Console/clsBandStackManager.cs`)
- **`.GetBandStack2Entries()`** — L610 — `public static List<BandStackEntry> GetBandStack2Entries()`
  Returns band stack2 entries.
  Called by: `.initLists()` (`Console/clsBandStackManager.cs`)
- **`.GetBandStack2Filters()`** — L638 — `public static Dictionary<string, BandStackFilter> GetBandStack2Filters()`
  Returns band stack2 filters.
  Called by: `.initLists()` (`Console/clsBandStackManager.cs`)
- **`.checkForPrimaryKeys()`** — L732 — `private static void checkForPrimaryKeys(string name)`
  Called by: `.AddFormTable()` (same file), `.RemoveVarsList()` (same file), `.SaveVarsDictionary()` (same file), `.SaveVars()` (same file)
- **`.AddFormTable()`** — L742 — `private static void AddFormTable(string name)`
  Adds form table.
  Called by: `.RemoveVarsList()` (same file), `.SaveVarsDictionary()` (same file), `.SaveVars()` (same file)
- **`.AddBandTextTable()`** — L751 — `private static void AddBandTextTable()`
  Adds band text table.
  Called by: `.VerifyTables()` (same file)
- **`.AddRegion2BandText()`** — L1074 — `private static void AddRegion2BandText()`
  Adds region2 band text.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionJapanBandText()`** — L1401 — `private static void AddRegionJapanBandText()`
  Adds region japan band text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddBandTextSWB()`** — L1484 — `private static void AddBandTextSWB()`
  Adds band text swb.
  Called by: `.AddBandTextTable()` (same file), `.AddRegion2BandText()` (same file), `.AddRegionJapanBandText()` (same file), `.UpdateRegion()` (same file)
- **`.ClearBandText()`** — L1636 — `private static void ClearBandText()`
  Clears band text.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText160m()`** — L1641 — `private static void AddRegion1BandText160m()`
  Adds region1 band text160m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText80m()`** — L1664 — `private static void AddRegion1BandText80m()`
  Adds region1 band text80m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText60m()`** — L1690 — `private static void AddRegion1BandText60m()`
  Adds region1 band text60m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText40m()`** — L1709 — `private static void AddRegion1BandText40m()`
  Adds region1 band text40m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText30m()`** — L1735 — `private static void AddRegion1BandText30m()`
  Adds region1 band text30m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText20m()`** — L1756 — `private static void AddRegion1BandText20m()`
  Adds region1 band text20m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText17m()`** — L1788 — `private static void AddRegion1BandText17m()`
  Adds region1 band text17m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText15m()`** — L1813 — `private static void AddRegion1BandText15m()`
  Adds region1 band text15m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText12m()`** — L1844 — `private static void AddRegion1BandText12m()`
  Adds region1 band text12m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText10m()`** — L1868 — `private static void AddRegion1BandText10m()`
  Adds region1 band text10m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText6m()`** — L1911 — `private static void AddRegion1BandText6m()`
  Adds region1 band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText160m()`** — L1952 — `private static void AddRegionIndiaBandText160m()`
  Adds region india band text160m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText80m()`** — L1975 — `private static void AddRegionIndiaBandText80m()`
  Adds region india band text80m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText40m()`** — L2000 — `private static void AddRegionIndiaBandText40m()`
  Adds region india band text40m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText20m()`** — L2026 — `private static void AddRegionIndiaBandText20m()`
  Adds region india band text20m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText17m()`** — L2058 — `private static void AddRegionIndiaBandText17m()`
  Adds region india band text17m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText15m()`** — L2083 — `private static void AddRegionIndiaBandText15m()`
  Adds region india band text15m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText12m()`** — L2114 — `private static void AddRegionIndiaBandText12m()`
  Adds region india band text12m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText10m()`** — L2138 — `private static void AddRegionIndiaBandText10m()`
  Adds region india band text10m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIndiaBandText6m()`** — L2181 — `private static void AddRegionIndiaBandText6m()`
  Adds region india band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText160m()`** — L2222 — `private static void AddRegionIsraelBandText160m()`
  Adds region israel band text160m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText80m()`** — L2243 — `private static void AddRegionIsraelBandText80m()`
  Adds region israel band text80m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText60m()`** — L2271 — `private static void AddRegionIsraelBandText60m()`
  Adds region israel band text60m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText40m()`** — L2290 — `private static void AddRegionIsraelBandText40m()`
  Adds region israel band text40m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText30m()`** — L2311 — `private static void AddRegionIsraelBandText30m()`
  Adds region israel band text30m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText20m()`** — L2330 — `private static void AddRegionIsraelBandText20m()`
  Adds region israel band text20m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText17m()`** — L2358 — `private static void AddRegionIsraelBandText17m()`
  Adds region israel band text17m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText15m()`** — L2379 — `private static void AddRegionIsraelBandText15m()`
  Adds region israel band text15m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText12m()`** — L2404 — `private static void AddRegionIsraelBandText12m()`
  Adds region israel band text12m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText10m()`** — L2425 — `private static void AddRegionIsraelBandText10m()`
  Adds region israel band text10m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegionIsraelBandText6m()`** — L2455 — `private static void AddRegionIsraelBandText6m()`
  Adds region israel band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandText4m()`** — L2482 — `private static void AddRegion1BandText4m()`
  Adds region1 band text4m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion1BandTextVHFplus()`** — L2517 — `private static void AddRegion1BandTextVHFplus()`
  Adds region1 band text vhfplus.
  Called by: `.UpdateRegion()` (same file)
- **`.AddBulgariaBandText160m()`** — L2722 — `private static void AddBulgariaBandText160m()`
  Adds bulgaria band text160m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddNetherlandsBandText160m()`** — L2746 — `private static void AddNetherlandsBandText160m()`
  Adds netherlands band text160m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddUK_PlusBandText60m()`** — L2770 — `private static void AddUK_PlusBandText60m()`
  Adds uk plus band text60m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddNorwayBandText60m()`** — L2811 — `private static void AddNorwayBandText60m()`
  Adds norway band text60m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddSwedenBandText60m()`** — L2830 — `private static void AddSwedenBandText60m()`
  Adds sweden band text60m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddHungaryBandText40m()`** — L2857 — `private static void AddHungaryBandText40m()`
  Adds hungary band text40m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRussiaBandText12m()`** — L2884 — `private static void AddRussiaBandText12m()`
  Adds russia band text12m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRussiaBandText11m()`** — L2908 — `private static void AddRussiaBandText11m()`
  Adds russia band text11m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddEUBandText6m()`** — L2926 — `private static void AddEUBandText6m()`
  Adds euband text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddFranceBandText6m()`** — L2960 — `private static void AddFranceBandText6m()`
  Adds france band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddLatviaBandText6m()`** — L2994 — `private static void AddLatviaBandText6m()`
  Adds latvia band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddBulgariaBandText6m()`** — L3028 — `private static void AddBulgariaBandText6m()`
  Adds bulgaria band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddGreeceBandText6m()`** — L3062 — `private static void AddGreeceBandText6m()`
  Adds greece band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion3BandText160m()`** — L3538 — `private static void AddRegion3BandText160m()`
  Adds region3 band text160m.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddRegion3BandText80m()`** — L3559 — `private static void AddRegion3BandText80m()`
  Adds region3 band text80m.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddRegion3BandText60m()`** — L3583 — `private static void AddRegion3BandText60m()`
  Adds region3 band text60m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion3BandText40m()`** — L3602 — `private static void AddRegion3BandText40m()`
  Adds region3 band text40m.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddRegion3BandText30m()`** — L3625 — `private static void AddRegion3BandText30m()`
  Adds region3 band text30m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion3BandText20m()`** — L3644 — `private static void AddRegion3BandText20m()`
  Adds region3 band text20m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion3BandText17m()`** — L3673 — `private static void AddRegion3BandText17m()`
  Adds region3 band text17m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion3BandText15m()`** — L3698 — `private static void AddRegion3BandText15m()`
  Adds region3 band text15m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion3BandText12m()`** — L3726 — `private static void AddRegion3BandText12m()`
  Adds region3 band text12m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddRegion3BandText10m()`** — L3749 — `private static void AddRegion3BandText10m()`
  Adds region3 band text10m.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddRegion3BandText6m()`** — L3777 — `private static void AddRegion3BandText6m()`
  Adds region3 band text6m.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddRegion3BandTextVHFplus()`** — L3797 — `private static void AddRegion3BandTextVHFplus()`
  Adds region3 band text vhfplus.
  Called by: `.UpdateRegion()` (same file)
- **`.AddJapanBandText160m()`** — L3915 — `private static void AddJapanBandText160m()`
  Adds japan band text160m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddJapanBandText80m()`** — L3937 — `private static void AddJapanBandText80m()`
  Adds japan band text80m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddJapanBandText40m()`** — L3970 — `private static void AddJapanBandText40m()`
  Adds japan band text40m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddJapanBandText10m()`** — L3995 — `private static void AddJapanBandText10m()`
  Adds japan band text10m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddJapanBandText6m()`** — L4024 — `private static void AddJapanBandText6m()`
  Adds japan band text6m.
  Called by: `.UpdateRegion()` (same file)
- **`.AddJapanBandTextEmergency()`** — L4048 — `private static void AddJapanBandTextEmergency()`
  Adds japan band text emergency.
  Called by: `.UpdateRegion()` (same file)
- **`.AddMemoryTable()`** — L4263 — `private static void AddMemoryTable()`
  Adds memory table.
  Called by: `.VerifyTables()` (same file)
- **`.AddGroupListTable()`** — L4284 — `private static void AddGroupListTable()`
  Adds group list table.
  Called by: `.VerifyTables()` (same file)
- **`.AddTXProfileTable()`** — L4303 — `private static void AddTXProfileTable(string sTableName, bool bIndcludeExtraProfiles = false)`
  Adds txprofile table.
  Called by: `.VerifyTables()` (same file)
- **`.CheckBandTextValid()`** — L9430 — `private static void CheckBandTextValid()`
  Checks band text valid.
  Called by: `.Init()` (same file), `.UpdateRegion()` (same file)
- **`.Init()`** — L9471 — `public static bool Init()`
  Called by: `.LoadDB()` (`Console/clsDBMan.cs`), `.createNewDB()` (`Console/clsDBMan.cs`), `.ImportAsAvailable()` (`Console/clsDBMan.cs`), `.MakeBackupAvailable()` (`Console/clsDBMan.cs`)
- **`.ConvertFromDBVal()`** — L9522 — `public static T ConvertFromDBVal<T>(object obj)`
  Converts from dbval.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WriteDB()`** — L9534 — `public static void WriteDB()`
  Writes db.
  Called by: `.VerifyTables()` (same file), `.Exit()` (same file), `.UpdateRegion()` (same file), `.Export()` (`Console/clsDBMan.cs`), `.SaveOptions()` (`Console/setup.cs`)
- **`.Exit()`** — L9564 — `public static void Exit()`
  Called by: `.createNewDB()` (`Console/clsDBMan.cs`), `.ImportAsAvailable()` (`Console/clsDBMan.cs`), `.MakeBackupAvailable()` (`Console/clsDBMan.cs`), `.ExitConsole()` (`Console/console.cs`)
- **`.BandText()`** — L9570 — `public static bool BandText(double freq, out string outStr)`
  Called by: `.UpdateBandText()` (`Console/MeterManager.cs`), `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`)
- **`.PurgeNotches()`** — L9819 — `public static void PurgeNotches()`
  This removes the notches from the state database so we can rewrite all of them without having one that was previously deleted staying in the database
  Called by: `.SaveState()` (`Console/console.cs`)
- **`.PurgeMeters()`** — L9823 — `public static void PurgeMeters(List<string>formGuids)`
  Called by: `.SaveOptions()` (`Console/setup.cs`)
- **`.purgeTableEntries()`** — L9870 — `private static void purgeTableEntries(string sTable, string sKey)`
  // find all that match and remove them DataRow[] rows = ds.Tables["State"].Select("Key like '" + sKey + "'"); if (rows != null) { foreach (DataRow row in rows) row.Delete(); } }
  Called by: `.PurgeNotches()` (same file), `.PurgeMeters()` (same file)
- **`.RemoveVarsList()`** — L9885 — `public static void RemoveVarsList(string tableName, List<string> list)`
  Removes vars list.
  Called by: `.removeOutdatedOptions()` (`Console/setup.cs`)
- **`.SaveVarsDictionary()`** — L9910 — `public static void SaveVarsDictionary(string tableName, ref Dictionary<string,string> dict, bool bSaveEmptyValues = true)`
  Saves vars dictionary.
  Called by: `.Init()` (same file), `.SaveOptions()` (`Console/setup.cs`)
- **`.SaveVars()`** — L9952 — `public static void SaveVars(string tableName, List<string> list, bool bSaveEmptyValues = true)`
  Saves vars.
  Called by: `.SaveForm()` (`Console/common.cs`), `.SaveState()` (`Console/console.cs`), `.SaveWideBand()` (`Console/wideband.cs`)
- **`.GetVarsDictionary()`** — L10001 — `public static Dictionary<string, string> GetVarsDictionary(string table_name)`
  Returns vars dictionary.
  Called by: `.Init()` (same file), `.LoadDB()` (`Console/clsDBMan.cs`), `.checkVersion()` (`Console/clsDBMan.cs`), `.DBWritten()` (`Console/clsDBMan.cs`), `.createNewDB()` (`Console/clsDBMan.cs`), `.ImportAsAvailable()` (`Console/clsDBMan.cs`) — and 3 more
- **`.GetVars()`** — L10037 — `public static List<string> GetVars(string tableName)`
  return dict; }
  Called by: `.RestoreForm()` (`Console/common.cs`), `.GetState()` (`Console/console.cs`), `.GetWideBand()` (`Console/wideband.cs`)
- **`.ImportAndMergeDatabase()`** — L10623 — `public static bool ImportAndMergeDatabase(string filename, out string log, bool ignore_merged)`
  Called by: `.checkVersion()` (`Console/clsDBMan.cs`), `.Import()` (`Console/clsDBMan.cs`)
- **`.getRadioSelectedFromOldRadButton()`** — L11219 — `private static bool getRadioSelectedFromOldRadButton(ref DataTable tempTable, string sRadButtonName)`
  --MW0LGE
  Called by: `.ImportAndMergeDatabase()` (same file)
- **`.ValidateImportedDatabase()`** — L11236 — `private static string ValidateImportedDatabase(DataSet oldDB)`
  -W2PA Basic validity checks of imported DataSet xml file
  Called by: `.ImportAndMergeDatabase()` (same file)
- **`.ExpandOldTxProfileTable()`** — L11269 — `private static DataTable ExpandOldTxProfileTable(DataTable oldTable)`
  -W2PA Expand an old TxProfile table into a newer one with more colunms. Fill in missing ones with default values.
  Called by: `.ImportAndMergeDatabase()` (same file)
- **`.WriteImportLog()`** — L11300 — `private static void WriteImportLog(string logFN, string s)`
  -W2PA Write a message to the ImportLog file during the import process
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ImportDatabase()`** — L11307 — `public static bool ImportDatabase(string filename)`
  -W2PA Original version of ImportDatabase
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRegion()`** — L11332 — `public static void UpdateRegion(FRSRegion current_region)`
  Updates region.
  Called by: `.comboFRSRegion_SelectedIndexChanged()` (`Console/setup.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/database.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
