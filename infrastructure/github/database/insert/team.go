package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
)

func Team(db *sql.DB, c *config.Config) {
	for _, team := range c.TeamArgs {
		_, err := db.Exec("INSERT INTO team VALUES(0, ?, ?);", team.Name, team.Size)
		if err != nil {
			fmt.Printf("INSERT INTO team VALUES(%v, %v) error: %v\n", team.Name, team.Size, err)
		}
	}
}

func TeamIssue(db *sql.DB, issue *model.Issue, team string) {

	_, err := db.Exec(`
		INSERT INTO team_issue(team_id, issue_id)
		SELECT team.id,?
		FROM team WHERE team.name = ?;`,
		issue.DatabaseID, team)
	if err != nil {
		fmt.Printf("INSERT INTO team_issue (%v, %v) error: %v\n", issue.DatabaseID, team, err)
	}
}
