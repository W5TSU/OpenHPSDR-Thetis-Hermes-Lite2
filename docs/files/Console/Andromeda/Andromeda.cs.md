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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Console` (type, L43)

- **`.SendAriesMsgViaTCPIPcat()`** — L72 — `private void SendAriesMsgViaTCPIPcat(string message)`
  send a CAT message over TCP/IP
  Called by: `.MakeAriesTuneRequestMsg()` (same file), `.MakeAriesAntennaChangeMsg()` (same file), `.MakeAriesRXAntennaChangeMsg()` (same file), `.MakeAriesATUEnableRequestMsg()` (same file), `.MakeAriesQuickTuneRequestMsg()` (same file), `.AriesErasePressed()` (same file) — and 2 more
- **`.UpdateAriesDisplayLabel()`** — L80 — `private void UpdateAriesDisplayLabel()`
  set the ATU display label appropriately
  Called by: `.CATHandleAriesTuneMessage()` (same file), `.SetAriesTXFrequency()` (same file), `.SetAriesTXAntenna()` (same file), `.InitialiseAries()` (same file)
- **`.CATHandleAriesTuneMessage()`** — L103 — `public void CATHandleAriesTuneMessage(bool TuneState)`
  handle a CAT ZZOX tune status message from ATU
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATHandleAriesEraseMessage()`** — L110 — `public void CATHandleAriesEraseMessage(bool EraseState)`
  handle a CAT ZZOZ erase status message from ATU
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MakeAriesVersionRequestMsg()`** — L121 — `private void MakeAriesVersionRequestMsg()`
  Called by: `.InitialiseAries()` (same file)
- **`.MakeAriesTuneRequestMsg()`** — L137 — `private void MakeAriesTuneRequestMsg(bool IsTune)`
  send CAT message to Aries to set the Tune on/off state
  Called by: `.SetAriesTuneState()` (same file)
- **`.MakeAriesAntennaChangeMsg()`** — L160 — `private void MakeAriesAntennaChangeMsg(int Ant)`
  send CAT message to Aries to set the TX antenna set the "Aries is in this state" variable only if the message is sent
  Called by: `.CheckAriesEnabled()` (same file)
- **`.MakeAriesRXAntennaChangeMsg()`** — L185 — `private void MakeAriesRXAntennaChangeMsg(int Ant)`
  send CAT message to Aries to set the RX antenna set the "Aries is in this state" variable only if the message is sent
  Called by: `.SendAriesRXAntennaMsg()` (same file)
- **`.MakeAriesATUEnableRequestMsg()`** — L209 — `private void MakeAriesATUEnableRequestMsg(bool IsEnabled)`
  send CAT message to Aries to set the ATU on/off state set the "Aries is in this state" variable only if the message is sent
  Called by: `.CheckAriesEnabled()` (same file), `.InitialiseAries()` (same file)
- **`.MakeAriesQuickTuneRequestMsg()`** — L230 — `private void MakeAriesQuickTuneRequestMsg(bool IsEnabled)`
  send CAT message to Aries to set the tune algorithm tweak
  Called by: `.InitialiseAries()` (same file)
- **`.AriesErasePressed()`** — L251 — `public void AriesErasePressed(int Channel)`
  send a CAT message to Aries to erase stored solutions for an antenna
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAriesTXFrequency()`** — L368 — `private void SetAriesTXFrequency(double NewFrequency)`
  function to check whether a new TX frequency needs to be signaled to Aries find out if frequency has moved by more than 10KHz and send mew message if so
  Called by: `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`)
- **`.SetAriesTXAntenna()`** — L403 — `public void SetAriesTXAntenna(int Antenna, Band band)`
  function to set the TX antenna for a given band this is called when an antenna control changes in setup. Antenna = 1,2 or 3
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAriesRXAntenna()`** — L423 — `public void SetAriesRXAntenna(int Antenna, Band band)`
  function to set the RX antenna for a given band this is called when an antenna control changes in setup. Antenna = 1,2 or 3
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAriesRXAntennaName()`** — L436 — `public void SetAriesRXAntennaName(string Antenna, Band band)`
  function to set the RX antenna name as a string for a given band this is called when an antenna control changes in setup. Antenna = a radio dependent string for check box selections Note that several can be selected!
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAriesRXAuxAntenna()`** — L450 — `public void SetAriesRXAuxAntenna(bool IsAux, Band band)`
  function to set whether an aux input is selected as RX antenna for a given band this is called when an antenna control changes in setup. parameter true if aux input (ext1, ext2, xvtr) selected
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAriesTuneState()`** — L461 — `private void SetAriesTuneState(bool IsTune)`
  function to tell Aries that TUNE is active at the moment, it just sends a CAT message
  Called by: `.chkTUN_CheckedChanged()` (`Console/console.cs`)
