package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
)

// insert.Repository insert Data into Repository table REPOSITORY
func Repository(db *sql.DB, repo *model.Repository, owner string) {
	_, err := db.Exec(`
insert into  repository (id,owner,url,repo_name) values (?,?,?,?);`,
		repo.DatabaseID, owner, repo.Url, repo.Name)
	if err != nil {
		fmt.Println("Insert fail while insert into  repository (id,owner,url,repo_name)", err)
	}
}
