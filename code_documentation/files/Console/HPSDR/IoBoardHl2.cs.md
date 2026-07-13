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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.getIOBoard()`** — L96 — `public static IOBoard getIOBoard(Console c)`
  Returns ioboard.
  Called by: `.UpdateIOBoard()` (`Console/console.cs`)

### Types

#### `IOBoard` (type, L37)

- **`.readRequest()`** — L129 — `public byte readRequest(Registers readRequest)`
  Called by: `.UpdateIOBoard()` (`Console/console.cs`)
- **`.readResponse()`** — L148 — `public byte readResponse()`
  Called by: `.UpdateIOBoard()` (`Console/console.cs`)
- **`.readRegister()`** — L172 — `public byte readRegister(Registers readRequest)`
  Called by: `.UpdateIOBoard()` (`Console/console.cs`)
- **`.writeRequest()`** — L177 — `public void writeRequest(Registers writeRequest, byte writeData)`
  Called by: `.AutoTuningHL2()` (`Console/console.cs`), `.UpdateIOBoard()` (`Console/console.cs`)
- **`.setFrequency()`** — L184 — `public void setFrequency(long frequency)`
  Sets frequency.
  Called by: `.UpdateIOBoard()` (`Console/console.cs`)

#### `Registers` (type, L39)

_No extracted members._

#### `HardwareVersion` (type, L82)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/IoBoardHl2.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
