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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

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

- **`.Enqueue()`** — L387 — `public void Enqueue(float[] source, int sourceOffset, int count)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CopyTo()`** — L397 — `public void CopyTo(float[] destination, int destinationOffset, int count)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Peek()`** — L405 — `public float Peek(int index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Advance()`** — L410 — `public void Advance(int count)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ensureCapacity()`** — L436 — `private void ensureCapacity(int additionalCount)`
  Called by: `.Enqueue()` (same file)

#### `clsTCISensorManager` (type, L462)

- **`.clampIntervalMs()`** — L500 — `private static int clampIntervalMs(int intervalMs)`
  Called by: `.ConfigureRxSensors()` (same file), `.ConfigureTxSensors()` (same file)
- **`.ConfigureRxSensors()`** — L527 — `public void ConfigureRxSensors(bool enabled, int intervalMs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ConfigureTxSensors()`** — L544 — `public void ConfigureTxSensors(bool enabled, int intervalMs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RequiresRxChannelUpdate()`** — L554 — `public bool RequiresRxChannelUpdate(int receiver, int channel)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RequiresTxUpdate()`** — L566 — `public bool RequiresTxUpdate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SensorRequiresUpdate()`** — L574 — `public bool SensorRequiresUpdate(int receiver, Reading reading)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRxChannelReading()`** — L599 — `public void SetRxChannelReading(int receiver, int channel, double signal, double avg_signal, double peak_bin_signal)`
  Sets rx channel reading.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTxReadings()`** — L613 — `public void SetTxReadings(double micLevelDbm, double powerWatts, double peakPowerWatts, double swr)`
  Sets tx readings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetRxChannelReadingForSend()`** — L625 — `public bool TryGetRxChannelReadingForSend(int receiver, int channel, out double signal, out double avg_signal, out double peak_bin_signal)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ConsumeRxChannelReading()`** — L645 — `public void ConsumeRxChannelReading(int receiver, int channel)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetTxReadingsForSend()`** — L656 — `public bool TryGetTxReadingsForSend(out double micLevelDbm, out double powerWatts, out double peakPowerWatts, out double swr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ConsumeTxReadings()`** — L675 — `public void ConsumeTxReadings()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsRxReadingState` (type, L464)

_No extracted members._

#### `clsTxReadingState` (type, L472)

_No extracted members._

#### `TCPIPtciSocketListener` (type, L684)

- **`.ClickedOnSpot()`** — L826 — `public void ClickedOnSpot(string callsign, long frequency, int rx = -1, int chan = -1)`
  Called by: `.OnSpotClicked()` (same file), `.SendSpotSimulationClickToAll()` (same file)
- **`.ThetisFocusChange()`** — L837 — `public void ThetisFocusChange(bool focus)`
  Called by: `.OnThetisFocusChanged()` (same file)
- **`.RX2EnabledChange()`** — L842 — `public void RX2EnabledChange(bool enabled)`
  Called by: `.OnRX2EnabledChanged()` (same file)
- **`.HWSampleRateChange()`** — L848 — `public void HWSampleRateChange(int rx, int oldSampleRate, int newSampleRate)`
  Called by: `.OnHWSampleRateChanged()` (same file)
- **`.RequiresRxSensorUpdate()`** — L875 — `internal bool RequiresRxSensorUpdate(int receiver, int channel)`
  Called by: `.RequiresRxSensorUpdate()` (same file)
- **`.SensorRequiresUpdate()`** — L880 — `internal bool SensorRequiresUpdate(int receiver, Reading reading)`
  Called by: `.SensorRequiresUpdate()` (same file)
- **`.RequiresTxSensorUpdate()`** — L885 — `internal bool RequiresTxSensorUpdate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MeterReadingsChanged()`** — L890 — `internal void MeterReadingsChanged(int rx, bool tx, ref Dictionary<Reading, float> readings)`
  Called by: `.OnMeterReadingsChanged()` (same file)
- **`.MinimumRequiredRxSensorInterval()`** — L915 — `internal int MinimumRequiredRxSensorInterval()`
  Called by: `.MinimumRequiredRxSensorInterval()` (same file)
- **`.MinimumRequiredTxSensorInterval()`** — L920 — `internal int MinimumRequiredTxSensorInterval()`
  Called by: `.MinimumRequiredTxSensorInterval()` (same file)
- **`.getPublishedIQSampleRate()`** — L925 — `private int getPublishedIQSampleRate()`
  Returns published iqsample rate.
  Called by: `.sendInitialRadioState()` (same file), `.parseTextFrame()` (same file)
- **`.getPublishedIQSampleRateLocked()`** — L933 — `private int getPublishedIQSampleRateLocked()`
  Returns published iqsample rate locked.
  Called by: `.HWSampleRateChange()` (same file), `.getPublishedIQSampleRate()` (same file)
- **`.destroyRxAudioResamplerState()`** — L945 — `private unsafe void destroyRxAudioResamplerState(TCIRxAudioResamplerState state)`
  Called by: `.clearRxAudioStateForReceiver()` (same file), `.clearRxAudioStreamState()` (same file), `.resampleRxAudioSamples()` (same file)
- **`.clearRxAudioStateForReceiver()`** — L966 — `private void clearRxAudioStateForReceiver(int receiver)`
  Called by: `.handleAudioStart()` (same file)
- **`.clearRxAudioStreamState()`** — L981 — `private void clearRxAudioStreamState()`
  Called by: `.StopSocketListener()` (same file), `.handleAudioSampleRate()` (same file)
- **`.resampleRxAudioSamples()`** — L995 — `private unsafe int resampleRxAudioSamples(int receiver, int inputRate, int targetRate, float[] left, float[] right, int samples, out float[] leftOut, out float[] rightOut, out bool`
  Called by: `.PublishRxAudioSamples()` (same file)
- **`.DrivePowerChange()`** — L1088 — `public void DrivePowerChange(int rx, int newPower, bool tune)`
  Called by: `.OnDrivePowerChanged()` (same file)
- **`.TuneChange()`** — L1096 — `public void TuneChange(int rx, bool oldTune, bool newTune)`
  Called by: `.OnTuneChanged()` (same file)
- **`.SplitChange()`** — L1101 — `public void SplitChange(int rx, bool newSplit)`
  Called by: `.OnSplitChanged()` (same file)
- **`.MuteChanged()`** — L1107 — `public void MuteChanged(int rx, bool newState)`
  Called by: `.OnMuteChanged()` (same file)
- **`.AnfChanged()`** — L1113 — `public void AnfChanged(int rx, bool newState)`
  Called by: `.OnAnfChanged()` (same file)
- **`.RxAfGainChanged()`** — L1118 — `public void RxAfGainChanged(int rx, bool is_subrx, int gain)`
  Called by: `.OnRxAfGainChanged()` (same file)
- **`.CTUNChanged()`** — L1125 — `public void CTUNChanged(int rx, bool enabled)`
  Called by: `.OnCTUNChanged()` (same file)
- **`.VFOSyncChanged()`** — L1130 — `public void VFOSyncChanged(bool enabled)`
  Called by: `.OnVFOSyncChanged()` (same file)
- **`.FMDeviationChanged()`** — L1135 — `public void FMDeviationChanged(int rx, int deviationHz)`
  Called by: `.OnFMDeviationChanged()` (same file)
- **`.AGCModeChanged()`** — L1140 — `public void AGCModeChanged(int rx, AGCMode mode)`
  Called by: `.OnAGCModeChanged()` (same file)
- **`.AGCAutoChanged()`** — L1145 — `public void AGCAutoChanged(int rx, bool enabled)`
  Called by: `.OnAGCAutoModeChanged()` (same file)
- **`.TXProfileChanged()`** — L1150 — `public void TXProfileChanged(string profile)`
  Called by: `.OnTXProfileChanged()` (same file)
- **`.TXProfilesChanged()`** — L1155 — `public void TXProfilesChanged()`
  Called by: `.OnTXProfilesChanged()` (same file)
- **`.CalibrationChanged()`** — L1160 — `public void CalibrationChanged(int rx)`
  Called by: `.sendInitialRadioState()` (same file), `.handleCalibration()` (same file), `.OnCalibrationChanged()` (same file)
- **`.MONChanged()`** — L1178 — `public void MONChanged(bool newState)`
  Called by: `.OnMONChanged()` (same file)
- **`.MONVolumeChanged()`** — L1183 — `public void MONVolumeChanged(int newVolume)`
  Called by: `.OnMONVolumeChanged()` (same file)
- **`.VolumeChanged()`** — L1188 — `public void VolumeChanged(int newVolume)`
  Called by: `.OnVolumeChanged()` (same file)
- **`.BalanceChanged()`** — L1193 — `public void BalanceChanged(int rx, bool is_subrx, int newBalance)`
  Called by: `.OnBalanceChanged()` (same file)
- **`.RxStepAttChanged()`** — L1200 — `public void RxStepAttChanged(int rx, int attenuation)`
  Called by: `.OnAttenuatorDataChanged()` (same file)
- **`.RxPreampAttChanged()`** — L1205 — `public void RxPreampAttChanged(int rx, PreampMode preamp_mode)`
  Called by: `.OnPreampModeChanged()` (same file)
- **`.RxStepAttEnabledChanged()`** — L1212 — `public void RxStepAttEnabledChanged(int rx, bool enabled)`
  Called by: `.OnStepAttEnabledChanged()` (same file)
- **`.AGCGainChanged()`** — L1217 — `public void AGCGainChanged(int rx, int newGain)`
  Called by: `.OnAGCGainChanged()` (same file)
- **`.RITChanged()`** — L1222 — `public void RITChanged(bool newState)`
  Called by: `.OnRITChanged()` (same file)
- **`.XITChanged()`** — L1228 — `public void XITChanged(bool newState)`
  Called by: `.OnXITChanged()` (same file)
- **`.RITValueChanged()`** — L1234 — `public void RITValueChanged(int newValue)`
  Called by: `.OnRITValueChanged()` (same file)
- **`.XITValueChanged()`** — L1240 — `public void XITValueChanged(int newValue)`
  Called by: `.OnXITValueChanged()` (same file)
