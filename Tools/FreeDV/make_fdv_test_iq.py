#!/usr/bin/env python3
"""Convert FreeDV/RADE modem audio into a Thetis-playable I/Q wav test signal.

Takes real (not complex) modem-band audio -- either the raw 8 kHz 16-bit mono
output of codec2's `freedv_tx`, or an arbitrary mono/stereo wav (PCM16 or
float32) -- and writes a stereo I/Q wav (left = I, right = Q) that Thetis's
RX wave playback injects ahead of the whole RXA DSP chain, exactly where
antenna samples enter (ChannelMaster pipe.c, "IQ data" position). The input
audio is converted to a single-sideband complex signal (FFT Hilbert
transform, then conjugated to match wdsp's spectral sign convention): with
the VFO on the playback centre and mode DIGU, the signal appears at
+500..+2500 Hz in the passband, image-free.

Usage (freedv_tx raw modem audio, the original use case -- FreeDV-Plan.md
Phase 3):
    python3 make_fdv_test_iq.py modem_8k.raw out.wav \
        [--rate-in 8000] [--rate-out 48000] [--peak-dbfs -50] [--noise-dbfs -95]

Usage (--input-wav, added for the RADE off-air sanity check -- FreeDV-Plan.md
Stage C): re-synthesizes I/Q from an already-demodulated audio capture, e.g.
Quick-Rec output. This is NOT the same as feeding a genuine I/Q wav straight
to Quick Play -- Quick-Rec taps a different, post-demod pipeline point than
Quick-Play's pre-DSP IQ-injection point (confirmed empirically: the existing
offair_14236000_RADEV1_20260808.wav's left/right channels are bit-identical,
corr(I,Q)=1.0, i.e. real mono audio duplicated into a stereo container --
not analytic I/Q). Feeding that file to Quick Play directly would replay
duplicated real audio as if it were I/Q, which is wrong. This mode instead
takes the real captured audio as the *target demodulated signal* and
resynthesizes matching I/Q so that after Thetis's own RX DSP chain
demodulates it back down, the result should approximate the original
capture -- same trick as the freedv_tx raw-modem case above, just with a
wav (any sample rate, mono or stereo -- stereo is averaged to mono, with a
warning if the channels differ enough to suggest the input wasn't actually
duplicated-mono) as the source instead of raw int16:
    python3 make_fdv_test_iq.py --input-wav offair_..._RADEV1_20260808.wav out.wav

Deploy the result on the Windows machine as:
    Music\\Thetis\\quickrecord\\SDRQuickAudio.wav
then press the console's quick Play button (plays into RX1), or drive it via
`thetisctl cat radae-sanity` (RADE) / the manual `quickplay`+`freedv` CAT
commands (700E) -- see Tools/FreeDV/README.md and
.claude/skills/thetis-control/SKILL.md.

Requires numpy only.
"""

import argparse
import os
import shutil
import struct
import sys
import warnings

import numpy as np


def read_wav_mono(path: str) -> tuple[np.ndarray, int]:
    """Minimal WAV reader (PCM16 or IEEE float32, mono or stereo) returning
    (mono float64 samples in [-1, 1], sample rate). Stereo is averaged to
    mono; if the channels differ enough that averaging would be lossy (i.e.
    this probably isn't the duplicated-mono case Quick-Rec produces), warns
    rather than silently discarding a real stereo signal. No dependency
    beyond numpy/struct -- avoids requiring soundfile/scipy just for this."""
    with open(path, "rb") as f:
        data = f.read()
    if data[0:4] != b"RIFF" or data[8:12] != b"WAVE":
        raise ValueError(f"{path}: not a RIFF/WAVE file")

    fmt = None
    audio_bytes = None
    off = 12
    while off + 8 <= len(data):
        chunk_id = data[off : off + 4]
        chunk_size = struct.unpack("<I", data[off + 4 : off + 8])[0]
        body = data[off + 8 : off + 8 + chunk_size]
        if chunk_id == b"fmt ":
            fmt = struct.unpack("<HHIIHH", body[:16])
        elif chunk_id == b"data":
            audio_bytes = body
        off += 8 + chunk_size + (chunk_size % 2)  # chunks are word-aligned

    if fmt is None or audio_bytes is None:
        raise ValueError(f"{path}: missing fmt or data chunk")
    tag, channels, rate, _, _, bits = fmt
    if tag == 1 and bits == 16:
        samples = np.frombuffer(audio_bytes, dtype="<i2").astype(np.float64) / 32768.0
    elif tag == 3 and bits == 32:
        samples = np.frombuffer(audio_bytes, dtype="<f4").astype(np.float64)
    else:
        raise ValueError(f"{path}: unsupported format tag={tag} bits={bits} (want PCM16 or float32)")

    if channels == 1:
        mono = samples
    elif channels == 2:
        stereo = samples.reshape(-1, 2)
        l, r = stereo[:, 0], stereo[:, 1]
        denom = max(np.max(np.abs(l)), np.max(np.abs(r)), 1e-12)
        if np.max(np.abs(l - r)) / denom > 0.01:
            warnings.warn(
                f"{path}: L/R channels differ by >1% of peak -- this may be a genuine "
                "stereo or I/Q signal, not duplicated-mono audio; averaging anyway, "
                "but double-check the input is what you expect"
            )
        mono = (l + r) / 2.0
    else:
        raise ValueError(f"{path}: unsupported channel count {channels} (want 1 or 2)")

    return mono, rate


