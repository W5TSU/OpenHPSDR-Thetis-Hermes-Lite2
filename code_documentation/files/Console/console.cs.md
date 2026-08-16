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

- **`.LockAsync()`** — L55045 — `public async Task<AsyncLock> LockAsync()`
  Called by: `.ATUTune()` (same file)

### Types

#### `Console` (type, L89)

- **`.SendMessage()`** — L582 — `[DllImport("user32.dll")] private static extern int SendMessage(IntPtr hWnd, Int32 wMsg, bool wParam, Int32 lParam)`
  Sends message.
  Called by: `.SuspendDrawing()` (same file), `.ResumeDrawing()` (same file)
- **`.RedrawWindow()`** — L584 — `[DllImport("user32.dll")] static extern bool RedrawWindow(IntPtr hWnd, IntPtr lprcUpdate, IntPtr hrgnUpdate, uint flags)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initialisePortAudio()`** — L1134 — `private void initialisePortAudio()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowReleaseNotes()`** — L1168 — `public void ShowReleaseNotes()`
  Shows release notes.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SuspendDrawing()`** — L1206 — `private void SuspendDrawing(Control control)`
  MW0LGE helper functions to suspend drawing/painting of controls. This is to help limit flicker during resize of console window
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file), `.pnlResizeMeter_MouseMove()` (same file)
- **`.ResumeDrawing()`** — L1212 — `private void ResumeDrawing(Control control, bool refresh = true)`
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file), `.pnlResizeMeter_MouseMove()` (same file)
- **`.OnAutoStartTimerEvent()`** — L1220 — `private void OnAutoStartTimerEvent(Object source, ElapsedEventArgs e)`
  Handles/raises the auto start timer event event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.gmh_MouseUp()`** — L1244 — `void gmh_MouseUp()`
  WinForms event handler: runs when `gmh` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Dispose()`** — L1249 — `protected override void Dispose(bool disposing)`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AttachConsole()`** — L1301 — `[DllImport("kernel32.dll")] private static extern bool AttachConsole(int dwProcessId)`
  Called by: `.showHelpInfo()` (same file)
- **`.FreeConsole()`** — L1303 — `[DllImport("kernel32.dll")] static extern bool FreeConsole()`
  Called by: `.showHelpInfo()` (same file)
- **`.showHelpInfo()`** — L1308 — `private static bool showHelpInfo()`
  Called by: `.Main()` (same file)
- **`.Application_ThreadException()`** — L1358 — `static void Application_ThreadException(object sender, ThreadExceptionEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CurrentDomain_UnhandledException()`** — L1363 — `static void CurrentDomain_UnhandledException(object sender, UnhandledExceptionEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cleanArg()`** — L1374 — `static string cleanArg(string input)`
  Called by: `.Main()` (same file)
- **`.Main()`** — L1398 — `[STAThread] static void Main(string[] args)`
  [DllImport("shcore.dll")] private static extern int SetProcessDpiAwareness(int awareness); Main
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.build_exception_text()`** — L1546 — `static string build_exception_text(Exception ex)`
  Called by: `.Main()` (same file)
- **`.getTitleWithFWVersion()`** — L1627 — `private string getTitleWithFWVersion()`
  Returns title with fwversion.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onNotchDelete()`** — L1669 — `private void onNotchDelete(int notch_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onBWChanged()`** — L1677 — `private void onBWChanged(int notch_index, double width)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onActiveChanged()`** — L1685 — `private void onActiveChanged(int notch_index, bool active)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onClearButton()`** — L1693 — `private void onClearButton()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitConsole()`** — L1699 — `private void InitConsole()`
  Inits console.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupInfoBarButton()`** — L2248 — `public void SetupInfoBarButton(ucInfoBar.ActionTypes action, bool bEnabled)`
  Setups info bar button.
  Called by: `.chkRX2_CheckedChanged()` (same file)
- **`.SetupInfoBar()`** — L2252 — `private void SetupInfoBar()`
  Setups info bar.
  Called by: `.InitConsole()` (same file)
- **`.InfoBarFeedbackLevel()`** — L2280 — `public void InfoBarFeedbackLevel(int level, bool bFeedbackLevelOk, bool bCorrectionsBeingApplied, bool bCalibrationAttemptsChanged, Color feedbackColour)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowTCPIPCatLog()`** — L2312 — `public void ShowTCPIPCatLog()`
  Shows tcpipcat log.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatClientConnect()`** — L2316 — `private void OnTCPIIPcatClientConnect()`
  Handles/raises the tcpiipcat client connect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatClientDisconnect()`** — L2322 — `private void OnTCPIIPcatClientDisconnect()`
  Handles/raises the tcpiipcat client disconnect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatClientError()`** — L2328 — `private void OnTCPIIPcatClientError(SocketException se)`
  Handles/raises the tcpiipcat client error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCPIIPcatServerError()`** — L2332 — `private void OnTCPIIPcatServerError(SocketException se)`
  Handles/raises the tcpiipcat server error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addTCPIPcatDelegates()`** — L2337 — `private void addTCPIPcatDelegates()`
  Called by: `.SetupTCPIPCat()` (same file)
- **`.removeTCPIPcatDelegates()`** — L2346 — `private void removeTCPIPcatDelegates()`
  Called by: `.SetupTCPIPCat()` (same file), `.Console_Closing()` (same file)
- **`.SetupTCPIPCat()`** — L2369 — `public void SetupTCPIPCat(bool bOn)`
  Setups tcpipcat.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.closeTcpIpCatServer()`** — L2415 — `private void closeTcpIpCatServer()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addTCIDelegates()`** — L2445 — `private void addTCIDelegates()`
  Called by: `.SetupTCI()` (same file)
- **`.removeTCIDelegates()`** — L2454 — `private void removeTCIDelegates()`
  Called by: `.SetupTCI()` (same file), `.Console_Closing()` (same file)
- **`.OnTCIClientConnect()`** — L2465 — `private void OnTCIClientConnect()`
  Handles/raises the tciclient connect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCIClientDisconnect()`** — L2471 — `private void OnTCIClientDisconnect()`
  Handles/raises the tciclient disconnect event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCIClientError()`** — L2477 — `private void OnTCIClientError(SocketException se)`
  Handles/raises the tciclient error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCIServerError()`** — L2481 — `private void OnTCIServerError(SocketException se)`
  Handles/raises the tciserver error event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowTCILog()`** — L2485 — `public void ShowTCILog()`
  Shows tcilog.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupTCI()`** — L2489 — `public void SetupTCI(bool bOn, int rateLimit)`
  Setups tci.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.selectFilters()`** — L2539 — `private void selectFilters()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.selectModes()`** — L2585 — `private void selectModes()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Init60mChannels()`** — L2611 — `public void Init60mChannels()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SyncDSP()`** — L2646 — `private void SyncDSP()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExitConsole()`** — L2695 — `public void ExitConsole()`
  Called by: `.Dispose()` (same file)
- **`.StateListToBase64()`** — L2753 — `public string StateListToBase64()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetStateList()`** — L2759 — `public List<string> GetStateList()`
  Returns state list.
  Called by: `.StateListToBase64()` (same file), `.SaveState()` (same file)
- **`.addControlState()`** — L3243 — `private void addControlState(List<string> a, Control c, List<string> always_save, bool combo_use_items_count, bool print_not_saving)`
  Called by: `.GetStateList()` (same file)
- **`.SaveState()`** — L3295 — `public void SaveState()`
  Saves state.
  Called by: `.Console_Closing()` (same file)
- **`.GetState()`** — L3889 — `public void GetState()`
  Returns state.
  Called by: `.InitConsole()` (same file)
- **`.InitFilterPresets()`** — L5147 — `private void InitFilterPresets()`
  Inits filter presets.
  Called by: `.InitConsole()` (same file), `.toolStripMenuItemRX1FilterReset_Click()` (same file), `.toolStripMenuItemRX2FilterReset_Click()` (same file)
- **`.InitDisplayModes()`** — L5550 — `private void InitDisplayModes()`
  Inits display modes.
  Called by: `.InitConsole()` (same file)
- **`.InitAGCModes()`** — L5561 — `private void InitAGCModes()`
  Inits agcmodes.
  Called by: `.InitConsole()` (same file)
- **`.InitMultiMeterModes()`** — L5579 — `private void InitMultiMeterModes()`
  Inits multi meter modes.
  Called by: `.InitConsole()` (same file)
- **`.DisableAllFilters()`** — L5611 — `private void DisableAllFilters()`
  Disables all filters.
  Called by: `.CalibratePAGain()` (same file), `.SetRX1Mode()` (same file)
- **`.DisableAllRX2Filters()`** — L5623 — `private void DisableAllRX2Filters()`
  Disables all rx2 filters.
  Called by: `.SetRX2Mode()` (same file)
- **`.EnableAllFilters()`** — L5634 — `private void EnableAllFilters()`
  Enables all filters.
  Called by: `.CalibratePAGain()` (same file), `.SetRX1Mode()` (same file)
- **`.EnableAllRX2Filters()`** — L5651 — `private void EnableAllRX2Filters()`
  Enables all rx2 filters.
  Called by: `.SetRX2Mode()` (same file)
- **`.DisableAllBands()`** — L5667 — `private void DisableAllBands()`
  Disables all bands.
  Called by: `.UIMOXChangedTrue()` (same file)
- **`.EnableAllBands()`** — L5711 — `private void EnableAllBands()`
  Enables all bands.
  Called by: `.UIMOXChangedFalse()` (same file)
- **`.DisableAllModes()`** — L5760 — `private void DisableAllModes()`
  Disables all modes.
  Called by: `.CalibratePAGain()` (same file), `.UIMOXChangedTrue()` (same file)
- **`.EnableAllModes()`** — L5770 — `private void EnableAllModes()`
  Enables all modes.
  Called by: `.CalibratePAGain()` (same file), `.UIMOXChangedFalse()` (same file)
- **`.GetVFOCharWidth()`** — L5781 — `private void GetVFOCharWidth()`
  Returns vfochar width.
  Called by: `.Console_MouseWheel()` (same file), `.txtVFOAFreq_MouseMove()` (same file), `.txtVFOBFreq_MouseMove()` (same file)
- **`.GetVFOSubCharWidth()`** — L5814 — `private void GetVFOSubCharWidth()`
  Returns vfosub char width.
  Called by: `.txtVFOABand_MouseMove()` (same file)
- **`.SetBand()`** — L5840 — `public void SetBand(string mode, string filter, double freq, bool CTUN, int zoomFactor, double centerFreq)`
  Sets band.
  Called by: `.setRX1BandFromBandStackEntry()` (same file)
- **`.getButtonForBand()`** — L5949 — `private RadioButtonTS getButtonForBand(Band b)`
  Returns button for band.
  Called by: `.SetBand()` (same file)
- **`.ChangeTuneStepUp()`** — L6089 — `public void ChangeTuneStepUp()`
  Called by: `.InitConsole()` (same file), `.CATTuneStepUp()` (same file), `.Console_KeyDown()` (same file), `.WheelTune_MouseDown()` (same file), `.btnChangeTuneStepLarger_Click()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.ChangeTuneStepDown()`** — L6095 — `public void ChangeTuneStepDown()`
  Called by: `.CATTuneStepDown()` (same file), `.Console_KeyDown()` (same file), `.btnChangeTuneStepSmaller_Click()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.UpdateBandButtonColors()`** — L6101 — `private void UpdateBandButtonColors()`
  Updates band button colors.
  Called by: `.SetRX1Band()` (same file), `.SetRX2Band()` (same file), `.SetTXBand()` (same file), `.UpdateWaterfallLevelValues()` (same file)
- **`.DeselectHF()`** — L6106 — `private void DeselectHF()`
  Called by: `.SetRX1BandButton()` (same file)
- **`.DeselectGEN()`** — L6125 — `private void DeselectGEN()`
  Called by: `.SetRX1BandButton()` (same file)
- **`.DeselectVHF()`** — L6144 — `private void DeselectVHF()`
  Called by: `.SetRX1BandButton()` (same file)
- **`.SetRX1BandButton()`** — L6162 — `private void SetRX1BandButton(Band b)`
  Sets rx1 band button.
  Called by: `.UpdateBandButtonColors()` (same file)
- **`.RX1BandForVFOB()`** — L6381 — `public Band RX1BandForVFOB()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTransverterTranslatedRXBand()`** — L6402 — `public Band GetTransverterTranslatedRXBand(double freq)`
  Returns transverter translated rxband.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTransverterTranslatedTXBand()`** — L6409 — `public Band GetTransverterTranslatedTXBand()`
  Returns transverter translated txband.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandByFreq()`** — L6416 — `private Band BandByFreq(double freq, int xvtr_index, FRSRegion region)`
  Called by: `.SetBand()` (same file), `.RX1BandForVFOB()` (same file), `.GetTransverterTranslatedRXBand()` (same file), `.GetTransverterTranslatedTXBand()` (same file), `.VFOASubUpdate()` (same file), `.HdwMOXChanged()` (same file) — and 10 more
- **`.SetRX1Band()`** — L6428 — `private void SetRX1Band(Band b)`
  Sets rx1 band.
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.SetRX2Band()`** — L6463 — `private void SetRX2Band(Band b)`
  Sets rx2 band.
  Called by: `.txtVFOBFreq_LostFocus()` (same file)
- **`.SetTXBand()`** — L6475 — `private void SetTXBand(Band b, bool bIngoreBandChange = false)`
  Sets txband.
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.GainByBand()`** — L6495 — `private float GainByBand(Band b, int nDriveValue)`
  Called by: `.SetPowerUsingTargetDBM()` (same file)
- **`.CheckSelectedButtonColor()`** — L6502 — `public void CheckSelectedButtonColor()`
  Checks selected button color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PABandOffset()`** — L6572 — `private double PABandOffset(Band b)`
  Called by: `.ScaledVoltage()` (same file), `.ADCtodBm()` (same file)
- **`.SWR()`** — L6615 — `private double SWR(int adc_fwd, int adc_rev)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ScaledVoltage()`** — L6630 — `private double ScaledVoltage(int adc)`
  Called by: `.SWR()` (same file), `.PAPower()` (same file)
- **`.ADCtodBm()`** — L6637 — `private double ADCtodBm(int adc_data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PAPower()`** — L6646 — `private double PAPower(int adc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WattsTodBm()`** — L6654 — `private double WattsTodBm(double watts)`
  Called by: `.CalibratePAGain()` (same file)
- **`.dBmToWatts()`** — L6659 — `private double dBmToWatts(double dBm)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibratedPAPower()`** — L6664 — `public float CalibratedPAPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.PowerKernel()`** — L6730 — `private float PowerKernel(float watts, float interval, int entries, float[] table)`
  Called by: `.CalibratedPAPower()` (same file)
- **`.CheckValidTXFreq()`** — L6744 — `public bool CheckValidTXFreq(FRSRegion r, double f, DSPMode mode, bool bIgnoreFilter = false)`
  Checks valid txfreq.
  Called by: `.chkMOX_CheckedChanged2()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.checkValidTXFreq_local()`** — L6782 — `private bool checkValidTXFreq_local(FRSRegion r, double f)`
  Called by: `.CheckValidTXFreq()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.setAlex1HPF()`** — L6793 — `private void setAlex1HPF(double freq)`
  Sets alex1 hpf.
  Called by: `.UpdateRX1DDSFreq()` (same file), `.UpdateTXDDSFreq()` (same file), `.UpdateAlexRXFilter()` (same file)
- **`.setAlexHPF()`** — L6805 — `private void setAlexHPF(double freq)`
  Sets alex hpf.
  Called by: `.setAlex1HPF()` (same file)
- **`.setBPF1ForOrionIISaturn()`** — L6919 — `private void setBPF1ForOrionIISaturn(double freq)`
  Sets bpf1 for orion iisaturn.
  Called by: `.setAlex1HPF()` (same file)
- **`.setAlex2HPF()`** — L7035 — `private void setAlex2HPF(double freq)`
  Sets alex2 hpf.
  Called by: `.UpdateRX2DDSFreq()` (same file)
- **`.setAlexLPF()`** — L7143 — `private void setAlexLPF(double freq, bool freqIsTX)`
  Sets alex lpf.
  Called by: `.UpdateTXDDSFreq()` (same file), `.UpdateAlexTXFilter()` (same file)
- **`.setAlex2LPF()`** — L7211 — `private void setAlex2LPF(double freq)`
  Sets alex2 lpf.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectRX1VarFilter()`** — L7273 — `public void SelectRX1VarFilter(bool update = true, bool prevent_update = false)`
  Selects rx1 var filter.
  Called by: `.Console_KeyDown()` (same file), `.ptbFilterShift_Scroll()` (same file), `.ptbFilterWidth_Scroll()` (same file), `.pnlDisplay_MouseMove()` (same file), `.ExecuteEncoderStep()` (`Console/Andromeda/Andromeda.cs`)
- **`.SelectRX2VarFilter()`** — L7286 — `public void SelectRX2VarFilter(bool update = true, bool prevent_update = false)`
  Selects rx2 var filter.
  Called by: `.pnlDisplay_MouseMove()` (same file), `.ExecuteEncoderStep()` (`Console/Andromeda/Andromeda.cs`)
- **`.UpdateRXADCCtrlP1()`** — L7300 — `public void UpdateRXADCCtrlP1()`
  Updates rxadcctrl p1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMemoryQS()`** — L7306 — `public void CATMemoryQS()`
  Added 06/24/05 BT for CAT commands
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMemoryQR()`** — L7312 — `public void CATMemoryQR()`
  Added 06/25/05 BT for CAT commands
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCATBand()`** — L7361 — `public void SetCATBand(Band pBand)`
  BT 06/17/05 added for CAT commands
  Called by: `.HandleFrontPanelButtonPress()` (`Console/Andromeda/Andromeda.cs`), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.SetVHFText()`** — L7458 — `public void SetVHFText(int index, string text)`
  Sets vhftext.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetVHFEnabled()`** — L7465 — `public void SetVHFEnabled(int index, bool b)`
  Sets vhfenabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetVHFText()`** — L7473 — `public string GetVHFText(int index)`
  G8NJJ added to allow labelling of buttons in popup form
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetVHFEnabled()`** — L7477 — `public bool GetVHFEnabled(int index)`
  Returns vhfenabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1Filters()`** — L7493 — `public void UpdateRX1Filters(int low, int high, bool force = false, bool from_change_event = false)`
  Updates rx1 filters.
  Called by: `.SelectRX1VarFilter()` (same file), `.UpdateRX1FilterPresetLow()` (same file), `.UpdateRX1FilterPresetHigh()` (same file), `.Console_KeyDown()` (same file), `.SetRX1Mode()` (same file), `.SetRX1Filter()` (same file) — and 9 more
- **`.UpdateRX2Filters()`** — L7616 — `public void UpdateRX2Filters(int low, int high, bool force = false, bool from_change_event = false)`
  Updates rx2 filters.
  Called by: `.SelectRX2VarFilter()` (same file), `.UpdateRX2FilterPresetLow()` (same file), `.UpdateRX2FilterPresetHigh()` (same file), `.SetRX2Mode()` (same file), `.SetRX2Filter()` (same file), `.udRX2FilterLow_ValueChanged()` (same file) — and 5 more
- **`.UpdateRX1FilterNames()`** — L7714 — `public void UpdateRX1FilterNames(Filter f, string old_name, string new_name)`
  Updates rx1 filter names.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1FilterPresetLow()`** — L7762 — `public void UpdateRX1FilterPresetLow(int val)`
  Updates rx1 filter preset low.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1FilterPresetHigh()`** — L7767 — `public void UpdateRX1FilterPresetHigh(int val)`
  Updates rx1 filter preset high.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2FilterNames()`** — L7772 — `public void UpdateRX2FilterNames(Filter f, string old_name, string new_name)`
  Updates rx2 filter names.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2FilterPresetLow()`** — L7811 — `public void UpdateRX2FilterPresetLow(int val)`
  Updates rx2 filter preset low.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2FilterPresetHigh()`** — L7816 — `public void UpdateRX2FilterPresetHigh(int val)`
  Updates rx2 filter preset high.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateVFOAFreq()`** — L7821 — `public void UpdateVFOAFreq(string freq)`
  Updates vfoafreq.
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.sioPut()`** — L7832 — `private bool sioPut(object sio_listener, string msg)`
  Called by: `.on_send_floodcontrol_message()` (same file)
