package crawler

import (
	"context"
	"dashboard/infrastructure/github/client"
	"dashboard/infrastructure/github/model"
	"github.com/pkg/math"
)

const maxGithubPageSize = 100

type FetchOption struct {
	owner        string
	repoName     string
	first        *int
	issueFilters *map[string]interface{}
}

func FetchByRepo(request client.Request, opt FetchOption) *model.Query {

	v := map[string]interface{}{
		"owner":           opt.owner,
		"repo_name":       opt.repoName,
		"issueFilters":    opt.issueFilters,
		"issueAfter":      nil,
		"commentAfter":    nil,
		"tagAfter":        nil,
		"IssuePageSize":   0,
		"CommentPageSize": 0,
		"tagPageSize":     0,
	}
	totalCountData, err := pingCountByRepo(request, v)
	if err != nil {
		panic(err)
	}

	var totalData model.Query

	v["IssuePageSize"] = maxGithubPageSize
	v["CommentPageSize"] = maxGithubPageSize
	totalCount := totalCountData.Repository.Issues.TotalCount
	if opt.first != nil {
		totalCount = math.Min(totalCount, *opt.first)
	}
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		v["IssuePageSize"] = math.Min(totalCount-count, maxGithubPageSize)
		var respData model.Query
		err := request.QueryWithAuthPool(context.Background(), &respData, v)
		if err != nil {
			panic(err)
		}
		totalData.Repository.Issues.Nodes = append(totalData.Repository.Issues.Nodes, respData.Repository.Issues.Nodes...)
		if !respData.Repository.Issues.PageInfo.HasNextPage {
			break
		}
		v["issueAfter"] = respData.Repository.Issues.PageInfo.EndCursor
	}

	v["tagPageSize"] = maxGithubPageSize
	v["IssuePageSize"] = 0
	v["CommentPageSize"] = 0
	totalCount = totalCountData.Repository.Refs.TotalCount
	if opt.first != nil {
		totalCount = math.Min(totalCount, *opt.first)
	}
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		var respData model.Query
		err := request.QueryWithAuthPool(context.Background(), &respData, v)
		if err != nil {
			panic(err)
		}
		totalData.Repository.Refs.Nodes = append(totalData.Repository.Refs.Nodes, respData.Repository.Refs.Nodes...)
		if !respData.Repository.Refs.PageInfo.HasNextPage {
			break
		}
		v["tagAfter"] = respData.Repository.Refs.PageInfo.EndCursor
	}

	return &totalData
}

func pingCountByRepo(request client.Request, variable map[string]interface{}) (model.Query, error) {
	var data model.Query
	err := request.QueryWithAuthPool(context.Background(), &data, variable)
	if err != nil {
		panic(err)
	}
	return data, nil
}
