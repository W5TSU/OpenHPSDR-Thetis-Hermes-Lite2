# `ChannelMaster/bandwidth_monitor.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Network bandwidth statistics and high-resolution timestamps.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/network.c` (calls ×2)
  - `ChannelMaster/netInterface.c` (calls ×1)
  - `ChannelMaster/networkproto1.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/bandwidth_monitor.h` (imports ×1)
- Most-referenced symbols from other files: `bandwidth_monitor_in()` (×2), `bandwidth_monitor_reset()` (×1), `bandwidth_monitor_out()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`now_ms()`** — L54 — `static int64_t now_ms(void)`
  Called by: `compute_bps()` (same file)
- **`bandwidth_monitor_reset()`** — L59 — `void bandwidth_monitor_reset(void)`
  Called by: `create_rnet()` (`ChannelMaster/netInterface.c`)
- **`bandwidth_monitor_in()`** — L74 — `void bandwidth_monitor_in(int bytes)`
  Called by: `ReadUDPFrame()` (`ChannelMaster/network.c`), `MetisReadDirect()` (`ChannelMaster/networkproto1.c`)
- **`bandwidth_monitor_out()`** — L80 — `void bandwidth_monitor_out(int bytes)`
  Called by: `sendPacket()` (`ChannelMaster/network.c`)
- **`compute_bps()`** — L86 — `static double compute_bps(volatile LONG64* total_bytes, volatile LONG64* last_bytes, volatile LONG64* last_ms, volatile double* last_bps)`
  Called by: `GetInboundBps()` (same file), `GetOutboundBps()` (same file)
- **`GetInboundBps()`** — L115 — `PORT double GetInboundBps(void)`
  Returns inbound bps — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.
- **`GetOutboundBps()`** — L120 — `PORT double GetOutboundBps(void)`
  Returns outbound bps — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/NetworkIOImports.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/bandwidth_monitor.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
