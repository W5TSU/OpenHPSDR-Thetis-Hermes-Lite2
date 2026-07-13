# NR3 — RNNoise Neural-Network Noise Reduction

NR3 is one of Thetis's four receive noise-reduction options. It wraps
**[RNNoise](https://gitlab.xiph.org/xiph/rnnoise)** (Xiph.Org / Jean-Marc Valin), a hybrid
DSP/deep-learning noise suppressor that uses a small recurrent neural network (GRU-based) to
estimate per-band gains and a pitch filter, as described in
[*A Hybrid DSP/Deep Learning Approach to Real-Time Full-Band Speech Enhancement*
(arXiv:1709.08243)](https://arxiv.org/pdf/1709.08243.pdf). Unlike NR (LMS, `wdsp/anr.c`) and NR2
(MMSE spectral, `wdsp/emnr.c`), which adapt blindly to whatever they hear, NR3 applies a **trained
model** — by default one trained on speech — which is why it can be strikingly effective on SSB
voice and why the model can be swapped for one trained on different material.

The Thetis integration was written by Richard Samphire MW0LGE (building on ideas from
[vu3rdd/wdsp](https://github.com/vu3rdd/wdsp)) and uses an **unmodified** RNNoise library, with a
ring-buffer shim to reconcile frame sizes.

---

## 1. Architecture — every layer, top to bottom

```
Main window NR button / DSP menu / CAT ZZNE/ZZNF / MIDI / Andromeda
        │  SelectNR(rx, sub, 0..4)                Console/console.cs
        ▼
RadioDSPRX properties: RXANR3Run, RXANR3Position,
                       RXANR3FixedGain             Console/radio.cs (L2259–2312)
        │  P/Invoke externs                        Console/dsp.cs (L247–257)
        ▼
wdsp.dll: SetRXARNNRRun / SetRXARNNRPosition /
          SetRXARNNRUseDefaultGain / RNNRloadModel wdsp/rnnr.c
        │  rnnr block in the RXA receive chain     wdsp/RXA.c (create L335, run L660/L668)
        ▼
rnnoise.dll: rnnoise_create / rnnoise_process_frame /
             rnnoise_model_from_filename           Project Files/lib/NR_Algorithms_x64
```

| Component | File(s) | What it does |
|-----------|---------|--------------|
| NR3 wrapper block | `Project Files/Source/wdsp/rnnr.c`, `rnnr.h` | Owns the RNNoise state per receiver, input AGC, ring buffers, model hot-swap. |
| RXA chain integration | `Project Files/Source/wdsp/RXA.c` | Creates one `rnnr` per receiver channel (L335) on the mid buffer; runs it pre-AGC (L660) or post-AGC (L668) depending on position; `RXAbp1Check/Set` widens the bandpass when any NR runs. |
| RNNoise library | `Project Files/lib/NR_Algorithms_x64/` (`rnnoise.dll`, `rnnoise_avx2.dll`, `rnnoise.h`, `rnnoise.lib`) | The unmodified Xiph library, statically linked into `wdsp.dll` (see `wdsp.vcxproj` AdditionalDependencies). Full source bundled under `src/rnnoise/`. |
| P/Invoke surface | `Project Files/Source/Console/dsp.cs` L247–257, L253–254 | `extern` declarations for the four NR3 entry points. |
| Per-RX state | `Project Files/Source/Console/radio.cs` L2259–2312 | `RXANR3Run/Position/FixedGain` properties push changes down only when they differ (or on `force`). |
| UI selection | `Project Files/Source/Console/console.cs` (`SelectNR`, ~L44330) | Mutually-exclusive NR selection 0–4 per RX; updates button text (`chkNR` → "NR3"), menus, and CAT status. |
| Model management UI | `Project Files/Source/Console/setup.cs` L35775–35870 | Load/Default buttons, filename label, model validation, persistence (`nr3_model_file` database key, L1631/L1993). |
| Installer | `Project Files/Source/Thetis-Installer/Product.wxs` L361–375 | Ships `rnnoise_weights_small.bin` and `rnnoise_weights_large.bin` into the Thetis install folder. |

## 2. How the signal is processed (`xrnnr`, rnnr.c L238–313)

RNNoise has hard expectations: **mono, 48 kHz, fixed 480-sample (10 ms) frames, 16-bit-PCM-scale
sample values**. The RXA chain delivers none of those directly, so the wrapper adapts:

1. **Mono tap** — only the real (I) component of the complex mid-buffer is taken
   (`a->in[2*i+0]`); after demodulation the audio is real, so nothing is lost. On output the
   imaginary part is written as 0.
2. **Frame-size adaptation** — Thetis's DSP block size (`dsp_size`) rarely equals 480, so input
   samples accumulate in a ring buffer; whenever ≥480 are available a frame is processed, and
   denoised samples queue in an output ring. Until the output ring holds a full DSP block, the
   input is passed through unmodified — so NR3 adds roughly **one RNNoise frame (~10 ms) of
   buffering latency** when engaged.
3. **Level matching (input AGC)** — RNNoise expects 16-bit-scale amplitudes, while wdsp samples
   are ~±1.0. Two modes (`SetRXARNNRUseDefaultGain`, toggled by Setup's *fixed gain* checkbox,
   default **on**):
   - **Fixed gain** — multiply by 500,000 (`VU3RDD_DEFAULT_GAIN`, the original vu3rdd approach).
   - **Adaptive AGC** — per-frame RMS is servo'd toward a 75 dB target (`AGC_TARGET_DB`) with
     10 ms attack / 200 ms release, clamped to −12…+220 dB.
   Either way samples are safety-clipped at ±30,000, `rnnoise_process_frame()` runs, and the
   **inverse gain** restores normal levels — so the AGC is transparent to the rest of the chain.
4. **Chain position** — `position` 0 runs NR3 **before** the AGC (`xwcpagc`), 1 (default) runs it
   **after**; the same Setup radio button moves NR/NR2/NR3/NR4 and ANF together
   (`setup.cs radANFPreAGC_CheckedChanged`, L8994). Pre-AGC generally behaves better because the
   AGC then rides the *denoised* signal level.

Thread safety: every instance carries a critical section; `SetRXARNNRRun` also takes the channel
DSP lock and re-evaluates the shared bandpass (`RXAbp1Check`) since NR3 (like NR2/NR4) forces the
wider `bp1` filter arrangement.

> **Sample-rate caveat:** the wrapper feeds RNNoise at the channel's DSP rate and Thetis's RXA
> chain runs at 48 kHz by default — exactly what the models are trained for. `a->rate` is stored
> ("for future use") but no resampling is performed, so a non-48 kHz DSP rate would shift the
> model's frequency axis.

## 3. How to use NR3

**From the main window** — click the **NR** button on the DSP panel until it reads **NR3** (the
selections are mutually exclusive: OFF → NR → NR2 → NR3 → NR4), or pick NR3 from the DSP menu.
RX1 and RX2 have independent selections; a sub-receiver follows its main receiver's DSP settings.

**From Setup → DSP → Noise Reduction:**
- **Pre-AGC / Post-AGC** radio — chain position for all NR variants and ANF (see above).
- **RNNoise fixed gain** checkbox (`chkNR3_RNNoiseFixedGain`) — checked (default) = fixed
  ×500,000 input gain; unchecked = adaptive input AGC. Try unchecking if very strong or very weak
  signals seem to pump or distort.
- **Model** — a label showing the active model (`Default (large)` unless a file is loaded) with
  **Load** and **Default** buttons (see §4).

**Remotely** — CAT `ZZNE` (RX1) and `ZZNF` (RX2) set/read the NR selection as `0`–`4` (`3` =
NR3); TCI, MIDI (Midi2Cat), and Andromeda button bars can all bind the NR3 toggle
(`OtherButtonId.NR3`).

**What to expect** — the default models are trained on *speech*, so NR3 shines on SSB/AM voice
under white-ish band noise. It is the wrong tool for CW or digital modes (the network will treat
tones as noise or artifacts), and unusual noise types (ignition, static crashes) are only removed
as well as the training data covered them — which is exactly what custom models are for.

## 4. Models — where they live and how to add one

**The built-in model.** `rnnoise.dll` is compiled with a default weight set baked in; Thetis
labels it *Default (large)*. Passing an empty path to `RNNRloadModel("")` reverts to it (this also
happens at radio shutdown, `radio.cs` L182).

**Shipped loadable models.** The installer places two ready-to-load blobs in the Thetis install
folder (typically `C:\Program Files\Thetis-HL2\`):

| File | Notes |
|------|-------|
| `rnnoise_weights_large.bin` | Same quality class as the built-in default. |
| `rnnoise_weights_small.bin` | Smaller/faster variant, lower CPU at some quality cost. |

**Loading a model:** *Setup → DSP → Noise Reduction → NR3 → Load*, pick a `.bin` file. What
happens (`setup.cs setNR3Model` → `wdsp rnnr.c RNNRloadModel`):

1. The file is **validated** in C# first (`validateRnnoiseModel`, L35840): the RNNoise weight
   blob is a sequence of 64-byte block headers + payloads (mirroring the library's
   `parse_lpcnet_weights.c`); malformed files are rejected with a warning rather than crashing
   the DSP.
2. `RNNRloadModel(path)` then hot-swaps the model **globally**: every live RNNR instance is
   paused and its state destroyed, the previous `RNNModel` is freed, the new blob is loaded via
   `rnnoise_model_from_filename()`, and each instance is recreated and resumed. The model is
   process-wide — all receivers use the same model; there is no per-RX model.
3. The choice persists in the settings database under `nr3_model_file` and is re-applied at
   startup. **Default** clears it back to the built-in model.

**Format constraints (important):**
- Since RNNoise v0.1.1 the blob is a **binary, machine-endian** format — old text/other formats
  will fail validation.
- A loadable blob must match the **architecture and layer sizes** the library was built with
  (the build's `rnnoise_data.h` fixes the sizes). In practice: only load blobs dumped from the
  same RNNoise source generation as the bundled one (`Project Files/lib/NR_Algorithms_x64/src/rnnoise`).
  A large-architecture DLL loads large-architecture blobs; that is why the two shipped files are
  known-good references.

## 5. Creating new models

Everything needed is bundled under `Project Files/lib/NR_Algorithms_x64/src/rnnoise/` — including
**prebuilt Windows tools** `dump_features.exe` and `dump_weights_blob.exe`, the PyTorch training
code (`torch/rnnoise/`), dataset pointers (`datasets.txt`), and the upstream `README` whose
"Training" section this summarizes. (`src/HowTo/how.txt` has MW0LGE's exact
MSYS2 + VS2022/ClangCL build commands for the DLLs, including the AVX2 variant.)

**You need:** clean speech and noise recordings, both **48 kHz, 16-bit raw PCM** (headerless).
Xiph provides ready-made data: concatenated TTS speech
(`https://media.xiph.org/rnnoise/data/tts_speech_48k.sw`), background/foreground noise files, and
the community noise collection (`rnnoise_contributions.tar.gz`).

**The pipeline:**

```bash
# 1. Mix speech+noise into training features (≥10,000 sequences; 200,000+ recommended)
./dump_features speech.pcm background_noise.pcm foreground_noise.pcm features.f32 <count>
#    (optionally add -rir_list rir_list.txt for reverberation augmentation;
#     scripts/dump_features_parallel.sh parallelizes this)

# 2. Train (PyTorch, torch/rnnoise/) — pick --epochs to reach ~75,000 weight updates
python3 train_rnnoise.py features.f32 output_dir        # → rnnoise_50.pth etc.

# 3. Convert the checkpoint to C source
python3 dump_rnnoise_weights.py --quantize rnnoise_50.pth rnnoise_c
#    → rnnoise_data.c / rnnoise_data.h ; copy into src/ and rebuild RNNoise

# 4. Dump the loadable blob from that build
./dump_weights_blob                                      # → weights_blob.bin
```

The resulting `weights_blob.bin` is what Thetis's **Load** button consumes. (Upstream tip: one
round of denoising your "clean" speech with a trained model and retraining improves results
slightly. `scripts/shrink_model.sh` is used when producing the small-architecture variant.)

**Why a ham would bother:** the stock models learned domestic/office noise. A model trained with
**HF band noise** — recorded static, QRN crashes, power-line buzz, typical QRM — as the noise
corpus, and SSB-bandwidth (≤ ~3 kHz) filtered speech as the clean corpus, can substantially
outperform the default on the air. Record noise off your own radio (48 kHz 16-bit mono raw),
concatenate, and substitute it for `background_noise.pcm` in step 1.

Step 3's rebuild requirement matters: if your training changes layer sizes, the blob will no
longer match the shipped `rnnoise.dll` and Thetis will refuse or misload it — keep the
architecture unchanged (default training config) so the blob stays loadable, or rebuild and
replace `rnnoise.dll` alongside it (build commands in `src/HowTo/how.txt`).

## 6. API quick reference

| Export (wdsp.dll) | Declared | Purpose |
|-------------------|----------|---------|
| `SetRXARNNRRun(channel, run)` | `dsp.cs` L247 | Enable/disable NR3 on a channel (locks DSP, re-checks bandpass). |
| `SetRXARNNRPosition(channel, pos)` | `dsp.cs` L250 | 0 = pre-AGC, 1 = post-AGC (moves `bp1` with it). |
| `SetRXARNNRUseDefaultGain(channel, use)` | `dsp.cs` L256 | 1 = fixed ×500,000 input gain, 0 = adaptive input AGC. |
| `RNNRloadModel(path)` | `dsp.cs` L253 | Global model hot-swap; `""` = built-in default. |

| RNNoise API (rnnoise.h) | Used at | Purpose |
|--------------------------|---------|---------|
| `rnnoise_get_frame_size()` | `create_rnnr` | Frame size (480 samples). |
| `rnnoise_create(model)` | `create_rnnr`, `RNNRloadModel` | New denoiser state (`NULL` = built-in model). |
| `rnnoise_process_frame(st, out, in)` | `xrnnr` | Denoise one frame; returns speech probability (unused here). |
| `rnnoise_model_from_filename(path)` / `rnnoise_model_free()` | `RNNRloadModel` | Load/free a weights blob. The model must outlive every state created from it — which is why all states are torn down before the swap. |

## 7. Related documentation

- [`docs/files/wdsp/rnnr.c.md`](files/wdsp/rnnr.c.md) — generated symbol outline of the wrapper
- [`docs/CODE_OUTLINE.md` §7](CODE_OUTLINE.md#7-wdsp--the-dsp-engine) — NR3's place among the wdsp noise-reduction blocks
- NR4 (`wdsp/sbnr.c`, libspecbleach) shares the same UI slot pattern and the same
  `NR_Algorithms_x64` library folder — the natural comparison when NR3's speech bias is a problem
