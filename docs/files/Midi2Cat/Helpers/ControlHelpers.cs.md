# `Midi2Cat/Helpers/ControlHelpers.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** UI helpers and MIDI diagnostics.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Midi2Cat.Helpers` (namespace, L30)

_No extracted members._

#### `ControlHelpers` (type, L32)

- **`.InvokeIfRequired()`** — L34 — `public static void InvokeIfRequired<T>(this T control, Action<T> action) where T : ISynchronizeInvoke`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Helpers/ControlHelpers.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
