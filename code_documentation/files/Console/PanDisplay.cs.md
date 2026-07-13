# `Console/PanDisplay.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Panadapter display user control hosting the render surface, mouse tuning, and zoom/pan interaction.

## How this file is used

- Used by (incoming references from other files):
  - `Console/rxa.Designer.cs` (references ×1)
  - `Console/rxa.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×8)
  - `Console/hiperftimer.cs` (calls ×2, references ×1)
  - `Console/HPSDR/specHPSDR.cs` (calls ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Channel.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `PanDisplay` (type, L47)

- **`.Init()`** — L999 — `public void Init()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawBackground()`** — L1020 — `public void DrawBackground()`
  Draws background.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.drawChannelBar()`** — L1041 — `void drawChannelBar(Graphics g, Channel chan, int left, int right, int top, int height, Color c, Color h)`
  draws the vertical bar to highlight where a channel is on the panadapter
  Called by: `.DrawPanadapterGrid()` (same file)
- **`.RenderGDIPlus()`** — L1065 — `unsafe public void RenderGDIPlus(int rx, Graphics e)`
  Renders gdiplus.
  Called by: `.StartDisplay()` (same file)
- **`.UpdateDisplayPeak()`** — L1080 — `private void UpdateDisplayPeak(float[] buffer, float[] new_data)`
  Updates display peak.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawPanadapterGrid()`** — L1105 — `unsafe private void DrawPanadapterGrid(Graphics g, int rx)`
  Draws panadapter grid.
  Called by: `.DrawPanadapter()` (same file)
- **`.dBToPixel()`** — L2000 — `private float dBToPixel(float dB)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawOffBackground()`** — L2005 — `private void DrawOffBackground(Graphics g, int W, int H, bool bottom)`
  Draws off background.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawPanadapter()`** — L2014 — `unsafe private bool DrawPanadapter(Graphics g, int rx)`
  Draws panadapter.
  Called by: `.RenderGDIPlus()` (same file)
- **`.DrawWaterfall()`** — L2247 — `unsafe private bool DrawWaterfall(Graphics g, int rx)`
  Draws waterfall.
  Called by: `.RenderGDIPlus()` (same file)
- **`.CreateDisplayRegions()`** — L3387 — `public void CreateDisplayRegions()`
  Creates display regions.
  Called by: `.Init()` (same file)
- **`.getRegion()`** — L3465 — `private void getRegion(Point p)`
  Returns region.
  Called by: `.PanDisplay_MouseMove()` (same file), `.PanDisplay_MouseUp()` (same file), `.PanDisplay_MouseDown()` (same file)
- **`.UpdateGraphicsBuffer()`** — L3510 — `public void UpdateGraphicsBuffer()`
  Updates graphics buffer.
  Called by: `.StartDisplay()` (same file)
- **`.Cancel_Display1()`** — L3522 — `public bool Cancel_Display1()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StartDisplay()`** — L3528 — `public void StartDisplay(int rx)`
  Starts display.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RunDisplay()`** — L3559 — `unsafe private void RunDisplay(int rx)`
  Called by: `.StartDisplay()` (same file)
- **`.PanDisplay_MouseMove()`** — L3584 — `private void PanDisplay_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseUp()`** — L3753 — `private void PanDisplay_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseEnter()`** — L3884 — `private void PanDisplay_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `PanDisplay` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseWheel()`** — L3890 — `private void PanDisplay_MouseWheel(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives a mouse wheel event.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PanDisplay_MouseDown()`** — L4051 — `private void PanDisplay_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `PanDisplay` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.initAnalyzer()`** — L4674 — `public void initAnalyzer()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/PanDisplay.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
