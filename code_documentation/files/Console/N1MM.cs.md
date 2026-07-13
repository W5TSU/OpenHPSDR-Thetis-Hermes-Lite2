# `Console/N1MM.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Streams spectrum display data over UDP to the N1MM+ logger's spectrum window.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×10)
  - `Console/console.cs` (calls ×3)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.SetEnabled()` (×2), `.Stop()` (×2), `.Resize()` (×2), `.SetScale()` (×2), `.GetID()` (×1), `.SetID()` (×1), `.IsEnabled()` (×1), `.CopyData()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `N1MM` (type, L53)

- **`.GetID()`** — L79 — `public static string GetID(int rx)`
  Returns id.
  Called by: `.txtN1MM_RXn_ID_TextChanged()` (`Console/setup.cs`)
- **`.SetID()`** — L84 — `public static void SetID(int rx, string id)`
  Sets id.
  Called by: `.txtN1MM_RXn_ID_TextChanged()` (`Console/setup.cs`)
- **`.IsEnabled()`** — L90 — `public static bool IsEnabled(int rx)`
  Called by: `.Resize()` (same file), `.UpdateStatusBarStatusIcons()` (`Console/console.cs`)
- **`.SetEnabled()`** — L96 — `public static void SetEnabled(int rx, bool enable)`
  Sets enabled.
  Called by: `.chkN1MMEnableRX1_CheckedChanged()` (`Console/setup.cs`), `.chkN1MMEnableRX2_CheckedChanged()` (`Console/setup.cs`)
- **`.setLowFrequencyMHz()`** — L104 — `private static void setLowFrequencyMHz(int rx, double freq)`
  Sets low frequency mhz.
  Called by: `.Resize()` (same file)
- **`.setHighFrequencyMHz()`** — L110 — `private static void setHighFrequencyMHz(int rx, double freq)`
  Sets high frequency mhz.
  Called by: `.Resize()` (same file)
- **`.Stop()`** — L127 — `public static void Stop()`
  Called by: `.ExitConsole()` (`Console/console.cs`), `.stopStartN1MMSpectrum()` (`Console/setup.cs`)
- **`.setMaxRXs()`** — L138 — `private static void setMaxRXs(int rxNumber)`
  Sets max rxs.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Resize()`** — L152 — `public static void Resize()`
  public static void Init(Console c) { _console = c; }
  Called by: `.SetEnabled()` (same file), `.chkN1MMEnableRX1_CheckedChanged()` (`Console/setup.cs`), `.chkN1MMEnableRX2_CheckedChanged()` (`Console/setup.cs`)
- **`.CopyData()`** — L227 — `public static void CopyData(int rx, float[] newData)`
  Called by: `.RunDisplay()` (`Console/console.cs`)
- **`.Start()`** — L252 — `public static void Start()`
  Called by: `.stopStartN1MMSpectrum()` (`Console/setup.cs`)
- **`.SetScale()`** — L310 — `public static void SetScale(int rx, float fScale)`
  Sets scale.
  Called by: `.udN1MMRX1Scaling_ValueChanged()` (`Console/setup.cs`), `.udN1MMRX2Scaling_ValueChanged()` (`Console/setup.cs`)
- **`.sendUDPData()`** — L316 — `private static void sendUDPData()`
  Called by: `.Start()` (same file)

#### `ReceiverStoredData` (type, L57)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/N1MM.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
