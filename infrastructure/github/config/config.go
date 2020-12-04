package config

import (
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/martian/log"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/util"
	dbConfig "github.com/PingCAP-QE/dashboard/infrastructure/github/database/config"
)

type Config struct {
	RepositoryArgs      []RepositoryArg
	CrawlerConfig       config.Config
	LabelSeverityWeight []LabelSeverityWeight
	CreateAtGlobal      time.Time
	DatabaseConfig      dbConfig.Config
	TeamArgs            []TeamArg
}

func GetConfig(path string) Config {
	content, err := util.ReadFile(path)
	if err != nil {
		absPath, _ := filepath.Abs(path)
		log.Errorf("%s", absPath)
		panic(err)
	}

	var result Config
	_, err = toml.Decode(string(content), &result)
	if err != nil {
		panic(err)
	}
	return result
}
