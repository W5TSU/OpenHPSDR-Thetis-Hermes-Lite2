# `Console/clsProgressLog.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Diagnostic/status logging windows.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×6)
  - `Console/console.cs` (calls ×4)
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.Completed()` (×3), `.AddLogEntry()` (×2), `.ShowLog()` (×1), `.Shutdown()` (×1), `.HideAndSave()` (×1), `.SetRegistryToShow()` (×1), `.GetRegistryToShow()` (×1)

## Outline

### Types

#### `LogTool` (type, L51)

- `.SetWindowLong32()` — L53
- `.SetWindowLongPtr64()` — L56
- `.setWindowLongAuto()` — L59
- `.ShowScrollBar()` — L84
- `.ShowNewLog()` — L87
- `.ShowLog()` — L118
- `.setOwner()` — L129
- `.AddLogEntry()` — L137
- `.Shutdown()` — L167
- `.Completed()` — L179
- `.Finish()` — L220
- `.HideAndSave()` — L231
- `.SetRegistryToShow()` — L241
- `.GetRegistryToShow()` — L245
- `.addCore()` — L250
- `.ensureForm()` — L296
- `.runOnUiThreadSync()` — L324
- `.runOnUiThread()` — L331
- `.readRegistryShow()` — L338
- `.writeRegistryShow()` — L376
- `.tryReadLocation()` — L383
- `.writeLocation()` — L400
- `.hideHorizontalScrollBar()` — L408

#### `Entry` (type, L67)

_No extracted members._

#### `LogForm` (type, L413)

- `.OnFormClosing()` — L524
- `.layoutColumns()` — L530
- `.centerClose()` — L539

#### `NoSelectListView` (type, L415)

- `.SendMessage()` — L417
- `.OnCreateControl()` — L424
- `.OnItemSelectionChanged()` — L430
- `.OnGotFocus()` — L436
- `.OnMouseDown()` — L444
- `.OnKeyDown()` — L451

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsProgressLog.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
