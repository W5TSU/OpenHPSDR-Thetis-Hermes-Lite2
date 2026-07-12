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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `SerialPortPTT` (type, L33)

- **`.Init()`** — L67 — `public void Init()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isPTT()`** — L84 — `public bool isPTT()`
  Called by: `.PollPTT()` (`Console/console.cs`)
- **`.isCTS()`** — L96 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L105 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L114 — `public void setDTR(bool v)`
  Sets dtr.
  Called by: `.HdwMOXChanged()` (`Console/console.cs`)
- **`.Destroy()`** — L123 — `public void Destroy()`
  Called by: `.chkPower_CheckedChanged()` (`Console/console.cs`)
- **`.SerialReceivedData()`** — L140 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/SerialPortPTT.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
