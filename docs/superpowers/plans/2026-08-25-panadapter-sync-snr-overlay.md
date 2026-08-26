# Panadapter Sync/SNR Overlay Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Show the active Digital Voice mode's live SYNC/SNR status directly
on Thetis's panadapter (not just in Setup), for RX1 and RX2 independently.

**Architecture:** A new static `Display.ShowRadeSyncOverlay` toggle (mirroring
the existing `ShowBandStackOverlays` pattern) gates a small new block inside
`drawPanadapterAndWaterfallGridDX2D` — the shared helper that already draws
the RX filter-passband overlay for both the panadapter pass
(`bIsWaterfall=false`) and the waterfall pass (`bIsWaterfall=true`). The new
block sits inside that helper's existing `!bIsWaterfall`- and
`!local_mox`-gated branch, reusing `filter_left_x`/`top`/`nVerticalShift`
already computed there, so it draws only on the panadapter, never on the
waterfall or during TX, with no new parameters threaded through the
function. A new checkbox in the Digital Voice Setup tab
(`chkShowRadeSyncOverlay`) drives the toggle.

**Tech Stack:** C# / WinForms, SharpDX Direct2D (this file's existing
rendering stack — no SkiaSharp, no new native code), existing P/Invoke
exports into `wdsp.dll`/`ChannelMaster.dll`.

**Spec:** [docs/superpowers/specs/2026-08-25-panadapter-sync-snr-overlay-design.md](../specs/2026-08-25-panadapter-sync-snr-overlay-design.md)

## Global Constraints

- **CRLF-only files, plain-text Edit tool risk**: `display.cs`, `setup.cs`,
  and `setup.designer.cs` are all confirmed 100% CRLF line endings (verified
  by direct byte count this session: 0 LF-only lines in any of the three).
  This project's prior sub-projects (#1-#3) repeatedly hit an Edit-tool
  defect where writing into one of these files flattens its line endings
  wholesale, producing a spurious multi-thousand-line diff. **Every edit to
  these three files in this plan must be done via Python byte-level
  splicing** (`data = open(path,'rb').read(); assert data.count(old_bytes)
  == 1; data = data.replace(old_bytes, new_bytes); open(path,'wb').write(data)`),
  never the plain Edit tool. Verify with `git diff --stat` after every edit
  — a diff far larger than the intended change means line endings were
  flattened; `git checkout -- <file>` and redo via splicing.
