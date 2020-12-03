package model

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimelineFromCreateAt(t *testing.T) {
	createAt, err := time.Parse(time.RFC3339, "2015-09-06T04:01:52Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(GetTimelineFromCreateAt(createAt))
}
