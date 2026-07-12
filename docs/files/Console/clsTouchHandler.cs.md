# `Console/clsTouchHandler.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Application-wide mouse message filtering and touch-screen gesture support.

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.DisableTouchSupport()` (×1)

## Outline

### Types

#### `TouchHandler` (type, L54)

- `.SetWindowLongPtr()` — L99
- `.CallWindowProc()` — L102
- `.GetTouchInputInfo()` — L105
- `.CloseTouchInputHandle()` — L108
- `.RegisterTouchWindow()` — L111
- `.UnregisterTouchWindow()` — L114
- `.ScreenToClient()` — L124
- `.EnableTouchSupport()` — L127
- `.DisableTouchSupport()` — L165
- `.CustomWndProc()` — L183

#### `ControlTouchInfo` (type, L56)

_No extracted members._

#### `TOUCHINPUT` (type, L84)

_No extracted members._

#### `POINT` (type, L117)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsTouchHandler.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
