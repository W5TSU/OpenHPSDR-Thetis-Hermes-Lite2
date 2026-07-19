#!/usr/bin/env python3
"""Convert FreeDV modem audio into a Thetis-playable I/Q wav test signal.

Takes the raw 8 kHz 16-bit mono modem audio produced by codec2's `freedv_tx`
and writes a stereo I/Q wav (left = I, right = Q) that Thetis's RX wave
playback injects ahead of the whole RXA DSP chain, exactly where antenna
samples enter (ChannelMaster pipe.c, "IQ data" position). The modem audio is
converted to a single-sideband complex signal (FFT Hilbert transform, then
conjugated to match wdsp's spectral sign convention): with the VFO on the
playback centre and mode DIGU, the signal appears at +500..+2500 Hz in the
passband, image-free.

Usage:
    python3 make_fdv_test_iq.py modem_8k.raw out.wav \
        [--rate-in 8000] [--rate-out 48000] [--peak-dbfs -50] [--noise-dbfs -95]

Deploy on the Windows machine as:
    Music\\Thetis\\quickrecord\\SDRQuickAudio.wav
then press the console's quick Play button (plays into RX1).

Requires numpy only.
"""

import argparse
import struct
import sys

import numpy as np


def resample_to_analytic(x: np.ndarray, rate_in: int, rate_out: int) -> np.ndarray:
    """Band-limited resample of real signal x and conversion to its analytic
    (complex, positive-frequencies-only) form, in a single FFT round trip."""
    if rate_out % rate_in != 0:
        raise ValueError("rate_out must be an integer multiple of rate_in")
    up = rate_out // rate_in
    n_in = len(x)
    n_out = n_in * up

    spec_in = np.fft.rfft(x)
    # Analytic-signal spectrum at the output length: positive bins doubled,
    # negative bins absent. DC (and the input Nyquist bin, which lands mid-
    # spectrum after upsampling) stay single.
    spec_out = np.zeros(n_out, dtype=complex)
    spec_out[: len(spec_in)] = spec_in * 2.0
    spec_out[0] = spec_in[0]
    if n_in % 2 == 0:
        spec_out[n_in // 2] = spec_in[-1]
    return np.fft.ifft(spec_out) * up


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    ap.add_argument("modem_raw", help="raw 16-bit signed LE mono modem audio from freedv_tx")
    ap.add_argument("out_wav", help="output stereo float32 I/Q wav")
    ap.add_argument("--rate-in", type=int, default=8000, help="modem audio sample rate (default 8000)")
    ap.add_argument("--rate-out", type=int, default=48000, help="output sample rate (default 48000)")
    ap.add_argument("--peak-dbfs", type=float, default=-50.0,
                    help="peak level of the I/Q signal in dBFS (default -50)")
    ap.add_argument("--noise-dbfs", type=float, default=-95.0,
                    help="added complex white noise floor in dBFS RMS; large negative disables visibly (default -95)")
    args = ap.parse_args()

    modem = np.fromfile(args.modem_raw, dtype="<i2").astype(np.float64)
    if len(modem) == 0:
        print("error: empty input file", file=sys.stderr)
        return 1
    modem /= 32768.0

    iq = resample_to_analytic(modem, args.rate_in, args.rate_out)

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
    return 0


if __name__ == "__main__":
    sys.exit(main())
