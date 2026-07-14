# `Console/clsSpectrumProcessor.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Pulls pixel-ready FFT data from the wdsp analyzer and post-processes it (averaging, peak/blend detection) for the display.

## How this file is used

- Used by (incoming references from other files):
  - `Console/ColorButton.cs` (calls ×2)
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/HPSDR/specHPSDR.cs` (calls ×9, references ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/enums.cs` (references ×1)
- Most-referenced symbols from other files: `.Refresh()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `clsSpectrumProcessor` (type, L50)

- **`.AddReceiver()`** — L100 — `public bool AddReceiver(int receiverId, int pixels, int frameRate, int fftSize)`
  Adds receiver.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddTransmitter()`** — L105 — `public bool AddTransmitter(int transmitterId, int pixels, int frameRate, int fftSize)`
  Adds transmitter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddSource()`** — L110 — `private bool AddSource(SpectrumSourceType sourceType, int sourceId, int pixels, int frameRate, int fftSize)`
  Adds source.
  Called by: `.AddReceiver()` (same file), `.AddTransmitter()` (same file), `.ShowTestForm()` (same file)
- **`.RemoveReceiver()`** — L137 — `public bool RemoveReceiver(int receiverId)`
  Removes receiver.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveTransmitter()`** — L142 — `public bool RemoveTransmitter(int transmitterId)`
  Removes transmitter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveSource()`** — L147 — `private bool RemoveSource(SpectrumSourceType sourceType, int sourceId)`
  Removes source.
  Called by: `.RemoveReceiver()` (same file), `.RemoveTransmitter()` (same file)
- **`.Clear()`** — L157 — `public void Clear()`
  Called by: `.Dispose()` (same file)
- **`.ContainsSource()`** — L172 — `private bool ContainsSource(SpectrumSourceType sourceType, int sourceId)`
  Called by: `.ShowTestForm()` (same file)
- **`.SetReceiverPixelResolution()`** — L182 — `public bool SetReceiverPixelResolution(int receiverId, int pixels)`
  Sets receiver pixel resolution.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTransmitterPixelResolution()`** — L187 — `public bool SetTransmitterPixelResolution(int transmitterId, int pixels)`
  Sets transmitter pixel resolution.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPixelResolution()`** — L192 — `private bool SetPixelResolution(SpectrumSourceType sourceType, int sourceId, int pixels)`
  Sets pixel resolution.
  Called by: `.SetReceiverPixelResolution()` (same file), `.SetTransmitterPixelResolution()` (same file)
- **`.SetReceiverFrameRate()`** — L201 — `public bool SetReceiverFrameRate(int receiverId, int frameRate)`
  Sets receiver frame rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTransmitterFrameRate()`** — L206 — `public bool SetTransmitterFrameRate(int transmitterId, int frameRate)`
  Sets transmitter frame rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetFrameRate()`** — L211 — `private bool SetFrameRate(SpectrumSourceType sourceType, int sourceId, int frameRate)`
  Sets frame rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetReceiverFFTSize()`** — L220 — `public bool SetReceiverFFTSize(int receiverId, int fftSize)`
  Sets receiver fftsize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTransmitterFFTSize()`** — L225 — `public bool SetTransmitterFFTSize(int transmitterId, int fftSize)`
  Sets transmitter fftsize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetFFTSize()`** — L230 — `private bool SetFFTSize(SpectrumSourceType sourceType, int sourceId, int fftSize)`
  Sets fftsize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetReceiverSampleRate()`** — L239 — `public bool SetReceiverSampleRate(int receiverId, int sampleRate)`
  Sets receiver sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTransmitterSampleRate()`** — L244 — `public bool SetTransmitterSampleRate(int transmitterId, int sampleRate)`
  Sets transmitter sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetSampleRate()`** — L249 — `private bool SetSampleRate(SpectrumSourceType sourceType, int sourceId, int sampleRate)`
  Sets sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetReceiverZoomSlider()`** — L258 — `public bool SetReceiverZoomSlider(int receiverId, double zoomSlider)`
  Sets receiver zoom slider.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTransmitterZoomSlider()`** — L263 — `public bool SetTransmitterZoomSlider(int transmitterId, double zoomSlider)`
  Sets transmitter zoom slider.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetZoomSlider()`** — L268 — `private bool SetZoomSlider(SpectrumSourceType sourceType, int sourceId, double zoomSlider)`
  Sets zoom slider.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetReceiverPanSlider()`** — L277 — `public bool SetReceiverPanSlider(int receiverId, double panSlider)`
  Sets receiver pan slider.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTransmitterPanSlider()`** — L282 — `public bool SetTransmitterPanSlider(int transmitterId, double panSlider)`
  Sets transmitter pan slider.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPanSlider()`** — L287 — `private bool SetPanSlider(SpectrumSourceType sourceType, int sourceId, double panSlider)`
  Sets pan slider.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CentreReceiverPan()`** — L296 — `public bool CentreReceiverPan(int receiverId)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CentreTransmitterPan()`** — L301 — `public bool CentreTransmitterPan(int transmitterId)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CentrePan()`** — L306 — `private bool CentrePan(SpectrumSourceType sourceType, int sourceId)`
  Called by: `.CentreReceiverPan()` (same file), `.CentreTransmitterPan()` (same file)
