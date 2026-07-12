# `Console/clsFlagAtlas.cs`

**Functional area:** [16. DX spots and cluster display](../../CODE_OUTLINE.md#16-dx-spots-and-cluster-display)

**Role:** Country flag atlas and web image fetching (e.g., QRZ pictures).

## How this file is used

- Used by (incoming references from other files):
  - `Console/SpotManager2.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetFlag()` (×1)

## Outline

### Types

#### `FlagAtlas` (type, L49)

- `.Init()` — L55
- `.GetFlag()` — L104
- `.ContainsFlag()` — L126
- `.GetFlagBounds()` — L140
- `.Clear()` — L157
- `.ensureInitialised()` — L171

#### `AtlasDefinition` (type, L177)

_No extracted members._

#### `SpriteDefinition` (type, L183)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsFlagAtlas.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