- **`initializing` guard**: every new/modified event handler that can fire
  during Setup's own startup sequence must start with
  `if (initializing) return;` — a defect class this project's final reviews
  caught twice before (sub-projects #2 and #3).
- **`WDSP.id()` doubled-thread vs. plain `radae.c` rx-index**: 700E's
  `WDSP.GetRXAFDVSync`/`GetRXAFDVSnr` take a `WDSP.id(thread, subrx)`
  channel index where RX2 is `WDSP.id(2, 0)`, NOT `WDSP.id(1, 0)` (that
  resolves to the TX channel — `dsp.cs:1161`'s `2*thread+subrx` mapping).
  RADE V1/V2's `WDSP.GetRadaeSync`/`GetRadaeSnrDb` take a plain 0/1 index
  instead — RX2 is plain `1`, never doubled. Getting this wrong caused two
  Critical bugs in sub-project #3. Task 1 below copies the already-correct,
  already-live-verified calls from `setup.cs` verbatim; do not re-derive
  them.
- **Default state**: the new toggle defaults to **checked/on** — this is
  the opposite default from the `ShowBandStackOverlays` precedent it
  otherwise mirrors (which defaults off). Do not copy that default.
- No new native exports, no new P/Invoke declarations, no new timer, no
  CAT/TCI changes — every backend call this plan uses already exists.

---

### Task 1: `Display.ShowRadeSyncOverlay` flag + panadapter draw code

**Files:**
- Modify: `Project Files/Source/Console/display.cs:1000` (new property,
  right after `ShowBandStackOverlays`)
- Modify: `Project Files/Source/Console/display.cs:9066` (new draw block,
  inside `drawPanadapterAndWaterfallGridDX2D`)

**Interfaces:**
- Consumes: `console.radio.GetDSPRX(int thread, int subrx)` returning
  `RadioDSPRX` with `int RXRadaeEnabled`/`int RXAFDVRun` fields (both
  already exist, `radio.cs`); `WDSP.id(uint thread, uint subrx)` returning
  `int` (`dsp.cs:1161`); `WDSP.GetRXAFDVSync(int channel)` returning `int`,
  `WDSP.GetRXAFDVSnr(int channel)` returning `double`, `WDSP.GetRadaeSync(int rx)`
  returning `int`, `WDSP.GetRadaeSnrDb(int rx)` returning `int` (all
  `dsp.cs`, all already exist); `getDXBrushForColour(Color c, int
  replaceAlpha = -1)` returning `SharpDX.Direct2D1.Brush`
  (`display.cs:11570`, already exists); `measureStringDX2D(string s,
  SharpDX.DirectWrite.TextFormat tf, bool cacheStringLength = false)`
  returning `System.Drawing.SizeF` (`display.cs:8542`, already exists);
  `drawStringDX2D(string s, SharpDX.DirectWrite.TextFormat tf,
  SharpDX.Direct2D1.Brush b, float x, float y)` (`display.cs:8504`,
  already exists); the existing static field `m_bDX2_grid_text_brush`
  (`display.cs:7865`) and existing static font `fontDX2d_font9`.
- Produces: `public static bool Display.ShowRadeSyncOverlay { get; set; }`
  — Task 2's checkbox handler sets this; Task 2's `InitRadePanelFromBackend`
  reads it.

- [ ] **Step 1: Add the new static property**

Read `display.cs` first (required before any edit). Then, using Python
byte-splicing, insert immediately after the existing `ShowBandStackOverlays`
property (locate this exact text — it is unique in the file):

Old bytes:
```
        private static bool m_bShowBandStackOverlays = false;
        public static bool ShowBandStackOverlays
        {
            get { return m_bShowBandStackOverlays; }
            set { m_bShowBandStackOverlays = value; }
        }
```

New bytes (old bytes + this appended immediately after, preserving the
exact CRLF line endings of the surrounding file — write the replacement
using `\r\n` line endings, matching every other line in this file):

```
        private static bool m_bShowBandStackOverlays = false;
        public static bool ShowBandStackOverlays
        {
            get { return m_bShowBandStackOverlays; }
            set { m_bShowBandStackOverlays = value; }
        }

        // W5TSU: panadapter sync/SNR overlay -- default ON (opposite of
        // ShowBandStackOverlays's own default), see
        // drawPanadapterAndWaterfallGridDX2D for the draw code.
        private static bool m_bShowRadeSyncOverlay = true;
        public static bool ShowRadeSyncOverlay
        {
            get { return m_bShowRadeSyncOverlay; }
            set { m_bShowRadeSyncOverlay = value; }
        }
```

Python splice script (adjust only the `path` if your checkout root
differs):

```python
path = "Project Files/Source/Console/display.cs"
data = open(path, "rb").read()
old = (
    b"        private static bool m_bShowBandStackOverlays = false;\r\n"
    b"        public static bool ShowBandStackOverlays\r\n"
    b"        {\r\n"
    b"            get { return m_bShowBandStackOverlays; }\r\n"
    b"            set { m_bShowBandStackOverlays = value; }\r\n"
    b"        }\r\n"
)
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"\r\n"
    b"        // W5TSU: panadapter sync/SNR overlay -- default ON (opposite of\r\n"
    b"        // ShowBandStackOverlays's own default), see\r\n"
    b"        // drawPanadapterAndWaterfallGridDX2D for the draw code.\r\n"
    b"        private static bool m_bShowRadeSyncOverlay = true;\r\n"
    b"        public static bool ShowRadeSyncOverlay\r\n"
    b"        {\r\n"
    b"            get { return m_bShowRadeSyncOverlay; }\r\n"
    b"            set { m_bShowRadeSyncOverlay = value; }\r\n"
    b"        }\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

- [ ] **Step 2: Run it, then verify with git diff**

Run the script above (e.g. `python3 splice1.py` after saving it to a
scratch file). Then run `git diff --stat -- "Project Files/Source/Console/display.cs"`
— expected output is a diff of roughly +9/-0 lines. If it instead reports
thousands of lines changed, line endings were flattened: run
`git checkout -- "Project Files/Source/Console/display.cs"` and redo the
splice, double-checking the byte string above matches exactly (including
the trailing `\r\n` on each line).

- [ ] **Step 3: Add the draw code**

Read `display.cs` again (it changed in Step 1). Using Python byte-splicing,
locate this exact block (it is unique in the file — it's the RX filter
edge highlight `switch` inside `drawPanadapterAndWaterfallGridDX2D`):

Old bytes:
```
                    if (!bIsWaterfall)
                    {
                        int nFilterEdge = 0;

                        if (rx == 1)
                            nFilterEdge = m_nHightlightFilterEdgeRX1;
                        else if (rx == 2)
                            nFilterEdge = m_nHightlightFilterEdgeRX2;

                        switch (nFilterEdge)
                        {
                            case -1:
                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_left_x, nVerticalShift + top, filter_left_x, nVerticalShift + H, 2);
                                break;
                            case 1:
                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_right_x, nVerticalShift + top, filter_right_x, nVerticalShift + H, 2);
                                break;
                            default:
                                break;
                        }
                    }
                }
            }
