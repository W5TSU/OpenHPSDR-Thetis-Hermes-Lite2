# `Console/HPSDR/Alex.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** Alex RF filter board control (antenna and band-filter relay selection). Retained from upstream; antenna switching from the console is disabled in this HL2 fork.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×4)
  - `Console/setup.cs` (calls ×2)
  - `Console/Andromeda/Andromeda.cs` (calls ×1)
  - `Console/MeterManager.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×9)
- Most-referenced symbols from other files: `.getAlex()` (×7), `.AntBandFromFreq()` (×1)

## Outline

### Functions

- `setBandBitMask()` — L480
- `ExtCtrlEnable()` — L495
- `UpdateExtCtrl()` — L507
- `.getAlex()` — L38

### Types

#### `Alex` (type, L34)

- `.SetAntennasTo1()` — L77
- `.setRxAnt()` — L82
- `.setRxOnlyAnt()` — L92
- `.setTxAnt()` — L102
- `.AntBandFromFreq()` — L112
- `.AntBandFromFreqB()` — L238
- `.UpdateAlexAntSelection()` — L303

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/Alex.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
