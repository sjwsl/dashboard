package github

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PingCAP-QE/libs/crawler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-github/v32/github"
	"github.com/shurcooL/githubv4"
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

// initClient Set link with githubV4 client & github client
func initClient() (crawler.ClientV4, *github.Client) {
	tokenEnvString := os.Getenv("GITHUB_TOKEN")
	tokens := strings.Split(tokenEnvString, ":")
	crawler.InitGithubV4Client(tokens)
	clientV4 := crawler.NewGithubV4Client()
	client := crawler.NewGithubClient(tokens[0])
	return clientV4, client
}

// insertData insert all the data fetched from database
func insertData(owner, repoName string, since githubv4.DateTime) {
	clientV4, client := initClient()
	repo, _, err := client.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		log.Printf("err while fetching repo data in %s / %s\n err : %v \n", owner, repoName, err)
		time.Sleep(time.Hour)
		repo, _, err = client.Repositories.Get(context.Background(), owner, repoName)
		if err != nil {
			log.Fatal(err)
		}
	}
	insertRepositoryData(db, repo)
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})

	tags := crawler.ListTags(client, owner, repoName)
	InsertTags(tx, tags, int(*repo.ID))

	issueWithComments, errs := crawler.FetchIssueWithCommentsByLabels(clientV4, owner, repoName, []string{"type/bug"}, since)
	if errs != nil {
		log.Printf("err while fetching issues data in %s / %s\n err : %v \n", owner, repoName, err)
		time.Sleep(time.Hour)
		issueWithComments, errs = crawler.FetchIssueWithCommentsByLabels(clientV4, owner, repoName, []string{"type/bug"}, since)
		if errs != nil {
			log.Fatal(errs[0])
		}
	}

	for _, issueWithComment := range *issueWithComments {
		deleteIssueData(tx, &issueWithComment)
		insertIssueData(tx, repo, &issueWithComment)
		insertUserDataAndRelationshipWithIssue(tx, &issueWithComment)
		insertLabelDataAndRelationshipWithIssue(tx, &issueWithComment)
		insertCommentData(tx, &issueWithComment)
		insertCrossReferenceEvent(tx, &issueWithComment)
	}
	insertAssignedIssueNumTimeLine(tx, repo, issueWithComments)

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
	for true {
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
