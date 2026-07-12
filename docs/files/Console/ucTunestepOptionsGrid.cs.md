# `Console/ucTunestepOptionsGrid.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Tuning-step definitions and their configuration grid.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/TuneStep.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucTunestepOptionsGrid` (type, L48)

- **`.initialize_checkboxes()`** — L62 — `private void initialize_checkboxes()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Init()`** — L99 — `public void Init(List<TuneStep> tune_steps)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.hook_up_checkbox_events()`** — L122 — `private void hook_up_checkbox_events()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkbox_checked_changed()`** — L130 — `private void checkbox_checked_changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCheckedCount()`** — L135 — `public int GetCheckedCount()`
  Returns checked count.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucTunestepOptionsGrid.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
