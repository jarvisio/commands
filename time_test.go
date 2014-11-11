package commands

import "testing"

func TestNow(t *testing.T) {
	const out = "01:20AM"
	NowForce(1234)
	if x := Now(); x != out {
		t.Errorf("Now() = %v, want %v", x, out)
	}
}
