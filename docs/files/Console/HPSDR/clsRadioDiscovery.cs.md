# `Console/HPSDR/clsRadioDiscovery.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** UDP broadcast discovery of HPSDR radios on all NICs; produces the discovered-radio list.

## How this file is used

- Used by (incoming references from other files):
  - `Console/ucRadioList.cs` (references ×12)
  - `Console/setup.cs` (references ×4, calls ×3)
  - `Console/clsDiscoveredRadioPicker.cs` (references ×5)
  - `Console/HPSDR/NetworkIO.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×3)
- Most-referenced symbols from other files: `.DiscoverUsingSingleNic()` (×2), `.DiscoverUsingAllNics()` (×1), `.ListUsableNics()` (×1)

## Outline

### Types

#### `IfTypeSyntax` (type, L80)

_No extracted members._

#### `RadioDiscoveryProtocolMode` (type, L383)

_No extracted members._

#### `RadioDiscoveryRadioProtocol` (type, L390)

_No extracted members._

#### `ScanPerformanceProfile` (type, L397)

_No extracted members._

#### `RadioDiscoveryOptions` (type, L407)

- `.applyScanPerformanceProfile()` — L465

#### `RadioInfo` (type, L514)

_No extracted members._

#### `DiscoveryDiagnostics` (type, L544)

_No extracted members._

#### `NicRadioScanResult` (type, L564)

- `.ToString()` — L620

#### `RadioDiscoveryService` (type, L626)

- `.DiscoverUsingAllNics()` — L638
- `.DiscoverUsingSingleNic()` — L671
- `.ListUsableNics()` — L710
- `.createNicResultSkeleton()` — L738
- `.hydrateNicNetworkProps()` — L757
- `.sanitizeLoopbackNicFields()` — L803
- `.formatNicMac()` — L829
- `.enumerateNicIpv4Bindings()` — L845
- `.discoverOnNic()` — L919
- `.parseDiscoveryReply()` — L1137
- `.mapP1DeviceType()` — L1236
- `.buildTargets()` — L1296
- `.buildDiscoveryPacketP1()` — L1335
- `.buildDiscoveryPacketP2()` — L1345
- `.isNicCandidate()` — L1353
- `.sameSubnet()` — L1390
- `.isApipa()` — L1419
- `.getBroadcastAddress()` — L1435

#### `NicIpv4Binding` (type, L631)

_No extracted members._

#### `DiscoveryParseResult` (type, L1115)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/clsRadioDiscovery.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
