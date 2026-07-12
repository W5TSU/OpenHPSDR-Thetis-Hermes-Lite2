# `wdsp/gen.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal generators (tone, two-tone, noise, sweep) for testing and tune.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_gen()` (×2), `destroy_gen()` (×2), `flush_gen()` (×2), `xgen()` (×2), `setSamplerate_gen()` (×2), `setBuffers_gen()` (×2), `setSize_gen()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_tone()`** — L29 — `void calc_tone (GEN a)`
  Called by: `calc_gen()` (same file), `SetRXAPreGenToneFreq()` (same file), `SetTXAPreGenToneFreq()` (same file), `SetTXAPostGenToneFreq()` (same file)
- **`calc_tt()`** — L37 — `void calc_tt (GEN a)`
  Called by: `calc_gen()` (same file), `SetTXAPostGenTTFreq()` (same file)
- **`calc_sweep()`** — L49 — `void calc_sweep (GEN a)`
  Called by: `calc_gen()` (same file), `SetRXAPreGenSweepFreq()` (same file), `SetRXAPreGenSweepRate()` (same file), `SetTXAPreGenSweepFreq()` (same file), `SetTXAPreGenSweepRate()` (same file), `SetTXAPostGenSweepFreq()` (same file) — and 1 more
- **`calc_sawtooth()`** — L57 — `void calc_sawtooth (GEN a)`
  Called by: `calc_gen()` (same file), `SetTXAPreGenSawtoothFreq()` (same file)
- **`calc_triangle()`** — L64 — `void calc_triangle (GEN a)`
  Called by: `calc_gen()` (same file), `SetTXAPreGenTriangleFreq()` (same file)
- **`calc_pulse()`** — L73 — `void calc_pulse (GEN a)`
  Called by: `calc_gen()` (same file), `SetTXAPreGenPulseFreq()` (same file), `SetTXAPreGenPulseDutyCycle()` (same file), `SetTXAPreGenPulseToneFreq()` (same file), `SetTXAPreGenPulseTransition()` (same file), `SetTXAPostGenPulseFreq()` (same file) — and 3 more
- **`calc_ttpulse()`** — L98 — `void calc_ttpulse(GEN a)`
  Called by: `calc_gen()` (same file), `SetTXAPostGenTTPulseFreq()` (same file), `SetTXAPostGenTTPulseDutyCycle()` (same file), `SetTXAPostGenTTPulseToneFreq()` (same file), `SetTXAPostGenTTPulseTransition()` (same file)
- **`calc_gen()`** — L127 — `void calc_gen (GEN a)`
  Called by: `create_gen()` (same file), `setSamplerate_gen()` (same file)
- **`decalc_gen()`** — L138 — `void decalc_gen (GEN a)`
  Called by: `destroy_gen()` (same file), `setSamplerate_gen()` (same file)
- **`create_gen()`** — L144 — `GEN create_gen (int run, int size, double* in, double* out, int rate, int mode)`
  Constructor for the `gen` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`)
- **`destroy_gen()`** — L195 — `void destroy_gen (GEN a)`
  Destroys the `gen` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_gen()`** — L201 — `void flush_gen (GEN a)`
  Flushes (zeroes) the `gen` block’s internal buffers/state.
  Called by: `setSize_gen()` (same file), `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`)
