package commands

import "testing"

func TestNow(t *testing.T) {
	const out = "09:20PM"
	const fullout = "21:20"
	NowForce(1415737200)
	if x := Now(false); x != out {
		t.Errorf("Now() = %v, want %v", x, out)
	}
	if y := Now(true); y != fullout {
		t.Errorf("Now() = %v, want %v", y, fullout)
	}
}
