// Package safety gates every code path in thetisctl that can key the
// transmitter. Neither of Thetis's network control protocols (CAT-over-TCP,
// TCI-over-WebSocket) has authentication, and TCI TX audio genuinely
// modulates and transmits RF (see Console/cmaster.cs TCITxThreadProc) — so
// anything TX-capable defaults to a dry run, and real keying requires an
// explicit, hard-to-fat-finger confirmation.
package safety

import (
	"bufio"
	"fmt"
	"strings"
)

// ConfirmPhrase is the exact value --confirm-tx must equal to key the
// transmitter. A bare boolean flag is intentionally rejected so that no
// other tool's "--confirm-tx" convention can accidentally authorize TX here.
const ConfirmPhrase = "I-UNDERSTAND-THIS-KEYS-THE-RADIO"

// Decision is the outcome of a TX confirmation check.
type Decision struct {
	Proceed bool
	DryRun  bool
}

// Check validates a --confirm-tx flag value. If confirmFlag does not equal
// ConfirmPhrase, the caller must run in dry-run mode (Proceed=false,
// DryRun=true). If it matches and isTTY, an interactive y/N prompt is also
// required before Proceed becomes true; non-TTY (scripted) use is allowed to
// proceed on the phrase match alone, since the phrase itself is the explicit,
// deliberate opt-in for automation.
func Check(confirmFlag string, isTTY bool, prompt func(string) (bool, error)) (Decision, error) {
	if confirmFlag != ConfirmPhrase {
		return Decision{Proceed: false, DryRun: true}, nil
	}
	if isTTY {
		ok, err := prompt(fmt.Sprintf(
			"About to key the transmitter. Confirm? (only 'yes' proceeds) "))
		if err != nil {
			return Decision{}, fmt.Errorf("safety: confirmation prompt: %w", err)
		}
		if !ok {
			return Decision{Proceed: false, DryRun: true}, nil
		}
	}
	return Decision{Proceed: true, DryRun: false}, nil
}

// PromptStdin implements the prompt func signature Check expects, reading a
// single line from stdin and treating exactly "yes" (case-insensitive,
// trimmed) as confirmation.
func PromptStdin(r *bufio.Reader, question string) (bool, error) {
	fmt.Print(question)
	line, err := r.ReadString('\n')
	if err != nil {
		return false, err
	}
	return strings.EqualFold(strings.TrimSpace(line), "yes"), nil
}
