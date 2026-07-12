# `Console/Path_Illustrator.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Interactive block diagram of the whole signal path (what's enabled where, RX/TX routing).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Types

#### `Path_Illustrator` (type, L12)

- `.canvas_Paint()` — L1143
- `.Update_control_settings()` — L1158
- `.draw_HPSDR()` — L1333
- `.draw_HERMES()` — L1914
- `.draw_ANAN_10E()` — L2339
- `.draw_ANAN_100_PA_rev15()` — L2605
- `.draw_ANAN_100_PA_rev24()` — L3112
- `.draw_ANAN_100D_PA_rev15()` — L3566
- `.draw_ANAN_100D_PA_rev24()` — L4279
- `.ADC0_to_Rx0()` — L4903
- `.ADC0_to_Rx1()` — L4908
- `.ADC0_to_Rx2()` — L4915
- `.ADC0_to_Rx3()` — L4922
- `.ADC0_to_Rx4()` — L4929
- `.ADC0_to_Rx5()` — L4936
- `.ADC0_to_Rx6()` — L4943
- `.ADC1_to_ground()` — L4950
- `.ADC1_to_Rx0()` — L4960
- `.ADC1_to_Rx1()` — L4967
- `.ADC1_to_Rx2()` — L4974
- `.ADC1_to_Rx3()` — L4981
- `.ADC1_to_Rx4()` — L4988
- `.ADC1_to_Rx5()` — L4995
- `.ADC1_to_Rx6()` — L5002
- `.ALEX_ANT_to_HPF_B()` — L5016
- `.ALEX_2_ANT_to_HPF_B()` — L5035
- `.ALEX_TX_ANT()` — L5054
- `.ALEX_2_RX_out_to_ADC()` — L5075
- `.AMPF_TX_path_PA10()` — L5082
- `.AMPF_to_PA15()` — L5089
- `.AMPF_TX_path_PA15()` — L5095
- `.AMPF_XVTR_TX()` — L5102
- `.basic_Tx_path()` — L5110
- `.BYPASS_to_ADC0()` — L5125
- `.C2_to_Rx0()` — L5132
- `.C3_to_Rx0()` — L5137
- `.C4_to_HPF_PA15_TX()` — L5145
- `.C4_to_Rx0()` — L5152
- `.C2_to_LPF()` — L5159
- `.C3_to_LPF()` — L5167
- `.C4_to_LPF()` — L5175
- `.C5_to_ADC0()` — L5183
- `.C5_to_HPF_PA15_TX()` — L5190
- `.C6_to_HPF_PA15_TX()` — L5197
- `.C7_to_ground()` — L5205
- `.C9_to_LPF_L()` — L5214
- `.C10_to_LPF_L()` — L5222
- `.C11_to_LPF_L()` — L5228
- `.CODEC2_to_AUDIO_MIXER()` — L5236
- `.draw_diversity_connection()` — L5242
- `.DSP_in1_to_out1_crossconnect()` — L5248
- `.DSP_in2_to_out2_crossconnect()` — L5253
- `.DSP_out1_to_RX1_DISPLAY()` — L5258
- `.DSP_out2_to_RX2_DISPLAY()` — L5265
- `.DUC0_to_Rx1()` — L5270
- `.EXT1_to_HPF_PA15()` — L5276
- `.EXT2_to_HPF_PA15()` — L5282
- `.HEADPHONES_to_CODEC2()` — L5288
- `.HERMES_PA_to_ALEX_LPF()` — L5295
- `.HPF_to_ground()` — L5303
- `.HPSDR_DSP_to_AUDIO_MIXER_input_1()` — L5316
- `.HPSDR_DSP_to_AUDIO_MIXER_input_2()` — L5322
- `.line_to_ground()` — L5328
- `.loopback_to_RX1_DISPLAY()` — L5336
- `.LPF_to_ADC0()` — L5344
- `.LPF_to_C2()` — L5351
- `.LPF_to_C3()` — L5357
- `.LPF_to_C4()` — L5364
- `.LPF_to_HPF_PA15()` — L5371
- `.MERCURY_ADC_to_DDCs()` — L5378
- `.MERCURY_DDC0_to_RX1_DISPLAY()` — L5387
- `.MERCURY_DDC1_to_RX2_DISPLAY()` — L5392
- `.MERCURY_2_ADC_to_DDCs()` — L5397
- `.MERCURY_2_DDC0_to_METIS_diversity()` — L5406
- `.MERCURY_2_DDC0_to_RX2_DISPLAY()` — L5415
- `.MERCURY_2_AUDIO_INPUT()` — L5422
- `.MERCURY_RX1_to_AUDIO_MIXER()` — L5430
- `.MERCURY_RX2_to_AUDIO_MIXER()` — L5441
- `.PA_to_LPF()` — L5451
- `.LPF_to_BYPASS()` — L5456
- `.PA15_to_LPF()` — L5462
- `.PENELOPE_PA_to_ALEX_LPF()` — L5472
- `.PENELOPE_PA_to_DSP()` — L5480
- `.PENELOPE_DSP_to_CODEC()` — L5495
- `.Rx0_to_DSP()` — L5502
- `.Rx1_to_DSP()` — L5510
- `.Rx2_to_DSP()` — L5515
- `.DSP_Rx2_to_RX1_DISPLAY()` — L5520
- `.Rx3_to_DSP()` — L5527
- `.SPKR_to_DSP1()` — L5536
- `.SPKR_to_DSP2()` — L5543
- `.SPKR_to_RX1_DISPLAY()` — L5550
- `.SPKR_to_RX1_DISPLAY_2()` — L5566
- `.SPKR_to_RX2_DISPLAY()` — L5582
- `.SPKR_to_RX2_DISPLAY_2()` — L5598
- `.SWR_to_ADC0()` — L5614
- `.update_labels_PA10()` — L5620
- `.update_labels_PA()` — L5687
- `.XVTR_to_HPF_PA15()` — L5778
- `.pi_Changed()` — L5790
- `.do_platform_prep()` — L5802
- `.hide_all_labels()` — L5814
- `.hide_rear_panel_labels()` — L5958
- `.hide_controls()` — L6000
- `.TX_AUDIO_OUT_2_RX_MODELS()` — L6005
- `.cb_DUAL_MERCURY_ALEX_CheckedChanged()` — L6290
- `.rb_rx_CheckedChanged()` — L6297
- `.rb_tx_CheckedChanged()` — L6304
- `.PI_Resize()` — L6311
- `.PI_Disposed()` — L6318

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Path_Illustrator.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
