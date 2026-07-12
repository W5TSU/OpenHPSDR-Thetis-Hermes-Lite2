# `Console/CW/CWInput.cs`

**Functional area:** [11. CW keying](../../../CODE_OUTLINE.md#11-cw-keying)

**Role:** Abstraction over CW key input sources (which line/device is dot, dash, PTT).

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/CAT/SDRSerialPortII.cs` (calls ×4, references ×1)
  - `Console/enums.cs` (references ×1)
- Most-referenced symbols from other files: `.SetPrimaryInput()` (×1), `.SetSecondaryInput()` (×1)

## Outline

### Types

#### `CWInput` (type, L39)

- `.SetPrimaryInput()` — L127
- `.SetSecondaryInput()` — L182
- `.SerialReceivedData()` — L239

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CW/CWInput.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
