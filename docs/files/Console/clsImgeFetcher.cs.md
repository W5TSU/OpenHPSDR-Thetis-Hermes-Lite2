# `Console/clsImgeFetcher.cs`

**Functional area:** [16. DX spots and cluster display](../../CODE_OUTLINE.md#16-dx-spots-and-cluster-display)

**Role:** Country flag atlas and web image fetching (e.g., QRZ pictures).

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (references ×1, calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Shutdown()` (×1)

## Outline

### Types

#### `ImageFetcher` (type, L58)

- `.RegisterURL()` — L105
- `.LatestImages()` — L129
- `.UpdateInterval()` — L140
- `.UpdateBypassCache()` — L152
- `.clearAllImages()` — L164
- `.cleanupResources()` — L173
- `.StopFetching()` — L188
- `.Shutdown()` — L199
- `.fetch_images()` — L218
- `.ProcessMultipartContent()` — L458
- `.IsSvgImage()` — L556
- `.ProcessSvgImage()` — L563
- `.findBoundary()` — L590
- `.CheckForBoundary()` — L611
- `.getBoundary()` — L621
- `.ExtractImageUrls()` — L631
- `.OnImagesObtained()` — L655

#### `State` (type, L60)

_No extracted members._

#### `StateEventArgs` (type, L80)

_No extracted members._

#### `ImageStore` (type, L660)

- `.AddImage()` — L672
- `.GetImages()` — L688
- `.ClearImages()` — L696

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsImgeFetcher.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