- **`.CwMacrosSpeedChanged()`** — L1246 — `public void CwMacrosSpeedChanged(int newSpeed)`
  Called by: `.OnCwMacrosSpeedChanged()` (same file)
- **`.CwMacrosDelayChanged()`** — L1251 — `public void CwMacrosDelayChanged(int newDelay)`
  Called by: `.OnCwMacrosDelayChanged()` (same file)
- **`.CwKeyerSpeedChanged()`** — L1256 — `public void CwKeyerSpeedChanged(int newSpeed)`
  Called by: `.OnCwKeyerSpeedChanged()` (same file)
- **`.CwMacrosEmpty()`** — L1261 — `public void CwMacrosEmpty(int rx)`
  Called by: `.OnCwMacrosEmpty()` (same file)
- **`.CwCallsignSent()`** — L1266 — `public void CwCallsignSent(string callsign)`
  Called by: `.OnCwCallsignSent()` (same file)
- **`.NBChanged()`** — L1271 — `public void NBChanged(int rx, int newNB)`
  Called by: `.OnNbChanged()` (same file)
- **`.NRChanged()`** — L1278 — `public void NRChanged(int rx, int newNR)`
  Called by: `.OnNrChanged()` (same file)
- **`.BinChanged()`** — L1285 — `public void BinChanged(int rx, bool newState)`
  Called by: `.OnBinChanged()` (same file)
- **`.LockChanged()`** — L1290 — `public void LockChanged(int rx, bool newState)`
  Called by: `.OnVfoALockChanged()` (same file), `.OnVfoBLockChanged()` (same file)
- **`.VFOLocksChanged()`** — L1295 — `public void VFOLocksChanged()`
  Called by: `.OnVfoALockChanged()` (same file), `.OnVfoBLockChanged()` (same file)
- **`.SqlChanged()`** — L1300 — `public void SqlChanged(int rx, SquelchState newState)`
  Called by: `.OnSqlChanged()` (same file)
- **`.SqlLevelChanged()`** — L1305 — `public void SqlLevelChanged(int rx, int newValue)`
  Called by: `.OnSqlLevelChanged()` (same file)
- **`.ApfChanged()`** — L1310 — `public void ApfChanged(int rx, bool newState)`
  Called by: `.OnApfChanged()` (same file)
- **`.NfChanged()`** — L1315 — `public void NfChanged(bool newState)`
  Called by: `.OnTnfChanged()` (same file)
- **`.DiglOffsetChanged()`** — L1321 — `public void DiglOffsetChanged(int newValue)`
  Called by: `.OnDiglOffsetChanged()` (same file)
- **`.DiguOffsetChanged()`** — L1326 — `public void DiguOffsetChanged(int newValue)`
  Called by: `.OnDiguOffsetChanged()` (same file)
- **`.TXFrequencyChanged()`** — L1331 — `public void TXFrequencyChanged(long new_frequency, Band new_band, bool rx2_enabled, bool tx_vfob)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.limitList()`** — L1336 — `private void limitList()`
  Called by: `.vfoFrequencyChange()` (same file), `.centreFrequencyChange()` (same file), `.txFrequencyChange()` (same file)
- **`.VFOdata()`** — L1348 — `private async void VFOdata()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.vfoFrequencyChange()`** — L1417 — `private void vfoFrequencyChange(VFOData vfod)`
  Called by: `.VFOcallback()` (same file), `.VFOChange()` (same file)
- **`.centreFrequencyChange()`** — L1426 — `private void centreFrequencyChange(VFOData vfod)`
  Called by: `.Centrecallback()` (same file), `.CentreChange()` (same file)
- **`.txFrequencyChange()`** — L1435 — `private void txFrequencyChange(VFOData vfod)`
  Called by: `.OnTXFrequencyChanged()` (same file)
- **`.MoxChange()`** — L1444 — `public void MoxChange(int rx, bool oldMox, bool newMox)`
  Called by: `.OnMoxChangeHandler()` (same file)
- **`.ModeChange()`** — L1473 — `public void ModeChange(int rx, DSPMode oldMode, DSPMode newMode, Band oldBand, Band newBand)`
  Called by: `.OnModeChangeHandler()` (same file)
- **`.BandChange()`** — L1478 — `public void BandChange(int rx, Band oldBand, Band newBand)`
  Called by: `.OnBandChangeHandler()` (same file)
- **`.FilterChange()`** — L1485 — `public void FilterChange(int rx, Filter oldFilter, Filter newFilter, Band band, int low, int high)`
  Called by: `.OnFilterChanged()` (same file)
- **`.FilterEdgesChange()`** — L1490 — `public void FilterEdgesChange(int rx, Filter filter, Band band, int low, int high)`
  Called by: `.OnFilterEdgesChanged()` (same file)
- **`.TXFilterBandChanged()`** — L1495 — `public void TXFilterBandChanged(int low, int high)`
  Called by: `.OnTXFiltersChanged()` (same file)
- **`.PowerChange()`** — L1500 — `public void PowerChange(bool oldPower, bool newPower)`
  Called by: `.OnPowerChangeHander()` (same file)
- **`.StartSocketListener()`** — L1507 — `public void StartSocketListener()`
  Starts socket listener.
  Called by: `.ServerThreadStart()` (same file)
- **`.getFrameFromString()`** — L1538 — `private static byte[] getFrameFromString(string Message, EOpcodeType Opcode = EOpcodeType.Text)`
  Returns frame from string.
  Called by: `.sendPingFrame()` (same file), `.sendPongFrame()` (same file), `.sendTextFrame()` (same file), `.sendCloseFrame()` (same file)
- **`.GetFrameFromBytes()`** — L1597 — `private static byte[] GetFrameFromBytes(byte[] payload, EOpcodeType opcode = EOpcodeType.Binary)`
  Returns frame from bytes.
  Called by: `.sendBinaryFrame()` (same file)
- **`.getCoalescedTextFrameKey()`** — L1638 — `private static string getCoalescedTextFrameKey(string message)`
  Returns coalesced text frame key.
  Called by: `.sendTextFrame()` (same file)
- **`.hasPendingOutboundFramesLocked()`** — L1680 — `private bool hasPendingOutboundFramesLocked()`
  Called by: `.flushOutboundFrames()` (same file), `.SendThreadProc()` (same file)
- **`.tryDequeueNextOutboundFrameLocked()`** — L1688 — `private bool tryDequeueNextOutboundFrameLocked(out TCIOutboundFrame frame)`
  Called by: `.SendThreadProc()` (same file)
- **`.clearOutboundFrames()`** — L1724 — `private void clearOutboundFrames()`
  Called by: `.StopSocketListener()` (same file)
- **`.flushOutboundFrames()`** — L1737 — `private void flushOutboundFrames(int timeoutMs)`
  Called by: `.StopSocketListener()` (same file)
- **`.enqueueOutboundFrame()`** — L1752 — `private void enqueueOutboundFrame(byte[] frameBytes, string logText, TCIOutboundPriority priority, string coalescedKey = null)`
  Called by: `.sendPingFrame()` (same file), `.sendPongFrame()` (same file), `.sendTextFrame()` (same file), `.sendBinaryFrame()` (same file), `.sendCloseFrame()` (same file)
