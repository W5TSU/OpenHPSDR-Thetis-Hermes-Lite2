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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `SpotManager2` (type, L54)

- **`.FreeUpFlags()`** — L171 — `public static void FreeUpFlags()`
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.compareByFrequency()`** — L183 — `private static int compareByFrequency(smSpot left, smSpot right)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.markSortedSpotsDirty()`** — L188 — `private static void markSortedSpotsDirty()`
  Called by: `.onTick()` (same file), `.AddSpot()` (same file), `.ClearAllSpots()` (same file), `.DeleteSpot()` (same file)
- **`.RebuildSortedSpotsCache()`** — L193 — `private static void RebuildSortedSpotsCache()`
  Called by: `.GetFrequencySortedSpots()` (same file)
- **`.clearHighlightedReference()`** — L210 — `private static void clearHighlightedReference(smSpot spot)`
  Called by: `.AddSpot()` (same file), `.ClearAllSpots()` (same file), `.DeleteSpot()` (same file)
- **`.pruneHighlightedReferences()`** — L223 — `private static void pruneHighlightedReferences()`
  Called by: `.onTick()` (same file)
- **`.onTick()`** — L245 — `private static void onTick(Object source, ElapsedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HighlightSpot()`** — L262 — `public static smSpot HighlightSpot(int x, int y)`
  Called by: `.pnlDisplay_MouseMove()` (`Console/console.cs`)
- **`.SpotModeNumberToDSPMode()`** — L302 — `public static DSPMode SpotModeNumberToDSPMode(int number, double freq = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterForRawMode()`** — L349 — `public static string FilterForRawMode(string raw_mode)`
  Called by: `.handleSpot()` (`Console/TCIServer.cs`)
- **`.getSpotTextColour()`** — L368 — `private static Color getSpotTextColour(string text_colour)`
  Returns spot text colour.
  Called by: `.AddSpot()` (same file)
- **`.getFlagImage()`** — L429 — `private static Image getFlagImage(string flag)`
  Returns flag image.
  Called by: `.getFlagImageFromCallsign()` (same file), `.AddSpot()` (same file)
- **`.getFlagImageFromCallsign()`** — L458 — `private static Image getFlagImageFromCallsign(string callsign, out string country)`
  Returns flag image from callsign.
  Called by: `.AddSpot()` (same file)
- **`.AddSpot()`** — L474 — `public static void AddSpot(string callsign, DSPMode mode, long frequencyHz, Color colour, string additionalText, JsonSpotData jsonSpotData = null)`
  Adds spot.
  Called by: `.handleSpot()` (`Console/TCIServer.cs`)
- **`.GetFrequencySortedSpots()`** — L676 — `public static smSpot[] GetFrequencySortedSpots()`
  Returns frequency sorted spots.
  Called by: `.drawSpots()` (`Console/display.cs`)
- **`.ClearAllSpots()`** — L685 — `public static void ClearAllSpots(bool non_swl, bool swl)`
  Clears all spots.
  Called by: `.handleSpotClear()` (`Console/TCIServer.cs`), `.btnClearTCISpots_Click()` (`Console/setup.cs`), `.btnClearTCISpotsSWL_Click()` (`Console/setup.cs`)
- **`.DeleteSpot()`** — L706 — `public static void DeleteSpot(string callsign)`
  Deletes spot.
  Called by: `.handleDeleteSpot()` (`Console/TCIServer.cs`)
- **`.OwnCallApearance()`** — L730 — `public static void OwnCallApearance(bool bEnabled, string sCall, Color replacementColorBackground)`
  Called by: `.chkSpotOwnCallAppearance_CheckedChanged()` (`Console/setup.cs`), `.txtOwnCallsign_TextChanged()` (`Console/setup.cs`), `.clrbtnOwnCallApearance_Changed()` (`Console/setup.cs`)

#### `JsonSpotData` (type, L69)

_No extracted members._

#### `smSpot` (type, L86)

- **`.BrowseQRZ()`** — L149 — `public void BrowseQRZ()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BrowseHamQTH()`** — L153 — `public void BrowseHamQTH()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/SpotManager2.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
