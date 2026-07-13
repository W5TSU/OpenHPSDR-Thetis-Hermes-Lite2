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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ParsedMidiMessage` (type, L44)

_No extracted members._

#### `MidiDevice` (type, L54)

- **`.GetDeviceName()`** — L98 — `public string GetDeviceName()`
  Returns device name.
  Called by: `.IsBehringerCMD()` (`Console/Midi2CatCommands.cs`), `.RIT_inc()` (`Console/Midi2CatCommands.cs`), `.XIT_inc()` (`Console/Midi2CatCommands.cs`), `.ChangeFreqVfoA()` (`Console/Midi2CatCommands.cs`), `.ChangeFreqVfoB()` (`Console/Midi2CatCommands.cs`), `.AGCLevel()` (`Console/Midi2CatCommands.cs`) — and 5 more
- **`.getUniqueDevice()`** — L102 — `private MidiUniqueDevices getUniqueDevice(string sDeviceName)`
  Returns unique device.
  Called by: `.OpenMidiIn()` (same file)
- **`.OpenMidiIn()`** — L120 — `public bool OpenMidiIn(int deviceIndex,string deviceName)`
  Opens midi in.
  Called by: `.OpenMidiDevice()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`), `.InitDevice()` (`Midi2Cat/MidiMessageManager.cs`)
- **`.OpenMidiOut()`** — L169 — `public bool OpenMidiOut()`
  Opens midi out.
  Called by: `.OpenMidiIn()` (same file), `.SendMsg()` (same file)
- **`.CloseMidiIn()`** — L203 — `public void CloseMidiIn()`
  Closes midi in.
  Called by: `.Reset()` (same file), `.CloseMidiDevice()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)
- **`.CloseMidiOut()`** — L217 — `public void CloseMidiOut()`
  Closes midi out.
  Called by: `.Reset()` (same file), `.CloseMidiDevice()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)
- **`.Reset()`** — L226 — `private void Reset()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DebugByte()`** — L343 — `public static void DebugByte(byte[] b)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InCallback()`** — L354 — `private int InCallback(int hMidiIn, int wMsg, int dwInstance, int dwParam1, int dwParam2)`
  private Hashtable midi_in_table = new Hashtable(10);
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SwapBytes()`** — L565 — `private static byte[] SwapBytes(byte[] b)`
  Called by: `.SendMsg()` (same file)
- **`.SendMsg()`** — L577 — `public static int SendMsg(int handle, ushort msg_id, byte protocol_id, ushort opcode, uint data1, uint data2)`
  Sends msg.
  Called by: `.OpenMidiIn()` (same file), `.SetPL1ButtonLight()` (same file), `.SetPL1KnobLight()` (same file), `.SendMidiCommand()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`), `.PL1InitialButtonLights()` (`Midi2Cat/MidiMessageManager.cs`), `.MicroInitialButtonLights()` (`Midi2Cat/MidiMessageManager.cs`) — and 1 more
- **`.SendLongMessage()`** — L592 — `public static int SendLongMessage(int handle, byte[] data)`
  Sends long message.
  Called by: `.SendMsg()` (same file)
- **`.inDevice_ChannelMessageReceived()`** — L642 — `public void inDevice_ChannelMessageReceived(int ControlId, int Data, int Status, int Event, int Channel)`
  Called by: `.InCallback()` (same file)
- **`.filterAndMap()`** — L657 — `private bool filterAndMap(int controlId, int status, int channel, int data, out int controlIDmapped, out int dataMapped)`
  Called by: `.inDevice_ChannelMessageReceived()` (same file)
- **`.FixBehringerCtlID()`** — L709 — `public int FixBehringerCtlID(int ControlId, int Status)`
  Called by: `.inDevice_ChannelMessageReceived()` (same file), `.OnMidiInput()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)
- **`.DebugMsg()`** — L742 — `private void DebugMsg(Direction direction, Status status, string msg1 = "", string msg2 = "")`
  Called by: `.OpenMidiIn()` (same file), `.OpenMidiOut()` (same file), `.CloseMidiIn()` (same file), `.CloseMidiOut()` (same file)
- **`.UnmapControlID()`** — L774 — `public int UnmapControlID(int inControl, out int byteCount)`
  Called by: `.ParseMsg()` (same file), `.OnMidiInput()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)
- **`.ParseMsg()`** — L803 — `public ParsedMidiMessage ParseMsg(int inChannel, int inValue, int inStatus, int inControl, string inMsg)`
  Parses msg.
  Called by: `.SendMsg()` (same file), `.ValidateMidiMessages()` (same file)
- **`.ValidateMidiMessages()`** — L877 — `public string[] ValidateMidiMessages(string inMessages)`
  Called by: `.validateDialogInput()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)
- **`.SetPL1ButtonLight()`** — L893 — `public void SetPL1ButtonLight(int n)`
  Sets pl1 button light.
  Called by: `.ToggleVFOWheel()` (`Console/Midi2CatCommands.cs`)
- **`.SetPL1KnobLight()`** — L914 — `public void SetPL1KnobLight(int n, int inCtlID)`
  Sets pl1 knob light.
  Called by: `.SendUpdateToMidi()` (`Midi2Cat/MidiMessageManager.cs`)

#### `MidiUniqueDevices` (type, L83)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.IO/MidiDevice.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
