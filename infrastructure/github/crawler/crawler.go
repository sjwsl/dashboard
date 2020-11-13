package crawler

import (
	"context"
	"dashboard/infrastructure/github/client"
	"dashboard/infrastructure/github/config"
	"dashboard/infrastructure/github/model"
)

func FetchIssuesByRepo(
	config config.Config,
	owner, repoName string,
	first int,
	issueFilters, issueOrder map[string]interface{}) *model.Query {

	client.InitClient(config)
	request := client.NewClient()

	v := map[string]interface{}{
		"owner":           owner,
		"repo_name":       repoName,
		"issueFilters":    issueFilters,
		"issueOrder":      issueOrder,
		"issueAfter":      nil,
		"commentAfter":    nil,
		"IssuePageSize":   100,
		"CommentPageSize": 100,
	}

	var totalData model.Query

	for {
		var respData model.Query
		err := request.QueryWithAuthPool(context.Background(), &respData, v)
		if err != nil {
			panic(err)
		}
		totalData.Repository.Issues.Nodes = append(totalData.Repository.Issues.Nodes, respData.Repository.Issues.Nodes...)
		if !respData.Repository.Issues.PageInfo.HasNextPage {
			return &totalData
		}
		v["issueAfter"] = respData.Repository.Issues.PageInfo.EndCursor
	}

	return nil
}
