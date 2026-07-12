# `Console/console.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** The main window and central hub (~50k lines). Owns VFOs, band/mode/filter state, PTT/MOX sequencing, menus, and wires every other subsystem together. The graph's second-biggest god node (1,285 edges).

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/Andromeda.cs` (calls ×23)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×106)
  - `Console/radio.cs` (calls ×91, references ×8)
  - `Console/clsBandStackManager.cs` (calls ×59, references ×10)
  - `Console/Andromeda/Andromeda.cs` (calls ×45, references ×1)
  - `Console/xvtr.cs` (calls ×40, references ×1)
  - `Console/HPSDR/specHPSDR.cs` (calls ×26, references ×1)
  - `Console/common.cs` (calls ×23)
  - `Console/MeterManager.cs` (calls ×20, references ×1)
  - `Console/hiperftimer.cs` (calls ×14, references ×1)
  - `Console/TCIServer.cs` (calls ×12, references ×1)
  - `Console/display.cs` (calls ×11)
  - `Console/CAT/SIOListenerII.cs` (references ×7, calls ×3)
  - …and 63 more files
- Most-referenced symbols from other files: `.SetCATBand()` (×2), `.SetupRX2Band()` (×2), `.BandByFreq()` (×1), `.SelectRX1VarFilter()` (×1), `.SelectRX2VarFilter()` (×1), `.UpdateRX1Filters()` (×1), `.UpdateRX2Filters()` (×1), `.btnFilterShiftReset_Click()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.LockAsync()`** — L54926 — `public async Task<AsyncLock> LockAsync()`
  Called by: `.ATUTune()` (same file)

### Types

#### `Console` (type, L89)

- **`.SendMessage()`** — L582 — `[DllImport("user32.dll")] private static extern int SendMessage(IntPtr hWnd, Int32 wMsg, bool wParam, Int32 lParam)`
  Sends message.
  Called by: `.SuspendDrawing()` (same file), `.ResumeDrawing()` (same file)
- **`.RedrawWindow()`** — L584 — `[DllImport("user32.dll")] static extern bool RedrawWindow(IntPtr hWnd, IntPtr lprcUpdate, IntPtr hrgnUpdate, uint flags)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initialisePortAudio()`** — L1204 — `private void initialisePortAudio()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowReleaseNotes()`** — L1238 — `public void ShowReleaseNotes()`
  Shows release notes.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SuspendDrawing()`** — L1276 — `private void SuspendDrawing(Control control)`
  MW0LGE helper functions to suspend drawing/painting of controls. This is to help limit flicker during resize of console window
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file), `.pnlResizeMeter_MouseMove()` (same file)
- **`.ResumeDrawing()`** — L1282 — `private void ResumeDrawing(Control control, bool refresh = true)`
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file), `.pnlResizeMeter_MouseMove()` (same file)
- **`.OnAutoStartTimerEvent()`** — L1290 — `private void OnAutoStartTimerEvent(Object source, ElapsedEventArgs e)`
  Handles/raises the auto start timer event event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.gmh_MouseUp()`** — L1314 — `void gmh_MouseUp()`
  WinForms event handler: runs when `gmh` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Dispose()`** — L1319 — `protected override void Dispose(bool disposing)`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AttachConsole()`** — L1371 — `[DllImport("kernel32.dll")] private static extern bool AttachConsole(int dwProcessId)`
  Called by: `.showHelpInfo()` (same file)
- **`.FreeConsole()`** — L1373 — `[DllImport("kernel32.dll")] static extern bool FreeConsole()`
  Called by: `.showHelpInfo()` (same file)
- **`.showHelpInfo()`** — L1378 — `private static bool showHelpInfo()`
  Called by: `.Main()` (same file)
- **`.Application_ThreadException()`** — L1428 — `static void Application_ThreadException(object sender, ThreadExceptionEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CurrentDomain_UnhandledException()`** — L1433 — `static void CurrentDomain_UnhandledException(object sender, UnhandledExceptionEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cleanArg()`** — L1444 — `static string cleanArg(string input)`
  Called by: `.Main()` (same file)
- **`.Main()`** — L1468 — `[STAThread] static void Main(string[] args)`
  [DllImport("shcore.dll")] private static extern int SetProcessDpiAwareness(int awareness); Main
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.build_exception_text()`** — L1616 — `static string build_exception_text(Exception ex)`
  Called by: `.Main()` (same file)
- **`.getTitleWithFWVersion()`** — L1697 — `private string getTitleWithFWVersion()`
  Returns title with fwversion.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onNotchDelete()`** — L1739 — `private void onNotchDelete(int notch_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onBWChanged()`** — L1747 — `private void onBWChanged(int notch_index, double width)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onActiveChanged()`** — L1755 — `private void onActiveChanged(int notch_index, bool active)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onClearButton()`** — L1763 — `private void onClearButton()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitConsole()`** — L1769 — `private void InitConsole()`
  Inits console.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupInfoBarButton()`** — L2318 — `public void SetupInfoBarButton(ucInfoBar.ActionTypes action, bool bEnabled)`
  Setups info bar button.
  Called by: `.chkRX2_CheckedChanged()` (same file)
- **`.SetupInfoBar()`** — L2322 — `private void SetupInfoBar()`
  Setups info bar.
  Called by: `.InitConsole()` (same file)
- **`.InfoBarFeedbackLevel()`** — L2350 — `public void InfoBarFeedbackLevel(int level, bool bFeedbackLevelOk, bool bCorrectionsBeingApplied, bool bCalibrationAttemptsChanged, Color feedbackColour)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowTCPIPCatLog()`** — L2382 — `public void ShowTCPIPCatLog()`
  Shows tcpipcat log.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatClientConnect()`** — L2386 — `private void OnTCPIIPcatClientConnect()`
  Handles/raises the tcpiipcat client connect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatClientDisconnect()`** — L2392 — `private void OnTCPIIPcatClientDisconnect()`
  Handles/raises the tcpiipcat client disconnect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatClientError()`** — L2398 — `private void OnTCPIIPcatClientError(SocketException se)`
  Handles/raises the tcpiipcat client error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatServerError()`** — L2402 — `private void OnTCPIIPcatServerError(SocketException se)`
  Handles/raises the tcpiipcat server error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addTCPIPcatDelegates()`** — L2407 — `private void addTCPIPcatDelegates()`
  Called by: `.SetupTCPIPCat()` (same file)
- **`.removeTCPIPcatDelegates()`** — L2416 — `private void removeTCPIPcatDelegates()`
  Called by: `.SetupTCPIPCat()` (same file), `.Console_Closing()` (same file)
- **`.SetupTCPIPCat()`** — L2439 — `public void SetupTCPIPCat(bool bOn)`
  Setups tcpipcat.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.closeTcpIpCatServer()`** — L2485 — `private void closeTcpIpCatServer()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addTCIDelegates()`** — L2515 — `private void addTCIDelegates()`
  Called by: `.SetupTCI()` (same file)
- **`.removeTCIDelegates()`** — L2524 — `private void removeTCIDelegates()`
  Called by: `.SetupTCI()` (same file), `.Console_Closing()` (same file)
- **`.OnTCIClientConnect()`** — L2535 — `private void OnTCIClientConnect()`
  Handles/raises the tciclient connect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCIClientDisconnect()`** — L2541 — `private void OnTCIClientDisconnect()`
  Handles/raises the tciclient disconnect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCIClientError()`** — L2547 — `private void OnTCIClientError(SocketException se)`
  Handles/raises the tciclient error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCIServerError()`** — L2551 — `private void OnTCIServerError(SocketException se)`
  Handles/raises the tciserver error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowTCILog()`** — L2555 — `public void ShowTCILog()`
  Shows tcilog.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupTCI()`** — L2559 — `public void SetupTCI(bool bOn, int rateLimit)`
  Setups tci.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.selectFilters()`** — L2609 — `private void selectFilters()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.selectModes()`** — L2655 — `private void selectModes()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Init60mChannels()`** — L2681 — `public void Init60mChannels()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SyncDSP()`** — L2716 — `private void SyncDSP()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExitConsole()`** — L2765 — `public void ExitConsole()`
  Called by: `.Dispose()` (same file)
- **`.StateListToBase64()`** — L2823 — `public string StateListToBase64()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetStateList()`** — L2829 — `public List<string> GetStateList()`
  Returns state list.
  Called by: `.StateListToBase64()` (same file), `.SaveState()` (same file)
- **`.addControlState()`** — L3313 — `private void addControlState(List<string> a, Control c, List<string> always_save, bool combo_use_items_count, bool print_not_saving)`
  Called by: `.GetStateList()` (same file)
- **`.SaveState()`** — L3365 — `public void SaveState()`
  Saves state.
  Called by: `.Console_Closing()` (same file)
- **`.GetState()`** — L3959 — `public void GetState()`
  Returns state.
  Called by: `.InitConsole()` (same file)
- **`.InitFilterPresets()`** — L5217 — `private void InitFilterPresets()`
  Inits filter presets.
  Called by: `.InitConsole()` (same file), `.toolStripMenuItemRX1FilterReset_Click()` (same file), `.toolStripMenuItemRX2FilterReset_Click()` (same file)
- **`.InitDisplayModes()`** — L5620 — `private void InitDisplayModes()`
  Inits display modes.
  Called by: `.InitConsole()` (same file)
- **`.InitAGCModes()`** — L5631 — `private void InitAGCModes()`
  Inits agcmodes.
  Called by: `.InitConsole()` (same file)
- **`.InitMultiMeterModes()`** — L5649 — `private void InitMultiMeterModes()`
  Inits multi meter modes.
  Called by: `.InitConsole()` (same file)
- **`.DisableAllFilters()`** — L5681 — `private void DisableAllFilters()`
  Disables all filters.
  Called by: `.CalibratePAGain()` (same file), `.SetRX1Mode()` (same file)
- **`.DisableAllRX2Filters()`** — L5693 — `private void DisableAllRX2Filters()`
  Disables all rx2 filters.
  Called by: `.SetRX2Mode()` (same file)
- **`.EnableAllFilters()`** — L5704 — `private void EnableAllFilters()`
  Enables all filters.
  Called by: `.CalibratePAGain()` (same file), `.SetRX1Mode()` (same file)
- **`.EnableAllRX2Filters()`** — L5721 — `private void EnableAllRX2Filters()`
  Enables all rx2 filters.
  Called by: `.SetRX2Mode()` (same file)
- **`.DisableAllBands()`** — L5737 — `private void DisableAllBands()`
  Disables all bands.
  Called by: `.UIMOXChangedTrue()` (same file)
- **`.EnableAllBands()`** — L5781 — `private void EnableAllBands()`
  Enables all bands.
  Called by: `.UIMOXChangedFalse()` (same file)
- **`.DisableAllModes()`** — L5830 — `private void DisableAllModes()`
  Disables all modes.
  Called by: `.CalibratePAGain()` (same file), `.UIMOXChangedTrue()` (same file)
- **`.EnableAllModes()`** — L5840 — `private void EnableAllModes()`
  Enables all modes.
  Called by: `.CalibratePAGain()` (same file), `.UIMOXChangedFalse()` (same file)
- **`.GetVFOCharWidth()`** — L5851 — `private void GetVFOCharWidth()`
  Returns vfochar width.
  Called by: `.Console_MouseWheel()` (same file), `.txtVFOAFreq_MouseMove()` (same file), `.txtVFOBFreq_MouseMove()` (same file)
- **`.GetVFOSubCharWidth()`** — L5884 — `private void GetVFOSubCharWidth()`
  Returns vfosub char width.
  Called by: `.txtVFOABand_MouseMove()` (same file)
- **`.SetBand()`** — L5910 — `public void SetBand(string mode, string filter, double freq, bool CTUN, int zoomFactor, double centerFreq)`
  Sets band.
  Called by: `.setRX1BandFromBandStackEntry()` (same file)
- **`.getButtonForBand()`** — L6019 — `private RadioButtonTS getButtonForBand(Band b)`
  Returns button for band.
  Called by: `.SetBand()` (same file)
- **`.ChangeTuneStepUp()`** — L6159 — `public void ChangeTuneStepUp()`
  Called by: `.InitConsole()` (same file), `.CATTuneStepUp()` (same file), `.Console_KeyDown()` (same file), `.WheelTune_MouseDown()` (same file), `.btnChangeTuneStepLarger_Click()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.ChangeTuneStepDown()`** — L6165 — `public void ChangeTuneStepDown()`
  Called by: `.CATTuneStepDown()` (same file), `.Console_KeyDown()` (same file), `.btnChangeTuneStepSmaller_Click()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.UpdateBandButtonColors()`** — L6171 — `private void UpdateBandButtonColors()`
  Updates band button colors.
  Called by: `.SetRX1Band()` (same file), `.SetRX2Band()` (same file), `.SetTXBand()` (same file), `.UpdateWaterfallLevelValues()` (same file)
- **`.DeselectHF()`** — L6176 — `private void DeselectHF()`
  Called by: `.SetRX1BandButton()` (same file)
- **`.DeselectGEN()`** — L6195 — `private void DeselectGEN()`
  Called by: `.SetRX1BandButton()` (same file)
- **`.DeselectVHF()`** — L6214 — `private void DeselectVHF()`
  Called by: `.SetRX1BandButton()` (same file)
- **`.SetRX1BandButton()`** — L6232 — `private void SetRX1BandButton(Band b)`
  Sets rx1 band button.
  Called by: `.UpdateBandButtonColors()` (same file)
- **`.RX1BandForVFOB()`** — L6451 — `public Band RX1BandForVFOB()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTransverterTranslatedRXBand()`** — L6472 — `public Band GetTransverterTranslatedRXBand(double freq)`
  Returns transverter translated rxband.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTransverterTranslatedTXBand()`** — L6479 — `public Band GetTransverterTranslatedTXBand()`
  Returns transverter translated txband.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandByFreq()`** — L6486 — `private Band BandByFreq(double freq, int xvtr_index, FRSRegion region)`
  Called by: `.SetBand()` (same file), `.RX1BandForVFOB()` (same file), `.GetTransverterTranslatedRXBand()` (same file), `.GetTransverterTranslatedTXBand()` (same file), `.VFOASubUpdate()` (same file), `.HdwMOXChanged()` (same file) — and 10 more
- **`.SetRX1Band()`** — L6498 — `private void SetRX1Band(Band b)`
  Sets rx1 band.
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.SetRX2Band()`** — L6533 — `private void SetRX2Band(Band b)`
  Sets rx2 band.
  Called by: `.txtVFOBFreq_LostFocus()` (same file)
- **`.SetTXBand()`** — L6545 — `private void SetTXBand(Band b, bool bIngoreBandChange = false)`
  Sets txband.
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.GainByBand()`** — L6565 — `private float GainByBand(Band b, int nDriveValue)`
  Called by: `.SetPowerUsingTargetDBM()` (same file)
- **`.CheckSelectedButtonColor()`** — L6572 — `public void CheckSelectedButtonColor()`
  Checks selected button color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PABandOffset()`** — L6642 — `private double PABandOffset(Band b)`
  Called by: `.ScaledVoltage()` (same file), `.ADCtodBm()` (same file)
- **`.SWR()`** — L6685 — `private double SWR(int adc_fwd, int adc_rev)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ScaledVoltage()`** — L6700 — `private double ScaledVoltage(int adc)`
  Called by: `.SWR()` (same file), `.PAPower()` (same file)
- **`.ADCtodBm()`** — L6707 — `private double ADCtodBm(int adc_data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PAPower()`** — L6716 — `private double PAPower(int adc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WattsTodBm()`** — L6724 — `private double WattsTodBm(double watts)`
  Called by: `.CalibratePAGain()` (same file)
- **`.dBmToWatts()`** — L6729 — `private double dBmToWatts(double dBm)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibratedPAPower()`** — L6734 — `public float CalibratedPAPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.PowerKernel()`** — L6800 — `private float PowerKernel(float watts, float interval, int entries, float[] table)`
  Called by: `.CalibratedPAPower()` (same file)
- **`.CheckValidTXFreq()`** — L6814 — `public bool CheckValidTXFreq(FRSRegion r, double f, DSPMode mode, bool bIgnoreFilter = false)`
  Checks valid txfreq.
  Called by: `.chkMOX_CheckedChanged2()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.checkValidTXFreq_local()`** — L6852 — `private bool checkValidTXFreq_local(FRSRegion r, double f)`
  Called by: `.CheckValidTXFreq()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.setAlex1HPF()`** — L6863 — `private void setAlex1HPF(double freq)`
  Sets alex1 hpf.
  Called by: `.UpdateRX1DDSFreq()` (same file), `.UpdateTXDDSFreq()` (same file), `.UpdateAlexRXFilter()` (same file)
- **`.setAlexHPF()`** — L6875 — `private void setAlexHPF(double freq)`
  Sets alex hpf.
  Called by: `.setAlex1HPF()` (same file)
- **`.setBPF1ForOrionIISaturn()`** — L6989 — `private void setBPF1ForOrionIISaturn(double freq)`
  Sets bpf1 for orion iisaturn.
  Called by: `.setAlex1HPF()` (same file)
- **`.setAlex2HPF()`** — L7105 — `private void setAlex2HPF(double freq)`
  Sets alex2 hpf.
  Called by: `.UpdateRX2DDSFreq()` (same file)
- **`.setAlexLPF()`** — L7213 — `private void setAlexLPF(double freq, bool freqIsTX)`
  Sets alex lpf.
  Called by: `.UpdateTXDDSFreq()` (same file), `.UpdateAlexTXFilter()` (same file)
- **`.setAlex2LPF()`** — L7281 — `private void setAlex2LPF(double freq)`
  Sets alex2 lpf.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectRX1VarFilter()`** — L7343 — `public void SelectRX1VarFilter(bool update = true, bool prevent_update = false)`
  Selects rx1 var filter.
  Called by: `.Console_KeyDown()` (same file), `.ptbFilterShift_Scroll()` (same file), `.ptbFilterWidth_Scroll()` (same file), `.pnlDisplay_MouseMove()` (same file), `.ExecuteEncoderStep()` (`Console/Andromeda/Andromeda.cs`)
- **`.SelectRX2VarFilter()`** — L7356 — `public void SelectRX2VarFilter(bool update = true, bool prevent_update = false)`
  Selects rx2 var filter.
  Called by: `.pnlDisplay_MouseMove()` (same file), `.ExecuteEncoderStep()` (`Console/Andromeda/Andromeda.cs`)
- **`.UpdateRXADCCtrlP1()`** — L7370 — `public void UpdateRXADCCtrlP1()`
  Updates rxadcctrl p1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMemoryQS()`** — L7376 — `public void CATMemoryQS()`
  Added 06/24/05 BT for CAT commands
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMemoryQR()`** — L7382 — `public void CATMemoryQR()`
  Added 06/25/05 BT for CAT commands
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCATBand()`** — L7431 — `public void SetCATBand(Band pBand)`
  BT 06/17/05 added for CAT commands
  Called by: `.HandleFrontPanelButtonPress()` (`Console/Andromeda/Andromeda.cs`), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.SetVHFText()`** — L7528 — `public void SetVHFText(int index, string text)`
  Sets vhftext.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetVHFEnabled()`** — L7535 — `public void SetVHFEnabled(int index, bool b)`
  Sets vhfenabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetVHFText()`** — L7543 — `public string GetVHFText(int index)`
  G8NJJ added to allow labelling of buttons in popup form
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetVHFEnabled()`** — L7547 — `public bool GetVHFEnabled(int index)`
  Returns vhfenabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1Filters()`** — L7563 — `public void UpdateRX1Filters(int low, int high, bool force = false, bool from_change_event = false)`
  Updates rx1 filters.
  Called by: `.SelectRX1VarFilter()` (same file), `.UpdateRX1FilterPresetLow()` (same file), `.UpdateRX1FilterPresetHigh()` (same file), `.Console_KeyDown()` (same file), `.SetRX1Mode()` (same file), `.SetRX1Filter()` (same file) — and 9 more
- **`.UpdateRX2Filters()`** — L7686 — `public void UpdateRX2Filters(int low, int high, bool force = false, bool from_change_event = false)`
  Updates rx2 filters.
  Called by: `.SelectRX2VarFilter()` (same file), `.UpdateRX2FilterPresetLow()` (same file), `.UpdateRX2FilterPresetHigh()` (same file), `.SetRX2Mode()` (same file), `.SetRX2Filter()` (same file), `.udRX2FilterLow_ValueChanged()` (same file) — and 5 more
- **`.UpdateRX1FilterNames()`** — L7784 — `public void UpdateRX1FilterNames(Filter f, string old_name, string new_name)`
  Updates rx1 filter names.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1FilterPresetLow()`** — L7832 — `public void UpdateRX1FilterPresetLow(int val)`
  Updates rx1 filter preset low.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1FilterPresetHigh()`** — L7837 — `public void UpdateRX1FilterPresetHigh(int val)`
  Updates rx1 filter preset high.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2FilterNames()`** — L7842 — `public void UpdateRX2FilterNames(Filter f, string old_name, string new_name)`
  Updates rx2 filter names.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2FilterPresetLow()`** — L7881 — `public void UpdateRX2FilterPresetLow(int val)`
  Updates rx2 filter preset low.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2FilterPresetHigh()`** — L7886 — `public void UpdateRX2FilterPresetHigh(int val)`
  Updates rx2 filter preset high.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateVFOAFreq()`** — L7891 — `public void UpdateVFOAFreq(string freq)`
  Updates vfoafreq.
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.sioPut()`** — L7902 — `private bool sioPut(object sio_listener, string msg)`
  Called by: `.on_send_floodcontrol_message()` (same file)
- **`.BroadcastFreqChange()`** — L7959 — `private void BroadcastFreqChange(string vfo, double freq)`
  Called by: `.OnVFOAFrequencyChangeHandler()` (same file), `.OnVFOBFrequencyChangeHandler()` (same file)
- **`.UpdateVFOBFreq()`** — L7968 — `public void UpdateVFOBFreq(string freq)`
  Updates vfobfreq.
  Called by: `.txtVFOBFreq_LostFocus()` (same file)
- **`.CalcDisplayFreq()`** — L7979 — `public void CalcDisplayFreq()`
  Called by: `.InitConsole()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ptbDisplayPan_Scroll()` (same file), `.ptbDisplayZoom_Scroll()` (same file), `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file) — and 1 more
- **`.CalcRX2DisplayFreq()`** — L7991 — `public void CalcRX2DisplayFreq()`
  Called by: `.InitConsole()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ptbDisplayPan_Scroll()` (same file), `.ptbDisplayZoom_Scroll()` (same file)
- **`.CalcTXDisplayFreq()`** — L8002 — `public void CalcTXDisplayFreq()`
  Called by: `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ptbDisplayPan_Scroll()` (same file), `.ptbDisplayZoom_Scroll()` (same file)
- **`.UpdateRXSpectrumDisplayVars()`** — L8008 — `public void UpdateRXSpectrumDisplayVars()`
  Updates rxspectrum display vars.
  Called by: `.chkDisplayAVG_CheckedChanged()` (same file), `.chkDisplayPeak_CheckedChanged()` (same file), `.UpdateDSP()` (same file), `.pnlDisplay_Resize()` (same file)
- **`.UpdateTXSpectrumDisplayVars()`** — L8013 — `public void UpdateTXSpectrumDisplayVars()`
  Updates txspectrum display vars.
  Called by: `.comboDisplayMode_SelectedIndexChanged()` (same file), `.chkDisplayAVG_CheckedChanged()` (same file), `.chkDisplayPeak_CheckedChanged()` (same file), `.UpdateDSP()` (same file), `.pnlDisplay_Resize()` (same file)
- **`.UpdateRXDisplayVars()`** — L8018 — `private void UpdateRXDisplayVars(int l, int h)`
  Updates rxdisplay vars.
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRXSpectrumDisplayVars()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.comboRX2DisplayMode_SelectedIndexChanged()` (same file)
- **`.UpdateTXDisplayVars()`** — L8060 — `private void UpdateTXDisplayVars(int l, int h)`
  Updates txdisplay vars.
  Called by: `.UpdateTXSpectrumDisplayVars()` (same file), `.SetTXFilters()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file)
- **`.UpdateTXLowHighFilterForMode()`** — L8096 — `public void UpdateTXLowHighFilterForMode(DSPMode mode, ref int low, ref int high)`
  Updates txlow high filter for mode.
  Called by: `.SetTXFilters()` (same file), `.SetTXFilter()` (same file)
- **`.SetTXFilters()`** — L8136 — `public void SetTXFilters(DSPMode mode, int low, int high, bool force = false)`
  Sets txfilters.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file), `.fmDeviation2k()` (same file), `.fmDeviation5k()` (same file)
- **`.BuildTXProfileCombos()`** — L8170 — `public void BuildTXProfileCombos()`
  Builds txprofile combos.
  Called by: `.InitConsole()` (same file), `.UpdateTXProfile()` (same file)
- **`.UpdateTXProfile()`** — L8184 — `public void UpdateTXProfile(string name)`
  Updates txprofile.
  Called by: `.InitConsole()` (same file)
- **`.UpdateDDCs()`** — L8231 — `public void UpdateDDCs(bool rx2_enabled)`
  Diversity operation is on RX1; therefore, the 'rx1_rate' will be used as the diversity rate;
  Called by: `.SetupForHPSDRModel()` (same file), `.chkPower_CheckedChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.txtVFOAFreq_LostFocus()` (same file)
- **`.GetDDC()`** — L8597 — `public void GetDDC(out int DDCrx1, out int DDCrx2, out int DDCsync1, out int DDCsync2, out int DDCpsrx, out int DDCpstx)`
  Returns ddc.
  Called by: `.handleOverload()` (same file), `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.UpdateDiversityMenuItem()`** — L8863 — `public void UpdateDiversityMenuItem()`
  Updates diversity menu item.
  Called by: `.showHideDiversity()` (same file)
- **`.UpdateDiversityValues()`** — L8871 — `private void UpdateDiversityValues()`
  Updates diversity values.
  Called by: `.SetRX1Band()` (same file), `.showHideDiversity()` (same file), `.setRX1BandFromBandStackEntry()` (same file)
- **`.UpdateWaterfallLevelValues()`** — L8953 — `public void UpdateWaterfallLevelValues()`
  Updates waterfall level values.
  Called by: `.SetRX1Band()` (same file), `.SetRX2Band()` (same file), `.setRX1BandFromBandStackEntry()` (same file)
- **`.updateDisplayGridLevelValues()`** — L9150 — `private void updateDisplayGridLevelValues()`
  Called by: `.SetRX1Band()` (same file), `.SetRX2Band()` (same file), `.setRX1BandFromBandStackEntry()` (same file)
- **`.setWaterfallGainsIfLinkedToSpectrum()`** — L9155 — `private void setWaterfallGainsIfLinkedToSpectrum(int rx)`
  Sets waterfall gains if linked to spectrum.
  Called by: `.UpdateWaterfallLevelValues()` (same file), `.UpdateDisplayGridLevelMinValues()` (same file), `.UpdateDisplayGridLevelMaxValues()` (same file)
- **`.CheckForMinMaxGridUpdatesTX()`** — L9170 — `public void CheckForMinMaxGridUpdatesTX()`
  Checks for min max grid updates tx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CheckForMinMaxGridUpdatesRX()`** — L9192 — `public void CheckForMinMaxGridUpdatesRX(int rx)`
  Checks for min max grid updates rx.
  Called by: `.UpdateDisplayGridLevelMinValues()` (same file), `.UpdateDisplayGridLevelMaxValues()` (same file)
- **`.CheckForMinMaxWaterfallUpdatesTX()`** — L9249 — `public void CheckForMinMaxWaterfallUpdatesTX()`
  Checks for min max waterfall updates tx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CheckForMinMaxWaterfallUpdatesRX()`** — L9271 — `public void CheckForMinMaxWaterfallUpdatesRX(int rx)`
  Checks for min max waterfall updates rx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateDisplayGridLevelMinValues()`** — L9328 — `public void UpdateDisplayGridLevelMinValues(bool bDoBandInfoAndWaterFallSync)`
  Updates display grid level min values.
  Called by: `.updateDisplayGridLevelValues()` (same file)
- **`.UpdateDisplayGridLevelMaxValues()`** — L9465 — `public void UpdateDisplayGridLevelMaxValues(bool bDoBandInfoAndWaterFallSync)`
  Updates display grid level max values.
  Called by: `.updateDisplayGridLevelValues()` (same file)
- **`.RX1IsIn60m()`** — L9603 — `public bool RX1IsIn60m()`
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.RX1IsOn60mChannel()`** — L9609 — `public bool RX1IsOn60mChannel(Channel c)`
  Called by: `.UpdatePeakText()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.SetRX1Mode()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.RX1IsIn60mChannel()`** — L9616 — `public bool RX1IsIn60mChannel(Channel c)`
  Called by: `.SetRX1Mode()` (same file)
- **`.RX2IsIn60m()`** — L9643 — `public bool RX2IsIn60m()`
  Called by: `.txtVFOBFreq_LostFocus()` (same file)
- **`.RX2IsOn60mChannel()`** — L9649 — `public bool RX2IsOn60mChannel(Channel c)`
  Called by: `.SetRX2Mode()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.RX2IsIn60mChannel()`** — L9656 — `public bool RX2IsIn60mChannel(Channel c)`
  Called by: `.SetRX2Mode()` (same file)
- **`.ModeFreqOffset()`** — L9682 — `private double ModeFreqOffset(DSPMode mode)`
  Called by: `.RX1IsOn60mChannel()` (same file), `.RX2IsOn60mChannel()` (same file), `.UpdatePeakText()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file), `.SetRX1Mode()` (same file) — and 1 more
- **`.CATGetXVTRBandNames()`** — L9704 — `public string CATGetXVTRBandNames()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetFilterPresets()`** — L9717 — `public string GetFilterPresets(int mode_ndx)`
  Returns filter presets.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radBandGEN0_Click()`** — L9740 — `private void radBandGEN0_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN0` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN1_Click()`** — L9746 — `private void radBandGEN1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN2_Click()`** — L9752 — `private void radBandGEN2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN3_Click()`** — L9758 — `private void radBandGEN3_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN3` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN4_Click()`** — L9764 — `private void radBandGEN4_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN4` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN5_Click()`** — L9770 — `private void radBandGEN5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN6_Click()`** — L9776 — `private void radBandGEN6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN6` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN7_Click()`** — L9782 — `private void radBandGEN7_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN7` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN8_Click()`** — L9788 — `private void radBandGEN8_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN8` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN9_Click()`** — L9794 — `private void radBandGEN9_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN9` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN10_Click()`** — L9800 — `private void radBandGEN10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN10` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN11_Click()`** — L9806 — `private void radBandGEN11_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN11` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN12_Click()`** — L9812 — `private void radBandGEN12_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN12` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN13_Click()`** — L9818 — `private void radBandGEN13_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN13` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CalibrateFreq()`** — L9828 — `unsafe public bool CalibrateFreq(float freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibrateLevel()`** — L9905 — `unsafe public bool CalibrateLevel(float level, float freq, Progress progress, bool suppress_errors)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibratePAGain()`** — L10290 — `public bool CalibratePAGain(Progress progress, bool[] run, int target_watts)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibratePAGain2()`** — L10451 — `public bool CalibratePAGain2(Progress progress, bool[] run, bool suppress_warnings)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LowPowerPASweep()`** — L10456 — `public bool LowPowerPASweep(Progress progress, int power)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.validateRX1StepAttData()`** — L10989 — `private int validateRX1StepAttData(int att)`
  [2.10.3.9]MW0LGE validate step attenuator values, so many HL2 issues with these being out of range
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.validateRX2StepAttData()`** — L10995 — `private int validateRX2StepAttData(int att)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.validateTXStepAttData()`** — L11001 — `private int validateTXStepAttData(int att)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TuneStepLookup()`** — L11384 — `public int TuneStepLookup(string s)`
  Called by: `.RecallMemory()` (same file)
- **`.isDSPModeValid()`** — L11394 — `private bool isDSPModeValid(DSPMode mode)`
  Called by: `.updateStepIndexForMode()` (same file)
- **`.updateStepIndexForMode()`** — L11399 — `private void updateStepIndexForMode(int rx, DSPMode mode)`
  Called by: `.OnModeChangeHandler()` (same file)
- **`.SetAFLinks()`** — L11492 — `public void SetAFLinks(int source, bool state)`
  Sets aflinks.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsMasterAFLinked()`** — L11517 — `public bool IsMasterAFLinked(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setLinkedAF()`** — L11524 — `private void setLinkedAF(int source, int value)`
  Sets linked af.
  Called by: `.SetAFLinks()` (same file), `.ptbAF_Scroll()` (same file), `.ptbRX0Gain_Scroll()` (same file), `.ptbRX1Gain_Scroll()` (same file), `.ptbRX2Gain_Scroll()` (same file)
- **`.HighlightTXProfileSaveItems()`** — L12010 — `public void HighlightTXProfileSaveItems(bool bHighlight)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPower()`** — L12201 — `public void SetPower(Band b, int pwr)`
  Sets power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPower()`** — L12207 — `public int GetPower(Band b)`
  Returns power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1DisplayOffsets()`** — L12387 — `private void UpdateRX1DisplayOffsets()`
  Updates rx1 display offsets.
  Called by: `.InitConsole()` (same file), `.CalibrateLevel()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ResetLevelCalibration()` (same file), `.udTXStepAttData_ValueChanged()` (same file)
- **`.UpdateRX2DisplayOffsets()`** — L12395 — `private void UpdateRX2DisplayOffsets()`
  Updates rx2 display offsets.
  Called by: `.InitConsole()` (same file), `.CalibrateLevel()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ResetLevelCalibration()` (same file), `.udTXStepAttData_ValueChanged()` (same file)
- **`.SafeTXProfileSet()`** — L12484 — `public void SafeTXProfileSet(string profile)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCWSideToneVolume()`** — L13127 — `private void setCWSideToneVolume()`
  Sets cwside tone volume.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupDisplayEngine()`** — L13618 — `public void SetupDisplayEngine(int decimation)`
  MW0LGE_21k9
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupForHPSDRModel()`** — L14864 — `public void SetupForHPSDRModel()`
  Setups for hpsdrmodel.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateOCTXPins()`** — L15007 — `private void updateOCTXPins(bool tx)`
  Called by: `.OnMoxChangeHandler()` (same file)
- **`.UpdateTRXAnt()`** — L15056 — `private void UpdateTRXAnt()`
  Updates trxant.
  Called by: `.HdwMOXChanged()` (same file), `.txtVFOAFreq_LostFocus()` (same file)
- **`.enableMONForCW()`** — L15103 — `private void enableMONForCW()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetADCInUse()`** — L15209 — `public int GetADCInUse(int ddc)`
  Returns adcin use.
  Called by: `.handleOverload()` (same file), `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.SetWavePlayback()`** — L15378 — `public void SetWavePlayback(int id, bool enabled)`
  Sets wave playback.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getWavePlayback()`** — L15398 — `private bool getWavePlayback(int id)`
  Returns wave playback.
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.adjustForSnapClickTuning()` (same file)
- **`.getWavePlaybackFreq()`** — L15403 — `public double getWavePlaybackFreq(int id)`
  Returns wave playback freq.
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.UpdateRX1DDSFreq()`** — L15497 — `private void UpdateRX1DDSFreq()`
  Updates rx1 ddsfreq.
  Called by: `.HdwMOXChanged()` (same file)
- **`.UpdateRX2DDSFreq()`** — L15527 — `private void UpdateRX2DDSFreq()`
  Updates rx2 ddsfreq.
  Called by: `.HdwMOXChanged()` (same file)
- **`.UpdateTXDDSFreq()`** — L15565 — `private void UpdateTXDDSFreq()`
  Updates txddsfreq.
  Called by: `.HdwMOXChanged()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.UpdateAlexTXFilter()`** — L15588 — `private void UpdateAlexTXFilter()`
  Updates alex txfilter.
  Called by: `.UpdateRX1DDSFreq()` (same file), `.UpdateRX2DDSFreq()` (same file)
- **`.UpdateAlexRXFilter()`** — L15601 — `private void UpdateAlexRXFilter()`
  Updates alex rxfilter.
  Called by: `.UpdateRX1DDSFreq()` (same file), `.UpdateRX2DDSFreq()` (same file)
- **`.ThreadSafeCatParse()`** — L15758 — `public string ThreadSafeCatParse(string msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.safeCat()`** — L15767 — `private string safeCat(string msg)`
  Called by: `.ThreadSafeCatParse()` (same file)
- **`.CATVFOAtoB()`** — L15884 — `public void CATVFOAtoB()`
  -W2PA Added three new functions to make CAT functions match behavior of equivalent console functions. i.e. not just copy frequency alone
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATVFOBtoA()`** — L15888 — `public void CATVFOBtoA()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATVFOABSwap()`** — L15892 — `public void CATVFOABSwap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATTuneStepUp()`** — L16127 — `public void CATTuneStepUp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATTuneStepDown()`** — L16132 — `public void CATTuneStepDown()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMidiMessagesPerTuneStepUp()`** — L16186 — `public void CATMidiMessagesPerTuneStepUp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMidiMessagesPerTuneStepDown()`** — L16191 — `public void CATMidiMessagesPerTuneStepDown()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMidiMessagesPerTuneStepToggle()`** — L16196 — `public void CATMidiMessagesPerTuneStepToggle()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATSingleCal()`** — L16255 — `public void CATSingleCal()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATRX2BandUpDown()`** — L16482 — `public void CATRX2BandUpDown(int direction)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandToString()`** — L17457 — `private string BandToString(Band b)`
  Called by: `.SetupRX2Band()` (same file)
- **`.StringToBand()`** — L17510 — `private Band StringToBand(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadSigStrength()`** — L17713 — `public string CATReadSigStrength()`
  Added 07/30/05 BT for cat commands next 8 functions
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadAvgStrength()`** — L17720 — `public string CATReadAvgStrength()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadADC_L()`** — L17727 — `public string CATReadADC_L()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadADC_R()`** — L17733 — `public string CATReadADC_R()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadALC()`** — L17739 — `public string CATReadALC()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadFwdPwr()`** — L17751 — `public string CATReadFwdPwr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadPeakPwr()`** — L17767 — `public string CATReadPeakPwr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadRevPwr()`** — L17803 — `public string CATReadRevPwr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadSWR()`** — L17809 — `public string CATReadSWR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXFilter()`** — L17949 — `public bool SetTXFilter(Filter filter)`
  Sets txfilter.
  Called by: `.SetTXFilters()` (same file), `.MatchTXFilterToRXFilter()` (same file)
- **`.GetDSPcwPitchShiftToZero()`** — L18289 — `public int GetDSPcwPitchShiftToZero(int rx)`
  Returns dspcw pitch shift to zero.
  Called by: `.AddNotch()` (same file)
- **`.freqFromString()`** — L18425 — `static double freqFromString(string s)`
  Called by: `.UpdateVFOAFreq()` (same file), `.UpdateVFOBFreq()` (same file), `.btnMemoryQuickRestore_Click()` (same file)
- **`.VFOAUpdate()`** — L18457 — `private void VFOAUpdate(double freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOBUpdate()`** — L18463 — `private void VFOBUpdate(double freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOASubUpdate()`** — L18469 — `private void VFOASubUpdate(double freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PanCentre()`** — L19712 — `public void PanCentre()`
  Called by: `.InitConsole()` (same file), `.zoomToBandBandwidth()` (same file), `.displayZoom05()` (same file), `.displayZoom1()` (same file), `.displayZoom2()` (same file), `.displayZoom4()` (same file) — and 1 more
- **`.ZoomFullyOut()`** — L19716 — `public void ZoomFullyOut()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTimer()`** — L19886 — `private void SetTimer(System.Windows.Forms.Timer t, bool enable)`
  Sets timer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CpuUsage()`** — L20824 — `private void CpuUsage()`
  Called by: `.systemToolStripMenuItem_Click()` (same file), `.thetisOnlyToolStripMenuItem_Click()` (same file)
- **`.disableCpuVoltsUsage()`** — L20863 — `private void disableCpuVoltsUsage()`
  Called by: `.CpuUsage()` (same file)
- **`.isBitSet()`** — L21127 — `private static bool isBitSet(int n, int pos)`
  Called by: `.checkSeqErrors()` (same file)
- **`.ShowSEQLog()`** — L21132 — `public void ShowSEQLog()`
  Shows seqlog.
  Called by: `.toolStripStatusLabel_SeqWarning_Click()` (same file)
- **`.RXPreampOffset()`** — L21150 — `public float RXPreampOffset(int rx)`
  Called by: `.UpdateRX1DisplayOffsets()` (same file), `.UpdateRX2DisplayOffsets()` (same file), `.RXOffset()` (same file)
- **`.RXCalibrationOffset()`** — L21183 — `public float RXCalibrationOffset(int rx)`
  Called by: `.UpdateRX1DisplayOffsets()` (same file), `.UpdateRX2DisplayOffsets()` (same file), `.RXOffset()` (same file)
- **`.RXOffset()`** — L21201 — `public float RXOffset(int rx)`
  Called by: `.CATReadSigStrength()` (same file), `.CATReadAvgStrength()` (same file), `.RXPBsnr()` (same file), `.UpdatePeakText()` (same file), `.UpdateMultimeter()` (same file), `.UpdateRX2Multimeter()` (same file) — and 6 more
- **`.RXPBsnr()`** — L21259 — `public double RXPBsnr(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.spectralCalculations()`** — L21317 — `private SpectralResult spectralCalculations(int rx, double signal)`
  Called by: `.RXPBsnr()` (same file), `.UpdatePeakText()` (same file), `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.checkOverloadsAndSync()`** — L21471 — `private async void checkOverloadsAndSync()`
  Called by: `.pollOverloadSyncSeqErr()` (same file)
- **`.keep_att_entries_for_band()`** — L21598 — `private void keep_att_entries_for_band(Stack<HistoricAttenuatorReading> readings_stack, Band target_band)`
  Called by: `.handleOverload()` (same file)
- **`.handleOverload()`** — L21616 — `private void handleOverload()`
  Called by: `.checkOverloadsAndSync()` (same file)
- **`.pollOverloadSyncSeqErr()`** — L21917 — `private async void pollOverloadSyncSeqErr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkSeqErrors()`** — L21972 — `private void checkSeqErrors()`
  Called by: `.pollOverloadSyncSeqErr()` (same file)
- **`.UpdatePeakText()`** — L22046 — `private void UpdatePeakText()`
  Updates peak text.
  Called by: `.timer_peak_text_Tick()` (same file)
- **`.HzInNPixels()`** — L22228 — `private int HzInNPixels(int nPixelCount, int rx)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.getLowHighForRXn()`** — L22237 — `private void getLowHighForRXn(int rx, out int low, out int high, bool bIncludeRitXit = true)`
  Returns low high for rxn.
  Called by: `.HzInNPixels()` (same file), `.PixelToHz()` (same file), `.HzToPixel()` (same file)
- **`.PixelToHz()`** — L22331 — `private float PixelToHz(float x)`
  Called by: `.UpdatePeakText()` (same file), `.getFrequencyAtPixel()` (same file), `.pnlDisplay_MouseDown()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.HzToPixel()`** — L22345 — `private int HzToPixel(float freq)`
  Called by: `.getFilterEdgesInPixels()` (same file), `.pnlDisplay_MouseDown()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.PixelToDb()`** — L22377 — `private float PixelToDb(float y)`
  Called by: `.pnlDisplay_DoubleClick()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.PixelToRx2Db()`** — L22382 — `private float PixelToRx2Db(float y)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.WaterfallPixelToTime()`** — L22394 — `private float WaterfallPixelToTime(float y, int rx)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.measureStringFromCache()`** — L22469 — `private SizeF measureStringFromCache(string str, Font font, int width, StringFormat format, Graphics g)`
  Called by: `.GetVFOCharWidth()` (same file), `.GetVFOSubCharWidth()` (same file), `.getMeterPixelPosAndDrawScales()` (same file)
- **`.getMeterPixelPosAndDrawScales()`** — L22479 — `private void getMeterPixelPosAndDrawScales(int rx, Graphics g, int H, int W, double num, out int pixel_x, out int pixel_x_swr, int nStringOffsetY, bool bDrawMarkers)`
  Returns meter pixel pos and draw scales.
  Called by: `.picMultiMeterDigital_Paint()` (same file), `.picRX2Meter_Paint()` (same file)
- **`.storeRX1SignalPixels_X()`** — L23595 — `private void storeRX1SignalPixels_X(float x)`
  Called by: `.picMultiMeterDigital_Paint()` (same file)
- **`.storeRX2SignalPixels_X()`** — L23612 — `private void storeRX2SignalPixels_X(float x)`
  Called by: `.picRX2Meter_Paint()` (same file)
- **`.clearRXSignalPixels()`** — L23630 — `private void clearRXSignalPixels(int rx)`
  Called by: `.picMultiMeterDigital_Paint()` (same file), `.picRX2Meter_Paint()` (same file), `.ResetMultiMeterPeak()` (same file), `.ResetRX2MeterPeak()` (same file), `.OnPowerChangeHander()` (same file), `.OnBandChangeHandler()` (same file) — and 2 more
- **`.picMultiMeterDigital_Paint()`** — L23666 — `private void picMultiMeterDigital_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picMultiMeterDigital` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picRX2Meter_Paint()`** — L23971 — `private void picRX2Meter_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picRX2Meter` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ResetMultiMeterPeak()`** — L24173 — `private void ResetMultiMeterPeak()`
  Resets multi meter peak.
  Called by: `.UIMOXChangedTrue()` (same file), `.UIMOXChangedFalse()` (same file), `.comboMeterRXMode_SelectedIndexChanged()` (same file), `.comboMeterTXMode_SelectedIndexChanged()` (same file)
- **`.ResetRX2MeterPeak()`** — L24180 — `private void ResetRX2MeterPeak()`
  Resets rx2 meter peak.
  Called by: `.comboRX2MeterMode_SelectedIndexChanged()` (same file)
- **`.panelVFOAHover_Paint()`** — L24187 — `private void panelVFOAHover_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `panelVFOAHover` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panelVFOBHover_Paint()`** — L24219 — `private void panelVFOBHover_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `panelVFOBHover` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.resetWDSPdisplayBuffers()`** — L24266 — `private void resetWDSPdisplayBuffers(int rx, bool tx)`
  Called by: `.RunDisplay()` (same file)
- **`.RunDisplay()`** — L24315 — `unsafe private void RunDisplay()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateMultimeter()`** — L24744 — `private async void UpdateMultimeter()`
  Updates multimeter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2Multimeter()`** — L24904 — `private async void UpdateRX2Multimeter()`
  Updates rx2 multimeter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeHermesDCVoltage()`** — L24955 — `public float computeHermesDCVoltage()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.readMKIIPAVoltsAmps()`** — L24974 — `private async void readMKIIPAVoltsAmps()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeMKIIPAVoltsAmps()`** — L25121 — `private void computeMKIIPAVoltsAmps()`
  Called by: `.timer_cpu_volts_meter_Tick()` (same file)
- **`.convertToVolts()`** — L25136 — `private float convertToVolts(float IOreading)`
  Called by: `.readMKIIPAVoltsAmps()` (same file), `.computeMKIIPAVoltsAmps()` (same file)
- **`.convertToAmps()`** — L25166 — `private float convertToAmps(float IOreading)`
  Called by: `.readMKIIPAVoltsAmps()` (same file), `.computeMKIIPAVoltsAmps()` (same file)
- **`.computeRefPower()`** — L25195 — `public float computeRefPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeAlexFwdPower()`** — L25281 — `public float computeAlexFwdPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeExciterPower()`** — L25353 — `public float computeExciterPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeOrionMkIIExciterPower()`** — L25412 — `public float computeOrionMkIIExciterPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeOrionExciterPower()`** — L25471 — `public float computeOrionExciterPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeANANExciterPower()`** — L25530 — `public float computeANANExciterPower()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeHermesLiteTemp()`** — L25591 — `public void computeHermesLiteTemp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeHermesLitePAAmps()`** — L25606 — `public void computeHermesLitePAAmps()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateSQL()`** — L25622 — `private async void UpdateSQL()`
  Updates sql.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2SQL()`** — L25639 — `private async void UpdateRX2SQL()`
  Updates rx2 sql.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateNoiseGate()`** — L25656 — `private async void UpdateNoiseGate()`
  Updates noise gate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIOBoardAerialPorts()`** — L25670 — `public void SetIOBoardAerialPorts(int rx_only_ant, int rx_ant, int tx_ant, bool tx)`
  Sets ioboard aerial ports.
  Called by: `.modifyXVTRantenna()` (same file)
- **`.SetI2CPollingPause()`** — L25697 — `public void SetI2CPollingPause( bool pause )`
  Sets i2 cpolling pause.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AutoTuningHL2()`** — L25728 — `bool AutoTuningHL2(ProtocolEvent protocolEvent)`
  Called by: `.UpdateIOBoard()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.UpdateIOBoard()`** — L25837 — `private async void UpdateIOBoard()`
  Updates ioboard.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateVOX()`** — L26003 — `private async void UpdateVOX()`
  Updates vox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getFallbackPTTModeAfterTCIRelease()`** — L26025 — `private PTTMode getFallbackPTTModeAfterTCIRelease(DSPMode tx_mode, bool mic_ptt, bool cw_ptt, bool cat_ptt, bool vox_ptt)`
  Returns fallback pttmode after tcirelease.
  Called by: `.PollPTT()` (same file)
- **`.PollPTT()`** — L26059 — `private async void PollPTT()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PollCW()`** — L26224 — `private async void PollCW()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cwAutoModeTick()`** — L26256 — `private void cwAutoModeTick(object o)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.update_for_auto_mode_return()`** — L26295 — `private void update_for_auto_mode_return(bool enabled)`
  Called by: `.PollCW()` (same file)
- **`.UpdatePreamps()`** — L26379 — `private void UpdatePreamps()`
  Updates preamps.
  Called by: `.comboPreamp_SelectedIndexChanged()` (same file), `.comboRX2Preamp_SelectedIndexChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.txtVFOAFreq_LostFocus()` (same file)
- **`.PollTXInhibit()`** — L26445 — `private async void PollTXInhibit()`
  bool audio_amp_mute;
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PollN1MMPacket()`** — L26486 — `private async void PollN1MMPacket()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.n1mm_delay_Elapsed()`** — L26520 — `private void n1mm_delay_Elapsed(object sender, ElapsedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleXml()`** — L26556 — `void HandleXml(string str)`
  Handles xml.
  Called by: `.PollN1MMPacket()` (same file)
- **`.ToggleFocusMasterTimer()`** — L26567 — `private void ToggleFocusMasterTimer()`
  Toggles focus master timer.
  Called by: `.gmh_MouseUp()` (same file), `.n1mm_delay_Elapsed()` (same file), `.Console_KeyUp()` (same file), `.Console_MouseWheel()` (same file), `.SetFocusMaster()` (same file)
- **`.PollPAPWR()`** — L26577 — `private async void PollPAPWR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkAntennaWarning()`** — L26773 — `private void checkAntennaWarning()`
  Called by: `.PollPAPWR()` (same file)
- **`.SWRScale()`** — L26795 — `private double SWRScale(double ref_pow)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.timer_cpu_volts_meter_Tick()`** — L26808 — `private void timer_cpu_volts_meter_Tick(object sender, System.EventArgs e)`
  WinForms event handler: runs when `timer_cpu_volts_meter` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.timer_peak_text_Tick()`** — L26875 — `private void timer_peak_text_Tick(object sender, System.EventArgs e)`
  WinForms event handler: runs when `timer_peak_text` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.timer_clock_Tick()`** — L26880 — `private void timer_clock_Tick(object sender, System.EventArgs e)`
  WinForms event handler: runs when `timer_clock` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_KeyPress()`** — L26910 — `private void Console_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `Console` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_KeyUp()`** — L26920 — `private void Console_KeyUp(object sender, System.Windows.Forms.KeyEventArgs e)`
  WinForms event handler: runs when `Console` receives a key-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.enableOutsideSpectral()`** — L26933 — `private void enableOutsideSpectral()`
  Called by: `.Console_KeyDown()` (same file), `.OnMouseWheelChanged()` (same file)
- **`.restoreOutsideSpectral()`** — L26938 — `private void restoreOutsideSpectral()`
  Called by: `.Console_KeyDown()` (same file), `.OnMouseWheelChanged()` (same file)
- **`.Console_KeyDown()`** — L26946 — `private void Console_KeyDown(object sender, System.Windows.Forms.KeyEventArgs e)`
  WinForms event handler: runs when `Console` receives a key-down.
  Called by: `.ProcessDialogKey()` (same file)
- **`.setupLegacyMeterThreads()`** — L27791 — `private void setupLegacyMeterThreads(int rx)`
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.chkPower_CheckedChanged()`** — L27823 — `private void chkPower_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPower` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateAAudioMixerStates()`** — L28276 — `unsafe public void UpdateAAudioMixerStates()`
  MW0LGE [2.9.0.8] re-implemented by Warren
  Called by: `.chkPower_CheckedChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.comboDisplayMode_SelectedIndexChanged()`** — L28429 — `public void comboDisplayMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDisplayMode` selection changes.
  Called by: `.SetupDisplayEngine()` (same file), `.chkPower_CheckedChanged()` (same file)
- **`.chkBIN_CheckedChanged()`** — L28648 — `private void chkBIN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkBIN` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAGC_SelectedIndexChanged()`** — L28662 — `private void comboAGC_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAGC` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.run_console_closing_handlers_async()`** — L28767 — `private Task run_console_closing_handlers_async()`
  Called by: `.Console_Closing()` (same file)
- **`.Console_Closing()`** — L28781 — `private void Console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `Console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getErrorLogSize()`** — L29015 — `private long getErrorLogSize()`
  Returns error log size.
  Called by: `.Dispose()` (same file)
- **`.shutdownLogStringToPath()`** — L29025 — `private void shutdownLogStringToPath(string entry)`
  Called by: `.Dispose()` (same file), `.ExitConsole()` (same file), `.Console_Closing()` (same file)
- **`.removeShutdownLog()`** — L29039 — `private void removeShutdownLog()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboPreamp_SelectedIndexChanged()`** — L29048 — `private void comboPreamp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboPreamp` selection changes.
  Called by: `.SetComboPreampForHPSDR()` (same file)
- **`.comboRX2Preamp_SelectedIndexChanged()`** — L29115 — `private void comboRX2Preamp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2Preamp` selection changes.
  Called by: `.SetComboPreampForHPSDR()` (same file)
- **`.chkMUT_CheckedChanged()`** — L29191 — `private void chkMUT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMUT` checked state changes.
  Called by: `.InitConsole()` (same file), `.SetRX1Mode()` (same file)
- **`.ModelIsHPSDRorHermes()`** — L29235 — `public bool ModelIsHPSDRorHermes()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDriveSliderUpdateTimerTick()`** — L29251 — `private void OnDriveSliderUpdateTimerTick(object sender, ElapsedEventArgs e)`
  Handles/raises the drive slider update timer tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTuneSliderUpdateTimerTick()`** — L29255 — `private void OnTuneSliderUpdateTimerTick(object sender, ElapsedEventArgs e)`
  Handles/raises the tune slider update timer tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateDriveLabel()`** — L29261 — `public void UpdateDriveLabel(bool bShowLimitValue, System.EventArgs e)`
  Updates drive label.
  Called by: `.OnDriveSliderUpdateTimerTick()` (same file), `.ptbPWR_MouseUp()` (same file), `.ptbPWR_Scroll()` (same file)
- **`.ptbPWR_MouseUp()`** — L29363 — `private void ptbPWR_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbPWR` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbPWR_Scroll()`** — L29368 — `private void ptbPWR_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbPWR` is scrolled.
  Called by: `.InitConsole()` (same file), `.checkOverloadsAndSync()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.ptbAF_Scroll()`** — L29407 — `private void ptbAF_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbAF` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file), `.chkMUT_CheckedChanged()` (same file), `.chkMON_CheckedChanged()` (same file), `.AudioMOXChanged()` (same file)
- **`.ptbRF_Scroll()`** — L29455 — `private void ptbRF_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRF` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.chkMicMute_CheckedChanged()`** — L29504 — `private void chkMicMute_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMicMute` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbMic_Scroll()`** — L29511 — `private void ptbMic_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbMic` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file), `.chkMicMute_CheckedChanged()` (same file), `.SetRX1Mode()` (same file), `.radModeButton_CheckedChanged()` (same file), `.SetRX2Mode()` (same file) — and 3 more
- **`.setAudioMicGain()`** — L29565 — `private void setAudioMicGain(double gain_db)`
  Sets audio mic gain.
  Called by: `.ptbMic_Scroll()` (same file), `.ptbFMMic_Scroll()` (same file)
- **`.ptbCWSpeed_Scroll()`** — L29595 — `private void ptbCWSpeed_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWSpeed` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file)
- **`.chkVOX_CheckedChanged()`** — L29620 — `private void chkVOX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVOX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picSquelch_Paint()`** — L29640 — `private void picSquelch_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picSquelch` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNoiseGate_CheckedChanged()`** — L29651 — `private void chkNoiseGate_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkNoiseGate` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbVACRXGain_Scroll()`** — L29661 — `private void ptbVACRXGain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbVACRXGain` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkVAC2_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.chkVFOATX_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.ptbVACTXGain_Scroll()`** — L29686 — `private void ptbVACTXGain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbVACTXGain` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkVAC2_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.chkVFOATX_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.ptbVOX_Scroll()`** — L29715 — `private void ptbVOX_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbVOX` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file)
- **`.picVOX_Paint()`** — L29726 — `unsafe private void picVOX_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picVOX` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbNoiseGate_Scroll()`** — L29739 — `private void ptbNoiseGate_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbNoiseGate` is scrolled.
  Called by: `.InitConsole()` (same file), `.pnlDisplay_DoubleClick()` (same file)
- **`.picNoiseGate_Paint()`** — L29749 — `private void picNoiseGate_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picNoiseGate` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.WheelTune_MouseDown()`** — L29760 — `private void WheelTune_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `WheelTune` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMON_CheckedChanged()`** — L29766 — `private void chkMON_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMON` checked state changes.
  Called by: `.chkVFOATX_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.AudioMOXChanged()`** — L29794 — `private void AudioMOXChanged(bool tx)`
  Called by: `.PollCW()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.HdwMOXChanged()`** — L29810 — `private void HdwMOXChanged(bool tx, double freq)`
  Called by: `.chkMOX_CheckedChanged2()` (same file)
- **`.UIMOXChangedTrue()`** — L29902 — `private void UIMOXChangedTrue()`
  Called by: `.chkMOX_CheckedChanged2()` (same file)
- **`.UIMOXChangedFalse()`** — L29937 — `private void UIMOXChangedFalse()`
  Called by: `.chkMOX_CheckedChanged2()` (same file)
- **`.updateAttNudsCombos()`** — L29989 — `private void updateAttNudsCombos()`
  Called by: `.UIMOXChangedTrue()` (same file), `.UIMOXChangedFalse()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.ExpandDisplay()` (same file)
- **`.chkMOX_CheckedChanged2()`** — L30119 — `private void chkMOX_CheckedChanged2(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkMOX_Click()`** — L30495 — `private void chkMOX_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMOX` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboMeterRXMode_SelectedIndexChanged()`** — L30516 — `private void comboMeterRXMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboMeterRXMode` selection changes.
  Called by: `.UIMOXChangedFalse()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.comboMeterTXMode_SelectedIndexChanged()`** — L30566 — `private void comboMeterTXMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboMeterTXMode` selection changes.
  Called by: `.UIMOXChangedTrue()` (same file), `.UIMOXChangedFalse()` (same file), `.chkTUN_CheckedChanged()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.isMeterModeAvailableWhenTune()`** — L30644 — `private bool isMeterModeAvailableWhenTune(MeterTXMode meterMode)`
  Called by: `.comboMeterTXMode_SelectedIndexChanged()` (same file)
- **`.chkDisplayAVG_CheckedChanged()`** — L30679 — `private void chkDisplayAVG_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDisplayAVG` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.chkDisplayPeak_CheckedChanged()`** — L30703 — `private void chkDisplayPeak_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDisplayPeak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateVFOFreqs()`** — L30727 — `private void updateVFOFreqs(bool tx, bool isTune = false)`
  Called by: `.HdwMOXChanged()` (same file), `.chkTUN_CheckedChanged()` (same file), `.chkXIT_CheckedChanged()` (same file), `.udXIT_ValueChanged()` (same file), `.chkRX2SR_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.chkTUN_CheckedChanged()`** — L30794 — `private async void chkTUN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTUN` checked state changes.
  Called by: `.AutoTuningHL2()` (same file), `.chk2TONE_CheckedChanged()` (same file)
- **`.SetupTunePulse()`** — L30998 — `public void SetupTunePulse()`
  Setups tune pulse.
  Called by: `.chkTUN_CheckedChanged()` (same file)
- **`.ATUTune()`** — L31034 — `private async void ATUTune(CancellationToken t)`
  Called by: `.chkTUN_CheckedChanged()` (same file)
- **`.comboTuneMode_SelectedIndexChanged()`** — L31058 — `private void comboTuneMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboTuneMode` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.HideFocus()`** — L31064 — `private void HideFocus(object sender, EventArgs e)`
  Hides focus.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.textbox_GotFocus()`** — L31069 — `private void textbox_GotFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `textbox` gains focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.textbox_LostFocus()`** — L31074 — `private void textbox_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `textbox` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.combo_OpenDropDown()`** — L31079 — `private void combo_OpenDropDown(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.combo_CloseDropDown()`** — L31084 — `private void combo_CloseDropDown(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkVFOLock_CheckedChanged()`** — L31089 — `private void chkVFOLock_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOLock` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOBLock_CheckedChanged()`** — L31095 — `private void chkVFOBLock_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVFOBLock` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.repopulateForms()`** — L31100 — `private void repopulateForms()`
  Called by: `.SetCATBand()` (same file), `.btnBandVHF_Click()` (same file), `.btnBandHF_Click()` (same file), `.btnBandGEN_Click()` (same file), `.radBand_CheckedChanged()` (same file)
- **`.BandPanelVisible()`** — L31106 — `public void BandPanelVisible(bool all_hidden = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModePanelVisible()`** — L31124 — `public void ModePanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOAVisible()`** — L31136 — `public void VFOAVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOBVisible()`** — L31140 — `public void VFOBVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOSyncVisible()`** — L31144 — `public void VFOSyncVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterPanelVisible()`** — L31148 — `public void FilterPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PowerRxPanelVisible()`** — L31158 — `public void PowerRxPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MonTunePanelVisible()`** — L31162 — `public void MonTunePanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SplitRitVacPanelVisible()`** — L31166 — `public void SplitRitVacPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NoiseMnfPanelVisible()`** — L31175 — `public void NoiseMnfPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MicCompVoxPanelVisible()`** — L31191 — `public void MicCompVoxPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisplayControlsPanelVisible()`** — L31199 — `public void DisplayControlsPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExtendPanelDisplaySizeRight()`** — L31204 — `public void ExtendPanelDisplaySizeRight(bool expand)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExtendPanelDisplaySizeTop()`** — L31223 — `public void ExtendPanelDisplaySizeTop(bool expand)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setBandPanelVisible()`** — L31248 — `private void setBandPanelVisible(bool gen, bool hf, bool vhf, bool force = false)`
  Sets band panel visible.
  Called by: `.SetRX1Band()` (same file), `.BandPanelVisible()` (same file), `.btnBandVHF_Click()` (same file), `.btnBandHF_Click()` (same file), `.btnBandGEN_Click()` (same file), `.OnBandChangeHandler()` (same file)
- **`.btnBandVHF_Click()`** — L31305 — `private void btnBandVHF_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnBandVHF` is clicked.
  Called by: `.ExpandDisplay()` (same file)
- **`.btnBandHF_Click()`** — L31318 — `private void btnBandHF_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnBandHF` is clicked.
  Called by: `.ExpandDisplay()` (same file)
- **`.btnBandGEN_Click()`** — L31331 — `private void btnBandGEN_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnBandGEN` is clicked.
  Called by: `.ExpandDisplay()` (same file)
- **`.udFilterLow_LostFocus()`** — L31375 — `private void udFilterLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udFilterLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udFilterHigh_LostFocus()`** — L31381 — `private void udFilterHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udFilterHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterLow_LostFocus()`** — L31387 — `private void udTXFilterLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXFilterLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterHigh_LostFocus()`** — L31393 — `private void udTXFilterHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXFilterHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2FilterLow_LostFocus()`** — L31399 — `private void udRX2FilterLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2FilterLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2FilterHigh_LostFocus()`** — L31405 — `private void udRX2FilterHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2FilterHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRIT_LostFocus()`** — L31413 — `private void udRIT_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udRIT` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udXIT_LostFocus()`** — L31418 — `private void udXIT_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udXIT` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnChangeTuneStepSmaller_Click()`** — L31423 — `private void btnChangeTuneStepSmaller_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnChangeTuneStepSmaller` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnChangeTuneStepLarger_Click()`** — L31428 — `private void btnChangeTuneStepLarger_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnChangeTuneStepLarger` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboTXProfile_SelectedIndexChanged()`** — L31433 — `private void comboTXProfile_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.comboDigTXProfile_SelectedIndexChanged()`** — L31446 — `private void comboDigTXProfile_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDigTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.comboFMTXProfile_SelectedIndexChanged()`** — L31456 — `private void comboFMTXProfile_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFMTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.LoadedTXProfile()`** — L31468 — `public void LoadedTXProfile()`
  MW0LGE_21j used by setup form whenever a TX profile is loaded When a digimode is selected, a number of settings are disabled. These are restored if leaving a digimode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboAMTXProfile_SelectedIndexChanged()`** — L31499 — `private void comboAMTXProfile_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboAMTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.chkShowTXFilter_CheckedChanged()`** — L31509 — `private void chkShowTXFilter_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowTXFilter` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVACStereo_CheckedChanged()`** — L31516 — `private void chkVACStereo_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVACStereo` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWIambic_CheckedChanged()`** — L31533 — `private void chkCWIambic_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWIambic` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWSidetone_CheckedChanged()`** — L31538 — `private void chkCWSidetone_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWSidetone` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCWPitch_ValueChanged()`** — L31543 — `private void udCWPitch_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWPitch` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboVACSampleRate_SelectedIndexChanged()`** — L31549 — `private void comboVACSampleRate_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboVACSampleRate` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkX2TR_CheckedChanged()`** — L31562 — `private void chkX2TR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkX2TR` checked state changes.
  Called by: `.GetState()` (same file)
- **`.chkShowTXCWFreq_CheckedChanged()`** — L31583 — `private void chkShowTXCWFreq_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowTXCWFreq` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowCWZero_CheckedChanged()`** — L31588 — `private void chkShowCWZero_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowCWZero` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCWBreakInDelay_ValueChanged()`** — L31595 — `private void udCWBreakInDelay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWBreakInDelay` value changes.
  Called by: `.udCWBreakInDelay_LostFocus()` (same file)
- **`.udCWBreakInDelay_LostFocus()`** — L31602 — `private void udCWBreakInDelay_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWBreakInDelay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWAPFEnabled_CheckedChanged()`** — L31607 — `private void chkCWAPFEnabled_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWAPFEnabled` checked state changes.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.ptbCWAPFFreq_Scroll()`** — L31648 — `private void ptbCWAPFFreq_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWAPFFreq` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.ptbCWAPFBandwidth_Scroll()`** — L31666 — `private void ptbCWAPFBandwidth_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWAPFBandwidth` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.ptbCWAPFGain_Scroll()`** — L31685 — `private void ptbCWAPFGain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWAPFGain` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.DisableDAX()`** — L31706 — `public void DisableDAX()`
  Disables dax.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableDAX()`** — L31715 — `public void EnableDAX()`
  Enables dax.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkVAC1_CheckedChanged()`** — L31722 — `private void chkVAC1_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2_CheckedChanged()`** — L31770 — `private void chkVAC2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRXEQ_CheckedChanged()`** — L31817 — `private void chkRXEQ_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRXEQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXEQ_CheckedChanged()`** — L31827 — `private void chkTXEQ_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTXEQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TuneHitTest()`** — L31856 — `private TuneLocation TuneHitTest(int x, int y)`
  Called by: `.Console_MouseWheel()` (same file)
- **`.Console_MouseWheel()`** — L31898 — `private void Console_MouseWheel(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `Console` receives a mouse wheel event.
  Called by: `.Console_KeyDown()` (same file), `.OnMouseWheelChanged()` (same file)
- **`.SnapTune()`** — L32073 — `public double SnapTune(double freq_mhz, int step_size_hz, int num_steps)`
  Calculates a "Snapped" frequency that lies on an integer multiple of the Tune Step.
  Called by: `.Console_MouseWheel()` (same file)
- **`.txtVFOAFreq_LostFocus()`** — L32111 — `private void txtVFOAFreq_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` loses focus.
  Called by: `.CalibratePAGain()` (same file), `.LowPowerPASweep()` (same file), `.SetupForHPSDRModel()` (same file), `.SetWavePlayback()` (same file), `.VFOAUpdate()` (same file), `.chkPower_CheckedChanged()` (same file) — and 15 more
- **`.setupModifyXVTRantennaArray()`** — L32755 — `private void setupModifyXVTRantennaArray()`
  Called by: `.InitConsole()` (same file)
- **`.modifyXVTRantenna()`** — L32766 — `private void modifyXVTRantenna(int rx, double freq, int rx_xvtr_index)`
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.undoXVTRantennaModify()`** — L32807 — `private void undoXVTRantennaModify(int rx)`
  Called by: `.Console_Closing()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.modifyXVTRantenna()` (same file)
- **`.getTXBandWhenExtended()`** — L32821 — `private Band getTXBandWhenExtended(Band b, double frequency = -1)`
  Returns txband when extended.
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.txtVFOAFreq_KeyPress()`** — L32860 — `private void txtVFOAFreq_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAFreq_MouseMove()`** — L32890 — `private void txtVFOAFreq_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` receives mouse movement.
  Called by: `.panelVFOAHover_MouseMove()` (same file), `.txtVFOALSD_MouseMove()` (same file), `.txtVFOAMSD_MouseMove()` (same file)
- **`.txtVFOAFreq_MouseLeave()`** — L32934 — `private void txtVFOAFreq_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` is left by the mouse.
  Called by: `.txtVFOAMSD_MouseLeave()` (same file)
- **`.txtVFOABand_LostFocus()`** — L32940 — `private void txtVFOABand_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOABand` loses focus.
  Called by: `.VFOASubUpdate()` (same file), `.UpdateVFOASub()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.txtVFOABand_KeyPress()`** — L33084 — `private void txtVFOABand_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `txtVFOABand` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBFreq_LostFocus()`** — L33121 — `private void txtVFOBFreq_LostFocus(object sender, System.EventArgs e)`
  txtVFOBFreq
  Called by: `.SetWavePlayback()` (same file), `.VFOBUpdate()` (same file), `.updateVFOFreqs()` (same file), `.chkX2TR_CheckedChanged()` (same file), `.zoomToBandBandwidth()` (same file), `.ptbDisplayZoom_Scroll()` (same file) — and 12 more
- **`.txtVFOBFreq_KeyPress()`** — L33780 — `private void txtVFOBFreq_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `txtVFOBFreq` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBFreq_MouseMove()`** — L33810 — `private void txtVFOBFreq_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBFreq` receives mouse movement.
  Called by: `.panelVFOBHover_MouseMove()` (same file), `.txtVFOBMSD_MouseMove()` (same file), `.txtVFOBLSD_MouseMove()` (same file)
- **`.txtVFOBFreq_MouseLeave()`** — L33854 — `private void txtVFOBFreq_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOBFreq` is left by the mouse.
  Called by: `.txtVFOBMSD_MouseLeave()` (same file)
- **`.panelVFOAHover_MouseMove()`** — L33860 — `private void panelVFOAHover_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `panelVFOAHover` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panelVFOBHover_MouseMove()`** — L33871 — `private void panelVFOBHover_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `panelVFOBHover` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOALSD_MouseDown()`** — L33882 — `private void txtVFOALSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOALSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOALSD_MouseMove()`** — L33890 — `private void txtVFOALSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOALSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAMSD_MouseDown()`** — L33901 — `private void txtVFOAMSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOAMSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAMSD_MouseMove()`** — L33909 — `private void txtVFOAMSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOAMSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAMSD_MouseLeave()`** — L33915 — `private void txtVFOAMSD_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOAMSD` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBMSD_MouseDown()`** — L33920 — `private void txtVFOBMSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBMSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBMSD_MouseLeave()`** — L33928 — `private void txtVFOBMSD_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOBMSD` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBMSD_MouseMove()`** — L33933 — `private void txtVFOBMSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBMSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBLSD_MouseDown()`** — L33939 — `private void txtVFOBLSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBLSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBLSD_MouseMove()`** — L33947 — `private void txtVFOBLSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBLSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.overRX()`** — L34009 — `private bool overRX(int x, int y, int rx, bool bIgnorePanafallWaterfall = true)`
  Called by: `.UpdatePeakText()` (same file), `.pnlDisplay_MouseDown()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.notchMouseWheel()`** — L34103 — `private void notchMouseWheel(int wheelDelta)`
  Called by: `.Console_MouseWheel()` (same file)
- **`.CurrentDSPhasTwoSidebands()`** — L34128 — `public bool CurrentDSPhasTwoSidebands(int rx, bool tx = false)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.agcCalOffset()`** — L34150 — `private float agcCalOffset(int rx)`
  Called by: `.setAGCThresholdPoint()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.getFilterEdgesInPixels()`** — L34181 — `private void getFilterEdgesInPixels(MouseEventArgs e, ref int low_x, ref int high_x, ref int vfoa_sub_x, ref int vfoa_sub_low_x, ref int vfoa_sub_high_x)`
  Returns filter edges in pixels.
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.dragWholeFilter()`** — L34240 — `private void dragWholeFilter(MouseEventArgs e)`
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.adjustForSnapClickTuning()`** — L34274 — `private double adjustForSnapClickTuning(int rx, double freq)`
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.getFrequencyAtPixel()`** — L34303 — `private double getFrequencyAtPixel(int x, int nRX)`
  Returns frequency at pixel.
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.ShowNotchPopup()`** — L34345 — `public void ShowNotchPopup(int x, int y, MNotch notch, int min_width, int max_width, bool on_top, int notch_index = -1)`
  Shows notch popup.
  Called by: `.pnlDisplay_MouseUp()` (same file)
- **`.ptbDisplayPan_Scroll()`** — L34365 — `private void ptbDisplayPan_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbDisplayPan` is scrolled.
  Called by: `.btnDisplayPanCenter_Click()` (same file)
- **`.btnDisplayPanCenter_Click()`** — L34379 — `private void btnDisplayPanCenter_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnDisplayPanCenter` is clicked.
  Called by: `.PanCentre()` (same file), `.SetPanAdjust()` (same file)
- **`.zoomToBandBandwidth()`** — L34404 — `private bool zoomToBandBandwidth(Band b, int rx)`
  Called by: `.ZoomToBand()` (same file)
- **`.ptbDisplayZoom_Scroll()`** — L34492 — `private void ptbDisplayZoom_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbDisplayZoom` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.radDisplayZoom05_CheckedChanged()`** — L34567 — `private void radDisplayZoom05_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom05` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom05()`** — L34574 — `private void displayZoom05()`
  Called by: `.radDisplayZoom05_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radDisplayZoom1x_CheckedChanged()`** — L34580 — `private void radDisplayZoom1x_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom1x` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom1()`** — L34587 — `private void displayZoom1()`
  Called by: `.radDisplayZoom1x_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radDisplayZoom2x_CheckedChanged()`** — L34593 — `private void radDisplayZoom2x_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom2x` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom2()`** — L34600 — `private void displayZoom2()`
  Called by: `.radDisplayZoom2x_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radDisplayZoom4x_CheckedChanged()`** — L34605 — `private void radDisplayZoom4x_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom4x` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom4()`** — L34612 — `private void displayZoom4()`
  Called by: `.radDisplayZoom4x_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radBand160_Click()`** — L34624 — `private void radBand160_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand160` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand80_Click()`** — L34630 — `private void radBand80_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand80` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand60_Click()`** — L34636 — `private void radBand60_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand60` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand40_Click()`** — L34642 — `private void radBand40_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand40` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand30_Click()`** — L34648 — `private void radBand30_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand30` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand20_Click()`** — L34654 — `private void radBand20_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand20` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand17_Click()`** — L34660 — `private void radBand17_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand17` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand15_Click()`** — L34666 — `private void radBand15_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand15` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand12_Click()`** — L34672 — `private void radBand12_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand12` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand10_Click()`** — L34678 — `private void radBand10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand10` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand6_Click()`** — L34684 — `private void radBand6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand6` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand2_Click()`** — L34690 — `private void radBand2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand2` is clicked.
  Called by: `.SetCATBand()` (same file)
- **`.radBandWWV_Click()`** — L34696 — `private void radBandWWV_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandWWV` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBandGEN_Click()`** — L34702 — `private void radBandGEN_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBandVHF_Click()`** — L34707 — `private void radBandVHF_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandVHF` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setVFOAFreqNoUpdate()`** — L34722 — `private void setVFOAFreqNoUpdate(double freq)`
  Sets vfoafreq no update.
  Called by: `.SetRX1Mode()` (same file)
- **`.setVFOBFreqNoUpdate()`** — L34727 — `private void setVFOBFreqNoUpdate(double freq)`
  Sets vfobfreq no update.
  Called by: `.SetRX2Mode()` (same file)
- **`.initControlBackColours()`** — L34735 — `private void initControlBackColours(Control c)`
  MW0LGE_21d used to default colours of all button+radio controls, and inside other panels or groups an issue was noticed where text change colour on buttons that had be selected/deseleted
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRX1Mode()`** — L34761 — `private void SetRX1Mode(DSPMode new_mode)`
  Sets rx1 mode.
  Called by: `.radModeButton_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.radModeButton_CheckedChanged()`** — L35465 — `private void radModeButton_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radModeButton` checked state changes.
  Called by: `.selectModes()` (same file)
- **`.SetRX1Filter()`** — L35557 — `public void SetRX1Filter(Filter new_filter)`
  Sets rx1 filter.
  Called by: `.radFilter_CheckedChanged()` (same file)
- **`.filterAndDspModeValid()`** — L35684 — `private bool filterAndDspModeValid(int rx)`
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRX2Filters()` (same file), `.UpdateRX1FilterNames()` (same file), `.UpdateRX2FilterNames()` (same file), `.SetRX1Filter()` (same file), `.SetRX2Filter()` (same file) — and 2 more
- **`.radRX2Filter_CheckedChanged()`** — L35696 — `private void radRX2Filter_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX2Filter` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MatchTXFilterToRXFilter()`** — L35757 — `public void MatchTXFilterToRXFilter()`
  Called by: `.radRX2Filter_CheckedChanged()` (same file), `.radFilter_CheckedChanged()` (same file), `.radFilter_rx1_MouseUp()` (same file), `.radFilter_rx2_MouseUp()` (same file)
- **`.radFilter_CheckedChanged()`** — L35779 — `private void radFilter_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFilter` checked state changes.
  Called by: `.selectFilters()` (same file)
- **`.udFilterLow_ValueChanged()`** — L35858 — `private void udFilterLow_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udFilterLow` value changes.
  Called by: `.udFilterLow_LostFocus()` (same file)
- **`.udFilterHigh_ValueChanged()`** — L35879 — `private void udFilterHigh_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udFilterHigh` value changes.
  Called by: `.udFilterHigh_LostFocus()` (same file)
- **`.ConstrainFilter()`** — L35900 — `public bool ConstrainFilter(ref int nNewLow, ref int nNewHigh, int rx, bool filterShift = false)`
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRX2Filters()` (same file), `.ptbFilterShift_Scroll()` (same file), `.ptbFilterWidth_Scroll()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.ptbFilterShift_Scroll()`** — L35993 — `private void ptbFilterShift_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbFilterShift` is scrolled.
  Called by: `.btnFilterShiftReset_Click()` (same file)
- **`.ptbFilterShift_Update()`** — L36071 — `private void ptbFilterShift_Update(int low, int high)`
  Called by: `.UpdateRX1Filters()` (same file)
- **`.btnFilterShiftReset_Click()`** — L36128 — `private void btnFilterShiftReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnFilterShiftReset` is clicked.
  Called by: `.Console_KeyDown()` (same file), `.btnIFtoVFO_Click()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.ptbFilterWidth_Update()`** — L36148 — `private void ptbFilterWidth_Update(int low, int high)`
  Called by: `.UpdateRX1Filters()` (same file)
- **`.ptbFilterWidth_Scroll()`** — L36186 — `private void ptbFilterWidth_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbFilterWidth` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbFilterWidthScroll_newMode()`** — L36285 — `private void tbFilterWidthScroll_newMode()`
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.CATVFOSwap()`** — L36308 — `public void CATVFOSwap(string pChangec)`
  Added 6/20/05 BT for CAT commands
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CopyVFOAtoB()`** — L36331 — `public void CopyVFOAtoB()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOAtoB_Click()`** — L36336 — `private void btnVFOAtoB_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnVFOAtoB` is clicked.
  Called by: `.CATVFOAtoB()` (same file), `.Console_KeyDown()` (same file), `.CATVFOSwap()` (same file), `.CopyVFOAtoB()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.CopyVFOBtoA()`** — L36372 — `public void CopyVFOBtoA()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOBtoA_Click()`** — L36377 — `private void btnVFOBtoA_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnVFOBtoA` is clicked.
  Called by: `.CATVFOBtoA()` (same file), `.Console_KeyDown()` (same file), `.CATVFOSwap()` (same file), `.CopyVFOBtoA()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.VFOSwap()`** — L36411 — `public void VFOSwap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOSwap_Click()`** — L36416 — `private void btnVFOSwap_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnVFOSwap` is clicked.
  Called by: `.CATVFOABSwap()` (same file), `.Console_KeyDown()` (same file), `.CATVFOSwap()` (same file), `.VFOSwap()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.UpdateVFOASub()`** — L36480 — `private void UpdateVFOASub()`
  Updates vfoasub.
  Called by: `.chkPower_CheckedChanged()` (same file), `.txtVFOBFreq_LostFocus()` (same file), `.chkVFOSplit_CheckedChanged()` (same file), `.chkEnableMultiRX_CheckedChanged()` (same file)
- **`.chkVFOSplit_CheckedChanged()`** — L36560 — `private void chkVFOSplit_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOSplit` checked state changes.
  Called by: `.chkPower_CheckedChanged()` (same file), `.chkEnableMultiRX_CheckedChanged()` (same file)
- **`.SetQuickSplit()`** — L36683 — `public void SetQuickSplit()`
  Sets quick split.
  Called by: `.chkVFOSplit_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file)
- **`.chkXIT_CheckedChanged()`** — L36793 — `private void chkXIT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkXIT` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRIT_CheckedChanged()`** — L36821 — `private void chkRIT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRIT` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRIT_ValueChanged()`** — L36856 — `private void udRIT_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRIT` value changes.
  Called by: `.udRIT_LostFocus()` (same file)
- **`.udXIT_ValueChanged()`** — L36883 — `private void udXIT_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udXIT` value changes.
  Called by: `.udXIT_LostFocus()` (same file)
- **`.btnXITReset_Click()`** — L36910 — `private void btnXITReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnXITReset` is clicked.
  Called by: `.Console_KeyDown()` (same file), `.DoGeneralSettingAction()` (same file)
- **`.btnRITReset_Click()`** — L36915 — `private void btnRITReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnRITReset` is clicked.
  Called by: `.Console_KeyDown()` (same file), `.DoGeneralSettingAction()` (same file)
- **`.setRIT_LEDs()`** — L36920 — `private void setRIT_LEDs()`
  Sets rit leds.
  Called by: `.udRIT_ValueChanged()` (same file), `.udXIT_ValueChanged()` (same file)
- **`.setXIT_LEDs()`** — L36939 — `private void setXIT_LEDs()`
  Sets xit leds.
  Called by: `.udRIT_ValueChanged()` (same file), `.udXIT_ValueChanged()` (same file)
- **`.btnZeroBeat_Click()`** — L36958 — `private void btnZeroBeat_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnZeroBeat` is clicked.
  Called by: `.CalibrateLevel()` (same file), `.Console_KeyDown()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.FindPeakFreqInPassband()`** — L37048 — `unsafe private int FindPeakFreqInPassband()`
  Finds peak freq in passband.
  Called by: `.btnZeroBeat_Click()` (same file)
- **`.btnIFtoVFO_Click()`** — L37125 — `private void btnIFtoVFO_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnIFtoVFO` is clicked.
  Called by: `.CATVFOSwap()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.chkANF_CheckedChanged()`** — L37224 — `private void chkANF_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkANF` checked state changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.chkDSPNB2_CheckedChanged()`** — L37251 — `private void chkDSPNB2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPNB2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NB2_CheckedChanged()`** — L37278 — `private void chkRX2NB2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2NB2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCPDR_CheckedChanged()`** — L37314 — `private void chkCPDR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCPDR` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbCPDR_Scroll()`** — L37350 — `private void ptbCPDR_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCPDR` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkCPDR_CheckedChanged()` (same file)
- **`.chkDX_CheckedChanged()`** — L37363 — `private void chkDX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMemoryQuickSave_Click()`** — L37375 — `private void btnMemoryQuickSave_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnMemoryQuickSave` is clicked.
  Called by: `.CATMemoryQS()` (same file), `.Console_KeyDown()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.btnMemoryQuickRestore_Click()`** — L37382 — `private void btnMemoryQuickRestore_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnMemoryQuickRestore` is clicked.
  Called by: `.CATMemoryQR()` (same file), `.Console_KeyDown()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.ptbPanMainRX_Scroll()`** — L37393 — `private void ptbPanMainRX_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbPanMainRX` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkPanSwap_CheckedChanged()` (same file), `.ptbPanMainRX_DoubleClick()` (same file)
- **`.ptbPanSubRX_Scroll()`** — L37413 — `private void ptbPanSubRX_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbPanSubRX` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkPanSwap_CheckedChanged()` (same file), `.ptbPanSubRX_DoubleClick()` (same file)
- **`.chkEnableMultiRX_CheckedChanged()`** — L37437 — `unsafe private void chkEnableMultiRX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnableMultiRX` checked state changes.
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.chkPanSwap_CheckedChanged()`** — L37536 — `private void chkPanSwap_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPanSwap` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX0Gain_Scroll()`** — L37551 — `private void ptbRX0Gain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX0Gain` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.ptbRX1Gain_Scroll()`** — L37592 — `private void ptbRX1Gain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX1Gain` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.chkFullDuplex_CheckedChanged()`** — L37635 — `private void chkFullDuplex_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkFullDuplex` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getConsole()`** — L37661 — `public static Console getConsole()`
  Returns console.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WndProc()`** — L37666 — `protected override void WndProc(ref Message m)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkFWCATUBypass_Click()`** — L37679 — `private void chkFWCATUBypass_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkFWCATUBypass` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSplitDisplay_CheckedChanged()`** — L37683 — `private void chkSplitDisplay_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkSplitDisplay` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ckQuickPlay_CheckedChanged()`** — L37694 — `private async void ckQuickPlay_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ckQuickPlay` checked state changes.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.ckQuickRec_CheckedChanged()`** — L37738 — `private void ckQuickRec_CheckedChanged(object sender, System.EventArgs e)`
  private bool _updated_from_wave_form = false; public bool UpdatedFromWaveForm { // prevent wave form changes causing loop get { return _updated_from_wave_form; } set { _updated_from_wave_form = value; } }
  Called by: `.DoOtherButtonAction()` (same file)
- **`.moveModeSpecificPanels()`** — L37784 — `private void moveModeSpecificPanels()`
  Called by: `.ResizeConsole()` (same file), `.SelectModeDependentPanel()` (`Console/Andromeda/Andromeda.cs`)
- **`.ResizeConsole()`** — L37791 — `private void ResizeConsole(int h_delta, int v_delta)`
  Called by: `.Console_Resize()` (same file)
- **`.GrabConsoleSizeBasis()`** — L37926 — `public void GrabConsoleSizeBasis()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2_CheckedChanged()`** — L38284 — `private void chkRX2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2` checked state changes.
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.setSmallRX2ModeFilterLabels()`** — L38390 — `private void setSmallRX2ModeFilterLabels()`
  Sets small rx2 mode filter labels.
  Called by: `.radModeButton_CheckedChanged()` (same file), `.radFilter_CheckedChanged()` (same file), `.chkEnableMultiRX_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.radRX2ModeButton_CheckedChanged()` (same file)
- **`.chkRX2SR_CheckedChanged()`** — L38418 — `private void chkRX2SR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2SR` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.panelVFOASubHover_Paint()`** — L38440 — `private void panelVFOASubHover_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `panelVFOASubHover` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panelVFOASubHover_MouseMove()`** — L38456 — `private void panelVFOASubHover_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `panelVFOASubHover` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOABand_MouseMove()`** — L38468 — `private void txtVFOABand_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOABand` receives mouse movement.
  Called by: `.panelVFOASubHover_MouseMove()` (same file)
- **`.txtVFOABand_MouseLeave()`** — L38499 — `private void txtVFOABand_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOABand` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetRX2Mode()`** — L38505 — `private void SetRX2Mode(DSPMode new_mode)`
  Sets rx2 mode.
  Called by: `.radRX2ModeButton_CheckedChanged()` (same file), `.radRX2ModeLSB_CheckedChanged()` (same file), `.radRX2ModeUSB_CheckedChanged()` (same file), `.radRX2ModeDSB_CheckedChanged()` (same file), `.radRX2ModeCWL_CheckedChanged()` (same file), `.radRX2ModeCWU_CheckedChanged()` (same file) — and 7 more
- **`.radRX2ModeButton_CheckedChanged()`** — L39042 — `private void radRX2ModeButton_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeButton` checked state changes.
  Called by: `.selectModes()` (same file)
- **`.radRX2ModeLSB_CheckedChanged()`** — L39111 — `private void radRX2ModeLSB_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeLSB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeUSB_CheckedChanged()`** — L39119 — `private void radRX2ModeUSB_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeUSB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDSB_CheckedChanged()`** — L39127 — `private void radRX2ModeDSB_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDSB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeCWL_CheckedChanged()`** — L39135 — `private void radRX2ModeCWL_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeCWL` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeCWU_CheckedChanged()`** — L39143 — `private void radRX2ModeCWU_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeCWU` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeFMN_CheckedChanged()`** — L39151 — `private void radRX2ModeFMN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeFMN` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeAM_CheckedChanged()`** — L39159 — `private void radRX2ModeAM_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeAM` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeSAM_CheckedChanged()`** — L39167 — `private void radRX2ModeSAM_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeSAM` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDIGL_CheckedChanged()`** — L39175 — `private void radRX2ModeDIGL_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDIGL` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDIGU_CheckedChanged()`** — L39183 — `private void radRX2ModeDIGU_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDIGU` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDRM_CheckedChanged()`** — L39191 — `private void radRX2ModeDRM_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDRM` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetRX2Filter()`** — L39199 — `public void SetRX2Filter(Filter new_filter, bool update = true)`
  Sets rx2 filter.
  Called by: `.radRX2Filter_CheckedChanged()` (same file), `.radRX2Filter1_CheckedChanged()` (same file), `.radRX2Filter2_CheckedChanged()` (same file), `.radRX2Filter3_CheckedChanged()` (same file), `.radRX2Filter4_CheckedChanged()` (same file), `.radRX2Filter5_CheckedChanged()` (same file) — and 4 more
- **`.radRX2Filter1_CheckedChanged()`** — L39297 — `private void radRX2Filter1_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter2_CheckedChanged()`** — L39303 — `private void radRX2Filter2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter3_CheckedChanged()`** — L39309 — `private void radRX2Filter3_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter4_CheckedChanged()`** — L39315 — `private void radRX2Filter4_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter4` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter5_CheckedChanged()`** — L39321 — `private void radRX2Filter5_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter5` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter6_CheckedChanged()`** — L39327 — `private void radRX2Filter6_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter6` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter7_CheckedChanged()`** — L39333 — `private void radRX2Filter7_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter7` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2FilterVar1_CheckedChanged()`** — L39339 — `private void radRX2FilterVar1_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2FilterVar1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2FilterVar2_CheckedChanged()`** — L39345 — `private void radRX2FilterVar2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2FilterVar2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2FilterLow_ValueChanged()`** — L39351 — `private void udRX2FilterLow_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2FilterLow` value changes.
  Called by: `.udRX2FilterLow_LostFocus()` (same file)
- **`.udRX2FilterHigh_ValueChanged()`** — L39378 — `private void udRX2FilterHigh_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2FilterHigh` value changes.
  Called by: `.udRX2FilterHigh_LostFocus()` (same file)
- **`.chkRX2ANF_CheckedChanged()`** — L39399 — `private void chkRX2ANF_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2ANF` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2BIN_CheckedChanged()`** — L39426 — `private void chkRX2BIN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2BIN` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboRX2MeterMode_SelectedIndexChanged()`** — L39440 — `private void comboRX2MeterMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2MeterMode` selection changes.
  Called by: `.UIMOXChangedFalse()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.chkRX2Preamp_CheckedChanged()`** — L39487 — `private void chkRX2Preamp_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2Preamp` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2RF_Scroll()`** — L39500 — `private void ptbRX2RF_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2RF` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.picRX2Squelch_Paint()`** — L39533 — `private void picRX2Squelch_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picRX2Squelch` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX1Preamp_CheckedChanged()`** — L39543 — `private void chkRX1Preamp_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX1Preamp` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2Pan_Scroll()`** — L39559 — `private void ptbRX2Pan_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2Pan` is scrolled.
  Called by: `.InitConsole()` (same file), `.ptbRX2Pan_DoubleClick()` (same file)
- **`.ptbRX2Gain_Scroll()`** — L39578 — `private void ptbRX2Gain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2Gain` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkRX2Mute_CheckedChanged()` (same file)
- **`.chkRX2Mute_CheckedChanged()`** — L39617 — `private void chkRX2Mute_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2Mute` checked state changes.
  Called by: `.InitConsole()` (same file), `.SetRX2Mode()` (same file)
- **`.comboRX2DisplayMode_SelectedIndexChanged()`** — L39657 — `private void comboRX2DisplayMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2DisplayMode` selection changes.
  Called by: `.SetupDisplayEngine()` (same file)
- **`.chkRX2DisplayAVG_CheckedChanged()`** — L39705 — `private void chkRX2DisplayAVG_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2DisplayAVG` checked state changes.
  Called by: `.chkRX2_CheckedChanged()` (same file)
- **`.chkRX2DisplayPeak_CheckedChanged()`** — L39727 — `private void chkRX2DisplayPeak_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2DisplayPeak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateDSP()`** — L39758 — `private void UpdateDSP()`
  Updates dsp.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.BuildFilterCharacteristics()`** — L39984 — `public void BuildFilterCharacteristics()`
  Builds filter characteristics.
  Called by: `.UpdateDSP()` (same file)
- **`.calcFilterCharacteristics()`** — L40016 — `private (double[], int, int) calcFilterCharacteristics(int id, double rate, int filter_size, int w_type, double corner_freq, bool hi_res)`
  Called by: `.BuildFilterCharacteristics()` (same file)
- **`.SetupRX2Band()`** — L40269 — `public void SetupRX2Band(Band b, bool is_for_rx1_vfo_b = false)`
  Setups rx2 band.
  Called by: `.comboRX2Band_SelectedIndexChanged()` (same file), `.HandleFrontPanelButtonPress()` (`Console/Andromeda/Andromeda.cs`), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.comboRX2Band_SelectedIndexChanged()`** — L40306 — `private void comboRX2Band_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2Band` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateResolutionStatusBarText()`** — L40341 — `private void updateResolutionStatusBarText()`
  Called by: `.Console_Resize()` (same file), `.Console_Shown()` (same file)
- **`.Console_Resize()`** — L40359 — `private void Console_Resize(object sender, System.EventArgs e)`
  WinForms event handler: runs when `Console` is resized.
  Called by: `.chkRX2_CheckedChanged()` (same file)
- **`.comboRX2AGC_SelectedIndexChanged()`** — L40431 — `private void comboRX2AGC_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2AGC` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOSync_CheckedChanged()`** — L40520 — `private void chkVFOSync_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOSync` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOATX_CheckedChanged()`** — L40614 — `private void chkVFOATX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOATX` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.BroadcastVFOChange()`** — L40678 — `private void BroadcastVFOChange(string ndx)`
  Called by: `.OnVFOTXChanged()` (same file)
- **`.chkVFOBTX_CheckedChanged()`** — L40691 — `private void chkVFOBTX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOBTX` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.toolStripMenuItemRX1FilterConfigure_Click()`** — L40793 — `private void toolStripMenuItemRX1FilterConfigure_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX1FilterConfigure` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItemRX1FilterReset_Click()`** — L40806 — `private void toolStripMenuItemRX1FilterReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX1FilterReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItemRX2FilterConfigure_Click()`** — L40853 — `private void toolStripMenuItemRX2FilterConfigure_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX2FilterConfigure` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchDelete_Click()`** — L40866 — `private void toolStripNotchDelete_Click(Object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchDelete` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchRemember_Click()`** — L40870 — `private void toolStripNotchRemember_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchRemember` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchNormal_Click()`** — L40874 — `private void toolStripNotchNormal_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchNormal` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchDeep_Click()`** — L40878 — `private void toolStripNotchDeep_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchDeep` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchVeryDeep_Click()`** — L40882 — `private void toolStripNotchVeryDeep_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchVeryDeep` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItemRX2FilterReset_Click()`** — L40886 — `private void toolStripMenuItemRX2FilterReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX2FilterReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTNF_CheckedChanged()`** — L40933 — `private void chkTNF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTNF` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.ChangeNotchBW()`** — L40957 — `unsafe public bool ChangeNotchBW(MNotch notch, double newWidth, int notch_index = -1)`
  Called by: `.onBWChanged()` (same file), `.notchMouseWheel()` (same file), `.pnlDisplay_MouseMove()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.ChangeNotchCentreFrequency()`** — L41000 — `unsafe public bool ChangeNotchCentreFrequency(MNotch notch, double newCentreFrequencyHz, int sourceRX, int notch_index = -1)`
  Called by: `.pnlDisplay_MouseMove()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.changeNotchActive()`** — L41073 — `unsafe private bool changeNotchActive(MNotch notch, bool bActive)`
  Called by: `.onActiveChanged()` (same file)
- **`.toggleNotchActive()`** — L41109 — `unsafe private bool toggleNotchActive(MNotch notch)`
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.removeNotch()`** — L41148 — `private bool removeNotch(MNotch notch)`
  Called by: `.onNotchDelete()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.AddNotch()`** — L41172 — `public void AddNotch(double fFreqHZ, int sourceRX)`
  Adds notch.
  Called by: `.TNFAdd()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.notchSidebandShift()`** — L41231 — `private int notchSidebandShift(int rx)`
  Called by: `.TNFAdd()` (same file)
- **`.btnTNFAdd_Click()`** — L41259 — `private void btnTNFAdd_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTNFAdd` is clicked.
  Called by: `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.TNFAdd()`** — L41263 — `public void TNFAdd(int rx)`
  Called by: `.btnTNFAdd_Click()` (same file), `.DoOtherButtonAction()` (same file)
- **`.ptbFMMic_Scroll()`** — L41283 — `private void ptbFMMic_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbFMMic` is scrolled.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.chkFMCTCSS_CheckedChanged()`** — L41313 — `private void chkFMCTCSS_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMCTCSS` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboFMCTCSS_SelectedIndexChanged()`** — L41318 — `private void comboFMCTCSS_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFMCTCSS` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.InitCTCSS()`** — L41324 — `private void InitCTCSS()`
  Inits ctcss.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitMemoryFrontPanel()`** — L41331 — `private void InitMemoryFrontPanel()`
  Inits memory front panel.
  Called by: `.InitConsole()` (same file)
- **`.radFMDeviation2kHz_CheckedChanged()`** — L41338 — `private void radFMDeviation2kHz_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFMDeviation2kHz` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.fmDeviation2k()`** — L41342 — `private void fmDeviation2k(bool force)`
  Called by: `.InitConsole()` (same file), `.radFMDeviation2kHz_CheckedChanged()` (same file)
- **`.radFMDeviation5kHz_CheckedChanged()`** — L41376 — `private void radFMDeviation5kHz_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFMDeviation5kHz` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.fmDeviation5k()`** — L41380 — `private void fmDeviation5k(bool force)`
  Called by: `.InitConsole()` (same file), `.radFMDeviation5kHz_CheckedChanged()` (same file)
- **`.udFMOffset_ValueChanged()`** — L41415 — `private void udFMOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udFMOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXHigh_CheckedChanged()`** — L41420 — `private void chkFMTXHigh_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXHigh` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXSimplex_CheckedChanged()`** — L41434 — `private void chkFMTXSimplex_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXSimplex` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXLow_CheckedChanged()`** — L41448 — `private void chkFMTXLow_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXLow` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXRev_CheckedChanged()`** — L41462 — `private void chkFMTXRev_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXRev` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMMode_Click()`** — L41490 — `private void chkFMMode_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMMode` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuMemory_Click()`** — L41539 — `private void mnuMemory_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuMemory` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.RecallMemory()`** — L41547 — `public void RecallMemory(MemoryRecord record)`
  Called by: `.comboFMMemory_SelectedIndexChanged()` (same file), `.changeComboFMMemory()` (same file)
- **`.comboFMMemory_SelectedIndexChanged()`** — L41577 — `private void comboFMMemory_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFMMemory` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnFMMemoryUp_Click()`** — L41585 — `private void btnFMMemoryUp_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFMMemoryUp` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnFMMemoryDown_Click()`** — L41591 — `private void btnFMMemoryDown_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFMMemoryDown` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.changeComboFMMemory()`** — L41597 — `public void changeComboFMMemory(int index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnFMMemory_Click()`** — L41613 — `private void btnFMMemory_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFMMemory` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.memoryToolStripMenuItem_Click()`** — L41632 — `private void memoryToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `memoryToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.waveToolStripMenuItem_Click()`** — L41653 — `private void waveToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `waveToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.CollapseToolStripMenuItem_Click()`** — L41672 — `private void CollapseToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `CollapseToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.equalizerToolStripMenuItem_Click()`** — L41684 — `private void equalizerToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `equalizerToolStripMenuItem` is clicked.
  Called by: `.chkRXEQ_MouseDown()` (same file), `.chkTXEQ_MouseDown()` (same file), `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.xVTRsToolStripMenuItem_Click()`** — L41705 — `private void xVTRsToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `xVTRsToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.cWXToolStripMenuItem_Click()`** — L41724 — `private void cWXToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `cWXToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.SetComboPreampForHPSDR()`** — L41777 — `public void SetComboPreampForHPSDR()`
  Sets combo preamp for hpsdr.
  Called by: `.SetupForHPSDRModel()` (same file)
- **`.MakeLineInList()`** — L41852 — `private void MakeLineInList()`
  Called by: `.SetMicGain()` (same file)
- **`.SetMicXlr()`** — L41864 — `public void SetMicXlr()`
  Sets mic xlr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMicGain()`** — L41870 — `public void SetMicGain()`
  Sets mic gain.
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.BPF1ToolStripMenuItem_Click()`** — L41885 — `private void BPF1ToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BPF1ToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BPF2ToolStripMenuItem_Click()`** — L41890 — `private void BPF2ToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BPF2ToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ExpandDisplay()`** — L42009 — `private void ExpandDisplay(bool bSuspendDraw = true)`
  Called by: `.GetState()` (same file), `.CollapseToolStripMenuItem_Click()` (same file)
- **`.setPAProfileLabelPos()`** — L42427 — `private void setPAProfileLabelPos()`
  Sets paprofile label pos.
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.CollapseDisplay()`** — L42475 — `public void CollapseDisplay(bool bSuspendDraw = true)`
  modified G8NJJ to add alternate top/button controls for Andromeda optimised for 1024x600 touchscreen display
  Called by: `.GetState()` (same file), `.btnBandVHF_Click()` (same file), `.btnBandHF_Click()` (same file), `.btnBandGEN_Click()` (same file), `.CollapseToolStripMenuItem_Click()` (same file), `.radRX1Show_CheckedChanged()` (same file) — and 1 more
- **`.RepositionControlsForCollapsedlDisplay()`** — L42978 — `private void RepositionControlsForCollapsedlDisplay()`
  relocate the controls on the collapsed display
  Called by: `.ResizeConsole()` (same file), `.CollapseDisplay()` (same file)
- **`.mnuFilter_Click()`** — L43390 — `private void mnuFilter_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuFilter` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuFilterRX2_Click()`** — L43431 — `private void mnuFilterRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuFilterRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuDSP_Click()`** — L43463 — `private void mnuDSP_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuDSP` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuDSPRX2_Click()`** — L43524 — `private void mnuDSPRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuDSPRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuBand_Click()`** — L43582 — `private void mnuBand_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuBand` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuBandRX2_Click()`** — L43632 — `private void mnuBandRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuBandRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupHiddenButton()`** — L43713 — `private void setupHiddenButton()`
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.RepositionControlsForCollapsedlDisplay()` (same file)
- **`.mnuMode_Click()`** — L43720 — `private void mnuMode_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuMode` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuModeRX2_Click()`** — L43767 — `private void mnuModeRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuModeRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuShowTopControls_Click()`** — L43811 — `private void mnuShowTopControls_Click(object sender, EventArgs e)`
  handlers for menu display controls events. The persistent state is held on the setup form matching controls
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuShowBandControls_Click()`** — L43816 — `private void mnuShowBandControls_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuShowBandControls` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuShowModeControls_Click()`** — L43821 — `private void mnuShowModeControls_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuShowModeControls` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.AndromedaTopControlsToolStripMenuItem_Click()`** — L43826 — `private void AndromedaTopControlsToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `AndromedaTopControlsToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.AndromedaButtonBarToolStripMenuItem_Click()`** — L43831 — `private void AndromedaButtonBarToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `AndromedaButtonBarToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBand_CheckedChanged()`** — L43836 — `private void radBand_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.eSCToolStripMenuItem_Click()`** — L43862 — `private void eSCToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `eSCToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.showHideDiversity()`** — L43867 — `private void showHideDiversity(bool show, bool starting_up = false)`
  Called by: `.eSCToolStripMenuItem_Click()` (same file), `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.ptbRX1AF_Scroll()`** — L43925 — `private void ptbRX1AF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX1AF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2AF_Scroll()`** — L43937 — `private void ptbRX2AF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2AF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX1Show_CheckedChanged()`** — L43949 — `private void radRX1Show_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX1Show` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.radRX2Show_CheckedChanged()`** — L43970 — `private void radRX2Show_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX2Show` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.ptbAF_DoubleClick()`** — L43991 — `private void ptbAF_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbAF` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX1AF_DoubleClick()`** — L43996 — `private void ptbRX1AF_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX1AF` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2AF_DoubleClick()`** — L44002 — `private void ptbRX2AF_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2AF` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX1StepAttData_ValueChanged()`** — L44008 — `private void udRX1StepAttData_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX1StepAttData` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2StepAttData_ValueChanged()`** — L44020 — `private void udRX2StepAttData_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2StepAttData` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblPreamp_MouseDoubleClick()`** — L44031 — `private void lblPreamp_MouseDoubleClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lblRX2Preamp_MouseDoubleClick()`** — L44054 — `private void lblRX2Preamp_MouseDoubleClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetFocusMaster()`** — L44064 — `public void SetFocusMaster(bool state)`
  Sets focus master.
  Called by: `.n1mm_delay_Elapsed()` (same file), `.textbox_GotFocus()` (same file), `.textbox_LostFocus()` (same file), `.combo_OpenDropDown()` (same file), `.combo_CloseDropDown()` (same file), `.memoryToolStripMenuItem_Click()` (same file) — and 2 more
- **`.chkFWCATU_CheckedChanged()`** — L44109 — `private void chkFWCATU_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFWCATU` checked state changes.
  Called by: `.GetState()` (same file)
- **`.linearityToolStripMenuItem_Click()`** — L44138 — `private void linearityToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `linearityToolStripMenuItem` is clicked.
  Called by: `.chkFWCATUBypass_MouseDown()` (same file), `.DoOtherButtonAction()` (same file)
- **`.RAtoolStripMenuItem_Click()`** — L44146 — `private void RAtoolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `RAtoolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.SetDigiMode()`** — L44165 — `private void SetDigiMode(int rx, DigiMode.DigiModeSettingState mode, bool bFromTXProfile = false)`
  Sets digi mode.
  Called by: `.LoadedTXProfile()` (same file), `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.chkCWFWKeyer_CheckedChanged()`** — L44250 — `private void chkCWFWKeyer_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCWFWKeyer` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nr_selected_from_text()`** — L44256 — `private void nr_selected_from_text(string text)`
  Called by: `.GetState()` (same file)
- **`.nr_selected_to_text()`** — L44283 — `private string nr_selected_to_text()`
  Called by: `.GetStateList()` (same file)
- **`.incrementNR()`** — L44293 — `private void incrementNR(int rx)`
  Called by: `.chkNR_Click()` (same file), `.chkRX2NR_Click()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.SelectNR()`** — L44307 — `public void SelectNR(int rx, bool incude_sub, int nr)`
  Selects nr.
  Called by: `.Console_KeyDown()` (same file), `.SetRX1Mode()` (same file), `.mnuDSP_Click()` (same file), `.mnuDSPRX2_Click()` (same file), `.SetDigiMode()` (same file), `.DoOtherButtonAction()` (same file)
- **`.GetSelectedNR()`** — L44321 — `public int GetSelectedNR(int rx)`
  Returns selected nr.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.setupNR()`** — L44326 — `private void setupNR(int rx, bool sub)`
  Called by: `.nr_selected_from_text()` (same file), `.incrementNR()` (same file), `.SelectNR()` (same file)
- **`.wbClosing()`** — L44512 — `public void wbClosing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.wBToolStripMenuItem_Click()`** — L44518 — `private void wBToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `wBToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.UpdatePIVisibilty()`** — L44533 — `public void UpdatePIVisibilty()`
  Updates pivisibilty.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.pIToolStripMenuItem_Click()`** — L44544 — `private void pIToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pIToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNB_CheckStateChanged()`** — L44553 — `private void chkNB_CheckStateChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2NB_CheckStateChanged()`** — L44601 — `private void chkRX2NB_CheckStateChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2NB2_CheckStateChanged()`** — L44650 — `private void chkRX2NB2_CheckStateChanged(object sender, EventArgs e)`
  RX2 Spectral Noise Blanker (SNB)
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LoadLEDFont()`** — L44667 — `private void LoadLEDFont()`
  Loads ledfont.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddFontMemResourceEx()`** — L44674 — `[DllImport("gdi32.dll", ExactSpelling = true)] private static extern IntPtr AddFontMemResourceEx(byte[] pbFont, int cbFont, IntPtr pdv, out uint pcFonts)`
  Adds font mem resource ex.
  Called by: `.GetCustomFont()` (same file)
- **`.GetCustomFont()`** — L44677 — `static public Font GetCustomFont(byte[] fontData, float size, FontStyle style)`
  Returns custom font.
  Called by: `.LoadLEDFont()` (same file)
- **`.regBox1_Click()`** — L44712 — `private void regBox1_Click(object sender, EventArgs e)`
  ke9ns add open up bandstack window when you click on the bandstack index
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.regBox_Click()`** — L44719 — `private void regBox_Click(object sender, EventArgs e)`
  ke9ns add open up bandstack window when you click on the bandstack index
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterHigh_ValueChanged()`** — L44728 — `private void udTXFilterHigh_ValueChanged(object sender, EventArgs e)`
  ke9ns add to allow TX filter on main console SSB panel
  Called by: `.udTXFilterHigh_LostFocus()` (same file)
- **`.udTXFilterLow_ValueChanged()`** — L44735 — `private void udTXFilterLow_ValueChanged(object sender, EventArgs e)`
  ke9ns add
  Called by: `.udTXFilterLow_LostFocus()` (same file)
- **`.ForcePureSignalAutoCalDisable()`** — L44742 — `public void ForcePureSignalAutoCalDisable()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkFWCATUBypass_CheckedChanged()`** — L44748 — `private void chkFWCATUBypass_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFWCATUBypass` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRxAnt_CheckedChanged()`** — L44781 — `private void chkRxAnt_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRxAnt` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkQSK_CheckStateChanged()`** — L44793 — `private void chkQSK_CheckStateChanged(object sender, EventArgs e)`
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.ptbPanMainRX_DoubleClick()`** — L44799 — `private void ptbPanMainRX_DoubleClick(object sender, EventArgs e)`
  MW0LGE
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbPanSubRX_DoubleClick()`** — L44812 — `private void ptbPanSubRX_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbPanSubRX` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2Pan_DoubleClick()`** — L44825 — `private void ptbRX2Pan_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2Pan` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setBackground()`** — L44839 — `private void setBackground()`
  MW0LGE
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initialiseRawInput()`** — L44934 — `private void initialiseRawInput()`
  MW0LGE RAWINPUT
  Called by: `.InitConsole()` (same file)
- **`.updateRawInputDevices()`** — L44957 — `private void updateRawInputDevices()`
  Called by: `.initialiseRawInput()` (same file), `.OnDevicesChanged()` (same file)
- **`.OnDevicesChanged()`** — L45040 — `private void OnDevicesChanged(object sender)`
  Handles/raises the devices changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnKeyPressedRaw()`** — L45046 — `private void OnKeyPressedRaw(object sender, RawInputEventArg e)`
  Handles/raises the key pressed raw event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseWheelChanged()`** — L45073 — `private void OnMouseWheelChanged(object sender, RawInputEventArg e)`
  Handles/raises the mouse wheel changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.incrementMutliMeterDisplayMode()`** — L45099 — `private void incrementMutliMeterDisplayMode()`
  Called by: `.txtMultiText_Click()` (same file), `.txtRX2Meter_Click()` (same file)
- **`.txtMultiText_Click()`** — L45111 — `private void txtMultiText_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMultiText` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtRX2Meter_Click()`** — L45116 — `private void txtRX2Meter_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `txtRX2Meter` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripStatusLabel_SeqWarning_Click()`** — L45121 — `private void toolStripStatusLabel_SeqWarning_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripStatusLabel_SeqWarning` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItem_4by3_DropDownItemClicked()`** — L45126 — `private void toolStripMenuItem_4by3_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolStripMenuItem_16by9_DropDownItemClicked()`** — L45131 — `private void toolStripMenuItem_16by9_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolStripMenuItem_16by10_DropDownItemClicked()`** — L45136 — `private void toolStripMenuItem_16by10_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.youTubeToolStripMenuItem_DropDownItemClicked()`** — L45141 — `private void youTubeToolStripMenuItem_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetResolution()`** — L45145 — `public void SetResolution(string resolutionString)`
  Sets resolution.
  Called by: `.toolStripMenuItem_4by3_DropDownItemClicked()` (same file), `.toolStripMenuItem_16by9_DropDownItemClicked()` (same file), `.toolStripMenuItem_16by10_DropDownItemClicked()` (same file), `.youTubeToolStripMenuItem_DropDownItemClicked()` (same file)
- **`.includeBordersToolStripMenuItem_Click()`** — L45207 — `private void includeBordersToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `includeBordersToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseDown()`** — L45216 — `private void pnlResizeMeter_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseMove()`** — L45225 — `private void pnlResizeMeter_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseUp()`** — L45260 — `private void pnlResizeMeter_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseEnter()`** — L45265 — `private void pnlResizeMeter_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseLeave()`** — L45270 — `private void pnlResizeMeter_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.systemToolStripMenuItem_Click()`** — L45275 — `private void systemToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `systemToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.thetisOnlyToolStripMenuItem_Click()`** — L45281 — `private void thetisOnlyToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `thetisOnlyToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.basicAudioLoadCompletedEvent()`** — L45347 — `private void basicAudioLoadCompletedEvent(bool bLoadedOk)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QSOTimerReset()`** — L45380 — `public void QSOTimerReset(bool bAutoReset = false)`
  Called by: `.HdwMOXChanged()` (same file), `.toolStripStatusLabel_timer_MouseUp()` (same file)
- **`.toolStripStatusLabel_timer_Click()`** — L45398 — `private void toolStripStatusLabel_timer_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripStatusLabel_timer` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripStatusLabel_timer_MouseUp()`** — L45411 — `private void toolStripStatusLabel_timer_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `toolStripStatusLabel_timer` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateQSOTimer()`** — L45418 — `private void updateQSOTimer()`
  Called by: `.timer_clock_Tick()` (same file)
- **`.updateQSOTimerStatusbar()`** — L45426 — `private void updateQSOTimerStatusbar()`
  Called by: `.timer_clock_Tick()` (same file), `.QSOTimerReset()` (same file)
- **`.chkVFOSync_MouseDown()`** — L45472 — `private void chkVFOSync_MouseDown(object sender, MouseEventArgs e)`
  -- RIGHT click on control shows related setup page // refactored
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNR_MouseDown()`** — L45476 — `private void chkNR_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkNR` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NR_MouseDown()`** — L45481 — `private void chkRX2NR_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2NR` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNB_MouseDown()`** — L45486 — `private void chkNB_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkNB` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDSPNB2_MouseDown()`** — L45490 — `private void chkDSPNB2_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkDSPNB2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NB_MouseDown()`** — L45494 — `private void chkRX2NB_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2NB` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NB2_MouseDown()`** — L45498 — `private void chkRX2NB2_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2NB2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeCWL_MouseDown()`** — L45502 — `private void radModeCWL_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeCWL` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeCWU_MouseDown()`** — L45506 — `private void radModeCWU_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeCWU` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC1_MouseDown()`** — L45510 — `private void chkVAC1_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVAC1` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2_MouseDown()`** — L45514 — `private void chkVAC2_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVAC2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeAM_MouseDown()`** — L45518 — `private void radModeAM_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeAM` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeSAM_MouseDown()`** — L45522 — `private void radModeSAM_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeSAM` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWAPFEnabled_MouseDown()`** — L45526 — `private void chkCWAPFEnabled_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkCWAPFEnabled` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTNF_MouseDown()`** — L45530 — `private void chkTNF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkTNF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVOX_MouseDown()`** — L45534 — `private void chkVOX_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVOX` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCPDR_MouseDown()`** — L45538 — `private void chkCPDR_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkCPDR` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNoiseGate_MouseDown()`** — L45542 — `private void chkNoiseGate_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkNoiseGate` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeFMN_MouseDown()`** — L45546 — `private void radModeFMN_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeFMN` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAGC_MouseDown()`** — L45550 — `private void comboAGC_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `comboAGC` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMicMute_MouseDown()`** — L45554 — `private void chkMicMute_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkMicMute` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboTXProfile_MouseDown()`** — L45558 — `private void comboTXProfile_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `comboTXProfile` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblRF_MouseDown()`** — L45562 — `private void lblRF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblRF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.IsRightButton()`** — L45566 — `private bool IsRightButton(MouseEventArgs e)`
  Called by: `.chkVFOSync_MouseDown()` (same file), `.chkNR_MouseDown()` (same file), `.chkRX2NR_MouseDown()` (same file), `.chkNB_MouseDown()` (same file), `.chkDSPNB2_MouseDown()` (same file), `.chkRX2NB_MouseDown()` (same file) — and 31 more
- **`.ProcessDialogKey()`** — L45571 — `protected override bool ProcessDialogKey(Keys keyData)`
  Processes dialog key.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkVersions()`** — L45588 — `private bool checkVersions()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToolStripMenuItem15_Click()`** — L45689 — `private void ToolStripMenuItem15_Click(object sender, EventArgs e)`
  set TX antenna to 1,2 or 3
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem16_Click()`** — L45694 — `private void ToolStripMenuItem16_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem16` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem17_Click()`** — L45699 — `private void ToolStripMenuItem17_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem17` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem18_Click()`** — L45704 — `private void ToolStripMenuItem18_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem18` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem19_Click()`** — L45709 — `private void ToolStripMenuItem19_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem19` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem20_Click()`** — L45714 — `private void ToolStripMenuItem20_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem20` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_Shown()`** — L45719 — `private void Console_Shown(object sender, EventArgs e)`
  WinForms event handler: runs when `Console` is shown.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chk2TONE_CheckedChanged()`** — L45767 — `private async void chk2TONE_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chk2TONE` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucQuickRecallPad_ButtonClicked()`** — L45801 — `private void ucQuickRecallPad_ButtonClicked(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lblBandStack_Click()`** — L45807 — `private void lblBandStack_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblBandStack` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.getControl()`** — L45815 — `private void getControl(Control cc, Point p, string sub)`
  Returns control.
  Called by: `.gmh_MouseMove()` (same file)
- **`.gmh_MouseMove()`** — L45847 — `private void gmh_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `gmh` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.addDelegates()`** — L46214 — `private void addDelegates()`
  Called by: `.InitConsole()` (same file)
- **`.removeDelegates()`** — L46246 — `private void removeDelegates()`
  Called by: `.Console_Closing()` (same file)
- **`.StopAllTx()`** — L46291 — `public void StopAllTx(string msg = "")`
  Stops all tx.
  Called by: `.timeOutTimer()` (same file)
- **`.timeOutTimer()`** — L46310 — `private void timeOutTimer(string msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXInhibitChanged()`** — L46319 — `private void OnTXInhibitChanged(bool oldState, bool newState)`
  Handles/raises the txinhibit changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOTXChanged()`** — L46324 — `private void OnVFOTXChanged(bool vfoB, bool oldState, bool newState)`
  Handles/raises the vfotxchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnIgnoreDupes()`** — L46330 — `private void OnIgnoreDupes(bool ignore)`
  Handles/raises the ignore dupes event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnHideOnSelect()`** — L46334 — `private void OnHideOnSelect(bool hideOnSelect)`
  Handles/raises the hide on select event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnShowInSpectrum()`** — L46338 — `private void OnShowInSpectrum(bool show)`
  Handles/raises the show in spectrum event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPowerChangeHander()`** — L46345 — `private void OnPowerChangeHander(bool oldPower, bool newPower)`
  Handles/raises the power change hander event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleBSFChange()`** — L46357 — `private void handleBSFChange(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double oldCentreF, do`
  Called by: `.OnSetBandChangeHander()` (same file), `.OnVFOAFrequencyChangeHandler()` (same file)
- **`.updateLastVisited()`** — L46398 — `private void updateLastVisited(BandStackFilter bsf, Band band, DSPMode mode, Filter filter, double freq, double centreF, bool cTUN, int zoomSlider)`
  Called by: `.handleBSFChange()` (same file)
- **`.OnSetBandChangeHander()`** — L46409 — `private void OnSetBandChangeHander(int rx, Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double `
  Handles/raises the set band change hander event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEntryAdd()`** — L46414 — `private void OnEntryAdd(BandStackFilter bsf)`
  Handles/raises the entry add event.
  Called by: `.preBandSelect()` (same file)
- **`.OnEntryUpdate()`** — L46433 — `private void OnEntryUpdate(BandStackFilter bsf, BandStackEntry bse)`
  Handles/raises the entry update event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEntryDelete()`** — L46450 — `private void OnEntryDelete(BandStackFilter bsf, BandStackEntry bse)`
  Handles/raises the entry delete event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCentreFrequencyChanged()`** — L46468 — `private void OnCentreFrequencyChanged(int rx, double oldFreq, double newFreq, Band band, double offset)`
  Handles/raises the centre frequency changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCTUNChanged()`** — L46489 — `private void OnCTUNChanged(int rx, bool oldCTUN, bool newCTUN, Band band)`
  Handles/raises the ctunchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFilterChanged()`** — L46501 — `private void OnFilterChanged(int rx, Filter oldFilter, Filter newFilter, Band band, int low, int high, string sName)`
  Handles/raises the filter changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateBandstackOverlay()`** — L46515 — `private void updateBandstackOverlay(int rx)`
  Called by: `.SetupDisplayEngine()` (same file), `.OnShowInSpectrum()` (same file), `.handleBSFChange()` (same file), `.OnEntryAdd()` (same file), `.OnEntryUpdate()` (same file), `.OnEntryDelete()` (same file) — and 4 more
- **`.OnZoomChanged()`** — L46545 — `private void OnZoomChanged(double oldZoomFactor, double newZoomFactor, int sliderValue)`
  Handles/raises the zoom changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEntryClicked()`** — L46559 — `private void OnEntryClicked(BandStackFilter bsf, BandStackEntry bse, bool updateLastVisited = true, bool obeyHide = true)`
  Handles/raises the entry clicked event.
  Called by: `.OnEntryDelete()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.preBandSelect()`** — L46587 — `private void preBandSelect(int rx, Band band, int dir = 0)`
  Called by: `.OnBandBeforeChangeHandler()` (same file), `.SetBandStack()` (same file)
- **`.OnBandBeforeChangeHandler()`** — L46701 — `private void OnBandBeforeChangeHandler(int rx, Band band)`
  Handles/raises the band before change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setRX1BandFromBandStackEntry()`** — L46705 — `private void setRX1BandFromBandStackEntry(in BandStackEntry bse)`
  Sets rx1 band from band stack entry.
  Called by: `.Console_KeyDown()` (same file), `.OnEntryClicked()` (same file), `.preBandSelect()` (same file)
- **`.OnBandChangeHandler()`** — L46720 — `private void OnBandChangeHandler(int rx, Band oldBand, Band newBand)`
  Handles/raises the band change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnModeChangeHandler()`** — L46759 — `private void OnModeChangeHandler(int rx, DSPMode oldMode, DSPMode newMode, Band oldBand, Band newBand)`
  Handles/raises the mode change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOAFrequencyChangeHandler()`** — L46786 — `private void OnVFOAFrequencyChangeHandler(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double o`
  Handles/raises the vfoafrequency change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOBFrequencyChangeHandler()`** — L46807 — `private void OnVFOBFrequencyChangeHandler(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double o`
  Handles/raises the vfobfrequency change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMoxChangeHandler()`** — L46827 — `private void OnMoxChangeHandler(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateStackNumberDisplay()`** — L46854 — `private void updateStackNumberDisplay(BandStackFilter bsf)`
  Called by: `.handleBSFChange()` (same file), `.OnEntryAdd()` (same file), `.OnEntryUpdate()` (same file), `.OnEntryDelete()` (same file), `.OnEntryClicked()` (same file), `.preBandSelect()` (same file) — and 2 more
- **`.RepositionExternalPAButton()`** — L46881 — `public void RepositionExternalPAButton(bool bShow)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkExternalPA_CheckedChanged()`** — L46907 — `private void chkExternalPA_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkExternalPA` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setAGCThresholdPoint()`** — L47005 — `public void setAGCThresholdPoint(double agc_thresh_point, int rx)`
  Sets agcthreshold point.
  Called by: `.tmrAutoAGC_Tick()` (same file)
- **`.tmrAutoAGC_Tick()`** — L47102 — `private void tmrAutoAGC_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `tmrAutoAGC` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRF_Click()`** — L47164 — `private void ptbRF_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRF` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2RF_Click()`** — L47170 — `private void ptbRX2RF_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2RF` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRXEQ_MouseDown()`** — L47176 — `private void chkRXEQ_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRXEQ` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXEQ_MouseDown()`** — L47185 — `private void chkTXEQ_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkTXEQ` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFWCATUBypass_MouseDown()`** — L47194 — `private void chkFWCATUBypass_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkFWCATUBypass` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTUN_MouseDown()`** — L47199 — `private void chkTUN_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkTUN` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chk2TONE_MouseDown()`** — L47204 — `private void chk2TONE_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chk2TONE` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkExternalPA_MouseDown()`** — L47209 — `private void chkExternalPA_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkExternalPA` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRF_MouseDown()`** — L47214 — `private void ptbRF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbRF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2RF_MouseDown()`** — L47220 — `private void ptbRX2RF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbRX2RF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MouseWheelAGCRX1()`** — L47226 — `private void MouseWheelAGCRX1(object sender, System.Windows.Forms.MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseWheelAGCRX2()`** — L47231 — `private void MouseWheelAGCRX2(object sender, System.Windows.Forms.MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomToBand()`** — L47253 — `public void ZoomToBand(bool bStore)`
  Called by: `.btnDisplayZTB_Click()` (same file), `.DoOtherButtonAction()` (same file)
- **`.btnDisplayZTB_Click()`** — L47347 — `private void btnDisplayZTB_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDisplayZTB` is clicked.
  Called by: `.btnDisplayZTB_MouseUp()` (same file)
- **`.setupZTBButton()`** — L47354 — `private void setupZTBButton()`
  Called by: `.InitConsole()` (same file), `.chkX2TR_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.chkFWCATU_CheckedChanged()` (same file)
- **`.btnDisplayZTB_MouseUp()`** — L47362 — `private void btnDisplayZTB_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `btnDisplayZTB` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_Activated()`** — L47367 — `private void Console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Console_Deactivate()`** — L47374 — `private void Console_Deactivate(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_Button1Clicked()`** — L47390 — `private void infoBar_Button1Clicked(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_Button2Clicked()`** — L47395 — `private void infoBar_Button2Clicked(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleInfoBarButtonClick()`** — L47400 — `private void handleInfoBarButtonClick(ucInfoBar.InfoBarAction e)`
  Called by: `.infoBar_Button1Clicked()` (same file), `.infoBar_Button2Clicked()` (same file)
- **`.infoBar_Button1MouseDown()`** — L47436 — `private void infoBar_Button1MouseDown(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_Button2MouseDown()`** — L47441 — `private void infoBar_Button2MouseDown(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showSetupFromInfoBar()`** — L47446 — `private void showSetupFromInfoBar(ucInfoBar.ActionTypes action)`
  Called by: `.infoBar_Button1MouseDown()` (same file), `.infoBar_Button2MouseDown()` (same file)
- **`.infoBar_HideFeedbackChanged()`** — L47488 — `private void infoBar_HideFeedbackChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_SwapRedBlueChanged()`** — L47492 — `private void infoBar_SwapRedBlueChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateTuneLabel()`** — L47511 — `public void UpdateTuneLabel(bool bShowLimitValue, System.EventArgs e)`
  Updates tune label.
  Called by: `.OnTuneSliderUpdateTimerTick()` (same file), `.ptbTune_Scroll()` (same file), `.ptbTune_MouseUp()` (same file)
- **`.ptbTune_Scroll()`** — L47587 — `private void ptbTune_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbTune` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.setupTuneDriveSlider()`** — L47657 — `private void setupTuneDriveSlider()`
  Called by: `.InitConsole()` (same file), `.chkTUN_CheckedChanged()` (same file), `.chk2TONE_CheckedChanged()` (same file)
- **`.setPowerFromDriveSlider()`** — L47690 — `private int setPowerFromDriveSlider(out bool bConstrain, bool bAdjustedBySliderControl)`
  Sets power from drive slider.
  Called by: `.ptbPWR_Scroll()` (same file)
- **`.setPowerFromTuneSlider()`** — L47697 — `private int setPowerFromTuneSlider(out bool bConstrain, bool bAdjustedBySliderControl)`
  Sets power from tune slider.
  Called by: `.ptbTune_Scroll()` (same file)
- **`.SetPowerUsingTargetDBM()`** — L47704 — `public int SetPowerUsingTargetDBM(out bool bConstrain, out double targetdBm, bool bSetPower, bool bFromTune, bool bTwoTone)`
  Sets power using target dbm.
  Called by: `.chkTUN_CheckedChanged()` (same file), `.setPowerFromDriveSlider()` (same file), `.setPowerFromTuneSlider()` (same file)
- **`.enableAudioAmplfier()`** — L47885 — `private void enableAudioAmplfier()`
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.ptbTune_MouseUp()`** — L47893 — `private void ptbTune_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbTune` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ResetLevelCalibration()`** — L47897 — `public void ResetLevelCalibration(bool ignoreSet = false)`
  Resets level calibration.
  Called by: `.InitConsole()` (same file)
- **`.chkEnableMultiRX_MouseDown()`** — L47917 — `private void chkEnableMultiRX_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkEnableMultiRX` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MultiMeter2UpdateRX1()`** — L47938 — `private async void MultiMeter2UpdateRX1()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeter2UpdateRX2()`** — L48141 — `private async void MultiMeter2UpdateRX2()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateMetersReading()`** — L48305 — `private void updateMetersReading(Reading reading, float value, int rx)`
  Called by: `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.picMultiMeterDigital_Click()`** — L48310 — `private void picMultiMeterDigital_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `picMultiMeterDigital` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picRX2Meter_Click()`** — L48314 — `private void picRX2Meter_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `picRX2Meter` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucVAC1UnderOver_ClearIssuesClick()`** — L48318 — `private void ucVAC1UnderOver_ClearIssuesClick(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucVAC2UnderOver_ClearIssuesClick()`** — L48324 — `private void ucVAC2UnderOver_ClearIssuesClick(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ptbSquelch_Scroll()`** — L48331 — `private void ptbSquelch_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbSquelch` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkSquelch_CheckStateChanged()` (same file), `.pnlDisplay_DoubleClick()` (same file)
- **`.chkSquelch_CheckStateChanged()`** — L48401 — `private void chkSquelch_CheckStateChanged(object sender, EventArgs e)`
  Called by: `.SetRX1Mode()` (same file)
- **`.handleSqlFM()`** — L48500 — `private void handleSqlFM(int rx, bool bFM, SquelchState force_to_state = SquelchState.LAST)`
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file), `.SetSqlMode()` (same file)
- **`.chkRX2Squelch_CheckStateChanged()`** — L48583 — `private void chkRX2Squelch_CheckStateChanged(object sender, EventArgs e)`
  Called by: `.SetRX2Mode()` (same file)
- **`.ptbRX2Squelch_Scroll()`** — L48681 — `private void ptbRX2Squelch_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2Squelch` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkRX2Squelch_CheckStateChanged()` (same file)
- **`.chkSquelch_MouseDown()`** — L48749 — `private void chkSquelch_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkSquelch` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2Squelch_MouseDown()`** — L48754 — `private void chkRX2Squelch_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2Squelch` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOSplit_MouseClick()`** — L48759 — `private void chkVFOSplit_MouseClick(object sender, MouseEventArgs e)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.chkVFOSplit_MouseDown()`** — L48774 — `private void chkVFOSplit_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVFOSplit` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblPAProfile_MouseDown()`** — L48779 — `private void lblPAProfile_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblPAProfile` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateLegacyMeterControls()`** — L48811 — `private void updateLegacyMeterControls(bool expanded)`
  Called by: `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.InitFFTFillTime()`** — L48884 — `public void InitFFTFillTime(int rx)`
  Inits fftfill time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.finderMenuItem_Click()`** — L48920 — `private void finderMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `finderMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.setupCMasioStatusBar()`** — L48938 — `private void setupCMasioStatusBar()`
  Called by: `.UpdateStatusBarStatusIcons()` (same file)
- **`.setupSerialCatStatusBar()`** — L48964 — `private void setupSerialCatStatusBar()`
  Called by: `.UpdateStatusBarStatusIcons()` (same file)
- **`.UpdateStatusBarStatusIcons()`** — L49032 — `public void UpdateStatusBarStatusIcons(StatusBarIconGroup iconGroup)`
  Updates status bar status icons.
  Called by: `.SetupTCPIPCat()` (same file), `.SetupTCI()` (same file)
- **`.addStatusStripToolTipHandlers()`** — L49060 — `private void addStatusStripToolTipHandlers()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolTipItemMouseHover()`** — L49080 — `private void toolTipItemMouseHover(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolTipItemMouseLeave()`** — L49112 — `private void toolTipItemMouseLeave(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getRX1stepAttenuatorForBand()`** — L49120 — `private int getRX1stepAttenuatorForBand(Band b)`
  [2.10.3.6]MW0LGE moved all this to functions to make it easier to diagnose issues
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.SetComboPreampForHPSDR()` (same file)
- **`.setRX1stepAttenuatorForBand()`** — L49126 — `private void setRX1stepAttenuatorForBand(Band b, int att)`
  Sets rx1step attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.GetState()` (same file)
- **`.getRX2stepAttenuatorForBand()`** — L49131 — `private int getRX2stepAttenuatorForBand(Band b)`
  Returns rx2step attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.SetComboPreampForHPSDR()` (same file)
- **`.setRX2stepAttenuatorForBand()`** — L49137 — `private void setRX2stepAttenuatorForBand(Band b, int att)`
  Sets rx2step attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.GetState()` (same file)
- **`.getTXstepAttenuatorForBand()`** — L49142 — `private int getTXstepAttenuatorForBand(Band b)`
  Returns txstep attenuator for band.
  Called by: `.GetStateList()` (same file), `.chkPower_CheckedChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.setTXstepAttenuatorForBand()`** — L49148 — `private void setTXstepAttenuatorForBand(Band b, int att)`
  Sets txstep attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetState()` (same file)
- **`.udTXStepAttData_ValueChanged()`** — L49155 — `private void udTXStepAttData_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXStepAttData` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetAutoFormStartSetting()`** — L49177 — `public void SetAutoFormStartSetting(string form, bool show)`
  Sets auto form start setting.
  Called by: `.setAutoStartData()` (same file)
- **`.GetAutoFormStartSetting()`** — L49189 — `public bool GetAutoFormStartSetting(string form)`
  Returns auto form start setting.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getAutoStartData()`** — L49203 — `private string getAutoStartData()`
  Returns auto start data.
  Called by: `.GetStateList()` (same file)
- **`.setAutoStartData()`** — L49216 — `private void setAutoStartData(string data)`
  Sets auto start data.
  Called by: `.GetState()` (same file)
- **`.handleShowOnStartWindowsForms()`** — L49234 — `private void handleShowOnStartWindowsForms()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnOpenWindowsFormsTimerEvent()`** — L49249 — `private void OnOpenWindowsFormsTimerEvent(Object source, EventArgs e)`
  Handles/raises the open windows forms timer event event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showOnStartup()`** — L49274 — `private void showOnStartup(string form)`
  Called by: `.OnOpenWindowsFormsTimerEvent()` (same file), `.DoOtherButtonAction()` (same file)
- **`.handleLaunchOnStartUp()`** — L49329 — `private void handleLaunchOnStartUp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsProcessRunning()`** — L49390 — `private bool IsProcessRunning(string processName)`
  Called by: `.handleLaunchOnStartUp()` (same file)
- **`.FindAllWindowHandlesByProcessId()`** — L49394 — `private static List<IntPtr> FindAllWindowHandlesByProcessId(int processId)`
  Finds all window handles by process id.
  Called by: `.autoLaunchTryToClose()` (same file)
- **`.EnumWindows()`** — L49407 — `[DllImport("user32.dll", SetLastError = true)] private static extern bool EnumWindows(EnumWindowsProc lpEnumFunc, IntPtr lParam)`
  Called by: `.FindAllWindowHandlesByProcessId()` (same file)
- **`.GetWindowThreadProcessId()`** — L49409 — `[DllImport("user32.dll", SetLastError = true)] private static extern uint GetWindowThreadProcessId(IntPtr hWnd, out uint lpdwProcessId)`
  Returns window thread process id.
  Called by: `.FindAllWindowHandlesByProcessId()` (same file)
- **`.PostMessage()`** — L49411 — `[DllImport("user32.dll")] private static extern bool PostMessage(IntPtr hWnd, uint Msg, IntPtr wParam, IntPtr lParam)`
  Called by: `.autoLaunchTryToClose()` (same file)
- **`.PostThreadMessage()`** — L49413 — `[DllImport("user32.dll", SetLastError = true)] private static extern bool PostThreadMessage(uint idThread, uint Msg, IntPtr wParam, IntPtr lParam)`
  Called by: `.autoLaunchTryToClose()` (same file)
- **`.autoLaunchTryToClose()`** — L49420 — `private void autoLaunchTryToClose()`
  Called by: `.Console_Closing()` (same file)
- **`.databaseManagerToolStripMenuItem_Click()`** — L49543 — `private void databaseManagerToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `databaseManagerToolStripMenuItem` is clicked.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.setupToolStripMenuItem1_Click()`** — L49549 — `private void setupToolStripMenuItem1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `setupToolStripMenuItem1` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.ToggleRxTxAnt()`** — L49569 — `public void ToggleRxTxAnt()`
  Toggles rx tx ant.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PopupFilterContextMenu()`** — L49573 — `public void PopupFilterContextMenu(int rx, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PopupBandstack()`** — L49580 — `public void PopupBandstack(int rx, Band b, bool is_on_top)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showBandStack()`** — L49594 — `private void showBandStack()`
  Called by: `.regBox1_Click()` (same file), `.lblBandStack_Click()` (same file)
- **`.miAbout_Click()`** — L49620 — `private void miAbout_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `miAbout` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.resizeBackgroundImage()`** — L49721 — `private void resizeBackgroundImage()`
  Called by: `.Console_Resize()` (same file)
- **`.setupToolStripMenuItem_MouseUp()`** — L49764 — `private void setupToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `setupToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayControlsToolStripMenuItem_MouseUp()`** — L49794 — `private void displayControlsToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `displayControlsToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.dSPToolStripMenuItem_MouseUp()`** — L49799 — `private void dSPToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `dSPToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.bandToolStripMenuItem_MouseUp()`** — L49804 — `private void bandToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `bandToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.modeToolStripMenuItem_MouseUp()`** — L49809 — `private void modeToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `modeToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.filterToolStripMenuItem_MouseUp()`** — L49814 — `private void filterToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `filterToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rX2ToolStripMenuItem_MouseUp()`** — L49819 — `private void rX2ToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `rX2ToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BPFToolStripMenuItem_MouseUp()`** — L49824 — `private void BPFToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `BPFToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.GetMinimumRXNotchWidth()`** — L49831 — `public double GetMinimumRXNotchWidth(int rx)`
  Returns minimum rxnotch width.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMinimumTXNotchWidth()`** — L49838 — `public double GetMinimumTXNotchWidth()`
  Returns minimum txnotch width.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateMinimumNotchWidthRX()`** — L49842 — `public void UpdateMinimumNotchWidthRX(int rx)`
  Updates minimum notch width rx.
  Called by: `.UpdateDSP()` (same file)
- **`.UpdateMinimumNotchWidthTX()`** — L49874 — `public void UpdateMinimumNotchWidthTX()`
  Updates minimum notch width tx.
  Called by: `.UpdateDSP()` (same file)
- **`.chkFWCATU_MouseUp()`** — L49895 — `private void chkFWCATU_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkFWCATU` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkX2TR_MouseUp()`** — L49904 — `private void chkX2TR_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkX2TR` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.WaterfallRXGradient()`** — L49913 — `public Color[] WaterfallRXGradient()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaterfallTXGradient()`** — L49919 — `public Color[] WaterfallTXGradient()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clampFilterMinMax()`** — L49926 — `private void clampFilterMinMax(int rx, bool use_lowHigh = false, int low = 0, int high = 0)`
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRX2Filters()` (same file)
- **`.clampFilterShift()`** — L49952 — `private void clampFilterShift(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.pnlDisplay_DoubleClick()`** — L49965 — `private void pnlDisplay_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlDisplay` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseDown()`** — L49993 — `private void pnlDisplay_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlDisplay` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseLeave()`** — L50811 — `private void pnlDisplay_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlDisplay` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseMove()`** — L50837 — `unsafe private void pnlDisplay_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlDisplay` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseUp()`** — L52080 — `private void pnlDisplay_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlDisplay` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_Resize()`** — L52211 — `private async void pnlDisplay_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlDisplay` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupDisplayMaxBinDetect()`** — L52238 — `private void setupDisplayMaxBinDetect(int rx, bool sub_rx, bool enabled, bool update_enabled_state = true)`
  Called by: `.chkRIT_CheckedChanged()` (same file), `.udRIT_ValueChanged()` (same file), `.OnCentreFrequencyChanged()` (same file), `.OnCTUNChanged()` (same file), `.OnVFOAFrequencyChangeHandler()` (same file), `.OnVFOBFrequencyChangeHandler()` (same file) — and 6 more
- **`.OnFilterEdgesChanged()`** — L52288 — `private void OnFilterEdgesChanged(int rx, Filter newFilter, Band band, int low, int high, string sName, int max_width, int max_shift)`
  Handles/raises the filter edges changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSampleRateChanged()`** — L52293 — `private void OnSampleRateChanged(int rx, int oldSampleRate, int newSampleRate)`
  Handles/raises the sample rate changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFSPChanged()`** — L52298 — `private void OnFSPChanged(int old_fpr, int new_fps)`
  Handles/raises the fspchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkNR_Click()`** — L52304 — `private void chkNR_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNR` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NR_Click()`** — L52309 — `private void chkRX2NR_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2NR` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.requires_reposition()`** — L52314 — `private bool requires_reposition()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.on_send_floodcontrol_message()`** — L52347 — `private void on_send_floodcontrol_message(string msg, string uid)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoOtherButtonAction()`** — L52392 — `public void DoOtherButtonAction(int rx, OtherButtonId id, MouseButtons button, bool force = false, bool current_state = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleDoOtherButtonActionRightClick()`** — L52678 — `private bool handleDoOtherButtonActionRightClick(int rx, OtherButtonId id)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetNFEnabled()`** — L52815 — `public void SetNFEnabled(int rx, bool state)`
  Sets nfenabled.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetNFEnabled()`** — L52830 — `public bool GetNFEnabled(int rx)`
  Returns nfenabled.
  Called by: `.DoOtherButtonAction()` (same file), `.GetGeneralSetting()` (same file)
- **`.SetBandStack()`** — L52844 — `public void SetBandStack(int rx, int dir)`
  Sets band stack.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAgcT()`** — L52858 — `public int GetAgcT(int rx)`
  Returns agc t.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetAgcT()`** — L52871 — `public void SetAgcT(int rx, int value)`
  Sets agc t.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSql()`** — L52890 — `public int GetSql(int rx)`
  Returns sql.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetSql()`** — L52903 — `public void SetSql(int rx, int value)`
  Sets sql.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetBal()`** — L52922 — `public int GetBal(int rx, bool subrx = false)`
  Returns bal.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetBal()`** — L52938 — `public void SetBal(int rx, int value, bool subrx = false)`
  Sets bal.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAF()`** — L52960 — `public int GetAF(int rx, bool subrx = false)`
  Returns af.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetAF()`** — L52976 — `public bool SetAF(int rx, int value, bool subrx = false)`
  Sets af.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetDisplayZoomGeneralSettings()`** — L53008 — `public DisplayZoomButton GetDisplayZoomGeneralSettings(int rx)`
  Returns display zoom general settings.
  Called by: `.GetGeneralSetting()` (same file)
- **`.SetDisplayZoomGeneralSettings()`** — L53017 — `public void SetDisplayZoomGeneralSettings(int rx, DisplayZoomButton dzb)`
  Sets display zoom general settings.
  Called by: `.ptbDisplayZoom_Scroll()` (same file)
- **`.SetPanAdjust()`** — L53048 — `private void SetPanAdjust(int adjust, bool centre = false)`
  Sets pan adjust.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetOtherButtonState()`** — L53053 — `public bool GetOtherButtonState(OtherButtonId id, int rx)`
  Returns other button state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetSqlMode()`** — L53177 — `public bool SetSqlMode(int rx, SquelchState state)`
  Sets sql mode.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSqlMode()`** — L53210 — `public SquelchState GetSqlMode(int rx)`
  Returns sql mode.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file), `.SetSqlMode()` (same file)
- **`.GetPanSwap()`** — L53223 — `public bool GetPanSwap(int rx)`
  Returns pan swap.
  Called by: `.GetOtherButtonState()` (same file)
- **`.GetSubRX()`** — L53237 — `public bool GetSubRX(int rx)`
  Returns sub rx.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetBin()`** — L53250 — `public bool GetBin(int rx)`
  Returns bin.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetMute()`** — L53263 — `public bool GetMute(int rx)`
  Returns mute.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file), `.SetMute()` (same file)
- **`.GetSelectedNB()`** — L53286 — `public int GetSelectedNB(int rx)`
  Returns selected nb.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetSelectedNB()`** — L53327 — `public bool SetSelectedNB(int rx, int nb)`
  Sets selected nb.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSplit()`** — L53358 — `public bool GetSplit(int rx)`
  Returns split.
  Called by: `.GetOtherButtonState()` (same file)
- **`.GetMNF()`** — L53372 — `public bool GetMNF(int rx)`
  Returns mnf.
  Called by: `.GetOtherButtonState()` (same file)
- **`.GetANF()`** — L53386 — `public bool GetANF(int rx)`
  Returns anf.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetANF()`** — L53399 — `public bool SetANF(int rx, bool state)`
  Sets anf.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSNB()`** — L53414 — `public bool GetSNB(int rx)`
  Returns snb.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetSNB()`** — L53428 — `public bool SetSNB(int rx, bool state)`
  Sets snb.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAVG()`** — L53443 — `public bool GetAVG(int rx)`
  Returns avg.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetPeak()`** — L53456 — `public bool GetPeak(int rx)`
  Returns peak.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetCTUN()`** — L53469 — `public bool GetCTUN(int rx)`
  Returns ctun.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetAGCMode()`** — L53482 — `public AGCMode GetAGCMode(int rx)`
  Returns agcmode.
  Called by: `.GetOtherButtonState()` (same file)
- **`.SetAGCMode()`** — L53495 — `public bool SetAGCMode(int rx, AGCMode mode)`
  Sets agcmode.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAGCAuto()`** — L53510 — `public bool GetAGCAuto(int rx)`
  Returns agcauto.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetAGCAuto()`** — L53523 — `public bool SetAGCAuto(int rx, bool state)`
  Sets agcauto.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetAVG()`** — L53538 — `public bool SetAVG(int rx, bool state)`
  Sets avg.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetPeak()`** — L53554 — `public bool SetPeak(int rx, bool state)`
  Sets peak.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetCTUN()`** — L53570 — `public bool SetCTUN(int rx, bool state)`
  Sets ctun.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetSubRX()`** — L53586 — `public bool SetSubRX(int rx, bool state)`
  Sets sub rx.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetBin()`** — L53602 — `public bool SetBin(int rx, bool state)`
  Sets bin.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.setMuteAllGeneralSettings()`** — L53617 — `private void setMuteAllGeneralSettings()`
  Sets mute all general settings.
  Called by: `.chkMUT_CheckedChanged()` (same file), `.chkRX2Mute_CheckedChanged()` (same file)
- **`.SetMute()`** — L53621 — `public bool SetMute(int rx, bool state)`
  Sets mute.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetXPAStatus()`** — L53653 — `public (bool in_use, bool enabled) GetXPAStatus()`
  Returns xpastatus.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetDisplayMode()`** — L53658 — `public DisplayMode GetDisplayMode(int rx)`
  Returns display mode.
  Called by: `.GetOtherButtonState()` (same file)
- **`.SetDisplayMode()`** — L53670 — `public void SetDisplayMode(int rx, DisplayMode mode)`
  Sets display mode.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetGeneralSetting()`** — L53724 — `public bool GetGeneralSetting(int rx, OtherButtonId id)`
  Returns general setting.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file), `.initGeneralSettings()` (same file)
- **`.initGeneralSettings()`** — L53822 — `private void initGeneralSettings(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetGeneralSetting()`** — L53873 — `public void SetGeneralSetting(int rx, OtherButtonId id, bool state)`
  Sets general setting.
  Called by: `.chkMicMute_CheckedChanged()` (same file), `.chkVOX_CheckedChanged()` (same file), `.chkNoiseGate_CheckedChanged()` (same file), `.chkShowTXFilter_CheckedChanged()` (same file), `.chkRXEQ_CheckedChanged()` (same file), `.chkTXEQ_CheckedChanged()` (same file) — and 11 more
- **`.DoGeneralSettingAction()`** — L53896 — `public bool DoGeneralSettingAction(int rx, OtherButtonId id, bool state)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetHWSampleRateSetting()`** — L53985 — `public void SetHWSampleRateSetting(int rx, int rate)`
  Sets hwsample rate setting.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetATT()`** — L54020 — `public int GetATT(int rx)`
  Returns att.
  Called by: `.GetGeneralSetting()` (same file)
- **`.maxAtt()`** — L54074 — `private int maxAtt()`
  Called by: `.SetATT()` (same file), `.IncrementATT()` (same file)
- **`.SetATT()`** — L54100 — `public bool SetATT(int rx, int att, SetAttMode mode)`
  Sets att.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.setATTGeneralSetting()`** — L54176 — `private void setATTGeneralSetting(int rx)`
  Sets attgeneral setting.
  Called by: `.initGeneralSettings()` (same file)
- **`.IncrementATT()`** — L54262 — `public void IncrementATT(int rx)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.DecrementATT()`** — L54353 — `public void DecrementATT(int rx)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.handleVfoSyncInitial()`** — L54462 — `private void handleVfoSyncInitial()`
  Called by: `.chkVFOSync_CheckedChanged()` (same file)
- **`.handleVfoSyncFrequency()`** — L54502 — `private void handleVfoSyncFrequency(int rx, bool b_to_a)`
  Called by: `.OnVFOAFrequencyChangeHandler()` (same file), `.OnVFOBFrequencyChangeHandler()` (same file)
- **`.handleVfoSyncMode()`** — L54536 — `private void handleVfoSyncMode(int rx, DSPMode mode)`
  Called by: `.OnModeChangeHandler()` (same file), `.handleVfoSyncFrequency()` (same file)
- **`.handleVfoSyncFilter()`** — L54570 — `private void handleVfoSyncFilter(int rx, Filter newFilter)`
  Called by: `.OnFilterChanged()` (same file), `.handleVfoSyncMode()` (same file)
- **`.btnAPF_type_Click()`** — L54589 — `private void btnAPF_type_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAPF_type` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAPF_type_MouseDown()`** — L54596 — `private void btnAPF_type_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `btnAPF_type` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilter_rx1_MouseUp()`** — L54601 — `private void radFilter_rx1_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radFilter_rx1` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilter_rx2_MouseUp()`** — L54606 — `private void radFilter_rx2_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radFilter_rx2` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripStatusLabel_PAstatus_MouseUp()`** — L54613 — `private void toolStripStatusLabel_PAstatus_MouseUp(object sender, MouseEventArgs e)`
  Support for Ganymede PA status
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.arp_PowerChanged()`** — L54690 — `private void arp_PowerChanged(bool old_power, bool new_power)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.arp_PlayingingChanged()`** — L54698 — `private void arp_PlayingingChanged(bool playing, string id, string filename, bool isWdsp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaveRecording()`** — L54736 — `public bool WaveRecording(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.arp_RecordingChanged()`** — L54741 — `private void arp_RecordingChanged(bool recording, string id, string filename)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setPlayRecordStatusBar()`** — L54793 — `private void setPlayRecordStatusBar()`
  Sets play record status bar.
  Called by: `.arp_PlayingingChanged()` (same file), `.arp_RecordingChanged()` (same file)
- **`.waveRecord()`** — L54811 — `private void waveRecord(int rx, bool recording)`
  Called by: `.DoOtherButtonAction()` (same file)

#### `ztb_data` (type, L206)

_No extracted members._

#### `SpectralResult` (type, L21306)

_No extracted members._

#### `HistoricAttenuatorReading` (type, L21365)

_No extracted members._

#### `MeasureKey` (type, L22431)

- **`.Equals()`** — L22446 — `public bool Equals(MeasureKey other) => Text == other.Text && Font.Equals(other.Font)`
  Called by: `.GetState()` (same file), `.SafeTXProfileSet()` (same file), `.Console_KeyDown()` (same file), `.txtVFOAFreq_KeyPress()` (same file), `.txtVFOABand_KeyPress()` (same file), `.txtVFOBFreq_KeyPress()` (same file)
- **`.GetHashCode()`** — L22455 — `public override int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `AutoTuneState` (type, L25703)

_No extracted members._

#### `ProtocolEvent` (type, L25713)

_No extracted members._

#### `TuneLocation` (type, L31846)

_No extracted members._

#### `ModeSpecificPanel` (type, L34753)

_No extracted members._

#### `DisplayZoomButton` (type, L53000)

_No extracted members._

#### `SetAttMode` (type, L54094)

_No extracted members._

#### `DigiMode` (type, L54896)

_No extracted members._

#### `DigiModeSettingState` (type, L54902)

_No extracted members._

#### `AsyncLock` (type, L54922)

- **`.Dispose()`** — L54932 — `public void Dispose()`
  Releases the object’s resources.
  Called by: `.Dispose()` (same file), `.resizeBackgroundImage()` (same file)

#### `MessageFloodControl` (type, L54939)

- **`.FloodControl()`** — L54950 — `public static void FloodControl(string message, string uid, bool ignore_flood = false)`
  Called by: `.BroadcastFreqChange()` (same file), `.BroadcastVFOChange()` (same file)
- **`.Shutdown()`** — L55011 — `public static void Shutdown()`
  Called by: `.Console_Closing()` (same file)
- **`.timer_callback()`** — L55042 — `static void timer_callback(object state_obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.raise_send_message()`** — L55075 — `static void raise_send_message(string message, string uid)`
  Called by: `.FloodControl()` (same file), `.timer_callback()` (same file)

#### `State` (type, L55100)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/console.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
