# `Console/clsFlagAtlas.cs`

**Functional area:** [16. DX spots and cluster display](../../CODE_OUTLINE.md#16-dx-spots-and-cluster-display)

**Role:** Country flag atlas and web image fetching (e.g., QRZ pictures).

## How this file is used

- Used by (incoming references from other files):
  - `Console/SpotManager2.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetFlag()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `FlagAtlas` (type, L49)

- **`.Init()`** — L55 — `public static void Init(Image atlas_image, string json)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetFlag()`** — L104 — `public static Bitmap GetFlag(string flag_name)`
  Returns flag.
  Called by: `.getFlagImage()` (`Console/SpotManager2.cs`)
- **`.ContainsFlag()`** — L126 — `public static bool ContainsFlag(string flag_name)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetFlagBounds()`** — L140 — `public static Rectangle GetFlagBounds(string flag_name)`
  Returns flag bounds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Clear()`** — L157 — `public static void Clear()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ensureInitialised()`** — L171 — `private static void ensureInitialised()`
  Called by: `.GetFlag()` (same file), `.GetFlagBounds()` (same file)

#### `AtlasDefinition` (type, L177)

_No extracted members._

#### `SpriteDefinition` (type, L183)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsFlagAtlas.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
