# `Console/Path_Illustrator.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Interactive block diagram of the whole signal path (what's enabled where, RX/TX routing).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Path_Illustrator` (type, L12)

- **`.canvas_Paint()`** — L1143 — `private void canvas_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `canvas` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Update_control_settings()`** — L1158 — `private void Update_control_settings()`
  Called by: `.canvas_Paint()` (same file), `.draw_HPSDR()` (same file), `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file) — and 2 more
- **`.draw_HPSDR()`** — L1333 — `private void draw_HPSDR()`
  Called by: `.canvas_Paint()` (same file), `.do_platform_prep()` (same file)
- **`.draw_HERMES()`** — L1914 — `private void draw_HERMES()`
  Called by: `.canvas_Paint()` (same file), `.do_platform_prep()` (same file)
- **`.draw_ANAN_10E()`** — L2339 — `private void draw_ANAN_10E()`
  Called by: `.canvas_Paint()` (same file), `.do_platform_prep()` (same file)
- **`.draw_ANAN_100_PA_rev15()`** — L2605 — `private void draw_ANAN_100_PA_rev15()`
  Called by: `.canvas_Paint()` (same file), `.do_platform_prep()` (same file)
- **`.draw_ANAN_100_PA_rev24()`** — L3112 — `private void draw_ANAN_100_PA_rev24()`
  Called by: `.canvas_Paint()` (same file), `.do_platform_prep()` (same file)
- **`.draw_ANAN_100D_PA_rev15()`** — L3566 — `private void draw_ANAN_100D_PA_rev15()`
  Called by: `.canvas_Paint()` (same file), `.do_platform_prep()` (same file)
- **`.draw_ANAN_100D_PA_rev24()`** — L4279 — `private void draw_ANAN_100D_PA_rev24()`
  Called by: `.canvas_Paint()` (same file), `.do_platform_prep()` (same file)
- **`.ADC0_to_Rx0()`** — L4903 — `private void ADC0_to_Rx0(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC0_to_Rx1()`** — L4908 — `private void ADC0_to_Rx1(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC0_to_Rx2()`** — L4915 — `private void ADC0_to_Rx2(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC0_to_Rx3()`** — L4922 — `private void ADC0_to_Rx3(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC0_to_Rx4()`** — L4929 — `private void ADC0_to_Rx4(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC0_to_Rx5()`** — L4936 — `private void ADC0_to_Rx5(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC0_to_Rx6()`** — L4943 — `private void ADC0_to_Rx6(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC1_to_ground()`** — L4950 — `private void ADC1_to_ground(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file)
- **`.ADC1_to_Rx0()`** — L4960 — `private void ADC1_to_Rx0(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC1_to_Rx1()`** — L4967 — `private void ADC1_to_Rx1(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC1_to_Rx2()`** — L4974 — `private void ADC1_to_Rx2(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC1_to_Rx3()`** — L4981 — `private void ADC1_to_Rx3(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC1_to_Rx4()`** — L4988 — `private void ADC1_to_Rx4(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC1_to_Rx5()`** — L4995 — `private void ADC1_to_Rx5(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ADC1_to_Rx6()`** — L5002 — `private void ADC1_to_Rx6(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.ALEX_ANT_to_HPF_B()`** — L5016 — `private void ALEX_ANT_to_HPF_B(Pen pen)`
  Called by: `.draw_HPSDR()` (same file), `.draw_HERMES()` (same file)
- **`.ALEX_2_ANT_to_HPF_B()`** — L5035 — `private void ALEX_2_ANT_to_HPF_B(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.ALEX_TX_ANT()`** — L5054 — `private void ALEX_TX_ANT(Pen pen)`
  Called by: `.draw_HPSDR()` (same file), `.draw_HERMES()` (same file)
- **`.ALEX_2_RX_out_to_ADC()`** — L5075 — `private void ALEX_2_RX_out_to_ADC(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.AMPF_TX_path_PA10()`** — L5082 — `private void AMPF_TX_path_PA10(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.AMPF_to_PA15()`** — L5089 — `private void AMPF_to_PA15(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.AMPF_TX_path_PA15()`** — L5095 — `private void AMPF_TX_path_PA15(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AMPF_XVTR_TX()`** — L5102 — `private void AMPF_XVTR_TX(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.basic_Tx_path()`** — L5110 — `private void basic_Tx_path(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.BYPASS_to_ADC0()`** — L5125 — `private void BYPASS_to_ADC0(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.C2_to_Rx0()`** — L5132 — `private void C2_to_Rx0(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.C3_to_Rx0()`** — L5137 — `private void C3_to_Rx0(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.C4_to_HPF_PA15_TX()`** — L5145 — `private void C4_to_HPF_PA15_TX(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.C4_to_Rx0()`** — L5152 — `private void C4_to_Rx0(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.C2_to_LPF()`** — L5159 — `private void C2_to_LPF(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.C3_to_LPF()`** — L5167 — `private void C3_to_LPF(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.C4_to_LPF()`** — L5175 — `private void C4_to_LPF(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.C5_to_ADC0()`** — L5183 — `private void C5_to_ADC0(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.C5_to_HPF_PA15_TX()`** — L5190 — `private void C5_to_HPF_PA15_TX(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file)
- **`.C6_to_HPF_PA15_TX()`** — L5197 — `private void C6_to_HPF_PA15_TX(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file)
- **`.C7_to_ground()`** — L5205 — `private void C7_to_ground(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file)
- **`.C9_to_LPF_L()`** — L5214 — `private void C9_to_LPF_L(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.C10_to_LPF_L()`** — L5222 — `private void C10_to_LPF_L(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.C11_to_LPF_L()`** — L5228 — `private void C11_to_LPF_L(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.CODEC2_to_AUDIO_MIXER()`** — L5236 — `private void CODEC2_to_AUDIO_MIXER(Pen pen)`
  Called by: `.SPKR_to_RX1_DISPLAY()` (same file), `.SPKR_to_RX1_DISPLAY_2()` (same file), `.SPKR_to_RX2_DISPLAY()` (same file), `.SPKR_to_RX2_DISPLAY_2()` (same file)
- **`.draw_diversity_connection()`** — L5242 — `private void draw_diversity_connection(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.DSP_in1_to_out1_crossconnect()`** — L5248 — `private void DSP_in1_to_out1_crossconnect(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.DSP_in2_to_out2_crossconnect()`** — L5253 — `private void DSP_in2_to_out2_crossconnect(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.DSP_out1_to_RX1_DISPLAY()`** — L5258 — `private void DSP_out1_to_RX1_DISPLAY(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.DSP_out2_to_RX2_DISPLAY()`** — L5265 — `private void DSP_out2_to_RX2_DISPLAY(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.DUC0_to_Rx1()`** — L5270 — `private void DUC0_to_Rx1(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.EXT1_to_HPF_PA15()`** — L5276 — `private void EXT1_to_HPF_PA15(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.EXT2_to_HPF_PA15()`** — L5282 — `private void EXT2_to_HPF_PA15(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.HEADPHONES_to_CODEC2()`** — L5288 — `private void HEADPHONES_to_CODEC2(Pen pen)`
  Called by: `.SPKR_to_RX1_DISPLAY()` (same file), `.SPKR_to_RX1_DISPLAY_2()` (same file), `.SPKR_to_RX2_DISPLAY()` (same file), `.SPKR_to_RX2_DISPLAY_2()` (same file)
- **`.HERMES_PA_to_ALEX_LPF()`** — L5295 — `private void HERMES_PA_to_ALEX_LPF(Pen pen)`
  Called by: `.draw_HERMES()` (same file)
- **`.HPF_to_ground()`** — L5303 — `private void HPF_to_ground(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file)
- **`.HPSDR_DSP_to_AUDIO_MIXER_input_1()`** — L5316 — `private void HPSDR_DSP_to_AUDIO_MIXER_input_1(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.HPSDR_DSP_to_AUDIO_MIXER_input_2()`** — L5322 — `private void HPSDR_DSP_to_AUDIO_MIXER_input_2(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.line_to_ground()`** — L5328 — `private void line_to_ground(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.loopback_to_RX1_DISPLAY()`** — L5336 — `private void loopback_to_RX1_DISPLAY(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file)
- **`.LPF_to_ADC0()`** — L5344 — `private void LPF_to_ADC0(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LPF_to_C2()`** — L5351 — `private void LPF_to_C2(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.LPF_to_C3()`** — L5357 — `private void LPF_to_C3(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.LPF_to_C4()`** — L5364 — `private void LPF_to_C4(Pen pen)`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.LPF_to_HPF_PA15()`** — L5371 — `private void LPF_to_HPF_PA15(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.MERCURY_ADC_to_DDCs()`** — L5378 — `private void MERCURY_ADC_to_DDCs(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_DDC0_to_RX1_DISPLAY()`** — L5387 — `private void MERCURY_DDC0_to_RX1_DISPLAY(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_DDC1_to_RX2_DISPLAY()`** — L5392 — `private void MERCURY_DDC1_to_RX2_DISPLAY(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_2_ADC_to_DDCs()`** — L5397 — `private void MERCURY_2_ADC_to_DDCs(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_2_DDC0_to_METIS_diversity()`** — L5406 — `private void MERCURY_2_DDC0_to_METIS_diversity(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_2_DDC0_to_RX2_DISPLAY()`** — L5415 — `private void MERCURY_2_DDC0_to_RX2_DISPLAY(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_2_AUDIO_INPUT()`** — L5422 — `private void MERCURY_2_AUDIO_INPUT(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_RX1_to_AUDIO_MIXER()`** — L5430 — `private void MERCURY_RX1_to_AUDIO_MIXER(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.MERCURY_RX2_to_AUDIO_MIXER()`** — L5441 — `private void MERCURY_RX2_to_AUDIO_MIXER(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.PA_to_LPF()`** — L5451 — `private void PA_to_LPF(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file)
- **`.LPF_to_BYPASS()`** — L5456 — `private void LPF_to_BYPASS(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file)
- **`.PA15_to_LPF()`** — L5462 — `private void PA15_to_LPF(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.PENELOPE_PA_to_ALEX_LPF()`** — L5472 — `private void PENELOPE_PA_to_ALEX_LPF(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.PENELOPE_PA_to_DSP()`** — L5480 — `private void PENELOPE_PA_to_DSP(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.PENELOPE_DSP_to_CODEC()`** — L5495 — `private void PENELOPE_DSP_to_CODEC(Pen pen)`
  Called by: `.draw_HPSDR()` (same file)
- **`.Rx0_to_DSP()`** — L5502 — `private void Rx0_to_DSP(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.Rx1_to_DSP()`** — L5510 — `private void Rx1_to_DSP(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.Rx2_to_DSP()`** — L5515 — `private void Rx2_to_DSP(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.DSP_Rx2_to_RX1_DISPLAY()`** — L5520 — `private void DSP_Rx2_to_RX1_DISPLAY(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.Rx3_to_DSP()`** — L5527 — `private void Rx3_to_DSP(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.SPKR_to_DSP1()`** — L5536 — `private void SPKR_to_DSP1(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SPKR_to_DSP2()`** — L5543 — `private void SPKR_to_DSP2(Pen pen)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SPKR_to_RX1_DISPLAY()`** — L5550 — `private void SPKR_to_RX1_DISPLAY(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file) — and 1 more
- **`.SPKR_to_RX1_DISPLAY_2()`** — L5566 — `private void SPKR_to_RX1_DISPLAY_2(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file), `.TX_AUDIO_OUT_2_RX_MODELS()` (same file)
- **`.SPKR_to_RX2_DISPLAY()`** — L5582 — `private void SPKR_to_RX2_DISPLAY(Pen pen)`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.TX_AUDIO_OUT_2_RX_MODELS()` (same file)
- **`.SPKR_to_RX2_DISPLAY_2()`** — L5598 — `private void SPKR_to_RX2_DISPLAY_2(Pen pen)`
  Called by: `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.SWR_to_ADC0()`** — L5614 — `private void SWR_to_ADC0(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.update_labels_PA10()`** — L5620 — `private void update_labels_PA10()`
  Called by: `.draw_ANAN_10E()` (same file)
- **`.update_labels_PA()`** — L5687 — `private void update_labels_PA()`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.XVTR_to_HPF_PA15()`** — L5778 — `private void XVTR_to_HPF_PA15(Pen pen)`
  Called by: `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file), `.draw_ANAN_100D_PA_rev24()` (same file)
- **`.pi_Changed()`** — L5790 — `public void pi_Changed()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.do_platform_prep()`** — L5802 — `private void do_platform_prep()`
  Called by: `.pi_Changed()` (same file), `.cb_DUAL_MERCURY_ALEX_CheckedChanged()` (same file), `.rb_rx_CheckedChanged()` (same file), `.rb_tx_CheckedChanged()` (same file), `.PI_Resize()` (same file)
- **`.hide_all_labels()`** — L5814 — `private void hide_all_labels()`
  Called by: `.draw_HPSDR()` (same file), `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file) — and 1 more
- **`.hide_rear_panel_labels()`** — L5958 — `private void hide_rear_panel_labels()`
  Called by: `.update_labels_PA10()` (same file), `.update_labels_PA()` (same file)
- **`.hide_controls()`** — L6000 — `private void hide_controls()`
  Called by: `.draw_HPSDR()` (same file), `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file), `.draw_ANAN_100D_PA_rev15()` (same file) — and 1 more
- **`.TX_AUDIO_OUT_2_RX_MODELS()`** — L6005 — `private void TX_AUDIO_OUT_2_RX_MODELS()`
  Called by: `.draw_HERMES()` (same file), `.draw_ANAN_10E()` (same file), `.draw_ANAN_100_PA_rev15()` (same file), `.draw_ANAN_100_PA_rev24()` (same file)
- **`.cb_DUAL_MERCURY_ALEX_CheckedChanged()`** — L6290 — `private void cb_DUAL_MERCURY_ALEX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `cb_DUAL_MERCURY_ALEX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rb_rx_CheckedChanged()`** — L6297 — `private void rb_rx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `rb_rx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rb_tx_CheckedChanged()`** — L6304 — `private void rb_tx_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `rb_tx` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PI_Resize()`** — L6311 — `private void PI_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `PI` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.PI_Disposed()`** — L6318 — `private void PI_Disposed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Path_Illustrator.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
