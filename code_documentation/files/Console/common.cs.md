# `Console/common.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Grab-bag of static helpers: string/number formatting, control lookup, debugging aids, exception reporting.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×49)
  - `Console/console.cs` (calls ×23)
  - `Console/MeterManager.cs` (calls ×16)
  - `Console/Memory/MemoryForm.cs` (calls ×7)
  - `Console/frmAbout.cs` (calls ×6)
  - `Console/clsDBMan.cs` (calls ×5)
  - `Console/frmBandStack2.cs` (calls ×4)
  - `Console/AmpView.cs` (calls ×3)
  - `Console/SpotManager2.cs` (calls ×3)
  - `Console/clsAudioRecordPlayback.cs` (calls ×3)
  - `Console/frmCFCConfig.cs` (calls ×3)
  - `Console/frmFinder.cs` (calls ×3)
  - …and 25 more files
- Uses (outgoing references to other files):
  - `Console/hiperftimer.cs` (references ×1, calls ×1)
  - `Console/database.cs` (calls ×2)
  - `Console/clsDPISafeTools.cs` (calls ×1)
- Most-referenced symbols from other files: `.SaveForm()` (×28), `.OpenUri()` (×15), `.ForceFormOnScreen()` (×9), `.FindNextPowerOf2()` (×9), `.GetComPortNumber()` (×9), `.FindPreviousPowerOf2()` (×8), `.GetVerNum()` (×7), `.RestoreForm()` (×6)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.Add()`** — L191 — `public TypeRenameBinder Add(string oldFullName, Type newType)`
  Called by: `.TabControlInsert()` (same file), `.DeserializeFromBase64()` (same file)
- **`.Create()`** — L207 — `public static TypeRenameBinder Create()`
  Called by: `.DeserializeFromBase64()` (same file)

### Types

#### `FloatExtensions` (type, L70)

- **`.Clamp()`** — L72 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static float Clamp(this float v, float min, float max) => (v < min) ? min : (v > max) ? max : v`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `StringExtensions` (type, L75)

- **`.Truncate()`** — L77 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static string Truncate(this string source, int maxLength)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Contains()`** — L87 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static bool Contains(this string source, string toCheck, StringComparison comp)`
  extend contains to be able to ignore case etc MW0LGE
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Left()`** — L101 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static string Left(this string source, int length)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Right()`** — L116 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static string Right(this string source, int length)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReplaceIgnoreTokenCase()`** — L131 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static string ReplaceIgnoreTokenCase(this string source, string token, string replacement)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `ControlExtentions` (type, L155)

- **`.GetFullName()`** — L157 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static string GetFullName(this Control control)`
  Returns full name.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `TypeRenameBinder` (type, L182)

