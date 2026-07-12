# `wdsp/analyzer.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** The multi-instance FFT spectrum analyzer behind every panadapter/waterfall.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
  - `ChannelMaster/analyzers.c` (calls ×2)
  - `ChannelMaster/network.c` (calls ×1)
  - `ChannelMaster/pipe.c` (calls ×1)
  - `wdsp/sender.c` (calls ×1)
  - `wdsp/siphon.c` (calls ×1)
- Uses (outgoing references to other files):
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/combase.h` (calls ×3)
  - `wdsp/meterlog10.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `XCreateAnalyzer()` (×3), `DestroyAnalyzer()` (×3), `Spectrum0()` (×3), `Spectrum()` (×1), `Spectrum2()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`bessi0()`** — L33 — `double bessi0(double x)`
  Called by: `new_window()` (same file)
- **`new_window()`** — L52 — `void new_window(int disp, int type, int size, double PiAlpha)`
  Called by: `SetAnalyzer()` (same file)
- **`eliminate()`** — L180 — `void eliminate(int disp, int ss, int LO)`
  spur elimination, REAL input data
  Called by: `spectra()` (same file)
- **`Celiminate()`** — L215 — `void Celiminate(int disp, int ss, int LO)`
  spur elimination, COMPLEX input data
  Called by: `Cspectra()` (same file)
- **`detector()`** — L283 — `void detector ( int det_type, int m, int num_pixels, double pix_per_bin, double bin_per_pix,`
  Called by: `stitch()` (same file)
- **`avenger()`** — L464 — `void avenger ( int av_mode, int num_pixels, int* avail_frames, int num_average, int* av_in_idx,`
  Called by: `stitch()` (same file)
- **`stitch()`** — L556 — `void stitch(int disp)`
  Called by: `spectra()` (same file), `Cspectra()` (same file)
- **`spectra()`** — L612 — `DWORD WINAPI spectra (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`Init_DetectMaxBin()`** — L688 — `void Init_DetectMaxBin(int disp)`
  Call in XCreateAnalyzer(...) This gets initialized for each 'disp' that is set up.
  Called by: `XCreateAnalyzer()` (same file)
- **`Destroy_DetectMaxBin()`** — L707 — `void Destroy_DetectMaxBin(int disp)`
  Call in DestroyAnalyzer(...)
  Called by: `DestroyAnalyzer()` (same file)
- **`calc_dmb()`** — L714 — `void calc_dmb(int disp, int size)`
  Called from SetupDetectMaxBin(...) AND anytime 'size' changes, e.g., in SetAnalyzer(...)
  Called by: `SetupDetectMaxBin()` (same file), `SetAnalyzer()` (same file)
- **`SetupDetectMaxBin()`** — L774 — `PORT void SetupDetectMaxBin(int run, int disp, int ss, int LO, double rate, double fLow, double fHigh, double tau, int frame_rate)`
  rate: Sample_rate of display data, e.g., '192000.0'. fLow: Lowest frequency of frequency range to evaluate, referenced to center_frequency to which DDC is tuned. For example, for LSB, not using CTUN, this might be '-3000.0'. fHigh: Highest frequency of frequency range to evaluate, referenced to…
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`DetectMaxBin()`** — L795 — `void DetectMaxBin(int disp, int ss, int LO)`
  Call this function in 'Cspectra(...)', after the FFT.
  Called by: `Cspectra()` (same file)
- **`GetDetectMaxBin()`** — L829 — `PORT double GetDetectMaxBin(int disp)`
  Call from console, for each 'disp' for which this function is desired. Always returns the value from the most recent display frame.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`Cspectra()`** — L846 — `DWORD WINAPI Cspectra (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`interpolate()`** — L928 — `void interpolate(int disp, int set, double fmin, double fmax, int num_pixels)`
  Called by: `SetAnalyzer()` (same file)
- **`build_interpolants()`** — L982 — `int build_interpolants(int disp, int set, int n, int m, double *x, double (*y)[dMAX_M])`
  Called by: `SetCalibration()` (same file)
- **`sendbuf()`** — L1066 — `void __cdecl sendbuf(void *arg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CalcBandwidthNormalization()`** — L1101 — `void CalcBandwidthNormalization (DP a)`
  Called by: `SetAnalyzer()` (same file), `SetDisplaySampleRate()` (same file)
- **`ResetPixelBuffers()`** — L1108 — `PORT void ResetPixelBuffers(int disp)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetAnalyzer()`** — L1188 — `PORT void SetAnalyzer ( int disp, int n_pixout, int n_fft, int typ,`
  Sets analyzer — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`XCreateAnalyzer()`** — L1339 — `PORT void XCreateAnalyzer( int disp, int *success, int m_size, int m_num_fft,`
  Called by: `alloc_analyzer()` (`ChannelMaster/analyzers.c`), `create_rcvr()` (`ChannelMaster/cmaster.c`), `create_xmtr()` (`ChannelMaster/cmaster.c`)
- **`DestroyAnalyzer()`** — L1448 — `PORT void DestroyAnalyzer(int disp)`
  Called by: `free_analyzer()` (`ChannelMaster/analyzers.c`), `destroy_rcvr()` (`ChannelMaster/cmaster.c`), `destroy_xmtr()` (`ChannelMaster/cmaster.c`)
- **`SetPixelRef()`** — L1529 — `PORT void SetPixelRef(int disp, double pixel_ref)`
  Sets pixel ref — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`GetPixels()`** — L1545 — `PORT void GetPixels ( int disp, int pixout, dOUTREAL *pix, int *flag,`
  Returns pixels — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SnapSpectrum()`** — L1588 — `PORT void SnapSpectrum( int disp, int ss, int LO, double *snap_buff)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SnapSpectrumTimeout()`** — L1600 — `PORT void SnapSpectrumTimeout(int disp, int ss, int LO, double* snap_buff,`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`calcompare()`** — L1621 — `int calcompare (const void * a, const void * b)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetCalibration()`** — L1631 — `PORT void SetCalibration ( int disp, int set_num, int n_points, double (*cal)[dMAX_M+1]`
  Sets calibration — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`OpenBuffer()`** — L1663 — `PORT void OpenBuffer(int disp, int ss, int LO, void **Ipointer, void **Qpointer)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CloseBuffer()`** — L1673 — `PORT void CloseBuffer(int disp, int ss, int LO)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`Spectrum()`** — L1702 — `PORT void Spectrum(int disp, int ss, int LO, dINREAL* pI, dINREAL* pQ)`
  Called by: `ReadUDPFrame()` (`ChannelMaster/network.c`)
- **`Spectrum2()`** — L1741 — `PORT void Spectrum2(int run, int disp, int ss, int LO, dINREAL* pbuff)`
  Called by: `xsender()` (`wdsp/sender.c`)
- **`Spectrum0()`** — L1787 — `PORT void Spectrum0(int run, int disp, int ss, int LO, double* pbuff)`
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`), `xpipe()` (`ChannelMaster/pipe.c`), `xsiphon()` (`wdsp/siphon.c`)
- **`SetDisplayDetectorMode()`** — L1833 — `PORT void SetDisplayDetectorMode (int disp, int pixout, int mode)`
  Sets display detector mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetDisplayAverageMode()`** — L1845 — `PORT void SetDisplayAverageMode (int disp, int pixout, int mode)`
  Sets display average mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetDisplayNumAverage()`** — L1877 — `PORT void SetDisplayNumAverage (int disp, int pixout, int num)`
  Sets display num average — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetDisplayAvBackmult()`** — L1892 — `PORT void SetDisplayAvBackmult (int disp, int pixout, double mult)`
  Sets display av backmult — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetDisplaySampleRate()`** — L1904 — `PORT void SetDisplaySampleRate (int disp, int rate)`
  Sets display sample rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetDisplayNormOneHz()`** — L1917 — `PORT void SetDisplayNormOneHz (int disp, int pixout, int norm)`
  Sets display norm one hz — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`GetDisplayENB()`** — L1929 — `PORT double GetDisplayENB (int disp)`
  Returns display enb — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/analyzer.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
