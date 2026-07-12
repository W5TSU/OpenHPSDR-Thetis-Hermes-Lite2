# `wdsp/iir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** IIR biquad sections (notches, peaking filters) and double-pole building blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×14)
  - `wdsp/apfshadow.c` (calls ×8)
  - `wdsp/TXA.c` (calls ×7)
  - `wdsp/fmd.c` (calls ×6)
  - `wdsp/ssql.c` (calls ×4)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×14)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRXABiQuadBandwidth()` (×2), `SetRXABiQuadFreq()` (×2), `SetRXABiQuadGain()` (×2), `SetRXABiQuadRun()` (×2), `create_mpeak()` (×1), `create_speak()` (×1), `destroy_mpeak()` (×1), `destroy_speak()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_snotch()`** — L35 — `void calc_snotch (SNOTCH a)`
  Called by: `create_snotch()` (same file), `setSamplerate_snotch()` (same file), `SetSNCTCSSFreq()` (same file)
- **`create_snotch()`** — L50 — `SNOTCH create_snotch (int run, int size, double* in, double* out, int rate, double f, double bw)`
  Constructor for the `snotch` block: allocates its state/buffers and computes initial coefficients.
  Called by: `calc_fmd()` (`wdsp/fmd.c`)
- **`destroy_snotch()`** — L65 — `void destroy_snotch (SNOTCH a)`
  Destroys the `snotch` block, freeing its allocated buffers.
  Called by: `decalc_fmd()` (`wdsp/fmd.c`)
- **`flush_snotch()`** — L71 — `void flush_snotch (SNOTCH a)`
  Flushes (zeroes) the `snotch` block’s internal buffers/state.
  Called by: `calc_snotch()` (same file), `setSize_snotch()` (same file), `flush_fmd()` (`wdsp/fmd.c`)
- **`xsnotch()`** — L76 — `void xsnotch (SNOTCH a)`
  Runs the `snotch` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xfmd()` (`wdsp/fmd.c`)
- **`setBuffers_snotch()`** — L97 — `void setBuffers_snotch (SNOTCH a, double* in, double* out)`
  Re-points the `snotch` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_snotch()`** — L103 — `void setSamplerate_snotch (SNOTCH a, int rate)`
  Reconfigures the `snotch` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_snotch()`** — L109 — `void setSize_snotch (SNOTCH a, int size)`
  Reconfigures the `snotch` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetSNCTCSSFreq()`** — L121 — `void SetSNCTCSSFreq (SNOTCH a, double freq)`
  Sets snctcssfreq — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXACTCSSFreq()` (`wdsp/fmd.c`)
- **`SetSNCTCSSRun()`** — L129 — `void SetSNCTCSSRun (SNOTCH a, int run)`
  Sets snctcssrun — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXACTCSSRun()` (`wdsp/fmd.c`)
- **`calc_speak()`** — L143 — `void calc_speak (SPEAK a)`
  Called by: `create_speak()` (same file), `setSamplerate_speak()` (same file), `SetRXABiQuadFreq()` (same file), `SetRXABiQuadBandwidth()` (same file), `SetRXABiQuadGain()` (same file), `SetRXAmpeakFilFreq()` (same file) — and 2 more
- **`create_speak()`** — L218 — `SPEAK create_speak (int run, int size, double* in, double* out, int rate, double f, double bw, double gain, int nstages, int design)`
  Constructor for the `speak` block: allocates its state/buffers and computes initial coefficients.
  Called by: `calc_mpeak()` (same file), `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_speak()`** — L242 — `void destroy_speak (SPEAK a)`
  Destroys the `speak` block, freeing its allocated buffers.
  Called by: `decalc_mpeak()` (same file), `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_speak()`** — L254 — `void flush_speak (SPEAK a)`
  Flushes (zeroes) the `speak` block’s internal buffers/state.
  Called by: `calc_speak()` (same file), `setSize_speak()` (same file), `flush_mpeak()` (same file), `flush_rxa()` (`wdsp/RXA.c`)
- **`xspeak()`** — L264 — `void xspeak (SPEAK a)`
  Runs the `speak` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xmpeak()` (same file), `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_speak()`** — L297 — `void setBuffers_speak (SPEAK a, double* in, double* out)`
  Re-points the `speak` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_speak()`** — L303 — `void setSamplerate_speak (SPEAK a, int rate)`
  Reconfigures the `speak` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_speak()`** — L309 — `void setSize_speak (SPEAK a, int size)`
  Reconfigures the `speak` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXABiQuadRun()`** — L321 — `PORT void SetRXABiQuadRun (int channel, int run)`
  Sets rxabi quad run — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWRun()` (`wdsp/apfshadow.c`)
