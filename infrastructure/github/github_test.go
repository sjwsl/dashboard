package github

import (
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"testing"
)

func TestRunInfrastructure(t *testing.T) {
	c := config.GetConfig("../../config.toml")
	c.CrawlerConfig.GraphqlPath = "./crawler/graphql/query.graphql"
	RunInfrastructure(c)
}

func TestName(t *testing.T) {

}
