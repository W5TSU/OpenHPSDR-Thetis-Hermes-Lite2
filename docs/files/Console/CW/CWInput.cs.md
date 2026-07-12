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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CWInput` (type, L39)

- **`.SetPrimaryInput()`** — L127 — `public static bool SetPrimaryInput(string s)`
  Sets primary input.
  Called by: `.comboKeyerConnPrimary_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.SetSecondaryInput()`** — L182 — `public static bool SetSecondaryInput(string s)`
  Sets secondary input.
  Called by: `.comboKeyerConnSecondary_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.SerialReceivedData()`** — L239 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CW/CWInput.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
