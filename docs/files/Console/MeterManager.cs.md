# `Console/MeterManager.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** The metering subsystem (~30k lines): collects signal/power/SWR/voltage readings, renders configurable meter faces (analog, bar, LED) via DirectX, and manages multiple meter containers.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×68, references ×5)
  - `Console/console.cs` (calls ×20, references ×1)
  - `Console/TCIServer.cs` (references ×5, calls ×2)
  - `Console/ucSignalSelect.cs` (references ×4)
  - `Console/CAT/TCPIPcatServer.cs` (calls ×2)
  - `Console/dsp.cs` (references ×2)
  - `Console/frmVariablePicker.cs` (calls ×2)
  - `Console/CAT/CATCommands.cs` (calls ×1)
  - `Console/clsMeterScriptEngine.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×83)
  - `Console/ucOtherButtonsOptionsGrid.cs` (calls ×11, references ×8)
  - `Console/common.cs` (calls ×16)
  - `Console/clsCATMessageQueue.cs` (calls ×7, references ×1)
  - `Console/clsBandStackManager.cs` (calls ×8)
  - `Console/clsMeterScriptEngine.cs` (calls ×8)
  - `Console/ucMeter.Designer.cs` (references ×7)
  - `Console/clsCatAtonic.cs` (references ×4, calls ×3)
  - `Console/Andromeda/Andromeda.cs` (references ×6)
  - `Console/HPSDR/specHPSDR.cs` (calls ×5, references ×1)
  - `Console/clsImgeFetcher.cs` (references ×1, calls ×1)
  - `Console/clsDiscord.cs` (calls ×2)
  - …and 9 more files
- Most-referenced symbols from other files: `.RequiresUpdate()` (×9), `.FilterItemFrequencies()` (×3), `.GetContainerNotes()` (×3), `.QuickestUpdateInterval()` (×3), `.RemoveMMIO()` (×3), `.Start()` (×2), `.Stop()` (×2), `.MeterName()` (×2)

## Outline

### Types

#### `Reading` (type, L78)

_No extracted members._

#### `MeterType` (type, L178)

_No extracted members._

#### `BandGroups` (type, L234)

_No extracted members._

#### `Globals` (type, L241)

_No extracted members._

#### `MeterManager` (type, L245)

- `.SetLedVariable()` — L404
- `.provide_variables()` — L410
- `.CatVariables()` — L441
- `.OnCatQmessage()` — L454
- `.guard_holds()` — L485
- `.eval_now()` — L509
- `.normalise_var()` — L552
- `.OnCatQstate()` — L559
- `.rebuildAllLedReadings()` — L565
- `.GetLedFrom4Char()` — L576
- `.GetVoiceRecordPlayFrom4Char()` — L592
- `.GetOtherButtonsFromID()` — L608
- `.GetWebImageIDsFrom4Char()` — L629
- `.IsWebImageBackgroundShown()` — L647
- `.FilterItemFrequencies()` — L662
- `.GetFilterItemFrequencies()` — L715
- `.ReadingsCustom()` — L2010
- `.ZeroReading()` — L2015
- `.GetMeterTXRXType()` — L2144
- `.MeterName()` — L2206
- `.ReadingName()` — L2261
- `.ReadingUnits()` — L2322
- `.UpdateMeters()` — L2376
- `.normaliseWaveRecordPath()` — L2423
- `.waveRecordPathsEqual()` — L2437
- `.loadImages()` — L2445
- `.loadImage()` — L2495
- `.loadResouceImages()` — L2519
- `.clearAllCachedImageData()` — L2574
- `.removeImageCacheData()` — L2594
- `.addBitmap()` — L2619
- `.GetBitmap()` — L2631
- `.ContainsBitmap()` — L2641
- `.AddStreamData()` — L2649
- `.GetStreamData()` — L2659
- `.ContainsStreamData()` — L2669
- `.RemoveStreamData()` — L2677
- `.ContainerBorder()` — L2695
- `.NoTitle()` — L2705
- `.AutoContainerHeight()` — L2716
- `.SetContainerHeight()` — L2727
- `.ContainerMinimises()` — L2745
- `.RefreshContainerVisible()` — L2766
- `.containerVisible()` — L2771
- `.ContainerHidesWhenRXNotUsed()` — L2784
- `.LockContainer()` — L2827
- `.SetContainerRX()` — L2837
- `.GetContainerRX()` — L2870
- `.updateHiddenMacroInfo()` — L2882
- `.HiddenByMacro()` — L2894
- `.ShowContainerOnRX()` — L2946
- `.ShowContainerOnTX()` — L2965
- `.enableContainer()` — L2984
- `.ContainerHasBorder()` — L3071
- `.ContainerNoTitleBar()` — L3082
- `.ContainerAutoHeight()` — L3093
- `.ContainerLocked()` — L3104
- `.ContainerIsHidden()` — L3126
- `.ContainerShowOnRX()` — L3162
- `.ContainerShowOnTX()` — L3173
- `.ContainerNotes()` — L3195
- `.GetContainerNotes()` — L3207
- `.ContainerBackgroundColour()` — L3218
- `.GetContainerBackgroundColour()` — L3232
- `.HighlightContainer()` — L3243
- `.DisposeImageData()` — L3257
- `.SetAntennaAuxText()` — L3274
- `.initAntennaArrays()` — L3284
- `.Init()` — L3336
- `.addRenderer()` — L3362
- `.UpdateS9()` — L3369
- `.RefreshAllImages()` — L3378
- `.loadDXSkinImages()` — L3393
- `.RunRendererDisplay()` — L3408
- `.RunAllRendererDisplays()` — L3419
- `.SetVsync()` — L3428
- `.removeRenderer()` — L3453
- `.GetMeterIDsFromSaveData()` — L3462
- `.Shutdown()` — L3524
- `.addDelegates()` — L3568
- `.removeDelegates()` — L3679
- `.OnContainerHiddenByMacro()` — L3795
- `.OnContainerEnabled()` — L3807
- `.OnCWXShown()` — L3819
- `.OnCWPitchChanged()` — L3831
- `.OnTNFChanged()` — L3844
- `.OnRXSpecGridMinMaxChanged()` — L3856
- `.OnTXSpecGridMinMaxChanged()` — L3869
- `.OnWaterfallRXGradientChanged()` — L3883
- `.OnWaterfallTXGradientChanged()` — L3895
- `.OnRXWaterfallMinMaxChanged()` — L3907
- `.OnTXWaterfallMinMaxChanged()` — L3920
- `.OnPAProfileChanged()` — L3934
- `.OnTXProfileChanged()` — L3948
- `.OnTXFiltersChanged()` — L3962
- `.OnTuneStepIndexChanged()` — L3977
- `.OnVFOSyncChanged()` — L3990
- `.OnVFOaLockChanged()` — L4002
- `.OnVFObLockChanged()` — L4014
- `.updateAntennaMeters()` — L4027
- `.OnTXFrequencyChanged()` — L4038
- `.OnAntennaRXChanged()` — L4066
- `.OnAntennaTXChanged()` — L4078
- `.OnAntennaAuxChanged()` — L4090
- `.OnAntennaDoNotTX()` — L4102
- `.OnAntennaRxTx()` — L4113
- `.GetNoTX()` — L4130
- `.GetRXAntState()` — L4138
- `.GetTXAntState()` — L4147
- `.GetRXAuxState()` — L4156
- `.GetRXAuxName()` — L4165
- `.GetTXEnabled()` — L4173
- `.GetRXAnt()` — L4181
- `.GetTXAnt()` — L4210
- `.OnSplitChanged()` — L4220
- `.OnMinimumNotchWidthChangedRX()` — L4236
- `.OnMinimumNotchWidthChangedTX()` — L4245
- `.OnFilterEdgesChanged()` — L4254
- `.OnFilterChanged()` — L4258
- `.updateFilterInfo()` — L4262
- `.OnMultiRxChanged()` — L4311
- `.OnTXBandChanged()` — L4326
- `.OnVFOTXChanged()` — L4338
- `.OnModeChangeHandler()` — L4352
- `.OnVHFDetailsChanged()` — L4377
- `.OnBandPanelChanged()` — L4391
- `.bandChange()` — L4405
- `.OnBandChange()` — L4476
- `.OnPreBandChange()` — L4480
- `.OnTransverterIndexChanged()` — L4484
- `.OnAlexPresentChanged()` — L4488
- `.OnPAPresentChanged()` — L4492
- `.OnApolloPresentChanged()` — L4496
- `.OnCurrentModelChanged()` — L4500
- `.OnVFOA()` — L4504
- `.OnVFOB()` — L4545
- `.OnVFOASub()` — L4578
- `.OnDuplexChanged()` — L4597
- `.OnPSAChanged()` — L4614
- `.OnQuickPlayChanged()` — L4627
- `.OnANFChanged()` — L4640
- `.OnSNBChanged()` — L4652
- `.OnAVGChanged()` — L4664
- `.OnPeakChanged()` — L4676
- `.OnCTUNChanged()` — L4688
- `.OnVACEnabledChanged()` — L4700
- `.OnMuteChanged()` — L4720
- `.OnBINChanged()` — L4732
- `.OnPanSwapChanged()` — L4744
- `.OnDisplayModeChanged()` — L4757
- `.OnAGCModeChanged()` — L4769
- `.OnAGCAutoModeChanged()` — L4781
- `.OnSqlChanged()` — L4793
- `.OnGeneralSettingsChanged()` — L4812
- `.OnXPAChanged()` — L4825
- `.OnQuickRecordChanged()` — L4839
- `.OnWaveRecordChanged()` — L4852
- `.OnTwoToneChanged()` — L4865
- `.OnTuneChanged()` — L4877
- `.OnNRChanged()` — L4889
- `.OnNBChanged()` — L4901
- `.OnMONChanged()` — L4913
- `.initAllConsoleData()` — L4926
- `.initConsoleData()` — L4946
- `.getFilterName()` — L5167
- `.OnPower()` — L5223
- `.OnPreMox()` — L5242
- `.OnMox()` — L5252
- `.OnEQChanged()` — L5314
- `.OnLevelerChanged()` — L5327
- `.OnCFCChanged()` — L5341
- `.OnCompandChanged()` — L5355
- `.OnQuickSplitChanged()` — L5368
- `.dbmOffsetForAboveS9Frequency()` — L5419
- `.IsAboveS9Frequency()` — L5426
- `.getReading()` — L5447
- `.setReading()` — L5456
- `.setReadingForced()` — L5461
- `.OnMeterReadings()` — L5468
- `.RequiresUpdate()` — L5543
- `.QuickestUpdateInterval()` — L5552
- `.KeycodeInUse()` — L5569
- `.AbortAllVoiceRecordRepeatPlaybacks()` — L5587
- `.MeterFromId()` — L5598
- `.MeterExists()` — L5608
- `.onGlobalKeyDown()` — L5616
- `.onGlobalKeyUp()` — L5627
- `.AddMeterContainer()` — L5638
- `.IsOnTop()` — L5699
- `.FinishSetupAndDisplay()` — L5712
- `.BringToFront()` — L5774
- `.ucMeter_SettingsClicked()` — L5809
- `.ucMeter_FloatingDockedClicked()` — L5817
- `.ucMeter_FloatingDockedMoved()` — L5831
- `.SetPositionOfDockedMeters()` — L5837
- `.setPoisitionOfDockedMeter()` — L5847
- `.returnMeterFromFloating()` — L5892
- `.setMeterFloating()` — L5919
- `.OnRX2EnabledChanged()` — L5944
- `.OnRX2EnabledPreChanged()` — L5983
- `.zeroAllMeters()` — L6026
- `.RestoreSettings()` — L6037
- `.GetFormGuidList()` — L6131
- `.ContainerToString()` — L6147
- `.ContainerFromString()` — L6217
- `.StoreSettings2()` — L6416
- `.RecoverContainer()` — L6539
- `.RemoveMeterContainer()` — L6558

#### `FilterItemSnapFrequencies` (type, L248)

_No extracted members._

#### `CustomReadings` (type, L732)

- `.setupReadings()` — L750
- `.formatNumber()` — L839
- `.formatElapsedTimeCompact()` — L849
- `.formatElapsedTime()` — L870
- `.GetPlaceholders()` — L899
- `.TakeReading()` — L923
- `.IsCustomString()` — L930
- `.GetReading()` — L999
- `.getIntPart()` — L1249
- `.returnTuneStep()` — L1267
- `.returnPAProfile()` — L1278
- `.UpdateReadings()` — L1289
- `.addReading()` — L1387
- `.addReadingText()` — L1391
- `.GetAvailableReadings()` — L1395

#### `clsIGSettings` (type, L1433)

- `.SetSetting()` — L1510
- `.GetSetting()` — L1524
- `.ToString2()` — L1573
- `.TryParse2()` — L1630
- `.TryParse()` — L1734
- `.GetMMIOGuid()` — L1972
- `.SetMMIOGuid()` — L1976
- `.GetMMIOVariable()` — L1980
- `.SetMMIOVariable()` — L1984

#### `GeneralOtherButtonSettings` (type, L4805)

_No extracted members._

#### `clsMeterItem` (type, L6625)

- `.PrepareCalibration()` — L6815
- `.AddPerc()` — L6883
- `.GetPerc()` — L6899
- `.HasPerc()` — L6908
- `.Update()` — L7049
- `.ClearHistory()` — L7072
- `.History()` — L7075
- `.ToString()` — L7124
- `.TryParse()` — L7128
- `.HandleIncrement()` — L7140
- `.HandleDecrement()` — L7143
- `.ZeroOut()` — L7146
- `.MouseClick()` — L7152
- `.MouseDown()` — L7156
- `.MouseUp()` — L7160
- `.KeyDown()` — L7198
- `.KeyUp()` — L7202
- `.MouseWheel()` — L7206
- `.Removing()` — L7210
- `.BandPanelsChanged()` — L7214
- `.BandChanged()` — L7218
- `.Initialise()` — L7222
- `.ModeChanged()` — L7226
- `.TuneStepIndexChanged()` — L7230
- `.FilterChanged()` — L7234
- `.TXFilterChanged()` — L7238
- `.PAProfileChanged()` — L7242
- `.TXProfileChanged()` — L7246
- `.SetRXSpectrumGridMin()` — L7250
- `.SetRXSpectrumGridMax()` — L7254
- `.SetTXSpectrumGridMin()` — L7258
- `.SetTXSpectrumGridMax()` — L7262
- `.SetRXWaterfallMin()` — L7266
- `.SetRXWaterfallMax()` — L7270
- `.SetTXWaterfallMin()` — L7274
- `.SetTXWaterfallMax()` — L7278
- `.WaterfallRXGradient()` — L7282
- `.WaterfallTXGradient()` — L7286
- `.ContainerEnabled()` — L7340
- `.ContainerHiddenByMacro()` — L7344

#### `MeterItemType` (type, L6627)

_No extracted members._

#### `clsPercCache` (type, L6666)

_No extracted members._

#### `clsItemGroup` (type, L7521)

- `.ToString()` — L7543
- `.TryParse()` — L7565

#### `clsClickBox` (type, L7597)

_No extracted members._

#### `clsSolidColour` (type, L7657)

_No extracted members._

#### `clsFadeCover` (type, L7691)

_No extracted members._

#### `clsFilterButtonBox` (type, L7700)

- `.Initialise()` — L7714
- `.FilterChanged()` — L7731
- `.InitFilterButtons()` — L7751
- `.setupButtons()` — L7768
- `.setupClick()` — L7861
- `.MouseDown()` — L7910
- `.MouseUp()` — L7923
- `.setFilter()` — L7956

#### `clsTunestepButtons` (type, L8025)

- `.setupButtons()` — L8059
- `.setupClick()` — L8166
- `.MouseDown()` — L8215
- `.MouseUp()` — L8228
- `.TuneStepIndexChanged()` — L8245

#### `clsOtherButtons` (type, L8251)

- `.Removing()` — L8332
- `.onPlayingChanged()` — L8340
- `.onRecordingChanged()` — L8356
- `.OnContainerVisible()` — L8404
- `.OnCatState()` — L8456
- `.GetMacroSettings()` — L8812
- `.SetMacroSettings()` — L8820
- `.ContainerHiddenByMacro()` — L8850
- `.updateOn()` — L8880
- `.Initialise()` — L8886
- `.GetVisibleBits()` — L8890
- `.SetVisibleBits()` — L8895
- `.isVisible()` — L8902
- `.try_index_from_group_bit()` — L8908
- `.setupButtons()` — L8971
- `.setupClick()` — L9273
- `.MouseDown()` — L9322
- `.MouseUp()` — L9346
- `.handleMacroButtonPress()` — L9428
- `.MoveButton()` — L9546
- `.Update()` — L9551

#### `clsAntennaButtonBox` (type, L9622)

- `.setupButtons()` — L9669
- `.formatAux()` — L9826
- `.AntennasChanged()` — L9834
- `.setupClick()` — L9899
- `.MouseDown()` — L9948
- `.MouseUp()` — L9961
- `.toggleTxRxAnt()` — L9986
- `.setRXAntenna()` — L9995
- `.setAuxAntenna()` — L10005
- `.setTXAntenna()` — L10015

#### `clsModeButtonBox` (type, L10071)

- `.Initialise()` — L10084
- `.ModeChanged()` — L10097
- `.setupButtons()` — L10109
- `.setupClick()` — L10183
- `.MouseDown()` — L10232
- `.MouseUp()` — L10245
- `.abortForLockedVFO()` — L10262
- `.setMode()` — L10275

#### `clsVoiceRecordPlay` (type, L10342)

- `.Removing()` — L10492
- `.RecordToSlot()` — L10509
- `.PlayFromSlot()` — L10519
- `.KeyUp()` — L10597
- `.KeyDown()` — L10641
- `.SlotDuration()` — L10703
- `.onJsonWritten()` — L10708
- `.onPlayingChanged()` — L10728
- `.onRecordingChanged()` — L10765
- `.setupRepeatTimer()` — L10798
- `.enableAllSlots()` — L10844
- `.AbortVoiceRecordRepeatPlayback()` — L10851
- `.IsPlaying()` — L10865
- `.IsRecording()` — L10871
- `.AtivateTime()` — L10877
- `.SetSlotLocked()` — L10883
- `.GetSlotLocked()` — L10889
- `.GetKeybind()` — L10895
- `.SetKeybind()` — L10900
- `.GetUsesKeybind()` — L10905
- `.SetUsesKeybind()` — L10910
- `.UsesKeybind()` — L10927
- `.SetCanRepeat()` — L10940
- `.GetCanRepeat()` — L10948
- `.SetRepeatDelay()` — L10954
- `.GetRepeatDelay()` — L10961
- `.GetRepeatEnabled()` — L10967
- `.SetRepeatEnabled()` — L10973
- `.GetDelayElapsed()` — L10979
- `.GetGainAdjust()` — L10988
- `.SetGainAdjust()` — L10994
- `.GetIgnorePlayTempChanges()` — L11000
- `.SetIgnorePlayTempChanges()` — L11006
- `.GetIgnoreRecordTempChanges()` — L11012
- `.SetIgnoreRecordTempChanges()` — L11018
- `.Initialise()` — L11040
- `.setupButtons()` — L11069
- `.MoveButton()` — L11207
- `.setupClick()` — L11217
- `.MouseDown()` — L11266
- `.MouseUp()` — L11295
- `.clearRunningRepeat()` — L11379
- `.handleClicked()` — L11386

#### `clsWaveRecord` (type, L11602)

- `.Initialise()` — L11754
- `.Removing()` — L11759
- `.onPlayingChanged()` — L11782
- `.onRecordingChanged()` — L11798
- `.CanAcceptFiles()` — L12073
- `.AddFiles()` — L12078
- `.IsPlaying()` — L12107
- `.SetRenderLayout()` — L12115
- `.UpdateDragFromMouse()` — L12137
- `.HitTest()` — L12143
- `.Update()` — L12157
- `.MouseWheel()` — L12162
- `.MouseDown()` — L12189
- `.MouseUp()` — L12236
- `.RefreshEntries()` — L12285
- `.buildJsonDataDisplay()` — L12347
- `.sanitiseJsonDataValue()` — L12370
- `.formatJsonDataDuration()` — L12392
- `.formatJsonDataAudio()` — L12401
- `.formatJsonDataUtc()` — L12412
- `.containsFilePathLocked()` — L12426
- `.sanitiseStoredPaths()` — L12432
- `.pathArraysEqual()` — L12452
- `.shortArraysEqual()` — L12466
- `.sanitiseOrderMap()` — L12482
- `.removeFromOrderMap()` — L12520
- `.handleHit()` — L12547
- `.handlePlay()` — L12566
- `.handleDelete()` — L12614
- `.moveEntry()` — L12660
- `.adjustScroll()` — L12707
- `.updateReorderDrag()` — L12718
- `.updateScrollDrag()` — L12738
- `.scrollTrackToLocked()` — L12755
- `.rowIndexFromPointLocked()` — L12772
- `.clampScrollLocked()` — L12790

#### `WaveRecordHitType` (type, L11604)

_No extracted members._

#### `WaveRecordEntry` (type, L11614)

_No extracted members._

#### `WaveRecordJsonDataDisplay` (type, L11622)

_No extracted members._

#### `WaveRecordHitRegion` (type, L11645)

_No extracted members._

#### `clsBandButtonBox` (type, L12802)

- `.Initialise()` — L12817
- `.BandPanelsChanged()` — L12852
- `.VHFUpdate()` — L12861
- `.BandChanged()` — L12876
- `.setupButtons()` — L12930
- `.setupClick()` — L13084
- `.MouseDown()` — L13133
- `.MouseUp()` — L13146
- `.setBand()` — L13224
- `.abortForLockedVFO()` — L13244

#### `clsDiscordButtonBox` (type, L13303)

- `.OnReady()` — L13324
- `.OnDisconnected()` — L13334
- `.setupButtons()` — L13345
- `.formatNumber()` — L13413
- `.setupClick()` — L13427
- `.MouseDown()` — L13476
- `.MouseUp()` — L13489
- `.sendMsg()` — L13573

#### `clsButtonBox` (type, L13627)

- `.setupArrays()` — L13727
- `.ResetButtons()` — L13797
- `.GetVisibleBits()` — L13893
- `.SetVisibleBits()` — L13894
- `.SetMacroSettings()` — L13895
- `.GetMacroSettings()` — L13896
- `.SetFillColour()` — L13922
- `.GetFillColour()` — L13927
- `.SetHoverColour()` — L13933
- `.GetHoverColour()` — L13938
- `.SetClickColour()` — L13944
- `.GetClickColour()` — L13949
- `.SetBorderColour()` — L13955
- `.GetBorderColour()` — L13960
- `.SetUseOffColour()` — L13966
- `.GetUseOffColour()` — L13972
- `.SetUseIndicator()` — L13978
- `.GetUseIndicator()` — L13983
- `.SetOn()` — L13989
- `.GetOn()` — L13994
- `.SetIndicatorWidth()` — L14000
- `.GetIndicatorWidth()` — L14005
- `.SetOnColour()` — L14011
- `.GetOnColour()` — L14016
- `.SetOffColour()` — L14022
- `.GetOffColour()` — L14027
- `.SetFontFamily()` — L14033
- `.GetFontFamily()` — L14038
- `.SetFontStyle()` — L14044
- `.GetFontStyle()` — L14049
- `.SetFontSize()` — L14055
- `.GetFontSize()` — L14060
- `.SetFontColour()` — L14066
- `.GetFontColour()` — L14071
- `.SetText()` — L14077
- `.GetText()` — L14084
- `.GetTextUpdateChanged()` — L14090
- `.GetTextChanged()` — L14099
- `.SetIconOn()` — L14105
- `.GetIconOn()` — L14110
- `.SetIconOff()` — L14116
- `.GetIconOff()` — L14121
- `.SetUseIcon()` — L14127
- `.GetUseIcon()` — L14132
- `.SetEnabled()` — L14138
- `.GetEnabled()` — L14143
- `.SetVisible()` — L14148
- `.GetVisible()` — L14164
- `.SetIndicatorType()` — L14170
- `.GetIndicatorType()` — L14175
- `.MoveButton()` — L14191

#### `IndicatorType` (type, L13629)

_No extracted members._

#### `clsVfoDisplay` (type, L14201)

- `.getVfo()` — L14397
- `.setVfo()` — L14430
- `.abortForLockedVFO()` — L14462
- `.adjustVfo()` — L14495
- `.KeyDown()` — L14527
- `.findOnePosition()` — L14560
- `.MouseWheel()` — L14572
- `.PopBandStack()` — L14584
- `.PopFilterMenu()` — L14599
- `.MouseDown()` — L14613
- `.MouseUp()` — L14631
- `.setTuneStep()` — L14809
- `.setLock()` — L14879
- `.setVfoSync()` — L14897
- `.setTX()` — L14905
- `.toggleSplit()` — L14933
- `.setFilter()` — L14948
- `.setMode()` — L14991
- `.setBand()` — L15030

#### `renderState` (type, L14203)

_No extracted members._

#### `buttonState` (type, L14211)

_No extracted members._

#### `DSPModeForModeDisplay` (type, L14232)

_No extracted members._

#### `VFODisplayMode` (type, L14247)

_No extracted members._

#### `clsClock` (type, L15395)

_No extracted members._

#### `clsWebImage` (type, L15485)

- `.OnWebImageRemoved()` — L15571
- `.OnShowWebImageBackground()` — L15578
- `.timerCallback()` — L15644
- `.Update()` — L15667
- `.ClearBackgroundFourChar()` — L15737
- `.Removing()` — L15767
- `.OnState()` — L15919
- `.OnImage()` — L15940

#### `clsImage` (type, L16025)

_No extracted members._

#### `clsScaleItem` (type, L16105)

- `.ToString()` — L16191

#### `clsNeedleScalePwrItem` (type, L16208)

- `.ToString()` — L16311

#### `clsRotatorItem` (type, L16362)

- `.SendRotatorMessage()` — L16475
- `.Update()` — L16564
- `.ZeroOut()` — L16696
- `.MouseUp()` — L16704
- `.MouseDown()` — L16711

#### `RotatorMode` (type, L16364)

_No extracted members._

#### `clsDialDisplay` (type, L16719)

- `.undoTuneStep()` — L17099
- `.Update()` — L17116
- `.MouseUp()` — L17134
- `.MouseWheel()` — L17157

#### `clsMagicEyeItem` (type, L17175)

- `.Update()` — L17208
- `.ClearHistory()` — L17319
- `.History()` — L17328
- `.ZeroOut()` — L17351

#### `clsDataOut` (type, L17367)

- `.Update()` — L17376

#### `clsSpacerItem` (type, L17436)

_No extracted members._

#### `clsHistoryItem` (type, L17469)

- `.ZeroOut()` — L17594
- `.addReading()` — L17755
- `.Update()` — L18096

#### `HistoryData` (type, L17471)

_No extracted members._

#### `clsFilterItem` (type, L18172)

- `.Initialise()` — L18366
- `.WaterfallRXGradient()` — L18498
- `.WaterfallTXGradient()` — L18509
- `.findNearestVGrid()` — L18586
- `.Removing()` — L18609
- `.TXProfileChanged()` — L18796
- `.TXFilterChanged()` — L18800
- `.FilterChanged()` — L18819
- `.MouseUp()` — L18858
- `.MouseWheel()` — L19013
- `.AdjustNotch()` — L19351
- `.Shift()` — L19378
- `.buildSpectrumGreyScale()` — L19560
- `.Update()` — L19616
- `.SetRXSpectrumGridMin()` — L19931
- `.SetRXSpectrumGridMax()` — L19939
- `.SetTXSpectrumGridMin()` — L19947
- `.SetTXSpectrumGridMax()` — L19955
- `.SetRXWaterfallMin()` — L19964
- `.SetRXWaterfallMax()` — L19972
- `.SetTXWaterfallMin()` — L19980
- `.SetTXWaterfallMax()` — L19988

#### `FIWaterfallPalette` (type, L18174)

_No extracted members._

#### `FIDisplayMode` (type, L18185)

_No extracted members._

#### `clsTextOverlay` (type, L20066)

- `.Update()` — L20529
- `.parseText()` — L20587
- `.ZeroOut()` — L20718

#### `clsLed` (type, L20768)

- `.Removing()` — L20935
- `.onTimerElapsedCondition()` — L20942
- `.UpdateReadings()` — L20961
- `.add_readings()` — L20965
- `.expand_placeholders()` — L21007
- `.RebuildCondition()` — L21066
- `.Update()` — L21177
- `.stopMox()` — L21206
- `.ZeroOut()` — L21213

#### `Led_Shape` (type, L20770)

_No extracted members._

#### `Led_Style` (type, L20776)

_No extracted members._

#### `clsBarItem` (type, L21237)

- `.Update()` — L21323
- `.ClearHistory()` — L21449
- `.History()` — L21458
- `.HandleIncrement()` — L21539
- `.HandleDecrement()` — L21557
- `.ZeroOut()` — L21584

#### `Units` (type, L21239)

_No extracted members._

#### `BarStyle` (type, L21246)

_No extracted members._

#### `clsSignalText` (type, L21606)

- `.Update()` — L21677
- `.ClearHistory()` — L21763
- `.History()` — L21772
- `.HandleIncrement()` — L21836
- `.HandleDecrement()` — L21852
- `.ZeroOut()` — L21870

#### `clsNeedleItem` (type, L21900)

- `.Update()` — L21974
- `.ClearHistory()` — L22131
- `.History()` — L22140
- `.ZeroOut()` — L22185

#### `NeedlePlacement` (type, L21902)

_No extracted members._

#### `NeedleStyle` (type, L21909)

_No extracted members._

#### `NeedleDirection` (type, L21915)

_No extracted members._

#### `clsText` (type, L22251)

- `.updateReadingText()` — L22274
- `.Update()` — L22293

#### `clsMeter` (type, L22359)

- `.addMeterItem()` — L22620
- `.MeterVariablesReadingString()` — L22628
- `.MeterVariables()` — L22684
- `.AddMeter()` — L22740
- `.AddSMeterBarSignal()` — L22822
- `.AddSMeterBarSignalAvg()` — L22826
- `.AddSMeterBarMaxBin()` — L22830
- `.getFadeCover()` — L22834
- `.addSMeterBar()` — L22846
- `.AddADCMaxMag()` — L22940
- `.AddSMeterBarText()` — L23001
- `.AddADCBar()` — L23063
- `.AddPBSNRBar()` — L23151
- `.AddAGCGainBar()` — L23222
- `.AddAGCBar()` — L23284
- `.AddCustomBar()` — L23374
- `.AddRotator()` — L23466
- `.AddMagicEye()` — L23572
- `.AddDial()` — L23619
- `.AddSpacer()` — L23647
- `.AddWebImage()` — L23676
- `.AddDataOut()` — L23704
- `.AddTextOverlay()` — L23728
- `.AddLed()` — L23756
- `.AddAnanMM()` — L23784
- `.AddCrossNeedle()` — L24140
- `.AddMicBar()` — L24327
- `.AddEQBar()` — L24414
- `.AddLevelerBar()` — L24502
- `.AddLevelerGainBar()` — L24588
- `.AddALCBar()` — L24649
- `.AddALCGainBar()` — L24735
- `.AddALCGroupBar()` — L24796
- `.AddCFCBar()` — L24857
- `.AddCFCGainBar()` — L24943
- `.AddCompBar()` — L25004
- `.AddPWRBar()` — L25177
- `.AddREVPWRBar()` — L25181
- `.AddSWRBar()` — L25313
- `.AddBandButtons()` — L25376
- `.AddDiscordButtons()` — L25412
- `.AddModeButtons()` — L25448
- `.AddVoiceRecordPlay()` — L25484
- `.AddWaveRecord()` — L25520
- `.AddFilterButtons()` — L25549
- `.AddHistory()` — L25585
- `.AddAntennaButtons()` — L25614
- `.AddTunestepButtons()` — L25650
- `.AddOtherButtons()` — L25686
- `.AddFilterDisplay()` — L25722
- `.AddVFODisplay()` — L25751
- `.AddClock()` — L25815
- `.GetBandGroupFromBand()` — L25864
- `.SetBandPanel()` — L25917
- `.GetWebImageIDsFrom4Char()` — L25947
- `.IsWebImageBackgroundShown()` — L25966
- `.UpdateFilterDetails()` — L25983
- `.PAProfileChanged()` — L25997
- `.TXProfileChanged()` — L26258
- `.UpdateTXFilterDetails()` — L26272
- `.InitFilterButtons()` — L26286
- `.ModeChanged()` — L26314
- `.TuneStepIndexChanged()` — L26328
- `.BandChanged()` — L26342
- `.AntennasChanged()` — L26356
- `.BandPanelsChanged()` — L26370
- `.VHFDetailsChanged()` — L26384
- `.ZeroOut()` — L26398
- `.RebuildLedReadings()` — L26421
- `.RebuildLedConditions()` — L26436
- `.KeyDown()` — L26452
- `.KeyUp()` — L26463
- `.TryParse()` — L26474
- `.numberOfMeterGroups()` — L26517
- `.removeMeterItem()` — L26541
- `.RemoveMeterType()` — L26570
- `.GetMeterItem()` — L26613
- `.MetersVoiceItemsUseThisKeycode()` — L26649
- `.AbortAllVoiceRecordRepeatPlaybacks()` — L26671
- `.MeterHasLockedVoiceRecords()` — L26691
- `.RemoveAllMeterTypes()` — L26713
- `.HasMeterType()` — L26746
- `.DisableMeterType()` — L26773
- `.MeterGroupID()` — L26814
- `.Find4Chars()` — L26844
- `.ApplySettingsForMeterGroup()` — L26908
- `.GetSettingsForMeterGroup()` — L28355
- `.GetOrderForMeterType()` — L29384
- `.SetOrderForMeterType()` — L29414
- `.UpdateAlways()` — L29500
- `.itemFromID()` — L29531
- `.getBounds()` — L29541
- `.LedIndicatorFromFourChar()` — L29567
- `.VoiceRecordPlayFromFourChar()` — L29585
- `.itemsFromID()` — L29603
- `.hasReading()` — L29620
- `.GetBottom()` — L29636
- `.UpdateItems()` — L29672
- `.getMeterGroups()` — L29729
- `.UpdateIntervals()` — L29756
- `.Rebuild()` — L29761
- `.setupSortedZOrder()` — L29795
- `.MouseUp()` — L29814
- `.updateVfoABandText()` — L29976
- `.updateVfoBBandText()` — L29996
- `.UpdateBandText()` — L30240
- `.ContainerEnabled()` — L30625
- `.ContainerHiddenByMacro()` — L30636
- `.QuickestUpdateInterval()` — L30913
- `.addUpdateReading()` — L31017
- `.Update()` — L31024
- `.clearB()` — L31134
- `.DelayForUpdate()` — L31139
- `.incrementDisplayGroup()` — L31165
- `.decrementDisplayGroup()` — L31170
- `.incrementMeterItem()` — L31175
- `.decrementMeterItem()` — L31181
- `.AddDisplayGroup()` — L31201
- `.RemoveDisplayGroup()` — L31208
- `.ToString()` — L31225

#### `clsReading` (type, L31244)

_No extracted members._

#### `clsReadings` (type, L31250)

- `.GetReading()` — L31263
- `.SetReading()` — L31269
- `.RequiresUpdate()` — L31275
- `.UseReading()` — L31280

#### `DXRenderer` (type, L31290)

- `.HandleTouchDown()` — L31454
- `.HandleTouchUp()` — L31460
- `.HandleTouchMove()` — L31466
- `.RunDisplay()` — L31483
- `.RemoveAnySkinImages()` — L31530
- `.convertImageToDX()` — L31545
- `.getMaxSamples()` — L31573
- `.dxInit()` — L31586
- `.setupFilterWaterfallBitmap()` — L31809
- `.resetDX2DModeDescription()` — L31845
- `.ShutdownDX()` — L31870
- `.RemoveAllDXImages()` — L31984
- `.RemoveDXImage()` — L32004
- `.dxRender()` — L32036
- `.setupAliasing()` — L32193
- `.releaseDXResources()` — L32245
- `.buildDXFonts()` — L32252
- `.releaseDXFonts()` — L32258
- `.clearAllDynamicBrushes()` — L32274
- `.clearAllDynamicTextFormats()` — L32288
- `.getDXBrushForColour()` — L32305
- `.convertColour()` — L32333
- `.getDXTextFormatForFont()` — L32338
- `.resizeDX()` — L32374
- `.target_Resize()` — L32429
- `.target_VisibleChanged()` — L32436
- `.OnMouseEnter()` — L32442
- `.OnMouseLeave()` — L32462
- `.OnMouseCaptureChanged()` — L32483
- `.OnMouseMove()` — L32504
- `.OnMouseWheel()` — L32554
- `.OnMouseClick()` — L32604
- `.OnMouseDown()` — L32658
- `.OnMouseUp()` — L32712
- `.OnDragEnter()` — L32787
- `.OnDragOver()` — L32800
- `.OnDragDrop()` — L32813
- `.tryGetWaveRecordDropTarget()` — L32824
- `.drawMeters()` — L32869
- `.measureString()` — L33023
- `.fade()` — L33153
- `.renderNeedleScale()` — L33179
- `.tidyPower()` — L33385
- `.renderScale()` — L33393
- `.generalScale()` — L33872
- `.renderGroup()` — L33959
- `.slits()` — L33977
- `.renderLed()` — L34010
- `.renderTextOverlay()` — L34135
- `.buildWaterfall()` — L34289
- `.renderDialDisplay()` — L35284
- `.renderFilterDisplay()` — L35466
- `.hzToPixels()` — L36272
- `.pixelsToHz()` — L36277
- `.isMouseNearLine()` — L36282
- `.pointToSegmentDistance()` — L36288
- `.distanceBetweenPoints()` — L36310
- `.renderHistory()` — L36315
- `.renderSpacer()` — L36544
- `.renderWebImage()` — L36578
- `.convertDegreesToCardinal()` — L36671
- `.renderRotator()` — L36704
- `.calculateDistance()` — L37227
- `.renderEye()` — L37233
- `.renderText()` — L37357
- `.renderHBarMarkersOnly()` — L37388
- `.renderHBar()` — L37424
- `.renderSolidColour()` — L37813
- `.renderFadeCover()` — L37826
- `.getParts()` — L37847
- `.plotText()` — L37857
- `.highlightBox()` — L38052
- `.drawTuneStep()` — L38063
- `.drawBand()` — L38170
- `.drawMode()` — L38343
- `.drawFilter()` — L38438
- `.shrinkRectangle()` — L38552
- `.drawRoundedRectangle()` — L38570
- `.fillRoundedRectangle()` — L38579
- `.drawSafeLine()` — L38587
- `.dimColour()` — L38600
- `.renderWaveRecord()` — L38611
- `.drawWaveRecordButton()` — L38948
- `.drawWaveRecordIcon()` — L38962
- `.clipRect()` — L38990
- `.rectEmpty()` — L39000
- `.samePath()` — L39004
- `.renderButtonBox()` — L39009
- `.adjustTextColourForContrast()` — L39820
- `.calculateContrastRatio()` — L39849
- `.renderVfoDisplay()` — L39857
- `.renderClock()` — L40474
- `.renderSignalTextDisplay()` — L40538
- `.renderImage()` — L40629
- `.renderNeedle()` — L40740
- `.degToRad()` — L41033
- `.radToDeg()` — L41037
- `.getPerc()` — L41041
- `.bitmapFromSystemBitmap()` — L41249

#### `MultiMeterIO` (type, L41321)

- `.StartListeningUDP()` — L42777
- `.StartListeningTCPIP()` — L42805
- `.StartTcpClient()` — L42831
- `.StartSerialPort()` — L42859
- `.StopConnection()` — L42887
- `.StopConnections()` — L42912
- `.AlreadyConfigured()` — L42938
- `.GetSaveData()` — L42951
- `.RestoreSaveData2()` — L43006
- `.RestoreSaveData()` — L43036
- `.AddMMIO()` — L43163
- `.RemoveMMIO()` — L43176
- `.SendDataMMIO()` — L43186
- `.GuidfromFourChar()` — L43194
- `.MultiMeterIO_ListenerRunning()` — L43203
- `.MultiMeterIO_TransmittedData()` — L43209
- `.IsValidXml()` — L43212
- `.MultiMeterIO_ReceivedDataString()` — L43238
- `.parseJsonToken()` — L43301
- `.parseXMLElement()` — L43336

#### `MMIODirection` (type, L41324)

_No extracted members._

#### `MMIOFormat` (type, L41331)

_No extracted members._

#### `MMIOType` (type, L41339)

_No extracted members._

#### `MMIOTerminator` (type, L41348)

_No extracted members._

#### `clsMMIO` (type, L41358)

- `.OnDeserialized()` — L41398
- `.init()` — L41405
- `.refreshUdpEndpoint()` — L41538
- `.EnqueueOutbound()` — L41606
- `.DequeueOutbound()` — L41616
- `.StartConnection()` — L41644
- `.StopConnection()` — L41665
- `.SetVariable()` — L41669
- `.GetVariable()` — L41683
- `.DetermineType()` — L41707
- `.ConvertToType()` — L41757
- `.Variables()` — L41791
- `.VariableValueType()` — L41795
- `.RemoveVariable()` — L41826

#### `TcpListener` (type, L41831)

- `.Start()` — L41855
- `.Stop()` — L41864
- `.listen()` — L41889

#### `TcpClientHandler` (type, L42092)

- `.Start()` — L42116
- `.Stop()` — L42125
- `.Connect()` — L42148

#### `UdpListener` (type, L42335)

- `.Start()` — L42358
- `.Stop()` — L42367
- `.listen()` — L42375

#### `SerialPortHandler` (type, L42521)

- `.Start()` — L42552
- `.Stop()` — L42561
- `.Connect()` — L42572
- `.GetAvailableComPorts()` — L42738

#### `MiniSpec` (type, L43359)

- `.Init()` — L43417
- `.UpdateRXFilterCharacteristics()` — L43469
- `.UpdateTXFilterCharacteristics()` — L43491
- `.GetRXCharacteristic()` — L43519
- `.GetTXCharacteristic()` — L43527
- `.UsingAFilter()` — L43534
- `.UsingFilter()` — L43540
- `.StopUsingFilter()` — L43564
- `.OnSpectrumDetailsChanged()` — L43584
- `.OnAVGChanged()` — L43592
- `.OnFilterEdgesChanged()` — L43600
- `.OnMinNotchWidth()` — L43610
- `.OnTNFChanged()` — L43618
- `.OnMox()` — L43659
- `.Add()` — L43667
- `.GetMiniRX()` — L43693
- `.shutdownRX()` — L43714
- `.ShutdownAllRX()` — L43721
- `.OnPowerChanged()` — L43748
- `.OnCTUNChanged()` — L43760
- `.OnModeChanged()` — L43768
- `.OnCWPitchChanged()` — L43776
- `.OnSpectrumSettingsChanged()` — L43784
- `.OnHWSampleRateChanged()` — L43792
- `.OnCentreFrequency()` — L43800
- `.OnNotchChanged()` — L43811
- `.GetNotches()` — L43866
- `.GetNotch()` — L43878

#### `FilterCharacteristics` (type, L43367)

_No extracted members._

#### `Notch` (type, L43394)

_No extracted members._

#### `clsMiniSpec` (type, L43886)

- `.setupSpecDetails()` — L44118
- `.resetBuffers()` — L44150
- `.ClearData()` — L44162
- `.runDisplay()` — L44173
- `.attenuateData()` — L44251
- `.UpdateSpecSettings()` — L44265
- `.Shutdown()` — L44335
- `.rateLimitSetPan()` — L44385
- `.updatePan()` — L44406
- `.setPan()` — L44429
- `.zoom()` — L44473

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/MeterManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
