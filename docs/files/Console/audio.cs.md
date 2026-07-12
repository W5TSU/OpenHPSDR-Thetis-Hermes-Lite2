# `Console/audio.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** Audio device enumeration and configuration: sample rates, buffer sizes, VAC on/off, device selection; drives ChannelMaster accordingly.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×34)
  - `Console/console.cs` (calls ×2)
  - `Console/cmaster.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/portaudio.cs` (calls ×18)
  - `Console/enums.cs` (references ×3)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
- Most-referenced symbols from other files: `.EnableVAC2()` (×14), `.EnableVAC1()` (×13), `.GetPAInputDevices()` (×3), `.GetPAOutputDevices()` (×3), `.GetPAHosts()` (×1), `.Start()` (×1), `.Stop()` (×1), `.DoScope()` (×1)

## Outline

### Types

#### `Audio` (type, L55)

- `.setupIVACforMon()` — L386
- `.isPowerOfTwo()` — L953
- `.SetOutCount()` — L1346
- `.SetOutCountRX2()` — L1354
- `.SetOutCountTX()` — L1362
- `.GetPAHosts()` — L1475
- `.GetPAInputDevices()` — L1487
- `.CheckPAInputDevices()` — L1523
- `.GetPAOutputDevices()` — L1536
- `.CheckPAOutputDevices()` — L1572
- `.EnableVAC1()` — L1585
- `.EnableVAC2()` — L1665
- `.Start()` — L1764
- `.Stop()` — L1849
- `.DoScope()` — L1978
- `.DoScope2()` — L2076

#### `AudioState` (type, L63)

_No extracted members._

#### `SignalSource` (type, L69)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/audio.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
