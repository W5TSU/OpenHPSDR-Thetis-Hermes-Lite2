# `Console/hiperftimer.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** High-resolution performance timer used for timing-critical UI work.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×14, references ×1)
  - `Console/PanDisplay.cs` (calls ×2, references ×1)
  - `Console/MeterManager.cs` (references ×1, calls ×1)
  - `Console/common.cs` (references ×1, calls ×1)
  - `Console/cwkeyer.cs` (calls ×2)
  - `Console/display.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Stop()` (×8), `.Start()` (×7), `.Reset()` (×5)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `HiPerfTimer` (type, L52)

- **`.QueryPerformanceCounter()`** — L54 — `[DllImport("Kernel32.dll")] private static extern bool QueryPerformanceCounter( out long lpPerformanceCount)`
  Called by: `.Start()` (same file), `.Stop()` (same file)
- **`.QueryPerformanceFrequency()`** — L58 — `[DllImport("Kernel32.dll")] private static extern bool QueryPerformanceFrequency( out long lpFrequency)`
  Called by: `.GetFreq()` (same file)
- **`.Start()`** — L80 — `public void Start()`
  Start the timer
  Called by: `.Reset()` (same file), `.DrawWaterfall()` (`Console/PanDisplay.cs`), `.ProcessCPUUsage()` (`Console/common.cs`), `.picMultiMeterDigital_Paint()` (`Console/console.cs`), `.picRX2Meter_Paint()` (`Console/console.cs`), `.UpdateMultimeter()` (`Console/console.cs`) — and 2 more
- **`.Stop()`** — L89 — `public void Stop()`
  Stop the timer
  Called by: `.DrawWaterfall()` (`Console/PanDisplay.cs`), `.storeRX1SignalPixels_X()` (`Console/console.cs`), `.storeRX2SignalPixels_X()` (`Console/console.cs`), `.picMultiMeterDigital_Paint()` (`Console/console.cs`), `.picRX2Meter_Paint()` (`Console/console.cs`), `.MultiMeter2UpdateRX1()` (`Console/console.cs`) — and 2 more
- **`.Reset()`** — L118 — `public void Reset()`
  Called by: `.dxRender()` (`Console/MeterManager.cs`), `.clearRXSignalPixels()` (`Console/console.cs`), `.RunDisplay()` (`Console/console.cs`), `.MultiMeter2UpdateRX1()` (`Console/console.cs`), `.MultiMeter2UpdateRX2()` (`Console/console.cs`)
- **`.GetFreq()`** — L130 — `public long GetFreq()`
  Returns freq.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/hiperftimer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
