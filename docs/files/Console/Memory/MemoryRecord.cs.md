# `Console/Memory/MemoryRecord.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Memory channel list UI and its record/list model (frequency, mode, filter, tones per memory).

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/CATCommands.cs` (references ×2)
  - `Console/Memory/MemoryList.cs` (references ×1)
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×4)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `MemoryRecord` (type, L37)

- **`.OnPropertyChanged()`** — L171 — `private void OnPropertyChanged(object sender, PropertyChangedEventArgs e)`
  Handles/raises the property changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CompareTo()`** — L526 — `public int CompareTo(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Memory/MemoryRecord.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
