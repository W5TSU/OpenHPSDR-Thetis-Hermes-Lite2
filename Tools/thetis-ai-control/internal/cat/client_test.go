package cat

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

// newTestClient wires a Client to an in-memory net.Pipe() fake CAT server
// that replies to any request found in replies (request text without the
// trailing ';') and silently drops anything else, mirroring how Thetis
// treats fire-and-forget set commands.
func newTestClient(t *testing.T, replies map[string]string) (*Client, func()) {
	t.Helper()
	clientConn, serverConn := net.Pipe()

	done := make(chan struct{})
	go func() {
		defer close(done)
		r := bufio.NewReader(serverConn)
		for {
			line, err := r.ReadString(';')
			if err != nil {
				return
			}
			req := strings.TrimSuffix(line, ";")
			if reply, ok := replies[req]; ok {
				if _, err := serverConn.Write([]byte(reply + ";")); err != nil {
					return
				}
			}
		}
	}()

	c := newClient(clientConn, 2*time.Second)
	cleanup := func() {
		c.Close()
		serverConn.Close()
		<-done
	}
	return c, cleanup
}

func TestFreqRoundTrip(t *testing.T) {
	c, cleanup := newTestClient(t, map[string]string{"FA": "FA00014074000"})
	defer cleanup()

	if err := c.SetVFOFreqHz("A", 14074000); err != nil {
		t.Fatalf("SetVFOFreqHz: %v", err)
	}
	hz, err := c.GetVFOFreqHz("A")
	if err != nil {
		t.Fatalf("GetVFOFreqHz: %v", err)
	}
	if hz != 14074000 {
		t.Errorf("GetVFOFreqHz = %d, want 14074000", hz)
	}
}

func TestModeRoundTrip(t *testing.T) {
	c, cleanup := newTestClient(t, map[string]string{"MD": "MD2"})
	defer cleanup()

	if err := c.SetMode("USB"); err != nil {
		t.Fatalf("SetMode: %v", err)
	}
	mode, err := c.GetMode()
	if err != nil {
		t.Fatalf("GetMode: %v", err)
	}
	if mode != "USB" {
		t.Errorf("GetMode = %q, want USB", mode)
	}
}

func TestSplitRoundTrip(t *testing.T) {
	c, cleanup := newTestClient(t, map[string]string{"ZZSP": "ZZSP1"})
	defer cleanup()

	if err := c.SetSplit(true); err != nil {
		t.Fatalf("SetSplit: %v", err)
	}
	on, err := c.GetSplit()
	if err != nil {
		t.Fatalf("GetSplit: %v", err)
	}
	if !on {
		t.Errorf("GetSplit = false, want true")
	}
}

func TestQueryUnexpectedReplyIsError(t *testing.T) {
	c, cleanup := newTestClient(t, map[string]string{"ZZ": "XX1"})
	defer cleanup()

	if _, err := c.Query("ZZ"); err == nil {
		t.Fatal("Query with mismatched reply prefix: want error, got nil")
	}
}

func TestSetVFOFreqHzOutOfRange(t *testing.T) {
	c, cleanup := newTestClient(t, nil)
	defer cleanup()

	if err := c.SetVFOFreqHz("A", 999999999999); err == nil {
		t.Fatal("SetVFOFreqHz with 12-digit Hz: want error, got nil")
	}
}

func TestParseIF(t *testing.T) {
	// Field layout per CATCommands.cs IF() (lines 378-401):
	// freq(11) step(4) incr(6) rit(1) xit(1) dummy(3) tx(1) mode(1) dummy(2) split(1) dummy(4)
	s := "00014074000" + "0000" + "+00100" + "1" + "0" + "000" + "1" + "2" + "00" + "1" + "0000"
	if len(s) != 35 {
		t.Fatalf("test fixture length = %d, want 35", len(s))
	}

	st, err := parseIF(s)
	if err != nil {
		t.Fatalf("parseIF: %v", err)
	}
	if st.FreqHz != 14074000 {
		t.Errorf("FreqHz = %d, want 14074000", st.FreqHz)
	}
	if st.RITXITHz != 100 {
		t.Errorf("RITXITHz = %d, want 100", st.RITXITHz)
	}
	if !st.RIT {
		t.Errorf("RIT = false, want true")
	}
	if st.XIT {
		t.Errorf("XIT = true, want false")
	}
	if !st.TXActive {
		t.Errorf("TXActive = false, want true")
	}
	if st.Mode != "USB" {
		t.Errorf("Mode = %q, want USB", st.Mode)
	}
	if !st.Split {
		t.Errorf("Split = false, want true")
	}
}

func TestPTTWireCommands(t *testing.T) {
	var got []string
	clientConn, serverConn := net.Pipe()
	done := make(chan struct{})
	go func() {
		defer close(done)
		r := bufio.NewReader(serverConn)
		for {
			line, err := r.ReadString(';')
			if err != nil {
				return
			}
			got = append(got, strings.TrimSuffix(line, ";"))
		}
	}()
	c := newClient(clientConn, 2*time.Second)

	if err := c.SetPTT(true); err != nil {
		t.Fatalf("SetPTT(true): %v", err)
	}
	if err := c.SetPTT(false); err != nil {
		t.Fatalf("SetPTT(false): %v", err)
	}
	c.Close()
	serverConn.Close()
	<-done

	want := []string{"TX", "RX"}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("wire commands = %v, want %v", got, want)
	}
}
