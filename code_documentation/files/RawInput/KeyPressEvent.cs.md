# `RawInput/KeyPressEvent.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Key mapping and event argument types.

## How this file is used

- Used by (incoming references from other files):
  - `RawInput/RawInputEventArg.cs` (references ×1)
  - `RawInput/Win32.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.ToString()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `KeyPressEvent` (type, L5)

- **`.ToString()`** — L24 — `public override string ToString()`
  Returns the string representation.
  Called by: `.DeviceAudit()` (`RawInput/Win32.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/KeyPressEvent.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
