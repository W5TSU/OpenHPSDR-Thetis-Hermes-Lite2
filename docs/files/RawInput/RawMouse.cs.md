# `RawInput/RawMouse.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Per-device keyboard and mouse message decoding.

## How this file is used

- Used by (incoming references from other files):
  - `RawInput/RawInput.cs` (references ×1)
- Uses (outgoing references to other files):
  - `RawInput/DataStructures.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `RawMouse` (type, L10)

- **`.EnumerateDevices()`** — L49 — `public void EnumerateDevices(string id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ProcessRawInput()`** — L128 — `public void ProcessRawInput(IntPtr hdevice)`
  Processes raw input.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/RawMouse.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
