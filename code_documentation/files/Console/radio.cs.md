# `Console/radio.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Radio/receiver object model — bands, modes, filter presets per mode, and per-RX DSP state that the console manipulates.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×91, references ×8)
  - `Console/setup.cs` (calls ×3)
  - `Console/display.cs` (references ×1, calls ×1)
  - `Console/MeterManager.cs` (calls ×1)
  - `Console/database.cs` (calls ×1)
  - `Console/frmNotchPopup.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×6)
- Most-referenced symbols from other files: `.GetDSPRX()` (×45), `.GetDSPTX()` (×27), `.NotchFromIndex()` (×6), `.IndexOf()` (×5), `.GetFirstNotchThatMatches()` (×4), `.NotchesInBW()` (×2), `.Parse()` (×2), `.Add()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.Copy()`** — L282 — `public void Copy(RadioDSPRX rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Parse()`** — L4358 — `public static MNotch Parse(string s)`
  Called by: `.GetState()` (`Console/console.cs`), `.ImportAndMergeDatabase()` (`Console/database.cs`)

### Types

#### `Radio` (type, L54)

- **`.Shutdown()`** — L80 — `public void Shutdown()`
  Called by: `.ExitConsole()` (`Console/console.cs`)
- **`.GetDSPRX()`** — L84 — `public RadioDSPRX GetDSPRX(int thread, int subrx)`
  Returns dsprx.
  Called by: `.SyncDSP()` (`Console/console.cs`), `.UpdateRX1Filters()` (`Console/console.cs`), `.UpdateRX2Filters()` (`Console/console.cs`), `.UpdateRXDisplayVars()` (`Console/console.cs`), `.VFOASubUpdate()` (`Console/console.cs`), `.chkPower_CheckedChanged()` (`Console/console.cs`) — and 39 more
- **`.GetDSPTX()`** — L89 — `public RadioDSPTX GetDSPTX(int thread)`
  Returns dsptx.
  Called by: `.SyncDSP()` (`Console/console.cs`), `.UpdateTXSpectrumDisplayVars()` (`Console/console.cs`), `.UpdateTXDisplayVars()` (`Console/console.cs`), `.UpdateTXLowHighFilterForMode()` (`Console/console.cs`), `.SetTXFilters()` (`Console/console.cs`), `.ptbMic_Scroll()` (`Console/console.cs`) — and 21 more

#### `RadioDSP` (type, L99)

- **`.CreateDSP()`** — L103 — `public static void CreateDSP()`
  Creates dsp.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DestroyDSP()`** — L163 — `public static void DestroyDSP()`
  Called by: `.Shutdown()` (same file)

#### `RadioDSPRX` (type, L270)

- **`.SyncAll()`** — L388 — `private void SyncAll()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXFilter()`** — L597 — `public void SetRXFilter(int low, int high)`
  Sets rxfilter.
  Called by: `.Copy()` (same file), `.SyncAll()` (same file)
- **`.SetNRVals()`** — L683 — `public void SetNRVals(int taps, int delay, double gain, double leak)`
  Sets nrvals.
  Called by: `.Copy()` (same file), `.SyncAll()` (same file)
- **`.SetANFVals()`** — L730 — `public void SetANFVals(int taps, int delay, double gain, double leak)`
  Sets anfvals.
  Called by: `.Copy()` (same file), `.SyncAll()` (same file)
- **`.GetNotchOn()`** — L1449 — `public bool GetNotchOn(int index)`
  Returns notch on.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNotchOn()`** — L1454 — `public void SetNotchOn(uint index, bool b)`
  Sets notch on.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNotchFreq()`** — L1468 — `public double GetNotchFreq(uint index)`
  Returns notch freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNotchFreq()`** — L1473 — `public void SetNotchFreq(uint index, double freq)`
  Sets notch freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNotchBW()`** — L1487 — `public double GetNotchBW(uint index)`
  Returns notch bw.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNotchBW()`** — L1497 — `public void SetNotchBW(uint index, double bw)`
  Sets the notch bandwidth
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `RadioDSPTX` (type, L2471)

- **`.SyncAll()`** — L2480 — `private void SyncAll()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXFilter()`** — L2732 — `public void SetTXFilter(int low, int high)`
  Sets txfilter.
  Called by: `.SyncAll()` (same file)

#### `MNotchDB` (type, L4192)

- **`.Clear()`** — L4205 — `public static void Clear()`
  get { lock (_listLock) { return _lstNotches; } } }
  Called by: `.SaveNotchesToDatabase()` (`Console/setup.cs`)
- **`.IndexOf()`** — L4212 — `public static int IndexOf(MNotch mNotch)`
  Called by: `.ChangeNotchBW()` (`Console/console.cs`), `.ChangeNotchCentreFrequency()` (`Console/console.cs`), `.changeNotchActive()` (`Console/console.cs`), `.toggleNotchActive()` (`Console/console.cs`), `.removeNotch()` (`Console/console.cs`)
- **`.Add()`** — L4219 — `public static void Add(MNotch mNotch)`
  Called by: `.GetState()` (`Console/console.cs`), `.SaveNotchesToDatabase()` (`Console/setup.cs`)
- **`.NotchFromIndex()`** — L4236 — `public static MNotch NotchFromIndex(int index)`
  Called by: `.onNotchDelete()` (`Console/console.cs`), `.onBWChanged()` (`Console/console.cs`), `.onActiveChanged()` (`Console/console.cs`), `.GetStateList()` (`Console/console.cs`), `.ShowNotchPopup()` (`Console/console.cs`), `.RestoreNotchesFromDatabase()` (`Console/setup.cs`)
- **`.GetFirstNotchThatMatches()`** — L4246 — `public static MNotch GetFirstNotchThatMatches(double freqHz, double fwidth, bool bActive)`
  MW0LGE return a notch that matches
  Called by: `.ChangeNotchBW()` (`Console/console.cs`), `.ChangeNotchCentreFrequency()` (`Console/console.cs`), `.changeNotchActive()` (`Console/console.cs`), `.toggleNotchActive()` (`Console/console.cs`)
- **`.NotchNearFreq()`** — L4261 — `public static bool NotchNearFreq(double freqHz, int deltaHz)`
  MW0LGE check if notch close by
  Called by: `.AddNotch()` (`Console/console.cs`)
- **`.NotchesInBW()`** — L4276 — `public static List<MNotch> NotchesInBW(double centreBWFreqHz, int lowHz, int highHz)`
  MW0LGE return list of notches in given bandwidth notch is included if filter width is enough to be within the BW
  Called by: `.NotchThatSurroundsFrequencyInBW()` (same file), `.Init()` (`Console/MeterManager.cs`), `.handleNotches()` (`Console/display.cs`)
- **`.NotchThatSurroundsFrequencyInBW()`** — L4297 — `public static MNotch NotchThatSurroundsFrequencyInBW(double centreBWFreqHz, int lowHz, int highHz, double freqHz, int nPadWidth = 0)`
  MW0LGE return first notch found that surrounds a given frequency in the given bandwidth
  Called by: `.pnlDisplay_MouseMove()` (`Console/console.cs`)

#### `MNotch` (type, L4328)

- **`.ToString()`** — L4376 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CompareTo()`** — L4381 — `public int CompareTo(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/radio.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
