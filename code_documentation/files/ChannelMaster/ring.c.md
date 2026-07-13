# `ChannelMaster/ring.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`ringbuffer_create()`** — L36 — `ringbuffer_t * ringbuffer_create (int sz)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`ringbuffer_free()`** — L62 — `void ringbuffer_free (ringbuffer_t * rb)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`ringbuffer_reset_size()`** — L69 — `void ringbuffer_reset_size (ringbuffer_t * rb, int sz)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`ringbuffer_reset()`** — L80 — `void ringbuffer_reset (ringbuffer_t * rb)`
  Called by: `ringbuffer_restart()` (same file)
- **`ringbuffer_clear()`** — L89 — `void ringbuffer_clear (ringbuffer_t * rb, int sz)`
  Called by: `ringbuffer_restart()` (same file)
- **`ringbuffer_restart()`** — L99 — `void ringbuffer_restart (ringbuffer_t * rb, int sz)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`ringbuffer_read_space()`** — L106 — `int ringbuffer_read_space (const ringbuffer_t * rb)`
  Called by: `ringbuffer_read()` (same file)
- **`ringbuffer_write_space()`** — L125 — `int ringbuffer_write_space (const ringbuffer_t * rb)`
  Called by: `ringbuffer_write()` (same file)
- **`ringbuffer_write()`** — L144 — `int ringbuffer_write (ringbuffer_t * rb, const double *src, int cnt)`
  Called by: `ringbuffer_clear()` (same file)
- **`ringbuffer_read()`** — L181 — `int ringbuffer_read (ringbuffer_t * rb, double *dest, int cnt)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/ring.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