- **`.BindToType()`** — L197 — `public override Type BindToType(string assemblyName, string typeName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `Common` (type, L214)

- **`.HightlightControl()`** — L230 — `public static void HightlightControl(Control c, bool bHighlight, bool bFromFinder = false)`
  Called by: `.HighlightTXProfileSaveItems()` (`Console/console.cs`), `.HighlightTXProfileSaveItems()` (`Console/eqform.cs`), `.HighlightTXProfileSaveItems()` (`Console/frmCFCConfig.cs`), `.lstResults_SelectedIndexChanged()` (`Console/frmFinder.cs`), `.highlightTXProfileSaveItems()` (`Console/setup.cs`)
- **`.DwmGetWindowAttribute()`** — L304 — `[DllImport("dwmapi.dll")] private static extern int DwmGetWindowAttribute(IntPtr hwnd, int dwAttribute, out RECT pvAttribute, int cbAttribute)`
  Called by: `.DropShadowSize()` (same file)
- **`.DropShadowSize()`** — L315 — `public static Size DropShadowSize(Form f)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ControlList()`** — L339 — `public static void ControlList(Control c, ref List<Control> a)`
  Called by: `.SaveForm()` (same file), `.RestoreForm()` (same file)
- **`.SaveForm()`** — L359 — `public static void SaveForm(Form form, string tablename)`
  Saves form.
  Called by: `.AmpView_FormClosing()` (`Console/AmpView.cs`), `.BandButtonsPopup_FormClosing()` (`Console/Andromeda/BandButtonsPopup.cs`), `.FilterButtonsPopup_FormClosing()` (`Console/Andromeda/FilterButtonsPopup.cs`), `.ModeButtonsPopup_FormClosing()` (`Console/Andromeda/ModeButtonsPopup.cs`), `.ModeDependentSettingsForm_FormClosing()` (`Console/Andromeda/ModeDependentSettingsForm.cs`), `.SliderSettingsForm_Closing()` (`Console/Andromeda/SliderSettingsForm.cs`) — and 22 more
- **`.RestoreForm()`** — L416 — `public static void RestoreForm(Form form, string tablename, bool restore_size)`
  Restores form.
  Called by: `.AmpView_Load()` (`Console/AmpView.cs`), `.InitConsole()` (`Console/console.cs`), `.InitForm()` (`Console/frmBandStack2.cs`), `.RecoverShow()` (`Console/frmBandwidth.cs`), `.Restore()` (`Console/frmDBMan.cs`), `.Show()` (`Console/frmFinder.cs`)
- **`.ForceFormOnScreen()`** — L510 — `public static (bool resized, bool relocated) ForceFormOnScreen(Form f, bool shrink_to_fit = false, bool keep_on_screen = false)`
  Called by: `.RestoreForm()` (same file), `.ContainerFromString()` (`Console/MeterManager.cs`), `.ShowNewLog()` (`Console/clsProgressLog.cs`), `.GetState()` (`Console/console.cs`), `.ShowNotchPopup()` (`Console/console.cs`), `.InitForm()` (`Console/frmBandStack2.cs`) — and 4 more
- **`.TabControlInsert()`** — L522 — `public static void TabControlInsert(TabControl tc, TabPage tp, int index)`
  Called by: `.AddHPSDRPages()` (`Console/setup.cs`)
- **`.SortedComPorts()`** — L547 — `public static string[] SortedComPorts()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RevToString()`** — L567 — `public static string RevToString(uint rev)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLogPath()`** — L576 — `public static void SetLogPath(string sPath)`
  Sets log path.
  Called by: `.Main()` (`Console/console.cs`)
- **`.LogString()`** — L580 — `public static void LogString(string entry)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LogException()`** — L602 — `public static void LogException(Exception e)`
  Called by: `.Application_ThreadException()` (`Console/console.cs`), `.CurrentDomain_UnhandledException()` (`Console/console.cs`), `.RunDisplay()` (`Console/console.cs`)
- **`.GetVerNum()`** — L640 — `public static string GetVerNum(bool include_revision = false, bool include_build = false)`
  Returns ver num.
  Called by: `.MouseUp()` (`Console/MeterManager.cs`), `.checkVersion()` (`Console/clsDBMan.cs`), `.showHelpInfo()` (`Console/console.cs`), `.GetStateList()` (`Console/console.cs`), `.miAbout_Click()` (`Console/console.cs`), `.Init()` (`Console/database.cs`) — and 1 more
- **`.GetFileVersion()`** — L661 — `public static string GetFileVersion()`
  Returns file version.
  Called by: `.ZZVN()` (`Console/CAT/CATCommands.cs`), `.updateSelectedSkin()` (`Console/setup.cs`)
- **`.GetRevision()`** — L669 — `public static string GetRevision()`
  Returns revision.
  Called by: `.GetString()` (`Console/titlebar.cs`)
- **`.setupVersions()`** — L677 — `private static void setupVersions()`
  Called by: `.GetVerNum()` (same file), `.GetFileVersion()` (same file), `.GetRevision()` (same file)
- **`.IsAdministrator()`** — L692 — `public static bool IsAdministrator()`
  Called by: `.Setup()` (`Console/Firewall.cs`), `.OnDragEnter()` (`Console/MeterManager.cs`), `.OnDragOver()` (`Console/MeterManager.cs`), `.OnDragDrop()` (`Console/MeterManager.cs`), `.renderWaveRecord()` (`Console/MeterManager.cs`), `.SetNetworkThrottle()` (`Console/NetworkThrottle.cs`)
- **`.DoubleBuffered()`** — L729 — `public static void DoubleBuffered(Control control, bool enabled)`
  Called by: `.DoubleBufferAll()` (same file)
- **`.FiveDigitHash()`** — L739 — `public static int FiveDigitHash(string str)`
  Called by: `.setTitle()` (`Console/frmMeterDisplay.cs`)
- **`.ColourToString()`** — L758 — `public static string ColourToString(System.Drawing.Color c)`
  Called by: `.ToString()` (`Console/ucMeter.cs`)
- **`.ColourFromString()`** — L762 — `public static System.Drawing.Color ColourFromString(string str)`
  Called by: `.TryParse()` (`Console/MeterManager.cs`), `.TryParse()` (`Console/ucMeter.cs`)
- **`.UVfromDBM()`** — L781 — `public static double UVfromDBM(double dbm)`
  Called by: `.renderHBar()` (`Console/MeterManager.cs`), `.renderSignalTextDisplay()` (`Console/MeterManager.cs`), `.picMultiMeterDigital_Paint()` (`Console/console.cs`), `.picRX2Meter_Paint()` (`Console/console.cs`)
- **`.SMeterFromDBM()`** — L786 — `public static string SMeterFromDBM(double dbm, bool bAboveS9Frequency)`
  Called by: `.picMultiMeterDigital_Paint()` (`Console/console.cs`), `.picRX2Meter_Paint()` (`Console/console.cs`)
- **`.SMeterFromDBM_Spaceless()`** — L834 — `public static string SMeterFromDBM_Spaceless(double dbm, bool bAboveS9Frequency)`
  Called by: `.renderHBar()` (`Console/MeterManager.cs`)
- **`.GetSMeterUnits()`** — L881 — `public static double GetSMeterUnits(double dbm, bool bAboveS9Frequency)`
  Returns smeter units.
  Called by: `.UpdatePeakText()` (`Console/console.cs`)
- **`.SMeterFromDBM2()`** — L888 — `public static void SMeterFromDBM2(double dbm, bool bAboveS9Frequency, out int S, out int over9dBm)`
  Called by: `.renderSignalTextDisplay()` (`Console/MeterManager.cs`)
- **`.DwmSetWindowAttribute()`** — L939 — `[DllImport("dwmapi.dll")] private static extern int DwmSetWindowAttribute(IntPtr hwnd, int attr, ref int attrValue, int attrSize)`
  MW0LGE [2.9.0.8] https://stackoverflow.com/questions/57124243/winforms-dark-title-bar-on-windows-10
  Called by: `.UseImmersiveDarkMode()` (same file)
- **`.UseImmersiveDarkMode()`** — L943 — `public static bool UseImmersiveDarkMode(IntPtr handle, bool enabled)`
  Called by: `.chkConsoleDarkModeTitleBar_CheckedChanged()` (`Console/setup.cs`)
- **`.IsWindows10OrGreater()`** — L959 — `public static bool IsWindows10OrGreater(int build = -1)`
  Called by: `.UseImmersiveDarkMode()` (same file), `.AfterConstructor()` (`Console/setup.cs`)
- **`.DateTimeStringForFile()`** — L965 — `public static string DateTimeStringForFile(string cultureName = "")`
  Called by: `.Export()` (`Console/clsDBMan.cs`), `.ExportBackup()` (`Console/clsDBMan.cs`)
- **`.FadeIn()`** — L988 — `public static async void FadeIn(Form frm, int msTimeToFade = 500, int steps = 20)`
  Called by: `.AmpView_Load()` (`Console/AmpView.cs`), `.ShowAtStartup_LinearityForm()` (`Console/PSForm.cs`), `.showHideDiversity()` (`Console/console.cs`), `.ShowReleaseNotes()` (`Console/frmReleaseNotes.cs`)
- **`.FadeOut()`** — L1000 — `public static async void FadeOut(Form frm, int msTimeToFade = 500, int steps = 20)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CompareVersions()`** — L1013 — `public static int CompareVersions(string version1, string version2)`
  Called by: `.ImportAndMergeDatabase()` (`Console/database.cs`), `.handleVersionInfo()` (`Console/frmAbout.cs`), `.updateSelectedSkin()` (`Console/setup.cs`)
- **`.tryParseVersionPart()`** — L1035 — `private static string tryParseVersionPart(string part)`
  Called by: `.CompareVersions()` (same file)
- **`.IsValidUri()`** — L1044 — `public static bool IsValidUri(string uri)`
  Called by: `.OpenUri()` (same file), `.skinDataReceivedHandler()` (`Console/setup.cs`), `.updateSelectedSkin()` (`Console/setup.cs`)
- **`.OpenUri()`** — L1060 — `public static bool OpenUri(string uri, bool check_uri = true)`
  Opens uri.
  Called by: `.BrowseQRZ()` (`Console/SpotManager2.cs`), `.BrowseHamQTH()` (`Console/SpotManager2.cs`), `.btnSysInfo_Click()` (`Console/frmAbout.cs`), `.btnDXDiag_Click()` (`Console/frmAbout.cs`), `.lnkLicence_LinkClicked()` (`Console/frmAbout.cs`), `.btnVisit_Click()` (`Console/frmAbout.cs`) — and 9 more
- **`.FindNextPowerOf2()`** — L1074 — `public static int FindNextPowerOf2(int n)`
  Finds next power of2.
  Called by: `.InitFFTFillTime()` (`Console/console.cs`), `.udVAC1PropMaxIn_ValueChanged()` (`Console/setup.cs`), `.udVAC1FFMaxIn_ValueChanged()` (`Console/setup.cs`), `.udVAC1PropMaxOut_ValueChanged()` (`Console/setup.cs`), `.udVAC1FFMaxOut_ValueChanged()` (`Console/setup.cs`), `.udVAC2PropMaxIn_ValueChanged()` (`Console/setup.cs`) — and 3 more
- **`.FindPreviousPowerOf2()`** — L1084 — `public static int FindPreviousPowerOf2(int n)`
  Finds previous power of2.
  Called by: `.udVAC1PropMaxIn_ValueChanged()` (`Console/setup.cs`), `.udVAC1FFMaxIn_ValueChanged()` (`Console/setup.cs`), `.udVAC1PropMaxOut_ValueChanged()` (`Console/setup.cs`), `.udVAC1FFMaxOut_ValueChanged()` (`Console/setup.cs`), `.udVAC2PropMaxIn_ValueChanged()` (`Console/setup.cs`), `.udVAC2FFMaxIn_ValueChanged()` (`Console/setup.cs`) — and 2 more
- **`.IsIpv4Valid()`** — L1094 — `public static bool IsIpv4Valid(string ip, int port)`
  Called by: `.StartListeningUDP()` (`Console/MeterManager.cs`), `.StartListeningTCPIP()` (`Console/MeterManager.cs`), `.StartTcpClient()` (`Console/MeterManager.cs`), `.setupUDPEndpoint()` (`Console/setup.cs`)
- **`.SerializeToBase64()`** — L1119 — `public static string SerializeToBase64<T>(T obj)`
  serilisation for any object type
  Called by: `.ContainerToString()` (`Console/MeterManager.cs`)
- **`.DeserializeFromBase64()`** — L1132 — `public static T DeserializeFromBase64<T>(string base64String)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HasArg()`** — L1156 — `public static bool HasArg(string[] args, string arg)`
  Called by: `.Main()` (`Console/console.cs`)
- **`.ArgParam()`** — L1166 — `public static string ArgParam(string[] args, string arg)`
  Called by: `.Main()` (`Console/console.cs`)
- **`.GetLuminance()`** — L1186 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static int GetLuminance(Color c)`
  Returns luminance.
  Called by: `.adjustTextColourForContrast()` (`Console/MeterManager.cs`), `.AddSpot()` (`Console/SpotManager2.cs`)
