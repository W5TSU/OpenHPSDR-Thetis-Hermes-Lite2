# Panadapter Sync/SNR Overlay — Design Spec

**Date**: 2026-08-25
**Status**: Approved, ready for implementation planning
**Sub-project**: 5 of 6 (see [2026-08-24-rade-setup-panel-design.md](2026-08-24-rade-setup-panel-design.md)'s
"Larger context" section for #1-#5; a sixth sub-project — Digital Voice tab
renamed to "FreeDV" plus FreeDV Reporter enable/self-report UI — was added
during this spec's brainstorming and will get its own design cycle after
this one ships)

## Goal

Show the currently active Digital Voice mode's live SYNC/SNR status
directly on Thetis's own panadapter, so an operator can see decode quality
at a glance without opening Setup. This is the "Panadapter SNR/sync overlay
display" originally scoped in sub-project #1's "Larger context" section —
unrelated to sub-project #4's FreeDV Reporter spotting (which shows *other*
stations' activity); this overlay shows *this station's own* RX decode
quality.

## What already exists (why this is small)

All of the underlying data and every rendering primitive this needs are
already in the codebase — this sub-project is new wiring between two
already-complete halves, not new subsystems on either side.

1. **The sync/SNR data is already read and formatted** by sub-projects
   #1-#3's Setup-panel status timers (`setup.cs`'s `radeStatusTimer_Tick`
   for RX1, `radeRX2StatusTimer_Tick` for RX2). Both already know exactly
   which backend to query and how to format the result — this spec's draw
   path reuses that same logic verbatim (see "Data source" below), it does
   not invent a new way to determine sync/SNR.
2. **`display.cs` already has a directly-analogous overlay** —
   `BandStackOverlay` (`m_bShowBandStackOverlays` flag +
   `m_bandStackOverlays` static data + `drawFilterOverlayDX2D`, gated
   `rx == 1`) is drawn today from exactly this same per-RX, per-frame draw
   path. **Correction from an earlier draft of this spec, verified by
   reading the actual function boundaries**: this code — and the RX's own
   filter-passband overlay it sits beside — lives inside
   `drawPanadapterAndWaterfallGridDX2D` (`display.cs:8802`), a single
   shared helper called once for the actual panadapter/spectrum pass
   (`DrawPanadapterDX2D:4980`, with its own `bIsWaterfall` parameter
   `false`) and once for the waterfall pass (`DrawWaterfallDX2D:7707`,
   `bIsWaterfall` `true`) — not a function named `DrawPanadapterDX2D`
   itself, which only handles a few other overlay elements (the
   noise-floor readout). This sub-project follows the same shape: a
   static show/hide flag, read at draw time, drawn per-rx, inside this
   shared helper's own `!bIsWaterfall`-gated branch (so it never appears on
   the waterfall) and its own `if (!local_mox)`-gated branch (so it's
   already guaranteed hidden during TX, no separate MOX check needed).
3. **Text drawing, string measurement, and color-brush helpers all
   already exist** and are used throughout `display.cs`'s DX2D draw
   functions for similar small in-spectrum labels (the noise-floor
   readout, the FPS/CPU diagnostic overlay, waterfall time labels,
   frequency-scale grid labels) — `drawStringDX2D`,
   `measureStringDX2D`, and `getDXBrushForColour` (a self-caching
   Color→Brush helper, `display.cs:11570`, already used for one-off fixed
   status colors like `m_bDX2_Red`/`m_bDX2_Yellow` at setup time and
   directly at draw time elsewhere, e.g. `display.cs:6113-6115`). No new
   brush lifecycle management (manual dispose/null-out) is needed — that
   pattern exists for brushes that must survive skin changes; a fixed
   semantic color like "sync is good" doesn't need it.
4. **`display.cs` is Direct2D-only** — confirmed by inspection: 240
   `*DX2D` drawing functions, zero SkiaSharp usage in this file. Only one
   rendering path needs this change, not two.

The actual remaining gap: nothing currently reads this data from
`display.cs`'s draw path or renders it on the panadapter. That's the whole
scope here.

## Design

### Data source — reusing the exact existing determination, not a new one

