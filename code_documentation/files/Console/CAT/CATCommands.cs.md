# `Console/CAT/CATCommands.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Implements the CAT command set (ZZxx extended + Kenwood TS-2000 subset) — 399-edge god node touching most console state.

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/CATParser.cs` (calls ×350)
  - `Console/Midi2CatCommands.cs` (calls ×273, references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×7)
  - `Console/Memory/MemoryRecord.cs` (references ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/CAT/CATParser.cs` (references ×1)
  - `Console/clsBandStackManager.cs` (calls ×1)
  - `Console/MeterManager.cs` (calls ×1)
  - `Console/common.cs` (calls ×1)
  - `Console/SortableBindingList.cs` (references ×1)
- Most-referenced symbols from other files: `.ZZMD()` (×24), `.ZZME()` (×17), `.ZZBS()` (×14), `.ZZBT()` (×14), `.ZZKM()` (×10), `.ZZAC()` (×5), `.ZZAY()` (×5), `.ZZFA()` (×5)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CATCommands` (type, L48)

- **`.AG()`** — L88 — `public string AG(string s)`
  Sets or reads the Audio Gain control
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.AI()`** — L109 — `public string AI(string s)`
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.BD()`** — L138 — `public string BD()`
  Moves one band down from the currently selected band write only
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.BU()`** — L147 — `public string BU()`
  Moves one band up from the currently selected band write only
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.CN()`** — L155 — `public string CN(string s)`
  Reads or sets the CTCSS frequency
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.CT()`** — L161 — `public string CT(string s)`
  Reads or sets the CTCSS enable button
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.DN()`** — L167 — `public string DN()`
  Moves the VFO A frequency by the step size set on the console
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.FA()`** — L193 — `public string FA(string s)`
  Sets or reads the frequency of VFO A
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.FB()`** — L209 — `public string FB(string s)`
  Sets or reads the frequency of VFO B
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.FR()`** — L227 — `public string FR(string s)`
  Sets VFO A to control rx this is a dummy command to keep other software happy since the SDR-1000 always uses VFO A for rx
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.FT()`** — L241 — `public string FT(string s, bool bFromCatDirect = false)`
  Sets or reads VFO B to control tx another "happiness" command
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.FW()`** — L266 — `public string FW(string s)`
  Sets or reads the DSP filter width OBSOLETE
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.GT()`** — L285 — `public string GT(string s)`
  Sets or reads the AGC constant this is a wrapper that calls ZZGT
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.ID()`** — L295 — `public string ID()`
  Reads the transceiver ID number this needs changing when 3rd party folks on line.
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.IF()`** — L321 — `public string IF()`
  Reads the transceiver status needs work in the split area
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.KS()`** — L482 — `public string KS(string s)`
  Sets or reads the CWX CW speed
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.KY()`** — L516 — `public string KY(string s)`
  Sends text data to CWX for conversion to Morse
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.MD()`** — L580 — `public string MD(string s)`
  Sets or reads the transceiver mode
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.MG()`** — L611 — `public string MG(string s)`
  Sets or reads the Mic Gain thumbwheel
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.MO()`** — L636 — `public string MO(string s)`
  Sets or reads the Monitor status
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.NB()`** — L664 — `public string NB(string s)`
  Sets or reads the Noise Blanker 1 status
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.NT()`** — L683 — `public string NT(string s)`
  Sets or reads the Automatic Notch Filter status
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.OF()`** — L702 — `public string OF(string s)`
  Sets or reads the FM repeater offset frequency
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.OS()`** — L708 — `public string OS(string s)`
  Sets or reads the repeater offset direction
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.PC()`** — L715 — `public string PC(string s)`
  Sets or reads the PA output thumbwheel
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.PR()`** — L738 — `public string PR(string s)`
  Sets or reads the Speech Compressor status Reactivated 10/21/2012 for HRD compatibility BT
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.PS()`** — L751 — `public string PS(string s)`
  Sets or reads the console power on/off status
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.QI()`** — L777 — `public string QI()`
  Sets the Quick Memory with the current contents of VFO A
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.RC()`** — L790 — `public string RC()`
  Clears the RIT value write only
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.RD()`** — L798 — `public string RD(string s)`
  Decrements RIT
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.RT()`** — L805 — `public string RT(string s)`
  Sets or reads the RIT status (on/off)
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.RU()`** — L831 — `public string RU(string s)`
  Increments RIT
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.RX()`** — L839 — `public string RX(string s)`
  Sets or reads the transceiver receive mode status write only but spec shows an answer parameter for a read???
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.SH()`** — L847 — `public string SH(string s)`
  Sets or reads the variable DSP filter high side
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.SL()`** — L880 — `public string SL(string s)`
  Sets or reads the variable DSP filter low side
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.SM()`** — L913 — `public string SM(string s)`
  Reads the S Meter value //TODO modify to consider console.S9Frequency
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.SQ()`** — L952 — `public string SQ(string s)`
  Sets or reads the Squelch value
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.TX()`** — L987 — `public string TX(string s)`
  Sets the transmitter on, write only will eventually need eiter Commander change or ZZ code since it is not CAT compliant as it is
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.UP()`** — L995 — `public string UP()`
  Moves the VFO A frequency up by the step size set on the console
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.XT()`** — L1022 — `public string XT(string s)`
  Sets or reads the transceiver XIT status (on/off)
  Called by: `.Get()` (`Console/CAT/CATParser.cs`)
