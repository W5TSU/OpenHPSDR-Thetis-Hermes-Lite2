#!/usr/bin/env bash
# Plays a captured off-air wav into freedv-gui's actual GUI (waterfall, sync
# LED, decoded audio) via a PulseAudio/PipeWire virtual sink, instead of the
# headless `-ut rx` test mode used elsewhere in this repo's testing. Useful
# when you want to *watch* sync attempt to happen live, or just listen to
# whatever the GUI manages to decode, rather than trust log output alone.
#
# Usage: ./listen_offair_in_freedv_gui.sh [--mode radev1|700e] [path/to/capture.wav]
#   Defaults to Raw-FreeDV-Off-Air.wav in this same directory if no path given.
#
# What it does:
#   1. Creates a "FreeDV_Test_Sink" virtual audio sink (idempotent -- skips
#      if one already exists).
#   2. Launches freedv-gui if it isn't already running. freedv-gui's own
#      config (~/.freedv) is left alone -- if it's already pointed at
#      freedv_test_sink.monitor for its receive input (soundCard1InDeviceName),
#      no manual device selection is needed; check Tools -> Audio Config in
#      the GUI if audio doesn't seem to be reaching it.
#   3. Waits for you to press Enter (so you have time to click the mode
#      button -- see --mode below -- and click Start in the GUI first).
#   4. Plays the wav file into the virtual sink -- freedv-gui hears it as
#      if it were live receive audio.
#
# --mode is NOT automated: this script cannot click freedv-gui's UI for
# you (no xdotool/wmctrl available in this environment, and this session
# is Wayland, where that kind of automation is unreliable even when those
# tools are installed). Instead --mode prints the exact control to click,
# confirmed from freedv-gui's own source (src/topFrame.cpp): there is a
# small boxed group titled "Mode" (mnemonic Alt+M) containing RADIO
# BUTTONS labelled RADEV1 / 700D / 700E / 1600 -- not a dropdown. Click
# the one matching --mode directly.
#
# You will NOT hear the raw file through your speakers this way -- it goes
# into the virtual sink, not your real output device. Use plain `paplay
# <file>` separately if you want to just listen to the raw capture.

set -euo pipefail

MODE=""
WAV_FILE=""
while [ $# -gt 0 ]; do
    case "$1" in
        --mode)
            MODE="${2:-}"
            shift 2
            ;;
        *)
            WAV_FILE="$1"
            shift
            ;;
    esac
done

WAV_FILE="${WAV_FILE:-$(dirname "$0")/Raw-FreeDV-Off-Air_iq.wav}"
SINK_NAME="freedv_test_sink"
FREEDV_BIN="$HOME/Development/freedv-gui/build_linux/src/freedv"

MODE_LABEL=""
case "${MODE,,}" in
    radev1) MODE_LABEL="RADEV1" ;;
    700e)   MODE_LABEL="700E" ;;
    "")     MODE_LABEL="" ;;
    *)
        echo "Error: --mode must be 'radev1' or '700e' (got '$MODE')" >&2
        exit 1
        ;;
esac

if [ ! -f "$WAV_FILE" ]; then
    echo "Error: wav file not found: $WAV_FILE" >&2
    exit 1
fi

if [ ! -x "$FREEDV_BIN" ]; then
    echo "Error: freedv-gui binary not found/executable at $FREEDV_BIN" >&2
    echo "(Build it first -- see FreeDV-Plan.md's freedv-gui build steps.)" >&2
    exit 1
fi

# --- 1. virtual sink, idempotent ---
if pactl list short sinks | grep -q "\b${SINK_NAME}\b"; then
    echo "Virtual sink '${SINK_NAME}' already exists, reusing it."
else
    echo "Creating virtual sink '${SINK_NAME}'..."
    pactl load-module module-null-sink \
        sink_name="${SINK_NAME}" \
        sink_properties=device.description="FreeDV_Test_Sink"
fi

# --- 2. launch freedv-gui if not already running ---
if pgrep -f "$FREEDV_BIN" > /dev/null; then
    echo "freedv-gui is already running, not launching a second instance."
else
    echo "Launching freedv-gui..."
    "$FREEDV_BIN" > /tmp/freedv_gui_launch.log 2>&1 &
    disown
    sleep 3
fi

# --- 3. wait for the operator to set mode + press Start in the GUI ---
echo ""
echo "In the freedv-gui window that just opened:"
if [ -n "$MODE_LABEL" ]; then
    echo "  1. Find the small boxed group titled 'Mode' and click its"
    echo "     '${MODE_LABEL}' RADIO BUTTON (not a dropdown -- the group also"
    echo "     has RADEV1/700D/700E/1600 as separate buttons)."
else
    echo "  1. Find the small boxed group titled 'Mode' and click the radio"
    echo "     button for whichever mode you want to test (RADEV1/700D/700E/1600)."
    echo "     Tip: pass --mode radev1 or --mode 700e next time to have this"
    echo "     printed for the exact mode you want."
fi
echo "  2. (Only if audio doesn't reach it) Tools -> Audio Config -> set the"
echo "     'From Radio' (receive) input device to 'Monitor of FreeDV_Test_Sink'."
echo "  3. Click the green Start button -- it should show 'Searching...'"
echo "     with a live waterfall."
echo ""
read -rp "Press Enter once you've clicked Start and are ready to play the file... "

# --- 4. play the capture into the virtual sink ---
echo "Playing '${WAV_FILE}' into ${SINK_NAME}..."
paplay --device="${SINK_NAME}" "$WAV_FILE"
echo "Playback finished."
