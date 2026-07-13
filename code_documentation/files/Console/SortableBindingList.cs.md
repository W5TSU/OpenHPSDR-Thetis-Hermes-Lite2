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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `SortableBindingList` (type, L33)

- **`.RemoveSortCore()`** — L56 — `protected override void RemoveSortCore()`
  methods
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ApplySortCore()`** — L57 — `protected override void ApplySortCore(System.ComponentModel.PropertyDescriptor prop, System.ComponentModel.ListSortDirection direction)`
  Applys sort core.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `PropertyComparer` (type, L78)

- **`.Compare()`** — L90 — `public int Compare(T x, T y)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/SortableBindingList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