- **`.BroadcastFreqChange()`** — L7889 — `private void BroadcastFreqChange(string vfo, double freq)`
  Called by: `.OnVFOAFrequencyChangeHandler()` (same file), `.OnVFOBFrequencyChangeHandler()` (same file)
- **`.UpdateVFOBFreq()`** — L7898 — `public void UpdateVFOBFreq(string freq)`
  Updates vfobfreq.
  Called by: `.txtVFOBFreq_LostFocus()` (same file)
- **`.CalcDisplayFreq()`** — L7909 — `public void CalcDisplayFreq()`
  Called by: `.InitConsole()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ptbDisplayPan_Scroll()` (same file), `.ptbDisplayZoom_Scroll()` (same file), `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file) — and 1 more
- **`.CalcRX2DisplayFreq()`** — L7921 — `public void CalcRX2DisplayFreq()`
  Called by: `.InitConsole()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ptbDisplayPan_Scroll()` (same file), `.ptbDisplayZoom_Scroll()` (same file)
- **`.CalcTXDisplayFreq()`** — L7932 — `public void CalcTXDisplayFreq()`
  Called by: `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ptbDisplayPan_Scroll()` (same file), `.ptbDisplayZoom_Scroll()` (same file)
- **`.UpdateRXSpectrumDisplayVars()`** — L7938 — `public void UpdateRXSpectrumDisplayVars()`
  Updates rxspectrum display vars.
  Called by: `.chkDisplayAVG_CheckedChanged()` (same file), `.chkDisplayPeak_CheckedChanged()` (same file), `.UpdateDSP()` (same file), `.pnlDisplay_Resize()` (same file)
- **`.UpdateTXSpectrumDisplayVars()`** — L7943 — `public void UpdateTXSpectrumDisplayVars()`
  Updates txspectrum display vars.
  Called by: `.comboDisplayMode_SelectedIndexChanged()` (same file), `.chkDisplayAVG_CheckedChanged()` (same file), `.chkDisplayPeak_CheckedChanged()` (same file), `.UpdateDSP()` (same file), `.pnlDisplay_Resize()` (same file)
- **`.UpdateRXDisplayVars()`** — L7948 — `private void UpdateRXDisplayVars(int l, int h)`
  Updates rxdisplay vars.
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRXSpectrumDisplayVars()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.comboRX2DisplayMode_SelectedIndexChanged()` (same file)
- **`.UpdateTXDisplayVars()`** — L7990 — `private void UpdateTXDisplayVars(int l, int h)`
  Updates txdisplay vars.
  Called by: `.UpdateTXSpectrumDisplayVars()` (same file), `.SetTXFilters()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file)
- **`.UpdateTXLowHighFilterForMode()`** — L8026 — `public void UpdateTXLowHighFilterForMode(DSPMode mode, ref int low, ref int high)`
  Updates txlow high filter for mode.
  Called by: `.SetTXFilters()` (same file), `.SetTXFilter()` (same file)
- **`.SetTXFilters()`** — L8066 — `public void SetTXFilters(DSPMode mode, int low, int high, bool force = false)`
  Sets txfilters.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file), `.fmDeviation2k()` (same file), `.fmDeviation5k()` (same file)
- **`.BuildTXProfileCombos()`** — L8100 — `public void BuildTXProfileCombos()`
  Builds txprofile combos.
  Called by: `.InitConsole()` (same file), `.UpdateTXProfile()` (same file)
- **`.UpdateTXProfile()`** — L8114 — `public void UpdateTXProfile(string name)`
  Updates txprofile.
  Called by: `.InitConsole()` (same file)
- **`.UpdateDDCs()`** — L8161 — `public void UpdateDDCs(bool rx2_enabled)`
  Diversity operation is on RX1; therefore, the 'rx1_rate' will be used as the diversity rate;
  Called by: `.SetupForHPSDRModel()` (same file), `.chkPower_CheckedChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.txtVFOAFreq_LostFocus()` (same file)
- **`.GetDDC()`** — L8527 — `public void GetDDC(out int DDCrx1, out int DDCrx2, out int DDCsync1, out int DDCsync2, out int DDCpsrx, out int DDCpstx)`
  Returns ddc.
  Called by: `.handleOverload()` (same file), `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.UpdateDiversityMenuItem()`** — L8793 — `public void UpdateDiversityMenuItem()`
  Updates diversity menu item.
  Called by: `.showHideDiversity()` (same file)
- **`.UpdateDiversityValues()`** — L8801 — `private void UpdateDiversityValues()`
  Updates diversity values.
  Called by: `.SetRX1Band()` (same file), `.showHideDiversity()` (same file), `.setRX1BandFromBandStackEntry()` (same file)
- **`.UpdateWaterfallLevelValues()`** — L8883 — `public void UpdateWaterfallLevelValues()`
  Updates waterfall level values.
  Called by: `.SetRX1Band()` (same file), `.SetRX2Band()` (same file), `.setRX1BandFromBandStackEntry()` (same file)
- **`.updateDisplayGridLevelValues()`** — L9080 — `private void updateDisplayGridLevelValues()`
  Called by: `.SetRX1Band()` (same file), `.SetRX2Band()` (same file), `.setRX1BandFromBandStackEntry()` (same file)
- **`.setWaterfallGainsIfLinkedToSpectrum()`** — L9085 — `private void setWaterfallGainsIfLinkedToSpectrum(int rx)`
  Sets waterfall gains if linked to spectrum.
  Called by: `.UpdateWaterfallLevelValues()` (same file), `.UpdateDisplayGridLevelMinValues()` (same file), `.UpdateDisplayGridLevelMaxValues()` (same file)
- **`.CheckForMinMaxGridUpdatesTX()`** — L9100 — `public void CheckForMinMaxGridUpdatesTX()`
  Checks for min max grid updates tx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CheckForMinMaxGridUpdatesRX()`** — L9122 — `public void CheckForMinMaxGridUpdatesRX(int rx)`
  Checks for min max grid updates rx.
  Called by: `.UpdateDisplayGridLevelMinValues()` (same file), `.UpdateDisplayGridLevelMaxValues()` (same file)
- **`.CheckForMinMaxWaterfallUpdatesTX()`** — L9179 — `public void CheckForMinMaxWaterfallUpdatesTX()`
  Checks for min max waterfall updates tx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CheckForMinMaxWaterfallUpdatesRX()`** — L9201 — `public void CheckForMinMaxWaterfallUpdatesRX(int rx)`
  Checks for min max waterfall updates rx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateDisplayGridLevelMinValues()`** — L9258 — `public void UpdateDisplayGridLevelMinValues(bool bDoBandInfoAndWaterFallSync)`
  Updates display grid level min values.
  Called by: `.updateDisplayGridLevelValues()` (same file)
- **`.UpdateDisplayGridLevelMaxValues()`** — L9395 — `public void UpdateDisplayGridLevelMaxValues(bool bDoBandInfoAndWaterFallSync)`
  Updates display grid level max values.
  Called by: `.updateDisplayGridLevelValues()` (same file)
- **`.RX1IsIn60m()`** — L9533 — `public bool RX1IsIn60m()`
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.RX1IsOn60mChannel()`** — L9539 — `public bool RX1IsOn60mChannel(Channel c)`
  Called by: `.UpdatePeakText()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.SetRX1Mode()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.RX1IsIn60mChannel()`** — L9546 — `public bool RX1IsIn60mChannel(Channel c)`
  Called by: `.SetRX1Mode()` (same file)
- **`.RX2IsIn60m()`** — L9573 — `public bool RX2IsIn60m()`
  Called by: `.txtVFOBFreq_LostFocus()` (same file)
- **`.RX2IsOn60mChannel()`** — L9579 — `public bool RX2IsOn60mChannel(Channel c)`
  Called by: `.SetRX2Mode()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.RX2IsIn60mChannel()`** — L9586 — `public bool RX2IsIn60mChannel(Channel c)`
  Called by: `.SetRX2Mode()` (same file)
- **`.ModeFreqOffset()`** — L9612 — `private double ModeFreqOffset(DSPMode mode)`
  Called by: `.RX1IsOn60mChannel()` (same file), `.RX2IsOn60mChannel()` (same file), `.UpdatePeakText()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file), `.SetRX1Mode()` (same file) — and 1 more
- **`.CATGetXVTRBandNames()`** — L9634 — `public string CATGetXVTRBandNames()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetFilterPresets()`** — L9647 — `public string GetFilterPresets(int mode_ndx)`
  Returns filter presets.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.radBandGEN0_Click()`** — L9670 — `private void radBandGEN0_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN0` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN1_Click()`** — L9676 — `private void radBandGEN1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN2_Click()`** — L9682 — `private void radBandGEN2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN3_Click()`** — L9688 — `private void radBandGEN3_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN3` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN4_Click()`** — L9694 — `private void radBandGEN4_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN4` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN5_Click()`** — L9700 — `private void radBandGEN5_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN6_Click()`** — L9706 — `private void radBandGEN6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN6` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN7_Click()`** — L9712 — `private void radBandGEN7_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN7` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN8_Click()`** — L9718 — `private void radBandGEN8_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN8` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN9_Click()`** — L9724 — `private void radBandGEN9_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN9` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN10_Click()`** — L9730 — `private void radBandGEN10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN10` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN11_Click()`** — L9736 — `private void radBandGEN11_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN11` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN12_Click()`** — L9742 — `private void radBandGEN12_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN12` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBandGEN13_Click()`** — L9748 — `private void radBandGEN13_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN13` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CalibrateFreq()`** — L9758 — `unsafe public bool CalibrateFreq(float freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibrateLevel()`** — L9835 — `unsafe public bool CalibrateLevel(float level, float freq, Progress progress, bool suppress_errors)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibratePAGain()`** — L10220 — `public bool CalibratePAGain(Progress progress, bool[] run, int target_watts)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalibratePAGain2()`** — L10381 — `public bool CalibratePAGain2(Progress progress, bool[] run, bool suppress_warnings)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LowPowerPASweep()`** — L10386 — `public bool LowPowerPASweep(Progress progress, int power)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.validateRX1StepAttData()`** — L10919 — `private int validateRX1StepAttData(int att)`
  [2.10.3.9]MW0LGE validate step attenuator values, so many HL2 issues with these being out of range
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.validateRX2StepAttData()`** — L10925 — `private int validateRX2StepAttData(int att)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.validateTXStepAttData()`** — L10931 — `private int validateTXStepAttData(int att)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TuneStepLookup()`** — L11314 — `public int TuneStepLookup(string s)`
  Called by: `.RecallMemory()` (same file)
- **`.isDSPModeValid()`** — L11324 — `private bool isDSPModeValid(DSPMode mode)`
  Called by: `.updateStepIndexForMode()` (same file)
- **`.updateStepIndexForMode()`** — L11329 — `private void updateStepIndexForMode(int rx, DSPMode mode)`
  Called by: `.OnModeChangeHandler()` (same file)
- **`.SetAFLinks()`** — L11422 — `public void SetAFLinks(int source, bool state)`
  Sets aflinks.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsMasterAFLinked()`** — L11447 — `public bool IsMasterAFLinked(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setLinkedAF()`** — L11454 — `private void setLinkedAF(int source, int value)`
  Sets linked af.
  Called by: `.SetAFLinks()` (same file), `.ptbAF_Scroll()` (same file), `.ptbRX0Gain_Scroll()` (same file), `.ptbRX1Gain_Scroll()` (same file), `.ptbRX2Gain_Scroll()` (same file)
- **`.HighlightTXProfileSaveItems()`** — L11954 — `public void HighlightTXProfileSaveItems(bool bHighlight)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPower()`** — L12145 — `public void SetPower(Band b, int pwr)`
  Sets power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPower()`** — L12151 — `public int GetPower(Band b)`
  Returns power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX1DisplayOffsets()`** — L12331 — `private void UpdateRX1DisplayOffsets()`
  Updates rx1 display offsets.
  Called by: `.InitConsole()` (same file), `.CalibrateLevel()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ResetLevelCalibration()` (same file), `.udTXStepAttData_ValueChanged()` (same file)
- **`.UpdateRX2DisplayOffsets()`** — L12339 — `private void UpdateRX2DisplayOffsets()`
  Updates rx2 display offsets.
  Called by: `.InitConsole()` (same file), `.CalibrateLevel()` (same file), `.comboDisplayMode_SelectedIndexChanged()` (same file), `.ResetLevelCalibration()` (same file), `.udTXStepAttData_ValueChanged()` (same file)
- **`.SafeTXProfileSet()`** — L12428 — `public void SafeTXProfileSet(string profile)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCWSideToneVolume()`** — L13071 — `private void setCWSideToneVolume()`
  Sets cwside tone volume.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupDisplayEngine()`** — L13562 — `public void SetupDisplayEngine(int decimation)`
  MW0LGE_21k9
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupForHPSDRModel()`** — L14808 — `public void SetupForHPSDRModel()`
  Setups for hpsdrmodel.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateOCTXPins()`** — L14951 — `private void updateOCTXPins(bool tx)`
  Called by: `.OnMoxChangeHandler()` (same file)
- **`.UpdateTRXAnt()`** — L15000 — `private void UpdateTRXAnt()`
  Updates trxant.
  Called by: `.HdwMOXChanged()` (same file), `.txtVFOAFreq_LostFocus()` (same file)
- **`.enableMONForCW()`** — L15047 — `private void enableMONForCW()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetADCInUse()`** — L15153 — `public int GetADCInUse(int ddc)`
  Returns adcin use.
  Called by: `.handleOverload()` (same file), `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.SetWavePlayback()`** — L15322 — `public void SetWavePlayback(int id, bool enabled)`
  Sets wave playback.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getWavePlayback()`** — L15342 — `private bool getWavePlayback(int id)`
  Returns wave playback.
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.adjustForSnapClickTuning()` (same file)
- **`.getWavePlaybackFreq()`** — L15347 — `public double getWavePlaybackFreq(int id)`
  Returns wave playback freq.
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.UpdateRX1DDSFreq()`** — L15441 — `private void UpdateRX1DDSFreq()`
  Updates rx1 ddsfreq.
  Called by: `.HdwMOXChanged()` (same file)
- **`.UpdateRX2DDSFreq()`** — L15471 — `private void UpdateRX2DDSFreq()`
  Updates rx2 ddsfreq.
  Called by: `.HdwMOXChanged()` (same file)
- **`.UpdateTXDDSFreq()`** — L15509 — `private void UpdateTXDDSFreq()`
  Updates txddsfreq.
  Called by: `.HdwMOXChanged()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.UpdateAlexTXFilter()`** — L15532 — `private void UpdateAlexTXFilter()`
  Updates alex txfilter.
  Called by: `.UpdateRX1DDSFreq()` (same file), `.UpdateRX2DDSFreq()` (same file)
- **`.UpdateAlexRXFilter()`** — L15545 — `private void UpdateAlexRXFilter()`
  Updates alex rxfilter.
  Called by: `.UpdateRX1DDSFreq()` (same file), `.UpdateRX2DDSFreq()` (same file)
- **`.ThreadSafeCatParse()`** — L15702 — `public string ThreadSafeCatParse(string msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.safeCat()`** — L15711 — `private string safeCat(string msg)`
  Called by: `.ThreadSafeCatParse()` (same file)
