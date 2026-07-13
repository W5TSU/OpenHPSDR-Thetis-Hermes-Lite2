# `RawInput/KeyMapper.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Key mapping and event argument types.

## How this file is used

- Used by (incoming references from other files):
  - `RawInput/RawKeyboard.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetKeyName()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `KeyMapper` (type, L6)

- **`.GetKeyName()`** — L10 — `public static string GetKeyName(int value)`
  I prefer to have control over the key mapping This mapping could be loading from file to allow mapping changes without a recompile
  Called by: `.ProcessRawInput()` (`RawInput/RawKeyboard.cs`)
- **`.GetMicrosoftKeyName()`** — L203 — `public static string GetMicrosoftKeyName(int virtualKey)`
  If you prefer the virtualkey converted into a Microsoft virtualkey code use this
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/KeyMapper.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
