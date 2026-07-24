package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"thetisctl/internal/cat"
	"thetisctl/internal/safety"
)

func runCAT(rawArgs []string) error {
	a := parseArgs(rawArgs)
	if len(a.pos) == 0 {
		return fmt.Errorf("cat: missing command; run 'thetisctl help'")
	}
	host := a.flag("host", "")
	if host == "" {
		return fmt.Errorf("cat: --host is required (Thetis is a separate, remote instance — never assume localhost)")
	}
	port := a.flag("port", "13013")
	timeout := parseDuration(a.flag("timeout", "3s"), 3*time.Second)

	addr := net.JoinHostPort(host, port)
	c, err := cat.Dial(addr, timeout)
	if err != nil {
		return err
	}
	defer c.Close()

	cmd, args := a.pos[0], a.pos[1:]
	switch cmd {
	case "freq":
		return catFreq(c, args)
	case "mode":
		return catMode(c, args)
	case "rit":
		return catToggle("rit", args, c.SetRIT, c.GetRIT)
	case "xit":
		return catToggle("xit", args, c.SetXIT, c.GetXIT)
	case "split":
		return catToggle("split", args, c.SetSplit, c.GetSplit)
	case "agc":
		return catAGC(c, args)
	case "atten":
		return catAtten(c, args)
	case "preamp":
		return catPreamp(c, args)
	case "band":
		return catBand(c, args)
	case "power":
		return catToggle("power", args, c.SetPowerOn, c.GetPowerOn)
	case "status":
		return catStatus(c)
	case "ptt":
		return catPTT(c, args, a)
	default:
		return fmt.Errorf("cat: unknown command %q", cmd)
	}
}

func catFreq(c *cat.Client, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("freq: usage: freq get A|B | freq set A|B <hz>")
	}
	switch args[0] {
	case "get":
		hz, err := c.GetVFOFreqHz(args[1])
		if err != nil {
			return err
		}
		fmt.Printf("VFO %s: %d Hz\n", strings.ToUpper(args[1]), hz)
	case "set":
		if len(args) != 3 {
			return fmt.Errorf("freq set: usage: freq set A|B <hz>")
		}
		hz, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			return fmt.Errorf("freq set: invalid Hz value %q: %w", args[2], err)
		}
		if err := c.SetVFOFreqHz(args[1], hz); err != nil {
			return err
		}
		got, err := c.GetVFOFreqHz(args[1])
		if err != nil {
			return err
		}
		fmt.Printf("VFO %s set to %d Hz (confirmed: %d Hz)\n", strings.ToUpper(args[1]), hz, got)
	default:
		return fmt.Errorf("freq: unknown subcommand %q", args[0])
	}
	return nil
}

func catMode(c *cat.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("mode: usage: mode get | mode set <name>")
	}
	switch args[0] {
	case "get":
		mode, err := c.GetMode()
		if err != nil {
			return err
		}
		fmt.Println("Mode:", mode)
	case "set":
		if len(args) != 2 {
			return fmt.Errorf("mode set: usage: mode set <USB|LSB|CW|CWL|FM|AM|DIGU|DIGL>")
		}
		if err := c.SetMode(args[1]); err != nil {
			return err
		}
		got, err := c.GetMode()
		if err != nil {
			return err
		}
		fmt.Println("Mode set to (confirmed):", got)
	default:
		return fmt.Errorf("mode: unknown subcommand %q", args[0])
	}
	return nil
}

// catToggle implements the common "<name> on|off|get" shape shared by
// rit/xit/split, always printing the confirmed state read back from Thetis.
func catToggle(name string, args []string, set func(bool) error, get func() (bool, error)) error {
	if len(args) != 1 {
		return fmt.Errorf("%s: usage: %s on|off|get", name, name)
	}
	switch args[0] {
	case "on":
		if err := set(true); err != nil {
			return err
		}
	case "off":
		if err := set(false); err != nil {
			return err
		}
	case "get":
		// read-only, fall through to the status print below
	default:
		return fmt.Errorf("%s: unknown value %q (want on|off|get)", name, args[0])
	}
	v, err := get()
	if err != nil {
		return err
	}
	fmt.Printf("%s: %v\n", name, v)
	return nil
}

