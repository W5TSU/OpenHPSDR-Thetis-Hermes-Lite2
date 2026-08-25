package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
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

	if doSpot && tciHost == "" {
		return fmt.Errorf("freedv-reporter watch: --spot requires --tci <host> (spots are pushed over TCI; there is nowhere else to send them)")
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
