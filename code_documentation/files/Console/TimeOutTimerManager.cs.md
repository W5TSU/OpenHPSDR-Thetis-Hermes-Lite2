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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `TimeOutTimerManager` (type, L53)

- **`.Initialise()`** — L77 — `public static void Initialise(Console c)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Shutdown()`** — L100 — `public static void Shutdown()`
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.SetCallback()`** — L110 — `public static void SetCallback(ToTOccured cb)`
  Sets callback.
  Called by: `.addDelegates()` (`Console/console.cs`)
- **`.RemoveCallback()`** — L115 — `public static void RemoveCallback(ToTOccured cb)`
  Removes callback.
  Called by: `.removeDelegates()` (`Console/console.cs`)
- **`.PingTimeOut()`** — L119 — `public static void PingTimeOut(string hostAddress, int timeOutSeconds, bool enabled)`
  Called by: `.chkToTPing_CheckedChanged()` (`Console/setup.cs`)
- **`.MoxTimeOut()`** — L128 — `public static void MoxTimeOut(int timeOutSeconds, bool enabled)`
  Called by: `.chkToTMox_CheckedChanged()` (`Console/setup.cs`)
- **`.startSecondTicker()`** — L136 — `private static void startSecondTicker()`
  Called by: `.Initialise()` (same file)
- **`.onMox()`** — L146 — `public static void onMox(int rx, bool oldMox, bool newMox)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tickLoop()`** — L159 — `private static void tickLoop()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/TimeOutTimerManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
