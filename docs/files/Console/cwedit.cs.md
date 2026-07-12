# `Console/cwedit.cs`

**Functional area:** [11. CW keying](../../CODE_OUTLINE.md#11-cw-keying)

**Role:** Editor for CWX stored messages.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Types

#### `cwedit` (type, L50)

- `.Dispose()` — L82
- `.InitializeComponent()` — L100
- `.extract_fields()` — L252
- `.make_current()` — L260
- `.cwedit_Load()` — L266
- `.saveButton_Click()` — L279
- `.cancelButton_Click()` — L285
- `.slen()` — L297
- `.txtComments_Leave()` — L304
- `.txtElements_Leave()` — L317

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/cwedit.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
