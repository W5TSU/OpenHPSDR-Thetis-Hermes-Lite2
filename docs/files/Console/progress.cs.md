# `Console/progress.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Startup splash screen and progress reporting during initialization.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×4, calls ×3)
  - `Console/setup.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Invoke/buttonts.cs` (references ×1)
- Most-referenced symbols from other files: `.SetPercent()` (×3)

## Outline

### Types

#### `Progress` (type, L51)

- `.Dispose()` — L73
- `.InitializeComponent()` — L89
- `.SetPercent()` — L135
- `.btnAbort_Click()` — L145
- `.panel1_Paint()` — L170
- `.Progress_Closing()` — L200

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/progress.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