- **`.CATVFOAtoB()`** — L15843 — `public void CATVFOAtoB()`
  -W2PA Added three new functions to make CAT functions match behavior of equivalent console functions. i.e. not just copy frequency alone
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATVFOBtoA()`** — L15847 — `public void CATVFOBtoA()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATVFOABSwap()`** — L15851 — `public void CATVFOABSwap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATTuneStepUp()`** — L16086 — `public void CATTuneStepUp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATTuneStepDown()`** — L16091 — `public void CATTuneStepDown()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMidiMessagesPerTuneStepUp()`** — L16145 — `public void CATMidiMessagesPerTuneStepUp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMidiMessagesPerTuneStepDown()`** — L16150 — `public void CATMidiMessagesPerTuneStepDown()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATMidiMessagesPerTuneStepToggle()`** — L16155 — `public void CATMidiMessagesPerTuneStepToggle()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATSingleCal()`** — L16214 — `public void CATSingleCal()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATRX2BandUpDown()`** — L16441 — `public void CATRX2BandUpDown(int direction)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandToString()`** — L17416 — `private string BandToString(Band b)`
  Called by: `.SetupRX2Band()` (same file)
- **`.StringToBand()`** — L17469 — `private Band StringToBand(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadSigStrength()`** — L17672 — `public string CATReadSigStrength()`
  Added 07/30/05 BT for cat commands next 8 functions
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadAvgStrength()`** — L17679 — `public string CATReadAvgStrength()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadADC_L()`** — L17686 — `public string CATReadADC_L()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadADC_R()`** — L17692 — `public string CATReadADC_R()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadALC()`** — L17698 — `public string CATReadALC()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadFwdPwr()`** — L17710 — `public string CATReadFwdPwr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadPeakPwr()`** — L17726 — `public string CATReadPeakPwr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadRevPwr()`** — L17762 — `public string CATReadRevPwr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CATReadSWR()`** — L17768 — `public string CATReadSWR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXFilter()`** — L17908 — `public bool SetTXFilter(Filter filter)`
  Sets txfilter.
  Called by: `.SetTXFilters()` (same file), `.MatchTXFilterToRXFilter()` (same file)
- **`.GetDSPcwPitchShiftToZero()`** — L18248 — `public int GetDSPcwPitchShiftToZero(int rx)`
  Returns dspcw pitch shift to zero.
  Called by: `.AddNotch()` (same file)
- **`.freqFromString()`** — L18384 — `static double freqFromString(string s)`
  Called by: `.UpdateVFOAFreq()` (same file), `.UpdateVFOBFreq()` (same file), `.btnMemoryQuickRestore_Click()` (same file)
- **`.VFOAUpdate()`** — L18416 — `private void VFOAUpdate(double freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOBUpdate()`** — L18422 — `private void VFOBUpdate(double freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOASubUpdate()`** — L18428 — `private void VFOASubUpdate(double freq)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PanCentre()`** — L19671 — `public void PanCentre()`
  Called by: `.InitConsole()` (same file), `.zoomToBandBandwidth()` (same file), `.displayZoom05()` (same file), `.displayZoom1()` (same file), `.displayZoom2()` (same file), `.displayZoom4()` (same file) — and 1 more
- **`.ZoomFullyOut()`** — L19675 — `public void ZoomFullyOut()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTimer()`** — L19845 — `private void SetTimer(System.Windows.Forms.Timer t, bool enable)`
  Sets timer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CpuUsage()`** — L20783 — `private void CpuUsage()`
  Called by: `.systemToolStripMenuItem_Click()` (same file), `.thetisOnlyToolStripMenuItem_Click()` (same file)
- **`.disableCpuVoltsUsage()`** — L20822 — `private void disableCpuVoltsUsage()`
  Called by: `.CpuUsage()` (same file)
- **`.isBitSet()`** — L21093 — `private static bool isBitSet(int n, int pos)`
  Called by: `.checkSeqErrors()` (same file)
- **`.ShowSEQLog()`** — L21098 — `public void ShowSEQLog()`
  Shows seqlog.
  Called by: `.toolStripStatusLabel_SeqWarning_Click()` (same file)
- **`.RXPreampOffset()`** — L21116 — `public float RXPreampOffset(int rx)`
  Called by: `.UpdateRX1DisplayOffsets()` (same file), `.UpdateRX2DisplayOffsets()` (same file), `.RXOffset()` (same file)
- **`.RXCalibrationOffset()`** — L21149 — `public float RXCalibrationOffset(int rx)`
  Called by: `.UpdateRX1DisplayOffsets()` (same file), `.UpdateRX2DisplayOffsets()` (same file), `.RXOffset()` (same file)
- **`.RXOffset()`** — L21167 — `public float RXOffset(int rx)`
  Called by: `.CATReadSigStrength()` (same file), `.CATReadAvgStrength()` (same file), `.RXPBsnr()` (same file), `.UpdatePeakText()` (same file), `.UpdateMultimeter()` (same file), `.UpdateRX2Multimeter()` (same file) — and 6 more
- **`.RXPBsnr()`** — L21225 — `public double RXPBsnr(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.spectralCalculations()`** — L21283 — `private SpectralResult spectralCalculations(int rx, double signal)`
  Called by: `.RXPBsnr()` (same file), `.UpdatePeakText()` (same file), `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.checkOverloadsAndSync()`** — L21437 — `private async void checkOverloadsAndSync()`
  Called by: `.pollOverloadSyncSeqErr()` (same file)
- **`.keep_att_entries_for_band()`** — L21564 — `private void keep_att_entries_for_band(Stack<HistoricAttenuatorReading> readings_stack, Band target_band)`
  Called by: `.handleOverload()` (same file)
- **`.handleOverload()`** — L21582 — `private void handleOverload()`
  Called by: `.checkOverloadsAndSync()` (same file)
- **`.pollOverloadSyncSeqErr()`** — L21883 — `private async void pollOverloadSyncSeqErr()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkSeqErrors()`** — L21938 — `private void checkSeqErrors()`
  Called by: `.pollOverloadSyncSeqErr()` (same file)
- **`.UpdatePeakText()`** — L22012 — `private void UpdatePeakText()`
  Updates peak text.
  Called by: `.timer_peak_text_Tick()` (same file)
- **`.HzInNPixels()`** — L22194 — `private int HzInNPixels(int nPixelCount, int rx)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.getLowHighForRXn()`** — L22203 — `private void getLowHighForRXn(int rx, out int low, out int high, bool bIncludeRitXit = true)`
  Returns low high for rxn.
  Called by: `.HzInNPixels()` (same file), `.PixelToHz()` (same file), `.HzToPixel()` (same file)
- **`.PixelToHz()`** — L22297 — `private float PixelToHz(float x)`
  Called by: `.UpdatePeakText()` (same file), `.getFrequencyAtPixel()` (same file), `.pnlDisplay_MouseDown()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.HzToPixel()`** — L22311 — `private int HzToPixel(float freq)`
  Called by: `.getFilterEdgesInPixels()` (same file), `.pnlDisplay_MouseDown()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.PixelToDb()`** — L22343 — `private float PixelToDb(float y)`
  Called by: `.pnlDisplay_DoubleClick()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.PixelToRx2Db()`** — L22348 — `private float PixelToRx2Db(float y)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.WaterfallPixelToTime()`** — L22360 — `private float WaterfallPixelToTime(float y, int rx)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.measureStringFromCache()`** — L22435 — `private SizeF measureStringFromCache(string str, Font font, int width, StringFormat format, Graphics g)`
  Called by: `.GetVFOCharWidth()` (same file), `.GetVFOSubCharWidth()` (same file), `.getMeterPixelPosAndDrawScales()` (same file)
- **`.getMeterPixelPosAndDrawScales()`** — L22445 — `private void getMeterPixelPosAndDrawScales(int rx, Graphics g, int H, int W, double num, out int pixel_x, out int pixel_x_swr, int nStringOffsetY, bool bDrawMarkers)`
  Returns meter pixel pos and draw scales.
  Called by: `.picMultiMeterDigital_Paint()` (same file), `.picRX2Meter_Paint()` (same file)
- **`.storeRX1SignalPixels_X()`** — L23561 — `private void storeRX1SignalPixels_X(float x)`
  Called by: `.picMultiMeterDigital_Paint()` (same file)
- **`.storeRX2SignalPixels_X()`** — L23578 — `private void storeRX2SignalPixels_X(float x)`
  Called by: `.picRX2Meter_Paint()` (same file)
- **`.clearRXSignalPixels()`** — L23596 — `private void clearRXSignalPixels(int rx)`
  Called by: `.picMultiMeterDigital_Paint()` (same file), `.picRX2Meter_Paint()` (same file), `.ResetMultiMeterPeak()` (same file), `.ResetRX2MeterPeak()` (same file), `.OnPowerChangeHander()` (same file), `.OnBandChangeHandler()` (same file) — and 2 more
- **`.picMultiMeterDigital_Paint()`** — L23632 — `private void picMultiMeterDigital_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picMultiMeterDigital` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picRX2Meter_Paint()`** — L23937 — `private void picRX2Meter_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picRX2Meter` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ResetMultiMeterPeak()`** — L24139 — `private void ResetMultiMeterPeak()`
  Resets multi meter peak.
  Called by: `.UIMOXChangedTrue()` (same file), `.UIMOXChangedFalse()` (same file), `.comboMeterRXMode_SelectedIndexChanged()` (same file), `.comboMeterTXMode_SelectedIndexChanged()` (same file)
- **`.ResetRX2MeterPeak()`** — L24146 — `private void ResetRX2MeterPeak()`
  Resets rx2 meter peak.
  Called by: `.comboRX2MeterMode_SelectedIndexChanged()` (same file)
- **`.panelVFOAHover_Paint()`** — L24153 — `private void panelVFOAHover_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `panelVFOAHover` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panelVFOBHover_Paint()`** — L24185 — `private void panelVFOBHover_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `panelVFOBHover` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.resetWDSPdisplayBuffers()`** — L24232 — `private void resetWDSPdisplayBuffers(int rx, bool tx)`
  Called by: `.RunDisplay()` (same file)
- **`.RunDisplay()`** — L24281 — `unsafe private void RunDisplay()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateMultimeter()`** — L24710 — `private async void UpdateMultimeter()`
  Updates multimeter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2Multimeter()`** — L24870 — `private async void UpdateRX2Multimeter()`
  Updates rx2 multimeter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeHermesDCVoltage()`** — L24921 — `public float computeHermesDCVoltage()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.readMKIIPAVoltsAmps()`** — L24940 — `private async void readMKIIPAVoltsAmps()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeMKIIPAVoltsAmps()`** — L25087 — `private void computeMKIIPAVoltsAmps()`
  Called by: `.timer_cpu_volts_meter_Tick()` (same file)
- **`.convertToVolts()`** — L25102 — `private float convertToVolts(float IOreading)`
  Called by: `.readMKIIPAVoltsAmps()` (same file), `.computeMKIIPAVoltsAmps()` (same file)
- **`.convertToAmps()`** — L25132 — `private float convertToAmps(float IOreading)`
  Called by: `.readMKIIPAVoltsAmps()` (same file), `.computeMKIIPAVoltsAmps()` (same file)
- **`.computeRefPower()`** — L25161 — `public float computeRefPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeAlexFwdPower()`** — L25247 — `public float computeAlexFwdPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeExciterPower()`** — L25319 — `public float computeExciterPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeOrionMkIIExciterPower()`** — L25378 — `public float computeOrionMkIIExciterPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeOrionExciterPower()`** — L25437 — `public float computeOrionExciterPower()`
  Called by: `.PollPAPWR()` (same file)
- **`.computeANANExciterPower()`** — L25496 — `public float computeANANExciterPower()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeHermesLiteTemp()`** — L25557 — `public void computeHermesLiteTemp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.computeHermesLitePAAmps()`** — L25572 — `public void computeHermesLitePAAmps()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateSQL()`** — L25588 — `private async void UpdateSQL()`
  Updates sql.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateRX2SQL()`** — L25605 — `private async void UpdateRX2SQL()`
  Updates rx2 sql.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateNoiseGate()`** — L25622 — `private async void UpdateNoiseGate()`
  Updates noise gate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIOBoardAerialPorts()`** — L25636 — `public void SetIOBoardAerialPorts(int rx_only_ant, int rx_ant, int tx_ant, bool tx)`
  Sets ioboard aerial ports.
  Called by: `.modifyXVTRantenna()` (same file)
- **`.SetI2CPollingPause()`** — L25663 — `public void SetI2CPollingPause( bool pause )`
  Sets i2 cpolling pause.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AutoTuningHL2()`** — L25694 — `bool AutoTuningHL2(ProtocolEvent protocolEvent)`
  Called by: `.UpdateIOBoard()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.UpdateIOBoard()`** — L25803 — `private async void UpdateIOBoard()`
  Updates ioboard.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateVOX()`** — L25969 — `private async void UpdateVOX()`
  Updates vox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getFallbackPTTModeAfterTCIRelease()`** — L25991 — `private PTTMode getFallbackPTTModeAfterTCIRelease(DSPMode tx_mode, bool mic_ptt, bool cw_ptt, bool cat_ptt, bool vox_ptt)`
  Returns fallback pttmode after tcirelease.
  Called by: `.PollPTT()` (same file)
- **`.PollPTT()`** — L26025 — `private async void PollPTT()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PollCW()`** — L26190 — `private async void PollCW()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cwAutoModeTick()`** — L26222 — `private void cwAutoModeTick(object o)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.update_for_auto_mode_return()`** — L26261 — `private void update_for_auto_mode_return(bool enabled)`
  Called by: `.PollCW()` (same file)
- **`.UpdatePreamps()`** — L26345 — `private void UpdatePreamps()`
  Updates preamps.
  Called by: `.comboPreamp_SelectedIndexChanged()` (same file), `.comboRX2Preamp_SelectedIndexChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.txtVFOAFreq_LostFocus()` (same file)
- **`.PollTXInhibit()`** — L26411 — `private async void PollTXInhibit()`
  bool audio_amp_mute;
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PollN1MMPacket()`** — L26452 — `private async void PollN1MMPacket()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.n1mm_delay_Elapsed()`** — L26486 — `private void n1mm_delay_Elapsed(object sender, ElapsedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleXml()`** — L26522 — `void HandleXml(string str)`
  Handles xml.
  Called by: `.PollN1MMPacket()` (same file)
- **`.ToggleFocusMasterTimer()`** — L26533 — `private void ToggleFocusMasterTimer()`
  Toggles focus master timer.
  Called by: `.gmh_MouseUp()` (same file), `.n1mm_delay_Elapsed()` (same file), `.Console_KeyUp()` (same file), `.Console_MouseWheel()` (same file), `.SetFocusMaster()` (same file)
- **`.PollPAPWR()`** — L26543 — `private async void PollPAPWR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkAntennaWarning()`** — L26739 — `private void checkAntennaWarning()`
  Called by: `.PollPAPWR()` (same file)
- **`.SWRScale()`** — L26761 — `private double SWRScale(double ref_pow)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.timer_cpu_volts_meter_Tick()`** — L26774 — `private void timer_cpu_volts_meter_Tick(object sender, System.EventArgs e)`
  WinForms event handler: runs when `timer_cpu_volts_meter` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.timer_peak_text_Tick()`** — L26841 — `private void timer_peak_text_Tick(object sender, System.EventArgs e)`
  WinForms event handler: runs when `timer_peak_text` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.timer_clock_Tick()`** — L26846 — `private void timer_clock_Tick(object sender, System.EventArgs e)`
  WinForms event handler: runs when `timer_clock` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_KeyPress()`** — L26876 — `private void Console_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `Console` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_KeyUp()`** — L26886 — `private void Console_KeyUp(object sender, System.Windows.Forms.KeyEventArgs e)`
  WinForms event handler: runs when `Console` receives a key-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.spacebarHoldEngaged()`** — L26905 — `private void spacebarHoldEngaged()`
  Called by: `.Console_KeyDown()` (same file)
- **`.spacebarHoldRelease()`** — L26921 — `private void spacebarHoldRelease()`
  Called by: `.Console_KeyUp()` (same file), `.spacebarHoldEngaged()` (same file)
- **`.enableOutsideSpectral()`** — L26947 — `private void enableOutsideSpectral()`
  Called by: `.Console_KeyDown()` (same file), `.OnMouseWheelChanged()` (same file)
- **`.restoreOutsideSpectral()`** — L26952 — `private void restoreOutsideSpectral()`
  Called by: `.Console_KeyDown()` (same file), `.OnMouseWheelChanged()` (same file)
- **`.Console_KeyDown()`** — L26960 — `private void Console_KeyDown(object sender, System.Windows.Forms.KeyEventArgs e)`
  WinForms event handler: runs when `Console` receives a key-down.
  Called by: `.ProcessDialogKey()` (same file)
- **`.setupLegacyMeterThreads()`** — L27826 — `private void setupLegacyMeterThreads(int rx)`
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.chkPower_CheckedChanged()`** — L27858 — `private void chkPower_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPower` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateAAudioMixerStates()`** — L28311 — `unsafe public void UpdateAAudioMixerStates()`
  MW0LGE [2.9.0.8] re-implemented by Warren
  Called by: `.chkPower_CheckedChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.comboDisplayMode_SelectedIndexChanged()`** — L28464 — `public void comboDisplayMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDisplayMode` selection changes.
  Called by: `.SetupDisplayEngine()` (same file), `.chkPower_CheckedChanged()` (same file)
- **`.chkBIN_CheckedChanged()`** — L28683 — `private void chkBIN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkBIN` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAGC_SelectedIndexChanged()`** — L28697 — `private void comboAGC_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboAGC` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.run_console_closing_handlers_async()`** — L28802 — `private Task run_console_closing_handlers_async()`
  Called by: `.Console_Closing()` (same file)
- **`.Console_Closing()`** — L28816 — `private void Console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `Console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getErrorLogSize()`** — L29050 — `private long getErrorLogSize()`
  Returns error log size.
  Called by: `.Dispose()` (same file)
- **`.shutdownLogStringToPath()`** — L29060 — `private void shutdownLogStringToPath(string entry)`
  Called by: `.Dispose()` (same file), `.ExitConsole()` (same file), `.Console_Closing()` (same file)
- **`.removeShutdownLog()`** — L29074 — `private void removeShutdownLog()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboPreamp_SelectedIndexChanged()`** — L29083 — `private void comboPreamp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboPreamp` selection changes.
  Called by: `.SetComboPreampForHPSDR()` (same file)
- **`.comboRX2Preamp_SelectedIndexChanged()`** — L29150 — `private void comboRX2Preamp_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2Preamp` selection changes.
  Called by: `.SetComboPreampForHPSDR()` (same file)
- **`.chkMUT_CheckedChanged()`** — L29226 — `private void chkMUT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMUT` checked state changes.
  Called by: `.InitConsole()` (same file), `.SetRX1Mode()` (same file)
- **`.ModelIsHPSDRorHermes()`** — L29270 — `public bool ModelIsHPSDRorHermes()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDriveSliderUpdateTimerTick()`** — L29286 — `private void OnDriveSliderUpdateTimerTick(object sender, ElapsedEventArgs e)`
  Handles/raises the drive slider update timer tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTuneSliderUpdateTimerTick()`** — L29290 — `private void OnTuneSliderUpdateTimerTick(object sender, ElapsedEventArgs e)`
  Handles/raises the tune slider update timer tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateDriveLabel()`** — L29296 — `public void UpdateDriveLabel(bool bShowLimitValue, System.EventArgs e)`
  Updates drive label.
  Called by: `.OnDriveSliderUpdateTimerTick()` (same file), `.ptbPWR_MouseUp()` (same file), `.ptbPWR_Scroll()` (same file)
- **`.ptbPWR_MouseUp()`** — L29398 — `private void ptbPWR_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbPWR` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbPWR_Scroll()`** — L29403 — `private void ptbPWR_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbPWR` is scrolled.
  Called by: `.InitConsole()` (same file), `.checkOverloadsAndSync()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.ptbAF_Scroll()`** — L29442 — `private void ptbAF_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbAF` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file), `.chkMUT_CheckedChanged()` (same file), `.chkMON_CheckedChanged()` (same file), `.AudioMOXChanged()` (same file)
- **`.ptbRF_Scroll()`** — L29490 — `private void ptbRF_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRF` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.chkMicMute_CheckedChanged()`** — L29539 — `private void chkMicMute_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMicMute` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbMic_Scroll()`** — L29546 — `private void ptbMic_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbMic` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file), `.chkMicMute_CheckedChanged()` (same file), `.SetRX1Mode()` (same file), `.radModeButton_CheckedChanged()` (same file), `.SetRX2Mode()` (same file) — and 3 more
- **`.setAudioMicGain()`** — L29600 — `private void setAudioMicGain(double gain_db)`
  Sets audio mic gain.
  Called by: `.ptbMic_Scroll()` (same file), `.ptbFMMic_Scroll()` (same file)
- **`.ptbCWSpeed_Scroll()`** — L29630 — `private void ptbCWSpeed_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWSpeed` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file)
- **`.chkVOX_CheckedChanged()`** — L29655 — `private void chkVOX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVOX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picSquelch_Paint()`** — L29675 — `private void picSquelch_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picSquelch` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNoiseGate_CheckedChanged()`** — L29686 — `private void chkNoiseGate_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkNoiseGate` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbVACRXGain_Scroll()`** — L29696 — `private void ptbVACRXGain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbVACRXGain` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkVAC2_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.chkVFOATX_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.ptbVACTXGain_Scroll()`** — L29721 — `private void ptbVACTXGain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbVACTXGain` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkVAC2_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.chkVFOATX_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.ptbVOX_Scroll()`** — L29750 — `private void ptbVOX_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbVOX` is scrolled.
  Called by: `.InitConsole()` (same file), `.Console_KeyDown()` (same file)
- **`.picVOX_Paint()`** — L29761 — `unsafe private void picVOX_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picVOX` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbNoiseGate_Scroll()`** — L29774 — `private void ptbNoiseGate_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbNoiseGate` is scrolled.
  Called by: `.InitConsole()` (same file), `.pnlDisplay_DoubleClick()` (same file)
- **`.picNoiseGate_Paint()`** — L29784 — `private void picNoiseGate_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picNoiseGate` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.WheelTune_MouseDown()`** — L29795 — `private void WheelTune_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `WheelTune` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMON_CheckedChanged()`** — L29801 — `private void chkMON_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMON` checked state changes.
  Called by: `.chkVFOATX_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.AudioMOXChanged()`** — L29829 — `private void AudioMOXChanged(bool tx)`
  Called by: `.PollCW()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.HdwMOXChanged()`** — L29845 — `private void HdwMOXChanged(bool tx, double freq)`
  Called by: `.chkMOX_CheckedChanged2()` (same file)
- **`.UIMOXChangedTrue()`** — L29937 — `private void UIMOXChangedTrue()`
  Called by: `.chkMOX_CheckedChanged2()` (same file)
- **`.UIMOXChangedFalse()`** — L29972 — `private void UIMOXChangedFalse()`
  Called by: `.chkMOX_CheckedChanged2()` (same file)
- **`.updateAttNudsCombos()`** — L30024 — `private void updateAttNudsCombos()`
  Called by: `.UIMOXChangedTrue()` (same file), `.UIMOXChangedFalse()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.ExpandDisplay()` (same file)
- **`.chkMOX_CheckedChanged2()`** — L30154 — `private void chkMOX_CheckedChanged2(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkMOX_Click()`** — L30530 — `private void chkMOX_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkMOX` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboMeterRXMode_SelectedIndexChanged()`** — L30551 — `private void comboMeterRXMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboMeterRXMode` selection changes.
  Called by: `.UIMOXChangedFalse()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.comboMeterTXMode_SelectedIndexChanged()`** — L30601 — `private void comboMeterTXMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboMeterTXMode` selection changes.
  Called by: `.UIMOXChangedTrue()` (same file), `.UIMOXChangedFalse()` (same file), `.chkTUN_CheckedChanged()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.isMeterModeAvailableWhenTune()`** — L30679 — `private bool isMeterModeAvailableWhenTune(MeterTXMode meterMode)`
  Called by: `.comboMeterTXMode_SelectedIndexChanged()` (same file)
- **`.chkDisplayAVG_CheckedChanged()`** — L30714 — `private void chkDisplayAVG_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDisplayAVG` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.chkDisplayPeak_CheckedChanged()`** — L30738 — `private void chkDisplayPeak_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDisplayPeak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateVFOFreqs()`** — L30762 — `private void updateVFOFreqs(bool tx, bool isTune = false)`
  Called by: `.HdwMOXChanged()` (same file), `.chkTUN_CheckedChanged()` (same file), `.chkXIT_CheckedChanged()` (same file), `.udXIT_ValueChanged()` (same file), `.chkRX2SR_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.chkTUN_CheckedChanged()`** — L30829 — `private async void chkTUN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTUN` checked state changes.
  Called by: `.AutoTuningHL2()` (same file), `.chk2TONE_CheckedChanged()` (same file)
