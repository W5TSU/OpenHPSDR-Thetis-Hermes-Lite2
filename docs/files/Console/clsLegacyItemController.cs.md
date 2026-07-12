# `Console/clsLegacyItemController.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Maps legacy/renamed UI items to their current equivalents so old databases and skins keep working.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
- Most-referenced symbols from other files: `.Update()` (×2)

## Outline

### Types

#### `LegacyItemController` (type, L49)

- `.Init()` — L90
- `.Update()` — L309

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsLegacyItemController.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
