# `Console/Andromeda/SliderSettingsForm.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/trackbarts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

### Types

#### `SliderSettingsForm` (type, L35)

- `.Dispose()` — L114
- `.InitializeComponent()` — L131
- `.SliderSettingsForm_Closing()` — L1240
- `.SliderSettingsForm_Activated()` — L1251
- `.tbRX1AF_Scroll()` — L1309
- `.tbRX2AF_Scroll()` — L1314
- `.tbSubRXAF_Scroll()` — L1319
- `.tbRX1RF_Scroll()` — L1324
- `.tbRX2RF_Scroll()` — L1330
- `.tbRX1Sql_Scroll()` — L1336
- `.tbRX2Sql_Scroll()` — L1341
- `.tbRX1Pan_Scroll()` — L1346
- `.tbRX2Pan_Scroll()` — L1351
- `.tbSubRXPan_Scroll()` — L1356
- `.tbMasterAF_Scroll()` — L1361
- `.tbDrive_Scroll()` — L1366
- `.chkSubRX_CheckedChanged()` — L1371
- `.btnClose_Click()` — L1379
- `.SliderSettingsForm_FormClosing()` — L1384
- `.tbRX1Atten_Scroll()` — L1389
- `.tbRX2Atten_Scroll()` — L1394
- `.chkRX1Mute_CheckedChanged()` — L1399
- `.chkRX2Mute_CheckedChanged()` — L1404
- `.ChkRX1VAC_CheckedChanged()` — L1409
- `.ChkRX2VAC_CheckedChanged()` — L1415
- `.TbMicGain_Scroll()` — L1421
- `.TbRX1VACRX_Scroll()` — L1426
- `.TbRX1VACTX_Scroll()` — L1432
- `.TbRX2VACRX_Scroll()` — L1438
- `.TbRX2VACTX_Scroll()` — L1444
- `.FormEncoderEvent()` — L1453
- `.Callback()` — L1472
- `.chkRX1Sql_CheckStateChanged()` — L1493
- `.chkRX2Sql_CheckStateChanged()` — L1509
- `.updateSQLButtons()` — L1525
- `.getSQLinfoOnFormActivate()` — L1542

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/SliderSettingsForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
