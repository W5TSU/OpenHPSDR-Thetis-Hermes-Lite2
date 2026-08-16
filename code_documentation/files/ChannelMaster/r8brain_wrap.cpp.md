# `ChannelMaster/r8brain_wrap.cpp`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** C-callable wrapper around the vendored r8brain-free-src `CDSPResampler24`, used by `radae.c` for its own internal sample-rate conversion.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/r8brain_wrap.h` (imports ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`r8b_prewarm()`** — L52 — `void r8b_prewarm(r8b_state* s)`
  Feed zero samples to the resampler until it is ready to emit output on its first real sample. r8brain absorbs the first getInLenBeforeOutPos(0) input samples into its polyphase filter delay-line before producing any output; without this pre-warm the first call to r8b_process_ff() with real audio…
  Called by: `r8b_create()` (same file), `r8b_clear()` (same file)
- **`r8b_create()`** — L68 — `r8b_handle r8b_create(double src_rate, double dst_rate, int max_in_len)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`r8b_destroy()`** — L104 — `void r8b_destroy(r8b_handle h)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`r8b_clear()`** — L119 — `void r8b_clear(r8b_handle h)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`r8b_process_ff()`** — L129 — `int r8b_process_ff(r8b_handle h, const float* in, int n_in, float* out, int out_cap)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

### Types

#### `r8b_state` (type, L38)


#### `resamp` (type, L40)

_No extracted members._

#### `in_d` (type, L41)

_No extracted members._

#### `max_in_len` (type, L42)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/r8brain_wrap.cpp`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
