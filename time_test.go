package commands

import "testing"

func TestNow12HoursAM(t *testing.T) {
	nowForce(1415694000)
	const outAM = "8:20AM"
	if y := Now(false); y != outAM {
		t.Errorf("Now() = %v, want %v", y, outAM)
	}
}

func TestNow12HoursPM(t *testing.T) {
	nowForce(1415737200)
	const outPM = "8:20PM"
	if x := Now(false); x != outPM {
		t.Errorf("Now() = %v, want %v", x, outPM)
	}
}

func TestNow24HoursAM(t *testing.T) {
	nowForce(1415694000)
	const outAM = "08:20"
	if y := Now(true); y != outAM {
		t.Errorf("Now() = %v, want %v", y, outAM)
	}
}

func TestNow24HoursPM(t *testing.T) {
	nowForce(1415737200)
	const outPM = "20:20"
	if x := Now(true); x != outPM {
		t.Errorf("Now() = %v, want %v", x, outPM)
	}
}
