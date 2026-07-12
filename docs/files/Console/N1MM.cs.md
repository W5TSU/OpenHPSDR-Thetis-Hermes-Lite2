# `Console/N1MM.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Streams spectrum display data over UDP to the N1MM+ logger's spectrum window.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×10)
  - `Console/console.cs` (calls ×3)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.SetEnabled()` (×2), `.Stop()` (×2), `.Resize()` (×2), `.SetScale()` (×2), `.GetID()` (×1), `.SetID()` (×1), `.IsEnabled()` (×1), `.CopyData()` (×1)

## Outline

### Types

#### `N1MM` (type, L53)

- `.GetID()` — L79
- `.SetID()` — L84
- `.IsEnabled()` — L90
- `.SetEnabled()` — L96
- `.setLowFrequencyMHz()` — L104
- `.setHighFrequencyMHz()` — L110
- `.Stop()` — L127
- `.setMaxRXs()` — L138
- `.Resize()` — L152
- `.CopyData()` — L227
- `.Start()` — L252
- `.SetScale()` — L310
- `.sendUDPData()` — L316

#### `ReceiverStoredData` (type, L57)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/N1MM.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