- **`.SetReceiverEnabled()`** — L311 — `public bool SetReceiverEnabled(int receiverId, bool enabled)`
  Sets receiver enabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTransmitterEnabled()`** — L316 — `public bool SetTransmitterEnabled(int transmitterId, bool enabled)`
  Sets transmitter enabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEnabled()`** — L321 — `private bool SetEnabled(SpectrumSourceType sourceType, int sourceId, bool enabled)`
  Sets enabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetReceiverBuffers()`** — L330 — `public bool ResetReceiverBuffers(int receiverId)`
  Resets receiver buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetTransmitterBuffers()`** — L335 — `public bool ResetTransmitterBuffers(int transmitterId)`
  Resets transmitter buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBuffers()`** — L340 — `private bool ResetBuffers(SpectrumSourceType sourceType, int sourceId)`
  Resets buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetReceiverPixels()`** — L349 — `public bool TryGetReceiverPixels(int receiverId, out float[] pixels, out int dataIndex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetTransmitterPixels()`** — L354 — `public bool TryGetTransmitterPixels(int transmitterId, out float[] pixels, out int dataIndex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryCopyReceiverPixels()`** — L359 — `public bool TryCopyReceiverPixels(int receiverId, float[] destination, out int pixelCount, out int dataIndex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryCopyTransmitterPixels()`** — L364 — `public bool TryCopyTransmitterPixels(int transmitterId, float[] destination, out int pixelCount, out int dataIndex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetLatestPixels()`** — L369 — `private bool TryGetLatestPixels(SpectrumSourceType sourceType, int sourceId, out float[] pixels, out int dataIndex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryCopyLatestPixels()`** — L382 — `private bool TryCopyLatestPixels(SpectrumSourceType sourceType, int sourceId, float[] destination, out int pixelCount, out int dataIndex)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetViewport()`** — L395 — `private bool TryGetViewport(SpectrumSourceType sourceType, int sourceId, out double zoomSlider, out double panSlider)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetSampleRate()`** — L408 — `private bool TryGetSampleRate(SpectrumSourceType sourceType, int sourceId, out int sampleRate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryGetFFTSize()`** — L420 — `private bool TryGetFFTSize(SpectrumSourceType sourceType, int sourceId, out int fftSize)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowReceiverTestForm()`** — L432 — `public Form ShowReceiverTestForm(int receiverId, float? min_dBm = null, float? max_dBm = null)`
  Shows receiver test form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowTransmitterTestForm()`** — L437 — `public Form ShowTransmitterTestForm(int transmitterId, float? min_dBm = null, float? max_dBm = null)`
  Shows transmitter test form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowTestForm()`** — L442 — `private Form ShowTestForm(SpectrumSourceType sourceType, int sourceId, float? min_dBm = null, float? max_dBm = null)`
  Shows test form.
  Called by: `.ShowReceiverTestForm()` (same file), `.ShowTransmitterTestForm()` (same file)
