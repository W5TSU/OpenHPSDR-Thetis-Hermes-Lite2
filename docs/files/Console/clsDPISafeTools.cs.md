# `Console/clsDPISafeTools.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** High-DPI/monitor-scaling safety helpers for WinForms layout.

## How this file is used

- Used by (incoming references from other files):
  - `Console/common.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.EnsureRectangleWithinNearestScreen()` (×1)

## Outline

### Types

#### `SafeScreens` (type, L49)

- `.GetDpiForMonitor()` — L54
- `.EnumDisplayMonitors()` — L57
- `.GetMonitorInfo()` — L60
- `.GetMonitorInfoEx()` — L62
- `.GetWindowRect()` — L67
- `.SetThreadDpiAwarenessContext()` — L70
- `.DwmGetWindowAttribute()` — L73
- `.EnsureRectangleWithinNearestScreen()` — L128
- `.getMonitorRects()` — L212
- `.chooseTargetMonitor()` — L219
- `.isFullyOnMonitors()` — L261
- `.isContainedByAnyScreen()` — L273
- `.getUnion()` — L282
- `.clamp()` — L298
- `.colorFromHue()` — L305
- `.tryGetExtendedFrameBounds()` — L328
- `.getDwmShadowMargins()` — L339
- `.tryGetWindowRectPhysical()` — L379
- `.createScreensBitmap()` — L404
- `.RenderScreensToPictureBox()` — L536
- `.getMonitorInfosPhysical()` — L560
- `.parseDisplayNumber()` — L630

#### `RECT` (type, L77)

- `.ToString()` — L85

#### `MONITOR_DPI_TYPE` (type, L91)

_No extracted members._

#### `MONITORINFO` (type, L99)

_No extracted members._

#### `MONITORINFOEX` (type, L108)

_No extracted members._

#### `monitor_info` (type, L119)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsDPISafeTools.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
