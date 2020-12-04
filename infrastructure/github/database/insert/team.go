package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
)

func Team(db *sql.DB, c *config.Config) {
	for _, team := range c.TeamArgs {
		_, err := db.Exec("INSERT INTO team VALUES(0, ?, ?);", team.Name, team.Size)
		if err != nil {
			fmt.Printf("INSERT INTO team VALUES(%v, %v) error: %v\n", team.Name, team.Size, err)
		}
	}
}
