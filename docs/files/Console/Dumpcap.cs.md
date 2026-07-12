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

### Types

#### `DumpCap` (type, L52)

- `.DumpCapExists()` — L66
- `.restartDumpcap()` — L126
- `.Initalise()` — L135
- `.ClearDumpFolder()` — L146
- `.dumpcapGO()` — L163
- `.StartDumpcap()` — L203
- `.StopDumpcap()` — L221
- `.isDumpcapRunning()` — L241
- `.ShowAppPathFolder()` — L260

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Dumpcap.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
