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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `NetworkIO` (type, L10)

- **`.InitRadio()`** — L26 — `public static int InitRadio()`
  Inits radio.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOutputPower()`** — L199 — `public static void SetOutputPower(float f)`
  Sets output power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VFOfreq()`** — L215 — `unsafe public static void VFOfreq(int id, double f, int tx)`
  Called by: `.FreqCorrectionChanged()` (same file)
- **`.FreqCorrectionChanged()`** — L237 — `public static void FreqCorrectionChanged()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Freq2PhaseWord()`** — L249 — `public static int Freq2PhaseWord(int freq)`
  Called by: `.VFOfreq()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/HPSDR/NetworkIO.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
