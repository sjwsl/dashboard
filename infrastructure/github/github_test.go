package github

import (
	"testing"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
)

func TestRunInfrastructure(t *testing.T) {
	c := config.GetConfig("../../config.toml")
	c.CrawlerConfig.GraphqlPath = "./crawler/graphql/query.graphql"
	RunInfrastructure(c)
}

