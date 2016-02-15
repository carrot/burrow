package util

import (
	"time"
)

const (
	ISO8601_LAYOUT = "2006-01-02T15:04:05Z07:00"
)

func TimeFromISO8601(timeString string) (time.Time, error) {
	return time.Parse(ISO8601_LAYOUT, timeString)
}
