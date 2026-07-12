# `Midi2Cat/Midi2Cat.IO/MidiDevice.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** MIDI device open/close and message receive/send over the Windows Multimedia (winmm) API.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Midi2CatCommands.cs` (references ×241, calls ×12)
  - `Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs` (calls ×7, references ×2)
  - `Midi2Cat/MidiMessageManager.cs` (calls ×5, references ×4)
  - `Midi2Cat/Midi2Cat.Data/ControllerBinding.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.IO/WinMM.cs` (calls ×13)
  - `Midi2Cat/Midi2Cat.Data/Enums.cs` (references ×2)
- Most-referenced symbols from other files: `.GetDeviceName()` (×11), `.SendMsg()` (×4), `.OpenMidiIn()` (×2), `.SetPL1ButtonLight()` (×1), `.CloseMidiIn()` (×1), `.CloseMidiOut()` (×1), `.FixBehringerCtlID()` (×1), `.UnmapControlID()` (×1)

## Outline

### Types

#### `ParsedMidiMessage` (type, L44)

_No extracted members._

#### `MidiDevice` (type, L54)

- `.GetDeviceName()` — L98
- `.getUniqueDevice()` — L102
- `.OpenMidiIn()` — L120
- `.OpenMidiOut()` — L169
- `.CloseMidiIn()` — L203
- `.CloseMidiOut()` — L217
- `.Reset()` — L226
- `.DebugByte()` — L343
- `.InCallback()` — L354
- `.SwapBytes()` — L565
- `.SendMsg()` — L577
- `.SendLongMessage()` — L592
- `.inDevice_ChannelMessageReceived()` — L642
- `.filterAndMap()` — L657
- `.FixBehringerCtlID()` — L709
- `.DebugMsg()` — L742
- `.UnmapControlID()` — L774
- `.ParseMsg()` — L803
- `.ValidateMidiMessages()` — L877
- `.SetPL1ButtonLight()` — L893
- `.SetPL1KnobLight()` — L914

#### `MidiUniqueDevices` (type, L83)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.IO/MidiDevice.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