```

New bytes (old bytes with a new sibling `if (!bIsWaterfall) { ... }` block
inserted right after the existing one, still inside the same enclosing
`if (!local_mox)` — note only ONE trailing `}` is removed from the old
block's end and re-added after the new block, since we're inserting a new
statement as a sibling, not nesting inside the old one):

```
                    if (!bIsWaterfall)
                    {
                        int nFilterEdge = 0;

                        if (rx == 1)
                            nFilterEdge = m_nHightlightFilterEdgeRX1;
                        else if (rx == 2)
                            nFilterEdge = m_nHightlightFilterEdgeRX2;

                        switch (nFilterEdge)
                        {
                            case -1:
                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_left_x, nVerticalShift + top, filter_left_x, nVerticalShift + H, 2);
                                break;
                            case 1:
                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_right_x, nVerticalShift + top, filter_right_x, nVerticalShift + H, 2);
                                break;
                            default:
                                break;
                        }
                    }

                    // W5TSU: panadapter sync/SNR overlay (sub-project 5 of 6)
                    // -- draws only here: not on the waterfall (!bIsWaterfall,
                    // this block), not during TX (enclosing !local_mox block).
                    if (!bIsWaterfall && m_bShowRadeSyncOverlay)
                    {
                        int thread = rx - 1; // GetDSPRX's plain thread index: 0 = RX1, 1 = RX2
                        RadioDSPRX dsp = console.radio.GetDSPRX(thread, 0);
                        bool radeOn = dsp.RXRadaeEnabled != 0;
                        bool fdv700eOn = dsp.RXAFDVRun != 0;

                        if (radeOn || fdv700eOn)
                        {
                            bool sync;
                            string snrText;

                            if (fdv700eOn)
                            {
                                // WDSP.id() folds thread+subrx into one wdsp.dll channel
                                // index (channel = 2*thread + subrx, dsp.cs:1161) -- RX2
                                // is WDSP.id(2, 0), NOT WDSP.id(1, 0) (that resolves to
                                // the TX channel). Verbatim from radeRX2StatusTimer_Tick /
                                // radeStatusTimer_Tick, setup.cs:36947-36948 / 36815-36816.
                                int channel = WDSP.id((uint)(thread == 0 ? 0 : 2), 0);
                                sync = WDSP.GetRXAFDVSync(channel) != 0;
                                snrText = sync ? string.Format("{0:F1}", WDSP.GetRXAFDVSnr(channel)) : "";
                            }
                            else // radeOn
                            {
                                // RADE V1/V2: ChannelMaster's radae.c uses its own plain
                                // 0/1 rx index, unrelated to WDSP.id()'s doubled
                                // convention. Verbatim from radeStatusTimer_Tick,
                                // setup.cs:36820-36821 (rx=0) / radeRX2StatusTimer_Tick,
                                // setup.cs:36952-36953 (rx=1).
                                sync = WDSP.GetRadaeSync(thread) != 0;
                                snrText = sync ? string.Format("{0}", WDSP.GetRadaeSnrDb(thread)) : "";
                            }

                            string overlayText = sync ? string.Format("SYNC  SNR {0} dB", snrText) : "no sync";
                            SharpDX.Direct2D1.Brush overlayBrush = sync ? getDXBrushForColour(Color.Green) : m_bDX2_grid_text_brush;

                            SizeF overlaySize = measureStringDX2D(overlayText, fontDX2d_font9);
                            drawStringDX2D(overlayText, fontDX2d_font9, overlayBrush, filter_left_x, nVerticalShift + top - overlaySize.Height);
                        }
                    }
                }
            }
```

Python splice script:

```python
path = "Project Files/Source/Console/display.cs"
data = open(path, "rb").read()
old = (
    b"                    if (!bIsWaterfall)\r\n"
    b"                    {\r\n"
    b"                        int nFilterEdge = 0;\r\n"
    b"\r\n"
    b"                        if (rx == 1)\r\n"
    b"                            nFilterEdge = m_nHightlightFilterEdgeRX1;\r\n"
    b"                        else if (rx == 2)\r\n"
    b"                            nFilterEdge = m_nHightlightFilterEdgeRX2;\r\n"
    b"\r\n"
    b"                        switch (nFilterEdge)\r\n"
    b"                        {\r\n"
    b"                            case -1:\r\n"
    b"                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_left_x, nVerticalShift + top, filter_left_x, nVerticalShift + H, 2);\r\n"
    b"                                break;\r\n"
    b"                            case 1:\r\n"
    b"                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_right_x, nVerticalShift + top, filter_right_x, nVerticalShift + H, 2);\r\n"
    b"                                break;\r\n"
    b"                            default:\r\n"
    b"                                break;\r\n"
    b"                        }\r\n"
    b"                    }\r\n"
    b"                }\r\n"
    b"            }\r\n"
)
assert data.count(old) == 1, "anchor not found or not unique"

