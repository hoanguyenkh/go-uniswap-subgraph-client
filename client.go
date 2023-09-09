package unigraphclient

import (
	"context"
	"net/http"

	"github.com/emersonmacro/go-uniswap-subgraph-client/graphql"

	"github.com/mitchellh/mapstructure"
)

type Client struct {
	hostUrl   string
	GqlClient *graphql.Client
}

func NewClient(url string, opts *ClientOptions) *Client {
	if opts == nil {
		opts = &ClientOptions{}
	}

	if opts.httpClient == nil {
		opts.httpClient = http.DefaultClient
	}

	var gqlClient *graphql.Client

	if opts.closeReq {
		gqlClient = graphql.NewClient(url, graphql.WithHTTPClient(opts.httpClient), graphql.ImmediatelyCloseReqBody())
	} else {
		gqlClient = graphql.NewClient(url, graphql.WithHTTPClient(opts.httpClient))
	}

	return &Client{
		hostUrl:   url,
		GqlClient: gqlClient,
	}
}

func (c *Client) GetPoolById(ctx context.Context, id string, opts *RequestOptions) (*PoolResult, error) {
	if opts == nil {
		opts = &RequestOptions{
			IncludeFields: []string{"*"},
		}
	}

	req, err := constructPoolByIdQuery(id, opts)
	if err != nil {
		return nil, err
	}

	var resp interface{}
	if err := c.GqlClient.Run(ctx, req, &resp); err != nil {
		return nil, err
	}

	var converted PoolResult
	if err := mapstructure.Decode(resp, &converted); err != nil {
		return nil, err
	}

	return &converted, nil
}

type ClientOptions struct {
	httpClient *http.Client
	closeReq   bool
}

type RequestOptions struct {
	IncludeFields []string
	ExcludeFields []string
}
