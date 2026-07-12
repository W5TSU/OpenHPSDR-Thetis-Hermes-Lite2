# `Console/TCIServer.cs`

**Functional area:** [10. CAT control and external program interfaces](../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** TCI WebSocket server (protocol used by SDC, LogHX, etc.): exposes VFOs, modes, spots, and audio to TCI clients.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×12, references ×1)
  - `Console/cmaster.cs` (calls ×6, references ×3)
  - `Console/setup.cs` (references ×2)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×37)
  - `Console/MeterManager.cs` (references ×5, calls ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×4)
  - `Console/SpotManager2.cs` (calls ×4)
  - `Console/frmLog.Designer.cs` (references ×1)
- Most-referenced symbols from other files: `.SensorRequiresUpdate()` (×3), `.StopServer()` (×2), `.MinimumRequiredRxSensorInterval()` (×2), `.MinimumRequiredTxSensorInterval()` (×2), `.StartServer()` (×1), `.ShowLog()` (×1), `.CloseLog()` (×1), `.PublishIQSamples()` (×1)

## Outline

### Types

#### `TCICWSpotForce` (type, L329)

_No extracted members._

#### `TCITxStereoInputMode` (type, L336)

_No extracted members._

#### `TCIStreamType` (type, L343)

_No extracted members._

#### `TCISampleType` (type, L352)

_No extracted members._

#### `TCIQueuedTxAudio` (type, L361)

_No extracted members._

#### `TCIPendingFloatBuffer` (type, L371)

- `.Enqueue()` — L387
- `.CopyTo()` — L397
- `.Peek()` — L405
- `.Advance()` — L410
- `.ensureCapacity()` — L436

#### `clsTCISensorManager` (type, L462)

- `.clampIntervalMs()` — L500
- `.ConfigureRxSensors()` — L527
- `.ConfigureTxSensors()` — L544
- `.RequiresRxChannelUpdate()` — L554
- `.RequiresTxUpdate()` — L566
- `.SensorRequiresUpdate()` — L574
- `.SetRxChannelReading()` — L599
- `.SetTxReadings()` — L613
- `.TryGetRxChannelReadingForSend()` — L625
- `.ConsumeRxChannelReading()` — L645
- `.TryGetTxReadingsForSend()` — L656
- `.ConsumeTxReadings()` — L675

#### `clsRxReadingState` (type, L464)

_No extracted members._

#### `clsTxReadingState` (type, L472)

_No extracted members._

#### `TCPIPtciSocketListener` (type, L684)