- **`SetRXABiQuadFreq()`** — L330 — `PORT void SetRXABiQuadFreq (int channel, double freq)`
  Sets rxabi quad freq — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWFreq()` (`wdsp/apfshadow.c`)
- **`SetRXABiQuadBandwidth()`** — L340 — `PORT void SetRXABiQuadBandwidth (int channel, double bw)`
  Sets rxabi quad bandwidth — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWBandwidth()` (`wdsp/apfshadow.c`)
- **`SetRXABiQuadGain()`** — L350 — `PORT void SetRXABiQuadGain (int channel, double gain)`
  Sets rxabi quad gain — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWGain()` (`wdsp/apfshadow.c`)
- **`calc_mpeak()`** — L366 — `void calc_mpeak (MPEAK a)`
  Called by: `create_mpeak()` (same file), `setBuffers_mpeak()` (same file), `setSamplerate_mpeak()` (same file), `setSize_mpeak()` (same file)
- **`decalc_mpeak()`** — L386 — `void decalc_mpeak (MPEAK a)`
  Called by: `destroy_mpeak()` (same file), `setBuffers_mpeak()` (same file), `setSamplerate_mpeak()` (same file), `setSize_mpeak()` (same file)
- **`create_mpeak()`** — L395 — `MPEAK create_mpeak (int run, int size, double* in, double* out, int rate, int npeaks, int* enable, double* f, double* bw, double* gain, int nstages)`
  Constructor for the `mpeak` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_mpeak()`** — L419 — `void destroy_mpeak (MPEAK a)`
  Destroys the `mpeak` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_mpeak()`** — L431 — `void flush_mpeak (MPEAK a)`
  Flushes (zeroes) the `mpeak` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xmpeak()`** — L438 — `void xmpeak (MPEAK a)`
  Runs the `mpeak` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_mpeak()`** — L461 — `void setBuffers_mpeak (MPEAK a, double* in, double* out)`
  Re-points the `mpeak` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_mpeak()`** — L469 — `void setSamplerate_mpeak (MPEAK a, int rate)`
  Reconfigures the `mpeak` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_mpeak()`** — L476 — `void setSize_mpeak (MPEAK a, int size)`
  Reconfigures the `mpeak` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAmpeakRun()`** — L489 — `PORT void SetRXAmpeakRun (int channel, int run)`
  Sets rxampeak run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAmpeakNpeaks()`** — L498 — `PORT void SetRXAmpeakNpeaks (int channel, int npeaks)`
  Sets rxampeak npeaks — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAmpeakFilEnable()`** — L507 — `PORT void SetRXAmpeakFilEnable (int channel, int fil, int enable)`
  Sets rxampeak fil enable — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAmpeakFilFreq()`** — L516 — `PORT void SetRXAmpeakFilFreq (int channel, int fil, double freq)`
  Sets rxampeak fil freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAmpeakFilBw()`** — L527 — `PORT void SetRXAmpeakFilBw (int channel, int fil, double bw)`
  Sets rxampeak fil bw — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAmpeakFilGain()`** — L538 — `PORT void SetRXAmpeakFilGain (int channel, int fil, double gain)`
  Sets rxampeak fil gain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`calc_phrot()`** — L556 — `void calc_phrot (PHROT a)`
  Called by: `create_phrot()` (same file), `setSamplerate_phrot()` (same file), `SetTXAPHROTCorner()` (same file), `SetTXAPHROTNstages()` (same file)
- **`create_phrot()`** — L569 — `PHROT create_phrot (int run, int size, double* in, double* out, int rate, double fc, int nstages)`
  Constructor for the `phrot` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`decalc_phrot()`** — L585 — `void decalc_phrot (PHROT a)`
  Called by: `destroy_phrot()` (same file), `setSamplerate_phrot()` (same file), `SetTXAPHROTCorner()` (same file), `SetTXAPHROTNstages()` (same file)
- **`destroy_phrot()`** — L593 — `void destroy_phrot (PHROT a)`
  Destroys the `phrot` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_phrot()`** — L600 — `void flush_phrot (PHROT a)`
  Flushes (zeroes) the `phrot` block’s internal buffers/state.
  Called by: `setSize_phrot()` (same file), `SetTXAPHROTRun()` (same file), `flush_txa()` (`wdsp/TXA.c`)
