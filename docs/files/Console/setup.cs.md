# `Console/setup.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** The Setup dialog (~90k lines) — every configurable option (radio, audio, display, DSP, CAT, VAC…) lives here. The graph's biggest god node (2,234 edges) and, per its cohesion scores, the prime refactoring candidate.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/HPSDR/Penny.cs` (calls ×88)
  - `Console/MeterManager.cs` (calls ×68, references ×5)
  - `Console/common.cs` (calls ×49)
  - `Console/enums.cs` (references ×35)
  - `Console/audio.cs` (calls ×34)
  - `Console/clsThetisSkinService.cs` (calls ×16, references ×4)
  - `Console/ucLGPicker.cs` (references ×16, calls ×2)
  - `Console/clsCMASIOConfig.cs` (calls ×15)
  - `Console/CAT/SDRSerialPortII.cs` (calls ×10)
  - `Console/N1MM.cs` (calls ×10)
  - `Console/CAT/UsbBCDCable.cs` (calls ×6, references ×1)
  - `Console/HPSDR/clsRadioDiscovery.cs` (references ×4, calls ×3)
  - …and 29 more files

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Setup` (type, L80)

- **`.AfterConstructor()`** — L133 — `internal void AfterConstructor()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addDelegates()`** — L664 — `private void addDelegates()`
  Called by: `.AfterConstructor()` (same file)
- **`.RemoveDelegates()`** — L689 — `public void RemoveDelegates()`
  Removes delegates.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXInhibit()`** — L715 — `private void OnTXInhibit(bool oldState, bool newState)`
  Handles/raises the txinhibit event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRX2EnabledChanged()`** — L721 — `private void OnRX2EnabledChanged(bool enabled)`
  Handles/raises the rx2 enabled changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitGeneralTab()`** — L742 — `private void InitGeneralTab(List<string> recoveryList = null)`
  note: any control that needs it settings recovered from the DB when cancel clicked, uses this 'needsRecovering' system
  Called by: `.getOptions()` (same file)
- **`.InitADC()`** — L752 — `private void InitADC(List<string> recoveryList = null)`
  Inits adc.
  Called by: `.getOptions()` (same file)
- **`.SetHWSampleRate()`** — L810 — `public void SetHWSampleRate(int rx, int rate)`
  Sets hwsample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHWSampleRate()`** — L828 — `public int GetHWSampleRate(int rx)`
  Returns hwsample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitAudioTab()`** — L842 — `public void InitAudioTab(List<string> recoveryList = null, bool only_rates = false)`
  Inits audio tab.
  Called by: `.AfterConstructor()` (same file), `.getOptions()` (same file), `.ucRadioList_Radios_SelectedRadioChanged()` (same file)
- **`.InitAdvancedAudioTab()`** — L887 — `private void InitAdvancedAudioTab(List<string> recoveryList = null)`
  Inits advanced audio tab.
  Called by: `.AfterConstructor()` (same file), `.getOptions()` (same file)
- **`.InitDisplayTab()`** — L926 — `private void InitDisplayTab(List<string> recoveryList = null)`
  Inits display tab.
  Called by: `.getOptions()` (same file)
- **`.SetMultiMeterMode()`** — L977 — `public void SetMultiMeterMode(MultiMeterMeasureMode mode)`
  Sets multi meter mode.
  Called by: `.InitDisplayTab()` (same file)
- **`.InitDSPTab()`** — L993 — `private void InitDSPTab(List<string> recoveryList = null)`
  Inits dsptab.
  Called by: `.getOptions()` (same file)
- **`.InitKeyboardTab()`** — L998 — `private void InitKeyboardTab(List<string> recoveryList = null)`
  Inits keyboard tab.
  Called by: `.getOptions()` (same file)
- **`.InitAppearanceTab()`** — L1045 — `private void InitAppearanceTab(List<string> recoveryList = null)`
  Inits appearance tab.
  Called by: `.getOptions()` (same file)
- **`.PerformDelayedInitalistion()`** — L1080 — `public void PerformDelayedInitalistion()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RefreshCOMPortLists()`** — L1137 — `private void RefreshCOMPortLists()`
  Refreshes comport lists.
  Called by: `.AfterConstructor()` (same file)
- **`.RefreshSkinList()`** — L1185 — `private void RefreshSkinList()`
  Refreshes skin list.
  Called by: `.AfterConstructor()` (same file), `.fileDownloadHandler()` (same file), `.btnRemoveSkin_Click()` (same file)
- **`.selectSkin()`** — L1227 — `private void selectSkin()`
  Called by: `.AfterConstructor()` (same file)
- **`.GetHosts()`** — L1243 — `private void GetHosts()`
  Returns hosts.
  Called by: `.AfterConstructor()` (same file)
- **`.GetDevices2()`** — L1260 — `private void GetDevices2()`
  Returns devices2.
  Called by: `.comboAudioDriver2_SelectedIndexChanged()` (same file)
- **`.GetDevices3()`** — L1274 — `private void GetDevices3()`
  Returns devices3.
  Called by: `.comboAudioDriver3_SelectedIndexChanged()` (same file)
- **`.getControlList()`** — L1288 — `private void getControlList(Control c, ref Dictionary<string, Control> a)`
  Returns control list.
  Called by: `.registerAllControlsForStateMonitoring()` (same file), `.SaveOptions()` (same file), `.getOptions()` (same file)
- **`.Show()`** — L1320 — `public new void Show()`
  MW0LGE_21d new code for selective cancel
  Called by: `.ShowSetupTab()` (same file), `.ShowMultiMeterSetupTab()` (same file)
- **`.Hide()`** — L1342 — `public new void Hide()`
  Called by: `.btnOK_Click()` (same file), `.btnCancel_Click()` (same file), `.Setup_Closing()` (same file)
- **`.inBothUpdateAndClick()`** — L1365 — `private bool inBothUpdateAndClick(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.inUpdateOnly()`** — L1369 — `private bool inUpdateOnly(string s)`
  Called by: `.removeSpecialCases()` (same file)
- **`.inClickOnly()`** — L1373 — `private bool inClickOnly(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.removeSpecialCases()`** — L1377 — `private void removeSpecialCases()`
  Called by: `.btnCancel_Click()` (same file)
- **`.needsRecovering()`** — L1435 — `private bool needsRecovering(List<string> recoveryList, string s)`
  Called by: `.InitGeneralTab()` (same file), `.InitADC()` (same file), `.InitAdvancedAudioTab()` (same file), `.InitDisplayTab()` (same file), `.InitDSPTab()` (same file), `.InitKeyboardTab()` (same file) — and 3 more
- **`.listClickedControls()`** — L1439 — `private void listClickedControls()`
  Called by: `.btnCancel_Click()` (same file)
- **`.listUpdatedControls()`** — L1446 — `private void listUpdatedControls()`
  Called by: `.btnCancel_Click()` (same file)
- **`.clearUpdatedClickControlList()`** — L1453 — `private void clearUpdatedClickControlList()`
  Called by: `.Show()` (same file), `.btnApply_Click()` (same file)
- **`.everyControlClickHandler()`** — L1458 — `private void everyControlClickHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkBoxCheckedChangeHandler()`** — L1464 — `private void checkBoxCheckedChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboboxSelectedIndexChangeHandler()`** — L1472 — `private void comboboxSelectedIndexChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.numericUDValueChangeHandler()`** — L1478 — `private void numericUDValueChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radioButtonCheckedChangeHandler()`** — L1484 — `private void radioButtonCheckedChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.textBoxTextChangeHandler()`** — L1490 — `private void textBoxTextChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.trackBarValueChangeHandler()`** — L1496 — `private void trackBarValueChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.colourButtonChangeHandler()`** — L1502 — `private void colourButtonChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgPickerChangeHandler()`** — L1508 — `private void lgPickerChangeHandler(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.registerAllControlsForStateMonitoring()`** — L1514 — `private void registerAllControlsForStateMonitoring()`
  Called by: `.AfterConstructor()` (same file)
- **`.SaveOptions()`** — L1546 — `public void SaveOptions()`
  Saves options.
  Called by: `.PreSaveOptions()` (same file), `.ApplyOptions()` (same file)
- **`.InitTransmitTab()`** — L1661 — `private void InitTransmitTab(List<string> recoveryList = null)`
  Inits transmit tab.
  Called by: `.getOptions()` (same file)
- **`.handleOutdatedOptions()`** — L1666 — `private void handleOutdatedOptions(ref Dictionary<string, string> getDict)`
  Called by: `.getOptions()` (same file)
- **`.removeOutdatedOptions()`** — L1699 — `private void removeOutdatedOptions()`
  Called by: `.SaveOptions()` (same file)
- **`.addToIgnore()`** — L1707 — `private void addToIgnore(ref List<string> controlNames, Control rootControl)`
  Called by: `.getOptions()` (same file)
- **`.getModelFromDB()`** — L1721 — `private HPSDRModel getModelFromDB()`
  Returns model from db.
  Called by: `.AfterConstructor()` (same file)
- **`.getOptions()`** — L1732 — `private void getOptions(List<string> recoveryList = null)`
  Returns options.
  Called by: `.AfterConstructor()` (same file), `.PreRestoreOptions()` (same file)
- **`.KeyToString()`** — L2068 — `private string KeyToString(Keys k)`
  Called by: `.InitKeyboardTab()` (same file), `.SetupKeyMap()` (same file)
- **`.SetupKeyMap()`** — L2095 — `private void SetupKeyMap()`
  Setups key map.
  Called by: `.AfterConstructor()` (same file)
- **`.SetPriorityClass()`** — L2203 — `public void SetPriorityClass()`
  Sets priority class.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ForceAllEvents()`** — L2207 — `private void ForceAllEvents()`
  Called by: `.AfterConstructor()` (same file)
- **`.GetTXProfileStrings()`** — L2943 — `public string[] GetTXProfileStrings()`
  Returns txprofile strings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTxProfiles()`** — L2960 — `public void GetTxProfiles()`
  Returns tx profiles.
  Called by: `.AfterConstructor()` (same file)
- **`.GetTxProfileDefs()`** — L2999 — `public void GetTxProfileDefs()`
  Returns tx profile defs.
  Called by: `.AfterConstructor()` (same file)
- **`.ConvertFromDBVal()`** — L3011 — `public static T ConvertFromDBVal<T>(object obj)`
  Converts from dbval.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isTXProfileSettingDifferent()`** — L3022 — `private bool isTXProfileSettingDifferent<T>(DataRow dr, string setting, T value, out string report)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getDataRowsForTXProfile()`** — L3039 — `private DataRow[] getDataRowsForTXProfile(string profile)`
  Returns data rows for txprofile.
  Called by: `.getTXProfileChangeReport()` (same file), `.checkTXProfileChanged2()` (same file), `.GetVACEnabledBitfield()` (same file), `.loadTXProfile()` (same file), `.btnTXProfileDelete_Click()` (same file), `.ExportCurrentTxProfile()` (same file)
- **`.getTXProfileChangeReport()`** — L3050 — `private string getTXProfileChangeReport(DataRow drToCheck = null, bool bOnlyVac = false)`
  Returns txprofile change report.
  Called by: `.lblTXProfileWarning_Click()` (same file)
- **`.checkTXProfileChanged2()`** — L3254 — `private bool checkTXProfileChanged2(DataRow drToCheck = null, bool bOnlyVac = false)`
  Called by: `.loadTXProfile()` (same file), `.comboTXProfileName_SelectedIndexChanged()` (same file), `.tmrCheckProfile_Tick()` (same file)
- **`.highlightTXProfileSaveItems()`** — L3464 — `private void highlightTXProfileSaveItems(bool bHighlight)`
  Called by: `.chkHighlightTXProfileSaveItems_CheckedChanged()` (same file)
- **`.updateTXProfileInDB()`** — L3644 — `private void updateTXProfileInDB(DataRow dr)`
  Called by: `.SaveTXProfileData()` (same file), `.btnTXProfileSave_Click()` (same file)
- **`.SaveTXProfileData()`** — L3825 — `public void SaveTXProfileData()`
  Saves txprofile data.
  Called by: `.comboTXProfileName_SelectedIndexChanged()` (same file)
- **`.UpdateWaterfallBandInfo()`** — L3857 — `public void UpdateWaterfallBandInfo()`
  Updates waterfall band info.
  Called by: `.udDisplayWaterfallLowLevel_ValueChanged()` (same file), `.udDisplayWaterfallHighLevel_ValueChanged()` (same file), `.udRX2DisplayWaterfallLowLevel_ValueChanged()` (same file), `.udRX2DisplayWaterfallHighLevel_ValueChanged()` (same file)
- **`.UpdateDisplayGridBandInfo()`** — L3867 — `public void UpdateDisplayGridBandInfo()`
  Updates display grid band info.
  Called by: `.udDisplayGridMax_ValueChanged()` (same file), `.udDisplayGridMin_ValueChanged()` (same file), `.udRX2DisplayGridMax_ValueChanged()` (same file), `.udRX2DisplayGridMin_ValueChanged()` (same file)
- **`.UpdateTXDisplayFFT()`** — L3877 — `public void UpdateTXDisplayFFT()`
  Updates txdisplay fft.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATSetRig()`** — L5808 — `public void CATSetRig(string rig)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initHPSDR()`** — L6381 — `private void initHPSDR()`
  Called by: `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.AddHPSDRPages()`** — L6617 — `public void AddHPSDRPages()`
  Adds hpsdrpages.
  Called by: `.initHPSDR()` (same file)
- **`.chkGeneralRXOnly_CheckedChanged()`** — L6725 — `private void chkGeneralRXOnly_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkGeneralRXOnly` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnGeneralCalFreqStart_Click()`** — L6750 — `private void btnGeneralCalFreqStart_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnGeneralCalFreqStart` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnGeneralCalLevelStart_Click()`** — L6762 — `private void btnGeneralCalLevelStart_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnGeneralCalLevelStart` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CalibrateFreq()`** — L6787 — `private void CalibrateFreq()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibrateLevel()`** — L6794 — `private void CalibrateLevel()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showCalibrateDone()`** — L6807 — `private void showCalibrateDone(string msg)`
  Called by: `.CalibrateFreq()` (same file), `.CalibrateLevel()` (same file), `.CalibratePAGain()` (same file)
- **`.chkGeneralDisablePTT_CheckedChanged()`** — L6815 — `private void chkGeneralDisablePTT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkGeneralDisablePTT` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboGeneralXVTR_SelectedIndexChanged()`** — L6821 — `private void comboGeneralXVTR_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboGeneralXVTR` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboGeneralProcessPriority_SelectedIndexChanged()`** — L6826 — `private void comboGeneralProcessPriority_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboGeneralProcessPriority` selection changes.
  Called by: `.SetPriorityClass()` (same file)
- **`.chkOptQuickQSY_CheckedChanged()`** — L6888 — `private void chkOptQuickQSY_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkOptQuickQSY` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOptAlwaysOnTop_CheckedChanged()`** — L6893 — `private void chkOptAlwaysOnTop_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkOptAlwaysOnTop` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptClickTuneOffsetDIGL_ValueChanged()`** — L6898 — `private void udOptClickTuneOffsetDIGL_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udOptClickTuneOffsetDIGL` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptClickTuneOffsetDIGU_ValueChanged()`** — L6903 — `private void udOptClickTuneOffsetDIGU_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udOptClickTuneOffsetDIGU` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptMaxFilterWidth_ValueChanged()`** — L6908 — `private void udOptMaxFilterWidth_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udOptMaxFilterWidth` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboOptFilterWidthMode_SelectedIndexChanged()`** — L6914 — `private void comboOptFilterWidthMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboOptFilterWidthMode` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptMaxFilterShift_ValueChanged()`** — L6930 — `private void udOptMaxFilterShift_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udOptMaxFilterShift` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkOptFilterSaveChanges_CheckedChanged()`** — L6936 — `private void chkOptFilterSaveChanges_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkOptFilterSaveChanges` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOptEnableKBShortcuts_CheckedChanged()`** — L6941 — `private void chkOptEnableKBShortcuts_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkOptEnableKBShortcuts` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udFilterDefaultLowCut_ValueChanged()`** — L6947 — `private void udFilterDefaultLowCut_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udFilterDefaultLowCut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2FilterDefaultLowCut_ValueChanged()`** — L6953 — `private void udRX2FilterDefaultLowCut_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2FilterDefaultLowCut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudioEnableVAC_CheckedChanged()`** — L6966 — `private void chkAudioEnableVAC_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioEnableVAC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2Enable_CheckedChanged()`** — L6985 — `private void chkVAC2Enable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC2Enable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboAudioDriver2_SelectedIndexChanged()`** — L7003 — `private void comboAudioDriver2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioDriver2` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAudioDriver3_SelectedIndexChanged()`** — L7054 — `private void comboAudioDriver3_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioDriver3` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAudioInput2_SelectedIndexChanged()`** — L7106 — `private void comboAudioInput2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioInput2` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAudioInput3_SelectedIndexChanged()`** — L7129 — `private void comboAudioInput3_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioInput3` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAudioOutput2_SelectedIndexChanged()`** — L7152 — `private void comboAudioOutput2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioOutput2` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAudioOutput3_SelectedIndexChanged()`** — L7175 — `private void comboAudioOutput3_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioOutput3` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ForceAudioReset()`** — L7198 — `public void ForceAudioReset()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboAudioSampleRate1_SelectedIndexChanged()`** — L7262 — `private void comboAudioSampleRate1_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioSampleRate1` selection changes.
  Called by: `.ForceAudioReset()` (same file)
