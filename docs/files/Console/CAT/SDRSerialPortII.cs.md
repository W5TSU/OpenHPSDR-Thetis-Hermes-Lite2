# `Console/CAT/SDRSerialPortII.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Serial-port wrapper and the per-port listener threads (CAT, PTT, keyer ports).

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/SIOListenerII.cs` (calls ×36, references ×7)
  - `Console/setup.cs` (calls ×10)
  - `Console/CAT/SerialPortPTT.cs` (calls ×7, references ×1)
  - `Console/CW/CWInput.cs` (calls ×4, references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Destroy()` (×15), `.put()` (×8), `.Create()` (×8), `.setCommParms()` (×7), `.StringToParity()` (×5), `.StringToStopBits()` (×5), `.Open()` (×2), `.Close()` (×2)

## Outline

### Types

#### `SDRSerialPort` (type, L31)

- `.Open()` — L50
- `.Close()` — L55
- `.StringToParity()` — L60
- `.StringToStopBits()` — L70
- `.setCommParms()` — L102
- `.put()` — L112
- `.Create()` — L119
- `.Destroy()` — L137
- `.isCTS()` — L150
- `.isDSR()` — L156
- `.isRI()` — L162
- `.isRLSD()` — L168
- `.setDTR()` — L174
- `.SerialErrorReceived()` — L180
- `.SerialPinChanged()` — L234
- `.SerialReceivedData()` — L339

#### `SDRSerialPort2` (type, L346)

- `.Open()` — L365
- `.Close()` — L370
- `.StringToParity()` — L375
- `.StringToStopBits()` — L385
- `.setCommParms()` — L418
- `.put()` — L428
- `.Create()` — L435
- `.Destroy()` — L453
- `.isCTS()` — L466
- `.isDSR()` — L472
- `.isRI()` — L478
- `.isRLSD()` — L484
- `.setDTR()` — L490
- `.SerialErrorReceived()` — L496
- `.SerialPinChanged()` — L550
- `.SerialReceivedData()` — L620

#### `SDRSerialPort3` (type, L627)

- `.Open()` — L646
- `.Close()` — L651
- `.StringToParity()` — L656
- `.StringToStopBits()` — L666
- `.setCommParms()` — L699
- `.put()` — L709
- `.Create()` — L716
- `.Destroy()` — L734
- `.isCTS()` — L747
- `.isDSR()` — L753
- `.isRI()` — L759
- `.isRLSD()` — L765
- `.setDTR()` — L771
- `.SerialErrorReceived()` — L777
- `.SerialPinChanged()` — L831
- `.SerialReceivedData()` — L901

#### `SDRSerialPort4` (type, L908)

- `.Open()` — L927
- `.Close()` — L932
- `.StringToParity()` — L937
- `.StringToStopBits()` — L947
- `.setCommParms()` — L980
- `.put()` — L990
- `.Create()` — L997
- `.Destroy()` — L1015
- `.isCTS()` — L1028
- `.isDSR()` — L1034
- `.isRI()` — L1040
- `.isRLSD()` — L1046
- `.setDTR()` — L1052
- `.SerialErrorReceived()` — L1058
- `.SerialPinChanged()` — L1112
- `.SerialReceivedData()` — L1182

#### `SDRSerialPort5` (type, L1190)

- `.Open()` — L1209
- `.Close()` — L1214
- `.StringToParity()` — L1219
- `.StringToStopBits()` — L1229
- `.setCommParms()` — L1262
- `.put()` — L1272
- `.Create()` — L1279
- `.Destroy()` — L1297
- `.isCTS()` — L1310
- `.isDSR()` — L1316
- `.isRI()` — L1322
- `.isRLSD()` — L1328
- `.setDTR()` — L1334
- `.SerialErrorReceived()` — L1340
- `.SerialPinChanged()` — L1394
- `.SerialReceivedData()` — L1464

#### `SDRSerialPort6` (type, L1472)

- `.Open()` — L1491
- `.Close()` — L1496
- `.StringToParity()` — L1501
- `.StringToStopBits()` — L1511
- `.setCommParms()` — L1544
- `.put()` — L1554
- `.Create()` — L1561
- `.Destroy()` — L1579
- `.isCTS()` — L1592
- `.isDSR()` — L1598
- `.isRI()` — L1604
- `.isRLSD()` — L1610
- `.setDTR()` — L1616
- `.SerialErrorReceived()` — L1622
- `.SerialPinChanged()` — L1676
- `.SerialReceivedData()` — L1746

#### `SDRSerialPort7` (type, L1754)

- `.Open()` — L1773
- `.Close()` — L1778
- `.StringToParity()` — L1783
- `.StringToStopBits()` — L1793
- `.setCommParms()` — L1826
- `.put()` — L1836
- `.Create()` — L1843
- `.Destroy()` — L1861
- `.isCTS()` — L1874
- `.isDSR()` — L1880
- `.isRI()` — L1886
- `.isRLSD()` — L1892
- `.setDTR()` — L1898
- `.SerialErrorReceived()` — L1904
- `.SerialPinChanged()` — L1958
- `.SerialReceivedData()` — L2028

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/SDRSerialPortII.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
