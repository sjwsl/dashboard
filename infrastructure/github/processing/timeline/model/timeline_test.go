package model

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimelineFromCreateAt(t *testing.T) {
	createAt := time.Now().AddDate(0, -2, -3)
	fmt.Println(GetTimelineFromCreateAt(createAt))
}
