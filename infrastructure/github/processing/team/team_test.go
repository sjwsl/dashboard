package team

import (
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"testing"
)

func TestGetTeams(t *testing.T) {
	repo := "tikv"
	issue := model.Issue{
		Labels: &model.LabelConnection{
			Nodes:      []*model.Label{{Name: "sig/transaction"}},
			PageInfo:   nil,
			TotalCount: 0,
		},
	}
	if len(GetTeams(repo, &issue)) != 1 || GetTeams(repo, &issue)[0] != "TP-Arch" {
		t.Error()
	}
}
