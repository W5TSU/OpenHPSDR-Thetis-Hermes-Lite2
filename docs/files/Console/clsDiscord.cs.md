# `Console/clsDiscord.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Discord rich-presence integration.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×4)
  - `Console/MeterManager.cs` (calls ×2)
  - `Console/console.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/clsCountryData.cs` (calls ×1)
- Most-referenced symbols from other files: `.IsValidCallsign()` (×2), `.GetMessagesString()` (×1), `.SendMessage()` (×1), `.Shutdown()` (×1), `.SetEnabled()` (×1), `.SetCallsign()` (×1)

## Outline

### Types

#### `ThetisBotDiscord` (type, L56)

- `.cleanupOldMessages()` — L164
- `.loadChannelInfoFromGitHub()` — L205
- `.adjustRetryInterval()` — L235
- `.ConnectStart()` — L270
- `.ConnectStop()` — L291
- `.Shutdown()` — L327
- `.tryConnect()` — L334
- `.OnReady()` — L358
- `.OnConnected()` — L380
- `.OnDisconnected()` — L404
- `.messageReceived()` — L419
- `.processQueue()` — L430
- `.handleQueuedMessage()` — L442
- `.getAuthorName()` — L542
- `.SendMessage()` — L576
- `.getLastMessages()` — L629
- `.GetMessagesString()` — L673
- `.GetReceivableChannels()` — L686
- `.GetSendableChannels()` — L705
- `.OnMessageDeleted()` — L724
- `.IsValidCallsign()` — L759
- `.SetEnabled()` — L793
- `.SetCallsign()` — L804

#### `BotConfig` (type, L73)

_No extracted members._

#### `ChannelInfo` (type, L82)

_No extracted members._

#### `MessageInfo` (type, L94)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsDiscord.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
