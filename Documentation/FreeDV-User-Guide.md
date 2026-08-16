# FreeDV in Thetis HL2 — User Guide

> **Status: prototype, `FreeDV` branch, RX-only.** Both modes below are decode-only —
> there is no FreeDV/RADE transmit path yet. This is not part of a released Thetis
> build; it's available on side-loaded `Thetis-Test` builds from the `FreeDV`
> development branch. See `Documentation/FreeDV-Plan.md` for build/implementation
> details and progress notes.

## What is FreeDV?

[FreeDV](https://freedv.org) is an open-source digital voice protocol for HF radio.
It packs a compressed speech codec and a modem into a normal SSB-width channel, so it
runs over the same rigs and antennas as regular voice — no dedicated digital hardware
needed. This build adds native FreeDV decode directly into Thetis's RX chain: no
external FreeDV app, no virtual audio cable, no second computer.

Two decode modes are available, from two different generations of the FreeDV project.
They sound different, behave differently at low SNR, and are implemented by different
libraries — the difference is worth understanding before picking one.

## FreeDV 700E vs. RADE V1

**FreeDV 700E** is the classical, all-DSP mode. Speech is compressed by the Codec2
vocoder to 700 bits/second, protected with LDPC forward error correction, and carried
on a coherent OFDM/PSK modem — the "E" revision is the most fading-resistant of the
700-series. Everything here is conventional, deterministic signal processing (matched
filters, coherent carrier/timing recovery, parity-check decoding), the same lineage
digital voice has used for years. It sounds compressed/synthetic — intelligible, but
clearly a vocoder — and it either syncs or it doesn't: sync drops cleanly under fast
fading or low SNR, and Thetis passes the raw modem audio through until it locks again.

**RADE V1** ("Radio Autoencoder") is FreeDV's newer machine-learning mode. Instead of
a hand-designed vocoder plus a separate modem, RADE trains a neural encoder and
decoder *end-to-end*, together with a model of the radio channel itself — there's no
discrete "700 bits/second" bitstream in the middle, just a continuous learned
representation optimized jointly with how HF actually distorts a signal. The payoff is
meant to be noticeably more natural-sounding speech than Codec2 at a comparable bit
rate — but **that payoff isn't audible yet on this build: see the callout below.**
Its "sync"/SNR figures come from pilot correlation and model confidence rather than an
LDPC parity-check pass, so they don't map onto 700E's numbers directly.

> ⚠️ **RADE V1 currently syncs but produces no decoded audio.** As of 2026-08-16,
> RX sync reproduces *reliably* against a controlled test signal — runs up to ~110
> seconds continuous, SNR 8–10 dB, not marginal — but the speaker/monitor output
> stays noise regardless of sync duration, confirmed by direct listening. This isn't
> the earlier local-speaker-routing bug (fixed 2026-08-15, `1c185f14`) — that fix is
> confirmed working correctly. The bug is upstream of it, in the actual decode/
> synthesis chain, and is under active investigation (`FreeDV-Plan.md`, Stage C).
> Until it's resolved, RADE V1 is not usable for actually listening to a QSO — 700E
> is the only mode that currently produces real decoded audio.

| | FreeDV 700E | RADE V1 |
|---|---|---|
| Speech engine | Codec2 vocoder, 700 bit/s | Neural autoencoder, continuous latent, trained jointly with the channel |
| Modem/FEC | Coherent OFDM/PSK + LDPC | Learned constellation, own pilot/sync scheme |
| Sound | Compressed/robotic, clearly a vocoder | No decoded audio yet (see callout above) — intended to be more natural ("AI voice"-ish) |
| Sync behavior | Locks/drops fairly quickly; passthrough audio while unsynced | Now reproduces reliably on a controlled signal (up to ~110 s continuous, SNR 8–10 dB) |
| Underlying library | `libcodec2.dll` | `radae_c` (`rade.lib`) — separate native library |
| Maturity here | RX prototype working, tested against bench + real-RF signals | Sync reliably reproduces; decoded audio does not yet reach the speaker — investigation ongoing |

Bottom line: **700E** is the only mode worth using to actually listen to a QSO right
now — mature, predictable, and it works. **RADE V1** syncs and reports a healthy SNR,
but don't expect to hear anything through it yet; it's a decode/synthesis-chain bug
away from being useful, not a tuning problem.

## Using it in Thetis

Both modes live in **Setup → DSP → FreeDV** tab, as two separate group boxes:

- **"FreeDV (prototype)"** — **Decode FreeDV 700E (RX1)** checkbox, with a live sync/SNR
  status label underneath.
- **"RADE V1 (prototype)"** — **Decode RADE V1 (RX1)** checkbox, with its own status
  label.

Only one should normally be enabled at a time, on RX1.

**Setup, both modes:**
1. RX1 mode: **DIGU** (USB-side digital voice convention).
2. Filter: roughly 3 kHz wide, centered normally.
3. Turn **off** NR, NR2, ANF, NB, and squelch — these DSP stages hurt a modem/codec
   signal rather than helping it.
4. Tune to a FreeDV or RADE V1 calling frequency (see the calling-frequency table
   referenced from the `thetisctl freedv-scan` tool, or use
   `thetisctl freedv-reporter watch` to auto-tune to live activity reported on
   [FreeDV Reporter](https://qso.freedv.org)).
5. Check the matching decode box. Audio passes through unmodified until sync is
   achieved, then switches to decoded speech. For 700E, decoded volume is tuned to
   roughly match passthrough loudness (see "Known issues / recent changes" below) —
   expect a similar level, not a sudden drop, when sync engages.

**Remote/CAT status:** the 700E sync/SNR state is also exposed over CAT as `ZZFD`
(run flag) / `ZZFS` (sync/SNR query), and RADE V1 status via `ZZDW`/`ZZDZ`, if you're
scripting or monitoring remotely (e.g. via `thetisctl`).

## Known issues / recent changes

- **RADE V1 produces no decoded audio (open, as of 2026-08-16).** Sync now
  reproduces reliably on a controlled test signal (see the callout above), a
  first for this project, but decoded speech never reaches the speaker at any
  tested sync duration — confirmed by direct listening, not just level meters.
  Not the local-speaker-routing bug fixed a day earlier (`1c185f14`, confirmed
  still working correctly); this is further upstream, in decode/synthesis
  itself. Under active investigation — see `FreeDV-Plan.md` Stage C for the
  evidence trail. **Use 700E if you actually want to hear something decoded.**
- **700E decoded volume (fixed 2026-08-16).** Earlier builds dropped noticeably
  quieter than passthrough audio the moment sync engaged (~28 dB quieter, RMS) —
  a jarring "why did it just go quiet" effect. Fixed by raising the decoder's
  internal speech gain; live-verified numbers: decoded speech now runs
  **-30.6 dBFS RMS / -6.0 dBFS peak** against passthrough's **-27.1 dBFS RMS /
  -20.2 dBFS peak** — RMS loudness is now essentially matched (within ~3.5 dB),
  and peaks are actually louder than passthrough's. No clipping observed. This
  was verified against one test signal, not yet broadly on-air across varying
  signal strengths — if 700E decode sounds unexpectedly loud or quiet on real
  traffic, that's worth reporting rather than assuming it's expected.
- **RX-only.** Neither mode transmits. There is no FreeDV/RADE keying path in
  this build yet.

See `Documentation/FreeDV-Plan.md` for the full dated history, evidence, and any
issues still open.

## See also

- `Documentation/FreeDV-Plan.md` — implementation details, architecture, and dated
  progress/testing notes for this feature.
- `Documentation/Thetis_VB-Audio_config.md` §7 — using the external `freedv-gui` app
  with Thetis instead of native decode.
- <https://freedv.org> and <https://freedv.org/radio-autoencoder/> — upstream FreeDV
  and RADE project pages.
