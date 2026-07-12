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

### Types

#### `PanDisplay` (type, L47)

- `.Init()` — L999
- `.DrawBackground()` — L1020
- `.drawChannelBar()` — L1041
- `.RenderGDIPlus()` — L1065
- `.UpdateDisplayPeak()` — L1080
- `.DrawPanadapterGrid()` — L1105
- `.dBToPixel()` — L2000
- `.DrawOffBackground()` — L2005
- `.DrawPanadapter()` — L2014
- `.DrawWaterfall()` — L2247
- `.CreateDisplayRegions()` — L3387
- `.getRegion()` — L3465
- `.UpdateGraphicsBuffer()` — L3510
- `.Cancel_Display1()` — L3522
- `.StartDisplay()` — L3528
- `.RunDisplay()` — L3559
- `.PanDisplay_MouseMove()` — L3584
- `.PanDisplay_MouseUp()` — L3753
- `.PanDisplay_MouseEnter()` — L3884
- `.PanDisplay_MouseWheel()` — L3890
- `.PanDisplay_MouseDown()` — L4051
- `.initAnalyzer()` — L4674

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/PanDisplay.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