- **`.AntennaBandFromFreq()`** — L469 — `public Band AntennaBandFromFreq(bool IsTX)`
  helper to find nearest band, if we are out of an amateur band. Same logic as Alex.
  Called by: `.DisplayAriesTXAntenna()` (same file), `.SendAriesRXAntennaMsg()` (same file), `.DisplayAriesRXAntenna()` (same file)
- **`.DisplayAriesTXAntenna()`** — L532 — `private void DisplayAriesTXAntenna()`
  show the TX antenna on the Andromeda screen
  Called by: `.SetAriesTXAntenna()` (same file), `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`)
- **`.SendAriesRXAntennaMsg()`** — L559 — `void SendAriesRXAntennaMsg()`
  send an Aries RX antenna message has to work out the band, and it could be TX antenna if the "RX on TX antenna" button pressed
  Called by: `.SetAriesTXAntenna()` (same file), `.DisplayAriesRXAntenna()` (same file), `.CheckAriesEnabled()` (same file), `.chkRxAnt_CheckedChanged()` (`Console/console.cs`)
- **`.DisplayAriesRXAntenna()`** — L599 — `private void DisplayAriesRXAntenna()`
  show the RX antenna on the Andromeda screen this is now displayed as 2 characters, so existing names have to be remapped
  Called by: `.SetAriesRXAntenna()` (same file), `.SetAriesRXAntennaName()` (same file), `.SetAriesRXAuxAntenna()` (same file), `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`)
- **`.TXAntennaStep()`** — L677 — `void TXAntennaStep()`
  antenna step. Called by the button bar handler to step the RX and TX antennas 1-2-3-1 etc G8NJJ_21h
  Called by: `.ExecuteButtonAction()` (same file)
- **`.SetNewTXAntenna()`** — L716 — `public void SetNewTXAntenna(int Ant)`
  set TX antenna to specified antenna (1=3) called by status bar handler
  Called by: `.ToolStripMenuItem15_Click()` (`Console/console.cs`), `.ToolStripMenuItem16_Click()` (`Console/console.cs`), `.ToolStripMenuItem17_Click()` (`Console/console.cs`)
- **`.SetNewRXAntenna()`** — L729 — `public void SetNewRXAntenna(int Ant)`
  set RX antenna to specified antenna (1=3) called by status bar handler
  Called by: `.ToolStripMenuItem18_Click()` (`Console/console.cs`), `.ToolStripMenuItem19_Click()` (`Console/console.cs`), `.ToolStripMenuItem20_Click()` (`Console/console.cs`)
- **`.CheckAriesEnabled()`** — L745 — `private void CheckAriesEnabled()`
  check if Aries is enabled. Find band and antenna and see if it should be enabled. find the band assigned for TX; find its antenna; and see if the ATU enabled for that antenna.
  Called by: `.SetAriesTXFrequency()` (same file), `.SetAriesTXAntenna()` (same file), `.InitialiseAries()` (same file)
- **`.SetAriesAlexMode()`** — L791 — `void SetAriesAlexMode(bool state)`
  set Alex to Aries mode: if Aries is in standalone mode, TX antenna = RX antenna = 1
  Called by: `.InitialiseAries()` (same file)
- **`.InitialiseAries()`** — L812 — `private void InitialiseAries()`
  initialise Aries this is called when the CAT port is enabled, or when ZZZS from TCP/IP detected work out if ATU enabled for current TX antenna assume setup will already have provided the antenna settings
  Called by: `.MakeAttachedVersionString()` (same file)
- **`.TweakAriesCapacitance()`** — L828 — `void TweakAriesCapacitance(int Steps)`
  functions to tweak Aries Inductance and Capacitance value set by Andromeda encoder; tells the ATU to adjust its tune
  Called by: `.ExecuteEncoderStep()` (same file)
- **`.TweakAriesInductance()`** — L833 — `void TweakAriesInductance(int Steps)`
  Called by: `.ExecuteEncoderStep()` (same file)
- **`.MakeAriesATUTuneTweakMsg()`** — L842 — `private void MakeAriesATUTuneTweakMsg(bool IsInductance, int TweakSteps)`
  send CAT message to Aries to adjust tune re-uses the "encoder step" message to send an encoder step to Aries encoder 1 = inductance; encoder 2 = capacitance ZZZEnnm; m= no. steps; nn = encoder number (51+ = encoder 1+, turned anticlockwise)
  Called by: `.TweakAriesCapacitance()` (same file), `.TweakAriesInductance()` (same file)
