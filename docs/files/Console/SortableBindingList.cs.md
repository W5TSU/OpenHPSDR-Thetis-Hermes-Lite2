# `Console/SortableBindingList.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Sortable list base used by memory grids.

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/CATCommands.cs` (references ×1)
  - `Console/Memory/DXMemList.cs` (references ×1)
  - `Console/Memory/MemoryList.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `SortableBindingList` (type, L33)

- `.RemoveSortCore()` — L56
- `.ApplySortCore()` — L57

#### `PropertyComparer` (type, L78)

- `.Compare()` — L90

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/SortableBindingList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
