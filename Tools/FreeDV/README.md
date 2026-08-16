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

**If Quick Play does nothing** (button press has no effect, nothing plays):
Thetis ships with Quick Play disabled by default at every startup
(`console.resx`'s `ckQuickPlay.Enabled` defaults `False` — a still-open bug,
not fixed as of this writing). Unstick it once per Thetis session by toggling
Quick Record on then off — `thetisctl cat quickrec on` then
`thetisctl cat quickrec off` if scripting, or the equivalent from the console
itself. (A second, related bug — the same disabled state coming back
*mid-session* after repeated on/off cycling — was fixed 2026-08-16, `8c1f07b0`;
only the startup-default case above still needs the manual kick.)

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

## Testing over real RF (GNU Radio Companion + HackRF) instead of Quick Play

Everything above injects the I/Q signal directly into Thetis's RX DSP chain
(Quick Play), bypassing the antenna entirely — useful, but it never actually
exercises the HL2's real RF front end. Two GNU Radio Companion (GRC)
flowgraphs transmit known-good FreeDV I/Q over the air via a HackRF instead,
as a genuine positive-control test of Thetis's real receive chain
(FreeDV-Plan.md Phase 3 step 6 and the RADE V1 "first confirmed decode over
real RF" section):

| Flowgraph | Signal | Default input wav |
|---|---|---|
| `tx_700e_hackrf.grc` | FreeDV 700E (from `freedv_tx`, via `make_fdv_test_iq.py`) | `fdv700e_test_iq.wav` |
| `tx_radev1_hackrf.grc` | RADE V1 (from freedv-gui's own encoder, via `make_fdv_test_iq.py --input-wav`) | `radev1_test_iq_long.wav` |

Both are gitignored companions of the wavs they transmit — regenerate per the
sections above if missing. Neither ships a canned RADE V1 encoder in this
repo; that wav came from a local `freedv-gui` checkout's `freedv -ut tx
-utmode RADEV1` (see `FreeDV-Plan.md`'s "First confirmed RADE V1 decode over
real RF" section for the exact build/generation steps — freedv-gui truncates
`-txfile` output to ~15 s regardless of input length, worked around there by
chunking a longer speech clip and concatenating the RADE-encoded pieces).

### Prerequisites

- **GNU Radio** (3.10+) with **gr-osmosdr** built/installed with HackRF
  support — both flowgraphs use an `osmosdr_sink` block, not a dedicated
  `hackrf_sink`, so gr-osmosdr (which wraps libhackrf) is the actual
  dependency, not just `hackrf-tools`. On Debian/Ubuntu:
  `sudo apt install gnuradio gr-osmosdr`.
- **A HackRF**, connected, with a working udev rule so it's accessible without
  root. Verify with `hackrf_info` before opening GRC — if that doesn't cleanly
  report the device, nothing downstream will work either. (This project's own
  bench HackRF shares a USB bus with other devices and has warned about
  problems at high sample rates as a result — if a run that previously synced
  suddenly doesn't, check `hackrf_info` and try a different USB port/hub
  before assuming a flowgraph or software regression.)
- An antenna, dummy load, or attenuator-terminated cable into the HL2's
  antenna port (see the licensing note below), and the HL2/Thetis on the
  receive side.

### Running one

```bash
gnuradio-companion tx_700e_hackrf.grc      # or tx_radev1_hackrf.grc
```

Press GRC's run (▶) button to start the flowgraph. Common structure for both:

- **Center Frequency** parameter defaults to 14.236 MHz (20m, matches the
  calling frequency used for this repo's off-air FreeDV/RADE V1 monitoring —
  see `thetisctl freedv-scan`/`freedv-reporter watch`'s built-in calling-
  frequency table, `Tools/thetis-ai-control`) — change it if you want a
  different dial frequency, but stay inside the amateur allocation for your
  license class.
- **TX VGA Gain** slider (`osmosdr_sink`'s IF gain) defaults to **20 dB** —
  found empirically (see "Hard-won lessons" below) as the level that reliably
  syncs both modes on this project's hardware. Raise or lower it while
  watching the TX spectrum plot and Thetis's RX; the HackRF's RF amp stage is
  left off entirely (too coarse a step, +14 dB, for this kind of test). Start
  lower (e.g. 0 dB, Part 97 minimum-necessary-power practice) on a new
  setup/antenna and work up rather than assuming 20 dB is safe for your
  situation.
- Sample rate is derived, not a separate control: `samp_rate = wav_samp_rate *
  50` (2.4 Msps for a 48 kHz wav) — HackRF TX quality/reliability drops below
  roughly 2 Msps, so both flowgraphs resample up from the wav's native rate
  rather than transmitting at 48 kHz directly.
- The wav plays **once and stops** (`repeat: False`) — a bounded, single
  transmission, not something that could be left running unattended.
- On the Thetis side: RX1 tuned to the same frequency, mode **DIGU**, ~3 kHz
  filter, Setup → DSP → FreeDV tab → the matching "Decode FreeDV 700E (RX1)"
  or "Decode RADE V1 (RX1)" checkbox, same as the Quick Play procedure above.

### Hard-won lessons (read before building a new flowgraph from these)

Both bugs below were found the hard way, over real air-test cycles, and are
now fixed *in the flowgraphs themselves* — but they're easy to reintroduce if
you build a new HackRF TX chain from scratch, so they're worth understanding
rather than just trusting the fix:

1. **DAC quantization floor.** `make_fdv_test_iq.py`'s default signal level
   (`-50 dBFS` peak) is calibrated for Quick-Play's direct float-sample
   injection, which has no real DAC in the loop. HackRF's TX DAC is 8-bit —
   one quantization step is `1/128` ≈ `-42 dBFS` — so a `-50 dBFS` signal sits
   *below* the DAC's own quantization noise floor and arrives at the receiver
   as an unmodulated-looking carrier, not the OFDM/RADE waveform, no matter
   how much `tx_gain` (analog IF gain, applied *after* the DAC) is raised.
   `tx_700e_hackrf.grc` fixes this with a `blocks_multiply_const_vxx` digital
   gain stage (≈+44.4 dB) between the I/Q assembly and the resampler, so the
   *wav file* stays correct for Quick-Play's unrelated use. `radev1_test_iq_long.wav`
   sidesteps the issue at generation time instead (`make_fdv_test_iq.py
   --peak-dbfs -6`), so `tx_radev1_hackrf.grc` doesn't carry the same gain
   block — if you point either flowgraph at a new wav, check its actual peak
   level (not just its intended one) rather than assuming either approach.
2. **An extra sideband inversion from the real hardware chain.**
   `make_fdv_test_iq.py`'s own conjugate correction (writing `-Q` so the
   signal lands on USB) was calibrated and verified only against Quick-Play's
   direct software injection point. Routing the same wav through a real
   HackRF TX up-converter and back in through the HL2's own RX front end adds
   two more independent mixer stages, and empirically one of them flips the
   sideband again on top of Quick-Play's already-applied correction — a
   signal built correctly per the README above showed up on LSB/DIGL instead
   of the expected USB/DIGU the first time this was tried live. Both
   flowgraphs fix it with a `blocks_conjugate_cc` right after the I/Q
   assembly, canceling the extra inversion back out — chosen over just
   switching Thetis to LSB/DIGL, to keep every flowgraph/doc/tool in this repo
   consistently assuming DIGU rather than carrying a HackRF-TX-specific
   exception. Confirm this is still correct if you retarget either flowgraph
   at different hardware — the inversion is a property of *this* TX/RX
   hardware pairing, not a universal constant.

### Licensing / safety

**This is a real amateur-radio transmission if run into an antenna** — a
licensed control operator must be present, and station identification (by
voice, CW, or another Part 97-approved method) is required at the start and
end of the transmission; this signal is FreeDV/RADE-encoded test speech, not
an ID by itself. For a purely RF-free bench test instead, run the HackRF's TX
port into a dummy load or an attenuator-terminated cable straight into the
HL2's antenna port rather than a real antenna — same flowgraphs, nothing
radiated.
