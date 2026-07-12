# `Console/frmNotchPopup.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Manual notch filter add/edit popup (backed by wdsp `nbp`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/radio.cs` (references ×1)

## Outline

### Types

#### `frmNotchPopup` (type, L48)

- `.Show()` — L74
- `.FrmNotchPopup_Deactivate()` — L117
- `.BtnDelete_Click()` — L125
- `.setBW()` — L160
- `.Btn25_Click()` — L167
- `.Btn50_Click()` — L172
- `.Btn100_Click()` — L177
- `.Btn200_Click()` — L182
- `.TrkWidth_Scroll()` — L187
- `.setText()` — L194
- `.ChkActive_CheckedChanged()` — L199

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmNotchPopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
