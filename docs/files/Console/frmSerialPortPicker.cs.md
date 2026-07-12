# `Console/frmSerialPortPicker.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Small shared picker dialogs.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `frmSerialPortPicker` (type, L48)

- `.Init()` — L89
- `.updateCombos()` — L110
- `.btnSelect_Click()` — L117
- `.btnCancel_Click()` — L154
- `.IsBaudRateSupported()` — L159
- `.IsDataBitsSupported()` — L179
- `.IsStopBitsSupported()` — L200
- `.IsParitySupported()` — L222
- `.IsComPortAvailable()` — L244
- `.comboComPort_SelectedIndexChanged()` — L263
- `.setupSubCombos()` — L270
- `.comboBaudRate_SelectedIndexChanged()` — L341
- `.comboDataBits_SelectedIndexChanged()` — L349
- `.comboStopBits_SelectedIndexChanged()` — L357
- `.comboParity_SelectedIndexChanged()` — L378

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmSerialPortPicker.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
