package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	model2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/timeline/model"
)

func Timeline(db *sql.Tx, config *config.Config) {
	timelines := model2.GetTimelineFromCreateAt(config.CreateAtGlobal)
	for _, time := range timelines.Times {
		_, err := db.Exec(`
insert into timeline (datetime)
values (?);`, time)
		if err != nil {
			fmt.Println("Insert fail while insert into timeline (datetime)")
		}
	}
}

func TimelineRepository(db *sql.Tx, repository *model.Repository) {
	timelines := model2.GetTimelineFromCreateAt(repository.CreatedAt)
	for _, time := range timelines.Times {
		_, err := db.Exec(`
insert into timeline_repository (datetime,repository_id)
values (?,?);`, time, repository.DatabaseID)
		if err != nil {
			fmt.Println("Insert fail while insert into repository_timeline (datetime,repository_id)")
		}
	}
}