- **`.SendThreadProc()`** — L1791 — `private void SendThreadProc()`
  Sends thread proc.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.abortSocketTransport()`** — L1836 — `private void abortSocketTransport()`
  Called by: `.SendThreadProc()` (same file)
- **`.isSocketReadTimeout()`** — L1865 — `private static bool isSocketReadTimeout(IOException ex)`
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.upgradeToWebSocket()`** — L1876 — `private bool upgradeToWebSocket(string msg)`
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.sendStart()`** — L1911 — `private void sendStart()`
  Called by: `.sendStartStop()` (same file)
- **`.sendStop()`** — L1915 — `private void sendStop()`
  Called by: `.sendStartStop()` (same file), `.StopSocketListener()` (same file)
- **`.sendSplit()`** — L1919 — `private void sendSplit(int rx, bool bSplit)`
  Called by: `.SplitChange()` (same file), `.sendInitialRadioState()` (same file), `.handleSplitEnableMessage()` (same file)
- **`.sendRITEnable()`** — L1924 — `private void sendRITEnable(int rx, bool enabled)`
  Called by: `.RITChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRITEnableMessage()` (same file)
- **`.sendXITEnable()`** — L1929 — `private void sendXITEnable(int rx, bool enabled)`
  Called by: `.XITChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleXITEnableMessage()` (same file)
- **`.sendRITOffset()`** — L1934 — `private void sendRITOffset(int rx, int offset)`
  Called by: `.RITValueChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRITOffsetMessage()` (same file)
- **`.sendXITOffset()`** — L1939 — `private void sendXITOffset(int rx, int offset)`
  Called by: `.XITValueChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleXITOffsetMessage()` (same file)
- **`.sendRxBinEnable()`** — L1944 — `private void sendRxBinEnable(int rx, bool enabled)`
  Called by: `.BinChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxBinEnable()` (same file)
- **`.sendRxApfEnable()`** — L1949 — `private void sendRxApfEnable(int rx, bool enabled)`
  Called by: `.ApfChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxApfEnable()` (same file)
- **`.sendRxNfEnable()`** — L1954 — `private void sendRxNfEnable(int rx, bool enabled)`
  Called by: `.NfChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxNfEnable()` (same file)
- **`.sendLock()`** — L1959 — `private void sendLock(int rx, bool enabled)`
  Called by: `.LockChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleLock()` (same file)
- **`.sendVFOLock()`** — L1964 — `private void sendVFOLock(int rx, int chan, bool enabled)`
  Called by: `.sendAllVFOLocks()` (same file), `.handleVFOLock()` (same file)
- **`.sendSqlEnable()`** — L1969 — `private void sendSqlEnable(int rx, bool enabled)`
  Called by: `.SqlChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleSqlEnable()` (same file)
- **`.sendSqlLevel()`** — L1974 — `private void sendSqlLevel(int rx, int level)`
  Called by: `.SqlLevelChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleSqlLevel()` (same file)
- **`.sendCwMacrosSpeed()`** — L1979 — `private void sendCwMacrosSpeed(int speed)`
  Called by: `.CwMacrosSpeedChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleCwMacrosSpeed()` (same file)
- **`.sendCwMacrosDelay()`** — L1983 — `private void sendCwMacrosDelay(int delayMs)`
  Called by: `.CwMacrosDelayChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleCwMacrosDelay()` (same file)
- **`.sendCwKeyerSpeed()`** — L1987 — `private void sendCwKeyerSpeed(int speed)`
  Called by: `.CwKeyerSpeedChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleCwKeyerSpeed()` (same file)
- **`.sendCwMacrosEmpty()`** — L1991 — `private void sendCwMacrosEmpty(int rx)`
  Called by: `.CwMacrosEmpty()` (same file)
- **`.sendCallsignSend()`** — L1995 — `private void sendCallsignSend(string callsign)`
  Called by: `.CwCallsignSent()` (same file)
- **`.tryGetVFOLockState()`** — L1999 — `private bool tryGetVFOLockState(int rx, int chan, out bool enabled)`
  Called by: `.sendAllVFOLocks()` (same file), `.handleVFOLock()` (same file)
- **`.trySetVFOLockState()`** — L2034 — `private bool trySetVFOLockState(int rx, int chan, bool enabled)`
  Called by: `.handleVFOLock()` (same file)
- **`.sendAllVFOLocks()`** — L2070 — `private void sendAllVFOLocks()`
  Called by: `.VFOLocksChanged()` (same file), `.sendInitialRadioState()` (same file)
- **`.sendDiglOffset()`** — L2089 — `private void sendDiglOffset(int offset)`
  Called by: `.DiglOffsetChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleDiglOffset()` (same file)
- **`.sendDiguOffset()`** — L2094 — `private void sendDiguOffset(int offset)`
  Called by: `.DiguOffsetChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleDiguOffset()` (same file)
- **`.sendVFO()`** — L2099 — `private void sendVFO(int rx, int chan, long vfo = -1)`
  Called by: `.VFOdata()` (same file), `.sendInitialRadioState()` (same file)
- **`.sendIF()`** — L2134 — `private void sendIF(int rx, int chan, int offset = -999999999)`
  Called by: `.VFOdata()` (same file), `.sendInitialRadioState()` (same file), `.handleIF()` (same file)
- **`.sendMOX()`** — L2159 — `private void sendMOX(int rx, bool mox, bool signalTCI = false)`
  Called by: `.MoxChange()` (same file), `.sendInitialRadioState()` (same file), `.handleTrxMessage()` (same file)
- **`.sendAudioStartStop()`** — L2170 — `private void sendAudioStartStop(int receiver, bool enable)`
  Called by: `.handleAudioStart()` (same file)
- **`.sendMode()`** — L2174 — `private void sendMode(int rx, DSPMode mode = DSPMode.FIRST)`
  Called by: `.ModeChange()` (same file), `.sendInitialRadioState()` (same file), `.handleModulationMessage()` (same file)
- **`.sendMute()`** — L2196 — `private void sendMute(bool mute)`
  Called by: `.MuteChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleMute()` (same file)
- **`.sendMuteRX()`** — L2201 — `private void sendMuteRX(int rx, bool mute)`
  Called by: `.MuteChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleMuteRX()` (same file)
- **`.sendMONEnable()`** — L2206 — `private void sendMONEnable(bool enable)`
  Called by: `.MONChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleMONEnable()` (same file)
- **`.sendVolume()`** — L2211 — `private void sendVolume(double volume)`
  Called by: `.VolumeChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleVolume()` (same file)
- **`.sendMONVolume()`** — L2218 — `private void sendMONVolume(double volume)`
  Called by: `.MONVolumeChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleMONVolume()` (same file)
- **`.sendRxBalance()`** — L2225 — `private void sendRxBalance(int rx, int chan, double balance)`
  Called by: `.BalanceChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxBalance()` (same file)
- **`.sendRxStepAttEx()`** — L2230 — `private void sendRxStepAttEx(int rx, int attenuation)`
  Called by: `.RxStepAttChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxStepAttEx()` (same file)
- **`.sendRxPreampAttEx()`** — L2235 — `private void sendRxPreampAttEx(int rx, int attenuation)`
  Called by: `.RxPreampAttChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxPreampAttEx()` (same file)
- **`.sendRxStepAttEnabledEx()`** — L2240 — `private void sendRxStepAttEnabledEx(int rx, bool enabled)`
  Called by: `.RxStepAttEnabledChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxStepAttEnabledEx()` (same file)
- **`.sendVFOSyncEx()`** — L2245 — `private void sendVFOSyncEx(bool enabled)`
  Called by: `.VFOSyncChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleVfoSyncEx()` (same file)
- **`.sendFMDeviationEx()`** — L2250 — `private void sendFMDeviationEx(int rx, int deviationHz)`
  Called by: `.FMDeviationChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleFMDeviationEx()` (same file)
- **`.sendAgcAutoEx()`** — L2255 — `private void sendAgcAutoEx(int rx, bool enabled)`
  Called by: `.AGCAutoChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleAgcAutoEx()` (same file)
- **`.agcModeToTciMode()`** — L2260 — `private string agcModeToTciMode(AGCMode mode)`
  Called by: `.sendAgcMode()` (same file)
- **`.tciModeToAgcMode()`** — L2280 — `private AGCMode tciModeToAgcMode(string mode)`
  Called by: `.handleAgcMode()` (same file)
- **`.sendAgcMode()`** — L2304 — `private void sendAgcMode(int rx, AGCMode mode)`
  Called by: `.AGCModeChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleAgcMode()` (same file)
- **`.sendAgcGain()`** — L2309 — `private void sendAgcGain(int rx, int gain)`
  Called by: `.AGCGainChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleAgcGain()` (same file)
- **`.sendTXFrequencyChanged()`** — L2314 — `private void sendTXFrequencyChanged(long new_frequency, Band new_band, bool rx2_enabled, bool tx_vfob)`
  Called by: `.TXFrequencyChanged()` (same file), `.VFOdata()` (same file), `.sendInitialRadioState()` (same file)
- **`.sendTunePower()`** — L2328 — `private void sendTunePower(int rx, int drive)`
  Called by: `.DrivePowerChange()` (same file), `.handleTuneDrive()` (same file)
- **`.sendDrivePower()`** — L2335 — `private void sendDrivePower(int rx, int drive)`
  Called by: `.DrivePowerChange()` (same file), `.handleDrive()` (same file)
- **`.sendTune()`** — L2342 — `private void sendTune(int rx, bool tune)`
  Called by: `.TuneChange()` (same file), `.sendInitialRadioState()` (same file), `.handleTune()` (same file)
- **`.sendRXEnable()`** — L2347 — `private void sendRXEnable(int rx, bool enable)`
  Called by: `.RX2EnabledChange()` (same file), `.sendInitialRadioState()` (same file), `.handleRXEnable()` (same file)
- **`.sendTXEnable()`** — L2352 — `private void sendTXEnable(int rx, bool bEnable)`
  Called by: `.RX2EnabledChange()` (same file), `.MoxChange()` (same file), `.BandChange()` (same file), `.sendInitialRadioState()` (same file)
- **`.sendVFOLimits()`** — L2357 — `private void sendVFOLimits(int low, int high)`
  Called by: `.sendInitialisationData()` (same file)
- **`.sendAppFocus()`** — L2362 — `private void sendAppFocus(bool focus)`
  Called by: `.ThetisFocusChange()` (same file)
- **`.sendIFLimits()`** — L2367 — `private void sendIFLimits(int low, int high)`
  Called by: `.HWSampleRateChange()` (same file), `.sendInitialisationData()` (same file)
- **`.sendClickedOnSpot()`** — L2372 — `private void sendClickedOnSpot(string callsign, long frequency)`
  Called by: `.ClickedOnSpot()` (same file)
- **`.sendClickedOnSpotRX()`** — L2377 — `private void sendClickedOnSpotRX(int rx, int chan, string callsign, long frequency)`
  Called by: `.ClickedOnSpot()` (same file)
- **`.sendRxSensors()`** — L2382 — `private void sendRxSensors(int rx, double levelDbm)`
  Called by: `.RxSensorsTimerCallback()` (same file)
- **`.sendRxChannelSensors()`** — L2386 — `private void sendRxChannelSensors(int rx, int channel, double levelDbm, double avgLevelDbm, double peakBinDbm)`
  Called by: `.RxSensorsTimerCallback()` (same file)
- **`.sendTxSensors()`** — L2391 — `private void sendTxSensors(int rx, double micLevelDbm, double rmsPowerWatts, double peakPowerWatts, double swr)`
  Called by: `.TxSensorsTimerCallback()` (same file)
- **`.sendDDS()`** — L2402 — `private void sendDDS(int rx, long ddsFreq = -1)`
  Called by: `.VFOdata()` (same file), `.sendInitialRadioState()` (same file), `.handleDDS()` (same file)
- **`.sendFilterBand()`** — L2417 — `private void sendFilterBand(int rx, int low, int high)`
  Called by: `.FilterChange()` (same file), `.FilterEdgesChange()` (same file), `.sendInitialRadioState()` (same file), `.handleRxFilterBand()` (same file)
- **`.normalizeTXFilterBandForSet()`** — L2422 — `private void normalizeTXFilterBandForSet(ref int low, ref int high)`
  Called by: `.normalizeTXFilterBandForSend()` (same file), `.handleTXFilterBandEx()` (same file)
- **`.normalizeTXFilterBandForSend()`** — L2437 — `private void normalizeTXFilterBandForSend(ref int low, ref int high)`
  Called by: `.sendTXFilterBandEx()` (same file)
- **`.sendTXFilterBandEx()`** — L2443 — `private void sendTXFilterBandEx(int low, int high)`
  Called by: `.TXFilterBandChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleTXFilterBandEx()` (same file)
- **`.sendStartStop()`** — L2449 — `private void sendStartStop(bool bPower)`
  Called by: `.PowerChange()` (same file), `.sendInitialRadioState()` (same file)
- **`.preampModeToAttenuation()`** — L2457 — `private int preampModeToAttenuation(PreampMode mode)`
  Called by: `.RxPreampAttChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxPreampAttEx()` (same file)
- **`.sendInitialRadioState()`** — L2486 — `private void sendInitialRadioState()`
  Called by: `.sendInitialisationData()` (same file)
- **`.sendInitialisationData()`** — L2662 — `private void sendInitialisationData()`
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.setRxSensorsEnabled()`** — L2704 — `private void setRxSensorsEnabled(bool enabled, int intervalMs, bool fireImmediately)`
  Sets rx sensors enabled.
  Called by: `.handleRxSensorsEnable()` (same file)
