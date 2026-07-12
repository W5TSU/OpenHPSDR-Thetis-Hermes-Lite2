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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ThetisBotDiscord` (type, L56)

- **`.cleanupOldMessages()`** — L164 — `private static void cleanupOldMessages()`
  Called by: `.ConnectStart()` (same file)
- **`.loadChannelInfoFromGitHub()`** — L205 — `public static async Task loadChannelInfoFromGitHub()`
  Called by: `.ConnectStart()` (same file)
- **`.adjustRetryInterval()`** — L235 — `private static void adjustRetryInterval()`
  Called by: `.loadChannelInfoFromGitHub()` (same file)
- **`.ConnectStart()`** — L270 — `public static void ConnectStart()`
  Connects start.
  Called by: `.SetEnabled()` (same file)
- **`.ConnectStop()`** — L291 — `public static void ConnectStop()`
  Connects stop.
  Called by: `.Shutdown()` (same file), `.SetEnabled()` (same file), `.SetCallsign()` (same file)
- **`.Shutdown()`** — L327 — `public static void Shutdown()`
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.tryConnect()`** — L334 — `private static async Task tryConnect()`
  Called by: `.loadChannelInfoFromGitHub()` (same file), `.ConnectStart()` (same file)
- **`.OnReady()`** — L358 — `private static Task OnReady()`
  Handles/raises the ready event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnConnected()`** — L380 — `private static Task OnConnected()`
  Handles/raises the connected event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDisconnected()`** — L404 — `private static Task OnDisconnected(Exception exception)`
  Handles/raises the disconnected event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.messageReceived()`** — L419 — `private static Task messageReceived(SocketMessage message)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.processQueue()`** — L430 — `private static void processQueue()`
  Called by: `.ConnectStart()` (same file), `.SendMessage()` (same file)
- **`.handleQueuedMessage()`** — L442 — `private static void handleQueuedMessage(IMessage message)`
  Called by: `.processQueue()` (same file)
- **`.getAuthorName()`** — L542 — `private static string getAuthorName(IMessage message)`
  Returns author name.
  Called by: `.handleQueuedMessage()` (same file)
- **`.SendMessage()`** — L576 — `public static async Task SendMessage(string message, ulong channel_id = 0)`
  Sends message.
  Called by: `.sendMsg()` (`Console/MeterManager.cs`)
- **`.getLastMessages()`** — L629 — `private static async Task<List<MessageInfo>> getLastMessages(int n)`
  Returns last messages.
  Called by: `.OnReady()` (same file)
- **`.GetMessagesString()`** — L673 — `public static string GetMessagesString(ulong channel_id, int message = 0, bool include_author = true)`
  Returns messages string.
  Called by: `.GetReading()` (`Console/MeterManager.cs`)
- **`.GetReceivableChannels()`** — L686 — `public static List<(ulong ChannelID, string ChannelName)> GetReceivableChannels()`
  Returns receivable channels.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetSendableChannels()`** — L705 — `public static List<(ulong ChannelID, string ChannelName)> GetSendableChannels()`
  Returns sendable channels.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMessageDeleted()`** — L724 — `private static Task OnMessageDeleted(Cacheable<IMessage, ulong> cachedMessage, Cacheable<IMessageChannel, ulong> cachedChannel)`
  Handles/raises the message deleted event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsValidCallsign()`** — L759 — `public static bool IsValidCallsign(string callsign, out string country)`
  Called by: `.SetEnabled()` (same file), `.SetCallsign()` (same file), `.chkDiscordEnabled_CheckedChanged()` (`Console/setup.cs`), `.txtDiscordCallsign_TextChanged()` (`Console/setup.cs`)
- **`.SetEnabled()`** — L793 — `public static void SetEnabled(bool enabled)`
  return match.Success; }
  Called by: `.chkDiscordEnabled_CheckedChanged()` (`Console/setup.cs`)
- **`.SetCallsign()`** — L804 — `public static void SetCallsign(string callsign)`
  Sets callsign.
  Called by: `.txtDiscordCallsign_TextChanged()` (`Console/setup.cs`)

#### `BotConfig` (type, L73)

_No extracted members._

#### `ChannelInfo` (type, L82)

_No extracted members._

#### `MessageInfo` (type, L94)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsDiscord.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