- **`.rGBtoLin()`** — L1196 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static int rGBtoLin(int col)`
  Called by: `.GetLuminance()` (same file)
- **`.DoubleBufferAll()`** — L1210 — `public static void DoubleBufferAll(Control control, bool enabled)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsValidFilename()`** — L1219 — `public static bool IsValidFilename(string filename)`
  Called by: `.LoadDB()` (`Console/clsDBMan.cs`)
- **`.IsValidPath()`** — L1234 — `public static bool IsValidPath(string path)`
  Called by: `.LoadDB()` (`Console/clsDBMan.cs`)
- **`.DebugPrintCallStack()`** — L1249 — `public static void DebugPrintCallStack(bool only_with_line = true)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FourChar()`** — L1261 — `public static string FourChar(string data1, int data2, Guid guid)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.convertToFourChar()`** — L1272 — `private static string convertToFourChar(string base64Hash)`
  Called by: `.FourChar()` (same file)
- **`.CanCreateFile()`** — L1288 — `public static bool CanCreateFile(string filePath)`
  Called by: `.RecordToFileFromWDSP()` (`Console/clsAudioRecordPlayback.cs`), `.RecordToFileFromPCAudio()` (`Console/clsAudioRecordPlayback.cs`)
- **`.hasWritePermissionOnDir()`** — L1330 — `private static bool hasWritePermissionOnDir(string path)`
  Called by: `.CanCreateFile()` (same file)
- **`.isFileWritable()`** — L1345 — `private static bool isFileWritable(string filePath)`
  Called by: `.CanCreateFile()` (same file)
- **`.GetComPortNumber()`** — L1359 — `public static bool GetComPortNumber(string comport, out int portNumber)`
  Returns com port number.
  Called by: `.initCATandPTTprops()` (`Console/setup.cs`), `.comboCATPort_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT2Port_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT3Port_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT4Port_SelectedIndexChanged()` (`Console/setup.cs`), `.ComboAndromedaCATPort_SelectedIndexChanged()` (`Console/setup.cs`) — and 3 more
- **`.SetProcessPriorityBoost()`** — L1373 — `[DllImport("kernel32.dll")] private static extern bool SetProcessPriorityBoost(IntPtr processHandle, bool disablePriorityBoost)`
  [2.10.3.9]MW0LGE performance related
  Called by: `.DisableForegroundPriorityBoost()` (same file)
- **`.DisableForegroundPriorityBoost()`** — L1375 — `public static void DisableForegroundPriorityBoost()`
  Disables foreground priority boost.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCpuName()`** — L1388 — `public static string GetCpuName()`
  [2.10.3.9]MW0LGE cpu/memory details
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetGpuNames()`** — L1407 — `public static List<string> GetGpuNames()`
  Returns gpu names.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTotalRam()`** — L1437 — `public static string GetTotalRam()`
  Returns total ram.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetInstalledRam()`** — L1460 — `public static string GetInstalledRam()`
  Returns installed ram.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDpiForMonitor()`** — L1487 — `[DllImport("Shcore.dll")] private static extern int GetDpiForMonitor( IntPtr hmonitor, MonitorDpiType dpiType, out uint dpiX,`
  Returns dpi for monitor.
  Called by: `.GetScalingForWindow()` (same file)