- **`.Dispose()`** — L455 — `public void Dispose()`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RunWorker()`** — L469 — `private void RunWorker()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SubscribeDelegates()`** — L487 — `private void SubscribeDelegates()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UnsubscribeDelegates()`** — L494 — `private void UnsubscribeDelegates()`
  Called by: `.Dispose()` (same file)
- **`.OnCentreFrequencyChanged()`** — L501 — `private void OnCentreFrequencyChanged(int rx, double oldFreq, double newFreq, Band band, double offset)`
  Handles/raises the centre frequency changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnHWSampleRateChanged()`** — L514 — `private void OnHWSampleRateChanged(int rx, int oldRate, int newRate)`
  Handles/raises the hwsample rate changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPowerChanged()`** — L526 — `private void OnPowerChanged(bool oldPower, bool newPower)`
  Handles/raises the power changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetEndpointArray()`** — L535 — `private SpectrumEndpoint[] GetEndpointArray()`
  Returns endpoint array.
  Called by: `.RunWorker()` (same file), `.OnCentreFrequencyChanged()` (same file), `.OnHWSampleRateChanged()` (same file), `.OnPowerChanged()` (same file)
- **`.TryGetEndpoint()`** — L540 — `private bool TryGetEndpoint(SpectrumSourceType sourceType, int sourceId, out SpectrumEndpoint endpoint)`
  Called by: `.SetPixelResolution()` (same file), `.SetFrameRate()` (same file), `.SetFFTSize()` (same file), `.SetSampleRate()` (same file), `.SetZoomSlider()` (same file), `.SetPanSlider()` (same file) — and 7 more
- **`.TryDetachEndpoint()`** — L551 — `private bool TryDetachEndpoint(SpectrumSourceType sourceType, int sourceId, out SpectrumEndpoint endpoint)`
  Called by: `.RemoveSource()` (same file)
- **`.UpdateEndpointSnapshotLocked()`** — L568 — `private void UpdateEndpointSnapshotLocked()`
  Updates endpoint snapshot locked.
  Called by: `.AddSource()` (same file), `.TryDetachEndpoint()` (same file)
- **`.ThrowIfDisposed()`** — L581 — `private void ThrowIfDisposed()`
  Called by: `.AddSource()` (same file), `.ShowTestForm()` (same file)
- **`.MakeSourceKey()`** — L586 — `private static string MakeSourceKey(SpectrumSourceType sourceType, int sourceId)`
  Called by: `.AddSource()` (same file), `.ContainsSource()` (same file), `.TryGetEndpoint()` (same file), `.TryDetachEndpoint()` (same file)
- **`.ValidateSourceId()`** — L591 — `private static void ValidateSourceId(int sourceId)`
  Called by: `.AddSource()` (same file), `.ContainsSource()` (same file), `.ShowTestForm()` (same file)
- **`.ClampInt()`** — L596 — `private static int ClampInt(int value, int minimum, int maximum)`
  Called by: `.AddSource()` (same file), `.SetPixelResolution()` (same file), `.SetFrameRate()` (same file), `.NormalizeFftSize()` (same file)
- **`.ValidateSampleRate()`** — L603 — `private static int ValidateSampleRate(int sampleRate)`
  Called by: `.SetSampleRate()` (same file), `.OnHWSampleRateChanged()` (same file)
- **`.ClampZoomSlider()`** — L611 — `private static double ClampZoomSlider(double value)`
  Called by: `.SetZoomSlider()` (same file)
- **`.ClampPanSlider()`** — L619 — `private static double ClampPanSlider(double value)`
  Called by: `.SetPanSlider()` (same file)
