# `Console/Memory/DXMemList.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Memory list variant used for DX cluster spot storage.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2, references ×1)
- Uses (outgoing references to other files):
  - `Console/SortableBindingList.cs` (references ×1)
- Most-referenced symbols from other files: `.Restore1()` (×1), `.CheckVersion1()` (×1)

## Outline

### Functions

- `.Restore1()` — L120

### Types

#### `DXMemList` (type, L38)

- `.Save1()` — L90
- `.CheckVersion1()` — L176

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Memory/DXMemList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
