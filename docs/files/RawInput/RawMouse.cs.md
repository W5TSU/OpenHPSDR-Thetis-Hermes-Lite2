# `RawInput/RawMouse.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Per-device keyboard and mouse message decoding.

## How this file is used

- Used by (incoming references from other files):
  - `RawInput/RawInput.cs` (references ×1)
- Uses (outgoing references to other files):
  - `RawInput/DataStructures.cs` (references ×1)

## Outline

### Types

#### `RawMouse` (type, L10)

- `.EnumerateDevices()` — L49
- `.ProcessRawInput()` — L128

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/RawMouse.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