- **`.NormalizeFftSize()`** — L627 — `private static int NormalizeFftSize(int fftSize)`
  Called by: `.AddSource()` (same file), `.SetFFTSize()` (same file)

#### `SpectrumSourceType` (type, L52)

_No extracted members._

#### `SpectrumEndpoint` (type, L646)

- **`.MatchesReceiverEvent()`** — L719 — `public bool MatchesReceiverEvent(int rx)`
  Called by: `.OnCentreFrequencyChanged()` (same file), `.OnHWSampleRateChanged()` (same file)
- **`.SetPixels()`** — L724 — `public void SetPixels(int pixels)`
  Sets pixels.
  Called by: `.SetPixelResolution()` (same file)
- **`.SetFrameRate()`** — L740 — `public void SetFrameRate(int frameRate)`
  Sets frame rate.
  Called by: `.SetReceiverFrameRate()` (same file), `.SetTransmitterFrameRate()` (same file), `.SetFrameRate()` (same file)
- **`.SetFFTSize()`** — L754 — `public void SetFFTSize(int fftSize)`
  Sets fftsize.
  Called by: `.SetReceiverFFTSize()` (same file), `.SetTransmitterFFTSize()` (same file), `.SetFFTSize()` (same file)
- **`.SetSampleRate()`** — L765 — `public void SetSampleRate(int sampleRate)`
  Sets sample rate.
  Called by: `.SetReceiverSampleRate()` (same file), `.SetTransmitterSampleRate()` (same file), `.SetSampleRate()` (same file), `.OnHWSampleRateChanged()` (same file)
- **`.SetZoomSlider()`** — L777 — `public void SetZoomSlider(double zoomSlider)`
  Sets zoom slider.
  Called by: `.SetReceiverZoomSlider()` (same file), `.SetTransmitterZoomSlider()` (same file), `.SetZoomSlider()` (same file)
- **`.SetPanSlider()`** — L792 — `public void SetPanSlider(double panSlider)`
  Sets pan slider.
  Called by: `.SetReceiverPanSlider()` (same file), `.SetTransmitterPanSlider()` (same file), `.SetPanSlider()` (same file), `.CentrePan()` (same file)
- **`.SetEnabled()`** — L807 — `public void SetEnabled(bool enabled)`
  Sets enabled.
  Called by: `.SetReceiverEnabled()` (same file), `.SetTransmitterEnabled()` (same file), `.SetEnabled()` (same file)
- **`.SyncFrequencies()`** — L819 — `public void SyncFrequencies()`
  Called by: `.OnCentreFrequencyChanged()` (same file)
- **`.Refresh()`** — L832 — `public void Refresh()`
  Called by: `.OnClick()` (`Console/ColorButton.cs`), `.MoveIndex()` (`Console/ColorButton.cs`)
- **`.ResetBuffers()`** — L842 — `public void ResetBuffers()`
  Resets buffers.
  Called by: `.ResetReceiverBuffers()` (same file), `.ResetTransmitterBuffers()` (same file), `.ResetBuffers()` (same file)
- **`.ClearData()`** — L855 — `public void ClearData()`
  Clears data.
  Called by: `.OnPowerChanged()` (same file)
- **`.ProcessFrame()`** — L865 — `public bool ProcessFrame(long nowUtcTicks)`
  Processes frame.
  Called by: `.RunWorker()` (same file)
- **`.TryGetLatestPixels()`** — L902 — `public bool TryGetLatestPixels(out float[] pixels, out int dataIndex)`
  Called by: `.TryGetReceiverPixels()` (same file), `.TryGetTransmitterPixels()` (same file), `.TryGetLatestPixels()` (same file)
- **`.TryCopyLatestPixels()`** — L913 — `public bool TryCopyLatestPixels(float[] destination, out int pixelCount, out int dataIndex)`
  Called by: `.TryCopyReceiverPixels()` (same file), `.TryCopyTransmitterPixels()` (same file), `.TryCopyLatestPixels()` (same file)
