# `wdsp/eq.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Graphic/parametric equalizer.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×9)
  - `wdsp/TXA.c` (calls ×9)
  - `wdsp/fmsq.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×22)
  - `wdsp/utilities.c` (calls ×10)
  - `wdsp/fir.c` (calls ×3)
  - `wdsp/impulse_cache.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_eqp()` (×2), `destroy_eqp()` (×2), `flush_eqp()` (×2), `xeqp()` (×2), `setSamplerate_eqp()` (×2), `setBuffers_eqp()` (×2), `setSize_eqp()` (×2), `eq_impulse()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`fEQcompare()`** — L54 — `int fEQcompare (const void * a, const void * b)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`fEQcompare3()`** — L64 — `static int fEQcompare3(const void* a, const void* b)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`eq_impulse()`** — L74 — `double* eq_impulse(int N, int nfreqs, double* F, double* G, double* Q, double samplerate, double scale, int ctfmode, int wintype)`
  Called by: `create_eqp()` (same file), `setSamplerate_eqp()` (same file), `setSize_eqp()` (same file), `SetRXAEQNC()` (same file), `SetRXAEQProfile()` (same file), `SetRXAEQCtfmode()` (same file) — and 12 more
- **`create_eqp()`** — L508 — `EQP create_eqp (int run, int size, int nc, int mp, double *in, double *out, int nfreqs, double* F, double* G, int ctfmode, int wintype, int samplerate)`
  Constructor for the `eqp` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`)
- **`destroy_eqp()`** — L534 — `void destroy_eqp (EQP a)`
  Destroys the `eqp` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_eqp()`** — L540 — `void flush_eqp (EQP a)`
  Flushes (zeroes) the `eqp` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`)
- **`xeqp()`** — L545 — `void xeqp (EQP a)`
  Runs the `eqp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_eqp()`** — L553 — `void setBuffers_eqp (EQP a, double* in, double* out)`
  Re-points the `eqp` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_eqp()`** — L560 — `void setSamplerate_eqp (EQP a, int rate)`
  Reconfigures the `eqp` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_eqp()`** — L569 — `void setSize_eqp (EQP a, int size)`
  Reconfigures the `eqp` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetRXAEQRun()`** — L585 — `PORT void SetRXAEQRun (int channel, int run)`
  Sets rxaeqrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEQNC()`** — L593 — `PORT void SetRXAEQNC (int channel, int nc)`
  Sets rxaeqnc — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetNC()` (`wdsp/RXA.c`)
- **`SetRXAEQMP()`** — L610 — `PORT void SetRXAEQMP (int channel, int mp)`
  Sets rxaeqmp — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetMP()` (`wdsp/RXA.c`)
- **`SetRXAEQProfile()`** — L622 — `PORT void SetRXAEQProfile (int channel, int nfreqs, double* F, double* G, double* Q)`
  Sets rxaeqprofile — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEQCtfmode()`** — L651 — `PORT void SetRXAEQCtfmode (int channel, int mode)`
  Sets rxaeqctfmode — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAEQWintype()`** — L663 — `PORT void SetRXAEQWintype (int channel, int wintype)`
  Sets rxaeqwintype — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAGrphEQ()`** — L675 — `PORT void SetRXAGrphEQ (int channel, int *rxeq)`
  Sets rxagrph eq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAGrphEQ10()`** — L703 — `PORT void SetRXAGrphEQ10 (int channel, int *rxeq)`
  Sets rxagrph eq10 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAEQRun()`** — L742 — `PORT void SetTXAEQRun (int channel, int run)`
  Sets txaeqrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAEQNC()`** — L750 — `PORT void SetTXAEQNC (int channel, int nc)`
  Sets txaeqnc — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetNC()` (`wdsp/TXA.c`)
- **`SetTXAEQMP()`** — L767 — `PORT void SetTXAEQMP (int channel, int mp)`
  Sets txaeqmp — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetMP()` (`wdsp/TXA.c`)
- **`SetTXAEQProfile()`** — L779 — `PORT void SetTXAEQProfile (int channel, int nfreqs, double* F, double* G, double* Q)`
  Sets txaeqprofile — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAEQCtfmode()`** — L807 — `PORT void SetTXAEQCtfmode (int channel, int mode)`
  Sets txaeqctfmode — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAEQWintype()`** — L819 — `PORT void SetTXAEQWintype (int channel, int wintype)`
  Sets txaeqwintype — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAGrphEQ()`** — L831 — `PORT void SetTXAGrphEQ (int channel, int *txeq)`
  Sets txagrph eq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAGrphEQ10()`** — L859 — `PORT void SetTXAGrphEQ10 (int channel, int *txeq)`
  Sets txagrph eq10 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`eq_mults()`** — L898 — `double* eq_mults (int size, int nfreqs, double* F, double* G, double* Q, double samplerate, double scale, int ctfmode, int wintype)`
  Called by: `calc_eq()` (same file)
- **`calc_eq()`** — L906 — `void calc_eq (EQ a)`
  Called by: `create_eq()` (same file), `setBuffers_eq()` (same file), `setSamplerate_eq()` (same file), `setSize_eq()` (same file)
- **`decalc_eq()`** — L916 — `void decalc_eq (EQ a)`
  Called by: `destroy_eq()` (same file), `setBuffers_eq()` (same file), `setSamplerate_eq()` (same file), `setSize_eq()` (same file)
- **`create_eq()`** — L925 — `EQ create_eq (int run, int size, double *in, double *out, int nfreqs, double* F, double* G, int ctfmode, int wintype, int samplerate)`
  Constructor for the `eq` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_eq()`** — L944 — `void destroy_eq (EQ a)`
  Destroys the `eq` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_eq()`** — L953 — `void flush_eq (EQ a)`
  Flushes (zeroes) the `eq` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xeq()`** — L958 — `void xeq (EQ a)`
  Runs the `eq` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_eq()`** — L980 — `void setBuffers_eq (EQ a, double* in, double* out)`
  Re-points the `eq` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_eq()`** — L988 — `void setSamplerate_eq (EQ a, int rate)`
  Reconfigures the `eq` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_eq()`** — L995 — `void setSize_eq (EQ a, int size)`
  Reconfigures the `eq` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/eq.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
