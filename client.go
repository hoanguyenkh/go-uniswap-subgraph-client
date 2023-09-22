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

	if opts.HttpClient == nil {
		opts.HttpClient = http.DefaultClient
	}

	var gqlClient *graphql.Client

	if opts.CloseReq {
		gqlClient = graphql.NewClient(url, graphql.WithHTTPClient(opts.HttpClient), graphql.ImmediatelyCloseReqBody())
	} else {
		gqlClient = graphql.NewClient(url, graphql.WithHTTPClient(opts.HttpClient))
	}

	return &Client{
		hostUrl:   url,
		GqlClient: gqlClient,
	}
}

func (c *Client) GetFactoryById(ctx context.Context, id string, opts *RequestOptions) (*FactoryResponse, error) {
	req, err := constructByIdQuery(id, FactoryFields, opts)
	if err != nil {
		return nil, err
	}

	return executeRequestAndConvert(ctx, req, FactoryResponse{}, c)
}

func (c *Client) GetPoolById(ctx context.Context, id string, opts *RequestOptions) (*PoolResponse, error) {
	req, err := constructByIdQuery(id, PoolFields, opts)
	if err != nil {
		return nil, err
	}

	return executeRequestAndConvert(ctx, req, PoolResponse{}, c)
}

func (c *Client) GetTokenById(ctx context.Context, id string, opts *RequestOptions) (*TokenResponse, error) {
	req, err := constructByIdQuery(id, TokenFields, opts)
	if err != nil {
		return nil, err
	}

	return executeRequestAndConvert(ctx, req, TokenResponse{}, c)
}

func executeRequestAndConvert[T Response](ctx context.Context, req *graphql.Request, converted T, c *Client) (*T, error) {
	var resp interface{}
	if err := c.GqlClient.Run(ctx, req, &resp); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(resp, &converted); err != nil {
		return nil, err
	}

	return &converted, nil
}

// type constraint for executeRequestAndConvert
type Response interface {
	FactoryResponse | PoolResponse | TokenResponse
}

// options when creating new Client
type ClientOptions struct {
	HttpClient *http.Client
	CloseReq   bool
}

// options when creating new Request
type RequestOptions struct {
	IncludeFields []string // fields to include in the query. '*' is a valid option meaning 'include all fields'. if any fields are listed in IncludeFields besides '*', ExcludeFields must be empty.
	ExcludeFields []string // fields to exclude from the query. only valid when '*' is in IncludeFields.
	Block         int      // query for data at a specific block number.
}
