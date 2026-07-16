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
adjust signal level and noise floor.)

## Using it in Thetis

1. Copy `fdv700e_test_iq.wav` to the Windows machine as
   `Music\Thetis\quickrecord\SDRQuickAudio.wav`.
2. Thetis: RX1 mode **DIGU**, ~3 kHz filter, NR/NB/ANF/squelch off; enable
   Setup → DSP → NR → "Decode FreeDV 700E (RX1)".
3. Press the console's quick **Play** button. Expect the modem waveform on the
   panadapter, warble → decoded voice, and a green "SYNC SNR x.x dB" label.

The generated `.wav`/`.raw` files are gitignored — only the generator is
tracked.
