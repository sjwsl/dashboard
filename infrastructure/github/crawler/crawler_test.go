package crawler

import (
	"dashboard/infrastructure/github/client"
	"dashboard/infrastructure/github/config"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFetchIssuesByRepo(t *testing.T) {
	tokenEnvString := os.Getenv("GITHUB_TOKEN")
	tokens := strings.Split(tokenEnvString, ":")

	client.InitClient(config.Config{
		GraphqlPath:   "../graphql/query.graphql",
		ServerUrl:     "https://api.github.com/graphql",
		Authorization: tokens,
	})
	request := client.NewClient()

	opt := FetchOption{
		owner:    "pingcap",
		repoName: "tidb",
		first:    nil,
		issueFilters: &map[string]interface{}{
			"states": []string{"CLOSED", "OPEN"},
			"labels": []string{"type/bug"}},
	}

	_ = FetchByRepo(request, opt)
	fmt.Println("ok")
}
