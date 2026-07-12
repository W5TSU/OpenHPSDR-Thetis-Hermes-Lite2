# `RawInput/RawInput.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Entry point: registers for raw input and dispatches device events.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×3, references ×1)
- Uses (outgoing references to other files):
  - `RawInput/MouseEvent.cs` (references ×1)
  - `RawInput/PreMessageFilter.cs` (references ×1)
  - `RawInput/RawKeyboard.cs` (references ×1)
  - `RawInput/RawMouse.cs` (references ×1)
- Most-referenced symbols from other files: `.AddMessageFilter()` (×1), `.RemoveMessageFilter()` (×1), `.MouseDevices()` (×1)

## Outline

### Types

#### `RawInput` (type, L18)

- `.MouseDevices()` — L50
- `.AddMessageFilter()` — L73
- `.RemoveMessageFilter()` — L81
- `.RegisterForDeviceNotifications()` — L105
- `.WndProc()` — L138

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/RawInput.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
