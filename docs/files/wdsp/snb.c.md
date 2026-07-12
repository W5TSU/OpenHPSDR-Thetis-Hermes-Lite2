# `wdsp/snb.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Spectral noise blanker.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×19)
  - `wdsp/nbp.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/nbp.c` (calls ×7)
  - `wdsp/resample.c` (calls ×6)
  - `wdsp/RXA.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×4)
  - `wdsp/lmath.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/firmin.c` (calls ×1)
- Most-referenced symbols from other files: `recalc_bpsnba_filter()` (×4), `create_bpsnba()` (×1), `create_snba()` (×1), `destroy_bpsnba()` (×1), `destroy_snba()` (×1), `flush_bpsnba()` (×1), `flush_snba()` (×1), `xbpsnbain()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_snba()`** — L31 — `void calc_snba (SNBA d)`
  Called by: `create_snba()` (same file), `setBuffers_snba()` (same file), `setSamplerate_snba()` (same file), `setSize_snba()` (same file), `SetRXASNBAovrlp()` (same file)
- **`create_snba()`** — L68 — `SNBA create_snba (int run, double* in, double* out, int inrate, int internalrate, int bsize, int ovrlp, int xsize, int asize, int npasses, double k1, double k2, int b, int pre, int`
  Constructor for the `snba` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`decalc_snba()`** — L121 — `void decalc_snba (SNBA d)`
  Called by: `destroy_snba()` (same file), `setBuffers_snba()` (same file), `setSamplerate_snba()` (same file), `setSize_snba()` (same file), `SetRXASNBAovrlp()` (same file)
- **`destroy_snba()`** — L131 — `void destroy_snba (SNBA d)`
  Destroys the `snba` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_snba()`** — L161 — `void flush_snba (SNBA d)`
  Flushes (zeroes) the `snba` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`setBuffers_snba()`** — L187 — `void setBuffers_snba (SNBA a, double* in, double* out)`
  Re-points the `snba` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_snba()`** — L195 — `void setSamplerate_snba (SNBA a, int rate)`
  Reconfigures the `snba` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_snba()`** — L202 — `void setSize_snba (SNBA a, int size)`
  Reconfigures the `snba` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`ATAc0()`** — L209 — `void ATAc0 (int n, int nr, double* A, double* r)`
  Called by: `xHat()` (same file)
- **`multA1TA2()`** — L218 — `void multA1TA2(double* a1, double* a2, int m, int n, int q, double* c)`
  Called by: `xHat()` (same file)
- **`multXKE()`** — L241 — `void multXKE(double* a, double* xk, int m, int q, int p, double* vout)`
  Called by: `xHat()` (same file)
- **`multAv()`** — L254 — `void multAv(double* a, double* v, int m, int q, double* vout)`
  Called by: `xHat()` (same file)
- **`xHat()`** — L265 — `void xHat(int xusize, int asize, double* xk, double* a, double* xout, double* r, double* ATAI, double* A1, double* A2, double* P1, double* P2, double* trI_y, double* trI_v, double*`
  Called by: `execFrame()` (same file)
- **`invf()`** — L306 — `void invf(int xsize, int asize, double* a, double* x, double* v)`
  Called by: `execFrame()` (same file)
- **`det()`** — L324 — `void det(SNBA d, int asize, double* v, int* detout)`
  Called by: `execFrame()` (same file)
- **`scanFrame()`** — L404 — `int scanFrame(int xsize, int pval, double pmultmin, int* det, int* bimp, int* limp, int* befimp, int* aftimp, int* p_opt, int* next)`
  Called by: `execFrame()` (same file)
- **`execFrame()`** — L492 — `void execFrame(SNBA d, double* x)`
  Called by: `xsnba()` (same file)
- **`xsnba()`** — L539 — `void xsnba (SNBA d)`
  Runs the `snba` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`SetRXASNBARun()`** — L579 — `PORT void SetRXASNBARun (int channel, int run)`
  Sets rxasnbarun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASNBAovrlp()`** — L598 — `PORT void SetRXASNBAovrlp (int channel, int ovrlp)`
  Sets rxasnbaovrlp — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXASNBAasize()`** — L607 — `PORT void SetRXASNBAasize (int channel, int size)`
  Sets rxasnbaasize — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXASNBAnpasses()`** — L614 — `PORT void SetRXASNBAnpasses (int channel, int npasses)`
  Sets rxasnbanpasses — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXASNBAk1()`** — L621 — `PORT void SetRXASNBAk1 (int channel, double k1)`
  Sets rxasnbak1 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASNBAk2()`** — L628 — `PORT void SetRXASNBAk2 (int channel, double k2)`
  Sets rxasnbak2 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASNBAbridge()`** — L635 — `PORT void SetRXASNBAbridge (int channel, int bridge)`
  Sets rxasnbabridge — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXASNBApresamps()`** — L642 — `PORT void SetRXASNBApresamps (int channel, int presamps)`
  Sets rxasnbapresamps — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXASNBApostsamps()`** — L649 — `PORT void SetRXASNBApostsamps (int channel, int postsamps)`
  Sets rxasnbapostsamps — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXASNBApmultmin()`** — L656 — `PORT void SetRXASNBApmultmin (int channel, double pmultmin)`
  Sets rxasnbapmultmin — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXASNBAOutputBandwidth()`** — L663 — `PORT void SetRXASNBAOutputBandwidth (int channel, double flow, double fhigh)`
  Sets rxasnbaoutput bandwidth — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetPassband()` (`wdsp/RXA.c`)
- **`calc_bpsnba()`** — L714 — `void calc_bpsnba (BPSNBA a)`
  Called by: `create_bpsnba()` (same file), `setBuffers_bpsnba()` (same file), `setSamplerate_bpsnba()` (same file), `setSize_bpsnba()` (same file)
- **`create_bpsnba()`** — L736 — `BPSNBA create_bpsnba (int run, int run_notches, int position, int size, int nc, int mp, double* in, double* out, int rate, double abs_low_freq, double abs_high_freq, double f_low, `
  Constructor for the `bpsnba` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`decalc_bpsnba()`** — L763 — `void decalc_bpsnba (BPSNBA a)`
  Called by: `destroy_bpsnba()` (same file), `setBuffers_bpsnba()` (same file), `setSamplerate_bpsnba()` (same file), `setSize_bpsnba()` (same file)
- **`destroy_bpsnba()`** — L769 — `void destroy_bpsnba (BPSNBA a)`
  Destroys the `bpsnba` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_bpsnba()`** — L775 — `void flush_bpsnba (BPSNBA a)`
  Flushes (zeroes) the `bpsnba` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`setBuffers_bpsnba()`** — L781 — `void setBuffers_bpsnba (BPSNBA a, double* in, double* out)`
  Re-points the `bpsnba` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_bpsnba()`** — L789 — `void setSamplerate_bpsnba (BPSNBA a, int rate)`
  Reconfigures the `bpsnba` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_bpsnba()`** — L796 — `void setSize_bpsnba (BPSNBA a, int size)`
  Reconfigures the `bpsnba` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`xbpsnbain()`** — L803 — `void xbpsnbain (BPSNBA a, int position)`
  Runs the `bpsnbain` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`xbpsnbaout()`** — L809 — `void xbpsnbaout (BPSNBA a, int position)`
  Runs the `bpsnbaout` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`recalc_bpsnba_filter()`** — L815 — `void recalc_bpsnba_filter (BPSNBA a, int update)`
  Called by: `RXAbpsnbaCheck()` (`wdsp/RXA.c`), `UpdateNBPFilters()` (`wdsp/nbp.c`), `RXANBPSetWindow()` (`wdsp/nbp.c`), `RXANBPSetAutoIncrease()` (`wdsp/nbp.c`)
- **`RXABPSNBASetNC()`** — L837 — `PORT void RXABPSNBASetNC (int channel, int nc)`
  RXA chain operation — bpsnbaset nc; part of the receive/transmit chain API.
  Called by: `RXASetNC()` (`wdsp/RXA.c`)
- **`RXABPSNBASetMP()`** — L852 — `PORT void RXABPSNBASetMP (int channel, int mp)`
  RXA chain operation — bpsnbaset mp; part of the receive/transmit chain API.
  Called by: `RXASetMP()` (`wdsp/RXA.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/snb.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
