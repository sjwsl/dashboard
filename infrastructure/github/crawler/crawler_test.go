package crawler

import (
	"dashboard/infrastructure/github/crawler/client"
	"dashboard/infrastructure/github/crawler/config"
	"os"
	"strings"
	"testing"
)

func TestFetchIssuesByRepo(t *testing.T) {
	tokenEnvString := os.Getenv("GITHUB_TOKEN")
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
		IssueFilters: &map[string]interface{}{
			"states": []string{"CLOSED", "OPEN"},
			"labels": []string{"type/bug"}},
	}

	totalData := FetchByRepo(request, opt)
	QueryCompletenessProof(totalData)
	QueryDataInvalidProof(totalData)
}
