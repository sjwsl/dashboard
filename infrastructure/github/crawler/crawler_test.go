package crawler

import (
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/util"
	"testing"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/client"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/config"
)

func TestFetchIssuesByRepo(t *testing.T) {
	client.InitClient(config.Config{
		GraphqlPath:   "./graphql/query.graphql",
		ServerUrl:     "https://api.github.com/graphql",
		Authorization: []string{""},
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
