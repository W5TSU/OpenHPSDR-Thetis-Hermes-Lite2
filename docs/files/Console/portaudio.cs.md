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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `PA19` (type, L60)

- **`.PA_GetVersion()`** — L202 — `[DllImport("PA19.dll")] public static extern int PA_GetVersion()`
  Called by: `.checkVersions()` (`Console/console.cs`), `.miAbout_Click()` (`Console/console.cs`)
- **`.PA_GetVersionText()`** — L205 — `[DllImport("PA19.dll")] public static extern String PA_GetVersionText()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IntPtr_PA_GetErrorText()`** — L213 — `[DllImport("PA19.dll", EntryPoint = "PA_GetErrorText")] public static extern IntPtr IntPtr_PA_GetErrorText(PaError error)`
  note that using the stock source and calling this function on errorCode = 0 will result in an Exception (no object reference. To fix this, I added a single statement in pa_front.c. The new line 444 is below. case paNoError: result = "1"; result = "Success"; break;
  Called by: `.PA_GetErrorText()` (same file)
- **`.PA_GetErrorText()`** — L216 — `public static string PA_GetErrorText(PaError error)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_Initialize()`** — L222 — `[DllImport("PA19.dll")] public static extern PaError PA_Initialize()`
  Called by: `.initialisePortAudio()` (`Console/console.cs`)
- **`.PA_Terminate()`** — L225 — `[DllImport("PA19.dll")] public static extern PaError PA_Terminate()`
  Called by: `.ExitConsole()` (`Console/console.cs`)
- **`.PA_GetHostApiCount()`** — L228 — `[DllImport("PA19.dll")] public static extern PaHostApiIndex PA_GetHostApiCount()`
  Called by: `.GetPAHosts()` (`Console/audio.cs`), `.GetPAInputDevices()` (`Console/audio.cs`), `.GetPAOutputDevices()` (`Console/audio.cs`), `.comboAudioDriver2_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioDriver3_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.PA_GetDefaultHostApi()`** — L231 — `[DllImport("PA19.dll")] public static extern PaHostApiIndex PA_GetDefaultHostApi()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetHostApiInfoPtr()`** — L236 — `[DllImport("PA19.dll", EntryPoint = "PA_GetHostApiInfo")] public static extern IntPtr PA_GetHostApiInfoPtr(int hostId)`
  Added layer to convert from the struct pointer to a C# struct automatically.
  Called by: `.PA_GetHostApiInfo()` (same file)
- **`.PA_GetHostApiInfo()`** — L238 — `public static PaHostApiInfo PA_GetHostApiInfo(int hostId)`
  Called by: `.GetPAHosts()` (`Console/audio.cs`), `.GetPAInputDevices()` (`Console/audio.cs`), `.CheckPAInputDevices()` (`Console/audio.cs`), `.GetPAOutputDevices()` (`Console/audio.cs`), `.CheckPAOutputDevices()` (`Console/audio.cs`), `.GetASIODevices()` (`Console/clsCMASIOConfig.cs`) — and 2 more
- **`.PA_HostApiTypeIdToHostApiIndex()`** — L245 — `[DllImport("PA19.dll")] public static extern PaHostApiIndex PA_HostApiTypeIdToHostApiIndex(PaHostApiTypeId type)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_HostApiDeviceIndexToDeviceIndex()`** — L248 — `[DllImport("PA19.dll")] public static extern PaDeviceIndex PA_HostApiDeviceIndexToDeviceIndex(int hostAPI, int hostApiDeviceIndex)`
  Called by: `.GetPAInputDevices()` (`Console/audio.cs`), `.CheckPAInputDevices()` (`Console/audio.cs`), `.GetPAOutputDevices()` (`Console/audio.cs`), `.CheckPAOutputDevices()` (`Console/audio.cs`)
- **`.PA_GetLastHostErrorInfoPtr()`** — L251 — `[DllImport("PA19.dll", EntryPoint = "PA_GetLastHostErrorInfo")] public static extern IntPtr PA_GetLastHostErrorInfoPtr()`
  Called by: `.PA_GetLastHostErrorInfo()` (same file)
- **`.PA_GetLastHostErrorInfo()`** — L253 — `public static PaHostErrorInfo PA_GetLastHostErrorInfo()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetDeviceCount()`** — L260 — `[DllImport("PA19.dll")] public static extern PaDeviceIndex PA_GetDeviceCount()`
  Called by: `.GetASIODevices()` (`Console/clsCMASIOConfig.cs`)
- **`.PA_GetDefaultInputDevice()`** — L263 — `[DllImport("PA19.dll")] public static extern PaDeviceIndex PA_GetDefaultInputDevice()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetDefaultOutputDevice()`** — L266 — `[DllImport("PA19.dll")] public static extern PaDeviceIndex PA_GetDefaultOutputDevice()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetDeviceInfoPtr()`** — L269 — `[DllImport("PA19.dll", EntryPoint = "PA_GetDeviceInfo")] public static extern IntPtr PA_GetDeviceInfoPtr(int device)`
  Called by: `.PA_GetDeviceInfo()` (same file)
- **`.PA_GetDeviceInfo()`** — L271 — `public static PaDeviceInfo PA_GetDeviceInfo(int device)`
  Called by: `.GetPAInputDevices()` (`Console/audio.cs`), `.CheckPAInputDevices()` (`Console/audio.cs`), `.GetPAOutputDevices()` (`Console/audio.cs`), `.CheckPAOutputDevices()` (`Console/audio.cs`), `.EnableVAC1()` (`Console/audio.cs`), `.EnableVAC2()` (`Console/audio.cs`) — and 1 more
- **`.PA_IsFormatSupported()`** — L278 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_IsFormatSupported( PaStreamParameters* inputParameters, PaStreamParameters* outputParameters, double sampleRate)`
  Called by: `.GetASIODevices()` (`Console/clsCMASIOConfig.cs`)
- **`.PA_OpenStream()`** — L284 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_OpenStream( out void* stream, PaStreamParameters* inputParameters, PaStreamParameters* outputParameters,`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_OpenDefaultStream()`** — L295 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_OpenDefaultStream( out void* stream, int numInputChannels, int numOutputChannels,`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_CloseStream()`** — L306 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_CloseStream(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_SetStreamFinishedCallback()`** — L309 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_SetStreamFinishedCallback( void* stream, PaStreamFinishedCallback streamFinishedCallback)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_StartStream()`** — L313 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_StartStream(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_StopStream()`** — L316 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_StopStream(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_AbortStream()`** — L319 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_AbortStream(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_IsStreamStopped()`** — L322 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_IsStreamStopped(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_IsStreamActive()`** — L325 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_IsStreamActive(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetStreamInfoPtr()`** — L328 — `[DllImport("PA19.dll", EntryPoint = "PA_GetStreamInfo")] unsafe public static extern IntPtr PA_GetStreamInfoPtr(void* stream)`
  Called by: `.PA_GetStreamInfo()` (same file)
- **`.PA_GetStreamInfo()`** — L330 — `unsafe public static PaStreamInfo PA_GetStreamInfo(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetStreamTime()`** — L337 — `[DllImport("PA19.dll")] unsafe public static extern PaTime PA_GetStreamTime(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetStreamCpuLoad()`** — L340 — `[DllImport("PA19.dll")] unsafe public static extern double PA_GetStreamCpuLoad(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_ReadStream()`** — L345 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_ReadStream(void* stream, void* buffer, uint frames)`
  note: These next 4 blocking IO functions are only currently implemented in MME (not DirectSound or ASIO)
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_WriteStream()`** — L348 — `[DllImport("PA19.dll")] unsafe public static extern PaError PA_WriteStream(void* stream, void* buffer, uint frames)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetStreamReadAvailable()`** — L351 — `[DllImport("PA19.dll")] unsafe public static extern int PA_GetStreamReadAvailable(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetStreamWriteAvailable()`** — L354 — `[DllImport("PA19.dll")] unsafe public static extern int PA_GetStreamWriteAvailable(void* stream)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_GetSampleSize()`** — L357 — `[DllImport("PA19.dll")] public static extern PaError PA_GetSampleSize(PaSampleFormat format)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PA_Sleep()`** — L360 — `[DllImport("PA19.dll")] public static extern void PA_Sleep(int msec)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

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