new_block = (
    b"                    if (!bIsWaterfall)\r\n"
    b"                    {\r\n"
    b"                        int nFilterEdge = 0;\r\n"
    b"\r\n"
    b"                        if (rx == 1)\r\n"
    b"                            nFilterEdge = m_nHightlightFilterEdgeRX1;\r\n"
    b"                        else if (rx == 2)\r\n"
    b"                            nFilterEdge = m_nHightlightFilterEdgeRX2;\r\n"
    b"\r\n"
    b"                        switch (nFilterEdge)\r\n"
    b"                        {\r\n"
    b"                            case -1:\r\n"
    b"                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_left_x, nVerticalShift + top, filter_left_x, nVerticalShift + H, 2);\r\n"
    b"                                break;\r\n"
    b"                            case 1:\r\n"
    b"                                drawLineDX2D(m_bDX2_cw_zero_pen, filter_right_x, nVerticalShift + top, filter_right_x, nVerticalShift + H, 2);\r\n"
    b"                                break;\r\n"
    b"                            default:\r\n"
    b"                                break;\r\n"
    b"                        }\r\n"
    b"                    }\r\n"
    b"\r\n"
    b"                    // W5TSU: panadapter sync/SNR overlay (sub-project 5 of 6)\r\n"
    b"                    // -- draws only here: not on the waterfall (!bIsWaterfall,\r\n"
    b"                    // this block), not during TX (enclosing !local_mox block).\r\n"
    b"                    if (!bIsWaterfall && m_bShowRadeSyncOverlay)\r\n"
    b"                    {\r\n"
    b"                        int thread = rx - 1; // GetDSPRX's plain thread index: 0 = RX1, 1 = RX2\r\n"
    b"                        RadioDSPRX dsp = console.radio.GetDSPRX(thread, 0);\r\n"
    b"                        bool radeOn = dsp.RXRadaeEnabled != 0;\r\n"
    b"                        bool fdv700eOn = dsp.RXAFDVRun != 0;\r\n"
    b"\r\n"
    b"                        if (radeOn || fdv700eOn)\r\n"
    b"                        {\r\n"
    b"                            bool sync;\r\n"
    b"                            string snrText;\r\n"
    b"\r\n"
    b"                            if (fdv700eOn)\r\n"
    b"                            {\r\n"
    b"                                // WDSP.id() folds thread+subrx into one wdsp.dll channel\r\n"
    b"                                // index (channel = 2*thread + subrx, dsp.cs:1161) -- RX2\r\n"
    b"                                // is WDSP.id(2, 0), NOT WDSP.id(1, 0) (that resolves to\r\n"
    b"                                // the TX channel). Verbatim from radeRX2StatusTimer_Tick /\r\n"
    b"                                // radeStatusTimer_Tick, setup.cs:36947-36948 / 36815-36816.\r\n"
    b"                                int channel = WDSP.id((uint)(thread == 0 ? 0 : 2), 0);\r\n"
    b"                                sync = WDSP.GetRXAFDVSync(channel) != 0;\r\n"
    b'                                snrText = sync ? string.Format("{0:F1}", WDSP.GetRXAFDVSnr(channel)) : "";\r\n'
    b"                            }\r\n"
    b"                            else // radeOn\r\n"
    b"                            {\r\n"
    b"                                // RADE V1/V2: ChannelMaster's radae.c uses its own plain\r\n"
    b"                                // 0/1 rx index, unrelated to WDSP.id()'s doubled\r\n"
    b"                                // convention. Verbatim from radeStatusTimer_Tick,\r\n"
    b"                                // setup.cs:36820-36821 (rx=0) / radeRX2StatusTimer_Tick,\r\n"
    b"                                // setup.cs:36952-36953 (rx=1).\r\n"
    b"                                sync = WDSP.GetRadaeSync(thread) != 0;\r\n"
    b'                                snrText = sync ? string.Format("{0}", WDSP.GetRadaeSnrDb(thread)) : "";\r\n'
    b"                            }\r\n"
    b"\r\n"
    b'                            string overlayText = sync ? string.Format("SYNC  SNR {0} dB", snrText) : "no sync";\r\n'
    b"                            SharpDX.Direct2D1.Brush overlayBrush = sync ? getDXBrushForColour(Color.Green) : m_bDX2_grid_text_brush;\r\n"
    b"\r\n"
    b"                            SizeF overlaySize = measureStringDX2D(overlayText, fontDX2d_font9);\r\n"
    b"                            drawStringDX2D(overlayText, fontDX2d_font9, overlayBrush, filter_left_x, nVerticalShift + top - overlaySize.Height);\r\n"
    b"                        }\r\n"
    b"                    }\r\n"
    b"                }\r\n"
    b"            }\r\n"
)
data = data.replace(old, new_block)
open(path, "wb").write(data)
```

- [ ] **Step 4: Run it, then verify with git diff**

Run the script, then `git diff --stat -- "Project Files/Source/Console/display.cs"`
(cumulative with Step 2's change, expect roughly +48/-0 lines total for
both splices). Same recovery procedure as Step 2 if the diff is far larger
than expected.

- [ ] **Step 5: Build**

This is a Windows-only C#/native solution — building requires the CI
pipeline (see Task 2's Step 8, which builds and verifies both tasks
together) or a local Visual Studio 2022 checkout if available. There is no
cross-platform build step to run here; do not attempt `dotnet build` or
`msbuild` on a non-Windows machine — it will not succeed and is not this
step's verification method. If you have Windows/VS2022 available locally,
run:
```
msbuild "Project Files/Source/Thetis_VS2026.sln" /p:Configuration=Release /p:Platform=x64 /p:PlatformToolset=v143 /m /v:minimal
```
and confirm zero errors. Otherwise, defer full build verification to
Task 2's Step 8 (both tasks are built together there via CI) — do not skip
building entirely, just don't duplicate it here if it's about to happen
anyway in Task 2.

- [ ] **Step 6: Commit**

```bash
git add "Project Files/Source/Console/display.cs"
git commit -m "feat(display): panadapter sync/SNR overlay for Digital Voice modes"
```

---

### Task 2: Setup checkbox + wiring

**Files:**
- Modify: `Project Files/Source/Console/setup.designer.cs` (3 insertion
  points: object instantiation, field declaration, properties block)
- Modify: `Project Files/Source/Console/setup.cs` (2 insertion points:
  `chkShowRadeSyncOverlay_CheckedChanged` handler,
  `InitRadePanelFromBackend` sync)

**Interfaces:**
- Consumes: `Display.ShowRadeSyncOverlay` (Task 1's new public static
  property — get to sync the checkbox on tab load, set from the checkbox's
  own handler).
- Produces: nothing further downstream — this is the outermost UI layer
  for this feature.

- [ ] **Step 1: Add the control's instantiation**

Read `setup.designer.cs` first. Using Python byte-splicing, locate this
exact, unique anchor (part of the RX2 Core controls' instantiation block):

Old bytes:
```
            this.lblRadeRX2Status = new System.Windows.Forms.LabelTS();
