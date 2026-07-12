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

### Types

#### `cmaster` (type, L61)

- `.SetRadioStructure()` — L68
- `.CreateRadio()` — L82
- `.DestroyRadio()` — L85
- `.SetCMDefaultRates()` — L88
- `.SetXcmInrate()` — L91
- `.SetXmtrChannelOutrate()` — L94
- `.GetBuffSize()` — L97
- `.inid()` — L100
- `.Inbound()` — L103
- `.chid()` — L106
- `.GetInputRate()` — L109
- `.GetChannelOutputRate()` — L112
- `.GetCMAstate()` — L115
- `.SendpOutboundTCIRxIQ()` — L119
- `.SendpOutboundTCIRxAudio()` — L122
- `.SendpInboundTCITxAudio()` — L125
- `.SetRXTCIRun()` — L128
- `.SetTXTCIAudioRun()` — L131
- `.SetTCIRxAudioMox()` — L134
- `.SetTCIRxAudioMon()` — L137
- `.SetTCIRxAudioMonVol()` — L140
- `.LoadRouterAll()` — L145
- `.LoadRouterControlBit()` — L149
- `.SetTopPan3Run()` — L154
- `.SetRunPanadapter()` — L157
- `.SetPSTxIdx()` — L162
- `.SetPSRxIdx()` — L165
- `.AllocAnalyzer()` — L169
- `.FreeAnalyzer()` — L171
- `.RunAnalyzer()` — L174
- `.SetTXAVoxThresh()` — L179
- `.GetDEXPPeakSignal()` — L182
- `.SetDEXPRun()` — L185
- `.SetDEXPDetectorTau()` — L188
- `.SetDEXPAttackTime()` — L191
- `.SetDEXPReleaseTime()` — L194
- `.SetDEXPHoldTime()` — L197
- `.SetDEXPExpansionRatio()` — L200
- `.SetDEXPHysteresisRatio()` — L203
- `.SetDEXPAttackThreshold()` — L206
- `.SetDEXPLowCut()` — L209
- `.SetDEXPHighCut()` — L212
- `.SetDEXPRunSideChannelFilter()` — L215
- `.SetDEXPRunVox()` — L218
- `.SetDEXPRunAudioDelay()` — L221
- `.SetDEXPAudioDelay()` — L224
- `.SetAntiVOXRun()` — L227
- `.SetAntiVOXGain()` — L230
- `.SetAntiVOXDetectorTau()` — L233
- `.SetAntiVOXSourceStates()` — L236
- `.SetAntiVOXSourceWhat()` — L239
- `.SetSiphonInsize()` — L244
- `.GetaSipF1EXT()` — L247
- `.SetAAudioMixWhat()` — L252
- `.SetAAudioMixState()` — L255
- `.SetAAudioMixStates()` — L258
- `.SetAAudioMixVolume()` — L261
- `.SetAAudioMixVol()` — L264
- `.SetTXVAC()` — L269
- `.SetTXFixedGainRun()` — L274
- `.SetTXFixedGain()` — L277
- `.SetEERRun()` — L282
- `.SetEERAMIQ()` — L285
- `.SetEERMgain()` — L288
- `.SetEERPgain()` — L291
- `.SetEERRunDelays()` — L294
- `.SetEERMdelay()` — L297
- `.SetEERPdelay()` — L300
- `.SetRCVRANBRun()` — L305
- `.SetRCVRANBTau()` — L308
- `.SetRCVRANBHangtime()` — L311
- `.SetRCVRANBAdvtime()` — L314
- `.SetRCVRANBBacktau()` — L317
- `.SetRCVRANBThreshold()` — L320
- `.SetRCVRNOBRun()` — L325
- `.SetRCVRNOBMode()` — L328
- `.SetRCVRNOBTau()` — L331
- `.SetRCVRNOBHangtime()` — L334
- `.SetRCVRNOBAdvtime()` — L337
- `.SetRCVRNOBBacktau()` — L340
- `.SetRCVRNOBThreshold()` — L343
- `.GetAndResetAmpProtect()` — L348
- `.SetAmpProtectRun()` — L351
- `.SetADCSupply()` — L354
- `.getLEDs()` — L359
- `.GetCMVersion()` — L363
- `.GetCMasioVersion()` — L366
- `.CMCreateCMaster()` — L500
- `.CMLoadRouterAll()` — L570
- `.CMSetAntiVoxSourceWhat()` — L934
- `.CMSetAudioVolume()` — L967
- `.CMSetFRXNBRun()` — L972
- `.CMSetFRXNB2Run()` — L989
- `.CMSetSRXWavePlayRun()` — L1006
- `.CMSetSRXWaveRecordRun()` — L1025
- `.CMSetEERRun()` — L1046
- `.CMSetTXAVoxRun()` — L1061
- `.CMSetTXAVoxThresh()` — L1076
- `.CMSetTXAPanelGain1()` — L1083
- `.CMSetScopeRun()` — L1125
- `.CMSetTXOutputLevelRun()` — L1131
- `.CMSetTXOutputLevel()` — L1137
- `.SendCallbacks()` — L1145
- `.ensureTCIStreamThreads()` — L1176
- `.StopTCIStreamThreads()` — L1198
- `.TCIRxThreadProc()` — L1208
- `.serviceTCIRxStreams()` — L1225
- `.TCITxThreadProc()` — L1257
- `.serviceTCITxProtocol()` — L1273
- `.resetTCITxState()` — L1379
- `.queueTCITxAudio()` — L1404
- `.resampleTCITxSamples()` — L1453
- `.destroyTCIIQResampler()` — L1500
- `.resampleTCIIQSamples()` — L1521
- `.rentTCIFloatBuffer()` — L1607
- `.returnTCIFloatBuffer()` — L1621
- `.rentTCIIQBlock()` — L1639
- `.returnTCIIQBlock()` — L1650
- `.rentTCIAudioBlock()` — L1668
- `.returnTCIAudioBlock()` — L1679
- `.enqueueTCIIQ()` — L1699
- `.enqueueTCIAudio()` — L1713
- `.tryDequeueTCIIQ()` — L1727
- `.tryDequeueTCIAudio()` — L1742
- `.OnTCIRxIQOutSamples()` — L1758
- `.OnTCIRxAudioOutSamples()` — L1796
- `.OnTCITxAudioInSamples()` — L1817
- `.SendCBPushVox()` — L1865
- `.create_wb()` — L1908
- `.Getwb()` — L1919
- `.Hidewb()` — L1930
- `.Closewb()` — L1936
- `.Savewb()` — L1942

