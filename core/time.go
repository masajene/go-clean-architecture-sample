package core

import "time"

func TruncateTime(t time.Time) time.Time {
	return t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour)
}