- **`.SetupTunePulse()`** — L31033 — `public void SetupTunePulse()`
  Setups tune pulse.
  Called by: `.chkTUN_CheckedChanged()` (same file)
- **`.ATUTune()`** — L31069 — `private async void ATUTune(CancellationToken t)`
  Called by: `.chkTUN_CheckedChanged()` (same file)
- **`.comboTuneMode_SelectedIndexChanged()`** — L31093 — `private void comboTuneMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboTuneMode` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.HideFocus()`** — L31099 — `private void HideFocus(object sender, EventArgs e)`
  Hides focus.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.textbox_GotFocus()`** — L31104 — `private void textbox_GotFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `textbox` gains focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.textbox_LostFocus()`** — L31109 — `private void textbox_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `textbox` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.combo_OpenDropDown()`** — L31114 — `private void combo_OpenDropDown(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.combo_CloseDropDown()`** — L31119 — `private void combo_CloseDropDown(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkVFOLock_CheckedChanged()`** — L31124 — `private void chkVFOLock_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOLock` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOBLock_CheckedChanged()`** — L31130 — `private void chkVFOBLock_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVFOBLock` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.repopulateForms()`** — L31135 — `private void repopulateForms()`
  Called by: `.SetCATBand()` (same file), `.btnBandVHF_Click()` (same file), `.btnBandHF_Click()` (same file), `.btnBandGEN_Click()` (same file), `.radBand_CheckedChanged()` (same file)
- **`.BandPanelVisible()`** — L31141 — `public void BandPanelVisible(bool all_hidden = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModePanelVisible()`** — L31159 — `public void ModePanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOAVisible()`** — L31171 — `public void VFOAVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOBVisible()`** — L31175 — `public void VFOBVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOSyncVisible()`** — L31179 — `public void VFOSyncVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterPanelVisible()`** — L31183 — `public void FilterPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PowerRxPanelVisible()`** — L31193 — `public void PowerRxPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MonTunePanelVisible()`** — L31197 — `public void MonTunePanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SplitRitVacPanelVisible()`** — L31201 — `public void SplitRitVacPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NoiseMnfPanelVisible()`** — L31210 — `public void NoiseMnfPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MicCompVoxPanelVisible()`** — L31226 — `public void MicCompVoxPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisplayControlsPanelVisible()`** — L31234 — `public void DisplayControlsPanelVisible(bool visible)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExtendPanelDisplaySizeRight()`** — L31239 — `public void ExtendPanelDisplaySizeRight(bool expand)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExtendPanelDisplaySizeTop()`** — L31258 — `public void ExtendPanelDisplaySizeTop(bool expand)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setBandPanelVisible()`** — L31283 — `private void setBandPanelVisible(bool gen, bool hf, bool vhf, bool force = false)`
  Sets band panel visible.
  Called by: `.SetRX1Band()` (same file), `.BandPanelVisible()` (same file), `.btnBandVHF_Click()` (same file), `.btnBandHF_Click()` (same file), `.btnBandGEN_Click()` (same file), `.OnBandChangeHandler()` (same file)
- **`.btnBandVHF_Click()`** — L31340 — `private void btnBandVHF_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnBandVHF` is clicked.
  Called by: `.ExpandDisplay()` (same file)
- **`.btnBandHF_Click()`** — L31353 — `private void btnBandHF_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnBandHF` is clicked.
  Called by: `.ExpandDisplay()` (same file)
- **`.btnBandGEN_Click()`** — L31366 — `private void btnBandGEN_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnBandGEN` is clicked.
  Called by: `.ExpandDisplay()` (same file)
- **`.udFilterLow_LostFocus()`** — L31410 — `private void udFilterLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udFilterLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udFilterHigh_LostFocus()`** — L31416 — `private void udFilterHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udFilterHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterLow_LostFocus()`** — L31422 — `private void udTXFilterLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXFilterLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterHigh_LostFocus()`** — L31428 — `private void udTXFilterHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXFilterHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2FilterLow_LostFocus()`** — L31434 — `private void udRX2FilterLow_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2FilterLow` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2FilterHigh_LostFocus()`** — L31440 — `private void udRX2FilterHigh_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2FilterHigh` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRIT_LostFocus()`** — L31448 — `private void udRIT_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udRIT` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udXIT_LostFocus()`** — L31453 — `private void udXIT_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udXIT` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnChangeTuneStepSmaller_Click()`** — L31458 — `private void btnChangeTuneStepSmaller_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnChangeTuneStepSmaller` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnChangeTuneStepLarger_Click()`** — L31463 — `private void btnChangeTuneStepLarger_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnChangeTuneStepLarger` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboTXProfile_SelectedIndexChanged()`** — L31468 — `private void comboTXProfile_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.comboDigTXProfile_SelectedIndexChanged()`** — L31481 — `private void comboDigTXProfile_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboDigTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.comboFMTXProfile_SelectedIndexChanged()`** — L31491 — `private void comboFMTXProfile_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFMTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.LoadedTXProfile()`** — L31503 — `public void LoadedTXProfile()`
  MW0LGE_21j used by setup form whenever a TX profile is loaded When a digimode is selected, a number of settings are disabled. These are restored if leaving a digimode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.comboAMTXProfile_SelectedIndexChanged()`** — L31534 — `private void comboAMTXProfile_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboAMTXProfile` selection changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.chkShowTXFilter_CheckedChanged()`** — L31544 — `private void chkShowTXFilter_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowTXFilter` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVACStereo_CheckedChanged()`** — L31551 — `private void chkVACStereo_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVACStereo` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWIambic_CheckedChanged()`** — L31568 — `private void chkCWIambic_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWIambic` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWSidetone_CheckedChanged()`** — L31573 — `private void chkCWSidetone_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWSidetone` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCWPitch_ValueChanged()`** — L31578 — `private void udCWPitch_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWPitch` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboVACSampleRate_SelectedIndexChanged()`** — L31584 — `private void comboVACSampleRate_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboVACSampleRate` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkX2TR_CheckedChanged()`** — L31597 — `private void chkX2TR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkX2TR` checked state changes.
  Called by: `.GetState()` (same file)
- **`.chkShowTXCWFreq_CheckedChanged()`** — L31618 — `private void chkShowTXCWFreq_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowTXCWFreq` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowCWZero_CheckedChanged()`** — L31623 — `private void chkShowCWZero_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkShowCWZero` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCWBreakInDelay_ValueChanged()`** — L31630 — `private void udCWBreakInDelay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWBreakInDelay` value changes.
  Called by: `.udCWBreakInDelay_LostFocus()` (same file)
- **`.udCWBreakInDelay_LostFocus()`** — L31637 — `private void udCWBreakInDelay_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udCWBreakInDelay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWAPFEnabled_CheckedChanged()`** — L31642 — `private void chkCWAPFEnabled_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCWAPFEnabled` checked state changes.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.ptbCWAPFFreq_Scroll()`** — L31683 — `private void ptbCWAPFFreq_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWAPFFreq` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.ptbCWAPFBandwidth_Scroll()`** — L31701 — `private void ptbCWAPFBandwidth_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWAPFBandwidth` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.ptbCWAPFGain_Scroll()`** — L31720 — `private void ptbCWAPFGain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCWAPFGain` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.DisableDAX()`** — L31741 — `public void DisableDAX()`
  Disables dax.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableDAX()`** — L31750 — `public void EnableDAX()`
  Enables dax.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkVAC1_CheckedChanged()`** — L31757 — `private void chkVAC1_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVAC1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2_CheckedChanged()`** — L31805 — `private void chkVAC2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVAC2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRXEQ_CheckedChanged()`** — L31852 — `private void chkRXEQ_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRXEQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXEQ_CheckedChanged()`** — L31862 — `private void chkTXEQ_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTXEQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TuneHitTest()`** — L31891 — `private TuneLocation TuneHitTest(int x, int y)`
  Called by: `.Console_MouseWheel()` (same file)
- **`.Console_MouseWheel()`** — L31933 — `private void Console_MouseWheel(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `Console` receives a mouse wheel event.
  Called by: `.Console_KeyDown()` (same file), `.OnMouseWheelChanged()` (same file)
- **`.SnapTune()`** — L32108 — `public double SnapTune(double freq_mhz, int step_size_hz, int num_steps)`
  Calculates a "Snapped" frequency that lies on an integer multiple of the Tune Step.
  Called by: `.Console_MouseWheel()` (same file)
- **`.txtVFOAFreq_LostFocus()`** — L32146 — `private void txtVFOAFreq_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` loses focus.
  Called by: `.CalibratePAGain()` (same file), `.LowPowerPASweep()` (same file), `.SetupForHPSDRModel()` (same file), `.SetWavePlayback()` (same file), `.VFOAUpdate()` (same file), `.chkPower_CheckedChanged()` (same file) — and 15 more
- **`.setupModifyXVTRantennaArray()`** — L32790 — `private void setupModifyXVTRantennaArray()`
  Called by: `.InitConsole()` (same file)
- **`.modifyXVTRantenna()`** — L32801 — `private void modifyXVTRantenna(int rx, double freq, int rx_xvtr_index)`
  Called by: `.txtVFOAFreq_LostFocus()` (same file)
- **`.undoXVTRantennaModify()`** — L32842 — `private void undoXVTRantennaModify(int rx)`
  Called by: `.Console_Closing()` (same file), `.txtVFOAFreq_LostFocus()` (same file), `.modifyXVTRantenna()` (same file)
- **`.getTXBandWhenExtended()`** — L32856 — `private Band getTXBandWhenExtended(Band b, double frequency = -1)`
  Returns txband when extended.
  Called by: `.txtVFOAFreq_LostFocus()` (same file), `.txtVFOABand_LostFocus()` (same file), `.txtVFOBFreq_LostFocus()` (same file)
- **`.txtVFOAFreq_KeyPress()`** — L32895 — `private void txtVFOAFreq_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAFreq_MouseMove()`** — L32925 — `private void txtVFOAFreq_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` receives mouse movement.
  Called by: `.panelVFOAHover_MouseMove()` (same file), `.txtVFOALSD_MouseMove()` (same file), `.txtVFOAMSD_MouseMove()` (same file)
- **`.txtVFOAFreq_MouseLeave()`** — L32969 — `private void txtVFOAFreq_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOAFreq` is left by the mouse.
  Called by: `.txtVFOAMSD_MouseLeave()` (same file)
- **`.txtVFOABand_LostFocus()`** — L32975 — `private void txtVFOABand_LostFocus(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOABand` loses focus.
  Called by: `.VFOASubUpdate()` (same file), `.UpdateVFOASub()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.txtVFOABand_KeyPress()`** — L33119 — `private void txtVFOABand_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `txtVFOABand` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBFreq_LostFocus()`** — L33156 — `private void txtVFOBFreq_LostFocus(object sender, System.EventArgs e)`
  txtVFOBFreq
  Called by: `.SetWavePlayback()` (same file), `.VFOBUpdate()` (same file), `.updateVFOFreqs()` (same file), `.chkX2TR_CheckedChanged()` (same file), `.zoomToBandBandwidth()` (same file), `.ptbDisplayZoom_Scroll()` (same file) — and 12 more
- **`.txtVFOBFreq_KeyPress()`** — L33815 — `private void txtVFOBFreq_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `txtVFOBFreq` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBFreq_MouseMove()`** — L33845 — `private void txtVFOBFreq_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBFreq` receives mouse movement.
  Called by: `.panelVFOBHover_MouseMove()` (same file), `.txtVFOBMSD_MouseMove()` (same file), `.txtVFOBLSD_MouseMove()` (same file)
- **`.txtVFOBFreq_MouseLeave()`** — L33889 — `private void txtVFOBFreq_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOBFreq` is left by the mouse.
  Called by: `.txtVFOBMSD_MouseLeave()` (same file)
- **`.panelVFOAHover_MouseMove()`** — L33895 — `private void panelVFOAHover_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `panelVFOAHover` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panelVFOBHover_MouseMove()`** — L33906 — `private void panelVFOBHover_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `panelVFOBHover` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOALSD_MouseDown()`** — L33917 — `private void txtVFOALSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOALSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOALSD_MouseMove()`** — L33925 — `private void txtVFOALSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOALSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAMSD_MouseDown()`** — L33936 — `private void txtVFOAMSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOAMSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAMSD_MouseMove()`** — L33944 — `private void txtVFOAMSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOAMSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOAMSD_MouseLeave()`** — L33950 — `private void txtVFOAMSD_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOAMSD` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBMSD_MouseDown()`** — L33955 — `private void txtVFOBMSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBMSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBMSD_MouseLeave()`** — L33963 — `private void txtVFOBMSD_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOBMSD` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBMSD_MouseMove()`** — L33968 — `private void txtVFOBMSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBMSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBLSD_MouseDown()`** — L33974 — `private void txtVFOBLSD_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBLSD` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOBLSD_MouseMove()`** — L33982 — `private void txtVFOBLSD_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOBLSD` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.overRX()`** — L34044 — `private bool overRX(int x, int y, int rx, bool bIgnorePanafallWaterfall = true)`
  Called by: `.UpdatePeakText()` (same file), `.pnlDisplay_MouseDown()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.notchMouseWheel()`** — L34138 — `private void notchMouseWheel(int wheelDelta)`
  Called by: `.Console_MouseWheel()` (same file)
- **`.CurrentDSPhasTwoSidebands()`** — L34163 — `public bool CurrentDSPhasTwoSidebands(int rx, bool tx = false)`
  Called by: `.pnlDisplay_MouseMove()` (same file)
- **`.agcCalOffset()`** — L34185 — `private float agcCalOffset(int rx)`
  Called by: `.setAGCThresholdPoint()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.getFilterEdgesInPixels()`** — L34216 — `private void getFilterEdgesInPixels(MouseEventArgs e, ref int low_x, ref int high_x, ref int vfoa_sub_x, ref int vfoa_sub_low_x, ref int vfoa_sub_high_x)`
  Returns filter edges in pixels.
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.dragWholeFilter()`** — L34275 — `private void dragWholeFilter(MouseEventArgs e)`
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.adjustForSnapClickTuning()`** — L34309 — `private double adjustForSnapClickTuning(int rx, double freq)`
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.getFrequencyAtPixel()`** — L34338 — `private double getFrequencyAtPixel(int x, int nRX)`
  Returns frequency at pixel.
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.ShowNotchPopup()`** — L34380 — `public void ShowNotchPopup(int x, int y, MNotch notch, int min_width, int max_width, bool on_top, int notch_index = -1)`
  Shows notch popup.
  Called by: `.pnlDisplay_MouseUp()` (same file)
- **`.ptbDisplayPan_Scroll()`** — L34400 — `private void ptbDisplayPan_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbDisplayPan` is scrolled.
  Called by: `.btnDisplayPanCenter_Click()` (same file)
- **`.btnDisplayPanCenter_Click()`** — L34414 — `private void btnDisplayPanCenter_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnDisplayPanCenter` is clicked.
  Called by: `.PanCentre()` (same file), `.SetPanAdjust()` (same file)
- **`.zoomToBandBandwidth()`** — L34439 — `private bool zoomToBandBandwidth(Band b, int rx)`
  Called by: `.ZoomToBand()` (same file)
- **`.ptbDisplayZoom_Scroll()`** — L34527 — `private void ptbDisplayZoom_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbDisplayZoom` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.radDisplayZoom05_CheckedChanged()`** — L34602 — `private void radDisplayZoom05_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom05` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom05()`** — L34609 — `private void displayZoom05()`
  Called by: `.radDisplayZoom05_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radDisplayZoom1x_CheckedChanged()`** — L34615 — `private void radDisplayZoom1x_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom1x` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom1()`** — L34622 — `private void displayZoom1()`
  Called by: `.radDisplayZoom1x_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radDisplayZoom2x_CheckedChanged()`** — L34628 — `private void radDisplayZoom2x_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom2x` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom2()`** — L34635 — `private void displayZoom2()`
  Called by: `.radDisplayZoom2x_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radDisplayZoom4x_CheckedChanged()`** — L34640 — `private void radDisplayZoom4x_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radDisplayZoom4x` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayZoom4()`** — L34647 — `private void displayZoom4()`
  Called by: `.radDisplayZoom4x_CheckedChanged()` (same file), `.DoOtherButtonAction()` (same file)
- **`.radBand160_Click()`** — L34659 — `private void radBand160_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand160` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand80_Click()`** — L34665 — `private void radBand80_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand80` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand60_Click()`** — L34671 — `private void radBand60_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand60` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand40_Click()`** — L34677 — `private void radBand40_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand40` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand30_Click()`** — L34683 — `private void radBand30_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand30` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand20_Click()`** — L34689 — `private void radBand20_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand20` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand17_Click()`** — L34695 — `private void radBand17_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand17` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand15_Click()`** — L34701 — `private void radBand15_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand15` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand12_Click()`** — L34707 — `private void radBand12_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand12` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand10_Click()`** — L34713 — `private void radBand10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand10` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand6_Click()`** — L34719 — `private void radBand6_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand6` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBand2_Click()`** — L34725 — `private void radBand2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand2` is clicked.
  Called by: `.SetCATBand()` (same file)
- **`.radBandWWV_Click()`** — L34731 — `private void radBandWWV_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandWWV` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBandGEN_Click()`** — L34737 — `private void radBandGEN_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandGEN` is clicked.
  Called by: `.SetCATBand()` (same file), `.mnuBand_Click()` (same file)
- **`.radBandVHF_Click()`** — L34742 — `private void radBandVHF_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `radBandVHF` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setVFOAFreqNoUpdate()`** — L34757 — `private void setVFOAFreqNoUpdate(double freq)`
  Sets vfoafreq no update.
  Called by: `.SetRX1Mode()` (same file)
- **`.setVFOBFreqNoUpdate()`** — L34762 — `private void setVFOBFreqNoUpdate(double freq)`
  Sets vfobfreq no update.
  Called by: `.SetRX2Mode()` (same file)
- **`.initControlBackColours()`** — L34770 — `private void initControlBackColours(Control c)`
  MW0LGE_21d used to default colours of all button+radio controls, and inside other panels or groups an issue was noticed where text change colour on buttons that had be selected/deseleted
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRX1Mode()`** — L34796 — `private void SetRX1Mode(DSPMode new_mode)`
  Sets rx1 mode.
  Called by: `.radModeButton_CheckedChanged()` (same file), `.chkVFOBTX_CheckedChanged()` (same file)
- **`.radModeButton_CheckedChanged()`** — L35500 — `private void radModeButton_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radModeButton` checked state changes.
  Called by: `.selectModes()` (same file)
- **`.SetRX1Filter()`** — L35592 — `public void SetRX1Filter(Filter new_filter)`
  Sets rx1 filter.
  Called by: `.radFilter_CheckedChanged()` (same file)
- **`.filterAndDspModeValid()`** — L35719 — `private bool filterAndDspModeValid(int rx)`
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRX2Filters()` (same file), `.UpdateRX1FilterNames()` (same file), `.UpdateRX2FilterNames()` (same file), `.SetRX1Filter()` (same file), `.SetRX2Filter()` (same file) — and 2 more
- **`.radRX2Filter_CheckedChanged()`** — L35731 — `private void radRX2Filter_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX2Filter` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MatchTXFilterToRXFilter()`** — L35792 — `public void MatchTXFilterToRXFilter()`
  Called by: `.radRX2Filter_CheckedChanged()` (same file), `.radFilter_CheckedChanged()` (same file), `.radFilter_rx1_MouseUp()` (same file), `.radFilter_rx2_MouseUp()` (same file)
- **`.radFilter_CheckedChanged()`** — L35814 — `private void radFilter_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFilter` checked state changes.
  Called by: `.selectFilters()` (same file)
- **`.udFilterLow_ValueChanged()`** — L35893 — `private void udFilterLow_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udFilterLow` value changes.
  Called by: `.udFilterLow_LostFocus()` (same file)
- **`.udFilterHigh_ValueChanged()`** — L35914 — `private void udFilterHigh_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udFilterHigh` value changes.
  Called by: `.udFilterHigh_LostFocus()` (same file)
- **`.ConstrainFilter()`** — L35935 — `public bool ConstrainFilter(ref int nNewLow, ref int nNewHigh, int rx, bool filterShift = false)`
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRX2Filters()` (same file), `.ptbFilterShift_Scroll()` (same file), `.ptbFilterWidth_Scroll()` (same file), `.pnlDisplay_MouseMove()` (same file)
- **`.ptbFilterShift_Scroll()`** — L36028 — `private void ptbFilterShift_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbFilterShift` is scrolled.
  Called by: `.btnFilterShiftReset_Click()` (same file)
- **`.ptbFilterShift_Update()`** — L36106 — `private void ptbFilterShift_Update(int low, int high)`
  Called by: `.UpdateRX1Filters()` (same file)
- **`.btnFilterShiftReset_Click()`** — L36163 — `private void btnFilterShiftReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnFilterShiftReset` is clicked.
  Called by: `.Console_KeyDown()` (same file), `.btnIFtoVFO_Click()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.ptbFilterWidth_Update()`** — L36183 — `private void ptbFilterWidth_Update(int low, int high)`
  Called by: `.UpdateRX1Filters()` (same file)
- **`.ptbFilterWidth_Scroll()`** — L36221 — `private void ptbFilterWidth_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbFilterWidth` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbFilterWidthScroll_newMode()`** — L36320 — `private void tbFilterWidthScroll_newMode()`
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.CATVFOSwap()`** — L36343 — `public void CATVFOSwap(string pChangec)`
  Added 6/20/05 BT for CAT commands
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CopyVFOAtoB()`** — L36366 — `public void CopyVFOAtoB()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOAtoB_Click()`** — L36371 — `private void btnVFOAtoB_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnVFOAtoB` is clicked.
  Called by: `.CATVFOAtoB()` (same file), `.Console_KeyDown()` (same file), `.CATVFOSwap()` (same file), `.CopyVFOAtoB()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.CopyVFOBtoA()`** — L36407 — `public void CopyVFOBtoA()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOBtoA_Click()`** — L36412 — `private void btnVFOBtoA_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnVFOBtoA` is clicked.
  Called by: `.CATVFOBtoA()` (same file), `.Console_KeyDown()` (same file), `.CATVFOSwap()` (same file), `.CopyVFOBtoA()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.VFOSwap()`** — L36446 — `public void VFOSwap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnVFOSwap_Click()`** — L36451 — `private void btnVFOSwap_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnVFOSwap` is clicked.
  Called by: `.CATVFOABSwap()` (same file), `.Console_KeyDown()` (same file), `.CATVFOSwap()` (same file), `.VFOSwap()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.UpdateVFOASub()`** — L36515 — `private void UpdateVFOASub()`
  Updates vfoasub.
  Called by: `.chkPower_CheckedChanged()` (same file), `.txtVFOBFreq_LostFocus()` (same file), `.chkVFOSplit_CheckedChanged()` (same file), `.chkEnableMultiRX_CheckedChanged()` (same file)
- **`.chkVFOSplit_CheckedChanged()`** — L36595 — `private void chkVFOSplit_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOSplit` checked state changes.
  Called by: `.chkPower_CheckedChanged()` (same file), `.chkEnableMultiRX_CheckedChanged()` (same file)
- **`.SetQuickSplit()`** — L36718 — `public void SetQuickSplit()`
  Sets quick split.
  Called by: `.chkVFOSplit_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file)
