package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"thetisctl/internal/freedvreporter"
	"thetisctl/internal/tci"
)

// 20m band edges, matching freedvreporter's own freqToBandClass in its
// index.js (14000000-14350000 Hz) - the default scope per the operator's
// call: our HL2 test setup lives on 20m, and retuning far outside the
// currently-connected antenna/band plan isn't useful.
const band20mMinHz, band20mMaxHz = 14000000, 14350000

// Fixed panadapter marker colors for FreeDV Reporter spots. Both are
// UNSIGNED 32-bit ARGB packed as 0xAARRGGBB — see Client.Spot's doc comment
// in internal/tci/control.go for why (TCIServer.cs:4375 uses uint.TryParse).
const (
	spotColorTX   uint32 = 4292618270 // 0xFFDC281E = (255,220,40,30), bright red: this station is transmitting right now
	spotColorIdle uint32 = 4282153160 // 0xFF3C78C8 = (255,60,120,200), calmer blue: connected/monitoring, not transmitting
)

func runFreeDVReporter(rawArgs []string) error {
	a := parseArgs(rawArgs)
	if len(a.pos) == 0 {
		return fmt.Errorf("freedv-reporter: missing command; try 'watch'")
	}
	switch a.pos[0] {
	case "watch":
		return freedvReporterWatch(a)
	default:
		return fmt.Errorf("freedv-reporter: unknown command %q (try 'watch')", a.pos[0])
	}
}

func freedvReporterWatch(a parsedArgs) error {
	minFreq := int64(band20mMinHz)
	maxFreq := int64(band20mMaxHz)
	if v := a.flag("min-freq", ""); v != "" {
		f, err := parseInt64(v)
		if err != nil {
			return fmt.Errorf("freedv-reporter watch: --min-freq: %w", err)
		}
		minFreq = f
	}
	if v := a.flag("max-freq", ""); v != "" {
		f, err := parseInt64(v)
		if err != nil {
			return fmt.Errorf("freedv-reporter watch: --max-freq: %w", err)
		}
		maxFreq = f
	}

	tciHost := a.flag("tci", "")
	tciPort := a.flag("tci-port", "50001")
	tuneMode := a.flag("mode", "digu")
	doSpot := a.has("spot")
	noTune := a.has("no-tune")
	doSelfReport := a.has("self-report")
	selfReportCallsign := a.flag("callsign", "")
	selfReportGrid := a.flag("grid", "")
	selfReportRxOnly := a.has("rx-only")

	if doSpot && tciHost == "" {
		return fmt.Errorf("freedv-reporter watch: --spot requires --tci <host> (spots are pushed over TCI; there is nowhere else to send them)")
	}
	if doSelfReport && tciHost == "" {
		return fmt.Errorf("freedv-reporter watch: --self-report requires --tci <host> (self-reporting reads Thetis's live VFO/mode/TX state over TCI)")
	}
	if doSelfReport && selfReportCallsign == "" {
		return fmt.Errorf("freedv-reporter watch: --self-report requires --callsign <call>")
	}

	fmt.Printf("freedv-reporter watch: connecting to %s ...\n", freedvreporter.ReporterHost)
	client, err := freedvreporter.Dial(10 * time.Second)
	if err != nil {
		return fmt.Errorf("freedv-reporter watch: %w", err)
	}
	defer client.Close()
	fmt.Printf("freedv-reporter watch: connected. Watching %.3f-%.3f MHz",
		float64(minFreq)/1e6, float64(maxFreq)/1e6)
	if tciHost != "" && !noTune {
		fmt.Printf(" — will auto-tune %s (mode %s) on activity", net.JoinHostPort(tciHost, tciPort), tuneMode)
	}
	if doSpot {
		fmt.Printf(" — pushing spots to %s panadapter", net.JoinHostPort(tciHost, tciPort))
	}
	if doSelfReport {
		fmt.Printf(" — reporting %s to FreeDV Reporter", selfReportCallsign)
	}
	fmt.Println(". Ctrl-C to stop.")

	// Ctrl-C closes the reporter connection cleanly rather than leaving the
	// process to be killed mid-frame; ReadEvent's blocking read then returns
	// an error and the loop below exits.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		fmt.Println("\nfreedv-reporter watch: stopping...")
		client.Close()
	}()

	if doSelfReport {
		go runSelfReport(tciHost, tciPort, selfReportCallsign, selfReportGrid, selfReportRxOnly)
	}

	tracker := freedvreporter.NewTracker()
	var tciClient *tci.Client // lazily dialed on first qualifying tune/spot action, kept for reuse

	// ensureTCI dials tciClient on first use and reuses it thereafter. Both
	// the auto-tune block and the spot-pushing block below call this —
	// they may run in the same watch session (doSpot and auto-tune are
	// independent flags) and share one TCI connection rather than each
	// dialing its own.
	ensureTCI := func() bool {
		if tciClient != nil {
			return true
		}
		addr := net.JoinHostPort(tciHost, tciPort)
		conn, derr := tci.Dial(addr, 5*time.Second)
		if derr != nil {
			fmt.Printf("  -> connect to Thetis TCI at %s failed: %v\n", addr, derr)
			return false
		}
		tciClient = tci.NewClient(conn)
		return true
	}

	for {
		ev, err := client.ReadEvent()
		if err != nil {
			if tciClient != nil {
				tciClient.Close()
			}
			return fmt.Errorf("freedv-reporter watch: %w", err)
		}

		started, err := tracker.Apply(ev)
		if err != nil {
			continue
		}

		for _, ts := range started {
			s := ts.Station
			if s.RXOnly || s.Callsign == "" {
				continue // monitor-only / not-yet-fully-known connections can't be "transmitting" in a meaningful sense
			}
			if s.FreqHz < minFreq || s.FreqHz > maxFreq {
				continue
			}

			fmt.Printf("[%s] %s TX-ing on %.4f MHz (FreeDV mode: %s)\n",
				time.Now().Format("15:04:05"), s.Callsign, float64(s.FreqHz)/1e6, s.Mode)

			if noTune || tciHost == "" {
				continue
			}
			if !ensureTCI() {
				continue
			}

			if serr := tciClient.SetVFOFreqHz(0, 0, s.FreqHz); serr != nil {
				fmt.Printf("  -> auto-tune failed: set VFO: %v\n", serr)
				tciClient.Close()
				tciClient = nil
				continue
			}
			if serr := tciClient.SetModulation(0, tuneMode); serr != nil {
				fmt.Printf("  -> auto-tune failed: set mode: %v\n", serr)
				tciClient.Close()
				tciClient = nil
				continue
			}
			fmt.Printf("  -> retuned RX1 to %.4f MHz / %s\n", float64(s.FreqHz)/1e6, tuneMode)
		}

		// Spot-pushing is independent of the TxStarted-only started slice
		// above: it re-sends every currently-tracked station (with a valid
		// callsign and a known frequency) on every event, not just the ones
		// that just started transmitting. SpotManager2's dedup-by-callsign-
		// and-frequency treats a repeat send as a harmless refresh, not a
		// duplicate (see the design spec's "Trigger mechanism, precisely").
		if doSpot {
			for _, s := range tracker.Stations() {
				if s.Callsign == "" || s.FreqHz == 0 {
					continue
				}
				if s.FreqHz < minFreq || s.FreqHz > maxFreq {
					continue
				}
				if !ensureTCI() {
					break // TCI is down for this whole pass; the next reporter event will retry once
				}

				color := spotColorIdle
				text := s.Mode
				if s.Transmitting {
					color = spotColorTX
					if text != "" {
						text += " TX"
					} else {
						text = "TX"
					}
				}

				if serr := tciClient.Spot(s.Callsign, "digu", s.FreqHz, color, text); serr != nil {
					fmt.Printf("  -> spot failed for %s: %v\n", s.Callsign, serr)
					tciClient.Close()
					tciClient = nil
				}
			}
		}
	}
}

