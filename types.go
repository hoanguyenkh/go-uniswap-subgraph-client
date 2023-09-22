package unigraphclient

import (
	"net/http"

	"github.com/emersonmacro/go-uniswap-subgraph-client/graphql"
)

// Uniswap model types are defined in models.go

// main uniswap subgraph client
type Client struct {
	hostUrl   string
	GqlClient *graphql.Client
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

// type constraint for executeRequestAndConvert
type Response interface {
	FactoryResponse | PoolResponse | TokenResponse
}

// intermediate struct used to construct queries
type fieldRefs struct {
	directs []string
	refs    []string
}
