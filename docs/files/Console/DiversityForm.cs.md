# `Console/DiversityForm.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Two-receiver diversity reception control (phase/gain mixing of RX1/RX2).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/numericupdownts.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/Invoke/textboxts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

### Types

#### `DiversityForm` (type, L59)

- `.applyControlStyles()` — L265
- `.Dispose()` — L359
- `.InitializeComponent()` — L377
- `.SerializeObjectToString()` — L1458
- `.DeserializeStringToObject()` — L1470
- `.initMemories()` — L1481
- `.picRadar_Paint()` — L1489
- `.getControlHandlePoint()` — L1603
- `.PolarToXY()` — L1629
- `.picRadar_MouseMove()` — L1636
- `.picRadar_MouseDown()` — L1730
- `.picRadar_MouseUp()` — L1757
- `.udR_ValueChanged()` — L1777
- `.udTheta_ValueChanged()` — L1783
- `.UpdateDirection()` — L1789
- `.UpdateDiversity()` — L1834
- `.chkEnable_CheckedChanged()` — L1898
- `.DiversityForm_Closing()` — L1911
- `.btnShiftUp45_Click()` — L1918
- `.btnShift180_Click()` — L1955
- `.btnShiftDwn45_Click()` — L1967
- `.radioButtonMerc1_CheckedChanged()` — L2004
- `.radioButtonMerc2_CheckedChanged()` — L2028
- `.chkLockAngle_CheckedChanged()` — L2048
- `.chkLockR_CheckedChanged()` — L2054
- `.groupBox_refMerc_Enter()` — L2060
- `.groupBox_udPhase_Enter()` — L2064
- `.udR2_ValueChanged()` — L2070
- `.udR1_ValueChanged()` — L2140
- `.udCalib_ValueChanged()` — L2341
- `.udAngle0_ValueChanged()` — L2348
- `.ConvertAngleToAngle0()` — L2360
- `.ConvertAngle0ToAngle()` — L2370
- `.panelDivControls_Enter()` — L2385
- `.udAntSpacing_ValueChanged_1()` — L2390
- `.CalcVrms()` — L2398
- `.chkCrossFire_CheckedChanged()` — L2442
- `.udFineNull_ValueChanged()` — L2453
- `.radRxSource1_CheckedChanged()` — L2534
- `.radRxSource2_CheckedChanged()` — L2548
- `.radRxSourceRx1Rx2_CheckedChanged()` — L2562
- `.chkEnableDiversity_CheckedChanged()` — L2576
- `.FormEncoderEvent()` — L2594
- `.Callback()` — L2613
- `.DiversityForm_Load()` — L2621
- `.DiversityForm_Resize()` — L2635
- `.udGainMulti_ValueChanged()` — L2640
- `.chkAlwaysOnTop_CheckedChanged()` — L2677
- `.chkNoAttLink_CheckedChanged()` — L2682
- `.chkVFOSync_CheckedChanged()` — L2687
- `.btnMemory_Click()` — L2707
- `.NormalizeAngle()` — L2770
- `.stepAngle()` — L2785
- `.btnShiftUp10_Click()` — L2798
- `.btnShift90_Click()` — L2803
- `.btnShiftDown10_Click()` — L2808
- `.udZoom_ValueChanged()` — L2813
- `.setNewNaming()` — L2828
- `.getMemoryIndexAtPoint()` — L2855
- `.updateHoverMemory()` — L2873
- `.recallMemory()` — L2885
- `.picRadar_MouseLeave()` — L2929

#### `memorySettings` (type, L1443)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/DiversityForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