- `.ClickedOnSpot()` — L826
- `.ThetisFocusChange()` — L837
- `.RX2EnabledChange()` — L842
- `.HWSampleRateChange()` — L848
- `.RequiresRxSensorUpdate()` — L875
- `.SensorRequiresUpdate()` — L880
- `.RequiresTxSensorUpdate()` — L885
- `.MeterReadingsChanged()` — L890
- `.MinimumRequiredRxSensorInterval()` — L915
- `.MinimumRequiredTxSensorInterval()` — L920
- `.getPublishedIQSampleRate()` — L925
- `.getPublishedIQSampleRateLocked()` — L933
- `.destroyRxAudioResamplerState()` — L945
- `.clearRxAudioStateForReceiver()` — L966
- `.clearRxAudioStreamState()` — L981
- `.resampleRxAudioSamples()` — L995
- `.DrivePowerChange()` — L1088
- `.TuneChange()` — L1096
- `.SplitChange()` — L1101
- `.MuteChanged()` — L1107
- `.AnfChanged()` — L1113
- `.RxAfGainChanged()` — L1118
- `.CTUNChanged()` — L1125
- `.VFOSyncChanged()` — L1130
- `.FMDeviationChanged()` — L1135
- `.AGCModeChanged()` — L1140
- `.AGCAutoChanged()` — L1145
- `.TXProfileChanged()` — L1150
- `.TXProfilesChanged()` — L1155
- `.CalibrationChanged()` — L1160
- `.MONChanged()` — L1178
- `.MONVolumeChanged()` — L1183
- `.VolumeChanged()` — L1188
- `.BalanceChanged()` — L1193
- `.RxStepAttChanged()` — L1200
- `.RxPreampAttChanged()` — L1205
- `.RxStepAttEnabledChanged()` — L1212
- `.AGCGainChanged()` — L1217
- `.RITChanged()` — L1222
- `.XITChanged()` — L1228
- `.RITValueChanged()` — L1234
- `.XITValueChanged()` — L1240
- `.CwMacrosSpeedChanged()` — L1246
- `.CwMacrosDelayChanged()` — L1251
- `.CwKeyerSpeedChanged()` — L1256
- `.CwMacrosEmpty()` — L1261
- `.CwCallsignSent()` — L1266
- `.NBChanged()` — L1271
- `.NRChanged()` — L1278
- `.BinChanged()` — L1285
- `.LockChanged()` — L1290
- `.VFOLocksChanged()` — L1295
- `.SqlChanged()` — L1300
- `.SqlLevelChanged()` — L1305
- `.ApfChanged()` — L1310
- `.NfChanged()` — L1315
- `.DiglOffsetChanged()` — L1321
- `.DiguOffsetChanged()` — L1326
- `.TXFrequencyChanged()` — L1331
- `.limitList()` — L1336
- `.VFOdata()` — L1348
- `.vfoFrequencyChange()` — L1417
- `.centreFrequencyChange()` — L1426
- `.txFrequencyChange()` — L1435
- `.MoxChange()` — L1444
- `.ModeChange()` — L1473
- `.BandChange()` — L1478
- `.FilterChange()` — L1485
- `.FilterEdgesChange()` — L1490
- `.TXFilterBandChanged()` — L1495
- `.PowerChange()` — L1500
- `.StartSocketListener()` — L1507
- `.getFrameFromString()` — L1538
- `.GetFrameFromBytes()` — L1597
- `.getCoalescedTextFrameKey()` — L1638
- `.hasPendingOutboundFramesLocked()` — L1680
- `.tryDequeueNextOutboundFrameLocked()` — L1688
- `.clearOutboundFrames()` — L1724
- `.flushOutboundFrames()` — L1737
- `.enqueueOutboundFrame()` — L1752
- `.SendThreadProc()` — L1791
- `.abortSocketTransport()` — L1836
- `.isSocketReadTimeout()` — L1865
- `.upgradeToWebSocket()` — L1876
- `.sendStart()` — L1911
- `.sendStop()` — L1915
- `.sendSplit()` — L1919
- `.sendRITEnable()` — L1924
- `.sendXITEnable()` — L1929
- `.sendRITOffset()` — L1934
- `.sendXITOffset()` — L1939
- `.sendRxBinEnable()` — L1944
- `.sendRxApfEnable()` — L1949
- `.sendRxNfEnable()` — L1954
- `.sendLock()` — L1959
- `.sendVFOLock()` — L1964
- `.sendSqlEnable()` — L1969
- `.sendSqlLevel()` — L1974
- `.sendCwMacrosSpeed()` — L1979
- `.sendCwMacrosDelay()` — L1983
- `.sendCwKeyerSpeed()` — L1987
- `.sendCwMacrosEmpty()` — L1991
- `.sendCallsignSend()` — L1995
- `.tryGetVFOLockState()` — L1999
- `.trySetVFOLockState()` — L2034
- `.sendAllVFOLocks()` — L2070
- `.sendDiglOffset()` — L2089
- `.sendDiguOffset()` — L2094
- `.sendVFO()` — L2099
- `.sendIF()` — L2134
- `.sendMOX()` — L2159
- `.sendAudioStartStop()` — L2170
- `.sendMode()` — L2174
- `.sendMute()` — L2196
- `.sendMuteRX()` — L2201
- `.sendMONEnable()` — L2206
- `.sendVolume()` — L2211
- `.sendMONVolume()` — L2218
- `.sendRxBalance()` — L2225
- `.sendRxStepAttEx()` — L2230
- `.sendRxPreampAttEx()` — L2235
- `.sendRxStepAttEnabledEx()` — L2240
- `.sendVFOSyncEx()` — L2245
- `.sendFMDeviationEx()` — L2250
- `.sendAgcAutoEx()` — L2255
- `.agcModeToTciMode()` — L2260
- `.tciModeToAgcMode()` — L2280
- `.sendAgcMode()` — L2304
- `.sendAgcGain()` — L2309
- `.sendTXFrequencyChanged()` — L2314
- `.sendTunePower()` — L2328
- `.sendDrivePower()` — L2335
- `.sendTune()` — L2342
- `.sendRXEnable()` — L2347
- `.sendTXEnable()` — L2352
- `.sendVFOLimits()` — L2357
- `.sendAppFocus()` — L2362
- `.sendIFLimits()` — L2367
- `.sendClickedOnSpot()` — L2372
- `.sendClickedOnSpotRX()` — L2377
- `.sendRxSensors()` — L2382
- `.sendRxChannelSensors()` — L2386
- `.sendTxSensors()` — L2391
- `.sendDDS()` — L2402
- `.sendFilterBand()` — L2417
- `.normalizeTXFilterBandForSet()` — L2422
- `.normalizeTXFilterBandForSend()` — L2437
- `.sendTXFilterBandEx()` — L2443
- `.sendStartStop()` — L2449
- `.preampModeToAttenuation()` — L2457
- `.sendInitialRadioState()` — L2486
- `.sendInitialisationData()` — L2662
- `.setRxSensorsEnabled()` — L2704
- `.setTxSensorsEnabled()` — L2719
- `.RxSensorsTimerCallback()` — L2734
- `.TxSensorsTimerCallback()` — L2769
- `.findEndOfHeader()` — L2781
- `.SocketListenerThreadStart()` — L2798
- `.notifyServerDisconnected()` — L2915
- `.sendPingFrame()` — L2926
- `.sendPongFrame()` — L2944
- `.sendTextFrame()` — L2962
- `.sendBinaryFrame()` — L2981
- `.sendCloseFrame()` — L3000
- `.StopSocketListener()` — L3014
- `.IsMarkedForDeletion()` — L3118
- `.IsDisconnected()` — L3122
- `.GetFrameLength()` — L3126
- `.ParseReceiveBuffer()` — L3156
- `.handleSetInFocus()` — L3223
- `.handleStart()` — L3227
- `.handleStop()` — L3232
- `.handleSpotClear()` — L3237
- `.handleSplitEnableMessage()` — L3241
- `.handleRITEnableMessage()` — L3278
- `.handleXITEnableMessage()` — L3294
- `.handleRITOffsetMessage()` — L3310
- `.handleXITOffsetMessage()` — L3326
- `.handleRxBinEnable()` — L3343
- `.handleRxApfEnable()` — L3359
- `.handleRxNfEnable()` — L3384
- `.handleLock()` — L3400
- `.handleVFOLock()` — L3419
- `.handleSqlEnable()` — L3436
- `.handleSqlLevel()` — L3452
- `.handleDiglOffset()` — L3469
- `.handleDiguOffset()` — L3481
- `.handleCwMacrosSpeed()` — L3493
- `.handleCwMacrosDelay()` — L3505
- `.handleCwKeyerSpeed()` — L3517
- `.handleCwMacrosSpeedUp()` — L3529
- `.handleCwMacrosSpeedDown()` — L3535
- `.handleCwMacros()` — L3541
- `.handleCwTerminal()` — L3550
- `.handleCwMsg()` — L3559
- `.handleCwMacrosStop()` — L3578
- `.handleKeyer()` — L3582
- `.handleTrxMessage()` — L3594
- `.shouldIgnoreTrxForCurrentCwBreakIn()` — L3696
- `.handleIF()` — L3710
- `.handleDDS()` — L3808
- `.handleVFOMessage()` — L3859
- `.handleModulationMessage()` — L3972
- `.handleDeleteSpot()` — L4076
- `.lineOutEnable()` — L4085
- `.handleLineOutStart()` — L4112
- `.handleLineOutStop()` — L4125
- `.handleDrive()` — L4138
- `.handleTuneDrive()` — L4165
- `.handleMute()` — L4214
- `.handleMuteRX()` — L4232
- `.handleMONEnable()` — L4256
- `.linearToDbVolume()` — L4273
- `.dbToLinearVolume()` — L4284
- `.handleMONVolume()` — L4296
- `.handleVolume()` — L4313
- `.handleSpotSimulateClick()` — L4325
- `.handleSpot()` — L4339
- `.handleTune()` — L4506
- `.handleRxFilterBand()` — L4529
- `.handleTXFilterBandEx()` — L4576
- `.handleRXEnable()` — L4595
- `.handleRxSensorsEnable()` — L4631
- `.handleTxSensorsEnable()` — L4642
- `.sendNREnable()` — L4654
- `.sendNBEnable()` — L4664
- `.handleNREnable()` — L4674
- `.handleRxNBEnable()` — L4707
- `.sendAnfEnable()` — L4740
- `.handleAnfEnable()` — L4746
- `.dbToAudioGain()` — L4767
- `.audioGainToDb()` — L4778
- `.sendRxVolume()` — L4788
- `.handleRxVolume()` — L4794
- `.handleRxBalance()` — L4856
- `.handleRxStepAttEnabledEx()` — L4883
- `.handleRxStepAttEx()` — L4905
- `.handleRxPreampAttEx()` — L4927
- `.handleVfoSyncEx()` — L4945
- `.handleVfoSwapEx()` — L4958
- `.handleFMDeviationEx()` — L4962
- `.handleAgcAutoEx()` — L4980
- `.handleAgcMode()` — L4996
- `.handleAgcGain()` — L5011
- `.sendCTUN()` — L5028
- `.handleCTUN()` — L5034
- `.sendTXProfile()` — L5053
- `.sendTXProfiles()` — L5059
- `.handleTXProfile()` — L5070
- `.handleTXProfiles()` — L5086
- `.handleShutdown()` — L5090
- `.sendCalibration()` — L5104
- `.handleCalibration()` — L5114
- `.handleRunCatCommand()` — L5122
- `.splitTextCommands()` — L5153
- `.parseTextFrame()` — L5259
- `.getDefaultAudioStreamSamples()` — L5606
- `.getBytesPerSample()` — L5622
- `.writeUInt32()` — L5637
- `.buildStreamPayload()` — L5645
- `.encodeSamples()` — L5669
- `.decodeSamples()` — L5712
- `.convertStreamSamplesToComplex()` — L5744
- `.sendIQSampleRate()` — L5769
- `.sendAudioSampleRate()` — L5777
- `.sendAudioStreamSampleType()` — L5782
- `.sendAudioStreamChannels()` — L5787
- `.sendAudioStreamSamples()` — L5792
- `.sendTxStreamAudioBuffering()` — L5797
- `.wantsIQStream()` — L5802
- `.wantsAudioStream()` — L5811
- `.IsReadyForStreaming()` — L5819
- `.WantsAnyRxStream()` — L5824
- `.PublishIQSamples()` — L5832
- `.PublishRxAudioSamples()` — L5842
- `.SendTxChrono()` — L5920
- `.UsesTCITxAudio()` — L5940
- `.UsesActiveTCITxAudio()` — L5948
- `.TryGetTxAudioRequestSettings()` — L5955
- `.SyncTciPttToMox()` — L5965
- `.clearQueuedTxAudio()` — L5983
- `.TryDequeueTxAudio()` — L5991
- `.handleBinaryFrame()` — L6007
- `.handleIQSampleRate()` — L6110
- `.getCurrentMaxHWSampleRate()` — L6129
- `.handleAudioSampleRate()` — L6145
- `.handleIQStart()` — L6202
- `.sendIQStartStop()` — L6219
- `.applyIQSampleRateToReceiver()` — L6224
- `.handleRxChannelEnable()` — L6252
- `.sendRxChannelEnable()` — L6292
- `.handleAudioStart()` — L6296
- `.handleAudioStreamSampleType()` — L6313
- `.handleAudioStreamChannels()` — L6340
- `.handleAudioStreamSamples()` — L6356
- `.handleTxStreamAudioBuffering()` — L6390
- `.PingFrameTimer()` — L6406
- `.VFOcallback()` — L6411
- `.Centrecallback()` — L6416
- `.VFOChange()` — L6421
- `.CentreChange()` — L6441

