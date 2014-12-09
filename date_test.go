package commands

import "testing"

func TestDate(t *testing.T) {
	forceTimeTo(1415694000)
	const expected = "November 11, 2014"
	if y := Date(); y != expected {
		t.Errorf("Date() = %v, want %v", y, expected)
	}
}
