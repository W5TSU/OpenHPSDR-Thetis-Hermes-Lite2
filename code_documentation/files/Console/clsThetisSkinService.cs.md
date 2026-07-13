# `Console/clsThetisSkinService.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** UI skin loading and application (SkiaSharp-backed image skins for console controls).

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×16, references ×4)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.CancelDownload()` (×4), `.GetThetisSkinsData()` (×1), `.GetSkinServers()` (×1), `.SubscribeForSkinData()` (×1), `.UnsubscribeFromSkinData()` (×1), `.SubscribeForSkinServerData()` (×1), `.UnsubscribeFromSkinServerData()` (×1), `.SubscribeForImageLoaded()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

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

- **`.GetThetisSkinsData()`** — L140 — `public static async void GetThetisSkinsData(string jsonUrl)`
  Returns thetis skins data.
  Called by: `.btnRefreshSkinsForServer_Click()` (`Console/setup.cs`)
- **`.GetSkinServers()`** — L200 — `public static async void GetSkinServers(string jsonUrl)`
  Returns skin servers.
  Called by: `.getSkinServers()` (`Console/setup.cs`)
- **`.SubscribeForSkinData()`** — L247 — `public static void SubscribeForSkinData(EventHandler<SkinsData> eventHandler)`
  Called by: `.addDelegates()` (`Console/setup.cs`)
- **`.UnsubscribeFromSkinData()`** — L252 — `public static void UnsubscribeFromSkinData(EventHandler<SkinsData> eventHandler)`
  Called by: `.RemoveDelegates()` (`Console/setup.cs`)
- **`.SubscribeForSkinServerData()`** — L256 — `public static void SubscribeForSkinServerData(EventHandler<SkinServersData> eventHandler)`
  Called by: `.addDelegates()` (`Console/setup.cs`)
- **`.UnsubscribeFromSkinServerData()`** — L261 — `public static void UnsubscribeFromSkinServerData(EventHandler<SkinServersData> eventHandler)`
  Called by: `.RemoveDelegates()` (`Console/setup.cs`)
- **`.SubscribeForImageLoaded()`** — L265 — `public static void SubscribeForImageLoaded(EventHandler<SkinHttpImage> eventHandler)`
  Called by: `.addDelegates()` (`Console/setup.cs`)
- **`.UnsubscribeFromImageLoaded()`** — L270 — `public static void UnsubscribeFromImageLoaded(EventHandler<SkinHttpImage> eventHandler)`
  Called by: `.RemoveDelegates()` (`Console/setup.cs`)
- **`.SubscribeForDownload()`** — L274 — `public static void SubscribeForDownload(EventHandler<SkinFileDownload> eventHandler)`
  Called by: `.addDelegates()` (`Console/setup.cs`)
- **`.UnsubscribeFromDownload()`** — L279 — `public static void UnsubscribeFromDownload(EventHandler<SkinFileDownload> eventHandler)`
  Called by: `.RemoveDelegates()` (`Console/setup.cs`)
- **`.LoadImageFromUrl()`** — L283 — `public static async void LoadImageFromUrl(string imageUrl, string sID)`
  Loads image from url.
  Called by: `.updateSelectedSkin()` (`Console/setup.cs`)
- **`.DownloadFile()`** — L337 — `public static async void DownloadFile(string fileUrl, string savePath, bool bypassFolderCheck, bool isMeterSkin)`
  Called by: `.downloadSkin()` (`Console/setup.cs`)
- **`.CancelDownload()`** — L427 — `public static void CancelDownload()`
  Called by: `.Hide()` (`Console/setup.cs`), `.tcAppearance_SelectedIndexChanged()` (`Console/setup.cs`), `.downloadSkin()` (`Console/setup.cs`), `.lstAvailableSkins_SelectedIndexChanged()` (`Console/setup.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsThetisSkinService.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