- **`.ZZAA()`** — L1055 — `public string ZZAA(string s)`
  -W2PA Sets or reads the APF gain (A for amplitude since G is taken)
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.APFGain()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAB()`** — L1090 — `public string ZZAB(string s)`
  -W2PA Sets or reads the APF bandwidth
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.APFBandwidth()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAC()`** — L1126 — `public string ZZAC(string s)`
  Sets or reads the console step size (also see zzst(read only)
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TuningStepUp()` (`Console/Midi2CatCommands.cs`), `.TuningStepDown()` (`Console/Midi2CatCommands.cs`), `.ChangeFreqVfoA()` (`Console/Midi2CatCommands.cs`), `.ChangeFreqVfoB()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAD()`** — L1150 — `public string ZZAD(string s)`
  Sets VFO A down nn Tune Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MoveVFOADown100Khz()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAE()`** — L1169 — `public string ZZAE(string s)`
  Sets VFO A down nn Pre-set Tune Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZAF()`** — L1193 — `public string ZZAF(string s)`
  Sets VFO A up nn pre-set Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZAG()`** — L1216 — `public string ZZAG(string s)`
  Sets or reads the SDR-1000 Audio Gain control
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.SetAFGain()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAI()`** — L1240 — `public string ZZAI(string s)`
  Called by: `.AI()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZAP()`** — L1267 — `public string ZZAP(string s)`
  -W2PA Sets or reads the APF button on/off status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.APF_OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAR()`** — L1290 — `public string ZZAR(string s)`
  Sets or reads the AGC RF gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.AGCLevel()` (`Console/Midi2CatCommands.cs`), `.AGCLevel_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAS()`** — L1328 — `public string ZZAS(string s)`
  Sets or reads the RX2 AGC-T
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2AGCLevel()` (`Console/Midi2CatCommands.cs`), `.RX2AGCLevel_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAT()`** — L1363 — `public string ZZAT(string s)`
  -W2PA Sets or reads the APF tune
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.APFFreq()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAU()`** — L1399 — `public string ZZAU(string s)`
  Sets VFO A up nn Tune Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MoveVFOAUp100Khz()` (`Console/Midi2CatCommands.cs`)
- **`.ZZAY()`** — L1417 — `public string ZZAY(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.APFType_doublepole()` (`Console/Midi2CatCommands.cs`), `.APFType_matched()` (`Console/Midi2CatCommands.cs`), `.APFType_gaussian()` (`Console/Midi2CatCommands.cs`), `.APFType_biquad()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBA()`** — L1447 — `public string ZZBA()`
  Moves the RX2 bandswitch down one band
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2BandUp()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBB()`** — L1454 — `public string ZZBB()`
  Moves the RX2 bandswitch up one band
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2BandDown()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBD()`** — L1462 — `public string ZZBD()`
  Moves the RX1 bandswitch down one band
  Called by: `.BD()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.BandDown()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBE()`** — L1469 — `public string ZZBE(string s)`
  Sets VFO B down nn Pre-set Tune Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZBF()`** — L1493 — `public string ZZBF(string s)`
  Sets VFO B up nn pre-set Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZBG()`** — L1516 — `public string ZZBG(string s)`
  Sets the Band Group (HF/VHF)
  Called by: `.SetBandGroup()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZBI()`** — L1534 — `public string ZZBI(string s)`
  Sets or reads the BIN button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.BinauralOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBM()`** — L1552 — `public string ZZBM(string s)`
  Sets VFO B down nn Tune Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MoveVFOBDown100Khz()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBP()`** — L1571 — `public string ZZBP(string s)`
  Sets VFO B up nn Tune Steps
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MoveVFOBUp100Khz()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBR()`** — L1590 — `public string ZZBR(string s)`
  Sets or reads the BCI Rejection button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZBS()`** — L1620 — `public string ZZBS(string s)`
  Sets or reads the current band setting
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Band160m()` (`Console/Midi2CatCommands.cs`), `.Band80m()` (`Console/Midi2CatCommands.cs`), `.Band60m()` (`Console/Midi2CatCommands.cs`), `.Band40m()` (`Console/Midi2CatCommands.cs`) — and 9 more
- **`.ZZBT()`** — L1626 — `public string ZZBT(string s)`
  Sets or gets the current RX2 band setting
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Band160mRX2()` (`Console/Midi2CatCommands.cs`), `.Band80mRX2()` (`Console/Midi2CatCommands.cs`), `.Band60mRX2()` (`Console/Midi2CatCommands.cs`), `.Band40mRX2()` (`Console/Midi2CatCommands.cs`) — and 9 more
- **`.ZZBU()`** — L1646 — `public string ZZBU()`
  Moves the bandswitch up one band
  Called by: `.BU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.BandUp()` (`Console/Midi2CatCommands.cs`)
- **`.ZZBY()`** — L1653 — `public string ZZBY()`
  Shuts down the console
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CloseConsole()` (`Console/Midi2CatCommands.cs`)
- **`.ZZCB()`** — L1668 — `public string ZZCB(string s)`
  Sets or reads the CW Break In Enabled checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CWBreakIn()` (`Console/Midi2CatCommands.cs`)
- **`.ZZCD()`** — L1688 — `public string ZZCD(string s)`
  Sets or reads the CW Break In Delay
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZCF()`** — L1714 — `public string ZZCF(string s)`
  Sets or reads the Show CW Frequency checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZCI()`** — L1745 — `public string ZZCI(string s)`
  Sets or reads the CW Iambic checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZCL()`** — L1770 — `public string ZZCL(string s)`
  Sets or reads the CW Pitch thumbwheel
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZCM()`** — L1793 — `public string ZZCM(string s)`
  Sets or reads the CW Monitor Disable button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZCN()`** — L1818 — `public string ZZCN(string s)`
  Sets or reads CTUN for RX1
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CTunOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZCO()`** — L1847 — `public string ZZCO(string s)`
  Sets or reads CTUN for RX2
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2CTunOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZCP()`** — L1875 — `public string ZZCP(string s)`
  Sets or reads the compander button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CompanderOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZCS()`** — L1896 — `public string ZZCS(string s)`
  Sets or reads the CW Speed thumbwheel
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CWSpeed()` (`Console/Midi2CatCommands.cs`), `.CWSpeed_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZCT()`** — L1920 — `public string ZZCT(string s)`
  Reads or sets the compander threshold
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CPDRLevel()` (`Console/Midi2CatCommands.cs`)
- **`.ZZCU()`** — L1946 — `public string ZZCU()`
  Reads the CPU Usage
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZDA()`** — L1954 — `public string ZZDA(string s)`
  Sets or reads the Display Average status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DisplayAverage()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDB()`** — L1974 — `public string ZZDB(string s)`
  Sets or reads the Diversity Form RX Reference radio buttons
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DiversityReference()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDC()`** — L2004 — `public string ZZDC(string s)`
  Sets or reads the Diversity Form RX2 gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZDD()`** — L2031 — `public string ZZDD(string s)`
  Sets or reads the Diversity Form phase
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DiversityPhase()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDE()`** — L2065 — `public string ZZDE(string s)`
  Sets or reads the Diversity Form Enable Button
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DiversityEnable()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDF()`** — L2095 — `public string ZZDF(string s)`
  Opens or closes the Diversity Form
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DiversityFormOpen()` (`Console/Midi2CatCommands.cs`), `.ESCFormOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDG()`** — L2124 — `public string ZZDG(string s)`
  Sets or reads the Diversity Form RX gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DiversityGain()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDH()`** — L2150 — `public string ZZDH(string s)`
  Sets or reads the Diversity Form RX Source radio buttons
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DiversitySource()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDM()`** — L2187 — `public string ZZDM(string s)`
  Sets or reads the current display mode
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DisplayModePrev()` (`Console/Midi2CatCommands.cs`), `.DisplayModeNext()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDN()`** — L2259 — `public string ZZDN(string s)`
  Reads or sets the setup form Waterfall Low value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.WaterfallLowLimit()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDO()`** — L2290 — `public string ZZDO(string s)`
  Reads or sets the setup form Waterfall High value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.WaterfallHighLimit()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDP()`** — L2322 — `public string ZZDP(string s)`
  Reads or sets the setup form Spectrum Grid Max Value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.WaterfallHighLimit()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDQ()`** — L2353 — `public string ZZDQ(string s)`
  Reads or sets the setup form Spectrum Grid Min Value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.WaterfallLowLimit()` (`Console/Midi2CatCommands.cs`)
- **`.ZZDR()`** — L2384 — `public string ZZDR(string s)`
  Sets or reads the Spectrum Grid Step
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZDU()`** — L2404 — `public string ZZDU()`
  Constructs the state word for DDUtil read only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZDX()`** — L2471 — `public string ZZDX(string s)`
  Sets or reads the DX button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.StereoDiversityOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZEA()`** — L2525 — `public string ZZEA(string s)`
  Reads or sets the RX equalizer. The CAT suffix string is 36 characters constant. Each value in the string occupies exactly three characters starting with the number of bands (003 or 010) followed by the preamp setting (-12 to 015) followed by 3 or 10 three digit EQ thumbwheel positions. If the…
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZEB()`** — L2566 — `public string ZZEB(string s)`
  Sets or reads the TX EQ settings
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZEM()`** — L2607 — `public string ZZEM(string s)`
  Provides verbose CAT error reporting
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZER()`** — L2630 — `public string ZZER(string s)`
  Sets or reads the RXEQ button statusl
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RXEQOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZET()`** — L2649 — `public string ZZET(string s)`
  Sets or reads the TXEQ button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TXEQOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZFA()`** — L2670 — `public string ZZFA(string s)`
  Sets or reads VFO A frequency
  Called by: `.FA()` (same file), `.IF()` (same file), `.ZZDU()` (same file), `.ZZIF()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ProcessStdMIDIWheelAsVFO()` (`Console/Midi2CatCommands.cs`) — and 3 more
- **`.ZZFB()`** — L2719 — `public string ZZFB(string s)`
  Sets or reads VFO B frequency
  Called by: `.FB()` (same file), `.IF()` (same file), `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ProcessStdMIDIWheelAsVFO()` (`Console/Midi2CatCommands.cs`), `.ProcessBehringerMainWheelAsVFO()` (`Console/Midi2CatCommands.cs`) — and 2 more
- **`.ZZFT()`** — L2770 — `public string ZZFT(string s)`
  Sets or reads TX frequency
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZFD()`** — L2788 — `public string ZZFD(string s)`
  Selects or reads the FM deviation radio button
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZFI()`** — L2813 — `public string ZZFI(string s)`
  Sets or reads the current filter index number
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx1FilterWider()` (`Console/Midi2CatCommands.cs`), `.Rx1FilterNarrower()` (`Console/Midi2CatCommands.cs`)
- **`.ZZFJ()`** — L2840 — `public string ZZFJ(string s)`
  Sets or reads the current RX2 DSP filter
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2FilterWider()` (`Console/Midi2CatCommands.cs`), `.Rx2FilterNarrower()` (`Console/Midi2CatCommands.cs`)
- **`.ZZFL()`** — L2867 — `public string ZZFL(string s)`
  Reads or sets the DSP Filter Low value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.FilterLow()` (`Console/Midi2CatCommands.cs`)
- **`.ZZFH()`** — L2904 — `public string ZZFH(string s)`
  Reads or sets the DSP Filter High value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.FilterHigh()` (`Console/Midi2CatCommands.cs`)
- **`.ZZFM()`** — L2935 — `public string ZZFM()`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.PreAmpSettingsKnob()` (`Console/Midi2CatCommands.cs`), `.PreampFlex5000()` (`Console/Midi2CatCommands.cs`)
- **`.ZZFR()`** — L2961 — `public string ZZFR(string s)`
  Reads or sets the RX2 DSP Filter High value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZFS()`** — L2997 — `public string ZZFS(string s)`
  Reads or sets the RX2 DSP Filter Low value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZFV()`** — L3028 — `public string ZZFV(string s)`
  Reads FlexWire single byte value commands
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZFW()`** — L3048 — `public string ZZFW(String s)`
  Reds FlexWire double byte value commands
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZFX()`** — L3071 — `public string ZZFX(string s)`
  Sends FlexWire single byte value commands
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZFY()`** — L3090 — `public string ZZFY(String s)`
  Sends FlexWire double byte value commands
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZGA()`** — L3118 — `public string ZZGA(string s)`
  Adds an id mainly used by tcpip cat ot direct messages back to a specific cat client, there is no get TCPIPcatserver will use the response to add an id against the client
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZGR()`** — L3138 — `public string ZZGR(string s)`
  Remove an id mainly used by tcpip cat ot direct messages back to a specific cat client, there is no get TCPIPcatserver will use the response to remove an id from the client
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZGE()`** — L3157 — `public string ZZGE(string s)`
  Sets or reads the noise gate enable button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DEXPOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZGL()`** — L3182 — `public string ZZGL(string s)`
  Sets or reads the noise gate level control
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DEXPThreshold()` (`Console/Midi2CatCommands.cs`)
- **`.ZZGT()`** — L3217 — `public string ZZGT(string s)`
  Sets or reads the AGC constant
  Called by: `.GT()` (same file), `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.AGCModeKnob()` (`Console/Midi2CatCommands.cs`), `.AGCModeUp()` (`Console/Midi2CatCommands.cs`), `.AGCModeDown()` (`Console/Midi2CatCommands.cs`)
- **`.ZZGU()`** — L3239 — `public string ZZGU(string s)`
  Sets or reads the RX2 AGC constant
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2AGCModeKnob()` (`Console/Midi2CatCommands.cs`), `.RX2AGCModeUp()` (`Console/Midi2CatCommands.cs`), `.RX2AGCModeDown()` (`Console/Midi2CatCommands.cs`)
- **`.ZZHA()`** — L3260 — `public string ZZHA(string s)`
  Sets or reads the Audio Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZHR()`** — L3277 — `public string ZZHR(string s)`
  Sets or reads the DSP Phone RX Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZHT()`** — L3295 — `public string ZZHT(string s)`
  Sets or reads the DSP Phone TX Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZHU()`** — L3313 — `public string ZZHU(string s)`
  Sets or reads the DSP CW RX Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZHV()`** — L3331 — `public string ZZHV(string s)`
  Sets or reads the DSP CW TX Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZHW()`** — L3349 — `public string ZZHW(string s)`
  Sets or reads the DSP Digital RX Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZHX()`** — L3367 — `public string ZZHX(string s)`
  Sets or reads the DSP Digital TX Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZID()`** — L3386 — `public string ZZID()`
  Sets the CAT Rig Type to SDR-1000 Modified 10/12/08 BT changed "SDR-1000" to "PowerSDR"
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZIF()`** — L3403 — `public string ZZIF(string s)`
  Reads the SDR-1000 transceiver status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZIO()`** — L3461 — `public string ZZIO()`
  Reads the installed options
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZIS()`** — L3467 — `public string ZZIS(string s)`
  Sets or reads the IF width
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.FilterBandwidth()` (`Console/Midi2CatCommands.cs`)
- **`.ZZIT()`** — L3490 — `public string ZZIT(string s)`
  Sets or reads the IF Shift
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.FilterShift()` (`Console/Midi2CatCommands.cs`)
- **`.ZZIU()`** — L3520 — `public string ZZIU()`
  Resets the Filter Shift to zero. Write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZKO()`** — L3526 — `public string ZZKO(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZKM()`** — L3556 — `public string ZZKM(string s)`
  Sends a CWX macro
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CWXMacro1()` (`Console/Midi2CatCommands.cs`), `.CWXMacro2()` (`Console/Midi2CatCommands.cs`), `.CWXMacro3()` (`Console/Midi2CatCommands.cs`), `.CWXMacro4()` (`Console/Midi2CatCommands.cs`), `.CWXMacro5()` (`Console/Midi2CatCommands.cs`) — and 4 more
- **`.ZZKS()`** — L3576 — `public string ZZKS(string s)`
  Sets or reads the CWX CW speed
  Called by: `.KS()` (same file), `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZKY()`** — L3599 — `public string ZZKY(string s)`
  Sends text to CWX for conversion to Morse
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZLA()`** — L3665 — `public string ZZLA(string s)`
  Sets or reads the RX0Gain level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VolumeVfoA()` (`Console/Midi2CatCommands.cs`), `.VolumeVfoA_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZLB()`** — L3691 — `public string ZZLB(string s)`
  Sets or reads the RX0 (Main RX) stereo balance
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RatioMainSubRx()` (`Console/Midi2CatCommands.cs`)
- **`.ZZLC()`** — L3722 — `public string ZZLC(string s)`
  Sets or reads the RX1 (SubRX) Gain level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VolumeVfoB()` (`Console/Midi2CatCommands.cs`)
- **`.ZZLD()`** — L3748 — `public string ZZLD(string s)`
  Sets or reads the RX1 (Sub RX) stereo balance
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RatioMainSubRx()` (`Console/Midi2CatCommands.cs`)
- **`.ZZLE()`** — L3779 — `public string ZZLE(string s)`
  Sets or reads the RX2 Gain level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VolumeVfoB_inc()` (`Console/Midi2CatCommands.cs`), `.RX2Volume()` (`Console/Midi2CatCommands.cs`)
- **`.ZZLF()`** — L3805 — `public string ZZLF(string s)`
  Sets or reads the RX2 stereo balance
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2Pan()` (`Console/Midi2CatCommands.cs`)
- **`.ZZLG()`** — L3836 — `public string ZZLG(string s)`
  Sets or reads the AutoMute RX1 on VFOB TX checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZLH()`** — L3859 — `public string ZZLH(string s)`
  Sets or reads the AutoMute RX2 on VFOA TX checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZLI()`** — L3882 — `public string ZZLI(string s)`
  Sets or reads the PS-A button on/off status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.PSOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZMA()`** — L3909 — `public string ZZMA(string s)`
  Sets or reads the MUT button on/off status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MuteOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZMB()`** — L3936 — `public string ZZMB(string s)`
  Sets or reads the RX2 MUT button on/off status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MuteRX2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZMD()`** — L3966 — `public string ZZMD(string s)`
  Sets or reads the SDR-1000 DSP mode
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx1ModeNext()` (`Console/Midi2CatCommands.cs`), `.Rx1ModePrev()` (`Console/Midi2CatCommands.cs`), `.XIT_inc()` (`Console/Midi2CatCommands.cs`), `.ChangeFreqVfoA()` (`Console/Midi2CatCommands.cs`) — and 19 more
- **`.ZZME()`** — L3988 — `public string ZZME(string s)`
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2ModeNext()` (`Console/Midi2CatCommands.cs`), `.Rx2ModePrev()` (`Console/Midi2CatCommands.cs`), `.RX2ModeSSB()` (`Console/Midi2CatCommands.cs`), `.RX2ModeLSB()` (`Console/Midi2CatCommands.cs`) — and 12 more
- **`.ZZMF()`** — L4040 — `public string ZZMF(string s)`
  ZZMFcccccccccccccccccccc; Set multifunction encoder text cc are 15 pairs of digits 0-99 each making up an ASCII code -32 (so 'A' is 33 for example)
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZD()`** — L4063 — `public string ZZZD(string s)`
  Andromeda front panel VFO encoder down write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZU()`** — L4077 — `public string ZZZU(string s)`
  Andromeda front panel VFO encoder up write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZS()`** — L4091 — `public string ZZZS(string s)`
  Andromeda front panel s/w version write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZE()`** — L4105 — `public string ZZZE(string s)`
  Andromeda front panel encoder step write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZP()`** — L4124 — `public string ZZZP(string s)`
  Andromeda front panel pushbutton press write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZA()`** — L4145 — `public string ZZZA(string s)`
  CATHandleAriesTuneMessage Ganymeda amplifier trip state write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOX()`** — L4160 — `public string ZZOX(string s)`
  ARIES ATU tune state message write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOZ()`** — L4178 — `public string ZZOZ(string s)`
  ARIES ATU erase state message write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMG()`** — L4194 — `public string ZZMG(string s)`
  Sets or reads the Mic gain control
  Called by: `.MG()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MicGain()` (`Console/Midi2CatCommands.cs`)
- **`.ZZML()`** — L4226 — `public string ZZML()`
  Returns a list of modes with index
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMN()`** — L4259 — `public string ZZMN(string s)`
  Reads the DSP filter presets for filter index (s) Returns 180 character length word for 12 filters x 15 characters each. Format is name high low: ZZMN 5.0k 5150 -160...
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMO()`** — L4270 — `public string ZZMO(string s)`
  Sets or reads the Monitor (MON) button status
  Called by: `.MO()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MONOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZMR()`** — L4293 — `public string ZZMR(string s)`
  Sets or reads the RX meter mode
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMS()`** — L4317 — `public string ZZMS(string s)`
  Sets or reads the MultiRX Swap checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMT()`** — L4336 — `public string ZZMT(string s)`
  Sets or reads the TX meter mode
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMU()`** — L4359 — `public string ZZMU(string s)`
  Sets or reads the MultiRX button status
  Called by: `.ZZDU()` (same file), `.ZZTV()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MultiRxOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZMV()`** — L4379 — `public string ZZMV()`
  Returns the count of memory records read only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMW()`** — L4393 — `public string ZZMW(string s)`
  Deletes a memory channel
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMX()`** — L4411 — `public string ZZMX(string s)`
  Restores memory channel n
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMY()`** — L4435 — `public string ZZMY()`
  Saves the current radio configuration to a new memory channel
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZMZ()`** — L4465 — `public string ZZMZ(string s)`
  Saves the radio configuration to a specific channel number (edit)
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZJS()`** — L4512 — `public string ZZJS()`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZJR()`** — L4523 — `public string ZZJR(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZJP()`** — L4631 — `public string ZZJP(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZJQ()`** — L4731 — `public string ZZJQ(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZNA()`** — L4803 — `public string ZZNA(string s)`
  Sets or reads Noise Blanker 2 status
  Called by: `.NB()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx1NoiseBlanker1OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNB()`** — L4821 — `public string ZZNB(string s)`
  Sets or reads the Noise Blanker 2 status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx1Noiseblanker2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNC()`** — L4840 — `public string ZZNC(string s)`
  Sets or reads the RX2 NB status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2NoiseBlanker1OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZND()`** — L4858 — `public string ZZND(string s)`
  Sets or reads the RX2 NB2 status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2Noiseblanker2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNL()`** — L4879 — `public string ZZNL(string s)`
  Sets or reads the Noise Blanker 1 threshold
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZNM()`** — L4898 — `public string ZZNM(string s)`
  Sets or reads the Noise Blanker 2 threshold
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZNN()`** — L4917 — `public string ZZNN(string s)`
  Sets or reads the Rx1 Spectral Noise Blanker
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.SpectralNoiseBlankerOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNO()`** — L4935 — `public string ZZNO(string s)`
  Sets or reads the Rx2 Spectral Noise Blanker
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.SpectralNoiseBlankerRx2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNR()`** — L4953 — `public string ZZNR(string s)`
  Sets or reads the RX1 Noise Reduction status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.NoiseReductionOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNS()`** — L4976 — `public string ZZNS(string s)`
  Sets or reads the RX1 NR2 button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.NoiseReduction2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNE()`** — L5000 — `public string ZZNE(string s)`
  Sets or reads the RX1 Noise Reduction status returns 0 for off, 1,2,3,4 depending on NR in use
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.NoiseReductionOnOff()` (`Console/Midi2CatCommands.cs`), `.NoiseReduction2OnOff()` (`Console/Midi2CatCommands.cs`), `.NoiseReduction3OnOff()` (`Console/Midi2CatCommands.cs`), `.NoiseReduction4OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNF()`** — L5024 — `public string ZZNF(string s)`
  Sets or reads the RX2 Noise Reduction status returns 0 for off, 1,2,3,4 depending on NR in use
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2NoiseReductionOnOff()` (`Console/Midi2CatCommands.cs`), `.Rx2NoiseReduction2OnOff()` (`Console/Midi2CatCommands.cs`), `.Rx2NoiseReduction3OnOff()` (`Console/Midi2CatCommands.cs`), `.Rx2NoiseReduction4OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNG()`** — L5048 — `public string ZZNG(string s)`
  Sets the RX1 NR4 reduction amount, supplied as int from 0 to 100, then converted to 0-20dB
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.NoiseReduction4Amount()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNH()`** — L5110 — `public string ZZNH(string s)`
  Sets the RX1 NR4 reduction amount, supplied as int from 0 to 100, then converted to 0-20dB
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2NoiseReduction4Amount()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNT()`** — L5173 — `public string ZZNT(string s)`
  Sets or reads the ANF button status
  Called by: `.NT()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.AutoNotchOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNU()`** — L5191 — `public string ZZNU(string s)`
  Sets or reads the RX2 ANF button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2AutoNotchOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNV()`** — L5208 — `public string ZZNV(string s)`
  Sets or reads the Noise Reduction status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2NoiseReductionOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZNW()`** — L5231 — `public string ZZNW(string s)`
  Sets or reads the RX2 NR2 button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2NoiseReduction2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZOA()`** — L5254 — `public string ZZOA(string s)`
  Sets or reads the RX1 antenna //[2.3.10.6]MW0LGE https://github.com/ramdor/Thetis/issues/385
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOB()`** — L5283 — `public string ZZOB(string s)`
  Sets or reads the RX2 antenna (if RX2 installed)
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOC()`** — L5290 — `public string ZZOC(string s)`
  Sets or reads the TX antenna //[2.3.10.6]MW0LGE https://github.com/ramdor/Thetis/issues/385
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOD()`** — L5319 — `public string ZZOD(string s)`
  Sets or reads the current Antenna Mode
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOE()`** — L5326 — `public string ZZOE(string s)`
  Sets or reads the RX1 External Antenna checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOF()`** — L5333 — `public string ZZOF(string s)`
  Sets or reads the TX relay RCA jack
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOG()`** — L5341 — `public string ZZOG(string s)`
  Sets or reads the TX Relay Delay enables
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOH()`** — L5348 — `public string ZZOH(string s)`
  Sets or reads the TX Relay Delays
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOJ()`** — L5354 — `public string ZZOJ(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOL()`** — L5361 — `public string ZZOL(string s)`
  Sets or reads the DigL Click Tune Offset
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOS()`** — L5386 — `public string ZZOS(string s)`
  Sets or reads the current repeater offset direction
  Called by: `.OS()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOT()`** — L5401 — `public string ZZOT(string s)`
  Sets for reads the repeater frequency offset need to resolve the negative offset question.
  Called by: `.OF()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOU()`** — L5420 — `public string ZZOU(string s)`
  Sets or reads the DigU Click Tune Offset
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZOV()`** — L5446 — `public string ZZOV(string s)`
  Sets or reads the console ATU button
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TunerOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZOW()`** — L5453 — `public string ZZOW(string s)`
  Sets or reads the console ATU Bypass button
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TunerBypassOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPA()`** — L5460 — `public string ZZPA(string s)`
  Sets or reads the Preamp thumbwheel
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.PreAmpSettingsKnob()` (`Console/Midi2CatCommands.cs`), `.PreampFlex5000()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPB()`** — L5515 — `public string ZZPB(string s)`
  Sets or reads the RX2 Preamp button
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.Rx2PreAmpOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPC()`** — L5587 — `public string ZZPC(string s)`
  Sets or reads the Drive level
  Called by: `.PC()` (same file), `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DriveLevel()` (`Console/Midi2CatCommands.cs`), `.DriveLevel_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPD()`** — L5612 — `public string ZZPD()`
  Centers the Display Pan scroll
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.PanCenter()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPE()`** — L5619 — `public string ZZPE(string s)`
  Sets or reads the Display Pan control
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.PanSliderInc()` (`Console/Midi2CatCommands.cs`), `.PanSlider()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPO()`** — L5707 — `public string ZZPO(string s)`
  Sets or reads the Display Peak button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DisplayPeak()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPS()`** — L5723 — `public string ZZPS(string s)`
  Sets or reads the Power button status
  Called by: `.PS()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.StartOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPY()`** — L5748 — `public string ZZPY(string s)`
  Sets or reads the Display Zoom control
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ZoomDec()` (`Console/Midi2CatCommands.cs`), `.ZoomInc()` (`Console/Midi2CatCommands.cs`), `.ZoomSliderInc()` (`Console/Midi2CatCommands.cs`), `.ZoomSliderFix()` (`Console/Midi2CatCommands.cs`)
- **`.ZZPZ()`** — L5774 — `public string ZZPZ(string s)`
  Sets the Display Zoom buttons
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ZoomDec()` (`Console/Midi2CatCommands.cs`), `.ZoomInc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZQK()`** — L5793 — `public string ZZQK(string s)`
  Sets or reads the CW Break-In for Semi/QSK modes
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CWQSK()` (`Console/Midi2CatCommands.cs`)
- **`.ZZQM()`** — L5812 — `public string ZZQM()`
  Reads the Quick Memory Save value
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZQR()`** — L5818 — `public string ZZQR()`
  Recalls Memory Quick Save
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.QuickModeRestore()` (`Console/Midi2CatCommands.cs`)
- **`.ZZQS()`** — L5825 — `public string ZZQS()`
  Saves Quick Memory value
  Called by: `.QI()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.QuickModeSave()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRA()`** — L5837 — `public string ZZRA(string s)`
  Sets or reads the RTTY Offset Enable VFO A checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ProcessStdMIDIWheelAsVFO()` (`Console/Midi2CatCommands.cs`), `.ProcessBehringerMainWheelAsVFO()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRB()`** — L5862 — `public string ZZRB(string s)`
  Sets or reads the RTTY Offset Enable VFO B checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ProcessStdMIDIWheelAsVFO()` (`Console/Midi2CatCommands.cs`), `.ProcessBehringerMainWheelAsVFO()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRC()`** — L5887 — `public string ZZRC()`
  Clears the RIT frequency
  Called by: `.RC()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RIT_inc()` (`Console/Midi2CatCommands.cs`), `.RIT_clear()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRD()`** — L5894 — `public string ZZRD(string s)`
  Decrements RIT
  Called by: `.RD()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RIT_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRF()`** — L5921 — `public string ZZRF(string s)`
  Sets or reads the RIT frequency value
  Called by: `.ZZDU()` (same file), `.ZZRD()` (same file), `.ZZRU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RIT()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRH()`** — L5957 — `public string ZZRH(string s)`
  Sets or reads the RTTY DIGH offset frequency ud counter
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ProcessStdMIDIWheelAsVFO()` (`Console/Midi2CatCommands.cs`), `.ProcessBehringerMainWheelAsVFO()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRL()`** — L5993 — `public string ZZRL(string s)`
  Sets or reads the RTTY DIGL offset frequency ud counter
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ProcessStdMIDIWheelAsVFO()` (`Console/Midi2CatCommands.cs`), `.ProcessBehringerMainWheelAsVFO()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRM()`** — L6029 — `public string ZZRM(string s)`
  Reads the Console RX meter
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZRS()`** — L6074 — `public string ZZRS(string s)`
  Sets or reads the RX2 button status
  Called by: `.ZZDU()` (same file), `.ZZTV()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRT()`** — L6099 — `public string ZZRT(string s)`
  Sets or reads the RIT button status
  Called by: `.RT()` (same file), `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RitOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRU()`** — L6124 — `public string ZZRU(string s)`
  Increments RIT
  Called by: `.RU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RIT_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZRV()`** — L6151 — `public string ZZRV()`
  Reads the primary input voltage
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZRX()`** — L6176 — `public string ZZRX(string s)`
  Sets or reads the RX1 step attenuation control, 0 to 31dB
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZRY()`** — L6200 — `public string ZZRY(string s)`
  Sets or reads the RX2 step attenuation control, 0 to 31dB
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSA()`** — L6224 — `public string ZZSA()`
  Moves VFO A down one Tune Step
  Called by: `.DN()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSB()`** — L6240 — `public string ZZSB()`
  Moves VFO A up one Tune Step
  Called by: `.UP()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSD()`** — L6257 — `public string ZZSD()`
  Moves the mouse wheel tuning step down
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TuningStepDown()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSF()`** — L6264 — `public string ZZSF(string s)`
  ZZSFccccwwww Set Filter, cccc=center freq www=width both in hz
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSG()`** — L6274 — `public string ZZSG()`
  Moves VFO B down one Tune Step
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSH()`** — L6291 — `public string ZZSH()`
  Moves VFO B up one Tune Step
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSM()`** — L6307 — `public string ZZSM(string s)`
  Reads the S Meter value
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSN()`** — L6360 — `public string ZZSN()`
  Reads the radio serial number
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSP()`** — L6370 — `public string ZZSP(string s, bool bFromCatDirect = false)`
  Sets or reads the VFO Split status
  Called by: `.FT()` (same file), `.ZZDU()` (same file), `.ZZTV()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.SplitOnOff()` (`Console/Midi2CatCommands.cs`), `.QuickSplitOnOffandSplitOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSO()`** — L6405 — `public string ZZSO(string s)`
  Sets or reads the Squelch on/off status
  Called by: `.ZZSQ()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.SquelchOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSQ()`** — L6421 — `public string ZZSQ(string s)`
  Sets or reads the SDR-1000 Squelch control
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.SquelchControl()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSR()`** — L6464 — `public string ZZSR(string s)`
  Reads or sets the Spur Reduction button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.SpurReductionOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSS()`** — L6487 — `public string ZZSS()`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.CWXStop()` (`Console/Midi2CatCommands.cs`)
- **`.ZZST()`** — L6494 — `public string ZZST()`
  Reads the current console step size (read-only property)
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZSU()`** — L6501 — `public string ZZSU()`
  Moves the mouse wheel step tune up
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TuningStepUp()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSV()`** — L6508 — `public string ZZSV(string s)`
  Sets or reads the RX2 Squelch button
  Called by: `.ZZSX()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2SquelchOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSW()`** — L6524 — `public string ZZSW(string s)`
  Swaps VFO A/B TX buttons
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ToggleTX()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSX()`** — L6551 — `public string ZZSX(string s)`
  Sets or reads the RX2 Squelch threshold
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2SquelchControl()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSY()`** — L6596 — `public string ZZSY(string s)`
  Reads or sets the VFO Sync button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VfoSyncOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZSZ()`** — L6619 — `public string ZZSZ(string s)`
  Syncs the chosen vfo to the selected tune step write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZTA()`** — L6637 — `public string ZZTA(string s)`
  Sets or reads the CTCSS enable button
  Called by: `.CT()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZTB()`** — L6665 — `public string ZZTB(string s)`
  Sets or reads the CTCSS tone frequency
  Called by: `.CN()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZTF()`** — L6691 — `public string ZZTF(string s)`
  Sets or reads the Show TX Filter checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.DisplayTxFilter()` (`Console/Midi2CatCommands.cs`)
- **`.ZZTH()`** — L6721 — `public string ZZTH(string s)`
  Sets or reads the TX filter high setting
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TXFilterHigh()` (`Console/Midi2CatCommands.cs`)
- **`.ZZTI()`** — L6744 — `public string ZZTI(string s)`
  Inhibits power output when using external antennas, tuners, etc.
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZTL()`** — L6765 — `public string ZZTL(string s)`
  Sets or reads the TX filter low setting
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TXFilterLow()` (`Console/Midi2CatCommands.cs`)
- **`.ZZTM()`** — L6788 — `public string ZZTM(string s)`
  Sets or reads the TX Monitor level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TXAFMonitor()` (`Console/Midi2CatCommands.cs`)
- **`.ZZTO()`** — L6812 — `public string ZZTO(string s)`
  Sets or reads the Tune Power level
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TUNPowerLevel()` (`Console/Midi2CatCommands.cs`)
- **`.ZZTP()`** — L6877 — `public string ZZTP(string s)`
  Sets or reads the TX Profile
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZTS()`** — L6898 — `public string ZZTS()`
  Reads the Flex 5000 temperature sensor
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZTU()`** — L6926 — `public string ZZTU(string s)`
  Sets or reads the TUN button on/off status
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TunOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZTV()`** — L6953 — `public string ZZTV(string s)`
  Sets or reads the transmit VFO frequency when in split with RX2 enabled
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZTX()`** — L6979 — `public string ZZTX(string s)`
  Sets or reads the MOX button status
  Called by: `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.MOXOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZUA()`** — L7002 — `public string ZZUA()`
  Reads the XVTR Band Names
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVA()`** — L7009 — `public string ZZVA(string s)`
  Reads or sets the VAC Enable checkbox (Setup Form)
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VACOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZUP()`** — L7032 — `public string ZZUP(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ExternalPAOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZUS()`** — L7061 — `public string ZZUS()`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZUT()`** — L7067 — `public string ZZUT(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.TwoToneOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVB()`** — L7103 — `public string ZZVB(string s)`
  Sets or reads the VAC RX Gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VACGainRX()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVC()`** — L7142 — `public string ZZVC(string s)`
  Sets or reads the VAC TX Gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VACGainTX()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVD()`** — L7181 — `public string ZZVD(string s)`
  Sets or reads the VAC Sample Rate
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVE()`** — L7267 — `public string ZZVE(string s)`
  Reads or sets the VOX Enable button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VOXOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZQA()`** — L7290 — `public string ZZQA(string s)`
  Reads or sets the Quick Play button status // DH1KLM
  Called by: `.QuickPlayOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZQB()`** — L7313 — `public string ZZQB(string s)`
  Reads or sets the Quick Rec button status // DH1KLM
  Called by: `.QuickRecOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVF()`** — L7340 — `public string ZZVF(string s)`
  Sets or reads the VAC Stereo checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVG()`** — L7365 — `public string ZZVG(string s)`
  Reads or set the VOX Gain control
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VOXGain()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVH()`** — L7399 — `public string ZZVH(string s)`
  Reads or sets the I/Q to VAC checkbox on the setup form
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.IQtoVAC()` (`Console/Midi2CatCommands.cs`), `.IQtoVACRX2()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVI()`** — L7424 — `public string ZZVI(string s)`
  Reads or sets the VAC Input cable
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVJ()`** — L7441 — `public string ZZVJ(string s)`
  Reads or sets the Direct I/Q Use RX2 checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.IQtoVACRX2()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVK()`** — L7465 — `public string ZZVK(string s)`
  Reads or sets the VAC2 Enable checkbox (Setup Form)
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VAC2OnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVL()`** — L7491 — `public string ZZVL(string s)`
  Reads or sets the VFO Lock button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.LockVFOOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZUX()`** — L7532 — `public string ZZUX(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.LockVFOAOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZUY()`** — L7563 — `public string ZZUY(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.LockVFOBOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVM()`** — L7595 — `public string ZZVM(string s)`
  Reads or sets the VAC Driver
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVN()`** — L7613 — `public string ZZVN()`
  Returns the version number of the PowerSDR program
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVO()`** — L7620 — `public string ZZVO(string s)`
  Reads or sets the VAC Output cable
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVP()`** — L7637 — `public string ZZVP(string s)`
  Reads or sets the VAC1 IQ Calibrate checkbox on the setup form
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVQ()`** — L7662 — `public string ZZVQ(string s)`
  Reads or sets the VAC2 Driver
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVR()`** — L7679 — `public string ZZVR(string s)`
  Reads or sets the VAC2 Input cable
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVS()`** — L7697 — `public string ZZVS(string s)`
  Sets the VFO swap status write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVT()`** — L7712 — `public string ZZVT(string s)`
  Reads or sets the VAC2 Output cable
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVU()`** — L7733 — `public string ZZVU(string s)`
  Sets or reads the VAC2 Sample Rate
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVV()`** — L7828 — `public string ZZVV(string s)`
  Sets or reads the VAC2 Stereo checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVW()`** — L7852 — `public string ZZVW(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VAC2GainRX()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVX()`** — L7891 — `public string ZZVX(string s)`
  Sets or reads the VAC2 TX Gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.VAC2GainTX()` (`Console/Midi2CatCommands.cs`)
- **`.ZZVY()`** — L7930 — `public string ZZVY(string s)`
  Sets or reads the VAC1 Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZVZ()`** — L7984 — `public string ZZVZ(string s)`
  Sets or reads the VAC1 Buffer Size
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWA()`** — L8034 — `public string ZZWA(string s)`
  Sets or reads the F5K Mixer Mic Gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWB()`** — L8041 — `public string ZZWB(string s)`
  Sets or reads the F5K Line In RCA level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWC()`** — L8048 — `public string ZZWC(string s)`
  Sets or reads the F5K Line In Phono level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWD()`** — L8055 — `public string ZZWD(string s)`
  Sets or reads the F5K Mixer Line In DB9 level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWE()`** — L8063 — `public string ZZWE(string s)`
  Sets or reads the F1500F5K Mixer Mic Selected Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWF()`** — L8070 — `public string ZZWF(string s)`
  Sets or reads the F5K Mixer Line In RCA Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWG()`** — L8077 — `public string ZZWG(string s)`
  Sets or reads the F5K Mixer Line In Phono Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWH()`** — L8084 — `public string ZZWH(string s)`
  Sets or reads the F1500/F5K Mixer Line In FlexWire/DB9 Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWJ()`** — L8092 — `public string ZZWJ(string s)`
  Sets or reads the F5K Mixer Mute All Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWK()`** — L8099 — `public string ZZWK(string s)`
  Sets or reads the F5K Mixer Internal Speaker level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWL()`** — L8106 — `public string ZZWL(string s)`
  Sets or reads the F5K Mixer External Speaker level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWM()`** — L8113 — `public string ZZWM(string s)`
  Sets or reads the F5K Mixer Headphone level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWN()`** — L8120 — `public string ZZWN(string s)`
  Sets or reads the F5K Mixer Line Out RCA level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWO()`** — L8127 — `public string ZZWO(string s)`
  Sets or reads the F5KC Mixer Internal Speaker Selected Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWP()`** — L8134 — `public string ZZWP(string s)`
  Sets or reads the F5K Mixer External Speaker Selected Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWQ()`** — L8141 — `public string ZZWQ(string s)`
  Sets or reads the F1500F5K Mixer Headphone Selected Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWR()`** — L8148 — `public string ZZWR(string s)`
  Sets or reads the F1500 FlexWire Out/F5K Mixer Line Out RCA Selected Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWS()`** — L8155 — `public string ZZWS(string s)`
  Sets or reads the F1500/F5K Mixer Output Mute All Checkbox
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWT()`** — L8163 — `public string ZZWT(string s)`
  Reads or sets the F1500 mixer form mic level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWU()`** — L8170 — `public string ZZWU(string s)`
  Reads or sets the F1500 Mixer Form FireWire Input Level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWV()`** — L8177 — `public string ZZWV(string s)`
  Sets ir reads the F1500 Mixer Form Phones level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZWW()`** — L8184 — `public string ZZWW(string s)`
  Sets or reads the F1500 Mixer Form FlexWire Out level
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXC()`** — L8192 — `public string ZZXC()`
  Clears the XIT frequency write only
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.XIT_inc()` (`Console/Midi2CatCommands.cs`), `.XIT_clear()` (`Console/Midi2CatCommands.cs`)
- **`.ZZXD()`** — L8199 — `public string ZZXD(string s)`
  Decrements XIT
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXF()`** — L8226 — `public string ZZXF(string s)`
  Sets or reads the XIT frequency value
  Called by: `.ZZDU()` (same file), `.ZZXD()` (same file), `.ZZXU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.XIT()` (`Console/Midi2CatCommands.cs`), `.XIT_inc()` (`Console/Midi2CatCommands.cs`)
- **`.ZZXH()`** — L8261 — `public string ZZXH(string s)`
  Reads or set the VOX Delay control
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXN()`** — L8290 — `public string ZZXN(string s)`
  Reads RX1 combined status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXO()`** — L8325 — `public string ZZXO(string s)`
  Reads RX2 combined status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXS()`** — L8360 — `public string ZZXS(string s)`
  Sets or reads the XIT button status
  Called by: `.XT()` (same file), `.ZZDU()` (same file), `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.XitOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZXT()`** — L8386 — `public string ZZXT(string s)`
  Sets or reads the X2TR button status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXU()`** — L8411 — `public string ZZXU(string s)`
  Increments XIT
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXV()`** — L8438 — `public string ZZXV(string s)`
  Reads VFO combined status
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZYA()`** — L8471 — `public string ZZYA(string s)`
  Reads or sets the VAC2 Direct I/Q checkbox on the setup form
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZYB()`** — L8496 — `public string ZZYB(string s)`
  Reads or sets the VAC2 IQ Calibrate checkbox on the setup form
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZYC()`** — L8521 — `public string ZZYC(string s)`
  Reads or sets the FM mic gain
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZYR()`** — L8546 — `public string ZZYR(string s)`
  Sets or reads the Rx1/RX2 radio button in collapsed mode
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZB()`** — L8563 — `public string ZZZB()`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ZeroBeatPress()` (`Console/Midi2CatCommands.cs`)
- **`.ZZZZ()`** — L8570 — `public string ZZZZ()`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZT()`** — L8581 — `public string ZZZT(string s)`
  Zooms to band MW0LGE_21k9 depending on ztb mode 0 will recall 1 will store query will return 1 if ztb is set up as store/recall
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.ZoomToBandRecall()` (`Console/Midi2CatCommands.cs`), `.ZoomToBandStore()` (`Console/Midi2CatCommands.cs`)
- **`.ZZZQ()`** — L8607 — `public string ZZZQ(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX1AutoAGC()` (`Console/Midi2CatCommands.cs`)
- **`.ZZZR()`** — L8633 — `public string ZZZR(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.RX2AutoAGC()` (`Console/Midi2CatCommands.cs`)
- **`.ZZZM()`** — L8660 — `public string ZZZM(string s)`
  hardware model string
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZV()`** — L8671 — `public string ZZZV(string s)`
  hardware version title string
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZW()`** — L8682 — `public string ZZZW(string s)`
  swap vfo wheels, vfoA becomes vfoB, B becomes A
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZZN()`** — L8701 — `public string ZZZN(string s)`
  [2.10.1.0]MW0LGE enable/disable quick split mode
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.QuickSplitOnOff()` (`Console/Midi2CatCommands.cs`), `.QuickSplitOnOffandSplitOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.ZZZO()`** — L8723 — `public string ZZZO(string s)`
  [2.10.1.0]MW0LGE enable/disable quick split and turn split on/off at same time
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`)
- **`.ZZXA()`** — L8748 — `public string ZZXA(string s)`
  Called by: `.ParseExtended()` (`Console/CAT/CATParser.cs`), `.AudioAmpOnOff()` (`Console/Midi2CatCommands.cs`)
- **`.AddLeadingZeros()`** — L8773 — `private string AddLeadingZeros(int n, int pad_len = -1)`
  Adds leading zeros.
  Called by: `.AG()` (same file), `.MG()` (same file), `.SQ()` (same file), `.ZZAA()` (same file), `.ZZAB()` (same file), `.ZZAC()` (same file) — and 79 more
- **`.JustSuffix()`** — L8793 — `private string JustSuffix(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OffsetDirection2String()`** — L8805 — `private string OffsetDirection2String()`
  Called by: `.ZZOS()` (same file)
- **`.String2OffsetDirection()`** — L8827 — `private void String2OffsetDirection(string s)`
  Called by: `.ZZOS()` (same file)
- **`.String2CTCSSFreq()`** — L8846 — `private double String2CTCSSFreq(string s)`
  Called by: `.ZZTB()` (same file)
- **`.CTCSSFreq2String()`** — L9005 — `private string CTCSSFreq2String(int freq)`
  Called by: `.ZZTB()` (same file)
- **`.GetMemoryList()`** — L9164 — `private SortableBindingList<MemoryRecord> GetMemoryList()`
  Returns memory list.
  Called by: `.GetChannelRecord()` (same file), `.GetIndex()` (same file)
- **`.GetChannelRecord()`** — L9178 — `private MemoryRecord GetChannelRecord(string channel)`
  Returns channel record.
  Called by: `.ZZMW()` (same file), `.ZZMZ()` (same file)
- **`.GetIndex()`** — L9197 — `private int GetIndex(string channel)`
  Returns index.
  Called by: `.ZZMX()` (same file), `.ZZMZ()` (same file)
- **`.GetNextChannelNumber()`** — L9216 — `private int GetNextChannelNumber()`
  Returns next channel number.
  Called by: `.ZZMY()` (same file)
- **`.StrVFOFreq()`** — L9239 — `private string StrVFOFreq(string vfo)`
  Converts a vfo frequency to a proper CAT frequency string
  Called by: `.ZZFA()` (same file), `.ZZFB()` (same file), `.ZZQM()` (same file)
- **`.Filter2String()`** — L9282 — `public string Filter2String(Filter f)`
  Called by: `.FW()` (same file)
- **`.String2Filter()`** — L9331 — `public Filter String2Filter(string f)`
  Called by: `.FW()` (same file)
- **`.SetFilterCenterAndWidth()`** — L9381 — `private void SetFilterCenterAndWidth(int center, int width)`
  set variable filter 1 to indicate center and width if either center or width is zero, current value of center or width is contained fixme ... what should this thing do for am, fm, dsb ... ignore width?
  Called by: `.ZZSF()` (same file)
- **`.Frequency2Code()`** — L9424 — `private string Frequency2Code(int f, string n)`
  Converts interger filter frequency into Kenwood SL/SH codes
  Called by: `.SH()` (same file), `.SL()` (same file), `.SetFilter()` (same file)
- **`.Code2Frequency()`** — L9525 — `private int Code2Frequency(string c, string n)`
  Converts a frequency code pair to frequency in hz according to the Kenwood TS-2000 spec. Receives code and calling methd as parameters
  Called by: `.SetFilter()` (same file)
- **`.SetFilter()`** — L9811 — `private void SetFilter(string c, string n)`
  Sets filter.
  Called by: `.SH()` (same file), `.SL()` (same file)
- **`.String2Mode()`** — L9874 — `public void String2Mode(string pIndex)`
  Called by: `.ZZMD()` (same file)
- **`.Mode2String()`** — L9919 — `public string Mode2String(DSPMode pMode)`
  Called by: `.ZZIF()` (same file), `.ZZMD()` (same file), `.ZZME()` (same file)
- **`.KString2Mode()`** — L9971 — `public void KString2Mode(string pIndex)`
  converts Kenwood single digit mode code to SDR mode
  Called by: `.MD()` (same file)
- **`.Mode2KString()`** — L10022 — `public string Mode2KString(DSPMode pMode)`
  converts SDR mode to Kenwood single digit mode code
  Called by: `.IF()` (same file), `.MD()` (same file)
- **`.MakeBandList()`** — L10072 — `private void MakeBandList()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBandGroup()`** — L10090 — `private void SetBandGroup(int band)`
  Sets band group.
  Called by: `.GetBand()` (same file)
- **`.GetBand()`** — L10102 — `private string GetBand(string b)`
  Returns band.
  Called by: `.ZZBS()` (same file)
- **`.BandUp()`** — L10129 — `private void BandUp()`
  Called by: `.ZZBU()` (same file)
- **`.BandDown()`** — L10142 — `private void BandDown()`
  Called by: `.ZZBD()` (same file)
- **`.Band2String()`** — L10155 — `private string Band2String(Band pBand)`
  Called by: `.ZZBT()` (same file), `.GetBand()` (same file)
- **`.String2Band()`** — L10253 — `private Band String2Band(string pBand)`
  Called by: `.ZZBT()` (same file), `.GetBand()` (same file)
- **`.Step2Freq()`** — L10355 — `private double Step2Freq(int step)`
  Called by: `.ZZAD()` (same file), `.ZZAU()` (same file), `.ZZBM()` (same file), `.ZZBP()` (same file)
- **`.Step2String()`** — L10412 — `private string Step2String(int pSize)`
  Called by: `.IF()` (same file), `.ZZIF()` (same file), `.ZZST()` (same file)
- **`.String2RXMeter()`** — L10508 — `private void String2RXMeter(int m)`
  Called by: `.ZZMR()` (same file)
- **`.RXMeter2String()`** — L10513 — `private string RXMeter2String()`
  Called by: `.ZZMR()` (same file)
- **`.String2TXMeter()`** — L10518 — `private void String2TXMeter(int m)`
  Called by: `.ZZMT()` (same file)
- **`.TXMeter2String()`** — L10523 — `private string TXMeter2String()`
  Called by: `.ZZMT()` (same file)
- **`.CAT2RigType()`** — L10532 — `private string CAT2RigType()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RigType2CAT()`** — L10537 — `private string RigType2CAT()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Width2Index()`** — L10546 — `private string Width2Index(int txt)`
  Called by: `.ZZHA()` (same file), `.ZZHR()` (same file), `.ZZHT()` (same file), `.ZZHU()` (same file), `.ZZHV()` (same file), `.ZZHW()` (same file) — and 1 more
- **`.Index2Width()`** — L10579 — `private int Index2Width(string ndx)`
  Called by: `.ZZHA()` (same file), `.ZZHR()` (same file), `.ZZHT()` (same file), `.ZZHU()` (same file), `.ZZHV()` (same file), `.ZZHW()` (same file) — and 1 more

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/CATCommands.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