```

New bytes:
```
            this.lblRadeRX2Status = new System.Windows.Forms.LabelTS();
            this.chkShowRadeSyncOverlay = new System.Windows.Forms.CheckBoxTS();
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()
old = b"            this.lblRadeRX2Status = new System.Windows.Forms.LabelTS();\r\n"
assert data.count(old) == 1, "anchor not found or not unique"
new = old + b"            this.chkShowRadeSyncOverlay = new System.Windows.Forms.CheckBoxTS();\r\n"
data = data.replace(old, new)
open(path, "wb").write(data)
```

- [ ] **Step 2: Verify with git diff**

`git diff --stat -- "Project Files/Source/Console/setup.designer.cs"` —
expect roughly +1/-0. Same flattening-recovery procedure as Task 1 if not.

- [ ] **Step 3: Add the control's properties and add it to the tab**

Read `setup.designer.cs` again. Using Python byte-splicing, locate this
exact, unique anchor (the end of `chkRadeRX1Loopback`'s own properties
block, immediately preceded by its declaration — this text appears once):

Old bytes:
```
            this.chkRadeRX1Loopback.AutoSize = true;
            this.chkRadeRX1Loopback.Image = null;
            this.chkRadeRX1Loopback.Location = new System.Drawing.Point(16, 48);
            this.chkRadeRX1Loopback.Name = "chkRadeRX1Loopback";
            this.chkRadeRX1Loopback.Size = new System.Drawing.Size(132, 17);
            this.chkRadeRX1Loopback.TabIndex = 1;
            this.chkRadeRX1Loopback.Text = "RX1 RADE Loopback Test";
            this.toolTip1.SetToolTip(this.chkRadeRX1Loopback, "TX encoder's modem output is bridged directly into RX1's decoder input -- no RF, radio never keys. For verifying the encode/decode round trip before any real on-air attempt.");
            this.chkRadeRX1Loopback.UseVisualStyleBackColor = true;
            this.chkRadeRX1Loopback.CheckedChanged += new System.EventHandler(this.chkRadeRX1Loopback_CheckedChanged);
```

New bytes (old bytes + a new block appended immediately after, defining
`chkShowRadeSyncOverlay` as a direct child of the tab page, positioned
below `grpRadeMicCond` which occupies `(16,16)` sized `320x300` — i.e.
ends at y=316):

```
            this.chkRadeRX1Loopback.AutoSize = true;
            this.chkRadeRX1Loopback.Image = null;
            this.chkRadeRX1Loopback.Location = new System.Drawing.Point(16, 48);
            this.chkRadeRX1Loopback.Name = "chkRadeRX1Loopback";
            this.chkRadeRX1Loopback.Size = new System.Drawing.Size(132, 17);
            this.chkRadeRX1Loopback.TabIndex = 1;
            this.chkRadeRX1Loopback.Text = "RX1 RADE Loopback Test";
            this.toolTip1.SetToolTip(this.chkRadeRX1Loopback, "TX encoder's modem output is bridged directly into RX1's decoder input -- no RF, radio never keys. For verifying the encode/decode round trip before any real on-air attempt.");
            this.chkRadeRX1Loopback.UseVisualStyleBackColor = true;
            this.chkRadeRX1Loopback.CheckedChanged += new System.EventHandler(this.chkRadeRX1Loopback_CheckedChanged);
            //
            // chkShowRadeSyncOverlay
            //
            this.chkShowRadeSyncOverlay.AutoSize = true;
            this.chkShowRadeSyncOverlay.Image = null;
            this.chkShowRadeSyncOverlay.Location = new System.Drawing.Point(16, 328);
            this.chkShowRadeSyncOverlay.Name = "chkShowRadeSyncOverlay";
            this.chkShowRadeSyncOverlay.Size = new System.Drawing.Size(190, 17);
            this.chkShowRadeSyncOverlay.TabIndex = 51;
            this.chkShowRadeSyncOverlay.Text = "Show sync/SNR on panadapter";
            this.toolTip1.SetToolTip(this.chkShowRadeSyncOverlay, "Shows the active Digital Voice mode's live SYNC/SNR status directly above the tuned filter on the panadapter, for RX1 and RX2 independently. Hidden during TX.");
            this.chkShowRadeSyncOverlay.UseVisualStyleBackColor = true;
            this.chkShowRadeSyncOverlay.CheckedChanged += new System.EventHandler(this.chkShowRadeSyncOverlay_CheckedChanged);
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()
old = (
    b"            this.chkRadeRX1Loopback.AutoSize = true;\r\n"
    b"            this.chkRadeRX1Loopback.Image = null;\r\n"
    b"            this.chkRadeRX1Loopback.Location = new System.Drawing.Point(16, 48);\r\n"
    b'            this.chkRadeRX1Loopback.Name = "chkRadeRX1Loopback";\r\n'
    b"            this.chkRadeRX1Loopback.Size = new System.Drawing.Size(132, 17);\r\n"
    b"            this.chkRadeRX1Loopback.TabIndex = 1;\r\n"
    b'            this.chkRadeRX1Loopback.Text = "RX1 RADE Loopback Test";\r\n'
    b'            this.toolTip1.SetToolTip(this.chkRadeRX1Loopback, "TX encoder\'s modem output is bridged directly into RX1\'s decoder input -- no RF, radio never keys. For verifying the encode/decode round trip before any real on-air attempt.");\r\n'
    b"            this.chkRadeRX1Loopback.UseVisualStyleBackColor = true;\r\n"
    b"            this.chkRadeRX1Loopback.CheckedChanged += new System.EventHandler(this.chkRadeRX1Loopback_CheckedChanged);\r\n"
)
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"            //\r\n"
    b"            // chkShowRadeSyncOverlay\r\n"
    b"            //\r\n"
    b"            this.chkShowRadeSyncOverlay.AutoSize = true;\r\n"
    b"            this.chkShowRadeSyncOverlay.Image = null;\r\n"
    b"            this.chkShowRadeSyncOverlay.Location = new System.Drawing.Point(16, 328);\r\n"
    b'            this.chkShowRadeSyncOverlay.Name = "chkShowRadeSyncOverlay";\r\n'
    b"            this.chkShowRadeSyncOverlay.Size = new System.Drawing.Size(190, 17);\r\n"
    b"            this.chkShowRadeSyncOverlay.TabIndex = 51;\r\n"
    b'            this.chkShowRadeSyncOverlay.Text = "Show sync/SNR on panadapter";\r\n'
    b'            this.toolTip1.SetToolTip(this.chkShowRadeSyncOverlay, "Shows the active Digital Voice mode\'s live SYNC/SNR status directly above the tuned filter on the panadapter, for RX1 and RX2 independently. Hidden during TX.");\r\n'
    b"            this.chkShowRadeSyncOverlay.UseVisualStyleBackColor = true;\r\n"
    b"            this.chkShowRadeSyncOverlay.CheckedChanged += new System.EventHandler(this.chkShowRadeSyncOverlay_CheckedChanged);\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

