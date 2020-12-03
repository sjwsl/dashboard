package client

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/machinebox/graphql"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/util"
)

var requests []*graphql.Request
var client *graphql.Client
var clientIsOpen bool

type Request struct {
	req   *graphql.Request
	index int
}

// NewClient return a client with request index 0
func NewClient() Request {
	if !clientIsOpen {
		panic(fmt.Errorf("clients need to be init before use it , you could init it by InitClient"))
	}
	return Request{requests[0], 0}
}

// InitClient init Client with all the config, you must InitClient before NewClient
func InitClient(config config.Config) {
	CheckConfig(config)
	content, err := util.ReadFile(config.GraphqlPath)
	if err != nil {
		panic(err)
	}

	requests = make([]*graphql.Request, len(config.Authorization))
	client = graphql.NewClient(config.ServerUrl)
	for i, s := range config.Authorization {
		req := graphql.NewRequest(string(content))
		req.Header.Set("authorization", "Bearer "+s)
		requests[i] = req
	}
	clientIsOpen = true
}

// CheckConfig check the init config.
func CheckConfig(c config.Config) {
	// Client use net/package to parse url too.
	_, err := url.Parse(c.ServerUrl)
	if err != nil {
		panic(err)
	}

	_, err = os.Open(c.GraphqlPath)
	if err != nil {
		absPath, _ := filepath.Abs(c.GraphqlPath)
		log.Fatal(absPath, err)
	}

	if len(c.Authorization) == 0 {
		panic(fmt.Errorf("there must have at least one token string in config"))
	}
}

type ErrAllRequestFailed struct {
	errStr string
}

func (a ErrAllRequestFailed) Error() string {
	return a.errStr
}

// QueryWithAuthPool package the requests pool, you could use it just like client.Run in machinebox/graphql package
func (req Request) QueryWithAuthPool(ctx context.Context, resp interface{}, variables map[string]interface{}) error {
	for {
		for key, arg := range variables {
			req.req.Var(key, arg)
		}
		err := client.Run(ctx, req.req, resp)

		if err != nil {
			if req.index == len(requests)-1 {
				var aerr ErrAllRequestFailed
				aerr.errStr = err.Error() + " ; all tokens has been used, but could not stop the steps of errors."
				return aerr
			} else {
				req.index++
				req.req = requests[req.index]
			}
		} else {
			return nil
		}
	}
}
