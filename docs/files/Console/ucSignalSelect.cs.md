# `Console/ucSignalSelect.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** Meter-related picker controls (open-collector LED strip, signal source, linear-gradient color pickers).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/MeterManager.cs` (references ×4)

## Outline

### Types

#### `ucSignalSelect` (type, L53)

- `.radSig_CheckedChanged()` — L72
- `.radSigAvg_CheckedChanged()` — L81
- `.getSignalTypeFromSelection()` — L90
- `.onSignalTypeChanged()` — L98
- `.radSigMaxBin_CheckedChanged()` — L130

#### `SignalTypeChangedEventArgs` (type, L55)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucSignalSelect.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
