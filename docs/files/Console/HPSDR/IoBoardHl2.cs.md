# `Console/HPSDR/IoBoardHl2.cs`

**Functional area:** [4. Hermes-Lite 2 specific hardware control](../../../CODE_OUTLINE.md#4-hermes-lite-2-specific-hardware-control)

**Role:** Register-level control of the HL2 I/O board: antenna tuner (ATU) commands and status, TX frequency bytes sent to the board, fault detection, and control-register read/write over the HL2's i2c-style interface.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×7, references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
- Most-referenced symbols from other files: `.writeRequest()` (×2), `.getIOBoard()` (×1), `.readRequest()` (×1), `.readResponse()` (×1), `.readRegister()` (×1), `.setFrequency()` (×1)

## Outline

### Functions

- `.getIOBoard()` — L96

### Types

#### `IOBoard` (type, L37)

- `.readRequest()` — L129
- `.readResponse()` — L148
- `.readRegister()` — L172
- `.writeRequest()` — L177
- `.setFrequency()` — L184

#### `Registers` (type, L39)

_No extracted members._

#### `HardwareVersion` (type, L82)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/IoBoardHl2.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
