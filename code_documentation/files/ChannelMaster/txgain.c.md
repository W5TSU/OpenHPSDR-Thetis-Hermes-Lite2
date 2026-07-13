# `ChannelMaster/txgain.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** VOX detection and TX gain staging.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×4)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_txgain()` (×1), `destroy_txgain()` (×1), `xtxgain()` (×1), `SetTXGainSize()` (×1), `SetAmpProtectADCValue()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_txgain()`** — L29 — `TXGAIN create_txgain( int run_fixed, int run_amp_protect, int size, double* in,`
  Constructor for the `txgain` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_xmtr()` (`ChannelMaster/cmaster.c`)
- **`destroy_txgain()`** — L58 — `void destroy_txgain(TXGAIN a)`
  Destroys the `txgain` block, freeing its allocated buffers.
  Called by: `destroy_xmtr()` (`ChannelMaster/cmaster.c`)
- **`xtxgain()`** — L65 — `void xtxgain(TXGAIN a)`
  Runs the `txgain` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`SetTXGainSize()`** — L111 — `void SetTXGainSize(TXGAIN p, int size)`
  Sets txgain size — API setter, typically called from the console via P/Invoke.
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetTXFixedGainRun()`** — L116 — `PORT void SetTXFixedGainRun(int txid, int run)`
  Sets txfixed gain run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTXFixedGain()`** — L126 — `PORT void SetTXFixedGain(int txid, double Igain, double Qgain)`
  Sets txfixed gain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetAmpProtectADCValue()`** — L137 — `void SetAmpProtectADCValue (int txid, int value)`
  call when new ADC value arrives from network
  Called by: `ReadThreadMainLoop()` (`ChannelMaster/network.c`)
- **`GetAndResetAmpProtect()`** — L146 — `PORT int GetAndResetAmpProtect(int txid)`
  Returns and reset amp protect — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetAmpProtectRun()`** — L153 — `PORT void SetAmpProtectRun(int txid, int run)`
  Sets amp protect run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetADCSupply()`** — L163 — `PORT void SetADCSupply(int txid, int v)`
  Sets adcsupply — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/txgain.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
