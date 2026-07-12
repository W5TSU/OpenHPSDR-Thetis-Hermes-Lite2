# `Console/HPSDR/NetworkIO.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** High-level radio session control: init/start/stop, VFO frequency-to-phase-word conversion, sample rate, and control-register updates to the radio.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/HPSDR/NetworkIOImports.cs` (calls ×3)
  - `Console/enums.cs` (references ×2)
  - `Console/HPSDR/clsRadioDiscovery.cs` (calls ×1)

## Outline

### Types

#### `NetworkIO` (type, L10)

- `.InitRadio()` — L26
- `.SetOutputPower()` — L199
- `.VFOfreq()` — L215
- `.FreqCorrectionChanged()` — L237
- `.Freq2PhaseWord()` — L249

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/NetworkIO.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
