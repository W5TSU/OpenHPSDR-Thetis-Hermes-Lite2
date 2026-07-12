# `Console/PSForm.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** PureSignal TX linearization control panel and the amplifier gain/phase view (backed by wdsp `calcc.c`/`iqc.c`).

## How this file is used

- Used by (incoming references from other files):
  - `Console/AmpView.cs` (references ×1)
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)
  - `Console/AmpView.Designer.cs` (references ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Types

#### `PSForm` (type, L61)

- `.startPSThread()` — L141
- `.StopPSThread()` — L168
- `.onConsoleClosingAsync()` — L182
- `.onPowerOn()` — L187
- `.PSLoop()` — L193
- `.psdefpeak()` — L375
- `.PSForm_Load()` — L390
- `.SetupForm()` — L395
- `.PSForm_Closing()` — L411
- `.CloseAmpView()` — L429
- `.RunAmpv()` — L450
- `.btnPSAmpView_Click()` — L458
- `.btnPSCalibrate_Click()` — L470
- `.SingleCalrun()` — L485
- `.btnPSReset_Click()` — L490
- `.udPSMoxDelay_ValueChanged()` — L497
- `.udPSCalWait_ValueChanged()` — L502
- `.udPSPhnum_ValueChanged()` — L507
- `.btnPSTwoToneGen_Click()` — L512
- `.btnPSSave_Click()` — L528
- `.btnPSRestore_Click()` — L538
- `.SetDefaultPeaks()` — L551
- `.timer1code()` — L559
- `.timer2code()` — L732
- `.PSpeak_TextChanged()` — L819
- `.UpdateWarningSetPk()` — L832
- `.chkPSRelaxPtol_CheckedChanged()` — L837
- `.chkPSAutoAttenuate_CheckedChanged()` — L845
- `.checkLoopback_CheckedChanged()` — L850
- `.chkPSPin_CheckedChanged()` — L865
- `.chkPSMap_CheckedChanged()` — L873
- `.chkPSStbl_CheckedChanged()` — L881
- `.comboPSTint_SelectedIndexChanged()` — L889
- `.btnPSAdvanced_Click()` — L921
- `.setAdvancedView()` — L926
- `.chkPSOnTop_CheckedChanged()` — L935
- `.ShowAtStartup_LinearityForm()` — L941
- `.ShowAtStartup_AmpViewForm()` — L948
- `.ForcePS()` — L956
- `.chkQuickAttenuate_CheckedChanged()` — L990
- `.btnDefaultPeaks_Click()` — L995
- `.chkShow2ToneMeasurements_CheckedChanged()` — L1000
- `.FixAmpViewOnTop()` — L1005

#### `eCMDState` (type, L115)

_No extracted members._

#### `eAAState` (type, L126)

_No extracted members._

#### `puresignal` (type, L1014)

- `.SetPSRunCal()` — L1018
- `.SetPSMox()` — L1021
- `.GetPSInfo()` — L1024
- `.SetPSReset()` — L1027
- `.SetPSMancal()` — L1030
- `.SetPSAutomode()` — L1033
- `.SetPSTurnon()` — L1036
- `.SetPSControl()` — L1039
- `.SetPSLoopDelay()` — L1042
- `.SetPSMoxDelay()` — L1045
- `.SetPSTXDelay()` — L1048
- `.psccF()` — L1051
- `.PSSaveCorr()` — L1054
- `.PSRestoreCorr()` — L1057
- `.SetPSHWPeak()` — L1060
- `.GetPSHWPeak()` — L1063
- `.GetPSMaxTX()` — L1066
- `.SetPSPtol()` — L1069
- `.GetPSDisp()` — L1072
- `.SetPSFeedbackRate()` — L1075
- `.SetPSPinMode()` — L1078
- `.SetPSMapMode()` — L1081
- `.SetPSStabilize()` — L1084
- `.SetPSIntsAndSpi()` — L1087
- `.GetInfo()` — L1108
- `.NeedToRecalibrate()` — L1141
- `.NeedToRecalibrate_HL2()` — L1146

#### `EngineState` (type, L1177)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/PSForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
