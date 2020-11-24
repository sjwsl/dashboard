package github

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"dashboard/infrastructure/github/crawler/model"
)

// insertRepositoryData insert Data into Repository table REPOSITORY
func insertRepositoryData(db *sql.DB, totalData *model.Query, owner string, repoName string) {
	_, err := db.Exec(`INSERT INTO  REPOSITORY (ID,OWNER, REPO_NAME) VALUES (?,?,?)`,
		totalData.Repository.DatabaseID, owner, repoName)
	if err != nil && !strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("Insert fail while REPLACE INTO  REPOSITORY (ID,OWNER, REPO_NAME) VALUES", err)
	}
}

// insertIssueData insert data into table ISSUE
func insertIssueData(db *sql.Tx, totalData *model.Query, issue *model.Issue) {
	closeAt := sql.NullTime{}
	if issue.Closed {
		closeAt = sql.NullTime{
			Time:  *issue.ClosedAt,
			Valid: true,
		}
	}

	_, err := db.Exec(
		`INSERT INTO ISSUE 
    	(ID,NUMBER, REPOSITORY_ID, CLOSED, CLOSED_AT, CREATED_AT, TITLE) 
    	VALUES (?,?,?,?,?,?,?);`,
		issue.DatabaseID,
		issue.Number, totalData.Repository.DatabaseID, issue.Closed,
		closeAt, issue.CreatedAt, issue.Title)
	if err != nil && !strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("Insert fail while REPLACE INTO ISSUE ", err)
	}
}

//insertLabelDataAndRelationshipWithIssue insert Data into Label And LABEL_ISSUE_RELATIONSHIP
func insertLabelDataAndRelationshipWithIssue(db *sql.Tx, issue *model.Issue) {
	for _, node := range issue.Labels.Nodes {
		_, err := db.Exec(
			`INSERT INTO LABEL (NAME) VALUES (?);`,
			node.Name)
		if err != nil && !strings.Contains(err.Error(), "Duplicate") {
			fmt.Println("INSERT INTO LABEL", err)
		}

		_, err = db.Exec(
			`INSERT INTO LABEL_ISSUE_RELATIONSHIP (LABEL_ID, ISSUE_ID)
				SELECT LABEL.ID,?
				FROM LABEL where LABEL.NAME = ?;`,
			issue.DatabaseID, node.Name)
		if err != nil && !strings.Contains(err.Error(), "Duplicate") {
			fmt.Println("INSERT INTO LABEL_ISSUE_RELATIONSHIP ", err)
		}
	}

}

// insertUserDataAndRelationshipWithIssue INSERT Data INTO USER table and INSERT INTO ASSIGNEE table
func insertUserDataAndRelationshipWithIssue(db *sql.Tx, issue *model.Issue) {
	for _, node := range issue.Assignees.Nodes {
		_, err := db.Exec(
			`INSERT INTO USER (LOGIN_NAME, EMAIL)VALUES (?,?);`,
			node.Login, node.Email)
		if err != nil && !strings.Contains(err.Error(), "Duplicate") {
			fmt.Println("INSERT INTO USER ", err)
		}

		_, err = db.Exec(
			`INSERT INTO ASSIGNEE (USER_ID, ISSUE_ID)
				SELECT USER.ID,?
				from USER where USER.LOGIN_NAME = ?;`,
			issue.DatabaseID, node.Login)
		if err != nil && !strings.Contains(err.Error(), "Duplicate") {
			fmt.Println("INSERT INTO ASSIGNEE ", err)
		}
	}
}

// insertCommentData INSERT Data INTO COMMENT table
// the comment data compose issue body and comments.
func insertCommentData(db *sql.Tx, issue *model.Issue) {
	stmt, err := db.Prepare(`INSERT INTO COMMENT (ISSUE_ID, BODY) VALUES (?,?)`)
	if err != nil && !strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("INSERT INTO COMMENT ", err)
		return
	}
	_, err = stmt.Exec(issue.DatabaseID, issue.Body)
	if err != nil && !strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("INSERT INTO COMMENT ", err)
	}
	for _, comment := range issue.Comments.Nodes {
		_, err := stmt.Exec(issue.DatabaseID, comment.Body)
		if err != nil && !strings.Contains(err.Error(), "Duplicate") {
			fmt.Println("INSERT INTO COMMENT ", err)
		}
	}
}

