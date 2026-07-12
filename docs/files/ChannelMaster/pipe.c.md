# `ChannelMaster/pipe.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/cmaster.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×7)
  - `ChannelMaster/ivac.c` (calls ×4)
  - `wdsp/nob.c` (calls ×3)
  - `wdsp/nobII.c` (calls ×3)
  - `ChannelMaster/tci.c` (calls ×3)
  - `wdsp/siphon.c` (calls ×3)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/analyzer.c` (calls ×1)
- Most-referenced symbols from other files: `xpipe()` (×1), `create_pipe()` (×1), `destroy_pipe()` (×1)

## Outline

### Functions

- `create_spc0()` — L32
- `destroy_spc0()` — L69
- `create_pipe()` — L79
- `destroy_pipe()` — L119
- `xplaywave()` — L133
- `xrecordwave()` — L142
- `xscope()` — L151
- `xpipe()` — L160
- `SendCBCreateScope()` — L257
- `SendCBScope()` — L263
- `SetScopeRun()` — L269
- `SendCBCreateWRecord()` — L275
- `SendCBWaveRecorder()` — L281
- `SetWaveRecorderRun()` — L287
- `SendCBCreateWPlay()` — L293
- `SendCBWavePlayer()` — L299
- `SetWavePlayerRun()` — L305
- `SetTopPan3Run()` — L311
- `SetTXVAC()` — L317

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/pipe.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
