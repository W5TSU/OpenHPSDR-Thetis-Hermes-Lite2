# `Console/filter.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX filter preset model per mode, the filter-edit form, and the filter-set manager.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×2)
  - `Console/FilterForm.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×10)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `FilterPreset` (type, L47)

- **`.SetLow()`** — L60 — `public void SetLow(Filter f, int val)`
  Sets low.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetHigh()`** — L65 — `public void SetHigh(Filter f, int val)`
  Sets high.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetName()`** — L70 — `public void SetName(Filter f, string n)`
  Sets name.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetFilter()`** — L75 — `public void SetFilter(Filter f, int l, int h, string n)`
  Sets filter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetLow()`** — L82 — `public int GetLow(Filter f)`
  Returns low.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHigh()`** — L87 — `public int GetHigh(Filter f)`
  Returns high.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetBW()`** — L92 — `public int GetBW(Filter f)`
  Returns bw.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetName()`** — L97 — `public string GetName(Filter f)`
  Returns name.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToString()`** — L109 — `public string ToString(Filter f)`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/filter.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
