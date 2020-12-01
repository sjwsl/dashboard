package crawler

import (
	"context"
	"github.com/pkg/math"
	"log"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/client"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/util"
)

const maxGithubPageSize = 100

type FetchOption struct {
	Owner        string
	RepoName     string
	First        *int
	IssueFilters map[string]interface{}
}

// FetchByRepoSafe Fetch all the data and then check the data.
func FetchByRepoSafe(request client.Request, opt FetchOption) model.Query {
	totalData := FetchRepo(request, opt)
	util.QueryCompletenessSpec(&totalData)
	util.QueryDataInvalidSpec(&totalData)
	return totalData
}

// pingCountByRepo ping the graphql server to get  count infos of issues and tags.
func pingCountByRepo(request client.Request, variable map[string]interface{}) (model.Query, error) {
	var data model.Query
	err := request.QueryWithAuthPool(context.Background(), &data, variable)
	if err != nil {
		panic(err)
	}
	return data, nil
}

// FetchRepo Fetch Repo necessary with FetchOption and request.
func FetchRepo(request client.Request, opt FetchOption) model.Query {
	// v is arguments rely on query in infrastructure/github/crawler/graphql/query.graphql
	v := map[string]interface{}{
		"owner":           opt.Owner,
		"repo_name":       opt.RepoName,
		"issueFilters":    opt.IssueFilters,
		"userAfter":       nil,
		"issueAfter":      nil,
		"commentAfter":    nil,
		"labelAfter":      nil,
		"tagAfter":        nil,
		"userPageSize":    0,
		"IssuePageSize":   0,
		"CommentPageSize": 0,
		"labelPageSize":   0,
		"tagPageSize":     0,
	}

	log.Printf("Ping count with %v\n", v)
	totalCountData, err := pingCountByRepo(request, v)
	if err != nil {
		panic(err)
	}
	log.Printf("rate limit : %v \n", totalCountData.RateLimit)

	var totalData model.Query
	totalData.Repository = totalCountData.Repository

	v["IssuePageSize"] = maxGithubPageSize
	v["CommentPageSize"] = maxGithubPageSize
	totalCount := totalCountData.Repository.Issues.TotalCount
	if opt.First != nil {
		totalCount = math.Min(totalCount, *opt.First)
	}
	log.Printf("Get data issue count: %d \n", totalCount)
	FetchIssueByRepo(request, totalCount, v, &totalData)

	v["tagPageSize"] = maxGithubPageSize
	v["IssuePageSize"] = 0
	v["CommentPageSize"] = 0
	totalCount = totalCountData.Repository.Refs.TotalCount
	if opt.First != nil {
		totalCount = math.Min(totalCount, *opt.First)
	}
	log.Printf("Get data Tag count: %d \n", totalCount)
	FetchTagByRepo(request, totalCount, v, &totalData)

	v["labelPageSize"] = maxGithubPageSize
	v["tagPageSize"] = 0
	totalCount = totalCountData.Repository.Labels.TotalCount
	if opt.First != nil {
		totalCount = math.Min(totalCount, *opt.First)
	}
	log.Printf("Get data label count: %d \n", totalCount)
	FetchLabelByRepo(request, totalCount, v, &totalData)

	v["userPageSize"] = maxGithubPageSize
	v["labelPageSize"] = 0
	totalCount = totalCountData.Repository.AssignableUsers.TotalCount
	if opt.First != nil {
		totalCount = math.Min(totalCount, *opt.First)
	}
	log.Printf("Get data label count: %d \n", totalCount)
	FetchUserByRepo(request, totalCount, v, &totalData)
	return totalData
}

// FetchIssueByRepo Fetch all the Query issues necessary.
func FetchIssueByRepo(request client.Request,
	totalCount int, v map[string]interface{}, query *model.Query) {
	log.Printf("Get data issue count: %d \n", totalCount)
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		log.Printf("<Fetching issue data %d to %d\n", count, count+math.Min(totalCount-count, maxGithubPageSize))
		v["IssuePageSize"] = math.Min(totalCount-count, maxGithubPageSize)
		var respData model.Query
		retryTimes := 10
		for {
			err := request.QueryWithAuthPool(context.Background(), &respData, v)
			if err != nil {
				log.Printf(err.Error()+" \n query variables: %v \n retry time: %d", v, 10-retryTimes)
			} else {
				break
			}
			retryTimes--
			if retryTimes == 0 {
				log.Fatal(err.Error()+" \n query variables: %v \n retry time: %d", v, 10-retryTimes)
			}
		}
		log.Printf("Fetch success.>\n")
		query.Repository.Issues.Nodes = append(query.Repository.Issues.Nodes, respData.Repository.Issues.Nodes...)
		if !respData.Repository.Issues.PageInfo.HasNextPage {
			break
		}
		v["issueAfter"] = respData.Repository.Issues.PageInfo.EndCursor
	}

	query.Repository.Issues.TotalCount = totalCount
}

