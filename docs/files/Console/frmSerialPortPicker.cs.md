# `Console/frmSerialPortPicker.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Small shared picker dialogs.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmSerialPortPicker` (type, L48)

- **`.Init()`** — L89 — `public void Init()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateCombos()`** — L110 — `private void updateCombos(bool enabled)`
  Called by: `.Init()` (same file), `.setupSubCombos()` (same file)
- **`.btnSelect_Click()`** — L117 — `private void btnSelect_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSelect` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCancel_Click()`** — L154 — `private void btnCancel_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCancel` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.IsBaudRateSupported()`** — L159 — `private bool IsBaudRateSupported(string portName, int baudRate)`
  Called by: `.btnSelect_Click()` (same file)
- **`.IsDataBitsSupported()`** — L179 — `private bool IsDataBitsSupported(string portName, int baudRate, int dataBits)`
  Called by: `.btnSelect_Click()` (same file)
- **`.IsStopBitsSupported()`** — L200 — `private bool IsStopBitsSupported(string portName, int baudRate, int dataBits, StopBits stopBits)`
  Called by: `.btnSelect_Click()` (same file)
- **`.IsParitySupported()`** — L222 — `private bool IsParitySupported(string portName, int baudRate, int dataBits, StopBits stopBits, Parity parity)`
  Called by: `.btnSelect_Click()` (same file)
- **`.IsComPortAvailable()`** — L244 — `public static bool IsComPortAvailable(string portName)`
  Called by: `.btnSelect_Click()` (same file)
- **`.comboComPort_SelectedIndexChanged()`** — L263 — `private void comboComPort_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboComPort` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupSubCombos()`** — L270 — `private void setupSubCombos(string com_port)`
  Called by: `.comboComPort_SelectedIndexChanged()` (same file)
- **`.comboBaudRate_SelectedIndexChanged()`** — L341 — `private void comboBaudRate_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboBaudRate` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboDataBits_SelectedIndexChanged()`** — L349 — `private void comboDataBits_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboDataBits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboStopBits_SelectedIndexChanged()`** — L357 — `private void comboStopBits_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboStopBits` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboParity_SelectedIndexChanged()`** — L378 — `private void comboParity_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboParity` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmSerialPortPicker.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
