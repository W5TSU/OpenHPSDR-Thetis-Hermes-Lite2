# `Console/CAT/SerialPortPTT.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** PTT via serial control lines; band-decoder output over a USB BCD cable.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×3, references ×1)
- Uses (outgoing references to other files):
  - `Console/CAT/SDRSerialPortII.cs` (calls ×7, references ×1)
- Most-referenced symbols from other files: `.isPTT()` (×1), `.setDTR()` (×1), `.Destroy()` (×1)

## Outline

### Types

#### `SerialPortPTT` (type, L33)

- `.Init()` — L67
- `.isPTT()` — L84
- `.isCTS()` — L96
- `.isDSR()` — L105
- `.setDTR()` — L114
- `.Destroy()` — L123
- `.SerialReceivedData()` — L140

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/SerialPortPTT.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
