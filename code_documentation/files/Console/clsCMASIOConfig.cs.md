# `Console/clsCMASIOConfig.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Configuration UI/state for the cmASIO ASIO driver connection.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×15)
- Uses (outgoing references to other files):
  - `Console/portaudio.cs` (calls ×4)
- Most-referenced symbols from other files: `.SetASIOdrivername()` (×2), `.SetASIOblocknum()` (×2), `.GetASIObaseinchannel()` (×2), `.GetASIOdrivername()` (×1), `.GetASIOblocknum()` (×1), `.GetASIOlockmode()` (×1), `.SetASIObaseinchannel()` (×1), `.SetASIObaseoutchannel()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `AsioDeviceInfo` (type, L48)

- **`.ToString()`** — L63 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `CMASIOConfig` (type, L69)

- **`.openRegistryKey()`** — L76 — `private static RegistryKey openRegistryKey()`
  Called by: `.DoesRegistryValueExist()` (same file), `.GetASIOdrivername()` (same file), `.SetASIOdrivername()` (same file), `.GetASIOblocknum()` (same file), `.GetASIOlockmode()` (same file), `.SetASIOblocknum()` (same file) — and 7 more
- **`.DoesRegistryValueExist()`** — L98 — `public static bool DoesRegistryValueExist(string valueName, RegistryKey key = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetASIOdrivername()`** — L109 — `public static string GetASIOdrivername(RegistryKey key = null)`
  Returns asiodrivername.
  Called by: `.SetupCMAsio()` (`Console/setup.cs`)
- **`.SetASIOdrivername()`** — L122 — `public static void SetASIOdrivername(string driverName, RegistryKey key = null)`
  Sets asiodrivername.
  Called by: `.btnCMASIOActive_Click()` (`Console/setup.cs`), `.btnCMASIODisable_Click()` (`Console/setup.cs`)
- **`.GetASIOblocknum()`** — L130 — `public static int GetASIOblocknum(RegistryKey key = null)`
  Returns asioblocknum.
  Called by: `.SetupCMAsio()` (`Console/setup.cs`)
- **`.GetASIOlockmode()`** — L143 — `public static bool GetASIOlockmode(RegistryKey key = null)`
  Returns asiolockmode.
  Called by: `.SetupCMAsio()` (`Console/setup.cs`)
- **`.SetASIOblocknum()`** — L156 — `public static void SetASIOblocknum(int blockNum, bool lock_mode, RegistryKey key = null)`
  Sets asioblocknum.
  Called by: `.nudAsioBlockNum_ValueChanged()` (`Console/setup.cs`), `.chkAsioLockMode_CheckedChanged()` (`Console/setup.cs`)
- **`.SetASIObaseinchannel()`** — L165 — `public static void SetASIObaseinchannel(int base_input_channel, RegistryKey key = null)`
  Sets asiobaseinchannel.
  Called by: `.comboCMASIO_inpair_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.SetASIObaseoutchannel()`** — L173 — `public static void SetASIObaseoutchannel(int base_output_channel, RegistryKey key = null)`
  Sets asiobaseoutchannel.
  Called by: `.comboCMASIO_outpair_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.GetASIObaseinchannel()`** — L181 — `public static int GetASIObaseinchannel(RegistryKey key = null)`
  Returns asiobaseinchannel.
  Called by: `.SetupCMAsio()` (`Console/setup.cs`), `.setupInOutBaseChannels()` (`Console/setup.cs`)
- **`.GetASIObaseoutchannel()`** — L194 — `public static int GetASIObaseoutchannel(RegistryKey key = null)`
  Returns asiobaseoutchannel.
  Called by: `.setupInOutBaseChannels()` (`Console/setup.cs`)
- **`.SetASIOinputmode()`** — L207 — `public static void SetASIOinputmode(int input_mode, RegistryKey key = null)`
  Sets asioinputmode.
  Called by: `.radCMASIO_mic_CheckedChanged()` (`Console/setup.cs`)
- **`.GetASIOinputmode()`** — L215 — `public static int GetASIOinputmode(RegistryKey key = null)`
  Returns asioinputmode.
  Called by: `.SetupCMAsio()` (`Console/setup.cs`)
- **`.deleteRegistryValue()`** — L229 — `private static void deleteRegistryValue(string valueName, RegistryKey key = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetASIODevices()`** — L242 — `public static List<AsioDeviceInfo> GetASIODevices()`
  Returns asiodevices.
  Called by: `.SetupCMAsio()` (`Console/setup.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCMASIOConfig.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
