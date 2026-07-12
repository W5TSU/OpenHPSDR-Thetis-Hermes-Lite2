# `Console/clsImgeFetcher.cs`

**Functional area:** [16. DX spots and cluster display](../../CODE_OUTLINE.md#16-dx-spots-and-cluster-display)

**Role:** Country flag atlas and web image fetching (e.g., QRZ pictures).

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (references ×1, calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Shutdown()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ImageFetcher` (type, L58)

- **`.RegisterURL()`** — L105 — `public Guid RegisterURL(string url, int timeout_secs, int image_limit, bool file, bool bypass_cache = false)`
  Registers url.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LatestImages()`** — L129 — `public List<Image> LatestImages(Guid id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateInterval()`** — L140 — `public void UpdateInterval(Guid id, int interval)`
  Updates interval.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateBypassCache()`** — L152 — `public void UpdateBypassCache(Guid id, bool bypass)`
  Updates bypass cache.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clearAllImages()`** — L164 — `private void clearAllImages()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.cleanupResources()`** — L173 — `private void cleanupResources(Guid id)`
  Called by: `.StopFetching()` (same file), `.fetch_images()` (same file)
- **`.StopFetching()`** — L188 — `public void StopFetching(Guid id)`
  Stops fetching.
  Called by: `.RegisterURL()` (same file), `.Shutdown()` (same file)
- **`.Shutdown()`** — L199 — `public void Shutdown()`
  Called by: `.Shutdown()` (`Console/MeterManager.cs`)
- **`.fetch_images()`** — L218 — `private void fetch_images(string url, ImageStore store, ManualResetEvent reset_event, Guid id, bool file)`
  Called by: `.RegisterURL()` (same file)
- **`.ProcessMultipartContent()`** — L458 — `private void ProcessMultipartContent(Stream stream, string boundary, ImageStore store, ref bool imagesAdded, ManualResetEvent reset_event, Guid id)`
  Processes multipart content.
  Called by: `.fetch_images()` (same file)
- **`.IsSvgImage()`** — L556 — `private bool IsSvgImage(byte[] imageData)`
  Method to check if the image data is SVG
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ProcessSvgImage()`** — L563 — `private void ProcessSvgImage(byte[] svgData, ImageStore store, ref bool imagesAdded, Guid id)`
  Method to process SVG images
  Called by: `.fetch_images()` (same file)
- **`.findBoundary()`** — L590 — `private int findBoundary(byte[] content, int start, byte[] boundary)`
  Called by: `.ProcessMultipartContent()` (same file)
- **`.CheckForBoundary()`** — L611 — `private bool CheckForBoundary(MemoryStream ms, string boundary)`
  Checks for boundary.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getBoundary()`** — L621 — `private string getBoundary(string contentType)`
  Returns boundary.
  Called by: `.fetch_images()` (same file)
- **`.ExtractImageUrls()`** — L631 — `private List<string> ExtractImageUrls(string html, string baseUrl)`
  Called by: `.fetch_images()` (same file)
- **`.OnImagesObtained()`** — L655 — `protected virtual void OnImagesObtained(Guid id)`
  Handles/raises the images obtained event.
  Called by: `.fetch_images()` (same file)

#### `State` (type, L60)

_No extracted members._

#### `StateEventArgs` (type, L80)

_No extracted members._

#### `ImageStore` (type, L660)

- **`.AddImage()`** — L672 — `public bool AddImage(Image image)`
  Adds image.
  Called by: `.fetch_images()` (same file), `.ProcessMultipartContent()` (same file), `.ProcessSvgImage()` (same file)
- **`.GetImages()`** — L688 — `public List<Image> GetImages()`
  Returns images.
  Called by: `.LatestImages()` (same file)
- **`.ClearImages()`** — L696 — `public void ClearImages()`
  Clears images.
  Called by: `.clearAllImages()` (same file), `.cleanupResources()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsImgeFetcher.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