- [ ] **Step 4: Add it to the tab's Controls collection and declare the field**

Read `setup.designer.cs` again. Two more splices in this step.

First — locate this exact, unique anchor (`tpDSPRADE`'s `Controls.Add`
list):

Old bytes:
```
            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX2Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
```

New bytes:
```
            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX2Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
            this.tpDSPRADE.Controls.Add(this.chkShowRadeSyncOverlay);
```

Second — locate this exact, unique anchor (the field declarations block,
end of the RADE tab's fields):

Old bytes:
```
        private CheckBoxTS chkRadeRX2Loopback;
        private LabelTS lblRadeRX2Level;
        private NumericUpDownTS udRadeRX2Level;
        private LabelTS lblRadeRX2Status;
```

New bytes:
```
        private CheckBoxTS chkRadeRX2Loopback;
        private LabelTS lblRadeRX2Level;
        private NumericUpDownTS udRadeRX2Level;
        private LabelTS lblRadeRX2Status;
        private CheckBoxTS chkShowRadeSyncOverlay;
```

```python
path = "Project Files/Source/Console/setup.designer.cs"
data = open(path, "rb").read()

old1 = (
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);\r\n"
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);\r\n"
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeRX2Core);\r\n"
    b"            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);\r\n"
)
assert data.count(old1) == 1, "anchor 1 not found or not unique"
new1 = old1 + b"            this.tpDSPRADE.Controls.Add(this.chkShowRadeSyncOverlay);\r\n"
data = data.replace(old1, new1)

old2 = (
    b"        private CheckBoxTS chkRadeRX2Loopback;\r\n"
    b"        private LabelTS lblRadeRX2Level;\r\n"
    b"        private NumericUpDownTS udRadeRX2Level;\r\n"
    b"        private LabelTS lblRadeRX2Status;\r\n"
)
assert data.count(old2) == 1, "anchor 2 not found or not unique"
new2 = old2 + b"        private CheckBoxTS chkShowRadeSyncOverlay;\r\n"
data = data.replace(old2, new2)

open(path, "wb").write(data)
```

- [ ] **Step 5: Verify with git diff**

`git diff --stat -- "Project Files/Source/Console/setup.designer.cs"` —
cumulative across Steps 1, 3, 4: expect roughly +17/-0 total. Recover via
`git checkout` + redo if the diff is far larger.

- [ ] **Step 6: Add the checkbox's event handler in `setup.cs`**

Read `setup.cs` first. Using Python byte-splicing, locate this exact,
unique anchor (`chkRadeRX1Loopback_CheckedChanged`'s full existing body):

Old bytes:
```
        // W5TSU: RX1 loopback bridge -- mode-aware, calls whichever
        // subsystem's loopback matches the current Mode selection
        // (radae.c SetRadaeLoopbackEnabled for RADE V1/V2, ChannelMaster
        // SetFDVLoopbackEnabled for 700E). Disabled entirely for Off, see
        // cmbRadeMode_SelectedIndexChanged.
        private void chkRadeRX1Loopback_CheckedChanged(object sender, EventArgs e)
        {
            int mode = cmbRadeMode.SelectedIndex;
            bool on = chkRadeRX1Loopback.Checked;

            if (mode == 1) // 700E
            {
                WDSP.SetFDVLoopbackEnabled(on ? 1 : 0);
            }
            else if (mode == 2 || mode == 3) // RADE V1/V2
            {
                WDSP.SetRadaeLoopbackEnabled(0, on ? 1 : 0);
            }
        }
```

New bytes (old bytes + a new handler appended immediately after):

```
        // W5TSU: RX1 loopback bridge -- mode-aware, calls whichever
        // subsystem's loopback matches the current Mode selection
        // (radae.c SetRadaeLoopbackEnabled for RADE V1/V2, ChannelMaster
        // SetFDVLoopbackEnabled for 700E). Disabled entirely for Off, see
        // cmbRadeMode_SelectedIndexChanged.
        private void chkRadeRX1Loopback_CheckedChanged(object sender, EventArgs e)
        {
            int mode = cmbRadeMode.SelectedIndex;
            bool on = chkRadeRX1Loopback.Checked;

            if (mode == 1) // 700E
            {
                WDSP.SetFDVLoopbackEnabled(on ? 1 : 0);
            }
            else if (mode == 2 || mode == 3) // RADE V1/V2
            {
                WDSP.SetRadaeLoopbackEnabled(0, on ? 1 : 0);
            }
        }

        // W5TSU: panadapter sync/SNR overlay toggle (sub-project 5 of 6) --
        // one-line passthrough to Display's static flag, guarded like every
        // other Digital Voice tab handler against firing during Setup's own
        // startup sequence.
        private void chkShowRadeSyncOverlay_CheckedChanged(object sender, EventArgs e)
        {
            if (initializing) return;
            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;
        }
```

```python
path = "Project Files/Source/Console/setup.cs"
data = open(path, "rb").read()
old = (
    b"        // W5TSU: RX1 loopback bridge -- mode-aware, calls whichever\r\n"
    b"        // subsystem's loopback matches the current Mode selection\r\n"
    b"        // (radae.c SetRadaeLoopbackEnabled for RADE V1/V2, ChannelMaster\r\n"
    b"        // SetFDVLoopbackEnabled for 700E). Disabled entirely for Off, see\r\n"
    b"        // cmbRadeMode_SelectedIndexChanged.\r\n"
    b"        private void chkRadeRX1Loopback_CheckedChanged(object sender, EventArgs e)\r\n"
    b"        {\r\n"
    b"            int mode = cmbRadeMode.SelectedIndex;\r\n"
    b"            bool on = chkRadeRX1Loopback.Checked;\r\n"
    b"\r\n"
    b"            if (mode == 1) // 700E\r\n"
    b"            {\r\n"
    b"                WDSP.SetFDVLoopbackEnabled(on ? 1 : 0);\r\n"
    b"            }\r\n"
    b"            else if (mode == 2 || mode == 3) // RADE V1/V2\r\n"
    b"            {\r\n"
    b"                WDSP.SetRadaeLoopbackEnabled(0, on ? 1 : 0);\r\n"
    b"            }\r\n"
    b"        }\r\n"
)
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"\r\n"
    b"        // W5TSU: panadapter sync/SNR overlay toggle (sub-project 5 of 6) --\r\n"
    b"        // one-line passthrough to Display's static flag, guarded like every\r\n"
    b"        // other Digital Voice tab handler against firing during Setup's own\r\n"
    b"        // startup sequence.\r\n"
    b"        private void chkShowRadeSyncOverlay_CheckedChanged(object sender, EventArgs e)\r\n"
    b"        {\r\n"
    b"            if (initializing) return;\r\n"
    b"            Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;\r\n"
    b"        }\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

- [ ] **Step 7: Sync the checkbox from backend state in `InitRadePanelFromBackend`**

Read `setup.cs` again. Using Python byte-splicing, locate this exact,
unique anchor (the end of `InitRadePanelFromBackend`'s `cmbRadeMode`
sync block):

Old bytes:
```
            cmbRadeMode.SelectedIndexChanged -= cmbRadeMode_SelectedIndexChanged;
            int mode;
            if (console.radio.GetDSPRX(0, 0).RXRadaeEnabled != 0)
                mode = (WDSP.GetRadaeProtocolV2(0) != 0) ? 3 : 2;
            else if (console.radio.GetDSPRX(0, 0).RXAFDVRun != 0)
                mode = 1;
            else
                mode = 0;
            cmbRadeMode.SelectedIndex = mode;
            cmbRadeMode.SelectedIndexChanged += cmbRadeMode_SelectedIndexChanged;
```

New bytes (old bytes + a new unsub/set/resub block for the new checkbox,
mirroring `cmbRadeMode`'s own exact pattern rather than relying solely on
the `initializing` guard, since this function runs after `initializing` is
already false):

```
            cmbRadeMode.SelectedIndexChanged -= cmbRadeMode_SelectedIndexChanged;
            int mode;
            if (console.radio.GetDSPRX(0, 0).RXRadaeEnabled != 0)
                mode = (WDSP.GetRadaeProtocolV2(0) != 0) ? 3 : 2;
            else if (console.radio.GetDSPRX(0, 0).RXAFDVRun != 0)
                mode = 1;
            else
                mode = 0;
            cmbRadeMode.SelectedIndex = mode;
            cmbRadeMode.SelectedIndexChanged += cmbRadeMode_SelectedIndexChanged;

            chkShowRadeSyncOverlay.CheckedChanged -= chkShowRadeSyncOverlay_CheckedChanged;
            chkShowRadeSyncOverlay.Checked = Display.ShowRadeSyncOverlay;
            chkShowRadeSyncOverlay.CheckedChanged += chkShowRadeSyncOverlay_CheckedChanged;
```

```python
path = "Project Files/Source/Console/setup.cs"
data = open(path, "rb").read()
old = (
    b"            cmbRadeMode.SelectedIndexChanged -= cmbRadeMode_SelectedIndexChanged;\r\n"
    b"            int mode;\r\n"
    b"            if (console.radio.GetDSPRX(0, 0).RXRadaeEnabled != 0)\r\n"
    b"                mode = (WDSP.GetRadaeProtocolV2(0) != 0) ? 3 : 2;\r\n"
    b"            else if (console.radio.GetDSPRX(0, 0).RXAFDVRun != 0)\r\n"
    b"                mode = 1;\r\n"
    b"            else\r\n"
    b"                mode = 0;\r\n"
    b"            cmbRadeMode.SelectedIndex = mode;\r\n"
    b"            cmbRadeMode.SelectedIndexChanged += cmbRadeMode_SelectedIndexChanged;\r\n"
)
assert data.count(old) == 1, "anchor not found or not unique"
new = old + (
    b"\r\n"
    b"            chkShowRadeSyncOverlay.CheckedChanged -= chkShowRadeSyncOverlay_CheckedChanged;\r\n"
    b"            chkShowRadeSyncOverlay.Checked = Display.ShowRadeSyncOverlay;\r\n"
    b"            chkShowRadeSyncOverlay.CheckedChanged += chkShowRadeSyncOverlay_CheckedChanged;\r\n"
)
data = data.replace(old, new)
open(path, "wb").write(data)
```

- [ ] **Step 8: Verify with git diff, then build and deploy to hl2winbox**

`git diff --stat -- "Project Files/Source/Console/setup.cs"` — expect
roughly +14/-0 (Steps 6+7 combined). Recover via `git checkout` + redo if
far larger.

Then run the full CI-build-and-deploy cycle this project has used for
every C#-touching sub-project so far:
1. `git push` (CI only sees the remote, not local commits).
2. `gh workflow run build.yml --ref FreeDV -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`,
   poll until it completes, download the `Thetis-HL2-installer` artifact.
3. `scp` the installer to `hl2winbox:Downloads/`.
4. Run the existing `extract.ps1` (msiexec admin-extract) → `stop_thetis.ps1`
   → `copy_build.ps1` (robocopy /MIR) → relaunch via `run_interactive.ps1`
   sequence already established in this project's prior sub-projects.
5. Verify the new build is live: `thetisctl cat --host 192.168.2.12 version`
   (the LAN address — the Tailscale address in some earlier session notes
   for this same box is unreliable; use the LAN address) should report a
   `git:` short SHA matching this task's own commit.

- [ ] **Step 9: Live verification on hl2winbox**

Per the spec's Testing section:
1. Enable RADE V1 with RX1 loopback active; confirm the overlay appears
   above RX1's filter passband showing `"SYNC  SNR X dB"` live-updating
   text that matches `lblRadeRX1Status`'s own value in Setup at the same
   moment.
2. Key PTT/MOX; confirm the overlay disappears for the duration of TX and
   reappears on drop to RX.
3. Uncheck `chkShowRadeSyncOverlay` in the Digital Voice tab; confirm the
   overlay disappears within a frame despite RX1 still being synced.
   Re-check; confirm it reappears.
4. Enable RX2 split with RADE V1 or 700E active on RX2; repeat steps 1-2
   for RX2's own pane, confirming both panes show independent, correct
   state simultaneously (e.g. RX1 synced, RX2 not, or vice versa).
5. Switch a mode to Off; confirm the overlay simply doesn't draw for that
   rx.

If step 4's remote visual confirmation is blocked by this project's
already-documented remote-desktop-screenshot unreliability (no interactive
desktop session reachable over plain SSH, hit repeatedly in sub-projects
#2, #3, and #4's own live checks), it's acceptable to defer only the
visual screen-capture proof as an explicitly-flagged gap — but do not skip
actually exercising the checkbox and mode changes over CAT/TCI or in
person if any hands-on access to the box is available this session; the
functional behavior (not merely "it compiles") must be confirmed live
before this task is done.

- [ ] **Step 10: Commit**

```bash
git add "Project Files/Source/Console/setup.designer.cs" "Project Files/Source/Console/setup.cs"
git commit -m "feat(setup): add panadapter sync/SNR overlay toggle to Digital Voice tab"
```

---
