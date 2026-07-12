# `Console/Memory/MemoryList.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Memory channel list UI and its record/list model (frequency, mode, filter, tones per memory).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×3, references ×1)
- Uses (outgoing references to other files):
  - `Console/Memory/MemoryRecord.cs` (references ×1)
  - `Console/SortableBindingList.cs` (references ×1)
- Most-referenced symbols from other files: `.Save()` (×1), `.Restore()` (×1), `.CheckVersion()` (×1)

## Outline

### Functions

- `.Restore()` — L123

### Types

#### `MemoryList` (type, L42)

- `.Save()` — L93
- `.CheckVersion()` — L179

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Memory/MemoryList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
