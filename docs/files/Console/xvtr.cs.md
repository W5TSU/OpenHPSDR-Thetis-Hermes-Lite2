# `Console/xvtr.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Transverter band setup (frequency offsets, power limits per transverter band).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×40, references ×1)
  - `Console/Andromeda/Andromeda.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/comboboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/numericupdownts.cs` (references ×1)
  - `Console/Invoke/textboxts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.TranslateFreq()` (×15), `.XVTRFreq()` (×7), `.GetBegin()` (×4), `.GetEnd()` (×4), `.GetEnabled()` (×3), `.GetRXOnly()` (×3), `.GetDisablePA()` (×2), `.GetRXGain()` (×2)

## Outline

### Types

#### `XVTRForm` (type, L51)

- `.Dispose()` — L334
- `.InitializeComponent()` — L353
- `.XVTRFreq()` — L5335
- `.TranslateFreq()` — L5356
- `.SetupControlArrays()` — L5374
- `.GetRXAntenna()` — L5581
- `.GetEnabled()` — L5588
- `.GetBegin()` — L5593
- `.GetEnd()` — L5598
- `.GetPower()` — L5603
- `.SetPower()` — L5608
- `.GetRXOnly()` — L5613
- `.SetRXOnly()` — L5618
- `.GetRXGain()` — L5623
- `.GetXVTRRF()` — L5635
- `.GetDisablePA()` — L5640
- `.XVTRForm_Closing()` — L5650
- `.chkEnable0_CheckedChanged()` — L5659
- `.chkEnable1_CheckedChanged()` — L5677
- `.chkEnable2_CheckedChanged()` — L5695
- `.chkEnable3_CheckedChanged()` — L5714
- `.chkEnable4_CheckedChanged()` — L5733
- `.chkEnable5_CheckedChanged()` — L5752
- `.chkEnable6_CheckedChanged()` — L5771
- `.chkEnable7_CheckedChanged()` — L5790
- `.chkEnable8_CheckedChanged()` — L5809
- `.chkEnable9_CheckedChanged()` — L5828
- `.chkEnable10_CheckedChanged()` — L5847
- `.chkEnable11_CheckedChanged()` — L5866
- `.chkEnable12_CheckedChanged()` — L5885
- `.chkEnable13_CheckedChanged()` — L5904
- `.chkEnable14_CheckedChanged()` — L5923
- `.chkEnable15_CheckedChanged()` — L5940
- `.txtButtonText_TextChanged()` — L5959
- `.chkUseXVTRTUNPWR_CheckedChanged()` — L5967
- `.chkXVTRRF_CheckedChanged()` — L5972
- `.udRXGain_ValueChanged()` — L5999
- `.chkAlexTRRelay_CheckedChanged()` — L6015
- `.XVTRForm_Paint()` — L6033
- `.udPower_ValueChanged()` — L6038
- `.comboAnt0_SelectedIndexChanged()` — L6044
- `.XVTRForm_Activated()` — L6056

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/xvtr.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
