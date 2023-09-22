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

// options when creating a new Client
type ClientOptions struct {
	HttpClient *http.Client
	CloseReq   bool
}

// options when creating a new Request
type RequestOptions struct {
	IncludeFields []string // fields to include in the query. '*' is a valid option meaning 'include all fields'. if any fields are listed in IncludeFields besides '*', ExcludeFields must be empty.
	ExcludeFields []string // fields to exclude from the query. only valid when '*' is in IncludeFields.
	Block         int      // query for data at a specific block number.
	First         int      // number of results to retrieve. `100` is the default. only valid for List queries.
	Skip          int      // number of results to skip. `0` is the default. only valid for List queries.
	OrderBy       string   // field to order by. `id` is the default. only valid for List queries.
	OrderDir      string   // order direction. `asc` for ascending and `desc` for descending are the only valid options. `asc` is the default. only valid for List queries.
}

// type constraint for executeRequestAndConvert
type Response interface {
	FactoryResponse | ListFactoriesResponse | PoolResponse | ListPoolsResponse | TokenResponse | ListTokensResponse
}

// intermediate struct used to construct queries
type fieldRefs struct {
	directs []string
	refs    []string
}

// query type enum
type QueryType int

const (
	ById QueryType = iota
	List
)
