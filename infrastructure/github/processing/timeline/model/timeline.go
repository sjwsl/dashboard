package model

import (
	"time"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/util"
)

type Timeline struct {
	Times []time.Time
}

func GetTimelineFromCreateAt(createdAt time.Time) Timeline {
	createTime := util.ParseDate(createdAt)
	duration := time.Now().Sub(createTime)
	hours := duration.Hours()
	dayNum := int(hours/24) + 1
	Timelines := make([]time.Time, dayNum)

	for tempTime, i := createTime, 0; i < int(hours/24)+1; i++ {
		tempTime = tempTime.AddDate(0, 0, 1)
		Timelines[i] = tempTime
	}
	return Timeline{Times: Timelines}
}
