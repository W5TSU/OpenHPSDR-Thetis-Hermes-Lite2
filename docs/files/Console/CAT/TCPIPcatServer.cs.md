# `Console/CAT/TCPIPcatServer.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** CAT over TCP/IP server (multiple client connections).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×6, references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×3)
  - `Console/MeterManager.cs` (calls ×2)
  - `Console/frmLog.Designer.cs` (references ×1)
- Most-referenced symbols from other files: `.StopServer()` (×2), `.StartServer()` (×1), `.SendToClients()` (×1), `.ShowLog()` (×1), `.CloseLog()` (×1)

## Outline

### Types

#### `TCPIPSocketListener` (type, L19)

- `.StartSocketListener()` — L67
- `.SocketListenerThreadStart()` — L79
- `.shouldSend()` — L169
- `.addClientId()` — L209
- `.removeClientId()` — L227
- `.StopSocketListener()` — L245
- `.IsMarkedForDeletion()` — L267
- `.IsDisconnected()` — L271
- `.ParseReceiveBuffer()` — L275
- `.processClientData()` — L298
- `.internal_send_data()` — L358
- `.SendData()` — L383
- `.checkClientCommInterval()` — L397

#### `BroadcastItem` (type, L43)

_No extracted members._

#### `TCPIPcatServer` (type, L416)

- `.Init()` — L489
- `.StartServer()` — L506
- `.SendToClients()` — L539
- `.StopServer()` — L569
- `.StopAllSocketListers()` — L632
- `.ServerThreadStart()` — L650
- `.PurgingThreadStart()` — L691
- `.ClientConnectedHandler()` — L727
- `.ClientDisconnectedHandler()` — L731
- `.ClientErrorHandler()` — L735
- `.ShowLog()` — L765
- `.CloseLog()` — L770

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/TCPIPcatServer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
