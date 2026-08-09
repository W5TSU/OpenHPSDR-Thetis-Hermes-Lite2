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

### 🟢 Phase 3 — verification *(bench decode achieved 2026-08-08; live decode + iteration remain)*

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
6. ⬜ **Live decode**: 14.236 MHz DIGU. Ground truth: before enabling the
   checkbox, confirm the external FreeDV GUI app (VAC path, §7) syncs on the same
   signal. Note SNR at sync acquire/drop (700E should hold to ~1 dB)
7. ⬜ **Iterate on findings** (once sync is achieved):
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
