package tci

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

// TestSpotWireFormat exercises Client.Spot end-to-end over a net.Pipe,
// verifying the exact wire text a real Thetis TCI server would receive —
// in particular, that the argb color is sent as the correct UNSIGNED
// decimal (TCIServer.cs:4375 parses via uint.TryParse, not a signed int).
// No "live" build tag: this needs no network or running Thetis instance,
// so it runs in normal `go test ./...` / CI, matching TestFrameRoundTrip's
// existing net.Pipe pattern in ws_test.go.
func TestSpotWireFormat(t *testing.T) {
	clientEnd, serverEnd := net.Pipe()
	defer clientEnd.Close()
	defer serverEnd.Close()

	c := NewClient(&Conn{tcp: clientEnd, r: bufio.NewReader(clientEnd), timeout: 2 * time.Second})
	s := &Conn{tcp: serverEnd, r: bufio.NewReader(serverEnd), timeout: 2 * time.Second}

	done := make(chan error, 1)
	go func() {
		op, payload, err := s.ReadFrame()
		if err != nil {
			done <- err
			return
		}
		want := "spot:W5TSU,digu,14070000,4292618270,RADEV1 TX;"
		if op != opText || string(payload) != want {
			done <- fmt.Errorf("got opcode %d payload %q, want text %q", op, payload, want)
			return
		}
		done <- nil
	}()

	if err := c.Spot("W5TSU", "digu", 14070000, 4292618270, "RADEV1 TX"); err != nil {
		t.Fatalf("Spot: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

// TestSpotWireFormatIdleColor checks the idle-color decimal value and an
// empty text field separately, so both of the Global Constraints' verified
// argb values are covered by a passing test, not just the TX one.
func TestSpotWireFormatIdleColor(t *testing.T) {
	clientEnd, serverEnd := net.Pipe()
	defer clientEnd.Close()
	defer serverEnd.Close()

	c := NewClient(&Conn{tcp: clientEnd, r: bufio.NewReader(clientEnd), timeout: 2 * time.Second})
	s := &Conn{tcp: serverEnd, r: bufio.NewReader(serverEnd), timeout: 2 * time.Second}

	done := make(chan error, 1)
	go func() {
		op, payload, err := s.ReadFrame()
		if err != nil {
			done <- err
			return
		}
		want := "spot:VK3ABC,digu,7177000,4282153160,700D;"
		if op != opText || string(payload) != want {
			done <- fmt.Errorf("got opcode %d payload %q, want text %q", op, payload, want)
			return
		}
		done <- nil
	}()

	if err := c.Spot("VK3ABC", "digu", 7177000, 4282153160, "700D"); err != nil {
		t.Fatalf("Spot: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

// TestSpotWireFormatInjection is a regression test for the TCI command
// injection finding: a hostile "mode"-shaped text value sourced from a
// public reporter feed (e.g. "RADEV1;trx:0,true") must not be able to
// terminate the spot command early and smuggle a second command — notably
// "trx:0,true", which would key the transmitter (SetTrx's doc comment) and
// bypass internal/safety's TX confirmation gate entirely. We deliberately
// don't pin the sanitizer's exact replacement character, since that's an
// implementation detail — only that the wire payload contains exactly one
// ';' and it's the terminating one.
func TestSpotWireFormatInjection(t *testing.T) {
	clientEnd, serverEnd := net.Pipe()
	defer clientEnd.Close()
	defer serverEnd.Close()

	c := NewClient(&Conn{tcp: clientEnd, r: bufio.NewReader(clientEnd), timeout: 2 * time.Second})
	s := &Conn{tcp: serverEnd, r: bufio.NewReader(serverEnd), timeout: 2 * time.Second}

	done := make(chan error, 1)
	go func() {
		op, payload, err := s.ReadFrame()
		if err != nil {
			done <- err
			return
		}
		if op != opText {
			done <- fmt.Errorf("got opcode %d, want text", op)
			return
		}
		payload2 := string(payload)
		if got := strings.Count(payload2, ";"); got != 1 {
			done <- fmt.Errorf("payload %q contains %d ';' characters, want exactly 1 (injection not neutralized)", payload2, got)
			return
		}
		if !strings.HasSuffix(payload2, ";") {
			done <- fmt.Errorf("payload %q does not end with ';' (injection split the frame)", payload2)
			return
		}
		done <- nil
	}()

	if err := c.Spot("N0CALL", "digu", 14236000, 4292618270, "RADEV1;trx:0,true"); err != nil {
		t.Fatalf("Spot: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}
