# `Console/radio.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Radio/receiver object model — bands, modes, filter presets per mode, and per-RX DSP state that the console manipulates.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×91, references ×8)
  - `Console/setup.cs` (calls ×3)
  - `Console/display.cs` (references ×1, calls ×1)
  - `Console/MeterManager.cs` (calls ×1)
  - `Console/database.cs` (calls ×1)
  - `Console/frmNotchPopup.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×6)
- Most-referenced symbols from other files: `.GetDSPRX()` (×45), `.GetDSPTX()` (×27), `.NotchFromIndex()` (×6), `.IndexOf()` (×5), `.GetFirstNotchThatMatches()` (×4), `.NotchesInBW()` (×2), `.Parse()` (×2), `.Add()` (×2)

## Outline

### Functions

- `.Copy()` — L282
- `.Parse()` — L4358

### Types

#### `Radio` (type, L54)

- `.Shutdown()` — L80
- `.GetDSPRX()` — L84
- `.GetDSPTX()` — L89

#### `RadioDSP` (type, L99)

- `.CreateDSP()` — L103
- `.DestroyDSP()` — L163

#### `RadioDSPRX` (type, L270)

- `.SyncAll()` — L388
- `.SetRXFilter()` — L597
- `.SetNRVals()` — L683
- `.SetANFVals()` — L730
- `.GetNotchOn()` — L1449
- `.SetNotchOn()` — L1454
- `.GetNotchFreq()` — L1468
- `.SetNotchFreq()` — L1473
- `.GetNotchBW()` — L1487
- `.SetNotchBW()` — L1497

#### `RadioDSPTX` (type, L2471)

- `.SyncAll()` — L2480
- `.SetTXFilter()` — L2732

#### `MNotchDB` (type, L4192)

- `.Clear()` — L4205
- `.IndexOf()` — L4212
- `.Add()` — L4219
- `.NotchFromIndex()` — L4236
- `.GetFirstNotchThatMatches()` — L4246
- `.NotchNearFreq()` — L4261
- `.NotchesInBW()` — L4276
- `.NotchThatSurroundsFrequencyInBW()` — L4297

#### `MNotch` (type, L4328)

- `.ToString()` — L4376
- `.CompareTo()` — L4381

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/radio.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
