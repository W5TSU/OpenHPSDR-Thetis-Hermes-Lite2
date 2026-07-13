# `Console/ucSignalSelect.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** Meter-related picker controls (open-collector LED strip, signal source, linear-gradient color pickers).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/MeterManager.cs` (references ×4)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucSignalSelect` (type, L53)

- **`.radSig_CheckedChanged()`** — L72 — `private void radSig_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSig` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radSigAvg_CheckedChanged()`** — L81 — `private void radSigAvg_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSigAvg` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getSignalTypeFromSelection()`** — L90 — `private Reading getSignalTypeFromSelection()`
  Returns signal type from selection.
  Called by: `.radSig_CheckedChanged()` (same file), `.radSigAvg_CheckedChanged()` (same file), `.radSigMaxBin_CheckedChanged()` (same file)
- **`.onSignalTypeChanged()`** — L98 — `private void onSignalTypeChanged(Reading signalType)`
  Called by: `.radSig_CheckedChanged()` (same file), `.radSigAvg_CheckedChanged()` (same file), `.radSigMaxBin_CheckedChanged()` (same file)
- **`.radSigMaxBin_CheckedChanged()`** — L130 — `private void radSigMaxBin_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radSigMaxBin` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `SignalTypeChangedEventArgs` (type, L55)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucSignalSelect.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