// insertCrossReferenceEvent  INSERT Data INTO Cross_Referenced_Event
func insertCrossReferenceEvent(db *sql.Tx, issue *model.Issue) {
	for _, Node := range issue.TimelineItems.Nodes {
		if *Node.Typename == "CrossReferenceEvent" {
			_, err := db.Exec(`INSERT INTO Cross_Referenced_Event (USER_ID,CREATE_AT,ISSUE_ID)
				SELECT USER.ID,?,?
				from USER where USER.LOGIN_NAME = ?;`,
				*Node.CreatedAt,
				issue.DatabaseID,
				*Node.Actor.Login)
			if err != nil && !strings.Contains(err.Error(), "Duplicate") {
				fmt.Println("INSERT INTO COMMENT ", err)
			}
		}
	}
}

// insertAssignedIssueNumTimeLine calculate the sum of assigned issue every day
//	and INSERT INTO ASSIGNED_ISSUE_NUM_TIMELINE table
func insertAssignedIssueNumTimeLine(db *sql.Tx, totalData *model.Query) {
	repoCreateTime := ParseDate(totalData.Repository.CreatedAt)
	assignedIssueNumTimeLine := time.Now().Sub(repoCreateTime)
	hours := assignedIssueNumTimeLine.Hours()
	assignedIssueNums := make([]int, int(hours/24)+1)

	for tempTime, i := repoCreateTime, 0; i < int(hours/24)+1; i++ {
		for _, issue := range totalData.Repository.Issues.Nodes {
			if issueAssignedBeforeDateTime(tempTime, issue) {
				assignedIssueNums[i]++
			}
		}

		_, err := db.Exec(`INSERT INTO ASSIGNED_ISSUE_NUM_TIMELINE (DATETIME,REPO_ID,ASSIGNED_ISSUE_NUM) VALUES (?,?,?)`,
			tempTime, totalData.Repository.DatabaseID, assignedIssueNums[i])
		if err != nil && !strings.Contains(err.Error(), "Duplicate") {
			fmt.Println("INSERT INTO ASSIGNED_ISSUE_NUM_TIMELINE ", err)
		}
		tempTime = tempTime.AddDate(0, 0, 1)
	}
}

// issueAssignedBeforeDateTime find if the issue assigned before giving datetime
func issueAssignedBeforeDateTime(dateTime time.Time, issue *model.Issue) bool {

	assigneeMap := make(map[string]bool)
	if issue.CreatedAt.Before(dateTime) {
		for _, node := range issue.TimelineItems.Nodes {
			switch *node.Typename {
			case "AssignedEvent":
				if node.CreatedAt.Before(dateTime) {
					assigneeMap[*node.Assignee.Login] = true
				}
			case "UnassignedEvent":
				if node.CreatedAt.Before(dateTime) {
					assigneeMap[*node.Assignee.Login] = false
				}
			}
		}
	}
	for _, Assigned := range assigneeMap {
		if Assigned {
			return true
		}
	}
	return false
}

// ParseDate return time with date and hour;min;sec;nsec is 0;0;0;0 in UTC
func ParseDate(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func InsertTags(tx *sql.Tx, totalData *model.Query) {
	for _, tag := range totalData.Repository.Refs.Nodes {
		_, err := tx.Exec(`INSERT INTO REPO_VERSION (TAG, REPO_ID) VALUES (?,?)`,
			tag.Name, totalData.Repository.DatabaseID)
		if err != nil && !strings.Contains(err.Error(), "Duplicate") {
			fmt.Println("INSERT INTO REPO_VERSION ", err)
		}
	}
}

func InsertCommentVersion(tx *sql.Tx) {
	_, err := tx.Exec(`INSERT INTO COMMENT_VERSION (COMMENT_ID,VERSIONS) SELECT COMMENT.ID AS C_ID, parse_affected_version(SUBSTRING(BODY,REGEXP_INSTR(BODY,"#### 5. Affected versions"),REGEXP_INSTR(BODY,"#### 6. Fixed versions") - REGEXP_INSTR(BODY,"#### 5. Affected versions"))) VERSIONS
FROM COMMENT
WHERE COMMENT.BODY REGEXP "#### 5. Affected versions" AND
      LENGTH(parse_affected_version(SUBSTRING(BODY,REGEXP_INSTR(BODY,"#### 5. Affected versions"),REGEXP_INSTR(BODY,"#### 6. Fixed versions") - REGEXP_INSTR(BODY,"#### 5. Affected versions")))) > 0;
`)
	if err != nil && !strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("INSERT INTO COMMENT_VERSION ", err)
	}
}
