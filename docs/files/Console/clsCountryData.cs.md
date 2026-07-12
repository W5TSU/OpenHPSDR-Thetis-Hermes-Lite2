# `Console/clsCountryData.cs`

**Functional area:** [16. DX spots and cluster display](../../CODE_OUTLINE.md#16-dx-spots-and-cluster-display)

**Role:** DXCC country/prefix lookup for spot flag/bearing data.

## How this file is used

- Used by (incoming references from other files):
  - `Console/SpotManager2.cs` (calls ×1)
  - `Console/clsDiscord.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetCallsignData()` (×2)

## Outline

### Types

#### `CountryData` (type, L51)

- `.GetCallsignData()` — L101
- `.LoadPrefixes()` — L143
- `.getCountryCode()` — L268
- `.getAssetCode()` — L280
- `.createAdifCountryCodeMap()` — L291
- `.createAdifAssetCodeMap()` — L395
- `.createCountryCodeAliasMap()` — L406
- `.createAssetCodeAliasMap()` — L499
- `.createRegionCountryCodeMap()` — L532
- `.addCountryCode()` — L594
- `.addAlias()` — L606
- `.normalizeCountryName()` — L621

#### `PrefixData` (type, L62)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCountryData.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
