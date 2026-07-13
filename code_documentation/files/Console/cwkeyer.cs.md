# `Console/cwkeyer.cs`

**Functional area:** [11. CW keying](../../CODE_OUTLINE.md#11-cw-keying)

**Role:** The CW keyer: iambic keying, key/paddle input sources (serial, radio, MIDI), break-in timing, and keying the radio.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/hiperftimer.cs` (calls ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `PowerSDR` (namespace, L36)

_No extracted members._

#### `CWKeyer2` (type, L43)

- **`.SetThreadAffinity()`** — L309 — `private bool SetThreadAffinity(int cpu)`
  Sets thread affinity.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.KeyThread()`** — L367 — `public void KeyThread()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/cwkeyer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
