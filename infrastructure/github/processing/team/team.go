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
		"SQL-Engine":   false,
		"TP-Arch":      false,
		"SQL-Metadata": false,
		"TP-Storage":   false,
		"Scheduling":   false,
	}
	switch repo {
	case "tidb":
		for _, label := range issue.Labels.Nodes {
			if exist(label.Name, []string{
				"sig/execution", "sig/planner", "component/coprocessor", "component/expression",
				"component/json", "component/executor", "component/statistics", "component/bindinfo"}) {
				teamsMap["SQL-Engine"] = true
			}

			if exist(label.Name, []string{
				"sig/transaction", "component/store", "component/tikv", "component/unistore",
				"component/mysql-protocol", "component/plugin", "component/privilege"}) {
				teamsMap["TP-Arch"] = true
			}

			if exist(label.Name, []string{"sig/DDL", "component/charset", "component/parser", "component/infoschema", "component/ddl"}) {
				teamsMap["SQL-Metadata"] = true
			}
		}
	case "tikv":
		teamsMap["TP-Storage"] = true
		for _, label := range issue.Labels.Nodes {
			if exist(label.Name, []string{"sig/transaction", "component/storage"}) {
				teamsMap["TP-Arch"] = true
				teamsMap["TP-Storage"] = false
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
