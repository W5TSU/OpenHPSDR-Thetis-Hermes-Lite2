# `Console/CAT/CATParser.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Tokenizes/validates incoming CAT strings and dispatches to `CATCommands`.

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/SIOListenerII.cs` (calls ×8, references ×7)
  - `Console/CAT/CATTester.cs` (references ×1, calls ×1)
  - `Console/console.cs` (references ×1, calls ×1)
  - `Console/CAT/CATCommands.cs` (references ×1)
  - `Console/Midi2CatCommands.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/CAT/CATCommands.cs` (calls ×350)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
- Most-referenced symbols from other files: `.Get()` (×10)

## Outline

### Types

#### `CATParser` (type, L42)

- `.GetCATData()` — L90
- `.Get()` — L105
- `.CheckFormat()` — L421
- `.FindPrefix()` — L468
- `.FindSuffix()` — L538
- `.ParseExtended()` — L596
- `.ProcessError()` — L1595

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/CATParser.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
