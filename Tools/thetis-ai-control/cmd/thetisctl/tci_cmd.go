package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"thetisctl/internal/safety"
	"thetisctl/internal/tci"
)

func runTCI(rawArgs []string) error {
	a := parseArgs(rawArgs)
	if len(a.pos) == 0 {
		return fmt.Errorf("tci: missing command; run 'thetisctl help'")
	}
	host := a.flag("host", "")
	if host == "" {
		return fmt.Errorf("tci: --host is required (Thetis is a separate, remote instance — never assume localhost)")
	}
	port := a.flag("port", "50001")
	timeout := parseDuration(a.flag("timeout", "5s"), 5*time.Second)

	addr := net.JoinHostPort(host, port)
	conn, err := tci.Dial(addr, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := tci.NewClient(conn)

	cmd, args := a.pos[0], a.pos[1:]
	switch cmd {
	case "vfo":
		return tciVFO(client, args)
	case "modulation":
		return tciSet2("modulation", args, client.SetModulation)
	case "split":
		return tciToggle("split", args, client.SetSplitEnable)
	case "rit":
		return tciToggle("rit", args, client.SetRITEnable)
	case "xit":
		return tciToggle("xit", args, client.SetXITEnable)
	case "rit-offset":
		return tciSetIntInt("rit-offset", args, client.SetRITOffsetHz)
	case "xit-offset":
		return tciSetIntInt("xit-offset", args, client.SetXITOffsetHz)
	case "filter":
		return tciFilter(client, args)
	case "atten":
		return tciSetIntInt("atten", args, client.SetStepAttenuatorDB)
	case "preamp":
		return tciSetIntInt("preamp", args, client.SetPreampAttenuatorDB)
	case "agc":
		return tciSet2("agc", args, client.SetAGCMode)
	case "agc-gain":
		return tciSetIntInt("agc-gain", args, client.SetAGCGain)
	case "drive":
		return tciSetIntInt("drive", args, client.SetDrive)
	case "tune":
		return tciTune(client, args, a)
	case "ptt":
		return tciPTT(client, args, a)
	case "rx-audio":
		return tciRxAudio(client, args, a)
	case "tx-audio":
		return tciTxAudio(client, args, a)
	case "query":
		return tciQuery(client, args)
	default:
		return fmt.Errorf("tci: unknown command %q", cmd)
	}
}

func parseRx(s string) (int, error) {
	rx, err := strconv.Atoi(s)
	if err != nil || (rx != 0 && rx != 1) {
		return 0, fmt.Errorf("rx must be 0 (RX1) or 1 (RX2), got %q", s)
	}
	return rx, nil
}

func tciVFO(client *tci.Client, args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("vfo: usage: vfo <rx> <chan 0|1> <hz>")
	}
	rx, err := parseRx(args[0])
	if err != nil {
		return fmt.Errorf("vfo: %w", err)
	}
	ch, err := strconv.Atoi(args[1])
	if err != nil || (ch != 0 && ch != 1) {
		return fmt.Errorf("vfo: chan must be 0 (VFO A) or 1 (VFO B), got %q", args[1])
	}
	hz, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return fmt.Errorf("vfo: invalid Hz value %q: %w", args[2], err)
	}
	if err := client.SetVFOFreqHz(rx, ch, hz); err != nil {
		return err
	}
	fmt.Printf("vfo %d chan %d set to %d Hz\n", rx, ch, hz)
	return nil
}

func tciSet2(name string, args []string, set func(int, string) error) error {
	if len(args) != 2 {
		return fmt.Errorf("%s: usage: %s <rx> <value>", name, name)
	}
	rx, err := parseRx(args[0])
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	if err := set(rx, args[1]); err != nil {
		return err
	}
	fmt.Printf("%s %d set to %s\n", name, rx, args[1])
	return nil
}

func tciSetIntInt(name string, args []string, set func(int, int) error) error {
	if len(args) != 2 {
		return fmt.Errorf("%s: usage: %s <rx> <value>", name, name)
	}
	rx, err := parseRx(args[0])
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	v, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%s: invalid value %q: %w", name, args[1], err)
	}
	if err := set(rx, v); err != nil {
		return err
	}
	fmt.Printf("%s %d set to %d\n", name, rx, v)
	return nil
}

