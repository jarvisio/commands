package commands

import "testing"

func TestClock12HoursAM(t *testing.T) {
	ForceTimeTo(1415694000)
	const outAM = "8:20AM"
	if y := Clock(false); y != outAM {
		t.Errorf("Clock() = %v, want %v", y, outAM)
	}
}

func TestClock12HoursPM(t *testing.T) {
	ForceTimeTo(1415737200)
	const outPM = "8:20PM"
	if x := Clock(false); x != outPM {
		t.Errorf("Clock() = %v, want %v", x, outPM)
	}
}

func TestClock24HoursAM(t *testing.T) {
	ForceTimeTo(1415694000)
	const outAM = "08:20"
	if y := Clock(true); y != outAM {
		t.Errorf("Clock() = %v, want %v", y, outAM)
	}
}

func TestClock24HoursPM(t *testing.T) {
	ForceTimeTo(1415737200)
	const outPM = "20:20"
	if x := Clock(true); x != outPM {
		t.Errorf("Clock() = %v, want %v", x, outPM)
	}
}
