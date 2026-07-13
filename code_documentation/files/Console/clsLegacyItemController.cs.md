# `Console/clsLegacyItemController.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Maps legacy/renamed UI items to their current equivalents so old databases and skins keep working.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
- Most-referenced symbols from other files: `.Update()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `LegacyItemController` (type, L49)

- **`.Init()`** — L90 — `public static void Init(Console c)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L309 — `public static void Update()`
  Called by: `.ExpandDisplay()` (`Console/console.cs`), `.CollapseDisplay()` (`Console/console.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsLegacyItemController.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
