# Hardware (H/W) Selection in Thetis — and How the Hermes-Lite 2 Is Defined

This document explains how Thetis models radio hardware: where the H/W Selection
data lives, how a selection is created and persisted, what it takes to add a new
radio, and exactly how the Hermes-Lite 2 (HL2) is wired into the system.

The UI entry point is **Setup → General → H/W Select** (`tpGeneralHardware`), whose
main control is the **Radio Model** dropdown (`comboRadioModel`).

---

## 1. Where is the H/W Selection data kept?

There is no data file or table of radios. The set of supported radios is defined
**in code**, spread across a small number of files, and the user's *choice* is
persisted in the settings database.

### 1.1 The two enums — `Project Files/Source/Console/enums.cs`

Thetis separates *marketed radio model* from *FPGA board type*:

| Enum | Line | Meaning |
|------|------|---------|
| `HPSDRModel` | `enums.cs:109` | The product the user owns (ANAN-100D, HERMES LITE, RED-PITAYA…). This is what the user picks in Setup. |
| `HPSDRHW` | `enums.cs:389` | The underlying FPGA board / firmware family (Hermes, Angelia, OrionMKII, HermesLite, Saturn…). Several models share one board — e.g. ANAN-7000DLE, ANAN-8000DLE, Anvelina-Pro3 and Red Pitaya are all `OrionMKII`. |

`HPSDRHW` values are **wire-protocol board IDs** — they must match the device-type
byte the radio firmware reports in its discovery reply (`HermesLite = 6`,
`Saturn = 10`, …). `HPSDRModel` values are ordinal and persisted, hence the
warning in the enum: *add new items before `LAST`, never reorder*.

### 1.2 The per-model behavior table — `Project Files/Source/Console/clsHardwareSpecific.cs`

The static class `HardwareSpecific` is the single source of truth for "what does
this model imply". It holds the current `Model` and `Hardware` and centralizes
model-specific defaults:

- `Model` setter (`clsHardwareSpecific.cs:76`) — the big switch that, for each
  `HPSDRModel`, configures ADC count (`NetworkIO.SetRxADC`), MkII band-pass
  filter presence, ADC supply voltage, L/R audio swap, and maps the model to its
  `HPSDRHW` board.
- `StringModelToEnum` / `EnumModelToString` (`clsHardwareSpecific.cs:352,392`) —
  the canonical display-string ↔ enum mapping (these strings are what the combo
  box shows and what gets saved to the database).
- Capability/default queries: `HasVolts`, `HasAmps`, `PSDefaultPeak` (PureSignal
  feedback level), `RXMeterCalbrationOffsetDefaults`, `DefaultPAGainsForBands`,
  `HasSteppedAttenuation`, `SupportsPathIllustrator`, `HasAudioAmplifier`.

### 1.3 The dropdown itself — `Project Files/Source/Console/setup.designer.cs`

`comboRadioModel.Items` is a hard-coded string list (`setup.designer.cs:8615`):

```
HERMES, HERMES LITE, ANAN-10, ANAN-10E, ANAN-100, ANAN-100B, ANAN-100D,
ANAN-200D, ANAN-7000DLE, ANAN-8000DLE, ANAN-G2E, ANAN-G2, ANAN-G2-1K,
ANVELINA-PRO3, RED-PITAYA
```

Each string must round-trip through `StringModelToEnum`/`EnumModelToString`.

### 1.4 Where the user's choice is persisted

Setup-form control values are saved by control name into the **`Options` table**
of the settings database — an XML `DataSet` stored at
`%AppData%\OpenHPSDR\Thetis-x64\DB\<GUID>\database.xml`, managed by
`clsDBMan.cs` and read/written by `database.cs` (`DB.GetVarsDictionary("Options")`,
`DB.SaveVars`). The saved key is literally `comboRadioModel` with the display
string as its value (see `Setup.getModelFromDB()` at `setup.cs:1721`).

On load, `getOptions()` validates the stored string against the combo's item
list; an unknown model (e.g. from a newer Thetis DB) pops a warning and resets
to HERMES (`setup.cs:1810`).

### 1.5 Discovery-time data — `Project Files/Source/Console/HPSDR/clsRadioDiscovery.cs`

Independently of the user's selection, radios announce themselves via UDP
discovery. The reply's board-ID byte is mapped to `HPSDRHW`:

- Protocol 1: `mapP1DeviceType()` (`clsRadioDiscovery.cs:1236`) — byte 10 of the
  reply; `boardId == 6 → HPSDRHW.HermesLite`.
- Protocol 2: byte 11 is cast directly to `HPSDRHW`.

The discovered board is shown in the radio picker (`ucRadioList.cs`,
`clsDiscoveredRadioPicker.cs`), and Setup shows a mismatch warning icon
(`picModelBoardWarning`, `setup.cs:34413`) if the selected model's board doesn't
match what was discovered.

