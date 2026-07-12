# `Console/Dumpcap.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Drives Wireshark's `dumpcap` to capture radio network traffic for debugging.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×5)
  - `Console/frmSeqLog.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
- Most-referenced symbols from other files: `.StopDumpcap()` (×2), `.DumpCapExists()` (×1), `.Initalise()` (×1), `.ClearDumpFolder()` (×1), `.StartDumpcap()` (×1), `.ShowAppPathFolder()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `DumpCap` (type, L52)

- **`.DumpCapExists()`** — L66 — `public static bool DumpCapExists()`
  Called by: `.dumpcapGO()` (same file), `.setupControlsDumpCap()` (`Console/frmSeqLog.cs`)
- **`.restartDumpcap()`** — L126 — `private static void restartDumpcap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Initalise()`** — L135 — `public static void Initalise(Console c)`
  Called by: `.InitConsole()` (`Console/console.cs`)
- **`.ClearDumpFolder()`** — L146 — `public static void ClearDumpFolder()`
  Clears dump folder.
  Called by: `.InitConsole()` (`Console/console.cs`)
- **`.dumpcapGO()`** — L163 — `private static void dumpcapGO()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StartDumpcap()`** — L203 — `public static void StartDumpcap(int nTimeOut)`
  Starts dumpcap.
  Called by: `.restartDumpcap()` (same file), `.checkSeqErrors()` (`Console/console.cs`)
- **`.StopDumpcap()`** — L221 — `public static void StopDumpcap()`
  Stops dumpcap.
  Called by: `.restartDumpcap()` (same file), `.checkSeqErrors()` (`Console/console.cs`), `.Console_Closing()` (`Console/console.cs`)
- **`.isDumpcapRunning()`** — L241 — `private static bool isDumpcapRunning()`
  Called by: `.restartDumpcap()` (same file), `.dumpcapGO()` (same file), `.StartDumpcap()` (same file), `.StopDumpcap()` (same file)
- **`.ShowAppPathFolder()`** — L260 — `public static void ShowAppPathFolder()`
  Shows app path folder.
  Called by: `.btnShowDumpCapFolder_Click()` (`Console/frmSeqLog.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Dumpcap.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
