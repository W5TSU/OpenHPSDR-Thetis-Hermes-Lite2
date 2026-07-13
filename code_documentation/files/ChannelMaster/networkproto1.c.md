# `ChannelMaster/networkproto1.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** HPSDR Protocol-1 UDP implementation: socket setup, packet build/parse, EP2/EP4/EP6 endpoint handling, sequence tracking.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/netInterface.c` (calls ×1)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/netInterface.c` (calls ×6)
  - `ChannelMaster/pro.c` (calls ×5)
  - `ChannelMaster/network.c` (calls ×3)
  - `ChannelMaster/router.c` (calls ×3)
  - `ChannelMaster/cmbuffs.c` (calls ×2)
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/bandwidth_monitor.c` (calls ×1)
  - `ChannelMaster/network.h` (imports ×1)
  - `ChannelMaster/pro.h` (imports ×1)
- Most-referenced symbols from other files: `SendStartToMetis()` (×1), `SendStopToMetis()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`SendStartToMetis()`** — L38 — `int SendStartToMetis(void)`
  Called by: `StartAudioNative()` (`ChannelMaster/netInterface.c`)
- **`SendStopToMetis()`** — L76 — `PORT int SendStopToMetis()`
  Called by: `StopReadThread()` (`ChannelMaster/network.c`)
- **`ForceCandCFrames()`** — L111 — `void ForceCandCFrames(int count, int c0, int vfofreq)`
  Called by: `ForceCandCFrame()` (same file)
- **`ForceCandCFrame()`** — L139 — `void ForceCandCFrame(int count)`
  Called by: `SendStartToMetis()` (same file), `MetisReadThreadMainLoop()` (same file), `MetisReadThreadMainLoop_HL2()` (same file)
- **`MetisReadDirect()`** — L146 — `int MetisReadDirect(unsigned char* bufp)`
  Called by: `SendStartToMetis()` (same file), `MetisReadThreadMainLoop()` (same file), `MetisReadThreadMainLoop_HL2()` (same file)
- **`MetisWriteFrame()`** — L221 — `int MetisWriteFrame(int endpoint, char* bufp)`
  Called by: `ForceCandCFrames()` (same file), `WriteMainLoop()` (same file), `WriteMainLoop_HL2()` (same file)
- **`MetisReadThreadMain()`** — L245 — `DWORD WINAPI MetisReadThreadMain(LPVOID n)`
  this is the main thread that reads data
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`twist()`** — L268 — `void twist (int nsamples, int stream0, int stream1, int source)`
  Called by: `MetisReadThreadMainLoop()` (same file), `MetisReadThreadMainLoop_HL2()` (same file)
- **`MetisReadThreadMainLoop()`** — L281 — `void MetisReadThreadMainLoop(void)`
  Called by: `MetisReadThreadMain()` (same file)
- **`MetisReadThreadMainLoop_HL2()`** — L427 — `void MetisReadThreadMainLoop_HL2(void)`
  Called by: `MetisReadThreadMain()` (same file)
- **`WriteMainLoop()`** — L593 — `void WriteMainLoop(char* bufp)`
  Called by: `sendProtocol1Samples()` (same file)
- **`WriteMainLoop_HL2()`** — L874 — `void WriteMainLoop_HL2(char* bufp)`
  Called by: `sendProtocol1Samples()` (same file)
- **`sendProtocol1Samples()`** — L1209 — `DWORD WINAPI sendProtocol1Samples(LPVOID n)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/networkproto1.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
