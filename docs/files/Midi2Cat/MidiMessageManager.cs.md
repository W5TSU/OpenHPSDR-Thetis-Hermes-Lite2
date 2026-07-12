# `Midi2Cat/MidiMessageManager.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Routes incoming MIDI messages to their mapped commands.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Midi2CatCommands.cs` (calls ×3, references ×1)
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.IO/MidiDevice.cs` (calls ×5, references ×4)
  - `Midi2Cat/Midi2Cat.Data/ControllerMapping.cs` (references ×2)
  - `Midi2Cat/Midi2Cat.Data/Enums.cs` (references ×2)
  - `Midi2Cat/Midi2Cat.Data/CatCmdDb.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/ControllerBinding.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/Database.cs` (references ×1)
- Most-referenced symbols from other files: `.Open()` (×1), `.Close()` (×1), `.SendUpdateToMidi()` (×1)

## Outline

### Types

#### `Midi2Cat` (namespace, L37)

_No extracted members._

#### `MidiMessageManager` (type, L39)

- `.Open()` — L55
- `.Close()` — L77
- `.PL1Device()` — L87
- `.PL1Index()` — L102
- `.SendUpdateToMidi()` — L118
- `.PL1InitialButtonLights()` — L139
- `.MicroInitialButtonLights()` — L197
- `.InitDevice()` — L206
- `.BindMappingHandlers()` — L224
- `.onMidiDebugMsg()` — L256
- `.OnMidiInput()` — L261

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/MidiMessageManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
