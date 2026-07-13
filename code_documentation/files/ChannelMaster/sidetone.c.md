# `ChannelMaster/sidetone.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** CW sidetone generation.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
  - `ChannelMaster/netInterface.c` (calls ×5)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `keySidetone()` (×2), `create_sidetone()` (×1), `destroy_sidetone()` (×1), `xsidetone()` (×1), `setSidetoneRate()` (×1), `setSidetoneSize()` (×1), `SetSidetoneWPM()` (×1), `SetSidetoneRun()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_tone1()`** — L31 — `void calc_tone1(SIDETONE a)`
  Called by: `calc_sidetone()` (same file), `SetSidetonePitch()` (same file)
- **`calc_rising_edge()`** — L41 — `void calc_rising_edge(SIDETONE a)`
  Called by: `calc_sidetone()` (same file), `SetSidetoneEdgetype()` (same file), `SetSidetoneEdgelength()` (same file)
- **`decalc_rising_edge()`** — L60 — `void decalc_rising_edge(SIDETONE a)`
  Called by: `decalc_sidetone()` (same file), `SetSidetoneEdgetype()` (same file), `SetSidetoneEdgelength()` (same file)
- **`calc_falling_edge()`** — L65 — `void calc_falling_edge(SIDETONE a)`
  Called by: `calc_sidetone()` (same file), `SetSidetoneEdgetype()` (same file), `SetSidetoneEdgelength()` (same file)
- **`decalc_falling_edge()`** — L84 — `void decalc_falling_edge(SIDETONE a)`
  Called by: `decalc_sidetone()` (same file), `SetSidetoneEdgetype()` (same file), `SetSidetoneEdgelength()` (same file)
- **`calc_wpm_times()`** — L89 — `void calc_wpm_times(SIDETONE a)`
  Called by: `calc_sidetone()` (same file), `SetSidetoneWPM()` (same file)
- **`calc_sidetone()`** — L98 — `void calc_sidetone(SIDETONE a)`
  Called by: `setSidetoneRate()` (same file)
- **`decalc_sidetone()`** — L108 — `void decalc_sidetone(SIDETONE a)`
  Called by: `destroy_sidetone()` (same file), `setSidetoneRate()` (same file)
- **`create_sidetone()`** — L114 — `SIDETONE create_sidetone( int id, int run_st, int run_tx, int rate,`
  Constructor for the `sidetone` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_xmtr()` (`ChannelMaster/cmaster.c`)
- **`destroy_sidetone()`** — L157 — `void destroy_sidetone(int id)`
  Destroys the `sidetone` block, freeing its allocated buffers.
  Called by: `destroy_xmtr()` (`ChannelMaster/cmaster.c`)
- **`osc_init()`** — L172 — `void osc_init(SIDETONE a)`
  Called by: `xsidetone()` (same file)
- **`osc()`** — L178 — `void osc(SIDETONE a)`
  Called by: `xsidetone()` (same file)
- **`xsidetone()`** — L193 — `void xsidetone(int id)`
  Runs the `sidetone` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`setSidetoneRate()`** — L290 — `void setSidetoneRate(int id, int rate)`
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`setSidetoneSize()`** — L300 — `void setSidetoneSize(int id, int size)`
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetSidetoneSelectKey()`** — L308 — `PORT void SetSidetoneSelectKey(int id, int select)`
  Sets sidetone select key — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`keySidetone()`** — L317 — `PORT void keySidetone(int id, int key_select, int state)`
  Called by: `SetCWX()` (`ChannelMaster/netInterface.c`), `ReadThreadMainLoop()` (`ChannelMaster/network.c`)
- **`makedotSidetone()`** — L327 — `PORT void makedotSidetone(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`makedashSidetone()`** — L336 — `PORT void makedashSidetone(int id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetCWtxIQpolarity()`** — L345 — `PORT void SetCWtxIQpolarity(int id, int polarity)`
  Sets cwtx iqpolarity — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetSidetoneVolume()`** — L354 — `PORT void SetSidetoneVolume(int id, double volume)`
  Sets sidetone volume — API setter, typically called from the console via P/Invoke.
  Called by: `SetCWSidetoneVolume()` (`ChannelMaster/netInterface.c`)
- **`SetCWtxVolume()`** — L363 — `PORT void SetCWtxVolume(int id, double volume)`
  Sets cwtx volume — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetSidetoneWPM()`** — L373 — `PORT void SetSidetoneWPM(int id, int wpm)`
  Sets sidetone wpm — API setter, typically called from the console via P/Invoke.
  Called by: `SetCWKeyerSpeed()` (`ChannelMaster/netInterface.c`)
- **`SetSidetoneRun()`** — L383 — `PORT void SetSidetoneRun(int id, int run)`
  Sets sidetone run — API setter, typically called from the console via P/Invoke.
  Called by: `EnableCWKeyer()` (`ChannelMaster/netInterface.c`)
- **`SetCWtxRun()`** — L392 — `PORT void SetCWtxRun(int id, int run)`
  Sets cwtx run — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetSidetonePitch()`** — L401 — `PORT void SetSidetonePitch(int id, double pitch)`
  Sets sidetone pitch — API setter, typically called from the console via P/Invoke.
  Called by: `SetCWSidetoneFreq()` (`ChannelMaster/netInterface.c`)
- **`SetSidetoneEdgetype()`** — L411 — `PORT void SetSidetoneEdgetype(int id, int type)`
  Sets sidetone edgetype — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetSidetoneEdgelength()`** — L424 — `PORT void SetSidetoneEdgelength(int id, double length)`
  Sets sidetone edgelength — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/sidetone.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
