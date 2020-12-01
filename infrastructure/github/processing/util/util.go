package util

import "time"

// ParseDate return time with date and hour;min;sec;nsec is 0;0;0;0 in UTC
func ParseDate(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
