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

### Types

#### `Splash` (type, L62)

- `.InvalidateRect()` — L65
- `.Dispose()` — L128
- `.InitializeComponent()` — L144
- `.ShowSplashScreen()` — L241
- `.ShowForm()` — L268
- `.CloseForm()` — L279
- `.HideForm()` — L292
- `.UnHideForm()` — L298
- `.SetStatus()` — L305
- `.SetReferencePoint()` — L325
- `.setVersion()` — L334
- `.setBackground()` — L338
- `.SetReferenceInternal()` — L375
- `.ElapsedMilliSeconds()` — L396
- `.ReadIncrements()` — L405
- `.StoreIncrements()` — L450
- `.timer1_Tick()` — L478
- `.pnlStatus_Paint()` — L546
- `.SplashScreen_DoubleClick()` — L562
- `.Splash_Load()` — L569

#### `StartParams` (type, L108)

_No extracted members._

#### `RegistryAccess` (type, L580)

- `.GetStringRegistryValue()` — L587
- `.SetStringRegistryValue()` — L613

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/splash.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
