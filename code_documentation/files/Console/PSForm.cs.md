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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `PSForm` (type, L61)

- **`.startPSThread()`** — L141 — `private void startPSThread()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StopPSThread()`** — L168 — `public void StopPSThread()`
  Stops psthread.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onConsoleClosingAsync()`** — L182 — `private async Task onConsoleClosingAsync()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onPowerOn()`** — L187 — `private void onPowerOn(bool oldPower, bool newPower)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PSLoop()`** — L193 — `private void PSLoop()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.psdefpeak()`** — L375 — `private void psdefpeak(double value)`
  Called by: `.SetDefaultPeaks()` (same file)
- **`.PSForm_Load()`** — L390 — `private void PSForm_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `PSForm` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetupForm()`** — L395 — `public void SetupForm()`
  Setups form.
  Called by: `.PSForm_Load()` (same file), `.ShowAtStartup_LinearityForm()` (same file)
- **`.PSForm_Closing()`** — L411 — `private void PSForm_Closing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `PSForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CloseAmpView()`** — L429 — `public void CloseAmpView()`
  Closes amp view.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RunAmpv()`** — L450 — `public void RunAmpv()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnPSAmpView_Click()`** — L458 — `private void btnPSAmpView_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPSAmpView` is clicked.
  Called by: `.ShowAtStartup_AmpViewForm()` (same file)
- **`.btnPSCalibrate_Click()`** — L470 — `private void btnPSCalibrate_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPSCalibrate` is clicked.
  Called by: `.SingleCalrun()` (same file)
- **`.SingleCalrun()`** — L485 — `public void SingleCalrun()`
  -W2PA Adds capability for CAT control via console
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnPSReset_Click()`** — L490 — `private void btnPSReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPSReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udPSMoxDelay_ValueChanged()`** — L497 — `private void udPSMoxDelay_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPSMoxDelay` value changes.
  Called by: `.ForcePS()` (same file)
- **`.udPSCalWait_ValueChanged()`** — L502 — `private void udPSCalWait_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPSCalWait` value changes.
  Called by: `.ForcePS()` (same file)
- **`.udPSPhnum_ValueChanged()`** — L507 — `private void udPSPhnum_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPSPhnum` value changes.
  Called by: `.ForcePS()` (same file)
- **`.btnPSTwoToneGen_Click()`** — L512 — `private void btnPSTwoToneGen_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPSTwoToneGen` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnPSSave_Click()`** — L528 — `private void btnPSSave_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPSSave` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnPSRestore_Click()`** — L538 — `private void btnPSRestore_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPSRestore` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetDefaultPeaks()`** — L551 — `public void SetDefaultPeaks()`
  Sets default peaks.
  Called by: `.btnDefaultPeaks_Click()` (same file)
- **`.timer1code()`** — L559 — `private void timer1code()`
  Called by: `.PSLoop()` (same file)
- **`.timer2code()`** — L732 — `private void timer2code()`
  Called by: `.PSLoop()` (same file)
- **`.PSpeak_TextChanged()`** — L819 — `private void PSpeak_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `PSpeak` text changes.
  Called by: `.psdefpeak()` (same file)
- **`.UpdateWarningSetPk()`** — L832 — `public void UpdateWarningSetPk()`
  Updates warning set pk.
  Called by: `.psdefpeak()` (same file), `.PSpeak_TextChanged()` (same file)
- **`.chkPSRelaxPtol_CheckedChanged()`** — L837 — `private void chkPSRelaxPtol_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPSRelaxPtol` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.chkPSAutoAttenuate_CheckedChanged()`** — L845 — `private void chkPSAutoAttenuate_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPSAutoAttenuate` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.checkLoopback_CheckedChanged()`** — L850 — `private void checkLoopback_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `checkLoopback` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPSPin_CheckedChanged()`** — L865 — `private void chkPSPin_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPSPin` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.chkPSMap_CheckedChanged()`** — L873 — `private void chkPSMap_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPSMap` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.chkPSStbl_CheckedChanged()`** — L881 — `private void chkPSStbl_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPSStbl` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.comboPSTint_SelectedIndexChanged()`** — L889 — `private void comboPSTint_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboPSTint` selection changes.
  Called by: `.ForcePS()` (same file)
- **`.btnPSAdvanced_Click()`** — L921 — `private void btnPSAdvanced_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPSAdvanced` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setAdvancedView()`** — L926 — `private void setAdvancedView()`
  Sets advanced view.
  Called by: `.SetupForm()` (same file), `.btnPSAdvanced_Click()` (same file)
- **`.chkPSOnTop_CheckedChanged()`** — L935 — `private void chkPSOnTop_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPSOnTop` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.ShowAtStartup_LinearityForm()`** — L941 — `public void ShowAtStartup_LinearityForm()`
  Shows at startup linearity form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowAtStartup_AmpViewForm()`** — L948 — `public void ShowAtStartup_AmpViewForm()`
  Shows at startup amp view form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ForcePS()`** — L956 — `public void ForcePS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkQuickAttenuate_CheckedChanged()`** — L990 — `private void chkQuickAttenuate_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkQuickAttenuate` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.btnDefaultPeaks_Click()`** — L995 — `private void btnDefaultPeaks_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDefaultPeaks` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShow2ToneMeasurements_CheckedChanged()`** — L1000 — `private void chkShow2ToneMeasurements_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShow2ToneMeasurements` checked state changes.
  Called by: `.ForcePS()` (same file)
- **`.FixAmpViewOnTop()`** — L1005 — `public void FixAmpViewOnTop()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `eCMDState` (type, L115)

_No extracted members._

