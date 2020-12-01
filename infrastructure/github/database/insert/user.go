package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"strings"
)

func User(db *sql.Tx, user *model.User) {
	_, err := db.Exec(`
insert into user (id,login,email) values (?,?,?);`,
		user.DatabaseID, user.Login, user.Email)
	if err != nil {
		fmt.Println("Insert fail while insert into insert into user (id,login,email) ", err)
	}
}

func UserIssue(db *sql.Tx, issue *model.Issue, user *model.User) {
	_, err := db.Exec(`
INSERT INTO user_issue (USER_ID, ISSUE_ID)
SELECT USER.ID,?
from USER where USER.LOGIN_NAME = ?;`,
		issue.DatabaseID, user.Login)
	if err != nil && !strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("INSERT INTO ASSIGNEE ", err)
	}
}
