# `Console/HPSDR/NetworkIOImports.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** The P/Invoke surface into ChannelMaster's network code — every `extern` for radio I/O (128 imports).

## How this file is used

- Used by (incoming references from other files):
  - `Console/HPSDR/NetworkIO.cs` (calls ×3)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.nativeInitMetis()` (×1), `.SetOutputPowerFactor()` (×1), `.SetVFOfreq()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `NetworkIO` (type, L7)

- **`.SetOutputPowerFactor()`** — L9 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetOutputPowerFactor(int i)`
  Sets output power factor.
  Called by: `.SetOutputPower()` (`Console/HPSDR/NetworkIO.cs`)
- **`.DeInitMetisSockets()`** — L12 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void DeInitMetisSockets()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nativeInitMetis()`** — L15 — `[DllImport("ChannelMaster.dll", CharSet = CharSet.Ansi, CallingConvention = CallingConvention.Cdecl)] public static extern int nativeInitMetis(String netaddr, int port, String loca`
  Called by: `.InitRadio()` (`Console/HPSDR/NetworkIO.cs`)
- **`.SetXVTREnable()`** — L18 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int SetXVTREnable(int enable)`
  Sets xvtrenable.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetWBPacketsPerFrame()`** — L21 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetWBPacketsPerFrame(int pps)`
  Sets wbpackets per frame.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetWBUpdateRate()`** — L24 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetWBUpdateRate(int ur)`
  Sets wbupdate rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetWBEnable()`** — L27 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetWBEnable(int adc, int enable)`
  Sets wbenable.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendHighPriority()`** — L30 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SendHighPriority(int enable)`
  Sends high priority.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDDCRate()`** — L33 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDDCRate(int id, int rate)`
  Sets ddcrate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CmdRx()`** — L36 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void CmdRx()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getOOO()`** — L39 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getOOO()`
  Returns ooo.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getSeqInDelta()`** — L42 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getSeqInDelta(bool bInit, int rx, int[] deltas, StringBuilder dateTimeStamp,`
  Returns seq in delta.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clearSnapshots()`** — L45 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void clearSnapshots()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateRNet()`** — L48 — `[DllImport("ChannelMaster.dll", EntryPoint = "create_rnet", CallingConvention = CallingConvention.Cdecl)] public static extern void CreateRNet()`
  Creates rnet.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DestroyRNet()`** — L51 — `[DllImport("ChannelMaster.dll", EntryPoint = "destroy_rnet", CallingConvention = CallingConvention.Cdecl)] public static extern void DestroyRNet()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMetisIPAddr()`** — L54 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int GetMetisIPAddr()`
  Returns metis ipaddr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StartAudioNative()`** — L66 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int StartAudioNative()`
  Starts audio native.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StopAudio()`** — L69 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int StopAudio()`
  Stops audio.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAlexHPFBits()`** — L72 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAlexHPFBits(int bits)`
  Sets alex hpfbits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAlexLPFBits()`** — L75 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAlexLPFBits(int bits, bool isTX, bool isMox)`
  Sets alex lpfbits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisablePA()`** — L78 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void DisablePA(int bit)`
  Disables pa.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTRXrelay()`** — L81 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTRXrelay(int bit)`
  Sets trxrelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAlex2HPFBits()`** — L84 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAlex2HPFBits(int bits)`
  Sets alex2 hpfbits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBPF2Gnd()`** — L87 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetBPF2Gnd(int bits)`
  Sets bpf2 gnd.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAlex2LPFBits()`** — L90 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAlex2LPFBits(int bits)`
  Sets alex2 lpfbits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableApolloFilter()`** — L93 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableApolloFilter(int bits)`
  Enables apollo filter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectApolloFilter()`** — L96 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SelectApolloFilter(int bits)`
  Selects apollo filter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableApolloTuner()`** — L99 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableApolloTuner(int bits)`
  Enables apollo tuner.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableApolloAutoTune()`** — L102 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableApolloAutoTune(int bits)`
  Enables apollo auto tune.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableEClassModulation()`** — L105 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableEClassModulation(int bits)`
  Enables eclass modulation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERPWMmin()`** — L108 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERPWMmin(int min)`
  Sets eerpwmmin.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEERPWMmax()`** — L111 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEERPWMmax(int max)`
  Sets eerpwmmax.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAudioAmpEnable()`** — L114 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAudioAmpEnable(bool enable)`
  Sets audio amp enable.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserADC0()`** — L118 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getUserADC0()`
  Returns user adc0.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserADC1()`** — L121 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getUserADC1()`
  Returns user adc1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserADC2()`** — L124 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getUserADC2()`
  Returns user adc2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserADC3()`** — L127 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getUserADC3()`
  Returns user adc3.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUserOut0()`** — L130 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetUserOut0(int bits)`
  Sets user out0.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUserOut1()`** — L133 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetUserOut1(int bits)`
  Sets user out1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUserOut2()`** — L136 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetUserOut2(int bits)`
  Sets user out2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUserOut3()`** — L139 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetUserOut3(int bits)`
  Sets user out3.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI01()`** — L143 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI01()`
  p1 versions
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI02()`** — L146 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI02()`
  Returns user i02.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI03()`** — L149 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI03()`
  Returns user i03.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI04()`** — L152 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI04()`
  Returns user i04.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI04_p2()`** — L157 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI04_p2()`
  p2 versions
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI05_p2()`** — L160 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI05_p2()`
  Returns user i05 p2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI06_p2()`** — L163 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI06_p2()`
  Returns user i06 p2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI08_p2()`** — L166 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI08_p2()`
  Returns user i08 p2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getUserI02_p2()`** — L169 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool getUserI02_p2()`
  Returns user i02 p2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPureSignal()`** — L173 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPureSignal(int enable)`
  Sets pure signal.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableRx()`** — L176 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableRx(int id, int enable)`
  Enables rx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableRxs()`** — L179 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableRxs(int Rxs)`
  Enables rxs.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableRxSync()`** — L182 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableRxSync(int id, int sync)`
  Enables rx sync.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Protocol1DDCConfig()`** — L185 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void Protocol1DDCConfig(int ddcconfig, int en_diversity, int rxcount, int nddc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nativeGetDotDashPTT()`** — L188 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int nativeGetDotDashPTT()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPttOut()`** — L191 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPttOut(int xmitbit)`
  Sets ptt out.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetVFOfreq()`** — L194 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetVFOfreq(int id, int freq, int tx)`
  Sets vfofreq.
  Called by: `.VFOfreq()` (`Console/HPSDR/NetworkIO.cs`)