- **`.chkXIT_CheckedChanged()`** — L36828 — `private void chkXIT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkXIT` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRIT_CheckedChanged()`** — L36856 — `private void chkRIT_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRIT` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRIT_ValueChanged()`** — L36891 — `private void udRIT_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRIT` value changes.
  Called by: `.udRIT_LostFocus()` (same file)
- **`.udXIT_ValueChanged()`** — L36918 — `private void udXIT_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udXIT` value changes.
  Called by: `.udXIT_LostFocus()` (same file)
- **`.btnXITReset_Click()`** — L36945 — `private void btnXITReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnXITReset` is clicked.
  Called by: `.Console_KeyDown()` (same file), `.DoGeneralSettingAction()` (same file)
- **`.btnRITReset_Click()`** — L36950 — `private void btnRITReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnRITReset` is clicked.
  Called by: `.Console_KeyDown()` (same file), `.DoGeneralSettingAction()` (same file)
- **`.setRIT_LEDs()`** — L36955 — `private void setRIT_LEDs()`
  Sets rit leds.
  Called by: `.udRIT_ValueChanged()` (same file), `.udXIT_ValueChanged()` (same file)
- **`.setXIT_LEDs()`** — L36974 — `private void setXIT_LEDs()`
  Sets xit leds.
  Called by: `.udRIT_ValueChanged()` (same file), `.udXIT_ValueChanged()` (same file)
- **`.btnZeroBeat_Click()`** — L36993 — `private void btnZeroBeat_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnZeroBeat` is clicked.
  Called by: `.CalibrateLevel()` (same file), `.Console_KeyDown()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.FindPeakFreqInPassband()`** — L37083 — `unsafe private int FindPeakFreqInPassband()`
  Finds peak freq in passband.
  Called by: `.btnZeroBeat_Click()` (same file)
- **`.btnIFtoVFO_Click()`** — L37160 — `private void btnIFtoVFO_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnIFtoVFO` is clicked.
  Called by: `.CATVFOSwap()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.chkANF_CheckedChanged()`** — L37259 — `private void chkANF_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkANF` checked state changes.
  Called by: `.SetRX1Mode()` (same file)
- **`.chkDSPNB2_CheckedChanged()`** — L37286 — `private void chkDSPNB2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDSPNB2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NB2_CheckedChanged()`** — L37313 — `private void chkRX2NB2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2NB2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCPDR_CheckedChanged()`** — L37349 — `private void chkCPDR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkCPDR` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbCPDR_Scroll()`** — L37385 — `private void ptbCPDR_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbCPDR` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkCPDR_CheckedChanged()` (same file)
- **`.chkDX_CheckedChanged()`** — L37398 — `private void chkDX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkDX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMemoryQuickSave_Click()`** — L37410 — `private void btnMemoryQuickSave_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnMemoryQuickSave` is clicked.
  Called by: `.CATMemoryQS()` (same file), `.Console_KeyDown()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.btnMemoryQuickRestore_Click()`** — L37417 — `private void btnMemoryQuickRestore_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnMemoryQuickRestore` is clicked.
  Called by: `.CATMemoryQR()` (same file), `.Console_KeyDown()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.ptbPanMainRX_Scroll()`** — L37428 — `private void ptbPanMainRX_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbPanMainRX` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkPanSwap_CheckedChanged()` (same file), `.ptbPanMainRX_DoubleClick()` (same file)
- **`.ptbPanSubRX_Scroll()`** — L37448 — `private void ptbPanSubRX_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbPanSubRX` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkPanSwap_CheckedChanged()` (same file), `.ptbPanSubRX_DoubleClick()` (same file)
- **`.chkEnableMultiRX_CheckedChanged()`** — L37472 — `unsafe private void chkEnableMultiRX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnableMultiRX` checked state changes.
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.chkPanSwap_CheckedChanged()`** — L37571 — `private void chkPanSwap_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPanSwap` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX0Gain_Scroll()`** — L37586 — `private void ptbRX0Gain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX0Gain` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.ptbRX1Gain_Scroll()`** — L37627 — `private void ptbRX1Gain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX1Gain` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.chkFullDuplex_CheckedChanged()`** — L37670 — `private void chkFullDuplex_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkFullDuplex` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getConsole()`** — L37696 — `public static Console getConsole()`
  Returns console.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WndProc()`** — L37701 — `protected override void WndProc(ref Message m)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkFWCATUBypass_Click()`** — L37714 — `private void chkFWCATUBypass_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkFWCATUBypass` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSplitDisplay_CheckedChanged()`** — L37718 — `private void chkSplitDisplay_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkSplitDisplay` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ckQuickPlay_CheckedChanged()`** — L37729 — `private async void ckQuickPlay_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ckQuickPlay` checked state changes.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.ckQuickRec_CheckedChanged()`** — L37794 — `private void ckQuickRec_CheckedChanged(object sender, System.EventArgs e)`
  private bool _updated_from_wave_form = false; public bool UpdatedFromWaveForm { // prevent wave form changes causing loop get { return _updated_from_wave_form; } set { _updated_from_wave_form = value; } }
  Called by: `.DoOtherButtonAction()` (same file)
- **`.moveModeSpecificPanels()`** — L37849 — `private void moveModeSpecificPanels()`
  Called by: `.ResizeConsole()` (same file), `.SelectModeDependentPanel()` (`Console/Andromeda/Andromeda.cs`)
- **`.ResizeConsole()`** — L37856 — `private void ResizeConsole(int h_delta, int v_delta)`
  Called by: `.Console_Resize()` (same file)
- **`.GrabConsoleSizeBasis()`** — L37991 — `public void GrabConsoleSizeBasis()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2_CheckedChanged()`** — L38349 — `private void chkRX2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2` checked state changes.
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.setSmallRX2ModeFilterLabels()`** — L38455 — `private void setSmallRX2ModeFilterLabels()`
  Sets small rx2 mode filter labels.
  Called by: `.radModeButton_CheckedChanged()` (same file), `.radFilter_CheckedChanged()` (same file), `.chkEnableMultiRX_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.radRX2ModeButton_CheckedChanged()` (same file)
- **`.chkRX2SR_CheckedChanged()`** — L38483 — `private void chkRX2SR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2SR` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.panelVFOASubHover_Paint()`** — L38505 — `private void panelVFOASubHover_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `panelVFOASubHover` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panelVFOASubHover_MouseMove()`** — L38521 — `private void panelVFOASubHover_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `panelVFOASubHover` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtVFOABand_MouseMove()`** — L38533 — `private void txtVFOABand_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `txtVFOABand` receives mouse movement.
  Called by: `.panelVFOASubHover_MouseMove()` (same file)
- **`.txtVFOABand_MouseLeave()`** — L38564 — `private void txtVFOABand_MouseLeave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtVFOABand` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetRX2Mode()`** — L38570 — `private void SetRX2Mode(DSPMode new_mode)`
  Sets rx2 mode.
  Called by: `.radRX2ModeButton_CheckedChanged()` (same file), `.radRX2ModeLSB_CheckedChanged()` (same file), `.radRX2ModeUSB_CheckedChanged()` (same file), `.radRX2ModeDSB_CheckedChanged()` (same file), `.radRX2ModeCWL_CheckedChanged()` (same file), `.radRX2ModeCWU_CheckedChanged()` (same file) — and 7 more
- **`.radRX2ModeButton_CheckedChanged()`** — L39107 — `private void radRX2ModeButton_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeButton` checked state changes.
  Called by: `.selectModes()` (same file)
- **`.radRX2ModeLSB_CheckedChanged()`** — L39176 — `private void radRX2ModeLSB_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeLSB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeUSB_CheckedChanged()`** — L39184 — `private void radRX2ModeUSB_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeUSB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDSB_CheckedChanged()`** — L39192 — `private void radRX2ModeDSB_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDSB` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeCWL_CheckedChanged()`** — L39200 — `private void radRX2ModeCWL_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeCWL` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeCWU_CheckedChanged()`** — L39208 — `private void radRX2ModeCWU_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeCWU` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeFMN_CheckedChanged()`** — L39216 — `private void radRX2ModeFMN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeFMN` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeAM_CheckedChanged()`** — L39224 — `private void radRX2ModeAM_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeAM` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeSAM_CheckedChanged()`** — L39232 — `private void radRX2ModeSAM_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeSAM` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDIGL_CheckedChanged()`** — L39240 — `private void radRX2ModeDIGL_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDIGL` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDIGU_CheckedChanged()`** — L39248 — `private void radRX2ModeDIGU_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDIGU` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2ModeDRM_CheckedChanged()`** — L39256 — `private void radRX2ModeDRM_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2ModeDRM` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetRX2Filter()`** — L39264 — `public void SetRX2Filter(Filter new_filter, bool update = true)`
  Sets rx2 filter.
  Called by: `.radRX2Filter_CheckedChanged()` (same file), `.radRX2Filter1_CheckedChanged()` (same file), `.radRX2Filter2_CheckedChanged()` (same file), `.radRX2Filter3_CheckedChanged()` (same file), `.radRX2Filter4_CheckedChanged()` (same file), `.radRX2Filter5_CheckedChanged()` (same file) — and 4 more
- **`.radRX2Filter1_CheckedChanged()`** — L39362 — `private void radRX2Filter1_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter2_CheckedChanged()`** — L39368 — `private void radRX2Filter2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter3_CheckedChanged()`** — L39374 — `private void radRX2Filter3_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter4_CheckedChanged()`** — L39380 — `private void radRX2Filter4_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter4` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter5_CheckedChanged()`** — L39386 — `private void radRX2Filter5_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter5` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter6_CheckedChanged()`** — L39392 — `private void radRX2Filter6_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter6` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2Filter7_CheckedChanged()`** — L39398 — `private void radRX2Filter7_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2Filter7` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2FilterVar1_CheckedChanged()`** — L39404 — `private void radRX2FilterVar1_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2FilterVar1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX2FilterVar2_CheckedChanged()`** — L39410 — `private void radRX2FilterVar2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `radRX2FilterVar2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2FilterLow_ValueChanged()`** — L39416 — `private void udRX2FilterLow_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2FilterLow` value changes.
  Called by: `.udRX2FilterLow_LostFocus()` (same file)
- **`.udRX2FilterHigh_ValueChanged()`** — L39443 — `private void udRX2FilterHigh_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRX2FilterHigh` value changes.
  Called by: `.udRX2FilterHigh_LostFocus()` (same file)
- **`.chkRX2ANF_CheckedChanged()`** — L39464 — `private void chkRX2ANF_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2ANF` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2BIN_CheckedChanged()`** — L39491 — `private void chkRX2BIN_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2BIN` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboRX2MeterMode_SelectedIndexChanged()`** — L39505 — `private void comboRX2MeterMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2MeterMode` selection changes.
  Called by: `.UIMOXChangedFalse()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.chkRX2Preamp_CheckedChanged()`** — L39552 — `private void chkRX2Preamp_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2Preamp` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2RF_Scroll()`** — L39565 — `private void ptbRX2RF_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2RF` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.picRX2Squelch_Paint()`** — L39598 — `private void picRX2Squelch_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picRX2Squelch` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX1Preamp_CheckedChanged()`** — L39608 — `private void chkRX1Preamp_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX1Preamp` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2Pan_Scroll()`** — L39624 — `private void ptbRX2Pan_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2Pan` is scrolled.
  Called by: `.InitConsole()` (same file), `.ptbRX2Pan_DoubleClick()` (same file)
- **`.ptbRX2Gain_Scroll()`** — L39643 — `private void ptbRX2Gain_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2Gain` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkRX2Mute_CheckedChanged()` (same file)
- **`.chkRX2Mute_CheckedChanged()`** — L39682 — `private void chkRX2Mute_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2Mute` checked state changes.
  Called by: `.InitConsole()` (same file), `.SetRX2Mode()` (same file)
- **`.comboRX2DisplayMode_SelectedIndexChanged()`** — L39722 — `private void comboRX2DisplayMode_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2DisplayMode` selection changes.
  Called by: `.SetupDisplayEngine()` (same file)
- **`.chkRX2DisplayAVG_CheckedChanged()`** — L39770 — `private void chkRX2DisplayAVG_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2DisplayAVG` checked state changes.
  Called by: `.chkRX2_CheckedChanged()` (same file)
- **`.chkRX2DisplayPeak_CheckedChanged()`** — L39792 — `private void chkRX2DisplayPeak_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRX2DisplayPeak` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateDSP()`** — L39823 — `private void UpdateDSP()`
  Updates dsp.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.BuildFilterCharacteristics()`** — L40049 — `public void BuildFilterCharacteristics()`
  Builds filter characteristics.
  Called by: `.UpdateDSP()` (same file)
- **`.calcFilterCharacteristics()`** — L40081 — `private (double[], int, int) calcFilterCharacteristics(int id, double rate, int filter_size, int w_type, double corner_freq, bool hi_res)`
  Called by: `.BuildFilterCharacteristics()` (same file)
- **`.SetupRX2Band()`** — L40334 — `public void SetupRX2Band(Band b, bool is_for_rx1_vfo_b = false)`
  Setups rx2 band.
  Called by: `.comboRX2Band_SelectedIndexChanged()` (same file), `.HandleFrontPanelButtonPress()` (`Console/Andromeda/Andromeda.cs`), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.comboRX2Band_SelectedIndexChanged()`** — L40371 — `private void comboRX2Band_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2Band` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateResolutionStatusBarText()`** — L40406 — `private void updateResolutionStatusBarText()`
  Called by: `.Console_Resize()` (same file), `.Console_Shown()` (same file)
- **`.Console_Resize()`** — L40424 — `private void Console_Resize(object sender, System.EventArgs e)`
  WinForms event handler: runs when `Console` is resized.
  Called by: `.chkRX2_CheckedChanged()` (same file)
- **`.comboRX2AGC_SelectedIndexChanged()`** — L40496 — `private void comboRX2AGC_SelectedIndexChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `comboRX2AGC` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOSync_CheckedChanged()`** — L40585 — `private void chkVFOSync_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOSync` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOATX_CheckedChanged()`** — L40679 — `private void chkVFOATX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOATX` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.BroadcastVFOChange()`** — L40743 — `private void BroadcastVFOChange(string ndx)`
  Called by: `.OnVFOTXChanged()` (same file)
