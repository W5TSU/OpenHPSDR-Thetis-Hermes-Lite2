# `Console/keyboard.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Keyboard shortcut handling (tune steps, band/mode changes, CW from keyboard).

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (calls ×1)
  - `Console/clsDBMan.cs` (calls ×1)
  - `Console/console.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.IsKeyDown()` (×3)

## Outline

### Types

#### `Keyboard` (type, L50)

- `.GetKeyState()` — L60
- `.IsKeyDown()` — L81
- `.IsKeyToggled()` — L86

#### `KeyStates` (type, L52)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/keyboard.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
