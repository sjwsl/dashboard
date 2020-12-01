package github

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/client"
	crawlerConfig "github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	dbConfig "github.com/PingCAP-QE/dashboard/infrastructure/github/database/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/insert"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/truncate"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// init Set mysql database connection
func initDB(c dbConfig.Config) {
	db, err = sql.Open("mysql", c.DatabaseUrl)
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
func FetchQuery(c crawlerConfig.Config, owner, name string) model.Query {
	client.InitClient(c)
	request := client.NewClient()

	opt := crawler.FetchOption{
		Owner:    owner,
		RepoName: name,
		IssueFilters: map[string]interface{}{
			"states": []string{"CLOSED", "OPEN"},
		},
	}

	totalData := crawler.FetchByRepoSafe(request, opt)
	return totalData
}

// insertData insert all the data fetched from database
func InsertQuery(tx *sql.Tx, totalData model.Query, owner string, c *config.Config) {
	insert.Repository(tx, totalData.Repository, owner)

	insert.Timeline(tx, c)
	insert.TimelineRepository(tx, totalData.Repository)

	for _, user := range totalData.Repository.AssignableUsers.Nodes {
		insert.User(tx, user)
	}

	for _, label := range totalData.Repository.Labels.Nodes {
		insert.Label(tx, totalData.Repository, label)
		insert.LabelSig(tx, label)
	}

	for _, weight := range c.LabelSeverityWeight {
		insert.LabelSeverityWeight(tx, weight)
	}

	for _, issue := range totalData.Repository.Issues.Nodes {
		insert.Issue(tx, totalData.Repository, issue)
		for _, issueComment := range issue.Comments.Nodes {
			insert.Comment(tx, issue, issueComment)
		}
	}

	for _, ref := range totalData.Repository.Refs.Nodes {
		insert.Tag(tx, totalData.Repository, ref)
		insert.Version(tx, ref)
	}

	for _, issue := range totalData.Repository.Issues.Nodes {
		for _, issueComment := range issue.Comments.Nodes {
			insert.Comment(tx, issue, issueComment)
			insert.IssueVersion(tx, issue, &issueComment.Body)
		}
	}

	for _, issue := range totalData.Repository.Issues.Nodes {
		for _, label := range issue.Labels.Nodes {
			insert.IssueLabel(tx, issue, label)
		}
	}

	for _, issue := range totalData.Repository.Issues.Nodes {
		for _, user := range issue.Assignees.Nodes {
			insert.UserIssue(tx, issue, user)
		}
	}

	err = tx.Commit()
	fmt.Println(err)
}

// RunInfrastructure fetch all the data first and then fetch data 10 days before.
func RunInfrastructure(c config.Config) {

	initDB(c.DatabaseConfig)

	queries := make([]model.Query, len(c.RepositoryArgs))
	for i, repositoryArg := range c.RepositoryArgs {
		queries[i] = FetchQuery(c.CrawlerConfig, repositoryArg.Owner, repositoryArg.Name)
	}

	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	if err != nil {
		panic(err)
	}

	truncate.AllClear(tx)

	for i, query := range queries {
		InsertQuery(tx, query, c.RepositoryArgs[i].Owner, &c)
	}

}
