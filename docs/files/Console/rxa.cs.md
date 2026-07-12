# `Console/rxa.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Typed wrappers for wdsp RXA (receiver chain) settings and the UI controls bound to them.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/PanDisplay.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

### Types

#### `rxa` (type, L15)

- `.create_rxa()` — L50
- `.ForceRxa()` — L67
- `.udRXAFreq_ValueChanged()` — L86
- `.udRXAAGCGain_ValueChanged()` — L92
- `.udRXAVolume_ValueChanged()` — L97
- `.udRXAMode_ValueChanged()` — L102
- `.rxa_FormClosing()` — L130
- `.panDisplay_Resize()` — L143
- `.rxa_Resize()` — L167

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/rxa.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
