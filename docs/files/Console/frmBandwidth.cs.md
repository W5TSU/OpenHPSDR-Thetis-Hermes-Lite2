# `Console/frmBandwidth.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Variable-bandwidth adjustment popup and its graphical bandwidth view.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)

## Outline

### Types

#### `frmBandwidth` (type, L46)

- `.timerReadBandwidth_Tick()` — L57
- `.RecoverShow()` — L64
- `.frmBandwidth_FormClosing()` — L71
- `.radUnits_CheckedChanged()` — L84
- `.chkOnTop_CheckedChanged()` — L89

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmBandwidth.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
