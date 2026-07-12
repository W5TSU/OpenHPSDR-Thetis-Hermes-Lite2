# `Console/eqform.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX/TX graphic and parametric equalizer forms (backed by wdsp `eq.c`).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2, references ×1)
- Uses (outgoing references to other files):
  - `Console/Invoke/panelts.cs` (references ×1, calls ×1)
  - `Console/common.cs` (calls ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/numericupdownts.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/Invoke/trackbarts.cs` (references ×1)
  - `Console/ucParametricEq.cs` (references ×1)
- Most-referenced symbols from other files: `.DSPOptionsChanged()` (×1), `.Show()` (×1)

## Outline

### Types

#### `EQForm` (type, L54)

- `.Dispose()` — L232
- `.InitializeComponent()` — L251
- `.HighlightTXProfileSaveItems()` — L2056
- `.EQForm_Closing()` — L2376
- `.tbRXEQ_Scroll()` — L2385
- `.setDBtip()` — L2407
- `.picRXEQ_Paint()` — L2504
- `.chkRXEQEnabled_CheckedChanged()` — L2548
- `.enableRxEq()` — L2555
- `.chkTXEQEnabled_CheckedChanged()` — L2568
- `.enableTxEq()` — L2574
- `.btnRXEQReset_Click()` — L2580
- `.rad3Band_CheckedChanged()` — L2622
- `.rad10Band_CheckedChanged()` — L2694
- `.SetTXProfile()` — L2765
- `.setTXEQProfile()` — L2777
- `.chkLegacyEQ_CheckedChanged()` — L2880
- `.chkUseQFactors_CheckedChanged()` — L2931
- `.setupWDSPdataFromParaEQ()` — L2950
- `.dspUpdateTimerTick()` — L2977
- `.sendRXDspUpdate()` — L2998
- `.sendTXDspUpdate()` — L3041
- `.btnParaEQReset_Click()` — L3083
- `.chkParaEQ_enabled_CheckedChanged()` — L3090
- `.radParaEQ_RX_CheckedChanged()` — L3110
- `.radParaEQ_CheckedChanged()` — L3121
- `.ucParametricEq1_GlobalGainChanged()` — L3167
- `.ucParametricEq1_PointDataChanged()` — L3189
- `.ucParametricEq1_PointsChanged()` — L3197
- `.setParaEQData()` — L3322
- `.radParaEQ_RXTX_CheckedChanged()` — L3380
- `.nudParaEQ_selected_band_ValueChanged()` — L3409
- `.updateBandData()` — L3418
- `.ucParametricEq1_PointSelected()` — L3447
- `.ucParametricEq1_PointUnselected()` — L3455
- `.nudParaEQ_f_ValueChanged()` — L3463
- `.nudParaEQ_gain_ValueChanged()` — L3476
- `.nudParaEQ_q_ValueChanged()` — L3489
- `.chkPanaEQ_live_CheckedChanged()` — L3502
- `.nudParaEQ_preamp_ValueChanged()` — L3508
- `.DSPOptionsChanged()` — L3527
- `.nudParaEQ_low_ValueChanged()` — L3539
- `.nudParaEQ_high_ValueChanged()` — L3559
- `.setupLowHigh()` — L3578
- `.UpdateEQEnabled()` — L3583
- `.setupTimer()` — L3604
- `.EQForm_VisibleChanged()` — L3618
- `.Show()` — L3626
- `.chkLogScale_CheckedChanged()` — L3640

#### `ParaEQState` (type, L2827)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/eqform.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
