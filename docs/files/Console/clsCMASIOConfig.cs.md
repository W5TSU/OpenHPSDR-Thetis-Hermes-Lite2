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

### Types

#### `AsioDeviceInfo` (type, L48)

- `.ToString()` — L63

#### `CMASIOConfig` (type, L69)

- `.openRegistryKey()` — L76
- `.DoesRegistryValueExist()` — L98
- `.GetASIOdrivername()` — L109
- `.SetASIOdrivername()` — L122
- `.GetASIOblocknum()` — L130
- `.GetASIOlockmode()` — L143
- `.SetASIOblocknum()` — L156
- `.SetASIObaseinchannel()` — L165
- `.SetASIObaseoutchannel()` — L173
- `.GetASIObaseinchannel()` — L181
- `.GetASIObaseoutchannel()` — L194
- `.SetASIOinputmode()` — L207
- `.GetASIOinputmode()` — L215
- `.deleteRegistryValue()` — L229
- `.GetASIODevices()` — L242

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCMASIOConfig.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
