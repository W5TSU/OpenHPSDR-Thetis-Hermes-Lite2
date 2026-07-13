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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`.getPenny()`** — L41 — `public static Penny getPenny()`
  Returns penny.
  Called by: `.HdwMOXChanged()` (`Console/console.cs`), `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`), `.chkVFOBTX_CheckedChanged()` (`Console/console.cs`), `.chkExternalPA_CheckedChanged()` (`Console/console.cs`), `.chkPenOCrcv160_CheckedChanged()` (`Console/setup.cs`) — and 87 more

### Types

#### `Penny` (type, L35)

- **`.setRXPinPA()`** — L77 — `public void setRXPinPA(int group, int pin, bool pa)`
  Sets rxpin pa.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTXPinPA()`** — L86 — `public void setTXPinPA(int group, int pin, bool pa)`
  Sets txpin pa.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTXPinAction()`** — L96 — `public void setTXPinAction(int group, int pin, TXPinActions action)`
  Sets txpin action.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setBandABitMask()`** — L106 — `public void setBandABitMask(Band band, byte mask, bool tx)`
  Sets band abit mask.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setBandBBitMask()`** — L121 — `public void setBandBBitMask(Band band, byte mask, bool tx)`
  Sets band bbit mask.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExtCtrlEnable()`** — L136 — `public int ExtCtrlEnable(Band band, Band bandb, bool tx, bool enable, bool tune, bool twoTone, bool pa)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateExtCtrl()`** — L158 — `public int UpdateExtCtrl(Band band, Band bandb, bool tx, bool tune, bool twoTone, bool pa)`
  Updates ext ctrl.
  Called by: `.ExtCtrlEnable()` (same file)
- **`.getGroup()`** — L224 — `private int getGroup(Band b)`
  Returns group.
  Called by: `.adjustForRX()` (same file), `.adjustForTXAction()` (same file)
- **`.adjustForRX()`** — L284 — `private int adjustForRX(Band band, Band bandb, int bits, bool tx, bool tune, bool twoTone, bool pa)`
  Called by: `.UpdateExtCtrl()` (same file)
- **`.adjustForTXAction()`** — L319 — `private int adjustForTXAction(Band band, Band bandb, int bits, bool tx, bool tune, bool twoTone, bool pa)`
  Called by: `.UpdateExtCtrl()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/Penny.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