#### `TCIIQBlock` (type, L444)

_No extracted members._

#### `TCIAudioBlock` (type, L452)

_No extracted members._

#### `VOX` (type, L1954)

- `.PushVox()` — L1956

#### `WaveThing` (type, L1965)

- `.SendCBCreateWPlay()` — L1976
- `.SetWavePlayerRun()` — L1979
- `.SendCBCreateWRecord()` — L1985
- `.SetWaveRecorderRun()` — L1988
- `.initWaves()` — L1991
- `.SendCBWavePlayer()` — L2006
- `.createWavePlayer()` — L2018
- `.SendCBWaveRecorder()` — L2037
- `.createWaveRecorder()` — L2045
- `.UpdateMox()` — L2052

#### `PlayWave` (type, L2064)

- `.wplay()` — L2092
- `.swizzle()` — L2118

#### `RecordWave` (type, L2133)

- `.wrecord()` — L2177
- `.deswizzle()` — L2265

#### `Scope` (type, L2282)

- `.SendCBCreateScope()` — L2287
- `.SendCBScope()` — L2292
- `.SetScopeRun()` — L2295
- `.createScope()` — L2302
- `.initScope()` — L2311

#### `DoScope` (type, L2317)

- `.xscope()` — L2338
- `.deswizzle()` — L2365

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/cmaster.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
