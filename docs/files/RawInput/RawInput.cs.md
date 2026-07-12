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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `RawInput` (type, L18)

- **`.MouseDevices()`** — L50 — `public Dictionary<IntPtr, MouseEvent> MouseDevices()`
  Called by: `.updateRawInputDevices()` (`Console/console.cs`)
- **`.AddMessageFilter()`** — L73 — `public void AddMessageFilter()`
  Adds message filter.
  Called by: `.initialiseRawInput()` (`Console/console.cs`)
- **`.RemoveMessageFilter()`** — L81 — `public void RemoveMessageFilter()`
  Removes message filter.
  Called by: `.initialiseRawInput()` (`Console/console.cs`)
- **`.RegisterForDeviceNotifications()`** — L105 — `static IntPtr RegisterForDeviceNotifications(IntPtr parent)`
  Registers for device notifications.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WndProc()`** — L138 — `protected override void WndProc(ref Message message)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/RawInput.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
