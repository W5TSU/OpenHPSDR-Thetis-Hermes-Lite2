# `wdsp/syncbuffs.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sample buffering between the audio callback world and DSP blocks.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`start_syncbthread()`** — L29 — `void start_syncbthread (SYNCB a)`
  Called by: `create_syncbuffs()` (same file), `SetSYNCBRingOutsize()` (same file)
- **`create_syncbuffs()`** — L35 — `SYNCB create_syncbuffs (int accept, int nstreams, int max_insize, int max_outsize, int outsize, double** out, void (*exf)(void))`
  Constructor for the `syncbuffs` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_syncbuffs()`** — L65 — `void destroy_syncbuffs (SYNCB a)`
  Destroys the `syncbuffs` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_syncbuffs()`** — L86 — `void flush_syncbuffs (SYNCB a)`
  Flushes (zeroes) the `syncbuffs` block’s internal buffers/state.
  Called by: `SetSYNCBRingOutsize()` (same file)
- **`Syncbound()`** — L97 — `void Syncbound (SYNCB a, int nsamples, double** in)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`syncbdata()`** — L132 — `void syncbdata (SYNCB a)`
  Called by: `syncb_main()` (same file)
- **`syncb_main()`** — L162 — `void syncb_main (void *p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetSYNCBRingOutsize()`** — L175 — `void SetSYNCBRingOutsize (SYNCB a, int size)`
  Sets syncbring outsize — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_dumfilt()`** — L199 — `DUMFILT create_dumfilt (int run, int delay, int opsize, double* in, double* out)`
  Constructor for the `dumfilt` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_dumfilt()`** — L214 — `void destroy_dumfilt (DUMFILT a)`
  Destroys the `dumfilt` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_dumfilt()`** — L220 — `void flush_dumfilt (DUMFILT a)`
  Flushes (zeroes) the `dumfilt` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xdumfilt()`** — L227 — `void xdumfilt (DUMFILT a)`
  Runs the `dumfilt` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/syncbuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
