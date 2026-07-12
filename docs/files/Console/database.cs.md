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

### Types

#### `DB` (type, L59)

- `.VerifyTables()` — L199
- `.AddBandStack2HiddenEntriesTable()` — L253
- `.AddBandStack2FilterBandsTable()` — L261
- `.AddBandStack2FilterSubModesTable()` — L269
- `.AddBandStack2FilterModesTable()` — L277
- `.AddBandStack2FilterFrequenciesTable()` — L285
- `.AddBandStack2FiltersTable()` — L303
- `.AddBandStack2EntriesTable()` — L341
- `.SaveBandStack2Filter()` — L366
- `.RemoveBandStack2Entry()` — L493
- `.removeReferencesToThisBSEGuid()` — L505
- `.RemoveAllBandStack2Entries()` — L544
- `.AddBandStack2Entry()` — L578
- `.GetBandStack2Entries()` — L610
- `.GetBandStack2Filters()` — L638
- `.checkForPrimaryKeys()` — L732
- `.AddFormTable()` — L742
- `.AddBandTextTable()` — L751
- `.AddRegion2BandText()` — L1074
- `.AddRegionJapanBandText()` — L1401
- `.AddBandTextSWB()` — L1484
- `.ClearBandText()` — L1636
- `.AddRegion1BandText160m()` — L1641
- `.AddRegion1BandText80m()` — L1664
- `.AddRegion1BandText60m()` — L1690
- `.AddRegion1BandText40m()` — L1709
- `.AddRegion1BandText30m()` — L1735
- `.AddRegion1BandText20m()` — L1756
- `.AddRegion1BandText17m()` — L1788
- `.AddRegion1BandText15m()` — L1813
- `.AddRegion1BandText12m()` — L1844
- `.AddRegion1BandText10m()` — L1868
- `.AddRegion1BandText6m()` — L1911
- `.AddRegionIndiaBandText160m()` — L1952
- `.AddRegionIndiaBandText80m()` — L1975
- `.AddRegionIndiaBandText40m()` — L2000
- `.AddRegionIndiaBandText20m()` — L2026
- `.AddRegionIndiaBandText17m()` — L2058
- `.AddRegionIndiaBandText15m()` — L2083
- `.AddRegionIndiaBandText12m()` — L2114
- `.AddRegionIndiaBandText10m()` — L2138
- `.AddRegionIndiaBandText6m()` — L2181
- `.AddRegionIsraelBandText160m()` — L2222
- `.AddRegionIsraelBandText80m()` — L2243
- `.AddRegionIsraelBandText60m()` — L2271
- `.AddRegionIsraelBandText40m()` — L2290
- `.AddRegionIsraelBandText30m()` — L2311
- `.AddRegionIsraelBandText20m()` — L2330
- `.AddRegionIsraelBandText17m()` — L2358
- `.AddRegionIsraelBandText15m()` — L2379
- `.AddRegionIsraelBandText12m()` — L2404
- `.AddRegionIsraelBandText10m()` — L2425
- `.AddRegionIsraelBandText6m()` — L2455
- `.AddRegion1BandText4m()` — L2482
- `.AddRegion1BandTextVHFplus()` — L2517
- `.AddBulgariaBandText160m()` — L2722
- `.AddNetherlandsBandText160m()` — L2746
- `.AddUK_PlusBandText60m()` — L2770
- `.AddNorwayBandText60m()` — L2811
- `.AddSwedenBandText60m()` — L2830
- `.AddHungaryBandText40m()` — L2857
- `.AddRussiaBandText12m()` — L2884
- `.AddRussiaBandText11m()` — L2908
- `.AddEUBandText6m()` — L2926
- `.AddFranceBandText6m()` — L2960
- `.AddLatviaBandText6m()` — L2994
- `.AddBulgariaBandText6m()` — L3028
- `.AddGreeceBandText6m()` — L3062
- `.AddRegion3BandText160m()` — L3538
- `.AddRegion3BandText80m()` — L3559
- `.AddRegion3BandText60m()` — L3583
- `.AddRegion3BandText40m()` — L3602
- `.AddRegion3BandText30m()` — L3625
- `.AddRegion3BandText20m()` — L3644
- `.AddRegion3BandText17m()` — L3673
- `.AddRegion3BandText15m()` — L3698
- `.AddRegion3BandText12m()` — L3726
- `.AddRegion3BandText10m()` — L3749
- `.AddRegion3BandText6m()` — L3777
- `.AddRegion3BandTextVHFplus()` — L3797
- `.AddJapanBandText160m()` — L3915
- `.AddJapanBandText80m()` — L3937
- `.AddJapanBandText40m()` — L3970
- `.AddJapanBandText10m()` — L3995
- `.AddJapanBandText6m()` — L4024
- `.AddJapanBandTextEmergency()` — L4048
- `.AddMemoryTable()` — L4263
- `.AddGroupListTable()` — L4284
- `.AddTXProfileTable()` — L4303
- `.CheckBandTextValid()` — L9430
- `.Init()` — L9471
- `.ConvertFromDBVal()` — L9522
- `.WriteDB()` — L9534
- `.Exit()` — L9564
- `.BandText()` — L9570
- `.PurgeNotches()` — L9819
- `.PurgeMeters()` — L9823
- `.purgeTableEntries()` — L9870
- `.RemoveVarsList()` — L9885
- `.SaveVarsDictionary()` — L9910
- `.SaveVars()` — L9952
- `.GetVarsDictionary()` — L10001
- `.GetVars()` — L10037
- `.ImportAndMergeDatabase()` — L10623
- `.getRadioSelectedFromOldRadButton()` — L11219
- `.ValidateImportedDatabase()` — L11236
- `.ExpandOldTxProfileTable()` — L11269
- `.WriteImportLog()` — L11300
- `.ImportDatabase()` — L11307
- `.UpdateRegion()` — L11332

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/database.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
