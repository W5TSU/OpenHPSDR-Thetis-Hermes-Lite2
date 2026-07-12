# `Console/ucOtherButtonsOptionsGrid.cs`

**Functional area:** [13. Andromeda control surface](../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** User-programmable macro buttons and their configuration grid.

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (calls ×11, references ×8)
  - `Console/console.cs` (references ×6)
  - `Console/frmMacroButtonConfig.cs` (references ×2)
- Uses (outgoing references to other files):
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
- Most-referenced symbols from other files: `.BitFromID()` (×7), `.BitToID()` (×2), `.BitToIcon()` (×1), `.BitToText()` (×1)

## Outline

### Types

#### `OtherButtonId` (type, L54)

_No extracted members._

#### `OtherButtonMacroSettings` (type, L259)

- `.deep_clone()` — L306

#### `OB_ButtonState` (type, L262)

_No extracted members._

#### `OtherButtonIdHelpers` (type, L521)

- `.OtherButtonIDToText()` — L776
- `.OtherButtonIDToIconOn()` — L781
- `.OtherButtonIDToIconOff()` — L786
- `.BitToID()` — L793
- `.BitFromID()` — L799
- `.BitToText()` — L805
- `.BitToIcon()` — L810
- `.OtherButtonIDToTooltip()` — L815
- `.checkImplemented()` — L820

#### `ucOtherButtonsOptionsGrid` (type, L842)

- `.initialise_checkboxes()` — L913
- `.checkbox_checked_changed()` — L1103
- `.button_clicked()` — L1108
- `.GetBitfield()` — L1119
- `.SetBitfield()` — L1131
- `.GetCheckedCount()` — L1147
- `.GetMacroSettings()` — L1159
- `.SetMacroSettings()` — L1165

#### `MacroButtonEventArgs` (type, L844)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucOtherButtonsOptionsGrid.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
