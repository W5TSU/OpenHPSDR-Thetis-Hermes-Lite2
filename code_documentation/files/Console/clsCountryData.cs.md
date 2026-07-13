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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CountryData` (type, L51)

- **`.GetCallsignData()`** — L101 — `public static PrefixData GetCallsignData(string callsign)`
  Returns callsign data.
  Called by: `.getFlagImageFromCallsign()` (`Console/SpotManager2.cs`), `.IsValidCallsign()` (`Console/clsDiscord.cs`)
- **`.LoadPrefixes()`** — L143 — `private static void LoadPrefixes(string filePath)`
  Loads prefixes.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getCountryCode()`** — L268 — `private static string getCountryCode(string country, int adif)`
  Returns country code.
  Called by: `.LoadPrefixes()` (same file)
- **`.getAssetCode()`** — L280 — `private static string getAssetCode(string country, int adif, string countryCode)`
  Returns asset code.
  Called by: `.LoadPrefixes()` (same file)
- **`.createAdifCountryCodeMap()`** — L291 — `private static Dictionary<int, string> createAdifCountryCodeMap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.createAdifAssetCodeMap()`** — L395 — `private static Dictionary<int, string> createAdifAssetCodeMap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.createCountryCodeAliasMap()`** — L406 — `private static Dictionary<string, string> createCountryCodeAliasMap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.createAssetCodeAliasMap()`** — L499 — `private static Dictionary<string, string> createAssetCodeAliasMap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.createRegionCountryCodeMap()`** — L532 — `private static Dictionary<string, string> createRegionCountryCodeMap()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addCountryCode()`** — L594 — `private static void addCountryCode(Dictionary<int, string> map, int adif, string code)`
  Called by: `.createAdifCountryCodeMap()` (same file), `.createAdifAssetCodeMap()` (same file)
- **`.addAlias()`** — L606 — `private static void addAlias(Dictionary<string, string> map, string name, string code)`
  Called by: `.createCountryCodeAliasMap()` (same file), `.createAssetCodeAliasMap()` (same file), `.createRegionCountryCodeMap()` (same file)
- **`.normalizeCountryName()`** — L621 — `private static string normalizeCountryName(string value)`
  Called by: `.getCountryCode()` (same file), `.getAssetCode()` (same file), `.addAlias()` (same file)

#### `PrefixData` (type, L62)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCountryData.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