For each RX (`rx` 1 or 2, i.e. `thread` 0 or 1 in `GetDSPRX`'s indexing):

```csharp
var dsp = console.radio.GetDSPRX(thread, 0);
bool radeOn = dsp.RXRadaeEnabled != 0;
bool fdv700eOn = dsp.RXAFDVRun != 0;
bool active = radeOn || fdv700eOn;
```

This is the same two-flag check `setup.cs`'s `InitRadePanelFromBackend`
already uses (`setup.cs:37135-37140`) to derive `cmbRadeMode`'s selection —
`RXRadaeEnabled` and `RXAFDVRun` are mutually exclusive by construction
(sub-project #2's interlock, `radio.cs`), so no third state is possible.
`thread` is 0 for RX1, 1 for RX2 — this is `GetDSPRX`'s own indexing
convention, the *plain* RX index, not the doubled `WDSP.id()` channel
index used for the DLL call below.

Once `active` is known, the sync/SNR call itself must branch on which
backend is active, and **the two backends take different index
conventions** — this exact confusion caused two Critical bugs in
sub-project #3, so the call sites below are copied verbatim from the
already-fixed, already-verified-live `setup.cs` status timers rather than
re-derived:

```csharp
bool sync; string snrText;
if (fdv700eOn)
{
    // 700E: WDSP.id() folds thread+subrx into one wdsp.dll channel index
    // (channel = 2*thread + subrx, dsp.cs:1161) — RX2 is WDSP.id(2, 0),
    // NOT WDSP.id(1, 0) (that resolves to the TX channel). Verbatim from
    // radeRX2StatusTimer_Tick, setup.cs:36947-36948 / radeStatusTimer_Tick,
    // setup.cs:36815-36816 for RX1 (WDSP.id(0, 0)).
    int channel = WDSP.id(thread == 0 ? 0 : 2, 0);
    sync = WDSP.GetRXAFDVSync(channel) != 0;
    snrText = sync ? string.Format("{0:F1}", WDSP.GetRXAFDVSnr(channel)) : "";
}
else // radeOn
{
    // RADE V1/V2: ChannelMaster's radae.c uses its own plain 0/1 rx index,
    // unrelated to WDSP.id()'s doubled convention. Verbatim from
    // radeStatusTimer_Tick, setup.cs:36820-36821 (rx=0) /
    // radeRX2StatusTimer_Tick, setup.cs:36952-36953 (rx=1).
    int rxIndex = thread; // 0 or 1, already the correct radae.c index
    sync = WDSP.GetRadaeSync(rxIndex) != 0;
    snrText = sync ? string.Format("{0}", WDSP.GetRadaeSnrDb(rxIndex)) : "";
}
```

`WDSP.GetRXAFDVSync(int channel)`/`GetRXAFDVSnr(int channel)` and
`WDSP.GetRadaeSync(int rx)`/`GetRadaeSnrDb(int rx)` are the exact existing
P/Invoke signatures (`dsp.cs:266,269,313,316`) — no new native exports are
needed.

No new timer is introduced in `display.cs`: the panadapter already
redraws continuously (it's a live spectrum), so this data is read fresh
directly inside the existing per-frame draw call — polling happens for
free as a side effect of the render loop already running.

### TX gating

`drawPanadapterAndWaterfallGridDX2D` already has a MOX/TX signal threaded
into its draw path for the existing filter overlay's own TX-awareness
(`local_mox`) — the RX filter overlay block and everything after it in
this function only runs inside `if (!local_mox)`. Placing this overlay's
code inside that same block (see "Rendering" below) means it is already
guaranteed hidden during TX as a structural side effect of where it lives
— no separate MOX check needed.

### Rendering

Inside `drawPanadapterAndWaterfallGridDX2D`, the new code sits immediately
after the existing "RX FILTER overlay + highlight edges" block's own
`if (!bIsWaterfall) { ... nFilterEdge switch ... }` sub-block (the one
that highlights a filter edge for CW), as a sibling `if (!bIsWaterfall)`
block at the same nesting level — still inside the enclosing
`if (!local_mox)` and `if ((bIsWaterfall && m_bShowRXFilterOnWaterfall) ||
!bIsWaterfall)` blocks, where `filter_left_x`, `filter_right_x`, `rx`,
`top`, `nVerticalShift`, `H`, and `W` are all already in scope from the
surrounding code — no new parameters need to be threaded in.

Per rx (1 and 2 — RX2's own pane naturally only reaches this code when
its call to `drawPanadapterAndWaterfallGridDX2D` happens at all, the same
way the existing filter-overlay code already only runs for whichever rx
is actually being drawn this call):

- Skip entirely if `!ShowRadeSyncOverlay` (new toggle, see below) or if
  `!active` for this rx (TX and waterfall are already excluded structurally
  by this insertion point, see "TX gating" above).
- Otherwise compute the display string: `"SYNC  SNR {snrText} dB"` when
  `sync`, else `"no sync"` — matching the Setup panel's own exact wording
  for consistency between the two surfaces.
- Color: `getDXBrushForColour(Color.Green)` when `sync`, otherwise reuse
  `m_bDX2_grid_text_brush` (the skin's own existing grid/frequency-label
  text color, already computed and cached each time the DX2D surface is
  set up, and already used directly as a `drawStringDX2D` brush argument
  elsewhere in this same file, e.g. `display.cs:10315-10316`'s frequency
  labels) — deliberately not a WinForms `SystemColors` value, since those
  are meant for form controls and would not read correctly against an
  arbitrary (often dark, often user-skinned) spectrum background.
- Position: `filter_left_x` (already computed at this point for this
  rx's own filter-passband shading — the exact X pixel position of the
  currently-tuned passband) for the X coordinate; vertically, a few
  pixels above the top of that same passband shading (`nVerticalShift +
  top`, the same top-left corner `drawFilterOverlayDX2D`'s own rectangle
  starts from) minus the text's own measured height via
  `measureStringDX2D`, so the label sits directly above the passband it
  describes without overlapping it — the same "measure then offset above"
  placement already used for the noise-floor readout label
  (`display.cs:5449`, `nf_box.Y - 6`).
- Draw via the existing `drawStringDX2D(text, font, brush, x, y)` helper
  (`display.cs:8504`) with `fontDX2d_font9` — the same small label font
  already used for the noise-floor value/waterfall time labels, keeping
  this visually consistent with the panadapter's other compact readouts
  rather than introducing a new font resource.

### Visibility toggle

A new checkbox in the Digital Voice Setup tab — `chkShowRadeSyncOverlay`,
label "Show sync/SNR on panadapter", default checked — sets a new public
static `Display.ShowRadeSyncOverlay` bool, following the exact existing
`ShowBandStackOverlays` pattern (`display.cs:995-1000`: a private static
backing field plus a public static bool property, read directly by the
draw path with no event/notification needed since the very next frame
picks up the new value). The checkbox's `CheckedChanged` handler is a
one-line `Display.ShowRadeSyncOverlay = chkShowRadeSyncOverlay.Checked;`,
guarded by the existing `if (initializing) return;` convention this
project's Setup-panel work has used in every sub-project so far (a defect
class the review process caught twice in sub-projects #2 and #3). Setup's
existing generic control-persistence mechanism handles saving/restoring
the checkbox's state across restarts the same way it does for every other
checkbox in that tab — no new persistence code.

### Testing

Live verification on hl2winbox, the same hardware-in-the-loop approach
used for sub-projects #1-#3's Setup-panel work:

1. Enable RADE V1 with RX1 loopback active (the existing test setup from
   sub-project #1/#2) and confirm the overlay appears above RX1's filter
   passband, showing live-updating `"SYNC  SNR X dB"` text that tracks the
   same values the Setup panel's own `lblRadeRX1Status` shows at the same
   moment (cross-check against the existing, already-verified label — this
   overlay must never show a different sync state or SNR value than the
   Setup panel does for the same RX at the same instant, since both read
   the same underlying calls).
2. Key PTT/MOX and confirm the overlay disappears for the duration of TX,
   then reappears on drop back to RX.
3. Uncheck `chkShowRadeSyncOverlay` and confirm the overlay disappears
   immediately (next frame) despite RX1 still being synced; re-check and
   confirm it reappears.
4. Enable RX2 split with RADE V1 or 700E active on RX2 and repeat steps
   1-2 for RX2's own pane, independently of RX1's state (e.g. RX1 in RADE
   V1 sync, RX2 in "no sync" — both panes must show their own correct,
   independent state simultaneously).
5. Switch a mode to Off and confirm the overlay simply doesn't draw for
   that rx (the `!active` gate), same as today's Setup-panel status label
   showing "off".

No automated/unit test is applicable here — this is a pixel-rendering
path inside a native-interop-heavy static rendering class with no existing
test harness in this codebase; hardware-in-the-loop visual verification
is this project's established practice for `display.cs`/Setup-panel work
(see sub-projects #1-#3's own testing sections).

## Explicitly out of scope

- Anything related to sub-project #4's FreeDV Reporter spotting
  (`SpotManager2.cs`, other stations' activity) — this overlay is this
  station's own RX decode quality only, a completely separate data source
  and a completely separate rendering path (this overlay lives inside
  `DrawPanadapterDX2D` itself; `SpotManager2`'s spots are a distinct
  overlay system already, per sub-project #4's own research).
- Sub-project #6 (Digital Voice tab renamed to "FreeDV", a UI toggle to
  launch/manage an external FreeDV Reporter helper process, and a new
  self-reporting-to-qso.freedv.org capability) — scoped as its own,
  separate design cycle after this one ships, per this session's
  brainstorming discussion.
- A TX-side overlay (showing anything about the *encoder*, e.g. a "TX
  active" indicator) — the existing MOX/PTT indicators elsewhere in the
  UI already cover this; this spec is RX-decode-quality only, matching the
  original "Panadapter SNR/sync overlay" scoping from sub-project #1.
- Waterfall-pane rendering — although the RX filter overlay this sub-project
  sits beside is drawn by a helper shared between the panadapter and
  waterfall passes (`drawPanadapterAndWaterfallGridDX2D`, see "Rendering"
  above), the new code is placed inside that helper's own
  `!bIsWaterfall`-gated branch, so it only ever draws on the panadapter
  pass, never the waterfall's — a compact SNR/sync text label doesn't fit
  the waterfall's continuously-scrolling presentation the way it fits the
  static spectrum view, matching the original "Panadapter SNR/sync
  overlay" scoping from sub-project #1.
- Persisting the overlay's on/off state anywhere beyond Setup's existing
  generic checkbox-persistence mechanism (no separate config file, no CAT
  command) — this is a display preference, not something CAT/TCI clients
  need to query or set remotely.
