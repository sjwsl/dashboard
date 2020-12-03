package insert

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	model2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/timeline/model"
)

func Timeline(db *sql.DB, config *config.Config) {
	timelines := model2.GetTimelineFromCreateAt(config.CreateAtGlobal)
	var wg sync.WaitGroup
	for _, t := range timelines.Times {
		wg.Add(1)
		go func(time time.Time) {
			_, err := db.Exec(`
insert into timeline (datetime)
values (?) on duplicate key update datetime=?;`, time, time)
			if err != nil {
				fmt.Println("Insert fail while insert into timeline (datetime)", err)
			}
			defer wg.Done()
		}(t)
	}
	wg.Wait()
}

func WeekLine(db *sql.DB, config *config.Config) {
	timelines := model2.GetTimelineFromCreateAt(config.CreateAtGlobal)
	var wg sync.WaitGroup
	for _, t := range timelines.Times {
		wg.Add(1)
		go func(time time.Time) {
			_, err := db.Exec(`
insert ignore into week_line (week)
values (date_add(?, interval 7 - weekday(?) day));`, time, time)
			if err != nil {
				fmt.Println("Insert fail while insert into week_line (week)", err)
			}
			defer wg.Done()
		}(t)
	}
	wg.Wait()
}

func TimelineRepository(db *sql.DB, repository *model.Repository) {
	timelines := model2.GetTimelineFromCreateAt(repository.CreatedAt)
	var wg sync.WaitGroup
	for _, t := range timelines.Times {
		wg.Add(1)
		go func(time time.Time) {
			_, err := db.Exec(`
insert into timeline_repository (datetime,repository_id)
values (?,?);`, time, repository.DatabaseID)
			if err != nil {
				fmt.Println("Insert fail while insert into repository_timeline (datetime,repository_id)", err)
			}
			defer wg.Done()
		}(t)
	}
	wg.Wait()
}
