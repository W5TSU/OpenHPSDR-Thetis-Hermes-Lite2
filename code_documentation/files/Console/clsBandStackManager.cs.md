# `Console/clsBandStackManager.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Per-band frequency stack (last-used frequencies per band) and its popup window.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×59, references ×10)
  - `Console/MeterManager.cs` (calls ×8)
  - `Console/database.cs` (references ×5, calls ×3)
  - `Console/frmBandStack2.cs` (calls ×4, references ×3)
  - `Console/setup.cs` (calls ×4)
  - `Console/CAT/CATCommands.cs` (calls ×1)
  - `Console/display.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×24)
  - `Console/database.cs` (calls ×7)
- Most-referenced symbols from other files: `.GetFilter()` (×17), `.BandToString()` (×10), `.GenerateFilteredList()` (×6), `.StringToBand()` (×6), `.BandToColour()` (×4), `.UpdateCurrentWithLastVisitedData()` (×4), `.IndexFromGUID()` (×3), `.SelectInitial()` (×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.Copy()`** — L99 — `public BandFrequencyData Copy()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Copy()`** — L178 — `public BandStackEntry Copy(bool bNewGUID = false)`
  Called by: `.UpdateEntry()` (same file), `.UpdateCurrentWithLastVisitedData()` (same file), `.SelectInitial()` (same file), `.Current()` (same file), `.FindEntriesForFrequency()` (same file), `.initLists()` (same file)
- **`.Copy()`** — L249 — `public BandStackFilter Copy()`
  Called by: `.UpdateEntry()` (same file), `.EntryByIndex()` (same file), `.SelectInitial()` (same file), `.First()` (same file), `.Current()` (same file), `.Next()` (same file) — and 2 more

### Types

#### `BandType` (type, L50)

_No extracted members._

#### `DSPSubMode` (type, L60)

_No extracted members._

#### `BandFrequencyData` (type, L80)

_No extracted members._

#### `BandStackEntry` (type, L114)