func catAGC(c *cat.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("agc: usage: agc get | agc set <FIXED|LONG|SLOW|MEDIUM|FAST|CUSTOM>")
	}
	switch args[0] {
	case "get":
		mode, err := c.GetAGC()
		if err != nil {
			return err
		}
		fmt.Println("AGC:", mode)
	case "set":
		if len(args) != 2 {
			return fmt.Errorf("agc set: usage: agc set <FIXED|LONG|SLOW|MEDIUM|FAST|CUSTOM>")
		}
		if err := c.SetAGC(args[1]); err != nil {
			return err
		}
		got, err := c.GetAGC()
		if err != nil {
			return err
		}
		fmt.Println("AGC set to (confirmed):", got)
	default:
		return fmt.Errorf("agc: unknown subcommand %q", args[0])
	}
	return nil
}

func catAtten(c *cat.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("atten: usage: atten get | atten set <0-31>")
	}
	switch args[0] {
	case "get":
		db, err := c.GetAttenuatorDB()
		if err != nil {
			return err
		}
		fmt.Printf("Attenuator: %d dB\n", db)
	case "set":
		if len(args) != 2 {
			return fmt.Errorf("atten set: usage: atten set <0-31>")
		}
		db, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("atten set: invalid dB value %q: %w", args[1], err)
		}
		if err := c.SetAttenuatorDB(db); err != nil {
			return err
		}
		got, err := c.GetAttenuatorDB()
		if err != nil {
			return err
		}
		fmt.Printf("Attenuator set to (confirmed): %d dB\n", got)
	default:
		return fmt.Errorf("atten: unknown subcommand %q", args[0])
	}
	return nil
}

func catPreamp(c *cat.Client, args []string) error {
	if len(args) != 2 || args[0] != "set" {
		return fmt.Errorf("preamp: usage: preamp set <0-9> (0=off 1=on 2..6=-10..-50dB 7..9=SA -10..-30dB)")
	}
	level, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("preamp set: invalid level %q: %w", args[1], err)
	}
	if err := c.SetPreamp(level); err != nil {
		return err
	}
	fmt.Println("Preamp level set to:", level)
	return nil
}

func catBand(c *cat.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("band: usage: band get | band set <name>")
	}
	switch args[0] {
	case "get":
		band, err := c.GetBand()
		if err != nil {
			return err
		}
		fmt.Println("Band:", band)
	case "set":
		if len(args) != 2 {
			return fmt.Errorf("band set: usage: band set <160|80|60|40|30|20|17|15|12|10|6|2|GEN|WWV|V0-V13>")
		}
		if err := c.SetBand(args[1]); err != nil {
			return err
		}
		got, err := c.GetBand()
		if err != nil {
			return err
		}
		fmt.Println("Band set to (confirmed):", got)
	default:
		return fmt.Errorf("band: unknown subcommand %q", args[0])
	}
	return nil
}

func catStatus(c *cat.Client) error {
	id, err := c.GetID()
	if err != nil {
		return err
	}
	st, err := c.GetIF()
	if err != nil {
		return err
	}
	fmt.Printf("Rig ID:   %s\n", id)
	fmt.Printf("Freq:     %d Hz\n", st.FreqHz)
	fmt.Printf("Mode:     %s\n", st.Mode)
	fmt.Printf("RIT/XIT:  RIT=%v XIT=%v offset=%+d Hz\n", st.RIT, st.XIT, st.RITXITHz)
	fmt.Printf("Split:    %v\n", st.Split)
	fmt.Printf("TX:       %v\n", st.TXActive)
	return nil
}

// catPTT is the sole TX-capable CAT command: it gates real keying behind the
// safety confirmation phrase and always auto-unkeys after --hold.
func catPTT(c *cat.Client, args []string, a parsedArgs) error {
	if len(args) != 1 {
		return fmt.Errorf("ptt: usage: ptt on --confirm-tx=<phrase> [--hold 3s] | ptt off")
	}
	switch args[0] {
	case "off":
		return c.SetPTT(false)
	case "on":
		hold := parseDuration(a.flag("hold", "3s"), 3*time.Second)
		dec := safety.Check(a.flag("confirm-tx", ""))
		if dec.DryRun {
			fmt.Printf("[dry-run] would send: TX; ... (hold %s) ... RX;\n", hold)
			fmt.Println("Pass --confirm-tx=" + safety.ConfirmPhrase + " to actually key the transmitter.")
			return nil
		}
		if err := c.SetPTT(true); err != nil {
			return err
		}
		fmt.Printf("PTT ON — auto-unkeying after %s\n", hold)
		time.Sleep(hold)
		if err := c.SetPTT(false); err != nil {
			return fmt.Errorf("PTT ON succeeded but auto-unkey failed, radio may still be keyed: %w", err)
		}
		fmt.Println("PTT OFF")
		return nil
	default:
		return fmt.Errorf("ptt: unknown value %q (want on|off)", args[0])
	}
}

func parseDuration(s string, def time.Duration) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return def
	}
	return d
}
