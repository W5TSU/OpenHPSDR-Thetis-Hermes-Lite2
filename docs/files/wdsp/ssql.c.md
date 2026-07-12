# `wdsp/ssql.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM squelch, FM squelch, and syllabic (voice-detecting) squelch.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/cblock.c` (calls ×4)
  - `wdsp/iir.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_ssql()` (×1), `destroy_ssql()` (×1), `flush_ssql()` (×1), `xssql()` (×1), `setSamplerate_ssql()` (×1), `setBuffers_ssql()` (×1), `setSize_ssql()` (×1)

## Outline

### Functions

- `create_ftov()` — L35
- `destroy_ftov()` — L56
- `flush_ftov()` — L62
- `xftov()` — L70
- `compute_ssql_slews()` — L113
- `calc_ssql()` — L133
- `decalc_ssql()` — L162
- `create_ssql()` — L177
- `destroy_ssql()` — L202
- `flush_ssql()` — L208
- `xssql()` — L230
- `setBuffers_ssql()` — L302
- `setSamplerate_ssql()` — L310
- `setSize_ssql()` — L317
- `SetRXASSQLRun()` — L330
- `SetRXASSQLThreshold()` — L338
- `SetRXASSQLTauMute()` — L348
- `SetRXASSQLTauUnMute()` — L360

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/ssql.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
