# `Console/titlebar.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Custom title bar text/version display.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×1)
  - `Console/database.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)
- Most-referenced symbols from other files: `.GetString()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `TitleBar` (type, L51)

- **`.GetString()`** — L54 — `public static string GetString(bool bWithFirmware = true)`
  Returns string.
  Called by: `.getTitleWithFWVersion()` (`Console/console.cs`), `.Init()` (`Console/database.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/titlebar.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
