# `Console/NetworkThrottle.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** Network send-rate throttling to smooth UDP bursts.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×3)
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.GetNetworkThrottle()` (×2), `.SetNetworkThrottle()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `NetworkThrottle` (type, L47)

- **`.GetNetworkThrottle()`** — L49 — `public static bool GetNetworkThrottle(out int throttle, bool showErrors = true)`
  Returns network throttle.
  Called by: `.chkNetworkThrottleIndexTweak_CheckedChanged()` (`Console/setup.cs`), `.updateNetworkThrottleCheckBox()` (`Console/setup.cs`)
- **`.SetNetworkThrottle()`** — L127 — `public static bool SetNetworkThrottle(int throttle)`
  Sets network throttle.
  Called by: `.chkNetworkThrottleIndexTweak_CheckedChanged()` (`Console/setup.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/NetworkThrottle.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
