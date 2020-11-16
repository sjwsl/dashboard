package client

import (
	"context"
	"dashboard/infrastructure/github/config"
	"dashboard/infrastructure/github/util"
	"fmt"
	"github.com/google/martian/log"
	"github.com/machinebox/graphql"
)

var requests []*graphql.Request
var client *graphql.Client
var clientIsOpen bool

type Request struct {
	req   *graphql.Request
	index int
}

func NewClient() Request {
	if !clientIsOpen {
		panic(fmt.Errorf("clients need to be init before use it , you could init it by InitClient"))
	}
	return Request{requests[0], 0}
}

func InitClient(config config.Config) {
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

// QueryWithAuthPool package the requests pool, you could use it just like client.Run in machinebox/graphql package
func (req Request) QueryWithAuthPool(ctx context.Context, resp interface{}, variables map[string]interface{}) error {
	reTryTimes := 3
	for {
		for key, arg := range variables {
			req.req.Var(key, arg)
		}
		err := client.Run(ctx, req.req, resp)

		if err != nil {
			if reTryTimes != 0 {
				log.Errorf("%v,err:%v", variables, err)
				reTryTimes--
				continue
			}
			if req.index == len(requests)-1 {
				log.Errorf("All tokens has been used, but could not stop the steps of errors.")
				return err
			} else {
				req.index++
				req.req = requests[req.index]
			}
		} else {
			return nil
		}
	}

}
