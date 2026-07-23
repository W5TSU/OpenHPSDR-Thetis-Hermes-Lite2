package safety

import "testing"

func TestCheckWrongPhraseIsDryRun(t *testing.T) {
	dec, err := Check("nope", false, nil)
	if err != nil {
		t.Fatalf("Check: %v", err)
	}
	if !dec.DryRun || dec.Proceed {
		t.Errorf("Check(wrong phrase) = %+v, want DryRun=true Proceed=false", dec)
	}
}

func TestCheckCorrectPhraseNonTTYProceeds(t *testing.T) {
	dec, err := Check(ConfirmPhrase, false, nil)
	if err != nil {
		t.Fatalf("Check: %v", err)
	}
	if dec.DryRun || !dec.Proceed {
		t.Errorf("Check(correct phrase, non-TTY) = %+v, want DryRun=false Proceed=true", dec)
	}
}

func TestCheckCorrectPhraseTTYRequiresPromptYes(t *testing.T) {
	dec, err := Check(ConfirmPhrase, true, func(string) (bool, error) { return true, nil })
	if err != nil {
		t.Fatalf("Check: %v", err)
	}
	if dec.DryRun || !dec.Proceed {
		t.Errorf("Check(correct phrase, TTY, yes) = %+v, want DryRun=false Proceed=true", dec)
	}
}

func TestCheckCorrectPhraseTTYPromptNoFallsBackToDryRun(t *testing.T) {
	dec, err := Check(ConfirmPhrase, true, func(string) (bool, error) { return false, nil })
	if err != nil {
		t.Fatalf("Check: %v", err)
	}
	if !dec.DryRun || dec.Proceed {
		t.Errorf("Check(correct phrase, TTY, no) = %+v, want DryRun=true Proceed=false", dec)
	}
}