---

## 2. How is a selection created (what happens when you pick a model)?

1. User picks a string in **Setup → General → H/W Select → Radio Model**.
2. `comboRadioModel_SelectedIndexChanged` fires (`setup.cs:20204`):
   - Converts the string via `HardwareSpecific.StringModelToEnum()`.
   - Assigns `HardwareSpecific.Model = new_model;` → the setter in
     `clsHardwareSpecific.cs` configures ADCs/BPF/supply/audio-swap and sets
     `HardwareSpecific.Hardware` to the matching `HPSDRHW`.
   - Calls `console.SetupForHPSDRModel()` (`console.cs:14794`) to adjust main
     console UI (RX2 preamp presence, DX button, …).
   - Runs a giant per-model `switch` that shows/hides/renames Setup controls
     (attenuator limits, antenna labels, Alex/Apollo panels, ADC→DDC routing
     radio buttons, extra tabs, tooltips…).
3. Elsewhere, the model drives runtime behavior:
   - `cmaster.cs` (~line 605) — DDC stream/router layout per model.
   - `NetworkIO.cs`, `Penny.cs`, `Alex.cs` — protocol and control-register
     details.
   - `audio.cs`, `PSForm.cs`, `MeterManager.cs`, `display.cs` — sample-rate,
     PureSignal, metering and display specifics.
4. When Setup is saved, `comboRadioModel`'s text is written to the `Options`
   table in `database.xml` like any other control.

---

## 3. How do you add a new radio?

Checklist, in dependency order (grep for an existing model such as `REDPITAYA`
or `HERMESLITE` to find every touch point — contributors tag their additions
with a callsign comment, e.g. `//DH1KLM`, `//N1GP G2E added`, `// MI0BOT: HL2`):

1. **`enums.cs`** — add a `HPSDRModel` member **immediately before `LAST`**
   (never reorder existing members; the int values are persisted). If the radio
   uses a new FPGA board, also add an `HPSDRHW` member whose value equals the
   board ID its firmware reports in discovery replies.
2. **`clsHardwareSpecific.cs`**:
   - Add a case to the `Model` setter switch: ADC count, MkII BPF flag, ADC
     supply, L/R audio swap, and the `HPSDRHW` mapping.
   - Add entries to `StringModelToEnum()` and `EnumModelToString()` with the
     display string.
   - Extend the capability properties as appropriate: `HasVolts`, `HasAmps`,
     `GetDefaultVoltCalibration`, `PSDefaultPeak`, RX meter/display calibration
     offsets, `DefaultPAGainsForBands` (per-band PA attenuation defaults),
     `HasSteppedAttenuation`, `SupportsPathIllustrator`, `HasAudioAmplifier`.
3. **`setup.designer.cs`** — add the display string to `comboRadioModel.Items`
   (usually done in the VS designer).
4. **`setup.cs`** — add a case to the `switch` in
   `comboRadioModel_SelectedIndexChanged` configuring every Setup control the
   radio needs (attenuator ranges, antenna checkbox labels, Alex/Apollo/Orion
   panels, ADC assignments, extra option tabs). Also sweep the file for other
   per-model switches (band edges, firmware checks, Apollo/Alex handling).
5. **`console.cs`** — add a case to `SetupForHPSDRModel()` (RX2 preamp, DX
   button, duplex) and check other `HPSDRModel` switches in the file.
6. **`cmaster.cs`** — add the model to the correct DDC/router configuration
   group (1/2/4/7-DDC layouts).
7. **Discovery** — if it's a new board ID: extend `mapP1DeviceType()` in
   `clsRadioDiscovery.cs` (Protocol 1) — Protocol 2 casts the byte directly, so
   the `HPSDRHW` value itself must match — and add display handling in
   `ucRadioList.cs` / `clsDiscoveredRadioPicker.cs` (version-string formatting,
   model matching).
8. **Sweep the rest** — `grep -rn "HPSDRModel\." "Project Files/Source/Console"`
   and review every switch: `NetworkIO.cs`, `Penny.cs`, `Alex.cs`, `audio.cs`,
   `PSForm.cs`, `MeterManager.cs`, `display.cs`, `CAT/CATCommands.cs`,
   `Path_Illustrator.cs`, `frmAbout.cs`. Most have sensible `default:` branches,
   but power/attenuation/protocol cases usually need explicit handling.

There is deliberately **no** plug-in or data-driven mechanism: a new radio is a
code change across these files.

---

## 4. How is the Hermes-Lite 2 defined?

The HL2 was added throughout the codebase by **MI0BOT** (all sites tagged
`// MI0BOT: HL2` or similar).

### 4.1 Identity

