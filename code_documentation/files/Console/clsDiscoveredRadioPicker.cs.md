# `Console/clsDiscoveredRadioPicker.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** UI for picking among discovered radios and defining custom/static radio addresses.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/HPSDR/clsRadioDiscovery.cs` (references ×5)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `clsDiscoveredRadioPicker` (type, L54)

- **`.PickRadios()`** — L62 — `public List<NicRadioScanResult> PickRadios(IWin32Window owner, List<NicRadioScanResult> discovered)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.buildNicKey()`** — L566 — `private string buildNicKey(NicRadioScanResult nic)`
  Called by: `.PickRadios()` (same file)
- **`.cloneNicWithoutRadios()`** — L573 — `private NicRadioScanResult cloneNicWithoutRadios(NicRadioScanResult src)`
  Called by: `.PickRadios()` (same file)

#### `RowRef` (type, L56)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsDiscoveredRadioPicker.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
