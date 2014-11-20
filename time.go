package commands

import (
	"time"
)

// Now returns the current time as a string
// you can choose between 12 and 24 hour time format
// using the use24 parameter.
func Now(use24 bool) string {
	layout := "3:04PM"
	if use24 {
		layout = "15:04"
	}
	t := timeNow()
	return t.Format(layout)
}

var timeNow = func() time.Time {
	return time.Now()
}

func nowForce(unix int) {
	timeNow = func() time.Time {
		t := time.Unix(int64(unix), 0)
		return t.UTC()
	}
}
