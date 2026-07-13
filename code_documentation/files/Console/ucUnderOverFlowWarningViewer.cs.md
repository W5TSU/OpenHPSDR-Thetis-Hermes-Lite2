# `Console/ucUnderOverFlowWarningViewer.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Audio buffer underflow/overflow warning indicator.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucUnderOverFlowWarningViewer` (type, L54)

- **`.UnderOverFlowWarningViewer_Load()`** — L72 — `private void UnderOverFlowWarningViewer_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `UnderOverFlowWarningViewer` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setColours()`** — L81 — `private void setColours(bool forceUpdate = false)`
  Sets colours.
  Called by: `.clearIssues()` (same file)
- **`.clearIssues()`** — L126 — `private void clearIssues()`
  Called by: `.UnderOverFlowWarningViewer_Load()` (same file), `.UnderOverFlowWarningViewer_Click()` (same file)
- **`.fadeBackground()`** — L172 — `private Color fadeBackground(Color c)`
  Called by: `.tmrFade_Tick()` (same file)
- **`.tmrFade_Tick()`** — L194 — `private void tmrFade_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `tmrFade` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UnderOverFlowWarningViewer_Paint()`** — L216 — `private void UnderOverFlowWarningViewer_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `UnderOverFlowWarningViewer` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UnderOverFlowWarningViewer_Click()`** — L263 — `private void UnderOverFlowWarningViewer_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `UnderOverFlowWarningViewer` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucUnderOverFlowWarningViewer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
