package commands

import (
	"time"
)

func Now() string {
	const layout = "03:04PM"
	t := TimeNow()
	return t.Format(layout)
}

var TimeNow = func() time.Time {
	return time.Now()
}

func NowForce(unix int) {
	TimeNow = func() time.Time {
		return time.Unix(int64(unix), 0)
	}
}