func tciToggle(name string, args []string, set func(int, bool) error) error {
	if len(args) != 2 {
		return fmt.Errorf("%s: usage: %s <rx> on|off", name, name)
	}
	rx, err := parseRx(args[0])
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	var on bool
	switch args[1] {
	case "on":
		on = true
	case "off":
		on = false
	default:
		return fmt.Errorf("%s: unknown value %q (want on|off)", name, args[1])
	}
	if err := set(rx, on); err != nil {
		return err
	}
	fmt.Printf("%s %d set to %v\n", name, rx, on)
	return nil
}

func tciFilter(client *tci.Client, args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("filter: usage: filter <rx> <lowHz> <highHz>")
	}
	rx, err := parseRx(args[0])
	if err != nil {
		return fmt.Errorf("filter: %w", err)
	}
	low, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("filter: invalid lowHz %q: %w", args[1], err)
	}
	high, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("filter: invalid highHz %q: %w", args[2], err)
	}
	if err := client.SetFilterBand(rx, low, high); err != nil {
		return err
	}
	fmt.Printf("filter %d set to [%d, %d] Hz\n", rx, low, high)
	return nil
}

// tciTune and tciPTT are TCI's TX-capable commands: both gate real keying
// behind the safety confirmation phrase and auto-unkey after --hold.

func tciTune(client *tci.Client, args []string, a parsedArgs) error {
	if len(args) != 2 {
		return fmt.Errorf("tune: usage: tune <rx> on|off --confirm-tx=<phrase> [--hold 3s]")
	}
	rx, err := parseRx(args[0])
	if err != nil {
		return fmt.Errorf("tune: %w", err)
	}
	if args[1] == "off" {
		return client.SetTune(rx, false)
	}
	if args[1] != "on" {
		return fmt.Errorf("tune: unknown value %q (want on|off)", args[1])
	}
	hold := parseDuration(a.flag("hold", "3s"), 3*time.Second)
	dec, err := safety.Check(a.flag("confirm-tx", ""), isTerminal(os.Stdin), stdinPrompt)
	if err != nil {
		return err
	}
	if dec.DryRun {
		fmt.Printf("[dry-run] would send: tune:%d,true; ... (hold %s) ... tune:%d,false;\n", rx, hold, rx)
		fmt.Println("Pass --confirm-tx=" + safety.ConfirmPhrase + " to actually key the transmitter.")
		return nil
	}
	if err := client.SetTune(rx, true); err != nil {
		return err
	}
	fmt.Printf("TUNE ON (rx %d) — auto-unkeying after %s\n", rx, hold)
	time.Sleep(hold)
	if err := client.SetTune(rx, false); err != nil {
		return fmt.Errorf("TUNE ON succeeded but auto-unkey failed, radio may still be keyed: %w", err)
	}
	fmt.Println("TUNE OFF")
	return nil
}

func tciPTT(client *tci.Client, args []string, a parsedArgs) error {
	if len(args) != 2 {
		return fmt.Errorf("ptt: usage: ptt <rx> on|off --confirm-tx=<phrase> [--hold 3s] [--audio]")
	}
	rx, err := parseRx(args[0])
	if err != nil {
		return fmt.Errorf("ptt: %w", err)
	}
	useAudio := a.has("audio")
	if args[1] == "off" {
		if useAudio {
			return client.SetTrxTCIAudio(rx, false)
		}
		return client.SetTrx(rx, false)
	}
	if args[1] != "on" {
		return fmt.Errorf("ptt: unknown value %q (want on|off)", args[1])
	}
	hold := parseDuration(a.flag("hold", "3s"), 3*time.Second)
	dec, err := safety.Check(a.flag("confirm-tx", ""), isTerminal(os.Stdin), stdinPrompt)
	if err != nil {
		return err
	}
	if dec.DryRun {
		fmt.Printf("[dry-run] would send: trx:%d,true%s; ... (hold %s) ... trx:%d,false;\n",
			rx, audioSuffix(useAudio), hold, rx)
		fmt.Println("Pass --confirm-tx=" + safety.ConfirmPhrase + " to actually key the transmitter.")
		return nil
	}
	setOn := client.SetTrx
	if useAudio {
		setOn = client.SetTrxTCIAudio
	}
	if err := setOn(rx, true); err != nil {
		return err
	}
	fmt.Printf("PTT ON (rx %d) — auto-unkeying after %s\n", rx, hold)
	time.Sleep(hold)
	if err := client.SetTrx(rx, false); err != nil {
		return fmt.Errorf("PTT ON succeeded but auto-unkey failed, radio may still be keyed: %w", err)
	}
	fmt.Println("PTT OFF")
	return nil
}

