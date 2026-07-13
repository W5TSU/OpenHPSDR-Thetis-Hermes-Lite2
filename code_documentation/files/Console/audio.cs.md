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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Audio` (type, L55)

- **`.setupIVACforMon()`** — L386 — `private static void setupIVACforMon()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isPowerOfTwo()`** — L953 — `private static bool isPowerOfTwo(int x)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOutCount()`** — L1346 — `private static void SetOutCount()`
  Sets out count.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOutCountRX2()`** — L1354 — `private static void SetOutCountRX2()`
  Sets out count rx2.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOutCountTX()`** — L1362 — `private static void SetOutCountTX()`
  Sets out count tx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPAHosts()`** — L1475 — `public static ArrayList GetPAHosts()`
  Returns pahosts.
  Called by: `.GetHosts()` (`Console/setup.cs`)
- **`.GetPAInputDevices()`** — L1487 — `public static ArrayList GetPAInputDevices(int hostIndex)`
  Returns painput devices.
  Called by: `.GetHosts()` (`Console/setup.cs`), `.GetDevices2()` (`Console/setup.cs`), `.GetDevices3()` (`Console/setup.cs`)
- **`.CheckPAInputDevices()`** — L1523 — `public static bool CheckPAInputDevices(int hostIndex, string name)`
  Checks painput devices.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPAOutputDevices()`** — L1536 — `public static ArrayList GetPAOutputDevices(int hostIndex)`
  Returns paoutput devices.
  Called by: `.GetHosts()` (`Console/setup.cs`), `.GetDevices2()` (`Console/setup.cs`), `.GetDevices3()` (`Console/setup.cs`)
- **`.CheckPAOutputDevices()`** — L1572 — `public static bool CheckPAOutputDevices(int hostIndex, string name)`
  Checks paoutput devices.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnableVAC1()`** — L1585 — `public static void EnableVAC1(bool enable)`
  Enables vac1.
  Called by: `.comboAudioDriver2_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioInput2_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioOutput2_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioSampleRate1_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioSampleRate2_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioBuffer2_SelectedIndexChanged()` (`Console/setup.cs`) — and 7 more
- **`.EnableVAC2()`** — L1665 — `public static void EnableVAC2(bool enable)`
  Enables vac2.
  Called by: `.comboAudioDriver3_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioInput3_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioOutput3_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioSampleRate1_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioSampleRateRX2_SelectedIndexChanged()` (`Console/setup.cs`), `.comboAudioSampleRate3_SelectedIndexChanged()` (`Console/setup.cs`) — and 8 more
- **`.Start()`** — L1764 — `public static bool Start()`
  Called by: `.chkPower_CheckedChanged()` (`Console/console.cs`)
- **`.Stop()`** — L1849 — `public static void Stop()`
  Called by: `.chkPower_CheckedChanged()` (`Console/console.cs`)
- **`.DoScope()`** — L1978 — `unsafe public static void DoScope(float* buf, int frameCount)`
  Called by: `.xscope()` (`Console/cmaster.cs`)
- **`.DoScope2()`** — L2076 — `unsafe public static void DoScope2(float* buf, int frameCount)`
  Called by: `.xscope()` (`Console/cmaster.cs`)

#### `AudioState` (type, L63)

_No extracted members._

#### `SignalSource` (type, L69)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/audio.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