#### `TCIOutboundPriority` (type, L686)

_No extracted members._

#### `TCIOutboundFrame` (type, L696)

_No extracted members._

#### `TCIRxAudioResamplerState` (type, L702)

_No extracted members._

#### `VFOData` (type, L719)

_No extracted members._

#### `EOpcodeType` (type, L1528)

_No extracted members._

#### `TCPIPtciServer` (type, L6483)

- `.Init()` — L6571
- `.StartServer()` — L6664
- `.StopServer()` — L6832
- `.GetCwMacrosSpeed()` — L6950
- `.SetCwMacrosSpeed()` — L6955
- `.GetCwMacrosDelay()` — L6960
- `.SetCwMacrosDelay()` — L6965
- `.GetCwKeyerSpeed()` — L6970
- `.SetCwKeyerSpeed()` — L6975
- `.IncreaseCwMacrosSpeed()` — L6980
- `.DecreaseCwMacrosSpeed()` — L6985
- `.SetCwTerminalEnabled()` — L6990
- `.SendCwMacro()` — L6995
- `.SendCwMessage()` — L7000
- `.UpdateCwMessageCallsign()` — L7005
- `.StopCwMacros()` — L7010
- `.HandleCwKeyer()` — L7015
- `.NotifyCwTciPttReleased()` — L7020
- `.OnSocketListenerDisconnected()` — L7025
- `.OnCwMacrosEmpty()` — L7030
- `.OnCwCallsignSent()` — L7043
- `.OnCwMacrosSpeedChanged()` — L7056
- `.OnCwMacrosDelayChanged()` — L7072
- `.OnCwRemoteCharacterStarted()` — L7085
- `.OnCwKeyerSpeedChanged()` — L7090
- `.StopAllSocketListers()` — L7124
- `.ServerThreadStart()` — L7145
- `.PurgingThreadStart()` — L7199
- `.ClientConnectedHandler()` — L7238
- `.ClientDisconnectedHandler()` — L7243
- `.ClientErrorHandler()` — L7248
- `.OnVFOAFrequencyChangeHandler()` — L7254
- `.OnVFOBFrequencyChangeHandler()` — L7285
- `.OnMoxChangeHandler()` — L7310
- `.OnMoxPreChangeHandler()` — L7326
- `.OnModeChangeHandler()` — L7340
- `.OnBandChangeHandler()` — L7352
- `.OnCentreFrequencyChanged()` — L7364
- `.OnFilterChanged()` — L7391
- `.OnFilterEdgesChanged()` — L7403
- `.OnTXFiltersChanged()` — L7415
- `.OnPowerChangeHander()` — L7427
- `.OnThetisFocusChanged()` — L7439
- `.OnRX2EnabledChanged()` — L7451
- `.OnHWSampleRateChanged()` — L7463
- `.OnDrivePowerChanged()` — L7475
- `.OnTuneChanged()` — L7487
- `.OnSplitChanged()` — L7499
- `.OnSpotClicked()` — L7512
- `.OnMuteChanged()` — L7525
- `.OnNrChanged()` — L7537
- `.OnNbChanged()` — L7549
- `.OnAnfChanged()` — L7561
- `.OnBinChanged()` — L7573
- `.OnAGCModeChanged()` — L7585
- `.OnAGCAutoModeChanged()` — L7597
- `.OnVFOSyncChanged()` — L7609
- `.OnVfoALockChanged()` — L7623
- `.OnVfoBLockChanged()` — L7636
- `.OnSqlChanged()` — L7650
- `.OnSqlLevelChanged()` — L7662
- `.OnApfChanged()` — L7674
- `.OnTnfChanged()` — L7686
- `.OnDiglOffsetChanged()` — L7698
- `.OnDiguOffsetChanged()` — L7710
- `.OnRxAfGainChanged()` — L7722
- `.OnCTUNChanged()` — L7734
- `.OnTXProfileChanged()` — L7746
- `.OnTXProfilesChanged()` — L7758
- `.OnCalibrationChanged()` — L7770
- `.OnMONChanged()` — L7782
- `.OnMONVolumeChanged()` — L7794
- `.OnVolumeChanged()` — L7806
- `.OnBalanceChanged()` — L7818
- `.OnAttenuatorDataChanged()` — L7830
- `.OnStepAttEnabledChanged()` — L7842
- `.OnPreampModeChanged()` — L7854
- `.OnFMDeviationChanged()` — L7866
- `.OnAGCGainChanged()` — L7878
- `.OnRITChanged()` — L7890
- `.OnXITChanged()` — L7902
- `.OnRITValueChanged()` — L7914
- `.OnXITValueChanged()` — L7926
- `.OnTXFrequencyChanged()` — L7938
- `.OnMeterReadingsChanged()` — L7969
- `.ShowLog()` — L7981
- `.CloseLog()` — L7986
- `.SendSpotSimulationClickToAll()` — L7990
- `.RefreshStreamRunState()` — L8003
- `.PublishIQSamples()` — L8034
- `.PublishRxAudioSamples()` — L8047
- `.RequiresRxSensorUpdate()` — L8060
- `.SensorRequiresUpdate()` — L8076
- `.MinimumRequiredRxSensorInterval()` — L8092
- `.MinimumRequiredTxSensorInterval()` — L8112
- `.GetActiveTxAudioListener()` — L8132
- `.TryAcquireActiveTxAudioListener()` — L8146
- `.ReleaseActiveTxAudioListener()` — L8167
- `.UsesActiveTCITxAudio()` — L8176
- `.TryGetTxAudioRequestSettings()` — L8184
- `.RefreshTxAudioSourceState()` — L8200
- `.SendTxChrono()` — L8205
- `.TryDequeueTxAudio()` — L8216

