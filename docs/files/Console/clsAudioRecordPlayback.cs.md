# `Console/clsAudioRecordPlayback.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** RX/TX audio and I/Q recording and playback (wave capture of what you hear/transmit).

## How this file is used

- Used by (incoming references from other files):
  - `Console/cmaster.cs` (references ×2)
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/ringbuffer.cs` (calls ×7, references ×2)
  - `Console/common.cs` (calls ×3)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Functions

- `.ReadChunk()` — L3074

### Types

#### `AudioRecordRxSource` (type, L54)

_No extracted members._

#### `AudioRecordTxSource` (type, L60)

_No extracted members._

#### `AudioBitDepthMode` (type, L66)

_No extracted members._

#### `AudioDeviceDriver` (type, L75)

_No extracted members._

#### `PCInputSource` (type, L81)

_No extracted members._

#### `RecordingDetails` (type, L88)

_No extracted members._

#### `AudioDeviceInfo` (type, L110)

- `.ToString()` — L135

#### `clsAudioRecordPlayback` (type, L154)

- `.initPlaybackSettings()` — L270
- `.tryGetWavDurationSeconds()` — L285
- `.tryParseUtcStamp()` — L332
- `.refreshExistingJsonFromWavIfNeeded()` — L345
- `.SetPlaybackSetting()` — L478
- `.GetPlaybackSetting()` — L485
- `.storeRestoreSettings()` — L493
- `.activatePlaybackRecordSettings()` — L533
- `.OnPreMox()` — L559
- `.Dispose()` — L568
- `.OkToRecord()` — L633
- `.startRecordSpaceTimer()` — L646
- `.stopRecordSpaceTimer()` — L653
- `.onRecordSpaceTimer()` — L662
- `.GetPcInputDevices()` — L694
- `.GetPcOutputDevices()` — L735
- `.DeleteRecording()` — L776
- `.hasTrailingSeparator()` — L910
- `.containsDirectorySeparator()` — L917
- `.hasRealExtension()` — L923
- `.ensureTrailingSeparator()` — L932
- `.isUnderBaseFolder()` — L939
- `.makeUniquePathInFolder()` — L948
- `.resolveRecordPath()` — L963
- `.resolvePlayPath()` — L1053
- `.ensureUtc()` — L1081
- `.ensureDetails()` — L1089
- `.GetJSONDetailsFromFile()` — L1102
- `.RecordToFileFromWDSP()` — L1154
- `.RecordToFileFromPCAudio()` — L1334
- `.StopRecord()` — L1579
- `.applyPcInputSourceStereoRemap()` — L1648
- `.CanBePlayed()` — L1694
- `.PlayFileViaWDSP()` — L1736
- `.adjustWavePreampByDB()` — L1852
- `.PlayFileViaPCAudio()` — L1862
- `.pcWaveOut_PlaybackStopped()` — L1953
- `.cleanupPcPlayback()` — L1985
- `.StopPlayback()` — L2052
- `.getMediaFailureMessage()` — L2106
- `.markActiveRecordFailureLocked()` — L2115
- `.stopFaultedPcRecording()` — L2134
- `.onPcDataAvailable()` — L2201
- `.onPcRecordingStopped()` — L2365
- `.cleanupPcRecording()` — L2396
- `.onWdspPlaybackFinished()` — L2413
- `.onWdspRecordFinished()` — L2440
- `.recordCompleted()` — L2489
- `.writeRecordingJson()` — L2584
- `.raiseRecordingJsonWritten()` — L2661
- `.raiseMediaError()` — L2679
- `.raiseRecordError()` — L2711
- `.raisePlaybackError()` — L2716
- `.ensureMediaFoundation()` — L2720
- `.generateMp3FromWav()` — L2737
- `.clearActiveRecordLocked()` — L2756
- `.isActiveRecordMatchLocked()` — L2769
- `.completeRecordStateIfCurrent()` — L2775
- `.setRecordingState()` — L2813
- `.setPlayingState()` — L2853
- `.generateFilename()` — L2898
- `.ensureFolderExists()` — L2905
- `.tryParseWaveHeader()` — L2915

#### `RecordingJsonModel` (type, L3046)

_No extracted members._

#### `Chunk` (type, L3070)

_No extracted members._

#### `RIFFChunk` (type, L3124)

_No extracted members._

#### `fmtChunk` (type, L3130)

_No extracted members._

#### `dataChunk` (type, L3141)

_No extracted members._

#### `WaveFileWriter` (type, L3146)

- `.UpdateMox()` — L3296
- `.ProcessRecordBuffers()` — L3312
- `.AddWriteBuffer()` — L3361
- `.Stop()` — L3368
- `.WaitForStop()` — L3373
- `.WriteBuffer()` — L3381
- `.Write_32()` — L3428
- `.Write_24()` — L3454
- `.Write_16()` — L3468
- `.Write_8()` — L3481
- `.getDitherAmp()` — L3493
- `.dither32()` — L3501
- `.dither24()` — L3514
- `.dither16()` — L3527
- `.dither8()` — L3540
- `.WriteWaveHeader()` — L3553

#### `WaveFileReader1` (type, L3572)

- `.buildCosineFade()` — L3749
- `.Stop()` — L3773
- `.WaitForStop()` — L3778
- `.ProcessBuffers()` — L3786
- `.ReadBuffer()` — L3817
- `.GetPlayBuffer()` — L3993
- `.queueFinish()` — L4023
- `.finishPlayback()` — L4035

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsAudioRecordPlayback.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
