# `Console/TimeOutTimerManager.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Transmit time-out timer (limits continuous TX time).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×3)
  - `Console/setup.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
- Most-referenced symbols from other files: `.Shutdown()` (×1), `.SetCallback()` (×1), `.RemoveCallback()` (×1), `.PingTimeOut()` (×1), `.MoxTimeOut()` (×1)

## Outline

### Types

#### `TimeOutTimerManager` (type, L53)

- `.Initialise()` — L77
- `.Shutdown()` — L100
- `.SetCallback()` — L110
- `.RemoveCallback()` — L115
- `.PingTimeOut()` — L119
- `.MoxTimeOut()` — L128
- `.startSecondTicker()` — L136
- `.onMox()` — L146
- `.tickLoop()` — L159

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/TimeOutTimerManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
