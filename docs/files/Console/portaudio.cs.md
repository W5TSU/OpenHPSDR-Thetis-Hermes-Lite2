# `Console/portaudio.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** P/Invoke wrapper for PortAudio device/host-API enumeration.

## How this file is used

- Used by (incoming references from other files):
  - `Console/audio.cs` (calls ×18)
  - `Console/clsCMASIOConfig.cs` (calls ×4)
  - `Console/console.cs` (calls ×4)
  - `Console/setup.cs` (calls ×4)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.PA_GetHostApiInfo()` (×8), `.PA_GetDeviceInfo()` (×7), `.PA_GetHostApiCount()` (×5), `.PA_HostApiDeviceIndexToDeviceIndex()` (×4), `.PA_GetVersion()` (×2), `.PA_GetDeviceCount()` (×1), `.PA_IsFormatSupported()` (×1), `.PA_Initialize()` (×1)

## Outline

### Types

#### `PA19` (type, L60)

- `.PA_GetVersion()` — L202
- `.PA_GetVersionText()` — L205
- `.IntPtr_PA_GetErrorText()` — L213
- `.PA_GetErrorText()` — L216
- `.PA_Initialize()` — L222
- `.PA_Terminate()` — L225
- `.PA_GetHostApiCount()` — L228
- `.PA_GetDefaultHostApi()` — L231
- `.PA_GetHostApiInfoPtr()` — L236
- `.PA_GetHostApiInfo()` — L238
- `.PA_HostApiTypeIdToHostApiIndex()` — L245
- `.PA_HostApiDeviceIndexToDeviceIndex()` — L248
- `.PA_GetLastHostErrorInfoPtr()` — L251
- `.PA_GetLastHostErrorInfo()` — L253
- `.PA_GetDeviceCount()` — L260
- `.PA_GetDefaultInputDevice()` — L263
- `.PA_GetDefaultOutputDevice()` — L266
- `.PA_GetDeviceInfoPtr()` — L269
- `.PA_GetDeviceInfo()` — L271
- `.PA_IsFormatSupported()` — L278
- `.PA_OpenStream()` — L284
- `.PA_OpenDefaultStream()` — L295
- `.PA_CloseStream()` — L306
- `.PA_SetStreamFinishedCallback()` — L309
- `.PA_StartStream()` — L313
- `.PA_StopStream()` — L316
- `.PA_AbortStream()` — L319
- `.PA_IsStreamStopped()` — L322
- `.PA_IsStreamActive()` — L325
- `.PA_GetStreamInfoPtr()` — L328
- `.PA_GetStreamInfo()` — L330
- `.PA_GetStreamTime()` — L337
- `.PA_GetStreamCpuLoad()` — L340
- `.PA_ReadStream()` — L345
- `.PA_WriteStream()` — L348
- `.PA_GetStreamReadAvailable()` — L351
- `.PA_GetStreamWriteAvailable()` — L354
- `.PA_GetSampleSize()` — L357
- `.PA_Sleep()` — L360

#### `PaErrorCode` (type, L91)

_No extracted members._

#### `PaHostApiTypeId` (type, L103)

_No extracted members._

#### `PaStreamCallbackResult` (type, L121)

_No extracted members._

#### `PaHostApiInfo` (type, L128)

_No extracted members._

#### `PaHostErrorInfo` (type, L143)

_No extracted members._

#### `PaDeviceInfo` (type, L152)

_No extracted members._

#### `PaStreamParameters` (type, L171)

_No extracted members._

#### `PaStreamCallbackTimeInfo` (type, L181)

_No extracted members._

#### `PaStreamInfo` (type, L189)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/portaudio.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