#### `TCICWController` (type, L8231)

- `.Dispose()` — L8303
- `.GetMacroSpeed()` — L8346
- `.SetMacroSpeed()` — L8357
- `.SetMacroSpeedSilently()` — L8370
- `.GetMacroDelayMs()` — L8383
- `.SetMacroDelayMs()` — L8388
- `.GetKeyerSpeed()` — L8393
- `.SetKeyerSpeed()` — L8398
- `.IncreaseMacroSpeed()` — L8403
- `.DecreaseMacroSpeed()` — L8408
- `.SetTerminalEnabled()` — L8413
- `.SendMacro()` — L8449
- `.SendMessage()` — L8464
- `.HandleKeyer()` — L8479
- `.UpdatePendingCallsign()` — L8521
- `.OnRemoteCharacterStarted()` — L8539
- `.Stop()` — L8563
- `.HandleTciPttReleased()` — L8612
- `.DisconnectClient()` — L8625
- `.clampMacroSpeed()` — L8642
- `.decodeTciText()` — L8647
- `.normalizeMessageField()` — L8653
- `.translateAbbreviationToken()` — L8659
- `.buildRepeatedCallsign()` — L8678
- `.parseCallsignBase()` — L8686
- `.parseMacroText()` — L8703
- `.buildMacroOperation()` — L8764
- `.buildMessageOperation()` — L8778
- `.PollCallback()` — L8817
- `.startNextOperationLocked()` — L8883
- `.queueNextSegmentLocked()` — L8906
- `.completeActiveOperationLocked()` — L8931
- `.isCWModeLocked()` — L8946
- `.abortOperationsForNonCWLocked()` — L8956
- `.isCurrentOwnerLocked()` — L8987
- `.tryAcquireOwnershipLocked()` — L8992
- `.releaseOwnershipIfIdleLocked()` — L9005
- `.KeyerSchedulerThreadProc()` — L9012
- `.scheduleKeyerReleaseLocked()` — L9050
- `.tryReleaseDirectKeyerFromPollLocked()` — L9066
- `.releaseKeyerLocked()` — L9095
- `.cancelKeyerReleaseScheduleLocked()` — L9114
- `.isCwTargetAvailableLocked()` — L9121
- `.selectCwTargetLocked()` — L9127
- `.ensureTerminalTciPttLocked()` — L9149
- `.releaseTerminalTciPttIfOwnedLocked()` — L9168
- `.isTerminalEnabledLocked()` — L9178
- `.isAnyTerminalEnabledLocked()` — L9183
- `.beginDirectKeyerLocked()` — L9194
- `.ensureDirectKeyerMoxLocked()` — L9212
- `.captureDirectKeyerMoxReleaseLocked()` — L9240
- `.releaseDirectKeyerMox()` — L9247
- `.waitForScheduledKeyerRelease()` — L9257
- `.millisecondsToStopwatchTicks()` — L9280
- `.stopwatchTicksToMilliseconds()` — L9286
- `.setDirectKeyerState()` — L9291
- `.InvokeOnConsole()` — L9305

#### `CWTxSegment` (type, L8235)

_No extracted members._

#### `CWTextParseResult` (type, L8241)

_No extracted members._

#### `CWTxOperation` (type, L8248)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/TCIServer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