- **`xphrot()`** — L608 — `void xphrot (PHROT a)`
  Runs the `phrot` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_phrot()`** — L639 — `void setBuffers_phrot (PHROT a, double* in, double* out)`
  Re-points the `phrot` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_phrot()`** — L645 — `void setSamplerate_phrot (PHROT a, int rate)`
  Reconfigures the `phrot` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_phrot()`** — L652 — `void setSize_phrot (PHROT a, int size)`
  Reconfigures the `phrot` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXAPHROTRun()`** — L664 — `PORT void SetTXAPHROTRun (int channel, int run)`
  Sets txaphrotrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPHROTCorner()`** — L674 — `PORT void SetTXAPHROTCorner (int channel, double corner)`
  Sets txaphrotcorner — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPHROTNstages()`** — L685 — `PORT void SetTXAPHROTNstages (int channel, int nstages)`
  Sets txaphrotnstages — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPHROTReverse()`** — L696 — `PORT void SetTXAPHROTReverse (int channel, int reverse)`
  Sets txaphrotreverse — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`calc_bqlp()`** — L711 — `void calc_bqlp(BQLP a)`
  Called by: `create_bqlp()` (same file), `setSamplerate_bqlp()` (same file)
- **`create_bqlp()`** — L726 — `BQLP create_bqlp(int run, int size, double* in, double* out, double rate, double fc, double Q, double gain, int nstages)`
  Constructor for the `bqlp` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_bqlp()`** — L749 — `void destroy_bqlp(BQLP a)`
  Destroys the `bqlp` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_bqlp()`** — L761 — `void flush_bqlp(BQLP a)`
  Flushes (zeroes) the `bqlp` block’s internal buffers/state.
  Called by: `calc_bqlp()` (same file), `setSize_bqlp()` (same file)
- **`xbqlp()`** — L771 — `void xbqlp(BQLP a)`
  Runs the `bqlp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_bqlp()`** — L804 — `void setBuffers_bqlp(BQLP a, double* in, double* out)`
  Re-points the `bqlp` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_bqlp()`** — L810 — `void setSamplerate_bqlp(BQLP a, int rate)`
  Reconfigures the `bqlp` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_bqlp()`** — L816 — `void setSize_bqlp(BQLP a, int size)`
  Reconfigures the `bqlp` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`calc_dbqlp()`** — L828 — `void calc_dbqlp(BQLP a)`
  Called by: `create_dbqlp()` (same file), `setSamplerate_dbqlp()` (same file)
- **`create_dbqlp()`** — L843 — `BQLP create_dbqlp(int run, int size, double* in, double* out, double rate, double fc, double Q, double gain, int nstages)`
  Constructor for the `dbqlp` block: allocates its state/buffers and computes initial coefficients.
  Called by: `calc_ssql()` (`wdsp/ssql.c`)
- **`destroy_dbqlp()`** — L866 — `void destroy_dbqlp(BQLP a)`
  Destroys the `dbqlp` block, freeing its allocated buffers.
  Called by: `decalc_ssql()` (`wdsp/ssql.c`)
- **`flush_dbqlp()`** — L878 — `void flush_dbqlp(BQLP a)`
  Flushes (zeroes) the `dbqlp` block’s internal buffers/state.
  Called by: `calc_dbqlp()` (same file), `setSize_dbqlp()` (same file), `flush_ssql()` (`wdsp/ssql.c`)
- **`xdbqlp()`** — L887 — `void xdbqlp(BQLP a)`
  Runs the `dbqlp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xssql()` (`wdsp/ssql.c`)
- **`setBuffers_dbqlp()`** — L917 — `void setBuffers_dbqlp(BQLP a, double* in, double* out)`
  Re-points the `dbqlp` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_dbqlp()`** — L923 — `void setSamplerate_dbqlp(BQLP a, int rate)`
  Reconfigures the `dbqlp` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_dbqlp()`** — L929 — `void setSize_dbqlp(BQLP a, int size)`
  Reconfigures the `dbqlp` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`calc_bqbp()`** — L942 — `void calc_bqbp(BQBP a)`
  Called by: `create_bqbp()` (same file), `setSamplerate_bqbp()` (same file)
- **`create_bqbp()`** — L961 — `BQBP create_bqbp(int run, int size, double* in, double* out, double rate, double f_low, double f_high, double gain, int nstages)`
  Constructor for the `bqbp` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_bqbp()`** — L984 — `void destroy_bqbp(BQBP a)`
  Destroys the `bqbp` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_bqbp()`** — L996 — `void flush_bqbp(BQBP a)`
  Flushes (zeroes) the `bqbp` block’s internal buffers/state.
  Called by: `calc_bqbp()` (same file), `setSize_bqbp()` (same file)
- **`xbqbp()`** — L1006 — `void xbqbp(BQBP a)`
  Runs the `bqbp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_bqbp()`** — L1039 — `void setBuffers_bqbp(BQBP a, double* in, double* out)`
  Re-points the `bqbp` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_bqbp()`** — L1045 — `void setSamplerate_bqbp(BQBP a, int rate)`
  Reconfigures the `bqbp` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_bqbp()`** — L1051 — `void setSize_bqbp(BQBP a, int size)`
  Reconfigures the `bqbp` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`calc_dbqbp()`** — L1063 — `void calc_dbqbp(BQBP a)`
  Called by: `create_dbqbp()` (same file), `setSamplerate_dbqbp()` (same file)
- **`create_dbqbp()`** — L1082 — `BQBP create_dbqbp(int run, int size, double* in, double* out, double rate, double f_low, double f_high, double gain, int nstages)`
  Constructor for the `dbqbp` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_dbqbp()`** — L1105 — `void destroy_dbqbp(BQBP a)`
  Destroys the `dbqbp` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_dbqbp()`** — L1117 — `void flush_dbqbp(BQBP a)`
  Flushes (zeroes) the `dbqbp` block’s internal buffers/state.
  Called by: `calc_dbqbp()` (same file), `setSize_dbqbp()` (same file)
- **`xdbqbp()`** — L1126 — `void xdbqbp(BQBP a)`
  Runs the `dbqbp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_dbqbp()`** — L1156 — `void setBuffers_dbqbp(BQBP a, double* in, double* out)`
  Re-points the `dbqbp` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_dbqbp()`** — L1162 — `void setSamplerate_dbqbp(BQBP a, int rate)`
  Reconfigures the `dbqbp` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_dbqbp()`** — L1168 — `void setSize_dbqbp(BQBP a, int size)`
  Reconfigures the `dbqbp` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`calc_sphp()`** — L1180 — `void calc_sphp(SPHP a)`
  Called by: `create_sphp()` (same file), `setSamplerate_sphp()` (same file)
- **`create_sphp()`** — L1193 — `SPHP create_sphp(int run, int size, double* in, double* out, double rate, double fc, int nstages)`
  Constructor for the `sphp` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`decalc_sphp()`** — L1208 — `void decalc_sphp(SPHP a)`
  Called by: `destroy_sphp()` (same file), `setSamplerate_sphp()` (same file)
- **`destroy_sphp()`** — L1216 — `void destroy_sphp(SPHP a)`
  Destroys the `sphp` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_sphp()`** — L1223 — `void flush_sphp(SPHP a)`
  Flushes (zeroes) the `sphp` block’s internal buffers/state.
  Called by: `setSize_sphp()` (same file)