- **`.MonitorFromWindow()`** — L1494 — `[DllImport("User32.dll")] private static extern IntPtr MonitorFromWindow( IntPtr hwnd, uint dwFlags)`
  Called by: `.GetScalingForWindow()` (same file)
- **`.GetScalingForWindow()`** — L1499 — `public static int GetScalingForWindow(IntPtr hwnd)`
  Returns scaling for window.
  Called by: `.btnFPSProfile_Click()` (`Console/setup.cs`), `.ValidFpsProfile()` (`Console/setup.cs`)
- **`.ProcessCPUUsage()`** — L1537 — `public static double ProcessCPUUsage()`
  Processes cpuusage.
  Called by: `.timer_cpu_volts_meter_Tick()` (`Console/console.cs`)
- **`.SetThreadExecutionState()`** — L1570 — `[DllImport("kernel32.dll", SetLastError = true)] private static extern ExecutionState SetThreadExecutionState(ExecutionState esFlags)`
  Sets thread execution state.
  Called by: `.PreventSleep()` (same file), `.PreventScreenSaver()` (same file), `.ResumeSleep()` (same file), `.ResumeScreenSaver()` (same file)
- **`.PreventSleep()`** — L1578 — `public static void PreventSleep()`
  Called by: `.chkPreventSleep_CheckedChanged()` (`Console/setup.cs`)
