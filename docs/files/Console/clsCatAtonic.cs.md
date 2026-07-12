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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

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

- **`.skip_ws_and_comments()`** — L115 — `private void skip_ws_and_comments()`
  Called by: `.next()` (same file)
- **`.next()`** — L131 — `public Token next()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `CATScriptInterpreter` (type, L183)

- **`.run()`** — L197 — `public ScriptResult run(string script)`
  Called by: `.OnContainerVisible()` (`Console/MeterManager.cs`), `.SetMacroSettings()` (`Console/MeterManager.cs`), `.handleMacroButtonPress()` (`Console/MeterManager.cs`), `.txtCatMacro_TextChanged()` (`Console/frmMacroButtonConfig.cs`)
- **`.filter_now()`** — L433 — `public List<ScriptCommand> filter_now(ScriptResult r, Func<string, bool> eval)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.total_wait_milliseconds_now()`** — L457 — `public int total_wait_milliseconds_now(ScriptResult r, Func<string, bool> eval)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cat_state_command_now()`** — L472 — `public string cat_state_command_now(ScriptResult r, Func<string, bool> eval)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.guard_holds()`** — L486 — `private static bool guard_holds(ScriptCommand c, Func<string, bool> eval)`
  Called by: `.filter_now()` (same file), `.total_wait_milliseconds_now()` (same file), `.cat_state_command_now()` (same file)
- **`.set_guard_from_stack()`** — L513 — `private static void set_guard_from_stack(ScriptCommand cmd, List<if_ctx> stack)`
  Sets guard from stack.
  Called by: `.run()` (same file)
- **`.contains_string()`** — L561 — `private static bool contains_string(List<string> list, string s)`
  Called by: `.set_guard_from_stack()` (same file)

#### `if_ctx` (type, L185)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCatAtonic.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