- **`.SetWatchdogTimer()`** — L197 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetWatchdogTimer(int bits)`
  Sets watchdog timer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMicXlr()`** — L200 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetMicXlr(int bits)`
  Sets mic xlr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMicBoost()`** — L203 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetMicBoost(int bits)`
  Sets mic boost.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLineIn()`** — L206 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetLineIn(int bits)`
  Sets line in.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLineBoost()`** — L209 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetLineBoost(int bits)`
  Sets line boost.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAlexAtten()`** — L212 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAlexAtten(int bits)`
  Sets alex atten.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADCDither()`** — L215 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADCDither(int bits)`
  Sets adcdither.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADCRandom()`** — L218 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADCRandom(int bits)`
  Sets adcrandom.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTxAttenData()`** — L221 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTxAttenData(int bits)`
  Sets tx atten data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRX1Preamp()`** — L224 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRX1Preamp(int bits)`
  Sets rx1 preamp.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRX2Preamp()`** — L227 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRX2Preamp(int bits)`
  Sets rx2 preamp.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADC1StepAttenData()`** — L230 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADC1StepAttenData(int bits)`
  Sets adc1 step atten data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADC2StepAttenData()`** — L233 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADC2StepAttenData(int bits)`
  Sets adc2 step atten data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADC3StepAttenData()`** — L236 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADC3StepAttenData(int bits)`
  Sets adc3 step atten data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMicTipRing()`** — L239 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetMicTipRing(int bits)`
  Sets mic tip ring.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMicBias()`** — L242 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetMicBias(int bits)`
  Sets mic bias.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMicPTT()`** — L245 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetMicPTT(int bits)`
  Sets mic ptt.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getAndResetADC_Overload()`** — L248 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getAndResetADC_Overload()`
  Returns and reset adc overload.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getADCmaxMagnitude()`** — L251 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern ushort getADCmaxMagnitude(int adc)`
  Returns adcmax magnitude.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getAndResetADCmaxMagnitudeAtOverload()`** — L254 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern ushort getAndResetADCmaxMagnitudeAtOverload(int adc)`
  Returns and reset adcmax magnitude at overload.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getHaveSync()`** — L257 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getHaveSync()`
  Returns have sync.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getExciterPower()`** — L260 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getExciterPower()`
  Returns exciter power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getRevPower()`** — L263 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern float getRevPower()`
  Returns rev power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getFwdPower()`** — L266 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern float getFwdPower()`
  Returns fwd power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getHermesDCVoltage()`** — L269 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int getHermesDCVoltage()`
  Returns hermes dcvoltage.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableCWKeyer()`** — L272 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableCWKeyer(int enable)`
  Enables cwkeyer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetSidetoneRun()`** — L275 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetSidetoneRun(int id, int enable)`
  Sets sidetone run.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetSidetoneVolume()`** — L278 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetSidetoneVolume(int id, double volume)`
  Sets sidetone volume.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWSidetoneVolume()`** — L281 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWSidetoneVolume(int vol)`
  Sets cwsidetone volume.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWPTTDelay()`** — L284 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWPTTDelay(int delay)`
  Sets cwpttdelay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWHangTime()`** — L287 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWHangTime(int hang)`
  Sets cwhang time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWSidetoneFreq()`** — L290 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWSidetoneFreq(int freq)`
  Sets cwsidetone freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWKeyerSpeed()`** — L293 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWKeyerSpeed(int speed)`
  Sets cwkeyer speed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWKeyerMode()`** — L296 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWKeyerMode(int mode)`
  Sets cwkeyer mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWKeyerWeight()`** — L299 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWKeyerWeight(int weight)`
  Sets cwkeyer weight.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWEdgeLength()`** — L302 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWEdgeLength(int edge_length)`
  Sets cwedge length.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableCWKeyerSpacing()`** — L305 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void EnableCWKeyerSpacing(int bits)`
  Enables cwkeyer spacing.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReversePaddles()`** — L308 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void ReversePaddles(int bits)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWDash()`** — L311 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWDash(int bit)`
  Sets cwdash.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWDot()`** — L314 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWDot(int bit)`
  Sets cwdot.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWX()`** — L317 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWX(int bit)`
  Sets cwx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWXPTT()`** — L320 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWXPTT(int bit)`
  Sets cwxptt.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWIambic()`** — L323 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWIambic(int bit)`
  Sets cwiambic.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWBreakIn()`** — L326 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWBreakIn(int bit)`
  Sets cwbreak in.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCWSidetone()`** — L329 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCWSidetone(int bit)`
  Sets cwsidetone.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOCBits()`** — L332 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetOCBits(int b)`
  Sets ocbits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOCExtraBits()`** — L335 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetOCExtraBits(int b)`
  Sets ocextra bits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAntBits()`** — L338 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAntBits(int rx_ant, int trx_ant, int tx_ant, int rx_out, bool tx)`
  Sets ant bits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMKIIBPF()`** — L341 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetMKIIBPF(int bpf)`
  Sets mkiibpf.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRxADC()`** — L344 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetRxADC(int n)`
  Sets rx adc.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADC_cntrl1()`** — L347 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADC_cntrl1(int g)`
  Sets adc cntrl1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetADC_cntrl1()`** — L350 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int GetADC_cntrl1()`
  Returns adc cntrl1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADC_cntrl2()`** — L353 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADC_cntrl2(int g)`
  Sets adc cntrl2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetADC_cntrl2()`** — L356 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int GetADC_cntrl2()`
  Returns adc cntrl2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetADC_cntrl_P1()`** — L359 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetADC_cntrl_P1(int g)`
  Sets adc cntrl p1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetADC_cntrl_P1()`** — L362 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int GetADC_cntrl_P1()`
  Returns adc cntrl p1.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPLLLock()`** — L365 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern bool GetPLLLock()`
  Returns plllock.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ATU_Tune()`** — L368 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void ATU_Tune(int tune)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendStartToMetis()`** — L371 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int SendStartToMetis()`
  Sends start to metis.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendStopToMetis()`** — L374 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int SendStopToMetis()`
  Sends stop to metis.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LRAudioSwap()`** — L377 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void LRAudioSwap(int swap)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCATPort()`** — L380 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCATPort(int port)`
  Sets catport.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTxLatency()`** — L383 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetTxLatency(int txLatency)`
  Sets tx latency.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPttHang()`** — L386 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPttHang(int pttHang)`
  Sets ptt hang.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetResetOnDisconnect()`** — L389 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetResetOnDisconnect(int bit)`
  Sets reset on disconnect.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SwapAudioChannels()`** — L392 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SwapAudioChannels(int bit)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.I2CReadInitiate()`** — L395 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int I2CReadInitiate(int bus, int address, int control)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.I2CWriteInitiate()`** — L398 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int I2CWriteInitiate(int bus, int address, int control, int data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.I2CWrite()`** — L401 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int I2CWrite(int bus, int address, int control, int data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.I2CResponse()`** — L404 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern int I2CResponse(byte[] read_data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetInboundBps()`** — L408 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern double GetInboundBps()`
  bandwdith monitoring
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetOutboundBps()`** — L411 — `[DllImport("ChannelMaster.dll", CallingConvention = CallingConvention.Cdecl)] public static extern double GetOutboundBps()`
  Returns outbound bps.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/NetworkIOImports.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
