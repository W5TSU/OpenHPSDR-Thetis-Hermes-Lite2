# `Console/frmBandStack2.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Per-band frequency stack (last-used frequencies per band) and its popup window.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmBandStack2.Designer.cs` (references ×1)
  - `Console/frmFilterManager.Designer.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/clsBandStackManager.cs` (calls ×4, references ×3)
  - `Console/common.cs` (calls ×4)
  - `Console/Andromeda/Andromeda.cs` (references ×2)

## Outline

### Types

#### `frmBandStack2` (type, L56)

- `.InitForm()` — L94
- `.OnMox()` — L118
- `.RemoveDelegates()` — L122
- `.InitBandStackFilter()` — L129
- `.UpdateSelected()` — L180
- `.setupSelectedButtons()` — L206
- `.setupRadioButtons()` — L233
- `.btnOptions_Click()` — L267
- `.radioLastUsedEntry_CheckedChanged()` — L281
- `.radioSpecific_CheckedChanged()` — L290
- `.radioLastUsed_CheckedChanged()` — L299
- `.btnSetSpecific_Click()` — L308
- `.btnLockSelected_Click()` — L322
- `.btnDeleteSelected_Click()` — L340
- `.bandStackListBox_SelectedIndexChanged()` — L351
- `.btnAddStackEntry_Click()` — L374
- `.chkAlwaysOnTop_CheckedChanged()` — L379
- `.HideClose()` — L384
- `.frmBandStack2_FormClosing()` — L396
- `.Show()` — L402
- `.Store()` — L424
- `.btnUpdateEntry_Click()` — L430
- `.chkIgnoreDuplicates_CheckedChanged()` — L441
- `.chkHideOnSelect_CheckedChanged()` — L446
- `.chkShowInSpectrum_CheckedChanged()` — L451
- `.frmBandStack2_LocationChanged()` — L463

#### `BandStackListBox` (type, L471)

- `.AddItem()` — L537
- `.ClearItems()` — L547
- `.OnFontChanged()` — L555
- `.OnSelectedIndexChanged()` — L565
- `.OnDrawItem()` — L571
- `.OnMouseEnter()` — L642
- `.OnMouseLeave()` — L646
- `.OnMouseMove()` — L653
- `.OnPaint()` — L669

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmBandStack2.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
