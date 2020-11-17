package crawler

import (
	"context"
	"dashboard/infrastructure/github/crawler/client"
	"dashboard/infrastructure/github/crawler/model"
	"dashboard/infrastructure/github/crawler/util"
	"fmt"
	"github.com/google/martian/log"
	"github.com/pkg/math"
)

const maxGithubPageSize = 100

type FetchOption struct {
	Owner        string
	RepoName     string
	First        *int
	IssueFilters *map[string]interface{}
}

func FetchByRepo(request client.Request, opt FetchOption) *model.Query {
	v := map[string]interface{}{
		"owner":           opt.Owner,
		"repo_name":       opt.RepoName,
		"issueFilters":    opt.IssueFilters,
		"issueAfter":      nil,
		"commentAfter":    nil,
		"tagAfter":        nil,
		"IssuePageSize":   0,
		"CommentPageSize": 0,
		"tagPageSize":     0,
	}
	fmt.Printf("Ping count with %v\n", v)
	totalCountData, err := pingCountByRepo(request, v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get data issue count: %d, tag count: %d\n",
		totalCountData.Repository.Issues.TotalCount, totalCountData.Repository.Refs.TotalCount)
	fmt.Printf("rate limit : %v \n", totalCountData.RateLimit)

	var totalData model.Query
	totalData.Repository = totalCountData.Repository

	v["IssuePageSize"] = maxGithubPageSize
	v["CommentPageSize"] = maxGithubPageSize
	totalCount := totalCountData.Repository.Issues.TotalCount
	if opt.First != nil {
		totalCount = math.Min(totalCount, *opt.First)
	}
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		fmt.Printf("<Fetching issue data %d to %d\n", count, count+math.Min(totalCount-count, maxGithubPageSize))
		v["IssuePageSize"] = math.Min(totalCount-count, maxGithubPageSize)
		var respData model.Query
		retryTimes := 10
		for retryTimes != 0 {
			err := request.QueryWithAuthPool(context.Background(), &respData, v)
			if err != nil {
				log.Errorf(err.Error()+" \n query variables: %v \n retry time: %d", v, 10-retryTimes)
			} else {
				break
			}
			retryTimes--
		}
		fmt.Printf("Fetch success.>\n")
		totalData.Repository.Issues.Nodes = append(totalData.Repository.Issues.Nodes, respData.Repository.Issues.Nodes...)
		if !respData.Repository.Issues.PageInfo.HasNextPage {
			break
		}
		v["issueAfter"] = respData.Repository.Issues.PageInfo.EndCursor
	}
	totalData.Repository.Issues.TotalCount = totalCount

	v["tagPageSize"] = maxGithubPageSize
	v["IssuePageSize"] = 0
	v["CommentPageSize"] = 0
	totalCount = totalCountData.Repository.Refs.TotalCount
	if opt.First != nil {
		totalCount = math.Min(totalCount, *opt.First)
	}
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		fmt.Printf("<Fetching tag data %d to %d\n", count, count+math.Min(totalCount-count, maxGithubPageSize))
		v["tagPageSize"] = math.Min(totalCount-count, maxGithubPageSize)
		var respData model.Query
		retryTimes := 10
		for retryTimes != 0 {
			err := request.QueryWithAuthPool(context.Background(), &respData, v)
			if err != nil {
				log.Errorf(err.Error()+" \n query variables: %v \n retry time: %d", v, retryTimes)
			} else {
				break
			}
			retryTimes--
		}
		fmt.Printf("Fetch success.>\n")
		totalData.Repository.Refs.Nodes = append(totalData.Repository.Refs.Nodes, respData.Repository.Refs.Nodes...)
		if !respData.Repository.Refs.PageInfo.HasNextPage {
			break
		}
		v["tagAfter"] = respData.Repository.Refs.PageInfo.EndCursor
	}
	totalData.Repository.Refs.TotalCount = totalCount

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

func QueryCompletenessProof(totalData *model.Query) {
	nums := make([]int, len(totalData.Repository.Issues.Nodes))
	for i, _ := range nums {
		nums[i] = totalData.Repository.Issues.Nodes[i].Number
	}
	err := util.IdCompletenessProof(totalData.Repository.Issues.TotalCount, nums)
	if err != nil {
		panic(err)
	}
	names := make([]string, len(totalData.Repository.Refs.Nodes))
	for i, _ := range names {
		names[i] = *totalData.Repository.Refs.Nodes[i].Name
	}
	err = util.NameCompletenessProof(totalData.Repository.Refs.TotalCount, names)
	if err != nil {
		panic(err)
	}
}
