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

### Functions

- `fnv1a_hash64()` — L55
- `fnv1a_hash32()` — L68
- `remove_impulse_cache_tail()` — L92
- `free_impulse_cache()` — L112
- `get_impulse_cache_entry()` — L127
- `add_impulse_to_cache()` — L163
- `save_impulse_cache()` — L186
- `read_impulse_cache()` — L215
- `use_impulse_cache()` — L261
- `init_impulse_cache()` — L269
- `destroy_impulse_cache()` — L282

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/impulse_cache.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
