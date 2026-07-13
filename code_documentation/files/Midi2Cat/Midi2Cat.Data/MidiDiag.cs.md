# `Midi2Cat/Midi2Cat.Data/MidiDiag.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Persistence and object model for controller-to-command mappings.

## How this file is used

- Used by (incoming references from other files):
  - `Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs` (references ×1, calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Add()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `MidiDiagItem` (type, L33)

_No extracted members._

#### `MidiDiagList` (type, L44)

- **`.Add()`** — L51 — `public new void Add(MidiDiagItem item)`
  Called by: `.OnMidiInput()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/MidiDiag.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
