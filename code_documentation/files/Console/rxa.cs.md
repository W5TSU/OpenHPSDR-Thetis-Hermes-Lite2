# `Console/rxa.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Typed wrappers for wdsp RXA (receiver chain) settings and the UI controls bound to them.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/PanDisplay.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `rxa` (type, L15)

- **`.create_rxa()`** — L50 — `private void create_rxa()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ForceRxa()`** — L67 — `private void ForceRxa()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.udRXAFreq_ValueChanged()`** — L86 — `private void udRXAFreq_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXAFreq` value changes.
  Called by: `.ForceRxa()` (same file)
- **`.udRXAAGCGain_ValueChanged()`** — L92 — `private void udRXAAGCGain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXAAGCGain` value changes.
  Called by: `.ForceRxa()` (same file)
- **`.udRXAVolume_ValueChanged()`** — L97 — `private void udRXAVolume_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXAVolume` value changes.
  Called by: `.ForceRxa()` (same file)
- **`.udRXAMode_ValueChanged()`** — L102 — `private void udRXAMode_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udRXAMode` value changes.
  Called by: `.ForceRxa()` (same file)
- **`.rxa_FormClosing()`** — L130 — `private void rxa_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `rxa` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.panDisplay_Resize()`** — L143 — `private void panDisplay_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `panDisplay` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rxa_Resize()`** — L167 — `private void rxa_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `rxa` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/rxa.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
