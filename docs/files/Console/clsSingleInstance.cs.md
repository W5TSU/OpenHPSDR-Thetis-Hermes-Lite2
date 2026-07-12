# `Console/clsSingleInstance.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Enforces a single running instance of Thetis.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.CheckAndPrompt()` (×1), `.Release()` (×1)

## Outline

### Types

#### `SingleInstance` (type, L47)

- `.CheckAndPrompt()` — L54
- `.Release()` — L97

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsSingleInstance.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
