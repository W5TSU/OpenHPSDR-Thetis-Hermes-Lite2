# `Console/dsp.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Central DSP settings hub: creates wdsp RX/TX channels and pushes every DSP parameter (NR, NB, AGC, filters, TX processing) down via P/Invoke.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×3)
  - `Console/MeterManager.cs` (references ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `WDSP` (type, L45)

- **`.OpenChannel()`** — L50 — `[DllImport("wdsp.dll", EntryPoint = "OpenChannel", CallingConvention = CallingConvention.Cdecl)] public static extern void OpenChannel(int channel, int in_size, int dsp_size, int i`
  Opens channel.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CloseChannel()`** — L53 — `[DllImport("wdsp.dll", EntryPoint = "CloseChannel", CallingConvention = CallingConvention.Cdecl)] public static extern void CloseChannel(int channel)`
  Closes channel.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetInputBuffsize()`** — L56 — `[DllImport("wdsp.dll", EntryPoint = "SetInputBuffsize", CallingConvention = CallingConvention.Cdecl)] public static extern void SetInputBuffsize(int channel, int in_size)`
  Sets input buffsize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDSPBuffsize()`** — L59 — `[DllImport("wdsp.dll", EntryPoint = "SetDSPBuffsize", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDSPBuffsize(int channel, int dsp_size)`
  Sets dspbuffsize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetInputSamplerate()`** — L62 — `[DllImport("wdsp.dll", EntryPoint = "SetInputSamplerate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetInputSamplerate(int channel, int rate)`
  Sets input samplerate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDSPSamplerate()`** — L65 — `[DllImport("wdsp.dll", EntryPoint = "SetDSPSamplerate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDSPSamplerate(int channel, int rate)`
  Sets dspsamplerate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOutputSamplerate()`** — L68 — `[DllImport("wdsp.dll", EntryPoint = "SetOutputSamplerate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetOutputSamplerate(int channel, int rate)`
  Sets output samplerate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAllRates()`** — L71 — `[DllImport("wdsp.dll", EntryPoint = "SetAllRates", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAllRates(int channel, int in_rate, int dsp_rate, int o`
  Sets all rates.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetChannelState()`** — L74 — `[DllImport("wdsp.dll", EntryPoint = "SetChannelState", CallingConvention = CallingConvention.Cdecl)] public static extern int SetChannelState(int channel, int state, int dmode)`
  Sets channel state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetChannelTDelayUp()`** — L77 — `[DllImport("wdsp.dll", EntryPoint = "SetChannelTDelayUp", CallingConvention = CallingConvention.Cdecl)] public static extern void SetChannelTDelayUp(int channel, double time)`
  Sets channel tdelay up.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetChannelTSlewUp()`** — L80 — `[DllImport("wdsp.dll", EntryPoint = "SetChannelTSlewUp", CallingConvention = CallingConvention.Cdecl)] public static extern void SetChannelTSlewUp(int channel, double time)`
  Sets channel tslew up.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetChannelTDelayDown()`** — L83 — `[DllImport("wdsp.dll", EntryPoint = "SetChannelTDelayDown", CallingConvention = CallingConvention.Cdecl)] public static extern void SetChannelTDelayDown(int channel, double time)`
  Sets channel tdelay down.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetChannelTSlewDown()`** — L86 — `[DllImport("wdsp.dll", EntryPoint = "SetChannelTSlewDown", CallingConvention = CallingConvention.Cdecl)] public static extern void SetChannelTSlewDown(int channel, double time)`
  Sets channel tslew down.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAuSlewTime()`** — L89 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAuSlewTime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAuSlewTime(int channel, double time)`
  Sets txau slew time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAMode()`** — L92 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAMode(int channel, DSPMode mode)`
  Sets rxamode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAMode()`** — L95 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAMode(int channel, DSPMode mode)`
  Sets txamode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.fexchange0()`** — L98 — `[DllImport("wdsp.dll", EntryPoint = "fexchange0", CallingConvention = CallingConvention.Cdecl)] public static extern void fexchange0 (int channel, double* Cin, double* Cout, int* e`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.fexchange2()`** — L101 — `[DllImport("wdsp.dll", EntryPoint = "fexchange2", CallingConvention = CallingConvention.Cdecl)] public static extern void fexchange2(int channel, float* Iin, float* Qin, float* Iou`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCMode()`** — L104 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCMode(int channel, AGCMode mode)`
  Sets rxaagcmode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCFixed()`** — L107 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCFixed", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCFixed(int channel, double fixed_agc)`
  Sets rxaagcfixed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXAAGCTop()`** — L110 — `[DllImport("wdsp.dll", EntryPoint = "GetRXAAGCTop", CallingConvention = CallingConvention.Cdecl)] public static extern void GetRXAAGCTop(int channel, double* max_agc)`
  Returns rxaagctop.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCTop()`** — L113 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCTop", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCTop(int channel, double max_agc)`
  Sets rxaagctop.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCAttack()`** — L116 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCAttack", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCAttack(int channel, int attack)`
  Sets rxaagcattack.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCDecay()`** — L119 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCDecay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCDecay(int channel, int decay)`
  Sets rxaagcdecay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCHang()`** — L122 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCHang", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCHang(int channel, int hang)`
  Sets rxaagchang.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCSlope()`** — L125 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCSlope", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCSlope(int channel, int slope)`
  Sets rxaagcslope.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXAAGCHangThreshold()`** — L128 — `[DllImport("wdsp.dll", EntryPoint = "GetRXAAGCHangThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void GetRXAAGCHangThreshold(int channel, int* hangt`
  Returns rxaagchang threshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCHangThreshold()`** — L131 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCHangThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCHangThreshold(int channel, int hangth`
  Sets rxaagchang threshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXAAGCThresh()`** — L134 — `[DllImport("wdsp.dll", EntryPoint = "GetRXAAGCThresh", CallingConvention = CallingConvention.Cdecl)] public static extern void GetRXAAGCThresh(int channel, double* thresh, double s`
  Returns rxaagcthresh.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCThresh()`** — L137 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCThresh", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCThresh(int channel, double thresh, double si`
  Sets rxaagcthresh.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXAAGCHangLevel()`** — L140 — `[DllImport("wdsp.dll", EntryPoint = "GetRXAAGCHangLevel", CallingConvention = CallingConvention.Cdecl)] public static extern void GetRXAAGCHangLevel(int channel, double* hanglevel)`
  Returns rxaagchang level.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAGCHangLevel()`** — L143 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAGCHangLevel", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAGCHangLevel(int channel, double hanglevel)`
  Sets rxaagchang level.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAALCDecay()`** — L146 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAALCDecay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAALCDecay(int channel, int decay)`
  Sets txaalcdecay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAALCMaxGain()`** — L149 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAALCMaxGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAALCMaxGain(int channel, double maxgain)`
  Sets txaalcmax gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAMDSBMode()`** — L152 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAMDSBMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAMDSBMode(int channel, int sbmode)`
  Sets rxaamdsbmode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAMDFadeLevel()`** — L155 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAMDFadeLevel", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAMDFadeLevel(int channel, int fadelevel)`
  Sets rxaamdfade level.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAMSQRun()`** — L158 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAMSQRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAMSQRun(int channel, bool run)`
  Sets rxaamsqrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAMSQThreshold()`** — L161 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAMSQThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAMSQThreshold(int channel, double threshold`
  Sets rxaamsqthreshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAAMSQMaxTail()`** — L164 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAAMSQMaxTail", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAAMSQMaxTail(int channel, double tail)`
  Sets rxaamsqmax tail.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAAMCarrierLevel()`** — L167 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAAMCarrierLevel", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAAMCarrierLevel(int channel, double carrier`
  Sets txaamcarrier level.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANFRun()`** — L170 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANFRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAANFRun(int channel, bool run)`
  Sets rxaanfrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANFVals()`** — L173 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANFVals", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAANFVals(int channel, int taps, int delay, double `
  Sets rxaanfvals.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANFTaps()`** — L176 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANFTaps", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANFTaps(int channel, int taps)`
  Sets rxaanftaps.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANFDelay()`** — L179 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANFDelay", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANFDelay(int channel, int delay)`
  Sets rxaanfdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANFGain()`** — L182 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANFGain", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANFGain(int channel, double gain)`
  Sets rxaanfgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANFLeakage()`** — L185 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANFLeakage", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANFLeakage(int channel, double leakage)`
  Sets rxaanfleakage.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANFPosition()`** — L188 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANFPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAANFPosition(int channel, int position)`
  Sets rxaanfposition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANRRun()`** — L191 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANRRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAANRRun(int channel, int run)`
  Sets rxaanrrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANRVals()`** — L194 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANRVals", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAANRVals(int channel, int taps, int delay, double `
  Sets rxaanrvals.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANRTaps()`** — L197 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANRTaps", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANRTaps(int channel, int taps)`
  Sets rxaanrtaps.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANRDelay()`** — L200 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANRDelay", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANRDelay(int channel, int delay)`
  Sets rxaanrdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANRGain()`** — L203 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANRGain", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANRGain(int channel, double gain)`
  Sets rxaanrgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANRLeakage()`** — L206 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANRLeakage", CallingConvention = CallingConvention.Cdecl)] public extern static void SetRXAANRLeakage(int channel, double leakage)`
  Sets rxaanrleakage.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAANRPosition()`** — L209 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAANRPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAANRPosition(int channel, int position)`
  Sets rxaanrposition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXABandpassFreqs()`** — L212 — `[DllImport("wdsp.dll", EntryPoint = "SetRXABandpassFreqs", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXABandpassFreqs(int channel, double low, doub`
  Sets rxabandpass freqs.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXABandpassWindow()`** — L215 — `[DllImport("wdsp.dll", EntryPoint = "SetRXABandpassWindow", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXABandpassWindow(int channel, int wintype)`
  Sets rxabandpass window.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXABandpassFreqs()`** — L218 — `[DllImport("wdsp.dll", EntryPoint = "SetTXABandpassFreqs", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXABandpassFreqs(int channel, double low, doub`
  Sets txabandpass freqs.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXABandpassWindow()`** — L221 — `[DllImport("wdsp.dll", EntryPoint = "SetTXABandpassWindow", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXABandpassWindow(int channel, int wintype)`
  Sets txabandpass window.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXACBLRun()`** — L224 — `[DllImport("wdsp.dll", EntryPoint = "SetRXACBLRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXACBLRun(int channel, bool run)`
  Sets rxacblrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXACBLPosition()`** — L227 — `[DllImport("wdsp.dll", EntryPoint = "SetRXACBLPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXACBLPosition(int channel, int position)`
  Sets rxacblposition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACFIRRun()`** — L231 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACFIRRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACFIRRun(int channel, bool run)`
  Sets txacfirrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACompressorRun()`** — L234 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACompressorRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACompressorRun(int channel, bool run)`
  Sets txacompressor run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACompressorGain()`** — L237 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACompressorGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACompressorGain(int channel, double gain)`
  Sets txacompressor gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAosctrlRun()`** — L240 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAosctrlRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAosctrlRun(int channel, bool run)`
  Sets txaosctrl run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRRun()`** — L243 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRRun(int channel, int run)`
  Sets rxaemnrrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXARNNRRun()`** — L247 — `[DllImport("wdsp.dll", EntryPoint = "SetRXARNNRRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXARNNRRun(int channel, int run)`
  rnnoise
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXARNNRPosition()`** — L250 — `[DllImport("wdsp.dll", EntryPoint = "SetRXARNNRPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXARNNRPosition(int channel, int position)`
  Sets rxarnnrposition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RNNRloadModel()`** — L253 — `[DllImport("wdsp.dll", EntryPoint = "RNNRloadModel", CallingConvention = CallingConvention.Cdecl)] public static extern void RNNRloadModel(string file_path)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXARNNRUseDefaultGain()`** — L256 — `[DllImport("wdsp.dll", EntryPoint = "SetRXARNNRUseDefaultGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXARNNRUseDefaultGain(int channel, int us`
  Sets rxarnnruse default gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRRun()`** — L261 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRRun(int channel, int run)`
  libspecbleach
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRPosition()`** — L264 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRPosition(int channel, int position)`
  Sets rxasbnrposition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRreductionAmount()`** — L267 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRreductionAmount", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRreductionAmount(int channel, floa`
  Sets rxasbnrreduction amount.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRsmoothingFactor()`** — L270 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRsmoothingFactor", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRsmoothingFactor(int channel, floa`
  Sets rxasbnrsmoothing factor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRwhiteningFactor()`** — L273 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRwhiteningFactor", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRwhiteningFactor(int channel, floa`
  Sets rxasbnrwhitening factor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRnoiseRescale()`** — L276 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRnoiseRescale", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRnoiseRescale(int channel, float fact`
  Sets rxasbnrnoise rescale.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRpostFilterThreshold()`** — L279 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRpostFilterThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRpostFilterThreshold(int chann`
  Sets rxasbnrpost filter threshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASBNRnoiseScalingType()`** — L282 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASBNRnoiseScalingType", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASBNRnoiseScalingType(int channel, in`
  Sets rxasbnrnoise scaling type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRPosition()`** — L286 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRPosition(int channel, int position)`
  Sets rxaemnrposition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRgainMethod()`** — L289 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRgainMethod", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRgainMethod(int channel, int method)`
  Sets rxaemnrgain method.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRnpeMethod()`** — L292 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRnpeMethod", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRnpeMethod(int channel, int method)`
  Sets rxaemnrnpe method.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRaeRun()`** — L296 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRaeRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRaeRun(int channel, int run)`
  post processing - page 55-57 v1.27 manual
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRpost2Run()`** — L299 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRpost2Run", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRpost2Run(int channel, int run)`
  Sets rxaemnrpost2 run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRpost2Nlevel()`** — L302 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRpost2Nlevel", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRpost2Nlevel(int channel, double nleve`
  Sets rxaemnrpost2 nlevel.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRpost2Factor()`** — L305 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRpost2Factor", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRpost2Factor(int channel, double facto`
  Sets rxaemnrpost2 factor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRpost2Rate()`** — L308 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRpost2Rate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRpost2Rate(int channel, double tc)`
  Sets rxaemnrpost2 rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRpost2Taper()`** — L311 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRpost2Taper", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRpost2Taper(int channel, int taper)`
  Sets rxaemnrpost2 taper.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRtrainZetaThresh()`** — L315 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRtrainZetaThresh", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRtrainZetaThresh(int channel, doub`
  Sets rxaemnrtrain zeta thresh.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEMNRtrainT2()`** — L318 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEMNRtrainT2", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEMNRtrainT2(int channel, double t2)`
  Sets rxaemnrtrain t2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEQRun()`** — L321 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEQRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEQRun(int channel, bool run)`
  Sets rxaeqrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAEQRun()`** — L324 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAEQRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAEQRun(int channel, bool run)`
  Sets txaeqrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAGrphEQ()`** — L327 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAGrphEQ", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAGrphEQ(int channel, int* ptr)`
  Sets rxagrph eq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAGrphEQ()`** — L330 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAGrphEQ", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAGrphEQ(int channel, int* ptr)`
  Sets txagrph eq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAGrphEQ10()`** — L333 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAGrphEQ10", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAGrphEQ10(int channel, int* ptr)`
  Sets rxagrph eq10.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAGrphEQ10()`** — L336 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAGrphEQ10", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAGrphEQ10(int channel, int* ptr)`
  Sets txagrph eq10.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAFMDeviation()`** — L339 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAFMDeviation", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAFMDeviation(int channel, double deviation)`
  Sets rxafmdeviation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAFMSQRun()`** — L342 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAFMSQRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAFMSQRun(int channel, bool run)`
  Sets rxafmsqrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAFMSQThreshold()`** — L345 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAFMSQThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAFMSQThreshold(int channel, double threshold`
  Sets rxafmsqthreshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAFMLimRun()`** — L348 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAFMLimRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAFMLimRun(int channel, bool run)`
  Sets rxafmlim run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAFMLimGain()`** — L351 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAFMLimGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAFMLimGain(int channel, double gaindB)`
  Sets rxafmlim gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAFMAFFilter()`** — L354 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAFMAFFilter", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAFMAFFilter(int channel, double low, double hig`
  Sets rxafmaffilter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAFMAFFilter()`** — L357 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAFMAFFilter", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAFMAFFilter(int channel, double low, double hig`
  Sets txafmaffilter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAFMDeviation()`** — L360 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAFMDeviation", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAFMDeviation(int channel, double deviation)`
  Sets txafmdeviation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAFMEmphPosition()`** — L363 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAFMEmphPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAFMEmphPosition(int channel, bool position)`
  Sets txafmemph position.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACTCSSRun()`** — L366 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACTCSSRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACTCSSRun(int channel, bool run)`
  Sets txactcssrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACTCSSFreq()`** — L369 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACTCSSFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACTCSSFreq(int channel, double freq_hz)`
  Sets txactcssfreq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXACTCSSRun()`** — L372 — `[DllImport("wdsp.dll", EntryPoint = "SetRXACTCSSRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXACTCSSRun(int channel, bool run)`
  Sets rxactcssrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXACTCSSFreq()`** — L375 — `[DllImport("wdsp.dll", EntryPoint = "SetRXACTCSSFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXACTCSSFreq(int channel, double freq_hz)`
  Sets rxactcssfreq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXALevelerTop()`** — L378 — `[DllImport("wdsp.dll", EntryPoint = "SetTXALevelerTop", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXALevelerTop(int channel, double maxgain)`
  Sets txaleveler top.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXALevelerDecay()`** — L381 — `[DllImport("wdsp.dll", EntryPoint = "SetTXALevelerDecay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXALevelerDecay(int channel, int decay)`
  Sets txaleveler decay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXALevelerSt()`** — L384 — `[DllImport("wdsp.dll", EntryPoint = "SetTXALevelerSt", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXALevelerSt(int channel, bool state)`
  Sets txaleveler st.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXAMeter()`** — L387 — `[DllImport("wdsp.dll", EntryPoint = "GetRXAMeter", CallingConvention = CallingConvention.Cdecl)] public static extern double GetRXAMeter(int channel, rxaMeterType meter)`
  Returns rxameter.
  Called by: `.CalculateRXMeter()` (same file)
- **`.GetTXAMeter()`** — L390 — `[DllImport("wdsp.dll", EntryPoint = "GetTXAMeter", CallingConvention = CallingConvention.Cdecl)] public static extern double GetTXAMeter(int channel, txaMeterType meter)`
  Returns txameter.
  Called by: `.CalculateTXMeter()` (same file)
- **`.SetRXAPanelRun()`** — L393 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPanelRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPanelRun(int channel, bool run)`
  Sets rxapanel run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPanelSelect()`** — L396 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPanelSelect", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPanelSelect(int channel, int select)`
  Sets rxapanel select.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPanelGain1()`** — L399 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPanelGain1", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPanelGain1(int channel, double gain)`
  Sets rxapanel gain1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPanelPan()`** — L402 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPanelPan", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPanelPan(int channel, double pan)`
  Sets rxapanel pan.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPanelBinaural()`** — L405 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPanelBinaural", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPanelBinaural(int channel, bool bin)`
  Sets rxapanel binaural.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPanelRun()`** — L408 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPanelRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPanelRun(int channel, bool run)`
  Sets txapanel run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPanelGain1()`** — L411 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPanelGain1", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPanelGain1(int channel, double gain)`
  Sets txapanel gain1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAShiftFreq()`** — L414 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAShiftFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAShiftFreq(int channel, double freq)`
  Sets rxashift freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASpectrum()`** — L417 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASpectrum", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASpectrum(int channel, int flag, int disp, int ss`
  Sets rxaspectrum.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXAGetSpecF1()`** — L420 — `[DllImport("wdsp.dll", EntryPoint = "TXAGetSpecF1", CallingConvention = CallingConvention.Cdecl)] public static extern void TXAGetSpecF1(int channel, float* results)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXAGetaSipF()`** — L423 — `[DllImport("wdsp.dll", EntryPoint = "RXAGetaSipF", CallingConvention = CallingConvention.Cdecl)] public static extern void RXAGetaSipF(int channel, float* results, int size)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXAGetaSipF1()`** — L426 — `[DllImport("wdsp.dll", EntryPoint = "RXAGetaSipF1", CallingConvention = CallingConvention.Cdecl)] public static extern void RXAGetaSipF1(int channel, float* results, int size)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXASetSipPosition()`** — L429 — `[DllImport("wdsp.dll", EntryPoint = "TXASetSipPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void TXASetSipPosition(int channel, int pos)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXASetSipMode()`** — L432 — `[DllImport("wdsp.dll", EntryPoint = "TXASetSipMode", CallingConvention = CallingConvention.Cdecl)] public static extern void TXASetSipMode(int channel, int mode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXASetSipDisplay()`** — L435 — `[DllImport("wdsp.dll", EntryPoint = "TXASetSipDisplay", CallingConvention = CallingConvention.Cdecl)] public static extern void TXASetSipDisplay(int channel, int disp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXAGetaSipF()`** — L438 — `[DllImport("wdsp.dll", EntryPoint = "TXAGetaSipF", CallingConvention = CallingConvention.Cdecl)] public static extern void TXAGetaSipF(int channel, float* results, int size)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXAGetaSipF1()`** — L441 — `[DllImport("wdsp.dll", EntryPoint = "TXAGetaSipF1", CallingConvention = CallingConvention.Cdecl)] public static extern void TXAGetaSipF1(int channel, float* results, int size)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.create_resampleFV()`** — L444 — `[DllImport("wdsp.dll", EntryPoint = "create_resampleFV", CallingConvention = CallingConvention.Cdecl)] public static extern void* create_resampleFV(int in_rate, int out_rate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.xresampleFV()`** — L447 — `[DllImport("wdsp.dll", EntryPoint = "xresampleFV", CallingConvention = CallingConvention.Cdecl)] public static extern void xresampleFV(float* input, float* output, int numsamps, in`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.destroy_resampleFV()`** — L450 — `[DllImport("wdsp.dll", EntryPoint = "destroy_resampleFV", CallingConvention = CallingConvention.Cdecl)] public static extern void destroy_resampleFV(void* ptr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WDSPwisdom()`** — L453 — `[DllImport("wdsp.dll", EntryPoint = "WDSPwisdom", CallingConvention = CallingConvention.Cdecl)] public static extern int WDSPwisdom(string directory)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenRun()`** — L456 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenRun(int channel, int run)`
  Sets rxapre gen run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenMode()`** — L459 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenMode(int channel, int mode)`
  Sets rxapre gen mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenToneMag()`** — L462 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenToneMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenToneMag(int channel, double mag)`
  Sets rxapre gen tone mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenToneFreq()`** — L465 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenToneFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenToneFreq(int channel, double freq)`
  Sets rxapre gen tone freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenNoiseMag()`** — L468 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenNoiseMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenNoiseMag(int channel, double mag)`
  Sets rxapre gen noise mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenSweepMag()`** — L471 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenSweepMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenSweepMag(int channel, double mag)`
  Sets rxapre gen sweep mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenSweepFreq()`** — L474 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenSweepFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenSweepFreq(int channel, double freq1`
  Sets rxapre gen sweep freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAPreGenSweepRate()`** — L477 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAPreGenSweepRate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAPreGenSweepRate(int channel, double rate)`
  Sets rxapre gen sweep rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenRun()`** — L480 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenRun(int channel, int run)`
  Sets txapre gen run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenMode()`** — L483 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenMode(int channel, int mode)`
  Sets txapre gen mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenToneMag()`** — L486 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenToneMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenToneMag(int channel, double mag)`
  Sets txapre gen tone mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenToneFreq()`** — L489 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenToneFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenToneFreq(int channel, double freq)`
  Sets txapre gen tone freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenNoiseMag()`** — L492 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenNoiseMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenNoiseMag(int channel, double mag)`
  Sets txapre gen noise mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenSweepMag()`** — L495 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenSweepMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenSweepMag(int channel, double mag)`
  Sets txapre gen sweep mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenSweepFreq()`** — L498 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenSweepFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenSweepFreq(int channel, double freq1`
  Sets txapre gen sweep freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenSweepRate()`** — L501 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenSweepRate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenSweepRate(int channel, double rate)`
  Sets txapre gen sweep rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenSawtoothMag()`** — L504 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenSawtoothMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenSawtoothMag(int channel, double m`
  Sets txapre gen sawtooth mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenSawtoothFreq()`** — L507 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenSawtoothFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenSawtoothFreq(int channel, double`
  Sets txapre gen sawtooth freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenTriangleMag()`** — L510 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenTriangleMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenTriangleMag(int channel, double m`
  Sets txapre gen triangle mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenTriangleFreq()`** — L513 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenTriangleFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenTriangleFreq(int channel, double`
  Sets txapre gen triangle freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenPulseMag()`** — L516 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenPulseMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenPulseMag(int channel, double mag)`
  Sets txapre gen pulse mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenPulseFreq()`** — L519 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenPulseFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenPulseFreq(int channel, double freq)`
  Sets txapre gen pulse freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenPulseDutyCycle()`** — L522 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenPulseDutyCycle", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenPulseDutyCycle(int channel, do`
  Sets txapre gen pulse duty cycle.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenPulseToneFreq()`** — L525 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenPulseToneFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenPulseToneFreq(int channel, doub`
  Sets txapre gen pulse tone freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPreGenPulseTransition()`** — L528 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPreGenPulseTransition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPreGenPulseTransition(int channel, `
  Sets txapre gen pulse transition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenRun()`** — L531 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenRun(int channel, int run)`
  Sets txapost gen run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenMode()`** — L534 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenMode(int channel, int mode)`
  Sets txapost gen mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenToneFreq()`** — L537 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenToneFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenToneFreq(int channel, double freq)`
  Sets txapost gen tone freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenToneMag()`** — L540 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenToneMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenToneMag(int channel, double mag)`
  Sets txapost gen tone mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTMag()`** — L543 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTMag(int channel, double mag1, doubl`
  Sets txapost gen ttmag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTFreq()`** — L546 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTFreq(int channel, double freq1, do`
  Sets txapost gen ttfreq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenSweepMag()`** — L549 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenSweepMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenSweepMag(int channel, double mag)`
  Sets txapost gen sweep mag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenSweepFreq()`** — L552 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenSweepFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenSweepFreq(int channel, double fre`
  Sets txapost gen sweep freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenSweepRate()`** — L555 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenSweepRate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenSweepRate(int channel, double rat`
  Sets txapost gen sweep rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenPulseMag()`** — L559 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenPulseMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenPulseMag(int channel, double mag)`
  post tune pulse
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenPulseFreq()`** — L562 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenPulseFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenPulseFreq(int channel, double fre`
  Sets txapost gen pulse freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenPulseDutyCycle()`** — L565 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenPulseDutyCycle", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenPulseDutyCycle(int channel, `
  Sets txapost gen pulse duty cycle.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenPulseToneFreq()`** — L568 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenPulseToneFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenPulseToneFreq(int channel, do`
  Sets txapost gen pulse tone freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenPulseTransition()`** — L571 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenPulseTransition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenPulseTransition(int channel`
  Sets txapost gen pulse transition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenPulseIQout()`** — L574 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenPulseIQout", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenPulseIQout(int channel, int IQou`
  Sets txapost gen pulse iqout.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTPulseMag()`** — L579 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTPulseMag", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTPulseMag(int channel, double m`
  post two tone pulse
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTPulseFreq()`** — L582 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTPulseFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTPulseFreq(int channel, double`
  Sets txapost gen ttpulse freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTPulseDutyCycle()`** — L585 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTPulseDutyCycle", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTPulseDutyCycle(int chann`
  Sets txapost gen ttpulse duty cycle.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTPulseToneFreq()`** — L588 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTPulseToneFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTPulseToneFreq(int channel`
  Sets txapost gen ttpulse tone freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTPulseTransition()`** — L591 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTPulseTransition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTPulseTransition(int cha`
  Sets txapost gen ttpulse transition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPostGenTTPulseIQout()`** — L594 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPostGenTTPulseIQout", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPostGenTTPulseIQout(int channel, int `
  Sets txapost gen ttpulse iqout.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetWDSPVersion()`** — L598 — `[DllImport("wdsp.dll", EntryPoint = "GetWDSPVersion", CallingConvention = CallingConvention.Cdecl)] public static extern int GetWDSPVersion()`
  Returns wdspversion.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.create_divEXT()`** — L603 — `[DllImport("wdsp.dll", EntryPoint = "create_divEXT", CallingConvention = CallingConvention.Cdecl)] public static extern void create_divEXT(int id, int run, int nr, int size)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.destroy_divEXT()`** — L606 — `[DllImport("wdsp.dll", EntryPoint = "destroy_divEXT", CallingConvention = CallingConvention.Cdecl)] public static extern void destroy_divEXT (int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTDIVRun()`** — L609 — `[DllImport("wdsp.dll", EntryPoint = "SetEXTDIVRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTDIVRun(int id, int run)`
  Sets extdivrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTDIVNr()`** — L612 — `[DllImport("wdsp.dll", EntryPoint = "SetEXTDIVNr", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTDIVNr (int id, int nr)`
  Sets extdivnr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTDIVOutput()`** — L615 — `[DllImport("wdsp.dll", EntryPoint = "SetEXTDIVOutput", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTDIVOutput (int id, int output)`
  Sets extdivoutput.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTDIVRotate()`** — L618 — `[DllImport("wdsp.dll", EntryPoint = "SetEXTDIVRotate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTDIVRotate(int id, int nr, double* Irotate, doub`
  Sets extdivrotate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.create_eerEXT()`** — L623 — `[DllImport("wdsp.dll", EntryPoint = "create_eerEXT", CallingConvention = CallingConvention.Cdecl)] public static extern void create_eerEXT(int id, int run, int size, int rate, doub`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.destroy_eerEXT()`** — L626 — `[DllImport("wdsp.dll", EntryPoint = "destroy_eerEXT", CallingConvention = CallingConvention.Cdecl)] public static extern void destroy_eerEXT(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.xeerEXTF()`** — L629 — `[DllImport("wdsp.dll", EntryPoint = "xeerEXTF", CallingConvention = CallingConvention.Cdecl)] public static extern void xeerEXTF(int id, float* inI, float* inQ, float* outI, float*`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERRun()`** — L632 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERRun(int id, bool run)`
  Sets eerrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERAMIQ()`** — L635 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERAMIQ", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERAMIQ(int id, bool amiq)`
  Sets eeramiq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERMgain()`** — L638 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERMgain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERMgain (int id, double gain)`
  Sets eermgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERPgain()`** — L641 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERPgain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERPgain(int id, double gain)`
  Sets eerpgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERRunDelays()`** — L644 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERRunDelays", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERRunDelays(int id, bool run)`
  Sets eerrun delays.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERMdelay()`** — L647 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERMdelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERMdelay(int id, double delay)`
  Sets eermdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERPdelay()`** — L650 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERPdelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERPdelay(int id, double delay)`
  Sets eerpdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERSize()`** — L653 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERSize", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERSize(int id, int size)`
  Sets eersize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERSamplerate()`** — L656 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERSamplerate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERSamplerate(int id, int rate)`
  Sets eersamplerate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASPCWRun()`** — L661 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASPCWRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASPCWRun(int channel, bool run)`
  Sets rxaspcwrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASPCWFreq()`** — L664 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASPCWFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASPCWFreq(int channel, double freq)`
  Sets rxaspcwfreq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASPCWBandwidth()`** — L667 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASPCWBandwidth", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASPCWBandwidth(int channel, double bw)`
  Sets rxaspcwbandwidth.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASPCWGain()`** — L670 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASPCWGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASPCWGain(int channel, double gain)`
  Sets rxaspcwgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASPCWSelection()`** — L673 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASPCWSelection", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASPCWSelection(int channel, int selection)`
  Sets rxaspcwselection.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAmpeakRun()`** — L678 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAmpeakRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAmpeakRun(int channel, bool run)`
  Sets rxampeak run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAmpeakFilFreq()`** — L681 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAmpeakFilFreq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAmpeakFilFreq(int channel, int fil, double fr`
  Sets rxampeak fil freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAmpeakFilBw()`** — L684 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAmpeakFilBw", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAmpeakFilBw(int channel, int fil, double bw)`
  Sets rxampeak fil bw.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAmpeakFilGain()`** — L687 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAmpeakFilGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAmpeakFilGain(int channel, int fil, double ga`
  Sets rxampeak fil gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASNBARun()`** — L692 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASNBARun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASNBARun(int channel, bool run)`
  Sets rxasnbarun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASNBAk1()`** — L695 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASNBAk1", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASNBAk1(int channel, double k1)`
  Sets rxasnbak1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASNBAk2()`** — L698 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASNBAk2", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASNBAk2(int channel, double k2)`
  Sets rxasnbak2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPAddNotch()`** — L703 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPAddNotch", CallingConvention = CallingConvention.Cdecl)] public static extern int RXANBPAddNotch(int channel, int notch, double fcenter, `
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPGetNotch()`** — L706 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPGetNotch", CallingConvention = CallingConvention.Cdecl)] public static extern int RXANBPGetNotch(int channel, int notch, double* fcenter,`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPDeleteNotch()`** — L709 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPDeleteNotch", CallingConvention = CallingConvention.Cdecl)] public static extern int RXANBPDeleteNotch(int channel, int notch)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPEditNotch()`** — L712 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPEditNotch", CallingConvention = CallingConvention.Cdecl)] public static extern int RXANBPEditNotch(int channel, int notch, double fcenter`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPGetNumNotches()`** — L715 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPGetNumNotches", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPGetNumNotches(int channel, int* nnotches)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPSetTuneFrequency()`** — L718 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPSetTuneFrequency", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPSetTuneFrequency(int channel, double tun`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPSetShiftFrequency()`** — L721 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPSetShiftFrequency", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPSetShiftFrequency(int channel, double s`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPSetNotchesRun()`** — L724 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPSetNotchesRun", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPSetNotchesRun(int channel, bool run)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPSetFreqs()`** — L727 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPSetFreqs", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPSetFreqs(int channel, double flow, double fhigh)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPSetWindow()`** — L730 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPSetWindow", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPSetWindow(int channel, int wintype)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPGetMinNotchWidth()`** — L733 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPGetMinNotchWidth", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPGetMinNotchWidth(int channel, double* mi`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXANBPSetAutoIncrease()`** — L736 — `[DllImport("wdsp.dll", EntryPoint = "RXANBPSetAutoIncrease", CallingConvention = CallingConvention.Cdecl)] public static extern void RXANBPSetAutoIncrease(int channel, bool autoinc`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASNBAOutputBandwidth()`** — L739 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASNBAOutputBandwidth", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASNBAOutputBandwidth(int channel, doub`
  Sets rxasnbaoutput bandwidth.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXASetMP()`** — L742 — `[DllImport("wdsp.dll", EntryPoint = "RXASetMP", CallingConvention = CallingConvention.Cdecl)] public static extern void RXASetMP(int channel, bool mp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXASetMP()`** — L745 — `[DllImport("wdsp.dll", EntryPoint = "TXASetMP", CallingConvention = CallingConvention.Cdecl)] public static extern void TXASetMP(int channel, bool mp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXASetNC()`** — L748 — `[DllImport("wdsp.dll", EntryPoint = "RXASetNC", CallingConvention = CallingConvention.Cdecl)] public static extern void RXASetNC(int channel, int nc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXASetNC()`** — L751 — `[DllImport("wdsp.dll", EntryPoint = "TXASetNC", CallingConvention = CallingConvention.Cdecl)] public static extern void TXASetNC(int channel, int nc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACFCOMPRun()`** — L755 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACFCOMPRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACFCOMPRun(int channel, int run)`
  cfcomp
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACFCOMPprofile()`** — L758 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACFCOMPprofile", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACFCOMPprofile(int channel, int nfreqs, doub`
  Sets txacfcompprofile.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACFCOMPPosition()`** — L761 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACFCOMPPosition", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACFCOMPPosition(int channel, int pos)`
  Sets txacfcompposition.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACFCOMPPrecomp()`** — L764 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACFCOMPPrecomp", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACFCOMPPrecomp(int channel, double precomp)`
  Sets txacfcompprecomp.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACFCOMPPeqRun()`** — L767 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACFCOMPPeqRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACFCOMPPeqRun(int channel, int run)`
  Sets txacfcomppeq run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXACFCOMPPrePeq()`** — L770 — `[DllImport("wdsp.dll", EntryPoint = "SetTXACFCOMPPrePeq", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXACFCOMPPrePeq(int channel, double prepeq)`
  Sets txacfcomppre peq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPHROTRun()`** — L774 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPHROTRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPHROTRun(int channel, int run)`
  phrot
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPHROTCorner()`** — L777 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPHROTCorner", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPHROTCorner(int channel, double corner)`
  Sets txaphrotcorner.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPHROTNstages()`** — L780 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPHROTNstages", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPHROTNstages(int channel, int nstages)`
  Sets txaphrotnstages.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAPHROTReverse()`** — L783 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAPHROTReverse", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAPHROTReverse(int channel, int reverse)`
  Sets txaphrotreverse.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAEQProfile()`** — L787 — `[DllImport("wdsp.dll", EntryPoint = "SetTXAEQProfile", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAEQProfile(int channel, int nfreqs, double* F, d`
  Sets txaeqprofile.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXAEQProfile()`** — L791 — `[DllImport("wdsp.dll", EntryPoint = "SetRXAEQProfile", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXAEQProfile(int channel, int nfreqs, double* F, d`
  Sets rxaeqprofile.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTXACFCOMPGainAndMask()`** — L795 — `[DllImport("wdsp.dll", EntryPoint = "GetTXACFCOMPGainAndMask", CallingConvention = CallingConvention.Cdecl)] public static extern void GetTXACFCOMPGainAndMask(int channel, double* `
  GetTXACFCOMPGainAndMask
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTXACFCOMPDisplayCompression()`** — L799 — `[DllImport("wdsp.dll", EntryPoint = "GetTXACFCOMPDisplayCompression", CallingConvention = CallingConvention.Cdecl)] public static extern void GetTXACFCOMPDisplayCompression(int cha`
  GetTXACFCOMPDisplayCompression
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASSQLThreshold()`** — L803 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASSQLThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASSQLThreshold(int channel, double threshold`
  Sets rxassqlthreshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASSQLRun()`** — L806 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASSQLRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASSQLRun(int channel, bool run)`
  Sets rxassqlrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASSQLTauMute()`** — L809 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASSQLTauMute", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASSQLTauMute(int channel, double tau_mute)`
  Sets rxassqltau mute.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXASSQLTauUnMute()`** — L812 — `[DllImport("wdsp.dll", EntryPoint = "SetRXASSQLTauUnMute", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXASSQLTauUnMute(int channel, double tau_unmut`
  Sets rxassqltau un mute.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.create_bfcu()`** — L816 — `[DllImport("wdsp.dll", EntryPoint = "create_bfcu", CallingConvention = CallingConvention.Cdecl)] public static extern void create_bfcu(int id, int min_size, int max_size, double ra`
  filter characteristics utility
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getFilterCorners()`** — L819 — `[DllImport("wdsp.dll", EntryPoint = "getFilterCorners", CallingConvention = CallingConvention.Cdecl)] public static extern void getFilterCorners(int id, int* lower_index, int* uppe`
  Returns filter corners.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getFilterCurve()`** — L822 — `[DllImport("wdsp.dll", EntryPoint = "getFilterCurve", CallingConvention = CallingConvention.Cdecl)] public static extern void getFilterCurve(int id, int size, int w_type, int index`
  Returns filter curve.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.destroy_bfcu()`** — L825 — `[DllImport("wdsp.dll", EntryPoint = "destroy_bfcu", CallingConvention = CallingConvention.Cdecl)] public static extern void destroy_bfcu(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.save_impulse_cache()`** — L829 — `[DllImport("wdsp.dll", EntryPoint = "save_impulse_cache", CallingConvention = CallingConvention.Cdecl)] public static extern void save_impulse_cache(string file)`
  WDSP impulse cache - MW0LGE
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.read_impulse_cache()`** — L832 — `[DllImport("wdsp.dll", EntryPoint = "read_impulse_cache", CallingConvention = CallingConvention.Cdecl)] public static extern void read_impulse_cache(string file)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.use_impulse_cache()`** — L835 — `[DllImport("wdsp.dll", EntryPoint = "use_impulse_cache", CallingConvention = CallingConvention.Cdecl)] public static extern void use_impulse_cache(int use)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.init_impulse_cache()`** — L838 — `[DllImport("wdsp.dll", EntryPoint = "init_impulse_cache", CallingConvention = CallingConvention.Cdecl)] public static extern void init_impulse_cache(int use)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.destroy_impulse_cache()`** — L841 — `[DllImport("wdsp.dll", EntryPoint = "destroy_impulse_cache", CallingConvention = CallingConvention.Cdecl)] public static extern void destroy_impulse_cache()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetupDetectMaxBin()`** — L846 — `[DllImport("wdsp.dll", EntryPoint = "SetupDetectMaxBin", CallingConvention = CallingConvention.Cdecl)] public static extern void SetupDetectMaxBin(int run, int disp, int ss, int LO`
  WDSP peak display bin
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDetectMaxBin()`** — L849 — `[DllImport("wdsp.dll", EntryPoint = "GetDetectMaxBin", CallingConvention = CallingConvention.Cdecl)] public static extern double GetDetectMaxBin(int disp)`
  Returns detect max bin.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.id()`** — L926 — `public static int id(uint thread, uint subrx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalculateRXMeter()`** — L947 — `public static float CalculateRXMeter (uint thread, uint subrx, MeterType MT)`
  Calculates rxmeter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalculateTXMeter()`** — L992 — `public static float CalculateTXMeter (uint thread, MeterType MT)`
  Calculates txmeter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `MeterType` (type, L857)

_No extracted members._

#### `rxaMeterType` (type, L887)

_No extracted members._

#### `txaMeterType` (type, L899)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/dsp.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
