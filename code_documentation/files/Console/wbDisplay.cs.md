# `Console/wbDisplay.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Wideband (full 0–61 MHz) spectrum display and its data acquisition from the radio's wideband sample stream.

## How this file is used

- Used by (incoming references from other files):
  - `Console/wideband.Designer.cs` (references ×1)
  - `Console/wideband.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×8)
  - `Console/HPSDR/specHPSDR.cs` (calls ×5)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Channel.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `wbDisplay` (type, L61)

- **`.Init()`** — L1037 — `public void Init()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawBackground()`** — L1056 — `public void DrawBackground()`
  Draws background.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.drawChannelBar()`** — L1077 — `void drawChannelBar(Graphics g, Channel chan, int left, int right, int top, int height, Color c, Color h)`
  draws the vertical bar to highlight where a channel is on the panadapter
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RenderGDIPlus()`** — L1101 — `unsafe public void RenderGDIPlus(int rx, Graphics e)`
  Renders gdiplus.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateDisplayPeak()`** — L1106 — `private void UpdateDisplayPeak(float[] buffer, float[] new_data)`
  Updates display peak.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawWideBandGrid()`** — L1131 — `unsafe private void DrawWideBandGrid(Graphics g, int rx)`
  Draws wide band grid.
  Called by: `.DrawWideBand()` (same file)
- **`.dBToPixel()`** — L2026 — `private float dBToPixel(float dB)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawOffBackground()`** — L2031 — `private void DrawOffBackground(Graphics g, int W, int H, bool bottom)`
  Draws off background.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawWideBand()`** — L2040 — `unsafe private bool DrawWideBand(Graphics g, int rx)`
  Draws wide band.
  Called by: `.RenderGDIPlus()` (same file), `.StartDisplay()` (same file)
- **`.CreateDisplayRegions()`** — L3333 — `public void CreateDisplayRegions()`
  Creates display regions.
  Called by: `.Init()` (same file)
- **`.getRegion()`** — L3411 — `private void getRegion(Point p)`
  Returns region.
  Called by: `.PanDisplay_MouseMove()` (same file), `.PanDisplay_MouseUp()` (same file), `.PanDisplay_MouseDown()` (same file)
- **`.UpdateGraphicsBuffer()`** — L3455 — `public void UpdateGraphicsBuffer()`
  Updates graphics buffer.
  Called by: `.StartDisplay()` (same file)
- **`.Cancel_Display()`** — L3465 — `public void Cancel_Display()`
  Called by: `.StartDisplay()` (same file)
- **`.StartDisplay()`** — L3481 — `public void StartDisplay(int rx)`
  Starts display.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RunDisplay()`** — L3528 — `unsafe private void RunDisplay(int rx)`
  Called by: `.StartDisplay()` (same file)
- **`.PanDisplay_MouseMove()`** — L3586 — `private void PanDisplay_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseUp()`** — L3759 — `private void PanDisplay_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseEnter()`** — L3890 — `private void PanDisplay_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `PanDisplay` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseWheel()`** — L3896 — `private void PanDisplay_MouseWheel(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives a mouse wheel event.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseDown()`** — L4057 — `private void PanDisplay_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.create_wideband()`** — L4651 — `public void create_wideband(int adc)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initWideband()`** — L4659 — `public void initWideband()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/wbDisplay.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
