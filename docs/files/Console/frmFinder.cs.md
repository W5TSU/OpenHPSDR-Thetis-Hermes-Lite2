# `Console/frmFinder.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Simple text-input dialog; searchable "find a setting" helper.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)

## Outline

### Types

#### `frmFinder` (type, L55)

- `.GatherSearchData()` — L98
- `.GatherCATStructData()` — L119
- `.gatherCATStructSearchDataThread()` — L148
- `.gatherSearchDataThread()` — L211
- `.stripPrefix()` — L236
- `.getControlList()` — L247
- `.txtSearch_TextChanged()` — L316
- `.lstResults_SelectedIndexChanged()` — L372
- `.lstResults_DrawItem()` — L395
- `.highlight()` — L468
- `.findSubstringOccurrences()` — L492
- `.applyTint()` — L516
- `.frmFinder_FormClosing()` — L525
- `.Show()` — L536
- `.lstResults_MeasureItem()` — L544
- `.showControl()` — L563
- `.selectRequiredTabs()` — L582
- `.chkFullDetails_CheckedChanged()` — L653
- `.ReadXmlFinderFile()` — L665
- `.WriteXmlFinderFile()` — L693
- `.frmFinder_KeyDown()` — L746
- `.chkHighlight_CheckedChanged()` — L777
- `.chkKeywords_CheckedChanged()` — L789

#### `SearchData` (type, L57)

_No extracted members._

#### `CatStructEntry` (type, L143)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmFinder.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
