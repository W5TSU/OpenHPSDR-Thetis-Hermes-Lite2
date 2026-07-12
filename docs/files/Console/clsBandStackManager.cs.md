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

### Functions

- `.Copy()` — L99
- `.Copy()` — L178
- `.Copy()` — L249

### Types

#### `BandType` (type, L50)

_No extracted members._

#### `DSPSubMode` (type, L60)

_No extracted members._

#### `BandFrequencyData` (type, L80)

_No extracted members._

#### `BandStackEntry` (type, L114)

- `.CompareTo()` — L148

#### `BandStackFilter` (type, L201)

- `.IndexFromGUID()` — L322
- `.UpdateEntry()` — L339
- `.UpdateCurrentWithLastVisitedData()` — L366
- `.Remove()` — L418
- `.RemoveCurrent()` — L456
- `.EntryByIndex()` — L465
- `.SelectInitial()` — L472
- `.First()` — L520
- `.Current()` — L531
- `.Next()` — L551
- `.Previous()` — L570
- `.FindEntriesForFrequency()` — L590
- `.FindForFrequencyRange()` — L610
- `.GenerateFilteredList()` — L622

#### `FilterReturnMode` (type, L203)

_No extracted members._

#### `BandStackManager` (type, L744)

- `.SaveToDB()` — L773
- `.RegionReset()` — L784
- `.initLists()` — L790
- `.IndexFromGUID()` — L833
- `.addStandardFilters()` — L849
- `.GetFilter()` — L876
- `.GetFilters()` — L890
- `.DoesFilterNameExist()` — L912
- `.internalAddFilter()` — L916
- `.AddFilter()` — L927
- `.AddEntry()` — L931
- `.DeleteEntry()` — L940
- `.GetBandFrequencyDataForFrequency()` — L996
- `.GetFrequencyRangesForBand()` — L1034
- `.IsFrequencyInBandType()` — L1048
- `.IsOKToTX()` — L1063
- `.addStandardFrequencies()` — L1084
- `.bandToBandType()` — L1159
- `.GetNearestBandForFrequency()` — L1210
- `.GetBandTypeForFrequency()` — L1262
- `.frequencyData()` — L1274
- `.ModeToString()` — L1766
- `.StringToMode()` — L1770
- `.FilterToString()` — L1776
- `.StringToFilter()` — L1780
- `.BandToColour()` — L1786
- `.BandToString()` — L1852
- `.StringToBand()` — L1915
- `.addBSObjectToEntries()` — L1981
- `.isGuidInList()` — L2013
- `.AddRegion1BandStack()` — L2022
- `.AddRegion2BandStack()` — L2099
- `.AddRegion3BandStack()` — L2178
- `.AddBandStackSWL()` — L2236
- `.AddUK_PlusBandStack()` — L2315
- `.AddUS_PlusBandStack()` — L2400
- `.AddSwedenBandStack()` — L2412
- `.AddRegionJapanBandStack()` — L2492

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsBandStackManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
