# `Console/cmaster.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** P/Invoke wrapper for ChannelMaster: channel setup, audio mixer (`aamix`), radio protocol selection, and stream control from C#.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/TCIServer.cs` (calls ×6, references ×3)
  - `Console/audio.cs` (calls ×2)
  - `Console/clsAudioRecordPlayback.cs` (references ×2)
  - `Console/enums.cs` (references ×2)
  - `Console/wideband.Designer.cs` (references ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `cmaster` (type, L61)

- **`.SetRadioStructure()`** — L68 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRadioStructure", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRadioStructure( int cmSTREAM, int cmRCV`
  Sets radio structure.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateRadio()`** — L82 — `[DllImport("ChannelMaster.dll", EntryPoint = "CreateRadio", CallingConvention = CallingConvention.Cdecl)] public static extern void CreateRadio()`
  Creates radio.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DestroyRadio()`** — L85 — `[DllImport("ChannelMaster.dll", EntryPoint = "DestroyRadio", CallingConvention = CallingConvention.Cdecl)] public static extern void DestroyRadio()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCMDefaultRates()`** — L88 — `[DllImport("ChannelMaster.dll", EntryPoint = "set_cmdefault_rates", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCMDefaultRates(int* xcm_inrates, int `
  Sets cmdefault rates.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetXcmInrate()`** — L91 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetXcmInrate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetXcmInrate(int in_id, int rate)`
  Sets xcm inrate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetXmtrChannelOutrate()`** — L94 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetXmtrChannelOutrate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetXmtrChannelOutrate(int xmtr_id, int`
  Sets xmtr channel outrate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetBuffSize()`** — L97 — `[DllImport("ChannelMaster.dll", EntryPoint = "getbuffsize", CallingConvention = CallingConvention.Cdecl)] public static extern int GetBuffSize (int rate)`
  Returns buff size.
  Called by: `.serviceTCITxProtocol()` (same file)
- **`.inid()`** — L100 — `[DllImport("ChannelMaster.dll", EntryPoint = "inid", CallingConvention = CallingConvention.Cdecl)] public static extern int inid(int stype, int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Inbound()`** — L103 — `[DllImport("ChannelMaster.dll", EntryPoint = "Inbound", CallingConvention = CallingConvention.Cdecl)] public static extern void Inbound(int id, int nsamples, double* data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chid()`** — L106 — `[DllImport("ChannelMaster.dll", EntryPoint = "chid", CallingConvention = CallingConvention.Cdecl)] public static extern int chid (int stream, int subrx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetInputRate()`** — L109 — `[DllImport("ChannelMaster.dll", EntryPoint = "getInputRate", CallingConvention = CallingConvention.Cdecl)] public static extern int GetInputRate (int stype, int id)`
  Returns input rate.
  Called by: `.serviceTCITxProtocol()` (same file), `.OnTCIRxIQOutSamples()` (same file)
- **`.GetChannelOutputRate()`** — L112 — `[DllImport("ChannelMaster.dll", EntryPoint = "getChannelOutputRate", CallingConvention = CallingConvention.Cdecl)] public static extern int GetChannelOutputRate(int stype, int id)`
  Returns channel output rate.
  Called by: `.OnTCIRxAudioOutSamples()` (same file)
- **`.GetCMAstate()`** — L115 — `[DllImport("ChannelMaster.dll", EntryPoint = "getCMAstate", CallingConvention = CallingConvention.Cdecl)] public static extern int GetCMAstate()`
  Returns cmastate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendpOutboundTCIRxIQ()`** — L119 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendpOutboundTCIRxIQ", CallingConvention = CallingConvention.Cdecl)] public static extern void SendpOutboundTCIRxIQ(TCIStreamSamples d`
  Called by: `.SendCallbacks()` (same file)
- **`.SendpOutboundTCIRxAudio()`** — L122 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendpOutboundTCIRxAudio", CallingConvention = CallingConvention.Cdecl)] public static extern void SendpOutboundTCIRxAudio(TCIStreamSam`
  Called by: `.SendCallbacks()` (same file)
- **`.SendpInboundTCITxAudio()`** — L125 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendpInboundTCITxAudio", CallingConvention = CallingConvention.Cdecl)] public static extern void SendpInboundTCITxAudio(TCITxInput del`
  Called by: `.SendCallbacks()` (same file)
- **`.SetRXTCIRun()`** — L128 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRXTCIRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRXTCIRun(int active)`
  Sets rxtcirun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXTCIAudioRun()`** — L131 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTXTCIAudioRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXTCIAudioRun(int txid, int active)`
  Sets txtciaudio run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTCIRxAudioMox()`** — L134 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTCIRxAudioMox", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTCIRxAudioMox(int id, int mox)`
  Sets tcirx audio mox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTCIRxAudioMon()`** — L137 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTCIRxAudioMon", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTCIRxAudioMon(int id, int mon)`
  Sets tcirx audio mon.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTCIRxAudioMonVol()`** — L140 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTCIRxAudioMonVol", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTCIRxAudioMonVol(int id, double vol)`
  Sets tcirx audio mon vol.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LoadRouterAll()`** — L145 — `[DllImport("ChannelMaster.dll", EntryPoint = "LoadRouterAll", CallingConvention = CallingConvention.Cdecl)] public static extern void LoadRouterAll(void* ptr, int id, int sources, `
  router
  Called by: `.CMLoadRouterAll()` (same file)
- **`.LoadRouterControlBit()`** — L149 — `[DllImport("ChannelMaster.dll", EntryPoint = "LoadRouterControlBit", CallingConvention = CallingConvention.Cdecl)] public static extern void LoadRouterControlBit(void* ptr, int id,`
  Loads router control bit.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTopPan3Run()`** — L154 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTopPan3Run", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTopPan3Run(bool run)`
  Sets top pan3 run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRunPanadapter()`** — L157 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRunPanadapter", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRunPanadapter(int id, bool run)`
  Sets run panadapter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPSTxIdx()`** — L162 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetPSTxIdx", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSTxIdx(int id, int idx)`
  Sets pstx idx.
  Called by: `.CMCreateCMaster()` (same file)
- **`.SetPSRxIdx()`** — L165 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetPSRxIdx", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPSRxIdx(int id, int idx)`
  Sets psrx idx.
  Called by: `.CMCreateCMaster()` (same file)
- **`.AllocAnalyzer()`** — L169 — `[DllImport("ChannelMaster.dll", EntryPoint = "alloc_analyzer", CallingConvention = CallingConvention.Cdecl)] public static extern int AllocAnalyzer(int stype, int id, int max_fft_s`
  cmaster multiple analyzers
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FreeAnalyzer()`** — L171 — `[DllImport("ChannelMaster.dll", EntryPoint = "free_analyzer", CallingConvention = CallingConvention.Cdecl)] public static extern int FreeAnalyzer(int disp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RunAnalyzer()`** — L174 — `[DllImport("ChannelMaster.dll", EntryPoint = "run_analyzer", CallingConvention = CallingConvention.Cdecl)] public static extern int RunAnalyzer(int disp, int run)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXAVoxThresh()`** — L179 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPAttackThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXAVoxThresh(int id, double thresh)`
  Sets txavox thresh.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDEXPPeakSignal()`** — L182 — `[DllImport("wdsp.dll", EntryPoint = "GetDEXPPeakSignal", CallingConvention = CallingConvention.Cdecl)] public static extern void GetDEXPPeakSignal(int id, double* peak)`
  Returns dexppeak signal.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPRun()`** — L185 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPRun (int id, bool run)`
  Sets dexprun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPDetectorTau()`** — L188 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPDetectorTau", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPDetectorTau(int id, double tau)`
  Sets dexpdetector tau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPAttackTime()`** — L191 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPAttackTime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPAttackTime(int id, double time)`
  Sets dexpattack time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPReleaseTime()`** — L194 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPReleaseTime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPReleaseTime(int id, double time)`
  Sets dexprelease time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPHoldTime()`** — L197 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPHoldTime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPHoldTime(int id, double time)`
  Sets dexphold time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPExpansionRatio()`** — L200 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPExpansionRatio", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPExpansionRatio(int id, double ratio)`
  Sets dexpexpansion ratio.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPHysteresisRatio()`** — L203 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPHysteresisRatio", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPHysteresisRatio(int id, double ratio)`
  Sets dexphysteresis ratio.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPAttackThreshold()`** — L206 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPAttackThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPAttackThreshold(int id, double thresh)`
  Sets dexpattack threshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPLowCut()`** — L209 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPLowCut", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPLowCut(int id, double lowcut)`
  Sets dexplow cut.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPHighCut()`** — L212 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPHighCut", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPHighCut(int id, double highcut)`
  Sets dexphigh cut.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPRunSideChannelFilter()`** — L215 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPRunSideChannelFilter", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPRunSideChannelFilter(int id, bool `
  Sets dexprun side channel filter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPRunVox()`** — L218 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPRunVox", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPRunVox(int id, bool run)`
  Sets dexprun vox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPRunAudioDelay()`** — L221 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPRunAudioDelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPRunAudioDelay(int id, bool run)`
  Sets dexprun audio delay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDEXPAudioDelay()`** — L224 — `[DllImport("wdsp.dll", EntryPoint = "SetDEXPAudioDelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDEXPAudioDelay(int id, double delay)`
  Sets dexpaudio delay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAntiVOXRun()`** — L227 — `[DllImport("wdsp.dll", EntryPoint = "SetAntiVOXRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAntiVOXRun(int id, bool run)`
  Sets anti voxrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAntiVOXGain()`** — L230 — `[DllImport("wdsp.dll", EntryPoint = "SetAntiVOXGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAntiVOXGain(int id, double gain)`
  Sets anti voxgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAntiVOXDetectorTau()`** — L233 — `[DllImport("wdsp.dll", EntryPoint = "SetAntiVOXDetectorTau", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAntiVOXDetectorTau(int id, double tau)`
  Sets anti voxdetector tau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAntiVOXSourceStates()`** — L236 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAntiVOXSourceStates", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAntiVOXSourceStates(int txid, int `
  Sets anti voxsource states.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAntiVOXSourceWhat()`** — L239 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAntiVOXSourceWhat", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAntiVOXSourceWhat(int txid, int stre`
  Sets anti voxsource what.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetSiphonInsize()`** — L244 — `[DllImport("wdsp.dll", EntryPoint = "SetSiphonInsize", CallingConvention = CallingConvention.Cdecl)] public static extern void SetSiphonInsize (int id, int size)`
  Sets siphon insize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetaSipF1EXT()`** — L247 — `[DllImport("wdsp.dll", EntryPoint = "GetaSipF1EXT", CallingConvention = CallingConvention.Cdecl)] public static extern void GetaSipF1EXT (int id, float* buff, int size)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAAudioMixWhat()`** — L252 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAAudioMixWhat", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAAudioMixWhat(void* ptr, int id, int str`
  Sets aaudio mix what.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAAudioMixState()`** — L255 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAAudioMixState", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAAudioMixState(void* ptr, int id, int s`
  Sets aaudio mix state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAAudioMixStates()`** — L258 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAAudioMixStates", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAAudioMixStates(void* ptr, int id, int`
  Sets aaudio mix states.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAAudioMixVolume()`** — L261 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAAudioMixVolume", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAAudioMixVolume(void* ptr, int id, dou`
  Sets aaudio mix volume.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAAudioMixVol()`** — L264 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAAudioMixVol", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAAudioMixVol(void* ptr, int id, int strea`
  Sets aaudio mix vol.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXVAC()`** — L269 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTXVAC", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXVAC(int txid, int txvac)`
  Sets txvac.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXFixedGainRun()`** — L274 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTXFixedGainRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXFixedGainRun(int id, bool run)`
  Sets txfixed gain run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXFixedGain()`** — L277 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetTXFixedGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTXFixedGain(int id, double Igain, double Q`
  Sets txfixed gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERRun()`** — L282 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERRun(int id, bool run)`
  Sets eerrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERAMIQ()`** — L285 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERAMIQ", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERAMIQ(int id, bool amiq)`
  Sets eeramiq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERMgain()`** — L288 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERMgain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERMgain(int id, double gain)`
  Sets eermgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERPgain()`** — L291 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERPgain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERPgain(int id, double gain)`
  Sets eerpgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERRunDelays()`** — L294 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERRunDelays", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERRunDelays(int id, bool run)`
  Sets eerrun delays.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERMdelay()`** — L297 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERMdelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERMdelay(int id, double delay)`
  Sets eermdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERPdelay()`** — L300 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetEERPdelay", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERPdelay(int id, double delay)`
  Sets eerpdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRANBRun()`** — L305 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRANBRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRANBRun(int stype, int id, bool run)`
  Sets rcvranbrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRANBTau()`** — L308 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRANBTau", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRANBTau(int stype, int id, double tau)`
  Sets rcvranbtau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRANBHangtime()`** — L311 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRANBHangtime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRANBHangtime(int stype, int id, dou`
  Sets rcvranbhangtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRANBAdvtime()`** — L314 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRANBAdvtime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRANBAdvtime(int stype, int id, doubl`
  Sets rcvranbadvtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRANBBacktau()`** — L317 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRANBBacktau", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRANBBacktau(int stype, int id, doubl`
  Sets rcvranbbacktau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRANBThreshold()`** — L320 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRANBThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRANBThreshold(int stype, int id, d`
  Sets rcvranbthreshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRNOBRun()`** — L325 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRNOBRun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRNOBRun(int stype, int id, bool run)`
  Sets rcvrnobrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRNOBMode()`** — L328 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRNOBMode", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRNOBMode(int stype, int id, int mode)`
  Sets rcvrnobmode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRNOBTau()`** — L331 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRNOBTau", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRNOBTau(int stype, int id, double tau)`
  Sets rcvrnobtau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRNOBHangtime()`** — L334 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRNOBHangtime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRNOBHangtime(int stype, int id, dou`
  Sets rcvrnobhangtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRNOBAdvtime()`** — L337 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRNOBAdvtime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRNOBAdvtime(int stype, int id, doubl`
  Sets rcvrnobadvtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRNOBBacktau()`** — L340 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRNOBBacktau", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRNOBBacktau(int stype, int id, doubl`
  Sets rcvrnobbacktau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRCVRNOBThreshold()`** — L343 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetRCVRNOBThreshold", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRCVRNOBThreshold(int stype, int id, d`
  Sets rcvrnobthreshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetAndResetAmpProtect()`** — L348 — `[DllImport("ChannelMaster.dll", EntryPoint = "GetAndResetAmpProtect", CallingConvention = CallingConvention.Cdecl)] extern public static int GetAndResetAmpProtect(int txid)`
  Returns and reset amp protect.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAmpProtectRun()`** — L351 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetAmpProtectRun", CallingConvention = CallingConvention.Cdecl)] extern public static void SetAmpProtectRun(int txid, int run)`
  Sets amp protect run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADCSupply()`** — L354 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetADCSupply", CallingConvention = CallingConvention.Cdecl)] extern public static void SetADCSupply(int txid, int v)`
  Sets adcsupply.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getLEDs()`** — L359 — `[DllImport("ChannelMaster.dll", EntryPoint = "getLEDs", CallingConvention = CallingConvention.Cdecl)] extern public static int getLEDs()`
  Returns leds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCMVersion()`** — L363 — `[DllImport("ChannelMaster.dll", EntryPoint = "GetCMVersion", CallingConvention = CallingConvention.Cdecl)] extern public static int GetCMVersion()`
  version
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCMasioVersion()`** — L366 — `[DllImport("cmASIO.dll", EntryPoint = "GetCMasioVersion", CallingConvention = CallingConvention.Cdecl)] extern public static int GetCMasioVersion()`
  Returns cmasio version.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMCreateCMaster()`** — L500 — `public static void CMCreateCMaster()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMLoadRouterAll()`** — L570 — `public static void CMLoadRouterAll(HPSDRModel model)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetAntiVoxSourceWhat()`** — L934 — `public static void CMSetAntiVoxSourceWhat()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetAudioVolume()`** — L967 — `public static void CMSetAudioVolume(double volume)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetFRXNBRun()`** — L972 — `public static void CMSetFRXNBRun(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetFRXNB2Run()`** — L989 — `public static void CMSetFRXNB2Run(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetSRXWavePlayRun()`** — L1006 — `public static void CMSetSRXWavePlayRun(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetSRXWaveRecordRun()`** — L1025 — `public static void CMSetSRXWaveRecordRun(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetEERRun()`** — L1046 — `public static void CMSetEERRun(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetTXAVoxRun()`** — L1061 — `public static void CMSetTXAVoxRun(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetTXAVoxThresh()`** — L1076 — `public static void CMSetTXAVoxThresh(int id, double thresh)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetTXAPanelGain1()`** — L1083 — `public static void CMSetTXAPanelGain1(int channel)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetScopeRun()`** — L1125 — `public static void CMSetScopeRun(int id, bool run)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetTXOutputLevelRun()`** — L1131 — `public static void CMSetTXOutputLevelRun()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CMSetTXOutputLevel()`** — L1137 — `public static void CMSetTXOutputLevel()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendCallbacks()`** — L1145 — `unsafe public static void SendCallbacks()`
  Sends callbacks.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ensureTCIStreamThreads()`** — L1176 — `private static void ensureTCIStreamThreads()`
  Called by: `.SendCallbacks()` (same file)
- **`.StopTCIStreamThreads()`** — L1198 — `public static void StopTCIStreamThreads()`
  Stops tcistream threads.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TCIRxThreadProc()`** — L1208 — `private static void TCIRxThreadProc()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.serviceTCIRxStreams()`** — L1225 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void serviceTCIRxStreams()`
  Called by: `.TCIRxThreadProc()` (same file)
- **`.TCITxThreadProc()`** — L1257 — `private static void TCITxThreadProc()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.serviceTCITxProtocol()`** — L1273 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void serviceTCITxProtocol()`
  Called by: `.TCITxThreadProc()` (same file)
- **`.resetTCITxState()`** — L1379 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void resetTCITxState()`
  Called by: `.serviceTCITxProtocol()` (same file)
- **`.queueTCITxAudio()`** — L1404 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void queueTCITxAudio(TCIQueuedTxAudio queuedAudio, int targetRate, TCITxStereoInputMode stereoInputMode)`
  Called by: `.serviceTCITxProtocol()` (same file)
- **`.resampleTCITxSamples()`** — L1453 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static unsafe float[] resampleTCITxSamples(float[] input, int inputRate, int targetRate)`
  Called by: `.queueTCITxAudio()` (same file)
- **`.destroyTCIIQResampler()`** — L1500 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static unsafe void destroyTCIIQResampler(int receiver)`
  Called by: `.resampleTCIIQSamples()` (same file)
- **`.resampleTCIIQSamples()`** — L1521 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static unsafe float[] resampleTCIIQSamples(int receiver, float[] input, int inputRate, int targetRate)`
  Called by: `.OnTCIRxIQOutSamples()` (same file)
- **`.rentTCIFloatBuffer()`** — L1607 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static float[] rentTCIFloatBuffer(int length)`
  Called by: `.OnTCIRxIQOutSamples()` (same file), `.OnTCIRxAudioOutSamples()` (same file)
- **`.returnTCIFloatBuffer()`** — L1621 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void returnTCIFloatBuffer(float[] buffer)`
  Called by: `.returnTCIIQBlock()` (same file), `.returnTCIAudioBlock()` (same file), `.OnTCIRxIQOutSamples()` (same file)
- **`.rentTCIIQBlock()`** — L1639 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static TCIIQBlock rentTCIIQBlock()`
  Called by: `.OnTCIRxIQOutSamples()` (same file)
- **`.returnTCIIQBlock()`** — L1650 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void returnTCIIQBlock(TCIIQBlock block)`
  Called by: `.serviceTCIRxStreams()` (same file), `.enqueueTCIIQ()` (same file)
- **`.rentTCIAudioBlock()`** — L1668 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static TCIAudioBlock rentTCIAudioBlock()`
  Called by: `.OnTCIRxAudioOutSamples()` (same file)
- **`.returnTCIAudioBlock()`** — L1679 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void returnTCIAudioBlock(TCIAudioBlock block)`
  Called by: `.serviceTCIRxStreams()` (same file), `.enqueueTCIAudio()` (same file)
- **`.enqueueTCIIQ()`** — L1699 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void enqueueTCIIQ(TCIIQBlock block)`
  Called by: `.OnTCIRxIQOutSamples()` (same file)
- **`.enqueueTCIAudio()`** — L1713 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void enqueueTCIAudio(TCIAudioBlock block)`
  Called by: `.OnTCIRxAudioOutSamples()` (same file)
- **`.tryDequeueTCIIQ()`** — L1727 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static bool tryDequeueTCIIQ(out TCIIQBlock block)`
  Called by: `.serviceTCIRxStreams()` (same file)
- **`.tryDequeueTCIAudio()`** — L1742 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static bool tryDequeueTCIAudio(out TCIAudioBlock block)`
  Called by: `.serviceTCIRxStreams()` (same file)
- **`.OnTCIRxIQOutSamples()`** — L1758 — `private static unsafe void OnTCIRxIQOutSamples(int id, int nsamples, double* data)`
  Handles/raises the tcirx iqout samples event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCIRxAudioOutSamples()`** — L1796 — `private static unsafe void OnTCIRxAudioOutSamples(int id, int nsamples, double* data)`
  Handles/raises the tcirx audio out samples event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTCITxAudioInSamples()`** — L1817 — `private static unsafe void OnTCITxAudioInSamples(int nsamples, double* data)`
  Handles/raises the tcitx audio in samples event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendCBPushVox()`** — L1865 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendCBPushVox", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SendCBPushVox(int id, PushVox Del)`
  set-up a method definition to send the function pointer to the dll
  Called by: `.SendCallbacks()` (same file)
- **`.create_wb()`** — L1908 — `private static void create_wb(int adc)`
  Called by: `.Getwb()` (same file)
- **`.Getwb()`** — L1919 — `public static wideband Getwb(int adc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Hidewb()`** — L1930 — `public static void Hidewb(int adc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Closewb()`** — L1936 — `public static void Closewb(int adc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Savewb()`** — L1942 — `public static void Savewb(int adc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `TCIIQBlock` (type, L444)

_No extracted members._

#### `TCIAudioBlock` (type, L452)

_No extracted members._

#### `VOX` (type, L1954)

- **`.PushVox()`** — L1956 — `unsafe public static void PushVox(int id, int active)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `WaveThing` (type, L1965)

- **`.SendCBCreateWPlay()`** — L1976 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendCBCreateWPlay", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SendCBCreateWPlay(createWplay del)`
  define the method to send the createWavePlayer function pointer
  Called by: `.initWaves()` (same file)
- **`.SetWavePlayerRun()`** — L1979 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetWavePlayerRun", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SetWavePlayerRun(int id, int run)`
  Sets wave player run.
  Called by: `.CMSetSRXWavePlayRun()` (same file)
- **`.SendCBCreateWRecord()`** — L1985 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendCBCreateWRecord", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SendCBCreateWRecord(createWrecord`
  Sends cbcreate wrecord.
  Called by: `.initWaves()` (same file)
- **`.SetWaveRecorderRun()`** — L1988 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetWaveRecorderRun", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SetWaveRecorderRun(int id, int run`
  Sets wave recorder run.
  Called by: `.CMSetSRXWaveRecordRun()` (same file)
- **`.initWaves()`** — L1991 — `unsafe public static void initWaves()`
  Called by: `.SendCallbacks()` (same file)
- **`.SendCBWavePlayer()`** — L2006 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendCBWavePlayer", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SendCBWavePlayer(int id, WPlay del)`
  define the method to send the player function pointer
  Called by: `.createWavePlayer()` (same file)
- **`.createWavePlayer()`** — L2018 — `unsafe public static void createWavePlayer(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendCBWaveRecorder()`** — L2037 — `[DllImport("ChannelMaster", EntryPoint = "SendCBWaveRecorder", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SendCBWaveRecorder(int id, WRecord del`
  Sends cbwave recorder.
  Called by: `.createWaveRecorder()` (same file)
- **`.createWaveRecorder()`** — L2045 — `unsafe public static void createWaveRecorder(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateMox()`** — L2052 — `public static void UpdateMox()`
  Updates mox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `PlayWave` (type, L2064)

- **`.wplay()`** — L2092 — `unsafe public void wplay(int state, double* data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.swizzle()`** — L2118 — `unsafe private static void swizzle(int n, float* I, float* Q, double* C)`
  Called by: `.wplay()` (same file)

#### `RecordWave` (type, L2133)

- **`.wrecord()`** — L2177 — `unsafe public void wrecord(int state, int pos, double* data)`
  'wrecord()' is called by ChannelMaster.dll
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.deswizzle()`** — L2265 — `unsafe private static void deswizzle(int n, double* C, float* I, float* Q)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `Scope` (type, L2282)

- **`.SendCBCreateScope()`** — L2287 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendCBCreateScope", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SendCBCreateScope(createscope del)`
  Sends cbcreate scope.
  Called by: `.initScope()` (same file)
- **`.SendCBScope()`** — L2292 — `[DllImport("ChannelMaster.dll", EntryPoint = "SendCBScope", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SendCBScope(int id, Xscope del)`
  Sends cbscope.
  Called by: `.createScope()` (same file)
- **`.SetScopeRun()`** — L2295 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetScopeRun", CallingConvention = CallingConvention.Cdecl)] unsafe public static extern void SetScopeRun(int id, int run)`
  Sets scope run.
  Called by: `.CMSetScopeRun()` (same file)
- **`.createScope()`** — L2302 — `unsafe public static void createScope(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initScope()`** — L2311 — `unsafe public static void initScope()`
  Called by: `.SendCallbacks()` (same file)

#### `DoScope` (type, L2317)

- **`.xscope()`** — L2338 — `unsafe public void xscope(int state, double* data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.deswizzle()`** — L2365 — `unsafe private static void deswizzle(int n, double* C, float* I, float* Q)`
  Called by: `.wrecord()` (same file), `.xscope()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/cmaster.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
