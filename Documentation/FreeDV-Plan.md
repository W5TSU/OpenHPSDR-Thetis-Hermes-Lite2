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

### 🟢 Phase 3 — verification *(bench + live decode both achieved; iteration on findings remains)*

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
3. ✅ **Bench decode — achieved 2026-08-08.** Signal plays and is
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
   - **Resampler-output dump added (this session, `3eb8fae0`)** — a third
     temporary diagnostic dump alongside `fdv_debug.txt`/
     `fdv_debug_audio.raw`: `fdv_debug_resamp.raw` captures `a->rs_down_out`
     (the 8 kHz signal exactly as `create_resampleF` produces it) as
     contiguous float32, written right after `xresampleF()` — before fdv's
     own RMS/AGC normalizer or the nin-block chunking touch it. Same
     reset-per-Quick-Play-session wiring via `ResetRXAFDVDebug()`, same
     150-call cap. **Still needs a Windows build + a Quick-Play run to
     produce data** — nothing has been diffed yet, this only adds the
     capture point.
   - **Frozen-gain experiment added, built, live-tested, and ruled out
     (this session)** — `56548b03` skipped the `FDV_GAIN_SMOOTH` re-lock
     entirely once `agc_seeded` was set, matching freedv-gui's own
     no-dynamic-AGC convention. Built via CI (`gh workflow run build.yml
     --ref FreeDV`), installed on the Windows box, and live-tested against
     the real instance (192.168.2.12): Quick-Play under AGC MEDIUM and
     again under AGC FIXED both still showed "no sync" across multiple
     `freedv status` polls over ~15s each. **Negative result — reverted in
     `653e3db2`**, `FDV_GAIN_SMOOTH` smoothing restored. Neither AGC
     precision nor AGC dynamics (frozen vs. re-locking) explain the "no
     sync" symptom; the bug is elsewhere.
   - **⚠️ Safety finding: `quickplay on` is TX-capable, not RX-only as
     documented (this session)** — while live-testing the frozen-gain
     build, the operator observed MOX engaging during Quick-Play. Root
     cause: `PlayFileViaWDSP` (`Console/clsAudioRecordPlayback.cs`, the
     function Quick Play calls) is shared with a genuine TX-audio-preview
     feature and contains `if (!_console.MOX && MoxOnPlayback) _console.MOX
     = true;`, and `MoxOnPlayback` **defaults to `true`** in this codebase.
     `thetisctl`'s `quickplay on` had no `--confirm-tx` gate at all before
     this was caught — every prior Quick-Play call in this project's
     testing history (including the "RF-free" framing used earlier in this
     document, and in the `thetis-control` skill) was unconfirmed TX, not
     confirmed RX-only. No RF is believed to have radiated this session —
     the antenna was on a dummy load, and the operator confirmed the PA/TX
     indicator didn't light even on manual TUN/2TON on this instance — but
     that's specific to this hardware setup, not a guarantee. Fixed in
     `f2ab8735`: `quickplay on` now requires `--confirm-tx` and auto-stops
     after `--hold` (default 15s), exactly like `ptt`/`tune`; `quickrec`
     was checked and confirmed to have no equivalent MOX side effect.
     Persisted to project memory (`quickplay-can-key-mox`) so this survives
     even outside this conversation. **Any future FreeDV bench-test
     `quickplay on` call must go through the full TX safety protocol** —
     dry-run first, explicit operator confirmation of the specific test in
     the current conversation, `--confirm-tx` only after that — not the
     fire-and-forget pattern used throughout Phase 3 above.
   - **Extended-stats debug logging (`57f8f029`) confirmed unable to
     diagnose the "still searching" case — a correction to that commit's
     own rationale.** Traced `freedv_get_modem_extended_stats()`'s dispatch
     for OFDM modes in `freedv_api.c`: it only `memcpy`s fresh data from
     `f->stats`, and `f->stats` is only populated by
     `ofdm_get_demod_stats()`, which `freedv_700.c` calls **exclusively
     inside the `if (sync_state == synced || trial)` branch**. While stuck
     in `search` state — our entire problem — `sync_metric`/`foff`/
     `rx_timing`/`clock_offset` never update at all, so the block-by-block
     dump this session added will show flat zeros for the whole capture
     regardless of whether the correlator is close to locking or nowhere
     near it. Left the logging in place (harmless, and instantly useful the
     moment sync is ever achieved even briefly), but it's not the diagnostic
     tool it was believed to be when added — the `fdv_debug_resamp.raw`
     diff is still the one live test that can see anything during `search`.
   - **CFO "dead zone" hypothesis — formed from source, then empirically
     refuted.** Re-read `ofdm_sync_search_stream()` (`ofdm.c`, the function
     700E's *voice*-mode search actually uses — not the burst-mode path
     read earlier in this project). Its coarse frequency search tests only
     three fixed candidates, `{-40, 0, +40}` Hz, against a ~180 ms
     correlation window (700E: `Ts=14ms`, `Tcp=6ms`, `Nc=21`); back-of-envelope
     phase-drift math over that window suggested a residual offset of just
     a few Hz between candidates could plausibly fail to lock at any of the
     three — a real, if narrow, structural "dead zone," and a clean
     candidate explanation not yet considered. **Tested directly and
     decisively, no radio needed**: built `libcodec2`/`freedv_rx` locally
     (already done this project) and swept a controlled, genuine
     single-sideband frequency shift (Hilbert-transform + `exp(j2πft)`,
     real part) into the known-good `ve9qrp_700e.raw` reference audio from
     −45 to +45 Hz in 5 Hz steps — **every single offset synced perfectly**
     (1405/1405 frames). Pushed to ±150 Hz and it still found sync. The
     hypothesis is refuted — codec2's real capture range is far more
     tolerant than the source-level math predicted, and nothing in our
     actual signal chain could plausibly produce anywhere near this much
     offset regardless. **CFO is now ruled out as decisively as AGC.**
   - **Virtual Audio Cable input investigated as an alternative test-signal
     path — ruled out, don't revisit.** Traced `xvacIN()`/`xvacOUT()` calls
     in `ChannelMaster/pipe.c`: `xvacIN` (a VAC's audio *feeding into*
     Thetis) is called exclusively inside the transmitter/mic block
     (`stream == inid(1,0)`, `case 0: // MIC data`) — there is no `xvacIN`
     anywhere in the RX chain. A VAC can only inject audio into Thetis's
     mic input → TX chain, never into RXA.c where `fdv.c` lives. Quick-Play
     remains the only mechanism in this codebase for injecting a controlled
     signal into the RX chain; VAC input cannot substitute for it. (VAC
     *output* — RX audio routed **out** to the external FreeDV app for a
     live differential test, step 5/§7 below — is the opposite direction
     and unaffected by this finding.)
   - **Clarified: "no output when Decode FreeDV is on" is likely expected
     behavior, not a new bug.** Operator observed normal RX audio going
     silent with the checkbox on. Root cause is probably two documented
     behaviors stacking, not silence: codec2's `freedv_rx()` echoes raw
     demod audio as "speech" when unsynced *unless* squelch is enabled
     (`freedv_api.c` doc comment: "useful for tuning FreeDV signals"), and
     `fdv.c` then scales whatever it receives by `FDV_SPEECH_GAIN = 0.30f`
     after two lossy 48k↔8k resample round-trips — likely real but very
     quiet audio, not literal zero output. Not chased further this session
     (downstream of, and lower priority than, the sync problem itself);
     worth a volume/meter check if it resurfaces. **The real proof sync is
     working is `freedv status` reporting `sync=1`, not audible output** —
     the CAT status read is ground truth from `ofdm->sync_state` directly,
     independent of the audio path entirely.
   - **New standing resource**: `sdr-for-engineers` (Claude Code skill,
     `~/.claude/skills/sdr-for-engineers/`) — built this session from
     *Software-Defined Radio for Engineers* (Collins/Getz/Pu/Wyglinski). Its
     Ch. 10 (OFDM: Schmidl & Cox, cyclic prefix, coarse/fine CFO) and Ch. 6/7
     (PLL structure, timing/carrier sync) directly informed the CFO
     dead-zone investigation above. Treat as a standing reference for any
     further sync-theory reasoning on this bug, not a one-off aside.
   - **Still-unresolved tooling gap**: no remote file access to the Windows
     box exists (`thetisctl` speaks only CAT/TCI, no SSH/SMB/etc.
     configured) — `fdv_debug_resamp.raw` and `fdv_debug.txt` currently
     require the operator to manually retrieve/relay their contents. Worth
     solving properly (e.g. a shared/synced folder) if this debugging phase
     continues much longer.
   - **Still open, ranked** (updated this session): (1) pull
     `fdv_debug_resamp.raw` off the Windows box (the dump point from
     `3eb8fae0` already landed; per the finding above, this is now the
     *only* live test left that can see anything about the "search" state)
     and diff it sample-for-sample (not just spectrally) against
     `Tools/FreeDV`'s known-good 8 kHz modem audio (e.g.
     `np.fromfile(path, dtype='<f4')`, scale ×32768 to compare against the
     int16 reference) to rule the untested `create_resampleF` path in or
     out directly; (2) confirm the actual DSP processing rate (Setup → DSP
     → Options) matches the 48 kHz `fdv.c` assumes; (3) run the same bench
     file/tuning through the **external FreeDV desktop app** (via the VAC
     *output* path, `Thetis_VB-Audio_config.md` §7) as a differential test
     — an independent decoder syncing on our signal would isolate the bug
     to Thetis's chain entirely, now less useful as a "bad synthetic file"
     or "bad frequency placement" check (both ruled out) but still useful
     as a full-chain sanity check. AGC precision, AGC dynamics, CFO, and bad
     test audio are now all ruled out — decisively, not just by inspection;
     item (1) is the only remaining hypothesis class (structural corruption
     of the sample stream itself — drops/duplicates/discontinuities) with
     a live test already wired up and ready to act on the moment the data
     can be retrieved.
   - **New: remote testing tooling** (`Tools/thetis-ai-control`,
     `.claude/skills/thetis-control/SKILL.md`) — CAT commands `quickplay
     on|off|get` (now TX-gated, see above) / `quickrec on|off|get` (revived
     orphaned `ZZQA`/`ZZQB`, previously implemented but never wired into
     the dispatch switch) and `freedv on|off|get` / `freedv status` (new
     `ZZDV`/`ZZDS`, reads
     `GetRXAFDVSync`/`GetRXAFDVSnr` — same calls `freedvStatusTimer_Tick`
     uses) let Quick-Play + sync/SNR be triggered and read entirely over the
     network, no one needing to sit at the Setup DSP tab. Steps 4-6 below can
     now be scripted once bench decode is unblocked.
   - **Tooling gap closed: SSH access to the Windows test box (this
     session)** — set up OpenSSH Server on 192.168.2.12 (`mark`, an
     Administrator account, key-based auth via
     `administrators_authorized_keys` + exact ACL — `SYSTEM:F` +
     `Administrators:F`, no inherited entries, or Windows silently ignores
     the file). `scp`/`ssh` now work directly (`ssh hl2winbox`, alias in
     `~/.ssh/config`); a `winps.sh` helper runs PowerShell remotely via
     `-EncodedCommand` (base64 UTF-16LE) to sidestep ssh→cmd.exe→powershell
     quoting — Windows OpenSSH's default shell is `cmd.exe`, not
     PowerShell, and naive quoting breaks silently or gets
     partially-reinterpreted. This unblocks direct retrieval of
     `%TEMP%\fdv_debug*.{txt,raw}` — no more manual relay — and can push
     files too (used below to fix the bench file).
   - **🎯 Root cause found for a *different*, previously-undiscovered bug:
     the wrong file was loaded in the Quick-Play slot (this session).**
     Pulled `C:\Users\mark\Music\Thetis\quickrecord\SDRQuickAudio.wav` via
     the new SSH access and inspected it directly: **mono, 16-bit PCM,
     8000 Hz, 49.8 s** — nothing like the properly-built bench file. The
     correct file (`Tools/FreeDV/fdv700e_test_iq.wav` /
     `ve9qrp_700e_golden_test_iq.wav`) is **stereo, 32-bit float, 48000 Hz,
     ~112 s** — a real analytic I/Q pair, matching what `xplaywave`
     (`ChannelMaster/pipe.c`, "IQ data" case) actually expects: it
     overwrites RX1's raw ADC-domain complex I/Q buffer in place, upstream
     of everything (mixer, SSB demod, `fdv.c`). A mono 8 kHz clip fed into
     that slot is read as if it were interleaved 48 kHz complex samples —
     every stage downstream (image-reject mixing by
     `(VFOAFreq*1e6) % sample_rate_rx1`, SSB demod, decimation) operates on
     nonsense. Confirmed quantitatively before fixing: normalized
     cross-correlation between `fdv_debug_resamp.raw` and the wrong
     source file's own samples never exceeded ~0.12 at any lag (a
     same-signal sanity check on the correlation method itself scored
     1.0000), i.e. no structural resemblance survived the chain — consistent
     with, not just consistent with but *explaining*, the total scrambling
     hypothesis. Copied the golden file over the wrong one via `scp`
     (`/Users/mark/Music/Thetis/quickrecord/SDRQuickAudio.wav`, 43,161,644
     bytes, confirmed via remote `Get-Item`). **Not yet reflected in the
     repo's own `SDRQuickAudio.wav` copy-on-generate step or any
     documentation warning** — if `make_fdv_test_iq.py` is ever re-run to
     regenerate the bench files, re-verify the live box's copy didn't get
     silently replaced by something else again; this class of bug (right
     file existing in the repo, wrong bytes on the actual test box) won't
     be visible from source alone.
   - **Live-tested with the corrected file — still "no sync."** Three more
     Quick-Play runs against the live instance (14236000 Hz / DIGU,
     30 s/15 s/20 s holds), monitored via repeated live `freedv status`
     CAT polls *during* playback (not just after) specifically to avoid
     trusting a stale post-hoc read. `quickplay get` confirmed genuinely
     "true" throughout each hold (ruling out a silent playback-start
     failure); every single live poll across all three runs still read
     "no sync." **The wrong-test-file bug was real, and fixing it was
     necessary, but it is not the (sole) explanation for the "no sync"
     symptom** — with a file already independently cross-validated via
     freedv-gui's own reference `freedv_rx` (full sync, 1405/1405 frames,
     step 3 above), Thetis's `fdv.c` still never syncs. The underlying
     bug is still somewhere in Thetis's own chain.
   - **New secondary bug found, not yet fixed: `fdv_debug.txt`/
     `fdv_debug_resamp.raw`/`fdv_debug_audio.raw` only ever capture the
     *first* Quick-Play session in a given Thetis process lifetime.**
     All three debug files came back byte-identical (same MD5) across all
     four Quick-Play runs this session, including the three run *after*
     the file fix — despite each run genuinely starting and stopping
     (`quickplay get` toggling true→false correctly each time, confirmed
     live). `ResetRXAFDVDebug()` (`console.cs`'s
     `ckQuickPlay_CheckedChanged`) should fire on every false→true
     transition and truncate+re-arm all three files, but evidently isn't
     taking effect on repeat sessions — file mtimes on the box stayed
     pinned to the very first run's timestamps throughout. Not yet
     root-caused (candidates worth checking first: `ckQuickPlay.Enabled` state
     across the `arp_PlayingingChanged`/`arp_RecordingChanged` handlers in
     `console.cs`, or a P/Invoke marshalling/threading issue on the
     `WDSP.ResetRXAFDVDebug()` call itself) — flagged here so the *next*
     `fdv_debug_resamp.raw` pull isn't trusted at face value without first
     confirming (via remote mtime check, now trivial with SSH access) that
     it's actually fresh. Practical workaround until fixed: restart Thetis
     between Quick-Play test sessions to guarantee a clean capture.
   - **Revised next step**: the sample-for-sample `fdv_debug_resamp.raw`
     diff (previously blocked on file access, now unblocked) is still the
     right next move, but needs a *fresh* capture — restart Thetis first,
     given the debug-log staleness bug above, then run one Quick-Play
     session and immediately pull all three debug files before running
     anything else.
   - **🎯🎯 Root cause found: Quick-Play was silently non-functional this
     entire debugging effort — every "no sync" result up to this point is
     void (this session, follow-up).** While chasing the debug-log
     staleness bug above, discovered the real explanation via the
     independent instrumentation added along the way (`fdv_debug_events.txt`,
     since removed): `ckQuickPlay.Enabled` is **`False` by default** —
     `console.resx` has an explicit `<data name="ckQuickPlay.Enabled">
     <value>False</value></data>` entry — and `ckQuickPlay_CheckedChanged`'s
     very first line is `if (!ckQuickPlay.Enabled) return;`. With `Enabled`
     false, every "Quick Play on" — whether from the physical GUI button,
     the Andromeda-style `OtherButtonId.PLAY` button-bar action
     (`DoOtherButtonAction`, also gated on the same `Enabled` check), or
     CAT's `ZZQA1` — silently no-ops: the `Checked` property still toggles
     (a separate flag CAT's `quickplay get` reads, which is why every prior
     poll looked "normal"), but `ResetRXAFDVDebug()` and
     `ARP.PlayFileViaWDSP()` — the actual file-load-and-inject call — never
     run. No audio was ever injected, by any test, this entire session,
     until this was found. The only thing that flips `Enabled` to `true` is
     completing a Quick-**Rec** session at least once
     (`arp_RecordingChanged`'s `id=="quick"` branch, `recording` going
     true→false) — apparently by design, though the intent (if any) isn't
     documented anywhere in the surrounding code. Confirmed practically:
     ran `quickrec on` then `quickrec off` (not TX-capable, no antenna/PA
     involvement) once via CAT — `Enabled` flipped `true` immediately and
     stayed true for the rest of the process's life. **Workaround for all
     future testing**: run one harmless Quick-Rec on/off cycle after every
     Thetis (re)start, before the first Quick-Play attempt. **Real fix
     still needed** (tracked as Phase 4 cleanup, not yet done): either flip
     the resx default to `true`, or add an explicit enable at whatever
     point Quick-Play is actually meant to become available (radio power-on
     completing seems the more sensible trigger than an unrelated
     recording feature) — needs a decision on intended UX, not just a
     bug fix.
   - **First genuine signal-in-chain confirmation of the whole project,
     immediately after the fix.** With `Enabled` unstuck, ran Quick-Play
     again (same bench file, 14236000 Hz / DIGU, 20 s hold) and — for the
     first time this entire debugging effort — the operator visually
     confirmed the expected OFDM "picket fence" signal on the panadapter/
     waterfall during playback. **`freedv status` still read "no sync"
     throughout, live-polled the whole time.** This is a clean, valuable
     isolation: injection → mixing → SSB demod → display are all now
     independently confirmed correct with a real, cross-validated signal;
     the bug is now known to be specifically in `fdv.c`'s
     resample/AGC/ring-buffer handling or codec2's `freedv_rx()`
     sync/acquisition itself, not anywhere upstream of it. Every hypothesis
     ruled out in earlier sessions (AGC precision, AGC dynamics, CFO, bad
     test file, wrong Quick-Play file) remains ruled out — this doesn't
     reopen any of them, it just confirms the signal reaching `fdv.c` now
     really is the intended one.
   - **First fresh `fdv_debug_resamp.raw` capture of the project — too
     short to be conclusive, cap raised (`e2ecd8c6`).** Pulled all three
     debug files immediately after the above run; confirmed genuinely
     fresh (new rms/sample values, current mtime — the staleness bug above
     wasn't hit this time). But `fdv_debug_resamp.raw` was only 6356 bytes
     = 1589 float32 samples ≈ **0.2 s** of real 8 kHz audio — the
     `fdv_dbg_resamp_count < 150` cap was tuned assuming a much lower
     samples-per-call rate than what `create_resampleF` actually produces
     (~10.7 samples/call at `dsp_size=64`, 48k→8k). A normalized
     cross-correlation against upstream's known-good
     `freedv-gui/wav/ve9qrp_700e.wav` (mono/8kHz/112s, the same reference
     independently confirmed to sync perfectly against `freedv_rx` earlier
     in this project) came back weak (~0.14) — but with only ~1600 samples
     to align against a 112 s reference, this result isn't trustworthy
     either way (the self-match sanity check that validated this exact
     correlation method earlier in the session doesn't rule out a
     short-window false negative here). Raised the cap to 4000 calls
     (~5 s of real audio) — not yet re-tested with the raised cap.
4. ✅ **Two more bugs found and fixed on the way to the fix above — the
   sample-for-sample diff turned out unnecessary once these were found.**
   - **`Enabled` guard blocked every Quick-Play attempt this entire
     debugging effort — worked around, real fix still pending.** See the
     item directly above for the discovery. Practical result: `quickrec
     on` then `quickrec off` once per Thetis (re)start reliably unsticks
     `ckQuickPlay.Enabled`.
   - **The Quick-Rec workaround silently destroys the bench file — must
     re-copy it every time before testing.** Quick-Rec and Quick-Play
     share the exact same file path
     (`Music\Thetis\quickrecord\SDRQuickAudio.wav`) — the workaround above
     records over and replaces the golden 43 MB bench file with a
     near-instant scratch recording of live RX (111,148 bytes observed,
     ~0.29 s). This produced a second, completely misleading symptom:
     after "fixing" `Enabled`, Quick-Play would key MOX for well under a
     second and drop — not a codec2/fdv.c bug at all, just testing against
     the wrong (tiny, garbage) file. Root-caused via one more targeted
     instrumentation pass (`ReadBuffer`'s short-read/EOF path,
     `a8c9a3aa`): logged `streamLen=111148` against an expected ~43M,
     immediately pointing at the file itself rather than the parser.
     **Standing test procedure**: after every `quickrec on`/`off`
     workaround cycle, re-`scp` the golden file
     (`Tools/FreeDV/ve9qrp_700e_golden_test_iq.wav`) back over
     `SDRQuickAudio.wav` on the box *before* the next Quick-Play test —
     forgetting this step silently invalidates the run.
   - With both worked around (Enabled unstuck, correct file back in
     place), **Quick-Play finally ran uninterrupted for its full
     duration, and `freedv status` reported real, stable sync**: `SYNC
     SNR 11.8 dB` and holding (11.3–12.9 dB across ten consecutive live
     polls over the full 20 s hold). Confirmed genuine (not a fluke) via
     the underlying codec2 state in the same `fdv_debug.txt` capture:
     `sync=1` from block 2 onward, `sync_metric` climbing 0.436 → 0.521 →
     0.559, `foff` a tiny, realistic 0.1–0.2 Hz, `rx_timing` converged,
     `clock_offset=0.0` — every internal diagnostic looks exactly like a
     properly-locked OFDM demodulator, not a marginal/lucky read.
   - **So the *original* "no sync" bug never actually existed in `fdv.c` or
     codec2.** Every hypothesis investigated and ruled out across this
     entire multi-session effort before this point (AGC precision, AGC
     dynamics, CFO, bad test file content, wrong Quick-Play file format)
     was correctly ruled out — none of them were ever the problem. The
     real blocker, this whole time, was that **Quick-Play itself was
     never successfully completing a full-duration test run** — first
     because `Enabled` silently no-op'd it entirely, then because the
     workaround for that silently corrupted the one file it needed to
     play. The lesson for next time: when a live test's *result* is
     suspicious for many sessions in a row despite every plausible
     internal-logic hypothesis being ruled out, question whether the test
     itself is actually running as intended before going deeper into the
     code under test — `Enabled`'s state and the file's actual bytes on
     disk were both directly, trivially checkable the whole time and
     would have shortened this significantly.
   - **A second, real `fdv.c` bug found and fixed once genuine signal
     finally reached the decoder (`15fe65c7`).** With the two blockers
     above worked around, bench decode was *intermittent* — reproduced
     2 clean syncs and 2 total failures across four back-to-back runs.
     Both failures shared an identical signature in the new
     `fdv_debug_nin.txt` trace: `freedv_nin()` returned **0** immediately
     after a sync gain/loss transition and never recovered for the rest
     of the session, while `demod_ring` kept filling, unconsumed, toward
     its capacity. Root cause: codec2's `ofdm_demod()` legitimately sets
     `nin=0` when its own internal `rxbuf` already has enough buffered
     samples for the next frame — signalling "call `freedv_rx()` again
     right now with zero new samples to drain me," not "stop." `fdv.c`'s
     modem-block loop required `nin > 0`, misreading that signal as
     terminal and permanently stalling. Confirmed against
     `freedv-gui/src/pipeline/FreeDVReceiveStep.cpp` (the actual reference
     RX loop, not just prior audits of `freedv_rx.c`): it has **no
     equivalent guard at all** — it simply checks a FIFO has `>= nin`
     bytes available (trivially true for `nin=0`) and keeps calling the
     modem. Fixed by changing the loop condition to `nin >= 0`, skipping
     the RMS/AGC normalisation block for the `nin==0` case (it would
     divide by zero), and adding a bounded safety cap (16 consecutive
     `nin==0` iterations) since nothing in the API contract actually
     guarantees the state self-resolves quickly, even though every case
     observed took only 1–2 iterations. **Verified**: 3 consecutive clean
     runs post-fix, 36/36 live `freedv status` polls synced (up from a
     ~50% failure rate immediately before the fix) — SNR held steady
     11–15 dB throughout every run.
5. ✅ **Off-air capture — done 2026-08-08, via the new `freedv-reporter watch`
   (Stage D).** Ran it with `--tci` against the live instance; it correctly
   auto-tuned RX1 to 14.236 MHz on every transmit-start over a real 10-minute
   ongoing QSO (ZL2MQ, W4MLN, VK4GRA, JH2WTQ — confirmed live by the reporter
   itself, not inferred). Quick-Rec'd ~2 minutes of it:
   `Tools/FreeDV/offair_14236000_RADEV1_20260808.wav` (local only, matches
   the existing `*.wav`/`.gitignore` convention for bench audio — not
   pushed). **Caveat**: this traffic was FreeDV's **RADE V1** mode, not
   700E — codec2 has no RADE support yet (Stage C, still upstream-blocked),
   so this file can't validate today's decode path and a spectral-only
   sanity check on it is inconclusive (RADE's waveform is broadband/
   noise-like by design, unlike 700E's clean OFDM comb — the same peak-
   frequency-drift heuristic that flagged the earlier wrong-file bug isn't a
   useful signal here). Real value: (a) proof `freedv-reporter watch`'s
   auto-tune genuinely works against live, moving traffic, not just the
   static bench file; (b) a real, dated regression sample ready the moment
   RADE support lands in codec2. A **700E** off-air capture (the mode this
   branch actually decodes) is still open — re-run the same watch/quick-rec
   combo whenever the reporter shows 700E activity specifically (its
   `mode` field distinguishes this, e.g. `KG7FMN` was seen on 700E during
   this session's live tests, per `internal/freedvreporter`'s output).

   **Re-attempted 2026-08-15, still open.** Ran a bounded version of the same
   watch/quick-rec combo (`watch_700e_bounded.sh`, a throwaway wrapper around
   `thetisctl freedv-reporter watch`, not committed — checks every 60s, gives
   up after 50 checks with no hit, auto-restores the Quick Play slot's backup
   either way) for ~80 minutes. Result: **27 transmissions tracked, 9 distinct
   callsigns, 100% RADE V1, zero 700E** — this calling frequency/window is
   currently saturated with RADE V1 traffic specifically, not evidence 700E
   itself is inactive or broken, just bad timing again. Still open; worth
   re-running at a different time of day or explicitly widening beyond
   14.236 MHz if RADE V1 keeps dominating that exact frequency.
6. ✅ **Live decode — done 2026-08-15, via a real HackRF positive control, not
   opportunistic off-air traffic.** `Tools/FreeDV/tx_700e_hackrf.grc` (new)
   transmits the same known-good `fdv700e_test_iq.wav` over real RF instead of
   Quick Play's direct RX-chain injection — genuinely exercises the HL2's RF
   front end for the first time in this branch's history. Three passes at
   14.236 MHz DIGU, over the air (licensed operator present, ID'd each pass):
   - **Pass 1, 0 dB VGA gain**: no sync — expected, deliberately minimal
     starting power (Part 97 practice).
   - **Pass 2, 20 dB VGA gain**: still no sync, but this time the operator
     confirmed **the signal was visible on Thetis's panadapter with no visible
     modulation** — a carrier, not the OFDM comb. Root-caused before a third
     blind attempt: the wav's `-50 dBFS` peak level was set for Quick-Play's
     direct float-sample injection (no real DAC involved there) — HackRF's TX
     DAC is 8-bit, one quantization step is `1/128` = `-42 dBFS`, so `-50 dBFS`
     sits *below* the DAC's own quantization noise floor. The OFDM structure
     was being crushed to quantization noise before the signal ever left the
     radio. Confirmed by directly inspecting the wav's actual sample
     statistics (peak 0.00302, matching the predicted -50.4 dBFS exactly).
   - **Fix**: added a `blocks_multiply_const_vxx` (166x ≈ +44.4 dB) to
     `tx_700e_hackrf.grc` between the I/Q assembly and the resampler — a
     flowgraph-side digital gain correction, not a wav-file change (the wav's
     level stays correct for Quick-Play's own, unrelated use).
   - **Pass 3, same 20 dB VGA gain, gain-fix applied**: **SYNC**, SNR
     2.1 dB → (one transient drop, consistent with the known ~125 ms priming/
     underrun-resets-priming behavior noted in step 7 below) → 12.7 dB →
     14.1 dB, climbing and holding through to the end of the transmission.
     First-ever confirmed 700E decode through Thetis's actual RX chain
     (antenna → HL2 → ChannelMaster → `fdv.c`), not a synthetic injection —
     **but only after the operator found it on LSB/DIGL, not the expected
     USB/DIGU.**
   - **Sideband bug, found from that observation**: `make_fdv_test_iq.py`'s
     own conjugate correction (`-iq.imag`, its comment: "a straight analytic
     signal displays in the lower sideband, so write the conjugate to land
     USB") was calibrated and verified specifically for Quick-Play's direct
     software injection point — it says nothing about a real hardware TX/RX
     path. Routing the same wav through a real HackRF TX up-converter and
     back in through the HL2's own RX front end adds two more independent
     mixer stages, and evidently one of them flips the sideband again on top
     of Quick-Play's already-applied correction. **Fix**: added a
     `blocks_conjugate_cc` in `tx_700e_hackrf.grc`, right after the I/Q
     assembly, to cancel that extra inversion back out — chosen over just
     switching to LSB/DIGL to keep this flowgraph consistent with every other
     place in the repo that assumes DIGU (README, SKILL.md,
     `make_fdv_test_iq.py`'s own docstring), rather than carrying a
     HackRF-TX-specific exception.
   - **Pass 4, same 20 dB VGA gain, sideband fix applied, Thetis back on
     DIGU**: **SYNC held for 5 consecutive polls (~75 s)**, SNR 2.4–4.4 dB,
     dropping only as the transmission itself ended. Confirms the fix and
     gives a second, independent, correctly-configured confirmation of live
     700E decode through the real RX chain.

   Ground truth cross-check against the external FreeDV GUI app (VAC path,
   §7) was not done this session — the HackRF positive-control result above
   is arguably stronger ground truth already (a signal proven, byte-for-byte,
   to sync via Quick-Play, now also proven to sync over real RF on the
   correct sideband), but the VAC cross-check remains open if wanted as an
   independent second confirmation.
7. 🟢 **Iterate on findings** (once sync is achieved) — run against `hl2winbox`
   (`git:1c185f14`), 2026-08-16:
   - ✅ **Robustness** — 8 rapid `freedv on`/`off` cycles, then mode switches
     (DIGU→USB→LSB→DIGL→DIGU) and band switches (20→40→20) all while decode
     was enabled, all via CAT. Same PID throughout, every `freedv status`
     call answered normally, nothing hung. MOX cycling (3× real PTT on/off,
     2 s hold, decode enabled, no mic audio) also clean — `TX: false`
     confirmed after each, same PID after. DSP buffer size/sample rate
     changes are Setup-tab-only (no CAT/TCI hook) — not exercised remotely,
     still open if wanted.
   - 🟡 **CPU load** — 10×1 s `typeperf` samples of Thetis.exe's
     `% Processor Time`, freedv off vs on: 95.96% mean off, 97.20% mean on
     (+1.2 points), inside the ~9-point sample-to-sample jitter this counter
     already shows on its own (this box's baseline DSP/audio threads already
     run near a full core). No statistically meaningful delta — consistent
     with "<1%," but the noise floor is wide enough that a small real cost
     could be hiding in it.
   - ✅ **Priming latency / swallowed syllables** — captured decoded RX audio
     via `tci rx-audio capture` alongside a live `freedv status` poll loop
     during a real Quick-Play run of the golden 700E bench file. Sync
     reported on the very first 1 s poll (SNR 12.6 dB); the capture shows
     ~1.2 s of near-silence (background/pre-content), then a **smooth,
     natural voice attack** (10 ms-resolution RMS ramps from -90 dBFS to a
     syllable peak over ~30 ms, no hard truncation) — no evidence of a
     swallowed first syllable. Consistent with sync + the ~125 ms priming
     buffer completing well under a second; the visible lead-in is most
     likely the bench speech clip's own leading pause, not decode latency.
   - 🟡 **Decoded speech level vs passthrough** — same capture, freedv-on
     decoded segment vs a freedv-off passthrough capture of the identical
     file/signal chain: decoded speech **-55.6 dBFS RMS / -31.5 dBFS peak**
     vs raw modem passthrough **-27.1 dBFS RMS / -20.2 dBFS peak** — decoded
     audio is **~28.5 dB quieter on average (≈11 dB quieter at peaks)** than
     what the operator hears while tuning. Envelope std/mean confirms the
     character difference too (0.761 decoded/bursty vs 0.024
     passthrough/flat — same discriminator used for the earlier RADE V1
     local-speaker diagnosis). This is a real, likely-annoying UX gap:
     `FDV_SPEECH_GAIN` (0.30f) leaves decoded audio much softer than the
     signal an operator was just listening to — worth a Phase 4 revisit
     (raise the gain, or add output-level normalization/AGC on the decoded
     stage) rather than leaving operators to ride the volume knob every time
     sync engages/drops.
   - 🔴 **Persistence — negative, and re-discovered a standing bug along the
     way.** First attempt: `quickplay on` reported `true` and held for the
     full hold but **`freedv status` never synced** — traced via the box's
     own `fdv_debug_events.txt` to the **same `ckQuickPlay.Enabled`-stuck-false
     bug already documented above** (last tripped by this session's own
     `quickrec on/off` cycles during MOX-adjacent testing); the CAT setter
     flips `Checked` but the handler bails out (`if (!ckQuickPlay.Enabled)
     return;`) so nothing actually plays. Confirmed no fix has landed for
     this yet. Re-ran the `quickrec on/off` unstick workaround, re-`scp`'d
     the golden file back over the now-clobbered `SDRQuickAudio.wav` (the
     workaround overwrites it with a ~0.4 s scratch capture, exactly as
     documented above), and the retest synced cleanly for the full 30 s
     hold (SNR 11–14 dB) — so this was a test-environment gotcha, not a new
     regression. Then, for the actual persistence check: set `freedv on`,
     killed and relaunched Thetis.exe on `hl2winbox` (a plain `Stop-Process
     -Force` had to be replaced with a scheduled-task-based relaunch —
     `Start-Process` over SSH crashes immediately, .NET exit code
     `0xE0434352`/unhandled CLR exception, because an SSH/service-context
     process has no interactive desktop or audio session; injecting via a
     temporary interactive-logon scheduled task worked). **Result: `freedv`
     read back `false` after the restart** — the checkbox did not come back
     checked, contrary to Phase 2's "(auto-persisted)" note. Caveat: the
     first kill was a hard `Stop-Process -Force`, so a follow-up attempt to
     retest via a graceful close (`CloseMainWindow()`) was tried instead —
     it found no valid main-window handle (empty `MainWindowTitle`, never
     exited within 15 s) and couldn't be made to trigger a clean shutdown
     remotely, so a true graceful-exit persistence test remained unconfirmed
     at the time. Practical read: `chkFreeDVDecode_CheckedChanged`
     (`setup.cs`) only sets `RXAFDVRun` in memory with no visible immediate
     settings write.

     **Resolved (2026-08-16, later same day): confirmed via a real graceful
     restart, done by the operator at the physical/RDP console, not remote
     scripting.** Set `freedv on` + `power on`, operator closed Thetis
     normally and relaunched it (new PID, same `git:d93e3bb0` build,
     confirmed via `Get-Process`/CAT version) — **both `freedv` and `power`
     read back `false`**, the same negative result as the earlier
     forced-kill test. Notably, **VFO frequency/mode *did* survive this
     restart** (14236000 DIGU carried over, unlike the forced-kill test
     where it reverted to the 14074000 default) — so this isn't a
     wholesale "nothing persists" issue, it's specific: freq/mode go
     through a real save path, `freedv`/`power` don't, regardless of
     whether the restart is clean or forced. Two independent restarts
     (one forced, one graceful) now agree: **an operator hitting any
     Thetis restart today will find FreeDV decode and radio power both
     off afterward**, even if they were on before. A real fix (persisting
     these like every other Setup checkbox) is Phase 4/5 follow-up work,
     not done here. Box left with `freedv`/`power` back on for continued
     testing.

**Exit criteria**: bench decode + at least one live off-air decode with believable
SNR readings — **met** (2026-08-15 HackRF positive control). Iteration findings
above are recorded; the two 🔴/🟡 items (persistence, speech level) are real
open items for Phase 4, not blockers on Phase 3 itself.

### ✅ Two of item 7's findings fixed and live-verified, same day (2026-08-16)

- **`8c1f07b0` — the `ckQuickPlay.Enabled`-stuck-false race, fixed.** Root
  cause: `arp_PlayingingChanged`/`arp_RecordingChanged`'s "quick" branches
  trusted their async playing/recording callback unconditionally. A fast
  on/off cycle (exactly what CAT scripting does) could have the "started"
  callback for an already-stopped session arrive late and re-disable
  `ckQuickPlay` with no matching "stopped" event ever following. Fixed by
  guarding both branches on the checkbox's own current `Checked` state, so a
  stale event can't override a stop that's already happened. **This is a
  distinct bug from the standing `console.resx`-default-`Enabled=False`-at-
  startup issue (`dce3fccf`/`359b44f5`) — that one is still open, still
  needs the `quickrec on/off` kick once per Thetis session.** Verified via
  the new debug logging added to the "quick" branches: 5 rapid `quickrec
  on`/`off` cycles back-to-back all paired cleanly, ending `Enabled=true`,
  followed by a real Quick-Play session running its full duration and
  syncing with no workaround needed.
- **`c8033819` — `FDV_SPEECH_GAIN` 0.30f → 0.75f, fixed and re-measured.**
  Redid the same passthrough-vs-decoded capture methodology from item 7
  against the new build: decoded speech now **-30.6 dBFS RMS / -6.0 dBFS
  peak** (was -55.6 / -31.5 dBFS) vs passthrough's -27.1 / -20.2 dBFS —
  RMS now within 3.5 dB of passthrough (the "why did it just go quiet" gap
  is closed) and peaks now ~14 dB *louder* than passthrough peaks. No
  clipping (checked directly: zero samples above 0.5 out of a 0.499 peak).
  The improvement is larger than the raw 0.30→0.75 gain-ratio math alone
  predicts (~8 dB expected, ~25 dB observed) — `fdv.c`'s own constant isn't
  the whole story once the signal passes through the rest of the RXA output
  chain, so the commit's own "~-2.5 dBFS ceiling" estimate was too narrow a
  model; the measured numbers are ground truth over that math. Verified
  against one test signal (the golden 700E bench file) only — not a final
  tuned value, worth revisiting if real on-air signals land differently.
- Both deployed via the established safe path (CI `workflow_dispatch` on
  `FreeDV` → MSI admin-extract → `robocopy` sync, not a real install) and
  verified against a fresh Thetis.exe relaunch on `hl2winbox`
  (`git:d93e3bb0`). Relaunching remotely needed a workaround of its own:
  Windows OpenSSH's session has no interactive desktop/audio, so a direct
  `Start-Process` crashes immediately (.NET exit code `0xE0434352`); a
  temporary interactive-logon scheduled task (`Register-ScheduledTask
  -Principal (New-ScheduledTaskPrincipal -LogonType Interactive) ->
  Start-ScheduledTask -> Unregister-ScheduledTask`) launches into the real
  session cleanly. Also confirmed **`power on` doesn't survive a restart
  either** (same pattern as `freedv`'s checkbox from item 7 above) — the
  radio engine came up off, silently producing "no sync"/0-sample TCI
  captures until noticed and powered back on; worth remembering as a
  first-check next time a fresh relaunch looks broken.

### 🟢 Phase 4 — prototype wrap-up

- ✅ **Decision gate resolved (2026-08-16): keep maturing on the branch, not
  merging to master yet.** Rationale: the persistence bugs (`freedv`/`power`
  checkboxes don't survive a restart — item 7 above) are still open, RADE V1
  hasn't synced against real off-air traffic yet (only a known-good HackRF
  signal), and `FDV_SPEECH_GAIN=0.75f` is verified against one test signal,
  not re-checked against real on-air variety. None of these are blocking for
  *further branch work*, but they're the kind of thing worth closing before
  calling this release-quality. Revisit this gate once those settle.
- ⬜ **Revert the Thetis-Test installer identity — deferred, not done.**
  Only actually needed right before a real merge to master (it's what keeps
  `Thetis-Test` side-by-side-installable on `hl2winbox` without touching the
  production install in the meantime) — doing it now would break the exact
  test setup this whole project has relied on. Do this as the last step
  before merging, not before.
- ⬜ **Release-notes entry — deferred, not done.** Only relevant at the
  actual release/tag step; nothing to write until there's a release this
  ships in.
- ✅ Docs: FreeDV-native section in `Documentation/` — done via
  `Documentation/FreeDV-User-Guide.md` (tech-writer agent, 2026-08-16,
  `7d19cb6d`/`6eb18ede`) — 700E vs RADE V1 explainer + how to use it + known
  issues.
- ✅ **code_documentation regeneration — done (2026-08-16, `53d65ab2`).**
  Graph rebuilt to pick up the 8 new source files added since the last
  regen (`fdv.c`/`.h`, `radae.c`/`.h`, `radae_micdsp.c`/`.h`,
  `r8brain_wrap.cpp`/`.h`); `CODE_OUTLINE.md` got a new "Digital voice
  (FreeDV)" wdsp subsection plus three new ChannelMaster rows; all 285
  per-file pages regenerated (up from 281). Verified: no broken links, no
  stale `docs/` refs, every outlined filename exists on disk.

## Stage B — Full FDV mode *(future, after the prototype proves out)*

- 🟡 TX path: mirror block in `TXA.c` using `freedv_tx()` — mic audio → 8 kHz →
  modem audio → SSB modulator. **Encoder built, wired, verified via real PTT
  tests, and a graceful-drain MOX/PTT arbiter added, 2026-08-18** — see the
  dated entries below. A no-RF loopback bridge (`ZZEG`) was built 2026-08-24
  to test decode confirmation without RF, but is not yet deployed/verified
  (`hl2winbox` offline) — still open until that runs for real.
- ⬜ Mode selection: 1600 / 700D / 700E (`SetRXAFDVMode`), RX2/subRX support
- ⬜ 700D/E text messages: callsign beacon TX + received-text display
- ⬜ Real UI: console mode button or info-bar sync light; entering FDV auto-applies the
  no-speech-processing TX rules (see `Documentation/Thetis_VB-Audio_config.md` §7)
- ⬜ Possible `DSPMode.FDV` enum member — the high-merge-cost step; only if Stage B
  ships for real
- ⬜ CAT/TCI exposure of FDV state

### 🟡 700E TX encoder built and wired — inert, no arming path yet (`2e295647`, 2026-08-18)

Unlike RADE V1 TX (which had a fully-built encoder sitting unwired from the
sv1eia port), 700E had **zero TX scaffolding** — this is genuinely new code,
not plumbing existing pieces together. Lives in wdsp's `fdv.c`/`fdv.h`
(alongside 700E RX decode, not in ChannelMaster where RADE V1 TX lives),
mirroring `fdv.c`'s existing RX architecture in the opposite direction:

- New `FDVTX` type: mic audio → `speech_rate` → accumulated into
  `freedv_get_n_speech_samples()`-sized blocks → `freedv_tx()` →
  `freedv_get_n_nom_modem_samples()` of modem audio → resampled to dsp rate →
  drained into the output buffer, using the same ring-buffer-plus-smoothed-AGC
  bridging pattern the RX side already uses for the dsp-block-size ↔
  modem-fixed-block-size mismatch. **Silence (never raw mic) on underrun** —
  the same never-leak-live-audio principle RADE V1 TX's own silence-on-underrun
  already follows.
- Wired into `TXA.c`'s full per-channel lifecycle (create/destroy/flush/
  `setDSPSamplerate`/`setDSPBuffsize`) and hooked into `xtxa()` right after the
  input resampler — before mic gain/panel/compressor/EQ/CESSB/ALC. When armed,
  it replaces mic audio in place; the modem waveform then flows through the
  rest of the normal TX chain unmodified.
- `SetTXAFDVRun`/`GetTXAFDVRun` enable control, mirroring `SetRXAFDVRun`.
  **Inert by default** (`run=0`) and nothing calls it yet.

**Deliberately left open, same session-split precedent as RADE V1 TX**: no
console/CAT wiring (no way for an operator to arm this), no MOX/PTT arbiter
(700E has no EOO-style burst designed yet, so this may end up simpler than
RADE V1's arbiter, or may not — undecided), and no test path — RADE V1 got a
pre-existing loopback bridge for free from the sv1eia port; 700E has no
equivalent, so testing this will need a different approach, probably a TCI
TX-audio capture plus spectral/waveform sanity check rather than a loopback.

**Open architectural question, shared with RADE V1 TX, not resolved either
place**: both features now inject a digital modem waveform early in the TX
chain and let it flow through the normal analog-voice processing
(compressor/EQ/CESSB/ALC) on the way out. Whether that dynamics processing
degrades a digital signal enough to matter is untested for either feature —
worth a real check before either one is considered on-air-ready, not just a
theoretical concern to note and move past.

**Verified tonight**: clean first-pass compile (new code, could easily have
had errors), deployed to `hl2winbox`, and a real PTT regression check (short
silent key/unkey cycle) confirmed normal TX operation is unaffected by the
new hook sitting inert in the hot path.

### 🟢 700E TX encoder proven structurally sound via real PTT test (`8cb9e3b8`, 2026-08-18)

Added CAT arming (**`ZZEF`**, mirrors `ZZDV`'s RX-side shape) plus a temporary
raw-audio dump of the encoder's modem output, aimed at a safe, zero-PTT
verification path — same idea as RADE V1 TX's loopback bridge.

**Correction, same session**: the zero-PTT plan didn't work, and the reason
is a real architectural difference from RADE V1 TX worth recording. The
assumption was "`xtxa()` runs continuously regardless of MOX state, confirmed
earlier via the mic capture path" — true for RADE V1 TX, which taps
ChannelMaster's raw mic buffer *before* wdsp runs at all, so it sees audio
whether or not the radio is keyed. **Not true for 700E TX**: it lives inside
wdsp's own `TXA` chain, gated by `fexchange0()`'s `ch[channel].exchange`
check, which `SetChannelState()` only flips on at a real MOX key-down
(`console.cs`'s `chkMOX_CheckedChanged2` calls
`WDSP.SetChannelState(WDSP.id(1,0), 1, 0)` on key-down / `(...,0,1)` on
key-up). The wdsp TX chain simply doesn't execute while idle — no equivalent
of RADE's loopback trick exists here.

So: a real 4-second silent PTT hold (no mic audio, `ZZEF=1` armed, same
`--confirm-tx` on-air-test protocol as every other test tonight), then the
debug dump (`fdvtx_debug_modem.raw`) pulled and analyzed locally.

**Result: genuinely confirms the encoder works.** 71,680 int16 samples, zero
NaN/garbage, zero clipping. Peak exactly **-6.02dBFS** — matches the
`FDV_TX_MODEM_GAIN=0.5` design constant's math precisely
(20·log10(0.5) = -6.02dB), a nice independent cross-check that the gain
staging does what the code comments claim. RMS -12.5dBFS, sane. Spectral
analysis (FFT + 10-band energy split) shows energy sharply concentrated
560–2450 Hz with essentially nothing outside that band — textbook
audio-passband OFDM occupancy, not noise, a stuck tone, or silence.

**One honest loose end, not a red flag**: at the confirmed 8000Hz modem rate
(cross-checked against codec2's `ofdm_mode.c`: `config->fs = 8000.0f`
default, 700E doesn't override it), 71,680 samples works out to ~8.96s of
audio from a 4s PTT hold. No clipping/distortion/NaN anywhere in the extra
content, so most likely a startup catch-up burst in the buffered DSP
pipeline (`fexchange0`'s semaphore-driven worker thread) rather than
anything wrong with the encoder — recorded here as unexplained-but-benign
rather than either hidden or overclaimed as understood.

> **Resolved, next session (`d9e82adb`, 2026-08-18)**: before trusting a
> timeout-based drain design on top of this unexplained gap, it got rechecked
> with tight timestamp bracketing and an immediate dump check — 3.92s of
> audio from a ~4.46s real window, right on target. It was a measurement
> artifact (this session's own delay before checking the dump), not a
> runaway backlog. No corruption theory needed after all.

`ZZEF` disarmed afterward, station back to idle. A good, real checkpoint:
the encoder itself is proven structurally sound. **Still open, same shape as
RADE V1 TX's remaining gaps**: a MOX/PTT arbiter for real end-to-end
operation (RADE V1 TX's EOO-flush-before-drop design is the template), and
eventually off-air decode confirmation — the same standing gap RADE V1 TX
also still has.

### 🟢 700E TX MOX/PTT arbiter — graceful drain on key-up, real-PTT verified (`d9e82adb`, 2026-08-18)

Closes the "no MOX/PTT arbiter at all yet" gap from the entry above. Turns
out 700E TX's *run* gate already is the real MOX transition (`xtxa()`/
`xfdvtx()` only execute while wdsp's `ch[channel].exchange` is on, flipped
by `SetChannelState()` inside `chkMOX_CheckedChanged2`) — unlike RADE V1 TX,
which taps ChannelMaster's raw mic buffer before wdsp runs and needed an
explicit MOX-state push. What 700E TX actually lacked was **graceful
drain**: the instant `SetChannelState(channel, 0, ...)` fired, whatever was
still sitting in `out_ring` (already-encoded modem audio not yet output) or
a partial `speech_ring` block got silently discarded, cutting off the tail
of the operator's last words.

**Design**: a new `draining` flag on `FDVTX` (`fdv.h`/`fdv.c`). While set,
`xfdvtx()` stops ingesting new mic audio but keeps draining what's already
buffered. New PORT functions `SetTXAFDVDraining`/`GetTXAFDVDrained`/
`FlushTXAFDV`. `console.cs`'s `OnMoxPreChangeHandler_FDVTX`, registered on
the same `MoxPreChangeHandlers` delegate as RADE V1 TX's arbiter — same
blocking-before-real-transition mechanism. Key-down flushes the ring
buffers (a stale-carry-over guard, mirroring `RadaeNotifyBeginOver`).
Key-up sets `draining=1` and blocks polling `GetTXAFDVDrained()`,
hard-capped at 2000ms — same principle as RADE V1 TX's EOO-flush cap.
Gated on `GetTXAFDVRun`, so normal SSB/CW/FM/etc PTT sees zero added delay
when 700E TX isn't armed.

**Verified tonight, precisely timed**:
1. First resolved the prior session's "~8.96s from a 4s hold" oddity before
   trusting a timeout design on top of it — see the resolution note above.
2. Regression check: `ZZEF=0` (disarmed), 2s hold, ~0.59s round-trip
   overhead — baseline.
3. Armed test: `ZZEF=1`, 3s hold, ~0.54s overhead — essentially identical
   to baseline, confirming the drain adds negligible delay (well under the
   2000ms cap, a small bounded backlog as expected).
4. Pulled and analyzed the resulting recording: 23,040 samples (2.88s),
   peak -6.03dBFS / RMS -12.6dBFS (matches the prior measurements exactly),
   zero NaN, and critically — the *last* samples right at the drain
   boundary are real, well-formed signal, not garbage/zero/clipped,
   confirming the drain logic doesn't corrupt the trailing edge.

`ZZEF` disarmed afterward, station idle. Closes both the "structurally
sound" and "don't lose the tail on key-up" pieces for 700E TX. **Still
open**: actual off-air decode confirmation (same standing gap RADE V1 TX
has), and the encoder currently reads raw pre-mic-gain audio with no
operator-facing UI/level meter — both future work, not blocking.

### ✅ 700E TX exercised with real speech, not just silence (2026-08-18)

Quick, no-code-change follow-up closing a gap between the two TX features:
every 700E TX test above used silent PTT holds, while RADE V1 TX had
already been tested with real speech. A real 5s PTT hold with the operator
actually talking, `ZZEF` armed — clean key/unkey, app responsive
throughout, 39,040 samples (4.88s) pulled from the recording. Two checks:

- **Per-chunk RMS is essentially flat (std=0.05dB)** across the recording.
  Initially looks suspicious — like the encoder is ignoring the input — but
  this is actually the *correct*, expected characteristic of an OFDM digital
  voice modem: transmitted power stays roughly constant regardless of the
  underlying speech dynamics (comfort/quantized frames during silence run
  at the same modem power as frames from loud speech), unlike an analog SSB
  envelope, which would show real amplitude variation. Worth remembering so
  a future reader doesn't misread flat RMS here as a bug.
- **More decisive**: compared this recording against the earlier
  silent-PTT-test recording over their overlapping length — not
  byte-identical, correlation only 0.37. Confirms the encoder is genuinely
  reactive to different input, not replaying fixed/static frames regardless
  of what's fed to it.

Disarmed afterward. Closes the "real speech, not just silence" testing gap
for 700E TX — now exercised with real voice at least once, matching RADE V1
TX's testing depth. **Still open, same for both TX features**: no second
receiver, so still no independent off-air decode confirmation.

### 🟡 700E RX1 loopback bridge — built, compiles clean, not yet deployed (`04927f76`, 2026-08-24)

A no-RF alternative to the real-off-air-capture route, same purpose as RADE
V1's own loopback (`radae.c`'s `SetRadaeLoopbackEnabled`/`g_loop_bridge`,
proven 2026-08-17/18 above) — but a different mechanism, since 700E's
encode/decode (`fdv.c`'s `FDVTX`/`FDV`) live inside wdsp with no shared file
the way RADE's TX+RX share `radae.c`. Bridges one level up instead, at
ChannelMaster's raw I/Q: `pipe.c`'s `xpipe()` already has both ends as
directly-reachable buffers — TX's `xmtr[0].out[2]` (the same tap `xMixAudio`
already uses for local monitor audio) into RX1's IQ-injection point (the
same one Quick-Play's `xplaywave` uses).

**A correction to the original design sketch, caught before implementing
rather than after**: the proposal (this doc, same session) described the TX
tap as "pre-hardware-domain." Checking the actual code first: wrong — wdsp's
own TXA output resampler has already converted `xmtr[tx].out[2]` to hardware
`ch_outrate` by that point; `xtxgain`/`xeer`/`xilv` (which run after) are
gain/EER/interleave stages, not rate converters. Both the TX tap and RX1's
injection point are already at their respective hardware `ch_outrate`s. Added
a runtime check (`xmtr[0].ch_outrate == rcvr[0].ch_outrate`) rather than
assuming equality — independently-configurable fields with no compile-time
guarantee they match, even though normal single-DSP-rate operation makes
them equal in practice. On a mismatch: silence to RX1, logged once — not
corrupted data, and not real antenna content leaking through to read as a
false "it worked" (the operator explicitly armed a loopback test).

New CAT command **`ZZEG`** (mirrors RADE V1's `ZZDL` exactly): arm `ZZEF`
(TX encode) and this both on, plus RX1's 700E decode, for the full round
trip with zero RF risk.

**Status: compiled clean via CI, not yet deployed or tested** — `hl2winbox`
has been unreachable since ~2026-08-21. First-pass verification (does it
actually sync/decode, matching RADE V1's own loopback proof) is next once
the box is back; this entry is a build checkpoint, not a working-loopback
claim.

### ⚠️ Safety finding: the loopback above did not actually stop real TX audio from reaching the antenna (`7098040a`, 2026-08-24)

Found while first bringing the design above up on `hl2winbox` — a real gap
in the design as proposed and first implemented, not just a build detail.
**The original version tapped `xmtr[0].out[2]` into the bridge but never
stopped that same TX chain from also reaching the real hardware output.**
Since `xtxa()`/`fexchange0` only produce anything at all while genuinely
keyed (the same `SetChannelState`-gated behavior found during the earlier
"no zero-PTT test path" 700E TX work), exercising this loopback as
originally written would have required a real, actually-transmitting MOX
key-down — defeating the entire "no RF needed" premise the design above
claims, the exact thing RADE V1's own loopback correctly avoids by
silencing `mic_io` while armed.

**No RF was radiated.** Caught during a passive `freedv status` check
*before* ever keying real PTT with the loopback armed — everything
disarmed immediately once found.

**Fix**: also silence `buffs[0]`/`buffs[1]` (`out[0]`/`out[1]`, the real
hardware-bound audio) in the same TX hook, confirmed correct by reading
`ilv.c` directly — `xilv()`'s normal (non-EER) path only sends `data[0]`
to `Outbound()`; the HL2 has no EER/dual-DAC hardware, so that path's
`a->run` is never true for this radio. `buffs[2]` (the loopback tap/local
monitor signal) is deliberately left alone — `xilv`'s real send path never
reads it anyway.

Worth being direct about: this gap traces back to the design proposed
earlier in this same entry — it named the right tap points but missed that
tapping isn't the same as silencing. Recorded here plainly rather than
folded quietly into the earlier entry, matching this log's own convention
for self-corrections. Still not yet deployed/tested for actual
sync/decode — this closes the RF-safety gap, not the "does it work" one.

## Stage C — RADE neural mode *(external dependency)*

- ⬜ Watch upstream: David Rowe's RADE V2 C port is in progress (classical DSP ported
  as of May 2026; C library distribution planned after their stored-file test campaign
  validates the algorithms — see freedv.org)

  **Update, 2026-08-21** (checked a fresh local clone of upstream
  [`drowe67/radae`](https://github.com/drowe67/radae) at `/home/w5tsu/Development/radae`,
  history through 2026-07-16): moved on from "planned" —
  - README now carries a dedicated **C port** pointer:
    [`freedv/rade_c`](https://github.com/freedv/rade_c), "for deployment and distribution
    of RADE" — a distinct repo from the one being watched originally.
  - Separately, [`peterbmarks/radae_nopy`](https://github.com/peterbmarks/radae_nopy) is now
    called out in-repo as **"our verified C port"** for the RADE V1 receiver specifically —
    upstream's own bug-report policy requires reproducing against "our verified C port or
    Python," i.e. it's not just a community reference anymore (compare the 2026-08-10
    sv1eia-eval entry below, which treated `radae_nopy` as a third-party reference at the
    time — that framing predates this).
  - A **"⚠️ RADE V2 Status"** notice is now in the README: V2 is under active development,
    **on-air use is explicitly not recommended yet**, known issues are under investigation,
    and official status will be announced on the FreeDV blog — worth checking before this
    project follows V2 any further, since it's a firmer "not yet" than "watch upstream" implied.
  - Net effect on this item: still nothing to build against yet (V1 is what this fork uses;
    V2 has no stable release), but the shape of "when it ships" has firmed up — two repos to
    watch now (`rade_c` for deployment, `radae_nopy` for the verified reference C
    implementation), not one, and V2 in particular has an explicit do-not-deploy notice from
    upstream itself.
- ⬜ When the C library ships: build it with a `build-codec2.yml`-style workflow and
  slot it into `fdv.c`'s mode handle — a dependency bump, not a research project
- Re-coding RADE ourselves was evaluated and rejected: months of duplicated effort and
  a validation problem only upstream's test campaign can solve

### 🟡 sv1eia/Thetis-RADE evaluated as a shortcut — throwaway branch, 2026-08-10

A community fork, [sv1eia/Thetis-RADE](https://github.com/sv1eia/Thetis-RADE) (Christos
Nikolaou, SV1EIA), already ships a full RADE V1 *and* V2 TX/RX pipeline (`radae_c`, a
from-scratch C port referencing `peterbmarks/radae_nopy`, not upstream's own C library —
that's still the thing Stage C above is watching for), a native FreeDV Reporter client,
RADE meters, TX mic conditioning, and its own HL2 support, active through June 2026.
Explored on a throwaway branch (`experiment/sv1eia-radae-eval`, not merged) to see
whether it could shortcut Stage C. A second candidate,
[ancosgrove/Thetis_FreeDV](https://github.com/ancosgrove/Thetis_FreeDV), was also
checked and ruled out immediately — despite the name, its tree has zero FreeDV-related
files; it's a stale `ramdor/Thetis` fork with no actual integration work.

**Mechanically compatible, confirmed.** Both this fork and sv1eia's share real git
history via `ramdor/Thetis` (common ancestor `ed4c27c9`, Oct 2020) — `git merge-base`
finds it, so this is cherry-pick territory, not a vendor-drop-from-scratch situation.
Pulling `Project Files/lib/{radae_c,opus_dnn,r8brain,freedv_text}` and
`ChannelMaster/radae*.{c,h}` verbatim (752 files) landed with **zero collisions** —
entirely additive against our tree.

**`radae_c` (the OFDM modem/DSP layer) is genuinely portable.** Builds standalone with
plain `gcc -O2` on Linux, no CMake, no Windows headers, no `opus_dnn` needed at all for
the acquisition/demod layer (`radc_modem.c`/`radc_acq.c`/`radc_demod.c` are self-
contained pure C — only the neural encoder/decoder, `radc_enc*.c`/`radc_dec*.c`, needs
the DNN).

**The neural codec layer (`opus_dnn`: LPCNet/FARGAN/DRED/OSCE) is *not* reproducibly
buildable from what sv1eia actually committed.** Their own `opus_dnn/commit_pin.txt`
documents this as a known, unresolved gap: it lists three possible build paths (CMake on
a Windows host, hand-authoring an MSVC project file, or WSL cross-compile) and ends with
"**decision deferred to a follow-up step**." Confirmed independently: their vendored
tree is missing `silk/x86/`, `celt/x86/`, `dnn/x86/` and the `arm/` equivalents entirely,
so upstream's own `CMakeLists.txt` fails at configure time (`Cannot find source file:
silk/x86/main_sse.h`) even with `OPUS_DISABLE_INTRINSICS=ON`. `radae_c.vcxproj` links
against `$(OpusDir)\build\$(Platform)\$(Configuration)\` — a prebuilt library directory
excluded by `.gitignore` — meaning SV1EIA builds `opus_dnn` locally and never commits
the result; a fresh clone of that repo cannot build RADE today without first solving
this. This is the real cost of adopting their work, not a one-line dependency bump.

**Sanity-check against `offair_14236000_RADEV1_20260808.wav` — inconclusive, not
negative.** Wrote a standalone harness driving only `radc_modem_init`/`radc_acq_init`/
`radc_acq_detect`/`radc_demod_frame` directly (bypassing the DNN entirely, stopping
right before where the real code would call `rade_core_decoder`) — a faithful port of
`radc_rx.c`'s SEARCH→CANDIDATE→SYNC state machine. Ran it against the full 133 s
capture: no sustained sync (candidate hits ~10% of frames, consistent with noise-floor
false positives, never promoted to SYNC). **Two important corrections/caveats, not a
verdict that RADE V1 modem is broken:**
- **Doc correction**: this capture is *not* raw I/Q despite matching the stereo
  float32/48 kHz container convention used elsewhere in this project — its left and
  right channels are bit-for-bit identical (`corr(I,Q)=1.0`, vs. `~0` for our genuine
  bench IQ files). It's real, demodulated mono RX audio duplicated into both channels.
  Quick-Rec apparently taps a different pipeline point than Quick-Play's IQ-injection
  point documented in Phase 3 above — worth confirming precisely if Quick-Rec is relied
  on again for bench audio.
- **No independent reference decoder exists to validate the test harness itself**
  against — unlike the 700E work (Phase 3), which had `freedv-gui`'s own `freedv_rx` as
  ground truth, there's no lightweight way to build a known-good RADE V1 decoder given
  the `opus_dnn` gap above, so this harness's own correctness is unverified. Also
  unconfirmed: whether this specific 2-minute slice actually overlapped a real
  transmission window (the reporter confirmed the *QSO* was live over 10 minutes, not
  that this particular recording segment did).

**Net assessment (superseded below)**: adopting sv1eia's RADE V1 work is mechanically
clean (file-level) but not a shortcut past Stage C's real work — it trades "wait for
upstream's official RADE C library" for "finish someone else's incomplete `opus_dnn`
build integration," which is bounded and known (their own doc scopes it to three
concrete options) but still real engineering, not a dependency bump. Reusable
regardless of that decision: their native `FreeDVReporter*.cs` (vs. our external
`thetisctl` CLI), RADE meters UI, and TX mic-conditioning chain are relevant
references for Stage B/D UI work independent of whether RADE V1 itself gets adopted.

> **Correction, 2026-08-24**: the "share real git history via `ramdor/Thetis`
> (common ancestor `ed4c27c9`)... cherry-pick territory" claim above does not
> hold up against a fresh check. `git merge-base HEAD sv1eia/main` returns
> **nothing** — sv1eia's repo is a single squash-imported root commit
> (`06c79f4f "Initial commit: Thetis-RADE fork from ramdor/Thetis
> v2.10.3.15-g2e"`, dated 2026-05-06 — *before* this Aug 10 entry was
> written), 36 commits total, no real ancestry with anything here. Either
> the original check ran against the wrong ref, or conflated "same code at
> one point in time" with "same git history." Practical effect: their RADE
> work is **vendor-drop-only** (targeted file copy + manual patch-reading),
> not cherry-pick territory — doesn't change anything about the actual file
> pulls already done (verbatim copies work either way), but changes how any
> *future* sv1eia sync should be approached. See the `thetis-fork-merge`
> skill (`.claude/skills/thetis-fork-merge/SKILL.md`) for the full current
> relationship-to-each-upstream picture, all four repos, not just this one.

### 🟢 `opus_dnn` build gap resolved, same session

Picked up option (a)/(c) from `commit_pin.txt` (CMake + real MSVC, not a hand-authored
vcxproj): root-caused the gap first rather than just working around it.

**Root cause**: not a deliberate deferral despite `commit_pin.txt`'s wording — the
**repo-root `.gitignore`'s standard Visual Studio "build results" patterns**
(`x64/`, `x86/`, `[Aa][Rr][Mm]/`, `[Aa][Rr][Mm]64/`) blanket-match *any* directory
with those names anywhere in the tree, silently swallowing `git add` of opus's own
upstream SIMD source-directory convention (`silk/x86/`, `celt/arm/`, `dnn/x86/`, nine
directories in total, including nested ones like `silk/fixed/arm/`). `commit_pin.txt`
lists them as intended `vendored_subset` entries that simply never landed. This
project's own throwaway-branch vendoring pass hit the exact same footgun on the first
`git add` — caught because `git status` showed nothing at all for genuinely-present
files on disk, the tell that a `.gitignore` pattern (not a real absence) was involved.

**Fix**: restored all nine directories verbatim from real upstream `xiph/opus` at the
exact commit `commit_pin.txt` already pins (`940d4e5af64351ca8ba8390df3f555484c567fbb`
— confirmed byte-identical elsewhere in the tree), and added scoped `!` negation rules
in `opus_dnn/.gitignore` so this can't silently regress. One git subtlety worth noting
for future vendoring: a bare `!/build/x64/` negation does **not** override a preceding
bare `build/` (trailing-slash directory match) — confirmed empirically with an isolated
test repo. The working idiom is `build/*` (glob one level down, doesn't prune the
directory from traversal) followed by the negations; that's what `opus_dnn/.gitignore`
now uses to un-ignore the vendored `build/x64/Release/opus.lib` itself.

**Verified end to end**:
- CMake now configures and builds a complete `libopus.a` on Linux with
  `-DOPUS_DEEP_PLC=ON -DOPUS_DRED=ON -DOPUS_OSCE=ON` — DNN/LPCNet/FARGAN/DRED/OSCE and
  x86 intrinsics all present and building cleanly.
- `radae_c` links against it with **zero undefined references** (previously:
  `compute_generic_dense`, `linear_init`, etc. all unresolved).
  `rade_open("", 0)` succeeds with real weights loaded: `protocol=1 nin_max=1120`.
- Added `.github/workflows/build-opus-dnn.yml` (dispatch-only, mirrors
  `build-codec2.yml`'s shape) to produce the actual MSVC x64 static lib
  `radae_c.vcxproj` expects. Unlike codec2, opus's `CMakeLists.txt` is MSVC-native —
  no MinGW/gendef import-library step needed, just `ilammy/msvc-dev-cmd` + Ninja
  (the plain `-G "Visual Studio 17 2022"` CMake generator no longer auto-detects on
  the current `windows-latest` image — same fix `build-codec2.yml` already uses for
  its own reasons). Workflow file had to be pushed to **master** to be dispatchable at
  all (a hard GitHub platform requirement — `workflow_dispatch` only recognizes files
  present on the default branch — not a scope decision; it dispatches against
  `--ref experiment/sv1eia-radae-eval` for the actual source/build). Run succeeded in
  6m14s: [run 31415733600](https://github.com/W5TSU/OpenHPSDR-Thetis-Hermes-Lite2/actions/runs/31415733600).
- Vendored the resulting `opus.lib` (7.4 MB, real COFF/MSVC archive, contains
  `compute_generic_dense`/`linear_init`/etc.) at
  `Project Files/lib/opus_dnn/build/x64/Release/opus.lib` — the exact path
  `radae_c.vcxproj`'s `AdditionalLibraryDirectories` already expects, so it should
  link unmodified. Not yet compiled *into* a full Thetis build (that needs
  `ChannelMaster.vcxproj`/console wiring beyond this branch's current scope) or
  tested on real Windows hardware — the Linux build is the correctness proof, the
  Windows artifact is the deployment target, and those are two different claims.

**Off-air sanity check re-run with the real (non-stub) DNN decoder**: same result as
the acquisition-only harness — no sync on `offair_14236000_RADEV1_20260808.wav`
across the full 133 s, at any of 6 tested input scales (0.01×–1000×, ruling out a
level/gain mismatch specifically). This is a stronger negative than before (the real
public API, not a hand-ported state machine, so harness fidelity is no longer a live
caveat) but the other caveats from the acquisition-only run stand unchanged: no
independent reference decoder to cross-check against, and unconfirmed whether this
particular 2-minute slice actually overlapped a real transmission window. **Read as**:
the `opus_dnn` build blocker is genuinely gone, but whether this specific recording
ever contained a lockable RADE V1 signal is still an open, separate question — a
fresh, verified-in-the-moment capture (or a real freedv-gui differential test) is the
next step if that question matters going forward, and is now easy to test against
since a working decoder exists.

**Revised net assessment**: the `opus_dnn` build gap that made this a "bounded but
real" adoption cost is now closed — cherry-picking sv1eia's RADE V1 work is both
mechanically clean *and* buildable end-to-end (source → Linux proof → Windows
artifact). What remains before RADE V1 could actually run in this fork is console/
ChannelMaster integration work (Stage B territory: TX/RX wiring, UI, meters) — a
separate, larger decision from the build question this session closed out. Branch
(`experiment/sv1eia-radae-eval`) and the two new CI workflow files are pushed to
`origin`; `build-opus-dnn.yml` is also on `master` (required for dispatchability, no
effect on the normal build). Revisit adoption-vs-wait-for-upstream now that the cost
side of that tradeoff is much better known.

### 🟢 ChannelMaster/console wiring — RX-only, same session

Went ahead and did the Stage B-territory wiring described above, scoped to RX-only
(matching this whole project's established precedent for the 700E prototype) rather
than sv1eia's full TX+RX+UI feature set. **Full solution builds green** end to end:
[run 31425440469](https://github.com/W5TSU/OpenHPSDR-Thetis-Hermes-Lite2/actions/runs/31425440469).

**Native dependency chain, fully vendored and CI-built.** Beyond `opus.lib` (above),
needed `rade.lib` (radae_c itself — already an MSVC project, no CMake needed) plus
three small TX-mic-conditioning libs `radae.c` links against regardless of whether TX
is ever used (C link requirements don't care that `xradae_tx` isn't wired yet):
`rnnoise.lib`, `ebur128.lib` (libebur128), `WebRTC_AGC.lib`. All four already had
vendored `.vcxproj` files from the original sv1eia pull. New workflow
`build-radae-c.yml` (mirrors `build-opus-dnn.yml`'s shape, on `master` for the same
GitHub dispatchability requirement) builds all four via plain `msbuild`, no CMake.
Hit and fixed three more vendoring gaps on the way, each a variant of a theme:
- **Same `.gitignore` footgun again**: `rnnoise/src/x86/` (the SIMD kernel headers
  `vec.h` needs, including `x86_arch_macros.h`) was caught by the identical repo-root
  `x86/`/`arm/` pattern as `opus_dnn`'s gap. Restored from real upstream — but
  `xiph/rnnoise`'s `master` branch is the old, simple GRU codebase; sv1eia's vendored
  copy (`nnet.c`/`vec.h`/`opus_types.h`, architecturally identical to `opus_dnn/dnn/`)
  is from rnnoise's **`main`** branch, a newer DNN-kernel rewrite. Confirmed via a
  byte-identical diff on the shared `vec.h` before trusting the source.
- **A genuinely build-time-fetched dependency, not a vendoring gap**: `rnnoise_data.h`
  (75 MB of trained weights) is deliberately excluded per `rnnoise/.gitignore`'s own
  comment — "fetched at build time from media.xiph.org." Found the real download URL
  pattern from `opus_dnn/dnn/download_model.sh`'s convention
  (`media.xiph.org/<project>/models/<name>-<sha256>.tar.gz`), verified it against
  `rnnoise/model_version`'s pinned hash (checksum matched), and added a fetch+verify
  step to `build-radae-c.yml` rather than vendoring the blob — respects the existing
  "kept out of the repo" decision instead of overriding it.
- **A missing single file, no pattern behind it**: `WebRTC_AGC/agc.h` resolves
  `../../util/sanitizers.h` (a `freedv-gui` 3rdparty-vendoring convention) to
  `Project Files/util/sanitizers.h` — a path outside `WebRTC_AGC/` entirely that
  simply wasn't pulled in the original vendoring pass. One small file from sv1eia.

**`ChannelMaster.vcxproj` wiring** (Debug|x64 and Release|x64 only, matching sv1eia's
own scope — Win32 configs never had RADE deps either): added `radae.c/.h`,
`radae_micdsp.c/.h`, `r8brain_wrap.cpp/.h` (with a `CompileAsCpp` override — the rest
of the project forces `CompileAsC`), r8brain's own `pffft*.c` (no separate r8brain
project, compiled directly in, matching sv1eia), and `freedv_text`'s `rade_text.c` +
its codec2 LDPC dependencies (found via a link error: `radae.c` uses `freedv_text`
for the EOO callsign codec, easy to miss since it only shows up at link time, not
compile time). `AdditionalIncludeDirectories`/`AdditionalDependencies` point at each
native lib's own `..\..\lib\<name>\build\$(Platform)\$(Configuration)\` explicitly
(not sv1eia's shared `$(SolutionDir)`-relative convention, since `rade.lib` etc.
aren't part of our `.sln`). Two more fixes found only by actually building:
`PFFFT_STATIC_DEFINE` (missing from `PreprocessorDefinitions` — without it `pffft.h`
declares its own functions `dllimport`, which collides with defining them in the same
translation unit, C2491) and the LDPC files above.

**`pipe.c`/`cmcomm.h` hook** (the actual hot-path wiring, five single-line `// W5TSU`
insertions, diff verified minimal per this project's mixed-line-endings convention):
`#include "radae.h"` via `cmcomm.h`; `create_radae()`/`destroy_radae()` from
`create_pipe()`/`destroy_pipe()`; `xradae_rx(rx, ppip->rbuff[rx])` right after
`xvacOUT`'s post-DSP audio-data hook, for both the RX1 and other-RX blocks in
`xpipe()` — matches `radae.h`'s own documented "called from xpipe() in pipe.c"
comment and dual-RX (`rx` 0/1) convention exactly. **`xradae_tx` deliberately not
wired** — RX-first, same as the 700E prototype.

**Console hook, deliberately minimal**: `dsp.cs` P/Invoke declarations
(`SetRadaeRxEnabled`/`GetRadaeRxEnabled`/`GetRadaeSync`/`GetRadaeSnrDb`, from
`ChannelMaster.dll` not `wdsp.dll`) and a `RXRadaeEnabled` cached property on
`RadioDSPRX` in `radio.cs`, mirroring `RXAFDVRun`'s survive-rebuild/delayed-update
pattern but calling `SetRadaeRxEnabled(thread, value)` directly (no `WDSP.id()`
channel handle — RADE's RX index is ChannelMaster's plain `thread`/`rx` numbering,
not a per-subrx wdsp channel). **No Setup-tab checkbox, no meters** — matches this
project's own Stage B "Real UI" being explicitly separate/future work even for the
original FreeDV prototype; this is the same infrastructure-layer scope, control-flow-
ready for a future checkbox/CAT/TCI hook but not a finished feature. (Hit the
project's own documented mixed-CRLF/LF hazard editing `radio.cs` — first attempt
flattened line endings into a 786-line spurious diff; reverted and redid the
insertion via direct byte-level splicing on the original content instead.)

**What's still open**: `xradae_tx` (TX path — mic conditioning is fully linked in but
unused), any actual UI, and CAT/TCI exposure. Functionally, `RXRadaeEnabled` defaults
to 0 (off) and nothing drives it yet, so this is inert by default — safe to have
merged into a build, but not yet something an operator can turn on without a debugger
or a follow-up UI/CAT patch. Next real test once there's a way to flip it on: repeat
the off-air sanity check (Stage C, above) through the actual Thetis pipeline instead
of a standalone harness.

**Merged to `FreeDV` 2026-08-10** (`190cac74`) — promoted out of the throwaway
branch once the full solution built green. `experiment/sv1eia-radae-eval` has
since been deleted (both local and `origin`, fully merged first — `git branch -d`
succeeded without needing `-D`); its full history lives on in `FreeDV`'s own log.

### 🟢 CAT toggle for `RXRadaeEnabled`, same session (2026-08-10)

The "not yet something an operator can turn on without a debugger" gap above is
closed via CAT, not a Setup-tab checkbox — lowest footprint, and `thetisctl`
already speaks CAT, so it doubles as the test harness for the off-air sanity check.
Two new commands, mirroring FreeDV-classic's `ZZDV`/`ZZDS` exactly:

- **`ZZDW`** — get/set `RXRadaeEnabled` (RX1/subrx0 only, same single-channel scope
  as `ZZDV`). `console.radio.GetDSPRX(0, 0).RXRadaeEnabled = ...`.
- **`ZZDZ`** — get-only sync/SNR status, `<sync 0|1><sign><snr dB, 3 digits>`. Calls
  `WDSP.GetRadaeSync`/`GetRadaeSnrDb` directly (ChannelMaster, plain `rx` index 0,
  not `WDSP.id()`) — unlike `ZZDS`, `GetRadaeSnrDb` already returns integer dB, no
  `*10` scaling needed.

Registered in `CATStructs.xml` (`nsetparms`/`ngetparms`/`nansparms` matching `ZZDV`/
`ZZDS`) and `CATParser.cs`'s switch. Hit the project's mixed-CRLF/LF hazard again
editing `CATCommands.cs` (a plain-text `Edit` flattened the whole file into a
2496-line spurious diff); reverted and redid the insertion via byte-level splicing,
landing a clean 36-line diff.

**Next real test**: run the off-air sanity check (Stage C, above) through the actual
Thetis pipeline via `thetisctl`'s CAT client (`ZZDW1` to enable, poll `ZZDZ` for
sync/SNR) instead of a standalone harness.

### 🟢 Scripted through `thetisctl`, same session (2026-08-10)

Two pieces, since the existing off-air capture turned out not to be directly
usable (see below):

- **`thetisctl cat radae`** / **`radae-sanity`** (`internal/cat/commands.go`,
  `cmd/thetisctl/{cat_cmd,radaesanity_cmd}.go`): `radae on|off|get`/`status` mirror
  `freedv`'s shape exactly, wrapping `ZZDW`/`ZZDZ`. `radae-sanity` scripts the whole
  check — radae on, Quick Play on (**TX-capable**, reuses `quickplay on`'s
  `--confirm-tx`/dry-run/auto-cleanup gate, documented in
  `.claude/skills/thetis-control/SKILL.md`), poll `radae status` every `--poll`
  (default 1s) for `--hold` (default 140s), always stops Quick Play and disables
  radae decode on the way out (even on error), then prints a sync/SNR summary
  (first-sync time, % of polls synced, max SNR) — optionally to `--csv` too.
- **Re-discovered, then fixed, the "not raw I/Q" problem from the `opus_dnn`
  session above** while building the test: `offair_14236000_RADEV1_20260808.wav`
  is Quick-Rec'd post-demod audio (duplicated into both channels), not the
  pre-DSP analytic I/Q Quick Play's injection point expects — playing it back
  directly would replay duplicated real audio as if it were I/Q. Generalized
  `Tools/FreeDV/make_fdv_test_iq.py` (previously freedv_tx-raw-modem-only) with
  an `--input-wav` mode: reads an arbitrary mono/stereo PCM16-or-float32 wav (own
  minimal reader, no new dependency — numpy only, same as before), averages
  stereo to mono (warns if the channels differ enough that averaging looks lossy,
  i.e. the input might not actually be duplicated-mono), and generalized
  `resample_to_analytic` to accept any `rate_out`/`rate_in` ratio, not just
  integer multiples (needed since a capture's rate won't generally divide evenly
  into 48 kHz, though this specific file is already 48 kHz so no resampling
  actually occurs). Verified against the real file: `--input-wav
  offair_14236000_RADEV1_20260808.wav` reproduces the exact 133.0 s duration
  reported in the `opus_dnn` session above, byte-identical container format to
  the existing 700E bench `.wav`s.

### 🟢 Run against the live `hl2winbox` instance, same session (2026-08-12)

**Deployment gap found and fixed first.** The box was running an old build
(`git:15fe65c7`, ~30 commits behind, predating this whole RADE thread) — CI had
never been asked to ship anything there. Deploying just `Thetis.exe` from the CI
binary artifact (the only file it packages) **crashed the app on the first RADE CAT
query**: `System.EntryPointNotFoundException` → the box's stale `ChannelMaster.dll`
didn't yet export `GetRadaeSync`/`GetRadaeSnrDb`, and the resulting P/Invoke failure
was unhandled, taking down the whole process from a network-triggered command — a
general Thetis robustness gap, not RADE-specific; filed standalone under "Known bugs
found along the way" below. Fixed the immediate deployment problem by extracting the
matching MSI (`msiexec /a ... TARGETDIR=...`,
admin-extract only — never registers/installs, doesn't touch the existing install's
identity) and `robocopy`-ing the full 51-file matching set (`ChannelMaster.dll`,
`wdsp.dll`, `libcodec2.dll`, etc.) into `Thetis-Test`, not just the one file.
Confirmed stable afterward: `radae on/off/get/status` round-tripped correctly
against the live instance with no crash, same PID throughout.

**The actual sanity check**: staged the `--input-wav`-converted
`offair_14236000_RADEV1_20260808.wav` as `SDRQuickAudio.wav` (backing up the
existing 700E bench file first, restored after), then ran `radae-sanity --freq
14236000 --mode DIGU --hold 140s --confirm-tx=...` after explicit per-instance
operator confirmation of the exact frequency/mode/duration (the safety protocol's
"ask again each time" rule, not satisfied by an earlier general go-ahead). **Result:
no sync across all 127 polls, 0.0–138.9 s, SNR 0 dB throughout** — full log/CSV
kept locally. Confirmed unkeyed (`TX: false`) and process-stable afterward.

**Read as**: consistent with the acquisition-only harness's earlier negative result
against this same audio (`opus_dnn` session above) — a second independent
confirmation, this time through the real Thetis pipeline (ChannelMaster's actual
`xradae_rx`/`radae.c` hook, not a standalone harness), that this specific 133 s
capture doesn't lock. Same caveats as before still stand: no independent reference
decoder to cross-check against, and unconfirmed whether this slice actually
overlapped a real transmission window. Not a verdict on the RADE V1 modem or this
branch's ChannelMaster wiring — a fresh, verified-in-the-moment capture is the next
thing that would actually move this question forward, and `radae-sanity` is now the
one-command way to test it whenever that's available.

### 🟢 Two fresh, verified-in-the-moment RADE V1 captures — live polling, still no sync (2026-08-15)

Answered the "fresh capture" ask directly: reused `freedv-reporter watch --tci`'s
auto-tune (Stage D) but reacted to it programmatically instead of watching by eye —
a throwaway wrapper script (not committed) triggered `quickrec on` the instant a
qualifying station started transmitting, polled `radae status` live twice mid-
recording, then `quickrec off` and pulled the file down over SSH. Two RADE V1
transmissions captured this way (`W4GOK`, 14.236 MHz, ~2 minutes each,
`offair_14236000_RADEV1_<timestamp>.wav` ×2, kept locally, not pushed — same
`.wav`/`.gitignore` convention as the earlier capture). **All 4 live `radae status`
polls across both captures: no sync.** This is different in kind from the earlier
negative results — this is sync checked *while the real transmission was actually
happening*, through the real ChannelMaster pipeline, removing the "was this slice
even really live" caveat that qualified every earlier attempt. Still not a verdict
on the modem itself (RADE V1's own reference implementation wasn't checked against
the same traffic), but it's the strongest negative data point so far: live,
verified-in-the-moment, real pipeline, still no lock. `quickrec`'s shared
`SDRQuickAudio.wav`/`.json` slot was backed up before this and restored after,
same as every other session that's used Quick Rec/Quick Play here.

### 🟢 Console UI controls added, "FreeDV" tab (2026-08-15)

CAT (`ZZDW`/`ZZDZ`) was the only way to enable `RXRadaeEnabled` up to this point —
no visible control anywhere in the console. Added a second group box, "RADE V1
(prototype)", right next to the existing "FreeDV (prototype)" one on the
`tpDSPFreeDV` tab (`Setup → DSP → FreeDV`) — same shape, same code pattern
(`chkFreeDVDecode`/`freedvStatusTimer_Tick` in `setup.cs`, mirrored exactly as
`chkRADEDecode`/`radeStatusTimer_Tick`): a checkbox wired straight to
`RXRadaeEnabled`, and a status label polled every 500 ms showing `SYNC SNR
<n> dB` (green) or `no sync`, matching the FreeDV label's own look. One real
difference from the FreeDV version: `WDSP.GetRadaeSync`/`GetRadaeSnrDb` take
ChannelMaster's plain `rx` index (`0` for RX1) directly, not `WDSP.id(0, 0)` —
same distinction `RXRadaeEnabled`'s own setter already had to account for.
Tooltip is explicit that this is experimental with no confirmed-working decode
yet, pointing back at this doc's Stage C — this is a control surface for
testing, not a claim that RADE V1 reception works. CAT (`ZZDW`/`ZZDZ`) and the
checkbox now both drive the exact same `RXRadaeEnabled` property, so either one
reflects the other's state live. Not yet persisted across restarts (matches
`RXRadaeEnabled`'s existing non-persisted behavior) and RX1-only, same scope as
everything else RADE in this branch so far.

### ✅ First confirmed RADE V1 decode over real RF (2026-08-15) — "not confirmed working" resolved

Immediately after the 700E HackRF positive control succeeded (above), applied
the exact same approach to RADE V1 — but this time there was no ready-made
known-good signal to transmit, since (per the caveat that's followed RADE V1
through this whole project) nothing had ever confirmed the modem/decode chain
worked *at all*, not even offline. Closed that gap directly:

- **Built a real RADE V1 encoder**: the local `~/Development/freedv-gui`
  checkout (a genuine, actively-developed reference implementation, not this
  project's own port) already has one — confirmed earlier via its own
  `fullduplex_RADEV1` self-test. Configured (`cmake`, system wxWidgets 3.2.9 —
  much faster than building it from source, `UNITTEST=1`) and built
  (`make -j16`) after installing the documented Ubuntu dependencies
  (`libwxgtk3.2-dev` et al. — needed the operator's `sudo`, since that can't be
  done non-interactively; written to a throwaway `~/install_freedv_gui_deps.sh`
  for them to run). `rade_open` confirmed real built-in DNN weights loaded
  (`V1 n_features_in=432 Nmf=960 Neoo=1152 n_eoo_bits=180`).
- **Generated real, known-good RADE V1 modem audio**, not synthetic: `freedv
  -ut tx -utmode RADEV1 -txfile <speech> -txoutfile <modem>` pipes an actual
  wav through freedv-gui's own TX pipeline and records the RADE-encoded
  result — no PulseAudio virtual-cable dance needed, unlike its `test_zeros.sh`
  harness (that's for a different, full-loopback test). Speech source:
  codec2's standard `ve9qrp.raw` test voice, vendored in the freedv-gui
  checkout (`codec2-1.2.0/raw/`), wrapped in a WAV header via `sox`. Result:
  14.85 s of real, varying, broadband RADE V1 audio at a healthy -4 dBFS peak
  (confirmed by inspecting the actual sample statistics) — shorter than the
  112 s speech input given to it (not yet root-caused why; freedv-gui's own
  internal behavior, third-party code, out of scope to chase further when the
  output is already real and sufficient for a sync test).
- **Converted to I/Q** via `make_fdv_test_iq.py --input-wav ... --peak-dbfs -6`
  — passing `--peak-dbfs -6` explicitly this time, applying the DAC-quantization
  lesson from the 700E flowgraph (above) at generation time instead of needing
  a separate gain-boost block afterward.
- **New flowgraph, `tx_radev1_hackrf.grc`**: built directly from the *fixed*
  700E template — sideband conjugate fix included from the start (the
  inversion is a property of the HackRF TX + HL2 RX hardware chain, not the
  FreeDV mode, so it was expected to apply identically), 20 dB VGA gain
  (the level that worked for 700E) as the starting point.
- **Result: SYNC on the first attempt** — SNR 7–8 dB held across 4 consecutive
  polls (~8 s) of the ~14.85 s transmission, dropping to "no sync" exactly as
  the signal ended. Over the air, licensed operator present, ID'd. No repeat
  attempts needed, unlike 700E's three tries — both known failure modes
  (quantization, sideband) were already fixed going in.

**This resolves the standing "not confirmed working" caveat** that's applied
to RADE V1 throughout this whole document, including as recently as this same
session (see the "So the RADE V1 receive is working?" exchange, answered "no"
at the time — every prior attempt really had failed). RX-only RADE V1 decode
through Thetis's real ChannelMaster/`radae.c` pipeline is now positively
confirmed against a real, known-good, independently-generated signal, over
real RF. Remaining open: `xradae_tx` (still unwired), any persistent/real UI
beyond the Setup-tab checkbox, and CAT/TCI exposure beyond `ZZDW`/`ZZDZ`.

### ✅ Longer-signal retest, same session — sustained sync, not just a brief lock

The 14.85 s clip above answered "does it sync at all," but not "does it hold
up." Built a proper long-form test signal to check:

- **`freedv -ut tx -utmode RADEV1 -txfile <112s speech> -txoutfile <modem>`
  truncates to ~15 s of output regardless of input length** — confirmed this
  isn't proportional (a 10 s input correctly produced 11.37 s of output, but
  the full 112 s speech input still only produced ~15 s), so it's some kind of
  fixed limit in freedv-gui's own `-txfile` UT-mode handling, not a simple
  timing thing. Not root-caused (third-party code, diminishing returns to dig
  further) — worked around instead: split the 112 s speech into ten ~12 s
  chunks (safely under the apparent threshold, confirmed each one completes
  proportionally), ran freedv-gui's TX pipeline on each separately, concatenated
  the ten real RADE V1-encoded outputs into one 126.3 s file. Spot-checked
  amplitude statistics across the full length to confirm it's genuine, varying
  content throughout, not silence or a repeated loop.
- Converted the same way (`make_fdv_test_iq.py --input-wav --peak-dbfs -6`),
  same sideband-fixed `tx_radev1_hackrf.grc` (now pointing at
  `radev1_test_iq_long.wav`), same 20 dB VGA gain.
- **Result: SYNC held continuously for 14 consecutive polls (~112 s) of the
  126.3 s transmission**, SNR ranging 5–14 dB and riding through the source
  audio's natural level variation (dropping to 5–6 dB exactly where the
  amplitude spot-check above found quieter passages, then recovering to
  12–14 dB) without ever dropping sync — dropping to "no sync" only right as
  the transmission itself ended. Over the air, licensed operator present,
  ID'd (single ID sufficient, well under the 10-minute re-ID threshold for a
  ~2-minute transmission).

This is a materially stronger result than the first pass: sustained lock
through real signal-level variation over close to two minutes, not just an
8-second burst. Between this and the two RADE V1 real-traffic captures with
live polling (Stage C, above), the branch's RX-only RADE V1 path has now been
exercised against both a known-good signal (sync achieved, holds) and real
off-air traffic (no sync, on the specific captures tried) — consistent with a
working decoder that simply hasn't yet been tested against off-air RADE V1
traffic that happens to be strong/clean enough to lock, rather than a broken
one.

### ✅ Fixed: decoded audio never reached the local speaker/monitor output (2026-08-15)

Live-testing the longer signal (above) surfaced a real, distinct bug: the
operator could see sync and a healthy SNR, but heard the raw undecoded RADE
V1 modem sound over the speakers, not decoded voice — different from "no
sync" (which is silence, by design) and different from a working decode.

**Diagnosed with evidence, not guesswork.** Two `thetisctl tci rx-audio
capture` recordings, taken from the exact same audio path, told the story
precisely:
- During a confirmed **no-sync** window: exactly 0.0 RMS, dead silence —
  matches `radae.c`'s own "pad with silence on underrun" logic exactly.
- During a confirmed **sync** window (SNR 0–6 dB, sustained): real, varying,
  speech-shaped content — dynamic range from near-silence to loud peaks, nothing
  like RADE's own raw modem signal (measured separately as remarkably flat,
  ~0.32 std throughout, since it's broadband encoded data, not speech).

This proved the decoder itself works — TCI clients already heard real decoded
speech. So the gap had to be downstream of the decode, specific to local
playback. Traced it in `pipe.c`/`cmaster.c`:

```
pipe.c xpipe(), case 1 "Audio data":
  xvacOUT(rx, 1, ppip->rbuff[rx]);   // VAC gets the ORIGINAL audio (before decode)
  xradae_rx(rx, ppip->rbuff[rx]);    // decode happens HERE, modifying rbuff[rx]
  xtciOUT(rx, 1, ppip->rbuff[rx]);   // TCI gets the DECODED audio (confirmed above)
  xrecordwave(rx, 0, 1, ppip->rbuff[rx]); // Quick-Rec gets it too

cmaster.c xcmaster(), case 0 "standard receiver" (runs AFTER xpipe returns):
  xMixAudio(0, 0, chid(stream, j), pcm->rcvr[rx].audio[j]);  // <- feeds local
                                                              //    speaker mix,
                                                              //    reads the
                                                              //    ORIGINAL,
                                                              //    UN-decoded
                                                              //    per-subrx
                                                              //    buffers
```

`ppip->rbuff[rx]` (what `xradae_rx` modifies) is a *separate, summed copy* that
`xpipe` builds from `buffs[0]` (= `pcm->rcvr[rx].audio[0]`) specifically for
VAC/TCI/recording — not the buffer the local-speaker mixer actually reads.
Writing the decode into the copy alone meant three of four consumers got it
(VAC actually didn't either, since it's called *before* `xradae_rx`) and the
one an operator actually listens to — the console's own speakers — never did.

**Fix**: in `pipe.c`, right after `xradae_rx`, added a `GetRadaeRxEnabled(rx)`-gated
block that copies the now-decoded `ppip->rbuff[rx]` back into `buffs[0]` (the
buffer `xMixAudio` reads) and zeroes any other subreceiver buffers
(`buffs[1..cmSubRCVR-1]`) so their raw, undecoded audio doesn't bleed back in
under the final per-subreceiver mix — applied at both `xpipe` call sites (RX1
and "other PowerSDR receivers"). Gated on `GetRadaeRxEnabled` specifically
(not unconditional) so behavior is provably unchanged when RADE decode is off,
including for diversity-RX configurations with `cmSubRCVR > 1`. Two
four-line, `// W5TSU`-tagged insertions, matching this project's usual hook
footprint for shared files.

**Deployed to `hl2winbox` same session** (`git:1c185f14`) via the established
safe path (MSI admin-extract, never registers/installs — plus a full
`robocopy` sync of every changed file, not just `Thetis.exe`, after the
earlier session's stale-`ChannelMaster.dll` crash taught that lesson). CAT
sanity-checked clean post-deploy: `radae on/off/get/status` round-tripped
correctly, no crash, same PID throughout.

**Re-test attempted, inconclusive — not a confirmed pass or fail.** Two runs
of the 126 s HackRF positive control immediately after deploying: **neither
synced at all**, a different failure mode than what this fix targets (the
fix is about audio *after* sync; these runs never reached sync in the first
place). Checked for an actual code-level regression first — traced through
the fix's logic again and it cannot affect sync acquisition: it only runs
*after* `xradae_rx` has already completed its decode and set the sync state
for that frame, and every buffer it touches (`buffs[0]`/`pcm->rcvr[rx].audio`)
gets freshly overwritten by `fexchange0` (the DSP call) on the *next* frame
regardless of what this fix wrote — no feedback path into `xradae_rx`'s own
input exists. Checked the mundane explanations too: `.grc` gain settings
unchanged (`tx_gain` still 20 dB, matching the successful pre-fix runs), no
stray processes holding the HackRF, `hackrf_info` responds cleanly. Operator
confirmed the transmitted signal itself "looked different or weaker" on the
panadapter during these two failed attempts, compared to the successful
pre-fix runs — consistent with real bench-setup RF/USB flakiness (this
HackRF has warned `"3 other devices on the same USB bus... problems at high
sample rates"` since the very first check this session) rather than anything
the fix changed. Paused here for the night rather than keep consuming
real-RF test cycles chasing what looks like hardware variability.

**Net status**: the fix is committed, CI-verified, and deployed — reasoned
through with high confidence and grounded in the same kind of precise,
evidence-based diagnosis (paired before/after audio captures) that found the
bug in the first place, but **not yet confirmed by actually hearing decoded
speech**. That confirmation — a HackRF positive-control run that both syncs
*and* is audibly checked for decoded voice, not just SNR — is the concrete
next step whenever this branch's RADE V1 work is picked back up. If sync
still won't reliably reproduce next time, that's the first thing to
re-diagnose, separately from the audio-routing fix itself.

### 🟡 Reliable sync achieved for the first time; decoded audio still never reaches the speaker — a distinct, deeper bug (2026-08-16)

Picked back up exactly where the entry above left off: "a HackRF positive-
control run that both syncs *and* is audibly checked for decoded voice."
Got a real answer on both halves, but they didn't match — sync now works
reliably; audio still doesn't, and it's a different bug than the one fixed
above.

**Live off-air, twice, before any deliberate testing started.** The operator
watched real off-air RADE V1 traffic and reported "seeing sync... but no
audio, just static and signal noise" — twice, independently, on different
real transmissions, hours apart. Both times a live CAT/TCI check confirmed
real signal reaching the receiver (`ZZDT`, new this session — see below —
read real, varying levels, e.g. -22 to -36 dBFS) but `radae status` had
already dropped back to "no sync" by the time it was checked, and a
concurrent TCI capture during the second event was **exact digital
silence throughout** despite the confirmed-live signal — consistent with
radae.c's documented "pad with silence while unsynced" behavior, i.e. sync
itself wasn't holding at the moment of either capture. Real, but not yet
the deliberately-controlled test the entry above called for.

**Built that deliberate test, from the operator's own audio.** Given a new
file, `Tools/Test-Audio-W5TSU.m4a` (28.18 s voice), rather than reusing the
existing `radev1_test_iq_long.wav`. Pipeline: `gst-launch-1.0`
`decodebin`/`audioconvert`/`wavenc` to decode the m4a (`sox` can't read
m4a directly) → `freedv -ut tx -utmode RADEV1` (freedv-gui's own encoder,
real not synthetic) → `make_fdv_test_iq.py --input-wav --peak-dbfs -6` →
`tx_radev1_hackrf_*.grc`/HackRF, same 14.236 MHz DIGU / 20 dB VGA gain as
every prior RADE V1 RF test.

- **Chunking questioned, tested, and dropped.** First pass split the 28 s
  clip into three ~9.4 s chunks before encoding, following the *documented*
  ~15 s freedv-gui UT-mode truncation quirk from the 2026-08-15 entries
  above. Operator directly asked why chunking was needed instead of one
  encode — the right question: **a single-shot encode of the full 28 s
  clip did *not* truncate** (28.18 s in → 29.61 s out, fully proportional).
  That prior truncation was specific to the earlier, longer (112 s) source,
  not a universal freedv-gui limit — chunking this file was unnecessary and
  added a real confound (each chunk gets its own PTT/EOO cycle from
  freedv-gui's TX pipeline, discontinuities a real transmission wouldn't
  have). Redid it single-shot for every test after this point.
- **Encode independently proven valid, offline, before touching RF.**
  Fed the single-shot modem file straight back into freedv-gui's own `-ut
  rx` reference decoder (no RF, no Thetis) — it synced immediately and held
  for the full clip (`Sync changed from 0 to 1` at start, `1 to 0` right as
  the file ended), and produced real, non-trivial decoded output (28.92 s,
  low but genuine amplitude, not silence). Ruled out "bad encode" entirely
  before spending any more RF time.
- **I/Q conversion also ruled out** — the new file's I/Q wav matched
  `radev1_test_iq_long.wav`'s peak/RMS/spectral stats closely (both
  generated the same way, both -6 dBFS peak).
- **First two attempts (chunked, then single-shot): no sync, `UUUU` HackRF
  USB-underrun markers in the TX log both times.** Operator explicitly
  asked to set the underrun explanation aside and keep digging rather than
  accept it as the answer.
- **Built real diagnostic visibility that didn't exist before.** `radae.c`
  already tracked RX decoder-input level and clip state internally
  (`GetRadaeRxLevelDb`/`GetRadaeClip`) but neither was reachable from
  anywhere — added CAT `ZZDT` (level/clip, works whether or not sync ever
  engages, unlike `ZZDZ`) to check "is signal reaching the decoder at all"
  independent of "did it sync." Found `ZZDT` pinned at exactly -120 dBFS
  (the code's true-silence floor, not just "quiet") through an entire
  transmission that a **concurrent TCI capture proved carried a real,
  strong, -28.5 dBFS signal for the right ~28 s window** — proof the
  decoder's own level tracking wasn't running at all, not that it saw real
  silence.
  - Added `ZZDI` (`GetRadaeDiag`): `g_initialized`, `g_rade[rx]!=NULL`,
    rx-in-range, `ch_outsize` — ruled out a silently-failed `rade_open()`
    and every other early-return guard in `xradae_rx()` (all read
    correctly: initialized=1, handle_valid=1, rx_in_range=1, outsize=64).
  - Added `ZZDJ`/`ArmRadaeRxDebug`: an unconditional, reset-on-demand
    file log at `xradae_rx()`'s very first line (fixed caps don't work
    here — outsize=64 @ 48 kHz implies ~750 calls/sec, exhausting a
    4000-entry cap, fdv.c's own convention, in ~5 s). This is where it got
    interesting: **the armed log showed a clean, sustained ~2.3x level
    jump exactly correlating with the transmission window** (meanabs
    ~0.013 quiet → ~0.030 during TX, real RX1 audio genuinely reaching the
    function), **every guard passing throughout**, yet `ZZDT` queried
    live during that same run *still* read frozen -120.
  - Chased two more code-level hypotheses and ruled both out with
    evidence, not guesswork: `SetRadaeMoxState`/`g_radae_mox_state` (MOX
    gate) and `SetRadaeRxScale`/`SetRadaeRxDialScale` (a per-sample gain
    that, at zero, would silently zero the decoder input while leaving the
    real signal visible everywhere else) both have **zero C# callers
    anywhere in the codebase** — dead code, permanently at their inert
    defaults (mox=0, scale=1.0), confirmed by grep, not assumption.
  - Also asked the operator to check a UI "Loopback" checkbox
    (`chkRADAELoopback`) as a live hypothesis, since a stray TX-loopback
    bridge would explain everything — **wrong call, corrected quickly**:
    that control doesn't actually exist in the UI (`SetRadaeLoopbackEnabled`
    is likewise uncalled dead code, comment-only), so nothing the operator
    found to check was real; re-enabled "Decode RADE V1" (a real checkbox
    they'd unchecked while looking for the nonexistent one) and moved on.
  - Added a second log at the actual write site
    (`radae_xrx_debug2.txt`, no separate CAT trigger — shares `ZZDJ`'s arm
    flag): the computed `blk_peak`/`db` plus an **immediate synchronous
    readback of the same variable in the same function call** — eliminates
    any cross-thread timing question entirely. This is where it resolved
    itself, not through finding a bug in this code: the write and readback
    agreed perfectly on every single line (`db=-21 readback=-21`,
    `db=-25 readback=-25`, ...), with real computed levels throughout the
    transmission (-20 to -25 dBFS, `loopback=0`, `rx_scale=1.0000`
    confirmed inert) — **the level-tracking code was correct all along.**
- **Immediately re-checked `ZZDT` live on the same process right after** —
  it now read a real, current value (-32.0 dBFS ambient), not frozen
  -120. The earlier frozen readings had occurred on a *different* process
  instance, mid-session, where a frequency drift (1.4 kHz off 14236000 Hz,
  most likely bumped at the console during the operator's own
  investigation) was also found and fixed around the same time — plausible
  that some combination of the drift and/or genuine HackRF/USB timing
  sensitivity explained the earlier frozen readings, rather than a
  standing code bug. Not fully isolated which factor mattered how much,
  and not worth further remote time chasing that specific attribution now
  that the practical result below settled the bigger question.
- **✅ Then it just worked, repeatedly.** With frequency corrected and a
  clean process state, re-ran the exact same single-shot 28 s transmission
  twice more: **first-ever reliable sync on a deliberately-built, real
  HackRF RADE V1 transmission** — 19 s continuous the first time (SNR
  8–10 dB), 23 s the second, both climbing quickly and holding steady, not
  a marginal one-off. This is a genuinely new, positive result: RADE V1
  sync had never reproduced this reliably on a controlled test before
  today.
- **🔴 But decoded audio still never reached the speaker — confirmed
  negative twice, at two very different sync durations.** During the first
  confirmed-sync run, operator reported hearing noise, not voice (MUT
  independently confirmed off via `ZZMA`/CAT beforehand). TCI capture
  during a second, longer run's full ~110 s of near-continuous sync
  (re-run with the original 126.3 s `radev1_test_iq_long.wav`, matching
  the file that produced this project's only prior confirmed real decoded
  RADE V1 content, back in the 2026-08-15 local-speaker-fix diagnosis)
  showed a *third*, different pattern from both "raw modem passthrough"
  (flat ~-27 dBFS, established signature) and "true silence" (exact
  -240 dBFS floor): **wide, dynamic swings between near-silence and
  -14 to -20 dBFS peaks** — read as possibly genuine speech dynamics at
  the time, floated to the operator as "maybe sync just needs more time
  before the vocoder engages." **Operator directly listened and
  corrected this: still just noise, not voice**, ruling the priming-time
  theory out too, and correctly overriding a statistical inference (varying
  ≠ necessarily speech) with direct ground truth.

**Where this actually leaves things**: the 2026-08-15 local-speaker fix
(`1c185f14`, copies `ppip->rbuff[rx]` into `buffs[0]` so decoded audio
reaches local speakers, not just TCI/VAC) is very likely fine on its own
terms — it faithfully copies *whatever's in the buffer*, and every piece of
evidence gathered today about that buffer's actual content (TCI, which
reads the identical buffer at the identical pipeline point) says it never
contains real decoded speech during these tests, sync duration
notwithstanding. **The bug is upstream of the routing fix, somewhere in the
actual decode/synthesis chain** (`rade_rx()` → FARGAN vocoder → the 16 kHz
speech resampler/FIFO) — not yet localized further; would need direct
tracing of decoded feature vectors or FARGAN's own output, not just buffer
levels, which is a real step beyond what today's diagnostics reached.
Deliberately paused here rather than continuing further tonight — this
was already a very long session with substantial, real progress (first
reliable sync ever + the audio bug newly and precisely isolated to a
specific pipeline stage) — pick the vocoder/synthesis trace back up fresh
next time.

**Debug instrumentation left in place, all explicitly temporary and
self-documenting** (`ZZDT`/`ZZDI`/`ZZDJ`, the two `radae_xrx_debug*.txt`
file logs, `ArmRadaeRxDebug`) — every addition is `// W5TSU: DEBUG`-tagged
with its own removal note, matching this project's established convention
(`fdv.c`'s own long-lived debug blocks). Remove once the decode/synthesis
bug is found; until then they're genuinely useful for the next session
picking this up. Commits: `8c1f07b0`..`b494d3f2` on `FreeDV` (see `git log`
for the full sequence — CAT diagnostics, single-shot-encode correction, and
this write-up).

### ✅ Resolved, same session: real bug (VAC routing) found and fixed, plus a genuine self-correction, ending in the first confirmed intelligible RADE V1 decode through Thetis

Continued past the pause above rather than stopping — real, decisive progress
came from two directions the CAT/file-log diagnostics alone couldn't reach:
a screenshot of the operator's actual audio setup, and a proper reference
comparison instead of more statistics.

- **Root cause of "no audio" found: VAC was tapped before the RADE V1
  decode, not after.** The operator directly reported the key clue —
  enabling FreeDV 700E audibly changes the sound, enabling RADE V1 decode
  never does, "static remains." That asymmetry was the tell. A screenshot of
  Setup → Audio → VAC 1 confirmed the operator's actual listening path is
  VAC 1 → Voicemeeter (matching the station's documented
  AT2020→Scarlett→Voicemeeter→VAC1 chain) — not the native local-speaker
  mixer every diagnostic that night had been checking. Reading `pipe.c`
  found it immediately once looking in the right place: `xvacOUT()` was
  called *before* `xradae_rx()` at both `xpipe()` call sites, so VAC always
  got the original, never-decoded audio regardless of RADE V1's state — the
  same bug the `1c185f14` local-speaker fix addressed for `xMixAudio`, but
  explicitly left unaddressed for VAC at the time ("VAC actually didn't
  either, since it's called before xradae_rx" — noted in that same commit's
  write-up, 2026-08-15) because the local speaker was the reported symptom
  back then. This also cleanly explains the 700E/RADE V1 asymmetry: 700E's
  decode runs inside the WDSP RXA chain (`fexchange0`, before `xpipe` is
  even called), already baked into what VAC receives; RADE V1's decode is a
  separate ChannelMaster-level hook that, until now, was wired into TCI and
  the local speaker but never VAC. **Fix**: moved `xvacOUT()` to after
  `xradae_rx()` at both call sites, matching where `xtciOUT()` already sits
  (`1425318d`). Verified with a direct proof-of-execution log at both the
  pre- and post-copy points — confirmed `buffs[0]` genuinely receives the
  decoded/silence content before `xMixAudio` reads it, ruling out a
  build/deployment mismatch as an alternative explanation before trusting
  the fix.
- **First live test after the VAC fix: correctly silent, not broken.**
  Sync held 26 s (SNR 7–10 dB) with total silence on VAC throughout —
  initially looked like a new problem, but it's actually the fix working
  exactly as designed: VAC now genuinely reflects RADE V1's decode state,
  and at that point the decode chain's actual output was still an open
  question (see below). A separate minor quirk — audio staying silent even
  after the transmission ended, needing a MUT cycle to recover — was read
  as a likely VAC/Voicemeeter stream-buffering artifact from extended
  silence, not a Thetis bug, and not chased further.
- **A real self-correction: the "FARGAN produces degenerate output"
  finding from earlier that session doesn't hold up.** Added feature/PCM
  tracing (`radae_fargan_debug.txt`, `radae_raderx_debug.txt`) right at the
  `rade_rx()`/`fargan_synthesize()` boundary. A small manual sample of the
  log looked alarming — `feat0` seemingly frozen near -12, `pcm_rms` near
  zero almost every call — and was reported as a likely decode/synthesis
  bug. **Proper statistics over the full log told a different story**:
  `feat0` actually ranges -13.4 to +8.8 (std 3.16), not frozen at all — the
  earlier read was an artifact of too small a sample. Settled it
  conclusively with real ground truth: freedv-gui's own `-rxfeaturefile`
  option dumps the reference decoder's actual feature vectors for the
  identical modem file. The reference shows the *same* pattern (feat0 mean
  -12.4, occasional excursions to +5.2) and, more tellingly, **the
  reference's own confirmed-correct decoded audio is silent in 95.2% of its
  10 ms frames** (median RMS exactly 0.0) — this vocoder's raw output is
  inherently peaky/mostly-silent even when producing genuine, correct
  speech, not a sign of anything broken. The right lesson: judge novel
  vocoder output against a real reference before calling it broken, not
  against an assumption of what "healthy" should look like.
- **The actual pattern: RADE V1 needs a long runway after sync before real
  audio appears — confirmed independently by the reference decoder
  itself.** The reference decode of a 28.8 s test clip was silent until
  *the last second* — not a Thetis artifact, since this is freedv-gui's own
  bundled reference implementation with no Thetis code involved at all.
  That reframed the whole session's short (28–32 s) test transmissions as
  simply too short to demonstrate real decode, and pointed straight back at
  the existing 126.3 s file (`radev1_test_iq_long.wav`) with a specific
  prediction: real content should show up in the *later* part of a long
  sustained sync window, not the start — matching a pattern already visible
  but previously mis-read in an earlier 126 s test that same session (rising
  energy/dynamics in the back half).
- **✅ Confirmed: first-ever intelligible RADE V1 decode through Thetis's
  real pipeline.** Re-ran the 126.3 s HackRF positive control (near-
  continuous sync, t≈2–119 s, SNR 6–10 dB) with a concurrent TCI capture,
  sent both that capture and the reference decoder's own output to the
  operator to listen to side by side. First pass (a shorter single-shot
  file) was unrecognizable even amplified — consistent with "too short."
  The 126 s capture, focused on the **last ~30 s of the sustained sync
  window**: **"I heard every word."** Audio level was fine, no gain tuning
  needed (unlike 700E's `FDV_SPEECH_GAIN` fix earlier the same session).
  This closes out the "sync but no audio" question that opened this entire
  investigation — the fix was real (VAC), and what looked like a second bug
  was a mix of that fix's own correctly-silent behavior plus a genuine
  analysis mistake, not an actual decode/synthesis defect.

**Net status**: RADE V1 RX is now confirmed working end-to-end through
Thetis's real pipeline — sync, decode, and audible speech via VAC — given a
long enough transmission (~2 minutes demonstrated; the exact minimum
runway isn't characterized yet). Real off-air confirmation (as opposed to
a HackRF positive control) is still open, same as 700E's own remaining
off-air-capture gap.

✅ **Debug instrumentation cleaned up, same session (`71f9b26e`).** Removed
`ZZDI`/`ZZDJ` and all the `radae_*_debug.txt`/`pipe_radae_debug.txt`
file-logging code from `radae.c`/`radae.h`/`pipe.c`/`dsp.cs`/
`CATCommands.cs`/`CATParser.cs`/`CATStructs.xml` — all were explicitly
tagged temporary. Kept `ZZDT` (RX decoder-input level/clip), which was
built as a real permanent diagnostic (same tier as `ZZDW`/`ZZDZ`), not
debug scaffolding. Verified after cleanup that both real fixes survived
intact: `pipe.c`'s `xvacOUT`-after-`xradae_rx` reordering, and `ZZDT`'s
P/Invoke + CAT wiring.

### 🟢 Session wrap-up (2026-08-16 → 2026-08-17) — pick up here next time

**What shipped tonight, all on `FreeDV`, all pushed:**
- `fix(freedv)`: `ckQuickPlay.Enabled` stale-async-callback race — fixed,
  live-verified.
- `feat(freedv)`: `FDV_SPEECH_GAIN` 0.30f → 0.75f (700E decoded speech was
  ~28.5 dB quieter than passthrough) — fixed, live-verified, no clipping.
- `docs`: code_documentation regenerated for the branch's new files.
- `docs`: FreeDV/power checkbox persistence confirmed broken across a real
  graceful restart (not just forced-kill) — documented, not fixed.
- **`fix(radae)`: VAC output was tapped before RADE V1 decode, not after
  — the actual root cause of "sync but no audio" for anyone listening via
  VAC (this station's real setup: VAC1→Voicemeeter). Fixed, and for the
  first time ever, real intelligible RADE V1 speech confirmed decoding
  through Thetis** ("I heard every word" — 126.3 s HackRF positive
  control, last ~30 s of a sustained sync window).
- `chore`: all temporary debug CAT commands/logging from that
  investigation removed again once it resolved.

**What's still open, roughly in priority order:**
1. **RADE V1 real off-air confirmation** — everything tonight was a HackRF
   positive control (a known-good signal, deliberately transmitted). No
   real over-the-air QSO has been confirmed to both sync *and* produce
   audible speech yet. Same standing gap 700E has for its own off-air
   capture (Phase 3 step 5).
2. **RADE V1's minimum required "runway"** isn't characterized — only know
   ~2 minutes reliably works and ~30 s doesn't. Worth narrowing if it
   matters for real-world usability (most real QSOs may or may not run
   that long).
3. **`freedv`/`power` checkbox persistence** — confirmed broken (doesn't
   survive any Thetis restart, graceful or forced) but not fixed. Low
   urgency, real annoyance for operators.
4. **DSP buffer size/rate robustness** — never exercised remotely (no
   CAT/TCI hook for it), still open from Phase 3 item 7.
5. **Phase 4 wrap-up items** (installer identity revert, release notes)
   — deliberately deferred until the merge-to-master decision, not before.
6. ~~RADE V1 TX (`xradae_tx`) is still entirely unwired~~ — **superseded**,
   see the loopback section below. TX encoder is now wired and proven via
   software loopback; real on-air TX (MOX/PTT arbiter wiring) remains open.

**Box state left at session end (2026-08-17)**: `hl2winbox` running the
final cleaned-up build, `git:71f9b26e`, confirmed via `thetisctl cat
version` — deployed and `ZZDT` re-verified working post-cleanup. RADE V1
decode on, power on, tuned 14.236 MHz DIGU, VAC 1 enabled per the
operator's normal setup. No TX in progress, nothing armed. Nothing left to
deploy next session — start straight into the open-items list above.

### 🟢 RADE V1 TX encoder wired + first loopback round-trip confirmed (2026-08-17/18)

Started from "are we ready to work on RADE TX?" Investigation before writing
any code found `xradae_tx()` was **not a stub** — a fully-built encoder
(mic conditioning: RNNoise/AGC/EQ -> LPCNet -> `rade_tx` -> modem waveform,
with MOX gating, an EOO/callsign end-of-over burst, and a PTT-hold-until-
flush arbiter already designed in) inherited from the original sv1eia port,
but with **zero callers anywhere** — never hooked into `pipe.c`, never
exposed to C#. Also found a fully-built **RX1 loopback bridge**
(`SetRadaeLoopbackEnabled`) already wired inside `radae.c`: routes the TX
encoder's modem output directly into RX1's decoder input instead of
`mic_io`, so the radio never keys — a safe, no-RF round-trip test path
someone had clearly designed in from the start but never activated.

**What shipped**: `xradae_tx(buff)` hooked into `pipe.c`'s TX mic-data hot
path (same placement pattern as `xradae_rx` on the RX side — after
wav-player/VAC mic-source mixing, before the wav recorder tap). New CAT
commands: `ZZDK` (TX encoder enable), `ZZDL` (RX1 loopback bridge enable),
`ZZDI` (TX mic-input level/clip meter, mirrors `ZZDT`'s RX-side format).
MOX/PTT wiring (`SetRadaeMoxState`/`SetRadaeTxRx`, and the EOO-flush arbiter)
deliberately **not** touched — this step alone cannot key real audio onto
the air.

**Bring-up bug and fix**: after deploying, `ZZDI` stayed frozen at the
-120dBFS floor even with the encoder + loopback both enabled and the
operator confirmed talking (native Thetis mic meter moving, audible to
themselves — ruling out a device/config problem). Added a temporary debug
tap (`ZZDJ`, call counter + last-gate-reached marker) to find which of
`xradae_tx`'s five early-return gates was firing. Root cause turned out to
be **operational, not a code bug**: the very first deploy attempt raced
`msiexec /a`'s administrative extraction — checking the extracted files
after only a 5s sleep caught the *previous* build's binaries still in
place (confirmed by comparing file timestamps against the two CI runs'
actual trigger times). Redone with `Start-Process -Wait`, the diagnostic
build (`git:398266e6`) landed correctly, `ZZDJ` immediately showed
`reason=9` ("reached step 1") with real varying mic-level readings, and
loopback produced **audible, intelligible synthesized speech** — "Yes I
hear my voice" — with `radae status` confirming `SYNC` and a real SNR
reading. Since loopback mode deliberately zeroes `mic_io` after encoding
(so live mic can't leak through the normal TX monitor path), hearing
anything at all is only possible via the encode -> bridge -> RX1-decode
chain actually working end-to-end — the first time this encoder has ever
produced audible output through Thetis. Debug tap removed once resolved
(commit reverting `ZZDJ`/`GetRadaeTxDebug`), same clean-up-after pattern
as every other debugging arc in this log.

**What's still open for real (on-air) TX**: MOX/PTT wiring
(`SetRadaeMoxState`/`SetRadaeTxRx` need real callers at the MOX 0->1/1->0
edge, matching the header comment's documented intent), and the PTT-hold-
until-EOO-flush arbiter (`GetRadaeEooFlushed`/`SetRadaeTxSilenceHold`) needs
to actually gate when Thetis drops PTT, not just exist as unused entry
points. Both are meaningfully riskier than loopback (real keying, real RF)
and deserve their own explicit go-ahead + `--confirm-tx` discipline before
any attempt, same as every other TX-capable `thetisctl` operation in this
project.

**Box state after this section's cleanup (2026-08-17/18)**: `hl2winbox`
running the debug-tap-removed build, `git:42f2dab8`, confirmed deployed via
`thetisctl`. Supersedes the `git:71f9b26e` box-state note above, which
predates this section's work. TX encoder + RX1 loopback bridge both wired
and available (`ZZDK`/`ZZDL`) but not enabled by default; no TX in progress,
nothing armed, MOX/PTT still unwired so real keying isn't possible from this
build regardless.

### 🟢 MOX/PTT wired; first real over-the-air RADE V1 transmission (`d4486c84`, 2026-08-18)

The piece explicitly deferred in the loopback milestone above — real keying —
is now wired. Hooked into `console.cs`'s existing `MoxPreChangeHandlers`
delegate (fires synchronously before the real hardware key transition in
`chkMOX_CheckedChanged2`), not into the core PTT function itself:

- **Key-down (RX→TX)**: unconditionally pushes `SetRadaeTxRx` /
  `SetRadaeMoxState(1)` / `RadaeNotifyBeginOver()` — matches `radae.h`'s
  documented "every MOX edge" intent; inert when `SetRadaeTxEnabled(0)`
  since `radae.c`'s own gate no-ops them.
- **Key-up (TX→RX)**: `RadaeNotifyEndOfOver()` fires unconditionally
  (cheap), but only *blocks* — synchronously, on the same call stack the
  real hardware key-up runs on right after — polling
  `GetRadaeEooFlushed()` when `GetRadaeTxEnabled()` is true, holding real
  PTT open long enough for the end-of-over burst to flush before the radio
  actually unkeys. Hard-capped at 2000ms so a stuck flush can never hang
  the transmitter keyed; normal SSB/CW/FM PTT release sees zero added
  delay.

**Test sequence, in order**: (1) loopback regression — SYNC, 30dB SNR, no
change from the MOX wiring, as expected since loopback bypasses the MOX
gate; (2) silent 3s PTT cycle via `thetisctl cat ptt on --hold 3s
--confirm-tx=...`, RADE TX armed, no mic audio — proved the MOX mechanics
alone: clean key/unkey, `TX: false` confirmed after, Thetis stayed
responsive (same PID) throughout, so the up-to-2s blocking wait never
hangs anything; (3) **a real 6-second over-the-air transmission with the
operator's live voice** at 14.236 MHz DIGU — the first time this station
has ever put an actual RADE V1 signal on the air (not loopback, not the
HackRF positive control). Clean key/unkey again, operator reported nothing
unusual on their end.

**What this does *not* prove — two open items, stated plainly so this
doesn't get read as more finished than it is:**

- **No decode confirmation.** There's no second receiver at this station,
  so nothing confirms the transmitted signal was actually decodable by
  anyone. The MOX/PTT mechanics are proven safe (clean key/unkey, no
  stuck-PTT, app stays responsive); whether the RF that went out is a
  correctly-formed RADE V1 signal is still unverified. Closing this needs
  either a second receiver/SDR listening to the transmission, or a
  cooperating station.
- **⚠️ The end-of-over callsign ID burst is not a real callsign.** The EOO
  burst fired as designed (2 repeats per `xradae_tx`'s scheduling), but
  `SetRadaeEooCallsign` has never been called from C# — the transmission
  went out with whatever bit pattern happened to be sitting in the
  never-initialized `g_rade_eoo_bits` buffer, not the operator's actual
  callsign. Worth being explicit about: this means the on-air test above
  did **not** include a valid station identification burst, which matters
  for anyone treating this milestone as "ready for normal operation" —
  it isn't yet. Wiring `SetRadaeEooCallsign` to a real callsign is an open
  item, not a formality.

TX encoder disarmed (`ZZDK=0`) after testing — station left in a safe
default state (normal QSOs can't accidentally get RADE-encoded), still at
14.236 MHz DIGU, power on, nothing armed.

### ✅ Callsign-ID gap closed — real EOO callsign wired, second on-air test (`3b6ef75b`, `7aa77f6d`, 2026-08-18)

Closes the "⚠️ not a real callsign" item flagged just above. `SetRadaeEooCallsign`
(present in `radae.c` since the original port, never called) finally got a
real caller: new CAT command **`ZZDJ`** — fixed-width 15-char field, Kenwood-CAT
convention (same shape as `KY`), space-padded on set, trimmed/re-padded on get.
Also added `GetRadaeEooCallsign` (a getter didn't exist before — no way to read
back what was set). *(Note: `ZZDJ` was previously used for the now-removed
`GetRadaeTxDebug` diagnostic tap from the loopback milestone above — command
letter reused after cleanup, same convention as elsewhere in this project;
unrelated to the current `ZZDJ`.)*

**Two real bugs found and fixed while wiring this in:**

1. `SetRadaeEooCallsign` called `EnterCriticalSection` unconditionally, missing
   the `g_radae_cs_inited` guard every other use of that critical section in
   the file has. Harmless while it had zero callers; a real crash risk now
   it's CAT-reachable before power-on (a `ZZDJ` sent too early would have hit
   an uninitialized critical section). Fixed both the new getter and the
   existing setter.
2. **`CATParser` response-framing quirk, found on first real test.** A raw
   socket probe showed `ZZDW`/`ZZDT` returning properly framed replies
   (`ZZDW1;`, `ZZDT-1200;`) while the first `ZZDJ` GET implementation came
   back as bare, unterminated text (`TEST`, no prefix, no `;`). Root cause:
   `ParseExtended()` only wraps a reply in its command-name/`;` framing when
   `rtncmd.Length == nAns` *exactly* — the first implementation returned a
   variable-length trimmed string instead of the full 15-char field `nAns=15`
   expects. Fixed by padding the GET response to the full width. Worth
   remembering for any future fixed-width CAT field: `nAns` isn't just
   documentation, `ParseExtended` silently drops framing if the actual
   response length doesn't match it exactly.

Also hit the project's known mixed-line-ending trap on `CATCommands.cs`
mid-fix (see `mixed-line-endings-warning` — caught this time via
`git diff --stat` showing 1370+ changed lines for a 2-line logic fix, before
committing; reverted and redid via byte-precise Python splicing).

**Test**: round-trip verified via raw socket + `thetisctl` after the framing
fix. Operator confirmed `W5TSU` as their real callsign, set via `ZZDJ`, then a
**second real 6-second over-the-air transmission** (same protocol as the
previous milestone — operator's live voice, `--confirm-tx`) — this one now
carries a valid EOO callsign burst instead of uninitialized bits. Clean
key/unkey again, operator confirmed nothing unusual. TX encoder disarmed
(`ZZDK=0`) afterward, station left idle at 14.236 MHz DIGU.

**Still open, unchanged from the milestone above**: no second receiver at
this station, so still no independent confirmation that a real RADE V1
signal transmitted here is actually decodable anywhere. Everything tested
so far proves the TX chain is mechanically correct and now carries a real
station ID — not that the signal is received/decoded elsewhere. That
remains the one open item standing between this and calling RADE V1 TX
genuinely finished.

### ✅ First independent off-air decode of Thetis's own RADE V1 TX (2026-08-18)

Closes the "no second receiver" gap above, using the HackRF RX capture tool
built the same session (`rx_offair_capture_hackrf.grc`/`.py`, `f350fff0`,
`Tools/FreeDV/README.md` "Testing Thetis's TX output"). RADE V1 TX armed
(`ZZDK=1`, loopback off), real PTT keyed via `thetisctl` (6s hold, operator's
live voice, 14.236 MHz DIGU), captured **simultaneously off-air** with the
new tool → `offair_radev1_real.wav` (demodulated mono, 8kHz, ~14s window).
Fed into freedv-gui's own reference decoder as the independent check:
`freedv -ut rx -utmode RADEV1 -rxfile offair_radev1_real.wav -rxoutfile
decoded_radev1_real.raw`.

**Result**: genuine `Sync changed from 0 to 1` then `1 to 0` in the decoder
log, ~6s stable lock — consistent with the PTT hold window, not a spurious
flicker. Decoded output: 264,982 samples @ 16kHz (~16.6s), peak -0.6dBFS,
RMS -24.3dBFS. RMS checked in 0.25s windows across the decode shows real,
varying speech dynamics (quiet gaps ~-50dBFS, voiced segments -14 to
-30dBFS) — genuinely speech-shaped, not flat noise or silence.

**Precision on what this does and doesn't prove**: this is a confirmed real,
intelligible-*shaped* speech decode — sync, timing, and audio dynamics all
check out independently of Thetis's own self-reported status. It is **not**
a content-verified result: no transcript/second-listener comparison was
done to confirm the decoded *words* matched what was actually said into the
mic. That distinction matters for how finished this gap should be considered
— worth a listen-and-compare pass before calling RADE V1 TX fully closed,
not just a numbers-check.

**Provenance note**: the capture (`offair_radev1_real.wav`) and decode
(`decoded_radev1_real.wav`) exist only in the session scratchpad that
produced them, not committed or copied anywhere durable in this repo — the
evidence above is what's citable; the files themselves aren't preserved.
Not affected by the `enableLegacyModes` tooling bug below (RADEV1 was
always the forced/default mode regardless of that flag).

### ✅ freedv-gui test-tooling bug found and fixed: `enableLegacyModes=0` silently forces every test to RADEV1 (2026-08-18/21)

While chasing 700E's still-open off-air decode confirmation, found that
`enableLegacyModes=0` in `~/.freedv` (freedv-gui's own config) silently
**forces every test run to RADEV1 regardless of `-utmode 700E`/`-utmode
700D`/`-utmode 1600` on the command line or clicking a different mode in
the GUI** — no error, no warning, and it hides the Mode selector UI
entirely so there's no visual cue anything is wrong. This invalidated every
700E test attempted in this line of investigation before it was found —
worth considering whether it also explains why the 700E off-air gap was
never closed in earlier sessions, though that's not confirmed, just a
plausible read given how silent the failure mode is.

**Fix**: `enableLegacyModes=1`. **Verified**: re-ran against a known-good
700E reference signal after the fix — real sync, real intelligible decoded
speech, confirmed by ear. Saved as a standing gotcha
(`freedv-gui-enablelegacymodes-gotcha` project memory) so it doesn't cost a
future session the same discovery process.

**Checked against this doc's existing record**: none of the previously
documented 700E claims in this log used the `-ut rx`/`-utmode` reference-
decoder path this bug affects (the RADEV1 test-signal-generation mentions
elsewhere in this doc use `-ut tx -utmode RADEV1`, a different flag and a
different mode, unaffected) — so nothing already on record here needs a
retroactive caveat. This bug only cost fresh 700E attempts made while it
was undiagnosed; the 700E off-air decode-confirmation gap itself **remains
open** — this entry is about the tooling bug and its fix, not a 700E
decode-confirmation milestone.

### 🟡 sv1eia revisited on `trial-merge-upstream`/`fork_merge_20260824` — full-solution build re-verified, Console-side feature comparison done (2026-08-24)

Separate thread from the RX-only wiring above (which lives on `FreeDV`, merged
2026-08-10). This one happened on a different lineage — `trial-merge-upstream`
(the fork's pre-FreeDV-work merge-base) via a new
[`thetis-fork-merge`](../.claude/skills/thetis-fork-merge/SKILL.md) skill built
the same session to formalize merging against all four related Thetis repos, not
just mi0bot/ramdor. Re-vendored sv1eia's RADE core engine fresh onto that lineage
(`f26e4c75`), specifically to evaluate sv1eia's *full* feature set as a possible
replacement candidate for this fork's own RADE work — not just the RX-only subset
already merged.

**Worth being upfront about**: most of the build-gap findings this pass hit
(`opus_dnn`'s SIMD-directory `.gitignore` footgun, `WebRTC_AGC/agc.h`'s escaping
`sanitizers.h` include, the `workflow_dispatch`-must-be-on-`master` requirement)
were **re-discoveries**, not new — all three were already found and fixed on the
`experiment/sv1eia-radae-eval` throwaway branch back on 2026-08-10 (above). Working
from a different branch lineage meant redoing that archaeology rather than reusing
it; noting this so it doesn't read as double the actual ground covered.

**What *is* new this pass: full CI re-verification, and the Console-side feature
comparison the Aug 10 entry only gestured at.**

**Build**: `build-radae-c.yml`/`build-opus-dnn.yml` (already on `master`, satisfying
the dispatch requirement) both green against `fork_merge_20260824`, artifacts
downloaded and committed (`d5d6cd63`), full solution + installer build green
end to end, ~8 minutes. First real confirmation that sv1eia's source, vendored
completely fresh (not reusing any of the Aug 10 branch's fixes), builds clean.

**Console-side comparison** — the real new work. Naive diffs against our tree were
useless (sv1eia normalized several files, `console.cs` included, from LF to CRLF at
some point, turning ~every line into a false "change" — same class of trap as the
ON7OFF diffstat issue in the `thetis-fork-merge` skill, now confirmed a second
time). Diffed with `--ignore-cr-at-eol` against **sv1eia's own root commit**
instead: 40 files, +8,436/−574 real lines.

*What sv1eia has that this fork doesn't*: independent dual RX1/RX2 RADE decode with
per-RX V1/V2 selection; a fully native FreeDV Reporter client (`FreeDVReporter*.cs`,
~1,200 lines, genuinely their own code — request-QSY button, row selection, error
logging — well past this fork's Stage D scope-down to `thetisctl` auto-tune-only);
a Direct2D panadapter RADE overlay (SNR/level bars, colour-coded); a real meter
subsystem integration with per-RX/TX visibility gating; and a mature
mic-conditioning DSP chain with its own Setup UI, iterated across ~30 of their own
commits.

*What this fork has that sv1eia doesn't — the important asymmetry*: sv1eia's RADE
integration has **zero CAT surface**. Checked `CATCommands.cs`/`CATParser.cs`
directly — every line they touched there is `VFO B` redirect plumbing for their
dual-RX feature, not RADE. No equivalent of this fork's `ZZDW`/`ZZDZ` (status),
`ZZDK` (arm), `ZZDJ` (callsign), `ZZEF` (700E) — all of which exist specifically
*because* the Aug 10 work found sv1eia's version lacked them (see "CAT toggle for
`RXRadaeEnabled`" above). Their RADE feature is GUI-only, matching what
`listen_offair_in_freedv_gui.sh` already implied by needing a human to click Start.
This fork's entire verification methodology — `thetisctl`, scripted PTT tests, the
HackRF capture-and-decode loop that closed the RADE V1 off-air gap — depends on
that CAT surface; a wholesale sv1eia swap would mean losing it, not gaining
anything on that axis. Also: sv1eia has no 700E TX work at all (RADE-only).

**A concrete risk, not yet resolved**: `HPSDR/IoBoardHl2.cs` — the file this fork
exists for — differs in a way that isn't cosmetic. This fork's version masks the
hardware-version byte (`hardwareVersion &= 0x0f`, with a comment noting the upper
nibble can vary) before comparing against `Version_1 = 0x01`; sv1eia's has neither
the mask nor the same constant (`0xf1`, compared unmasked). Both trace back to the
same MI0BOT-authored file (sv1eia's copy carries mi0bot's original copyright header
plus a light SV1EIA modification stamp — their HL2 support is a port of mi0bot's
work into their ramdor-based tree, not an independent implementation), so this is
most likely a version-timing difference between what each fork pulled from mi0bot
at different points, not a bug SV1EIA introduced — but it's unverified either way
and matters for real hardware. Don't trust sv1eia's HL2-adjacent files over this
fork's own without resolving this first.

**Net read**: sv1eia is well ahead on UI polish and capability (dual RX, V2,
Reporter GUI, panadapter overlay). But the CAT-scriptability gap is fundamental to
how this project operates and tests, not a nice-to-have — a full replacement would
be a real regression there. Porting sv1eia's UI/feature richness onto this fork's
existing CAT-first architecture looks like the higher-value path over a wholesale
swap, but that's a judgment call for whoever picks up Stage B/D UI work next, not a
decision made here. Branch `fork_merge_20260824` (off `trial-merge-upstream`) has
the full vendor-drop + build fixes + committed artifacts if that work gets picked
up; not merged anywhere yet.

### 🟢 Real Setup UI at last: RADE Setup panel, then unified into a "Digital Voice" mode selector with 700E (2026-08-24/25)

Direct follow-through on the "port sv1eia's UI richness onto this fork's own
architecture" read above — the operator has never had more than two bare
checkboxes for either mode before this. Run via a brainstorming-first process
(the "duplicate everything, then fold 700E in" ask decomposed into 5
independent sub-projects — FreeDV Reporter integration, RX2/dual-channel, and
panadapter overlays don't exist in this fork's backend at all yet, not just
missing UI, so those stay future work). Specs and plans live under
`docs/superpowers/{specs,plans}/` for anyone who wants the full detail; this
entry is the summary.

**Sub-project 1 — RADE Setup panel** (`3183427f` spec, `3ef97f20`/`34fb2689`
implementation, `510d90d9` bugfix). New `Setup → DSP → RADE` tab, three group
boxes: **Mic/TX Conditioning** (mic level, RNNoise, AGC + LUFS target, 3-band
EQ), **RX1 Core** (protocol select, loopback test, RX level, live sync
status), **Diagnostics** (the 5-checkbox bypass ladder, boots OFF). Backend
audit before writing any UI code found the real gap was narrower than it
looked: RX1 enable/TX enable/loopback/sync-SNR/level-meter/EOO-callsign were
already fully CAT-wired (all this project's own prior work); only mic
RNNoise/AGC/EQ, protocol selection, RX/mic gain, and the diagnostics ladder
had working C-level DSP logic with zero C#/CAT exposure — real wiring work,
not new signal processing. 16 new CAT commands. **A real bug found via live
testing**: `SetRadaeMicScale`/`SetRadaeRxScale` take a *linear* gain factor,
but the new spinners were feeding them dB values directly — fixed at both the
C# and CAT layers, reverified live (`-40`/`-20`/`0`/`+40` dB all round-trip
exactly now). Status per its own spec: "Implemented — UI, CAT, and backend
wiring all built, deployed, and verified live on `hl2winbox`."

**Sub-project 2 — unify 700E into one "Digital Voice" mode selector**
(`66ff5f94` spec, `3c9f2cef` plan, `b1689338`/`698f4582`/`cb664941`
implementation). The panel above only covered RADE; this folds FreeDV 700E in
as a peer rather than leaving it on its own separate, more primitive tab.
`cmbRadeMode`: **Off / 700E / RADE V1 / RADE V2** — one exclusive combo box
replacing three independent, previously-uninterlocked enable paths.
**Mutual exclusion is enforced in `radio.cs`'s property setters themselves**,
not just the UI — holds regardless of entry point (the new combo, or the old
per-mode CAT commands used directly). New CAT command `ZZEX` for the unified
mode index; existing per-subsystem commands (`ZZDK`, `ZZEF`, etc.) keep
working unchanged, now cross-disarming each other correctly. The old
`grpFreeDV`/`chkFreeDVDecode` FreeDV tab is removed entirely — same treatment
sub-project #1 gave the old bare `grpRADE` checkbox. Loopback Test is now
mode-aware (`SetFDVLoopbackEnabled` for 700E, `SetRadaeLoopbackEnabled` for
RADE, same checkbox). RADE-only controls (RX Level, Mic/TX Conditioning,
Diagnostics) gray out for Off/700E, since they have no meaning there. Tab
renamed **RADE → Digital Voice**.

**Worth being direct about, not just noted in passing**: picking a mode in
this combo box **arms TX encode, not just RX decode** — the tooltip says so
plainly ("Arming a mode starts its RX decode and TX encode together; picking
a different mode disarms the previous one automatically"). This is a real
behavior change from the old two-checkbox UI, which only ever armed RX. No
RF goes out from touching the dropdown alone — a real MOX/PTT event is still
required — but the *next* PTT after picking a mode transmits that mode's
encoded signal instead of normal SSB, not something the old UI could ever
do by accident. Worth flagging clearly in the User Guide, not just here.

**Verification status**: a live screenshot of the deployed, merged panel
(`Documentation/images/digital-voice-setup-panel.png`, captured on
`hl2winbox` 2026-08-25) confirms all three groups render correctly, the
`cmbRadeMode` combo shows all four options with **RADE V2** selectable and
selected, and the RX1 status label updates (reading "no sync" — a real,
running status read, not a placeholder string). That's UI-render
confirmation, not a full functional pass — actually cycling through all four
modes with real signals, confirming the `radio.cs`-level interlock holds
under real CAT traffic, and a real PTT test of the new TX-arms-on-mode-select
behavior are all still open next steps, not yet claimed here.

## Stage D — FreeDV Reporter spotting *(future, planned 2026-08-08; re-scoped same day)*

Motivation: off-air bench testing (Phase 3 step 5) is blocked on catching a real
FreeDV QSO in progress at exactly the moment we're recording — the
[reporter](https://qso.freedv.org) shows who's currently active, but only if a human
is watching it at the right time. A live spotting feed removes that timing luck, both
for testing and for normal on-air use once Stage B ships a real TX/RX mode.

**Correction (same session)**: initially scoped this as two pieces, the second being
"build a panadapter overlay, no existing rendering code for this anywhere in the
codebase." That was wrong — missed on the first search. **Thetis already has a full
spot-overlay renderer, `Console/SpotManager2.cs`**, and it's *already wired to accept
spots over TCI* — its only caller in the whole codebase is `TCIServer.cs`'s
`handleSpot()`. Thetis expects an *external* client to push spots to it; there is no
rendering work left to do. This collapses the project to one piece:

- **TCI wire format** (`TCIServer.cs::handleSpot`, dispatched on command name `spot`):
  `spot:<callsign>,<mode>,<freqHz>,<argb_color>,<additional_text>;` — `<mode>` is a
  `DSPMode` enum name (case-insensitive; `digu` is what FreeDV/PSK/RTTY-style spots
  use) or a raw string filtered through `SpotManager2.FilterForRawMode` first;
  `<argb_color>` is a signed 32-bit ARGB int (`Color.FromArgb(...)`); `<additional_text>`
  is free text, or a `[json]{...}` tag deserialized into `SpotManager2.JsonSpotData`
  (spotter, country, continent, heading, distance, flag, text colour, SWL fields) for
  richer display — this is genuinely a general-purpose, already-in-the-wild protocol
  (the `-swl[` handling comment in `FilterForRawMode` references "other spot sources,"
  i.e. other tools already push spots to Thetis this way).
- The reporter's own live feed is **Socket.IO**, not a REST/JSON endpoint
  (`cdn.socket.io/4.6.0`, connect to `qso.freedv.org` with
  `{ auth: { role: "view", protocol_version: 2 } }`). Relevant server-pushed events
  (from `/static/js/index.js`): `bulk_update` (full station table snapshot on
  connect), `new_connection`/`remove_connection`, `freq_change`, `rx_report`/
  `tx_report` (who's hearing/being heard, with SNR). A client needs an actual
  Socket.IO library (e.g. a Go client, or the raw socket.io v4 wire protocol over
  `gorilla/websocket` if no maintained Go client exists), not a plain HTTP poll.

✅ **Built and verified (2026-08-09), `ee88a402`** — re-scoped once more while building:
the operator's actual goal was catching a real on-air FreeDV transmission to test/
listen to, not just seeing labels on the waterfall, so `thetisctl freedv-reporter
watch [--min-freq/--max-freq, default 20m] [--tci <ip>] [--tci-port 50001]
[--mode digu]` **auto-tunes RX1 to the activity** rather than pushing `spot:...;`
markers: it tracks live station state from the reporter's feed
(`internal/freedvreporter`, a hand-rolled Socket.IO v4 client — no third-party
dependency, matching `internal/tci`'s existing convention, since the reporter has no
REST API) and on every transmit-start transition within range, retunes Thetis's VFO
A + mode there via the existing `tci` package (not TX-capable — read/tune only).
Verified against the live service in `internal/freedvreporter/live_test.go`
(`go:build live`, excluded from normal `go test ./...`/CI): a real 15s session
tracked 37 real stations with correct callsigns/frequencies, including genuine
`tx=true` transitions on live traffic. Pushing visual `spot:...;` markers onto the
panadapter (the original framing, still fully supported by
`SpotManager2.AddSpot`/`handleSpot`) remains available as a separate, easy follow-on
if wanted later — same event data, different action taken on it.

## Known bugs found along the way (not FreeDV/RADE-specific)

GitHub issues are disabled on this fork, so there's no tracker to file these
into — this doc is it. Findings here are cross-cutting Thetis console bugs
discovered as a side effect of FreeDV/RADE work, not part of either feature.

### ✅ Any CAT command that throws crashes the entire process, not just the CAT session

**Found**: 2026-08-12, working the RADE off-air sanity check (Stage C above) —
`radae status`/`radae get` (`ZZDZ`/`ZZDW`) against `hl2winbox` running a build
whose `ChannelMaster.dll` didn't yet export `GetRadaeSync`/`GetRadaeSnrDb` threw
`System.EntryPointNotFoundException`, and Thetis vanished entirely (process gone,
not just the socket) instead of returning a CAT error reply.

**Root cause — no exception boundary around CAT command dispatch, at any layer**:

```
TCPIPSocketListener.SocketListenerThreadStart()          CAT/TCPIPcatServer.cs:79
  try { ... } catch (SocketException se) { ... }          — only SocketException is caught
  → ParseReceiveBuffer(byteBuffer, size)                  CAT/TCPIPcatServer.cs:275
    → processClientData(msg)                              CAT/TCPIPcatServer.cs:298
      → console.ThreadSafeCatParse(sInboundCatCommand)     CAT/TCPIPcatServer.cs:350, unguarded
        → this.Invoke(() => safeCat(msg))                  console.cs:15702-15709
          → m_objTCPIPCatParser.Get(msg)                   console.cs:15711-15721, unguarded
            → CATParser.Get → ParseExtended → cmdlist.ZZxx(...)
              → whatever that handler does (P/Invoke, array index, cast, ...)
```

Nothing in this chain catches anything but `SocketException`. Any other
exception a command handler can throw — a P/Invoke `EntryPointNotFoundException`
(this case), but just as easily a malformed-argument `FormatException`/
`IndexOutOfRangeException` in a handler that doesn't validate its `suffix`
as carefully as `CATParser.FindSuffix()`'s regex does — propagates back
through `Control.Invoke` to the socket listener thread and becomes an
unhandled exception there. .NET's default policy for an unhandled exception
on *any* thread (not just the UI thread) is to terminate the whole process.
**Impact**: since CAT-over-TCP has no authentication (`.claude/skills/thetis-control/SKILL.md`'s
first line), literally anything that can open a TCP connection to port 13013
can crash a running Thetis instance — accidentally (a buggy client, a typo'd
command) or, worse, deliberately, with zero credentials. `TCIServer.cs:5137`
calls the same `ThreadSafeCatParse` from its own `ParseReceiveBuffer`/
`SocketListenerThreadStart` (the TCI listener) — same gap, second entry point,
not yet checked in detail for its own additional exposure.

**Suggested fix**: wrap the dispatch call, not just the socket read loop — either
`safeCat()` (console.cs:15711) or `TCPIPcatServer.cs`'s call site (line 350),
catching `Exception` (not just `SocketException`), logging it, and returning
a CAT error reply (`"?;"`, matching `processClientData`'s existing `ERROR`
convention) instead of letting it escape. Apply the same fix to `TCIServer.cs`'s
call site. A broader `AppDomain.CurrentDomain.UnhandledException`/
`Application.ThreadException` handler would be reasonable defense in depth on
top of that, but the targeted fix at the CAT/TCI dispatch boundary is the
actual bug — a network protocol handler should never be able to take the
whole app down on a bad request.

**Status**: ✅ fixed 2026-08-15. Applied the targeted fix at the single shared
boundary (`console.cs`'s `safeCat()`, called by `ThreadSafeCatParse` from both
`TCPIPcatServer.cs` and `TCIServer.cs`) rather than duplicating a try/catch at
each listener's call site — one guard covers both entry points. Catches
`Exception` around `m_objTCPIPCatParser.Get(msg)`, logs via `Debug.Print`
(matching `MeterManager.cs`'s existing listener-exception logging convention),
and returns `CATParser.Error1` (`"?;"`, the same wire-level error reply
`processClientData` already uses for its own validation failures) instead of
letting the exception propagate. Not yet re-verified against a live instance
with a command engineered to throw (the original repro — a stale-DLL P/Invoke
gap — no longer reproduces now that `hl2winbox` runs a matching build); the fix
itself is a small, direct, single-file change with an obvious correctness
argument, so this was accepted on that basis rather than manufacturing a new
crash to test against. The `AppDomain.CurrentDomain.UnhandledException`/
`Application.ThreadException` defense-in-depth suggested above was not added —
the targeted fix is the actual bug fix; that would be a separate, broader
hardening decision if wanted later. Was found and worked around at the
deployment level (matching build redeployed, not this bug) during the RADE
live test — see Stage C's "Run against the live `hl2winbox` instance" entry
above for that session's narrative.

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
- `sdr-for-engineers` skill (`~/.claude/skills/sdr-for-engineers/`) — SDR/DSP
  knowledge base built from *Software-Defined Radio for Engineers*
  (Collins/Getz/Pu/Wyglinski); standing reference for synchronization theory
  on this bug (PLL structure, coarse/fine acquisition, Schmidl & Cox OFDM
  sync) — see project memory `act-as-sdr-expert` for when to reach for it
