package commands

import (
	"time"
)

var currentTime = func() time.Time {
	return time.Now()
}

func ForceTimeTo(unix int) {
	currentTime = func() time.Time {
		t := time.Unix(int64(unix), 0)
		return t.UTC()
	}
}