func audioSuffix(useAudio bool) string {
	if useAudio {
		return ",tci"
	}
	return ""
}

func tciQuery(client *tci.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("query: usage: query <cmd> [args...]")
	}
	if err := client.SendCmd(args[0], args[1:]...); err != nil {
		return err
	}
	cmd, replyArgs, err := client.RecvCmd()
	if err != nil {
		return err
	}
	fmt.Printf("%s: %v\n", cmd, replyArgs)
	return nil
}

func tciRxAudio(client *tci.Client, args []string, a parsedArgs) error {
	if len(args) < 1 {
		return fmt.Errorf("rx-audio: usage: rx-audio capture|stream <rx> [--duration 10s] [--out file.wav] [--sample-type float32]")
	}
	mode, rest := args[0], args[1:]
	if mode != "capture" && mode != "stream" {
		return fmt.Errorf("rx-audio: unknown mode %q (want capture|stream)", mode)
	}
	if len(rest) != 1 {
		return fmt.Errorf("rx-audio %s: usage: rx-audio %s <rx> [--duration 10s] [--out file.wav]", mode, mode)
	}
	rx, err := parseRx(rest[0])
	if err != nil {
		return fmt.Errorf("rx-audio: %w", err)
	}
	duration := parseDuration(a.flag("duration", "10s"), 10*time.Second)
	sampleType := parseSampleType(a.flag("sample-type", "float32"))

	if mode == "capture" && a.flag("out", "") == "" {
		return fmt.Errorf("rx-audio capture: --out <file.wav> is required")
	}

	if err := client.SetAudioSampleType(sampleType); err != nil {
		return err
	}
	if err := client.StartAudio(rx); err != nil {
		return err
	}
	defer client.StopAudio(rx)

	var (
		buffered   []float32
		sampleRate = 48000
		channels   = 1
	)
	deadline := time.Now().Add(duration)
	for time.Now().Before(deadline) {
		h, data, err := client.RecvAudioFrame()
		if err != nil {
			if isTimeoutErr(err) {
				continue
			}
			return err
		}
		if h.StreamType != tci.StreamRXAudio || h.ReceiverID != rx {
			continue
		}
		if h.SampleRate > 0 {
			sampleRate = h.SampleRate
		}
		if h.Channels > 0 {
			channels = h.Channels
		}
		samples := tci.DecodeSamples(data, h.SampleType)
		if mode == "stream" {
			if err := writeFloat32LE(os.Stdout, samples); err != nil {
				return err
			}
		} else {
			buffered = append(buffered, samples...)
		}
	}

	if mode == "capture" {
		out := a.flag("out", "")
		format := tci.WAVFormat{SampleRate: sampleRate, Channels: channels, BitsPerSample: 32, Float: true}
		if err := tci.WriteWAV(out, format, buffered); err != nil {
			return err
		}
		fmt.Printf("captured %.2fs (%d samples, %d Hz, %d ch) to %s\n",
			float64(len(buffered))/float64(channels)/float64(sampleRate), len(buffered), sampleRate, channels, out)
	}
	return nil
}

func writeFloat32LE(w *os.File, samples []float32) error {
	buf := make([]byte, len(samples)*4)
	for i, s := range samples {
		binary.LittleEndian.PutUint32(buf[i*4:], math.Float32bits(s))
	}
	_, err := w.Write(buf)
	return err
}

