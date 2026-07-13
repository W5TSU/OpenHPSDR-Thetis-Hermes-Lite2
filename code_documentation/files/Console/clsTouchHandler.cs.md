# `Console/clsTouchHandler.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Application-wide mouse message filtering and touch-screen gesture support.

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.DisableTouchSupport()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `TouchHandler` (type, L54)

- **`.SetWindowLongPtr()`** — L99 — `[DllImport("user32.dll")] private static extern IntPtr SetWindowLongPtr(IntPtr hWnd, int nIndex, IntPtr dwNewLong)`
  Sets window long ptr.
  Called by: `.EnableTouchSupport()` (same file), `.DisableTouchSupport()` (same file)
- **`.CallWindowProc()`** — L102 — `[DllImport("user32.dll")] private static extern IntPtr CallWindowProc(IntPtr lpPrevWndFunc, IntPtr hWnd, uint Msg, IntPtr wParam, IntPtr lParam)`
  Called by: `.CustomWndProc()` (same file)
- **`.GetTouchInputInfo()`** — L105 — `[DllImport("user32.dll")] private static extern bool GetTouchInputInfo(IntPtr hTouchInput, int cInputs, [Out] TOUCHINPUT[] pInputs, int cbSize)`
  Returns touch input info.
  Called by: `.CustomWndProc()` (same file)
- **`.CloseTouchInputHandle()`** — L108 — `[DllImport("user32.dll")] private static extern void CloseTouchInputHandle(IntPtr lParam)`
  Closes touch input handle.
  Called by: `.CustomWndProc()` (same file)
- **`.RegisterTouchWindow()`** — L111 — `[DllImport("user32.dll")] private static extern bool RegisterTouchWindow(IntPtr hWnd, uint ulFlags)`
  Registers touch window.
  Called by: `.EnableTouchSupport()` (same file)
- **`.UnregisterTouchWindow()`** — L114 — `[DllImport("user32.dll")] private static extern bool UnregisterTouchWindow(IntPtr hWnd)`
  Unregisters touch window.
  Called by: `.DisableTouchSupport()` (same file)
- **`.ScreenToClient()`** — L124 — `[DllImport("user32.dll")] private static extern bool ScreenToClient(IntPtr hWnd, ref POINT lpPoint)`
  Called by: `.CustomWndProc()` (same file)
- **`.EnableTouchSupport()`** — L127 — `public static Guid EnableTouchSupport(Control control, Action<int, int> touchDown, Action<int, int> touchMove, Action<int, int> touchUp, int touch_mask, string id = "")`
  Enables touch support.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisableTouchSupport()`** — L165 — `public static void DisableTouchSupport(Guid id)`
  Disables touch support.
  Called by: `.ShutdownDX()` (`Console/MeterManager.cs`)
- **`.CustomWndProc()`** — L183 — `private static IntPtr CustomWndProc(IntPtr hWnd, uint msg, IntPtr wParam, IntPtr lParam)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `ControlTouchInfo` (type, L56)

_No extracted members._

#### `TOUCHINPUT` (type, L84)

_No extracted members._

#### `POINT` (type, L117)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsTouchHandler.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
