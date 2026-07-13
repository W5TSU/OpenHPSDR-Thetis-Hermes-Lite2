# `Console/HPSDR/specHPSDR.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** Configures the wdsp spectrum analyzer instances for HPSDR data streams.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×26, references ×1)
  - `Console/clsSpectrumProcessor.cs` (calls ×9, references ×1)
  - `Console/MeterManager.cs` (calls ×5, references ×1)
  - `Console/wbDisplay.cs` (calls ×5)
  - `Console/PanDisplay.cs` (calls ×2)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetSpecRX()` (×22), `.resetPixelBuffers()` (×8), `.GetPixels()` (×5), `.initAnalyzer()` (×2), `.SetAnalyzer()` (×2), `.SnapSpectrum()` (×2), `.ZoomToBandwidth()` (×1), `.GetFrequencyExtents()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `SpecRX` (type, L29)

- **`.GetSpecRX()`** — L46 — `public SpecHPSDR GetSpecRX(int disp)`
  Returns spec rx.
  Called by: `.CalcDisplayFreq()` (`Console/console.cs`), `.CalcRX2DisplayFreq()` (`Console/console.cs`), `.CalcTXDisplayFreq()` (`Console/console.cs`), `.UpdateRXDisplayVars()` (`Console/console.cs`), `.UpdateTXDisplayVars()` (`Console/console.cs`), `.CalibrateFreq()` (`Console/console.cs`) — and 16 more

#### `SpecHPSDR` (type, L52)

- **`.updateNormalizePan()`** — L288 — `void updateNormalizePan()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.resetPixelBuffers()`** — L490 — `public void resetPixelBuffers()`
  Called by: `.resetBuffers()` (`Console/MeterManager.cs`), `.SetPixels()` (`Console/clsSpectrumProcessor.cs`), `.SetFrameRate()` (`Console/clsSpectrumProcessor.cs`), `.SetZoomSlider()` (`Console/clsSpectrumProcessor.cs`), `.SetPanSlider()` (`Console/clsSpectrumProcessor.cs`), `.SyncFrequencies()` (`Console/clsSpectrumProcessor.cs`) — and 2 more
- **`.initAnalyzer()`** — L504 — `public void initAnalyzer()`
  Called by: `.setupSpecDetails()` (`Console/MeterManager.cs`), `.RefreshLocked()` (`Console/clsSpectrumProcessor.cs`)
- **`.ZoomToBandwidth()`** — L645 — `public void ZoomToBandwidth(double target_bandwidth_hz)`
  Called by: `.zoom()` (`Console/MeterManager.cs`)
- **`.GetFrequencyExtents()`** — L672 — `public (int, int) GetFrequencyExtents(double zslider, double panslider)`
  Returns frequency extents.
  Called by: `.setPan()` (`Console/MeterManager.cs`)
- **`.CalcSpectrum()`** — L738 — `public void CalcSpectrum(int filter_low, int filter_high, int spec_blocksize, int sample_rate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SpecHPSDRDLL` (type, L809)

