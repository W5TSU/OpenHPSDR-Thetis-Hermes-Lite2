# The Thetis Settings Database — and What "Export Database" Saves

This document explains Thetis's settings database: where it lives, what the
**Export Database** function writes, the exact file format, and the schema of
every table in the export.

---

## 1. Where Export Database lives in the UI

Database import/export is handled by the **Database Manager**
(`frmDBMan.cs`), opened from the main window menu: **Setup → Database Manager**.
(Older Thetis versions had Import/Export Database buttons directly on the
Setup → General tab; in this 2.10.x codebase those were replaced by the
Database Manager.)

The manager lists all databases and their backups. Two export paths exist:

| Button | Handler | What it does |
|--------|---------|--------------|
| Export (tooltip "Export the selected database") | `frmDBMan.cs:405` → `DBMan.Export()` (`clsDBMan.cs:1726`) | Exports one of the available databases |
| Export backup | `frmDBMan.cs:427` → `DBMan.ExportBackup()` (`clsDBMan.cs:1778`) | Exports a selected automatic/manual backup |

Both open a `SaveFileDialog` titled **"Export Database"**, defaulting to the
user's **My Documents** folder with a filename of:

```
Thetis_database_export_<description>_<date>_<time>.xml
Thetis_database_export_backup_<description>_<date>_<time>.xml   (backup export)
```

`<description>` is the database's name from its `dbman.json` metadata; the
date/time is the local short date/time with separators replaced by `_`
(`Common.DateTimeStringForFile()`, `common.cs:965`), e.g.
`Thetis_database_export_Default_7_13_2026_3_42_PM.xml`.

An important nuance in `DBMan.Export()`:

- If the selected database is the **active** one, Thetis serializes the live
  in-memory `DataSet` via `DB.WriteDB(filename)` (`database.cs:9540`) — so the
  export reflects the **current session state**, including changes not yet
  flushed to disk.
- If it's a non-active database (or a backup), the export is a **verbatim file
  copy** of that database's `database.xml` (or backup file).

Either way, the exported file is byte-for-byte the same format as the working
database file.

## 2. Where the working data is kept

Each database is a folder under:

```
%AppData%\OpenHPSDR\Thetis-x64\DB\<GUID>\
    database.xml     ← the actual settings database (what export produces)
    dbman.json       ← Database-Manager metadata (description, model, version,
                       backup-on-startup/shutdown flags) — NOT included in export
    backups\         ← timestamped copies of database.xml — NOT included in export
```

`clsDBMan.cs` manages the folders/metadata; `database.cs` (class `DB`) owns the
`DataSet` itself. Note that an export contains **only** `database.xml` content:
the description, hardware model tag, version info shown in the Database Manager
list, and all backups stay behind in `dbman.json`/`backups\`.

## 3. File format

The file is a standard **.NET `System.Data.DataSet` XML serialization with an
inline XSD schema** — written by `ds.WriteXml(fn, XmlWriteMode.WriteSchema)`
(`database.cs:9550`). Structure:

```xml
<?xml version="1.0" standalone="yes"?>
<Data>
  <xs:schema id="Data" xmlns:xs="http://www.w3.org/2001/XMLSchema" ...>
    <!-- inline XSD describing every table and column -->
  </xs:schema>

  <Options>
    <Key>comboRadioModel</Key>
    <Value>HERMES LITE</Value>
  </Options>
  <Options> ... </Options>

  <State>
    <Key>chkPower</Key>
    <Value>False</Value>
  </State>

  <TXProfile>
    <Name>Default</Name>
    <FilterLow>100</FilterLow>
    ...
  </TXProfile>
  ...
