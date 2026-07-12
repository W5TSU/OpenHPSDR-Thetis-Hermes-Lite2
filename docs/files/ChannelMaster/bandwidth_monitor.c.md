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

### Functions

- `now_ms()` — L54
- `bandwidth_monitor_reset()` — L59
- `bandwidth_monitor_in()` — L74
- `bandwidth_monitor_out()` — L80
- `compute_bps()` — L86
- `GetInboundBps()` — L115
- `GetOutboundBps()` — L120

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/bandwidth_monitor.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
