# `ChannelMaster/netInterface.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** HPSDR Protocol-1 UDP implementation: socket setup, packet build/parse, EP2/EP4/EP6 endpoint handling, sequence tracking.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/networkproto1.c` (calls ×6)
  - `ChannelMaster/network.c` (calls ×4)
- Uses (outgoing references to other files):
  - `ChannelMaster/network.c` (calls ×64)
  - `ChannelMaster/sidetone.c` (calls ×5)
  - `ChannelMaster/cmasio.c` (calls ×2)
  - `ChannelMaster/cmaster.c` (calls ×2)
  - `ChannelMaster/obbuffs.c` (calls ×2)
  - `ChannelMaster/bandwidth_monitor.c` (calls ×1)
  - `ChannelMaster/network.h` (imports ×1)
  - `ChannelMaster/obbuffs.h` (imports ×1)
  - `ChannelMaster/networkproto1.c` (calls ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `PrintTimeHack()` (×4), `PeakFwdPower()` (×3), `PeakRevPower()` (×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`StartAudioNative()`** — L35 — `PORT int StartAudioNative()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`StopAudio()`** — L100 — `PORT void StopAudio()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`nativeGetDotDashPTT()`** — L123 — `PORT int nativeGetDotDashPTT()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getAndResetADC_Overload()`** — L131 — `PORT int getAndResetADC_Overload()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getAndResetADCmaxMagnitudeAtOverload()`** — L148 — `PORT uint16_t getAndResetADCmaxMagnitudeAtOverload(int adc)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getADCmaxMagnitude()`** — L162 — `PORT uint16_t getADCmaxMagnitude(int adc)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getOOO()`** — L172 — `PORT int getOOO()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getSeqInDelta()`** — L211 — `PORT int getSeqInDelta(int nInit, int rx, int deltas[], char* dateTimeStamp, int *received_seqnum, int *last_seqnum)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`GetPLLLock()`** — L238 — `PORT int GetPLLLock()`
  Returns plllock — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI01()`** — L248 — `PORT int getUserI01()`
  NOTE: these 4 user get fuctions are named for P1 //MW0LGE_22b Bit [0] - User I/O (IO1) 1 = active, 0 = inactive Bit [1] - User I/O (IO2) 1 = active, 0 = inactive Bit [2] - User I/O (IO3) 1 = active, 0 = inactive Bit [3] - User I/O (IO4) 1 = active, 0 = inactive
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI02()`** — L254 — `PORT int getUserI02()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI03()`** — L260 — `PORT int getUserI03()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI04()`** — L266 — `PORT int getUserI04()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI04_p2()`** — L279 — `PORT int getUserI04_p2()`
  NOTE: these 5 user get functions are named for P2 //MW0LGE_22b Bit [0] - User I/O (IO4) 1 = active, 0 = inactive Bit [1] - User I/O (IO5) 1 = active, 0 = inactive Bit [2] - User I/O (IO6) 1 = active, 0 = inactive Bit [3] - User I/O (IO8) 1 = active, 0 = inactive Bit [4] - User I/O (IO2) 1 =…
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI05_p2()`** — L285 — `PORT int getUserI05_p2()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI06_p2()`** — L291 — `PORT int getUserI06_p2()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI08_p2()`** — L297 — `PORT int getUserI08_p2()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserI02_p2()`** — L302 — `PORT int getUserI02_p2()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getExciterPower()`** — L309 — `PORT int getExciterPower()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getFwdPower()`** — L316 — `PORT float getFwdPower()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getRevPower()`** — L323 — `PORT float getRevPower()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserADC0()`** — L330 — `PORT int getUserADC0()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserADC1()`** — L337 — `PORT int getUserADC1()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserADC2()`** — L344 — `PORT int getUserADC2()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getUserADC3()`** — L351 — `PORT int getUserADC3()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getHermesDCVoltage()`** — L358 — `PORT int getHermesDCVoltage()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetPttOut()`** — L365 — `PORT void SetPttOut(int xmit)`
  Sets ptt out — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetTRXrelay()`** — L377 — `PORT void SetTRXrelay(int bit)`
  Sets trxrelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`EnableEClassModulation()`** — L391 — `PORT void EnableEClassModulation(int bit)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetOCBits()`** — L402 — `PORT void SetOCBits(int b)`
  Sets ocbits — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetOCExtraBits()`** — L413 — `PORT void SetOCExtraBits(int b)`
  Sets ocextra bits — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAlexAtten()`** — L424 — `PORT void SetAlexAtten(int bits)`
  Sets alex atten — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADCDither()`** — L438 — `PORT void SetADCDither(int bits)`
  Sets adcdither — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADCRandom()`** — L450 — `PORT void SetADCRandom(int bits)`
  Sets adcrandom — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAntBits()`** — L462 — `PORT void SetAntBits(int rx_only_ant, int trx_ant, int tx_ant, int rx_out, char tx)`
  Sets ant bits — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetVFOfreq()`** — L505 — `PORT void SetVFOfreq(int id, int freq, int tx)`
  Sets vfofreq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetOutputPowerFactor()`** — L528 — `PORT void SetOutputPowerFactor(int u)`
  Sets output power factor — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetMicBoost()`** — L539 — `PORT void SetMicBoost(int bits)`
  Sets mic boost — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetMicXlr()`** — L550 — `PORT void SetMicXlr(int bits)`
  Sets mic xlr — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`, `Console/console.cs`.
- **`SetLineIn()`** — L561 — `PORT void SetLineIn(int bits)`
  Sets line in — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`EnableApolloFilter()`** — L572 — `PORT void EnableApolloFilter(int bits)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`EnableApolloTuner()`** — L581 — `PORT void EnableApolloTuner(int bits)`
  Called by: `DisablePA()` (same file)
- **`EnableApolloAutoTune()`** — L590 — `PORT void EnableApolloAutoTune(int bits)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SelectApolloFilter()`** — L599 — `PORT void SelectApolloFilter(int bits)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAlexHPFBits()`** — L608 — `PORT void SetAlexHPFBits(int bits)`
  Sets alex hpfbits — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`DisablePA()`** — L627 — `PORT void DisablePA(int bit)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAlex2HPFBits()`** — L641 — `PORT void SetAlex2HPFBits(int bits)`
  Sets alex2 hpfbits — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetBPF2Gnd()`** — L660 — `PORT void SetBPF2Gnd(int bits)`
  Sets bpf2 gnd — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAlex3HPFBits()`** — L671 — `PORT void SetAlex3HPFBits(int bits)`
  Sets alex3 hpfbits — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAlex4HPFBits()`** — L677 — `PORT void SetAlex4HPFBits(int bits)`
  Sets alex4 hpfbits — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAlexLPFBits()`** — L688 — `PORT void SetAlexLPFBits(int bits, bool isTX, bool isMox)`
  LPF bits can be used in older radioas as part of RX filtering too. Change to protocol 2 from 4.3 onwards: TX settings are encoded in the Alex1 word to remain comparible with older hardware, the logic will be: if MOX, write settings to alex0 and alex1 if not MOX, write to alex1 if a TX setting…
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAlex2LPFBits()`** — L735 — `PORT void SetAlex2LPFBits(int bits)`
  Sets alex2 lpfbits — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAlex3LPFBits()`** — L741 — `PORT void SetAlex3LPFBits(int bits)`
  Sets alex3 lpfbits — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAlex4LPFBits()`** — L747 — `PORT void SetAlex4LPFBits(int bits)`
  Sets alex4 lpfbits — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRX1Preamp()`** — L753 — `PORT void SetRX1Preamp(int bits)`
  Sets rx1 preamp — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetRX2Preamp()`** — L765 — `PORT void SetRX2Preamp(int bits)`
  Sets rx2 preamp — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetMicTipRing()`** — L776 — `PORT void SetMicTipRing(int bits)`
  Sets mic tip ring — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetMicBias()`** — L787 — `PORT void SetMicBias(int bits)`
  Sets mic bias — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetMicPTT()`** — L798 — `PORT void SetMicPTT(int bits)`
  Sets mic ptt — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetLineBoost()`** — L809 — `PORT void SetLineBoost(int bits)`
  Sets line boost — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetResetOnDisconnect()`** — L820 — `PORT void SetResetOnDisconnect(int bit)`
  Sets reset on disconnect — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SwapAudioChannels()`** — L829 — `PORT void SwapAudioChannels(int swap)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetPureSignal()`** — L839 — `PORT void SetPureSignal(int bit)`
  Sets pure signal — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetUserOut0()`** — L850 — `PORT void SetUserOut0(int out)`
  Sets user out0 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetUserOut1()`** — L856 — `PORT void SetUserOut1(int out)`
  Sets user out1 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetUserOut2()`** — L862 — `PORT void SetUserOut2(int out)`
  Sets user out2 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetUserOut3()`** — L868 — `PORT void SetUserOut3(int out)`
  Sets user out3 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADC1StepAttenData()`** — L874 — `PORT void SetADC1StepAttenData(int data)`
  Sets adc1 step atten data — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADC2StepAttenData()`** — L885 — `PORT void SetADC2StepAttenData(int data)`
  Sets adc2 step atten data — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADC3StepAttenData()`** — L896 — `PORT void SetADC3StepAttenData(int data)`
  Sets adc3 step atten data — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`ReversePaddles()`** — L907 — `PORT void ReversePaddles(int bits)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWKeyerSpeed()`** — L918 — `PORT void SetCWKeyerSpeed(int speed)`
  Sets cwkeyer speed — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`, `Console/TCIServer.cs`.
- **`SetCWKeyerMode()`** — L930 — `PORT void SetCWKeyerMode(int mode)`
  Sets cwkeyer mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWKeyerWeight()`** — L941 — `PORT void SetCWKeyerWeight(int weight)`
  Sets cwkeyer weight — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`EnableCWKeyerSpacing()`** — L952 — `PORT void EnableCWKeyerSpacing(int bits)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWEdgeLength()`** — L963 — `PORT void SetCWEdgeLength(int edge_length)`
  Sets cwedge length — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADC_cntrl1()`** — L974 — `PORT void SetADC_cntrl1(int bits)`
  Sets adc cntrl1 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`GetADC_cntrl1()`** — L990 — `PORT int GetADC_cntrl1()`
  Returns adc cntrl1 — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADC_cntrl2()`** — L996 — `PORT void SetADC_cntrl2(int bits)`
  Sets adc cntrl2 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`GetADC_cntrl2()`** — L1011 — `PORT int GetADC_cntrl2()`
  Returns adc cntrl2 — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetADC_cntrl_P1()`** — L1018 — `PORT void SetADC_cntrl_P1(int bits)`
  Sets adc cntrl p1 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`GetADC_cntrl_P1()`** — L1024 — `PORT int GetADC_cntrl_P1()`
  Returns adc cntrl p1 — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetTxAttenData()`** — L1031 — `PORT void SetTxAttenData(int bits)`
  Sets tx atten data — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`EnableCWKeyer()`** — L1047 — `PORT void EnableCWKeyer(int enable)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWSidetoneVolume()`** — L1059 — `PORT void SetCWSidetoneVolume(int vol)`
  Sets cwsidetone volume — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`, `Console/console.cs`.
- **`SetCWPTTDelay()`** — L1071 — `PORT void SetCWPTTDelay(int delay)`
  Sets cwpttdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWHangTime()`** — L1082 — `PORT void SetCWHangTime(int time)`
  Sets cwhang time — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWSidetoneFreq()`** — L1093 — `PORT void SetCWSidetoneFreq(int freq)`
  Sets cwsidetone freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetEERPWMmin()`** — L1105 — `PORT void SetEERPWMmin(int min)`
  Sets eerpwmmin — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetEERPWMmax()`** — L1116 — `PORT void SetEERPWMmax(int max)`
  Sets eerpwmmax — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetAudioAmpEnable()`** — L1128 — `PORT void SetAudioAmpEnable(int enable)`
  MW0LGE_22b
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWSidetone()`** — L1142 — `PORT void SetCWSidetone(int enable)`
  misc functions
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWIambic()`** — L1153 — `PORT void SetCWIambic(int enable)`
  Sets cwiambic — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWBreakIn()`** — L1164 — `PORT void SetCWBreakIn(int enable)`
  Sets cwbreak in — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWDash()`** — L1175 — `PORT void SetCWDash(int bit)`
  Sets cwdash — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWDot()`** — L1186 — `PORT void SetCWDot(int bit)`
  Sets cwdot — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWX()`** — L1197 — `PORT void SetCWX(int bit)`
  Sets cwx — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetCWXPTT()`** — L1219 — `PORT void SetCWXPTT(int bit)`
  Sets cwxptt — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getHaveSync()`** — L1230 — `PORT int getHaveSync()`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getControlByteIn()`** — L1236 — `PORT int getControlByteIn(int n)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`EnableRx()`** — L1246 — `PORT void EnableRx(int id, int enable)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`EnableRxs()`** — L1257 — `PORT void EnableRxs(int rxs)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`EnableRxSync()`** — L1286 — `PORT void EnableRxSync(int id, int sync)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`Protocol1DDCConfig()`** — L1295 — `PORT void Protocol1DDCConfig(int ddcconfig, int en_diversity, int rxcount, int inddc)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetDDCRate()`** — L1304 — `PORT void SetDDCRate(int id, int rate)`
  Sets ddcrate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetRxADC()`** — L1361 — `PORT void SetRxADC(int n)`
  Sets rx adc — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetWBPacketsPerFrame()`** — L1373 — `PORT void SetWBPacketsPerFrame(int ppf)`
  wideband data display
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetWBUpdateRate()`** — L1381 — `PORT void SetWBUpdateRate(int ur)`
  Sets wbupdate rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetWBEnable()`** — L1389 — `PORT void SetWBEnable(int adc, int enable)`
  Sets wbenable — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SendHighPriority()`** — L1398 — `PORT void SendHighPriority(int send)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetWatchdogTimer()`** — L1410 — `PORT void SetWatchdogTimer(int enable)`
  Sets watchdog timer — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetMKIIBPF()`** — L1421 — `PORT void SetMKIIBPF(int bpf)`
  Sets mkiibpf — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetXVTREnable()`** — L1427 — `PORT void SetXVTREnable(int enable)`
  Sets xvtrenable — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`ATU_Tune()`** — L1438 — `PORT void ATU_Tune(int tune)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`getLEDs()`** — L1449 — `PORT int getLEDs()`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`LRAudioSwap()`** — L1455 — `PORT void LRAudioSwap (int swap)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetTxLatency()`** — L1462 — `PORT void SetTxLatency (int txLatency)`
  Sets tx latency — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SetPttHang()`** — L1468 — `PORT void SetPttHang (int pttHang)`
  Sets ptt hang — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`I2CReadInitiate()`** — L1474 — `PORT int I2CReadInitiate(int bus, int address, int control)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`I2CWriteInitiate()`** — L1505 — `PORT int I2CWriteInitiate(int bus, int address, int control, int data)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`I2CWrite()`** — L1539 — `PORT int I2CWrite(int bus, int address, int control, int data)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`I2CResponse()`** — L1570 — `PORT int I2CResponse(unsigned char* read_data)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`create_rnet()`** — L1594 — `PORT void create_rnet()`
  Constructor for the `rnet` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`clearSnapshots()`** — L1769 — `PORT void clearSnapshots()`
  Called by: `destroy_rnet()` (same file)
- **`destroy_rnet()`** — L1788 — `PORT void destroy_rnet()`
  Destroys the `rnet` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`PrintTimeHack()`** — L1817 — `void PrintTimeHack()`
  Called by: `StopReadThread()` (`ChannelMaster/network.c`), `ReadThreadMainLoop()` (`ChannelMaster/network.c`), `MetisReadThreadMainLoop()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop_HL2()` (`ChannelMaster/networkproto1.c`)
- **`PeakFwdPower()`** — L1824 — `void PeakFwdPower(float fwd)`
  Called by: `ReadThreadMainLoop()` (`ChannelMaster/network.c`), `MetisReadThreadMainLoop()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop_HL2()` (`ChannelMaster/networkproto1.c`)
- **`PeakRevPower()`** — L1832 — `void PeakRevPower(float rev)`
  Called by: `ReadThreadMainLoop()` (`ChannelMaster/network.c`), `MetisReadThreadMainLoop()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop_HL2()` (`ChannelMaster/networkproto1.c`)
- **`UpdateRadioProtocolSampleSize()`** — L1840 — `void UpdateRadioProtocolSampleSize()`
  Called by: `StartAudioNative()` (same file)
- **`SetCATPort()`** — L1865 — `PORT void SetCATPort(int port)`
  set CAT over TCP port for remote communication with protocol client apps
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/netInterface.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
