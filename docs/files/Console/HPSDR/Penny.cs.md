# `Console/HPSDR/Penny.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** PennyLane/Penelope open-collector output and mic-gain control by band.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×88)
  - `Console/console.cs` (calls ×5)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×9)
- Most-referenced symbols from other files: `.getPenny()` (×93)

## Outline

### Functions

- `.getPenny()` — L41

### Types

#### `Penny` (type, L35)

- `.setRXPinPA()` — L77
- `.setTXPinPA()` — L86
- `.setTXPinAction()` — L96
- `.setBandABitMask()` — L106
- `.setBandBBitMask()` — L121
- `.ExtCtrlEnable()` — L136
- `.UpdateExtCtrl()` — L158
- `.getGroup()` — L224
- `.adjustForRX()` — L284
- `.adjustForTXAction()` — L319

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/Penny.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