#### `eAAState` (type, L126)

_No extracted members._

#### `puresignal` (type, L1014)

- **`.SetPSRunCal()`** — L1018 — `[DllImport("wdsp.dll", EntryPoint = "SetPSRunCal", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSRunCal(int channel, bool run)`
  Sets psrun cal.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSMox()`** — L1021 — `[DllImport("wdsp.dll", EntryPoint = "SetPSMox", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSMox(int channel, bool mox)`
  Sets psmox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPSInfo()`** — L1024 — `[DllImport("wdsp.dll", EntryPoint = "GetPSInfo", CallingConvention = CallingConvention.Cdecl)] public static extern void GetPSInfo(int channel, int* info)`
  Returns psinfo.
  Called by: `.GetInfo()` (same file)
- **`.SetPSReset()`** — L1027 — `[DllImport("wdsp.dll", EntryPoint = "SetPSReset", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSReset(int channel, int reset)`
  Sets psreset.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSMancal()`** — L1030 — `[DllImport("wdsp.dll", EntryPoint = "SetPSMancal", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSMancal(int channel, int mancal)`
  Sets psmancal.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSAutomode()`** — L1033 — `[DllImport("wdsp.dll", EntryPoint = "SetPSAutomode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSAutomode(int channel, int automode)`
  Sets psautomode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSTurnon()`** — L1036 — `[DllImport("wdsp.dll", EntryPoint = "SetPSTurnon", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSTurnon(int channel, int turnon)`
  Sets psturnon.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSControl()`** — L1039 — `[DllImport("wdsp.dll", EntryPoint = "SetPSControl", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSControl(int channel, int reset, int mancal, int aut`
  Sets pscontrol.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSLoopDelay()`** — L1042 — `[DllImport("wdsp.dll", EntryPoint = "SetPSLoopDelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSLoopDelay(int channel, double delay)`
  Sets psloop delay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSMoxDelay()`** — L1045 — `[DllImport("wdsp.dll", EntryPoint = "SetPSMoxDelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSMoxDelay(int channel, double delay)`
  Sets psmox delay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSTXDelay()`** — L1048 — `[DllImport("wdsp.dll", EntryPoint = "SetPSTXDelay", CallingConvention = CallingConvention.Cdecl)] public static extern double SetPSTXDelay(int channel, double delay)`
  Sets pstxdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.psccF()`** — L1051 — `[DllImport("wdsp.dll", EntryPoint = "psccF", CallingConvention = CallingConvention.Cdecl)] public static extern void psccF(int channel, int size, float* Itxbuff, float* Qtxbuff, fl`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PSSaveCorr()`** — L1054 — `[DllImport("wdsp.dll", EntryPoint = "PSSaveCorr", CallingConvention = CallingConvention.Cdecl)] public static extern void PSSaveCorr(int channel, string filename)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PSRestoreCorr()`** — L1057 — `[DllImport("wdsp.dll", EntryPoint = "PSRestoreCorr", CallingConvention = CallingConvention.Cdecl)] public static extern void PSRestoreCorr(int channel, string filename)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSHWPeak()`** — L1060 — `[DllImport("wdsp.dll", EntryPoint = "SetPSHWPeak", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSHWPeak(int channel, double peak)`
  Sets pshwpeak.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPSHWPeak()`** — L1063 — `[DllImport("wdsp.dll", EntryPoint = "GetPSHWPeak", CallingConvention = CallingConvention.Cdecl)] public static extern void GetPSHWPeak(int channel, double* peak)`
  Returns pshwpeak.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPSMaxTX()`** — L1066 — `[DllImport("wdsp.dll", EntryPoint = "GetPSMaxTX", CallingConvention = CallingConvention.Cdecl)] public static extern void GetPSMaxTX(int channel, double* maxtx)`
  Returns psmax tx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSPtol()`** — L1069 — `[DllImport("wdsp.dll", EntryPoint = "SetPSPtol", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSPtol(int channel, double ptol)`
  Sets psptol.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPSDisp()`** — L1072 — `[DllImport("wdsp.dll", EntryPoint = "GetPSDisp", CallingConvention = CallingConvention.Cdecl)] public static extern void GetPSDisp(int channel, IntPtr x, IntPtr ym, IntPtr yc, IntP`
  Returns psdisp.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSFeedbackRate()`** — L1075 — `[DllImport("wdsp.dll", EntryPoint = "SetPSFeedbackRate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSFeedbackRate(int channel, int rate)`
  Sets psfeedback rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSPinMode()`** — L1078 — `[DllImport("wdsp.dll", EntryPoint = "SetPSPinMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSPinMode(int channel, int pin)`
  Sets pspin mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSMapMode()`** — L1081 — `[DllImport("wdsp.dll", EntryPoint = "SetPSMapMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSMapMode(int channel, int map)`
  Sets psmap mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSStabilize()`** — L1084 — `[DllImport("wdsp.dll", EntryPoint = "SetPSStabilize", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSStabilize(int channel, int stbl)`
  Sets psstabilize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSIntsAndSpi()`** — L1087 — `[DllImport("wdsp.dll", EntryPoint = "SetPSIntsAndSpi", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSIntsAndSpi(int channel, int ints, int spi)`
  Sets psints and spi.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetInfo()`** — L1108 — `public static void GetInfo(int txachannel)`
  Returns info.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NeedToRecalibrate()`** — L1141 — `public static bool NeedToRecalibrate(int nCurrentATTonTX)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NeedToRecalibrate_HL2()`** — L1146 — `public static bool NeedToRecalibrate_HL2(int nCurrentATTonTX)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `EngineState` (type, L1177)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/PSForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
