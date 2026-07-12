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

### Types

#### `HiPerfTimer` (type, L52)

- `.QueryPerformanceCounter()` — L54
- `.QueryPerformanceFrequency()` — L58
- `.Start()` — L80
- `.Stop()` — L89
- `.Reset()` — L118
- `.GetFreq()` — L130

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/hiperftimer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
