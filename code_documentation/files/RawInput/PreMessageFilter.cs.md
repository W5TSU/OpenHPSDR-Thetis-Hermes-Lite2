# `RawInput/PreMessageFilter.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Win32 interop declarations and device-name registry lookup.

## How this file is used

- Used by (incoming references from other files):
  - `RawInput/RawInput.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `PreMessageFilter` (type, L6)

- **`.PreFilterMessage()`** — L17 — `public bool PreFilterMessage(ref Message m)`
  true to filter the message and stop it from being dispatched false to allow the message to continue to the next filter or control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/PreMessageFilter.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
