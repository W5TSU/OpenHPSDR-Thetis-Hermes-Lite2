# `Console/clsCatAtonic.cs`

**Functional area:** [10. CAT control and external program interfaces](../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Queued/asynchronous CAT message handling and scripted ("atomic") CAT command sequences.

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (references ×4, calls ×3)
  - `Console/clsCATMessageQueue.cs` (references ×2)
  - `Console/frmMacroButtonConfig.cs` (references ×1, calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.run()` (×4)

## Outline

### Types

#### `CatAtonic` (namespace, L45)

_No extracted members._

#### `ScriptCommandType` (type, L47)

_No extracted members._

#### `ScriptCommand` (type, L55)

_No extracted members._

#### `ScriptResult` (type, L68)

_No extracted members._

#### `TokenType` (type, L82)

_No extracted members._

#### `Token` (type, L90)

_No extracted members._

#### `Tokeniser` (type, L102)

- `.skip_ws_and_comments()` — L115
- `.next()` — L131

#### `CATScriptInterpreter` (type, L183)

- `.run()` — L197
- `.filter_now()` — L433
- `.total_wait_milliseconds_now()` — L457
- `.cat_state_command_now()` — L472
- `.guard_holds()` — L486
- `.set_guard_from_stack()` — L513
- `.contains_string()` — L561

#### `if_ctx` (type, L185)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCatAtonic.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
