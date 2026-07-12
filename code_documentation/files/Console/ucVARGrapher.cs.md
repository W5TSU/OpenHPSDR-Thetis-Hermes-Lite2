# `Console/ucVARGrapher.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Small graphing control (used for VAC variable-rate resampler diagnostics).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucVARGrapher` (type, L49)

- **`.AddDataPoint()`** — L104 — `public void AddDataPoint(double dataPoint)`
  Adds data point.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VARGraph_Paint()`** — L139 — `private void VARGraph_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `VARGraph` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.VARGraph_Resize()`** — L183 — `private void VARGraph_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `VARGraph` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucVARGrapher.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
