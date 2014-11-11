package commands

import "testing"

func TestNow12HoursAM(t *testing.T) {
	NowForce(1415694000)
	const outAM = "9:20AM"
	if y := Now(false); y != outAM {
		t.Errorf("Now() = %v, want %v", y, outAM)
	}
}

func TestNow12HoursPM(t *testing.T) {
	NowForce(1415737200)
	const outPM = "9:20PM"
	if x := Now(false); x != outPM {
		t.Errorf("Now() = %v, want %v", x, outPM)
	}
}

func TestNow24HoursAM(t *testing.T) {
	NowForce(1415694000)
	const outAM = "09:20"
	if y := Now(true); y != outAM {
		t.Errorf("Now() = %v, want %v", y, outAM)
	}
}

func TestNow24HoursPM(t *testing.T) {
	NowForce(1415737200)
	const outPM = "21:20"
	if x := Now(true); x != outPM {
		t.Errorf("Now() = %v, want %v", x, outPM)
	}
}
