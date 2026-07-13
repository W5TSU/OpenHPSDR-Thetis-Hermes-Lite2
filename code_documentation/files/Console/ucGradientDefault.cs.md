# `Console/ucGradientDefault.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** Meter-related picker controls (open-collector LED strip, signal source, linear-gradient color pickers).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucGradientDefault` (type, L49)

- **`.populateGradientList()`** — L93 — `private void populateGradientList()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnSet_Click()`** — L104 — `private void btnSet_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSet` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.OnEnabledChanged()`** — L116 — `protected override void OnEnabledChanged(EventArgs e)`
  Handles/raises the enabled changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setControlState()`** — L122 — `private void setControlState(Control parent, bool enabled)`
  Sets control state.
  Called by: `.OnEnabledChanged()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucGradientDefault.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
