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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.ReadChunk()`** — L3074 — `public static Chunk ReadChunk(ref BinaryReader reader)`
  Reads chunk.
  Called by: `.tryParseWaveHeader()` (same file)

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

- **`.ToString()`** — L135 — `public override string ToString()`
  Returns the string representation.
  Called by: `.writeRecordingJson()` (same file)

#### `clsAudioRecordPlayback` (type, L154)

- **`.initPlaybackSettings()`** — L270 — `private void initPlaybackSettings()`
  Called by: `.SetPlaybackSetting()` (same file), `.GetPlaybackSetting()` (same file)
- **`.tryGetWavDurationSeconds()`** — L285 — `private static double? tryGetWavDurationSeconds(string wavPath)`
  Called by: `.refreshExistingJsonFromWavIfNeeded()` (same file), `.recordCompleted()` (same file), `.writeRecordingJson()` (same file)
- **`.tryParseUtcStamp()`** — L332 — `private static bool tryParseUtcStamp(string s, out DateTime utc)`
  Called by: `.refreshExistingJsonFromWavIfNeeded()` (same file)
- **`.refreshExistingJsonFromWavIfNeeded()`** — L345 — `private void refreshExistingJsonFromWavIfNeeded(string unique_id, string wavPath, int formatTag, int sampleRate, int channels, int bitsPerSample)`
  Called by: `.PlayFileViaWDSP()` (same file), `.PlayFileViaPCAudio()` (same file)
- **`.SetPlaybackSetting()`** — L478 — `public void SetPlaybackSetting(string setting, bool value)`
  Sets playback setting.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPlaybackSetting()`** — L485 — `public bool GetPlaybackSetting(string setting)`
  Returns playback setting.
  Called by: `.activatePlaybackRecordSettings()` (same file)
- **`.storeRestoreSettings()`** — L493 — `private void storeRestoreSettings(bool store, bool playback)`
  Called by: `.RecordToFileFromWDSP()` (same file), `.StopRecord()` (same file), `.PlayFileViaWDSP()` (same file), `.StopPlayback()` (same file), `.completeRecordStateIfCurrent()` (same file)
- **`.activatePlaybackRecordSettings()`** — L533 — `private void activatePlaybackRecordSettings(bool playback, bool ignore_temp_changes = false)`
  Called by: `.RecordToFileFromWDSP()` (same file), `.PlayFileViaWDSP()` (same file)
- **`.OnPreMox()`** — L559 — `private void OnPreMox(int rx, bool oldMox, bool newMox)`
  Handles/raises the pre mox event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L568 — `public void Dispose()`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OkToRecord()`** — L633 — `public bool OkToRecord(string filepath, bool allowUnknown = true)`
  Called by: `.onRecordSpaceTimer()` (same file), `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file)
- **`.startRecordSpaceTimer()`** — L646 — `private void startRecordSpaceTimer()`
  Called by: `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file)
- **`.stopRecordSpaceTimer()`** — L653 — `private void stopRecordSpaceTimer()`
  Called by: `.Dispose()` (same file), `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file), `.StopRecord()` (same file), `.completeRecordStateIfCurrent()` (same file)
- **`.onRecordSpaceTimer()`** — L662 — `private void onRecordSpaceTimer(object state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPcInputDevices()`** — L694 — `public List<AudioDeviceInfo> GetPcInputDevices()`
  Returns pc input devices.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPcOutputDevices()`** — L735 — `public List<AudioDeviceInfo> GetPcOutputDevices()`
  Returns pc output devices.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DeleteRecording()`** — L776 — `public bool DeleteRecording(string full_path, out string error, bool delete_containing_folder_if_empty = false)`
  Deletes recording.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.hasTrailingSeparator()`** — L910 — `private static bool hasTrailingSeparator(string path)`
  Called by: `.ensureTrailingSeparator()` (same file), `.resolveRecordPath()` (same file)
- **`.containsDirectorySeparator()`** — L917 — `private static bool containsDirectorySeparator(string path)`
  Called by: `.resolveRecordPath()` (same file)
- **`.hasRealExtension()`** — L923 — `private static bool hasRealExtension(string path)`
  Called by: `.DeleteRecording()` (same file), `.resolveRecordPath()` (same file)
- **`.ensureTrailingSeparator()`** — L932 — `private static string ensureTrailingSeparator(string path)`
  Called by: `.isUnderBaseFolder()` (same file)
- **`.isUnderBaseFolder()`** — L939 — `private static bool isUnderBaseFolder(string baseFolder, string candidatePath)`
  Called by: `.DeleteRecording()` (same file), `.resolveRecordPath()` (same file), `.resolvePlayPath()` (same file)
- **`.makeUniquePathInFolder()`** — L948 — `private string makeUniquePathInFolder(string folder, string prefix, out string filename)`
  Called by: `.resolveRecordPath()` (same file)
- **`.resolveRecordPath()`** — L963 — `private string resolveRecordPath(string record_id, string input_path, string prefix, out string filename, out string error)`
  Called by: `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file)
