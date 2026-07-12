# `Midi2Cat/Midi2Cat.Data/Enums.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Persistence and object model for controller-to-command mappings.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Midi2CatCommands.cs` (references ×69)
  - `Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs` (references ×3)
  - `Midi2Cat/Midi2Cat.Data/Database.cs` (references ×1, calls ×1)
  - `Midi2Cat/Midi2Cat.Data/MappedCommands.cs` (calls ×1, references ×1)
  - `Midi2Cat/Midi2Cat.IO/MidiDevice.cs` (references ×2)
  - `Midi2Cat/MidiMessageManager.cs` (references ×2)
  - `Midi2Cat/Midi2Cat.Data/CatCommandAttribute.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/ControllerMapping.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.FixControlType()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `FixUp` (type, L33)

- **`.FixControlType()`** — L35 — `public static ControlType FixControlType(int controlType)`
  Called by: `.PopulateMapping()` (`Midi2Cat/Midi2Cat.Data/Database.cs`), `.GetDeviceMappings()` (`Midi2Cat/Midi2Cat.Data/MappedCommands.cs`)

#### `CmdState` (type, L47)

_No extracted members._

#### `MidiEvent` (type, L54)

_No extracted members._

#### `ControlType` (type, L65)

_No extracted members._

#### `Direction` (type, L75)

_No extracted members._

#### `Status` (type, L82)

_No extracted members._

#### `MappingFilter` (type, L90)

_No extracted members._

#### `Command` (type, L97)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/Enums.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
