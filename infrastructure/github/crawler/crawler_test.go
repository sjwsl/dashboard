package crawler

import (
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/util"
	"os"
	"strings"
	"testing"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/client"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/config"
)

func TestFetchIssuesByRepo(t *testing.T) {
	tokenEnvString := os.Getenv("GITHUB_TOKEN")
	tokenEnvString = "16e7d7c387fb1a53aa36010dcb466ddcd5b521ff"
	tokens := strings.Split(tokenEnvString, ":")

	client.InitClient(config.Config{
		GraphqlPath:   "./graphql/query.graphql",
		ServerUrl:     "https://api.github.com/graphql",
		Authorization: tokens,
	})
	request := client.NewClient()

	first := 101
	opt := FetchOption{
		Owner:    "pingcap",
		RepoName: "tidb",
		First:    &first,
		IssueFilters: map[string]interface{}{
			"states": []string{"CLOSED", "OPEN"},
			"labels": []string{"type/bug"}},
	}

	totalData := FetchRepo(request, opt)
	util.QueryCompletenessSpec(&totalData)
	util.QueryDataInvalidSpec(&totalData)
}
