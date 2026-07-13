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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `TCPIPSocketListener` (type, L19)

- **`.StartSocketListener()`** — L67 — `public void StartSocketListener()`
  Starts socket listener.
  Called by: `.ServerThreadStart()` (same file)
- **`.SocketListenerThreadStart()`** — L79 — `private void SocketListenerThreadStart()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.shouldSend()`** — L169 — `private bool shouldSend(List<string> id_limit)`
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.addClientId()`** — L209 — `public void addClientId(string id)`
  Called by: `.processClientData()` (same file)
- **`.removeClientId()`** — L227 — `public void removeClientId(string id)`
  Called by: `.processClientData()` (same file)
- **`.StopSocketListener()`** — L245 — `public void StopSocketListener()`
  Stops socket listener.
  Called by: `.checkClientCommInterval()` (same file), `.StopAllSocketListers()` (same file), `.PurgingThreadStart()` (same file)
- **`.IsMarkedForDeletion()`** — L267 — `public bool IsMarkedForDeletion()`
  Called by: `.PurgingThreadStart()` (same file)
- **`.IsDisconnected()`** — L271 — `public bool IsDisconnected()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ParseReceiveBuffer()`** — L275 — `private void ParseReceiveBuffer(Byte[] byteBuffer, int size)`
  Parses receive buffer.
  Called by: `.SocketListenerThreadStart()` (same file)
- **`.processClientData()`** — L298 — `private void processClientData(string sInboundCatCommand)`
  Called by: `.ParseReceiveBuffer()` (same file)
- **`.internal_send_data()`** — L358 — `private void internal_send_data(string oneLine)`
  Called by: `.SocketListenerThreadStart()` (same file), `.processClientData()` (same file)
- **`.SendData()`** — L383 — `public void SendData(string data, List<string> id_limit = null)`
  Sends data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkClientCommInterval()`** — L397 — `private void checkClientCommInterval(object o)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `BroadcastItem` (type, L43)

_No extracted members._

#### `TCPIPcatServer` (type, L416)

- **`.Init()`** — L489 — `private void Init(IPEndPoint ipNport)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StartServer()`** — L506 — `public void StartServer(Console c, bool bTCPIPcatWelcomeMessage = true)`
  Starts server.
  Called by: `.SetupTCPIPCat()` (`Console/console.cs`)
- **`.SendToClients()`** — L539 — `public void SendToClients(string sMsg, List<string> id_limit = null)`
  Sends to clients.
  Called by: `.on_send_floodcontrol_message()` (`Console/console.cs`)
- **`.StopServer()`** — L569 — `public void StopServer()`
  Stops server.
  Called by: `.StartServer()` (same file), `.SetupTCPIPCat()` (`Console/console.cs`), `.Console_Closing()` (`Console/console.cs`)
- **`.StopAllSocketListers()`** — L632 — `private void StopAllSocketListers()`
  Stops all socket listers.
  Called by: `.StopServer()` (same file)
- **`.ServerThreadStart()`** — L650 — `private void ServerThreadStart()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PurgingThreadStart()`** — L691 — `private void PurgingThreadStart()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClientConnectedHandler()`** — L727 — `private void ClientConnectedHandler()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClientDisconnectedHandler()`** — L731 — `private void ClientDisconnectedHandler()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClientErrorHandler()`** — L735 — `private void ClientErrorHandler(SocketException se)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowLog()`** — L765 — `public void ShowLog()`
  Shows log.
  Called by: `.ShowTCPIPCatLog()` (`Console/console.cs`)
- **`.CloseLog()`** — L770 — `public void CloseLog()`
  Closes log.
  Called by: `.SetupTCPIPCat()` (`Console/console.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/TCPIPcatServer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