func parseInt64(s string) (int64, error) {
	var v int64
	_, err := fmt.Sscanf(s, "%d", &v)
	if err != nil {
		return 0, fmt.Errorf("not an integer: %q", s)
	}
	return v, nil
}

// runSelfReport bridges Thetis's own live TCI state (VFO frequency, demod
// mode, TX/MOX) into FreeDV Reporter's write-side protocol, publishing
// this station's presence. It uses its OWN dedicated TCI connection,
// deliberately never sharing freedvReporterWatch's tciClient variable —
// that variable is written to and nilled out by the auto-tune/spot
// features' own single-goroutine error handling; a second goroutine
// reading it without synchronization would be a data race. Runs until the
// process exits (see the call site's own comment for why no explicit
// shutdown signal is threaded in here), reconnecting both sides on any
// error after a short pause.
//
// KNOWN SIMPLIFICATION: the tx_report "mode" field below carries Thetis's
// raw DSPMode string (e.g. "USB", "DIGU") from a plain "modulation:" TCI
// query, not the actual active FreeDV codec name ("RADEV1", "700E") a real
// freedv-gui report would show. Getting the real codec name needs a CAT
// query (the ZZEX family) in addition to TCI — a second protocol
// dependency deliberately out of scope for this pass.
func runSelfReport(tciHost, tciPort, callsign, gridSquare string, rxOnly bool) {
	for {
		addr := net.JoinHostPort(tciHost, tciPort)
		stateConn, derr := tci.Dial(addr, 5*time.Second)
		if derr != nil {
			fmt.Printf("[self-report] connect to Thetis TCI at %s failed: %v; retrying in 5s\n", addr, derr)
			time.Sleep(5 * time.Second)
			continue
		}
		stateClient := tci.NewClient(stateConn)

		reportClient, rerr := freedvreporter.DialReport(callsign, gridSquare, rxOnly, true, 10*time.Second)
		if rerr != nil {
			fmt.Printf("[self-report] connect to FreeDV Reporter failed: %v; retrying in 5s\n", rerr)
			stateClient.Close()
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Printf("[self-report] reporting %s to FreeDV Reporter\n", callsign)

		selfReportQueryInitialState(stateClient, reportClient)
		selfReportListenLoop(stateClient, reportClient)

		reportClient.Close()
		stateClient.Close()
		fmt.Println("[self-report] connection lost; reconnecting in 5s")
		time.Sleep(5 * time.Second)
	}
}

// selfReportQueryInitialState actively queries Thetis's current VFO A
// frequency, demod mode, and MOX state right after connecting (rather
// than only waiting passively for the next change), so the very first
// report reflects reality instead of a stale/empty value. Query wire
// forms confirmed directly against TCIServer.cs: "vfo:0,0;" (2 args) is
// the GET form replying "vfo:0,0,<hz>;" (TCIServer.cs:3859-3969,
// sendVFO/TCIServer.cs:2099-2132); "modulation:0;" (1 arg) is the GET form
// replying "modulation:0,<MODE>;" uppercase (TCIServer.cs:3972-4074,
// sendMode/TCIServer.cs:2174-2195); "trx:0;" (1 arg) is the GET form
// replying "trx:0,<true|false>[,tci];" (TCIServer.cs:3594-3694,
// sendMOX/TCIServer.cs:2159-2169). Errors here are logged and swallowed —
// a failed initial query just means the first report is skipped; the
// listen loop below will emit as soon as anything actually changes.
func selfReportQueryInitialState(stateClient *tci.Client, reportClient *freedvreporter.Client) {
	if err := stateClient.SendCmd("vfo", "0", "0"); err == nil {
		if cmd, args, rerr := stateClient.RecvCmd(); rerr == nil && cmd == "vfo" && len(args) >= 3 {
			if hz, perr := strconv.ParseInt(args[2], 10, 64); perr == nil {
				emitFreqChange(reportClient, hz)
			}
		}
	}
	if err := stateClient.SendCmd("modulation", "0"); err == nil {
		if cmd, args, rerr := stateClient.RecvCmd(); rerr == nil && cmd == "modulation" && len(args) >= 2 {
			lastKnownMode = args[1]
		}
	}
	if err := stateClient.SendCmd("trx", "0"); err == nil {
		if cmd, args, rerr := stateClient.RecvCmd(); rerr == nil && cmd == "trx" && len(args) >= 2 {
			if tx, perr := strconv.ParseBool(args[1]); perr == nil {
				emitTxReport(reportClient, lastKnownMode, tx)
			}
		}
	}
}

// lastKnownMode caches the most recently seen modulation string so a
// standalone "trx:" broadcast (which carries no mode of its own) can
// still emit a complete tx_report — mirrors freedv-gui's own
// FreeDVReporter::transmitImpl_, which similarly remembers mode_
// alongside tx_ state.
var lastKnownMode string

// selfReportListenLoop passively listens for Thetis's own unsolicited
// vfo:/modulation:/trx: broadcasts (sent whenever that state changes) and
// translates each into the matching Emit call. Returns (without erroring
// further) as soon as RecvCmd fails, letting the caller reconnect both
// connections.
func selfReportListenLoop(stateClient *tci.Client, reportClient *freedvreporter.Client) {
	for {
		cmd, args, err := stateClient.RecvCmd()
		if err != nil {
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				// tci.Dial's timeout is reapplied as a read deadline before
				// every frame (internal/tci/ws.go), so this loop's own
				// passive listen naturally times out during any idle radio
				// state -- not a real disconnect. Keep listening rather
				// than tearing down and reconnecting both connections on a
				// ~5s cadence.
				continue
			}
			fmt.Printf("[self-report] TCI connection lost: %v\n", err)
			return
		}
		switch cmd {
		case "vfo":
			if len(args) >= 3 {
				if hz, perr := strconv.ParseInt(args[2], 10, 64); perr == nil {
					emitFreqChange(reportClient, hz)
				}
			}
		case "modulation":
			if len(args) >= 2 {
				lastKnownMode = args[1]
			}
		case "trx":
			if len(args) >= 2 {
				if tx, perr := strconv.ParseBool(args[1]); perr == nil {
					emitTxReport(reportClient, lastKnownMode, tx)
				}
			}
		}
	}
}

func emitFreqChange(reportClient *freedvreporter.Client, hz int64) {
	type freqChangePayload struct {
		Freq int64 `json:"freq"`
	}
	if err := reportClient.Emit("freq_change", freqChangePayload{Freq: hz}); err != nil {
		fmt.Printf("[self-report] emit freq_change failed: %v\n", err)
	}
}

func emitTxReport(reportClient *freedvreporter.Client, mode string, transmitting bool) {
	type txReportPayload struct {
		Mode         string `json:"mode"`
		Transmitting bool   `json:"transmitting"`
	}
	if err := reportClient.Emit("tx_report", txReportPayload{Mode: mode, Transmitting: transmitting}); err != nil {
		fmt.Printf("[self-report] emit tx_report failed: %v\n", err)
	}
}
