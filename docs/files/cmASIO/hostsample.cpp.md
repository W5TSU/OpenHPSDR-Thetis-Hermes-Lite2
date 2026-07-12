# `cmASIO/hostsample.cpp`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** The cmASIO DLL: thin host wrapper around the Steinberg ASIO SDK giving ChannelMaster direct ASIO driver access. The bundled `asiosdk_2.3.3.../` tree is the vendored Steinberg SDK.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmasio.c` (calls ×9)
- Uses (outgoing references to other files):
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/asio.cpp` (calls ×16)
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/asiosys.h` (imports ×1)
  - `cmASIO/asiosdk_2.3.3_2019-06-14/host/asiodrivers.cpp` (calls ×1)
  - `cmASIO/hostsample.h` (imports ×1)
- Most-referenced symbols from other files: `getASIOBaseInputChannel()` (×1), `getASIOBaseOutputChannel()` (×1), `getASIOBlockNum()` (×1), `getASIODriverString()` (×1), `getASIOInputMode()` (×1), `prepareASIO()` (×1), `unloadASIO()` (×1), `asioStart()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`init_asio_static_data()`** — L137 — `long init_asio_static_data (DriverInfo *asioDriverInfo)`
  Called by: `prepareASIO()` (same file)
- **`bufferSwitchTimeInfo()`** — L212 — `ASIOTime *bufferSwitchTimeInfo(ASIOTime *timeInfo, long index, ASIOBool processNow)`
  Called by: `bufferSwitch()` (same file)
- **`bufferSwitch()`** — L231 — `void bufferSwitch(long index, ASIOBool processNow)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`sampleRateChanged()`** — L251 — `void sampleRateChanged(ASIOSampleRate sRate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`asioMessages()`** — L262 — `long asioMessages(long selector, long value, void* message, double* opt)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_asio_buffers()`** — L327 — `ASIOError create_asio_buffers(DriverInfo* asioDriverInfo, long input_ch0, long input_ch1, long output_ch0, long output_ch1)`
  Constructor for the `asio_buffers` block: allocates its state/buffers and computes initial coefficients.
  Called by: `prepareASIO()` (same file)
- **`prepareASIO()`** — L492 — `int prepareASIO(int blocksize, int samplerate, char* asioDriverName, void (*CallbackASIO)(void* inputL, void* inputR, void* outputL, void* outputR), long input_base_channel, long o`
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`)
- **`unloadASIO()`** — L598 — `void unloadASIO()`
  Called by: `destroy_cmasio()` (`ChannelMaster/cmasio.c`)
- **`getASIODriverString()`** — L609 — `long getASIODriverString(void* szData)`
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`)
- **`getASIOBlockNum()`** — L634 — `long getASIOBlockNum(void* dwData)`
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`)
- **`asioStart()`** — L657 — `long asioStart()`
  Called by: `cm_asioStart()` (`ChannelMaster/cmasio.c`)
- **`asioStop()`** — L674 — `long asioStop()`
  Called by: `cm_asioStop()` (`ChannelMaster/cmasio.c`)
- **`getASIOBaseInputChannel()`** — L691 — `long getASIOBaseInputChannel(void* dwData)`
  [2.10.3.13]MW0LGE get base channel numbers for input and output, and input mode
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`)
- **`getASIOBaseOutputChannel()`** — L713 — `long getASIOBaseOutputChannel(void* dwData)`
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`)
- **`getASIOInputMode()`** — L735 — `long getASIOInputMode(void* dwData)`
  Called by: `create_cmasio()` (`ChannelMaster/cmasio.c`)
- **`get_sys_reference_time()`** — L760 — `unsigned long get_sys_reference_time()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

### Types

#### `DriverInfo` (type, L58)


#### `driverInfo` (type, L61)

_No extracted members._

#### `inputChannels` (type, L64)

_No extracted members._

#### `outputChannels` (type, L65)

_No extracted members._

#### `minSize` (type, L68)

_No extracted members._

#### `maxSize` (type, L69)

_No extracted members._

#### `preferredSize` (type, L70)

_No extracted members._

#### `granularity` (type, L71)

_No extracted members._

#### `sampleRate` (type, L74)

_No extracted members._

#### `postOutput` (type, L77)

_No extracted members._

#### `inputLatency` (type, L80)

_No extracted members._

#### `outputLatency` (type, L81)

_No extracted members._

#### `inputBuffers` (type, L84)

_No extracted members._

#### `outputBuffers` (type, L85)

_No extracted members._

#### `bufferInfos` (type, L86)

_No extracted members._

#### `channelInfos` (type, L89)

_No extracted members._

#### `nanoSeconds` (type, L94)

_No extracted members._

#### `samples` (type, L95)

_No extracted members._

#### `tcSamples` (type, L96)

_No extracted members._

#### `tInfo` (type, L99)

_No extracted members._

#### `sysRefTime` (type, L100)

_No extracted members._

#### `stopped` (type, L103)

_No extracted members._

#### `requestedDriverName` (type, L106)

_No extracted members._

#### `requestedBufferSize` (type, L107)

_No extracted members._

#### `requestedSampleRate` (type, L108)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/cmASIO/hostsample.cpp`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
