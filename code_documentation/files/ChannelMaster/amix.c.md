# `ChannelMaster/amix.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Audio mixers (monitor mix, multi-RX audio combination) with per-input gain and slew.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_amix()`** — L29 — `AMIX create_amix (int id, int run, int size, double** in0, double** in1, double* out, unsigned int what0, unsigned int what1, double volume)`
  Constructor for the `amix` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_amixEXT()` (same file)
- **`destroy_amix()`** — L51 — `void destroy_amix (AMIX a)`
  Destroys the `amix` block, freeing its allocated buffers.
  Called by: `destroy_amixEXT()` (same file)
- **`xamix()`** — L57 — `void xamix (AMIX a)`
  Runs the `amix` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xamixEXT()` (same file)
- **`create_amixEXT()`** — L101 — `void create_amixEXT (int id, int run, int size, unsigned int what0, unsigned int what1, double volume)`
  Constructor for the `amixEXT` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_amixEXT()`** — L106 — `void destroy_amixEXT (int id)`
  Destroys the `amixEXT` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xamixEXT()`** — L111 — `void xamixEXT (int id, double** in0, double** in1, double* out)`
  Runs the `amixEXT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAudioMixWhat()`** — L121 — `PORT void SetAudioMixWhat (int id, int bank, unsigned int stream, int state)`
  Sets audio mix what — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAudioMixSize()`** — L144 — `PORT void SetAudioMixSize (int id, int size)`
  Sets audio mix size — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAudioMixVolume()`** — L153 — `PORT void SetAudioMixVolume (int id, double volume)`
  Sets audio mix volume — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetAudioMixVol()`** — L168 — `PORT void SetAudioMixVol (int id, int bank, unsigned int stream, double vol)`
  Sets audio mix vol — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/amix.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
