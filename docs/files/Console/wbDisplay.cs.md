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

### Types

#### `wbDisplay` (type, L61)

- `.Init()` — L1037
- `.DrawBackground()` — L1056
- `.drawChannelBar()` — L1077
- `.RenderGDIPlus()` — L1101
- `.UpdateDisplayPeak()` — L1106
- `.DrawWideBandGrid()` — L1131
- `.dBToPixel()` — L2026
- `.DrawOffBackground()` — L2031
- `.DrawWideBand()` — L2040
- `.CreateDisplayRegions()` — L3333
- `.getRegion()` — L3411
- `.UpdateGraphicsBuffer()` — L3455
- `.Cancel_Display()` — L3465
- `.StartDisplay()` — L3481
- `.RunDisplay()` — L3528
- `.PanDisplay_MouseMove()` — L3586
- `.PanDisplay_MouseUp()` — L3759
- `.PanDisplay_MouseEnter()` — L3890
- `.PanDisplay_MouseWheel()` — L3896
- `.PanDisplay_MouseDown()` — L4057
- `.create_wideband()` — L4651
- `.initWideband()` — L4659

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/wbDisplay.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
