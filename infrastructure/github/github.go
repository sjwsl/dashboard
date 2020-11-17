package github

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shurcooL/githubv4"

	"dashboard/infrastructure/github/crawler"
	"dashboard/infrastructure/github/crawler/client"
	"dashboard/infrastructure/github/crawler/config"
	"dashboard/infrastructure/github/crawler/model"
)

var db *sql.DB
var err error

// init Set mysql database connection
func init() {
	MYSQLEnvString := os.Getenv("MYSQL_URI")
	db, err = sql.Open("mysql", MYSQLEnvString)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(100)

	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
}

// Fetch fetch all data
func Fetch(owner, reponame string) *model.Query {
	tokenEnvString := os.Getenv("GITHUB_TOKEN")
	tokens := strings.Split(tokenEnvString, ":")

	client.InitClient(config.Config{
		GraphqlPath:   "./infrastructure/github/crawler/graphql/query.graphql",
		ServerUrl:     "https://api.github.com/graphql",
		Authorization: tokens,
	})
	request := client.NewClient()

	opt := crawler.FetchOption{
		Owner:    owner,
		RepoName: reponame,
		IssueFilters: &map[string]interface{}{
			"states": []string{"CLOSED", "OPEN"},
		},
	}

	totalData := crawler.FetchByRepo(request, opt)
	crawler.QueryCompletenessProof(totalData)
	return totalData
}

// insertData insert all the data fetched from database
func insertData(owner, repoName string, since githubv4.DateTime) {
	totalData := Fetch(owner, repoName)

	insertRepositoryData(db, totalData, owner, repoName)
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})

	InsertTags(tx, totalData)

	for _, issue := range totalData.Repository.Issues.Nodes {
		deleteIssueData(tx, issue)
		insertIssueData(tx, totalData, issue)
		insertUserDataAndRelationshipWithIssue(tx, issue)
		insertLabelDataAndRelationshipWithIssue(tx, issue)
		insertCommentData(tx, issue)
		insertCrossReferenceEvent(tx, issue)
	}
	insertAssignedIssueNumTimeLine(tx, totalData)
	InsertCommentVersion(tx)

	err = tx.Commit()
	fmt.Println(err)
}

type repository struct {
	owner string
	name  string
}

var DBList = []repository{{
	owner: "tikv",
	name:  "tikv",
}, {
	owner: "tikv",
	name:  "pd",
}, {
	owner: "pingcap",
	name:  "tidb",
}, {
	owner: "pingcap",
	name:  "dm",
}, {
	owner: "pingcap",
	name:  "ticdc",
}, {
	owner: "pingcap",
	name:  "br",
}, {
	owner: "pingcap",
	name:  "tidb-lightning",
}}

// RunInfrastructure fetch all the data first and then fetch data 10 days before.
func RunInfrastructure() {
	for _, r := range DBList {
		insertData(r.owner, r.name, githubv4.DateTime{})
	}

	fmt.Printf(
		`
###########################################################################################
init db ok %v
###########################################################################################
`, time.Now())
	for {
		time.Sleep(time.Hour)
		for _, r := range DBList {
			insertData(r.owner, r.name, githubv4.DateTime{})
		}
		fmt.Printf(
			`
###########################################################################################
update database %v
###########################################################################################
`, time.Now())
	}
}
