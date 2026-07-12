# `Console/clsProgressLog.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Diagnostic/status logging windows.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×6)
  - `Console/console.cs` (calls ×4)
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.Completed()` (×3), `.AddLogEntry()` (×2), `.ShowLog()` (×1), `.Shutdown()` (×1), `.HideAndSave()` (×1), `.SetRegistryToShow()` (×1), `.GetRegistryToShow()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `LogTool` (type, L51)

- **`.SetWindowLong32()`** — L53 — `[DllImport("user32.dll", EntryPoint = "SetWindowLong")] static extern IntPtr SetWindowLong32(IntPtr hWnd, int nIndex, IntPtr dwNewLong)`
  Sets window long32.
  Called by: `.setWindowLongAuto()` (same file)
- **`.SetWindowLongPtr64()`** — L56 — `[DllImport("user32.dll", EntryPoint = "SetWindowLongPtr")] static extern IntPtr SetWindowLongPtr64(IntPtr hWnd, int nIndex, IntPtr dwNewLong)`
  Sets window long ptr64.
  Called by: `.setWindowLongAuto()` (same file)
- **`.setWindowLongAuto()`** — L59 — `static IntPtr setWindowLongAuto(IntPtr hWnd, int nIndex, IntPtr dwNewLong)`
  Sets window long auto.
  Called by: `.setOwner()` (same file)
- **`.ShowScrollBar()`** — L84 — `[DllImport("user32.dll")] static extern bool ShowScrollBar(IntPtr hWnd, int wBar, bool bShow)`
  Shows scroll bar.
  Called by: `.hideHorizontalScrollBar()` (same file), `.layoutColumns()` (same file)
- **`.ShowNewLog()`** — L87 — `public static void ShowNewLog(IntPtr ownerHandle)`
  Shows new log.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowLog()`** — L118 — `public static void ShowLog(IntPtr ownerHandle)`
  Shows log.
  Called by: `.chkShowStartupLog_CheckedChanged()` (`Console/setup.cs`)
- **`.setOwner()`** — L129 — `private static void setOwner(IntPtr ownerHandle)`
  Sets owner.
  Called by: `.ShowNewLog()` (same file), `.ShowLog()` (same file)
- **`.AddLogEntry()`** — L137 — `public static string AddLogEntry(string text)`
  Adds log entry.
  Called by: `.InitConsole()` (`Console/console.cs`), `.AfterConstructor()` (`Console/setup.cs`)
- **`.Shutdown()`** — L167 — `public static void Shutdown()`
  Called by: `.Dispose()` (`Console/console.cs`)
- **`.Completed()`** — L179 — `public static void Completed(string id)`
  Called by: `.initialisePortAudio()` (`Console/console.cs`), `.InitConsole()` (`Console/console.cs`), `.AfterConstructor()` (`Console/setup.cs`)
- **`.Finish()`** — L220 — `public static void Finish()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HideAndSave()`** — L231 — `public static void HideAndSave()`
  Hides and save.
  Called by: `.OnFormClosing()` (same file), `.chkShowStartupLog_CheckedChanged()` (`Console/setup.cs`)
- **`.SetRegistryToShow()`** — L241 — `public static void SetRegistryToShow(bool show)`
  Sets registry to show.
  Called by: `.chkShowStartupLog_CheckedChanged()` (`Console/setup.cs`)
- **`.GetRegistryToShow()`** — L245 — `public static bool GetRegistryToShow(out bool show)`
  Returns registry to show.
  Called by: `.updateShowStartupLogCheckBox()` (`Console/setup.cs`)
- **`.addCore()`** — L250 — `static void addCore(string text, string id, bool colour_warn)`
  Called by: `.AddLogEntry()` (same file)
- **`.ensureForm()`** — L296 — `static void ensureForm()`
  Called by: `.ShowNewLog()` (same file), `.ShowLog()` (same file), `.AddLogEntry()` (same file), `.Completed()` (same file), `.Finish()` (same file), `.HideAndSave()` (same file)
- **`.runOnUiThreadSync()`** — L324 — `static void runOnUiThreadSync(MethodInvoker a)`
  Called by: `.ShowNewLog()` (same file), `.setOwner()` (same file)
- **`.runOnUiThread()`** — L331 — `static void runOnUiThread(MethodInvoker a)`
  Called by: `.ShowLog()` (same file), `.Shutdown()` (same file), `.Completed()` (same file), `.Finish()` (same file), `.HideAndSave()` (same file), `.addCore()` (same file)
- **`.readRegistryShow()`** — L338 — `static bool readRegistryShow(out bool show)`
  Called by: `.ShowNewLog()` (same file), `.GetRegistryToShow()` (same file)
- **`.writeRegistryShow()`** — L376 — `static void writeRegistryShow(bool show)`
  Called by: `.SetRegistryToShow()` (same file)
- **`.tryReadLocation()`** — L383 — `static bool tryReadLocation(out Point p)`
  Called by: `.ShowNewLog()` (same file)
- **`.writeLocation()`** — L400 — `static void writeLocation(Point p)`
  Called by: `.Shutdown()` (same file), `.HideAndSave()` (same file)
- **`.hideHorizontalScrollBar()`** — L408 — `static void hideHorizontalScrollBar(ListView lv)`
  Called by: `.addCore()` (same file)

#### `Entry` (type, L67)

_No extracted members._

#### `LogForm` (type, L413)

- **`.OnFormClosing()`** — L524 — `protected override void OnFormClosing(FormClosingEventArgs e)`
  Handles/raises the form closing event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.layoutColumns()`** — L530 — `void layoutColumns()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.centerClose()`** — L539 — `void centerClose()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `NoSelectListView` (type, L415)

- **`.SendMessage()`** — L417 — `[DllImport("user32.dll")] static extern IntPtr SendMessage(IntPtr hWnd, int msg, IntPtr wParam, IntPtr lParam)`
  Sends message.
  Called by: `.OnCreateControl()` (same file), `.OnGotFocus()` (same file)
- **`.OnCreateControl()`** — L424 — `protected override void OnCreateControl()`
  Handles/raises the create control event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnItemSelectionChanged()`** — L430 — `protected override void OnItemSelectionChanged(ListViewItemSelectionChangedEventArgs e)`
  Handles/raises the item selection changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnGotFocus()`** — L436 — `protected override void OnGotFocus(EventArgs e)`
  Handles/raises the got focus event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseDown()`** — L444 — `protected override void OnMouseDown(MouseEventArgs e)`
  Handles/raises the mouse down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnKeyDown()`** — L451 — `protected override void OnKeyDown(KeyEventArgs e)`
  Handles/raises the key down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsProgressLog.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
