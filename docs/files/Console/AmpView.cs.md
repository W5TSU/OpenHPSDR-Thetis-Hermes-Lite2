# `Console/AmpView.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** PureSignal TX linearization control panel and the amplifier gain/phase view (backed by wdsp `calcc.c`/`iqc.c`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)
  - `Console/PSForm.cs` (references ×1)

## Outline

### Types

#### `AmpView` (type, L58)

- `.AmpView_Load()` — L84
- `.disp_setup()` — L107
- `.init_data()` — L123
- `.disp_data_Update()` — L156
- `.chkStayOnTop_CheckedChanged()` — L329
- `.CloseDown()` — L335
- `.timer1_Tick()` — L355
- `.chkAVShowGain_CheckedChanged()` — L435
- `.chkAVLowRes_CheckedChanged()` — L457
- `.AmpView_FormClosing()` — L465
- `.chkAVPhaseZoom_CheckedChanged()` — L470
- `.AmpView_FormClosed()` — L484
- `.SetWindowPos()` — L490
- `.FixOnTop()` — L501
- `.OnShown()` — L521

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/AmpView.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
