# `ChannelMaster/network.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** HPSDR Protocol-1 UDP implementation: socket setup, packet build/parse, EP2/EP4/EP6 endpoint handling, sequence tracking.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/netInterface.c` (calls ×64)
  - `ChannelMaster/networkproto1.c` (calls ×3)
  - `ChannelMaster/obbuffs.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/netInterface.c` (calls ×4)
  - `ChannelMaster/bandwidth_monitor.c` (calls ×2)
  - `ChannelMaster/cmbuffs.c` (calls ×1)
  - `ChannelMaster/cmsetup.c` (calls ×1)
  - `ChannelMaster/network.h` (imports ×1)
  - `ChannelMaster/networkproto1.c` (calls ×1)
  - `wdsp/analyzer.c` (calls ×1)
  - `ChannelMaster/router.c` (calls ×1)
  - `ChannelMaster/sidetone.c` (calls ×1)
  - `ChannelMaster/txgain.c` (calls ×1)
- Most-referenced symbols from other files: `CmdHighPriority()` (×26), `CmdTx()` (×23), `CmdGeneral()` (×7), `CmdRx()` (×4), `sendPacket()` (×3), `StopReadThread()` (×2), `StartReadThread()` (×1), `IOThreadStop()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`initWSA()`** — L53 — `int initWSA(void)`
  Called by: `nativeInitMetis()` (same file)
- **`DeInitMetisSockets()`** — L75 — `PORT void DeInitMetisSockets()`
  Called by: `StopReadThread()` (same file)
- **`nativeInitMetis()`** — L87 — `PORT int nativeInitMetis(char* netaddr, int port, char* localaddr, int localport, int protocol, int model_id, int p2hw_uses_differnt_ports)`
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`GetMetisIPAddr()`** — L348 — `PORT int GetMetisIPAddr(void)`
  Returns metis ipaddr — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`SendStart()`** — L353 — `int SendStart(void)`
  Called by: `StartReadThread()` (same file)
- **`SendStop()`** — L363 — `int SendStop(void)`
  Called by: `StopReadThread()` (same file), `ReadThreadMainLoop()` (same file)
- **`StartReadThread()`** — L371 — `int StartReadThread(void)`
  returns 0 on success, !0 on failure */
  Called by: `StartAudioNative()` (`ChannelMaster/netInterface.c`)
- **`StopReadThread()`** — L389 — `void StopReadThread()`
  Called by: `StartReadThread()` (same file), `StartAudioNative()` (`ChannelMaster/netInterface.c`), `StopAudio()` (`ChannelMaster/netInterface.c`)
- **`addSnapShot()`** — L398 — `void addSnapShot(int rx, unsigned int received_seqnum, unsigned int last_seqnum)`
  Called by: `ReadUDPFrame()` (same file)
- **`storeRXSeqDelta()`** — L456 — `void storeRXSeqDelta(int rx, unsigned int received_seqnum)`
  Called by: `ReadUDPFrame()` (same file)
- **`ReadUDPFrame()`** — L472 — `int ReadUDPFrame(unsigned char* bufp)`
  Called by: `ReadThreadMainLoop()` (same file)
- **`ReadThreadMainLoop()`** — L635 — `void ReadThreadMainLoop()`
  Called by: `ReadThreadMain()` (same file)
- **`CmdGeneral()`** — L812 — `void CmdGeneral()`
  Called by: `SendStart()` (same file), `KeepAliveLoop()` (same file), `DisablePA()` (`ChannelMaster/netInterface.c`), `SetEERPWMmin()` (`ChannelMaster/netInterface.c`), `SetEERPWMmax()` (`ChannelMaster/netInterface.c`), `SetWBPacketsPerFrame()` (`ChannelMaster/netInterface.c`) — and 3 more
- **`CmdHighPriority()`** — L904 — `void CmdHighPriority()`
  Called by: `SendStart()` (same file), `SendStop()` (same file), `SetPttOut()` (`ChannelMaster/netInterface.c`), `SetTRXrelay()` (`ChannelMaster/netInterface.c`), `SetOCBits()` (`ChannelMaster/netInterface.c`), `SetOCExtraBits()` (`ChannelMaster/netInterface.c`) — and 22 more
- **`CmdRx()`** — L1056 — `PORT void CmdRx()`
  Called by: `SendStart()` (same file), `SetADCDither()` (`ChannelMaster/netInterface.c`), `SetADCRandom()` (`ChannelMaster/netInterface.c`), `EnableRx()` (`ChannelMaster/netInterface.c`), `SetRxADC()` (`ChannelMaster/netInterface.c`)
- **`CmdTx()`** — L1172 — `void CmdTx()`
  Called by: `SendStart()` (same file), `EnableEClassModulation()` (`ChannelMaster/netInterface.c`), `SetMicBoost()` (`ChannelMaster/netInterface.c`), `SetMicXlr()` (`ChannelMaster/netInterface.c`), `SetLineIn()` (`ChannelMaster/netInterface.c`), `SetMicTipRing()` (`ChannelMaster/netInterface.c`) — and 18 more
- **`sendOutbound()`** — L1241 — `void sendOutbound(int id, double* out)`
  Called by: `ob_main()` (`ChannelMaster/obbuffs.c`)
- **`WriteUDPFrame()`** — L1347 — `void WriteUDPFrame(int id, char* bufp, int buflen)`
  Called by: `sendOutbound()` (same file)
- **`sendPacket()`** — L1386 — `int sendPacket(SOCKET sock, char* data, int length, int port)`
  Called by: `CmdGeneral()` (same file), `CmdHighPriority()` (same file), `CmdRx()` (same file), `CmdTx()` (same file), `WriteUDPFrame()` (same file), `SendStartToMetis()` (`ChannelMaster/networkproto1.c`) — and 2 more
- **`KeepAliveLoop()`** — L1408 — `void KeepAliveLoop(void)`
  Called by: `KeepAliveMain()` (same file)
- **`IOThreadStop()`** — L1438 — `int IOThreadStop()`
  stops and kill's the IOThread
  Called by: `StopAudio()` (`ChannelMaster/netInterface.c`)
- **`ReadThreadMain()`** — L1478 — `DWORD WINAPI ReadThreadMain(LPVOID n)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`KeepAliveMain()`** — L1493 — `DWORD WINAPI KeepAliveMain(LPVOID n)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/network.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
