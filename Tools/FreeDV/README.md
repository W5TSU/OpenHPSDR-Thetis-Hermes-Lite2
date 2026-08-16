# FreeDV test-signal generator

Produces a Thetis-playable I/Q wav containing a perfect FreeDV 700E signal, for
bench-testing the `fdv.c` decoder (FreeDV-Plan.md, Phase 3) without RF or
off-air activity.

## How it works

Thetis's RX wave playback (the quick **Play** button on the console) injects a
stereo wav — left = I, right = Q — at the very input of the RX DSP chain,
replacing antenna samples (`ChannelMaster/pipe.c`, `xplaywave` at the "IQ data"
position). Any sample rate works; playback resamples to the DSP rate.
`make_fdv_test_iq.py` converts `freedv_tx` modem audio into the analytic
(Hilbert) signal, so it behaves as a clean USB transmission at +500..+2500 Hz
from the VFO: tune DIGU, ~3 kHz filter, and the 700E waveform sits centred in
the passband.

## Regenerating the wav (Linux or WSL)

```bash
git clone --depth 1 --branch 1.2.0 https://github.com/drowe67/codec2.git
cmake -S codec2 -B codec2/build -DCMAKE_BUILD_TYPE=Release
make -C codec2/build -j freedv_tx
codec2/build/src/freedv_tx 700E codec2/raw/ve9qrp.raw fdv700e_modem_8k.raw
python3 make_fdv_test_iq.py fdv700e_modem_8k.raw fdv700e_test_iq.wav
```

(~112 s of speech; needs Python 3 + numpy only. `--peak-dbfs`/`--noise-dbfs`
adjust signal level and noise floor. Alongside the named output the script also
writes `SDRQuickAudio.wav`, a ready-named copy for Thetis's quick-Play button.)

## Using it in Thetis

1. Copy `SDRQuickAudio.wav` to the Windows machine into
   `Music\Thetis\quickrecord\` (Setup → Audio → Recording → "Open Quick Record
   Folder" opens the exact folder; Thetis maintains the `.json` sidecar there
   itself).
2. Thetis: RX1 mode **DIGU**, ~3 kHz filter, NR/NB/ANF/squelch off; enable
   Setup → DSP → **FreeDV** tab → "Decode FreeDV 700E (RX1)" (the "FreeDV
   (prototype)" group box — added 2026-08-15, RADE V1's equivalent control
   sits right next to it in the same tab).
3. Press the console's quick **Play** button. Expect the modem waveform on the
   panadapter, warble → decoded voice, and a green "SYNC SNR x.x dB" label.

The generated `.wav`/`.raw` files are gitignored — only the generator is
tracked.

## Re-synthesizing I/Q from a captured recording (`--input-wav`)

`make_fdv_test_iq.py --input-wav <capture.wav> out.wav` takes an already-
demodulated audio capture (mono or stereo, PCM16 or float32, any sample
rate) instead of `freedv_tx`'s raw modem output, and re-synthesizes matching
analytic-signal I/Q from it the same way. This is what a RADE off-air sanity
check needs: `Tools/FreeDV/offair_14236000_RADEV1_20260808.wav` (Quick-Rec'd
off real traffic, FreeDV-Plan.md Stage C) is *not* raw I/Q despite matching
the stereo float32/48 kHz container convention — its channels are bit-
identical, i.e. real mono audio duplicated into both channels, because
Quick-Rec taps a different (post-demod) pipeline point than Quick-Play's
pre-DSP IQ-injection point. Feeding it to Quick Play directly would replay
that duplicated audio as if it were I/Q, which is wrong; `--input-wav`
resynthesizes the correct signal instead:

```bash
python3 make_fdv_test_iq.py --input-wav offair_14236000_RADEV1_20260808.wav radae_test_iq.wav
```

Then drive the result with `thetisctl cat radae-sanity` (see
`.claude/skills/thetis-control/SKILL.md`) instead of the manual Quick Play +
`freedv status` loop above — same idea, scripted, for RADE's `ZZDW`/`ZZDZ`
instead of 700E's `ZZDV`/`ZZDS`.

## Testing over real RF (HackRF) instead of Quick Play

Everything above injects the I/Q signal directly into Thetis's RX DSP chain
(Quick Play), bypassing the antenna entirely — useful, but it never actually
exercises the HL2's real RF front end. `tx_700e_hackrf.grc` transmits the same
known-good `fdv700e_test_iq.wav` for real, over the air, via a HackRF — a
genuine positive-control test of Thetis's real receive chain (FreeDV-Plan.md
Phase 3 step 6, "live decode"), not a replay.

```bash
gnuradio-companion tx_700e_hackrf.grc
```

- **Center Frequency** parameter defaults to 14.236 MHz (20m, matches the
  frequency used for the RADE V1 off-air monitoring elsewhere in this repo) —
  change it if you want a different dial frequency, but stay inside the
  amateur allocation for your license class.
- **TX VGA Gain** slider defaults to **0 dB** — Part 97 minimum-necessary-power
  practice. Raise it gradually while watching the TX spectrum plot and Thetis's
  RX; the HackRF's RF amp stage is left off entirely (too coarse a step, +14 dB,
  for this kind of test).
- The wav plays **once and stops** (no repeat) — a bounded, single transmission,
  not something that could be left running unattended.
- On the Thetis side: RX1 tuned to the same frequency, mode DIGU, ~3 kHz filter,
  Setup → DSP → FreeDV tab → "Decode FreeDV 700E (RX1)" checked, same as the
  Quick Play procedure above.

**This is a real amateur-radio transmission if run into an antenna** — a
licensed control operator must be present, and station identification (by
voice, CW, or another Part 97-approved method) is required at the start and
end of the transmission; this signal is FreeDV-encoded test speech, not an ID
by itself. For a purely RF-free bench test instead, run the HackRF's TX port
into a dummy load or an attenuator-terminated cable straight into the HL2's
antenna port rather than a real antenna — same flowgraph, nothing radiated.