- **`.SendGanymedeMsgViaTCPIPcat()`** — L902 — `private void SendGanymedeMsgViaTCPIPcat(string message)`
  send a CAT message over TCP/IP
  Called by: `.GanymedeResetPressed()` (same file), `.MakeGanymedeStatusRequestMsg()` (same file)
- **`.GetPAStatusText()`** — L908 — `public static string GetPAStatusText(PAstatusIndicatorState state)`
  Returns pastatus text.
  Called by: `.CATHandleAmplifierTripMessage()` (same file)
- **`.CATHandleAmplifierTripMessage()`** — L944 — `public void CATHandleAmplifierTripMessage(int TripState)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GanymedeResetPressed()`** — L979 — `public void GanymedeResetPressed()`
  send a CAT message to Ganymede to reset the tripped state
  Called by: `.toolStripStatusLabel_PAstatus_MouseUp()` (`Console/console.cs`)
- **`.MakeGanymedeVersionRequestMsg()`** — L1003 — `private void MakeGanymedeVersionRequestMsg()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MakeGanymedeStatusRequestMsg()`** — L1020 — `private void MakeGanymedeStatusRequestMsg()`
  send a CAT message to Ganymede asking for amplifier status when first connected
  Called by: `.InitialiseGanymede()` (same file)
- **`.InitialiseGanymede()`** — L1043 — `private void InitialiseGanymede()`
  initialise Ganymede: called when ZZZS identify message has been received request status message
  Called by: `.MakeAttachedVersionString()` (same file)
- **`.EditAndromedaDataSet()`** — L1311 — `public void EditAndromedaDataSet()`
  this is called from the setup form, to edit control assignments
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendG2V2PanelMsgViaTCPIPcat()`** — L1335 — `private void SendG2V2PanelMsgViaTCPIPcat(string message)`
  send a CAT message over TCP/IP
  Called by: `.MakeIndicatorCATMsg()` (same file)
- **`.InitialiseAndromedaMenus()`** — L1351 — `private void InitialiseAndromedaMenus()`
  initialise the data structures for andromeda. there are 5 tables: "Menu Bar Settings": softkey menus, for the bottom of the screen; "Indicators": front panel indicator actions "Encoders": front panel encoder actions "Pushbuttons": front panel pushbutton actions "Multifunction Settings": actions…
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MakeNewAndromedaDataset()`** — L1609 — `private void MakeNewAndromedaDataset()`
  create a new Andromeda dataset from scratch, if one doesn't yet exist this is like a "factory reset" for the dataset if it can't be read from file create tables, and add initial data to them
  Called by: `.InitialiseAndromedaMenus()` (same file), `.ResetAndromedaDataset()` (same file)
- **`.MakeNewG2PanelDataset()`** — L1872 — `private void MakeNewG2PanelDataset()`
  create a new G2 panel dataset from scratch, if one doesn't yet exist this is like a "factory reset" for the dataset if it can't be read from file create tables, and add initial data to them
  Called by: `.ResetG2PanelDataset()` (same file)
- **`.SaveAndromedaDataset()`** — L2133 — `void SaveAndromedaDataset()`
  Saves andromeda dataset.
  Called by: `.InitialiseAndromedaMenus()` (same file), `.ResetAndromedaDataset()` (same file), `.ResetG2PanelDataset()` (same file)
- **`.ResetAndromedaDataset()`** — L2142 — `public void ResetAndromedaDataset()`
  Resets andromeda dataset.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetG2PanelDataset()`** — L2151 — `public void ResetG2PanelDataset()`
  Resets g2 panel dataset.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AndromedaIndicatorCheck()`** — L2168 — `void AndromedaIndicatorCheck(EIndicatorActions IndicatorType, bool IsRX1, bool State)`
  indicator check this is called from control CheckedChanged() handlers to see if an Andromeda indicator needs to be updated this PUSHes the indicator setting to Andromeda rather than periodic polling IsRX1: true for controls belonging to RX1; ignored otherwise State: true if indicator should be…
  Called by: `.CheckAriesEnabled()` (same file), `.HandleFrontPanelButtonPress()` (same file), `.ExecuteButtonAction()` (same file), `.chkMOX_CheckedChanged2()` (`Console/console.cs`), `.chkTUN_CheckedChanged()` (`Console/console.cs`), `.chkX2TR_CheckedChanged()` (`Console/console.cs`) — and 18 more
