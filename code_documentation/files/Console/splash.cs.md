# `Console/splash.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Startup splash screen and progress reporting during initialization.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×1)
  - `Console/setup.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/Invoke/labelts.cs` (references ×1)
- Most-referenced symbols from other files: `.SetStatus()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Splash` (type, L62)

- **`.InvalidateRect()`** — L65 — `[DllImport("user32.dll")] static extern bool InvalidateRect(IntPtr hWnd, IntPtr lpRect, bool bErase)`
  MW0LGE
  Called by: `.timer1_Tick()` (same file)
- **`.Dispose()`** — L128 — `protected override void Dispose( bool disposing )`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L144 — `private void InitializeComponent()`
  Designer-generated UI construction (creates and lays out the form’s controls).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowSplashScreen()`** — L241 — `static public void ShowSplashScreen(string version, string splash_screen_folder = "")`
  A static method to create the thread and launch the SplashScreen.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowForm()`** — L268 — `static private void ShowForm()`
  A private entry point for the thread.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CloseForm()`** — L279 — `static public void CloseForm()`
  A static method to close the SplashScreen
  Called by: `.SplashScreen_DoubleClick()` (same file)
- **`.HideForm()`** — L292 — `static public void HideForm()`
  Hides form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UnHideForm()`** — L298 — `static public void UnHideForm()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetStatus()`** — L305 — `static public void SetStatus(string newStatus)`
  A static method to set the status and update the reference.
  Called by: `.InitConsole()` (`Console/console.cs`), `.AfterConstructor()` (`Console/setup.cs`)
- **`.SetReferencePoint()`** — L325 — `static public void SetReferencePoint()`
  Static method called from the initializing application to give the splash screen reference points. Not needed if you are using a lot of status strings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setVersion()`** — L334 — `private void setVersion(string version)`
  ************ Private methods ************
  Called by: `.ShowForm()` (same file)
- **`.setBackground()`** — L338 — `private void setBackground(string splash_screen_folder)`
  Sets background.
  Called by: `.ShowForm()` (same file)
- **`.SetReferenceInternal()`** — L375 — `private void SetReferenceInternal()`
  Internal method for setting reference points.
  Called by: `.SetStatus()` (same file), `.SetReferencePoint()` (same file)
- **`.ElapsedMilliSeconds()`** — L396 — `private double ElapsedMilliSeconds()`
  Utility function to return elapsed Milliseconds since the SplashScreen was launched.
  Called by: `.SetReferenceInternal()` (same file), `.StoreIncrements()` (same file)
- **`.ReadIncrements()`** — L405 — `private void ReadIncrements()`
  Function to read the checkpoint intervals from the previous invocation of the splashscreen from the registry.
  Called by: `.SetReferenceInternal()` (same file)
- **`.StoreIncrements()`** — L450 — `private void StoreIncrements()`
  Method to store the intervals (in percent complete) from the current invocation of the splash screen to the registry.
  Called by: `.timer1_Tick()` (same file)
- **`.timer1_Tick()`** — L478 — `private void timer1_Tick(object sender, System.EventArgs e)`
  Tick Event handler for the Timer control. Handle fade in and fade out. Also handle the smoothed progress bar.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlStatus_Paint()`** — L546 — `private void pnlStatus_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  Paint the portion of the panel invalidated during the tick event.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SplashScreen_DoubleClick()`** — L562 — `private void SplashScreen_DoubleClick(object sender, System.EventArgs e)`
  Close the form if they double click on it.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Splash_Load()`** — L569 — `private void Splash_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `Splash` loads.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `StartParams` (type, L108)

_No extracted members._

#### `RegistryAccess` (type, L580)

- **`.GetStringRegistryValue()`** — L587 — `static public string GetStringRegistryValue(string key, string defaultValue)`
  Method for retrieving a Registry Value.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetStringRegistryValue()`** — L613 — `static public void SetStringRegistryValue(string key, string stringValue)`
  Method for storing a Registry Value.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/splash.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