- **`xgen()`** — L215 — `void xgen (GEN a)`
  Runs the `gen` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_gen()`** — L530 — `void setBuffers_gen (GEN a, double* in, double* out)`
  Re-points the `gen` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_gen()`** — L536 — `void setSamplerate_gen (GEN a, int rate)`
  Reconfigures the `gen` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_gen()`** — L543 — `void setSize_gen (GEN a, int size)`
  Reconfigures the `gen` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetRXAPreGenRun()`** — L558 — `PORT void SetRXAPreGenRun (int channel, int run)`
  Sets rxapre gen run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPreGenMode()`** — L566 — `PORT void SetRXAPreGenMode (int channel, int mode)`
  Sets rxapre gen mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPreGenToneMag()`** — L574 — `PORT void SetRXAPreGenToneMag (int channel, double mag)`
  Sets rxapre gen tone mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPreGenToneFreq()`** — L582 — `PORT void SetRXAPreGenToneFreq (int channel, double freq)`
  Sets rxapre gen tone freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPreGenNoiseMag()`** — L591 — `PORT void SetRXAPreGenNoiseMag (int channel, double mag)`
  Sets rxapre gen noise mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPreGenSweepMag()`** — L599 — `PORT void SetRXAPreGenSweepMag (int channel, double mag)`
  Sets rxapre gen sweep mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPreGenSweepFreq()`** — L607 — `PORT void SetRXAPreGenSweepFreq (int channel, double freq1, double freq2)`
  Sets rxapre gen sweep freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPreGenSweepRate()`** — L617 — `PORT void SetRXAPreGenSweepRate (int channel, double rate)`
  Sets rxapre gen sweep rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenRun()`** — L635 — `PORT void SetTXAPreGenRun (int channel, int run)`
  Sets txapre gen run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenMode()`** — L643 — `PORT void SetTXAPreGenMode (int channel, int mode)`
  Sets txapre gen mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenToneMag()`** — L651 — `PORT void SetTXAPreGenToneMag (int channel, double mag)`
  Sets txapre gen tone mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenToneFreq()`** — L659 — `PORT void SetTXAPreGenToneFreq (int channel, double freq)`
  Sets txapre gen tone freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenNoiseMag()`** — L668 — `PORT void SetTXAPreGenNoiseMag (int channel, double mag)`
  Sets txapre gen noise mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenSweepMag()`** — L676 — `PORT void SetTXAPreGenSweepMag (int channel, double mag)`
  Sets txapre gen sweep mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenSweepFreq()`** — L684 — `PORT void SetTXAPreGenSweepFreq (int channel, double freq1, double freq2)`
  Sets txapre gen sweep freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenSweepRate()`** — L694 — `PORT void SetTXAPreGenSweepRate (int channel, double rate)`
  Sets txapre gen sweep rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenSawtoothMag()`** — L703 — `PORT void SetTXAPreGenSawtoothMag (int channel, double mag)`
  Sets txapre gen sawtooth mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenSawtoothFreq()`** — L711 — `PORT void SetTXAPreGenSawtoothFreq (int channel, double freq)`
  Sets txapre gen sawtooth freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenTriangleMag()`** — L720 — `PORT void SetTXAPreGenTriangleMag (int channel, double mag)`
  Sets txapre gen triangle mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenTriangleFreq()`** — L728 — `PORT void SetTXAPreGenTriangleFreq (int channel, double freq)`
  Sets txapre gen triangle freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenPulseMag()`** — L737 — `PORT void SetTXAPreGenPulseMag (int channel, double mag)`
  Sets txapre gen pulse mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenPulseFreq()`** — L745 — `PORT void SetTXAPreGenPulseFreq (int channel, double freq)`
  Sets txapre gen pulse freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenPulseDutyCycle()`** — L754 — `PORT void SetTXAPreGenPulseDutyCycle (int channel, double dc)`
  Sets txapre gen pulse duty cycle — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenPulseToneFreq()`** — L763 — `PORT void SetTXAPreGenPulseToneFreq (int channel, double freq)`
  Sets txapre gen pulse tone freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPreGenPulseTransition()`** — L772 — `PORT void SetTXAPreGenPulseTransition (int channel, double transtime)`
  Sets txapre gen pulse transition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenRun()`** — L783 — `PORT void SetTXAPostGenRun (int channel, int run)`
  Sets txapost gen run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenMode()`** — L791 — `PORT void SetTXAPostGenMode (int channel, int mode)`
  Sets txapost gen mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenToneMag()`** — L799 — `PORT void SetTXAPostGenToneMag (int channel, double mag)`
  Sets txapost gen tone mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenToneFreq()`** — L807 — `PORT void SetTXAPostGenToneFreq (int channel, double freq)`
  Sets txapost gen tone freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTMag()`** — L816 — `PORT void SetTXAPostGenTTMag (int channel, double mag1, double mag2)`
  Sets txapost gen ttmag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTFreq()`** — L825 — `PORT void SetTXAPostGenTTFreq (int channel, double freq1, double freq2)`
  Sets txapost gen ttfreq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenSweepMag()`** — L835 — `PORT void SetTXAPostGenSweepMag (int channel, double mag)`
  Sets txapost gen sweep mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenSweepFreq()`** — L843 — `PORT void SetTXAPostGenSweepFreq (int channel, double freq1, double freq2)`
  Sets txapost gen sweep freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenSweepRate()`** — L853 — `PORT void SetTXAPostGenSweepRate (int channel, double rate)`
  Sets txapost gen sweep rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenPulseMag()`** — L862 — `PORT void SetTXAPostGenPulseMag(int channel, double mag)`
  Sets txapost gen pulse mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenPulseFreq()`** — L870 — `PORT void SetTXAPostGenPulseFreq(int channel, double freq)`
  Sets txapost gen pulse freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenPulseDutyCycle()`** — L879 — `PORT void SetTXAPostGenPulseDutyCycle(int channel, double dc)`
  Sets txapost gen pulse duty cycle — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenPulseToneFreq()`** — L888 — `PORT void SetTXAPostGenPulseToneFreq(int channel, double freq)`
  Sets txapost gen pulse tone freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenPulseTransition()`** — L897 — `PORT void SetTXAPostGenPulseTransition(int channel, double transtime)`
  Sets txapost gen pulse transition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenPulseIQout()`** — L906 — `PORT void SetTXAPostGenPulseIQout(int channel, int IQout)`
  Sets txapost gen pulse iqout — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTPulseMag()`** — L914 — `PORT void SetTXAPostGenTTPulseMag(int channel, double mag1, double mag2)`
  Sets txapost gen ttpulse mag — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTPulseFreq()`** — L925 — `PORT void SetTXAPostGenTTPulseFreq(int channel, double freq)`
  Sets txapost gen ttpulse freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTPulseDutyCycle()`** — L934 — `PORT void SetTXAPostGenTTPulseDutyCycle(int channel, double dc)`
  Sets txapost gen ttpulse duty cycle — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTPulseToneFreq()`** — L943 — `PORT void SetTXAPostGenTTPulseToneFreq(int channel, double freq1, double freq2)`
  Sets txapost gen ttpulse tone freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTPulseTransition()`** — L954 — `PORT void SetTXAPostGenTTPulseTransition(int channel, double transtime)`
  Sets txapost gen ttpulse transition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPostGenTTPulseIQout()`** — L963 — `PORT void SetTXAPostGenTTPulseIQout(int channel, int IQout)`
  Sets txapost gen ttpulse iqout — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/gen.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