- **`.SetAnalyzer()`** — L812 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetAnalyzer( int disp, int n_pixout, int n_fft,`
  Sets analyzer.
  Called by: `.initAnalyzer()` (same file), `.CalcSpectrum()` (same file), `.initAnalyzer()` (`Console/PanDisplay.cs`), `.initWideband()` (`Console/wbDisplay.cs`)
- **`.XCreateAnalyzer()`** — L835 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void XCreateAnalyzer(int disp, ref int success, int m_size, int m_LO, int m_stitch, string`
  Called by: `.create_wideband()` (`Console/wbDisplay.cs`)
- **`.ResetPixelBuffers()`** — L839 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void ResetPixelBuffers(int disp)`
  Resets pixel buffers.
  Called by: `.resetPixelBuffers()` (same file)
- **`.DestroyAnalyzer()`** — L842 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void DestroyAnalyzer(int disp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPixelRef()`** — L845 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetPixelRef(int disp, double pixel_ref)`
  Sets pixel ref.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPixelsNative()`** — L848 — `[DllImport("WDSP.dll", EntryPoint = "GetPixels", CallingConvention = CallingConvention.Cdecl)] private static extern void GetPixelsNative(int disp, int pixout, float* pix, ref int `
  Returns pixels native.
  Called by: `.GetPixels()` (same file)
- **`.GetPixels()`** — L851 — `public static void GetPixels(int disp, int pixout, float* pix, ref int flag)`
  Returns pixels.
  Called by: `.runDisplay()` (`Console/MeterManager.cs`), `.RunDisplay()` (`Console/PanDisplay.cs`), `.ProcessFrame()` (`Console/clsSpectrumProcessor.cs`), `.RunDisplay()` (`Console/console.cs`), `.RunDisplay()` (`Console/wbDisplay.cs`)
- **`.Spectrum()`** — L862 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void Spectrum(int disp, int ss, int LO, float* pI, float* pQ)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCalibration()`** — L865 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetCalibration(int disp, int set, int points, IntPtr cal)`
  Sets calibration.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SnapSpectrum()`** — L868 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SnapSpectrum(int disp, int ss, int LO, double* snap_buff)`
  Called by: `.CalibrateFreq()` (`Console/console.cs`), `.CalibrateLevel()` (`Console/console.cs`)
- **`.SnapSpectrumTimeout()`** — L871 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SnapSpectrumTimeout(int disp, int ss, int LO, double* snap_buff, uint timeout, ref in`
  Called by: `.FindPeakFreqInPassband()` (`Console/console.cs`)
- **`.SetDisplayDetectorMode()`** — L874 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDisplayDetectorMode (int disp, int pixout, int mode)`
  Sets display detector mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDisplayAverageMode()`** — L877 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDisplayAverageMode (int disp, int pixout, int mode)`
  Sets display average mode.
  Called by: `.initWideband()` (`Console/wbDisplay.cs`)
- **`.SetDisplayNumAverage()`** — L880 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDisplayNumAverage (int disp, int pixout, int num)`
  Sets display num average.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDisplayAvBackmult()`** — L883 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDisplayAvBackmult(int disp, int pixout, double mult)`
  Sets display av backmult.
  Called by: `.initWideband()` (`Console/wbDisplay.cs`)
- **`.SetDisplayNormOneHz()`** — L886 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDisplayNormOneHz(int disp, int pixout, bool norm)`
  Sets display norm one hz.
  Called by: `.updateNormalizePan()` (same file)
- **`.GetDisplayENB()`** — L889 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern double GetDisplayENB(int disp)`
  Returns display enb.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetDisplaySampleRate()`** — L892 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetDisplaySampleRate (int disp, int rate)`
  Sets display sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.create_nobEXT()`** — L895 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void create_nobEXT( int id, int run, int mode,`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.destroy_nobEXT()`** — L909 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void destroy_nobEXT(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.xnobEXTF()`** — L912 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void xnobEXTF(int id, float* I, float* Q)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBBuffsize()`** — L915 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBBuffsize(int id, int size)`
  Sets extnobbuffsize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBSamplerate()`** — L918 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBSamplerate(int id, int rate)`
  Sets extnobsamplerate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBTau()`** — L921 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBTau(int id, double tau)`
  Sets extnobtau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBHangtime()`** — L924 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBHangtime(int id, double time)`
  Sets extnobhangtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBAdvtime()`** — L927 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBAdvtime(int id, double time)`
  Sets extnobadvtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBBacktau()`** — L930 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBBacktau(int id, double tau)`
  Sets extnobbacktau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBThreshold()`** — L933 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBThreshold(int id, double thresh)`
  Sets extnobthreshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTNOBMode()`** — L936 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTNOBMode(int id, int mode)`
  Sets extnobmode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.create_anbEXT()`** — L939 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void create_anbEXT( int id, int run, int buffsize,`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.destroy_anbEXT()`** — L952 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void destroy_anbEXT(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.xanbEXTF()`** — L955 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void xanbEXTF(int id, float* I, float* Q)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTANBBuffsize()`** — L958 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTANBBuffsize(int id, int size)`
  Sets extanbbuffsize.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTANBSamplerate()`** — L961 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTANBSamplerate(int id, int rate)`
  Sets extanbsamplerate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTANBTau()`** — L964 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTANBTau(int id, double tau)`
  Sets extanbtau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTANBHangtime()`** — L967 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTANBHangtime(int id, double time)`
  Sets extanbhangtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTANBAdvtime()`** — L970 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTANBAdvtime(int id, double time)`
  Sets extanbadvtime.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTANBBacktau()`** — L973 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTANBBacktau(int id, double tau)`
  Sets extanbbacktau.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEXTANBThreshold()`** — L976 — `[DllImport("WDSP.dll", CallingConvention = CallingConvention.Cdecl)] public static extern void SetEXTANBThreshold(int id, double thresh)`
  Sets extanbthreshold.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/specHPSDR.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