- **`xsphp()`** — L1231 — `void xsphp(SPHP a)`
  Runs the `sphp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_sphp()`** — L1260 — `void setBuffers_sphp(SPHP a, double* in, double* out)`
  Re-points the `sphp` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_sphp()`** — L1266 — `void setSamplerate_sphp(SPHP a, int rate)`
  Reconfigures the `sphp` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_sphp()`** — L1273 — `void setSize_sphp(SPHP a, int size)`
  Reconfigures the `sphp` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`calc_dsphp()`** — L1285 — `void calc_dsphp(SPHP a)`
  Called by: `create_dsphp()` (same file), `setSamplerate_dsphp()` (same file)
- **`create_dsphp()`** — L1298 — `SPHP create_dsphp(int run, int size, double* in, double* out, double rate, double fc, int nstages)`
  Constructor for the `dsphp` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`decalc_dsphp()`** — L1313 — `void decalc_dsphp(SPHP a)`
  Called by: `destroy_dsphp()` (same file), `setSamplerate_dsphp()` (same file)
- **`destroy_dsphp()`** — L1321 — `void destroy_dsphp(SPHP a)`
  Destroys the `dsphp` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_dsphp()`** — L1328 — `void flush_dsphp(SPHP a)`
  Flushes (zeroes) the `dsphp` block’s internal buffers/state.
  Called by: `setSize_dsphp()` (same file)
- **`xdsphp()`** — L1336 — `void xdsphp(SPHP a)`
  Runs the `dsphp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_dsphp()`** — L1362 — `void setBuffers_dsphp(SPHP a, double* in, double* out)`
  Re-points the `dsphp` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_dsphp()`** — L1368 — `void setSamplerate_dsphp(SPHP a, int rate)`
  Reconfigures the `dsphp` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_dsphp()`** — L1375 — `void setSize_dsphp(SPHP a, int size)`
  Reconfigures the `dsphp` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/iir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
