# `Console/SpotManager2.cs`

**Functional area:** [16. DX spots and cluster display](../../CODE_OUTLINE.md#16-dx-spots-and-cluster-display)

**Role:** DX spot store and on-panadapter spot rendering (callsigns on the spectrum); fed by TCI and cluster sources.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×5)
  - `Console/TCIServer.cs` (calls ×4)
  - `Console/console.cs` (calls ×2)
  - `Console/display.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×3)
  - `Console/common.cs` (calls ×3)
  - `Console/clsFlagAtlas.cs` (calls ×1)
  - `Console/clsCountryData.cs` (calls ×1)
- Most-referenced symbols from other files: `.ClearAllSpots()` (×3), `.OwnCallApearance()` (×3), `.FreeUpFlags()` (×1), `.HighlightSpot()` (×1), `.FilterForRawMode()` (×1), `.AddSpot()` (×1), `.GetFrequencySortedSpots()` (×1), `.DeleteSpot()` (×1)

## Outline

### Types

#### `SpotManager2` (type, L54)

- `.FreeUpFlags()` — L171
- `.compareByFrequency()` — L183
- `.markSortedSpotsDirty()` — L188
- `.RebuildSortedSpotsCache()` — L193
- `.clearHighlightedReference()` — L210
- `.pruneHighlightedReferences()` — L223
- `.onTick()` — L245
- `.HighlightSpot()` — L262
- `.SpotModeNumberToDSPMode()` — L302
- `.FilterForRawMode()` — L349
- `.getSpotTextColour()` — L368
- `.getFlagImage()` — L429
- `.getFlagImageFromCallsign()` — L458
- `.AddSpot()` — L474
- `.GetFrequencySortedSpots()` — L676
- `.ClearAllSpots()` — L685
- `.DeleteSpot()` — L706
- `.OwnCallApearance()` — L730

#### `JsonSpotData` (type, L69)

_No extracted members._

#### `smSpot` (type, L86)

- `.BrowseQRZ()` — L149
- `.BrowseHamQTH()` — L153

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/SpotManager2.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
