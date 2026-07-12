# `Console/clsDPISafeTools.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** High-DPI/monitor-scaling safety helpers for WinForms layout.

## How this file is used

- Used by (incoming references from other files):
  - `Console/common.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.EnsureRectangleWithinNearestScreen()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `SafeScreens` (type, L49)

- **`.GetDpiForMonitor()`** — L54 — `[DllImport("Shcore.dll")] private static extern int GetDpiForMonitor(IntPtr hmonitor, MONITOR_DPI_TYPE dpiType, out uint dpiX, out uint dpiY)`
  Returns dpi for monitor.
  Called by: `.getMonitorInfosPhysical()` (same file)
- **`.EnumDisplayMonitors()`** — L57 — `[DllImport("user32.dll")] private static extern bool EnumDisplayMonitors(IntPtr hdc, IntPtr lprcClip, MonitorEnumProc lpfnEnum, IntPtr dwData)`
  Called by: `.getMonitorInfosPhysical()` (same file)
- **`.GetMonitorInfo()`** — L60 — `[DllImport("user32.dll", CharSet = CharSet.Auto)] private static extern bool GetMonitorInfo(IntPtr hMonitor, ref MONITORINFO lpmi)`
  Returns monitor info.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMonitorInfoEx()`** — L62 — `[DllImport("user32.dll", CharSet = CharSet.Auto, EntryPoint = "GetMonitorInfo")] private static extern bool GetMonitorInfoEx(IntPtr hMonitor, ref MONITORINFOEX lpmi)`
  Returns monitor info ex.
  Called by: `.getMonitorInfosPhysical()` (same file)
- **`.GetWindowRect()`** — L67 — `[DllImport("user32.dll")] private static extern bool GetWindowRect(IntPtr hWnd, out RECT lpRect)`
  Returns window rect.
  Called by: `.getDwmShadowMargins()` (same file), `.tryGetWindowRectPhysical()` (same file)
- **`.SetThreadDpiAwarenessContext()`** — L70 — `[DllImport("user32.dll")] private static extern IntPtr SetThreadDpiAwarenessContext(IntPtr dpiContext)`
  Sets thread dpi awareness context.
  Called by: `.tryGetWindowRectPhysical()` (same file), `.getMonitorInfosPhysical()` (same file)
- **`.DwmGetWindowAttribute()`** — L73 — `[DllImport("dwmapi.dll")] private static extern int DwmGetWindowAttribute(IntPtr hwnd, int dwAttribute, out RECT pvAttribute, int cbAttribute)`
  Called by: `.tryGetExtendedFrameBounds()` (same file)
- **`.EnsureRectangleWithinNearestScreen()`** — L128 — `public static (Rectangle adjusted, bool resized, bool repositioned) EnsureRectangleWithinNearestScreen(Rectangle? rect = null, Form form = null, bool keep_on_screen = false, bool u`
  Called by: `.ForceFormOnScreen()` (`Console/common.cs`)
- **`.getMonitorRects()`** — L212 — `private static List<Rectangle> getMonitorRects(bool use_working_area)`
  Returns monitor rects.
  Called by: `.EnsureRectangleWithinNearestScreen()` (same file)
- **`.chooseTargetMonitor()`** — L219 — `private static Rectangle chooseTargetMonitor(Rectangle r, List<Rectangle> monitors, bool keep_on_screen)`
  Called by: `.EnsureRectangleWithinNearestScreen()` (same file)
- **`.isFullyOnMonitors()`** — L261 — `private static bool isFullyOnMonitors(Rectangle r, List<Rectangle> monitors)`
  Called by: `.EnsureRectangleWithinNearestScreen()` (same file)
- **`.isContainedByAnyScreen()`** — L273 — `private static bool isContainedByAnyScreen(Rectangle r, List<Rectangle> monitors)`
  Called by: `.EnsureRectangleWithinNearestScreen()` (same file)
- **`.getUnion()`** — L282 — `private static Rectangle getUnion(List<Rectangle> rects)`
  Returns union.
  Called by: `.createScreensBitmap()` (same file)
- **`.clamp()`** — L298 — `private static int clamp(int v, int a, int b)`
  Called by: `.chooseTargetMonitor()` (same file)
- **`.colorFromHue()`** — L305 — `private static Color colorFromHue(double hue_deg, double saturation, double value)`
  Called by: `.createScreensBitmap()` (same file)
- **`.tryGetExtendedFrameBounds()`** — L328 — `private static bool tryGetExtendedFrameBounds(Form form, out Rectangle rect)`
  Called by: `.EnsureRectangleWithinNearestScreen()` (same file), `.getDwmShadowMargins()` (same file), `.createScreensBitmap()` (same file)
- **`.getDwmShadowMargins()`** — L339 — `private static void getDwmShadowMargins(Form form, out int left, out int top, out int right, out int bottom)`
  Returns dwm shadow margins.
  Called by: `.EnsureRectangleWithinNearestScreen()` (same file)
- **`.tryGetWindowRectPhysical()`** — L379 — `private static bool tryGetWindowRectPhysical(IntPtr hwnd, out RECT rect)`
  Called by: `.getDwmShadowMargins()` (same file), `.createScreensBitmap()` (same file)
- **`.createScreensBitmap()`** — L404 — `public static Bitmap createScreensBitmap(Size target_size, IEnumerable<Form> forms, bool use_working_area = false)`
  Called by: `.RenderScreensToPictureBox()` (same file)
- **`.RenderScreensToPictureBox()`** — L536 — `public static void RenderScreensToPictureBox(PictureBox picture_box, IEnumerable<Form> forms = null, bool use_working_area = false)`
  Renders screens to picture box.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getMonitorInfosPhysical()`** — L560 — `private static List<monitor_info> getMonitorInfosPhysical(bool use_working_area)`
  Returns monitor infos physical.
  Called by: `.createScreensBitmap()` (same file)
- **`.parseDisplayNumber()`** — L630 — `private static int parseDisplayNumber(string device)`
  Called by: `.getMonitorInfosPhysical()` (same file)

#### `RECT` (type, L77)

- **`.ToString()`** — L85 — `public override string ToString()`
  Returns the string representation.
  Called by: `.createScreensBitmap()` (same file)

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
