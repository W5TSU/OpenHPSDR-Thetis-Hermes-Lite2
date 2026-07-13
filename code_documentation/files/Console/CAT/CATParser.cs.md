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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CATParser` (type, L42)

- **`.GetCATData()`** — L90 — `private void GetCATData()`
  Returns catdata.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Get()`** — L105 — `public string Get(byte[] pCmdString)`
  Overloaded Get method accepts either byte or string
  Called by: `.ExecuteCommand()` (`Console/CAT/CATTester.cs`), `ParseString()` (`Console/CAT/SIOListenerII.cs`), `.SerialRXEventHandler()` (`Console/CAT/SIOListenerII.cs`), `.SerialRX2EventHandler()` (`Console/CAT/SIOListenerII.cs`), `.SerialRX3EventHandler()` (`Console/CAT/SIOListenerII.cs`), `.SerialRX4EventHandler()` (`Console/CAT/SIOListenerII.cs`) — and 4 more
- **`.CheckFormat()`** — L421 — `private bool CheckFormat()`
  Checks format.
  Called by: `.Get()` (same file)
- **`.FindPrefix()`** — L468 — `private bool FindPrefix()`
  Finds prefix.
  Called by: `.CheckFormat()` (same file)
- **`.FindSuffix()`** — L538 — `private bool FindSuffix()`
  Finds suffix.
  Called by: `.CheckFormat()` (same file)
- **`.ParseExtended()`** — L596 — `private string ParseExtended()`
  Parses extended.
  Called by: `.Get()` (same file)
- **`.ProcessError()`** — L1595 — `private string ProcessError(string error)`
  Processes error.
  Called by: `.Get()` (same file), `.ParseExtended()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/CATParser.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
