# FreeDV-in-Thetis Project Plan (W5TSU)

## Goal

Native FreeDV digital voice support in the W5TSU Thetis build for the Hermes-Lite 2 —
starting with an RX-only 700E decode prototype to prove the wdsp insertion point, growing
toward a full FDV mode, with the neural RADE mode arriving later via upstream's C library
rather than our own port.

Development happens on the **`FreeDV` branch**. Status markers: ✅ done, ⬜ pending.

---

## Stage A — RX-only prototype (branch `FreeDV`)

### ✅ Phase 0 — codec2 toolchain *(done)*

- `.github/workflows/build-codec2.yml` dispatch workflow: MinGW64 build (codec2's CMake
  uses GCC-only flags, MSVC rejects it), `gendef` + `lib.exe` MSVC import library,
  `-static-libgcc` so the dll is self-contained (imports only KERNEL32/msvcrt)
- Vendored at `Project Files/lib/codec2_x64/`: `libcodec2.dll` (codec2 1.2.0, 121
  `freedv_*` exports), `codec2.lib`, public headers, LGPL-2.1 `COPYING`

### ✅ Phase 1 — wdsp DSP block *(done, CI green)*

- `wdsp/fdv.c` + `fdv.h`, modeled on the `rnnr.c` (NR3) block; sits **post-AGC** in the
  RXA chain (`RXA.c`, after the AGC meter)
- dsp_rate ↔ 8 kHz resampling via wdsp's `resampleF`, smoothed RMS gain to normalise
  wdsp's float levels into the modem's 16-bit domain, ring buffers bridging the dsp
  buffer size and `freedv_nin()`'s variable block size
- Passthrough-until-sync priming: raw modem audio stays audible for tuning; decode
  engages once synced and ~125 ms of output is buffered; underruns fall back gracefully
- Exports: `SetRXAFDVRun(channel, run)`, `GetRXAFDVSync(channel)`, `GetRXAFDVSnr(channel)`

### ✅ Phase 2 — console wiring *(done, CI green)*

- `dsp.cs`: P/Invoke declarations (rnnoise group pattern)
- `radio.cs`: `RXAFDVRun` cached property on `RadioDSPRX` (`RXANR3Run` pattern) — run
  flag survives wdsp channel rebuilds and applies via delayed update
- Setup → DSP → its own **"FreeDV"** tab (placed after AM/SAM and FM; originally
  tucked into the NR tab where it was easy to miss — moved `7289b20e`) — "FreeDV
  (prototype)" group, **"Decode FreeDV 700E (RX1)"** checkbox (auto-persisted) +
  live sync/SNR label polled at 500 ms
- Quick-Play failures (missing file, already-playing, bad header) used to silently
  uncheck the button with the error discarded; now shown in a message box with the
  attempted file path (`9006b38e`)

### ✅ Side-by-side test installer *(done — branch-only change)*