- **`.setTxSensorsEnabled()`** — L2719 — `private void setTxSensorsEnabled(bool enabled, int intervalMs, bool fireImmediately)`
  Sets tx sensors enabled.
  Called by: `.handleTxSensorsEnable()` (same file)
- **`.RxSensorsTimerCallback()`** — L2734 — `private void RxSensorsTimerCallback(object state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TxSensorsTimerCallback()`** — L2769 — `private void TxSensorsTimerCallback(object state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.findEndOfHeader()`** — L2781 — `private int findEndOfHeader(byte[] bytes)`
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.SocketListenerThreadStart()`** — L2798 — `private void SocketListenerThreadStart()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.notifyServerDisconnected()`** — L2915 — `private void notifyServerDisconnected(TCPIPtciServer server = null)`
  Called by: `.SocketListenerThreadStart()` (same file), `.StopSocketListener()` (same file)
- **`.sendPingFrame()`** — L2926 — `private void sendPingFrame(string sMsg)`
  Called by: `.PingFrameTimer()` (same file)
- **`.sendPongFrame()`** — L2944 — `private void sendPongFrame(string sMsg)`
  Called by: `.ParseReceiveBuffer()` (same file)
- **`.sendTextFrame()`** — L2962 — `private void sendTextFrame(string sMsg)`
  Called by: `.sendStart()` (same file), `.sendStop()` (same file), `.sendSplit()` (same file), `.sendRITEnable()` (same file), `.sendXITEnable()` (same file), `.sendRITOffset()` (same file) — and 69 more
- **`.sendBinaryFrame()`** — L2981 — `private void sendBinaryFrame(byte[] payload)`
  Called by: `.PublishIQSamples()` (same file), `.PublishRxAudioSamples()` (same file), `.SendTxChrono()` (same file)
- **`.sendCloseFrame()`** — L3000 — `private void sendCloseFrame()`
  Called by: `.StopSocketListener()` (same file), `.ParseReceiveBuffer()` (same file)
- **`.StopSocketListener()`** — L3014 — `public void StopSocketListener()`
  Stops socket listener.
  Called by: `.StopAllSocketListers()` (same file), `.PurgingThreadStart()` (same file)
- **`.IsMarkedForDeletion()`** — L3118 — `public bool IsMarkedForDeletion()`
  Called by: `.PurgingThreadStart()` (same file)
- **`.IsDisconnected()`** — L3122 — `public bool IsDisconnected()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetFrameLength()`** — L3126 — `private int GetFrameLength(Byte[] bytes)`
  Returns frame length.
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.ParseReceiveBuffer()`** — L3156 — `private void ParseReceiveBuffer(Byte[] bytes)`
  Parses receive buffer.
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.handleSetInFocus()`** — L3223 — `private void handleSetInFocus()`
  Called by: `.parseTextFrame()` (same file)
- **`.handleStart()`** — L3227 — `private void handleStart()`
  Called by: `.parseTextFrame()` (same file)
- **`.handleStop()`** — L3232 — `private void handleStop()`
  Called by: `.parseTextFrame()` (same file)
- **`.handleSpotClear()`** — L3237 — `private void handleSpotClear()`
  Called by: `.parseTextFrame()` (same file)
- **`.handleSplitEnableMessage()`** — L3241 — `private void handleSplitEnableMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRITEnableMessage()`** — L3278 — `private void handleRITEnableMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleXITEnableMessage()`** — L3294 — `private void handleXITEnableMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRITOffsetMessage()`** — L3310 — `private void handleRITOffsetMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleXITOffsetMessage()`** — L3326 — `private void handleXITOffsetMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxBinEnable()`** — L3343 — `private void handleRxBinEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxApfEnable()`** — L3359 — `private void handleRxApfEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxNfEnable()`** — L3384 — `private void handleRxNfEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleLock()`** — L3400 — `private void handleLock(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleVFOLock()`** — L3419 — `private void handleVFOLock(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleSqlEnable()`** — L3436 — `private void handleSqlEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleSqlLevel()`** — L3452 — `private void handleSqlLevel(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleDiglOffset()`** — L3469 — `private void handleDiglOffset(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleDiguOffset()`** — L3481 — `private void handleDiguOffset(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwMacrosSpeed()`** — L3493 — `private void handleCwMacrosSpeed(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwMacrosDelay()`** — L3505 — `private void handleCwMacrosDelay(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwKeyerSpeed()`** — L3517 — `private void handleCwKeyerSpeed(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwMacrosSpeedUp()`** — L3529 — `private void handleCwMacrosSpeedUp(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwMacrosSpeedDown()`** — L3535 — `private void handleCwMacrosSpeedDown(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwMacros()`** — L3541 — `private void handleCwMacros(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwTerminal()`** — L3550 — `private void handleCwTerminal(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwMsg()`** — L3559 — `private void handleCwMsg(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleCwMacrosStop()`** — L3578 — `private void handleCwMacrosStop()`
  Called by: `.parseTextFrame()` (same file)
- **`.handleKeyer()`** — L3582 — `private void handleKeyer(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleTrxMessage()`** — L3594 — `private void handleTrxMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.shouldIgnoreTrxForCurrentCwBreakIn()`** — L3696 — `private bool shouldIgnoreTrxForCurrentCwBreakIn()`
  Called by: `.handleTrxMessage()` (same file)
- **`.handleIF()`** — L3710 — `private void handleIF(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleDDS()`** — L3808 — `private void handleDDS(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleVFOMessage()`** — L3859 — `private void handleVFOMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleModulationMessage()`** — L3972 — `private void handleModulationMessage(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleDeleteSpot()`** — L4076 — `private void handleDeleteSpot(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.lineOutEnable()`** — L4085 — `private void lineOutEnable(int vac_number, bool enable)`
  line out to switch vacs
  Called by: `.handleLineOutStart()` (same file), `.handleLineOutStop()` (same file)
- **`.handleLineOutStart()`** — L4112 — `private void handleLineOutStart(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleLineOutStop()`** — L4125 — `private void handleLineOutStop(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleDrive()`** — L4138 — `private void handleDrive(string[] args)`
  Called by: `.sendInitialRadioState()` (same file), `.parseTextFrame()` (same file)
- **`.handleTuneDrive()`** — L4165 — `private void handleTuneDrive(string[] args)`
  Called by: `.sendInitialRadioState()` (same file), `.parseTextFrame()` (same file)
- **`.handleMute()`** — L4214 — `private void handleMute(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleMuteRX()`** — L4232 — `private void handleMuteRX(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleMONEnable()`** — L4256 — `private void handleMONEnable(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.linearToDbVolume()`** — L4273 — `private double linearToDbVolume(int volume)`
  Called by: `.MONVolumeChanged()` (same file), `.VolumeChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleMONVolume()` (same file), `.handleVolume()` (same file)
- **`.dbToLinearVolume()`** — L4284 — `private int dbToLinearVolume(double dBLevel)`
  Called by: `.handleMONVolume()` (same file), `.handleVolume()` (same file)
- **`.handleMONVolume()`** — L4296 — `private void handleMONVolume(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleVolume()`** — L4313 — `private void handleVolume(string[] args, bool hasArgs = true)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleSpotSimulateClick()`** — L4325 — `private void handleSpotSimulateClick(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleSpot()`** — L4339 — `private void handleSpot(string[] args, bool is_json, string msg)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleTune()`** — L4506 — `private void handleTune(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxFilterBand()`** — L4529 — `private void handleRxFilterBand(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleTXFilterBandEx()`** — L4576 — `private void handleTXFilterBandEx(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRXEnable()`** — L4595 — `private void handleRXEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxSensorsEnable()`** — L4631 — `private void handleRxSensorsEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleTxSensorsEnable()`** — L4642 — `private void handleTxSensorsEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.sendNREnable()`** — L4654 — `private void sendNREnable(int rx, bool enabled, bool is_extended, int nr)`
  Called by: `.NRChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleNREnable()` (same file)
- **`.sendNBEnable()`** — L4664 — `private void sendNBEnable(int rx, bool enabled, bool is_extended, int nb)`
  Called by: `.NBChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxNBEnable()` (same file)
- **`.handleNREnable()`** — L4674 — `private void handleNREnable(string[] args, bool is_extended)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxNBEnable()`** — L4707 — `private void handleRxNBEnable(string[] args, bool is_extended)`
  Called by: `.parseTextFrame()` (same file)
- **`.sendAnfEnable()`** — L4740 — `private void sendAnfEnable(int rx, bool enabled)`
  Called by: `.AnfChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleAnfEnable()` (same file)
- **`.handleAnfEnable()`** — L4746 — `private void handleAnfEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.dbToAudioGain()`** — L4767 — `private double dbToAudioGain(double db)`
  Called by: `.handleRxVolume()` (same file)
- **`.audioGainToDb()`** — L4778 — `private double audioGainToDb(double gain)`
  Called by: `.RxAfGainChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxVolume()` (same file)
- **`.sendRxVolume()`** — L4788 — `private void sendRxVolume(int rx, int chan, double volume)`
  Called by: `.RxAfGainChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleRxVolume()` (same file)
- **`.handleRxVolume()`** — L4794 — `private void handleRxVolume(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxBalance()`** — L4856 — `private void handleRxBalance(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxStepAttEnabledEx()`** — L4883 — `private void handleRxStepAttEnabledEx(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxStepAttEx()`** — L4905 — `private void handleRxStepAttEx(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRxPreampAttEx()`** — L4927 — `private void handleRxPreampAttEx(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleVfoSyncEx()`** — L4945 — `private void handleVfoSyncEx(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleVfoSwapEx()`** — L4958 — `private void handleVfoSwapEx()`
  Called by: `.parseTextFrame()` (same file)
- **`.handleFMDeviationEx()`** — L4962 — `private void handleFMDeviationEx(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleAgcAutoEx()`** — L4980 — `private void handleAgcAutoEx(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleAgcMode()`** — L4996 — `private void handleAgcMode(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleAgcGain()`** — L5011 — `private void handleAgcGain(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.sendCTUN()`** — L5028 — `private void sendCTUN(int rx, bool enabled)`
  Called by: `.CTUNChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleCTUN()` (same file)
- **`.handleCTUN()`** — L5034 — `private void handleCTUN(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.sendTXProfile()`** — L5053 — `private void sendTXProfile(string prof)`
  Called by: `.TXProfileChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleTXProfile()` (same file)
- **`.sendTXProfiles()`** — L5059 — `private void sendTXProfiles()`
  Called by: `.TXProfilesChanged()` (same file), `.sendInitialRadioState()` (same file), `.handleTXProfiles()` (same file)
- **`.handleTXProfile()`** — L5070 — `private void handleTXProfile(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleTXProfiles()`** — L5086 — `private void handleTXProfiles()`
  Called by: `.parseTextFrame()` (same file)
- **`.handleShutdown()`** — L5090 — `private void handleShutdown()`
  Called by: `.parseTextFrame()` (same file)
- **`.sendCalibration()`** — L5104 — `private void sendCalibration(int rx, float meter, float display, float xvtr, float six_meter, float tx_display_offset)`
  Called by: `.CalibrationChanged()` (same file)
- **`.handleCalibration()`** — L5114 — `private void handleCalibration(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleRunCatCommand()`** — L5122 — `private void handleRunCatCommand(string msg)`
  Called by: `.parseTextFrame()` (same file)
- **`.splitTextCommands()`** — L5153 — `private List<string> splitTextCommands(string msg)`
  Called by: `.parseTextFrame()` (same file)
- **`.parseTextFrame()`** — L5259 — `private void parseTextFrame(string msg)`
  Called by: `.ParseReceiveBuffer()` (same file)
- **`.getDefaultAudioStreamSamples()`** — L5606 — `private static int getDefaultAudioStreamSamples(int sampleRate)`
  Returns default audio stream samples.
  Called by: `.handleAudioSampleRate()` (same file)
- **`.getBytesPerSample()`** — L5622 — `private static int getBytesPerSample(TCISampleType sampleType)`
  Returns bytes per sample.
  Called by: `.encodeSamples()` (same file), `.handleBinaryFrame()` (same file)
- **`.writeUInt32()`** — L5637 — `private static void writeUInt32(byte[] buffer, int offset, uint value)`
  Called by: `.buildStreamPayload()` (same file)
- **`.buildStreamPayload()`** — L5645 — `private byte[] buildStreamPayload(int receiver, int sampleRate, TCISampleType sampleType, int length, TCIStreamType streamType, int channels, byte[] samplePayload)`
  Called by: `.PublishIQSamples()` (same file), `.PublishRxAudioSamples()` (same file), `.SendTxChrono()` (same file)
- **`.encodeSamples()`** — L5669 — `private byte[] encodeSamples(float[] samples, TCISampleType sampleType)`
  Called by: `.PublishIQSamples()` (same file), `.PublishRxAudioSamples()` (same file)
- **`.decodeSamples()`** — L5712 — `private static float[] decodeSamples(byte[] payload, int dataOffset, int length, TCISampleType sampleType)`
  Called by: `.handleBinaryFrame()` (same file)
- **`.convertStreamSamplesToComplex()`** — L5744 — `private static double[] convertStreamSamplesToComplex(float[] samples, int channels)`
  Called by: `.handleBinaryFrame()` (same file)
- **`.sendIQSampleRate()`** — L5769 — `private void sendIQSampleRate(int sampleRate)`
  Called by: `.HWSampleRateChange()` (same file), `.sendInitialRadioState()` (same file), `.parseTextFrame()` (same file), `.handleIQSampleRate()` (same file)
- **`.sendAudioSampleRate()`** — L5777 — `private void sendAudioSampleRate(int sampleRate)`
  Called by: `.sendInitialRadioState()` (same file), `.parseTextFrame()` (same file), `.handleAudioSampleRate()` (same file)
- **`.sendAudioStreamSampleType()`** — L5782 — `private void sendAudioStreamSampleType(TCISampleType sampleType)`
  Called by: `.sendInitialRadioState()` (same file), `.handleAudioStreamSampleType()` (same file)
- **`.sendAudioStreamChannels()`** — L5787 — `private void sendAudioStreamChannels(int channels)`
  Called by: `.sendInitialRadioState()` (same file), `.handleAudioStreamChannels()` (same file)
- **`.sendAudioStreamSamples()`** — L5792 — `private void sendAudioStreamSamples(int samples)`
  Called by: `.sendInitialRadioState()` (same file), `.handleAudioSampleRate()` (same file), `.handleAudioStreamSamples()` (same file)
- **`.sendTxStreamAudioBuffering()`** — L5797 — `private void sendTxStreamAudioBuffering(int milliseconds)`
  Called by: `.sendInitialRadioState()` (same file), `.parseTextFrame()` (same file), `.handleTxStreamAudioBuffering()` (same file)
- **`.wantsIQStream()`** — L5802 — `private bool wantsIQStream(int receiver)`
  Called by: `.PublishIQSamples()` (same file)
- **`.wantsAudioStream()`** — L5811 — `private bool wantsAudioStream(int receiver)`
  Called by: `.PublishRxAudioSamples()` (same file)
- **`.IsReadyForStreaming()`** — L5819 — `internal bool IsReadyForStreaming()`
  Called by: `.RefreshStreamRunState()` (same file)
- **`.WantsAnyRxStream()`** — L5824 — `internal bool WantsAnyRxStream()`
  Called by: `.RefreshStreamRunState()` (same file)
- **`.PublishIQSamples()`** — L5832 — `internal void PublishIQSamples(int receiver, int sampleRate, float[] iqSamples, int complexSamples = -1)`
  Called by: `.PublishIQSamples()` (same file)
- **`.PublishRxAudioSamples()`** — L5842 — `internal void PublishRxAudioSamples(int receiver, int sampleRate, float[] left, float[] right, int samples = -1)`
  Called by: `.PublishRxAudioSamples()` (same file)
- **`.SendTxChrono()`** — L5920 — `internal void SendTxChrono(int receiver)`
  Sends tx chrono.
  Called by: `.SendTxChrono()` (same file)
- **`.UsesTCITxAudio()`** — L5940 — `internal bool UsesTCITxAudio()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UsesActiveTCITxAudio()`** — L5948 — `internal bool UsesActiveTCITxAudio()`
  Called by: `.TryAcquireActiveTxAudioListener()` (same file), `.UsesActiveTCITxAudio()` (same file)
- **`.TryGetTxAudioRequestSettings()`** — L5955 — `internal bool TryGetTxAudioRequestSettings(out int sampleRate, out int samples, out int bufferingMs)`
  Called by: `.TryGetTxAudioRequestSettings()` (same file)
- **`.SyncTciPttToMox()`** — L5965 — `internal void SyncTciPttToMox(bool expectedMox)`
  Called by: `.OnMoxChangeHandler()` (same file), `.OnMoxPreChangeHandler()` (same file)
- **`.clearQueuedTxAudio()`** — L5983 — `private void clearQueuedTxAudio()`
  Called by: `.StopSocketListener()` (same file), `.handleTrxMessage()` (same file), `.SyncTciPttToMox()` (same file)
- **`.TryDequeueTxAudio()`** — L5991 — `internal bool TryDequeueTxAudio(out TCIQueuedTxAudio queuedAudio)`
  Called by: `.TryDequeueTxAudio()` (same file)
- **`.handleBinaryFrame()`** — L6007 — `private void handleBinaryFrame(byte[] payload)`
  Called by: `.ParseReceiveBuffer()` (same file)
- **`.handleIQSampleRate()`** — L6110 — `private void handleIQSampleRate(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.getCurrentMaxHWSampleRate()`** — L6129 — `private int getCurrentMaxHWSampleRate()`
  Returns current max hwsample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleAudioSampleRate()`** — L6145 — `private void handleAudioSampleRate(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleIQStart()`** — L6202 — `private void handleIQStart(string[] args, bool enable)`
  Called by: `.parseTextFrame()` (same file)
- **`.sendIQStartStop()`** — L6219 — `private void sendIQStartStop(int receiver, bool enable)`
  Called by: `.sendInitialRadioState()` (same file), `.handleIQStart()` (same file)
- **`.applyIQSampleRateToReceiver()`** — L6224 — `private void applyIQSampleRateToReceiver(int receiver, int sampleRate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleRxChannelEnable()`** — L6252 — `private void handleRxChannelEnable(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.sendRxChannelEnable()`** — L6292 — `private void sendRxChannelEnable(int rx, int channel, bool enabled)`
  Called by: `.sendInitialRadioState()` (same file), `.handleRxChannelEnable()` (same file)
- **`.handleAudioStart()`** — L6296 — `private void handleAudioStart(string[] args, bool enable)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleAudioStreamSampleType()`** — L6313 — `private void handleAudioStreamSampleType(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleAudioStreamChannels()`** — L6340 — `private void handleAudioStreamChannels(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleAudioStreamSamples()`** — L6356 — `private void handleAudioStreamSamples(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.handleTxStreamAudioBuffering()`** — L6390 — `private void handleTxStreamAudioBuffering(string[] args)`
  Called by: `.parseTextFrame()` (same file)
- **`.PingFrameTimer()`** — L6406 — `private void PingFrameTimer(object o)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOcallback()`** — L6411 — `private void VFOcallback(Object o)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Centrecallback()`** — L6416 — `private void Centrecallback(Object o)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOChange()`** — L6421 — `public void VFOChange(VFOData vfod)`
  Called by: `.handleVFOMessage()` (same file), `.OnVFOAFrequencyChangeHandler()` (same file), `.OnVFOBFrequencyChangeHandler()` (same file)
- **`.CentreChange()`** — L6441 — `public void CentreChange(VFOData vfod)`
  Called by: `.OnCentreFrequencyChanged()` (same file)

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

- **`.Init()`** — L6571 — `private void Init(IPEndPoint ipNport)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StartServer()`** — L6664 — `public void StartServer(Console c, int rateLimit = 0)`
  Starts server.
  Called by: `.SetupTCI()` (`Console/console.cs`)
- **`.StopServer()`** — L6832 — `public void StopServer()`
  Stops server.
  Called by: `.StartServer()` (same file), `.SetupTCI()` (`Console/console.cs`), `.Console_Closing()` (`Console/console.cs`)
- **`.GetCwMacrosSpeed()`** — L6950 — `internal int GetCwMacrosSpeed()`
  Returns cw macros speed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCwMacrosSpeed()`** — L6955 — `internal void SetCwMacrosSpeed(int wpm)`
  Sets cw macros speed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCwMacrosDelay()`** — L6960 — `internal int GetCwMacrosDelay()`
  Returns cw macros delay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCwMacrosDelay()`** — L6965 — `internal void SetCwMacrosDelay(int delayMs)`
  Sets cw macros delay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCwKeyerSpeed()`** — L6970 — `internal int GetCwKeyerSpeed()`
  Returns cw keyer speed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCwKeyerSpeed()`** — L6975 — `internal void SetCwKeyerSpeed(int wpm)`
  Sets cw keyer speed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IncreaseCwMacrosSpeed()`** — L6980 — `internal void IncreaseCwMacrosSpeed(int amount)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DecreaseCwMacrosSpeed()`** — L6985 — `internal void DecreaseCwMacrosSpeed(int amount)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCwTerminalEnabled()`** — L6990 — `internal void SetCwTerminalEnabled(TCPIPtciSocketListener socketListener, int rx, bool enabled)`
  Sets cw terminal enabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendCwMacro()`** — L6995 — `internal void SendCwMacro(TCPIPtciSocketListener socketListener, int rx, string text)`
  Sends cw macro.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendCwMessage()`** — L7000 — `internal void SendCwMessage(TCPIPtciSocketListener socketListener, int rx, string prefix, string callsign, string suffix)`
  Sends cw message.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateCwMessageCallsign()`** — L7005 — `internal void UpdateCwMessageCallsign(TCPIPtciSocketListener socketListener, string callsign)`
  Updates cw message callsign.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StopCwMacros()`** — L7010 — `internal void StopCwMacros(TCPIPtciSocketListener socketListener)`
  Stops cw macros.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleCwKeyer()`** — L7015 — `internal void HandleCwKeyer(TCPIPtciSocketListener socketListener, int rx, bool pressed, int durationMs)`
  Handles cw keyer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NotifyCwTciPttReleased()`** — L7020 — `internal void NotifyCwTciPttReleased(TCPIPtciSocketListener socketListener)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSocketListenerDisconnected()`** — L7025 — `internal void OnSocketListenerDisconnected(TCPIPtciSocketListener socketListener)`
  Handles/raises the socket listener disconnected event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCwMacrosEmpty()`** — L7030 — `internal void OnCwMacrosEmpty(int rx)`
  Handles/raises the cw macros empty event.
  Called by: `.OnRemoteCharacterStarted()` (same file), `.PollCallback()` (same file)
- **`.OnCwCallsignSent()`** — L7043 — `internal void OnCwCallsignSent(string callsign)`
  Handles/raises the cw callsign sent event.
  Called by: `.PollCallback()` (same file)
- **`.OnCwMacrosSpeedChanged()`** — L7056 — `private void OnCwMacrosSpeedChanged(int oldSpeed, int newSpeed)`
  Handles/raises the cw macros speed changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCwMacrosDelayChanged()`** — L7072 — `private void OnCwMacrosDelayChanged(int oldDelay, int newDelay)`
  Handles/raises the cw macros delay changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCwRemoteCharacterStarted()`** — L7085 — `private void OnCwRemoteCharacterStarted(int remainingRemoteCharacters, int pendingElements)`
  Handles/raises the cw remote character started event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCwKeyerSpeedChanged()`** — L7090 — `private void OnCwKeyerSpeedChanged(int oldSpeed, int newSpeed)`
  Handles/raises the cw keyer speed changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StopAllSocketListers()`** — L7124 — `private void StopAllSocketListers()`
  Stops all socket listers.
  Called by: `.StopServer()` (same file)
- **`.ServerThreadStart()`** — L7145 — `private void ServerThreadStart()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PurgingThreadStart()`** — L7199 — `private void PurgingThreadStart()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClientConnectedHandler()`** — L7238 — `private void ClientConnectedHandler()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClientDisconnectedHandler()`** — L7243 — `private void ClientDisconnectedHandler()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClientErrorHandler()`** — L7248 — `private void ClientErrorHandler(SocketException se)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOAFrequencyChangeHandler()`** — L7254 — `public void OnVFOAFrequencyChangeHandler(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double ol`
  Handles/raises the vfoafrequency change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOBFrequencyChangeHandler()`** — L7285 — `public void OnVFOBFrequencyChangeHandler(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double ol`
  Handles/raises the vfobfrequency change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMoxChangeHandler()`** — L7310 — `public void OnMoxChangeHandler(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMoxPreChangeHandler()`** — L7326 — `public void OnMoxPreChangeHandler(int rx, bool currentMox, bool expectedMox)`
  Handles/raises the mox pre change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnModeChangeHandler()`** — L7340 — `public void OnModeChangeHandler(int rx, DSPMode oldMode, DSPMode newMode, Band oldBand, Band newBand)`
  Handles/raises the mode change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnBandChangeHandler()`** — L7352 — `public void OnBandChangeHandler(int rx, Band oldBand, Band newBand)`
  Handles/raises the band change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCentreFrequencyChanged()`** — L7364 — `public void OnCentreFrequencyChanged(int rx, double oldFreq, double newFreq, Band band, double offset)`
  Handles/raises the centre frequency changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFilterChanged()`** — L7391 — `public void OnFilterChanged(int rx, Filter oldFilter, Filter newFilter, Band band, int low, int high, string sName)`
  Handles/raises the filter changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFilterEdgesChanged()`** — L7403 — `public void OnFilterEdgesChanged(int rx, Filter filter, Band band, int low, int high, string sName, int max_width, int max_shift)`
  Handles/raises the filter edges changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXFiltersChanged()`** — L7415 — `public void OnTXFiltersChanged(int low, int high)`
  Handles/raises the txfilters changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPowerChangeHander()`** — L7427 — `public void OnPowerChangeHander(bool oldPower, bool newPower)`
  Handles/raises the power change hander event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnThetisFocusChanged()`** — L7439 — `public void OnThetisFocusChanged(bool focus)`
  Handles/raises the thetis focus changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRX2EnabledChanged()`** — L7451 — `public void OnRX2EnabledChanged(bool enabled)`
  Handles/raises the rx2 enabled changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnHWSampleRateChanged()`** — L7463 — `private void OnHWSampleRateChanged(int rx, int oldSampleRate, int newSampleRate)`
  Handles/raises the hwsample rate changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDrivePowerChanged()`** — L7475 — `private void OnDrivePowerChanged(int rx, int newPower, bool tune)`
  Handles/raises the drive power changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTuneChanged()`** — L7487 — `private void OnTuneChanged(int rx, bool oldTune, bool newTune)`
  Handles/raises the tune changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSplitChanged()`** — L7499 — `private void OnSplitChanged(int rx, bool oldSplit, bool newSplit)`
  Handles/raises the split changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSpotClicked()`** — L7512 — `private void OnSpotClicked(string callsign, long frequencyHz, int rx = -1, bool vfoB = false)`
  Handles/raises the spot clicked event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMuteChanged()`** — L7525 — `private void OnMuteChanged(int rx, bool oldState, bool newState)`
  Handles/raises the mute changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnNrChanged()`** — L7537 — `private void OnNrChanged(int rx, int old_nr, int new_nr)`
  Handles/raises the nr changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnNbChanged()`** — L7549 — `private void OnNbChanged(int rx, int old_nb, int new_nb)`
  Handles/raises the nb changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAnfChanged()`** — L7561 — `private void OnAnfChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the anf changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnBinChanged()`** — L7573 — `private void OnBinChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the bin changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAGCModeChanged()`** — L7585 — `private void OnAGCModeChanged(int rx, AGCMode old_mode, AGCMode new_mode)`
  Handles/raises the agcmode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAGCAutoModeChanged()`** — L7597 — `private void OnAGCAutoModeChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the agcauto mode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOSyncChanged()`** — L7609 — `private void OnVFOSyncChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the vfosync changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVfoALockChanged()`** — L7623 — `private void OnVfoALockChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the vfo alock changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVfoBLockChanged()`** — L7636 — `private void OnVfoBLockChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the vfo block changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSqlChanged()`** — L7650 — `private void OnSqlChanged(int rx, SquelchState old_state, SquelchState new_state)`
  Handles/raises the sql changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSqlLevelChanged()`** — L7662 — `private void OnSqlLevelChanged(int rx, int oldValue, int newValue)`
  Handles/raises the sql level changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnApfChanged()`** — L7674 — `private void OnApfChanged(int rx, bool oldState, bool newState)`
  Handles/raises the apf changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTnfChanged()`** — L7686 — `private void OnTnfChanged(bool old_tnf, bool new_tnf)`
  Handles/raises the tnf changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDiglOffsetChanged()`** — L7698 — `private void OnDiglOffsetChanged(int oldValue, int newValue)`
  Handles/raises the digl offset changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDiguOffsetChanged()`** — L7710 — `private void OnDiguOffsetChanged(int oldValue, int newValue)`
  Handles/raises the digu offset changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRxAfGainChanged()`** — L7722 — `private void OnRxAfGainChanged(int rx, bool is_subrx, int old_gain, int new_gain)`
  Handles/raises the rx af gain changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCTUNChanged()`** — L7734 — `private void OnCTUNChanged(int rx, bool oldCTUN, bool newCTUN, Band band)`
  Handles/raises the ctunchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXProfileChanged()`** — L7746 — `private void OnTXProfileChanged(string old_name, string new_name)`
  Handles/raises the txprofile changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXProfilesChanged()`** — L7758 — `private void OnTXProfilesChanged()`
  Handles/raises the txprofiles changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCalibrationChanged()`** — L7770 — `private void OnCalibrationChanged(int rx, float oldcal, float newcal)`
  Handles/raises the calibration changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMONChanged()`** — L7782 — `private void OnMONChanged(bool oldState, bool newState)`
  Handles/raises the monchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMONVolumeChanged()`** — L7794 — `private void OnMONVolumeChanged(int oldVolume, int newVolume)`
  Handles/raises the monvolume changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVolumeChanged()`** — L7806 — `private void OnVolumeChanged(int oldVolume, int newVolume)`
  Handles/raises the volume changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnBalanceChanged()`** — L7818 — `private void OnBalanceChanged(int rx, bool is_subrx, int oldValue, int newValue)`
  Handles/raises the balance changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAttenuatorDataChanged()`** — L7830 — `private void OnAttenuatorDataChanged(int rx, int oldValue, int newValue)`
  Handles/raises the attenuator data changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnStepAttEnabledChanged()`** — L7842 — `private void OnStepAttEnabledChanged(int rx, bool oldEnabled, bool newEnabled)`
  Handles/raises the step att enabled changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPreampModeChanged()`** — L7854 — `private void OnPreampModeChanged(int rx, PreampMode oldMode, PreampMode newMode)`
  Handles/raises the preamp mode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFMDeviationChanged()`** — L7866 — `private void OnFMDeviationChanged(int rx, int oldValue, int newValue)`
  Handles/raises the fmdeviation changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAGCGainChanged()`** — L7878 — `private void OnAGCGainChanged(int rx, int oldValue, int newValue)`
  Handles/raises the agcgain changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRITChanged()`** — L7890 — `private void OnRITChanged(bool oldState, bool newState)`
  Handles/raises the ritchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnXITChanged()`** — L7902 — `private void OnXITChanged(bool oldState, bool newState)`
  Handles/raises the xitchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRITValueChanged()`** — L7914 — `private void OnRITValueChanged(int oldValue, int newValue)`
  Handles/raises the ritvalue changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnXITValueChanged()`** — L7926 — `private void OnXITValueChanged(int oldValue, int newValue)`
  Handles/raises the xitvalue changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXFrequencyChanged()`** — L7938 — `private void OnTXFrequencyChanged(double old_frequency, double new_frequency, Band old_band, Band new_band, bool rx2_enabled, bool tx_vfob, double centre_freq)`
  Handles/raises the txfrequency changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMeterReadingsChanged()`** — L7969 — `private void OnMeterReadingsChanged(int rx, bool tx, ref Dictionary<Reading, float> readings)`
  Handles/raises the meter readings changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowLog()`** — L7981 — `public void ShowLog()`
  Shows log.
  Called by: `.ShowTCILog()` (`Console/console.cs`)
- **`.CloseLog()`** — L7986 — `public void CloseLog()`
  Closes log.
  Called by: `.SetupTCI()` (`Console/console.cs`)
- **`.SendSpotSimulationClickToAll()`** — L7990 — `public void SendSpotSimulationClickToAll(string callsign, long freq)`
  Sends spot simulation click to all.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RefreshStreamRunState()`** — L8003 — `internal void RefreshStreamRunState()`
  Refreshes stream run state.
  Called by: `.ClientConnectedHandler()` (same file), `.ClientDisconnectedHandler()` (same file)
- **`.PublishIQSamples()`** — L8034 — `public void PublishIQSamples(int receiver, int sampleRate, float[] iqSamples, int complexSamples = -1)`
  Called by: `.serviceTCIRxStreams()` (`Console/cmaster.cs`)
- **`.PublishRxAudioSamples()`** — L8047 — `public void PublishRxAudioSamples(int receiver, int sampleRate, float[] left, float[] right, int samples = -1)`
  Called by: `.serviceTCIRxStreams()` (`Console/cmaster.cs`)
- **`.RequiresRxSensorUpdate()`** — L8060 — `public bool RequiresRxSensorUpdate(int receiver, int channel)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SensorRequiresUpdate()`** — L8076 — `public bool SensorRequiresUpdate(int receiver, Reading reading)`
  Called by: `.MultiMeter2UpdateRX1()` (`Console/console.cs`), `.MultiMeter2UpdateRX2()` (`Console/console.cs`), `.updateMetersReading()` (`Console/console.cs`)
- **`.MinimumRequiredRxSensorInterval()`** — L8092 — `public int MinimumRequiredRxSensorInterval()`
  Called by: `.MultiMeter2UpdateRX1()` (`Console/console.cs`), `.MultiMeter2UpdateRX2()` (`Console/console.cs`)
- **`.MinimumRequiredTxSensorInterval()`** — L8112 — `public int MinimumRequiredTxSensorInterval()`
  Called by: `.MultiMeter2UpdateRX1()` (`Console/console.cs`), `.MultiMeter2UpdateRX2()` (`Console/console.cs`)
- **`.GetActiveTxAudioListener()`** — L8132 — `private TCPIPtciSocketListener GetActiveTxAudioListener()`
  Returns active tx audio listener.
  Called by: `.TryAcquireActiveTxAudioListener()` (same file), `.UsesActiveTCITxAudio()` (same file), `.TryGetTxAudioRequestSettings()` (same file), `.SendTxChrono()` (same file), `.TryDequeueTxAudio()` (same file)
- **`.TryAcquireActiveTxAudioListener()`** — L8146 — `internal bool TryAcquireActiveTxAudioListener(TCPIPtciSocketListener socketListener)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReleaseActiveTxAudioListener()`** — L8167 — `internal void ReleaseActiveTxAudioListener(TCPIPtciSocketListener socketListener)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UsesActiveTCITxAudio()`** — L8176 — `internal bool UsesActiveTCITxAudio()`
  Called by: `.RefreshTxAudioSourceState()` (same file), `.serviceTCITxProtocol()` (`Console/cmaster.cs`)
- **`.TryGetTxAudioRequestSettings()`** — L8184 — `internal bool TryGetTxAudioRequestSettings(out int sampleRate, out int samples, out int bufferingMs)`
  Called by: `.serviceTCITxProtocol()` (`Console/cmaster.cs`)
- **`.RefreshTxAudioSourceState()`** — L8200 — `internal void RefreshTxAudioSourceState()`
  Refreshes tx audio source state.
  Called by: `.OnMoxChangeHandler()` (same file), `.OnMoxPreChangeHandler()` (same file)
- **`.SendTxChrono()`** — L8205 — `public void SendTxChrono(int receiver)`
  Sends tx chrono.
  Called by: `.serviceTCITxProtocol()` (`Console/cmaster.cs`)
- **`.TryDequeueTxAudio()`** — L8216 — `internal bool TryDequeueTxAudio(out TCIQueuedTxAudio queuedAudio)`
  Called by: `.serviceTCITxProtocol()` (`Console/cmaster.cs`)

#### `TCICWController` (type, L8231)

- **`.Dispose()`** — L8303 — `public void Dispose()`
  Releases the object’s resources.
  Called by: `.StartServer()` (same file), `.StopServer()` (same file)
- **`.GetMacroSpeed()`** — L8346 — `public int GetMacroSpeed()`
  Returns macro speed.
  Called by: `.GetCwMacrosSpeed()` (same file), `.IncreaseMacroSpeed()` (same file), `.DecreaseMacroSpeed()` (same file), `.buildMacroOperation()` (same file), `.buildMessageOperation()` (same file)
- **`.SetMacroSpeed()`** — L8357 — `public void SetMacroSpeed(int wpm)`
  Sets macro speed.
  Called by: `.IncreaseMacroSpeed()` (same file), `.DecreaseMacroSpeed()` (same file), `.Stop()` (same file)
- **`.SetMacroSpeedSilently()`** — L8370 — `private void SetMacroSpeedSilently(int wpm)`
  Sets macro speed silently.
  Called by: `.queueNextSegmentLocked()` (same file), `.completeActiveOperationLocked()` (same file), `.abortOperationsForNonCWLocked()` (same file)
- **`.GetMacroDelayMs()`** — L8383 — `public int GetMacroDelayMs()`
  Returns macro delay ms.
  Called by: `.GetCwMacrosDelay()` (same file)
- **`.SetMacroDelayMs()`** — L8388 — `public void SetMacroDelayMs(int delayMs)`
  Sets macro delay ms.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetKeyerSpeed()`** — L8393 — `public int GetKeyerSpeed()`
  Returns keyer speed.
  Called by: `.GetCwKeyerSpeed()` (same file)
- **`.SetKeyerSpeed()`** — L8398 — `public void SetKeyerSpeed(int wpm)`
  Sets keyer speed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IncreaseMacroSpeed()`** — L8403 — `public void IncreaseMacroSpeed(int amount)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DecreaseMacroSpeed()`** — L8408 — `public void DecreaseMacroSpeed(int amount)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTerminalEnabled()`** — L8413 — `public void SetTerminalEnabled(TCPIPtciSocketListener owner, int rx, bool enabled)`
  Sets terminal enabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendMacro()`** — L8449 — `public void SendMacro(TCPIPtciSocketListener owner, int rx, string text)`
  Sends macro.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendMessage()`** — L8464 — `public void SendMessage(TCPIPtciSocketListener owner, int rx, string prefix, string callsign, string suffix)`
  Sends message.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleKeyer()`** — L8479 — `public void HandleKeyer(TCPIPtciSocketListener owner, int rx, bool pressed, int durationMs)`
  Handles keyer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdatePendingCallsign()`** — L8521 — `public void UpdatePendingCallsign(TCPIPtciSocketListener owner, string callsign)`
  Updates pending callsign.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRemoteCharacterStarted()`** — L8539 — `public void OnRemoteCharacterStarted(int remainingRemoteCharacters, int pendingElements)`
  Handles/raises the remote character started event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Stop()`** — L8563 — `public void Stop(TCPIPtciSocketListener owner)`
  Called by: `.DisconnectClient()` (same file)
- **`.HandleTciPttReleased()`** — L8612 — `public void HandleTciPttReleased(TCPIPtciSocketListener owner)`
  Handles tci ptt released.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisconnectClient()`** — L8625 — `public void DisconnectClient(TCPIPtciSocketListener owner)`
  Disconnects client.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clampMacroSpeed()`** — L8642 — `private static int clampMacroSpeed(int speed)`
  Called by: `.SetMacroSpeed()` (same file), `.SetMacroSpeedSilently()` (same file), `.parseMacroText()` (same file)
- **`.decodeTciText()`** — L8647 — `private static string decodeTciText(string text)`
  Called by: `.UpdatePendingCallsign()` (same file), `.normalizeMessageField()` (same file), `.buildMacroOperation()` (same file)
- **`.normalizeMessageField()`** — L8653 — `private static string normalizeMessageField(string text)`
  Called by: `.buildMessageOperation()` (same file)
- **`.translateAbbreviationToken()`** — L8659 — `private static string translateAbbreviationToken(string token)`
  Called by: `.parseMacroText()` (same file)
- **`.buildRepeatedCallsign()`** — L8678 — `private static string buildRepeatedCallsign(string callsign, int repeatCount)`
  Called by: `.UpdatePendingCallsign()` (same file), `.buildMessageOperation()` (same file)
- **`.parseCallsignBase()`** — L8686 — `private static string parseCallsignBase(string callsign, out int repeatCount)`
  Called by: `.UpdatePendingCallsign()` (same file), `.buildMessageOperation()` (same file)
- **`.parseMacroText()`** — L8703 — `private static CWTextParseResult parseMacroText(string text, int startingSpeed)`
  Called by: `.buildMacroOperation()` (same file), `.buildMessageOperation()` (same file)
- **`.buildMacroOperation()`** — L8764 — `private CWTxOperation buildMacroOperation(int rx, string text)`
  Called by: `.SendMacro()` (same file)
- **`.buildMessageOperation()`** — L8778 — `private CWTxOperation buildMessageOperation(int rx, string prefix, string callsign, string suffix)`
  Called by: `.SendMessage()` (same file)
- **`.PollCallback()`** — L8817 — `private void PollCallback(object state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.startNextOperationLocked()`** — L8883 — `private void startNextOperationLocked()`
  Called by: `.SendMacro()` (same file), `.SendMessage()` (same file), `.PollCallback()` (same file), `.completeActiveOperationLocked()` (same file)
- **`.queueNextSegmentLocked()`** — L8906 — `private void queueNextSegmentLocked()`
  Called by: `.PollCallback()` (same file), `.startNextOperationLocked()` (same file)
- **`.completeActiveOperationLocked()`** — L8931 — `private void completeActiveOperationLocked()`
  Called by: `.PollCallback()` (same file)
- **`.isCWModeLocked()`** — L8946 — `private bool isCWModeLocked()`
  Called by: `.HandleKeyer()` (same file), `.PollCallback()` (same file), `.startNextOperationLocked()` (same file), `.queueNextSegmentLocked()` (same file)
- **`.abortOperationsForNonCWLocked()`** — L8956 — `private void abortOperationsForNonCWLocked()`
  Called by: `.PollCallback()` (same file), `.startNextOperationLocked()` (same file), `.queueNextSegmentLocked()` (same file)
- **`.isCurrentOwnerLocked()`** — L8987 — `private bool isCurrentOwnerLocked(TCPIPtciSocketListener owner)`
  Called by: `.SetTerminalEnabled()` (same file), `.HandleKeyer()` (same file), `.UpdatePendingCallsign()` (same file), `.Stop()` (same file), `.HandleTciPttReleased()` (same file), `.DisconnectClient()` (same file)
- **`.tryAcquireOwnershipLocked()`** — L8992 — `private bool tryAcquireOwnershipLocked(TCPIPtciSocketListener owner)`
  Called by: `.SetTerminalEnabled()` (same file), `.SendMacro()` (same file), `.SendMessage()` (same file), `.HandleKeyer()` (same file)
- **`.releaseOwnershipIfIdleLocked()`** — L9005 — `private void releaseOwnershipIfIdleLocked()`
  Called by: `.SetTerminalEnabled()` (same file), `.HandleKeyer()` (same file), `.Stop()` (same file), `.completeActiveOperationLocked()` (same file), `.abortOperationsForNonCWLocked()` (same file), `.releaseKeyerLocked()` (same file)
- **`.KeyerSchedulerThreadProc()`** — L9012 — `private void KeyerSchedulerThreadProc()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.scheduleKeyerReleaseLocked()`** — L9050 — `private void scheduleKeyerReleaseLocked(int durationMs)`
  Called by: `.HandleKeyer()` (same file)
- **`.tryReleaseDirectKeyerFromPollLocked()`** — L9066 — `private bool tryReleaseDirectKeyerFromPollLocked()`
  Called by: `.PollCallback()` (same file)
- **`.releaseKeyerLocked()`** — L9095 — `private void releaseKeyerLocked()`
  Called by: `.HandleKeyer()` (same file), `.HandleTciPttReleased()` (same file), `.KeyerSchedulerThreadProc()` (same file), `.scheduleKeyerReleaseLocked()` (same file), `.tryReleaseDirectKeyerFromPollLocked()` (same file)
- **`.cancelKeyerReleaseScheduleLocked()`** — L9114 — `private void cancelKeyerReleaseScheduleLocked()`
  Called by: `.Stop()` (same file), `.abortOperationsForNonCWLocked()` (same file), `.releaseKeyerLocked()` (same file)
- **`.isCwTargetAvailableLocked()`** — L9121 — `private bool isCwTargetAvailableLocked(int rx)`
  Called by: `.SendMacro()` (same file), `.SendMessage()` (same file)
- **`.selectCwTargetLocked()`** — L9127 — `private bool selectCwTargetLocked(int rx)`
  Called by: `.HandleKeyer()` (same file), `.PollCallback()` (same file), `.startNextOperationLocked()` (same file), `.queueNextSegmentLocked()` (same file)
- **`.ensureTerminalTciPttLocked()`** — L9149 — `private void ensureTerminalTciPttLocked()`
  Called by: `.SetTerminalEnabled()` (same file), `.startNextOperationLocked()` (same file), `.queueNextSegmentLocked()` (same file)
- **`.releaseTerminalTciPttIfOwnedLocked()`** — L9168 — `private void releaseTerminalTciPttIfOwnedLocked()`
  Called by: `.SetTerminalEnabled()` (same file), `.Stop()` (same file), `.PollCallback()` (same file), `.startNextOperationLocked()` (same file), `.completeActiveOperationLocked()` (same file), `.abortOperationsForNonCWLocked()` (same file)
- **`.isTerminalEnabledLocked()`** — L9178 — `private bool isTerminalEnabledLocked(int rx)`
  Called by: `.OnRemoteCharacterStarted()` (same file), `.PollCallback()` (same file), `.startNextOperationLocked()` (same file), `.completeActiveOperationLocked()` (same file), `.ensureTerminalTciPttLocked()` (same file)
- **`.isAnyTerminalEnabledLocked()`** — L9183 — `private bool isAnyTerminalEnabledLocked()`
  Called by: `.releaseOwnershipIfIdleLocked()` (same file)
- **`.beginDirectKeyerLocked()`** — L9194 — `private bool beginDirectKeyerLocked()`
  Called by: `.HandleKeyer()` (same file)
- **`.ensureDirectKeyerMoxLocked()`** — L9212 — `private bool ensureDirectKeyerMoxLocked()`
  Called by: `.beginDirectKeyerLocked()` (same file)
- **`.captureDirectKeyerMoxReleaseLocked()`** — L9240 — `private bool captureDirectKeyerMoxReleaseLocked()`
  Called by: `.Dispose()` (same file), `.Stop()` (same file), `.abortOperationsForNonCWLocked()` (same file), `.releaseKeyerLocked()` (same file), `.beginDirectKeyerLocked()` (same file)
- **`.releaseDirectKeyerMox()`** — L9247 — `private void releaseDirectKeyerMox()`
  Called by: `.Dispose()` (same file), `.Stop()` (same file), `.abortOperationsForNonCWLocked()` (same file), `.releaseKeyerLocked()` (same file), `.beginDirectKeyerLocked()` (same file)
- **`.waitForScheduledKeyerRelease()`** — L9257 — `private bool waitForScheduledKeyerRelease(long releaseAtTicks)`
  Called by: `.KeyerSchedulerThreadProc()` (same file)
- **`.millisecondsToStopwatchTicks()`** — L9280 — `private static long millisecondsToStopwatchTicks(int durationMs)`
  Called by: `.scheduleKeyerReleaseLocked()` (same file), `.tryReleaseDirectKeyerFromPollLocked()` (same file)
- **`.stopwatchTicksToMilliseconds()`** — L9286 — `private static double stopwatchTicksToMilliseconds(long ticks)`
  Called by: `.waitForScheduledKeyerRelease()` (same file)
- **`.setDirectKeyerState()`** — L9291 — `private static bool setDirectKeyerState(bool pressed)`
  Sets direct keyer state.
  Called by: `.Dispose()` (same file), `.Stop()` (same file), `.abortOperationsForNonCWLocked()` (same file), `.releaseKeyerLocked()` (same file), `.beginDirectKeyerLocked()` (same file)
- **`.InvokeOnConsole()`** — L9305 — `private T InvokeOnConsole<T>(Func<Console, T> action, T defaultValue)`
  Called by: `.Dispose()` (same file), `.GetMacroSpeed()` (same file), `.SetMacroSpeed()` (same file), `.SetMacroSpeedSilently()` (same file), `.GetMacroDelayMs()` (same file), `.SetMacroDelayMs()` (same file) — and 15 more

#### `CWTxSegment` (type, L8235)

_No extracted members._

#### `CWTextParseResult` (type, L8241)

_No extracted members._

#### `CWTxOperation` (type, L8248)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/TCIServer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
