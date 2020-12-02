package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/util"
	model2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/model"
)

func Tag(db *sql.DB, repo *model.Repository, tag *model.Ref) {
	_, err := db.Exec(`
insert into tag (name,repository_id) values (?,?);`,
		tag.Name, repo.DatabaseID)
	if err != nil {
		fmt.Println("Insert fail while insert into insert into tag (name,repository_id) ", err)
	}
}

func Version(db *sql.DB, tag *model.Ref) {
	version, err := model2.ParseVersionFromRegularStrMustHaveV(tag.Name)
	if err != nil {
		return
	}
	DatabaseID, err := util.GenIDFromVersion(version)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(`
insert into version (id,major, minor, patch)
VALUES (?,?,?,?) on duplicate key update id=?;;`,
		DatabaseID, version.Major, version.Minor, version.Patch, DatabaseID)
	if err != nil {
		fmt.Println("Insert fail while insert into insert into version (major, minor, patch, tag_id) ", err)
	}
}