func isTimeoutErr(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

func parseSampleType(s string) tci.SampleType {
	switch s {
	case "int16":
		return tci.SampleInt16
	case "int24":
		return tci.SampleInt24
	case "int32":
		return tci.SampleInt32
	default:
		return tci.SampleFloat32
	}
}

// tciTxAudio is TCI's most consequential TX-capable command: it streams a
// WAV file's audio as TX_AUDIO_STREAM frames while PTT is held via the
// "tci"-audio-source form of trx, which genuinely modulates and transmits
// RF (Console/cmaster.cs TCITxThreadProc drains this straight into the DSP
// TX chain). Gated behind the safety confirmation phrase; always applies a
// hard --max-duration cap and unkeys on completion, error, or Ctrl-C.
func tciTxAudio(client *tci.Client, args []string, a parsedArgs) error {
	if len(args) != 2 || args[0] != "send" {
		return fmt.Errorf("tx-audio: usage: tx-audio send <rx> --file tone.wav --confirm-tx=<phrase> [--max-duration 10s] [--sample-type int16]")
	}
	rx, err := parseRx(args[1])
	if err != nil {
		return fmt.Errorf("tx-audio: %w", err)
	}
	file := a.flag("file", "")
	if file == "" {
		return fmt.Errorf("tx-audio send: --file <wav> is required")
	}
	maxDuration := parseDuration(a.flag("max-duration", "10s"), 10*time.Second)
	sampleType := parseSampleType(a.flag("sample-type", "int16"))

	format, samples, err := tci.ReadWAV(file)
	if err != nil {
		return err
	}
	if format.Channels < 1 {
		format.Channels = 1
	}
	totalDuration := time.Duration(float64(len(samples)) / float64(format.Channels) / float64(format.SampleRate) * float64(time.Second))
	truncated := false
	if totalDuration > maxDuration {
		keepFrames := int(maxDuration.Seconds() * float64(format.SampleRate))
		keepSamples := keepFrames * format.Channels
		if keepSamples < len(samples) {
			samples = samples[:keepSamples]
			truncated = true
		}
		totalDuration = maxDuration
	}
	peak := peakAbs(samples)

	dec, err := safety.Check(a.flag("confirm-tx", ""), isTerminal(os.Stdin), stdinPrompt)
	if err != nil {
		return err
	}
	if dec.DryRun {
		fmt.Printf("[dry-run] would send: trx:%d,true,tci; then stream %s of TX audio from %s (%d Hz, %d ch, peak %.3f%s) as %s frames; then trx:%d,false,tci;\n",
			rx, totalDuration, file, format.SampleRate, format.Channels, peak, truncatedNote(truncated), sampleType.WireName(), rx)
		fmt.Println("Pass --confirm-tx=" + safety.ConfirmPhrase + " to actually transmit this audio.")
		return nil
	}

	if err := client.SetAudioSampleType(sampleType); err != nil {
		return err
	}

	// Always unkey on completion, error, or Ctrl-C — never leave the radio
	// keyed because this process exited unexpectedly.
	unkeyed := false
	unkey := func() {
		if unkeyed {
			return
		}
		unkeyed = true
		_ = client.SetTrxTCIAudio(rx, false)
	}
	defer unkey()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)
	go func() {
		if _, ok := <-sigCh; ok {
			unkey()
			fmt.Fprintln(os.Stderr, "\ntx-audio: interrupted, unkeyed")
			os.Exit(130)
		}
	}()

	if err := client.SetTrxTCIAudio(rx, true); err != nil {
		return err
	}
	fmt.Printf("TX ON (rx %d, TCI audio) — streaming %s%s\n", rx, totalDuration, truncatedNote(truncated))

	const chunkFrames = 2048
	chunkSamples := chunkFrames * format.Channels
	for off := 0; off < len(samples); off += chunkSamples {
		end := off + chunkSamples
		if end > len(samples) {
			end = len(samples)
		}
		chunk := samples[off:end]
		h := tci.StreamHeader{
			ReceiverID: rx,
			SampleRate: format.SampleRate,
			SampleType: sampleType,
			Length:     len(chunk),
			StreamType: tci.StreamTXAudio,
			Channels:   format.Channels,
		}
		if err := client.SendAudioFrame(h, tci.EncodeSamples(chunk, sampleType)); err != nil {
			return err
		}
		frames := len(chunk) / format.Channels
		time.Sleep(time.Duration(float64(frames)/float64(format.SampleRate)*1000) * time.Millisecond)
	}

	unkey()
	fmt.Println("TX OFF")
	return nil
}

func truncatedNote(truncated bool) string {
	if truncated {
		return " [truncated to --max-duration]"
	}
	return ""
}

func peakAbs(samples []float32) float32 {
	var peak float32
	for _, s := range samples {
		if s < 0 {
			s = -s
		}
		if s > peak {
			peak = s
		}
	}
	return peak
}
