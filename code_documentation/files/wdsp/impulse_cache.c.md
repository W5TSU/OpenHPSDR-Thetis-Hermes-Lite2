# `wdsp/impulse_cache.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FFTW wisdom generation/caching and FIR impulse-response caching for fast startup.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/fir.c` (calls ×4)
  - `wdsp/eq.c` (calls ×2)
  - `wdsp/fcurve.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `add_impulse_to_cache()` (×4), `get_impulse_cache_entry()` (×4)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`fnv1a_hash64()`** — L55 — `uint64_t fnv1a_hash64(const void* data, size_t len)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`fnv1a_hash32()`** — L68 — `uint32_t fnv1a_hash32(const void* data, size_t len)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`remove_impulse_cache_tail()`** — L92 — `void remove_impulse_cache_tail(size_t bucket)`
  Called by: `add_impulse_to_cache()` (same file)
- **`free_impulse_cache()`** — L112 — `void free_impulse_cache(void)`
  Called by: `read_impulse_cache()` (same file), `destroy_impulse_cache()` (same file)
- **`get_impulse_cache_entry()`** — L127 — `double* get_impulse_cache_entry(size_t bucket, HASH_T hash, int N)`
  Called by: `eq_impulse()` (`wdsp/eq.c`), `fc_impulse()` (`wdsp/fcurve.c`), `fir_bandpass()` (`wdsp/fir.c`), `mp_imp()` (`wdsp/fir.c`)
- **`add_impulse_to_cache()`** — L163 — `void add_impulse_to_cache(size_t bucket, HASH_T hash, int N, double* impulse)`
  Called by: `eq_impulse()` (`wdsp/eq.c`), `fc_impulse()` (`wdsp/fcurve.c`), `fir_bandpass()` (`wdsp/fir.c`), `mp_imp()` (`wdsp/fir.c`)
- **`save_impulse_cache()`** — L186 — `PORT int save_impulse_cache(const char* path)`
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`read_impulse_cache()`** — L215 — `PORT int read_impulse_cache(const char* path)`
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`use_impulse_cache()`** — L261 — `PORT void use_impulse_cache(int use)`
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`init_impulse_cache()`** — L269 — `PORT void init_impulse_cache(int use)`
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`destroy_impulse_cache()`** — L282 — `PORT void destroy_impulse_cache(void)`
  Destroys the `impulse_cache` block, freeing its allocated buffers.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/impulse_cache.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
