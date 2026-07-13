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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.Restore1()`** — L120 — `public static DXMemList Restore1()`
  Called by: `.InitConsole()` (`Console/console.cs`)

### Types

#### `DXMemList` (type, L38)

- **`.Save1()`** — L90 — `private void Save1(string file_name)`
  Called by: `.Restore1()` (same file)
- **`.CheckVersion1()`** — L176 — `public void CheckVersion1()`
  Checks version1.
  Called by: `.InitConsole()` (`Console/console.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Memory/DXMemList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
