# `Console/Channel.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Simple channel object used by scanning/memory features.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×5)
  - `Console/display.cs` (references ×1)
  - `Console/PanDisplay.cs` (references ×1)
  - `Console/wbDisplay.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.Copy()`** — L109 — `public Channel Copy()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

### Types

#### `Channel` (type, L46)

- **`.ToString()`** — L98 — `public override string ToString()`
  Displays the Channel details in a string
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CompareTo()`** — L103 — `public int CompareTo(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InBW()`** — L114 — `public bool InBW(double low, double high)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Channel.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
