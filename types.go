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
	Id            string   // query for data by id.
	Block         int      // query for data at a specific block number.
	First         int      // number of results to retrieve. `100` is the default. only valid for List queries.
	Skip          int      // number of results to skip. `0` is the default. only valid for List queries.
	OrderBy       string   // field to order by. `id` is the default. only valid for List queries.
	OrderDir      string   // order direction. `asc` for ascending and `desc` for descending are the only valid options. `asc` is the default. only valid for List queries.
}

// type constraint for executeRequestAndConvert
type Response interface {
	FactoryResponse | ListFactoriesResponse |
		PoolResponse | ListPoolsResponse |
		TokenResponse | ListTokensResponse |
		BundleResponse | ListBundlesResponse |
		TickResponse | ListTicksResponse |
		PositionResponse | ListPositionsResponse |
		PositionSnapshotResponse | ListPositionSnapshotsResponse |
		TransactionResponse | ListTransactionsResponse |
		MintResponse | ListMintsResponse |
		BurnResponse | ListBurnsResponse |
		SwapResponse | ListSwapsResponse |
		CollectResponse | ListCollectsResponse |
		FlashResponse | ListFlashesResponse |
		UniswapDayDataResponse | ListUniswapDayDatasResponse |
		PoolDayDataResponse | ListPoolDayDatasResponse |
		PoolHourDataResponse | ListPoolHourDatasResponse |
		TickHourDataResponse | ListTickHourDatasResponse |
		TickDayDataResponse | ListTickDayDatasResponse |
		TokenDayDataResponse | ListTokenDayDatasResponse |
		TokenHourDataResponse | ListTokenHourDatasResponse
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

// endpoints enum
type Endpoint int

const (
	Ethereum Endpoint = iota
	Arbitrum
	Optimism
	Polygon
	Celo
	Bnb
	Base
	Avalanche
)

// endpoints map (values from https://github.com/Uniswap/v3-info/blob/master/src/apollo/client.ts)
var Endpoints map[Endpoint]string = map[Endpoint]string{
	Ethereum:  "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3",
	Arbitrum:  "https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-arbitrum-one",
	Optimism:  "https://api.thegraph.com/subgraphs/name/ianlapham/optimism-post-regenesis",
	Polygon:   "https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-polygon",
	Celo:      "https://api.thegraph.com/subgraphs/name/jesse-sawa/uniswap-celo",
	Bnb:       "https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-bsc",
	Base:      "https://api.studio.thegraph.com/query/48211/uniswap-v3-base/version/latest",
	Avalanche: "https://api.thegraph.com/subgraphs/name/lynnshaoyu/uniswap-v3-avax",
}
