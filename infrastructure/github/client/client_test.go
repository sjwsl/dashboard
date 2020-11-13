package client

import (
	"context"
	"dashboard/infrastructure/github/config"
	"dashboard/infrastructure/github/model"
	"testing"
)

func TestName(t *testing.T) {
	config := config.Config{
		GraphqlPath:   "../graphql/query.graphql",
		ServerUrl:     "https://api.github.com/graphql",
		Authorization: []string{"81026979e15de49bd71b049d9418779d61837a5c"},
	}
	InitClient(config)
	request := NewClient()
	v := map[string]interface{}{
		"owner":        "pingcap",
		"repo_name":    "tidb",
		"issueFilters": map[string]interface{}{"states": []string{"OPEN", "CLOSED"}},
		"issueAfter":   nil,
		"commentAfter": nil,
	}
	var respData model.Query
	err := request.QueryWithAuthPool(context.Background(), &respData, v)
	if err != nil {
		panic(err)
	}

}