// FetchTagByRepo Fetch all the Tag necessary.
func FetchTagByRepo(request client.Request,
	totalCount int, v map[string]interface{}, query *model.Query) {
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		log.Printf("<Fetching tag data %d to %d\n", count, count+math.Min(totalCount-count, maxGithubPageSize))
		v["tagPageSize"] = math.Min(totalCount-count, maxGithubPageSize)
		var respData model.Query
		retryTimes := 10
		for {
			err := request.QueryWithAuthPool(context.Background(), &respData, v)
			if err != nil {
				log.Printf(err.Error()+" \n query variables: %v \n retry time: %d", v, retryTimes)
			} else {
				break
			}
			retryTimes--
			if retryTimes == 0 {
				log.Fatal(err.Error()+" \n query variables: %v \n retry time: %d", v, 10-retryTimes)
			}
		}
		log.Printf("Fetch success.>\n")
		query.Repository.Refs.Nodes = append(query.Repository.Refs.Nodes, respData.Repository.Refs.Nodes...)
		if !respData.Repository.Refs.PageInfo.HasNextPage {
			break
		}
		v["tagAfter"] = respData.Repository.Refs.PageInfo.EndCursor
	}
	query.Repository.Refs.TotalCount = totalCount
}

// FetchLabelByRepo Fetch all the Label necessary.
func FetchLabelByRepo(request client.Request,
	totalCount int, v map[string]interface{}, query *model.Query) {
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		log.Printf("<Fetching Label data %d to %d\n", count, count+math.Min(totalCount-count, maxGithubPageSize))
		v["labelPageSize"] = math.Min(totalCount-count, maxGithubPageSize)
		var respData model.Query
		retryTimes := 10
		for {
			err := request.QueryWithAuthPool(context.Background(), &respData, v)
			if err != nil {
				log.Printf(err.Error()+" \n query variables: %v \n retry time: %d", v, retryTimes)
			} else {
				break
			}
			retryTimes--
			if retryTimes == 0 {
				log.Fatal(err.Error()+" \n query variables: %v \n retry time: %d", v, 10-retryTimes)
			}
		}
		log.Printf("Fetch success.>\n")
		query.Repository.Labels.Nodes = append(query.Repository.Labels.Nodes, respData.Repository.Labels.Nodes...)
		if !respData.Repository.Labels.PageInfo.HasNextPage {
			break
		}
		v["labelAfter"] = respData.Repository.Labels.PageInfo.EndCursor
	}
	query.Repository.Labels.TotalCount = totalCount
}

// FetchUserByRepo Fetch all the User necessary.
func FetchUserByRepo(request client.Request,
	totalCount int, v map[string]interface{}, query *model.Query) {
	for count := 0; count < totalCount; count += math.Min(totalCount-count, maxGithubPageSize) {
		log.Printf("<Fetching User data %d to %d\n", count, count+math.Min(totalCount-count, maxGithubPageSize))
		v["userPageSize"] = math.Min(totalCount-count, maxGithubPageSize)
		var respData model.Query
		retryTimes := 10
		for {
			err := request.QueryWithAuthPool(context.Background(), &respData, v)
			if err != nil {
				log.Printf(err.Error()+" \n query variables: %v \n retry time: %d", v, retryTimes)
			} else {
				break
			}
			retryTimes--
			if retryTimes == 0 {
				log.Fatal(err.Error()+" \n query variables: %v \n retry time: %d", v, 10-retryTimes)
			}
		}
		log.Printf("Fetch success.>\n")
		query.Repository.AssignableUsers.Nodes = append(query.Repository.AssignableUsers.Nodes, respData.Repository.AssignableUsers.Nodes...)
		if !respData.Repository.AssignableUsers.PageInfo.HasNextPage {
			break
		}
		v["userAfter"] = respData.Repository.AssignableUsers.PageInfo.EndCursor
	}
	query.Repository.AssignableUsers.TotalCount = totalCount
}
