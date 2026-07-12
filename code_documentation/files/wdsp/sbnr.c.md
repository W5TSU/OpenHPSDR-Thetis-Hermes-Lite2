# `wdsp/sbnr.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** libspecbleach spectral noise reduction "NR4".

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×6)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_sbnr()` (×1), `destroy_sbnr()` (×1), `xsbnr()` (×1), `setSamplerate_sbnr()` (×1), `setBuffers_sbnr()` (×1), `setSize_sbnr()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`setSize_sbnr()`** — L52 — `void setSize_sbnr (SBNR a, int size)`
  Reconfigures the `sbnr` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setBuffers_sbnr()`** — L61 — `void setBuffers_sbnr (SBNR a, double* in, double* out)`
  Re-points the `sbnr` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`create_sbnr()`** — L67 — `SBNR create_sbnr (int run, int position, int size, double *in, double *out, int rate)`
  Constructor for the `sbnr` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_sbnr()`** — L90 — `void setSamplerate_sbnr(SBNR a, int rate)`
  Reconfigures the `sbnr` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`xsbnr()`** — L97 — `void xsbnr (SBNR a, int pos)`
  Runs the `sbnr` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`destroy_sbnr()`** — L136 — `void destroy_sbnr (SBNR a)`
  Destroys the `sbnr` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`SetRXASBNRRun()`** — L144 — `PORT void SetRXASBNRRun (int channel, int run)`
  Sets rxasbnrrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASBNRreductionAmount()`** — L162 — `PORT void SetRXASBNRreductionAmount (int channel, float amount)`
  Sets the amount of dBs that the noise will be attenuated. It goes from 0 dB to 20 dB */
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASBNRsmoothingFactor()`** — L176 — `PORT void SetRXASBNRsmoothingFactor (int channel, float factor)`
  Percentage of smoothing to apply. Averages the reduction calculation frame per frame so the rate of change is less resulting in less musical noise but if too strong it can blur transient and reduce high frequencies. It goes from 0 to 100 percent */
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASBNRwhiteningFactor()`** — L190 — `PORT void SetRXASBNRwhiteningFactor (int channel, float factor)`
  Percentage of whitening that is going to be applied to the residue of the reduction. It modifies the noise floor to be more like white noise. This can help hide musical noise when the noise is colored. It goes from 0 to 100 percent */
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASBNRnoiseRescale()`** — L205 — `PORT void SetRXASBNRnoiseRescale (int channel, float factor)`
  Strength in which the reduction will be applied. It uses the masking thresholds of the signal to determine where in the spectrum the reduction needs to be stronger. This parameter scales how much in each of the frequencies the reduction is going to be applied. It can be a positive dB value in…
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASBNRpostFilterThreshold()`** — L218 — `PORT void SetRXASBNRpostFilterThreshold (int channel, float threshold)`
  Sets the SNR threshold in dB in which the post-filter will start to blur musical noise. It can be a positive or negative dB value in between -10 dB and 10 dB */
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASBNRnoiseScalingType()`** — L233 — `PORT void SetRXASBNRnoiseScalingType(int channel, int noise_scaling_type)`
  Type of algorithm used to scale noise in order to apply over or under subtraction in different parts of the spectrum while calculating the reduction. 0 is a-posteriori snr scaling using the complete spectrum, 1 is a-posteriori using critical bands and 2 is using masking thresholds
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASBNRPosition()`** — L243 — `PORT void SetRXASBNRPosition(int channel, int position)`
  Sets rxasbnrposition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/sbnr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
