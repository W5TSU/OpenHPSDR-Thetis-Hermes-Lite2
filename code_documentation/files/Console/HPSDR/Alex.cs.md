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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`setBandBitMask()`** — L480 — `public void setBandBitMask(Band band, byte mask, bool tx)`
  Sets band bit mask.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`ExtCtrlEnable()`** — L495 — `public void ExtCtrlEnable(bool enable, Band band, bool tx )`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`UpdateExtCtrl()`** — L507 — `public void UpdateExtCtrl(Band band, bool tx)`
  Updates ext ctrl.
  Called by: `ExtCtrlEnable()` (same file)
- **`.getAlex()`** — L38 — `public static Alex getAlex()`
  Returns alex.
  Called by: `.SetAriesAlexMode()` (`Console/Andromeda/Andromeda.cs`), `.HdwMOXChanged()` (`Console/console.cs`), `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.chkVFOATX_CheckedChanged()` (`Console/console.cs`), `.chkVFOBTX_CheckedChanged()` (`Console/console.cs`), `.ProcessAlexAntCheckBox()` (`Console/setup.cs`) — and 1 more

### Types

#### `Alex` (type, L34)

- **`.SetAntennasTo1()`** — L77 — `public void SetAntennasTo1(bool IsSetTo1)`
  SetAntennasTo1 causes RX, TX antennas to be set to 1 the various RX "bypass" unaffected
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setRxAnt()`** — L82 — `public void setRxAnt(Band band, byte ant)`
  Sets rx ant.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setRxOnlyAnt()`** — L92 — `public void setRxOnlyAnt(Band band, byte ant)`
  Sets rx only ant.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTxAnt()`** — L102 — `public void setTxAnt(Band band, byte ant)`
  Sets tx ant.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AntBandFromFreq()`** — L112 — `public static Band AntBandFromFreq(double freq)`
  Called by: `.UpdateAlexAntSelection()` (same file), `.AntennasChanged()` (`Console/MeterManager.cs`)
- **`.AntBandFromFreqB()`** — L238 — `public static Band AntBandFromFreqB()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateAlexAntSelection()`** — L303 — `public void UpdateAlexAntSelection(Band band, bool tx, bool xvtr)`
  Updates alex ant selection.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/Alex.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
