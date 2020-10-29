package github

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	dateTime := time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
	if ParseDate(dateTime) != time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("ParseDate give %v, but expect %v , with arg (%d,%d,%d,%d,%d,%d,%d,%d,%v)",
			ParseDate(dateTime), time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			1, 1, 1, 1, 1, 1, 1, 1, time.UTC)
	}
}
