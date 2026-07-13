# `ChannelMaster/nanotime.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Network bandwidth statistics and high-resolution timestamps.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/nanotimer.h` (imports ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`getPerfTicks()`** — L32 — `NANOTIMER_API __int64 getPerfTicks(void)`
  returns current tick count from the performance counter. Convert this tick to nanosecs with perfTicksToNanos 0 is returned on failure
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`getPerfFreq()`** — L46 — `NANOTIMER_API __int64 getPerfFreq(void)`
  returns freq in hertz of perf counter returns 0 on failure
  Called by: `perfTicksToNanos()` (same file)
- **`perfTicksToNanos()`** — L57 — `NANOTIMER_API __int64 perfTicksToNanos(__int64 ticks)`
  Called by: `printHLANano()` (same file)
- **`updateHLA()`** — L70 — `NANOTIMER_API void updateHLA(HLA_COUNTER *p, __int64 v)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`initHLA()`** — L82 — `NANOTIMER_API void initHLA(HLA_COUNTER *p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`printHLA()`** — L90 — `NANOTIMER_API void printHLA(HLA_COUNTER *p, /* FILE *f, */ unsigned char *prefix)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`printHLANano()`** — L109 — `NANOTIMER_API void printHLANano(HLA_COUNTER *p, unsigned char *prefix)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/nanotime.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
