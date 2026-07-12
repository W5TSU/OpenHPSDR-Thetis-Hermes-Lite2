# `ChannelMaster/cmasio.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Bridge to the cmASIO DLL for direct ASIO device I/O.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/netInterface.c` (calls ×2)
- Uses (outgoing references to other files):
  - `cmASIO/hostsample.cpp` (calls ×9)
  - `wdsp/rmatch.c` (calls ×9)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `ChannelMaster/obbuffs.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `ChannelMaster/ivac.c` (calls ×1)
  - `ChannelMaster/obbuffs.c` (calls ×1)
- Most-referenced symbols from other files: `create_cmasio()` (×1), `destroy_cmasio()` (×1), `asioIN()` (×1), `cm_asioStart()` (×1), `cm_asioStop()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_cmasio()`** — L47 — `void create_cmasio()`
  Constructor for the `cmasio` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_cmaster()` (`ChannelMaster/cmaster.c`)
- **`destroy_cmasio()`** — L104 — `void destroy_cmasio()`
  Destroys the `cmasio` block, freeing its allocated buffers.
  Called by: `destroy_cmaster()` (`ChannelMaster/cmaster.c`)
- **`asioIN()`** — L120 — `void asioIN(double* in_tx)`
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`asioOUT()`** — L137 — `void asioOUT(int id, int nsamples, double* buff)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CallbackASIO()`** — L150 — `void CallbackASIO(void* inputL, void* inputR, void* outputL, void* outputR)`
  [2.10.3.13]MW0LGE added input mode, so can use ch1(L), ch2(R), or both for input clamp and speed refactor
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`cm_asioStart()`** — L337 — `long cm_asioStart(int protocol)`
  Called by: `StartAudioNative()` (`ChannelMaster/netInterface.c`)
- **`cm_asioStop()`** — L358 — `long cm_asioStop()`
  Called by: `StopAudio()` (`ChannelMaster/netInterface.c`)
- **`getCMAstate()`** — L369 — `PORT int getCMAstate()`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`getCMAevents()`** — L379 — `PORT void getCMAevents(long* overFlowsIn, long* overFlowsOut, long* underFlowsIn, long* underFlowsOut)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`resetCMAevents()`** — L396 — `PORT void resetCMAevents()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmasio.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
