# `Console/rxaControls.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Typed wrappers for wdsp RXA (receiver chain) settings and the UI controls bound to them.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `rxaControls` (type, L13)

- **`.comboAudioSampleRate_SelectedIndexChanged()`** — L31 — `private void comboAudioSampleRate_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboAudioSampleRate` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rxaControls_Paint()`** — L36 — `private void rxaControls_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `rxaControls` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rxaControls_FormClosing()`** — L41 — `private void rxaControls_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `rxaControls` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rxaControls_FormClosed()`** — L46 — `private void rxaControls_FormClosed(object sender, FormClosedEventArgs e)`
  WinForms event handler: runs when `rxaControls` has closed.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rxaControls_Activated()`** — L51 — `private void rxaControls_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.rxaControls_Deactivate()`** — L57 — `private void rxaControls_Deactivate(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.rxaControls_MdiChildActivate()`** — L63 — `private void rxaControls_MdiChildActivate(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/rxaControls.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
