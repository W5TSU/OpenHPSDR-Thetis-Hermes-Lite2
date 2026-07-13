# `Midi2Cat/Midi2Cat.Data/MappedCommands.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Persistence and object model for controller-to-command mappings.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.Data/Enums.cs` (calls ×1, references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `MappedCommand` (type, L29)

_No extracted members._

#### `MappedCommands` (type, L59)

- **`.GetDeviceMappings()`** — L76 — `private int GetDeviceMappings(MappedCommand mappedCmd, DataTable controllerDT)`
  Returns device mappings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/MappedCommands.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