- **`.comboAudioSampleRateRX2_SelectedIndexChanged()`** — L7453 — `private void comboAudioSampleRateRX2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioSampleRateRX2` selection changes.
  Called by: `.ForceAudioReset()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.comboAudioSampleRate2_SelectedIndexChanged()`** — L7532 — `private void comboAudioSampleRate2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioSampleRate2` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboAudioSampleRate3_SelectedIndexChanged()`** — L7567 — `private void comboAudioSampleRate3_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioSampleRate3` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboAudioBuffer2_SelectedIndexChanged()`** — L7601 — `private void comboAudioBuffer2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioBuffer2` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboAudioBuffer3_SelectedIndexChanged()`** — L7623 — `private void comboAudioBuffer3_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAudioBuffer3` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udAudioLatency2_ValueChanged()`** — L7645 — `private void udAudioLatency2_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udAudioLatency2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udAudioLatency2_Out_ValueChanged()`** — L7661 — `private void udAudioLatency2_Out_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udAudioLatency2_Out` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udAudioLatencyPAIn_ValueChanged()`** — L7678 — `private void udAudioLatencyPAIn_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udAudioLatencyPAIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udAudioLatencyPAOut_ValueChanged()`** — L7694 — `private void udAudioLatencyPAOut_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udAudioLatencyPAOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2Latency_ValueChanged()`** — L7710 — `private void udVAC2Latency_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udVAC2Latency` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2LatencyOut_ValueChanged()`** — L7725 — `private void udVAC2LatencyOut_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udVAC2LatencyOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2LatencyPAIn_ValueChanged()`** — L7740 — `private void udVAC2LatencyPAIn_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udVAC2LatencyPAIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2LatencyPAOut_ValueChanged()`** — L7755 — `private void udVAC2LatencyPAOut_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udVAC2LatencyPAOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudio2Stereo_CheckedChanged()`** — L7770 — `private void chkAudio2Stereo_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudio2Stereo` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAudioStereo3_CheckedChanged()`** — L7777 — `private void chkAudioStereo3_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioStereo3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAudioVACGainRX_ValueChanged()`** — L7784 — `private void udAudioVACGainRX_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udAudioVACGainRX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2GainRX_ValueChanged()`** — L7793 — `private void udVAC2GainRX_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udVAC2GainRX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udAudioVACGainTX_ValueChanged()`** — L7802 — `private void udAudioVACGainTX_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udAudioVACGainTX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2GainTX_ValueChanged()`** — L7811 — `private void udVAC2GainTX_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udVAC2GainTX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudioVACAutoEnable_CheckedChanged()`** — L7820 — `private void chkAudioVACAutoEnable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioVACAutoEnable` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2AutoEnable_CheckedChanged()`** — L7825 — `private void chkVAC2AutoEnable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC2AutoEnable` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAudioLatencyManual2_CheckedChanged()`** — L7830 — `private void chkAudioLatencyManual2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioLatencyManual2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudioLatencyManual2_Out_CheckedChanged()`** — L7847 — `private void chkAudioLatencyManual2_Out_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioLatencyManual2_Out` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudioLatencyPAInManual_CheckedChanged()`** — L7864 — `private void chkAudioLatencyPAInManual_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioLatencyPAInManual` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudioLatencyPAOutManual_CheckedChanged()`** — L7881 — `private void chkAudioLatencyPAOutManual_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioLatencyPAOutManual` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2LatencyManual_CheckedChanged()`** — L7898 — `private void chkVAC2LatencyManual_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC2LatencyManual` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2LatencyOutManual_CheckedChanged()`** — L7915 — `private void chkVAC2LatencyOutManual_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC2LatencyOutManual` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2LatencyPAInManual_CheckedChanged()`** — L7932 — `private void chkVAC2LatencyPAInManual_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC2LatencyPAInManual` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2LatencyPAOutManual_CheckedChanged()`** — L7949 — `private void chkVAC2LatencyPAOutManual_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC2LatencyPAOutManual` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chk20dbMicBoost_CheckedChanged()`** — L7965 — `private void chk20dbMicBoost_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chk20dbMicBoost` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayGridMax_LostFocus()`** — L7978 — `private void udDisplayGridMax_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayGridMax` loses focus.
  Called by: `.udDisplayGridMax_Click()` (same file), `.udDisplayGridMax_MouseWheel()` (same file)
- **`.udTXGridMax_LostFocus()`** — L7983 — `private void udTXGridMax_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXGridMax` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayGridMax_Click()`** — L7988 — `private void udDisplayGridMax_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayGridMax` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayGridMax_MouseWheel()`** — L7993 — `private void udDisplayGridMax_MouseWheel(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `udDisplayGridMax` receives a mouse wheel event.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayFPS_ValueChanged()`** — L7998 — `private void udDisplayFPS_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayFPS` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayGridMax_ValueChanged()`** — L8011 — `private void udDisplayGridMax_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayGridMax` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayGridMin_ValueChanged()`** — L8089 — `private void udDisplayGridMin_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayGridMin` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayGridStep_ValueChanged()`** — L8167 — `private void udDisplayGridStep_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayGridStep` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2DisplayGridMax_ValueChanged()`** — L8173 — `private void udRX2DisplayGridMax_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayGridMax` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2DisplayGridMin_ValueChanged()`** — L8252 — `private void udRX2DisplayGridMin_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayGridMin` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2DisplayGridStep_ValueChanged()`** — L8329 — `private void udRX2DisplayGridStep_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayGridStep` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDisplayLabelAlign_SelectedIndexChanged()`** — L8335 — `private void comboDisplayLabelAlign_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDisplayLabelAlign` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayPhasePts_ValueChanged()`** — L8360 — `private void udDisplayPhasePts_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayPhasePts` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayAVGTime_ValueChanged()`** — L8366 — `private void udDisplayAVGTime_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayAVGTime` value changes.
  Called by: `.ForceAllEvents()` (same file), `.udDisplayFPS_ValueChanged()` (same file)
- **`.udRX2DisplayAVGTime_ValueChanged()`** — L8383 — `private void udRX2DisplayAVGTime_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayAVGTime` value changes.
  Called by: `.ForceAllEvents()` (same file), `.udDisplayFPS_ValueChanged()` (same file)
- **`.udDisplayMeterDelay_ValueChanged()`** — L8395 — `private void udDisplayMeterDelay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayMeterDelay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayPeakText_ValueChanged()`** — L8401 — `private void udDisplayPeakText_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayPeakText` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayCPUMeter_ValueChanged()`** — L8407 — `private void udDisplayCPUMeter_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayCPUMeter` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnWaterfallLow_Changed()`** — L8413 — `private void clrbtnWaterfallLow_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayWaterfallLowLevel_ValueChanged()`** — L8419 — `private void udDisplayWaterfallLowLevel_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayWaterfallLowLevel` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayWaterfallHighLevel_ValueChanged()`** — L8491 — `private void udDisplayWaterfallHighLevel_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayWaterfallHighLevel` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayMultiPeakHoldTime_ValueChanged()`** — L8563 — `private void udDisplayMultiPeakHoldTime_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayMultiPeakHoldTime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayMultiTextHoldTime_ValueChanged()`** — L8569 — `private void udDisplayMultiTextHoldTime_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayMultiTextHoldTime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnRX2WaterfallLow_Changed()`** — L8576 — `private void clrbtnRX2WaterfallLow_Changed(object sender, System.EventArgs e)`
  RX2 WaterFall
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2DisplayWaterfallLowLevel_ValueChanged()`** — L8582 — `private void udRX2DisplayWaterfallLowLevel_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayWaterfallLowLevel` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2DisplayWaterfallHighLevel_ValueChanged()`** — L8654 — `private void udRX2DisplayWaterfallHighLevel_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayWaterfallHighLevel` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayScopeTime_ValueChanged()`** — L8726 — `private void udDisplayScopeTime_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayScopeTime` value changes.
  Called by: `.AfterConstructor()` (same file)
- **`.udDisplayMeterAvg_ValueChanged()`** — L8739 — `private void udDisplayMeterAvg_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayMeterAvg` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGridMax_ValueChanged()`** — L8747 — `private void udTXGridMax_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXGridMax` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXGridMin_ValueChanged()`** — L8757 — `private void udTXGridMin_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXGridMin` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXGridStep_ValueChanged()`** — L8767 — `private void udTXGridStep_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXGridStep` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXWFAmpMax_ValueChanged()`** — L8773 — `private void udTXWFAmpMax_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXWFAmpMax` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXWFAmpMin_ValueChanged()`** — L8780 — `private void udTXWFAmpMin_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXWFAmpMin` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboTXLabelAlign_SelectedIndexChanged()`** — L8787 — `private void comboTXLabelAlign_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboTXLabelAlign` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udLMSNR_ValueChanged()`** — L8820 — `private void udLMSNR_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udLMSNR` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udLMSNR2_ValueChanged()`** — L8835 — `private void udLMSNR2_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udLMSNR2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPNB_ValueChanged()`** — L8850 — `private void udDSPNB_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPNB` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDSPPhoneRXBuf_SelectedIndexChanged()`** — L8857 — `private void comboDSPPhoneRXBuf_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPPhoneRXBuf` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPPhoneTXBuf_SelectedIndexChanged()`** — L8862 — `private void comboDSPPhoneTXBuf_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPPhoneTXBuf` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPFMRXBuf_SelectedIndexChanged()`** — L8867 — `private void comboDSPFMRXBuf_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPFMRXBuf` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPFMTXBuf_SelectedIndexChanged()`** — L8872 — `private void comboDSPFMTXBuf_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPFMTXBuf` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPCWRXBuf_SelectedIndexChanged()`** — L8877 — `private void comboDSPCWRXBuf_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPCWRXBuf` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPDigRXBuf_SelectedIndexChanged()`** — L8882 — `private void comboDSPDigRXBuf_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPDigRXBuf` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPDigTXBuf_SelectedIndexChanged()`** — L8887 — `private void comboDSPDigTXBuf_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPDigTXBuf` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPPhoneRXFiltSize_SelectedIndexChanged()`** — L8892 — `private void comboDSPPhoneRXFiltSize_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPPhoneRXFiltSize` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPPhoneTXFiltSize_SelectedIndexChanged()`** — L8897 — `private void comboDSPPhoneTXFiltSize_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPPhoneTXFiltSize` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPFMRXFiltSize_SelectedIndexChanged()`** — L8902 — `private void comboDSPFMRXFiltSize_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPFMRXFiltSize` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPFMTXFiltSize_SelectedIndexChanged()`** — L8907 — `private void comboDSPFMTXFiltSize_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPFMTXFiltSize` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPCWRXFiltSize_SelectedIndexChanged()`** — L8912 — `private void comboDSPCWRXFiltSize_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPCWRXFiltSize` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPDigRXFiltSize_SelectedIndexChanged()`** — L8917 — `private void comboDSPDigRXFiltSize_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPDigRXFiltSize` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPDigTXFiltSize_SelectedIndexChanged()`** — L8922 — `private void comboDSPDigTXFiltSize_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPDigTXFiltSize` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPPhoneRXFiltType_SelectedIndexChanged()`** — L8927 — `private void comboDSPPhoneRXFiltType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPPhoneRXFiltType` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPPhoneTXFiltType_SelectedIndexChanged()`** — L8932 — `private void comboDSPPhoneTXFiltType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPPhoneTXFiltType` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPFMRXFiltType_SelectedIndexChanged()`** — L8937 — `private void comboDSPFMRXFiltType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPFMRXFiltType` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPFMTXFiltType_SelectedIndexChanged()`** — L8942 — `private void comboDSPFMTXFiltType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPFMTXFiltType` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPCWRXFiltType_SelectedIndexChanged()`** — L8947 — `private void comboDSPCWRXFiltType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPCWRXFiltType` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPDigRXFiltType_SelectedIndexChanged()`** — L8952 — `private void comboDSPDigRXFiltType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPDigRXFiltType` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboDSPDigTXFiltType_SelectedIndexChanged()`** — L8957 — `private void comboDSPDigTXFiltType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDSPDigTXFiltType` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.udLMSANF_ValueChanged()`** — L8964 — `private void udLMSANF_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udLMSANF` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udLMSANF2_ValueChanged()`** — L8979 — `private void udLMSANF2_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udLMSANF2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radANFPreAGC_CheckedChanged()`** — L8994 — `private void radANFPreAGC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radANFPreAGC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radANF2PreAGC_CheckedChanged()`** — L9015 — `private void radANF2PreAGC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radANF2PreAGC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPCWPitch_ValueChanged()`** — L9035 — `private void udDSPCWPitch_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPCWPitch` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkCWKeyerIambic_CheckedChanged()`** — L9041 — `private void chkCWKeyerIambic_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWKeyerIambic` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udCWKeyerWeight_ValueChanged()`** — L9050 — `private void udCWKeyerWeight_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWKeyerWeight` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udCWKeyerSemiBreakInDelay_ValueChanged()`** — L9056 — `private void udCWKeyerSemiBreakInDelay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWKeyerSemiBreakInDelay` value changes.
  Called by: `.chkCWBreakInEnabled_CheckStateChanged()` (same file), `.udCWKeyUpDelay_ValueChanged()` (same file)
- **`.chkDSPKeyerSemiBreakInEnabled_CheckedChanged()`** — L9061 — `private void chkDSPKeyerSemiBreakInEnabled_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDSPKeyerSemiBreakInEnabled` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWBreakInEnabled_CheckStateChanged()`** — L9065 — `private void chkCWBreakInEnabled_CheckStateChanged(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSideTones_CheckedChanged()`** — L9104 — `private void chkSideTones_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSideTones` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPKeyerSidetone_CheckedChanged()`** — L9115 — `private void chkDSPKeyerSidetone_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDSPKeyerSidetone` checked state changes.
  Called by: `.chkSideTones_CheckedChanged()` (same file)
- **`.chkDSPKeyerSidetone_software_CheckedChanged()`** — L9126 — `private void chkDSPKeyerSidetone_software_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPKeyerSidetone_software` checked state changes.
  Called by: `.chkSideTones_CheckedChanged()` (same file)
- **`.chkCWKeyerRevPdl_CheckedChanged()`** — L9138 — `private void chkCWKeyerRevPdl_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWKeyerRevPdl` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKeyerConnPrimary_SelectedIndexChanged()`** — L9144 — `private void comboKeyerConnPrimary_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKeyerConnPrimary` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboKeyerConnSecondary_SelectedIndexChanged()`** — L9157 — `private void comboKeyerConnSecondary_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKeyerConnSecondary` selection changes.
  Called by: `.AfterConstructor()` (same file)
- **`.comboKeyerConnKeyLine_SelectedIndexChanged()`** — L9225 — `private void comboKeyerConnKeyLine_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKeyerConnKeyLine` selection changes.
  Called by: `.comboKeyerConnSecondary_SelectedIndexChanged()` (same file)
- **`.comboKeyerConnPTTLine_SelectedIndexChanged()`** — L9250 — `private void comboKeyerConnPTTLine_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKeyerConnPTTLine` selection changes.
  Called by: `.comboKeyerConnSecondary_SelectedIndexChanged()` (same file)
- **`.udDSPAGCFixedGaindB_ValueChanged()`** — L9279 — `private void udDSPAGCFixedGaindB_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCFixedGaindB` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCMaxGaindB_ValueChanged()`** — L9289 — `private void udDSPAGCMaxGaindB_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCMaxGaindB` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCRX2MaxGaindB_ValueChanged()`** — L9299 — `private void udDSPAGCRX2MaxGaindB_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPAGCRX2MaxGaindB` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCRX2FixedGaindB_ValueChanged()`** — L9308 — `private void udDSPAGCRX2FixedGaindB_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPAGCRX2FixedGaindB` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCDecay_ValueChanged()`** — L9317 — `private void udDSPAGCDecay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCDecay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCRX2Decay_ValueChanged()`** — L9325 — `private void udDSPAGCRX2Decay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCRX2Decay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCSlope_ValueChanged()`** — L9332 — `private void udDSPAGCSlope_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCSlope` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCRX2Slope_ValueChanged()`** — L9339 — `private void udDSPAGCRX2Slope_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCRX2Slope` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCHangTime_ValueChanged()`** — L9345 — `private void udDSPAGCHangTime_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCHangTime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPAGCRX2HangTime_ValueChanged()`** — L9353 — `private void udDSPAGCRX2HangTime_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPAGCRX2HangTime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbDSPAGCHangThreshold_Scroll()`** — L9359 — `private void tbDSPAGCHangThreshold_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `tbDSPAGCHangThreshold` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbDSPAGCRX2HangThreshold_Scroll()`** — L9366 — `private void tbDSPAGCRX2HangThreshold_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `tbDSPAGCRX2HangThreshold` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPLevelerThreshold_ValueChanged()`** — L9376 — `private void udDSPLevelerThreshold_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPLevelerThreshold` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPLevelerDecay_ValueChanged()`** — L9382 — `private void udDSPLevelerDecay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPLevelerDecay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPLevelerEnabled_CheckedChanged()`** — L9389 — `private void chkDSPLevelerEnabled_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDSPLevelerEnabled` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPALCMaximumGain_ValueChanged()`** — L9410 — `private void udDSPALCMaximumGain_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPALCMaximumGain` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPALCDecay_ValueChanged()`** — L9417 — `private void udDSPALCDecay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDSPALCDecay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXFilterHigh_ValueChanged()`** — L9429 — `private void udTXFilterHigh_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXFilterHigh` value changes.
  Called by: `.ForceAllEvents()` (same file), `.udTXFilterHigh_LostFocus()` (same file)
- **`.TXBW()`** — L9460 — `private void TXBW()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.udTXFilterLow_ValueChanged()`** — L9470 — `private void udTXFilterLow_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXFilterLow` value changes.
  Called by: `.ForceAllEvents()` (same file), `.udTXFilterLow_LostFocus()` (same file)
- **`.udTransmitTunePower_ValueChanged()`** — L9500 — `private void udTransmitTunePower_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTransmitTunePower` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.GetVACEnabledBitfield()`** — L9515 — `public int GetVACEnabledBitfield(string profile_name = "")`
  Returns vacenabled bitfield.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.loadTXProfile()`** — L9542 — `private bool loadTXProfile(String sProfileName)`
  Called by: `.getOptions()` (same file), `.comboTXProfileName_SelectedIndexChanged()` (same file)
- **`.ForceTXProfileUpdate()`** — L9795 — `public void ForceTXProfileUpdate()`
  Called by: `.ForceAllEvents()` (same file)
- **`.comboTXProfileName_SelectedIndexChanged()`** — L9801 — `private void comboTXProfileName_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboTXProfileName` selection changes.
  Called by: `.ForceTXProfileUpdate()` (same file)
- **`.btnTXProfileSave_Click()`** — L9841 — `private void btnTXProfileSave_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnTXProfileSave` is clicked.
  Called by: `.comboTXProfileName_SelectedIndexChanged()` (same file)
- **`.btnTXProfileDelete_Click()`** — L9910 — `private void btnTXProfileDelete_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnTXProfileDelete` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udVOXGain_ValueChanged()`** — L9953 — `private void udVOXGain_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udVOXGain` value changes.
  Called by: `.chk20dbMicBoost_CheckedChanged()` (same file)
- **`.udTXAF_ValueChanged()`** — L9957 — `private void udTXAF_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXAF` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXAMCarrierLevel_ValueChanged()`** — L9962 — `private void udTXAMCarrierLevel_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXAMCarrierLevel` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSaveTXProfileOnExit_CheckedChanged()`** — L9968 — `private void chkSaveTXProfileOnExit_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSaveTXProfileOnExit` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udMicGainMin_ValueChanged()`** — L9974 — `private void udMicGainMin_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udMicGainMin` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udMicGainMax_ValueChanged()`** — L9980 — `private void udMicGainMax_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udMicGainMax` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udLineInBoost_ValueChanged()`** — L9986 — `private void udLineInBoost_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udLineInBoost` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowTopControls_CheckedChanged()`** — L9992 — `private void chkShowTopControls_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowTopControls` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowBandControls_CheckedChanged()`** — L9999 — `private void chkShowBandControls_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowBandControls` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkModeControls_CheckedChanged()`** — L10006 — `private void chkModeControls_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkModeControls` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowAndromedaTop_CheckedChanged()`** — L10016 — `private void chkShowAndromedaTop_CheckedChanged(object sender, EventArgs e)`
  G8NJJ: setup control to select an Andromeda top bar when display is collapsed
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowAndromedaBar_CheckedChanged()`** — L10026 — `private void chkShowAndromedaBar_CheckedChanged(object sender, EventArgs e)`
  G8NJJ: setup control to select an Andromeda top bar when display is collapsed
  Called by: `.ForceAllEvents()` (same file)
- **`.btnPAGainCalibration_Click()`** — L10038 — `private void btnPAGainCalibration_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnPAGainCalibration` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CalibratePAGain()`** — L10078 — `private void CalibratePAGain()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnBackground_Changed()`** — L10109 — `private void clrbtnBackground_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbBackgroundAlpha_Scroll()` (same file)
- **`.clrbtnTXBackground_Changed()`** — L10115 — `private void clrbtnTXBackground_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbTXBackgroundAlpha_Scroll()` (same file)
- **`.clrbtnGrid_Changed()`** — L10121 — `private void clrbtnGrid_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbGridCourseAlpha_Scroll()` (same file)
- **`.clrbtnTXVGrid_Changed()`** — L10127 — `private void clrbtnTXVGrid_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbTXVGridCourseAlpha_Scroll()` (same file)
- **`.clrbtnZeroLine_Changed()`** — L10133 — `private void clrbtnZeroLine_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnTXZeroLine_Changed()`** — L10139 — `private void clrbtnTXZeroLine_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbTXZeroLineAlpha_Scroll()` (same file)
- **`.clrbtnText_Changed()`** — L10145 — `private void clrbtnText_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnTXText_Changed()`** — L10151 — `private void clrbtnTXText_Changed(object sender, System.EventArgs e)`
  Called by: `.tbTXTextAlpha_Scroll()` (same file)
- **`.clrbtnDataLine_Changed()`** — L10156 — `private void clrbtnDataLine_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbDataLineAlpha_Scroll()` (same file)
- **`.RebuildLGBrushes()`** — L10163 — `public void RebuildLGBrushes()`
  Called by: `.clrbtnDataLine_Changed()` (same file), `.lgPickerRX1_Changed()` (same file), `.lgPickerRX1_GripperDBMChanged()` (same file)
- **`.clrbtnTXDataLine_Changed()`** — L10168 — `private void clrbtnTXDataLine_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbDataLineAlpha_tx_Scroll()` (same file)
- **`.clrbtnFilter_Changed()`** — L10175 — `private void clrbtnFilter_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbRX1FilterAlpha_Scroll()` (same file)
- **`.clrbtnGridTXFilter_Changed()`** — L10181 — `private void clrbtnGridTXFilter_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbTXFilterAlpha_Scroll()` (same file)
- **`.udDisplayLineWidth_ValueChanged()`** — L10187 — `private void udDisplayLineWidth_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayLineWidth` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXLineWidth_ValueChanged()`** — L10193 — `private void udTXLineWidth_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTXLineWidth` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnMeterLeft_Changed()`** — L10199 — `private void clrbtnMeterLeft_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnMeterRight_Changed()`** — L10205 — `private void clrbtnMeterRight_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnBtnSel_Changed()`** — L10211 — `private void clrbtnBtnSel_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnVFODark_Changed()`** — L10217 — `private void clrbtnVFODark_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnVFOLight_Changed()`** — L10223 — `private void clrbtnVFOLight_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnBandDark_Changed()`** — L10229 — `private void clrbtnBandDark_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnBandLight_Changed()`** — L10235 — `private void clrbtnBandLight_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnPeakText_Changed()`** — L10241 — `private void clrbtnPeakText_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnOutOfBand_Changed()`** — L10247 — `private void clrbtnOutOfBand_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVFOSmallLSD_CheckedChanged()`** — L10253 — `private void chkVFOSmallLSD_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOSmallLSD` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnVFOSmallColor_Changed()`** — L10259 — `private void clrbtnVFOSmallColor_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnInfoButtonsColor_Changed()`** — L10265 — `private void clrbtnInfoButtonsColor_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnPeakBackground_Changed()`** — L10271 — `private void clrbtnPeakBackground_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnMeterBackground_Changed()`** — L10277 — `private void clrbtnMeterBackground_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnBandBackground_Changed()`** — L10283 — `private void clrbtnBandBackground_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnVFOBackground_Changed()`** — L10289 — `private void clrbtnVFOBackground_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneUp1_SelectedIndexChanged()`** — L10299 — `private void comboKBTuneUp1_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneUp1` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneDown1_SelectedIndexChanged()`** — L10305 — `private void comboKBTuneDown1_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneDown1` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneUp2_SelectedIndexChanged()`** — L10311 — `private void comboKBTuneUp2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneUp2` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneDown2_SelectedIndexChanged()`** — L10317 — `private void comboKBTuneDown2_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneDown2` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneUp3_SelectedIndexChanged()`** — L10323 — `private void comboKBTuneUp3_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneUp3` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneDown3_SelectedIndexChanged()`** — L10329 — `private void comboKBTuneDown3_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneDown3` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneUp4_SelectedIndexChanged()`** — L10335 — `private void comboKBTuneUp4_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneUp4` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneDown4_SelectedIndexChanged()`** — L10341 — `private void comboKBTuneDown4_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneDown4` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneUp5_SelectedIndexChanged()`** — L10347 — `private void comboKBTuneUp5_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneUp5` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneDown5_SelectedIndexChanged()`** — L10353 — `private void comboKBTuneDown5_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneDown5` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneUp6_SelectedIndexChanged()`** — L10359 — `private void comboKBTuneUp6_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneUp6` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneDown6_SelectedIndexChanged()`** — L10365 — `private void comboKBTuneDown6_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneDown6` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBTuneUp7_SelectedIndexChanged()`** — L10371 — `private void comboKBTuneUp7_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneUp7` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboKBTuneDown7_SelectedIndexChanged()`** — L10376 — `private void comboKBTuneDown7_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBTuneDown7` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboKBBandUp_SelectedIndexChanged()`** — L10381 — `private void comboKBBandUp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBBandUp` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBBandDown_SelectedIndexChanged()`** — L10387 — `private void comboKBBandDown_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBBandDown` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBFilterUp_SelectedIndexChanged()`** — L10393 — `private void comboKBFilterUp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBFilterUp` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBFilterDown_SelectedIndexChanged()`** — L10399 — `private void comboKBFilterDown_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBFilterDown` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBModeUp_SelectedIndexChanged()`** — L10405 — `private void comboKBModeUp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBModeUp` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBModeDown_SelectedIndexChanged()`** — L10411 — `private void comboKBModeDown_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBModeDown` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBCWDot_SelectedIndexChanged()`** — L10417 — `private void comboKBCWDot_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBCWDot` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBCWDash_SelectedIndexChanged()`** — L10423 — `private void comboKBCWDash_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBCWDash` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBRITUp_SelectedIndexChanged()`** — L10429 — `private void comboKBRITUp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBRITUp` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboKBRITDown_SelectedIndexChanged()`** — L10434 — `private void comboKBRITDown_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBRITDown` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboKBXITUp_SelectedIndexChanged()`** — L10439 — `private void comboKBXITUp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBXITUp` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboKBXITDown_SelectedIndexChanged()`** — L10444 — `private void comboKBXITDown_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBXITDown` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboKBPTTTx_SelectedIndexChanged()`** — L10449 — `private void comboKBPTTTx_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBPTTTx` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKBPTTRx_SelectedIndexChanged()`** — L10455 — `private void comboKBPTTRx_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboKBPTTRx` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.initCATandPTTprops()`** — L10465 — `public void initCATandPTTprops()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.copyCATPropsToDialogVars()`** — L10538 — `public void copyCATPropsToDialogVars()`
  called in error cases to set the dialiog vars from the console properties -- sort of ugly, we should only have 1 copy of this stuff
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.copyCAT2PropsToDialogVars()`** — L10554 — `public void copyCAT2PropsToDialogVars()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.copyCAT3PropsToDialogVars()`** — L10562 — `public void copyCAT3PropsToDialogVars()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.copyCAT4PropsToDialogVars()`** — L10570 — `public void copyCAT4PropsToDialogVars()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.copyAndromedaCATPropsToDialogVars()`** — L10578 — `public void copyAndromedaCATPropsToDialogVars()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.copyAriesCATPropsToDialogVars()`** — L10586 — `public void copyAriesCATPropsToDialogVars()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.copyGanymedeCATPropsToDialogVars()`** — L10595 — `public void copyGanymedeCATPropsToDialogVars()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkCATEnable_CheckedChanged()`** — L10603 — `private void chkCATEnable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCATEnable` checked state changes.
  Called by: `.AfterConstructor()` (same file)
- **`.chkCAT2Enable_CheckedChanged()`** — L10680 — `private void chkCAT2Enable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCAT2Enable` checked state changes.
  Called by: `.AfterConstructor()` (same file)
- **`.chkCAT3Enable_CheckedChanged()`** — L10736 — `private void chkCAT3Enable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCAT3Enable` checked state changes.
  Called by: `.AfterConstructor()` (same file)
- **`.chkCAT4Enable_CheckedChanged()`** — L10792 — `private void chkCAT4Enable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCAT4Enable` checked state changes.
  Called by: `.AfterConstructor()` (same file)
- **`.ChkEnableAndromeda_CheckedChanged()`** — L10849 — `private void ChkEnableAndromeda_CheckedChanged(object sender, EventArgs e)`
  reguire a valid COM port for Andromeda; not needed for G2 panel
  Called by: `.AfterConstructor()` (same file)
- **`.enableCAT_HardwareFields()`** — L10926 — `private void enableCAT_HardwareFields(bool enable)`
  Called by: `.chkCATEnable_CheckedChanged()` (same file)
- **`.enableCAT2_HardwareFields()`** — L10934 — `private void enableCAT2_HardwareFields(bool enable)`
  Called by: `.chkCAT2Enable_CheckedChanged()` (same file)
- **`.enableCAT3_HardwareFields()`** — L10942 — `private void enableCAT3_HardwareFields(bool enable)`
  Called by: `.chkCAT3Enable_CheckedChanged()` (same file)
- **`.enableCAT4_HardwareFields()`** — L10950 — `private void enableCAT4_HardwareFields(bool enable)`
  Called by: `.chkCAT4Enable_CheckedChanged()` (same file)
- **`.doEnablementOnBitBangEnable()`** — L10958 — `private void doEnablementOnBitBangEnable()`
  Called by: `.chkCATPTT_RTS_CheckedChanged()` (same file), `.chkCATPTT_DTR_CheckedChanged()` (same file), `.comboCATPTTPort_SelectedIndexChanged()` (same file)
- **`.chkCATPTT_RTS_CheckedChanged()`** — L10972 — `private void chkCATPTT_RTS_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCATPTT_RTS` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCATPTT_DTR_CheckedChanged()`** — L10979 — `private void chkCATPTT_DTR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCATPTT_DTR` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCATPTTEnabled_CheckedChanged()`** — L10986 — `private void chkCATPTTEnabled_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCATPTTEnabled` checked state changes.
  Called by: `.AfterConstructor()` (same file), `.comboCATPTTPort_SelectedIndexChanged()` (same file)
- **`.comboCATparity_SelectedIndexChanged()`** — L11033 — `private void comboCATparity_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCATparity` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT2parity_SelectedIndexChanged()`** — L11041 — `private void comboCAT2parity_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT2parity` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT3parity_SelectedIndexChanged()`** — L11049 — `private void comboCAT3parity_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT3parity` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT4parity_SelectedIndexChanged()`** — L11057 — `private void comboCAT4parity_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT4parity` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCATPort_SelectedIndexChanged()`** — L11065 — `private void comboCATPort_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCATPort` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT2Port_SelectedIndexChanged()`** — L11094 — `private void comboCAT2Port_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT2Port` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT3Port_SelectedIndexChanged()`** — L11114 — `private void comboCAT3Port_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT3Port` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT4Port_SelectedIndexChanged()`** — L11134 — `private void comboCAT4Port_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT4Port` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboAndromedaCATPort_SelectedIndexChanged()`** — L11154 — `private void ComboAndromedaCATPort_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboAndromedaCATPort` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCATPTTPort_SelectedIndexChanged()`** — L11173 — `private void comboCATPTTPort_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCATPTTPort` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCATbaud_SelectedIndexChanged()`** — L11218 — `private void comboCATbaud_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCATbaud` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT2baud_SelectedIndexChanged()`** — L11224 — `private void comboCAT2baud_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT2baud` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT3baud_SelectedIndexChanged()`** — L11230 — `private void comboCAT3baud_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT3baud` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT4baud_SelectedIndexChanged()`** — L11236 — `private void comboCAT4baud_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT4baud` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCATdatabits_SelectedIndexChanged()`** — L11242 — `private void comboCATdatabits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCATdatabits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT2databits_SelectedIndexChanged()`** — L11248 — `private void comboCAT2databits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT2databits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT3databits_SelectedIndexChanged()`** — L11254 — `private void comboCAT3databits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT3databits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT4databits_SelectedIndexChanged()`** — L11260 — `private void comboCAT4databits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT4databits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCATstopbits_SelectedIndexChanged()`** — L11266 — `private void comboCATstopbits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCATstopbits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT2stopbits_SelectedIndexChanged()`** — L11272 — `private void comboCAT2stopbits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT2stopbits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT3stopbits_SelectedIndexChanged()`** — L11278 — `private void comboCAT3stopbits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT3stopbits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT4stopbits_SelectedIndexChanged()`** — L11284 — `private void comboCAT4stopbits_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboCAT4stopbits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCATTest_Click()`** — L11290 — `private void btnCATTest_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnCATTest` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCATRigType_SelectedIndexChanged()`** — L11298 — `private void comboCATRigType_SelectedIndexChanged(object sender, System.EventArgs e)`
  Modified 10/12/08 BT to change "SDR-1000" to "PowerSDR"
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTestIMD_CheckedChanged()`** — L11344 — `private async void chkTestIMD_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTestIMD` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.cmboSigGenRXMode_SelectedIndexChanged()`** — L11503 — `private void cmboSigGenRXMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `cmboSigGenRXMode` selection changes.
  Called by: `.chkSigGenRX2_CheckedChanged()` (same file), `.chkSigGenRX1_CheckedChanged()` (same file)
- **`.chkSigGenRX2_CheckedChanged()`** — L11543 — `private void chkSigGenRX2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkSigGenRX2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.cmboSigGenTXMode_SelectedIndexChanged()`** — L11548 — `private void cmboSigGenTXMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `cmboSigGenTXMode` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setButtonState()`** — L11610 — `private void setButtonState(bool bSaving, bool bLoading)`
  Sets button state.
  Called by: `.btnOK_Click()` (same file), `.btnCancel_Click()` (same file), `.btnApply_Click()` (same file), `.PreRestoreOptions()` (same file), `.PreSaveOptions()` (same file), `.ApplyOptions()` (same file) — and 1 more
- **`.WaitForSaveLoad()`** — L11658 — `public void WaitForSaveLoad(int nWait = -1)`
  Called by: `.btnOK_Click()` (same file), `.btnCancel_Click()` (same file), `.btnApply_Click()` (same file)
- **`.btnOK_Click()`** — L11669 — `private void btnOK_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnOK` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCancel_Click()`** — L11688 — `private void btnCancel_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnCancel` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnApply_Click()`** — L11717 — `private void btnApply_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnApply` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PreRestoreOptions()`** — L11735 — `private void PreRestoreOptions()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreSaveOptions()`** — L11741 — `private void PreSaveOptions()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ApplyOptions()`** — L11746 — `private void ApplyOptions()`
  Applys options.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.udGeneralLPTDelay_ValueChanged()`** — L11752 — `private void udGeneralLPTDelay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udGeneralLPTDelay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.Setup_Closing()`** — L11757 — `private void Setup_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `Setup` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnTXFilter_Changed()`** — L11778 — `private void clrbtnTXFilter_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tbDSPAudRX1APFGain_LostFocus()`** — L11785 — `private void tbDSPAudRX1APFGain_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDSPAudRX1APFGain` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbDSPAudRX1subAPFGain_LostFocus()`** — L11790 — `private void tbDSPAudRX1subAPFGain_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDSPAudRX1subAPFGain` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbDSPAudRX2APFGain_LostFocus()`** — L11795 — `private void tbDSPAudRX2APFGain_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDSPAudRX2APFGain` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udGeneralCalFreq1_LostFocus()`** — L11800 — `private void udGeneralCalFreq1_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udGeneralCalFreq1` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptClickTuneOffsetDIGL_LostFocus()`** — L11805 — `private void udOptClickTuneOffsetDIGL_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udOptClickTuneOffsetDIGL` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptClickTuneOffsetDIGU_LostFocus()`** — L11810 — `private void udOptClickTuneOffsetDIGU_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udOptClickTuneOffsetDIGU` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udGeneralCalLevel_LostFocus()`** — L11815 — `private void udGeneralCalLevel_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udGeneralCalLevel` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udGeneralCalFreq2_LostFocus()`** — L11820 — `private void udGeneralCalFreq2_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udGeneralCalFreq2` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udFilterDefaultLowCut_LostFocus()`** — L11825 — `private void udFilterDefaultLowCut_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udFilterDefaultLowCut` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptMaxFilterShift_LostFocus()`** — L11830 — `private void udOptMaxFilterShift_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udOptMaxFilterShift` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udOptMaxFilterWidth_LostFocus()`** — L11835 — `private void udOptMaxFilterWidth_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udOptMaxFilterWidth` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAudioVACGainTX_LostFocus()`** — L11840 — `private void udAudioVACGainTX_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udAudioVACGainTX` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAudioVACGainRX_LostFocus()`** — L11845 — `private void udAudioVACGainRX_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udAudioVACGainRX` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAudioLatency2_LostFocus()`** — L11850 — `private void udAudioLatency2_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udAudioLatency2` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayScopeTime_LostFocus()`** — L11855 — `private void udDisplayScopeTime_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayScopeTime` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayMeterAvg_LostFocus()`** — L11860 — `private void udDisplayMeterAvg_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayMeterAvg` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayMultiTextHoldTime_LostFocus()`** — L11865 — `private void udDisplayMultiTextHoldTime_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayMultiTextHoldTime` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayMultiPeakHoldTime_LostFocus()`** — L11870 — `private void udDisplayMultiPeakHoldTime_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayMultiPeakHoldTime` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayWaterfallLowLevel_LostFocus()`** — L11875 — `private void udDisplayWaterfallLowLevel_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayWaterfallLowLevel` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayWaterfallHighLevel_LostFocus()`** — L11880 — `private void udDisplayWaterfallHighLevel_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayWaterfallHighLevel` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayCPUMeter_LostFocus()`** — L11885 — `private void udDisplayCPUMeter_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayCPUMeter` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayPeakText_LostFocus()`** — L11890 — `private void udDisplayPeakText_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayPeakText` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayMeterDelay_LostFocus()`** — L11895 — `private void udDisplayMeterDelay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayMeterDelay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayFPS_LostFocus()`** — L11900 — `private void udDisplayFPS_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayFPS` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayAVGTime_LostFocus()`** — L11905 — `private void udDisplayAVGTime_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayAVGTime` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayPhasePts_LostFocus()`** — L11910 — `private void udDisplayPhasePts_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayPhasePts` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayGridStep_LostFocus()`** — L11915 — `private void udDisplayGridStep_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayGridStep` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGridMin_LostFocus()`** — L11920 — `private void udTXGridMin_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGridMin` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGridStep_LostFocus()`** — L11925 — `private void udTXGridStep_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGridStep` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayGridMin_LostFocus()`** — L11930 — `private void udDisplayGridMin_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayGridMin` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPNB_LostFocus()`** — L11935 — `private void udDSPNB_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNB` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udLMSNRgain_LostFocus()`** — L11940 — `private void udLMSNRgain_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udLMSNRgain` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udLMSNRdelay_LostFocus()`** — L11945 — `private void udLMSNRdelay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udLMSNRdelay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udLMSNRtaps_LostFocus()`** — L11950 — `private void udLMSNRtaps_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udLMSNRtaps` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udLMSANFgain_LostFocus()`** — L11955 — `private void udLMSANFgain_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udLMSANFgain` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udLMSANFdelay_LostFocus()`** — L11960 — `private void udLMSANFdelay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udLMSANFdelay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udLMSANFtaps_LostFocus()`** — L11965 — `private void udLMSANFtaps_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udLMSANFtaps` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPCWPitch_LostFocus()`** — L11970 — `private void udDSPCWPitch_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPCWPitch` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCWKeyerWeight_LostFocus()`** — L11975 — `private void udCWKeyerWeight_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udCWKeyerWeight` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCWBreakInDelay_LostFocus()`** — L11980 — `private void udCWBreakInDelay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udCWBreakInDelay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPLevelerThreshold_LostFocus()`** — L11985 — `private void udDSPLevelerThreshold_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPLevelerThreshold` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPLevelerDecay_LostFocus()`** — L11990 — `private void udDSPLevelerDecay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPLevelerDecay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPALCThreshold_LostFocus()`** — L11995 — `private void udDSPALCThreshold_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPALCThreshold` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPALCDecay_LostFocus()`** — L12000 — `private void udDSPALCDecay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPALCDecay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPAGCHangTime_LostFocus()`** — L12005 — `private void udDSPAGCHangTime_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPAGCHangTime` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPAGCMaxGaindB_LostFocus()`** — L12010 — `private void udDSPAGCMaxGaindB_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPAGCMaxGaindB` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPAGCSlope_LostFocus()`** — L12015 — `private void udDSPAGCSlope_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPAGCSlope` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPAGCDecay_LostFocus()`** — L12020 — `private void udDSPAGCDecay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPAGCDecay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPAGCFixedGaindB_LostFocus()`** — L12025 — `private void udDSPAGCFixedGaindB_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPAGCFixedGaindB` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXAMCarrierLevel_LostFocus()`** — L12030 — `private void udTXAMCarrierLevel_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXAMCarrierLevel` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXAF_LostFocus()`** — L12035 — `private void udTXAF_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXAF` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXTunePower_LostFocus()`** — L12040 — `private void udTXTunePower_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXTunePower` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterLow_LostFocus()`** — L12045 — `private void udTXFilterLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXFilterLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterHigh_LostFocus()`** — L12050 — `private void udTXFilterHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXFilterHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udMicGainMax_LostFocus()`** — L12055 — `private void udMicGainMax_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udMicGainMax` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udMicGainMin_LostFocus()`** — L12060 — `private void udMicGainMin_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udMicGainMin` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayLineWidth_LostFocus()`** — L12065 — `private void udDisplayLineWidth_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayLineWidth` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXLineWidth_LostFocus()`** — L12070 — `private void udTXLineWidth_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXLineWidth` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenScale_LostFocus()`** — L12075 — `private void udTXGenScale_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenScale` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenSweepRate_LostFocus()`** — L12080 — `private void udTXGenSweepRate_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenSweepRate` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenSweepHigh_LostFocus()`** — L12085 — `private void udTXGenSweepHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenSweepHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenSweepLow_LostFocus()`** — L12090 — `private void udTXGenSweepLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenSweepLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTestIMDFreq2_LostFocus()`** — L12095 — `private void udTestIMDFreq2_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTestIMDFreq2` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTestIMDPower_LostFocus()`** — L12100 — `private void udTestIMDPower_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTestIMDPower` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTestIMDFreq1_LostFocus()`** — L12105 — `private void udTestIMDFreq1_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTestIMDFreq1` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowFreqOffset_CheckedChanged()`** — L12112 — `private void chkShowFreqOffset_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowFreqOffset` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowZeroLine_CheckedChanged()`** — L12117 — `private void chkShowZeroLine_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowZeroLine` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnBandEdge_Changed()`** — L12122 — `private void clrbtnBandEdge_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnTXBandEdge_Changed()`** — L12128 — `private void clrbtnTXBandEdge_Changed(object sender, System.EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.comboMeterType_SelectedIndexChanged()`** — L12134 — `private void comboMeterType_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboMeterType` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnMeterEdgeLow_Changed()`** — L12152 — `private void clrbtnMeterEdgeLow_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterEdgeHigh_Changed()`** — L12157 — `private void clrbtnMeterEdgeHigh_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterEdgeBackground_Changed()`** — L12162 — `private void clrbtnMeterEdgeBackground_Changed(object sender, System.EventArgs e)`
  Called by: `.tbMeterEdgeBackgroundAlpha_Scroll()` (same file)
- **`.clrbtnEdgeIndicator_Changed()`** — L12167 — `private void clrbtnEdgeIndicator_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterDigText_Changed()`** — L12172 — `private void clrbtnMeterDigText_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterDigBackground_Changed()`** — L12177 — `private void clrbtnMeterDigBackground_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnSubRXFilter_Changed()`** — L12182 — `private void clrbtnSubRXFilter_Changed(object sender, System.EventArgs e)`
  Called by: `.tbMultiRXFilterAlpha_Scroll()` (same file)
- **`.clrbtnSubRXZero_Changed()`** — L12187 — `private void clrbtnSubRXZero_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkCWKeyerMode_CheckedChanged()`** — L12192 — `private void chkCWKeyerMode_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWKeyerMode` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDisableToolTips_CheckedChanged()`** — L12202 — `private void chkDisableToolTips_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDisableToolTips` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboColorPalette_SelectedIndexChanged()`** — L12208 — `private void comboColorPalette_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboColorPalette` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.showHideWaterfallControls()`** — L12253 — `private void showHideWaterfallControls(int rx, bool show)`
  Called by: `.comboColorPalette_SelectedIndexChanged()` (same file), `.comboRX2ColorPalette_SelectedIndexChanged()` (same file)
- **`.comboRX2ColorPalette_SelectedIndexChanged()`** — L12275 — `private void comboRX2ColorPalette_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRX2ColorPalette` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setWaterFallCalculatedDelayText()`** — L12321 — `private void setWaterFallCalculatedDelayText()`
  Sets water fall calculated delay text.
  Called by: `.udDisplayFPS_ValueChanged()` (same file), `.udDisplayWaterfallUpdatePeriod_ValueChanged()` (same file), `.udRX2DisplayWaterfallUpdatePeriod_ValueChanged()` (same file)
- **`.udDisplayWaterfallUpdatePeriod_ValueChanged()`** — L12327 — `private void udDisplayWaterfallUpdatePeriod_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDisplayWaterfallUpdatePeriod` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2DisplayWaterfallUpdatePeriod_ValueChanged()`** — L12334 — `private void udRX2DisplayWaterfallUpdatePeriod_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayWaterfallUpdatePeriod` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSnapClickTune_CheckedChanged()`** — L12341 — `private void chkSnapClickTune_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkSnapClickTune` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkClickTuneFilter_CheckedChanged()`** — L12346 — `private void chkClickTuneFilter_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkClickTuneFilter` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowCTHLine_CheckedChanged()`** — L12353 — `private void chkShowCTHLine_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowCTHLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radPACalAllBands_CheckedChanged()`** — L12359 — `private void radPACalAllBands_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radPACalAllBands` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkZeroBeatRIT_CheckedChanged()`** — L12370 — `private void chkZeroBeatRIT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkZeroBeatRIT` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPANewCal_CheckedChanged()`** — L12375 — `private void chkPANewCal_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPANewCal` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udMeterDigitalDelay_ValueChanged()`** — L12404 — `private void udMeterDigitalDelay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udMeterDigitalDelay` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMouseTuneStep_CheckedChanged()`** — L12409 — `private void chkMouseTuneStep_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMouseTuneStep` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWheelReverse_CheckedChanged()`** — L12414 — `private void chkWheelReverse_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWheelReverse` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtGenCustomTitle_TextChanged()`** — L12432 — `private void txtGenCustomTitle_TextChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtGenCustomTitle` text changes.
  Called by: `.UpdateGeneraHardware()` (same file), `.chkDisplayIPPort_CheckedChanged()` (same file)
- **`.chkGenAllModeMicPTT_CheckedChanged()`** — L12437 — `private void chkGenAllModeMicPTT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkGenAllModeMicPTT` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkKWAI_CheckedChanged()`** — L12442 — `private void chkKWAI_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkKWAI` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSplitOff_CheckedChanged()`** — L12447 — `private void chkSplitOff_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkSplitOff` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnableRFEPATR_CheckedChanged()`** — L12458 — `private void chkEnableRFEPATR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnableRFEPATR` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVACAllowBypass_CheckedChanged()`** — L12463 — `private void chkVACAllowBypass_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVACAllowBypass` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSPACEAllowBypass_CheckedChanged()`** — L12468 — `private void chkSPACEAllowBypass_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkSPACEAllowBypass` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMOXAllowBypass_CheckedChanged()`** — L12473 — `private void chkMOXAllowBypass_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMOXAllowBypass` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDSPTXMeterPeak_CheckedChanged()`** — L12478 — `private void chkDSPTXMeterPeak_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDSPTXMeterPeak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVACCombine_CheckedChanged()`** — L12483 — `private void chkVACCombine_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVACCombine` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWAutoSwitchMode_CheckedChanged()`** — L12488 — `private void chkCWAutoSwitchMode_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWAutoSwitchMode` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnGenBackground_Changed()`** — L12495 — `private void clrbtnGenBackground_Changed(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboTXTUNMeter_SelectedIndexChanged()`** — L12526 — `private void comboTXTUNMeter_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboTXTUNMeter` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisplayMeterShowDecimal_CheckedChanged()`** — L12548 — `private void chkDisplayMeterShowDecimal_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDisplayMeterShowDecimal` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRTTYOffsetEnableA_CheckedChanged()`** — L12553 — `private void chkRTTYOffsetEnableA_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRTTYOffsetEnableA` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRTTYOffsetEnableB_CheckedChanged()`** — L12558 — `private void chkRTTYOffsetEnableB_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRTTYOffsetEnableB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRTTYL_ValueChanged()`** — L12563 — `private void udRTTYL_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRTTYL` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRTTYU_ValueChanged()`** — L12568 — `private void udRTTYU_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRTTYU` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2AutoMuteTX_CheckedChanged()`** — L12573 — `private void chkRX2AutoMuteTX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2AutoMuteTX` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudioIQtoVAC_CheckedChanged()`** — L12580 — `private void chkAudioIQtoVAC_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAudioIQtoVAC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2DirectIQ_CheckedChanged()`** — L12601 — `private void chkVAC2DirectIQ_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC2DirectIQ` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAudioCorrectIQ_CheckChanged()`** — L12621 — `private void chkAudioCorrectIQ_CheckChanged(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkVAC2IQCal_CheckChanged()`** — L12626 — `private void chkVAC2IQCal_CheckChanged(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2AutoMuteRX1OnVFOBTX_CheckedChanged()`** — L12631 — `private void chkRX2AutoMuteRX1OnVFOBTX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2AutoMuteRX1OnVFOBTX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX1BlankDisplayOnVFOBTX_CheckedChanged()`** — L12636 — `private void chkRX1BlankDisplayOnVFOBTX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX1BlankDisplayOnVFOBTX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2BlankDisplayOnVFOATX_CheckedChanged()`** — L12641 — `private void chkRX2BlankDisplayOnVFOATX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2BlankDisplayOnVFOATX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXExpert_CheckedChanged()`** — L12646 — `private void chkTXExpert_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTXExpert` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnTXProfileDefImport_Click()`** — L12651 — `private void btnTXProfileDefImport_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnTXProfileDefImport` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ExportCurrentTxProfile()`** — L12720 — `private void ExportCurrentTxProfile()`
  -W2PA Export a single TX Profile to send to someone else for importing.
  Called by: `.btnExportCurrentTXProfile_Click()` (same file)
- **`.Setup_KeyDown()`** — L12794 — `private void Setup_KeyDown(object sender, System.Windows.Forms.KeyEventArgs e)`
  WinForms event handler: runs when `Setup` receives a key-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisplayPanFill_CheckedChanged()`** — L12817 — `private void chkDisplayPanFill_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDisplayPanFill` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXPanFill_CheckedChanged()`** — L12826 — `private void chkTXPanFill_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTXPanFill` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAppSkin_SelectedIndexChanged()`** — L12832 — `private void comboAppSkin_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboAppSkin` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnSkinExport_Click()`** — L12869 — `private void btnSkinExport_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSkinExport` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWDisableUI_CheckedChanged()`** — L12879 — `private void chkCWDisableUI_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCWDisableUI` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAudioRX2toVAC_CheckedChanged()`** — L12884 — `private void chkAudioRX2toVAC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAudioRX2toVAC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2UseRX2_CheckedChanged()`** — L12890 — `private void chkVAC2UseRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2UseRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX1FilterAlpha_Scroll()`** — L12896 — `private void tbRX1FilterAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1FilterAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbTXFilterAlpha_Scroll()`** — L12902 — `private void tbTXFilterAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXFilterAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbMultiRXFilterAlpha_Scroll()`** — L12908 — `private void tbMultiRXFilterAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbMultiRXFilterAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.chkWheelTuneVFOB_CheckedChanged()`** — L12914 — `private void chkWheelTuneVFOB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWheelTuneVFOB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.checkHPSDRDefaults()`** — L12919 — `private void checkHPSDRDefaults(object sender, System.EventArgs e)`
  Called by: `.AfterConstructor()` (same file)
- **`.chkAlexPresent_CheckedChanged()`** — L12923 — `private void chkAlexPresent_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlexPresent` checked state changes.
  Called by: `.AfterConstructor()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.UpdateOCBits()`** — L12948 — `private void UpdateOCBits()`
  Updates ocbits.
  Called by: `.radSplitPins4x3_CheckedChanged()` (same file), `.radSplitPins3x4_CheckedChanged()` (same file)
- **`.chkPenOCrcv160_CheckedChanged()`** — L12976 — `private void chkPenOCrcv160_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv160` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit160_CheckedChanged()`** — L12992 — `private void chkPenOCxmit160_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit160` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv80_CheckedChanged()`** — L13008 — `private void chkPenOCrcv80_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv80` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit80_CheckedChanged()`** — L13024 — `private void chkPenOCxmit80_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit80` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv60_CheckedChanged()`** — L13040 — `private void chkPenOCrcv60_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv60` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit60_CheckedChanged()`** — L13056 — `private void chkPenOCxmit60_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit60` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv40_CheckedChanged()`** — L13072 — `private void chkPenOCrcv40_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv40` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit40_CheckedChanged()`** — L13089 — `private void chkPenOCxmit40_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit40` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv30_CheckedChanged()`** — L13106 — `private void chkPenOCrcv30_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv30` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit30_CheckedChanged()`** — L13123 — `private void chkPenOCxmit30_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit30` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv20_CheckedChanged()`** — L13140 — `private void chkPenOCrcv20_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv20` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit20_CheckedChanged()`** — L13157 — `private void chkPenOCxmit20_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit20` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv17_CheckedChanged()`** — L13174 — `private void chkPenOCrcv17_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv17` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit17_CheckedChanged()`** — L13191 — `private void chkPenOCxmit17_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit17` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv15_CheckedChanged()`** — L13208 — `private void chkPenOCrcv15_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv15` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit15_CheckedChanged()`** — L13225 — `private void chkPenOCxmit15_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit15` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv12_CheckedChanged()`** — L13243 — `private void chkPenOCrcv12_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv12` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit12_CheckedChanged()`** — L13260 — `private void chkPenOCxmit12_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit12` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv10_CheckedChanged()`** — L13277 — `private void chkPenOCrcv10_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv10` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit10_CheckedChanged()`** — L13294 — `private void chkPenOCxmit10_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit10` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv6_CheckedChanged()`** — L13312 — `private void chkPenOCrcv6_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv6` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit6_CheckedChanged()`** — L13329 — `private void chkPenOCxmit6_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit6` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCrcv2_CheckedChanged()`** — L13346 — `private void chkPenOCrcv2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcv2` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPenOCxmit2_CheckedChanged()`** — L13363 — `private void chkPenOCxmit2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmit2` checked state changes.
  Called by: `.UpdateOCBits()` (same file)
- **`.chkPennyExtCtrl_CheckedChanged()`** — L13380 — `private void chkPennyExtCtrl_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPennyExtCtrl` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlexAntCtrl_CheckedChanged()`** — L13388 — `private void chkAlexAntCtrl_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlexAntCtrl` checked state changes.
  Called by: `.AfterConstructor()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.chkMercDither_CheckedChanged()`** — L13396 — `private void chkMercDither_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMercDither` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkMercRandom_CheckedChanged()`** — L13416 — `private void chkMercRandom_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMercRandom` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkHL2BandVolts_CheckedChanged()`** — L13448 — `private void chkHL2BandVolts_CheckedChanged(object sender, System.EventArgs e)`
  MI0BOT: Control band volts for the HL2
  Called by: `.ForceAllEvents()` (same file)
- **`.chkHL2PsSync_CheckedChanged()`** — L13456 — `private void chkHL2PsSync_CheckedChanged(object sender, System.EventArgs e)`
  MI0BOT: Control power supply sync for the HL2
  Called by: `.PerformDelayedInitalistion()` (same file), `.ForceAllEvents()` (same file)
- **`.InitAlexAntTables()`** — L13464 — `private void InitAlexAntTables()`
  Inits alex ant tables.
  Called by: `.AfterConstructor()` (same file)
- **`.SetAlexAntEnabled()`** — L13551 — `public bool SetAlexAntEnabled(bool state)`
  Sets alex ant enabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radAlexR_160_CheckedChanged()`** — L13558 — `private void radAlexR_160_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_160` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_80_CheckedChanged()`** — L13585 — `private void radAlexR_80_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_80` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_60_CheckedChanged()`** — L13612 — `private void radAlexR_60_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_60` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_40_CheckedChanged()`** — L13639 — `private void radAlexR_40_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_40` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_30_CheckedChanged()`** — L13666 — `private void radAlexR_30_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_30` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_20_CheckedChanged()`** — L13693 — `private void radAlexR_20_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_20` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_17_CheckedChanged()`** — L13720 — `private void radAlexR_17_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_17` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_15_CheckedChanged()`** — L13747 — `private void radAlexR_15_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_15` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_12_CheckedChanged()`** — L13774 — `private void radAlexR_12_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_12` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_10_CheckedChanged()`** — L13801 — `private void radAlexR_10_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_10` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexR_6_CheckedChanged()`** — L13828 — `private void radAlexR_6_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexR_6` checked state changes.
  Called by: `.chkBlockTxAnt2_CheckedChanged()` (same file), `.chkBlockTxAnt3_CheckedChanged()` (same file)
- **`.radAlexT_160_CheckedChanged()`** — L13855 — `private void radAlexT_160_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_160` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_80_CheckedChanged()`** — L13860 — `private void radAlexT_80_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_80` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_60_CheckedChanged()`** — L13865 — `private void radAlexT_60_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_60` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_40_CheckedChanged()`** — L13870 — `private void radAlexT_40_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_40` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_30_CheckedChanged()`** — L13875 — `private void radAlexT_30_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_30` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_20_CheckedChanged()`** — L13880 — `private void radAlexT_20_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_20` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_17_CheckedChanged()`** — L13885 — `private void radAlexT_17_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_17` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_15_CheckedChanged()`** — L13890 — `private void radAlexT_15_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_15` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_12_CheckedChanged()`** — L13895 — `private void radAlexT_12_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_12` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_10_CheckedChanged()`** — L13900 — `private void radAlexT_10_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_10` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radAlexT_6_CheckedChanged()`** — L13905 — `private void radAlexT_6_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radAlexT_6` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex160R_CheckedChanged()`** — L13910 — `private void chkAlex160R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex160R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex80R_CheckedChanged()`** — L13915 — `private void chkAlex80R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex80R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex60R_CheckedChanged()`** — L13920 — `private void chkAlex60R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex60R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex40R_CheckedChanged()`** — L13925 — `private void chkAlex40R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex40R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex30R_CheckedChanged()`** — L13930 — `private void chkAlex30R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex30R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex20R_CheckedChanged()`** — L13935 — `private void chkAlex20R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex20R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex17R_CheckedChanged()`** — L13940 — `private void chkAlex17R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex17R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex15R_CheckedChanged()`** — L13945 — `private void chkAlex15R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex15R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex12R_CheckedChanged()`** — L13950 — `private void chkAlex12R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex12R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex10R_CheckedChanged()`** — L13955 — `private void chkAlex10R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex10R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex6R_CheckedChanged()`** — L13960 — `private void chkAlex6R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex6R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex2R_CheckedChanged()`** — L13965 — `private void chkAlex2R_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlex2R` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ProcessAlexAntCheckBox()`** — L13970 — `private void ProcessAlexAntCheckBox(object sender, Band band)`
  Processes alex ant check box.
  Called by: `.chkAlex160R_CheckedChanged()` (same file), `.chkAlex80R_CheckedChanged()` (same file), `.chkAlex60R_CheckedChanged()` (same file), `.chkAlex40R_CheckedChanged()` (same file), `.chkAlex30R_CheckedChanged()` (same file), `.chkAlex20R_CheckedChanged()` (same file) — and 6 more
- **`.ProcessAlexAntRadioButton()`** — L14068 — `private void ProcessAlexAntRadioButton(object sender, Band band, bool is_xmit)`
  Processes alex ant radio button.
  Called by: `.radAlexR_160_CheckedChanged()` (same file), `.radAlexR_80_CheckedChanged()` (same file), `.radAlexR_60_CheckedChanged()` (same file), `.radAlexR_40_CheckedChanged()` (same file), `.radAlexR_30_CheckedChanged()` (same file), `.radAlexR_20_CheckedChanged()` (same file) — and 16 more
- **`.updateChangedAntAlexButton()`** — L14124 — `private void updateChangedAntAlexButton(bool is_xmit, int idx, Band band)`
  Called by: `.ProcessAlexAntRadioButton()` (same file)
- **`.GetRXAntenna()`** — L14153 — `public int GetRXAntenna(Band band)`
  get RX antenna in use for band this must also ignores ext input check boxes
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTXAntenna()`** — L14170 — `public int GetTXAntenna(Band band)`
  get TX antenna in use for band
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAntenna()`** — L14200 — `public void SetRXAntenna(int Antenna, Band band)`
  set RX antenna to new antenna 1-3 this must also cancel any ext input check boxes
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAuxAntenna()`** — L14226 — `public void SetAuxAntenna(int Antenna, Band band, bool byp, bool ext1)`
  set RX antenna to new byp/ext1 this must also cancel the other
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAntenna()`** — L14242 — `public void SetTXAntenna(int Antenna, Band band)`
  set TX antenna to new antenna 1-3 only select antennas 2 or 3 if not blocked by "do not TX" check boxes
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleRXAntennaChangeForNF()`** — L14275 — `private void handleRXAntennaChangeForNF(Band band)`
  Called by: `.ProcessAlexAntCheckBox()` (same file), `.ProcessAlexAntRadioButton()` (same file)
- **`.btnHPSDRFreqCalReset_Click()`** — L14291 — `private void btnHPSDRFreqCalReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnHPSDRFreqCalReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tpHPSDR_Paint()`** — L14297 — `private void tpHPSDR_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `tpHPSDR` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tpGeneralHardware_Paint()`** — L14309 — `private void tpGeneralHardware_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `tpGeneralHardware` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.GetFirmwareCodeVersionString()`** — L14314 — `public string GetFirmwareCodeVersionString()`
  MW0LGE_21g
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateGeneraHardware()`** — L14356 — `public void UpdateGeneraHardware()`
  Updates genera hardware.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.udMaxFreq_ValueChanged()`** — L14364 — `private void udMaxFreq_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udMaxFreq` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udHPSDRFreqCorrectFactor_ValueChanged()`** — L14369 — `private void udHPSDRFreqCorrectFactor_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udHPSDRFreqCorrectFactor` value changes.
  Called by: `.btnHPSDRFreqCalReset_Click()` (same file), `.chkUsing10MHzRef_CheckedChanged()` (same file), `.udHPSDRFreqCorrectFactor10MHz_ValueChanged()` (same file)
- **`.chkHERCULES_CheckedChanged()`** — L14382 — `private void chkHERCULES_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHERCULES` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnPennyCtrlReset_Click()`** — L14497 — `private void btnPennyCtrlReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPennyCtrlReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnPennyCtrlVHFReset_Click()`** — L14513 — `private void btnPennyCtrlVHFReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPennyCtrlVHFReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCtrlSWLReset_Click()`** — L14524 — `private void btnCtrlSWLReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCtrlSWLReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboFRSRegion_SelectedIndexChanged()`** — L14535 — `private void comboFRSRegion_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFRSRegion` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMicIn_CheckedChanged()`** — L14638 — `private void radMicIn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMicIn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radLineIn_CheckedChanged()`** — L14650 — `private void radLineIn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radLineIn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.IPStringFromInt()`** — L14662 — `private string IPStringFromInt(Int32 addr, StringBuilder sb)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.udGenPTTOutDelay_ValueChanged()`** — L14683 — `private void udGenPTTOutDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udGenPTTOutDelay` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udMoxDelay_ValueChanged()`** — L14688 — `private void udMoxDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udMoxDelay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udCWKeyUpDelay_ValueChanged()`** — L14694 — `private void udCWKeyUpDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udCWKeyUpDelay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRFDelay_ValueChanged()`** — L14701 — `private void udRFDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRFDelay` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXDisplayCalOffset_ValueChanged()`** — L14706 — `private void udTXDisplayCalOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXDisplayCalOffset` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXDisplayCalOffset_LostFocus()`** — L14712 — `private void udTXDisplayCalOffset_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXDisplayCalOffset` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTwoToneLevel_ValueChanged()`** — L14717 — `private void udTwoToneLevel_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTwoToneLevel` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnGridFine_Changed()`** — L14726 — `private void clrbtnGridFine_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbGridFineAlpha_Scroll()` (same file)
- **`.clrbtnTXVGridFine_Changed()`** — L14732 — `private void clrbtnTXVGridFine_Changed(object sender, EventArgs e)`
  Called by: `.tbTXVGridFineAlpha_Scroll()` (same file)
- **`.tbBackgroundAlpha_Scroll()`** — L14737 — `private void tbBackgroundAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbBackgroundAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbTXBackgroundAlpha_Scroll()`** — L14743 — `private void tbTXBackgroundAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXBackgroundAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbGridCourseAlpha_Scroll()`** — L14749 — `private void tbGridCourseAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbGridCourseAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbTXVGridCourseAlpha_Scroll()`** — L14755 — `private void tbTXVGridCourseAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXVGridCourseAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbGridFineAlpha_Scroll()`** — L14761 — `private void tbGridFineAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbGridFineAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbTXVGridFineAlpha_Scroll()`** — L14767 — `private void tbTXVGridFineAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXVGridFineAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.clrbtnHGridColor_Changed()`** — L14773 — `private void clrbtnHGridColor_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbHGridColorAlpha_Scroll()` (same file)
- **`.clrbtnTXHGridColor_Changed()`** — L14779 — `private void clrbtnTXHGridColor_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbTXHGridColorAlpha_Scroll()` (same file)
- **`.tbHGridColorAlpha_Scroll()`** — L14785 — `private void tbHGridColorAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbHGridColorAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbMeterEdgeBackgroundAlpha_Scroll()`** — L14791 — `private void tbMeterEdgeBackgroundAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbMeterEdgeBackgroundAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbTXHGridColorAlpha_Scroll()`** — L14797 — `private void tbTXHGridColorAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXHGridColorAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbTXZeroLineAlpha_Scroll()`** — L14803 — `private void tbTXZeroLineAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXZeroLineAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.tbTXTextAlpha_Scroll()`** — L14809 — `private void tbTXTextAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXTextAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.chkGridControl_CheckedChanged()`** — L14815 — `private void chkGridControl_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkGridControl` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkGridControl_minor_CheckedChanged()`** — L14820 — `private void chkGridControl_minor_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkGridControl_minor` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkTXGridControl_CheckedChanged()`** — L14825 — `private void chkTXGridControl_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTXGridControl` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radSpaceBarLastBtn_CheckedChanged()`** — L14830 — `private void radSpaceBarLastBtn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSpaceBarLastBtn` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radSpaceBarPTT_CheckedChanged()`** — L14835 — `private void radSpaceBarPTT_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSpaceBarPTT` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radSpaceBarVOX_CheckedChanged()`** — L14840 — `private void radSpaceBarVOX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSpaceBarVOX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radSpaceBarMicMute_CheckedChanged()`** — L14845 — `private void radSpaceBarMicMute_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSpaceBarMicMute` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF0_CheckedChanged()`** — L14850 — `private void chkPenOCrcvVHF0_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF0` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF1_CheckedChanged()`** — L14866 — `private void chkPenOCrcvVHF1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF2_CheckedChanged()`** — L14882 — `private void chkPenOCrcvVHF2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF3_CheckedChanged()`** — L14898 — `private void chkPenOCrcvVHF3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF4_CheckedChanged()`** — L14914 — `private void chkPenOCrcvVHF4_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF4` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF5_CheckedChanged()`** — L14930 — `private void chkPenOCrcvVHF5_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF5` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF6_CheckedChanged()`** — L14946 — `private void chkPenOCrcvVHF6_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF6` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF7_CheckedChanged()`** — L14962 — `private void chkPenOCrcvVHF7_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF7` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF8_CheckedChanged()`** — L14978 — `private void chkPenOCrcvVHF8_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF8` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF9_CheckedChanged()`** — L14994 — `private void chkPenOCrcvVHF9_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF9` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF10_CheckedChanged()`** — L15010 — `private void chkPenOCrcvVHF10_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF10` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCrcvVHF11_CheckedChanged()`** — L15026 — `private void chkPenOCrcvVHF11_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCrcvVHF11` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF0_CheckedChanged()`** — L15042 — `private void chkPenOCxmitVHF0_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF0` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF1_CheckedChanged()`** — L15058 — `private void chkPenOCxmitVHF1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF2_CheckedChanged()`** — L15074 — `private void chkPenOCxmitVHF2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF3_CheckedChanged()`** — L15090 — `private void chkPenOCxmitVHF3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF4_CheckedChanged()`** — L15106 — `private void chkPenOCxmitVHF4_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF4` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF5_CheckedChanged()`** — L15122 — `private void chkPenOCxmitVHF5_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF5` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF6_CheckedChanged()`** — L15138 — `private void chkPenOCxmitVHF6_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF6` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF7_CheckedChanged()`** — L15154 — `private void chkPenOCxmitVHF7_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF7` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF8_CheckedChanged()`** — L15170 — `private void chkPenOCxmitVHF8_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF8` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF9_CheckedChanged()`** — L15186 — `private void chkPenOCxmitVHF9_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF9` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF10_CheckedChanged()`** — L15202 — `private void chkPenOCxmitVHF10_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF10` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPenOCxmitVHF11_CheckedChanged()`** — L15218 — `private void chkPenOCxmitVHF11_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPenOCxmitVHF11` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcvLMW_CheckedChanged()`** — L15234 — `private void chkOCrcvLMW_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcvLMW` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv120_CheckedChanged()`** — L15250 — `private void chkOCrcv120_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv120` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv90_CheckedChanged()`** — L15266 — `private void chkOCrcv90_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv90` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv61_CheckedChanged()`** — L15282 — `private void chkOCrcv61_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv61` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv49_CheckedChanged()`** — L15298 — `private void chkOCrcv49_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv49` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv41_CheckedChanged()`** — L15314 — `private void chkOCrcv41_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv41` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv31_CheckedChanged()`** — L15330 — `private void chkOCrcv31_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv31` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv25_CheckedChanged()`** — L15346 — `private void chkOCrcv25_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv25` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv22_CheckedChanged()`** — L15362 — `private void chkOCrcv22_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv22` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv19_CheckedChanged()`** — L15378 — `private void chkOCrcv19_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv19` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv16_CheckedChanged()`** — L15394 — `private void chkOCrcv16_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv16` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv14_CheckedChanged()`** — L15410 — `private void chkOCrcv14_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv14` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv13_CheckedChanged()`** — L15426 — `private void chkOCrcv13_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv13` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCrcv11_CheckedChanged()`** — L15442 — `private void chkOCrcv11_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCrcv11` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmitLMW_CheckedChanged()`** — L15458 — `private void chkOCxmitLMW_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmitLMW` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit120_CheckedChanged()`** — L15474 — `private void chkOCxmit120_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit120` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit90_CheckedChanged()`** — L15490 — `private void chkOCxmit90_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit90` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit61_CheckedChanged()`** — L15506 — `private void chkOCxmit61_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit61` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit49_CheckedChanged()`** — L15522 — `private void chkOCxmit49_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit49` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit41_CheckedChanged()`** — L15538 — `private void chkOCxmit41_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit41` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit31_CheckedChanged()`** — L15554 — `private void chkOCxmit31_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit31` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit25_CheckedChanged()`** — L15570 — `private void chkOCxmit25_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit25` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit22_CheckedChanged()`** — L15586 — `private void chkOCxmit22_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit22` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit19_CheckedChanged()`** — L15602 — `private void chkOCxmit19_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit19` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit16_CheckedChanged()`** — L15618 — `private void chkOCxmit16_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit16` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit14_CheckedChanged()`** — L15634 — `private void chkOCxmit14_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit14` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit13_CheckedChanged()`** — L15650 — `private void chkOCxmit13_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit13` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOCxmit11_CheckedChanged()`** — L15666 — `private void chkOCxmit11_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOCxmit11` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisable6mLNAonTX_CheckedChanged()`** — L15682 — `private void chkDisable6mLNAonTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisable6mLNAonTX` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkDisable6mLNAonRX_CheckedChanged()`** — L15688 — `private void chkDisable6mLNAonRX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisable6mLNAonRX` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkDisableHPFonTX_CheckedChanged()`** — L15694 — `private void chkDisableHPFonTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisableHPFonTX` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlexHPFBypass_CheckedChanged()`** — L15716 — `private void chkAlexHPFBypass_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlexHPFBypass` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlex2HPFBypass_CheckedChanged()`** — L15737 — `private void chkAlex2HPFBypass_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex2HPFBypass` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowAGC_CheckedChanged()`** — L15742 — `private void chkShowAGC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowAGC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSpectrumLine_CheckedChanged()`** — L15748 — `private void chkSpectrumLine_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSpectrumLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAGCDisplayHangLine_CheckedChanged()`** — L15754 — `private void chkAGCDisplayHangLine_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAGCDisplayHangLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAGCHangSpectrumLine_CheckedChanged()`** — L15760 — `private void chkAGCHangSpectrumLine_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAGCHangSpectrumLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDisplayRX2GainLine_CheckedChanged()`** — L15766 — `private void chkDisplayRX2GainLine_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisplayRX2GainLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkRX2GainSpectrumLine_CheckedChanged()`** — L15772 — `private void chkRX2GainSpectrumLine_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2GainSpectrumLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDisplayRX2HangLine_CheckedChanged()`** — L15778 — `private void chkDisplayRX2HangLine_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisplayRX2HangLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkRX2HangSpectrumLine_CheckedChanged()`** — L15784 — `private void chkRX2HangSpectrumLine_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2HangSpectrumLine` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkFirmwareByp_CheckedChanged()`** — L15790 — `private void chkFirmwareByp_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFirmwareByp` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkStrictCharSpacing_CheckedChanged()`** — L15795 — `private void chkStrictCharSpacing_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkStrictCharSpacing` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkRxOutOnTx_CheckedChanged()`** — L15801 — `private void chkRxOutOnTx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRxOutOnTx` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSWRProtection_CheckedChanged()`** — L15814 — `private void chkSWRProtection_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSWRProtection` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkATTOnTX_CheckedChanged()`** — L15833 — `private void chkATTOnTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkATTOnTX` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex1_5BPHPF_CheckedChanged()`** — L15839 — `private void chkAlex1_5BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex1_5BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex6_5BPHPF_CheckedChanged()`** — L15845 — `private void chkAlex6_5BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex6_5BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex9_5BPHPF_CheckedChanged()`** — L15851 — `private void chkAlex9_5BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex9_5BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex13BPHPF_CheckedChanged()`** — L15857 — `private void chkAlex13BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex13BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex20BPHPF_CheckedChanged()`** — L15863 — `private void chkAlex20BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex20BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex6BPHPF_CheckedChanged()`** — L15869 — `private void chkAlex6BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex6BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex21_5BPHPF_CheckedChanged()`** — L15875 — `private void chkAlex21_5BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex21_5BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex26_5BPHPF_CheckedChanged()`** — L15881 — `private void chkAlex26_5BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex26_5BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex29_5BPHPF_CheckedChanged()`** — L15887 — `private void chkAlex29_5BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex29_5BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex213BPHPF_CheckedChanged()`** — L15893 — `private void chkAlex213BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex213BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex220BPHPF_CheckedChanged()`** — L15899 — `private void chkAlex220BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex220BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkAlex26BPHPF_CheckedChanged()`** — L15905 — `private void chkAlex26BPHPF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlex26BPHPF` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.tpPennyCtrl_Paint()`** — L15911 — `private void tpPennyCtrl_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `tpPennyCtrl` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkApolloPresent_CheckedChanged()`** — L15947 — `private void chkApolloPresent_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkApolloPresent` checked state changes.
  Called by: `.AfterConstructor()` (same file)
- **`.chkApolloFilter_CheckedChanged()`** — L15963 — `private void chkApolloFilter_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkApolloFilter` checked state changes.
  Called by: `.chkApolloPresent_CheckedChanged()` (same file)
- **`.chkApolloTuner_CheckedChanged()`** — L15969 — `private void chkApolloTuner_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkApolloTuner` checked state changes.
  Called by: `.chkApolloPresent_CheckedChanged()` (same file)
- **`.chkLevelFades_CheckedChanged()`** — L15974 — `private void chkLevelFades_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLevelFades` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radLSBUSB_CheckedChanged()`** — L15989 — `private void radLSBUSB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radLSBUSB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radLSB_CheckedChanged()`** — L15999 — `private void radLSB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radLSB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radUSB_CheckedChanged()`** — L16009 — `private void radUSB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUSB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkRX2LevelFades_CheckedChanged()`** — L16019 — `private void chkRX2LevelFades_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2LevelFades` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radRX2LSBUSB_CheckedChanged()`** — L16034 — `private void radRX2LSBUSB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX2LSBUSB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radRX2LSB_CheckedChanged()`** — L16044 — `private void radRX2LSB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX2LSB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radRX2USB_CheckedChanged()`** — L16054 — `private void radRX2USB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX2USB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAutoPACalibrate_CheckedChanged()`** — L16064 — `private void chkAutoPACalibrate_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoPACalibrate` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateAttenuationInfo()`** — L16086 — `private void updateAttenuationInfo()`
  Called by: `.AfterConstructor()` (same file)
- **`.chkHermesStepAttenuator_CheckedChanged()`** — L16111 — `private void chkHermesStepAttenuator_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHermesStepAttenuator` checked state changes.
  Called by: `.updateAttenuationInfo()` (same file)
- **`.udHermesStepAttenuatorData_ValueChanged()`** — L16146 — `private void udHermesStepAttenuatorData_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udHermesStepAttenuatorData` value changes.
  Called by: `.chkAlexPresent_CheckedChanged()` (same file), `.chkHermesStepAttenuator_CheckedChanged()` (same file)
- **`.chkRX2StepAtt_CheckedChanged()`** — L16175 — `private void chkRX2StepAtt_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2StepAtt` checked state changes.
  Called by: `.updateAttenuationInfo()` (same file), `.chkHermesStepAttenuator_CheckedChanged()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.udHermesStepAttenuatorDataRX2_ValueChanged()`** — L16214 — `private void udHermesStepAttenuatorDataRX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udHermesStepAttenuatorDataRX2` value changes.
  Called by: `.chkRX2StepAtt_CheckedChanged()` (same file)
- **`.udAlex160mLPFStart_ValueChanged()`** — L16239 — `private void udAlex160mLPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex160mLPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex160mLPFEnd_ValueChanged()`** — L16247 — `private void udAlex160mLPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex160mLPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex80mLPFStart_ValueChanged()`** — L16259 — `private void udAlex80mLPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex80mLPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex80mLPFEnd_ValueChanged()`** — L16267 — `private void udAlex80mLPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex80mLPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex40mLPFStart_ValueChanged()`** — L16275 — `private void udAlex40mLPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex40mLPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex40mLPFEnd_ValueChanged()`** — L16283 — `private void udAlex40mLPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex40mLPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex20mLPFStart_ValueChanged()`** — L16291 — `private void udAlex20mLPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex20mLPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex20mLPFEnd_ValueChanged()`** — L16299 — `private void udAlex20mLPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex20mLPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex15mLPFStart_ValueChanged()`** — L16307 — `private void udAlex15mLPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex15mLPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex15mLPFEnd_ValueChanged()`** — L16315 — `private void udAlex15mLPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex15mLPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex10mLPFStart_ValueChanged()`** — L16323 — `private void udAlex10mLPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex10mLPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex10mLPFEnd_ValueChanged()`** — L16331 — `private void udAlex10mLPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex10mLPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex6mLPFStart_ValueChanged()`** — L16339 — `private void udAlex6mLPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex6mLPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex1_5HPFStart_ValueChanged()`** — L16347 — `private void udAlex1_5HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex1_5HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex1_5HPFEnd_ValueChanged()`** — L16352 — `private void udAlex1_5HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex1_5HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex6_5HPFStart_ValueChanged()`** — L16360 — `private void udAlex6_5HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex6_5HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex6_5HPFEnd_ValueChanged()`** — L16368 — `private void udAlex6_5HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex6_5HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex9_5HPFStart_ValueChanged()`** — L16376 — `private void udAlex9_5HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex9_5HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex9_5HPFEnd_ValueChanged()`** — L16384 — `private void udAlex9_5HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex9_5HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex13HPFStart_ValueChanged()`** — L16392 — `private void udAlex13HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex13HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex13HPFEnd_ValueChanged()`** — L16400 — `private void udAlex13HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex13HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex20HPFStart_ValueChanged()`** — L16408 — `private void udAlex20HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex20HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex20HPFEnd_ValueChanged()`** — L16416 — `private void udAlex20HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex20HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex6BPFStart_ValueChanged()`** — L16424 — `private void udAlex6BPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex6BPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex6BPFEnd_ValueChanged()`** — L16432 — `private void udAlex6BPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex6BPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex21_5HPFStart_ValueChanged()`** — L16437 — `private void udAlex21_5HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex21_5HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex21_5HPFEnd_ValueChanged()`** — L16442 — `private void udAlex21_5HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex21_5HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex26_5HPFStart_ValueChanged()`** — L16450 — `private void udAlex26_5HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex26_5HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex26_5HPFEnd_ValueChanged()`** — L16458 — `private void udAlex26_5HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex26_5HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex29_5HPFStart_ValueChanged()`** — L16466 — `private void udAlex29_5HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex29_5HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex29_5HPFEnd_ValueChanged()`** — L16474 — `private void udAlex29_5HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex29_5HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex213HPFStart_ValueChanged()`** — L16482 — `private void udAlex213HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex213HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex213HPFEnd_ValueChanged()`** — L16490 — `private void udAlex213HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex213HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex220HPFStart_ValueChanged()`** — L16498 — `private void udAlex220HPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex220HPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex220HPFEnd_ValueChanged()`** — L16506 — `private void udAlex220HPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex220HPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex26BPFStart_ValueChanged()`** — L16514 — `private void udAlex26BPFStart_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex26BPFStart` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAlex26BPFEnd_ValueChanged()`** — L16522 — `private void udAlex26BPFEnd_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAlex26BPFEnd` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSWRTuneProtection_CheckedChanged()`** — L16527 — `private void chkSWRTuneProtection_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSWRTuneProtection` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbDisplayFFTSize_Scroll()`** — L16534 — `private void tbDisplayFFTSize_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDisplayFFTSize` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX2DisplayFFTSize_Scroll()`** — L16560 — `private void tbRX2DisplayFFTSize_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2DisplayFFTSize` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDispWinType_SelectedIndexChanged()`** — L16584 — `private void comboDispWinType_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDispWinType` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboRX2DispWinType_SelectedIndexChanged()`** — L16599 — `private void comboRX2DispWinType_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRX2DispWinType` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPNBTransition_ValueChanged()`** — L16611 — `private void udDSPNBTransition_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNBTransition` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPNBLead_ValueChanged()`** — L16618 — `private void udDSPNBLead_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNBLead` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPNBLag_ValueChanged()`** — L16625 — `private void udDSPNBLag_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNBLag` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkCBlock_CheckedChanged()`** — L16632 — `private void chkCBlock_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCBlock` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.chkCBlock_before_rx1_CheckedChanged()` (same file), `.chkCBlock_after_rx1_CheckedChanged()` (same file)
- **`.chkRX2CBlock_CheckedChanged()`** — L16649 — `private void chkRX2CBlock_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2CBlock` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.chkCBlock_before_rx2_CheckedChanged()` (same file), `.chkCBlock_after_rx2_CheckedChanged()` (same file)
- **`.chkCBlock_before_rx1_CheckedChanged()`** — L16666 — `private void chkCBlock_before_rx1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCBlock_before_rx1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCBlock_after_rx1_CheckedChanged()`** — L16672 — `private void chkCBlock_after_rx1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCBlock_after_rx1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCBlock_before_rx2_CheckedChanged()`** — L16678 — `private void chkCBlock_before_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCBlock_before_rx2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCBlock_after_rx2_CheckedChanged()`** — L16684 — `private void chkCBlock_after_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCBlock_after_rx2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnConfigure_Click()`** — L16690 — `private void btnConfigure_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnConfigure` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ConfigMidi2CatSetupClosed()`** — L16702 — `private void ConfigMidi2CatSetupClosed(object sender, FormClosedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX1WaterfallAGC_CheckedChanged()`** — L16711 — `private void chkRX1WaterfallAGC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX1WaterfallAGC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkRX2WaterfallAGC_CheckedChanged()`** — L16721 — `private void chkRX2WaterfallAGC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2WaterfallAGC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkPAValues_CheckedChanged()`** — L16731 — `private void chkPAValues_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPAValues` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnResetPAValues_Click()`** — L16737 — `private void btnResetPAValues_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetPAValues` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnResetWattMeterValues_Click()`** — L16750 — `private void btnResetWattMeterValues_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetWattMeterValues` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboDSPRxWindow_SelectedIndexChanged()`** — L16807 — `private void comboDSPRxWindow_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDSPRxWindow` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDSPTxWindow_SelectedIndexChanged()`** — L16827 — `private void comboDSPTxWindow_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDSPTxWindow` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radSaturn3p5mm_CheckedChanged()`** — L16842 — `private void radSaturn3p5mm_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSaturn3p5mm` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radOrionPTTOff_CheckedChanged()`** — L16849 — `private void radOrionPTTOff_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radOrionPTTOff` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radOrionMicTip_CheckedChanged()`** — L16855 — `private void radOrionMicTip_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radOrionMicTip` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radOrionBiasOn_CheckedChanged()`** — L16863 — `private void radOrionBiasOn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radOrionBiasOn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkEXT1OutOnTx_CheckedChanged()`** — L16871 — `private void chkEXT1OutOnTx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEXT1OutOnTx` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkEXT2OutOnTx_CheckedChanged()`** — L16884 — `private void chkEXT2OutOnTx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEXT2OutOnTx` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXGenScale_ValueChanged()`** — L16897 — `private void udTXGenScale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenScale` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXGenFreq_ValueChanged()`** — L16923 — `private void udTXGenFreq_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenFreq` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXGenSweepLow_ValueChanged()`** — L16943 — `private void udTXGenSweepLow_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenSweepLow` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenSweepHigh_ValueChanged()`** — L16948 — `private void udTXGenSweepHigh_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenSweepHigh` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenSweepRate_ValueChanged()`** — L16953 — `private void udTXGenSweepRate_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenSweepRate` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSigGenRX1_CheckedChanged()`** — L16958 — `private void chkSigGenRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSigGenRX1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRXGenScale_ValueChanged()`** — L16963 — `private void udRXGenScale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXGenScale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRXGenFreq_ValueChanged()`** — L16982 — `private void udRXGenFreq_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXGenFreq` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRXGenSweepLow_ValueChanged()`** — L16993 — `private void udRXGenSweepLow_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXGenSweepLow` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRXGenSweepHigh_ValueChanged()`** — L16999 — `private void udRXGenSweepHigh_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXGenSweepHigh` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRXGenSweepRate_ValueChanged()`** — L17005 — `private void udRXGenSweepRate_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXGenSweepRate` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXInhibit_CheckedChanged()`** — L17011 — `private void chkTXInhibit_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTXInhibit` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXInhibitReverse_CheckedChanged()`** — L17015 — `private void chkTXInhibitReverse_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTXInhibitReverse` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenPulseFreq_ValueChanged()`** — L17019 — `private void udTXGenPulseFreq_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenPulseFreq` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenPulseDutyCycle_ValueChanged()`** — L17024 — `private void udTXGenPulseDutyCycle_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenPulseDutyCycle` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXGenPulseTransition_ValueChanged()`** — L17029 — `private void udTXGenPulseTransition_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXGenPulseTransition` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEmphPos_CheckedChanged()`** — L17034 — `private void chkEmphPos_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEmphPos` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkRemoveTone_CheckedChanged()`** — L17040 — `private void chkRemoveTone_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRemoveTone` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkFMDetLimON_CheckedChanged()`** — L17048 — `private void chkFMDetLimON_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMDetLimON` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbDSPDetLimGain_Scroll()`** — L17056 — `private void tbDSPDetLimGain_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDSPDetLimGain` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPEERon_CheckedChanged()`** — L17064 — `private void chkDSPEERon_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPEERon` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPEERmgain_ValueChanged()`** — L17077 — `private void udDSPEERmgain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPEERmgain` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPEERpgain_ValueChanged()`** — L17083 — `private void udDSPEERpgain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPEERpgain` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPEERmdelay_ValueChanged()`** — L17089 — `private void udDSPEERmdelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPEERmdelay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPEERamIQ_CheckedChanged()`** — L17095 — `private void chkDSPEERamIQ_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPEERamIQ` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkHFTRRelay_CheckedChanged()`** — L17101 — `private void chkHFTRRelay_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHFTRRelay` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udHWKeyDownDelay_ValueChanged()`** — L17107 — `private void udHWKeyDownDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udHWKeyDownDelay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDDCADC_CheckedChanged()`** — L17274 — `private void radDDCADC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDDCADC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radP1DDCADC_CheckedChanged()`** — L17319 — `private void radP1DDCADC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radP1DDCADC` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDSPNOBmode_SelectedIndexChanged()`** — L17358 — `private void comboDSPNOBmode_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDSPNOBmode` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkClickTuneDrag_CheckedChanged()`** — L17372 — `private void chkClickTuneDrag_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkClickTuneDrag` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDSPRX1APFEnable_CheckedChanged()`** — L17382 — `private void chkDSPRX1APFEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPRX1APFEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPRX1subAPFEnable_CheckedChanged()`** — L17401 — `private void chkDSPRX1subAPFEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPRX1subAPFEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPRX2APFEnable_CheckedChanged()`** — L17421 — `private void chkDSPRX2APFEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPRX2APFEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbDSPAudRX1APFGain_ValueChanged()`** — L17436 — `private void tbDSPAudRX1APFGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDSPAudRX1APFGain` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbDSPAudRX1subAPFGain_ValueChanged()`** — L17444 — `private void tbDSPAudRX1subAPFGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDSPAudRX1subAPFGain` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbDSPAudRX2APFGain_ValueChanged()`** — L17452 — `private void tbDSPAudRX2APFGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDSPAudRX2APFGain` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX1APFTune_Scroll()`** — L17460 — `private void tbRX1APFTune_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1APFTune` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX1subAPFTune_Scroll()`** — L17468 — `private void tbRX1subAPFTune_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1subAPFTune` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX2APFTune_Scroll()`** — L17476 — `private void tbRX2APFTune_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2APFTune` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX1APFBW_Scroll()`** — L17484 — `private void tbRX1APFBW_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1APFBW` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX1subAPFBW_Scroll()`** — L17492 — `private void tbRX1subAPFBW_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1subAPFBW` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX2APFBW_Scroll()`** — L17500 — `private void tbRX2APFBW_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2APFBW` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPRX1APFControls_CheckedChanged()`** — L17508 — `private void radDSPRX1APFControls_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPRX1APFControls` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPRX1subAPFControls_CheckedChanged()`** — L17531 — `private void radDSPRX1subAPFControls_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPRX1subAPFControls` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPRX2APFControls_CheckedChanged()`** — L17554 — `private void radDSPRX2APFControls_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPRX2APFControls` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPRX1DollyEnable_CheckedChanged()`** — L17577 — `private void chkDSPRX1DollyEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPRX1DollyEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPRX1DollySubEnable_CheckedChanged()`** — L17583 — `private void chkDSPRX1DollySubEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPRX1DollySubEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPRX2DollyEnable_CheckedChanged()`** — L17589 — `private void chkDSPRX2DollyEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPRX2DollyEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPRX1DollyF0_ValueChanged()`** — L17595 — `private void udDSPRX1DollyF0_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPRX1DollyF0` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPRX1SubDollyF0_ValueChanged()`** — L17601 — `private void udDSPRX1SubDollyF0_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPRX1SubDollyF0` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPRX2DollyF0_ValueChanged()`** — L17607 — `private void udDSPRX2DollyF0_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPRX2DollyF0` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPRX1DollyF1_ValueChanged()`** — L17613 — `private void udDSPRX1DollyF1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPRX1DollyF1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPRX2DollyF1_ValueChanged()`** — L17619 — `private void udDSPRX2DollyF1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPRX2DollyF1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPRX1SubDollyF1_ValueChanged()`** — L17625 — `private void udDSPRX1SubDollyF1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPRX1SubDollyF1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udATTOnTX_ValueChanged()`** — L17630 — `private void udATTOnTX_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udATTOnTX` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud6mLNAGainOffset_ValueChanged()`** — L17635 — `private void ud6mLNAGainOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud6mLNAGainOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPEERpdelay_ValueChanged()`** — L17640 — `private void udDSPEERpdelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPEERpdelay` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPEERRunDelays_CheckedChanged()`** — L17646 — `private void chkDSPEERRunDelays_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPEERRunDelays` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPEERpwmMax_ValueChanged()`** — L17652 — `private void udDSPEERpwmMax_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPEERpwmMax` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPEERpwmMin_ValueChanged()`** — L17658 — `private void udDSPEERpwmMin_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPEERpwmMin` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboKeyerConnSecondary_MouseClick()`** — L17664 — `private void comboKeyerConnSecondary_MouseClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboCATPort_Click()`** — L17673 — `private void comboCATPort_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCATPort` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT2Port_Click()`** — L17681 — `private void comboCAT2Port_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCAT2Port` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT3Port_Click()`** — L17689 — `private void comboCAT3Port_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCAT3Port` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCAT4Port_Click()`** — L17697 — `private void comboCAT4Port_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCAT4Port` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCATPTTPort_Click()`** — L17705 — `private void comboCATPTTPort_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCATPTTPort` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboAndromedaCATPort_Click()`** — L17713 — `private void ComboAndromedaCATPort_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboAndromedaCATPort` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDSPCESSB_CheckedChanged()`** — L17722 — `private void chkDSPCESSB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPCESSB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRXAMSQMaxTail_ValueChanged()`** — L17737 — `private void udRXAMSQMaxTail_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXAMSQMaxTail` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2Linear_CheckedChanged()`** — L17746 — `private void radDSPNR2Linear_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2Linear` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2Log_CheckedChanged()`** — L17756 — `private void radDSPNR2Log_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2Log` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2OSMS_CheckedChanged()`** — L17766 — `private void radDSPNR2OSMS_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2OSMS` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2MMSE_CheckedChanged()`** — L17776 — `private void radDSPNR2MMSE_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2MMSE` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPNR2AE_CheckedChanged()`** — L17786 — `private void chkDSPNR2AE_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPNR2AE` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2LinearRX2_CheckedChanged()`** — L17798 — `private void radDSPNR2LinearRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2LinearRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2LogRX2_CheckedChanged()`** — L17808 — `private void radDSPNR2LogRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2LogRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2OSMSRX2_CheckedChanged()`** — L17818 — `private void radDSPNR2OSMSRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2OSMSRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2MMSERX2_CheckedChanged()`** — L17828 — `private void radDSPNR2MMSERX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2MMSERX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDSPNR2AERX2_CheckedChanged()`** — L17838 — `private void chkDSPNR2AERX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPNR2AERX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2Gamma_CheckedChanged()`** — L17850 — `private void radDSPNR2Gamma_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2Gamma` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radDSPNR2GammaRX2_CheckedChanged()`** — L17856 — `private void radDSPNR2GammaRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2GammaRX2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLimitExtAmpOnOverload_CheckedChanged()`** — L17862 — `private void chkLimitExtAmpOnOverload_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLimitExtAmpOnOverload` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboFocusMasterMode_SelectedIndexChanged()`** — L17872 — `private void comboFocusMasterMode_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFocusMasterMode` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtFocusMasterUDPPort_TextChanged()`** — L17913 — `private void txtFocusMasterUDPPort_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtFocusMasterUDPPort` text changes.
  Called by: `.comboFocusMasterMode_SelectedIndexChanged()` (same file)
- **`.txtFocusMasterDelay_TextChanged()`** — L17918 — `private void txtFocusMasterDelay_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtFocusMasterDelay` text changes.
  Called by: `.comboFocusMasterMode_SelectedIndexChanged()` (same file)
- **`.txtFocusMasterWinTitle_KeyDown()`** — L17923 — `private void txtFocusMasterWinTitle_KeyDown(object sender, KeyEventArgs e)`
  WinForms event handler: runs when `txtFocusMasterWinTitle` receives a key-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnableLEDFont_CheckedChanged()`** — L17955 — `private void chkEnableLEDFont_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEnableLEDFont` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisableRXOut_CheckedChanged()`** — L17960 — `private void chkDisableRXOut_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisableRXOut` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSplitPins_CheckedChanged()`** — L17967 — `private void chkSplitPins_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSplitPins` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radSplitPins4x3_CheckedChanged()`** — L17978 — `private void radSplitPins4x3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSplitPins4x3` checked state changes.
  Called by: `.chkSplitPins_CheckedChanged()` (same file)
- **`.radSplitPins3x4_CheckedChanged()`** — L17988 — `private void radSplitPins3x4_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSplitPins3x4` checked state changes.
  Called by: `.chkSplitPins_CheckedChanged()` (same file)
- **`.udDSPSNBThresh1_ValueChanged()`** — L17998 — `private void udDSPSNBThresh1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPSNBThresh1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPSNBThresh2_ValueChanged()`** — L18006 — `private void udDSPSNBThresh2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPSNBThresh2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnMNFAdd_Click()`** — L18021 — `private void btnMNFAdd_Click(object sender, EventArgs e)`
  accept input for a notch to be added
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMNFEdit_Click()`** — L18059 — `private void btnMNFEdit_Click(object sender, EventArgs e)`
  accept input for editing a notch
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMNFEnter_Click()`** — L18089 — `private void btnMNFEnter_Click(object sender, EventArgs e)`
  store the values from an Add or Edit operation
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMNFCancel_Click()`** — L18133 — `unsafe private void btnMNFCancel_Click(object sender, EventArgs e)`
  cancel the Add or Edit operation
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMNFDelete_Click()`** — L18176 — `private void btnMNFDelete_Click(object sender, EventArgs e)`
  delete a notch
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udMNFNotch_ValueChanged()`** — L18245 — `private void udMNFNotch_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udMNFNotch` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMNFAutoIncrease_CheckedChanged()`** — L18276 — `private void chkMNFAutoIncrease_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMNFAutoIncrease` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.SaveNotchesToDatabase()`** — L18285 — `unsafe public void SaveNotchesToDatabase()`
  Saves notches to database.
  Called by: `.btnMNFEnter_Click()` (same file), `.btnMNFDelete_Click()` (same file)
- **`.UpdateNotchDisplay()`** — L18307 — `unsafe public void UpdateNotchDisplay()`
  Updates notch display.
  Called by: `.RestoreNotchesFromDatabase()` (same file)
- **`.RestoreNotchesFromDatabase()`** — L18340 — `unsafe public void RestoreNotchesFromDatabase()`
  Restores notches from database.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOFreq_Click()`** — L18363 — `private void btnVFOFreq_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnVFOFreq` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAudioExpert_CheckedChanged()`** — L18370 — `private void chkAudioExpert_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAudioExpert` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNetworkWDT_CheckedChanged()`** — L18375 — `private void chkNetworkWDT_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNetworkWDT` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboRX2DispPanDetector_SelectedIndexChanged()`** — L18381 — `private void comboRX2DispPanDetector_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRX2DispPanDetector` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboRX2DispPanAveraging_SelectedIndexChanged()`** — L18396 — `private void comboRX2DispPanAveraging_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRX2DispPanAveraging` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboRX2DispWFDetector_SelectedIndexChanged()`** — L18408 — `private void comboRX2DispWFDetector_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRX2DispWFDetector` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboRX2DispWFAveraging_SelectedIndexChanged()`** — L18414 — `private void comboRX2DispWFAveraging_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRX2DispWFAveraging` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2DisplayWFAVTime_ValueChanged()`** — L18420 — `private void udRX2DisplayWFAVTime_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2DisplayWFAVTime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDispPanDetector_SelectedIndexChanged()`** — L18426 — `private void comboDispPanDetector_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDispPanDetector` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDispWFDetector_SelectedIndexChanged()`** — L18441 — `private void comboDispWFDetector_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDispWFDetector` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDispPanAveraging_SelectedIndexChanged()`** — L18447 — `private void comboDispPanAveraging_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDispPanAveraging` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboDispWFAveraging_SelectedIndexChanged()`** — L18461 — `private void comboDispWFAveraging_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDispWFAveraging` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayAVTimeWF_ValueChanged()`** — L18467 — `private void udDisplayAVTimeWF_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayAVTimeWF` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDispRX2Normalize_CheckedChanged()`** — L18473 — `private void chkDispRX2Normalize_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDispRX2Normalize` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDispNormalize_CheckedChanged()`** — L18485 — `private void chkDispNormalize_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDispNormalize` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboTXDispPanDetector_SelectedIndexChanged()`** — L18497 — `private void comboTXDispPanDetector_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboTXDispPanDetector` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboTXDispPanAveraging_SelectedIndexChanged()`** — L18507 — `private void comboTXDispPanAveraging_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboTXDispPanAveraging` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXDisplayAVGTime_ValueChanged()`** — L18514 — `private void udTXDisplayAVGTime_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXDisplayAVGTime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDispTXNormalize_CheckedChanged()`** — L18521 — `private void chkDispTXNormalize_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDispTXNormalize` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbTXDisplayFFTSize_Scroll()`** — L18528 — `private void tbTXDisplayFFTSize_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTXDisplayFFTSize` is scrolled.
  Called by: `.UpdateTXDisplayFFT()` (same file)
- **`.comboTXDispWinType_SelectedIndexChanged()`** — L18537 — `private void comboTXDispWinType_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboTXDispWinType` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboTXDispWFDetector_SelectedIndexChanged()`** — L18544 — `private void comboTXDispWFDetector_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboTXDispWFDetector` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.comboTXDispWFAveraging_SelectedIndexChanged()`** — L18551 — `private void comboTXDispWFAveraging_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboTXDispWFAveraging` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTXDisplayAVTime_ValueChanged()`** — L18558 — `private void udTXDisplayAVTime_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXDisplayAVTime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkBPF2Gnd_CheckedChanged()`** — L18565 — `private void chkBPF2Gnd_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBPF2Gnd` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.ud1_5BPF1Start_ValueChanged()`** — L18571 — `private void ud1_5BPF1Start_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud1_5BPF1Start` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud1_5BPF1End_ValueChanged()`** — L18576 — `private void ud1_5BPF1End_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud1_5BPF1End` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud6_5BPF1Start_ValueChanged()`** — L18585 — `private void ud6_5BPF1Start_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud6_5BPF1Start` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud6_5BPF1End_ValueChanged()`** — L18594 — `private void ud6_5BPF1End_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud6_5BPF1End` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud9_5BPF1Start_ValueChanged()`** — L18603 — `private void ud9_5BPF1Start_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud9_5BPF1Start` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud9_5BPF1End_ValueChanged()`** — L18612 — `private void ud9_5BPF1End_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud9_5BPF1End` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud13BPF1Start_ValueChanged()`** — L18621 — `private void ud13BPF1Start_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud13BPF1Start` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud13BPF1End_ValueChanged()`** — L18630 — `private void ud13BPF1End_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud13BPF1End` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud20BPF1Start_ValueChanged()`** — L18639 — `private void ud20BPF1Start_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud20BPF1Start` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud20BPF1End_ValueChanged()`** — L18648 — `private void ud20BPF1End_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud20BPF1End` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud6BPF1Start_ValueChanged()`** — L18657 — `private void ud6BPF1Start_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud6BPF1Start` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ud6BPF1End_ValueChanged()`** — L18666 — `private void ud6BPF1End_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud6BPF1End` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBPF1_1_5BP_CheckedChanged()`** — L18671 — `private void chkBPF1_1_5BP_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBPF1_1_5BP` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBPF1_6_5BP_CheckedChanged()`** — L18676 — `private void chkBPF1_6_5BP_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBPF1_6_5BP` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBPF1_9_5BP_CheckedChanged()`** — L18681 — `private void chkBPF1_9_5BP_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBPF1_9_5BP` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBPF1_13BP_CheckedChanged()`** — L18686 — `private void chkBPF1_13BP_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBPF1_13BP` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBPF1_20BP_CheckedChanged()`** — L18691 — `private void chkBPF1_20BP_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBPF1_20BP` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBPF1_6BP_CheckedChanged()`** — L18696 — `private void chkBPF1_6BP_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBPF1_6BP` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkANAN8000DLEDisplayVoltsAmps_CheckedChanged()`** — L18701 — `private void chkANAN8000DLEDisplayVoltsAmps_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkANAN8000DLEDisplayVoltsAmps` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.ud6mRx2LNAGainOffset_ValueChanged()`** — L18707 — `private void ud6mRx2LNAGainOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ud6mRx2LNAGainOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnableXVTRHF_CheckedChanged()`** — L18712 — `private void chkEnableXVTRHF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEnableXVTRHF` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkCFCEnable_CheckedChanged()`** — L18719 — `private void chkCFCEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCFCEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.setCFCProfile()`** — L18739 — `private void setCFCProfile(object sender, EventArgs e)`
  Sets cfcprofile.
  Called by: `.setLegacyCFCProfile()` (same file)
- **`.tbCFCPRECOMP_Scroll()`** — L18792 — `private void tbCFCPRECOMP_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbCFCPRECOMP` is scrolled.
  Called by: `.setLegacyCFCProfile()` (same file)
- **`.chkCFCPeqEnable_CheckedChanged()`** — L18802 — `private void chkCFCPeqEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCFCPeqEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkPHROTEnable_CheckedChanged()`** — L18816 — `private void chkPHROTEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPHROTEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udPhRotFreq_ValueChanged()`** — L18827 — `private void udPhRotFreq_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPhRotFreq` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udPHROTStages_ValueChanged()`** — L18833 — `private void udPHROTStages_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPHROTStages` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbCFCPEG_Scroll()`** — L18839 — `private void tbCFCPEG_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbCFCPEG` is scrolled.
  Called by: `.setLegacyCFCProfile()` (same file)
- **`.radTXDSB_CheckedChanged()`** — L18847 — `private void radTXDSB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radTXDSB` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnExportCurrentTXProfile_Click()`** — L18860 — `private void btnExportCurrentTXProfile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnExportCurrentTXProfile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.timer_VAC_Monitor_Tick()`** — L18882 — `[HandleProcessCorruptedStateExceptions] private void timer_VAC_Monitor_Tick(object sender, EventArgs e)`
  [2.10.3.9]MW0LGE this attribue, together with the app.config change 'legacyCorruptedStateExceptionsPolicy' enables the catch of address acceptions inside the try/catch block which is in an unsafe block
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC1_Force_CheckedChanged()`** — L19107 — `private void chkVAC1_Force_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1_Force` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC1_Force2_CheckedChanged()`** — L19115 — `private void chkVAC1_Force2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1_Force2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2_Force_CheckedChanged()`** — L19123 — `private void chkVAC2_Force_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2_Force` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2_Force2_CheckedChanged()`** — L19130 — `private void chkVAC2_Force2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2_Force2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBlockTxAnt2_CheckedChanged()`** — L19137 — `private void chkBlockTxAnt2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBlockTxAnt2` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkBlockTxAnt3_CheckedChanged()`** — L19160 — `private void chkBlockTxAnt3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBlockTxAnt3` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkLPFBypass_CheckedChanged()`** — L19183 — `private void chkLPFBypass_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLPFBypass` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.lblVAC1ovfl_Click()`** — L19189 — `private void lblVAC1ovfl_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC1ovfl` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblVAC1unfl_Click()`** — L19194 — `private void lblVAC1unfl_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC1unfl` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblVAC1ovfl2_Click()`** — L19199 — `private void lblVAC1ovfl2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC1ovfl2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblVAC1unfl2_Click()`** — L19204 — `private void lblVAC1unfl2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC1unfl2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblVAC2ovfl_Click()`** — L19209 — `private void lblVAC2ovfl_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC2ovfl` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblVAC2unfl_Click()`** — L19214 — `private void lblVAC2unfl_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC2unfl` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblVAC2ovfl2_Click()`** — L19219 — `private void lblVAC2ovfl2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC2ovfl2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblVAC2unfl2_Click()`** — L19224 — `private void lblVAC2unfl2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblVAC2unfl2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.timer_LED_Mirror_Tick()`** — L19229 — `private void timer_LED_Mirror_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `timer_LED_Mirror` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2onSplit_CheckedChanged()`** — L19258 — `private void chkVAC2onSplit_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2onSplit` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVOXEnable_CheckedChanged()`** — L19266 — `private void chkVOXEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVOXEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDEXPEnable_CheckedChanged()`** — L19274 — `private void chkDEXPEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDEXPEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDEXPAttack_ValueChanged()`** — L19282 — `private void udDEXPAttack_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPAttack` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDEXPHold_ValueChanged()`** — L19288 — `private void udDEXPHold_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPHold` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDEXPRelease_ValueChanged()`** — L19294 — `private void udDEXPRelease_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPRelease` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDEXPThreshold_ValueChanged()`** — L19300 — `private void udDEXPThreshold_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPThreshold` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDEXPExpansionRatio_ValueChanged()`** — L19307 — `private void udDEXPExpansionRatio_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPExpansionRatio` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDEXPHysteresisRatio_ValueChanged()`** — L19313 — `private void udDEXPHysteresisRatio_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPHysteresisRatio` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDEXPDetTau_ValueChanged()`** — L19319 — `private void udDEXPDetTau_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPDetTau` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSCFEnable_CheckedChanged()`** — L19325 — `private void chkSCFEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSCFEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udSCFLowCut_ValueChanged()`** — L19331 — `private void udSCFLowCut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udSCFLowCut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udSCFHighCut_ValueChanged()`** — L19337 — `private void udSCFHighCut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udSCFHighCut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkDEXPLookAheadEnable_CheckedChanged()`** — L19343 — `private void chkDEXPLookAheadEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDEXPLookAheadEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.chkVOXEnable_CheckedChanged()` (same file), `.chkDEXPEnable_CheckedChanged()` (same file)
- **`.udDEXPLookAhead_ValueChanged()`** — L19350 — `private void udDEXPLookAhead_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDEXPLookAhead` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udUpdatesPerStepMax_ValueChanged()`** — L19356 — `private void udUpdatesPerStepMax_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udUpdatesPerStepMax` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udUpdatesPerStepMin_ValueChanged()`** — L19364 — `private void udUpdatesPerStepMin_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udUpdatesPerStepMin` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAntiVoxEnable_CheckedChanged()`** — L19372 — `private void chkAntiVoxEnable_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAntiVoxEnable` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udAntiVoxGain_ValueChanged()`** — L19378 — `private void udAntiVoxGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAntiVoxGain` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udAntiVoxTau_ValueChanged()`** — L19384 — `private void udAntiVoxTau_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAntiVoxTau` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAntiVoxSource_CheckedChanged()`** — L19390 — `private void chkAntiVoxSource_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAntiVoxSource` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkITSync_CheckedChanged()`** — L19396 — `private void chkITSync_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkITSync` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisablePicDisplayBackgroundImage_CheckedChanged()`** — L19404 — `private void chkDisablePicDisplayBackgroundImage_CheckedChanged(object sender, EventArgs e)`
  MW0LGE
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnDataFill_Changed()`** — L19410 — `private void clrbtnDataFill_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.TbDataFillAlpha_Scroll()` (same file)
- **`.TbDataFillAlpha_Scroll()`** — L19416 — `private void TbDataFillAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `TbDataFillAlpha` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.chkWheelTunesOutsideSpectral_CheckedChanged()`** — L19424 — `private void chkWheelTunesOutsideSpectral_CheckedChanged(object sender, EventArgs e)`
  --- RAWINPUT
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboHIDMouseWheel_SelectedIndexChanged()`** — L19429 — `private void ComboHIDMouseWheel_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboHIDMouseWheel` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkGlobalListenForMouseWheel_CheckedChanged()`** — L19447 — `private void ChkGlobalListenForMouseWheel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkGlobalListenForMouseWheel` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkWheelOnlyAdjustsVFO_CheckedChanged()`** — L19452 — `private void ChkWheelOnlyAdjustsVFO_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkWheelOnlyAdjustsVFO` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupHIDControls()`** — L19457 — `private void setupHIDControls(bool bEnabled)`
  Called by: `.ChkAlsoUseSpecificMouseWheel_CheckedChanged()` (same file)
- **`.ChkAlsoUseSpecificMouseWheel_CheckedChanged()`** — L19467 — `private void ChkAlsoUseSpecificMouseWheel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkAlsoUseSpecificMouseWheel` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.UpdateRawInputMouseDevices()`** — L19474 — `public void UpdateRawInputMouseDevices(Dictionary<IntPtr, MouseEvent> mouseDevices)`
  Updates raw input mouse devices.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WheelChangeNotify()`** — L19489 — `public void WheelChangeNotify()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Timer_RawInputMouseWheel_Tick()`** — L19544 — `private void Timer_RawInputMouseWheel_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `Timer_RawInputMouseWheel` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TxtDeviceHID_hidden_TextChanged()`** — L19550 — `private void TxtDeviceHID_hidden_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `TxtDeviceHID_hidden` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowRXFilterOnWaterfall_CheckedChanged()`** — L19563 — `private void chkShowRXFilterOnWaterfall_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowRXFilterOnWaterfall` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.ChkShowTXFilterOnWaterfall_CheckedChanged()`** — L19569 — `private void ChkShowTXFilterOnWaterfall_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkShowTXFilterOnWaterfall` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowRXZeroLineOnWaterfall_CheckedChanged()`** — L19574 — `private void chkShowRXZeroLineOnWaterfall_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowRXZeroLineOnWaterfall` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.ChkShowTXZeroLineOnWaterfall_CheckedChanged()`** — L19580 — `private void ChkShowTXZeroLineOnWaterfall_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkShowTXZeroLineOnWaterfall` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowTXFilterOnRXWaterfall_CheckedChanged()`** — L19585 — `private void chkShowTXFilterOnRXWaterfall_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowTXFilterOnRXWaterfall` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.ChkZoomShiftModifier_CheckedChanged()`** — L19591 — `private void ChkZoomShiftModifier_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkZoomShiftModifier` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkReverseShiftZoomModifier_CheckedChanged()`** — L19597 — `private void ChkReverseShiftZoomModifier_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkReverseShiftZoomModifier` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkExtended_CheckedChanged()`** — L19602 — `private void ChkExtended_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkExtended` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWaterfallUseRX1SpectrumMinMax_CheckedChanged()`** — L19613 — `private void chkWaterfallUseRX1SpectrumMinMax_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWaterfallUseRX1SpectrumMinMax` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkWaterfallUseRX2SpectrumMinMax_CheckedChanged()`** — L19638 — `private void chkWaterfallUseRX2SpectrumMinMax_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWaterfallUseRX2SpectrumMinMax` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.chkRecenterOnZZFx_CheckedChanged()`** — L19663 — `private void chkRecenterOnZZFx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecenterOnZZFx` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDisplayPhasePtSize_ValueChanged()`** — L19669 — `private void udDisplayPhasePtSize_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayPhasePtSize` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowFPS_CheckedChanged()`** — L19674 — `private void chkShowFPS_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowFPS` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSmallModeFilteronVFOs_CheckedChanged()`** — L19679 — `private void chkSmallModeFilteronVFOs_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSmallModeFilteronVFOs` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboDisplayThreadPriority_SelectedIndexChanged()`** — L19684 — `private void comboDisplayThreadPriority_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDisplayThreadPriority` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShowSeqLog_Click()`** — L19702 — `private void btnShowSeqLog_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShowSeqLog` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radSReading_CheckedChanged()`** — L19707 — `private void radSReading_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSReading` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radDBM_CheckedChanged()`** — L19712 — `private void radDBM_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDBM` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radUV_CheckedChanged()`** — L19717 — `private void radUV_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUV` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAntiAlias_CheckedChanged()`** — L19722 — `private void chkAntiAlias_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAntiAlias` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnStatusBarBackground_Changed()`** — L19727 — `private void clrbtnStatusBarBackground_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnStatusBarText_Changed()`** — L19732 — `private void clrbtnStatusBarText_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkAccurateFrameTiming_CheckedChanged()`** — L19737 — `private void chkAccurateFrameTiming_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAccurateFrameTiming` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.GetActivePeakHoldsEnabledRX()`** — L19749 — `public bool GetActivePeakHoldsEnabledRX(int rx)`
  Returns active peak holds enabled rx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetActivePeakHoldsEnabledRX()`** — L19761 — `public void SetActivePeakHoldsEnabledRX(int rx, bool enabled)`
  Sets active peak holds enabled rx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkPeakBlobsEnabled_CheckedChanged()`** — L19787 — `private void chkPeakBlobsEnabled_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPeakBlobsEnabled` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udPeakBlobs_ValueChanged()`** — L19807 — `private void udPeakBlobs_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPeakBlobs` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkPeakBlobInsideFilterOnly_CheckedChanged()`** — L19813 — `private void chkPeakBlobInsideFilterOnly_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPeakBlobInsideFilterOnly` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSignalHistory_CheckedChanged()`** — L19819 — `private void chkSignalHistory_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSignalHistory` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnSignalHistoryColour_Changed()`** — L19827 — `private void clrbtnSignalHistoryColour_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbSignalHistoryAlpha_Scroll()` (same file)
- **`.tbSignalHistoryAlpha_Scroll()`** — L19833 — `private void tbSignalHistoryAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbSignalHistoryAlpha` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVSyncDX_CheckedChanged()`** — L19839 — `private void chkVSyncDX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVSyncDX` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkBlobPeakHold_CheckedChanged()`** — L19854 — `private void chkBlobPeakHold_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBlobPeakHold` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udBlobPeakHoldMS_ValueChanged()`** — L19863 — `private void udBlobPeakHoldMS_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udBlobPeakHoldMS` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udSignalHistoryDuration_ValueChanged()`** — L19869 — `private void udSignalHistoryDuration_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udSignalHistoryDuration` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udSpaceMoxDelay_ValueChanged()`** — L19875 — `private void udSpaceMoxDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udSpaceMoxDelay` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ButtonAndromeda_Click()`** — L19880 — `private void ButtonAndromeda_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ButtonAndromeda` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.LabelTS528_Click()`** — L19885 — `private void LabelTS528_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `LabelTS528` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboGanymedeCATPort_SelectedIndexChanged()`** — L19891 — `private void ComboGanymedeCATPort_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboGanymedeCATPort` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboAriesCATPort_SelectedIndexChanged()`** — L19909 — `private void ComboAriesCATPort_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboAriesCATPort` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkEnableGanymede_CheckedChanged()`** — L19927 — `private void ChkEnableGanymede_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkEnableGanymede` checked state changes.
  Called by: `.AfterConstructor()` (same file)
- **`.ChkEnableAries_CheckedChanged()`** — L19977 — `private void ChkEnableAries_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkEnableAries` checked state changes.
  Called by: `.AfterConstructor()` (same file)
- **`.ComboGanymedeCATPort_Click()`** — L20027 — `private void ComboGanymedeCATPort_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboGanymedeCATPort` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ComboAriesCATPort_Click()`** — L20035 — `private void ComboAriesCATPort_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ComboAriesCATPort` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BtnGanymedeReset_Click()`** — L20043 — `private void BtnGanymedeReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnGanymedeReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BtnAriesErase1_Click()`** — L20048 — `private void BtnAriesErase1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnAriesErase1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BtnAriesErase2_Click()`** — L20053 — `private void BtnAriesErase2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnAriesErase2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BtnAriesErase3_Click()`** — L20058 — `private void BtnAriesErase3_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnAriesErase3` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkAries1_CheckedChanged()`** — L20063 — `private void ChkAries1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkAries1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkAries2_CheckedChanged()`** — L20068 — `private void ChkAries2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkAries2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkAries3_CheckedChanged()`** — L20073 — `private void ChkAries3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkAries3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetAriesEnabledButton()`** — L20079 — `public void SetAriesEnabledButton(int Antenna, bool State)`
  set an aries enabled button for 1 of the 3 antennas
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkEnableAriesQuickTune_CheckedChanged()`** — L20089 — `private void checkEnableAriesQuickTune_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `checkEnableAriesQuickTune` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkQSOTimerEnabled_CheckedChanged()`** — L20094 — `private void chkQSOTimerEnabled_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQSOTimerEnabled` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkQSOTimerOnlyDuringMOX_CheckedChanged()`** — L20116 — `private void chkQSOTimerOnlyDuringMOX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQSOTimerOnlyDuringMOX` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkQSOTimerPlaySoundOnExpiry_CheckedChanged()`** — L20122 — `private void chkQSOTimerPlaySoundOnExpiry_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQSOTimerPlaySoundOnExpiry` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.chkQSOTimerEnabled_CheckedChanged()` (same file)
- **`.btnQSOTimerSelectWAV_Click()`** — L20133 — `private void btnQSOTimerSelectWAV_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnQSOTimerSelectWAV` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnQSOTimerPlaySelectedWAV_Click()`** — L20147 — `private void btnQSOTimerPlaySelectedWAV_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnQSOTimerPlaySelectedWAV` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.audioFileLoaded()`** — L20152 — `private void audioFileLoaded(bool bOk)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.udQSOTimerMinutes_ValueChanged()`** — L20157 — `private void udQSOTimerMinutes_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udQSOTimerMinutes` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udQSOTimerSeconds_ValueChanged()`** — L20163 — `private void udQSOTimerSeconds_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udQSOTimerSeconds` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkQSOTimerResetOnMOX_CheckedChanged()`** — L20169 — `private void chkQSOTimerResetOnMOX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQSOTimerResetOnMOX` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.setQSOTimerDuration()`** — L20175 — `private void setQSOTimerDuration()`
  Sets qsotimer duration.
  Called by: `.ForceAllEvents()` (same file), `.udQSOTimerMinutes_ValueChanged()` (same file), `.udQSOTimerSeconds_ValueChanged()` (same file)
- **`.chkQSOTimerResetOnExpiry_CheckedChanged()`** — L20189 — `private void chkQSOTimerResetOnExpiry_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQSOTimerResetOnExpiry` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.chkQSOTimerEnabled_CheckedChanged()` (same file)
- **`.chkQSOTimerFlashTimerIfResetOnExpiry_CheckedChanged()`** — L20199 — `private void chkQSOTimerFlashTimerIfResetOnExpiry_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQSOTimerFlashTimerIfResetOnExpiry` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.chkQSOTimerResetOnExpiry_CheckedChanged()` (same file)
- **`.comboRadioModel_SelectedIndexChanged()`** — L20204 — `private void comboRadioModel_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRadioModel` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.removeHL2Options()`** — L21085 — `private void removeHL2Options()`
  Called by: `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.setupADCRadioButtions()`** — L21091 — `private void setupADCRadioButtions()`
  Called by: `.AfterConstructor()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.setADCVisible()`** — L21116 — `private void setADCVisible(Control cc, bool bADC0, bool bADC1, bool bADC2)`
  Sets adcvisible.
  Called by: `.setupADCRadioButtions()` (same file)
- **`.ShowSetupTab()`** — L21165 — `public void ShowSetupTab(SetupTab eTab)`
  Shows setup tab.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnResetP2ADC_Click()`** — L21274 — `private void btnResetP2ADC_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetP2ADC` is clicked.
  Called by: `.AfterConstructor()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.btnResetP1ADC_Click()`** — L21305 — `private void btnResetP1ADC_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetP1ADC` is clicked.
  Called by: `.AfterConstructor()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file)
- **`.chkAndrG2Panel_CheckedChanged()`** — L21336 — `private void chkAndrG2Panel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAndrG2Panel` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.checkAriesStandalone_CheckedChanged()`** — L21351 — `private void checkAriesStandalone_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `checkAriesStandalone` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTxBufferLat_ValueChanged()`** — L21358 — `private void udTxBufferLat_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTxBufferLat` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udPTTHang_ValueChanged()`** — L21365 — `private void udPTTHang_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPTTHang` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCATtoVFOB_CheckedChanged()`** — L21372 — `private void chkCATtoVFOB_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCATtoVFOB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisconnectReset_CheckedChanged()`** — L21378 — `private void chkDisconnectReset_CheckedChanged(object sender, EventArgs e)`
  MI0BOT: Controls if the HL2 will reset after an Ethernet disconnect
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAndrBandBtnDefault_CheckedChanged()`** — L21384 — `private void chkAndrBandBtnDefault_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAndrBandBtnDefault` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAndrStickyShift_CheckedChanged()`** — L21389 — `private void chkAndrStickyShift_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAndrStickyShift` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAndrStickyMenus_CheckedChanged()`** — L21394 — `private void chkAndrStickyMenus_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAndrStickyMenus` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkActivePeakHoldRX1_CheckedChanged()`** — L21399 — `private void chkActivePeakHoldRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkActivePeakHoldRX1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udActivePeakHoldDurationRX1_ValueChanged()`** — L21415 — `private void udActivePeakHoldDurationRX1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udActivePeakHoldDurationRX1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkActivePeakHoldRX2_CheckedChanged()`** — L21421 — `private void chkActivePeakHoldRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkActivePeakHoldRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udActivePeakHoldDurationRX2_ValueChanged()`** — L21437 — `private void udActivePeakHoldDurationRX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udActivePeakHoldDurationRX2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnActiveSpectralPeak_Changed()`** — L21443 — `private void clrbtnActiveSpectralPeak_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbActiveSpectralPeakAlpha_Scroll()` (same file)
- **`.tbActiveSpectralPeakAlpha_Scroll()`** — L21449 — `private void tbActiveSpectralPeakAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbActiveSpectralPeakAlpha` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lgPickerRX1_Changed()`** — L21467 — `private void lgPickerRX1_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnDeleteColourGripper_Click()`** — L21472 — `private void btnDeleteColourGripper_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDeleteColourGripper` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnClearColourGrippers_Click()`** — L21478 — `private void btnClearColourGrippers_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClearColourGrippers` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lgPickerRX1_GripperSelected()`** — L21483 — `private void lgPickerRX1_GripperSelected(object sender, ColourEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnGripperColour_Changed()`** — L21489 — `private void clrbtnGripperColour_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkPanadpatorGradient_CheckedChanged()`** — L21496 — `private void chkPanadpatorGradient_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPanadpatorGradient` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSpecWarningLEDRenderDelay_CheckedChanged()`** — L21522 — `private void chkSpecWarningLEDRenderDelay_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSpecWarningLEDRenderDelay` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSpecWarningLEDGetPixels_CheckedChanged()`** — L21528 — `private void chkSpecWarningLEDGetPixels_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSpecWarningLEDGetPixels` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnDefaultGradient_Click()`** — L21534 — `private void btnDefaultGradient_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDefaultGradient` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbDataLineAlpha_Scroll()`** — L21540 — `private void tbDataLineAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDataLineAlpha` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPeakHoldDrop_CheckedChanged()`** — L21546 — `private void chkPeakHoldDrop_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPeakHoldDrop` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVFOSyncLinksCTUN_CheckedChanged()`** — L21555 — `private void chkVFOSyncLinksCTUN_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVFOSyncLinksCTUN` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnI2CRead_MouseDown()`** — L21561 — `private void btnI2CRead_MouseDown(object sender, MouseEventArgs e)`
  MI0BOT: HL2 access to I2C bus
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnI2CWrite_MouseDown()`** — L21625 — `private void btnI2CWrite_MouseDown(object sender, MouseEventArgs e)`
  MI0BOT: HL2 access to I2C bus
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkI2CWriteEnable_CheckedChanged()`** — L21653 — `private void chkI2CWriteEnable_CheckedChanged(object sender, EventArgs e)`
  MI0BOT: HL2 access to I2C bus
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkI2CEnable_CheckedChanged()`** — L21670 — `private void chkI2CEnable_CheckedChanged(object sender, EventArgs e)`
  MI0BOT: HL2 access to I2C bus
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.WriteVersaClockAsync()`** — L21752 — `private async Task WriteVersaClockAsync( byte[] registerData )`
  MI0BOT: Support for HL2 Cl2 clock output
  Called by: `.EnableCl1_10MHz()` (same file), `.DisableCl1_10MHz()` (same file), `.ControlCl2()` (same file)
- **`.EnableCl1_10MHz()`** — L21803 — `async public Task EnableCl1_10MHz()`
  MI0BOT: Support for HL2 10MHz clock input
  Called by: `.chkExt10MHz_CheckedChanged()` (same file)
- **`.DisableCl1_10MHz()`** — L21809 — `async public Task DisableCl1_10MHz()`
  MI0BOT: Support for HL2 10MHz clock input
  Called by: `.chkExt10MHz_CheckedChanged()` (same file)
- **`.ControlCl2()`** — L21815 — `async public Task ControlCl2(bool enable)`
  MI0BOT: Support for HL2 Cl2 clock output
  Called by: `.chkCl2Enable_CheckedChanged()` (same file), `.udCl2Freq_ValueChanged()` (same file), `.chkExt10MHz_CheckedChanged()` (same file)
- **`.chkCl2Enable_CheckedChanged()`** — L21852 — `private void chkCl2Enable_CheckedChanged(object sender, EventArgs e)`
  MI0BOT: Support for HL2 Cl2 clock output
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCl2Freq_ValueChanged()`** — L21858 — `private void udCl2Freq_ValueChanged(object sender, EventArgs e)`
  MI0BOT: Support for HL2 Cl2 clock output
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkExt10MHz_CheckedChanged()`** — L21864 — `async private void chkExt10MHz_CheckedChanged(object sender, EventArgs e)`
  MI0BOT: Support for HL2 10MHz clock input
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDataLineGradient_CheckedChanged()`** — L21878 — `private void chkDataLineGradient_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDataLineGradient` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lgPickerRX1_GripperDBMChanged()`** — L21885 — `private void lgPickerRX1_GripperDBMChanged(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgPickerRX1_GripperMouseLeave()`** — L21892 — `private void lgPickerRX1_GripperMouseLeave(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgPickerRX1_GripperMouseEnter()`** — L21897 — `private void lgPickerRX1_GripperMouseEnter(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboFRSRegion_MouseEnter()`** — L21902 — `private void comboFRSRegion_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFRSRegion` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.showRegionBandstackWarning()`** — L21906 — `private void showRegionBandstackWarning(bool bShow)`
  Called by: `.AfterConstructor()` (same file), `.comboFRSRegion_MouseEnter()` (same file), `.comboFRSRegion_MouseLeave()` (same file), `.chkExtended_MouseEnter()` (same file), `.chkExtended_MouseLeave()` (same file), `.comboFRSRegion_MouseClick()` (same file) — and 1 more
- **`.comboFRSRegion_MouseLeave()`** — L21920 — `private void comboFRSRegion_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFRSRegion` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkExtended_MouseEnter()`** — L21925 — `private void chkExtended_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `chkExtended` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkExtended_MouseLeave()`** — L21930 — `private void chkExtended_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `chkExtended` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboFRSRegion_MouseClick()`** — L21935 — `private void comboFRSRegion_MouseClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboFRSRegion_MouseHover()`** — L21940 — `private void comboFRSRegion_MouseHover(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkN1MMEnableRX1_CheckedChanged()`** — L21945 — `private void chkN1MMEnableRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkN1MMEnableRX1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.stopStartN1MMSpectrum()`** — L21964 — `private void stopStartN1MMSpectrum()`
  Called by: `.chkN1MMEnableRX1_CheckedChanged()` (same file), `.chkN1MMEnableRX2_CheckedChanged()` (same file)
- **`.chkN1MMEnableRX2_CheckedChanged()`** — L21975 — `private void chkN1MMEnableRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkN1MMEnableRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtN1MMSendTo_TextChanged()`** — L21994 — `private void txtN1MMSendTo_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtN1MMSendTo` text changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udN1MMSendRate_ValueChanged()`** — L22000 — `private void udN1MMSendRate_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udN1MMSendRate` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udN1MMRX1Scaling_ValueChanged()`** — L22006 — `private void udN1MMRX1Scaling_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udN1MMRX1Scaling` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udN1MMRX2Scaling_ValueChanged()`** — L22012 — `private void udN1MMRX2Scaling_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udN1MMRX2Scaling` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX1WaterfallOpacity_Scroll()`** — L22018 — `private void tbRX1WaterfallOpacity_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1WaterfallOpacity` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.tbRX2WaterfallOpacity_Scroll()`** — L22025 — `private void tbRX2WaterfallOpacity_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2WaterfallOpacity` is scrolled.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowControlDebug_CheckedChanged()`** — L22032 — `private void chkShowControlDebug_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowControlDebug` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.UpdateDDCTab()`** — L22041 — `public void UpdateDDCTab()`
  Updates ddctab.
  Called by: `.AfterConstructor()` (same file), `.ucRadioList_Radios_SelectedRadioChanged()` (same file)
- **`.chkHighlightTXProfileSaveItems_CheckedChanged()`** — L22149 — `private void chkHighlightTXProfileSaveItems_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHighlightTXProfileSaveItems` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnN1MMDefault_Click()`** — L22154 — `private void btnN1MMDefault_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnN1MMDefault` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnVAC1AdvancedDefault_Click()`** — L22160 — `private void btnVAC1AdvancedDefault_Click(object sender, EventArgs e)`
  MW0LGE_21h resampler advanced settings
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udVAC1FeedbackGainIn_ValueChanged()`** — L22175 — `private void udVAC1FeedbackGainIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FeedbackGainIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1SlewTimeIn_ValueChanged()`** — L22181 — `private void udVAC1SlewTimeIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1SlewTimeIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1FeedbackGainOut_ValueChanged()`** — L22186 — `private void udVAC1FeedbackGainOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FeedbackGainOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1SlewTimeOut_ValueChanged()`** — L22192 — `private void udVAC1SlewTimeOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1SlewTimeOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnVAC2AdvancedDefault_Click()`** — L22198 — `private void btnVAC2AdvancedDefault_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnVAC2AdvancedDefault` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udVAC2FeedbackGainIn_ValueChanged()`** — L22213 — `private void udVAC2FeedbackGainIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FeedbackGainIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2SlewTimeIn_ValueChanged()`** — L22219 — `private void udVAC2SlewTimeIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2SlewTimeIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2FeedbackGainOut_ValueChanged()`** — L22224 — `private void udVAC2FeedbackGainOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FeedbackGainOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2SlewTimeOut_ValueChanged()`** — L22230 — `private void udVAC2SlewTimeOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2SlewTimeOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.pbVAC1FeedbackGainInfo_Click()`** — L22235 — `private void pbVAC1FeedbackGainInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbVAC1FeedbackGainInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbVAC1SlewTimeInfo_Click()`** — L22240 — `private void pbVAC1SlewTimeInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbVAC1SlewTimeInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNetworkThrottleIndexTweak_CheckedChanged()`** — L22246 — `private void chkNetworkThrottleIndexTweak_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNetworkThrottleIndexTweak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateNetworkThrottleCheckBox()`** — L22279 — `private void updateNetworkThrottleCheckBox()`
  Called by: `.AfterConstructor()` (same file), `.chkNetworkThrottleIndexTweak_CheckedChanged()` (same file)
- **`.tbBandstackOverlayAlpha_Scroll()`** — L22291 — `private void tbBandstackOverlayAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbBandstackOverlayAlpha` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnBandstackOverlay_Changed()`** — L22297 — `private void clrbtnBandstackOverlay_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbBandstackOverlayAlpha_Scroll()` (same file)
- **`.udVAC1GrapherSwing_ValueChanged()`** — L22303 — `private void udVAC1GrapherSwing_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1GrapherSwing` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC1GrapherAuto_CheckedChanged()`** — L22308 — `private void chkVAC1GrapherAuto_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1GrapherAuto` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2GrapherAuto_CheckedChanged()`** — L22314 — `private void chkVAC2GrapherAuto_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2GrapherAuto` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udVAC2GrapherSwing_ValueChanged()`** — L22320 — `private void udVAC2GrapherSwing_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2GrapherSwing` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.isPowerOfTwo()`** — L22325 — `private bool isPowerOfTwo(int x)`
  Called by: `.udVAC1PropMaxIn_ValueChanged()` (same file), `.udVAC1FFMaxIn_ValueChanged()` (same file), `.udVAC1PropMaxOut_ValueChanged()` (same file), `.udVAC1FFMaxOut_ValueChanged()` (same file), `.udVAC2PropMaxIn_ValueChanged()` (same file), `.udVAC2FFMaxIn_ValueChanged()` (same file) — and 2 more
- **`.udVAC1PropMinIn_ValueChanged()`** — L22330 — `private void udVAC1PropMinIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1PropMinIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1PropMaxIn_ValueChanged()`** — L22336 — `private void udVAC1PropMaxIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1PropMaxIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1FFMinIn_ValueChanged()`** — L22362 — `private void udVAC1FFMinIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FFMinIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1FFMaxIn_ValueChanged()`** — L22368 — `private void udVAC1FFMaxIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FFMaxIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1FFAlphaIn_ValueChanged()`** — L22394 — `private void udVAC1FFAlphaIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FFAlphaIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC1OldVarIn_CheckedChanged()`** — L22400 — `private void chkVAC1OldVarIn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1OldVarIn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1PropMinOut_ValueChanged()`** — L22406 — `private void udVAC1PropMinOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1PropMinOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1PropMaxOut_ValueChanged()`** — L22412 — `private void udVAC1PropMaxOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1PropMaxOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1FFMinOut_ValueChanged()`** — L22438 — `private void udVAC1FFMinOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FFMinOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1FFMaxOut_ValueChanged()`** — L22444 — `private void udVAC1FFMaxOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FFMaxOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC1FFAlphaOut_ValueChanged()`** — L22470 — `private void udVAC1FFAlphaOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC1FFAlphaOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC1OldVarOut_CheckedChanged()`** — L22476 — `private void chkVAC1OldVarOut_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1OldVarOut` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtVAC1OldVarIn_TextChanged()`** — L22482 — `private void txtVAC1OldVarIn_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtVAC1OldVarIn` text changes.
  Called by: `.ForceAllEvents()` (same file), `.chkVAC1OldVarIn_CheckedChanged()` (same file)
- **`.txtVAC1OldVarOut_TextChanged()`** — L22489 — `private void txtVAC1OldVarOut_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtVAC1OldVarOut` text changes.
  Called by: `.ForceAllEvents()` (same file), `.chkVAC1OldVarOut_CheckedChanged()` (same file)
- **`.udVAC2PropMinIn_ValueChanged()`** — L22496 — `private void udVAC2PropMinIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2PropMinIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2PropMaxIn_ValueChanged()`** — L22502 — `private void udVAC2PropMaxIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2PropMaxIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2FFMinIn_ValueChanged()`** — L22528 — `private void udVAC2FFMinIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FFMinIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2FFMaxIn_ValueChanged()`** — L22534 — `private void udVAC2FFMaxIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FFMaxIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2FFAlphaIn_ValueChanged()`** — L22560 — `private void udVAC2FFAlphaIn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FFAlphaIn` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2OldVarIn_CheckedChanged()`** — L22566 — `private void chkVAC2OldVarIn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2OldVarIn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2PropMinOut_ValueChanged()`** — L22572 — `private void udVAC2PropMinOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2PropMinOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2PropMaxOut_ValueChanged()`** — L22578 — `private void udVAC2PropMaxOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2PropMaxOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2FFMinOut_ValueChanged()`** — L22604 — `private void udVAC2FFMinOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FFMinOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2FFMaxOut_ValueChanged()`** — L22610 — `private void udVAC2FFMaxOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FFMaxOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udVAC2FFAlphaOut_ValueChanged()`** — L22636 — `private void udVAC2FFAlphaOut_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVAC2FFAlphaOut` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2OldVarOut_CheckedChanged()`** — L22642 — `private void chkVAC2OldVarOut_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2OldVarOut` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtVAC2OldVarIn_TextChanged()`** — L22648 — `private void txtVAC2OldVarIn_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtVAC2OldVarIn` text changes.
  Called by: `.ForceAllEvents()` (same file), `.chkVAC2OldVarIn_CheckedChanged()` (same file)
- **`.txtVAC2OldVarOut_TextChanged()`** — L22655 — `private void txtVAC2OldVarOut_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtVAC2OldVarOut` text changes.
  Called by: `.ForceAllEvents()` (same file), `.chkVAC2OldVarOut_CheckedChanged()` (same file)
- **`.pbVAC1PropFeedbackMinInfo_Click()`** — L22662 — `private void pbVAC1PropFeedbackMinInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbVAC1PropFeedbackMinInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbVAC1PropFeedbackMaxInfo_Click()`** — L22667 — `private void pbVAC1PropFeedbackMaxInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbVAC1PropFeedbackMaxInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbVAC1FFMinInfo_Click()`** — L22672 — `private void pbVAC1FFMinInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbVAC1FFMinInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbVAC1FFMaxInfo_Click()`** — L22677 — `private void pbVAC1FFMaxInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbVAC1FFMaxInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbVAC1FFAlphaInfo_Click()`** — L22682 — `private void pbVAC1FFAlphaInfo_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbVAC1FFAlphaInfo` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboPin1TXActionHF_SelectedIndexChanged()`** — L22687 — `private void comboPin1TXActionHF_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboPin1TXActionHF` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboPin1TXActionVHF_SelectedIndexChanged()`** — L22698 — `private void comboPin1TXActionVHF_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboPin1TXActionVHF` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboPin1TXActionSWL_SelectedIndexChanged()`** — L22709 — `private void comboPin1TXActionSWL_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboPin1TXActionSWL` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateOCLedStrip()`** — L22720 — `public void UpdateOCLedStrip(bool tx, int bits)`
  Updates ocled strip.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateIOLedStrip()`** — L22726 — `public void UpdateIOLedStrip(bool tx, byte bits)`
  Updates ioled strip.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableIOLedStrip()`** — L22732 — `public void EnableIOLedStrip(bool state)`
  Enables ioled strip.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tcOCControl_SelectedIndexChanged()`** — L22739 — `private void tcOCControl_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `tcOCControl` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPin1TXPAHF_CheckedChanged()`** — L22758 — `private void chkPin1TXPAHF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPin1TXPAHF` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPin1TXPAVHF_CheckedChanged()`** — L22773 — `private void chkPin1TXPAVHF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPin1TXPAVHF` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPin1TXPASWL_CheckedChanged()`** — L22788 — `private void chkPin1TXPASWL_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPin1TXPASWL` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CheckForAnyExternalPACheckBoxes()`** — L22803 — `public bool CheckForAnyExternalPACheckBoxes()`
  Checks for any external pacheck boxes.
  Called by: `.AfterConstructor()` (same file), `.chkPin1TXPAHF_CheckedChanged()` (same file), `.chkPin1TXPAVHF_CheckedChanged()` (same file), `.chkPin1TXPASWL_CheckedChanged()` (same file), `.chkPin1RXPAHF_CheckedChanged()` (same file), `.chkPin1RXPAVHF_CheckedChanged()` (same file) — and 1 more
- **`.chkTuneStepPerModeRX1_CheckedChanged()`** — L22847 — `private void chkTuneStepPerModeRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTuneStepPerModeRX1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC1WillMute_CheckedChanged()`** — L22853 — `private void chkVAC1WillMute_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1WillMute` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2WillMute_CheckedChanged()`** — L22859 — `private void chkVAC2WillMute_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2WillMute` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAllowHotSwitching_CheckedChanged()`** — L22865 — `private void chkAllowHotSwitching_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAllowHotSwitching` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.UpdateForHotSwitch()`** — L22871 — `public void UpdateForHotSwitch(bool tx)`
  Updates for hot switch.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkShowRX1NoiseFloor_CheckedChanged()`** — L22906 — `private void chkShowRX1NoiseFloor_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowRX1NoiseFloor` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowRX2NoiseFloor_CheckedChanged()`** — L22914 — `private void chkShowRX2NoiseFloor_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowRX2NoiseFloor` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAutoAGCRX1_CheckedChanged()`** — L22922 — `private void chkAutoAGCRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoAGCRX1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAutoAGCRX2_CheckedChanged()`** — L22928 — `private void chkAutoAGCRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoAGCRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX1AutoAGCOffset_ValueChanged()`** — L22950 — `private void udRX1AutoAGCOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX1AutoAGCOffset` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udRX2AutoAGCOffset_ValueChanged()`** — L22956 — `private void udRX2AutoAGCOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2AutoAGCOffset` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkPin1RXPAHF_CheckedChanged()`** — L22962 — `private void chkPin1RXPAHF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPin1RXPAHF` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPin1RXPAVHF_CheckedChanged()`** — L22977 — `private void chkPin1RXPAVHF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPin1RXPAVHF` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPin1RXPASWL_CheckedChanged()`** — L22992 — `private void chkPin1RXPASWL_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPin1RXPASWL` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udNoiseFloorAttackRX1_ValueChanged()`** — L23007 — `private void udNoiseFloorAttackRX1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udNoiseFloorAttackRX1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udNoiseFloorAttackRX2_ValueChanged()`** — L23013 — `private void udNoiseFloorAttackRX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udNoiseFloorAttackRX2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnNoiseFloor_Changed()`** — L23019 — `private void clrbtnNoiseFloor_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.chkNoiseFloorShowDBM_CheckedChanged()`** — L23025 — `private void chkNoiseFloorShowDBM_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNoiseFloorShowDBM` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udNoiseFloorLineWidth_ValueChanged()`** — L23031 — `private void udNoiseFloorLineWidth_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udNoiseFloorLineWidth` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udWaterfallAGCOffsetRX1_ValueChanged()`** — L23037 — `private void udWaterfallAGCOffsetRX1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udWaterfallAGCOffsetRX1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkWaterfallUseNFForAGCRX1_CheckedChanged()`** — L23043 — `private void chkWaterfallUseNFForAGCRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWaterfallUseNFForAGCRX1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udWaterfallAGCOffsetRX2_ValueChanged()`** — L23049 — `private void udWaterfallAGCOffsetRX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udWaterfallAGCOffsetRX2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkWaterfallUseNFForAGCRX2_CheckedChanged()`** — L23055 — `private void chkWaterfallUseNFForAGCRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWaterfallUseNFForAGCRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tmrCFCOMPGain_Tick()`** — L23065 — `private void tmrCFCOMPGain_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `tmrCFCOMPGain` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picCFC_Paint()`** — L23096 — `private void picCFC_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `picCFC` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCFCDisplayAutoScale_CheckedChanged()`** — L23222 — `private void chkCFCDisplayAutoScale_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCFCDisplayAutoScale` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udCFCPicDBPerLine_ValueChanged()`** — L23228 — `private void udCFCPicDBPerLine_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udCFCPicDBPerLine` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowDisplayDebug_CheckedChanged()`** — L23235 — `private void chkShowDisplayDebug_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowDisplayDebug` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCTLimitDragToSpectral_CheckedChanged()`** — L23242 — `private void chkCTLimitDragToSpectral_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCTLimitDragToSpectral` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCTLimitDragMouseOnly_CheckedChanged()`** — L23248 — `private void chkCTLimitDragMouseOnly_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCTLimitDragMouseOnly` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udActivePeakHoldDropRX1_ValueChanged()`** — L23253 — `private void udActivePeakHoldDropRX1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udActivePeakHoldDropRX1` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udActivePeakHoldDropRX2_ValueChanged()`** — L23258 — `private void udActivePeakHoldDropRX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udActivePeakHoldDropRX2` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udPeakBlobDropDBMs_ValueChanged()`** — L23263 — `private void udPeakBlobDropDBMs_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPeakBlobDropDBMs` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFillActivePeakHoldRX1_CheckedChanged()`** — L23268 — `private void chkFillActivePeakHoldRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFillActivePeakHoldRX1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFillActivePeakHoldRX2_CheckedChanged()`** — L23273 — `private void chkFillActivePeakHoldRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFillActivePeakHoldRX2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAFSlidersMute_CheckedChanged()`** — L23278 — `private void chkAFSlidersMute_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAFSlidersMute` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkIgnoreMasterAFChangeForMonitor_CheckedChanged()`** — L23283 — `private void chkIgnoreMasterAFChangeForMonitor_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkIgnoreMasterAFChangeForMonitor` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowMHzOnCursor_CheckedChanged()`** — L23296 — `private void chkShowMHzOnCursor_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowMHzOnCursor` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLimitFilterEdgesToSidebands_CheckedChanged()`** — L23305 — `private void chkLimitFilterEdgesToSidebands_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLimitFilterEdgesToSidebands` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDisplayDecimation_ValueChanged()`** — L23310 — `private void udDisplayDecimation_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDisplayDecimation` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowPhaseAngularMean_CheckedChanged()`** — L23316 — `private void chkShowPhaseAngularMean_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowPhaseAngularMean` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDefaultBindPortForTCPIPCat_Click()`** — L23321 — `private void btnDefaultBindPortForTCPIPCat_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDefaultBindPortForTCPIPCat` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateTCPIPPort()`** — L23326 — `private void updateTCPIPPort()`
  Called by: `.btnDefaultBindPortForTCPIPCat_Click()` (same file), `.txtTCPIPCATServerBindIPPort_TextChanged()` (same file)
- **`.txtTCPIPCATServerBindIPPort_TextChanged()`** — L23362 — `private void txtTCPIPCATServerBindIPPort_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtTCPIPCATServerBindIPPort` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTCPIPCatServerListening_CheckedChanged()`** — L23369 — `private void chkTCPIPCatServerListening_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTCPIPCatServerListening` checked state changes.
  Called by: `.StartupTCPIPcatServer()` (same file)
- **`.chkZTBIsRecallStore_CheckedChanged()`** — L23378 — `private void chkZTBIsRecallStore_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkZTBIsRecallStore` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkZTBstoreLock_CheckedChanged()`** — L23384 — `private void chkZTBstoreLock_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkZTBstoreLock` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.DisableTCPIPCatServerDueToError()`** — L23389 — `public void DisableTCPIPCatServerDueToError()`
  Disables tcpipcat server due to error.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisableTCIServerDueToError()`** — L23394 — `public void DisableTCIServerDueToError()`
  Disables tciserver due to error.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkCTUNignore0beat_CheckedChanged()`** — L23400 — `private void chkCTUNignore0beat_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCTUNignore0beat` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.SetupDSPWarnings()`** — L23406 — `public void SetupDSPWarnings(bool bufferSizeDifferentRX, bool filterSizeDifferentRX, bool filterTypeDifferentRX, bool bufferSizeDifferentTX, bool filterSizeDifferentTX, bool filter`
  Setups dspwarnings.
  Called by: `.AfterConstructor()` (same file)
- **`.txtTCIServerBindIPPort_TextChanged()`** — L23425 — `private void txtTCIServerBindIPPort_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtTCIServerBindIPPort` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDefaultBindPortForTCI_Click()`** — L23432 — `private void btnDefaultBindPortForTCI_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDefaultBindPortForTCI` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTCIServerListening_CheckedChanged()`** — L23438 — `private void chkTCIServerListening_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTCIServerListening` checked state changes.
  Called by: `.StartupTCIServer()` (same file)
- **`.StartupTCIServer()`** — L23446 — `public void StartupTCIServer()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StartupTCPIPcatServer()`** — L23450 — `public void StartupTCPIPcatServer()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateTCIPort()`** — L23454 — `private void updateTCIPort()`
  Called by: `.txtTCIServerBindIPPort_TextChanged()` (same file), `.btnDefaultBindPortForTCI_Click()` (same file)
- **`.udTCIRateLimit_ValueChanged()`** — L23490 — `private void udTCIRateLimit_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTCIRateLimit` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCopyRX2VFObToVFOa_CheckedChanged()`** — L23495 — `private void chkCopyRX2VFObToVFOa_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCopyRX2VFObToVFOa` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkForgetRX2VfoBVFOinfo_CheckedChanged()`** — L23500 — `private void chkForgetRX2VfoBVFOinfo_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkForgetRX2VfoBVFOinfo` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkUseRX1vfoaForRX2vfoa_CheckedChanged()`** — L23504 — `private void chkUseRX1vfoaForRX2vfoa_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUseRX1vfoaForRX2vfoa` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTCIsendInitialStateOnConnect_CheckedChanged()`** — L23509 — `private void chkTCIsendInitialStateOnConnect_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTCIsendInitialStateOnConnect` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWelcomeMessageTCPIPCat_CheckedChanged()`** — L23514 — `private void chkWelcomeMessageTCPIPCat_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWelcomeMessageTCPIPCat` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udMaxTCISpots_ValueChanged()`** — L23519 — `private void udMaxTCISpots_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udMaxTCISpots` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udTCISpotLifetime_ValueChanged()`** — L23525 — `private void udTCISpotLifetime_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTCISpotLifetime` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowTCISpots_CheckedChanged()`** — L23531 — `private void chkShowTCISpots_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowTCISpots` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkFlashNewTCISpots_CheckedChanged()`** — L23541 — `private void chkFlashNewTCISpots_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFlashNewTCISpots` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSpotOwnCallAppearance_CheckedChanged()`** — L23553 — `private void chkSpotOwnCallAppearance_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSpotOwnCallAppearance` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtOwnCallsign_TextChanged()`** — L23562 — `private void txtOwnCallsign_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtOwnCallsign` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnOwnCallApearance_Changed()`** — L23567 — `private void clrbtnOwnCallApearance_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkCWLUbecomesCW_CheckedChanged()`** — L23572 — `private void chkCWLUbecomesCW_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCWLUbecomesCW` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShowLog_Click()`** — L23577 — `private void btnShowLog_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShowLog` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShowTCPIPCatLog_Click()`** — L23582 — `private void btnShowTCPIPCatLog_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShowTCPIPCatLog` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setDBtip()`** — L23587 — `private void setDBtip(object sender)`
  Sets dbtip.
  Called by: `.setCFCProfile()` (same file), `.tbCFCPRECOMP_Scroll()` (same file), `.tbCFCPEG_Scroll()` (same file)
- **`.chkNoFadeOverUnderWarning_CheckedChanged()`** — L23670 — `private void chkNoFadeOverUnderWarning_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNoFadeOverUnderWarning` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkUseOutlinedCross_CheckedChanged()`** — L23676 — `private void chkUseOutlinedCross_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUseOutlinedCross` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSwapREDBluePSAColours_CheckedChanged()`** — L23681 — `private void chkSwapREDBluePSAColours_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSwapREDBluePSAColours` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkHideFeebackLevel_CheckedChanged()`** — L23686 — `private void chkHideFeebackLevel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHideFeebackLevel` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SwapRedBlueChanged()`** — L23695 — `public void SwapRedBlueChanged()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkFirewallCheck_Click()`** — L23700 — `private void chkFirewallCheck_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFirewallCheck` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkUsing10MHzRef_CheckedChanged()`** — L23705 — `private void chkUsing10MHzRef_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUsing10MHzRef` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnHPSDRFreqCalReset10MHz_Click()`** — L23714 — `private void btnHPSDRFreqCalReset10MHz_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnHPSDRFreqCalReset10MHz` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udHPSDRFreqCorrectFactor10MHz_ValueChanged()`** — L23719 — `private void udHPSDRFreqCorrectFactor10MHz_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udHPSDRFreqCorrectFactor10MHz` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnLoadGradient_Click()`** — L23724 — `private async void btnLoadGradient_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLoadGradient` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSaveGradient_Click()`** — L23745 — `private async void btnSaveGradient_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSaveGradient` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkIgnoreSeqErrors_CheckedChanged()`** — L23760 — `private void chkIgnoreSeqErrors_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkIgnoreSeqErrors` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnIPv4TCI_Click()`** — L23765 — `private void btnIPv4TCI_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnIPv4TCI` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnIPv4N1MM_Click()`** — L23779 — `private void btnIPv4N1MM_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnIPv4N1MM` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnIPv4TCPCat_Click()`** — L23793 — `private void btnIPv4TCPCat_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnIPv4TCPCat` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEmulateSunSDR2Pro_CheckedChanged()`** — L23807 — `private void chkEmulateSunSDR2Pro_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEmulateSunSDR2Pro` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEmulateExpertSDR3Protocol_CheckedChanged()`** — L23812 — `private void chkEmulateExpertSDR3Protocol_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEmulateExpertSDR3Protocol` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisableRearSpeakerJacksAudioAmplifier_CheckedChanged()`** — L23824 — `private void chkDisableRearSpeakerJacksAudioAmplifier_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisableRearSpeakerJacksAudioAmplifier` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radUseDriveSliderTune_CheckedChanged()`** — L23829 — `private void radUseDriveSliderTune_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUseDriveSliderTune` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radUseTuneSliderTune_CheckedChanged()`** — L23834 — `private void radUseTuneSliderTune_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUseTuneSliderTune` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radUseFixedDriveTune_CheckedChanged()`** — L23839 — `private void radUseFixedDriveTune_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUseFixedDriveTune` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTestIMDPower_ValueChanged()`** — L23844 — `private void udTestIMDPower_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTestIMDPower` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.setupTuneAnd2ToneRadios()`** — L23849 — `private void setupTuneAnd2ToneRadios()`
  Called by: `.ForceAllEvents()` (same file)
- **`.radUseDriveSlider2Tone_CheckedChanged()`** — L23879 — `private void radUseDriveSlider2Tone_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUseDriveSlider2Tone` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radUseTuneSlider2Tone_CheckedChanged()`** — L23884 — `private void radUseTuneSlider2Tone_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUseTuneSlider2Tone` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radUseFixedDrive2Tone_CheckedChanged()`** — L23889 — `private void radUseFixedDrive2Tone_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radUseFixedDrive2Tone` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLimitPowerCATTCIMsgs_CheckedChanged()`** — L23893 — `private void chkLimitPowerCATTCIMsgs_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLimitPowerCATTCIMsgs` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updatePAControls()`** — L23899 — `private void updatePAControls(PAProfile p)`
  Called by: `.comboPAProfile_SelectedIndexChanged()` (same file)
- **`.comboPAProfile_SelectedIndexChanged()`** — L23922 — `private void comboPAProfile_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboPAProfile` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.validatePAProfileName()`** — L23959 — `private bool validatePAProfileName(string sProfileName)`
  Called by: `.btnCopyPAProfile_Click()` (same file), `.btnNewPAProfile_Click()` (same file)
- **`.btnCopyPAProfile_Click()`** — L23982 — `private void btnCopyPAProfile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCopyPAProfile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnNewPAProfile_Click()`** — L23999 — `private void btnNewPAProfile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnNewPAProfile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDeletePAProfile_Click()`** — L24013 — `private void btnDeletePAProfile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDeletePAProfile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateNUDgains()`** — L24042 — `private void updateNUDgains(PAProfile p)`
  Called by: `.comboPAProfile_SelectedIndexChanged()` (same file), `.btnResetPAProfile_Click()` (same file)
- **`.updateNUDAdjustgains()`** — L24085 — `private void updateNUDAdjustgains(PAProfile p = null)`
  Called by: `.comboPAProfile_SelectedIndexChanged()` (same file), `.btnResetPAProfile_Click()` (same file), `.setAdjustingBand()` (same file)
- **`.updateDriveLabels()`** — L24109 — `private void updateDriveLabels(PAProfile p = null)`
  Called by: `.comboPAProfile_SelectedIndexChanged()` (same file), `.btnResetPAProfile_Click()` (same file), `.setAdjustingBand()` (same file), `.nudMaxPowerForBandPA_ValueChanged()` (same file), `.chkUsePowerOnDrvTunPA_CheckedChanged()` (same file)
- **`.updateMaxPower()`** — L24160 — `private void updateMaxPower(PAProfile p = null)`
  Called by: `.comboPAProfile_SelectedIndexChanged()` (same file), `.btnResetPAProfile_Click()` (same file), `.setAdjustingBand()` (same file)
- **`.btnResetPAProfile_Click()`** — L24176 — `private void btnResetPAProfile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetPAProfile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mapBandMetersToIndex()`** — L24197 — `private int mapBandMetersToIndex(int nMeters)`
  Called by: `.nudPAProfileGain_ValueChanged()` (same file), `.enabledAllPAnuds()` (same file)
- **`.mapBandToMeters()`** — L24226 — `private int mapBandToMeters(Band b)`
  Called by: `.handleOldPAGainSettings()` (same file)
- **`.nudPAProfileGain_ValueChanged()`** — L24255 — `private void nudPAProfileGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPAProfileGain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.GetPAGain()`** — L24282 — `public float GetPAGain(Band b, int nDriveValue)`
  Returns pagain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPABandUsesMaxPower()`** — L24291 — `public bool GetPABandUsesMaxPower(Band b)`
  Returns paband uses max power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPABandMaxPower()`** — L24300 — `public float GetPABandMaxPower(Band b)`
  Returns paband max power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initPAProfiles()`** — L24311 — `private void initPAProfiles()`
  Called by: `.AfterConstructor()` (same file)
- **`.getPAProfile()`** — L24333 — `private PAProfile getPAProfile(string sName)`
  Returns paprofile.
  Called by: `.getOptions()` (same file), `.comboPAProfile_SelectedIndexChanged()` (same file), `.btnCopyPAProfile_Click()` (same file), `.btnDeletePAProfile_Click()` (same file), `.updateNUDgains()` (same file), `.updateNUDAdjustgains()` (same file) — and 14 more
- **`.updatePAProfileCombo()`** — L24339 — `private void updatePAProfileCombo(string sSelectProfile = "")`
  Called by: `.getOptions()` (same file), `.chkAutoPACalibrate_CheckedChanged()` (same file), `.comboRadioModel_SelectedIndexChanged()` (same file), `.btnCopyPAProfile_Click()` (same file), `.btnNewPAProfile_Click()` (same file), `.btnDeletePAProfile_Click()` (same file)
- **`.GetBypassGain()`** — L24392 — `public float GetBypassGain(Band b)`
  Returns bypass gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBypassGain()`** — L24399 — `public void SetBypassGain(Band b, float gain)`
  Sets bypass gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PAProfileEnableControls()`** — L24406 — `private void PAProfileEnableControls(bool tx)`
  Called by: `.OnMoxChangeHandler()` (same file)
- **`.getOldVariablePAgain()`** — L24450 — `private float getOldVariablePAgain(string sKey, ref Dictionary<string, string> getDict)`
  Returns old variable pagain.
  Called by: `.handleOldPAGainSettings()` (same file)
- **`.removeOldPASetting()`** — L24455 — `private void removeOldPASetting(string sSetting)`
  Called by: `.handleOldPAGainSettings()` (same file)
- **`.handleOldPAGainSettings()`** — L24459 — `private void handleOldPAGainSettings(ref Dictionary<string, string> getDict)`
  Called by: `.handleOutdatedOptions()` (same file)
- **`.OnMoxChangeHandler()`** — L24771 — `private void OnMoxChangeHandler(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXBandChanged()`** — L24780 — `private void OnTXBandChanged(Band oldBand, Band newBand, double tx_frequency)`
  Handles/raises the txband changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setAdjustingBand()`** — L24785 — `private void setAdjustingBand(Band b)`
  Sets adjusting band.
  Called by: `.OnTXBandChanged()` (same file)
- **`.enabledPAAdjust()`** — L25099 — `private void enabledPAAdjust(bool bEnabled)`
  Called by: `.setAdjustingBand()` (same file)
- **`.enabledAllPAnuds()`** — L25114 — `private void enabledAllPAnuds(bool bEnabled)`
  Called by: `.OnMoxChangeHandler()` (same file)
- **`.nudAdjustGain_ValueChanged()`** — L25144 — `private void nudAdjustGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudAdjustGain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateMaxPowerCheckbox()`** — L25172 — `private void updateMaxPowerCheckbox(PAProfile p = null)`
  Called by: `.updateMaxPower()` (same file), `.nudMaxPowerForBandPA_ValueChanged()` (same file)
- **`.nudMaxPowerForBandPA_ValueChanged()`** — L25193 — `private void nudMaxPowerForBandPA_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMaxPowerForBandPA` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkUsePowerOnDrvTunPA_CheckedChanged()`** — L25211 — `private void chkUsePowerOnDrvTunPA_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUsePowerOnDrvTunPA` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnSliderLimitBar_Changed()`** — L25227 — `private void clrbtnSliderLimitBar_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.chkUseSUnitsForPBNPPBSNR_CheckedChanged()`** — L25233 — `private void chkUseSUnitsForPBNPPBSNR_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUseSUnitsForPBNPPBSNR` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNFsensitivity_ValueChanged()`** — L25238 — `private void nudNFsensitivity_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNFsensitivity` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNFshift_ValueChanged()`** — L25243 — `private void nudNFshift_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNFshift` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNFShowDecimal_CheckedChanged()`** — L25248 — `private void chkNFShowDecimal_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNFShowDecimal` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAdjustGridMinToNFRX1_CheckedChanged()`** — L25253 — `private void chkAdjustGridMinToNFRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAdjustGridMinToNFRX1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAdjustGridMinToNFRX2_CheckedChanged()`** — L25260 — `private void chkAdjustGridMinToNFRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAdjustGridMinToNFRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudRX2NFoffsetGridFollow_ValueChanged()`** — L25267 — `private void nudRX2NFoffsetGridFollow_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRX2NFoffsetGridFollow` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRX1NFoffsetGridFollow_ValueChanged()`** — L25272 — `private void nudRX1NFoffsetGridFollow_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRX1NFoffsetGridFollow` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnResetLevelCal_Click()`** — L25277 — `private void btnResetLevelCal_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetLevelCal` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAmpDefault_Click()`** — L25291 — `private void btnAmpDefault_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAmpDefault` is clicked.
  Called by: `.initVoltsAmpsCalibration()` (same file)
- **`.udAmpVoff_ValueChanged()`** — L25298 — `private void udAmpVoff_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAmpVoff` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAmpSens_ValueChanged()`** — L25304 — `private void udAmpSens_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAmpSens` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.initVoltsAmpsCalibration()`** — L25310 — `private void initVoltsAmpsCalibration()`
  Called by: `.AfterConstructor()` (same file)
- **`.chkForceATTwhenPSAoff_CheckedChanged()`** — L25315 — `private void chkForceATTwhenPSAoff_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkForceATTwhenPSAoff` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVisualNotch_CheckedChanged()`** — L25321 — `private void chkVisualNotch_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVisualNotch` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRX1PBsnr_Click()`** — L25327 — `private void btnRX1PBsnr_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX1PBsnr` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRX2PBsnr_Click()`** — L25341 — `private void btnRX2PBsnr_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX2PBsnr` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnResetNFShift_Click()`** — L25355 — `private void btnResetNFShift_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetNFShift` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPreventTXonDifferentBandToRX_CheckedChanged()`** — L25359 — `private void chkPreventTXonDifferentBandToRX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPreventTXonDifferentBandToRX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.meterFromSelectedContainer()`** — L25402 — `private MeterManager.clsMeter meterFromSelectedContainer()`
  Called by: `.updateMeterLists()` (same file), `.btnAddMeterItem_Click()` (same file), `.btnRemoveMeterItem_Click()` (same file), `.btnMeterUp_Click()` (same file), `.btnMeterDown_Click()` (same file), `.meterItemGroupIDfromSelected()` (same file) — and 16 more
- **`.btnAddRX1Container_Click()`** — L25409 — `private void btnAddRX1Container_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAddRX1Container` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAddRX2Container_Click()`** — L25418 — `private void btnAddRX2Container_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAddRX2Container` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.containerNameFromId()`** — L25426 — `private string containerNameFromId(string id)`
  Called by: `.updateMeter2Controls()` (same file), `.ucOtherButtonsOptionsGrid_buttons_MacroSetupClicked()` (same file)
- **`.updateMeter2Controls()`** — L25447 — `private void updateMeter2Controls(string sId = "")`
  Called by: `.OnRX2EnabledChanged()` (same file), `.PerformDelayedInitalistion()` (same file), `.btnAddRX1Container_Click()` (same file), `.btnAddRX2Container_Click()` (same file), `.btnContainerDelete_Click()` (same file), `.btnContainer_load_Click()` (same file) — and 1 more
- **`.updateMeterLists()`** — L25527 — `private void updateMeterLists()`
  Called by: `.updateMeter2Controls()` (same file), `.comboContainerSelect_SelectedIndexChanged()` (same file), `.btnAddMeterItem_Click()` (same file), `.btnRemoveMeterItem_Click()` (same file), `.btnMeterUp_Click()` (same file), `.btnMeterDown_Click()` (same file)
- **`.findIndexForInsertOfSpecialItem()`** — L25573 — `private int findIndexForInsertOfSpecialItem(clsMeterTypeComboboxItem mtci, ListBox lb)`
  Called by: `.updateMeterLists()` (same file)
- **`.btnContainerDelete_Click()`** — L25618 — `private void btnContainerDelete_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnContainerDelete` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboContainerSelect_SelectedIndexChanged()`** — L25640 — `private void comboContainerSelect_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboContainerSelect` selection changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkContainerHighlight_CheckedChanged()`** — L25683 — `private void chkContainerHighlight_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkContainerHighlight` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkContainerShowRX_CheckedChanged()`** — L25699 — `private void chkContainerShowRX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkContainerShowRX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkContainerShowTX_CheckedChanged()`** — L25708 — `private void chkContainerShowTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkContainerShowTX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtContainerNotes_TextChanged()`** — L25717 — `private void txtContainerNotes_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtContainerNotes` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAddMeterItem_Click()`** — L25758 — `private void btnAddMeterItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAddMeterItem` is clicked.
  Called by: `.lstMetersAvailable_DoubleClick()` (same file)
- **`.lstMetersAvailable_SelectedIndexChanged()`** — L25775 — `private void lstMetersAvailable_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstMetersAvailable` selection changes.
  Called by: `.updateMeterLists()` (same file), `.btnAddMeterItem_Click()` (same file)
- **`.lstMetersInUse_SelectedIndexChanged()`** — L25781 — `private void lstMetersInUse_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstMetersInUse` selection changes.
  Called by: `.ForceAllEvents()` (same file), `.updateMeterLists()` (same file), `.btnRemoveMeterItem_Click()` (same file)
- **`.btnRemoveMeterItem_Click()`** — L25800 — `private void btnRemoveMeterItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRemoveMeterItem` is clicked.
  Called by: `.lstMetersInUse_DoubleClick()` (same file)
- **`.btnMeterUp_Click()`** — L25818 — `private void btnMeterUp_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMeterUp` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMeterDown_Click()`** — L25837 — `private void btnMeterDown_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMeterDown` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstMetersAvailable_DoubleClick()`** — L25856 — `private void lstMetersAvailable_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `lstMetersAvailable` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstMetersInUse_DoubleClick()`** — L25861 — `private void lstMetersInUse_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `lstMetersInUse` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstMetersAvailable_DrawItem()`** — L25866 — `private void lstMetersAvailable_DrawItem(object sender, DrawItemEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lstMetersInUse_DrawItem()`** — L25871 — `private void lstMetersInUse_DrawItem(object sender, DrawItemEventArgs e)`
  Called by: `.lstMetersAvailable_DrawItem()` (same file)
- **`.meterItemGroupIDfromSelected()`** — L25914 — `private string meterItemGroupIDfromSelected()`
  Called by: `.updateMeterType()` (same file), `.updateItemSettingsControlsForSelected()` (same file), `.variableInUse()` (same file), `.mmioSetupVariable()` (same file), `.updateWebImageState()` (same file)
- **`.meterItemGroupTypefromSelected()`** — L25925 — `private MeterType meterItemGroupTypefromSelected()`
  Called by: `.updateMeterType()` (same file), `.updateItemSettingsControlsForSelected()` (same file), `.btnMeterCopySettings_Click()` (same file), `.btnMeterPasteSettings_Click()` (same file), `.canPasteSettings()` (same file), `.variableInUse()` (same file) — and 3 more
- **`.chkMeterItemHistory_CheckedChanged()`** — L25936 — `private void chkMeterItemHistory_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemHistory` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateMeterType()`** — L25944 — `private MeterManager.clsIGSettings updateMeterType()`
  Called by: `.chkMeterItemHistory_CheckedChanged()` (same file), `.clrbtnMeterItemHistory_Changed()` (same file), `.tbMeterItemHistoryAlpha_Scroll()` (same file), `.chkMeterItemFadeOnRx_CheckedChanged()` (same file), `.chkMeterItemFadeOnTx_CheckedChanged()` (same file), `.chkMeterItemSegmented_CheckedChanged()` (same file) — and 268 more
- **`.updateItemSettingsControlsForSelected()`** — L26531 — `private void updateItemSettingsControlsForSelected()`
  Called by: `.lstMetersInUse_SelectedIndexChanged()` (same file), `.btnMeterPasteSettings_Click()` (same file), `.nudVoiceRecordingPlayback_slots_ValueChanged()` (same file), `.updateSlotSettings()` (same file), `.nudRecording_slot_settings_ValueChanged()` (same file)
- **`.updateHistoryControls()`** — L27389 — `private void updateHistoryControls(bool showHistory, Color c, bool updateColor = false)`
  Called by: `.chkMeterItemHistory_CheckedChanged()` (same file), `.updateItemSettingsControlsForSelected()` (same file)
- **`.updateSegmentedSolidControls()`** — L27401 — `private void updateSegmentedSolidControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMeterItemSegmented_CheckedChanged()` (same file), `.chkMeterItemSolid_CheckedChanged()` (same file)
- **`.updateTitleControls()`** — L27409 — `private void updateTitleControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMeterItemTitle_CheckedChanged()` (same file)
- **`.updateTitleControlsClock()`** — L27413 — `private void updateTitleControlsClock()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMMClockTitle_CheckedChanged()` (same file)
- **`.updatePeakValueControls()`** — L27417 — `private void updatePeakValueControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMeterItemPeakValue_CheckedChanged()` (same file)
- **`.updatePeakHoldControls()`** — L27421 — `private void updatePeakHoldControls(bool showPeakHold, Color c, bool updateColor = false)`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMeterItemPeakHold_CheckedChanged()` (same file)
- **`.clrbtnMeterItemHistory_Changed()`** — L27432 — `private void clrbtnMeterItemHistory_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tbMeterItemHistoryAlpha_Scroll()`** — L27437 — `private void tbMeterItemHistoryAlpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbMeterItemHistoryAlpha` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemFadeOnRx_CheckedChanged()`** — L27442 — `private void chkMeterItemFadeOnRx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemFadeOnRx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemFadeOnTx_CheckedChanged()`** — L27447 — `private void chkMeterItemFadeOnTx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemFadeOnTx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemSegmented_CheckedChanged()`** — L27452 — `private void chkMeterItemSegmented_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemSegmented` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemTitle_CheckedChanged()`** — L27460 — `private void chkMeterItemTitle_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemTitle` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemPeakValue_CheckedChanged()`** — L27466 — `private void chkMeterItemPeakValue_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemPeakValue` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemPeakHold_CheckedChanged()`** — L27472 — `private void chkMeterItemPeakHold_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemPeakHold` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemPeakHold_Changed()`** — L27481 — `private void clrbtnMeterItemPeakHold_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudMeterItemHistoryDuration_ValueChanged()`** — L27486 — `private void nudMeterItemHistoryDuration_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemHistoryDuration` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItemUpdateRate_ValueChanged()`** — L27491 — `private void nudMeterItemUpdateRate_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemUpdateRate` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItemAttackRate_ValueChanged()`** — L27496 — `private void nudMeterItemAttackRate_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemAttackRate` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItemDecayRate_ValueChanged()`** — L27501 — `private void nudMeterItemDecayRate_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemDecayRate` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemShadow_CheckedChanged()`** — L27506 — `private void chkMeterItemShadow_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemShadow` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemLow_Changed()`** — L27511 — `private void clrbtnMeterItemLow_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemHigh_Changed()`** — L27516 — `private void clrbtnMeterItemHigh_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemHBackground_Changed()`** — L27520 — `private void clrbtnMeterItemHBackground_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemMeterTitle_Changed()`** — L27524 — `private void clrbtnMeterItemMeterTitle_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemPeakValueColour_Changed()`** — L27529 — `private void clrbtnMeterItemPeakValueColour_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudMeterItemEyeScale_ValueChanged()`** — L27534 — `private void nudMeterItemEyeScale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemEyeScale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemDarkMode_CheckedChanged()`** — L27539 — `private void chkMeterItemDarkMode_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemDarkMode` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMaintainNFAdjustDeltaRX2_CheckedChanged()`** — L27544 — `private void chkMaintainNFAdjustDeltaRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMaintainNFAdjustDeltaRX2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMaintainNFAdjustDeltaRX1_CheckedChanged()`** — L27549 — `private void chkMaintainNFAdjustDeltaRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMaintainNFAdjustDeltaRX1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkContainerBorder_CheckedChanged()`** — L27554 — `private void chkContainerBorder_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkContainerBorder` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnContainerBackground_Changed()`** — L27564 — `private void clrbtnContainerBackground_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudMeterItemsPowerLimit_ValueChanged()`** — L27574 — `private void nudMeterItemsPowerLimit_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemsPowerLimit` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemSolid_CheckedChanged()`** — L27579 — `private void chkMeterItemSolid_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemSolid` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemSegmentedSolidColourHigh_Changed()`** — L27587 — `private void clrbtnMeterItemSegmentedSolidColourHigh_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemSegmentedSolidColourLow_Changed()`** — L27592 — `private void clrbtnMeterItemSegmentedSolidColourLow_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkMeterItemShowIndicator_CheckedChanged()`** — L27597 — `private void chkMeterItemShowIndicator_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemShowIndicator` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemShowSubIndicator_CheckedChanged()`** — L27602 — `private void chkMeterItemShowSubIndicator_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemShowSubIndicator` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.waveRecordSettingControlChanged()`** — L27606 — `private void waveRecordSettingControlChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.moveButtonBoxSettings()`** — L27610 — `private void moveButtonBoxSettings()`
  Called by: `.AfterConstructor()` (same file)
- **`.setupMMSettingsGroupBoxes()`** — L27636 — `private void setupMMSettingsGroupBoxes(MeterType mt, bool all = true)`
  Called by: `.btnContainerDelete_Click()` (same file), `.lstMetersInUse_SelectedIndexChanged()` (same file), `.updateItemSettingsControlsForSelected()` (same file), `.moveButtonBoxSettings()` (same file)
- **`.radMM12Clock_CheckedChanged()`** — L27810 — `private void radMM12Clock_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMM12Clock` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMM24Clock_CheckedChanged()`** — L27815 — `private void radMM24Clock_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMM24Clock` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMMClockTitle_CheckedChanged()`** — L27820 — `private void chkMMClockTitle_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMMClockTitle` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMMClockTitle_Changed()`** — L27826 — `private void clrbtnMMClockTitle_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMTime_Changed()`** — L27831 — `private void clrbtnMMTime_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMDate_Changed()`** — L27836 — `private void clrbtnMMDate_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMClockBackground_Changed()`** — L27841 — `private void clrbtnMMClockBackground_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayBackground_Changed()`** — L27846 — `private void clrbtnMMVfoDisplayBackground_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayTitle_Changed()`** — L27851 — `private void clrbtnMMVfoDisplayTitle_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayFrequency_Changed()`** — L27856 — `private void clrbtnMMVfoDisplayFrequency_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayMode_Changed()`** — L27861 — `private void clrbtnMMVfoDisplayMode_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplaySplitBack_Changed()`** — L27866 — `private void clrbtnMMVfoDisplaySplitBack_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplaySplit_Changed()`** — L27871 — `private void clrbtnMMVfoDisplaySplit_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayRx_Changed()`** — L27876 — `private void clrbtnMMVfoDisplayRx_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayTx_Changed()`** — L27881 — `private void clrbtnMMVfoDisplayTx_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayFilter_Changed()`** — L27886 — `private void clrbtnMMVfoDisplayFilter_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMMVfoDisplayBand_Changed()`** — L27891 — `private void clrbtnMMVfoDisplayBand_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudMeterItemEyeBezelScale_ValueChanged()`** — L27896 — `private void nudMeterItemEyeBezelScale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemEyeBezelScale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemPowerScale_Changed()`** — L27901 — `private void clrbtnMeterItemPowerScale_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudMeterItemIgnoreHistoryDuration_ValueChanged()`** — L27906 — `private void nudMeterItemIgnoreHistoryDuration_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemIgnoreHistoryDuration` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMeterCopySettings_Click()`** — L27914 — `private void btnMeterCopySettings_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMeterCopySettings` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMeterPasteSettings_Click()`** — L27933 — `private void btnMeterPasteSettings_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMeterPasteSettings` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.canPasteSettings()`** — L28030 — `private bool canPasteSettings()`
  Called by: `.lstMetersInUse_SelectedIndexChanged()` (same file), `.btnMeterPasteSettings_Click()` (same file)
- **`.chkHL2IOBoardPresent_CheckedChanged()`** — L28075 — `private void chkHL2IOBoardPresent_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHL2IOBoardPresent` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ShowMultiMeterSetupTab()`** — L28080 — `public void ShowMultiMeterSetupTab(string sID = "")`
  Shows multi meter setup tab.
  Called by: `.btnWebImage_goto_next_Click()` (same file)
- **`.udVSQLMuteTimeConstant_ValueChanged()`** — L28105 — `private void udVSQLMuteTimeConstant_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVSQLMuteTimeConstant` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udVSQLUnMuteTimeConstant_ValueChanged()`** — L28120 — `private void udVSQLUnMuteTimeConstant_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udVSQLUnMuteTimeConstant` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkConsoleDarkModeTitleBar_CheckedChanged()`** — L28135 — `private void chkConsoleDarkModeTitleBar_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkConsoleDarkModeTitleBar` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.OnBCDBandChangeHandler()`** — L28175 — `private void OnBCDBandChangeHandler(int rx, Band old_band, Band new_band)`
  Handles/raises the bcdband change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateUsbBCDdevice()`** — L28180 — `private void updateUsbBCDdevice(Band rx1band)`
  Called by: `.OnBCDBandChangeHandler()` (same file)
- **`.chkUsbBCD_CheckedChanged()`** — L28189 — `private void chkUsbBCD_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUsbBCD` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.CloseUsbBcdDevice()`** — L28233 — `public void CloseUsbBcdDevice()`
  Closes usb bcd device.
  Called by: `.chkUsbBCD_CheckedChanged()` (same file), `.comboUsbDevices_SelectedIndexChanged()` (same file)
- **`.comboUsbDevices_SelectedIndexChanged()`** — L28250 — `private void comboUsbDevices_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboUsbDevices` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tmrCheckProfile_Tick()`** — L28303 — `private void tmrCheckProfile_Tick(object sender, EventArgs e)`
  by checkTXProfileChanged2() in tmrCheckProfile_Tick() reading values fix #301
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLimit2Subnet_CheckedChanged()`** — L28323 — `private void chkLimit2Subnet_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLimit2Subnet` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDisplayIPPort_CheckedChanged()`** — L28328 — `private void chkDisplayIPPort_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisplayIPPort` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnClearTCISpots_Click()`** — L28333 — `private void btnClearTCISpots_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClearTCISpots` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkQuickSplit_CheckedChanged()`** — L28376 — `private void chkQuickSplit_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQuickSplit` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnQuickSplitDown5_Click()`** — L28384 — `private void btnQuickSplitDown5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnQuickSplitDown5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnQuickSplitUp5_Click()`** — L28389 — `private void btnQuickSplitUp5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnQuickSplitUp5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLinkRX0AF_CheckedChanged()`** — L28394 — `private void chkLinkRX0AF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLinkRX0AF` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLinkRX1AF_CheckedChanged()`** — L28401 — `private void chkLinkRX1AF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLinkRX1AF` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLinkRX2AF_CheckedChanged()`** — L28408 — `private void chkLinkRX2AF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLinkRX2AF` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLinkMaster_CheckedChanged()`** — L28415 — `private void chkLinkMaster_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLinkMaster` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyMeters_CheckedChanged()`** — L28422 — `private void chkLegacyMeters_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyMeters` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkHideLegacyMeters_CheckedChanged()`** — L28426 — `private void chkHideLegacyMeters_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHideLegacyMeters` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radSpaceBarVFOBTX_CheckedChanged()`** — L28432 — `private void radSpaceBarVFOBTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSpaceBarVFOBTX` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.lblTXProfileWarning_Click()`** — L28439 — `private void lblTXProfileWarning_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblTXProfileWarning` is clicked.
  Called by: `.tmrCheckProfile_Tick()` (same file)
- **`.chkQuickSplitPanAudio_CheckedChanged()`** — L28460 — `private void chkQuickSplitPanAudio_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQuickSplitPanAudio` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.getTextBoxWidthForReport()`** — L28467 — `private int getTextBoxWidthForReport()`
  Returns text box width for report.
  Called by: `.lblTXProfileWarning_Click()` (same file)
- **`.chkJoinBandEdges_CheckedChanged()`** — L28481 — `private void chkJoinBandEdges_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkJoinBandEdges` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkContainerNoTitle_CheckedChanged()`** — L28487 — `private void chkContainerNoTitle_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkContainerNoTitle` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAutoModeSwitchCWReturn_CheckedChanged()`** — L28497 — `private void chkAutoModeSwitchCWReturn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoModeSwitchCWReturn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.setEnabledAutoModeSwitchCWReturn()`** — L28506 — `private void setEnabledAutoModeSwitchCWReturn(bool bEnable)`
  Sets enabled auto mode switch cwreturn.
  Called by: `.chkCWAutoSwitchMode_CheckedChanged()` (same file)
- **`.nudAutoModeSwitchCWReturn_ValueChanged()`** — L28512 — `private void nudAutoModeSwitchCWReturn_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudAutoModeSwitchCWReturn` value changes.
  Called by: `.chkAutoModeSwitchCWReturn_CheckedChanged()` (same file)
- **`.btnZipDebug_Click()`** — L28517 — `private void btnZipDebug_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnZipDebug` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.buildZipFile()`** — L28521 — `private void buildZipFile(string version, string sourceDirectory)`
  Called by: `.btnZipDebug_Click()` (same file)
- **`.chkLogVoltsAmps_CheckedChanged()`** — L28588 — `private void chkLogVoltsAmps_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLogVoltsAmps` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tcAppearance_SelectedIndexChanged()`** — L28602 — `private void tcAppearance_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `tcAppearance` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.isSkinServerTabVisible()`** — L28616 — `private bool isSkinServerTabVisible(bool checkFormVisible = true)`
  Called by: `.Show()` (same file), `.tcAppearance_SelectedIndexChanged()` (same file), `.skinServersDataReceivedHandler()` (same file), `.tcSetup_SelectedIndexChanged()` (same file)
- **`.getSkinServers()`** — L28630 — `private void getSkinServers()`
  Returns skin servers.
  Called by: `.Show()` (same file), `.tcAppearance_SelectedIndexChanged()` (same file), `.tcSetup_SelectedIndexChanged()` (same file)
- **`.hideAllSkinServerRelatedControls()`** — L28643 — `private void hideAllSkinServerRelatedControls()`
  Called by: `.Hide()` (same file), `.tcAppearance_SelectedIndexChanged()` (same file), `.getSkinServers()` (same file), `.comboSkinServerList_SelectedIndexChanged()` (same file)
- **`.skinDataReceivedHandler()`** — L28653 — `private void skinDataReceivedHandler(object sender, SkinsData e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateSelectedSkin()`** — L28696 — `private void updateSelectedSkin()`
  Called by: `.lstAvailableSkins_SelectedIndexChanged()` (same file)
- **`.validateDate()`** — L28751 — `private string validateDate(string sDate)`
  Called by: `.updateSelectedSkin()` (same file)
- **`.validateVersion()`** — L28757 — `private string validateVersion(string sVersion)`
  Called by: `.updateSelectedSkin()` (same file)
- **`.skinServersDataReceivedHandler()`** — L28780 — `private void skinServersDataReceivedHandler(object sender, SkinServersData e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.imageLoadedHandler()`** — L28819 — `private void imageLoadedHandler(object sender, SkinHttpImage e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnRefreshSkinsForServer_Click()`** — L28849 — `private void btnRefreshSkinsForServer_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRefreshSkinsForServer` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSkinsHomepage_Click()`** — L28860 — `private void btnSkinsHomepage_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSkinsHomepage` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSkinsDonate_Click()`** — L28865 — `private void btnSkinsDonate_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSkinsDonate` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSkinHomepage_Click()`** — L28870 — `private void btnSkinHomepage_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSkinHomepage` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.fileDownloadHandler()`** — L28874 — `private void fileDownloadHandler(object sender, SkinFileDownload e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tryRemoveDownload()`** — L29026 — `private void tryRemoveDownload(string path)`
  Called by: `.fileDownloadHandler()` (same file)
- **`.getFileFromUrl()`** — L29035 — `private string getFileFromUrl(string sUrl, bool bWithoutExtenstion = true)`
  Returns file from url.
  Called by: `.fileDownloadHandler()` (same file)
- **`.btnDownloadSkin_Click()`** — L29048 — `private void btnDownloadSkin_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDownloadSkin` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.downloadSkin()`** — L29052 — `private void downloadSkin(bool bUseAfterDownload)`
  Called by: `.btnDownloadSkin_Click()` (same file)
- **`.lstAvailableSkins_SelectedIndexChanged()`** — L29078 — `private void lstAvailableSkins_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstAvailableSkins` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.isSkinZipFile()`** — L29086 — `private bool isSkinZipFile(string filePath, string sFilename, out bool usesFilesInRoot, out bool bSkinsFolderFoundInRoot, out bool bConsoleFolderFoundInRoot, out bool bMeterFolderF`
  Called by: `.fileDownloadHandler()` (same file)
- **`.doesFolderExistInZip()`** — L29149 — `private bool doesFolderExistInZip(string zipFilePath, string folderName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.extractImagesFilesFromZip()`** — L29157 — `private static string extractImagesFilesFromZip(string sourceZipFilePath, string outputPath)`
  Called by: `.fileDownloadHandler()` (same file)
- **`.comboSkinServerList_SelectedIndexChanged()`** — L29214 — `private void comboSkinServerList_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboSkinServerList` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRemoveSkin_Click()`** — L29222 — `private void btnRemoveSkin_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRemoveSkin` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnOpenSkinsFolder_Click()`** — L29284 — `private void btnOpenSkinsFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnOpenSkinsFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPurgeBuffers_CheckedChanged()`** — L29296 — `private void chkPurgeBuffers_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPurgeBuffers` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.tcSetup_SelectedIndexChanged()`** — L29302 — `private void tcSetup_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `tcSetup` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udFMLowCutRX_ValueChanged()`** — L29309 — `private void udFMLowCutRX_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udFMLowCutRX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udFMHighCutRX_ValueChanged()`** — L29322 — `private void udFMHighCutRX_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udFMHighCutRX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udFMLowCutTX_ValueChanged()`** — L29335 — `private void udFMLowCutTX_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udFMLowCutTX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udFMHighCutTX_ValueChanged()`** — L29344 — `private void udFMHighCutTX_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udFMHighCutTX` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSwapIQVac1_CheckedChanged()`** — L29353 — `private void chkSwapIQVac1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSwapIQVac1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkSwapIQVac2_CheckedChanged()`** — L29359 — `private void chkSwapIQVac2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSwapIQVac2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkPHROTReverse_CheckedChanged()`** — L29373 — `private void chkPHROTReverse_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPHROTReverse` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkRecoverPAProfileFromTXProfile_CheckedChanged()`** — L29380 — `private void chkRecoverPAProfileFromTXProfile_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecoverPAProfileFromTXProfile` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC1ExclusiveOut_CheckedChanged()`** — L29400 — `private void chkVAC1ExclusiveOut_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1ExclusiveOut` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2ExclusiveOut_CheckedChanged()`** — L29417 — `private void chkVAC2ExclusiveOut_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2ExclusiveOut` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC1ExclusiveIn_CheckedChanged()`** — L29434 — `private void chkVAC1ExclusiveIn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC1ExclusiveIn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkVAC2ExclusiveIn_CheckedChanged()`** — L29451 — `private void chkVAC2ExclusiveIn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2ExclusiveIn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnReleaseNotes_Click()`** — L29468 — `private void btnReleaseNotes_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnReleaseNotes` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkIgnore14bitMidiMessages_CheckedChanged()`** — L29473 — `private void chkIgnore14bitMidiMessages_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkIgnore14bitMidiMessages` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkMidiControlIDincludesChannel_CheckedChanged()`** — L29479 — `private void chkMidiControlIDincludesChannel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMidiControlIDincludesChannel` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkMidiControlIDincludesStatus_CheckedChanged()`** — L29487 — `private void chkMidiControlIDincludesStatus_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMidiControlIDincludesStatus` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkForceATTwhenOutPowerChanges_CheckedChanged()`** — L29493 — `private void chkForceATTwhenOutPowerChanges_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkForceATTwhenOutPowerChanges` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkToTMox_CheckedChanged()`** — L29500 — `private void chkToTMox_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkToTMox` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.udMoxToTSeconds_ValueChanged()` (same file)
- **`.udMoxToTSeconds_ValueChanged()`** — L29508 — `private void udMoxToTSeconds_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udMoxToTSeconds` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkToTPing_CheckedChanged()`** — L29513 — `private void chkToTPing_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkToTPing` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.udPingToTSeconds_ValueChanged()` (same file), `.txtToTPingIP_TextChanged()` (same file)
- **`.udPingToTSeconds_ValueChanged()`** — L29531 — `private void udPingToTSeconds_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPingToTSeconds` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtToTPingIP_TextChanged()`** — L29536 — `private void txtToTPingIP_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtToTPingIP` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnPingDef_Click()`** — L29541 — `private void btnPingDef_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPingDef` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnNoiseFloorText_Changed()`** — L29546 — `private void clrbtnNoiseFloorText_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.chkMICVOXAllowBypass_CheckedChanged()`** — L29582 — `private void chkMICVOXAllowBypass_CheckedChanged(object sender, EventArgs e)`
  [2.10.3.5]W4WMT implements #87
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkStopRX1WaterfallOnTx_CheckedChanged()`** — L29588 — `private void chkStopRX1WaterfallOnTx_CheckedChanged(object sender, EventArgs e)`
  [2.10.3.5]MW0LGE implements #306
  Called by: `.ForceAllEvents()` (same file)
- **`.chkStopRX2WaterfallOnTx_CheckedChanged()`** — L29594 — `private void chkStopRX2WaterfallOnTx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkStopRX2WaterfallOnTx` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.SerialCatState()`** — L29601 — `public int[] SerialCatState()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkAutoPowerOn_CheckedChanged()`** — L29615 — `private void chkAutoPowerOn_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoPowerOn` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnRX1CopyLowHighWaterfall_Click()`** — L29621 — `private void btnRX1CopyLowHighWaterfall_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX1CopyLowHighWaterfall` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRX1CopyMinMaxSpectrumGrid_Click()`** — L29757 — `private void btnRX1CopyMinMaxSpectrumGrid_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX1CopyMinMaxSpectrumGrid` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRX2CopyLowHighWaterfall_Click()`** — L29893 — `private void btnRX2CopyLowHighWaterfall_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX2CopyLowHighWaterfall` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRX2CopyMinMaxSpectrumGrid_Click()`** — L30029 — `private void btnRX2CopyMinMaxSpectrumGrid_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX2CopyMinMaxSpectrumGrid` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRX1PBsnrReset_Click()`** — L30165 — `private void btnRX1PBsnrReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX1PBsnrReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRX2PBsnrReset_Click()`** — L30170 — `private void btnRX2PBsnrReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRX2PBsnrReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudPBsnrShiftRx1_ValueChanged()`** — L30175 — `private void nudPBsnrShiftRx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPBsnrShiftRx1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudPBsnrShiftRx2_ValueChanged()`** — L30181 — `private void nudPBsnrShiftRx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPBsnrShiftRx2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkCWbecomesCWUabove10mhz_CheckedChanged()`** — L30187 — `private void chkCWbecomesCWUabove10mhz_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCWbecomesCWUabove10mhz` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucOutPinsLedStripHF_Click()`** — L30192 — `private void ucOutPinsLedStripHF_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ucOutPinsLedStripHF` is clicked.
  Called by: `.btnI2CWrite_MouseDown()` (same file), `.ucOutPinsLedStripHF_MouseDown()` (same file)
- **`.ucOutPinsLedStripHF_MouseDown()`** — L30225 — `private void ucOutPinsLedStripHF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ucOutPinsLedStripHF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkKWAI_port1_CheckedChanged()`** — L30246 — `private void chkKWAI_port1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkKWAI_port1` checked state changes.
  Called by: `.updatePortAIstate()` (same file)
- **`.chkKWAI_port2_CheckedChanged()`** — L30258 — `private void chkKWAI_port2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkKWAI_port2` checked state changes.
  Called by: `.updatePortAIstate()` (same file)
- **`.chkKWAI_port3_CheckedChanged()`** — L30270 — `private void chkKWAI_port3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkKWAI_port3` checked state changes.
  Called by: `.updatePortAIstate()` (same file)
- **`.chkKWAI_port4_CheckedChanged()`** — L30282 — `private void chkKWAI_port4_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkKWAI_port4` checked state changes.
  Called by: `.updatePortAIstate()` (same file)
- **`.chkKWAI_tcp_CheckedChanged()`** — L30294 — `private void chkKWAI_tcp_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkKWAI_tcp` checked state changes.
  Called by: `.updatePortAIstate()` (same file)
- **`.updatePortAIstate()`** — L30306 — `private void updatePortAIstate(bool enabled)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.udCWEdgeLength_ValueChanged()`** — L30330 — `private void udCWEdgeLength_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udCWEdgeLength` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2TRND_CheckedChanged()`** — L30347 — `private void radDSPNR2TRND_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2TRND` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2TRNDRX2_CheckedChanged()`** — L30361 — `private void radDSPNR2TRNDRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2TRNDRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2NSTAT_CheckedChanged()`** — L30375 — `private void radDSPNR2NSTAT_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2NSTAT` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radDSPNR2NSTATRX2_CheckedChanged()`** — L30385 — `private void radDSPNR2NSTATRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDSPNR2NSTATRX2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPNR2trainThresh_ValueChanged()`** — L30395 — `private void udDSPNR2trainThresh_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNR2trainThresh` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPNR2trainThreshRX2_ValueChanged()`** — L30402 — `private void udDSPNR2trainThreshRX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNR2trainThreshRX2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radBelow30_CheckedChanged()`** — L30409 — `private void radBelow30_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radBelow30` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBelow144_CheckedChanged()`** — L30415 — `private void radBelow144_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radBelow144` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udSwrProtectionLimit_ValueChanged()`** — L30421 — `private void udSwrProtectionLimit_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udSwrProtectionLimit` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTunePowerSwrIgnore_ValueChanged()`** — L30426 — `private void udTunePowerSwrIgnore_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTunePowerSwrIgnore` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowFormStartup_CheckedChanged()`** — L30431 — `private void chkShowFormStartup_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowFormStartup` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateAutoStartForms()`** — L30441 — `public void UpdateAutoStartForms()`
  Updates auto start forms.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkDisableHPFonPS_CheckedChanged()`** — L30454 — `private void chkDisableHPFonPS_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDisableHPFonPS` checked state changes.
  Called by: `.PerformDelayedInitalistion()` (same file)
- **`.txtAutoLaunchFile_TextChanged()`** — L30475 — `private void txtAutoLaunchFile_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtAutoLaunchFile` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAutoLaunchSelectFile_Click()`** — L30480 — `private void btnAutoLaunchSelectFile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAutoLaunchSelectFile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAutoLaunch_CheckedChanged()`** — L30499 — `private void chkAutoLaunch_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoLaunch` checked state changes.
  Called by: `.updateAutoLaunchControls()` (same file)
- **`.updateAutoLaunchControls()`** — L30511 — `private void updateAutoLaunchControls()`
  Called by: `.ForceAllEvents()` (same file)
- **`.GetAutoLaunchFiles()`** — L30521 — `public string[] GetAutoLaunchFiles()`
  Returns auto launch files.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnTXAttenuationBackground_Changed()`** — L30547 — `private void clrbtnTXAttenuationBackground_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.chkMeterItemFadeOnRxSpacer_CheckedChanged()`** — L30553 — `private void chkMeterItemFadeOnRxSpacer_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemFadeOnRxSpacer` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemFadeOnTxSpacer_CheckedChanged()`** — L30558 — `private void chkMeterItemFadeOnTxSpacer_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemFadeOnTxSpacer` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemHBackgroundSpacerRX_Changed()`** — L30563 — `private void clrbtnMeterItemHBackgroundSpacerRX_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudMeterItemSpacerPadding_ValueChanged()`** — L30568 — `private void nudMeterItemSpacerPadding_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemSpacerPadding` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemHBackgroundSpacerTX_Changed()`** — L30573 — `private void clrbtnMeterItemHBackgroundSpacerTX_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkTextOverlay_ShowPanel_CheckedChanged()`** — L30578 — `private void chkTextOverlay_ShowPanel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTextOverlay_ShowPanel` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateTextOverlayPanelControls()`** — L30583 — `private void updateTextOverlayPanelControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkTextOverlay_ShowPanel_CheckedChanged()` (same file)
- **`.clrbtnTextOverlay_PanelBackground_Changed()`** — L30595 — `private void clrbtnTextOverlay_PanelBackground_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudTextOverlay_PanelPadding_ValueChanged()`** — L30600 — `private void nudTextOverlay_PanelPadding_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudTextOverlay_PanelPadding` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtTextOverlay_RXText_TextChanged()`** — L30605 — `private void txtTextOverlay_RXText_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtTextOverlay_RXText` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtTextOverlay_TXText_TextChanged()`** — L30610 — `private void txtTextOverlay_TXText_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtTextOverlay_TXText` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnTextOverlay_Font1_Click()`** — L30616 — `private void btnTextOverlay_Font1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTextOverlay_Font1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.isFontTrueType()`** — L30629 — `public bool isFontTrueType(Font f)`
  Called by: `.btnTextOverlay_Font1_Click()` (same file), `.btnTextOverlay_Font2_Click()` (same file), `.btnBandButtons_font_Click()` (same file)
- **`.btnTextOverlay_Font2_Click()`** — L30650 — `private void btnTextOverlay_Font2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTextOverlay_Font2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnTextOverlay_TextColour1_Changed()`** — L30664 — `private void clrbtnTextOverlay_TextColour1_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnTextOverlay_TextColour2_Changed()`** — L30669 — `private void clrbtnTextOverlay_TextColour2_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudTextOverlay_RXxOffset_ValueChanged()`** — L30674 — `private void nudTextOverlay_RXxOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudTextOverlay_RXxOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudTextOverlay_RXyOffset_ValueChanged()`** — L30679 — `private void nudTextOverlay_RXyOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudTextOverlay_RXyOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudTextOverlay_TXxOffset_ValueChanged()`** — L30684 — `private void nudTextOverlay_TXxOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudTextOverlay_TXxOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudTextOverlay_TXyOffset_ValueChanged()`** — L30689 — `private void nudTextOverlay_TXyOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudTextOverlay_TXyOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTextOverlay_FadeOnRX_CheckedChanged()`** — L30694 — `private void chkTextOverlay_FadeOnRX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTextOverlay_FadeOnRX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTextOverlay_FadeOnTX_CheckedChanged()`** — L30699 — `private void chkTextOverlay_FadeOnTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTextOverlay_FadeOnTX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnTextOverlay_copyoffsets_Click()`** — L30704 — `private void btnTextOverlay_copyoffsets_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTextOverlay_copyoffsets` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnTextOverlay_PanelBackgroundTX_Changed()`** — L30710 — `private void clrbtnTextOverlay_PanelBackgroundTX_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnTextOverlay_TextBackColour1_Changed()`** — L30715 — `private void clrbtnTextOverlay_TextBackColour1_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnTextOverlay_TextBackColour2_Changed()`** — L30720 — `private void clrbtnTextOverlay_TextBackColour2_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkTextOverlay_textback1_CheckedChanged()`** — L30725 — `private void chkTextOverlay_textback1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTextOverlay_textback1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTextOverlay_textback2_CheckedChanged()`** — L30731 — `private void chkTextOverlay_textback2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTextOverlay_textback2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateTextOverlayBackTextControls()`** — L30737 — `private void updateTextOverlayBackTextControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkTextOverlay_textback1_CheckedChanged()` (same file), `.chkTextOverlay_textback2_CheckedChanged()` (same file)
- **`.init_lstMMIO()`** — L30746 — `private void init_lstMMIO()`
  Called by: `.ForceAllEvents()` (same file)
- **`.MultiMeterIO_ListenerRunning()`** — L30797 — `private void MultiMeterIO_ListenerRunning(Guid guid, MultiMeterIO.MMIOType type, bool enabled)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeterIO_ClientConnected()`** — L30809 — `private void MultiMeterIO_ClientConnected(Guid guid)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeterIO_ClientDisconnected()`** — L30822 — `private void MultiMeterIO_ClientDisconnected(Guid guid)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeterIOStopTimers()`** — L30838 — `public void MultiMeterIOStopTimers()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRxTimerTick()`** — L30855 — `private void OnRxTimerTick(object sender, ElapsedEventArgs e)`
  Handles/raises the rx timer tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTxTimerTick()`** — L30871 — `private void OnTxTimerTick(object sender, ElapsedEventArgs e)`
  Handles/raises the tx timer tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeterIO_ReceivedDataString()`** — L30887 — `private void MultiMeterIO_ReceivedDataString(Guid guid, string dataString)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeterIO_TransmittedData()`** — L30899 — `private void MultiMeterIO_TransmittedData(Guid guid)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateFromRecievedData()`** — L30911 — `private void updateFromRecievedData(Guid guid, string dataString)`
  Called by: `.MultiMeterIO_ReceivedDataString()` (same file)
- **`.updateFromTransmittedData()`** — L30924 — `private void updateFromTransmittedData(Guid guid)`
  Called by: `.MultiMeterIO_TransmittedData()` (same file)
- **`.updateVariableList()`** — L30935 — `private void updateVariableList()`
  Called by: `.tcSetup_SelectedIndexChanged()` (same file), `.updateFromRecievedData()` (same file), `.lstMMIO_network_list_SelectedIndexChanged()` (same file), `.btnMMIO_network_remove_variable_Click()` (same file), `.tcCAT_SelectedIndexChanged()` (same file)
- **`.lstMMIO_network_list_SelectedIndexChanged()`** — L30997 — `private void lstMMIO_network_list_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstMMIO_network_list` selection changes.
  Called by: `.ForceAllEvents()` (same file), `.MultiMeterIO_ListenerRunning()` (same file), `.MultiMeterIO_ClientConnected()` (same file), `.MultiMeterIO_ClientDisconnected()` (same file)
- **`.showSerialPortPicker()`** — L31069 — `private void showSerialPortPicker(MultiMeterIO.MMIOType type, string existsing_com_port, int baud_rate, int data_bits, StopBits stop_bits, Parity parity)`
  Called by: `.addEditConnector()` (same file)
- **`.addEditConnector()`** — L31176 — `private void addEditConnector(MultiMeterIO.MMIOType type, bool check_same, string existing_ip_port = "", string existsing_com_port = "", int baud_rate = 0, int data_bits = 0, StopB`
  Called by: `.radMMIO_network_add_udp_Click()` (same file), `.btnMMIO_network_add_tcpip_Click()` (same file), `.txtMMIO_network_ip_port_Click()` (same file), `.btnMMIO_network_ip_port_ip4_Click()` (same file), `.btnMMIO_network_add_tcpip_client_Click()` (same file), `.btnMMIO_network_add_serial_Click()` (same file)
- **`.radMMIO_network_add_udp_Click()`** — L31328 — `private void radMMIO_network_add_udp_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radMMIO_network_add_udp` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_add_tcpip_Click()`** — L31333 — `private void btnMMIO_network_add_tcpip_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_add_tcpip` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_delete_Click()`** — L31338 — `private void btnMMIO_network_delete_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_delete` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_ip_port_TextChanged()`** — L31349 — `private void txtMMIO_network_ip_port_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_ip_port` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_ip_port_Click()`** — L31355 — `private void txtMMIO_network_ip_port_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_ip_port` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_ip_port_ip4_Click()`** — L31374 — `private void btnMMIO_network_ip_port_ip4_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_ip_port_ip4` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_4char_TextChanged()`** — L31394 — `private void txtMMIO_network_4char_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_4char` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMMIO_network_in_CheckedChanged()`** — L31400 — `private void radMMIO_network_in_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMMIO_network_in` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SendMessage()`** — L31420 — `[DllImport("user32.dll")] public static extern int SendMessage(IntPtr hWnd, int wMsg, IntPtr wParam, IntPtr lParam)`
  Sends message.
  Called by: `.adjustHeightOfMMIOcomboItem()` (same file)
- **`.adjustHeightOfMMIOcomboItem()`** — L31422 — `private void adjustHeightOfMMIOcomboItem(int i, int height)`
  Called by: `.radMMIO_network_in_CheckedChanged()` (same file), `.radMMIO_network_out_CheckedChanged()` (same file), `.radMMIO_network_both_CheckedChanged()` (same file)
- **`.radMMIO_network_out_CheckedChanged()`** — L31429 — `private void radMMIO_network_out_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMMIO_network_out` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMMIO_network_both_CheckedChanged()`** — L31448 — `private void radMMIO_network_both_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMMIO_network_both` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateTerminators()`** — L31466 — `private void updateTerminators(MultiMeterIO.clsMMIO mmio)`
  Called by: `.lstMMIO_network_list_SelectedIndexChanged()` (same file), `.radMMIO_network_in_CheckedChanged()` (same file), `.radMMIO_network_out_CheckedChanged()` (same file), `.radMMIO_network_both_CheckedChanged()` (same file), `.comboMMIO_network_terminator_in_SelectedIndexChanged()` (same file), `.comboMMIO_network_terminator_out_SelectedIndexChanged()` (same file)
- **`.updateDirection()`** — L31473 — `private void updateDirection(MultiMeterIO.clsMMIO mmio)`
  Called by: `.lstMMIO_network_list_SelectedIndexChanged()` (same file), `.radMMIO_network_in_CheckedChanged()` (same file), `.radMMIO_network_out_CheckedChanged()` (same file), `.radMMIO_network_both_CheckedChanged()` (same file)
- **`.comboMMIO_network_format_in_SelectedIndexChanged()`** — L31599 — `private void comboMMIO_network_format_in_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboMMIO_network_format_in` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstMMIO_network_list_DrawItem()`** — L31695 — `private void lstMMIO_network_list_DrawItem(object sender, DrawItemEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.itemHeight()`** — L31772 — `private int itemHeight(clsMultiMeterIOComboboxItem mmioci)`
  Called by: `.radMMIO_network_in_CheckedChanged()` (same file), `.radMMIO_network_out_CheckedChanged()` (same file), `.radMMIO_network_both_CheckedChanged()` (same file), `.lstMMIO_network_list_MeasureItem()` (same file)
- **`.lstMMIO_network_list_MeasureItem()`** — L31793 — `private void lstMMIO_network_list_MeasureItem(object sender, MeasureItemEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkMMIO_network_enabled_CheckedChanged()`** — L31807 — `private void chkMMIO_network_enabled_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMMIO_network_enabled` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_ip_port_KeyPress()`** — L31832 — `private void txtMMIO_network_ip_port_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_ip_port` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_4char_KeyPress()`** — L31837 — `private void txtMMIO_network_4char_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_4char` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_copy4char_Click()`** — L31841 — `private void btnMMIO_network_copy4char_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_copy4char` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_remove_variable_Click()`** — L31845 — `private void btnMMIO_network_remove_variable_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_remove_variable` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstMMIO_network_variables_SelectedIndexChanged()`** — L31870 — `private void lstMMIO_network_variables_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstMMIO_network_variables` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_copyvariable_clipboard_Click()`** — L31884 — `private void btnMMIO_network_copyvariable_clipboard_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_copyvariable_clipboard` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_variable_Click()`** — L31891 — `private void btnMMIO_variable_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_variable` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboMMIO_network_terminator_in_SelectedIndexChanged()`** — L31895 — `private void comboMMIO_network_terminator_in_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboMMIO_network_terminator_in` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_variable_2_Click()`** — L31910 — `private void btnMMIO_variable_2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_variable_2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.variableInUse()`** — L31914 — `private bool variableInUse(int variable)`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.mmioSetupVariable()` (same file)
- **`.mmioSetupVariable()`** — L31933 — `private void mmioSetupVariable(int variable)`
  Called by: `.btnMMIO_variable_Click()` (same file), `.btnMMIO_variable_2_Click()` (same file), `.btnMMIO_variable_rotator_Click()` (same file), `.btnMMIO_variable_2_rotator_Click()` (same file), `.btnMMIO_variable_history_Click()` (same file), `.btnMMIO_variable_2_history_Click()` (same file)
- **`.comboMMIO_network_terminator_out_SelectedIndexChanged()`** — L32005 — `private void comboMMIO_network_terminator_out_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboMMIO_network_terminator_out` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboMMIO_network_format_out_SelectedIndexChanged()`** — L32020 — `private void comboMMIO_network_format_out_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboMMIO_network_format_out` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_terminator_in_custom_TextChanged()`** — L32031 — `private void txtMMIO_network_terminator_in_custom_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_terminator_in_custom` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_terminator_out_custom_TextChanged()`** — L32042 — `private void txtMMIO_network_terminator_out_custom_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_terminator_out_custom` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_udp_endpoint_ip_port_Click()`** — L32052 — `private void btnMMIO_network_udp_endpoint_ip_port_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_udp_endpoint_ip_port` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_udp_endpoint_ip_port_KeyPress()`** — L32071 — `private void txtMMIO_network_udp_endpoint_ip_port_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_udp_endpoint_ip_port` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setUDPEndpoint()`** — L32075 — `private bool setUDPEndpoint()`
  Sets udpendpoint.
  Called by: `.setupUDPEndpoint()` (same file), `.txtMMIO_network_udp_endpoint_ip_port_TextChanged()` (same file)
- **`.setupUDPEndpoint()`** — L32107 — `private void setupUDPEndpoint(string existing_ip_port = "")`
  Called by: `.btnMMIO_network_udp_endpoint_ip_port_Click()` (same file), `.txtMMIO_network_udp_endpoint_ip_port_Click()` (same file)
- **`.txtMMIO_network_udp_endpoint_ip_port_Click()`** — L32177 — `private void txtMMIO_network_udp_endpoint_ip_port_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_udp_endpoint_ip_port` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_network_udp_endpoint_ip_port_TextChanged()`** — L32182 — `private void txtMMIO_network_udp_endpoint_ip_port_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_network_udp_endpoint_ip_port` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_remove_all_variables_Click()`** — L32189 — `private void btnMMIO_network_remove_all_variables_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_remove_all_variables` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_network_add_tcpip_client_Click()`** — L32205 — `private void btnMMIO_network_add_tcpip_client_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_add_tcpip_client` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudDataOutNode_sendinterval_ValueChanged()`** — L32211 — `private void nudDataOutNode_sendinterval_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDataOutNode_sendinterval` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtDataOutNode_4charID_TextChanged()`** — L32216 — `private void txtDataOutNode_4charID_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtDataOutNode_4charID` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemRotatorCardinals_CheckedChanged()`** — L32222 — `private void chkMeterItemRotatorCardinals_CheckedChanged(object sender, EventArgs e)`
  rotator
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemFadeOnRxRotator_CheckedChanged()`** — L32227 — `private void chkMeterItemFadeOnRxRotator_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemFadeOnRxRotator` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemFadeOnTxRotator_CheckedChanged()`** — L32232 — `private void chkMeterItemFadeOnTxRotator_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemFadeOnTxRotator` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemDarkModeRotator_CheckedChanged()`** — L32237 — `private void chkMeterItemDarkModeRotator_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemDarkModeRotator` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItemUpdateRateRotator_ValueChanged()`** — L32242 — `private void nudMeterItemUpdateRateRotator_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemUpdateRateRotator` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemHBackgroundRotator_Changed()`** — L32247 — `private void clrbtnMeterItemHBackgroundRotator_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemRotatorArrow_Changed()`** — L32252 — `private void clrbtnMeterItemRotatorArrow_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemRotatorLargeDot_Changed()`** — L32257 — `private void clrbtnMeterItemRotatorLargeDot_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemRotatorSmallDot_Changed()`** — L32262 — `private void clrbtnMeterItemRotatorSmallDot_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemRotatorBeamWidth_Changed()`** — L32267 — `private void clrbtnMeterItemRotatorBeamWidth_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkMeterItemRotatorShowBeamWidth_CheckedChanged()`** — L32272 — `private void chkMeterItemRotatorShowBeamWidth_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemRotatorShowBeamWidth` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateShowBeamWidthControls()`** — L32277 — `private void updateShowBeamWidthControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMeterItemRotatorShowBeamWidth_CheckedChanged()` (same file)
- **`.nudMeterItemRotatorBeamWidth_ValueChanged()`** — L32286 — `private void nudMeterItemRotatorBeamWidth_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemRotatorBeamWidth` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemRotatorText_Changed()`** — L32291 — `private void clrbtnMeterItemRotatorText_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnMMIO_variable_rotator_Click()`** — L32296 — `private void btnMMIO_variable_rotator_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_variable_rotator` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_variable_2_rotator_Click()`** — L32301 — `private void btnMMIO_variable_2_rotator_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_variable_2_rotator` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnTextOverlay_copyfonts_Click()`** — L32305 — `private void btnTextOverlay_copyfonts_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTextOverlay_copyfonts` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMeterItemRotatorAllowControl_CheckedChanged()`** — L32315 — `private void chkMeterItemRotatorAllowControl_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMeterItemRotatorAllowControl` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMeterItemRotatorControlColour_Changed()`** — L32321 — `private void clrbtnMeterItemRotatorControlColour_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.txtMeterItemRotatorAZcommand_TextChanged()`** — L32326 — `private void txtMeterItemRotatorAZcommand_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMeterItemRotatorAZcommand` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMeterItemRotatorELEcommand_TextChanged()`** — L32331 — `private void txtMeterItemRotatorELEcommand_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMeterItemRotatorELEcommand` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateRotatorControlControls()`** — L32335 — `private void updateRotatorControlControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMeterItemRotatorAllowControl_CheckedChanged()` (same file)
- **`.bntMultiMeterItemRotator_default_pstRotator_Click()`** — L32351 — `private void bntMultiMeterItemRotator_default_pstRotator_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `bntMultiMeterItemRotator_default_pstRotator` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtRotator_4charID_TextChanged()`** — L32358 — `private void txtRotator_4charID_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtRotator_4charID` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLedIndicator_ShowPanel_CheckedChanged()`** — L32363 — `private void chkLedIndicator_ShowPanel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLedIndicator_ShowPanel` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnLedIndicator_PanelBackground_Changed()`** — L32369 — `private void clrbtnLedIndicator_PanelBackground_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudLedIndicator_PanelPadding_ValueChanged()`** — L32374 — `private void nudLedIndicator_PanelPadding_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudLedIndicator_PanelPadding` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLedIndicator_FadeOnRX_CheckedChanged()`** — L32379 — `private void chkLedIndicator_FadeOnRX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLedIndicator_FadeOnRX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLedIndicator_FadeOnTX_CheckedChanged()`** — L32384 — `private void chkLedIndicator_FadeOnTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLedIndicator_FadeOnTX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtLedIndicator_condition_TextChanged()`** — L32389 — `private void txtLedIndicator_condition_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtLedIndicator_condition` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnLedIndicator_true_Changed()`** — L32394 — `private void clrbtnLedIndicator_true_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnLedIndicator_false_Changed()`** — L32399 — `private void clrbtnLedIndicator_false_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnLedIndicator_copy_truefalse_colours_Click()`** — L32404 — `private void btnLedIndicator_copy_truefalse_colours_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLedIndicator_copy_truefalse_colours` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudLedIndicator_xOffset_ValueChanged()`** — L32409 — `private void nudLedIndicator_xOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudLedIndicator_xOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudLedIndicator_yOffset_ValueChanged()`** — L32414 — `private void nudLedIndicator_yOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudLedIndicator_yOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudLedIndicator_xSize_ValueChanged()`** — L32419 — `private void nudLedIndicator_xSize_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudLedIndicator_xSize` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudLedIndicator_ySize_ValueChanged()`** — L32424 — `private void nudLedIndicator_ySize_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudLedIndicator_ySize` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnLedIndicator_copy_sizex_to_y_Click()`** — L32429 — `private void btnLedIndicator_copy_sizex_to_y_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLedIndicator_copy_sizex_to_y` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateLedIndicatorPanelControls()`** — L32433 — `private void updateLedIndicatorPanelControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkLedIndicator_ShowPanel_CheckedChanged()` (same file)
- **`.updateLedValidControls()`** — L32445 — `private void updateLedValidControls()`
  Called by: `.updateMeterType()` (same file), `.updateItemSettingsControlsForSelected()` (same file), `.tmrLedValid_Tick()` (same file)
- **`.chkLed_show_true_CheckedChanged()`** — L32468 — `private void chkLed_show_true_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLed_show_true` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLed_show_false_CheckedChanged()`** — L32473 — `private void chkLed_show_false_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLed_show_false` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radLed_light_on_off_CheckedChanged()`** — L32478 — `private void radLed_light_on_off_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radLed_light_on_off` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radLed_light_blink_CheckedChanged()`** — L32484 — `private void radLed_light_blink_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radLed_light_blink` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radLed_light_pulsate_CheckedChanged()`** — L32490 — `private void radLed_light_pulsate_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radLed_light_pulsate` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMMVfoDigitHighlight_Changed()`** — L32496 — `private void clrbtnMMVfoDigitHighlight_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnMMIO_network_add_serial_Click()`** — L32501 — `private void btnMMIO_network_add_serial_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_network_add_serial` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtWebImage_url_TextChanged()`** — L32506 — `private void txtWebImage_url_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtWebImage_url` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudWebImage_width_scale_ValueChanged()`** — L32548 — `private void nudWebImage_width_scale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudWebImage_width_scale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudWebImage_update_interval_ValueChanged()`** — L32553 — `private void nudWebImage_update_interval_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudWebImage_update_interval` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWebImage_fade_rx_CheckedChanged()`** — L32558 — `private void chkWebImage_fade_rx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWebImage_fade_rx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWebImage_fade_tx_CheckedChanged()`** — L32563 — `private void chkWebImage_fade_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWebImage_fade_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnWebImage_hamqsl_donate_Click()`** — L32568 — `private void btnWebImage_hamqsl_donate_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnWebImage_hamqsl_donate` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setup_comboWebImage_HamQsl()`** — L32573 — `private void setup_comboWebImage_HamQsl()`
  Called by: `.AfterConstructor()` (same file)
- **`.setup_comboWebImage_BsdWorld()`** — L32582 — `private void setup_comboWebImage_BsdWorld()`
  Called by: `.AfterConstructor()` (same file)
- **`.setup_comboWebImage_nasa()`** — L32591 — `private void setup_comboWebImage_nasa()`
  Called by: `.AfterConstructor()` (same file)
- **`.setup_comboWebImage_noaa()`** — L32600 — `private void setup_comboWebImage_noaa()`
  Called by: `.AfterConstructor()` (same file)
- **`.comboWebImage_HamQsl_SelectedIndexChanged()`** — L32610 — `private void comboWebImage_HamQsl_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboWebImage_HamQsl` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateWebImageState()`** — L32621 — `private void updateWebImageState(ImageFetcher.State state, bool checkSelected = false, string id = "")`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.SetWebImageState()` (same file)
- **`.SetWebImageState()`** — L32665 — `public void SetWebImageState(string id, ImageFetcher.State state)`
  Sets web image state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkContainerMinimises_CheckedChanged()`** — L32677 — `private void chkContainerMinimises_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkContainerMinimises` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMeterItemRotator_show_az_CheckedChanged()`** — L32687 — `private void radMeterItemRotator_show_az_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMeterItemRotator_show_az` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMeterItemRotator_show_ele_CheckedChanged()`** — L32697 — `private void radMeterItemRotator_show_ele_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMeterItemRotator_show_ele` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMeterItemRotator_show_both_CheckedChanged()`** — L32706 — `private void radMeterItemRotator_show_both_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMeterItemRotator_show_both` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItemRotator_padding_ValueChanged()`** — L32715 — `private void nudMeterItemRotator_padding_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemRotator_padding` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.fixBSDWorldUrls()`** — L32802 — `private string fixBSDWorldUrls(string url)`
  Called by: `.txtWebImage_url_TextChanged()` (same file)
- **`.comboWebImage_BsdWorld_SelectedIndexChanged()`** — L32822 — `private void comboWebImage_BsdWorld_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboWebImage_BsdWorld` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnWebImage_bsdworld_visit_Click()`** — L32835 — `private void btnWebImage_bsdworld_visit_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnWebImage_bsdworld_visit` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboWebImage_nasa_SelectedIndexChanged()`** — L32840 — `private void comboWebImage_nasa_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboWebImage_nasa` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboWebImage_noaa_SelectedIndexChanged()`** — L32852 — `private void comboWebImage_noaa_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboWebImage_noaa` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMeterItemRotatorSTOPcommand_TextChanged()`** — L32864 — `private void txtMeterItemRotatorSTOPcommand_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMeterItemRotatorSTOPcommand` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWebImage_bypass_cache_CheckedChanged()`** — L32869 — `private void chkWebImage_bypass_cache_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWebImage_bypass_cache` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMultiMeter_vfo_display_both_CheckedChanged()`** — L32874 — `private void radMultiMeter_vfo_display_both_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMultiMeter_vfo_display_both` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMultiMeter_vfo_display_vfoa_CheckedChanged()`** — L32880 — `private void radMultiMeter_vfo_display_vfoa_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMultiMeter_vfo_display_vfoa` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radMultiMeter_vfo_display_vfob_CheckedChanged()`** — L32886 — `private void radMultiMeter_vfo_display_vfob_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMultiMeter_vfo_display_vfob` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMultiMeter_auto_container_height_CheckedChanged()`** — L32892 — `private void chkMultiMeter_auto_container_height_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMultiMeter_auto_container_height` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudBandButtons_columns_ValueChanged()`** — L32902 — `private void nudBandButtons_columns_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudBandButtons_columns` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudBandButtons_border_ValueChanged()`** — L32907 — `private void nudBandButtons_border_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudBandButtons_border` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudBandButtons_margin_ValueChanged()`** — L32912 — `private void nudBandButtons_margin_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudBandButtons_margin` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudBandButtons_radius_ValueChanged()`** — L32917 — `private void nudBandButtons_radius_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudBandButtons_radius` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudBandButtons_height_ratio_ValueChanged()`** — L32922 — `private void nudBandButtons_height_ratio_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudBandButtons_height_ratio` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBandButtons_use_indicator_CheckedChanged()`** — L32927 — `private void chkBandButtons_use_indicator_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBandButtons_use_indicator` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudBandButtons_indicator_border_ValueChanged()`** — L32933 — `private void nudBandButtons_indicator_border_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudBandButtons_indicator_border` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateButtonIndicatorControls()`** — L32937 — `private void updateButtonIndicatorControls()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkBandButtons_use_indicator_CheckedChanged()` (same file)
- **`.clrbtnBandButtons_indicator_on_Changed()`** — L32946 — `private void clrbtnBandButtons_indicator_on_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnBandButtons_indicator_off_Changed()`** — L32950 — `private void clrbtnBandButtons_indicator_off_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnBandButtons_border_Changed()`** — L32955 — `private void clrbtnBandButtons_border_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnBandButtons_fill_Changed()`** — L32960 — `private void clrbtnBandButtons_fill_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnBandButtons_hover_Changed()`** — L32965 — `private void clrbtnBandButtons_hover_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkBandButtons_band_inactive_use_CheckedChanged()`** — L32970 — `private void chkBandButtons_band_inactive_use_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBandButtons_band_inactive_use` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBandButtons_fade_rx_CheckedChanged()`** — L32975 — `private void chkBandButtons_fade_rx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBandButtons_fade_rx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkBandButtons_fade_tx_CheckedChanged()`** — L32980 — `private void chkBandButtons_fade_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBandButtons_fade_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnBandButtons_font_Click()`** — L32985 — `private void btnBandButtons_font_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnBandButtons_font` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudBandButtons_indicator_style_ValueChanged()`** — L32999 — `private void nudBandButtons_indicator_style_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudBandButtons_indicator_style` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItemRotatorBeamWidth_alpha_ValueChanged()`** — L33004 — `private void nudMeterItemRotatorBeamWidth_alpha_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItemRotatorBeamWidth_alpha` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLegacyItems_band_CheckedChanged()`** — L33009 — `private void chkLegacyItems_band_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_band` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_mode_CheckedChanged()`** — L33015 — `private void chkLegacyItems_mode_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_mode` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_filter_CheckedChanged()`** — L33021 — `private void chkLegacyItems_filter_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_filter` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_expand_spectral_CheckedChanged()`** — L33027 — `private void chkLegacyItems_expand_spectral_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_expand_spectral` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudButtonBox_font_scale_ValueChanged()`** — L33033 — `private void nudButtonBox_font_scale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudButtonBox_font_scale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudButtonBox_font_x_shift_ValueChanged()`** — L33038 — `private void nudButtonBox_font_x_shift_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudButtonBox_font_x_shift` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudButtonBox_font_y_shift_ValueChanged()`** — L33043 — `private void nudButtonBox_font_y_shift_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudButtonBox_font_y_shift` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.notifyAntennaState()`** — L33048 — `private void notifyAntennaState()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnLedIndicator_PanelBackgroundTX_Changed()`** — L33063 — `private void clrbtnLedIndicator_PanelBackgroundTX_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemIndicator_Changed()`** — L33068 — `private void clrbtnMeterItemIndicator_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMeterItemSubIndicator_Changed()`** — L33073 — `private void clrbtnMeterItemSubIndicator_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkButtonBox_antenna_rx1_CheckedChanged()`** — L33078 — `private void chkButtonBox_antenna_rx1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_rx1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_rx2_CheckedChanged()`** — L33083 — `private void chkButtonBox_antenna_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_rx2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_rx3_CheckedChanged()`** — L33088 — `private void chkButtonBox_antenna_rx3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_rx3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_tx1_CheckedChanged()`** — L33093 — `private void chkButtonBox_antenna_tx1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_tx1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_tx2_CheckedChanged()`** — L33098 — `private void chkButtonBox_antenna_tx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_tx2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_tx3_CheckedChanged()`** — L33103 — `private void chkButtonBox_antenna_tx3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_tx3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_byp_CheckedChanged()`** — L33108 — `private void chkButtonBox_antenna_byp_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_byp` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_ext1_CheckedChanged()`** — L33113 — `private void chkButtonBox_antenna_ext1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_ext1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_xvtr_CheckedChanged()`** — L33118 — `private void chkButtonBox_antenna_xvtr_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_xvtr` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_antenna_rxtxant_CheckedChanged()`** — L33123 — `private void chkButtonBox_antenna_rxtxant_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_antenna_rxtxant` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getTotalColumnsNeededForAntennaButtons()`** — L33127 — `private int getTotalColumnsNeededForAntennaButtons()`
  Returns total columns needed for antenna buttons.
  Called by: `.updateMeterType()` (same file), `.updateItemSettingsControlsForSelected()` (same file)
- **`.chkMultiMeter_vfo_show_bandtext_CheckedChanged()`** — L33144 — `private void chkMultiMeter_vfo_show_bandtext_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMultiMeter_vfo_show_bandtext` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMultiMeter_vfo_show_bandtext_Changed()`** — L33150 — `private void clrbtnMultiMeter_vfo_show_bandtext_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateVfoShowBandtextColour()`** — L33155 — `private void updateVfoShowBandtextColour()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkMultiMeter_vfo_show_bandtext_CheckedChanged()` (same file)
- **`.chkLegacyItems_vfoa_CheckedChanged()`** — L33160 — `private void chkLegacyItems_vfoa_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_vfoa` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_vfob_CheckedChanged()`** — L33166 — `private void chkLegacyItems_vfob_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_vfob` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_expand_spectral_top_CheckedChanged()`** — L33172 — `private void chkLegacyItems_expand_spectral_top_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_expand_spectral_top` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_vfosync_CheckedChanged()`** — L33178 — `private void chkLegacyItems_vfosync_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_vfosync` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnMMVfoDisplayFrequency_small_Changed()`** — L33184 — `private void clrbtnMMVfoDisplayFrequency_small_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOCopyColourFromMainNumbers_Click()`** — L33189 — `private void btnVFOCopyColourFromMainNumbers_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnVFOCopyColourFromMainNumbers` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudHistory_vertical_ratio_ValueChanged()`** — L33194 — `private void nudHistory_vertical_ratio_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudHistory_vertical_ratio` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnHistory_background_Changed()`** — L33199 — `private void clrbtnHistory_background_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkHistory_fade_rx_CheckedChanged()`** — L33204 — `private void chkHistory_fade_rx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHistory_fade_rx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkHistory_fade_tx_CheckedChanged()`** — L33209 — `private void chkHistory_fade_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHistory_fade_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudHistory_update_ValueChanged()`** — L33214 — `private void nudHistory_update_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudHistory_update` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudHistory_keep_for_ValueChanged()`** — L33219 — `private void nudHistory_keep_for_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudHistory_keep_for` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboHistory_reading_0_SelectedIndexChanged()`** — L33224 — `private void comboHistory_reading_0_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboHistory_reading_0` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.initComboHistoryReadings0()`** — L33253 — `private void initComboHistoryReadings0()`
  Called by: `.AfterConstructor()` (same file)
- **`.nudHistory_axis0_min_ValueChanged()`** — L33277 — `private void nudHistory_axis0_min_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudHistory_axis0_min` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudHistory_axis0_max_ValueChanged()`** — L33283 — `private void nudHistory_axis0_max_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudHistory_axis0_max` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkHistory_auto_0_scale_CheckedChanged()`** — L33289 — `private void chkHistory_auto_0_scale_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHistory_auto_0_scale` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboHistory_reading_1_SelectedIndexChanged()`** — L33294 — `private void comboHistory_reading_1_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboHistory_reading_1` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkHistory_auto_1_scale_CheckedChanged()`** — L33302 — `private void chkHistory_auto_1_scale_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHistory_auto_1_scale` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudHistory_axis1_min_ValueChanged()`** — L33307 — `private void nudHistory_axis1_min_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudHistory_axis1_min` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudHistory_axis1_max_ValueChanged()`** — L33313 — `private void nudHistory_axis1_max_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudHistory_axis1_max` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDSPNR2trainT2_ValueChanged()`** — L33319 — `private void udDSPNR2trainT2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNR2trainT2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.udDSPNR2trainT2RX2_ValueChanged()`** — L33326 — `private void udDSPNR2trainT2RX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udDSPNR2trainT2RX2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkHistory_1_show_axis_CheckedChanged()`** — L33333 — `private void chkHistory_1_show_axis_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHistory_1_show_axis` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnHistory_copy_minmax_from_0_Click()`** — L33338 — `private void btnHistory_copy_minmax_from_0_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnHistory_copy_minmax_from_0` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_variable_history_Click()`** — L33344 — `private void btnMMIO_variable_history_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_variable_history` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMMIO_variable_2_history_Click()`** — L33349 — `private void btnMMIO_variable_2_history_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMMIO_variable_2_history` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnMultiMeter_vfo_lock_Changed()`** — L33354 — `private void clrbtnMultiMeter_vfo_lock_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnMultiMeter_vfo_sync_Changed()`** — L33359 — `private void clrbtnMultiMeter_vfo_sync_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucTunestepOptionsGrid_buttons_checkbox_changed()`** — L33364 — `private void ucTunestepOptionsGrid_buttons_checkbox_changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnHistory_colour_0_Changed()`** — L33369 — `private void clrbtnHistory_colour_0_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnHistory_colour_1_Changed()`** — L33374 — `private void clrbtnHistory_colour_1_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnHistory_lines_Changed()`** — L33379 — `private void clrbtnHistory_lines_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnHistory_time_Changed()`** — L33384 — `private void clrbtnHistory_time_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnGetMonitorHz_Click()`** — L33389 — `private void btnGetMonitorHz_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnGetMonitorHz` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLockContainer_CheckedChanged()`** — L33395 — `private void chkLockContainer_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLockContainer` checked state changes.
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (same file)
- **`.SetupCMAsio()`** — L33410 — `public void SetupCMAsio(bool portaudio_issue, bool cmasio_config_flag)`
  Setups cmasio.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radCMASIO_mic_CheckedChanged()`** — L33492 — `private void radCMASIO_mic_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radCMASIO_mic` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCMASIO_inpair_SelectedIndexChanged()`** — L33516 — `private void comboCMASIO_inpair_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCMASIO_inpair` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCMASIO_outpair_SelectedIndexChanged()`** — L33535 — `private void comboCMASIO_outpair_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCMASIO_outpair` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupMicSource()`** — L33551 — `private void setupMicSource(int base_in, bool enabled)`
  Called by: `.SetupCMAsio()` (same file), `.comboCMASIO_inpair_SelectedIndexChanged()` (same file)
- **`.setupInOutBaseChannels()`** — L33566 — `private void setupInOutBaseChannels(bool select_zero = false)`
  Called by: `.SetupCMAsio()` (same file), `.comboASIODevicesAvailable_SelectedIndexChanged()` (same file)
- **`.btnCMASIOActive_Click()`** — L33619 — `private void btnCMASIOActive_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCMASIOActive` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCMASIODisable_Click()`** — L33631 — `private void btnCMASIODisable_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCMASIODisable` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setCMasioControls()`** — L33636 — `private void setCMasioControls(bool enabled)`
  Sets cmasio controls.
  Called by: `.SetupCMAsio()` (same file), `.btnCMASIOActive_Click()` (same file)
- **`.nudAsioBlockNum_ValueChanged()`** — L33647 — `private void nudAsioBlockNum_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudAsioBlockNum` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboASIODevicesAvailable_SelectedIndexChanged()`** — L33655 — `private void comboASIODevicesAvailable_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboASIODevicesAvailable` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAsioLockMode_CheckedChanged()`** — L33674 — `private void chkAsioLockMode_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAsioLockMode` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCMAsioDefaultBlockNum_Click()`** — L33682 — `private void btnCMAsioDefaultBlockNum_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCMAsioDefaultBlockNum` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateCMAsioInfo()`** — L33686 — `private void updateCMAsioInfo()`
  Called by: `.radCMASIO_mic_CheckedChanged()` (same file), `.comboCMASIO_inpair_SelectedIndexChanged()` (same file), `.comboCMASIO_outpair_SelectedIndexChanged()` (same file), `.btnCMASIOActive_Click()` (same file), `.btnCMASIODisable_Click()` (same file), `.nudAsioBlockNum_ValueChanged()` (same file) — and 1 more
- **`.tmrLedValid_Tick()`** — L33691 — `private void tmrLedValid_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `tmrLedValid` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLed_notx_true_CheckedChanged()`** — L33697 — `private void chkLed_notx_true_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLed_notx_true` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLed_notx_false_CheckedChanged()`** — L33702 — `private void chkLed_notx_false_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLed_notx_false` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudLedIndicator_UpdateInterval_ValueChanged()`** — L33706 — `private void nudLedIndicator_UpdateInterval_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudLedIndicator_UpdateInterval` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tcCAT_SelectedIndexChanged()`** — L33711 — `private void tcCAT_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `tcCAT` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAutoATTTXPsOff_CheckedChanged()`** — L33718 — `private void chkAutoATTTXPsOff_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoATTTXPsOff` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkUndoAutoATTTx_CheckedChanged()`** — L33725 — `private void chkUndoAutoATTTx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUndoAutoATTTx` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAutoATTRx1_CheckedChanged()`** — L33731 — `private void chkAutoATTRx1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoATTRx1` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.setupAttRXControls()`** — L33738 — `private void setupAttRXControls(int rx)`
  Called by: `.comboRadioModel_SelectedIndexChanged()` (same file), `.chkAutoATTRx1_CheckedChanged()` (same file), `.chkAutoATTRx2_CheckedChanged()` (same file)
- **`.chkAutoAttUndoRX1_CheckedChanged()`** — L33760 — `private void chkAutoAttUndoRX1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoAttUndoRX1` checked state changes.
  Called by: `.chkAutoATTRx1_CheckedChanged()` (same file)
- **`.nudAutoAttHoldRX1_ValueChanged()`** — L33767 — `private void nudAutoAttHoldRX1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudAutoAttHoldRX1` value changes.
  Called by: `.chkAutoAttUndoRX1_CheckedChanged()` (same file)
- **`.chkAutoATTRx2_CheckedChanged()`** — L33773 — `private void chkAutoATTRx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoATTRx2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAutoAttUndoRX2_CheckedChanged()`** — L33781 — `private void chkAutoAttUndoRX2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAutoAttUndoRX2` checked state changes.
  Called by: `.chkAutoATTRx2_CheckedChanged()` (same file)
- **`.nudAutoAttHoldRX2_ValueChanged()`** — L33788 — `private void nudAutoAttHoldRX2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudAutoAttHoldRX2` value changes.
  Called by: `.chkAutoAttUndoRX2_CheckedChanged()` (same file)
- **`.lnkDiscordJoin_LinkClicked()`** — L33794 — `private void lnkDiscordJoin_LinkClicked(object sender, LinkLabelLinkClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkDiscordEnabled_CheckedChanged()`** — L33800 — `private void chkDiscordEnabled_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDiscordEnabled` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtDiscordCallsign_TextChanged()`** — L33820 — `private void txtDiscordCallsign_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtDiscordCallsign` text changes.
  Called by: `.chkDiscordEnabled_CheckedChanged()` (same file)
- **`.updateDiscordState()`** — L33835 — `private void updateDiscordState()`
  Called by: `.AfterConstructor()` (same file), `.chkDiscordEnabled_CheckedChanged()` (same file), `.txtDiscordCallsign_TextChanged()` (same file), `.OnDiscordConnect()` (same file), `.OnDiscordDisconnect()` (same file), `.OnDiscordReady()` (same file)
- **`.OnDiscordConnect()`** — L33860 — `private void OnDiscordConnect()`
  Handles/raises the discord connect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDiscordDisconnect()`** — L33864 — `private void OnDiscordDisconnect()`
  Handles/raises the discord disconnect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDiscordReady()`** — L33868 — `private void OnDiscordReady()`
  Handles/raises the discord ready event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.txtDiscordUniqueIDs_KeyPress()`** — L33873 — `private void txtDiscordUniqueIDs_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtDiscordUniqueIDs` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtDiscordFilter_KeyPress()`** — L33882 — `private void txtDiscordFilter_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtDiscordFilter` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtDiscordUniqueIDs_TextChanged()`** — L33891 — `private void txtDiscordUniqueIDs_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtDiscordUniqueIDs` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtDiscordFilter_TextChanged()`** — L33896 — `private void txtDiscordFilter_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtDiscordFilter` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnButonBox_click_Changed()`** — L33901 — `private void clrbtnButonBox_click_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnButonBox_fontcolour_Changed()`** — L33906 — `private void clrbtnButonBox_fontcolour_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.txtDiscordIgnore_TextChanged()`** — L33911 — `private void txtDiscordIgnore_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtDiscordIgnore` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbDiscordInfo_tag_Click()`** — L33916 — `private void pbDiscordInfo_tag_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbDiscordInfo_tag` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbDiscordInfo_ignore_Click()`** — L33921 — `private void pbDiscordInfo_ignore_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbDiscordInfo_ignore` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbDiscordInfo_filter_Click()`** — L33926 — `private void pbDiscordInfo_filter_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbDiscordInfo_filter` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtDiscordIgnore_KeyPress()`** — L33931 — `private void txtDiscordIgnore_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtDiscordIgnore` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilterDisplay_vertical_ratio_ValueChanged()`** — L33940 — `private void nudFilterDisplay_vertical_ratio_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilterDisplay_vertical_ratio` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnFilterDisplay_backcolour_Changed()`** — L33944 — `private void clrbtnFilterDisplay_backcolour_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkFilterDisplay_fadeonrx_CheckedChanged()`** — L33948 — `private void chkFilterDisplay_fadeonrx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilterDisplay_fadeonrx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilterDisplay_fadeontx_CheckedChanged()`** — L33952 — `private void chkFilterDisplay_fadeontx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilterDisplay_fadeontx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilterDisplay_show_limits_CheckedChanged()`** — L33956 — `private void chkFilterDisplay_show_limits_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilterDisplay_show_limits` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilterDisplay_fixed_zoom_CheckedChanged()`** — L33960 — `private void chkFilterDisplay_fixed_zoom_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilterDisplay_fixed_zoom` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilterDisplay_fixed_zoom_level_ValueChanged()`** — L33965 — `private void nudFilterDisplay_fixed_zoom_level_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilterDisplay_fixed_zoom_level` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilterDisplay_fixed_tx_zoom_CheckedChanged()`** — L33970 — `private void chkFilterDisplay_fixed_tx_zoom_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilterDisplay_fixed_tx_zoom` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilterDisplay_fixed_tx_zoom_level_ValueChanged()`** — L33976 — `private void nudFilterDisplay_fixed_tx_zoom_level_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilterDisplay_fixed_tx_zoom_level` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDiscordTimeStamp_CheckedChanged()`** — L33981 — `private void chkDiscordTimeStamp_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDiscordTimeStamp` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.updateDiscordTimeStampVisibilty()`** — L33986 — `private void updateDiscordTimeStampVisibilty()`
  Called by: `.chkDiscordTimeStamp_CheckedChanged()` (same file)
- **`.nudFilterItem_sidebands_scale_ValueChanged()`** — L34013 — `private void nudFilterItem_sidebands_scale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilterItem_sidebands_scale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilterItem_cw_scale_ValueChanged()`** — L34018 — `private void nudFilterItem_cw_scale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilterItem_cw_scale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilterItem_others_scale_ValueChanged()`** — L34023 — `private void nudFilterItem_others_scale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilterItem_others_scale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilterItem_panadaptor_CheckedChanged()`** — L34028 — `private void radFilterItem_panadaptor_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFilterItem_panadaptor` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilterItem_waterfall_CheckedChanged()`** — L34034 — `private void radFilterItem_waterfall_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFilterItem_waterfall` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilterItem_panafall_CheckedChanged()`** — L34040 — `private void radFilterItem_panafall_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFilterItem_panafall` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilterItem_none_CheckedChanged()`** — L34046 — `private void radFilterItem_none_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFilterItem_none` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilterItem_font_scale_ValueChanged()`** — L34052 — `private void nudFilterItem_font_scale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilterItem_font_scale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilter_fill_spec_CheckedChanged()`** — L34057 — `private void chkFilter_fill_spec_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilter_fill_spec` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnFilter_data_line_Changed()`** — L34062 — `private void clrbtnFilter_data_line_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_data_fill_Changed()`** — L34067 — `private void clrbtnFilter_data_fill_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboFilter_wf_palette_SelectedIndexChanged()`** — L34072 — `private void comboFilter_wf_palette_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFilter_wf_palette` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnFilter_wf_low_Changed()`** — L34077 — `private void clrbtnFilter_wf_low_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_text_Changed()`** — L34082 — `private void clrbtnFilter_text_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_number_highlight_Changed()`** — L34087 — `private void clrbtnFilter_number_highlight_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_edges_Changed()`** — L34092 — `private void clrbtnFilter_edges_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_edges_tx_Changed()`** — L34097 — `private void clrbtnFilter_edges_tx_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_edge_highlight_Changed()`** — L34102 — `private void clrbtnFilter_edge_highlight_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_meter_back_Changed()`** — L34107 — `private void clrbtnFilter_meter_back_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_notch_Changed()`** — L34112 — `private void clrbtnFilter_notch_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_notch_highlight_Changed()`** — L34117 — `private void clrbtnFilter_notch_highlight_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_extents_Changed()`** — L34122 — `private void clrbtnFilter_extents_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkFilter_sideband_mode_CheckedChanged()`** — L34127 — `private void chkFilter_sideband_mode_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilter_sideband_mode` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilter_waterfall_frame_update_ValueChanged()`** — L34132 — `private void nudFilter_waterfall_frame_update_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilter_waterfall_frame_update` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilter_grey_outsidepb_CheckedChanged()`** — L34137 — `private void chkFilter_grey_outsidepb_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilter_grey_outsidepb` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtFilter_sideband_frequencies_TextChanged()`** — L34142 — `private void txtFilter_sideband_frequencies_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtFilter_sideband_frequencies` text changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtFilter_cw_frequencies_TextChanged()`** — L34148 — `private void txtFilter_cw_frequencies_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtFilter_cw_frequencies` text changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.txtFilter_other_frequencies_TextChanged()`** — L34154 — `private void txtFilter_other_frequencies_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtFilter_other_frequencies` text changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnFilter_sideband_default_Click()`** — L34160 — `private void btnFilter_sideband_default_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFilter_sideband_default` is clicked.
  Called by: `.initFilterSnapFrequencies()` (same file)
- **`.btnFilter_cw_default_Click()`** — L34165 — `private void btnFilter_cw_default_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFilter_cw_default` is clicked.
  Called by: `.initFilterSnapFrequencies()` (same file)
- **`.btnFilter_others_default_Click()`** — L34170 — `private void btnFilter_others_default_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFilter_others_default` is clicked.
  Called by: `.initFilterSnapFrequencies()` (same file)
- **`.initFilterSnapFrequencies()`** — L34174 — `private void initFilterSnapFrequencies()`
  Called by: `.AfterConstructor()` (same file)
- **`.clrbtnFilter_snap_line_Changed()`** — L34181 — `private void clrbtnFilter_snap_line_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.defaultAlexSettings()`** — L34185 — `private void defaultAlexSettings()`
  Called by: `.AfterConstructor()` (same file)
- **`.clrbtnFilter_setting_on_Changed()`** — L34206 — `private void clrbtnFilter_setting_on_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnFilter_button_highlight_Changed()`** — L34211 — `private void clrbtnFilter_button_highlight_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.txtWebImage_4char_KeyPress()`** — L34216 — `private void txtWebImage_4char_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtWebImage_4char` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnFilter_4char_copy_Click()`** — L34221 — `private void btnFilter_4char_copy_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFilter_4char_copy` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWebImage_background_CheckedChanged()`** — L34226 — `private void chkWebImage_background_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWebImage_background` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudWebImage_background_time_ValueChanged()`** — L34232 — `private void nudWebImage_background_time_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudWebImage_background_time` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtWebImage_background_4char_TextChanged()`** — L34237 — `private void txtWebImage_background_4char_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtWebImage_background_4char` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateWebImageBackground()`** — L34241 — `private void updateWebImageBackground()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkWebImage_background_CheckedChanged()` (same file)
- **`.chkMaintainBackgroundAspectRatio_CheckedChanged()`** — L34251 — `private void chkMaintainBackgroundAspectRatio_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkMaintainBackgroundAspectRatio` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnWebImage_goto_next_Click()`** — L34256 — `private void btnWebImage_goto_next_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnWebImage_goto_next` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilter_characteristic_CheckedChanged()`** — L34280 — `private void chkFilter_characteristic_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilter_characteristic` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudFilter_lower_characteristic_ValueChanged()`** — L34286 — `private void nudFilter_lower_characteristic_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudFilter_lower_characteristic` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFilter_high_resolution_characteristics_CheckedChanged()`** — L34291 — `private void chkFilter_high_resolution_characteristics_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFilter_high_resolution_characteristics` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudDialDisplay_vertical_ratio_ValueChanged()`** — L34296 — `private void nudDialDisplay_vertical_ratio_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDialDisplay_vertical_ratio` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudDialDisplay_font_scale_ValueChanged()`** — L34301 — `private void nudDialDisplay_font_scale_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDialDisplay_font_scale` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnDialDisplay_background_Changed()`** — L34306 — `private void clrbtnDialDisplay_background_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkDialDisplay_fade_rx_CheckedChanged()`** — L34311 — `private void chkDialDisplay_fade_rx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDialDisplay_fade_rx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDialDisplay_fade_tx_CheckedChanged()`** — L34316 — `private void chkDialDisplay_fade_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDialDisplay_fade_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDialDisplay_alwaysshow_vfos_CheckedChanged()`** — L34321 — `private void chkDialDisplay_alwaysshow_vfos_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDialDisplay_alwaysshow_vfos` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDial_align_CheckedChanged()`** — L34326 — `private void chkDial_align_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDial_align` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudDial_increment_ValueChanged()`** — L34331 — `private void nudDial_increment_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDial_increment` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudDial_decrement_ValueChanged()`** — L34341 — `private void nudDial_decrement_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDial_decrement` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudDial_interval_ValueChanged()`** — L34351 — `private void nudDial_interval_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDial_interval` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnDial_colours_changed()`** — L34356 — `private void clrbtnDial_colours_changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudDial_max_increments_ValueChanged()`** — L34361 — `private void nudDial_max_increments_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDial_max_increments` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudDial_degrees_for_change_ValueChanged()`** — L34366 — `private void nudDial_degrees_for_change_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudDial_degrees_for_change` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnClearTCISpotsSWL_Click()`** — L34371 — `private void btnClearTCISpotsSWL_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClearTCISpotsSWL` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWindBackPowerSWR_CheckedChanged()`** — L34376 — `private void chkWindBackPowerSWR_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWindBackPowerSWR` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkForceATTwhenOutPowerChanges_decreased_CheckedChanged()`** — L34382 — `private void chkForceATTwhenOutPowerChanges_decreased_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkForceATTwhenOutPowerChanges_decreased` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowFrequencyNumbers_CheckedChanged()`** — L34388 — `private void chkShowFrequencyNumbers_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowFrequencyNumbers` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkShowLedMirror_CheckedChanged()`** — L34394 — `private void chkShowLedMirror_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowLedMirror` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.showLedMirror()`** — L34398 — `private void showLedMirror()`
  Called by: `.ForceAllEvents()` (same file), `.UpdateDDCTab()` (same file)
- **`.btnDefaultGradient_waterfall_Click()`** — L34424 — `private void btnDefaultGradient_waterfall_Click(object sender, EventArgs e)`
  LG waterfall
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDeleteColourGripper_waterfall_Click()`** — L34430 — `private void btnDeleteColourGripper_waterfall_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDeleteColourGripper_waterfall` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnClearColourGrippers_waterfall_Click()`** — L34436 — `private void btnClearColourGrippers_waterfall_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClearColourGrippers_waterfall` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnLoadGradient_waterfall_Click()`** — L34441 — `private async void btnLoadGradient_waterfall_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLoadGradient_waterfall` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSaveGradient_waterfall_Click()`** — L34462 — `private async void btnSaveGradient_waterfall_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSaveGradient_waterfall` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lgLinearGradient_waterfall_Changed()`** — L34477 — `private void lgLinearGradient_waterfall_Changed(object sender, EventArgs e)`
  Called by: `.lgLinearGradient_waterfall_GripperDBMChanged()` (same file)
- **`.WaterfallRXGradient()`** — L34485 — `public Color[] WaterfallRXGradient()`
  Called by: `.lgLinearGradient_waterfall_Changed()` (same file)
- **`.WaterfallTXGradient()`** — L34495 — `public Color[] WaterfallTXGradient()`
  Called by: `.lgLinearGradientTX_waterfall_Changed()` (same file)
- **`.lgLinearGradient_waterfall_GripperDBMChanged()`** — L34505 — `private void lgLinearGradient_waterfall_GripperDBMChanged(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradient_waterfall_GripperMouseEnter()`** — L34513 — `private void lgLinearGradient_waterfall_GripperMouseEnter(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradient_waterfall_GripperMouseLeave()`** — L34518 — `private void lgLinearGradient_waterfall_GripperMouseLeave(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradient_waterfall_GripperSelected()`** — L34523 — `private void lgLinearGradient_waterfall_GripperSelected(object sender, ColourEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnGripperColour_waterfall_Changed()`** — L34529 — `private void clrbtnGripperColour_waterfall_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.defaultLinearGradients()`** — L34535 — `private void defaultLinearGradients(bool pana, bool waterfall, bool tx)`
  Called by: `.AfterConstructor()` (same file), `.InitDisplayTab()` (same file)
- **`.btnWaterfallToClipboard_Click()`** — L34568 — `private void btnWaterfallToClipboard_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnWaterfallToClipboard` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lgLinearGradientTX_Changed()`** — L34573 — `private void lgLinearGradientTX_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradientTX_GripperDBMChanged()`** — L34578 — `private void lgLinearGradientTX_GripperDBMChanged(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradientTX_GripperMouseEnter()`** — L34585 — `private void lgLinearGradientTX_GripperMouseEnter(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradientTX_GripperMouseLeave()`** — L34590 — `private void lgLinearGradientTX_GripperMouseLeave(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradientTX_GripperSelected()`** — L34595 — `private void lgLinearGradientTX_GripperSelected(object sender, ColourEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnGripperColour_tx_Changed()`** — L34601 — `private void clrbtnGripperColour_tx_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkDataLineGradient_tx_CheckedChanged()`** — L34608 — `private void chkDataLineGradient_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDataLineGradient_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDefaultGradient_tx_Click()`** — L34615 — `private void btnDefaultGradient_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDefaultGradient_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDeleteColourGripper_tx_Click()`** — L34621 — `private void btnDeleteColourGripper_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDeleteColourGripper_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnClearColourGrippers_tx_Click()`** — L34627 — `private void btnClearColourGrippers_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClearColourGrippers_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnLoadGradient_tx_Click()`** — L34632 — `private async void btnLoadGradient_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLoadGradient_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSaveGradient_tx_Click()`** — L34653 — `private async void btnSaveGradient_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSaveGradient_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnDataFill_tx_Changed()`** — L34668 — `private void clrbtnDataFill_tx_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file), `.tbDataFillAlpha_tx_Scroll()` (same file)
- **`.tbDataFillAlpha_tx_Scroll()`** — L34674 — `private void tbDataFillAlpha_tx_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDataFillAlpha_tx` is scrolled.
  Called by: `.AfterConstructor()` (same file)
- **`.rebuildTXLGBrushes()`** — L34680 — `private void rebuildTXLGBrushes()`
  Called by: `.clrbtnTXDataLine_Changed()` (same file), `.lgLinearGradientTX_Changed()` (same file), `.lgLinearGradientTX_GripperDBMChanged()` (same file)
- **`.tbDataLineAlpha_tx_Scroll()`** — L34685 — `private void tbDataLineAlpha_tx_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDataLineAlpha_tx` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPanadpatorGradient_tx_CheckedChanged()`** — L34691 — `private void chkPanadpatorGradient_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPanadpatorGradient_tx` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.lgLinearGradientTX_waterfall_Changed()`** — L34717 — `private void lgLinearGradientTX_waterfall_Changed(object sender, EventArgs e)`
  Called by: `.lgLinearGradientTX_waterfall_GripperDBMChanged()` (same file)
- **`.lgLinearGradientTX_waterfall_GripperDBMChanged()`** — L34724 — `private void lgLinearGradientTX_waterfall_GripperDBMChanged(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradientTX_waterfall_GripperMouseEnter()`** — L34732 — `private void lgLinearGradientTX_waterfall_GripperMouseEnter(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradientTX_waterfall_GripperMouseLeave()`** — L34737 — `private void lgLinearGradientTX_waterfall_GripperMouseLeave(object sender, GripperEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lgLinearGradientTX_waterfall_GripperSelected()`** — L34742 — `private void lgLinearGradientTX_waterfall_GripperSelected(object sender, ColourEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clrbtnGripperColour_waterfall_tx_Changed()`** — L34748 — `private void clrbtnGripperColour_waterfall_tx_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnDefaultGradient_waterfall_tx_Click()`** — L34755 — `private void btnDefaultGradient_waterfall_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDefaultGradient_waterfall_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDeleteColourGripper_waterfall_tx_Click()`** — L34761 — `private void btnDeleteColourGripper_waterfall_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDeleteColourGripper_waterfall_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnClearColourGrippers_waterfall_tx_Click()`** — L34767 — `private void btnClearColourGrippers_waterfall_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClearColourGrippers_waterfall_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnLoadGradient_waterfall_tx_Click()`** — L34772 — `private async void btnLoadGradient_waterfall_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLoadGradient_waterfall_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSaveGradient_waterfall_tx_Click()`** — L34793 — `private async void btnSaveGradient_waterfall_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSaveGradient_waterfall_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboColorPalette_tx_SelectedIndexChanged()`** — L34808 — `private void comboColorPalette_tx_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboColorPalette_tx` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnWaterfallLow_tx_Changed()`** — L34852 — `private void clrbtnWaterfallLow_tx_Changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.ucGradientDefault_rx_pana_SetGradient()`** — L34858 — `private void ucGradientDefault_rx_pana_SetGradient(bool arg1, string arg2)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucGradientDefault_rx_waterfall_SetGradient()`** — L34863 — `private void ucGradientDefault_rx_waterfall_SetGradient(bool arg1, string arg2)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucGradientDefault_tx_pana_SetGradient()`** — L34868 — `private void ucGradientDefault_tx_pana_SetGradient(bool arg1, string arg2)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucGradientDefault_tx_waterfall_SetGradient()`** — L34873 — `private void ucGradientDefault_tx_waterfall_SetGradient(bool arg1, string arg2)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radMeterItemSettings_CheckedChanged()`** — L34878 — `private void radMeterItemSettings_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMeterItemSettings` checked state changes.
  Called by: `.updateItemSettingsControlsForSelected()` (same file)
- **`.radMeterItemSettings_custom_CheckedChanged()`** — L34884 — `private void radMeterItemSettings_custom_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radMeterItemSettings_custom` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupCustomItemSettings()`** — L34890 — `private void setupCustomItemSettings(bool show_settings)`
  Called by: `.radMeterItemSettings_CheckedChanged()` (same file), `.radMeterItemSettings_custom_CheckedChanged()` (same file)
- **`.nudMeterItem_custom_min_ValueChanged()`** — L34907 — `private void nudMeterItem_custom_min_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItem_custom_min` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItem_custom_max_ValueChanged()`** — L34928 — `private void nudMeterItem_custom_max_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItem_custom_max` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudMeterItem_custom_high_ValueChanged()`** — L34949 — `private void nudMeterItem_custom_high_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudMeterItem_custom_high` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMeterItem_custom_units_TextChanged()`** — L34970 — `private void txtMeterItem_custom_units_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMeterItem_custom_units` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMeterItem_custom_title_TextChanged()`** — L34975 — `private void txtMeterItem_custom_title_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMeterItem_custom_title` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnFPSProfile_Click()`** — L34981 — `private void btnFPSProfile_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFPSProfile` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ValidFpsProfile()`** — L35278 — `public bool ValidFpsProfile()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkWDSP_cache_impulse_CheckedChanged()`** — L35393 — `private void chkWDSP_cache_impulse_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWDSP_cache_impulse` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkWDSP_save_restore_cache_impulse_CheckedChanged()`** — L35399 — `private void chkWDSP_save_restore_cache_impulse_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWDSP_save_restore_cache_impulse` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnTwoToneF_defaults_Click()`** — L35405 — `private void btnTwoToneF_defaults_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTwoToneF_defaults` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnTwoToneF_stealth_Click()`** — L35411 — `private void btnTwoToneF_stealth_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTwoToneF_stealth` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudPulsedTune_window_ValueChanged()`** — L35417 — `private void nudPulsedTune_window_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPulsedTune_window` value changes.
  Called by: `.chkPulsedTune_CheckedChanged()` (same file)
- **`.nudPulsedTune_percent_ValueChanged()`** — L35425 — `private void nudPulsedTune_percent_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPulsedTune_percent` value changes.
  Called by: `.chkPulsedTune_CheckedChanged()` (same file)
- **`.chkPulsedTune_CheckedChanged()`** — L35433 — `private void chkPulsedTune_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPulsedTune` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudPulsedTune_ramp_ValueChanged()`** — L35447 — `private void nudPulsedTune_ramp_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPulsedTune_ramp` value changes.
  Called by: `.chkPulsedTune_CheckedChanged()` (same file)
- **`.updateTunePulseInfo()`** — L35455 — `private void updateTunePulseInfo()`
  Called by: `.nudPulsedTune_window_ValueChanged()` (same file), `.nudPulsedTune_percent_ValueChanged()` (same file), `.nudPulsedTune_ramp_ValueChanged()` (same file)
- **`.chkPreventSleep_CheckedChanged()`** — L35478 — `private void chkPreventSleep_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPreventSleep` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkPreventScreenSaver_CheckedChanged()`** — L35488 — `private void chkPreventScreenSaver_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPreventScreenSaver` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radContainer_rx1_data_CheckedChanged()`** — L35498 — `private void radContainer_rx1_data_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radContainer_rx1_data` checked state changes.
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (same file)
- **`.radContainer_rx2_data_CheckedChanged()`** — L35512 — `private void radContainer_rx2_data_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radContainer_rx2_data` checked state changes.
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (same file)
- **`.chkContainer_hidewhennotused_CheckedChanged()`** — L35526 — `private void chkContainer_hidewhennotused_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkContainer_hidewhennotused` checked state changes.
  Called by: `.radContainer_rx1_data_CheckedChanged()` (same file), `.radContainer_rx2_data_CheckedChanged()` (same file)
- **`.chkPulsed_TwoTone_CheckedChanged()`** — L35536 — `private void chkPulsed_TwoTone_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPulsed_TwoTone` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudPulsed_TwoTone_window_ValueChanged()`** — L35547 — `private void nudPulsed_TwoTone_window_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPulsed_TwoTone_window` value changes.
  Called by: `.chkPulsed_TwoTone_CheckedChanged()` (same file)
- **`.nudPulsed_TwoTone_percent_ValueChanged()`** — L35554 — `private void nudPulsed_TwoTone_percent_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPulsed_TwoTone_percent` value changes.
  Called by: `.chkPulsed_TwoTone_CheckedChanged()` (same file)
- **`.nudPulsed_TwoTone_ramp_ValueChanged()`** — L35561 — `private void nudPulsed_TwoTone_ramp_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudPulsed_TwoTone_ramp` value changes.
  Called by: `.chkPulsed_TwoTone_CheckedChanged()` (same file)
- **`.updateTwoTonePulseInfo()`** — L35568 — `private void updateTwoTonePulseInfo()`
  Called by: `.nudPulsed_TwoTone_window_ValueChanged()` (same file), `.nudPulsed_TwoTone_percent_ValueChanged()` (same file), `.nudPulsed_TwoTone_ramp_ValueChanged()` (same file)
- **`.setupTwoTonePulse()`** — L35590 — `private void setupTwoTonePulse()`
  Called by: `.chkTestIMD_CheckedChanged()` (same file), `.nudPulsed_TwoTone_window_ValueChanged()` (same file), `.nudPulsed_TwoTone_percent_ValueChanged()` (same file), `.nudPulsed_TwoTone_ramp_ValueChanged()` (same file)
- **`.ucMeterItemSignalType_SignalTypeChanged()`** — L35601 — `private void ucMeterItemSignalType_SignalTypeChanged(object sender, ucSignalSelect.SignalTypeChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkOverrideSpotFlashColour_CheckedChanged()`** — L35606 — `private void chkOverrideSpotFlashColour_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOverrideSpotFlashColour` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.clrbtnSpotFlashColour_Changed()`** — L35614 — `private void clrbtnSpotFlashColour_Changed(object sender, EventArgs e)`
  Called by: `.chkOverrideSpotFlashColour_CheckedChanged()` (same file)
- **`.nudNR4_red_rx1_ValueChanged()`** — L35649 — `private void nudNR4_red_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_red_rx1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_smo_rx1_ValueChanged()`** — L35656 — `private void nudNR4_smo_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_smo_rx1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_whi_rx1_ValueChanged()`** — L35663 — `private void nudNR4_whi_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_whi_rx1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_res_rx1_ValueChanged()`** — L35670 — `private void nudNR4_res_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_res_rx1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_snr_rx1_ValueChanged()`** — L35677 — `private void nudNR4_snr_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_snr_rx1` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.setupNR4algorithm()`** — L35683 — `private void setupNR4algorithm()`
  Called by: `.ForceAllEvents()` (same file)
- **`.radNR4_algo1_CheckedChanged()`** — L35692 — `private void radNR4_algo1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radNR4_algo1` checked state changes.
  Called by: `.setupNR4algorithm()` (same file)
- **`.radNR4_algo2_CheckedChanged()`** — L35700 — `private void radNR4_algo2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radNR4_algo2` checked state changes.
  Called by: `.setupNR4algorithm()` (same file)
- **`.radNR4_algo3_CheckedChanged()`** — L35708 — `private void radNR4_algo3_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radNR4_algo3` checked state changes.
  Called by: `.setupNR4algorithm()` (same file)
- **`.nudNR4_red_rx2_ValueChanged()`** — L35716 — `private void nudNR4_red_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_red_rx2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_smo_rx2_ValueChanged()`** — L35723 — `private void nudNR4_smo_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_smo_rx2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_whi_rx2_ValueChanged()`** — L35730 — `private void nudNR4_whi_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_whi_rx2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_res_rx2_ValueChanged()`** — L35737 — `private void nudNR4_res_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_res_rx2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.nudNR4_snr_rx2_ValueChanged()`** — L35744 — `private void nudNR4_snr_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR4_snr_rx2` value changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radNR4_algo1_rx2_CheckedChanged()`** — L35751 — `private void radNR4_algo1_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radNR4_algo1_rx2` checked state changes.
  Called by: `.setupNR4algorithm()` (same file)
- **`.radNR4_algo2_rx2_CheckedChanged()`** — L35759 — `private void radNR4_algo2_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radNR4_algo2_rx2` checked state changes.
  Called by: `.setupNR4algorithm()` (same file)
- **`.radNR4_algo3_rx2_CheckedChanged()`** — L35767 — `private void radNR4_algo3_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radNR4_algo3_rx2` checked state changes.
  Called by: `.setupNR4algorithm()` (same file)
- **`.btnNR3_model_default_Click()`** — L35785 — `private void btnNR3_model_default_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnNR3_model_default` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnNR3_model_load_Click()`** — L35790 — `private void btnNR3_model_load_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnNR3_model_load` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setNR3Model()`** — L35801 — `private void setNR3Model()`
  Sets nr3 model.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.validateRnnoiseModel()`** — L35840 — `private bool validateRnnoiseModel(string path)`
  Called by: `.setNR3Model()` (same file)
- **`.setupNR2PostProcessing()`** — L35873 — `private void setupNR2PostProcessing(int rx)`
  Called by: `.ForceAllEvents()` (same file), `.chkNR2PostProc_enable_rx1_CheckedChanged()` (same file), `.nudNR2PostProc_level_rx1_ValueChanged()` (same file), `.nudNR2PostProc_factor_rx1_ValueChanged()` (same file), `.nudNR2PostProc_rate_rx1_ValueChanged()` (same file), `.nudNR2PostProc_taper_rx1_ValueChanged()` (same file) — and 5 more
- **`.chkNR2PostProc_enable_rx1_CheckedChanged()`** — L35937 — `private void chkNR2PostProc_enable_rx1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNR2PostProc_enable_rx1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_level_rx1_ValueChanged()`** — L35943 — `private void nudNR2PostProc_level_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_level_rx1` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_factor_rx1_ValueChanged()`** — L35949 — `private void nudNR2PostProc_factor_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_factor_rx1` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_rate_rx1_ValueChanged()`** — L35955 — `private void nudNR2PostProc_rate_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_rate_rx1` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_taper_rx1_ValueChanged()`** — L35961 — `private void nudNR2PostProc_taper_rx1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_taper_rx1` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNR2PostProc_enable_rx2_CheckedChanged()`** — L35967 — `private void chkNR2PostProc_enable_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNR2PostProc_enable_rx2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_level_rx2_ValueChanged()`** — L35973 — `private void nudNR2PostProc_level_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_level_rx2` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_factor_rx2_ValueChanged()`** — L35979 — `private void nudNR2PostProc_factor_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_factor_rx2` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_rate_rx2_ValueChanged()`** — L35985 — `private void nudNR2PostProc_rate_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_rate_rx2` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudNR2PostProc_taper_rx2_ValueChanged()`** — L35991 — `private void nudNR2PostProc_taper_rx2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudNR2PostProc_taper_rx2` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnFormLocationHelper_Click()`** — L35998 — `private void btnFormLocationHelper_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFormLocationHelper` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnLedIndicator_4char_copy_Click()`** — L36005 — `private void btnLedIndicator_4char_copy_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLedIndicator_4char_copy` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtLedIndicator_4char_KeyPress()`** — L36010 — `private void txtLedIndicator_4char_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtLedIndicator_4char` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateTextOverlayLedIndicator()`** — L36015 — `private void updateTextOverlayLedIndicator()`
  Called by: `.updateItemSettingsControlsForSelected()` (same file), `.chkTextOverlay_rx_on_led_CheckedChanged()` (same file), `.chkTextOverlay_tx_on_led_CheckedChanged()` (same file)
- **`.chkTextOverlay_rx_on_led_CheckedChanged()`** — L36021 — `private void chkTextOverlay_rx_on_led_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTextOverlay_rx_on_led` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtTextOverlay_rx_on_led_4char_TextChanged()`** — L36027 — `private void txtTextOverlay_rx_on_led_4char_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtTextOverlay_rx_on_led_4char` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTextOverlay_tx_on_led_CheckedChanged()`** — L36032 — `private void chkTextOverlay_tx_on_led_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTextOverlay_tx_on_led` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtTextOverlay_tx_on_led_4char_TextChanged()`** — L36038 — `private void txtTextOverlay_tx_on_led_4char_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtTextOverlay_tx_on_led_4char` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucOtherButtonsOptionsGrid_buttons_CheckboxChanged()`** — L36043 — `private void ucOtherButtonsOptionsGrid_buttons_CheckboxChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkButtonBox_fix_text_size_CheckedChanged()`** — L36048 — `private void chkButtonBox_fix_text_size_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_fix_text_size` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButtonBox_use_icons_CheckedChanged()`** — L36053 — `private void chkButtonBox_use_icons_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButtonBox_use_icons` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLegacyItems_power_rx2_CheckedChanged()`** — L36058 — `private void chkLegacyItems_power_rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_power_rx2` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_mon_tun_CheckedChanged()`** — L36064 — `private void chkLegacyItems_mon_tun_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_mon_tun` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_hide_split_rit_CheckedChanged()`** — L36070 — `private void chkLegacyItems_hide_split_rit_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_hide_split_rit` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_hide_noise_mnf_CheckedChanged()`** — L36076 — `private void chkLegacyItems_hide_noise_mnf_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_hide_noise_mnf` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_hide_mic_comp_CheckedChanged()`** — L36082 — `private void chkLegacyItems_hide_mic_comp_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_hide_mic_comp` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkLegacyItems_hide_avg_peak_CheckedChanged()`** — L36088 — `private void chkLegacyItems_hide_avg_peak_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyItems_hide_avg_peak` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnOtherButtons_reset_layout_Click()`** — L36096 — `private void btnOtherButtons_reset_layout_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnOtherButtons_reset_layout` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnWaveRecord_reset_layout_Click()`** — L36102 — `private void btnWaveRecord_reset_layout_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnWaveRecord_reset_layout` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnContainer_save_Click()`** — L36109 — `private void btnContainer_save_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnContainer_save` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnContainer_load_Click()`** — L36134 — `private void btnContainer_load_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnContainer_load` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnContainer_dupe_Click()`** — L36213 — `private void btnContainer_dupe_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnContainer_dupe` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucOtherButtonsOptionsGrid_buttons_MacroSetupClicked()`** — L36244 — `private void ucOtherButtonsOptionsGrid_buttons_MacroSetupClicked(object sender, ucOtherButtonsOptionsGrid.MacroButtonEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkLed_process_when_hidden_CheckedChanged()`** — L36276 — `private void chkLed_process_when_hidden_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLed_process_when_hidden` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecoverContainer_Click()`** — L36281 — `private void btnRecoverContainer_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecoverContainer` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnTextOverlayVarPicker_Click()`** — L36294 — `private void btnTextOverlayVarPicker_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTextOverlayVarPicker` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.showVarPickerForClipboard()`** — L36298 — `private string showVarPickerForClipboard()`
  Called by: `.btnTextOverlayVarPicker_Click()` (same file), `.btnLedIndicatorVarPicker_Click()` (same file)
- **`.btnLedIndicatorVarPicker_Click()`** — L36312 — `private void btnLedIndicatorVarPicker_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLedIndicatorVarPicker` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOsync_settings_changed()`** — L36318 — `private void chkVFOsync_settings_changed(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.txtN1MM_RXn_ID_TextChanged()`** — L36357 — `private void txtN1MM_RXn_ID_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtN1MM_RXn_ID` text changes.
  Called by: `.chkN1MMEnableRX1_CheckedChanged()` (same file), `.chkN1MMEnableRX2_CheckedChanged()` (same file)
- **`.chkShowStartupLog_CheckedChanged()`** — L36419 — `private void chkShowStartupLog_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowStartupLog` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateShowStartupLogCheckBox()`** — L36434 — `private void updateShowStartupLogCheckBox()`
  Called by: `.AfterConstructor()` (same file)
- **`.radTCI_spot_force_CheckedChanged()`** — L36448 — `private void radTCI_spot_force_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radTCI_spot_force` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAPF_type_SelectedIndexChanged()`** — L36546 — `private void comboAPF_type_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboAPF_type` selection changes.
  Called by: `.ForceAllEvents()` (same file), `.SetAPFType()` (same file)
- **`.CycleAPFType()`** — L36583 — `public void CycleAPFType()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAPFType()`** — L36609 — `public void SetAPFType(int type)`
  Sets apftype.
  Called by: `.radDSPRX1APFControls_CheckedChanged()` (same file), `.radDSPRX1subAPFControls_CheckedChanged()` (same file), `.radDSPRX2APFControls_CheckedChanged()` (same file), `.CycleAPFType()` (same file)
- **`.chkIgnoreATTOffset_CheckedChanged()`** — L36649 — `private void chkIgnoreATTOffset_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkIgnoreATTOffset` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnShowBandwidth_Click()`** — L36655 — `private void btnShowBandwidth_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShowBandwidth` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CloseBandwidthForm()`** — L36660 — `public void CloseBandwidthForm()`
  Closes bandwidth form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkNR3_RNNoiseFixedGain_CheckedChanged()`** — L36669 — `private void chkNR3_RNNoiseFixedGain_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNR3_RNNoiseFixedGain` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.pbCMasio_InOut_Info_Click()`** — L36679 — `private void pbCMasio_InOut_Info_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pbCMasio_InOut_Info` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getRadioDiscoveryOptions()`** — L36685 — `private RadioDiscoveryOptions getRadioDiscoveryOptions(bool nics_only = false)`
  Returns radio discovery options.
  Called by: `.rebuildNicCombo()` (same file), `.tryDiscoverRadios()` (same file)
- **`.setupNeworking()`** — L36734 — `private void setupNeworking()`
  Called by: `.AfterConstructor()` (same file)
- **`.rebuildNicCombo()`** — L36743 — `private void rebuildNicCombo()`
  Called by: `.setupNeworking()` (same file), `.btnRefreshNics_Click()` (same file), `.chkLimitInterfacesToEthernetWifi_CheckedChanged()` (same file)
- **`.applyAnyOrSpecificRadio()`** — L36777 — `private bool applyAnyOrSpecificRadio(RadioDiscoveryOptions options)`
  Called by: `.getRadioDiscoveryOptions()` (same file)
- **`.tryParseIpPort()`** — L36808 — `private bool tryParseIpPort(string text, int defaultPort, out IPAddress ip, out int port)`
  Called by: `.applyAnyOrSpecificRadio()` (same file), `.btnAddCustomRadio_Click()` (same file)
- **`.tryResolveHostToIPv4()`** — L36853 — `private bool tryResolveHostToIPv4(string host, int p, out IPAddress ip, out int port)`
  Called by: `.tryParseIpPort()` (same file)
- **`.looksLikeHostname()`** — L36898 — `private bool looksLikeHostname(string host)`
  Called by: `.tryResolveHostToIPv4()` (same file)
- **`.tryDiscoverRadios()`** — L36947 — `private bool tryDiscoverRadios(out List<NicRadioScanResult> discovered, bool showNoRadiosMessage)`
  Called by: `.btnDiscoverRadios_Click()` (same file), `.ScanForFirstFoundRadio()` (same file)
- **`.btnDiscoverRadios_Click()`** — L37046 — `private void btnDiscoverRadios_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDiscoverRadios` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScanForFirstFoundRadio()`** — L37074 — `public bool ScanForFirstFoundRadio()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnRefreshNics_Click()`** — L37097 — `private void btnRefreshNics_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRefreshNics` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radViaNics_CheckedChanged()`** — L37102 — `private void radViaNics_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radViaNics` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.radAnyOrSpecificRadio_CheckedChanged()`** — L37118 — `private void radAnyOrSpecificRadio_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radAnyOrSpecificRadio` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.ucRadioList_Radios_RadioListChanged()`** — L37157 — `private void ucRadioList_Radios_RadioListChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucRadioList_Radios_SelectedRadioChanged()`** — L37162 — `private void ucRadioList_Radios_SelectedRadioChanged(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.OnPowerChangeHandler()`** — L37184 — `private void OnPowerChangeHandler(bool oldPower, bool newPower)`
  Handles/raises the power change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radDefaultOrRandomListenPort_CheckedChanged()`** — L37190 — `private void radDefaultOrRandomListenPort_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radDefaultOrRandomListenPort` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.chkAdvancedNetworkingSettings_CheckedChanged()`** — L37196 — `private void chkAdvancedNetworkingSettings_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAdvancedNetworkingSettings` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnAddCustomRadio_Click()`** — L37214 — `private void btnAddCustomRadio_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAddCustomRadio` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLimitInterfacesToEthernetWifi_CheckedChanged()`** — L37276 — `private void chkLimitInterfacesToEthernetWifi_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLimitInterfacesToEthernetWifi` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.initRecordingPlaybackAudio()`** — L37284 — `private void initRecordingPlaybackAudio()`
  Called by: `.AfterConstructor()` (same file)
- **`.initPCAudioDevicesComobs()`** — L37290 — `private void initPCAudioDevicesComobs()`
  Called by: `.initRecordingPlaybackAudio()` (same file), `.btnRecording_refreshDevices_Click()` (same file)
- **`.radRecordingBits_CheckedChanged()`** — L37364 — `private void radRecordingBits_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRecordingBits` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_playbackMox_CheckedChanged()`** — L37388 — `private void chkRecording_playbackMox_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_playbackMox` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_generateMP3s_CheckedChanged()`** — L37393 — `private void chkRecording_generateMP3s_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_generateMP3s` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRecording_storage_CheckedChanged()`** — L37398 — `private void radRecording_storage_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRecording_storage` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.txtRecording_customFolder_TextChanged()` (same file)
- **`.txtRecording_customFolder_TextChanged()`** — L37425 — `private void txtRecording_customFolder_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtRecording_customFolder` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_selectCustomFolder_Click()`** — L37431 — `private void btnRecording_selectCustomFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_selectCustomFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.OpenWaveRecordFolder()`** — L37450 — `public void OpenWaveRecordFolder()`
  Opens wave record folder.
  Called by: `.btnRecording_openWaverecordFolder_Click()` (same file)
- **`.btnRecording_openWaverecordFolder_Click()`** — L37472 — `private void btnRecording_openWaverecordFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_openWaverecordFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_openQuickFolder_Click()`** — L37477 — `private void btnRecording_openQuickFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_openQuickFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_openRecordingsFolder_Click()`** — L37500 — `private void btnRecording_openRecordingsFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_openRecordingsFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboRecording_samplerate_SelectedIndexChanged()`** — L37513 — `private void comboRecording_samplerate_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboRecording_samplerate` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_dither_CheckedChanged()`** — L37550 — `private void chkRecording_dither_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_dither` checked state changes.
  Called by: `.nudRecording_dither_ValueChanged()` (same file)
- **`.nudRecording_dither_ValueChanged()`** — L37557 — `private void nudRecording_dither_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_dither` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRecording_RXing_CheckedChanged()`** — L37562 — `private void radRecording_RXing_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRecording_RXing` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRecording_TXing_CheckedChanged()`** — L37574 — `private void radRecording_TXing_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRecording_TXing` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboPCAudioDevices_IN_SelectedIndexChanged()`** — L37586 — `private void comboPCAudioDevices_IN_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboPCAudioDevices_IN` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboPCAudioDevices_OUT_SelectedIndexChanged()`** — L37594 — `private void comboPCAudioDevices_OUT_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboPCAudioDevices_OUT` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_gainInput_ValueChanged()`** — L37602 — `private void nudRecording_gainInput_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_gainInput` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_gainOutput_ValueChanged()`** — L37607 — `private void nudRecording_gainOutput_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_gainOutput` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_refreshDevices_Click()`** — L37612 — `private void btnRecording_refreshDevices_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_refreshDevices` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_txGain_ValueChanged()`** — L37616 — `private void nudRecording_txGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_txGain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_monoPlaybackGain_ValueChanged()`** — L37620 — `private void nudRecording_monoPlaybackGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_monoPlaybackGain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_disable_during_playback()`** — L37624 — `private void chkRecording_disable_during_playback(object sender, EventArgs e)`
  Called by: `.ForceAllEvents()` (same file)
- **`.chkBypassVACPlayingRecording_CheckedChanged()`** — L37636 — `private void chkBypassVACPlayingRecording_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkBypassVACPlayingRecording` checked state changes.
  Called by: `.ForceAllEvents()` (same file)
- **`.btnResetFMAF_rx_Click()`** — L37643 — `private void btnResetFMAF_rx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetFMAF_rx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnResetFMAF_tx_Click()`** — L37648 — `private void btnResetFMAF_tx_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetFMAF_tx` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudVoiceRecordingPlayback_slots_ValueChanged()`** — L37658 — `private void nudVoiceRecordingPlayback_slots_ValueChanged(object sender, EventArgs e)`
  used by nudVoiceRecordingPlayback_slots_ValueChanged
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateSlotSettings()`** — L37704 — `private void updateSlotSettings(int slots)`
  Called by: `.nudVoiceRecordingPlayback_slots_ValueChanged()` (same file)
- **`.preventIfContainerContainsLockedRecordings()`** — L37718 — `private bool preventIfContainerContainsLockedRecordings()`
  Called by: `.btnContainerDelete_Click()` (same file)
- **`.preventIfItemContainsLockedRecordings()`** — L37741 — `private bool preventIfItemContainsLockedRecordings()`
  Called by: `.btnRemoveMeterItem_Click()` (same file)
- **`.txtRecording_labelText_TextChanged()`** — L37770 — `private void txtRecording_labelText_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtRecording_labelText` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_canRepeat_CheckedChanged()`** — L37776 — `private void chkRecording_canRepeat_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_canRepeat` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_repeatDelay_ValueChanged()`** — L37782 — `private void nudRecording_repeatDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_repeatDelay` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_slot_locked_CheckedChanged()`** — L37788 — `private void chkRecording_slot_locked_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_slot_locked` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_playkeybind_CheckedChanged()`** — L37794 — `private void chkRecording_playkeybind_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_playkeybind` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_globalkeybind_CheckedChanged()`** — L37802 — `private void chkRecording_globalkeybind_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_globalkeybind` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_globalkeybind_assign_Click()`** — L37814 — `private void btnRecording_globalkeybind_assign_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_globalkeybind_assign` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_assingnkeybind_Click()`** — L37822 — `private void btnRecording_assingnkeybind_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_assingnkeybind` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.handleAssignKeybind()`** — L37828 — `private void handleAssignKeybind()`
  Called by: `.btnRecording_globalkeybind_assign_Click()` (same file), `.btnRecording_assingnkeybind_Click()` (same file)
- **`.recordingKeybindTimer_Tick()`** — L37877 — `private void recordingKeybindTimer_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `recordingKeybindTimer` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.onGlobalKeyUp()`** — L37893 — `private void onGlobalKeyUp(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onGlobalKeyDown()`** — L37915 — `private void onGlobalKeyDown(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnRecording_openStorageFolder_Click()`** — L38000 — `private void btnRecording_openStorageFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_openStorageFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtRecording_customFolder_KeyDown()`** — L38032 — `private void txtRecording_customFolder_KeyDown(object sender, KeyEventArgs e)`
  WinForms event handler: runs when `txtRecording_customFolder` receives a key-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtRecording_customFolder_KeyPress()`** — L38036 — `private void txtRecording_customFolder_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtRecording_customFolder` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_tx_gain_adjust_ValueChanged()`** — L38040 — `private void nudRecording_tx_gain_adjust_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_tx_gain_adjust` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_ignore_play_tempchanges_CheckedChanged()`** — L38045 — `private void chkRecording_ignore_play_tempchanges_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_ignore_play_tempchanges` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRecording_ignore_record_tempchanges_CheckedChanged()`** — L38050 — `private void chkRecording_ignore_record_tempchanges_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRecording_ignore_record_tempchanges` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRecording_in_chan_CheckedChanged()`** — L38055 — `private void radRecording_in_chan_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRecording_in_chan` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_4char_copy_Click()`** — L38061 — `private void btnRecording_4char_copy_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_4char_copy` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_slot_settings_ValueChanged()`** — L38066 — `private void nudRecording_slot_settings_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_slot_settings` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtRecording_4char_KeyPress()`** — L38073 — `private void txtRecording_4char_KeyPress(object sender, KeyPressEventArgs e)`
  WinForms event handler: runs when `txtRecording_4char` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudRecording_stop_free_space_ValueChanged()`** — L38078 — `private void nudRecording_stop_free_space_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudRecording_stop_free_space` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateRecordingFreeSpace()`** — L38085 — `private void updateRecordingFreeSpace()`
  Called by: `.radRecording_storage_CheckedChanged()` (same file), `.nudRecording_stop_free_space_ValueChanged()` (same file), `.tmrCheckStorageSpace_Tick()` (same file)
- **`.tmrCheckStorageSpace_Tick()`** — L38120 — `private void tmrCheckStorageSpace_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `tmrCheckStorageSpace` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_load_wav_to_slot_Click()`** — L38131 — `private void btnRecording_load_wav_to_slot_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_load_wav_to_slot` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnRecording_export_wav_from_slot_Click()`** — L38212 — `private void btnRecording_export_wav_from_slot_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnRecording_export_wav_from_slot` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkActivePeakRX1_tx_CheckedChanged()`** — L38302 — `private void chkActivePeakRX1_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkActivePeakRX1_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkActivePeakRX2_tx_CheckedChanged()`** — L38307 — `private void chkActivePeakRX2_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkActivePeakRX2_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSwapAudioChannels_CheckedChanged()`** — L38313 — `private void chkSwapAudioChannels_CheckedChanged(object sender, EventArgs e)`
  MI0BOT: Controls if the audio over the Over Protocol 1. There was a fudge to correct a problem in some HPSDR hardware
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCFC_legacy_CheckedChanged()`** — L38320 — `private void chkCFC_legacy_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCFC_legacy` checked state changes.
  Called by: `.ForceAllEvents()` (same file), `.loadTXProfile()` (same file)
- **`.btnCFCConfig_Click()`** — L38360 — `private void btnCFCConfig_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCFCConfig` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.closeCFCConfig()`** — L38372 — `private void closeCFCConfig()`
  Called by: `.btnOK_Click()` (same file), `.btnCancel_Click()` (same file), `.Setup_Closing()` (same file), `.chkCFC_legacy_CheckedChanged()` (same file)
- **`.setLegacyCFCProfile()`** — L38380 — `private void setLegacyCFCProfile()`
  Sets legacy cfcprofile.
  Called by: `.ForceAllEvents()` (same file), `.chkCFC_legacy_CheckedChanged()` (same file)
- **`.chkTCISwapIQ_CheckedChanged()`** — L38390 — `private void chkTCISwapIQ_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTCISwapIQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTCIAlwaysStreamIQ_CheckedChanged()`** — L38395 — `private void chkTCIAlwaysStreamIQ_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTCIAlwaysStreamIQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radTCITXchannel_CheckedChanged()`** — L38400 — `private void radTCITXchannel_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radTCITXchannel` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radWaterfall_timelab_CheckedChanged()`** — L38418 — `private void radWaterfall_timelab_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radWaterfall_timelab` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radWaterfall_timelab_time_CheckedChanged()`** — L38434 — `private void radWaterfall_timelab_time_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radWaterfall_timelab_time` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTCI_spot_flags_CheckedChanged()`** — L38446 — `private void chkTCI_spot_flags_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTCI_spot_flags` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clrbtnWaterfall_time_label_colour_Changed()`** — L38451 — `private void clrbtnWaterfall_time_label_colour_Changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tbTCISpotBackPanel_alpha_Scroll()`** — L38456 — `private void tbTCISpotBackPanel_alpha_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbTCISpotBackPanel_alpha` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkWaterfall_smear_CheckedChanged()`** — L38461 — `private void chkWaterfall_smear_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkWaterfall_smear` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `HIDComboItem` (type, L19496)

- **`.ToString()`** — L19510 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SetupTab` (type, L21138)

_No extracted members._

#### `PAProfile` (type, L24814)

- **`.base64Encode()`** — L24862 — `private string base64Encode(string plainText)`
  Called by: `.DataFromString()` (same file), `.DataToString()` (same file)
- **`.base64Decode()`** — L24874 — `private string base64Decode(string base64EncodedData)`
  Called by: `.DataFromString()` (same file)
- **`.DataFromString()`** — L24886 — `public void DataFromString(string sData)`
  Called by: `.getOptions()` (same file)
- **`.DataToString()`** — L24935 — `public string DataToString()`
  Called by: `.SaveOptions()` (same file)
- **`.GetGainForBand()`** — L24965 — `public float GetGainForBand(Band b, int nDriveValue = -1)`
  Returns gain for band.
  Called by: `.updateNUDgains()` (same file), `.nudPAProfileGain_ValueChanged()` (same file), `.GetPAGain()` (same file), `.GetBypassGain()` (same file)
- **`.calcDriveAdjust()`** — L24974 — `private float calcDriveAdjust(Band b, int nDriveValue)`
  Called by: `.GetGainForBand()` (same file)
- **`.lerp()`** — L25011 — `private float lerp(float a, float b, float frac)`
  Called by: `.calcDriveAdjust()` (same file)
- **`.SetMaxPower()`** — L25015 — `public void SetMaxPower(Band b, float maxPower)`
  Sets max power.
  Called by: `.nudMaxPowerForBandPA_ValueChanged()` (same file)
- **`.GetMaxPower()`** — L25020 — `public float GetMaxPower(Band b)`
  Returns max power.
  Called by: `.updateDriveLabels()` (same file), `.updateMaxPower()` (same file), `.GetPABandUsesMaxPower()` (same file), `.GetPABandMaxPower()` (same file)
- **`.SetMaxPowerUse()`** — L25025 — `public void SetMaxPowerUse(Band b, bool bUse)`
  Sets max power use.
  Called by: `.chkUsePowerOnDrvTunPA_CheckedChanged()` (same file)
- **`.GetMaxPowerUse()`** — L25030 — `public bool GetMaxPowerUse(Band b)`
  Returns max power use.
  Called by: `.updateDriveLabels()` (same file), `.GetPABandUsesMaxPower()` (same file), `.updateMaxPowerCheckbox()` (same file)
- **`.SetAdjust()`** — L25035 — `public void SetAdjust(Band b, int stepIndex, float gain)`
  Sets adjust.
  Called by: `.nudAdjustGain_ValueChanged()` (same file)
- **`.GetAdjust()`** — L25042 — `public float GetAdjust(Band b, int stepIndex)`
  Returns adjust.
  Called by: `.updateNUDAdjustgains()` (same file)
- **`.SetGainForBand()`** — L25049 — `public void SetGainForBand(Band b, float gain)`
  Sets gain for band.
  Called by: `.nudPAProfileGain_ValueChanged()` (same file), `.SetBypassGain()` (same file), `.handleOldPAGainSettings()` (same file), `.ResetGainDefaultsForModel()` (same file)
- **`.CopySettings()`** — L25055 — `public void CopySettings(PAProfile sourceProfile)`
  Called by: `.btnCopyPAProfile_Click()` (same file)
- **`.ResetGainDefaultsForModel()`** — L25078 — `public void ResetGainDefaultsForModel(HPSDRModel model)`
  Resets gain defaults for model.
  Called by: `.btnResetPAProfile_Click()` (same file)

#### `clsContainerComboboxItem` (type, L25367)

- **`.ToString()`** — L25373 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsMeterTypeComboboxItem` (type, L25378)

- **`.ToString()`** — L25397 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsMultiMeterIOComboboxItem` (type, L31610)

- **`.ToString()`** — L31683 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsComboHistoryItem` (type, L33231)

- **`.ToString()`** — L33248 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SelectFormsForReposition` (type, L38469)

- **`.getAllOpenForms()`** — L38528 — `public static List<Form> getAllOpenForms()`
  Returns all open forms.
  Called by: `.btnFormLocationHelper_Click()` (same file)
- **`.repositionForms()`** — L38545 — `private static void repositionForms(List<Form> forms, int start_x = 100, int start_y = 100, int step = 20)`
  Called by: `.on_ok_clicked()` (same file)
- **`.on_ok_clicked()`** — L38560 — `private void on_ok_clicked(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.on_cancel_clicked()`** — L38576 — `private void on_cancel_clicked(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `RepositionItem` (type, L38582)

- **`.ToString()`** — L38589 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `PADeviceInfo` (type, L38603)

- **`.ToString()`** — L38624 — `public override string ToString()`
  Returns the string representation.
  Called by: `.SaveOptions()` (same file), `.getOptions()` (same file), `.getTXProfileChangeReport()` (same file), `.checkTXProfileChanged2()` (same file), `.updateTXProfileInDB()` (same file), `.comboAudioSampleRate1_SelectedIndexChanged()` (same file) — and 72 more

#### `Channel60m` (type, L38637)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/setup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