- **`.PreventScreenSaver()`** — L1589 — `public static void PreventScreenSaver()`
  Called by: `.chkPreventScreenSaver_CheckedChanged()` (`Console/setup.cs`)
- **`.ResumeSleep()`** — L1600 — `public static ExecutionState ResumeSleep()`
  Called by: `.ExitConsole()` (`Console/console.cs`), `.chkPreventSleep_CheckedChanged()` (`Console/setup.cs`)
- **`.ResumeScreenSaver()`** — L1620 — `public static ExecutionState ResumeScreenSaver()`
  Called by: `.ExitConsole()` (`Console/console.cs`), `.chkPreventScreenSaver_CheckedChanged()` (`Console/setup.cs`)
- **`.GenerateKeyBase64()`** — L1658 — `public static string GenerateKeyBase64()`
  encryped stuff
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EncryptAndCombineIvToBase64()`** — L1667 — `public static string EncryptAndCombineIvToBase64(string plaintext, byte[] key)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DecryptFromCombinedIvBase64()`** — L1704 — `public static string DecryptFromCombinedIvBase64(string combinedBase64, byte[] key)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Compress_gzip()`** — L1745 — `public static string Compress_gzip(string uncompressed_input)`
  string compress
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Decompress_gzip()`** — L1764 — `public static string Decompress_gzip(string compressed_input)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDiskFreeSpaceExW()`** — L1794 — `[DllImport("kernel32.dll", EntryPoint = "GetDiskFreeSpaceExW")] private static extern bool GetDiskFreeSpaceExW( [MarshalAs(UnmanagedType.LPWStr)] string lpDirectoryName, out ulong `
  Returns disk free space ex w.
  Called by: `.TryGetDriveTotalAndFreeBytes()` (same file)
- **`.TryGetDriveTotalAndFreeBytes()`** — L1801 — `public static bool TryGetDriveTotalAndFreeBytes(string folderPath, out ulong totalBytes, out ulong freeBytes)`
  Called by: `.OkToRecord()` (`Console/clsAudioRecordPlayback.cs`), `.updateRecordingFreeSpace()` (`Console/setup.cs`)
- **`.getUncShareRoot()`** — L1917 — `private static string getUncShareRoot(string uncPath)`
  Returns unc share root.
  Called by: `.TryGetDriveTotalAndFreeBytes()` (same file)

#### `HighlightData` (type, L220)

_No extracted members._

#### `RECT` (type, L306)

_No extracted members._

#### `MonitorDpiType` (type, L1480)

_No extracted members._

#### `ExecutionState` (type, L1562)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/common.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
