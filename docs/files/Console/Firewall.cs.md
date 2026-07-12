# `Console/Firewall.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Adds Windows Firewall rules so radio UDP traffic is not blocked.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.Setup()` (×1)

## Outline

### Types

#### `Firewall` (type, L77)

- `.Setup()` — L84
- `.findRule()` — L138
- `.addApplicationRule()` — L155

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Firewall.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
