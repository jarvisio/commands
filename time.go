package commands

import (
	"time"
)

func Now(use24 bool) string {
	layout := "3:04PM"
	if use24 {
		layout = "15:04"
	}
	t := TimeNow()
	return t.Format(layout)
}

var TimeNow = func() time.Time {
	return time.Now()
}

func NowForce(unix int) {
	TimeNow = func() time.Time {
		t := time.Unix(int64(unix), 0)
		return t.UTC()
	}
}