- On this branch the MSI installs as **"Thetis HL2 Test"** into
  `Program Files\OpenHPSDR\Thetis-Test\`, with its own x64 UpgradeCode (so it can
  never upgrade/remove the production Thetis HL2 install), `Thetis-Test` shortcut
  names, and `-dbid:HL2TEST` on the shortcuts so the test build keeps its own
  active settings database. Output file: `Thetis-Test-v<version>.x64.msi`
  (commit `d304a2df`)
- **Must be reverted (or identity restored) before merging to master** — the release
  installer keeps the production UpgradeCode, name, and `Thetis-HL2` folder

### 🟡 Phase 3 — verification *(in progress; blocked on "no sync")*

> Note: audio can't be injected into the RX chain via Voicemeeter/VAC — VAC input
> feeds the TX mic path only. The bench route is Thetis's RX wave playback, which
> replaces antenna I/Q at the head of the RX DSP chain (`ChannelMaster/pipe.c`,
> `xplaywave`, "IQ data" position), upstream of the fdv block.

**Common settings** (all steps): RX1 mode **DIGU**, ~3 kHz filter, NR/NR2/ANF/NB
and squelch **off** (`Thetis_VB-Audio_config.md` §7).

1. ✅ **Bench signal built** — `Tools/FreeDV/fdv700e_test_iq.wav` (~112 s, 700E via
   codec2's `freedv_tx`, ve9qrp sample). Regenerate any time with
   `Tools/FreeDV/make_fdv_test_iq.py` (see its README); also drops a ready-named
   `SDRQuickAudio.wav` copy. **Bug found & fixed in this step**: the generator's
   original analytic-signal (Hilbert) construction landed the signal in the *lower*
   sideband — wdsp's I/Q spectral sign convention is the conjugate of the textbook
   one. Fixed by writing `-Q` (`df591ac0`); confirmed on the panadapter afterward.
2. ✅ **Smoke test** — `Thetis-Test` installs and runs the FreeDV tab/checkbox
   without fault; `libcodec2.dll` loads fine.
3. 🟡 **Bench decode — in progress, currently blocked.** Signal plays and is
   visually/numerically confirmed correct on every axis checked so far, but the
   sync label sticks on **"no sync"** and never flips. Ruled out so far:
   - Sideband/position: DIGU confirmed, dial re-anchors correctly on replay,
     signal sits dead-centered in the 300–2700 Hz filter (~1500 Hz above dial,
     matching 700E's designed OFDM centre frequency) — confirmed via panadapter
     screenshots, not just the frequency math
   - Filter width: 300–2700 Hz (2400 Hz) is ample; 700E's actual occupied
     bandwidth is ~1500 Hz (21 carriers × ~71 Hz spacing, codec2 `ofdm_mode.c`),
     centred at 1500 Hz — no clipping
   - Level: signal peaks -85 dBFS against a -145 dBFS floor (60 dB), far more
     than needed for OFDM sync
   - ANF, NR/NR2/NR3, NB, SNB, MNF (manual notch), CTUN — all confirmed off via
     console screenshots (ANF/MNF were the leading suspects — a notch filter
     sitting on the OFDM carriers — but both were off)
   - AGC set to Slow; not yet tested with AGC fully off/fixed gain — **now the
     top suspect, see below**
   - Code audit: `fdv.c`'s resampler direction, RMS/gain-normalization math, and
     the `freedv_nin()`/`freedv_rx()` block loop were checked line-by-line against
     codec2's own reference `freedv_rx.c` CLI tool — structurally identical, no
     bug found by inspection
   - **Deeper code audit (this session)** — re-checked with fresh eyes, all
     confirmed correct, not the bug: `freedv_open(FREEDV_MODE_700E)` is right;
     the 48 kHz→8 kHz resample uses `create_resampleF` with the real channel
     rate read at runtime (`ch[channel].dsp_rate`, not hardcoded), not naive
     decimation; `freedv_nin()` is called before every `freedv_rx()` and the
     code correctly waits for a full `nin` before consuming, respecting
     codec2's variable-block-size contract; `xfdv()`'s post-AGC/post-NR
     placement in `RXA.c`'s chain is confirmed correct and its channel
     mono-extraction is byte-identical to `rnnr.c`'s established convention.
     One residual risk flagged, not yet a confirmed bug: `fdv.c` is the
     **only caller of `create_resampleF` anywhere in this codebase** — that
     decimate-by-6 path has never been exercised by any other working
     feature, so it's unverified-elsewhere code even though inspection found
     no bug in it.
   - **Runtime debug data captured (this session)** — used the new CAT-based
     remote tooling (see below) to trigger two fresh Quick-Play runs and pull
     `%TEMP%\fdv_debug.txt`. Finding: `rms` at `fdv.c`'s own AGC input swings
     by up to ~4.7× (≈13 dB) between *consecutive* 80 ms blocks (e.g. block
     3→4: 2687.8→12644.3 in run 1; 921.2→5816.3 in run 2) — well before
     `fdv.c`'s own smoothing gets a chance to act. A continuous, pause-free
     OFDM test signal shouldn't naturally swing like that; this is consistent
     with the channel AGC (upstream of `fdv.c`, still set to Slow) fighting a
     signal shape it isn't tuned for. `sync`/`snr` stay completely flat
     (`sync=0`, `snr=-5.0`) across every block in both runs — no partial
     correlation visible at all.
   - **Cross-check against freedv-gui (this session)** — a local checkout of
     upstream's `freedv-gui` (with vendored `codec2-1.2.0`) was used as a
     second, independent reference, and its own bundled sample WAVs as
     ground truth. Built `libcodec2`/`freedv_tx`/`freedv_rx` on Linux (no
     Windows box needed) and ran real decodes:
     - `freedv-gui/wav/ve9qrp_700e.wav` (upstream-shipped, pre-encoded 700E
       modem audio, same `ve9qrp` voice source, ~112 s) decodes cleanly
       against the reference `freedv_rx` — full sync, all 1405 frames, full
       899200 samples of real speech out. Added to the repo (gitignored) as
       a second bench file: `Tools/FreeDV/ve9qrp_700e_golden_test_iq.wav`,
       built via the existing `make_fdv_test_iq.py` from this file — play it
       through Quick-Play as an alternate to `fdv700e_test_iq.wav`.
     - Reproduced `fdv700e_test_iq.wav`'s own recipe exactly (`freedv_tx
       700E codec2/raw/ve9qrp.raw`, no `--clip`) and decoded that too —
       also clean, full sync throughout, RMS 4458/peak 16383. **This rules
       out "bad test file" as the bug**: the modem audio `fdv.c` is being
       fed is provably correct where `freedv_tx` produces it. It also
       confirms `FDV_TARGET_RMS_DB` (~4000 counts) is well-matched to this
       signal's real level (within ~1 dB), so the AGC *target* isn't
       mistuned either — the bug is downstream, somewhere in
       IQ → wdsp SSB demod → `fdv.c` resample/AGC/ring buffers →
       `freedv_rx`.
     - Read codec2's `ofdm_sync_search_shorts()` (`ofdm.c`, the function
       `freedv_rx()` actually calls while hunting for 700E sync): it
       normalises input as `rxbuf_in[j] / 32767.0f` and its own comment
       states **"Gain is not used here"** — the `gain` value `fdv.c`
       carefully computes only affects `ofdm_demod_shorts()`, which only
       runs *after* sync is already found. **`fdv.c`'s per-block AGC
       precision cannot be what's blocking initial sync acquisition**, as
       long as levels aren't clipped or near-zero. This tempers the
       AGC-off hypothesis below — still worth ruling out, but a null result
       there wouldn't be surprising — and re-weights suspicion toward
       something structural: discontinuities/dropped samples across
       `fdv.c`'s ring buffers, or the unverified `create_resampleF`
       decimate-by-6 path.
     - Diffed against `freedv-gui/src/pipeline/FreeDVReceiveStep.cpp`
       (upstream's own RX chain): it applies **no dynamic per-block AGC at
       all** — raw samples go straight into a FIFO at unity gain, relying on
       codec2's documented tolerance for "wide but not clipping" levels.
       `fdv.c` instead re-locks its gain target *every* ~80 ms `nin` block
       (`FDV_GAIN_SMOOTH` = 30% step). If that produces a real amplitude
       discontinuity at block boundaries, it lands inside
       `ofdm->rxbuf`'s multi-block sliding correlation window, which wants
       amplitude-consistent samples — not segments independently rescaled
       up to 13 dB apart, which is exactly the swing the runtime debug data
       above already captured. Plausible way `fdv.c`'s own AGC could be
       self-sabotaging sync even if the upstream channel AGC turns out
       innocent.
   - **AGC-off tested live (this session) — negative result.** Using the
     remote tooling against the live instance (192.168.2.12): confirmed
     mode DIGU / filter 300–2700 Hz, ran a baseline Quick-Play under the
     instance's then-current AGC (MEDIUM) — "no sync" throughout, matching
     the documented symptom. Set `agc set FIXED`, re-ran Quick-Play for
     ~12 s, polled `freedv status` three times — **"no sync" every time**.
     Restored AGC to MEDIUM afterward. This confirms the
     `ofdm_sync_search_shorts()` code reading above: AGC state does not
     gate initial sync acquisition. **AGC is no longer a live suspect at
     all** (not just demoted) — the bug is elsewhere in the chain.
   - **Still open, ranked** (updated this session): (1) dump `fdv.c`'s own
     resampler output (`a->rs_down_out`, the 8 kHz signal *before* the RMS
     normalizer) and diff it sample-for-sample (not just spectrally) against
     `Tools/FreeDV`'s known-good 8 kHz modem audio — now backed by an actual
     reference decode, not just inspection, to rule the untested
     `create_resampleF` path in or out directly; (2) freeze `fdv.c`'s gain
     after the first block (skip the `FDV_GAIN_SMOOTH` re-lock loop
     entirely) to match freedv-gui's no-AGC convention and test whether a
     constant per-session gain changes anything; (3) confirm the actual DSP
     processing rate (Setup → DSP → Options) matches the 48 kHz `fdv.c`
     assumes; (4) run the same bench file/tuning through the **external
     FreeDV desktop app** (via the VAC path, `Thetis_VB-Audio_config.md`
     §7) as a differential test — an independent decoder syncing on our
     signal would isolate the bug to Thetis's chain entirely, now less
     useful as a "bad synthetic file" check (already ruled out above) but
     still useful as a full-chain sanity check. Items (1) and (2) both need
     a `fdv.c` code change and a Windows rebuild before they're testable —
     the remote tooling alone can't make further progress on this list
     until one of those lands.
   - **New: remote testing tooling** (`Tools/thetis-ai-control`,
     `.claude/skills/thetis-control/SKILL.md`) — CAT commands `quickplay
     on|off|get` / `quickrec on|off|get` (revived orphaned `ZZQA`/`ZZQB`,
     previously implemented but never wired into the dispatch switch) and
     `freedv on|off|get` / `freedv status` (new `ZZDV`/`ZZDS`, reads
     `GetRXAFDVSync`/`GetRXAFDVSnr` — same calls `freedvStatusTimer_Tick`
     uses) let Quick-Play + sync/SNR be triggered and read entirely over the
     network, no one needing to sit at the Setup DSP tab. Steps 4-6 below can
     now be scripted once bench decode is unblocked.
4. ⬜ **Off-air capture** — next step in progress: quick-**Rec** a few minutes of
   live 14.236 MHz traffic (check qso.freedv.org first). Doubles as a
   more-realistic differential test signal (real channel effects) and the
   permanent regression file once decode is working.
5. ⬜ **Live decode**: 14.236 MHz DIGU. Ground truth: before enabling the
   checkbox, confirm the external FreeDV GUI app (VAC path, §7) syncs on the same
   signal. Note SNR at sync acquire/drop (700E should hold to ~1 dB)
6. ⬜ **Iterate on findings** (once sync is achieved):
   - decoded speech level — `FDV_SPEECH_GAIN` (0.30f, fdv.c) vs passthrough/SSB
     loudness
   - priming latency — ~125 ms buffer before decode engages; listen for swallowed
     first syllables; underrun resets priming → watch for passthrough/voice
     flapping on fades (fix would be sync-loss hysteresis, a Phase 4 item)
   - robustness — change DSP buffer size/rate, switch bands/modes, toggle checkbox
     rapidly, MOX cycles: no crash, sync recovers, run flag survives rebuilds
   - CPU load — checkbox on vs off (expect <1%); persistence — restart with
     checkbox on, still decoding

**Exit criteria**: bench decode + at least one live off-air decode with believable
SNR readings; findings fixed or recorded as Phase 4 work.

### ⬜ Phase 4 — prototype wrap-up

- Docs: FreeDV-native section in `Documentation/`, code_documentation regeneration
  (`fdv.c` needs a `CODE_OUTLINE.md` table row first), release-notes entry
- Revert the Thetis-Test installer identity (see above) so master's MSI remains the
  production installer
- Decision gate: merge to master as an experimental feature in the next release,
  or keep maturing on the branch

## Stage B — Full FDV mode *(future, after the prototype proves out)*

- ⬜ TX path: mirror block in `TXA.c` using `freedv_tx()` — mic audio → 8 kHz →
  modem audio → SSB modulator
- ⬜ Mode selection: 1600 / 700D / 700E (`SetRXAFDVMode`), RX2/subRX support
- ⬜ 700D/E text messages: callsign beacon TX + received-text display
- ⬜ Real UI: console mode button or info-bar sync light; entering FDV auto-applies the
  no-speech-processing TX rules (see `Documentation/Thetis_VB-Audio_config.md` §7)
- ⬜ Possible `DSPMode.FDV` enum member — the high-merge-cost step; only if Stage B
  ships for real
- ⬜ CAT/TCI exposure of FDV state

## Stage C — RADE neural mode *(external dependency)*

- ⬜ Watch upstream: David Rowe's RADE V2 C port is in progress (classical DSP ported
  as of May 2026; C library distribution planned after their stored-file test campaign
  validates the algorithms — see freedv.org)
- ⬜ When the C library ships: build it with a `build-codec2.yml`-style workflow and
  slot it into `fdv.c`'s mode handle — a dependency bump, not a research project
- Re-coding RADE ourselves was evaluated and rejected: months of duplicated effort and
  a validation problem only upstream's test campaign can solve

## Standing constraints

- Everything new lives in new files; hooks into shared files are single-line and
  `// W5TSU`-tagged (upstream merge survival — see `Project Files/Source/AGENTS.md`)
- Older Console files can have mixed CRLF/LF line endings (`radio.cs`) — check diff
  size after editing; rebuild byte-exact if an editor flattens them
- Windows-only build: CI (`gh workflow run build.yml --ref FreeDV -R
  W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`) verifies compile; behavior is verified on-air
- codec2 is LGPL-2.1, dynamically linked; its license ships beside the dll

## Reference

- FreeDV / RADE: <https://freedv.org>, <https://freedv.org/radio-autoencoder/>
- codec2: <https://github.com/drowe67/codec2>
- Activity: FreeDV Reporter <https://qso.freedv.org>
- Using the external FreeDV app with Thetis instead:
  `Documentation/Thetis_VB-Audio_config.md` §7
- Local `freedv-gui` checkout (`~/Development/freedv-gui`, vendored
  `codec2-1.2.0`) — used as a second reference implementation and source of
  known-good sample audio (`wav/ve9qrp_700e.wav` etc.); its bundled codec2
  can be built standalone on Linux (`cmake` + `make freedv_tx freedv_rx`) for
  fast ground-truth decodes without the Windows toolchain