| Item | Value | Where |
|------|-------|-------|
| Model enum | `HPSDRModel.HERMESLITE` | `enums.cs:124` |
| Board enum / discovery ID | `HPSDRHW.HermesLite = 6` | `enums.cs:395` |
| Display string | `"HERMES LITE"` (combo item; `"HERMES-LITE"` also accepted on read) | `clsHardwareSpecific.cs:380-383,422` |
| Protocol | HPSDR Protocol 1 (USB framing over UDP); discovery reply board-ID byte `6` | `clsRadioDiscovery.cs:1243` |

### 4.2 Hardware profile (`HardwareSpecific.Model` setter, `clsHardwareSpecific.cs:98`)

- 1 RX ADC (`SetRxADC(1)`), no MkII BPF, ADC supply 3.3 V, L/R audio swapped —
  identical low-level profile to HERMES, but mapped to `HPSDRHW.HermesLite`.
- `HasVolts` / `HasAmps` = true — the HL2 reports temperature and PA current,
  displayed via the (renamed) "Show Temp/Current" console option.
- `PSDefaultPeak` = **0.233** (PureSignal feedback default, vs 0.4072 for other
  Protocol-1 radios).
- `DefaultPAGainsForBands`: HF bands default to **100** (i.e. full attenuation /
  no drive) — the HL2's 5 W PA is calibrated per-band by the user rather than
  shipping ANAN-style defaults; 6 m and VHF default to 38.8.
- `HasSteppedAttenuation`: RX1 yes (as all radios), RX2 yes (default branch).

### 4.3 Discovery extras (`clsRadioDiscovery.cs:1172`)

For HL2 replies, Thetis parses extra fields beyond the standard Protocol-1
discovery data: the fixed-IP address stored in EEPROM (bytes 13–16), the EEPROM
config bytes (11–12), number of receivers (19) and a beta/sub-version byte
(21). `ucRadioList.cs:1974` formats the HL2 gateware version as
`major.minor.beta` (e.g. `7.2.5`).

### 4.4 Setup UI specialization (`setup.cs:20308`, case `HPSDRModel.HERMESLITE`)

Selecting HERMES LITE repurposes large parts of the Setup form:

- Adds the HL2-only options tab `tpHL2Options` (removed for all other models).
- Attenuator: forced on, range remapped to the HL2's LNA gain model
  (`udHermesStepAttenuatorData.Maximum = 31`, TX ATT minimum −28 dB).
- TX Tune power becomes a −16.5…0 dB slider in 0.5 dB steps (the HL2 drive is
  an attenuation).
- Max frequency set to 38.4 MHz (HL2 ADC Nyquist with its 76.8 MHz clock).
- Apollo controls are re-labeled for the HL2 PA: "PA Control", "Enable PA",
  "Enable Full Duplex"; Penny tab becomes "Hermes Lite Control"; the
  `chkHERCULES` checkbox becomes "N2ADR Filter" (presets the IO pins for the
  N2ADR filter board); Alex tab becomes "Ant/Filters" with RX-only antenna rows.
- All seven DDCs are pinned to ADC0 (single-ADC radio).
- Enables the HL2 I/O-board checkbox (`chkHL2IOBoardPresent`) and the IO-pin
  LED strip (6 bits).
- Console "8000DLE" meter group is retitled "Hermes-Lite" with
  "Show Temp/Current".

### 4.5 HL2-specific runtime code

- **`Console/HPSDR/IoBoardHl2.cs`** — the core HL2 fork feature: register-level
  control of the HL2 I/O board (antenna tuner control, TX frequency bytes,
  fault detection, control registers). This file exists only in this fork.
- `cmaster.cs:605` — HL2 shares the 4-DDC router layout with HERMES/ANAN-10/100.
- `setup.cs` contains ~40 further `HERMESLITE` branches (band-plan limits,
  firmware-version checks, power-limit tables around `setup.cs:5500-5761`, PA
  calibration behavior).

---

## 5. Quick file map

| Concern | File |
|---------|------|
| Model & board enums | `Project Files/Source/Console/enums.cs` |
| Per-model behavior/defaults | `Project Files/Source/Console/clsHardwareSpecific.cs` |
| H/W Select tab & combo | `Project Files/Source/Console/setup.cs`, `setup.designer.cs` |
| Console per-model setup | `Project Files/Source/Console/console.cs` (`SetupForHPSDRModel`) |
| Persistence (Options table, `database.xml`) | `Project Files/Source/Console/database.cs`, `clsDBMan.cs` |
| UDP discovery & board-ID mapping | `Project Files/Source/Console/HPSDR/clsRadioDiscovery.cs` |
| Discovered-radio picker | `Project Files/Source/Console/ucRadioList.cs`, `clsDiscoveredRadioPicker.cs` |
| DDC stream routing per model | `Project Files/Source/Console/cmaster.cs` |
| HL2 I/O board (fork-specific) | `Project Files/Source/Console/HPSDR/IoBoardHl2.cs` |
