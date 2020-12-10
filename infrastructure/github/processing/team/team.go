package team

import (
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
)

func exist(target string, list []string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}
	return false
}

func GetTeams(repo string, issue *model.Issue) []string {
	teamsMap := map[string]bool{
		"Runtime":                     false,
		"Optimizer":                   false,
		"Transaction-A Transaction-B": false,
		"General":                     false,
		"Scheduling":                  false,
		"TP-Storage Cloud-Storage":    false,
	}
	switch repo {
	case "tidb":
		for _, label := range issue.Labels.Nodes {
			if exist(label.Name, []string{
				"sig/execution", "component/coprocessor", "component/expression", "component/json", "component/executor",
			}) {
				teamsMap["Runtime"] = true
			}

			if exist(label.Name, []string{
				"sig/planner", "component/statistics", "component/bindinfo"}) {
				teamsMap["Optimizer"] = true
			}

			if exist(label.Name, []string{"sig/transaction", "component/store", "component/tikv", "component/unistore"}) {
				teamsMap["Transaction-A Transaction-B"] = true
			}

			if exist(label.Name, []string{"sig/DDL", "component/charset", "component/parser", "component/infoscheme", "component/ddl", "component/mysql-protocol", "component/plugin",
				"component/privilege", "component/config", "component/server"}) {
				teamsMap["General"] = true
			}
		}
	case "tikv":
		teamsMap["TP-Storage Cloud-Storage"] = true
		for _, label := range issue.Labels.Nodes {
			if exist(label.Name, []string{"sig/transaction", "component/storage"}) {
				teamsMap["TP-Storage Cloud-Storage"] = false
				teamsMap["Transaction-A Transaction-B"] = true
				break
			}
		}
	case "pd":
		teamsMap["Scheduling"] = true
	default:
	}

	teams := make([]string, 0)

	for team, vis := range teamsMap {
		if vis {
			teams = append(teams, team)
		}
	}
	return teams
}
