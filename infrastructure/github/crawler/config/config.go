package config

import (
	"dashboard/infrastructure/github/crawler/util"
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	GraphqlPath   string   `json:"graphql_path"`
	ServerUrl     string   `json:"server_url"`
	Authorization []string `json:"authorization"`
}

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

func FromEnv() Config {
	var result Config
	result.GraphqlPath = os.Getenv("graphql_path")
	result.ServerUrl = os.Getenv("server_url")
	AuthorizationString := os.Getenv("authorization")
	result.Authorization = strings.Split(AuthorizationString, ":")
	return result
}
