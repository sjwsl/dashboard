package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/util"
)

type Config struct {
	GraphqlPath   string
	ServerUrl     string
	Authorization []string
}

// FromJSON form config by json file path
func FromJSON(path string) Config {
	content, err := util.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var result Config
	err = json.Unmarshal(content, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// FromEnv form config by Env
func FromEnv() Config {
	var result Config
	result.GraphqlPath = os.Getenv("graphql_path")
	result.ServerUrl = os.Getenv("server_url")
	AuthorizationString := os.Getenv("authorization")
	result.Authorization = strings.Split(AuthorizationString, ":")
	return result
}
