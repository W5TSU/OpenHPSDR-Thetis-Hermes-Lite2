# `Console/HPSDR/NetworkIOImports.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** The P/Invoke surface into ChannelMaster's network code — every `extern` for radio I/O (128 imports).

## How this file is used

- Used by (incoming references from other files):
  - `Console/HPSDR/NetworkIO.cs` (calls ×3)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.nativeInitMetis()` (×1), `.SetOutputPowerFactor()` (×1), `.SetVFOfreq()` (×1)

## Outline

### Types

#### `NetworkIO` (type, L7)

- `.SetOutputPowerFactor()` — L9
- `.DeInitMetisSockets()` — L12
- `.nativeInitMetis()` — L15
- `.SetXVTREnable()` — L18
- `.SetWBPacketsPerFrame()` — L21
- `.SetWBUpdateRate()` — L24
- `.SetWBEnable()` — L27
- `.SendHighPriority()` — L30
- `.SetDDCRate()` — L33
- `.CmdRx()` — L36
- `.getOOO()` — L39
- `.getSeqInDelta()` — L42
- `.clearSnapshots()` — L45
- `.CreateRNet()` — L48
- `.DestroyRNet()` — L51
- `.GetMetisIPAddr()` — L54
- `.StartAudioNative()` — L66
- `.StopAudio()` — L69
- `.SetAlexHPFBits()` — L72
- `.SetAlexLPFBits()` — L75
- `.DisablePA()` — L78
- `.SetTRXrelay()` — L81
- `.SetAlex2HPFBits()` — L84
- `.SetBPF2Gnd()` — L87
- `.SetAlex2LPFBits()` — L90
- `.EnableApolloFilter()` — L93
- `.SelectApolloFilter()` — L96
- `.EnableApolloTuner()` — L99
- `.EnableApolloAutoTune()` — L102
- `.EnableEClassModulation()` — L105
- `.SetEERPWMmin()` — L108
- `.SetEERPWMmax()` — L111
- `.SetAudioAmpEnable()` — L114
- `.getUserADC0()` — L118
- `.getUserADC1()` — L121
- `.getUserADC2()` — L124
- `.getUserADC3()` — L127
- `.SetUserOut0()` — L130
- `.SetUserOut1()` — L133
- `.SetUserOut2()` — L136
- `.SetUserOut3()` — L139
- `.getUserI01()` — L143
- `.getUserI02()` — L146
- `.getUserI03()` — L149
- `.getUserI04()` — L152
- `.getUserI04_p2()` — L157
- `.getUserI05_p2()` — L160
- `.getUserI06_p2()` — L163
- `.getUserI08_p2()` — L166
- `.getUserI02_p2()` — L169
- `.SetPureSignal()` — L173
- `.EnableRx()` — L176
- `.EnableRxs()` — L179
- `.EnableRxSync()` — L182
- `.Protocol1DDCConfig()` — L185
- `.nativeGetDotDashPTT()` — L188
- `.SetPttOut()` — L191
- `.SetVFOfreq()` — L194
- `.SetWatchdogTimer()` — L197
- `.SetMicXlr()` — L200
- `.SetMicBoost()` — L203
- `.SetLineIn()` — L206
- `.SetLineBoost()` — L209
- `.SetAlexAtten()` — L212
- `.SetADCDither()` — L215
- `.SetADCRandom()` — L218
- `.SetTxAttenData()` — L221
- `.SetRX1Preamp()` — L224
- `.SetRX2Preamp()` — L227
- `.SetADC1StepAttenData()` — L230
- `.SetADC2StepAttenData()` — L233
- `.SetADC3StepAttenData()` — L236
- `.SetMicTipRing()` — L239
- `.SetMicBias()` — L242
- `.SetMicPTT()` — L245
- `.getAndResetADC_Overload()` — L248
- `.getADCmaxMagnitude()` — L251
- `.getAndResetADCmaxMagnitudeAtOverload()` — L254
- `.getHaveSync()` — L257
- `.getExciterPower()` — L260
- `.getRevPower()` — L263
- `.getFwdPower()` — L266
- `.getHermesDCVoltage()` — L269
- `.EnableCWKeyer()` — L272
- `.SetSidetoneRun()` — L275
- `.SetSidetoneVolume()` — L278
- `.SetCWSidetoneVolume()` — L281
- `.SetCWPTTDelay()` — L284
- `.SetCWHangTime()` — L287
- `.SetCWSidetoneFreq()` — L290
- `.SetCWKeyerSpeed()` — L293
- `.SetCWKeyerMode()` — L296
- `.SetCWKeyerWeight()` — L299
- `.SetCWEdgeLength()` — L302
- `.EnableCWKeyerSpacing()` — L305
- `.ReversePaddles()` — L308
- `.SetCWDash()` — L311
- `.SetCWDot()` — L314
- `.SetCWX()` — L317
- `.SetCWXPTT()` — L320
- `.SetCWIambic()` — L323
- `.SetCWBreakIn()` — L326
- `.SetCWSidetone()` — L329
- `.SetOCBits()` — L332
- `.SetOCExtraBits()` — L335
- `.SetAntBits()` — L338
- `.SetMKIIBPF()` — L341
- `.SetRxADC()` — L344
- `.SetADC_cntrl1()` — L347
- `.GetADC_cntrl1()` — L350
- `.SetADC_cntrl2()` — L353
- `.GetADC_cntrl2()` — L356
- `.SetADC_cntrl_P1()` — L359
- `.GetADC_cntrl_P1()` — L362
- `.GetPLLLock()` — L365
- `.ATU_Tune()` — L368
- `.SendStartToMetis()` — L371
- `.SendStopToMetis()` — L374
- `.LRAudioSwap()` — L377
- `.SetCATPort()` — L380
- `.SetTxLatency()` — L383
- `.SetPttHang()` — L386
- `.SetResetOnDisconnect()` — L389
- `.SwapAudioChannels()` — L392
- `.I2CReadInitiate()` — L395
- `.I2CWriteInitiate()` — L398
- `.I2CWrite()` — L401
- `.I2CResponse()` — L404
- `.GetInboundBps()` — L408
- `.GetOutboundBps()` — L411

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/NetworkIOImports.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
