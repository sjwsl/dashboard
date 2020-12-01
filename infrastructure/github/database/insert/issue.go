package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/util"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions"
	model2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/model"
)

// insert.Issue insert data into table ISSUE
func Issue(db *sql.Tx, repo *model.Repository, issue *model.Issue) {
	closeAt := util.GetIssueClosedTime(issue.Closed, issue.ClosedAt)
	_, err := db.Exec(`
insert into issue 
(id,number, repository_id, closed, closed_at, created_at, title, url) 
values (?,?,?,?,?,?,?,?);`,
		issue.DatabaseID, issue.Number, repo.DatabaseID, issue.Closed,
		closeAt, issue.CreatedAt, issue.Title, issue.Url)
	if err != nil {
		fmt.Println("Insert fail while insert into issue (id,number, repository_id, closed, closed_at, created_at, title, url) ", err)
	}
}

func IssueLabel(db *sql.Tx, repo *model.Repository, issue *model.Issue, label *model.Label) {
	_, err := db.Exec(`
insert into issue_label (issue_id,label_id)
select ?, label.id
from label where label.name = ? and
                 repository_id = ?;`,
		issue.DatabaseID, label.Name, repo.DatabaseID)
	if err != nil {
		fmt.Println("Insert fail while insert into issue_label (issue_id,label_id)", err)
	}
}

func IssueVersion(db *sql.Tx, issue *model.Issue, body *string) {
	affectedVersions, fixedVersions, err := versions.GetVersions(body)
	if err != nil {
		return
	}

	for _, version := range affectedVersions {
		err := issueVersionAffected(db, issue, &version)
		if err != nil {
			fmt.Println(err)
			fmt.Println(*body)
			fmt.Println(issue.Number)
			fmt.Println("###########################")
		}
	}
	for _, version := range fixedVersions {
		err := issueVersionFixed(db, issue, &version)
		if err != nil {
			fmt.Println(err)
			fmt.Println(*body)
			fmt.Println(issue.Number)
			fmt.Println("###########################")
		}
	}
}

func issueVersionAffected(db *sql.Tx, issue *model.Issue, version *model2.Version) error {
	versionDatabaseID, err := util.GenIDFromVersion(*version)
	if err != nil {
		return nil
	}
	_, err = db.Exec(`
insert into issue_version_affected 
(issue_id, version_id)
values (?,?)`,
		issue.DatabaseID, versionDatabaseID)
	if err != nil {
		return fmt.Errorf(" insert fail while insert into issue_version_affected (issue_id, version_id) , %v", err)
	}
	return nil
}

func issueVersionFixed(db *sql.Tx, issue *model.Issue, version *model2.Version) error {
	versionDatabaseID, err := util.GenIDFromVersion(*version)
	if err != nil {
		return nil
	}
	_, err = db.Exec(`
insert into issue_version_fixed
(issue_id, version_id)
values (?,?);`,
		issue.DatabaseID, versionDatabaseID)
	if err != nil {
		return fmt.Errorf(" insert fail while insert into issue_version_fixed (issue_id, version_id), %v", err)
	}
	return nil
}