- **`.resolvePlayPath()`** — L1053 — `private string resolvePlayPath(string input_path, out string error)`
  Called by: `.DeleteRecording()` (same file), `.CanBePlayed()` (same file), `.PlayFileViaWDSP()` (same file), `.PlayFileViaPCAudio()` (same file)
- **`.ensureUtc()`** — L1081 — `private static DateTime ensureUtc(DateTime dt)`
  Called by: `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file), `.recordCompleted()` (same file), `.writeRecordingJson()` (same file)
- **`.ensureDetails()`** — L1089 — `private static RecordingDetails ensureDetails(RecordingDetails details, bool needed)`
  Called by: `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file)
- **`.GetJSONDetailsFromFile()`** — L1102 — `public bool GetJSONDetailsFromFile(string full_path_file, out RecordingJsonModel json_data)`
  Returns jsondetails from file.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RecordToFileFromWDSP()`** — L1154 — `public string RecordToFileFromWDSP(string record_id, string full_path, int wfw_id, out string error, bool remove_if_file_exists = false, RecordingDetails details = null, bool ignor`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RecordToFileFromPCAudio()`** — L1334 — `public string RecordToFileFromPCAudio(string record_id, string full_path, int pcAudioDeviceInputId, out string error, bool remove_if_file_exists = false, RecordingDetails details =`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StopRecord()`** — L1579 — `public bool StopRecord(out string error)`
  Stops record.
  Called by: `.OnPreMox()` (same file), `.Dispose()` (same file), `.onRecordSpaceTimer()` (same file)
- **`.applyPcInputSourceStereoRemap()`** — L1648 — `private void applyPcInputSourceStereoRemap(byte[] buffer, int bytesRecorded)`
  Called by: `.onPcDataAvailable()` (same file)
- **`.CanBePlayed()`** — L1694 — `public bool CanBePlayed(string filepath)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PlayFileViaWDSP()`** — L1736 — `public bool PlayFileViaWDSP(string play_id, string full_path, int wfw_id, out string error, double adjustGain_dB = 0, bool ignore_temp_changes = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.adjustWavePreampByDB()`** — L1852 — `private double adjustWavePreampByDB(double nonDBfloor, double db_adjust)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PlayFileViaPCAudio()`** — L1862 — `public bool PlayFileViaPCAudio(string play_id, string full_path, int pcAudioDeviceOutputId, out string error)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.pcWaveOut_PlaybackStopped()`** — L1953 — `private void pcWaveOut_PlaybackStopped(object sender, StoppedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cleanupPcPlayback()`** — L1985 — `private void cleanupPcPlayback()`
  Called by: `.PlayFileViaPCAudio()` (same file), `.StopPlayback()` (same file)
- **`.StopPlayback()`** — L2052 — `public bool StopPlayback(out string error)`
  Stops playback.
  Called by: `.OnPreMox()` (same file), `.Dispose()` (same file), `.pcWaveOut_PlaybackStopped()` (same file), `.onWdspPlaybackFinished()` (same file)
- **`.getMediaFailureMessage()`** — L2106 — `private static string getMediaFailureMessage(Exception ex)`
  Returns media failure message.
  Called by: `.pcWaveOut_PlaybackStopped()` (same file), `.markActiveRecordFailureLocked()` (same file), `.onWdspPlaybackFinished()` (same file)
- **`.markActiveRecordFailureLocked()`** — L2115 — `private void markActiveRecordFailureLocked(Exception ex)`
  Called by: `.stopFaultedPcRecording()` (same file), `.onPcDataAvailable()` (same file), `.onPcRecordingStopped()` (same file), `.onWdspRecordFinished()` (same file)
- **`.stopFaultedPcRecording()`** — L2134 — `private void stopFaultedPcRecording(Exception ex)`
  Called by: `.onPcDataAvailable()` (same file)
- **`.onPcDataAvailable()`** — L2201 — `private void onPcDataAvailable(object sender, WaveInEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onPcRecordingStopped()`** — L2365 — `private void onPcRecordingStopped(object sender, StoppedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cleanupPcRecording()`** — L2396 — `private void cleanupPcRecording()`
  Called by: `.RecordToFileFromPCAudio()` (same file), `.stopFaultedPcRecording()` (same file), `.onPcRecordingStopped()` (same file)
- **`.onWdspPlaybackFinished()`** — L2413 — `private void onWdspPlaybackFinished(Exception ex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onWdspRecordFinished()`** — L2440 — `private void onWdspRecordFinished(string wavPath, Exception ex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.recordCompleted()`** — L2489 — `private void recordCompleted(string unique_id, string wav, string json, string mp3, RecordingDetails details, bool recordSucceeded, string failureMessage)`
  Called by: `.stopFaultedPcRecording()` (same file), `.onPcRecordingStopped()` (same file), `.onWdspRecordFinished()` (same file)
- **`.writeRecordingJson()`** — L2584 — `private bool writeRecordingJson(string unique_id, string jsonPath, RecordingDetails details)`
  Called by: `.refreshExistingJsonFromWavIfNeeded()` (same file), `.recordCompleted()` (same file)
- **`.raiseRecordingJsonWritten()`** — L2661 — `private void raiseRecordingJsonWritten(string unique_id, RecordingJsonModel json_data)`
  Called by: `.writeRecordingJson()` (same file)
- **`.raiseMediaError()`** — L2679 — `private void raiseMediaError(Action<string, string, string> handler, string unique_id, string filename, string errorMessage)`
  Called by: `.raiseRecordError()` (same file), `.raisePlaybackError()` (same file)
- **`.raiseRecordError()`** — L2711 — `private void raiseRecordError(string unique_id, string filename, string errorMessage)`
  Called by: `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file), `.StopRecord()` (same file), `.recordCompleted()` (same file)
- **`.raisePlaybackError()`** — L2716 — `private void raisePlaybackError(string unique_id, string filename, string errorMessage)`
  Called by: `.PlayFileViaWDSP()` (same file), `.PlayFileViaPCAudio()` (same file), `.pcWaveOut_PlaybackStopped()` (same file), `.StopPlayback()` (same file), `.onWdspPlaybackFinished()` (same file)
- **`.ensureMediaFoundation()`** — L2720 — `private static void ensureMediaFoundation()`
  Called by: `.generateMp3FromWav()` (same file)
- **`.generateMp3FromWav()`** — L2737 — `private bool generateMp3FromWav(string wavPath, string mp3Path)`
  Called by: `.recordCompleted()` (same file)
- **`.clearActiveRecordLocked()`** — L2756 — `private void clearActiveRecordLocked()`
  Called by: `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file), `.StopRecord()` (same file), `.completeRecordStateIfCurrent()` (same file)
- **`.isActiveRecordMatchLocked()`** — L2769 — `private bool isActiveRecordMatchLocked(string unique_id, string wav)`
  Called by: `.completeRecordStateIfCurrent()` (same file)
- **`.completeRecordStateIfCurrent()`** — L2775 — `private void completeRecordStateIfCurrent(string unique_id, string wav)`
  Called by: `.StopRecord()` (same file), `.recordCompleted()` (same file)
- **`.setRecordingState()`** — L2813 — `private void setRecordingState(bool recording)`
  Sets recording state.
  Called by: `.RecordToFileFromWDSP()` (same file), `.RecordToFileFromPCAudio()` (same file), `.StopRecord()` (same file)
- **`.setPlayingState()`** — L2853 — `private void setPlayingState(bool playing)`
  Sets playing state.
  Called by: `.PlayFileViaWDSP()` (same file), `.PlayFileViaPCAudio()` (same file), `.StopPlayback()` (same file)
- **`.generateFilename()`** — L2898 — `private static string generateFilename(string prefix, int suffixNumber, string ext)`
  Called by: `.makeUniquePathInFolder()` (same file)
- **`.ensureFolderExists()`** — L2905 — `private static void ensureFolderExists(string folder)`
  Called by: `.makeUniquePathInFolder()` (same file), `.resolveRecordPath()` (same file), `.writeRecordingJson()` (same file)
- **`.tryParseWaveHeader()`** — L2915 — `private static bool tryParseWaveHeader(BinaryReader reader, out int formatTag, out int sampleRate, out int channels, out int bitsPerSample, out long dataStart, out long dataLengthB`
  Called by: `.tryGetWavDurationSeconds()` (same file), `.CanBePlayed()` (same file), `.PlayFileViaWDSP()` (same file), `.PlayFileViaPCAudio()` (same file)

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

- **`.UpdateMox()`** — L3296 — `public void UpdateMox()`
  Updates mox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ProcessRecordBuffers()`** — L3312 — `private void ProcessRecordBuffers()`
  Processes record buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddWriteBuffer()`** — L3361 — `unsafe public void AddWriteBuffer(float* left, float* right, int nsamps)`
  Adds write buffer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Stop()`** — L3368 — `public void Stop()`
  Called by: `.StopRecord()` (same file)
- **`.WaitForStop()`** — L3373 — `public bool WaitForStop(int timeout_ms)`
  Called by: `.StopRecord()` (same file)
- **`.WriteBuffer()`** — L3381 — `private void WriteBuffer(ref BinaryWriter w, ref int count)`
  Writes buffer.
  Called by: `.ProcessRecordBuffers()` (same file)
- **`.Write_32()`** — L3428 — `private void Write_32(int length, ref int count, int out_cnt, BinaryWriter w)`
  Called by: `.WriteBuffer()` (same file)
- **`.Write_24()`** — L3454 — `private void Write_24(int length, ref int count, int out_cnt, BinaryWriter w)`
  Called by: `.WriteBuffer()` (same file)
- **`.Write_16()`** — L3468 — `private void Write_16(int length, ref int count, int out_cnt, BinaryWriter w)`
  Called by: `.WriteBuffer()` (same file)
- **`.Write_8()`** — L3481 — `private void Write_8(int length, ref int count, int out_cnt, BinaryWriter w)`
  Called by: `.WriteBuffer()` (same file)
- **`.getDitherAmp()`** — L3493 — `private float getDitherAmp()`
  Returns dither amp.
  Called by: `.dither32()` (same file), `.dither24()` (same file), `.dither16()` (same file), `.dither8()` (same file)
- **`.dither32()`** — L3501 — `private int dither32(float sample)`
  Called by: `.Write_32()` (same file)
- **`.dither24()`** — L3514 — `private int dither24(float sample)`
  Called by: `.Write_24()` (same file)
- **`.dither16()`** — L3527 — `private int dither16(float sample)`
  Called by: `.Write_16()` (same file)
- **`.dither8()`** — L3540 — `private sbyte dither8(float sample)`
  Called by: `.Write_8()` (same file)
- **`.WriteWaveHeader()`** — L3553 — `private static void WriteWaveHeader(ref BinaryWriter w, short channels, int sample_rate, short format_tag, short bit_depth, int data_length)`
  Writes wave header.
  Called by: `.ProcessRecordBuffers()` (same file)

#### `WaveFileReader1` (type, L3572)

- **`.buildCosineFade()`** — L3749 — `private static float[] buildCosineFade(int frames)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Stop()`** — L3773 — `public void Stop()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaitForStop()`** — L3778 — `public bool WaitForStop(int timeout_ms)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ProcessBuffers()`** — L3786 — `private void ProcessBuffers()`
  Processes buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadBuffer()`** — L3817 — `private void ReadBuffer(ref BinaryReader r)`
  Reads buffer.
  Called by: `.ProcessBuffers()` (same file)
- **`.GetPlayBuffer()`** — L3993 — `unsafe public void GetPlayBuffer(float* left, float* right)`
  Returns play buffer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.queueFinish()`** — L4023 — `private void queueFinish(Exception ex)`
  Called by: `.ProcessBuffers()` (same file), `.GetPlayBuffer()` (same file)
- **`.finishPlayback()`** — L4035 — `private void finishPlayback()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsAudioRecordPlayback.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
