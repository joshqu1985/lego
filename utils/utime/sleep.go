package utime

import (
	"time"
)

func Sleep(d time.Duration) {
	if d == 0 {
		return
	}

	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C:
		return
	}
}
