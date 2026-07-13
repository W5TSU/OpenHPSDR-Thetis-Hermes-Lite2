# `wdsp/meterlog10.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/analyzer.c` (calls ×3)
  - `wdsp/cfcomp.c` (calls ×1)
  - `wdsp/emnr.c` (calls ×1)
  - `wdsp/meter.c` (calls ×1)
  - `wdsp/siphon.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `mlog10()` (×7)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`mlog10()`** — L547 — `inline double mlog10 (double val)`
  Called by: `avenger()` (`wdsp/analyzer.c`), `DetectMaxBin()` (`wdsp/analyzer.c`), `CalcBandwidthNormalization()` (`wdsp/analyzer.c`), `GetTXACFCOMPDisplayCompression()` (`wdsp/cfcomp.c`), `getZeta()` (`wdsp/emnr.c`), `xmeter()` (`wdsp/meter.c`) — and 1 more

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/meterlog10.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
