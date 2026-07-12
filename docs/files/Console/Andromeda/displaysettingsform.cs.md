# `Console/Andromeda/displaysettingsform.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/comboboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

### Types

#### `DisplaySettingsForm` (type, L40)

- `.Dispose()` — L78
- `.InitializeComponent()` — L96
- `.RepopulateForm()` — L324
- `.BtnClose_Click()` — L373
- `.DisplaySettingsForm_FormClosing()` — L378
- `.DisplaySettingsForm_Activated()` — L386
- `.ComboRX1Meter_SelectedIndexChanged()` — L392
- `.ComboRX2Meter_SelectedIndexChanged()` — L397
- `.ComboTXMeter_SelectedIndexChanged()` — L402
- `.ComboRX1Display_SelectedIndexChanged()` — L407
- `.ComboRX2Display_SelectedIndexChanged()` — L412
- `.ChkRX1Avg_CheckedChanged()` — L417
- `.ChkRX2Avg_CheckedChanged()` — L425
- `.ChkRX1Peak_CheckedChanged()` — L433
- `.ChkRX2Peak_CheckedChanged()` — L441

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/displaysettingsform.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
