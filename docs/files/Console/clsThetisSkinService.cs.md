# `Console/clsThetisSkinService.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** UI skin loading and application (SkiaSharp-backed image skins for console controls).

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×16, references ×4)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.CancelDownload()` (×4), `.GetThetisSkinsData()` (×1), `.GetSkinServers()` (×1), `.SubscribeForSkinData()` (×1), `.UnsubscribeFromSkinData()` (×1), `.SubscribeForSkinServerData()` (×1), `.UnsubscribeFromSkinServerData()` (×1), `.SubscribeForImageLoaded()` (×1)

## Outline

### Types

#### `ThetisSkin` (type, L54)

_No extracted members._

#### `SkinsData` (type, L74)

_No extracted members._

#### `SkinServer` (type, L84)

_No extracted members._

#### `SkinServersData` (type, L102)

_No extracted members._

#### `SkinHttpImage` (type, L107)

_No extracted members._

#### `SkinFileDownload` (type, L112)

_No extracted members._

#### `ThetisSkinService` (type, L125)

- `.GetThetisSkinsData()` — L140
- `.GetSkinServers()` — L200
- `.SubscribeForSkinData()` — L247
- `.UnsubscribeFromSkinData()` — L252
- `.SubscribeForSkinServerData()` — L256
- `.UnsubscribeFromSkinServerData()` — L261
- `.SubscribeForImageLoaded()` — L265
- `.UnsubscribeFromImageLoaded()` — L270
- `.SubscribeForDownload()` — L274
- `.UnsubscribeFromDownload()` — L279
- `.LoadImageFromUrl()` — L283
- `.DownloadFile()` — L337
- `.CancelDownload()` — L427

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsThetisSkinService.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