- **`.chkVFOBTX_CheckedChanged()`** — L40756 — `private void chkVFOBTX_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkVFOBTX` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.toolStripMenuItemRX1FilterConfigure_Click()`** — L40858 — `private void toolStripMenuItemRX1FilterConfigure_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX1FilterConfigure` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItemRX1FilterReset_Click()`** — L40871 — `private void toolStripMenuItemRX1FilterReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX1FilterReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItemRX2FilterConfigure_Click()`** — L40918 — `private void toolStripMenuItemRX2FilterConfigure_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX2FilterConfigure` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchDelete_Click()`** — L40931 — `private void toolStripNotchDelete_Click(Object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchDelete` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchRemember_Click()`** — L40935 — `private void toolStripNotchRemember_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchRemember` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchNormal_Click()`** — L40939 — `private void toolStripNotchNormal_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchNormal` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchDeep_Click()`** — L40943 — `private void toolStripNotchDeep_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchDeep` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripNotchVeryDeep_Click()`** — L40947 — `private void toolStripNotchVeryDeep_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripNotchVeryDeep` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItemRX2FilterReset_Click()`** — L40951 — `private void toolStripMenuItemRX2FilterReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripMenuItemRX2FilterReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTNF_CheckedChanged()`** — L40998 — `private void chkTNF_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkTNF` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.ChangeNotchBW()`** — L41022 — `unsafe public bool ChangeNotchBW(MNotch notch, double newWidth, int notch_index = -1)`
  Called by: `.onBWChanged()` (same file), `.notchMouseWheel()` (same file), `.pnlDisplay_MouseMove()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.ChangeNotchCentreFrequency()`** — L41065 — `unsafe public bool ChangeNotchCentreFrequency(MNotch notch, double newCentreFrequencyHz, int sourceRX, int notch_index = -1)`
  Called by: `.pnlDisplay_MouseMove()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.changeNotchActive()`** — L41138 — `unsafe private bool changeNotchActive(MNotch notch, bool bActive)`
  Called by: `.onActiveChanged()` (same file)
- **`.toggleNotchActive()`** — L41174 — `unsafe private bool toggleNotchActive(MNotch notch)`
  Called by: `.pnlDisplay_MouseDown()` (same file)
- **`.removeNotch()`** — L41213 — `private bool removeNotch(MNotch notch)`
  Called by: `.onNotchDelete()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.AddNotch()`** — L41237 — `public void AddNotch(double fFreqHZ, int sourceRX)`
  Adds notch.
  Called by: `.TNFAdd()` (same file), `.pnlDisplay_MouseDown()` (same file)
- **`.notchSidebandShift()`** — L41296 — `private int notchSidebandShift(int rx)`
  Called by: `.TNFAdd()` (same file)
- **`.btnTNFAdd_Click()`** — L41324 — `private void btnTNFAdd_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnTNFAdd` is clicked.
  Called by: `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.TNFAdd()`** — L41328 — `public void TNFAdd(int rx)`
  Called by: `.btnTNFAdd_Click()` (same file), `.DoOtherButtonAction()` (same file)
- **`.ptbFMMic_Scroll()`** — L41348 — `private void ptbFMMic_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbFMMic` is scrolled.
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.chkFMCTCSS_CheckedChanged()`** — L41378 — `private void chkFMCTCSS_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMCTCSS` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboFMCTCSS_SelectedIndexChanged()`** — L41383 — `private void comboFMCTCSS_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFMCTCSS` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.InitCTCSS()`** — L41389 — `private void InitCTCSS()`
  Inits ctcss.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitMemoryFrontPanel()`** — L41396 — `private void InitMemoryFrontPanel()`
  Inits memory front panel.
  Called by: `.InitConsole()` (same file)
- **`.radFMDeviation2kHz_CheckedChanged()`** — L41403 — `private void radFMDeviation2kHz_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFMDeviation2kHz` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.fmDeviation2k()`** — L41407 — `private void fmDeviation2k(bool force)`
  Called by: `.InitConsole()` (same file), `.radFMDeviation2kHz_CheckedChanged()` (same file)
- **`.radFMDeviation5kHz_CheckedChanged()`** — L41441 — `private void radFMDeviation5kHz_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radFMDeviation5kHz` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.fmDeviation5k()`** — L41445 — `private void fmDeviation5k(bool force)`
  Called by: `.InitConsole()` (same file), `.radFMDeviation5kHz_CheckedChanged()` (same file)
- **`.udFMOffset_ValueChanged()`** — L41480 — `private void udFMOffset_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udFMOffset` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXHigh_CheckedChanged()`** — L41485 — `private void chkFMTXHigh_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXHigh` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXSimplex_CheckedChanged()`** — L41499 — `private void chkFMTXSimplex_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXSimplex` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXLow_CheckedChanged()`** — L41513 — `private void chkFMTXLow_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXLow` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMTXRev_CheckedChanged()`** — L41527 — `private void chkFMTXRev_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMTXRev` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFMMode_Click()`** — L41555 — `private void chkFMMode_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFMMode` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuMemory_Click()`** — L41604 — `private void mnuMemory_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuMemory` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.RecallMemory()`** — L41612 — `public void RecallMemory(MemoryRecord record)`
  Called by: `.comboFMMemory_SelectedIndexChanged()` (same file), `.changeComboFMMemory()` (same file)
- **`.comboFMMemory_SelectedIndexChanged()`** — L41642 — `private void comboFMMemory_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboFMMemory` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnFMMemoryUp_Click()`** — L41650 — `private void btnFMMemoryUp_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFMMemoryUp` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnFMMemoryDown_Click()`** — L41656 — `private void btnFMMemoryDown_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFMMemoryDown` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.changeComboFMMemory()`** — L41662 — `public void changeComboFMMemory(int index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnFMMemory_Click()`** — L41678 — `private void btnFMMemory_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFMMemory` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.memoryToolStripMenuItem_Click()`** — L41697 — `private void memoryToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `memoryToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.waveToolStripMenuItem_Click()`** — L41718 — `private void waveToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `waveToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.CollapseToolStripMenuItem_Click()`** — L41737 — `private void CollapseToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `CollapseToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.equalizerToolStripMenuItem_Click()`** — L41749 — `private void equalizerToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `equalizerToolStripMenuItem` is clicked.
  Called by: `.chkRXEQ_MouseDown()` (same file), `.chkTXEQ_MouseDown()` (same file), `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.xVTRsToolStripMenuItem_Click()`** — L41770 — `private void xVTRsToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `xVTRsToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.cWXToolStripMenuItem_Click()`** — L41789 — `private void cWXToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `cWXToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.SetComboPreampForHPSDR()`** — L41842 — `public void SetComboPreampForHPSDR()`
  Sets combo preamp for hpsdr.
  Called by: `.SetupForHPSDRModel()` (same file)
- **`.MakeLineInList()`** — L41917 — `private void MakeLineInList()`
  Called by: `.SetMicGain()` (same file)
- **`.SetMicXlr()`** — L41929 — `public void SetMicXlr()`
  Sets mic xlr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMicGain()`** — L41935 — `public void SetMicGain()`
  Sets mic gain.
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.BPF1ToolStripMenuItem_Click()`** — L41950 — `private void BPF1ToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BPF1ToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BPF2ToolStripMenuItem_Click()`** — L41955 — `private void BPF2ToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BPF2ToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ExpandDisplay()`** — L42074 — `private void ExpandDisplay(bool bSuspendDraw = true)`
  Called by: `.GetState()` (same file), `.CollapseToolStripMenuItem_Click()` (same file)
- **`.setPAProfileLabelPos()`** — L42492 — `private void setPAProfileLabelPos()`
  Sets paprofile label pos.
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.CollapseDisplay()`** — L42540 — `public void CollapseDisplay(bool bSuspendDraw = true)`
  modified G8NJJ to add alternate top/button controls for Andromeda optimised for 1024x600 touchscreen display
  Called by: `.GetState()` (same file), `.btnBandVHF_Click()` (same file), `.btnBandHF_Click()` (same file), `.btnBandGEN_Click()` (same file), `.CollapseToolStripMenuItem_Click()` (same file), `.radRX1Show_CheckedChanged()` (same file) — and 1 more
- **`.RepositionControlsForCollapsedlDisplay()`** — L43043 — `private void RepositionControlsForCollapsedlDisplay()`
  relocate the controls on the collapsed display
  Called by: `.ResizeConsole()` (same file), `.CollapseDisplay()` (same file)
- **`.mnuFilter_Click()`** — L43455 — `private void mnuFilter_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuFilter` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuFilterRX2_Click()`** — L43496 — `private void mnuFilterRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuFilterRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuDSP_Click()`** — L43528 — `private void mnuDSP_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuDSP` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuDSPRX2_Click()`** — L43589 — `private void mnuDSPRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuDSPRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuBand_Click()`** — L43647 — `private void mnuBand_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuBand` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuBandRX2_Click()`** — L43697 — `private void mnuBandRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuBandRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupHiddenButton()`** — L43778 — `private void setupHiddenButton()`
  Called by: `.ResizeConsole()` (same file), `.ExpandDisplay()` (same file), `.RepositionControlsForCollapsedlDisplay()` (same file)
- **`.mnuMode_Click()`** — L43785 — `private void mnuMode_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuMode` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuModeRX2_Click()`** — L43832 — `private void mnuModeRX2_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuModeRX2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuShowTopControls_Click()`** — L43876 — `private void mnuShowTopControls_Click(object sender, EventArgs e)`
  handlers for menu display controls events. The persistent state is held on the setup form matching controls
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuShowBandControls_Click()`** — L43881 — `private void mnuShowBandControls_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuShowBandControls` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.mnuShowModeControls_Click()`** — L43886 — `private void mnuShowModeControls_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `mnuShowModeControls` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.AndromedaTopControlsToolStripMenuItem_Click()`** — L43891 — `private void AndromedaTopControlsToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `AndromedaTopControlsToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.AndromedaButtonBarToolStripMenuItem_Click()`** — L43896 — `private void AndromedaButtonBarToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `AndromedaButtonBarToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radBand_CheckedChanged()`** — L43901 — `private void radBand_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radBand` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.eSCToolStripMenuItem_Click()`** — L43927 — `private void eSCToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `eSCToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.showHideDiversity()`** — L43932 — `private void showHideDiversity(bool show, bool starting_up = false)`
  Called by: `.eSCToolStripMenuItem_Click()` (same file), `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.ptbRX1AF_Scroll()`** — L43990 — `private void ptbRX1AF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX1AF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2AF_Scroll()`** — L44002 — `private void ptbRX2AF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2AF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRX1Show_CheckedChanged()`** — L44014 — `private void radRX1Show_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX1Show` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.radRX2Show_CheckedChanged()`** — L44035 — `private void radRX2Show_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRX2Show` checked state changes.
  Called by: `.InitConsole()` (same file)
- **`.ptbAF_DoubleClick()`** — L44056 — `private void ptbAF_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbAF` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX1AF_DoubleClick()`** — L44061 — `private void ptbRX1AF_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX1AF` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2AF_DoubleClick()`** — L44067 — `private void ptbRX2AF_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2AF` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX1StepAttData_ValueChanged()`** — L44073 — `private void udRX1StepAttData_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX1StepAttData` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRX2StepAttData_ValueChanged()`** — L44085 — `private void udRX2StepAttData_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRX2StepAttData` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblPreamp_MouseDoubleClick()`** — L44096 — `private void lblPreamp_MouseDoubleClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lblRX2Preamp_MouseDoubleClick()`** — L44119 — `private void lblRX2Preamp_MouseDoubleClick(object sender, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetFocusMaster()`** — L44129 — `public void SetFocusMaster(bool state)`
  Sets focus master.
  Called by: `.n1mm_delay_Elapsed()` (same file), `.textbox_GotFocus()` (same file), `.textbox_LostFocus()` (same file), `.combo_OpenDropDown()` (same file), `.combo_CloseDropDown()` (same file), `.memoryToolStripMenuItem_Click()` (same file) — and 2 more
- **`.chkFWCATU_CheckedChanged()`** — L44174 — `private void chkFWCATU_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFWCATU` checked state changes.
  Called by: `.GetState()` (same file)
- **`.linearityToolStripMenuItem_Click()`** — L44203 — `private void linearityToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `linearityToolStripMenuItem` is clicked.
  Called by: `.chkFWCATUBypass_MouseDown()` (same file), `.DoOtherButtonAction()` (same file)
- **`.RAtoolStripMenuItem_Click()`** — L44211 — `private void RAtoolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `RAtoolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.SetDigiMode()`** — L44230 — `private void SetDigiMode(int rx, DigiMode.DigiModeSettingState mode, bool bFromTXProfile = false)`
  Sets digi mode.
  Called by: `.LoadedTXProfile()` (same file), `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file)
- **`.chkCWFWKeyer_CheckedChanged()`** — L44315 — `private void chkCWFWKeyer_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCWFWKeyer` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nr_selected_from_text()`** — L44321 — `private void nr_selected_from_text(string text)`
  Called by: `.GetState()` (same file)
- **`.nr_selected_to_text()`** — L44348 — `private string nr_selected_to_text()`
  Called by: `.GetStateList()` (same file)
- **`.incrementNR()`** — L44358 — `private void incrementNR(int rx)`
  Called by: `.chkNR_Click()` (same file), `.chkRX2NR_Click()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.SelectNR()`** — L44372 — `public void SelectNR(int rx, bool incude_sub, int nr)`
  Selects nr.
  Called by: `.Console_KeyDown()` (same file), `.SetRX1Mode()` (same file), `.mnuDSP_Click()` (same file), `.mnuDSPRX2_Click()` (same file), `.SetDigiMode()` (same file), `.DoOtherButtonAction()` (same file)
- **`.GetSelectedNR()`** — L44386 — `public int GetSelectedNR(int rx)`
  Returns selected nr.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.setupNR()`** — L44391 — `private void setupNR(int rx, bool sub)`
  Called by: `.nr_selected_from_text()` (same file), `.incrementNR()` (same file), `.SelectNR()` (same file)
- **`.wbClosing()`** — L44577 — `public void wbClosing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.wBToolStripMenuItem_Click()`** — L44583 — `private void wBToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `wBToolStripMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file)
- **`.UpdatePIVisibilty()`** — L44598 — `public void UpdatePIVisibilty()`
  Updates pivisibilty.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.pIToolStripMenuItem_Click()`** — L44609 — `private void pIToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `pIToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNB_CheckStateChanged()`** — L44618 — `private void chkNB_CheckStateChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2NB_CheckStateChanged()`** — L44666 — `private void chkRX2NB_CheckStateChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2NB2_CheckStateChanged()`** — L44715 — `private void chkRX2NB2_CheckStateChanged(object sender, EventArgs e)`
  RX2 Spectral Noise Blanker (SNB)
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LoadLEDFont()`** — L44732 — `private void LoadLEDFont()`
  Loads ledfont.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddFontMemResourceEx()`** — L44739 — `[DllImport("gdi32.dll", ExactSpelling = true)] private static extern IntPtr AddFontMemResourceEx(byte[] pbFont, int cbFont, IntPtr pdv, out uint pcFonts)`
  Adds font mem resource ex.
  Called by: `.GetCustomFont()` (same file)
- **`.GetCustomFont()`** — L44742 — `static public Font GetCustomFont(byte[] fontData, float size, FontStyle style)`
  Returns custom font.
  Called by: `.LoadLEDFont()` (same file)
- **`.regBox1_Click()`** — L44777 — `private void regBox1_Click(object sender, EventArgs e)`
  ke9ns add open up bandstack window when you click on the bandstack index
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.regBox_Click()`** — L44784 — `private void regBox_Click(object sender, EventArgs e)`
  ke9ns add open up bandstack window when you click on the bandstack index
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTXFilterHigh_ValueChanged()`** — L44793 — `private void udTXFilterHigh_ValueChanged(object sender, EventArgs e)`
  ke9ns add to allow TX filter on main console SSB panel
  Called by: `.udTXFilterHigh_LostFocus()` (same file)
- **`.udTXFilterLow_ValueChanged()`** — L44800 — `private void udTXFilterLow_ValueChanged(object sender, EventArgs e)`
  ke9ns add
  Called by: `.udTXFilterLow_LostFocus()` (same file)
- **`.ForcePureSignalAutoCalDisable()`** — L44807 — `public void ForcePureSignalAutoCalDisable()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkFWCATUBypass_CheckedChanged()`** — L44813 — `private void chkFWCATUBypass_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFWCATUBypass` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRxAnt_CheckedChanged()`** — L44846 — `private void chkRxAnt_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRxAnt` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkQSK_CheckStateChanged()`** — L44858 — `private void chkQSK_CheckStateChanged(object sender, EventArgs e)`
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.ptbPanMainRX_DoubleClick()`** — L44864 — `private void ptbPanMainRX_DoubleClick(object sender, EventArgs e)`
  MW0LGE
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbPanSubRX_DoubleClick()`** — L44877 — `private void ptbPanSubRX_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbPanSubRX` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2Pan_DoubleClick()`** — L44890 — `private void ptbRX2Pan_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2Pan` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setBackground()`** — L44904 — `private void setBackground()`
  MW0LGE
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initialiseRawInput()`** — L44999 — `private void initialiseRawInput()`
  MW0LGE RAWINPUT
  Called by: `.InitConsole()` (same file)
- **`.updateRawInputDevices()`** — L45022 — `private void updateRawInputDevices()`
  Called by: `.initialiseRawInput()` (same file), `.OnDevicesChanged()` (same file)
- **`.OnDevicesChanged()`** — L45105 — `private void OnDevicesChanged(object sender)`
  Handles/raises the devices changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnKeyPressedRaw()`** — L45111 — `private void OnKeyPressedRaw(object sender, RawInputEventArg e)`
  Handles/raises the key pressed raw event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseWheelChanged()`** — L45138 — `private void OnMouseWheelChanged(object sender, RawInputEventArg e)`
  Handles/raises the mouse wheel changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.incrementMutliMeterDisplayMode()`** — L45164 — `private void incrementMutliMeterDisplayMode()`
  Called by: `.txtMultiText_Click()` (same file), `.txtRX2Meter_Click()` (same file)
- **`.txtMultiText_Click()`** — L45176 — `private void txtMultiText_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMultiText` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtRX2Meter_Click()`** — L45181 — `private void txtRX2Meter_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `txtRX2Meter` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripStatusLabel_SeqWarning_Click()`** — L45186 — `private void toolStripStatusLabel_SeqWarning_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripStatusLabel_SeqWarning` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripMenuItem_4by3_DropDownItemClicked()`** — L45191 — `private void toolStripMenuItem_4by3_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolStripMenuItem_16by9_DropDownItemClicked()`** — L45196 — `private void toolStripMenuItem_16by9_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolStripMenuItem_16by10_DropDownItemClicked()`** — L45201 — `private void toolStripMenuItem_16by10_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.youTubeToolStripMenuItem_DropDownItemClicked()`** — L45206 — `private void youTubeToolStripMenuItem_DropDownItemClicked(object sender, ToolStripItemClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetResolution()`** — L45210 — `public void SetResolution(string resolutionString)`
  Sets resolution.
  Called by: `.toolStripMenuItem_4by3_DropDownItemClicked()` (same file), `.toolStripMenuItem_16by9_DropDownItemClicked()` (same file), `.toolStripMenuItem_16by10_DropDownItemClicked()` (same file), `.youTubeToolStripMenuItem_DropDownItemClicked()` (same file)
- **`.includeBordersToolStripMenuItem_Click()`** — L45272 — `private void includeBordersToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `includeBordersToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseDown()`** — L45281 — `private void pnlResizeMeter_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseMove()`** — L45290 — `private void pnlResizeMeter_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseUp()`** — L45325 — `private void pnlResizeMeter_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseEnter()`** — L45330 — `private void pnlResizeMeter_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlResizeMeter_MouseLeave()`** — L45335 — `private void pnlResizeMeter_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlResizeMeter` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.systemToolStripMenuItem_Click()`** — L45340 — `private void systemToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `systemToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.thetisOnlyToolStripMenuItem_Click()`** — L45346 — `private void thetisOnlyToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `thetisOnlyToolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.basicAudioLoadCompletedEvent()`** — L45412 — `private void basicAudioLoadCompletedEvent(bool bLoadedOk)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QSOTimerReset()`** — L45445 — `public void QSOTimerReset(bool bAutoReset = false)`
  Called by: `.HdwMOXChanged()` (same file), `.toolStripStatusLabel_timer_MouseUp()` (same file)
- **`.toolStripStatusLabel_timer_Click()`** — L45463 — `private void toolStripStatusLabel_timer_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `toolStripStatusLabel_timer` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripStatusLabel_timer_MouseUp()`** — L45476 — `private void toolStripStatusLabel_timer_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `toolStripStatusLabel_timer` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateQSOTimer()`** — L45483 — `private void updateQSOTimer()`
  Called by: `.timer_clock_Tick()` (same file)
- **`.updateQSOTimerStatusbar()`** — L45491 — `private void updateQSOTimerStatusbar()`
  Called by: `.timer_clock_Tick()` (same file), `.QSOTimerReset()` (same file)
- **`.chkVFOSync_MouseDown()`** — L45537 — `private void chkVFOSync_MouseDown(object sender, MouseEventArgs e)`
  -- RIGHT click on control shows related setup page // refactored
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNR_MouseDown()`** — L45541 — `private void chkNR_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkNR` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NR_MouseDown()`** — L45546 — `private void chkRX2NR_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2NR` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNB_MouseDown()`** — L45551 — `private void chkNB_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkNB` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDSPNB2_MouseDown()`** — L45555 — `private void chkDSPNB2_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkDSPNB2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NB_MouseDown()`** — L45559 — `private void chkRX2NB_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2NB` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NB2_MouseDown()`** — L45563 — `private void chkRX2NB2_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2NB2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeCWL_MouseDown()`** — L45567 — `private void radModeCWL_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeCWL` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeCWU_MouseDown()`** — L45571 — `private void radModeCWU_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeCWU` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC1_MouseDown()`** — L45575 — `private void chkVAC1_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVAC1` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVAC2_MouseDown()`** — L45579 — `private void chkVAC2_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVAC2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeAM_MouseDown()`** — L45583 — `private void radModeAM_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeAM` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeSAM_MouseDown()`** — L45587 — `private void radModeSAM_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeSAM` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCWAPFEnabled_MouseDown()`** — L45591 — `private void chkCWAPFEnabled_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkCWAPFEnabled` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTNF_MouseDown()`** — L45595 — `private void chkTNF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkTNF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVOX_MouseDown()`** — L45599 — `private void chkVOX_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVOX` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCPDR_MouseDown()`** — L45603 — `private void chkCPDR_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkCPDR` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNoiseGate_MouseDown()`** — L45607 — `private void chkNoiseGate_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkNoiseGate` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radModeFMN_MouseDown()`** — L45611 — `private void radModeFMN_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radModeFMN` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAGC_MouseDown()`** — L45615 — `private void comboAGC_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `comboAGC` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkMicMute_MouseDown()`** — L45619 — `private void chkMicMute_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkMicMute` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboTXProfile_MouseDown()`** — L45623 — `private void comboTXProfile_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `comboTXProfile` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblRF_MouseDown()`** — L45627 — `private void lblRF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblRF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.IsRightButton()`** — L45631 — `private bool IsRightButton(MouseEventArgs e)`
  Called by: `.chkVFOSync_MouseDown()` (same file), `.chkNR_MouseDown()` (same file), `.chkRX2NR_MouseDown()` (same file), `.chkNB_MouseDown()` (same file), `.chkDSPNB2_MouseDown()` (same file), `.chkRX2NB_MouseDown()` (same file) — and 31 more
- **`.ProcessDialogKey()`** — L45636 — `protected override bool ProcessDialogKey(Keys keyData)`
  Processes dialog key.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkVersions()`** — L45653 — `private bool checkVersions()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToolStripMenuItem15_Click()`** — L45754 — `private void ToolStripMenuItem15_Click(object sender, EventArgs e)`
  set TX antenna to 1,2 or 3
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem16_Click()`** — L45759 — `private void ToolStripMenuItem16_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem16` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem17_Click()`** — L45764 — `private void ToolStripMenuItem17_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem17` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem18_Click()`** — L45769 — `private void ToolStripMenuItem18_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem18` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem19_Click()`** — L45774 — `private void ToolStripMenuItem19_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem19` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem20_Click()`** — L45779 — `private void ToolStripMenuItem20_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem20` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_Shown()`** — L45784 — `private void Console_Shown(object sender, EventArgs e)`
  WinForms event handler: runs when `Console` is shown.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chk2TONE_CheckedChanged()`** — L45832 — `private async void chk2TONE_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chk2TONE` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucQuickRecallPad_ButtonClicked()`** — L45866 — `private void ucQuickRecallPad_ButtonClicked(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lblBandStack_Click()`** — L45872 — `private void lblBandStack_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblBandStack` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.getControl()`** — L45880 — `private void getControl(Control cc, Point p, string sub)`
  Returns control.
  Called by: `.gmh_MouseMove()` (same file)
- **`.gmh_MouseMove()`** — L45912 — `private void gmh_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `gmh` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.addDelegates()`** — L46279 — `private void addDelegates()`
  Called by: `.InitConsole()` (same file)
- **`.removeDelegates()`** — L46311 — `private void removeDelegates()`
  Called by: `.Console_Closing()` (same file)
- **`.StopAllTx()`** — L46356 — `public void StopAllTx(string msg = "")`
  Stops all tx.
  Called by: `.timeOutTimer()` (same file)
- **`.timeOutTimer()`** — L46375 — `private void timeOutTimer(string msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXInhibitChanged()`** — L46384 — `private void OnTXInhibitChanged(bool oldState, bool newState)`
  Handles/raises the txinhibit changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOTXChanged()`** — L46389 — `private void OnVFOTXChanged(bool vfoB, bool oldState, bool newState)`
  Handles/raises the vfotxchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnIgnoreDupes()`** — L46395 — `private void OnIgnoreDupes(bool ignore)`
  Handles/raises the ignore dupes event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnHideOnSelect()`** — L46399 — `private void OnHideOnSelect(bool hideOnSelect)`
  Handles/raises the hide on select event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnShowInSpectrum()`** — L46403 — `private void OnShowInSpectrum(bool show)`
  Handles/raises the show in spectrum event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPowerChangeHander()`** — L46410 — `private void OnPowerChangeHander(bool oldPower, bool newPower)`
  Handles/raises the power change hander event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleBSFChange()`** — L46422 — `private void handleBSFChange(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double oldCentreF, do`
  Called by: `.OnSetBandChangeHander()` (same file), `.OnVFOAFrequencyChangeHandler()` (same file)
- **`.updateLastVisited()`** — L46463 — `private void updateLastVisited(BandStackFilter bsf, Band band, DSPMode mode, Filter filter, double freq, double centreF, bool cTUN, int zoomSlider)`
  Called by: `.handleBSFChange()` (same file)
- **`.OnSetBandChangeHander()`** — L46474 — `private void OnSetBandChangeHander(int rx, Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double `
  Handles/raises the set band change hander event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEntryAdd()`** — L46479 — `private void OnEntryAdd(BandStackFilter bsf)`
  Handles/raises the entry add event.
  Called by: `.preBandSelect()` (same file)
- **`.OnEntryUpdate()`** — L46498 — `private void OnEntryUpdate(BandStackFilter bsf, BandStackEntry bse)`
  Handles/raises the entry update event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEntryDelete()`** — L46515 — `private void OnEntryDelete(BandStackFilter bsf, BandStackEntry bse)`
  Handles/raises the entry delete event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCentreFrequencyChanged()`** — L46533 — `private void OnCentreFrequencyChanged(int rx, double oldFreq, double newFreq, Band band, double offset)`
  Handles/raises the centre frequency changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCTUNChanged()`** — L46554 — `private void OnCTUNChanged(int rx, bool oldCTUN, bool newCTUN, Band band)`
  Handles/raises the ctunchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFilterChanged()`** — L46566 — `private void OnFilterChanged(int rx, Filter oldFilter, Filter newFilter, Band band, int low, int high, string sName)`
  Handles/raises the filter changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateBandstackOverlay()`** — L46580 — `private void updateBandstackOverlay(int rx)`
  Called by: `.SetupDisplayEngine()` (same file), `.OnShowInSpectrum()` (same file), `.handleBSFChange()` (same file), `.OnEntryAdd()` (same file), `.OnEntryUpdate()` (same file), `.OnEntryDelete()` (same file) — and 4 more
- **`.OnZoomChanged()`** — L46610 — `private void OnZoomChanged(double oldZoomFactor, double newZoomFactor, int sliderValue)`
  Handles/raises the zoom changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEntryClicked()`** — L46624 — `private void OnEntryClicked(BandStackFilter bsf, BandStackEntry bse, bool updateLastVisited = true, bool obeyHide = true)`
  Handles/raises the entry clicked event.
  Called by: `.OnEntryDelete()` (same file), `.pnlDisplay_MouseUp()` (same file)
- **`.preBandSelect()`** — L46652 — `private void preBandSelect(int rx, Band band, int dir = 0)`
  Called by: `.OnBandBeforeChangeHandler()` (same file), `.SetBandStack()` (same file)
- **`.OnBandBeforeChangeHandler()`** — L46766 — `private void OnBandBeforeChangeHandler(int rx, Band band)`
  Handles/raises the band before change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setRX1BandFromBandStackEntry()`** — L46770 — `private void setRX1BandFromBandStackEntry(in BandStackEntry bse)`
  Sets rx1 band from band stack entry.
  Called by: `.Console_KeyDown()` (same file), `.OnEntryClicked()` (same file), `.preBandSelect()` (same file)
- **`.OnBandChangeHandler()`** — L46785 — `private void OnBandChangeHandler(int rx, Band oldBand, Band newBand)`
  Handles/raises the band change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnModeChangeHandler()`** — L46824 — `private void OnModeChangeHandler(int rx, DSPMode oldMode, DSPMode newMode, Band oldBand, Band newBand)`
  Handles/raises the mode change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOAFrequencyChangeHandler()`** — L46851 — `private void OnVFOAFrequencyChangeHandler(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double o`
  Handles/raises the vfoafrequency change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOBFrequencyChangeHandler()`** — L46872 — `private void OnVFOBFrequencyChangeHandler(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double o`
  Handles/raises the vfobfrequency change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMoxChangeHandler()`** — L46892 — `private void OnMoxChangeHandler(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateStackNumberDisplay()`** — L46919 — `private void updateStackNumberDisplay(BandStackFilter bsf)`
  Called by: `.handleBSFChange()` (same file), `.OnEntryAdd()` (same file), `.OnEntryUpdate()` (same file), `.OnEntryDelete()` (same file), `.OnEntryClicked()` (same file), `.preBandSelect()` (same file) — and 2 more
- **`.RepositionExternalPAButton()`** — L46946 — `public void RepositionExternalPAButton(bool bShow)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkExternalPA_CheckedChanged()`** — L46972 — `private void chkExternalPA_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkExternalPA` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setAGCThresholdPoint()`** — L47070 — `public void setAGCThresholdPoint(double agc_thresh_point, int rx)`
  Sets agcthreshold point.
  Called by: `.tmrAutoAGC_Tick()` (same file)
- **`.tmrAutoAGC_Tick()`** — L47167 — `private void tmrAutoAGC_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `tmrAutoAGC` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRF_Click()`** — L47229 — `private void ptbRF_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRF` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2RF_Click()`** — L47235 — `private void ptbRX2RF_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbRX2RF` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRXEQ_MouseDown()`** — L47241 — `private void chkRXEQ_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRXEQ` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTXEQ_MouseDown()`** — L47250 — `private void chkTXEQ_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkTXEQ` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkFWCATUBypass_MouseDown()`** — L47259 — `private void chkFWCATUBypass_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkFWCATUBypass` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkTUN_MouseDown()`** — L47264 — `private void chkTUN_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkTUN` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chk2TONE_MouseDown()`** — L47269 — `private void chk2TONE_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chk2TONE` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkExternalPA_MouseDown()`** — L47274 — `private void chkExternalPA_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkExternalPA` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRF_MouseDown()`** — L47279 — `private void ptbRF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbRF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ptbRX2RF_MouseDown()`** — L47285 — `private void ptbRX2RF_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbRX2RF` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MouseWheelAGCRX1()`** — L47291 — `private void MouseWheelAGCRX1(object sender, System.Windows.Forms.MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseWheelAGCRX2()`** — L47296 — `private void MouseWheelAGCRX2(object sender, System.Windows.Forms.MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomToBand()`** — L47318 — `public void ZoomToBand(bool bStore)`
  Called by: `.btnDisplayZTB_Click()` (same file), `.DoOtherButtonAction()` (same file)
- **`.btnDisplayZTB_Click()`** — L47412 — `private void btnDisplayZTB_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDisplayZTB` is clicked.
  Called by: `.btnDisplayZTB_MouseUp()` (same file)
- **`.setupZTBButton()`** — L47419 — `private void setupZTBButton()`
  Called by: `.InitConsole()` (same file), `.chkX2TR_CheckedChanged()` (same file), `.chkRX2_CheckedChanged()` (same file), `.chkFWCATU_CheckedChanged()` (same file)
- **`.btnDisplayZTB_MouseUp()`** — L47427 — `private void btnDisplayZTB_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `btnDisplayZTB` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Console_Activated()`** — L47432 — `private void Console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Console_Deactivate()`** — L47439 — `private void Console_Deactivate(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_Button1Clicked()`** — L47455 — `private void infoBar_Button1Clicked(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_Button2Clicked()`** — L47460 — `private void infoBar_Button2Clicked(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleInfoBarButtonClick()`** — L47465 — `private void handleInfoBarButtonClick(ucInfoBar.InfoBarAction e)`
  Called by: `.infoBar_Button1Clicked()` (same file), `.infoBar_Button2Clicked()` (same file)
- **`.infoBar_Button1MouseDown()`** — L47501 — `private void infoBar_Button1MouseDown(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_Button2MouseDown()`** — L47506 — `private void infoBar_Button2MouseDown(object sender, ucInfoBar.InfoBarAction e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showSetupFromInfoBar()`** — L47511 — `private void showSetupFromInfoBar(ucInfoBar.ActionTypes action)`
  Called by: `.infoBar_Button1MouseDown()` (same file), `.infoBar_Button2MouseDown()` (same file)
- **`.infoBar_HideFeedbackChanged()`** — L47553 — `private void infoBar_HideFeedbackChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.infoBar_SwapRedBlueChanged()`** — L47557 — `private void infoBar_SwapRedBlueChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateTuneLabel()`** — L47576 — `public void UpdateTuneLabel(bool bShowLimitValue, System.EventArgs e)`
  Updates tune label.
  Called by: `.OnTuneSliderUpdateTimerTick()` (same file), `.ptbTune_Scroll()` (same file), `.ptbTune_MouseUp()` (same file)
- **`.ptbTune_Scroll()`** — L47652 — `private void ptbTune_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `ptbTune` is scrolled.
  Called by: `.InitConsole()` (same file)
- **`.setupTuneDriveSlider()`** — L47722 — `private void setupTuneDriveSlider()`
  Called by: `.InitConsole()` (same file), `.chkTUN_CheckedChanged()` (same file), `.chk2TONE_CheckedChanged()` (same file)
- **`.setPowerFromDriveSlider()`** — L47755 — `private int setPowerFromDriveSlider(out bool bConstrain, bool bAdjustedBySliderControl)`
  Sets power from drive slider.
  Called by: `.ptbPWR_Scroll()` (same file)
- **`.setPowerFromTuneSlider()`** — L47762 — `private int setPowerFromTuneSlider(out bool bConstrain, bool bAdjustedBySliderControl)`
  Sets power from tune slider.
  Called by: `.ptbTune_Scroll()` (same file)
- **`.SetPowerUsingTargetDBM()`** — L47769 — `public int SetPowerUsingTargetDBM(out bool bConstrain, out double targetdBm, bool bSetPower, bool bFromTune, bool bTwoTone)`
  Sets power using target dbm.
  Called by: `.chkTUN_CheckedChanged()` (same file), `.setPowerFromDriveSlider()` (same file), `.setPowerFromTuneSlider()` (same file)
- **`.enableAudioAmplfier()`** — L47950 — `private void enableAudioAmplfier()`
  Called by: `.chkPower_CheckedChanged()` (same file)
- **`.ptbTune_MouseUp()`** — L47958 — `private void ptbTune_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `ptbTune` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ResetLevelCalibration()`** — L47962 — `public void ResetLevelCalibration(bool ignoreSet = false)`
  Resets level calibration.
  Called by: `.InitConsole()` (same file)
- **`.chkEnableMultiRX_MouseDown()`** — L47982 — `private void chkEnableMultiRX_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkEnableMultiRX` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MultiMeter2UpdateRX1()`** — L48003 — `private async void MultiMeter2UpdateRX1()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeter2UpdateRX2()`** — L48206 — `private async void MultiMeter2UpdateRX2()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateMetersReading()`** — L48370 — `private void updateMetersReading(Reading reading, float value, int rx)`
  Called by: `.MultiMeter2UpdateRX1()` (same file), `.MultiMeter2UpdateRX2()` (same file)
- **`.picMultiMeterDigital_Click()`** — L48375 — `private void picMultiMeterDigital_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `picMultiMeterDigital` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picRX2Meter_Click()`** — L48379 — `private void picRX2Meter_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `picRX2Meter` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucVAC1UnderOver_ClearIssuesClick()`** — L48383 — `private void ucVAC1UnderOver_ClearIssuesClick(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucVAC2UnderOver_ClearIssuesClick()`** — L48389 — `private void ucVAC2UnderOver_ClearIssuesClick(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ptbSquelch_Scroll()`** — L48396 — `private void ptbSquelch_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbSquelch` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkSquelch_CheckStateChanged()` (same file), `.pnlDisplay_DoubleClick()` (same file)
- **`.chkSquelch_CheckStateChanged()`** — L48466 — `private void chkSquelch_CheckStateChanged(object sender, EventArgs e)`
  Called by: `.SetRX1Mode()` (same file)
- **`.handleSqlFM()`** — L48565 — `private void handleSqlFM(int rx, bool bFM, SquelchState force_to_state = SquelchState.LAST)`
  Called by: `.SetRX1Mode()` (same file), `.SetRX2Mode()` (same file), `.SetSqlMode()` (same file)
- **`.chkRX2Squelch_CheckStateChanged()`** — L48648 — `private void chkRX2Squelch_CheckStateChanged(object sender, EventArgs e)`
  Called by: `.SetRX2Mode()` (same file)
- **`.ptbRX2Squelch_Scroll()`** — L48746 — `private void ptbRX2Squelch_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `ptbRX2Squelch` is scrolled.
  Called by: `.InitConsole()` (same file), `.chkRX2Squelch_CheckStateChanged()` (same file)
- **`.chkSquelch_MouseDown()`** — L48814 — `private void chkSquelch_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkSquelch` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2Squelch_MouseDown()`** — L48819 — `private void chkRX2Squelch_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkRX2Squelch` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOSplit_MouseClick()`** — L48824 — `private void chkVFOSplit_MouseClick(object sender, MouseEventArgs e)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.chkVFOSplit_MouseDown()`** — L48839 — `private void chkVFOSplit_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkVFOSplit` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblPAProfile_MouseDown()`** — L48844 — `private void lblPAProfile_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblPAProfile` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateLegacyMeterControls()`** — L48876 — `private void updateLegacyMeterControls(bool expanded)`
  Called by: `.ExpandDisplay()` (same file), `.CollapseDisplay()` (same file)
- **`.InitFFTFillTime()`** — L48949 — `public void InitFFTFillTime(int rx)`
  Inits fftfill time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.finderMenuItem_Click()`** — L48985 — `private void finderMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `finderMenuItem` is clicked.
  Called by: `.showOnStartup()` (same file)
- **`.setupCMasioStatusBar()`** — L49003 — `private void setupCMasioStatusBar()`
  Called by: `.UpdateStatusBarStatusIcons()` (same file)
- **`.setupSerialCatStatusBar()`** — L49029 — `private void setupSerialCatStatusBar()`
  Called by: `.UpdateStatusBarStatusIcons()` (same file)
- **`.UpdateStatusBarStatusIcons()`** — L49097 — `public void UpdateStatusBarStatusIcons(StatusBarIconGroup iconGroup)`
  Updates status bar status icons.
  Called by: `.SetupTCPIPCat()` (same file), `.SetupTCI()` (same file)
- **`.addStatusStripToolTipHandlers()`** — L49125 — `private void addStatusStripToolTipHandlers()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolTipItemMouseHover()`** — L49145 — `private void toolTipItemMouseHover(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toolTipItemMouseLeave()`** — L49177 — `private void toolTipItemMouseLeave(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getRX1stepAttenuatorForBand()`** — L49185 — `private int getRX1stepAttenuatorForBand(Band b)`
  [2.10.3.6]MW0LGE moved all this to functions to make it easier to diagnose issues
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.chkMOX_CheckedChanged2()` (same file), `.SetComboPreampForHPSDR()` (same file)
- **`.setRX1stepAttenuatorForBand()`** — L49191 — `private void setRX1stepAttenuatorForBand(Band b, int att)`
  Sets rx1step attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.GetState()` (same file)
- **`.getRX2stepAttenuatorForBand()`** — L49196 — `private int getRX2stepAttenuatorForBand(Band b)`
  Returns rx2step attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.SetComboPreampForHPSDR()` (same file)
- **`.setRX2stepAttenuatorForBand()`** — L49202 — `private void setRX2stepAttenuatorForBand(Band b, int att)`
  Sets rx2step attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetStateList()` (same file), `.GetState()` (same file)
- **`.getTXstepAttenuatorForBand()`** — L49207 — `private int getTXstepAttenuatorForBand(Band b)`
  Returns txstep attenuator for band.
  Called by: `.GetStateList()` (same file), `.chkPower_CheckedChanged()` (same file), `.chkMOX_CheckedChanged2()` (same file)
- **`.setTXstepAttenuatorForBand()`** — L49213 — `private void setTXstepAttenuatorForBand(Band b, int att)`
  Sets txstep attenuator for band.
  Called by: `.InitConsole()` (same file), `.GetState()` (same file)
- **`.udTXStepAttData_ValueChanged()`** — L49220 — `private void udTXStepAttData_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udTXStepAttData` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetAutoFormStartSetting()`** — L49242 — `public void SetAutoFormStartSetting(string form, bool show)`
  Sets auto form start setting.
  Called by: `.setAutoStartData()` (same file)
- **`.GetAutoFormStartSetting()`** — L49254 — `public bool GetAutoFormStartSetting(string form)`
  Returns auto form start setting.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getAutoStartData()`** — L49268 — `private string getAutoStartData()`
  Returns auto start data.
  Called by: `.GetStateList()` (same file)
- **`.setAutoStartData()`** — L49281 — `private void setAutoStartData(string data)`
  Sets auto start data.
  Called by: `.GetState()` (same file)
- **`.handleShowOnStartWindowsForms()`** — L49299 — `private void handleShowOnStartWindowsForms()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnOpenWindowsFormsTimerEvent()`** — L49314 — `private void OnOpenWindowsFormsTimerEvent(Object source, EventArgs e)`
  Handles/raises the open windows forms timer event event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showOnStartup()`** — L49339 — `private void showOnStartup(string form)`
  Called by: `.OnOpenWindowsFormsTimerEvent()` (same file), `.DoOtherButtonAction()` (same file)
- **`.handleLaunchOnStartUp()`** — L49394 — `private void handleLaunchOnStartUp()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsProcessRunning()`** — L49455 — `private bool IsProcessRunning(string processName)`
  Called by: `.handleLaunchOnStartUp()` (same file)
- **`.FindAllWindowHandlesByProcessId()`** — L49459 — `private static List<IntPtr> FindAllWindowHandlesByProcessId(int processId)`
  Finds all window handles by process id.
  Called by: `.autoLaunchTryToClose()` (same file)
- **`.EnumWindows()`** — L49472 — `[DllImport("user32.dll", SetLastError = true)] private static extern bool EnumWindows(EnumWindowsProc lpEnumFunc, IntPtr lParam)`
  Called by: `.FindAllWindowHandlesByProcessId()` (same file)
- **`.GetWindowThreadProcessId()`** — L49474 — `[DllImport("user32.dll", SetLastError = true)] private static extern uint GetWindowThreadProcessId(IntPtr hWnd, out uint lpdwProcessId)`
  Returns window thread process id.
  Called by: `.FindAllWindowHandlesByProcessId()` (same file)
- **`.PostMessage()`** — L49476 — `[DllImport("user32.dll")] private static extern bool PostMessage(IntPtr hWnd, uint Msg, IntPtr wParam, IntPtr lParam)`
  Called by: `.autoLaunchTryToClose()` (same file)
- **`.PostThreadMessage()`** — L49478 — `[DllImport("user32.dll", SetLastError = true)] private static extern bool PostThreadMessage(uint idThread, uint Msg, IntPtr wParam, IntPtr lParam)`
  Called by: `.autoLaunchTryToClose()` (same file)
- **`.autoLaunchTryToClose()`** — L49485 — `private void autoLaunchTryToClose()`
  Called by: `.Console_Closing()` (same file)
- **`.databaseManagerToolStripMenuItem_Click()`** — L49608 — `private void databaseManagerToolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `databaseManagerToolStripMenuItem` is clicked.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.setupToolStripMenuItem1_Click()`** — L49614 — `private void setupToolStripMenuItem1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `setupToolStripMenuItem1` is clicked.
  Called by: `.showOnStartup()` (same file), `.DoOtherButtonAction()` (same file), `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`)
- **`.ToggleRxTxAnt()`** — L49634 — `public void ToggleRxTxAnt()`
  Toggles rx tx ant.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PopupFilterContextMenu()`** — L49638 — `public void PopupFilterContextMenu(int rx, MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PopupBandstack()`** — L49645 — `public void PopupBandstack(int rx, Band b, bool is_on_top)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showBandStack()`** — L49659 — `private void showBandStack()`
  Called by: `.regBox1_Click()` (same file), `.lblBandStack_Click()` (same file)
- **`.miAbout_Click()`** — L49685 — `private void miAbout_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `miAbout` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.resizeBackgroundImage()`** — L49786 — `private void resizeBackgroundImage()`
  Called by: `.Console_Resize()` (same file)
- **`.setupToolStripMenuItem_MouseUp()`** — L49829 — `private void setupToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `setupToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.displayControlsToolStripMenuItem_MouseUp()`** — L49859 — `private void displayControlsToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `displayControlsToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.dSPToolStripMenuItem_MouseUp()`** — L49864 — `private void dSPToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `dSPToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.bandToolStripMenuItem_MouseUp()`** — L49869 — `private void bandToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `bandToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.modeToolStripMenuItem_MouseUp()`** — L49874 — `private void modeToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `modeToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.filterToolStripMenuItem_MouseUp()`** — L49879 — `private void filterToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `filterToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rX2ToolStripMenuItem_MouseUp()`** — L49884 — `private void rX2ToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `rX2ToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.BPFToolStripMenuItem_MouseUp()`** — L49889 — `private void BPFToolStripMenuItem_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `BPFToolStripMenuItem` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.GetMinimumRXNotchWidth()`** — L49896 — `public double GetMinimumRXNotchWidth(int rx)`
  Returns minimum rxnotch width.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMinimumTXNotchWidth()`** — L49903 — `public double GetMinimumTXNotchWidth()`
  Returns minimum txnotch width.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateMinimumNotchWidthRX()`** — L49907 — `public void UpdateMinimumNotchWidthRX(int rx)`
  Updates minimum notch width rx.
  Called by: `.UpdateDSP()` (same file)
- **`.UpdateMinimumNotchWidthTX()`** — L49939 — `public void UpdateMinimumNotchWidthTX()`
  Updates minimum notch width tx.
  Called by: `.UpdateDSP()` (same file)
- **`.chkFWCATU_MouseUp()`** — L49960 — `private void chkFWCATU_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkFWCATU` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkX2TR_MouseUp()`** — L49969 — `private void chkX2TR_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkX2TR` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.WaterfallRXGradient()`** — L49978 — `public Color[] WaterfallRXGradient()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaterfallTXGradient()`** — L49984 — `public Color[] WaterfallTXGradient()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clampFilterMinMax()`** — L49991 — `private void clampFilterMinMax(int rx, bool use_lowHigh = false, int low = 0, int high = 0)`
  Called by: `.UpdateRX1Filters()` (same file), `.UpdateRX2Filters()` (same file)
- **`.clampFilterShift()`** — L50017 — `private void clampFilterShift(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.pnlDisplay_DoubleClick()`** — L50030 — `private void pnlDisplay_DoubleClick(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlDisplay` is double-clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseDown()`** — L50058 — `private void pnlDisplay_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlDisplay` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseLeave()`** — L50876 — `private void pnlDisplay_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlDisplay` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseMove()`** — L50902 — `unsafe private void pnlDisplay_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlDisplay` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_MouseUp()`** — L52145 — `private void pnlDisplay_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlDisplay` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlDisplay_Resize()`** — L52276 — `private async void pnlDisplay_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlDisplay` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupDisplayMaxBinDetect()`** — L52303 — `private void setupDisplayMaxBinDetect(int rx, bool sub_rx, bool enabled, bool update_enabled_state = true)`
  Called by: `.chkRIT_CheckedChanged()` (same file), `.udRIT_ValueChanged()` (same file), `.OnCentreFrequencyChanged()` (same file), `.OnCTUNChanged()` (same file), `.OnVFOAFrequencyChangeHandler()` (same file), `.OnVFOBFrequencyChangeHandler()` (same file) — and 6 more
- **`.OnFilterEdgesChanged()`** — L52353 — `private void OnFilterEdgesChanged(int rx, Filter newFilter, Band band, int low, int high, string sName, int max_width, int max_shift)`
  Handles/raises the filter edges changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSampleRateChanged()`** — L52358 — `private void OnSampleRateChanged(int rx, int oldSampleRate, int newSampleRate)`
  Handles/raises the sample rate changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFSPChanged()`** — L52363 — `private void OnFSPChanged(int old_fpr, int new_fps)`
  Handles/raises the fspchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkNR_Click()`** — L52369 — `private void chkNR_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNR` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2NR_Click()`** — L52374 — `private void chkRX2NR_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2NR` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.requires_reposition()`** — L52379 — `private bool requires_reposition()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.on_send_floodcontrol_message()`** — L52412 — `private void on_send_floodcontrol_message(string msg, string uid)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoOtherButtonAction()`** — L52457 — `public void DoOtherButtonAction(int rx, OtherButtonId id, MouseButtons button, bool force = false, bool current_state = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleDoOtherButtonActionRightClick()`** — L52743 — `private bool handleDoOtherButtonActionRightClick(int rx, OtherButtonId id)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetNFEnabled()`** — L52880 — `public void SetNFEnabled(int rx, bool state)`
  Sets nfenabled.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetNFEnabled()`** — L52895 — `public bool GetNFEnabled(int rx)`
  Returns nfenabled.
  Called by: `.DoOtherButtonAction()` (same file), `.GetGeneralSetting()` (same file)
- **`.SetBandStack()`** — L52909 — `public void SetBandStack(int rx, int dir)`
  Sets band stack.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAgcT()`** — L52923 — `public int GetAgcT(int rx)`
  Returns agc t.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetAgcT()`** — L52936 — `public void SetAgcT(int rx, int value)`
  Sets agc t.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSql()`** — L52955 — `public int GetSql(int rx)`
  Returns sql.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetSql()`** — L52968 — `public void SetSql(int rx, int value)`
  Sets sql.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetBal()`** — L52987 — `public int GetBal(int rx, bool subrx = false)`
  Returns bal.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetBal()`** — L53003 — `public void SetBal(int rx, int value, bool subrx = false)`
  Sets bal.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAF()`** — L53025 — `public int GetAF(int rx, bool subrx = false)`
  Returns af.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetAF()`** — L53041 — `public bool SetAF(int rx, int value, bool subrx = false)`
  Sets af.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetDisplayZoomGeneralSettings()`** — L53073 — `public DisplayZoomButton GetDisplayZoomGeneralSettings(int rx)`
  Returns display zoom general settings.
  Called by: `.GetGeneralSetting()` (same file)
- **`.SetDisplayZoomGeneralSettings()`** — L53082 — `public void SetDisplayZoomGeneralSettings(int rx, DisplayZoomButton dzb)`
  Sets display zoom general settings.
  Called by: `.ptbDisplayZoom_Scroll()` (same file)
- **`.SetPanAdjust()`** — L53113 — `private void SetPanAdjust(int adjust, bool centre = false)`
  Sets pan adjust.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetOtherButtonState()`** — L53118 — `public bool GetOtherButtonState(OtherButtonId id, int rx)`
  Returns other button state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetSqlMode()`** — L53242 — `public bool SetSqlMode(int rx, SquelchState state)`
  Sets sql mode.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSqlMode()`** — L53275 — `public SquelchState GetSqlMode(int rx)`
  Returns sql mode.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file), `.SetSqlMode()` (same file)
- **`.GetPanSwap()`** — L53288 — `public bool GetPanSwap(int rx)`
  Returns pan swap.
  Called by: `.GetOtherButtonState()` (same file)
- **`.GetSubRX()`** — L53302 — `public bool GetSubRX(int rx)`
  Returns sub rx.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetBin()`** — L53315 — `public bool GetBin(int rx)`
  Returns bin.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetMute()`** — L53328 — `public bool GetMute(int rx)`
  Returns mute.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file), `.SetMute()` (same file)
- **`.GetSelectedNB()`** — L53351 — `public int GetSelectedNB(int rx)`
  Returns selected nb.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetSelectedNB()`** — L53392 — `public bool SetSelectedNB(int rx, int nb)`
  Sets selected nb.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSplit()`** — L53423 — `public bool GetSplit(int rx)`
  Returns split.
  Called by: `.GetOtherButtonState()` (same file)
- **`.GetMNF()`** — L53437 — `public bool GetMNF(int rx)`
  Returns mnf.
  Called by: `.GetOtherButtonState()` (same file)
- **`.GetANF()`** — L53451 — `public bool GetANF(int rx)`
  Returns anf.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetANF()`** — L53464 — `public bool SetANF(int rx, bool state)`
  Sets anf.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetSNB()`** — L53479 — `public bool GetSNB(int rx)`
  Returns snb.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetSNB()`** — L53493 — `public bool SetSNB(int rx, bool state)`
  Sets snb.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAVG()`** — L53508 — `public bool GetAVG(int rx)`
  Returns avg.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetPeak()`** — L53521 — `public bool GetPeak(int rx)`
  Returns peak.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetCTUN()`** — L53534 — `public bool GetCTUN(int rx)`
  Returns ctun.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetAGCMode()`** — L53547 — `public AGCMode GetAGCMode(int rx)`
  Returns agcmode.
  Called by: `.GetOtherButtonState()` (same file)
- **`.SetAGCMode()`** — L53560 — `public bool SetAGCMode(int rx, AGCMode mode)`
  Sets agcmode.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetAGCAuto()`** — L53575 — `public bool GetAGCAuto(int rx)`
  Returns agcauto.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.SetAGCAuto()`** — L53588 — `public bool SetAGCAuto(int rx, bool state)`
  Sets agcauto.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetAVG()`** — L53603 — `public bool SetAVG(int rx, bool state)`
  Sets avg.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetPeak()`** — L53619 — `public bool SetPeak(int rx, bool state)`
  Sets peak.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetCTUN()`** — L53635 — `public bool SetCTUN(int rx, bool state)`
  Sets ctun.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetSubRX()`** — L53651 — `public bool SetSubRX(int rx, bool state)`
  Sets sub rx.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetBin()`** — L53667 — `public bool SetBin(int rx, bool state)`
  Sets bin.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.setMuteAllGeneralSettings()`** — L53682 — `private void setMuteAllGeneralSettings()`
  Sets mute all general settings.
  Called by: `.chkMUT_CheckedChanged()` (same file), `.chkRX2Mute_CheckedChanged()` (same file)
- **`.SetMute()`** — L53686 — `public bool SetMute(int rx, bool state)`
  Sets mute.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetXPAStatus()`** — L53718 — `public (bool in_use, bool enabled) GetXPAStatus()`
  Returns xpastatus.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file)
- **`.GetDisplayMode()`** — L53723 — `public DisplayMode GetDisplayMode(int rx)`
  Returns display mode.
  Called by: `.GetOtherButtonState()` (same file)
- **`.SetDisplayMode()`** — L53735 — `public void SetDisplayMode(int rx, DisplayMode mode)`
  Sets display mode.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.GetGeneralSetting()`** — L53789 — `public bool GetGeneralSetting(int rx, OtherButtonId id)`
  Returns general setting.
  Called by: `.DoOtherButtonAction()` (same file), `.GetOtherButtonState()` (same file), `.initGeneralSettings()` (same file)
- **`.initGeneralSettings()`** — L53887 — `private void initGeneralSettings(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetGeneralSetting()`** — L53938 — `public void SetGeneralSetting(int rx, OtherButtonId id, bool state)`
  Sets general setting.
  Called by: `.chkMicMute_CheckedChanged()` (same file), `.chkVOX_CheckedChanged()` (same file), `.chkNoiseGate_CheckedChanged()` (same file), `.chkShowTXFilter_CheckedChanged()` (same file), `.chkRXEQ_CheckedChanged()` (same file), `.chkTXEQ_CheckedChanged()` (same file) — and 11 more
- **`.DoGeneralSettingAction()`** — L53961 — `public bool DoGeneralSettingAction(int rx, OtherButtonId id, bool state)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.SetHWSampleRateSetting()`** — L54050 — `public void SetHWSampleRateSetting(int rx, int rate)`
  Sets hwsample rate setting.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetATT()`** — L54085 — `public int GetATT(int rx)`
  Returns att.
  Called by: `.GetGeneralSetting()` (same file)
- **`.maxAtt()`** — L54139 — `private int maxAtt()`
  Called by: `.SetATT()` (same file), `.IncrementATT()` (same file)
- **`.SetATT()`** — L54165 — `public bool SetATT(int rx, int att, SetAttMode mode)`
  Sets att.
  Called by: `.DoOtherButtonAction()` (same file)
- **`.setATTGeneralSetting()`** — L54241 — `private void setATTGeneralSetting(int rx)`
  Sets attgeneral setting.
  Called by: `.initGeneralSettings()` (same file)
- **`.IncrementATT()`** — L54327 — `public void IncrementATT(int rx)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.DecrementATT()`** — L54418 — `public void DecrementATT(int rx)`
  Called by: `.DoOtherButtonAction()` (same file)
- **`.handleVfoSyncInitial()`** — L54527 — `private void handleVfoSyncInitial()`
  Called by: `.chkVFOSync_CheckedChanged()` (same file)
- **`.handleVfoSyncFrequency()`** — L54567 — `private void handleVfoSyncFrequency(int rx, bool b_to_a)`
  Called by: `.OnVFOAFrequencyChangeHandler()` (same file), `.OnVFOBFrequencyChangeHandler()` (same file)
- **`.handleVfoSyncMode()`** — L54601 — `private void handleVfoSyncMode(int rx, DSPMode mode)`
  Called by: `.OnModeChangeHandler()` (same file), `.handleVfoSyncFrequency()` (same file)
- **`.handleVfoSyncFilter()`** — L54635 — `private void handleVfoSyncFilter(int rx, Filter newFilter)`
  Called by: `.OnFilterChanged()` (same file), `.handleVfoSyncMode()` (same file)
- **`.btnAPF_type_Click()`** — L54654 — `private void btnAPF_type_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAPF_type` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAPF_type_MouseDown()`** — L54661 — `private void btnAPF_type_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `btnAPF_type` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilter_rx1_MouseUp()`** — L54666 — `private void radFilter_rx1_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radFilter_rx1` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radFilter_rx2_MouseUp()`** — L54671 — `private void radFilter_rx2_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `radFilter_rx2` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.toolStripStatusLabel_PAstatus_MouseUp()`** — L54678 — `private void toolStripStatusLabel_PAstatus_MouseUp(object sender, MouseEventArgs e)`
  Support for Ganymede PA status
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.arp_PowerChanged()`** — L54755 — `private void arp_PowerChanged(bool old_power, bool new_power)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.arp_PlayingingChanged()`** — L54763 — `private void arp_PlayingingChanged(bool playing, string id, string filename, bool isWdsp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaveRecording()`** — L54831 — `public bool WaveRecording(int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.arp_RecordingChanged()`** — L54836 — `private void arp_RecordingChanged(bool recording, string id, string filename)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setPlayRecordStatusBar()`** — L54912 — `private void setPlayRecordStatusBar()`
  Sets play record status bar.
  Called by: `.arp_PlayingingChanged()` (same file), `.arp_RecordingChanged()` (same file)
- **`.waveRecord()`** — L54930 — `private void waveRecord(int rx, bool recording)`
  Called by: `.DoOtherButtonAction()` (same file)

#### `ztb_data` (type, L206)

_No extracted members._

#### `SpectralResult` (type, L21272)

_No extracted members._

#### `HistoricAttenuatorReading` (type, L21331)

_No extracted members._

#### `MeasureKey` (type, L22397)

- **`.Equals()`** — L22412 — `public bool Equals(MeasureKey other) => Text == other.Text && Font.Equals(other.Font)`
  Called by: `.GetState()` (same file), `.SafeTXProfileSet()` (same file), `.Console_KeyDown()` (same file), `.txtVFOAFreq_KeyPress()` (same file), `.txtVFOABand_KeyPress()` (same file), `.txtVFOBFreq_KeyPress()` (same file)
- **`.GetHashCode()`** — L22421 — `public override int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `AutoTuneState` (type, L25669)

_No extracted members._

#### `ProtocolEvent` (type, L25679)

_No extracted members._

#### `TuneLocation` (type, L31881)

_No extracted members._

#### `ModeSpecificPanel` (type, L34788)

_No extracted members._

#### `DisplayZoomButton` (type, L53065)

_No extracted members._

#### `SetAttMode` (type, L54159)

_No extracted members._

#### `DigiMode` (type, L55015)

_No extracted members._

#### `DigiModeSettingState` (type, L55021)

_No extracted members._

#### `AsyncLock` (type, L55041)

- **`.Dispose()`** — L55051 — `public void Dispose()`
  Releases the object’s resources.
  Called by: `.Dispose()` (same file), `.resizeBackgroundImage()` (same file)

#### `MessageFloodControl` (type, L55058)

- **`.FloodControl()`** — L55069 — `public static void FloodControl(string message, string uid, bool ignore_flood = false)`
  Called by: `.BroadcastFreqChange()` (same file), `.BroadcastVFOChange()` (same file)
- **`.Shutdown()`** — L55130 — `public static void Shutdown()`
  Called by: `.Console_Closing()` (same file)
- **`.timer_callback()`** — L55161 — `static void timer_callback(object state_obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.raise_send_message()`** — L55194 — `static void raise_send_message(string message, string uid)`
  Called by: `.FloodControl()` (same file), `.timer_callback()` (same file)

#### `State` (type, L55219)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/console.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
