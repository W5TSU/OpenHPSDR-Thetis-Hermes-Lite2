# `Console/FilterForm.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX filter preset model per mode, the filter-edit form, and the filter-set manager.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/filter.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

### Types

#### `FilterForm` (type, L52)

- `.Dispose()` — L116
- `.InitializeComponent()` — L135
- `.GetFilterInfo()` — L646
- `.HzToPixel()` — L658
- `.PixelToHz()` — L666
- `.UpdateFilter()` — L675
- `.radFilter_CheckedChanged()` — L705
- `.comboDSPMode_SelectedIndexChanged()` — L722
- `.txtName_LostFocus()` — L727
- `.udLow_ValueChanged()` — L784
- `.udHigh_ValueChanged()` — L806
- `.udLow_LostFocus()` — L828
- `.udHigh_LostFocus()` — L833
- `.picDisplay_Paint()` — L838
- `.picDisplay_MouseMove()` — L855
- `.picDisplay_MouseDown()` — L884
- `.picDisplay_MouseUp()` — L905
- `.udWidth_ValueChanged()` — L918
- `.FilterForm_FormClosing()` — L963

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/FilterForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
