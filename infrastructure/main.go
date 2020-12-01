package main

import (
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"time"
)

func main() {
	c := config.GetConfig("./config.toml")
	github.RunInfrastructure(c)
	fmt.Printf(
		`
	###########################################################################################
	init db ok %v
	###########################################################################################
	`, time.Now())
	for {
		time.Sleep(time.Hour)
		github.RunInfrastructure(c)
		fmt.Printf(
			`
	###########################################################################################
	update database %v
	###########################################################################################
	`, time.Now())
	}
}
