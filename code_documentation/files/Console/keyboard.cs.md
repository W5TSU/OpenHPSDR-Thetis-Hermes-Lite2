# `Console/keyboard.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Static key-state helpers over `GetKeyState`/`GetAsyncKeyState` P/Invoke; the async variant backs the spacebar PTT-hold release detection (shortcut handling itself lives in `console.cs` `Console_KeyDown`).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2)
  - `Console/MeterManager.cs` (calls ×1)
  - `Console/clsDBMan.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.IsKeyDown()` (×3), `.IsKeyDownAsync()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Keyboard` (type, L50)

- **`.GetKeyState()`** — L60 — `[DllImport("user32.dll", CharSet = CharSet.Auto, ExactSpelling = true)] private static extern short GetKeyState(int keyCode)`
  Returns key state.
  Called by: `.IsKeyDown()` (same file), `.IsKeyToggled()` (same file)
- **`.IsKeyDown()`** — L81 — `public static bool IsKeyDown(Keys key)`
  Called by: `.toggleSplit()` (`Console/MeterManager.cs`), `.LoadDB()` (`Console/clsDBMan.cs`), `.chkVFOSplit_MouseClick()` (`Console/console.cs`)
- **`.GetAsyncKeyState()`** — L86 — `[DllImport("user32.dll", CharSet = CharSet.Auto, ExactSpelling = true)] private static extern short GetAsyncKeyState(int keyCode)`
  Returns async key state.
  Called by: `.IsKeyDownAsync()` (same file)
- **`.IsKeyDownAsync()`** — L91 — `public static bool IsKeyDownAsync(Keys key)`
  W5TSU: physical key state - unlike GetKeyState this is not tied to the calling thread's input queue, so it stays correct when another window has keyboard focus
  Called by: `.spacebarHoldEngaged()` (`Console/console.cs`)
- **`.IsKeyToggled()`** — L96 — `public static bool IsKeyToggled(Keys key)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `KeyStates` (type, L52)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/keyboard.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