- **`.TryGetViewport()`** — L930 — `public bool TryGetViewport(out double zoomSlider, out double panSlider)`
  Called by: `.TryGetViewport()` (same file)
- **`.TryGetSampleRate()`** — L940 — `public bool TryGetSampleRate(out int sampleRate)`
  Called by: `.TryGetSampleRate()` (same file)
- **`.TryGetFFTSize()`** — L949 — `public bool TryGetFFTSize(out int fftSize)`
  Called by: `.TryGetFFTSize()` (same file)
- **`.Shutdown()`** — L958 — `public void Shutdown()`
  Called by: `.AddSource()` (same file), `.RemoveSource()` (same file), `.Clear()` (same file)
- **`.RefreshLocked()`** — L974 — `private void RefreshLocked()`
  Refreshes locked.
  Called by: `.SetFFTSize()` (same file), `.SetSampleRate()` (same file), `.Refresh()` (same file)
- **`.ResolveDefaultBlockSize()`** — L1003 — `private static int ResolveDefaultBlockSize(int sampleRate)`
  Called by: `.SetSampleRate()` (same file)
- **`.GetPixelOffset()`** — L1009 — `private float GetPixelOffset()`
  Returns pixel offset.
  Called by: `.ProcessFrame()` (same file)
- **`.EnsurePixelBuffersLocked()`** — L1022 — `private void EnsurePixelBuffersLocked(int pixels)`
  Called by: `.ClearPixelBuffersLocked()` (same file)
- **`.ClearPixelBuffersLocked()`** — L1037 — `private void ClearPixelBuffersLocked()`
  Clears pixel buffers locked.
  Called by: `.SetPixels()` (same file), `.SetZoomSlider()` (same file), `.SetPanSlider()` (same file), `.SyncFrequencies()` (same file), `.ResetBuffers()` (same file), `.ClearData()` (same file) — and 1 more
- **`.RentPixelBuffer()`** — L1044 — `private static float[] RentPixelBuffer(int pixels)`
  Called by: `.EnsurePixelBuffersLocked()` (same file)
- **`.ReturnPixelBuffer()`** — L1049 — `private static void ReturnPixelBuffer(ref float[] buffer)`
  Called by: `.Shutdown()` (same file), `.EnsurePixelBuffersLocked()` (same file)
- **`.FillPixelBuffer()`** — L1057 — `private static void FillPixelBuffer(float[] data, int pixels)`
  Called by: `.ClearPixelBuffersLocked()` (same file)

#### `SpectrumTestForm` (type, L1064)

- **`.BuildTitle()`** — L1105 — `private static string BuildTitle(SpectrumSourceType sourceType, int sourceId, float? minDbm, float? maxDbm)`
  Builds title.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RefreshTimer_Tick()`** — L1117 — `private void RefreshTimer_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `RefreshTimer` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SpectrumTestForm_FormClosed()`** — L1122 — `private void SpectrumTestForm_FormClosed(object sender, FormClosedEventArgs e)`
  WinForms event handler: runs when `SpectrumTestForm` has closed.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `SpectrumGraphPanel` (type, L1129)

- **`.Dispose()`** — L1164 — `protected override void Dispose(bool disposing)`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L1177 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryCopyPixels()`** — L1294 — `private bool TryCopyPixels(out float[] pixels, out int pixelCount, out int dataIndex)`
  Called by: `.OnPaint()` (same file)
- **`.TryCopyPixelsWithBuffer()`** — L1310 — `private bool TryCopyPixelsWithBuffer(float[] destination, out int pixelCount, out int dataIndex)`
  Called by: `.TryCopyPixels()` (same file)
- **`.EnsurePointBuffer()`** — L1317 — `private void EnsurePointBuffer(int pixelCount)`
  Called by: `.OnPaint()` (same file)
- **`.DrawGrid()`** — L1323 — `private void DrawGrid(Graphics g, Rectangle plotRect)`
  Draws grid.
  Called by: `.OnPaint()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsSpectrumProcessor.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
