# `Console/Andromeda/Andromeda.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Core handler: decodes Andromeda encoder/button CAT messages and applies them to the console; manages button-bar text feedback.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×45, references ×1)
  - `Console/CAT/SIOListenerII.cs` (references ×7)
  - `Console/MeterManager.cs` (references ×6)
  - `Console/TCIServer.cs` (references ×4)
  - `Console/CAT/TCPIPcatServer.cs` (references ×3)
  - `Console/clsLegacyItemController.cs` (references ×2)
  - `Console/clsSpectrumProcessor.cs` (imports ×1, references ×1)
  - `Console/Dumpcap.cs` (references ×2)
  - `Console/frmBandStack2.cs` (references ×2)
  - `Console/frmMacroButtonConfig.cs` (references ×2)
  - `Console/HPSDR/IoBoardHl2.cs` (references ×2)
  - `Console/Skin.cs` (references ×2)
  - …and 33 more files
- Uses (outgoing references to other files):
  - `Console/console.cs` (calls ×23)
  - `Console/enums.cs` (references ×6)
  - `Console/Andromeda/AndromedaEditForm.cs` (calls ×2, references ×1)
  - `Console/xvtr.cs` (calls ×2)
  - `Console/Skin.cs` (calls ×2)
  - `Console/HPSDR/Alex.cs` (calls ×1)
- Most-referenced symbols from other files: `.AndromedaIndicatorCheck()` (×21), `.SelectModeDependentPanel()` (×4), `.SetNewTXAntenna()` (×3), `.SetNewRXAntenna()` (×3), `.SetAriesTXFrequency()` (×2), `.DisplayAriesTXAntenna()` (×2), `.DisplayAriesRXAntenna()` (×2), `.InitialiseAndromedaIndicators()` (×2)

## Outline

### Types

#### `Console` (type, L43)

- `.SendAriesMsgViaTCPIPcat()` — L72
- `.UpdateAriesDisplayLabel()` — L80
- `.CATHandleAriesTuneMessage()` — L103
- `.CATHandleAriesEraseMessage()` — L110
- `.MakeAriesVersionRequestMsg()` — L121
- `.MakeAriesTuneRequestMsg()` — L137
- `.MakeAriesAntennaChangeMsg()` — L160
- `.MakeAriesRXAntennaChangeMsg()` — L185
- `.MakeAriesATUEnableRequestMsg()` — L209
- `.MakeAriesQuickTuneRequestMsg()` — L230
- `.AriesErasePressed()` — L251
- `.SetAriesTXFrequency()` — L368
- `.SetAriesTXAntenna()` — L403
- `.SetAriesRXAntenna()` — L423
- `.SetAriesRXAntennaName()` — L436
- `.SetAriesRXAuxAntenna()` — L450
- `.SetAriesTuneState()` — L461
- `.AntennaBandFromFreq()` — L469
- `.DisplayAriesTXAntenna()` — L532
- `.SendAriesRXAntennaMsg()` — L559
- `.DisplayAriesRXAntenna()` — L599
- `.TXAntennaStep()` — L677
- `.SetNewTXAntenna()` — L716
- `.SetNewRXAntenna()` — L729
- `.CheckAriesEnabled()` — L745
- `.SetAriesAlexMode()` — L791
- `.InitialiseAries()` — L812
- `.TweakAriesCapacitance()` — L828
- `.TweakAriesInductance()` — L833
- `.MakeAriesATUTuneTweakMsg()` — L842
- `.SendGanymedeMsgViaTCPIPcat()` — L902
- `.GetPAStatusText()` — L908
- `.CATHandleAmplifierTripMessage()` — L944
- `.GanymedeResetPressed()` — L979
- `.MakeGanymedeVersionRequestMsg()` — L1003
- `.MakeGanymedeStatusRequestMsg()` — L1020
- `.InitialiseGanymede()` — L1043
- `.EditAndromedaDataSet()` — L1311
- `.SendG2V2PanelMsgViaTCPIPcat()` — L1335
- `.InitialiseAndromedaMenus()` — L1351
- `.MakeNewAndromedaDataset()` — L1609
- `.MakeNewG2PanelDataset()` — L1872
- `.SaveAndromedaDataset()` — L2133
- `.ResetAndromedaDataset()` — L2142
- `.ResetG2PanelDataset()` — L2151
- `.AndromedaIndicatorCheck()` — L2168
- `.MakeAndromedaVersionRequestMsg()` — L2201
- `.MakeIndicatorCATMsg()` — L2219
- `.InitialiseAndromedaIndicators()` — L2247
- `.btnAndrBar1_Click()` — L2398
- `.btnAndrBar2_Click()` — L2403
- `.btnAndrBar3_Click()` — L2408
- `.btnAndrBar4_Click()` — L2413
- `.btnAndrBar5_Click()` — L2418
- `.btnAndrBar6_Click()` — L2423
- `.btnAndrBar7_Click()` — L2428
- `.btnAndrBar8_Click()` — L2433
- `.HandleFrontPanelEncoderStep()` — L2442
- `.HandleFrontPanelVFOEncoderStep()` — L2498
- `.CalculateFastTuneSteps()` — L2531
- `.HandleFrontPanelButtonPress()` — L2585
- `.MakeAttachedVersionString()` — L2745
- `.ExecuteEncoderStep()` — L2804
- `.CheckGainFormAutoShow()` — L3250
- `.CheckDiversityFormAutoShow()` — L3264
- `.ShowAndromedaSlider()` — L3280
- `.Callback()` — L3301
- `.EncoderUpdate()` — L3315
- `.ExecuteButtonBarPress()` — L3328
- `.AndromedaMenuTimerCallback()` — L3368
- `.ExecuteButtonAction()` — L3380
- `.UpdateAndromedaSkins()` — L4069
- `.setupModePanels()` — L4082
- `.ExecuteButtonLongpress()` — L4103
- `.SelectModeDependentPanel()` — L4148
- `.UpdateButtonBarLabel()` — L4251
- `.CheckButtonHighlight()` — L4393
- `.UpdateButtonBarButtons()` — L4612
- `.panelButtonBar_Layout()` — L4703

#### `EButtonBarActions` (type, L1057)

_No extracted members._

#### `EEncoderActions` (type, L1177)

_No extracted members._

#### `EIndicatorActions` (type, L1212)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/Andromeda.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