- **`.CompareTo()`** — L148 — `public int CompareTo(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `BandStackFilter` (type, L201)

- **`.IndexFromGUID()`** — L322 — `public int IndexFromGUID(string sGUID)`
  Called by: `.OnEntryAdd()` (`Console/console.cs`), `.OnEntryUpdate()` (`Console/console.cs`), `.OnEntryClicked()` (`Console/console.cs`)
- **`.UpdateEntry()`** — L339 — `public bool UpdateEntry(BandStackEntry bse)`
  Updates entry.
  Called by: `.btnLockSelected_Click()` (`Console/frmBandStack2.cs`)
- **`.UpdateCurrentWithLastVisitedData()`** — L366 — `public bool UpdateCurrentWithLastVisitedData(bool bCheckForFreqDupe = false)`
  Updates current with last visited data.
  Called by: `.Console_Closing()` (`Console/console.cs`), `.OnEntryUpdate()` (`Console/console.cs`), `.OnEntryClicked()` (`Console/console.cs`), `.preBandSelect()` (`Console/console.cs`)
- **`.Remove()`** — L418 — `public void Remove(int index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveCurrent()`** — L456 — `public void RemoveCurrent()`
  Removes current.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EntryByIndex()`** — L465 — `public BandStackEntry EntryByIndex(int index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectInitial()`** — L472 — `public BandStackEntry SelectInitial()`
  Selects initial.
  Called by: `.internalAddFilter()` (same file), `.OnEntryDelete()` (`Console/console.cs`), `.preBandSelect()` (`Console/console.cs`), `.OnBandChangeHandler()` (`Console/console.cs`)
- **`.First()`** — L520 — `public BandStackEntry First()`
  Called by: `.CATRX2BandUpDown()` (`Console/console.cs`), `.SetupRX2Band()` (`Console/console.cs`), `.mnuBandRX2_Click()` (`Console/console.cs`)
- **`.Current()`** — L531 — `public BandStackEntry Current()`
  Called by: `.Console_KeyDown()` (`Console/console.cs`), `.OnEntryDelete()` (`Console/console.cs`), `.UpdateSelected()` (`Console/frmBandStack2.cs`)
- **`.Next()`** — L551 — `public BandStackEntry Next()`
  Called by: `.Console_KeyDown()` (`Console/console.cs`), `.preBandSelect()` (`Console/console.cs`)
- **`.Previous()`** — L570 — `public BandStackEntry Previous()`
  Called by: `.Console_KeyDown()` (`Console/console.cs`), `.preBandSelect()` (`Console/console.cs`)
- **`.FindEntriesForFrequency()`** — L590 — `public List<BandStackEntry> FindEntriesForFrequency(double frequency)`
  Finds entries for frequency.
  Called by: `.UpdateCurrentWithLastVisitedData()` (same file), `.OnEntryAdd()` (`Console/console.cs`)
- **`.FindForFrequencyRange()`** — L610 — `public List<BandStackEntry> FindForFrequencyRange(double frequencyLow, double frequencyHigh)`
  Finds for frequency range.
  Called by: `.updateBandstackOverlay()` (`Console/console.cs`)
- **`.GenerateFilteredList()`** — L622 — `public void GenerateFilteredList(bool bMaintainSelected, bool bInitalising = false)`
  Called by: `.internalAddFilter()` (same file), `.Console_Closing()` (`Console/console.cs`), `.OnEntryAdd()` (`Console/console.cs`), `.OnEntryUpdate()` (`Console/console.cs`), `.OnEntryDelete()` (`Console/console.cs`), `.OnEntryClicked()` (`Console/console.cs`) — and 1 more

#### `FilterReturnMode` (type, L203)

_No extracted members._

#### `BandStackManager` (type, L744)

- **`.SaveToDB()`** — L773 — `public static void SaveToDB()`
  Saves to db.
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.RegionReset()`** — L784 — `public static void RegionReset()`
  Called by: `.comboFRSRegion_SelectedIndexChanged()` (`Console/setup.cs`), `.ChkExtended_CheckedChanged()` (`Console/setup.cs`)
- **`.initLists()`** — L790 — `private static void initLists()`
  Called by: `.RegionReset()` (same file)
- **`.IndexFromGUID()`** — L833 — `public static int IndexFromGUID(string sGUID)`
  Called by: `.UpdateEntry()` (same file), `.UpdateCurrentWithLastVisitedData()` (same file), `.SelectInitial()` (same file), `.GenerateFilteredList()` (same file), `.DeleteEntry()` (same file)
- **`.addStandardFilters()`** — L849 — `private static void addStandardFilters()`
  Called by: `.initLists()` (same file)
- **`.GetFilter()`** — L876 — `public static BandStackFilter GetFilter(string sFilterName, bool bIncludeUserDefined = true)`
  Returns filter.
  Called by: `.CATRX2BandUpDown()` (`Console/console.cs`), `.Console_KeyDown()` (`Console/console.cs`), `.Console_Closing()` (`Console/console.cs`), `.SetupRX2Band()` (`Console/console.cs`), `.mnuBandRX2_Click()` (`Console/console.cs`), `.handleBSFChange()` (`Console/console.cs`) — and 11 more
- **`.GetFilters()`** — L890 — `public static List<BandStackFilter> GetFilters(Band b, bool onlyFirst = false, bool bIncludeUserDefined = true)`
  Returns filters.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoesFilterNameExist()`** — L912 — `public static bool DoesFilterNameExist(string sFilterName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.internalAddFilter()`** — L916 — `private static bool internalAddFilter(BandStackFilter bsf, bool bInitalising = false)`
  Called by: `.initLists()` (same file), `.addStandardFilters()` (same file), `.AddFilter()` (same file)
- **`.AddFilter()`** — L927 — `public static bool AddFilter(BandStackFilter bsf)`
  Adds filter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddEntry()`** — L931 — `public static void AddEntry(BandStackEntry bse)`
  Adds entry.
  Called by: `.OnEntryAdd()` (`Console/console.cs`)
- **`.DeleteEntry()`** — L940 — `public static bool DeleteEntry(BandStackEntry bse)`
  Deletes entry.
  Called by: `.OnEntryDelete()` (`Console/console.cs`)
- **`.GetBandFrequencyDataForFrequency()`** — L996 — `public static List<BandFrequencyData> GetBandFrequencyDataForFrequency(double frequency, bool extended, FRSRegion region, Band band = Band.LAST)`
  Returns band frequency data for frequency.
  Called by: `.IsFrequencyInBandType()` (same file), `.IsOKToTX()` (same file), `.BandByFreq()` (`Console/console.cs`)
- **`.GetFrequencyRangesForBand()`** — L1034 — `public static List<BandFrequencyData> GetFrequencyRangesForBand(Band band, bool extended, FRSRegion region)`
  Returns frequency ranges for band.
  Called by: `.addStandardFilters()` (same file), `.zoomToBandBandwidth()` (`Console/console.cs`), `.preBandSelect()` (`Console/console.cs`)
- **`.IsFrequencyInBandType()`** — L1048 — `public static bool IsFrequencyInBandType(double frequency, BandType bandType, bool extended, FRSRegion region)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsOKToTX()`** — L1063 — `public static bool IsOKToTX(double frequency, bool extended, FRSRegion region)`
  Called by: `.checkValidTXFreq_local()` (`Console/console.cs`)
- **`.addStandardFrequencies()`** — L1084 — `private static void addStandardFrequencies()`
  Called by: `.initLists()` (same file)
- **`.bandToBandType()`** — L1159 — `private static BandType bandToBandType(Band b)`
  Called by: `.GetBandFrequencyDataForFrequency()` (same file)
- **`.GetNearestBandForFrequency()`** — L1210 — `public static Band GetNearestBandForFrequency(double freq, bool ignoreGen, bool ignoreWWV)`
  Returns nearest band for frequency.
  Called by: `.getTXBandWhenExtended()` (`Console/console.cs`)
- **`.GetBandTypeForFrequency()`** — L1262 — `public static BandType GetBandTypeForFrequency(double frequency)`
  Returns band type for frequency.
  Called by: `.OnBandChangeHandler()` (`Console/console.cs`)
- **`.frequencyData()`** — L1274 — `private static List<BandFrequencyData> frequencyData(FRSRegion region)`
  Called by: `.GetBandFrequencyDataForFrequency()` (same file), `.GetFrequencyRangesForBand()` (same file)
- **`.ModeToString()`** — L1766 — `public static string ModeToString(DSPMode mode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToMode()`** — L1770 — `public static DSPMode StringToMode(string mode)`
  Called by: `.addBSObjectToEntries()` (same file), `.ImportAndMergeDatabase()` (`Console/database.cs`)
- **`.FilterToString()`** — L1776 — `public static string FilterToString(Filter filter)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToFilter()`** — L1780 — `public static Filter StringToFilter(string filter)`
  Called by: `.addBSObjectToEntries()` (same file), `.ImportAndMergeDatabase()` (`Console/database.cs`)
- **`.BandToColour()`** — L1786 — `public static Color BandToColour(Band b)`
  Called by: `.VHFUpdate()` (`Console/MeterManager.cs`), `.setupButtons()` (`Console/MeterManager.cs`), `.drawBand()` (`Console/MeterManager.cs`), `.InitBandStackFilter()` (`Console/frmBandStack2.cs`)
- **`.BandToString()`** — L1852 — `public static string BandToString(Band b)`
  Called by: `.ZZJR()` (`Console/CAT/CATCommands.cs`), `.GetReading()` (`Console/MeterManager.cs`), `.handleClicked()` (`Console/MeterManager.cs`), `.setupButtons()` (`Console/MeterManager.cs`), `.drawBand()` (`Console/MeterManager.cs`), `.renderVfoDisplay()` (`Console/MeterManager.cs`) — and 4 more
- **`.StringToBand()`** — L1915 — `public static Band StringToBand(string s)`
  Called by: `.addBSObjectToEntries()` (same file), `.CATRX2BandUpDown()` (`Console/console.cs`), `.radBandVHF_Click()` (`Console/console.cs`), `.SetupRX2Band()` (`Console/console.cs`), `.mnuBandRX2_Click()` (`Console/console.cs`), `.ImportAndMergeDatabase()` (`Console/database.cs`) — and 1 more
- **`.addBSObjectToEntries()`** — L1981 — `private static void addBSObjectToEntries(object[] o, bool bIgnore60m = false)`
  Called by: `.AddRegion1BandStack()` (same file), `.AddRegion2BandStack()` (same file), `.AddRegion3BandStack()` (same file), `.AddBandStackSWL()` (same file), `.AddUK_PlusBandStack()` (same file), `.AddUS_PlusBandStack()` (same file) — and 2 more
- **`.isGuidInList()`** — L2013 — `private static bool isGuidInList(string sGUID)`
  Called by: `.addBSObjectToEntries()` (same file)
- **`.AddRegion1BandStack()`** — L2022 — `private static void AddRegion1BandStack()`
  Adds region1 band stack.
  Called by: `.addStandardFrequencies()` (same file)
- **`.AddRegion2BandStack()`** — L2099 — `private static void AddRegion2BandStack(bool bIgnore60m = false)`
  Adds region2 band stack.
  Called by: `.addStandardFrequencies()` (same file)
- **`.AddRegion3BandStack()`** — L2178 — `private static void AddRegion3BandStack()`
  Adds region3 band stack.
  Called by: `.addStandardFrequencies()` (same file)
- **`.AddBandStackSWL()`** — L2236 — `private static void AddBandStackSWL()`
  Adds band stack swl.
  Called by: `.addStandardFrequencies()` (same file)
- **`.AddUK_PlusBandStack()`** — L2315 — `private static void AddUK_PlusBandStack()`
  Adds uk plus band stack.
  Called by: `.addStandardFrequencies()` (same file)
- **`.AddUS_PlusBandStack()`** — L2400 — `private static void AddUS_PlusBandStack()`
  Adds us plus band stack.
  Called by: `.addStandardFrequencies()` (same file)
- **`.AddSwedenBandStack()`** — L2412 — `private static void AddSwedenBandStack()`
  Adds sweden band stack.
  Called by: `.addStandardFrequencies()` (same file)
- **`.AddRegionJapanBandStack()`** — L2492 — `private static void AddRegionJapanBandStack()`
  Adds region japan band stack.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsBandStackManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
