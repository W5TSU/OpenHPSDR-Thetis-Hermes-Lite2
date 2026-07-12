# `Console/ucQuickRecall.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Quick recall (recent frequencies) list.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×7)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/frmQuickRecallPopupList.Designer.cs` (references ×1)

## Outline

### Types

#### `ucQuickRecall` (type, L49)

- `.resizeAndReposition()` — L142
- `.OnPopupClosed()` — L153
- `.OnMoxChanged()` — L158
- `.OnModeChanged()` — L174
- `.OnVFOAChange()` — L192
- `.buttonClicked()` — L204
- `.formatFrequencyToString()` — L210
- `.findExistingInVFOA()` — L221
- `.addVFOAEntry()` — L230
- `.OnBackgroundColourPingerTick()` — L238
- `.OnVFOAModeTick()` — L244
- `.OnVFOATick()` — L267
- `.OnDispose()` — L299
- `.OnEntrySelected()` — L310
- `.btnPrevious_Click()` — L318
- `.btnNext_Click()` — L329
- `.selectVFOAEntry()` — L340
- `.btnList_Click()` — L348
- `.buildAndShowPopup()` — L361
- `.ucQuickRecall_BackColorChanged()` — L378
- `.setButtonBackColour()` — L383
- `.ucQuickRecall_Resize()` — L390
- `.lblFlashColour_Click()` — L395

#### `QuickInfo` (type, L64)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucQuickRecall.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