</Data>
```

- The DataSet (and therefore the XML root element) is named **`Data`**
  (`database.cs:9474`). On import, Thetis validates `DataSetName == "Data"`
  (`database.cs:11242`) — files with a different root are rejected.
- Each table row is one XML element named after its table; each column is a
  child element. Missing/null columns are simply omitted.
- Everything is text; the inline schema records the declared .NET column types
  (`xs:string`, `xs:int`, `xs:double`, `xs:boolean`).

## 4. What is saved — the tables

Two kinds of tables exist: generic **Key/Value "form" tables** (created on
demand by `DB.SaveVars()` via `AddFormTable()`, `database.cs:742` — two string
columns, `Key` and `Value`) and **structured tables** with typed columns
(created by `VerifyTables()`, `database.cs:199`).

### 4.1 Key/Value tables

| Table | Written by | Contents |
|-------|-----------|----------|
| `Options` | `Setup.saveOptions()` (`setup.cs:1654`) | **Every control on the Setup form**, keyed by control name — checkboxes (`"True"/"False"`), numeric up/downs, combo selections (display text, e.g. `comboRadioModel = HERMES LITE`), text boxes, radio buttons, plus keyed data blobs such as multimeter layouts (`meterData_*`, `meterIGSettings_*`), PA profiles and similar serialized strings. This is by far the largest table (thousands of rows). |
| `State` | `console.SaveState()` (`console.cs:3307`) | **Every control on the main console window** (VFO frequencies, mode/filter/band buttons, sliders like AF/RF gain, window size/position) plus non-control keys: manual notch list (`mnotchdb[i]` = centre/width/active tuples, `console.cs:3000`), band-stack registers per band-button, DB-manager options (`PruneBackups`), and other console state. |
| `WideBand` | `wideband.cs:144` | Settings of the wideband spectrum display form. |

Values are stored via `InvariantCulture` string conversion of each control's
value; on import/startup the same mapping restores each control by name. A key
whose control no longer exists is ignored (and pruned), which is what makes
databases portable across most version upgrades.

### 4.2 Structured tables

Created/verified at startup by `VerifyTables()` (`database.cs:199`):

| Table | Schema (columns) | Contents |
|-------|------------------|----------|
| `BandText` | `Low, High, Name, TX` (double, double, string, bool) | The frequency-vs-band-label/TX-allowed lookup used for the band text display, per FRS region (`database.cs:753`). |
| `Memory` | `GroupID, Freq, ModeID, FilterID, Callsign, Comments, Scan, Squelch, StepSizeID, AGCID, Gain, FilterLow, FilterHigh, CreateDate` | Memory channels (Memory form, `database.cs:4263`). |
| `GroupList` | `GroupID, GroupName` | Memory-channel groups; seeded with AM, FM, SSB, SSTV, CW, PSK, RTTY (`database.cs:4284`). |
| `TXProfile` | `Name` + ~80 columns: TX filter low/high, 10-band TX EQ (+ legacy/parametric EQ data strings), compander/CESSB, mic gain, DEXP, CFC settings, PureSignal options, VAC1/VAC2 audio routing… (`database.cs:4303`) | All user TX profiles selected in Setup → Transmit. |
| `TXProfileDef` | same schema | The factory-default TX profiles (used for reset-to-default). |
| `BandStack2Entries` | `GUID, Description, Locked, Frequency, CentreFrequency, Band, Mode, SubMode, CTUNEnabled, Filter, ZoomFactor, ZoomSlider, PowerLevel, AGCLevel` | The BandStack2 spot entries (`database.cs:342`). |
| `BandStack2Filters` | `GUID, FilterName, FilterDescription, FilterOn* flags, UserDefined, FilterReturnMode, SpecificReturnGUID, CurrentSelected*, LastVisited*` (18 cols) | BandStack2 filter definitions incl. last-visited state per filter (`database.cs:305`). |
| `BandStack2FilterFrequencies` | `FilterGUID, Low, High, LowOnly, Band, BandType, Region` | Frequency ranges attached to a filter. |
| `BandStack2FilterModes` / `...SubModes` / `...Bands` | `FilterGUID` + `Mode`/`SubMode`/`Band` (int) | Mode/submode/band criteria attached to a filter. |
| `BandStack2HiddenEntries` | `FilterGUID, EntryGUID` | Entries hidden from a given filter. |

In short: **an export is the complete radio personality** — every Setup option
(including radio model, PA calibration, attenuation and line-gain settings),
console state, TX profiles, memories, band stacks, band text and manual
notches — in one XML file.

## 5. What is *not* in the export

- `dbman.json` metadata (database description, hardware tag, Thetis version,
  backup flags) and the `backups\` folder.
- Anything stored outside the database: skins (`%AppData%...\Skins`), MIDI
  mappings DB if stored separately, TCI/CAT port state held only in `Options`
  is included, but log/error files are not.
- The DB-manager's own settings (`dbman_settings.json`).

## 6. Importing

The counterpart is the Database Manager's import buttons
(`DBMan.Import()` `clsDBMan.cs:1386`, `ImportAsAvailable()` `clsDBMan.cs:1454`):
the XML is read back with `DataSet.ReadXml`, validated (root must be `Data`;
the stored `comboRadioModel` must be known to this build — otherwise the model
resets to HERMES with a warning, `setup.cs:1810`), merged table-by-table, and a
restart applies it. Because the export is just `database.xml`, an exported file
can also be dropped manually into a `DB\<GUID>\` folder.

## 7. Quick file map

| Concern | File |
|---------|------|
| Export/import/backup UI | `Project Files/Source/Console/frmDBMan.cs` (+ `.Designer.cs`) |
| Export logic, DB folders, `dbman.json` | `Project Files/Source/Console/clsDBMan.cs` (`Export()` at 1726, `ExportBackup()` at 1778) |
| DataSet, tables, XML read/write | `Project Files/Source/Console/database.cs` (`WriteDB` at 9540, `VerifyTables` at 199) |
| Options table writer (Setup controls) | `Project Files/Source/Console/setup.cs` (`saveOptions`) |
| State table writer (console controls) | `Project Files/Source/Console/console.cs` (`SaveState`) |
| Export filename timestamp | `Project Files/Source/Console/common.cs` (`DateTimeStringForFile`, 965) |