def resample_to_analytic(x: np.ndarray, rate_in: int, rate_out: int) -> np.ndarray:
    """Band-limited resample of real signal x and conversion to its analytic
    (complex, positive-frequencies-only) form, in a single FFT round trip.
    rate_out need not be an integer multiple of rate_in (needed for
    --input-wav, whose source rate is whatever the capture used) -- the
    output length is simply rounded to the nearest sample, which is exact
    whenever it is an integer multiple (the original, still-common case)."""
    n_in = len(x)
    n_out = int(round(n_in * rate_out / rate_in))

    spec_in = np.fft.rfft(x)
    # Analytic-signal spectrum at the output length: positive bins doubled,
    # negative bins absent. DC (and the input Nyquist bin, which lands mid-
    # spectrum after upsampling) stay single. If downsampling (n_out <
    # len(spec_in)), crop to the output's own Nyquist instead of aliasing.
    spec_out = np.zeros(n_out, dtype=complex)
    take = min(len(spec_in), n_out)
    spec_out[:take] = spec_in[:take] * 2.0
    spec_out[0] = spec_in[0]
    if n_in % 2 == 0 and n_in // 2 < n_out:
        spec_out[n_in // 2] = spec_in[-1]
    return np.fft.ifft(spec_out) * (n_out / n_in)


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    ap.add_argument("modem_raw", nargs="?", default=None,
                    help="raw 16-bit signed LE mono modem audio from freedv_tx (mutually exclusive with --input-wav)")
    ap.add_argument("out_wav", help="output stereo float32 I/Q wav")
    ap.add_argument("--input-wav", default=None,
                    help="mono/stereo wav (PCM16 or float32), e.g. a Quick-Rec capture, instead of modem_raw -- "
                         "see the --input-wav usage note in this script's module docstring")
    ap.add_argument("--rate-in", type=int, default=None,
                    help="input sample rate (default: 8000 for modem_raw; the wav's own rate for --input-wav)")
    ap.add_argument("--rate-out", type=int, default=48000, help="output sample rate (default 48000)")
    ap.add_argument("--peak-dbfs", type=float, default=-50.0,
                    help="peak level of the I/Q signal in dBFS (default -50)")
    ap.add_argument("--noise-dbfs", type=float, default=-95.0,
                    help="added complex white noise floor in dBFS RMS; large negative disables visibly (default -95)")
    args = ap.parse_args()

    if (args.modem_raw is None) == (args.input_wav is None):
        print("error: pass exactly one of modem_raw or --input-wav", file=sys.stderr)
        return 1

    if args.input_wav is not None:
        modem, wav_rate = read_wav_mono(args.input_wav)
        rate_in = args.rate_in if args.rate_in is not None else wav_rate
    else:
        modem = np.fromfile(args.modem_raw, dtype="<i2").astype(np.float64) / 32768.0
        rate_in = args.rate_in if args.rate_in is not None else 8000
    if len(modem) == 0:
        print("error: empty input file", file=sys.stderr)
        return 1

    iq = resample_to_analytic(modem, rate_in, args.rate_out)

    peak = np.max(np.abs(iq))
    if peak > 0:
        iq *= (10.0 ** (args.peak_dbfs / 20.0)) / peak

    if args.noise_dbfs > -200.0:
        rng = np.random.default_rng(1)
        sigma = 10.0 ** (args.noise_dbfs / 20.0) / np.sqrt(2.0)
        iq += rng.normal(0.0, sigma, len(iq)) + 1j * rng.normal(0.0, sigma, len(iq))

    frames = np.empty((len(iq), 2), dtype="<f4")
    # Thetis/wdsp use the opposite spectral sign convention: a straight analytic
    # signal displays in the lower sideband, so write the conjugate to land USB
    # (verified on the panadapter against the FreeDV branch build).
    frames[:, 0] = iq.real  # left  = I
    frames[:, 1] = -iq.imag  # right = Q (conjugate)
    data = frames.tobytes()

    with open(args.out_wav, "wb") as f:
        byte_rate = args.rate_out * 2 * 4
        f.write(b"RIFF")
        f.write(struct.pack("<I", 36 + len(data)))
        f.write(b"WAVE")
        f.write(b"fmt ")
        # format tag 3 = IEEE float, 2 channels, 32 bits — what Thetis records
        f.write(struct.pack("<IHHIIHH", 16, 3, 2, args.rate_out, byte_rate, 8, 32))
        f.write(b"data")
        f.write(struct.pack("<I", len(data)))
        f.write(data)

    dur = len(iq) / args.rate_out
    print(f"wrote {args.out_wav}: {dur:.1f} s, {args.rate_out} Hz stereo float32 I/Q, "
          f"peak {args.peak_dbfs:.0f} dBFS")

    # Also drop a copy under the fixed name Thetis's quick-Play button expects,
    # ready to place in Music\Thetis\quickrecord\ (Thetis maintains the .json
    # sidecar itself on first play).
    quick = os.path.join(os.path.dirname(os.path.abspath(args.out_wav)), "SDRQuickAudio.wav")
    if os.path.abspath(args.out_wav) != quick:
        shutil.copyfile(args.out_wav, quick)
        print(f"wrote {quick} (copy for Thetis quick-Play)")
    return 0


if __name__ == "__main__":
    sys.exit(main())
