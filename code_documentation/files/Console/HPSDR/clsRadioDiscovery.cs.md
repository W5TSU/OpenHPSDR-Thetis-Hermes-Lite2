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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

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

- **`.applyScanPerformanceProfile()`** — L465 — `private void applyScanPerformanceProfile(ScanPerformanceProfile profile)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `RadioInfo` (type, L514)

_No extracted members._

#### `DiscoveryDiagnostics` (type, L544)

_No extracted members._

#### `NicRadioScanResult` (type, L564)

- **`.ToString()`** — L620 — `public override string ToString()`
  Returns the string representation.
  Called by: `.discoverOnNic()` (same file)

#### `RadioDiscoveryService` (type, L626)

- **`.DiscoverUsingAllNics()`** — L638 — `public List<NicRadioScanResult> DiscoverUsingAllNics(RadioDiscoveryOptions options)`
  Called by: `.tryDiscoverRadios()` (`Console/setup.cs`)
- **`.DiscoverUsingSingleNic()`** — L671 — `public NicRadioScanResult DiscoverUsingSingleNic(RadioDiscoveryOptions options, IPAddress localIPv4)`
  Called by: `.InitRadio()` (`Console/HPSDR/NetworkIO.cs`), `.tryDiscoverRadios()` (`Console/setup.cs`)
- **`.ListUsableNics()`** — L710 — `public List<NicRadioScanResult> ListUsableNics(RadioDiscoveryOptions options)`
  Called by: `.rebuildNicCombo()` (`Console/setup.cs`)
- **`.createNicResultSkeleton()`** — L738 — `private NicRadioScanResult createNicResultSkeleton(NicIpv4Binding b)`
  Called by: `.DiscoverUsingAllNics()` (same file), `.DiscoverUsingSingleNic()` (same file), `.ListUsableNics()` (same file)
- **`.hydrateNicNetworkProps()`** — L757 — `private void hydrateNicNetworkProps(NicIpv4Binding b, NicRadioScanResult nicResult)`
  Called by: `.DiscoverUsingAllNics()` (same file), `.DiscoverUsingSingleNic()` (same file), `.ListUsableNics()` (same file)
- **`.sanitizeLoopbackNicFields()`** — L803 — `private void sanitizeLoopbackNicFields(NicRadioScanResult nicResult)`
  Called by: `.DiscoverUsingAllNics()` (same file), `.DiscoverUsingSingleNic()` (same file), `.ListUsableNics()` (same file)
- **`.formatNicMac()`** — L829 — `private string formatNicMac(PhysicalAddress pa)`
  Called by: `.createNicResultSkeleton()` (same file)
- **`.enumerateNicIpv4Bindings()`** — L845 — `private List<NicIpv4Binding> enumerateNicIpv4Bindings(RadioDiscoveryOptions options)`
  Called by: `.DiscoverUsingAllNics()` (same file), `.DiscoverUsingSingleNic()` (same file), `.ListUsableNics()` (same file)
- **`.discoverOnNic()`** — L919 — `private List<RadioInfo> discoverOnNic(IPAddress localIPv4, IPAddress localMaskIPv4, RadioDiscoveryOptions options, out DiscoveryDiagnostics diagnostics)`
  Called by: `.DiscoverUsingAllNics()` (same file), `.DiscoverUsingSingleNic()` (same file)
- **`.parseDiscoveryReply()`** — L1137 — `private DiscoveryParseResult parseDiscoveryReply(byte[] data, int len, IPAddress senderIp, RadioDiscoveryOptions options)`
  Called by: `.discoverOnNic()` (same file)
- **`.mapP1DeviceType()`** — L1236 — `private HPSDRHW mapP1DeviceType(byte boardId)`
  Called by: `.parseDiscoveryReply()` (same file)
- **`.buildTargets()`** — L1296 — `private List<IPEndPoint> buildTargets(IPAddress localIPv4, IPAddress mask, RadioDiscoveryOptions options)`
  Called by: `.discoverOnNic()` (same file)
- **`.buildDiscoveryPacketP1()`** — L1335 — `private byte[] buildDiscoveryPacketP1()`
  Called by: `.discoverOnNic()` (same file)
- **`.buildDiscoveryPacketP2()`** — L1345 — `private byte[] buildDiscoveryPacketP2()`
  Called by: `.discoverOnNic()` (same file)
- **`.isNicCandidate()`** — L1353 — `private bool isNicCandidate(NetworkInterface nic, RadioDiscoveryOptions options)`
  Called by: `.enumerateNicIpv4Bindings()` (same file)
- **`.sameSubnet()`** — L1390 — `private bool sameSubnet(IPAddress a, IPAddress b, IPAddress mask)`
  Called by: `.discoverOnNic()` (same file)
- **`.isApipa()`** — L1419 — `private bool isApipa(IPAddress ip)`
  Called by: `.createNicResultSkeleton()` (same file), `.enumerateNicIpv4Bindings()` (same file), `.discoverOnNic()` (same file)
- **`.getBroadcastAddress()`** — L1435 — `private IPAddress getBroadcastAddress(IPAddress address, IPAddress subnetMask)`
  Returns broadcast address.
  Called by: `.buildTargets()` (same file)

#### `NicIpv4Binding` (type, L631)

_No extracted members._

#### `DiscoveryParseResult` (type, L1115)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/clsRadioDiscovery.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