- **`.MakeAndromedaVersionRequestMsg()`** — L2201 — `private void MakeAndromedaVersionRequestMsg()`
  send a CAT message to Andromeda asking for h/w and s/w versions only required if USB CAT port
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MakeIndicatorCATMsg()`** — L2219 — `void MakeIndicatorCATMsg(int IndicatorNumber, bool State)`
  send a CAT message to Andromeda if enabled indicator number = number to send (0-19) State true if indicator should be lit
  Called by: `.AndromedaIndicatorCheck()` (same file), `.InitialiseAndromedaIndicators()` (same file)
- **`.InitialiseAndromedaIndicators()`** — L2247 — `private void InitialiseAndromedaIndicators(bool InitialiseAll)`
  initialise indicators after console initialised. scan through list and send required state. the CAT port should have been created by this point. InitialiseAll: if true, updates all; otherwise only indicators which should change when RX1/RX2 is toggled
  Called by: `.SaveAndromedaDataset()` (same file), `.ResetAndromedaDataset()` (same file), `.ResetG2PanelDataset()` (same file), `.MakeAttachedVersionString()` (same file), `.radRX1Show_CheckedChanged()` (`Console/console.cs`), `.radRX2Show_CheckedChanged()` (`Console/console.cs`)
- **`.btnAndrBar1_Click()`** — L2398 — `private void btnAndrBar1_Click(object sender, EventArgs e)`
  andromeda button bar button event handlers
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.btnAndrBar2_Click()`** — L2403 — `private void btnAndrBar2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAndrBar2` is clicked.
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.btnAndrBar3_Click()`** — L2408 — `private void btnAndrBar3_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAndrBar3` is clicked.
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.btnAndrBar4_Click()`** — L2413 — `private void btnAndrBar4_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAndrBar4` is clicked.
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.btnAndrBar5_Click()`** — L2418 — `private void btnAndrBar5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAndrBar5` is clicked.
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.btnAndrBar6_Click()`** — L2423 — `private void btnAndrBar6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAndrBar6` is clicked.
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.btnAndrBar7_Click()`** — L2428 — `private void btnAndrBar7_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAndrBar7` is clicked.
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.btnAndrBar8_Click()`** — L2433 — `private void btnAndrBar8_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAndrBar8` is clicked.
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.HandleFrontPanelEncoderStep()`** — L2442 — `public void HandleFrontPanelEncoderStep(int Encoder, int Steps)`
  handle a rotational step from a front panel encoder, sent by CAT command encoder = 0-19; steps = number of turns; +ve = clockwise
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleFrontPanelVFOEncoderStep()`** — L2498 — `public void HandleFrontPanelVFOEncoderStep(int Steps)`
  handle a rotational step from a front panel VFO encoder steps = number of turns; +ve = clockwise in FM mode, count several clicks before stepping VFO, to make the dial "slower"
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalculateFastTuneSteps()`** — L2531 — `int CalculateFastTuneSteps(int RawSteps)`
  calculate fast tune rate. This uses a cubic calculation to increase apparent tune rate if tune knob moved more quickly.
  Called by: `.HandleFrontPanelVFOEncoderStep()` (same file)
- **`.HandleFrontPanelButtonPress()`** — L2585 — `public void HandleFrontPanelButtonPress(int Button, bool State, bool LongPress)`
  handle a button press from a front panel physical button, sent by CAT command button = 0-98 state true if pressed normally; LongPress = true if a "long press" event if AndromedaBandButtonsEnabled, treat as a band keypad instead
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MakeAttachedVersionString()`** — L2745 — `public void MakeAttachedVersionString(int Product, int HardwareVersion, int SoftwareVersion)`
  make version string for Andromeda, Aries and Ganymede these arise from ZZZS; CAT messages
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExecuteEncoderStep()`** — L2804 — `private void ExecuteEncoderStep(EEncoderActions Action, int Steps, int OverrideRX)`
  handler for encoder steps OverrideRX sets how to select the RX control used OverrideRX=0: use setting set by show_rx1 variable (from radio button) OverrideRX = 1: use RX1 OverrideRX = 2: use RX2 OverrideRX = 3: use Sub-RX these are only relevant to some controls!
  Called by: `.HandleFrontPanelEncoderStep()` (same file)
- **`.CheckGainFormAutoShow()`** — L3250 — `void CheckGainFormAutoShow()`
  method to check of the gain control form should be auto-shown if an encoder event occurred called whenever an encoder event for a control on this form happens
  Called by: `.ExecuteEncoderStep()` (same file)
- **`.CheckDiversityFormAutoShow()`** — L3264 — `void CheckDiversityFormAutoShow()`
  method to check of the gain control form should be auto-shown if an encoder event occurred called whenever an encoder event for a control on this form happens
  Called by: `.ExecuteEncoderStep()` (same file)
- **`.ShowAndromedaSlider()`** — L3280 — `void ShowAndromedaSlider(int Value, int Min, int Max, string Label)`
  helper to display relative position of the adjusted control on the Andromeda display this is a value i percent between 0 and 100, because min and max values are data type dependent ideally the control will be turned off after 2s delay
  Called by: `.ExecuteEncoderStep()` (same file)
- **`.Callback()`** — L3301 — `private void Callback(object source, ElapsedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EncoderUpdate()`** — L3315 — `public void EncoderUpdate(int Steps, ref int Value, int Minimum, int Maximum)`
  helper to clip an encoder adjusted value at the min and max allowed Steps number of encoder steps to add/subtract value the control value being edited Minimum min allowed value for control maximum max allowed value for control
  Called by: `.HandleFrontPanelEncoderStep()` (same file), `.ExecuteEncoderStep()` (same file)
- **`.ExecuteButtonBarPress()`** — L3328 — `private void ExecuteButtonBarPress(int ButtonNumber)`
  execute events for menu button bar presses called with a button number; finds the action assigned and implements it
  Called by: `.btnAndrBar1_Click()` (same file), `.btnAndrBar2_Click()` (same file), `.btnAndrBar3_Click()` (same file), `.btnAndrBar4_Click()` (same file), `.btnAndrBar5_Click()` (same file), `.btnAndrBar6_Click()` (same file) — and 2 more
- **`.AndromedaMenuTimerCallback()`** — L3368 — `private void AndromedaMenuTimerCallback(object source, ElapsedEventArgs e)`
  callback for andromeda menu timer. If it times out, a menu hasn't been touched for 10 seconds so reset to 1st menu
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExecuteButtonAction()`** — L3380 — `private void ExecuteButtonAction(EButtonBarActions assignedAction, int OverrideRX)`
  execute a single button press action this can be invoked by a menu button press or front panel button press
  Called by: `.HandleFrontPanelButtonPress()` (same file), `.ExecuteButtonBarPress()` (same file), `.CollapseDisplay()` (`Console/console.cs`)
- **`.UpdateAndromedaSkins()`** — L4069 — `public void UpdateAndromedaSkins()`
  Updates andromeda skins.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupModePanels()`** — L4082 — `private void setupModePanels()`
  Called by: `.ExecuteButtonAction()` (same file), `.SelectModeDependentPanel()` (same file)
- **`.ExecuteButtonLongpress()`** — L4103 — `private void ExecuteButtonLongpress(EButtonBarActions assignedAction, int OverrideRX)`
  execute a single button longpress action this can be invoked by a menu button press or front panel button press only a few buttons support long press)
  Called by: `.HandleFrontPanelButtonPress()` (same file)
- **`.SelectModeDependentPanel()`** — L4148 — `private void SelectModeDependentPanel()`
  bring the right mode dependent panel to the front
  Called by: `.ExecuteButtonAction()` (same file), `.SetRX1Mode()` (`Console/console.cs`), `.SetRX2Mode()` (`Console/console.cs`), `.ExpandDisplay()` (`Console/console.cs`), `.CollapseDisplay()` (`Console/console.cs`)
- **`.UpdateButtonBarLabel()`** — L4251 — `private String UpdateButtonBarLabel(EButtonBarActions assignedAction, String DefaultString, bool UseRX1)`
  update text for button bar buttons most will be labelled with default text; but they get an opportunity to rename themselves
  Called by: `.UpdateButtonBarButtons()` (same file)
- **`.CheckButtonHighlight()`** — L4393 — `private bool CheckButtonHighlight(EButtonBarActions assignedAction, bool UseRX1)`
  check if buttons on button bar should be highlighted (generally to indicate on/off state). if response is true, button will be highlighted
  Called by: `.UpdateButtonBarButtons()` (same file)
- **`.UpdateButtonBarButtons()`** — L4612 — `private void UpdateButtonBarButtons()`
  put text onto buttons for current menu
  Called by: `.SetNewTXAntenna()` (same file), `.SaveAndromedaDataset()` (same file), `.ExecuteButtonBarPress()` (same file), `.AndromedaMenuTimerCallback()` (same file), `.ExecuteButtonAction()` (same file), `.ExecuteButtonLongpress()` (same file) — and 3 more
- **`.panelButtonBar_Layout()`** — L4703 — `private void panelButtonBar_Layout(object sender, LayoutEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `EButtonBarActions` (type, L1057)

_No extracted members._

#### `EEncoderActions` (type, L1177)

_No extracted members._

#### `EIndicatorActions` (type, L1212)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/Andromeda.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
