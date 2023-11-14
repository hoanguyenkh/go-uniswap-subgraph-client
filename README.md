# go-uniswap-subgraph-client

A Go library for querying Uniswap v3 subgraphs.

## Installation

```
$ go get github.com/emersonmacro/go-uniswap-subgraph-client
```

## Basic Usage

```go
import (
  "context"
  "fmt"
  unigraphclient "github.com/emersonmacro/go-uniswap-subgraph-client"
)

poolId := "0xc2e9f25be6257c210d7adf0d4cd6e3e881ba25f8"
endpoint := unigraphclient.Endpoints[unigraphclient.Ethereum]

client := unigraphclient.NewClient(endpoint, nil)

requestOpts := &unigraphclient.RequestOptions{
  IncludeFields: []string{"*"},
}
response, err := client.GetPoolById(context.Background(), poolId, requestOpts)

fmt.Println(response.Pool.ID) // 0xc2e9f25be6257c210d7adf0d4cd6e3e881ba25f8
fmt.Println(response.Pool.Token0.Symbol) // DAI
fmt.Println(response.Pool.Token1.Symbol) // WETH
fmt.Println(response.Pool.VolumeUSD) // 12790614690.20250028366283473022774
```

All [Uniswap v3 models](https://github.com/Uniswap/v3-subgraph/blob/main/schema.graphql) are supported. Each model can by queried by ID (`Get<Model>ById`) or as a list (`List<Model>`). This is the full list of models:

```
Factory
Pool
Token
Bundle
Tick
Position
PositionSnapshot
Transaction
Mint
Burn
Swap
Collect
Flash
UniswapDayData
PoolDayData
PoolHourData
TickHourData
TickDayData
TokenDayData
TokenHourData
```

See the [`test` directory](https://github.com/emersonmacro/go-uniswap-subgraph-client/tree/master/test) for more usage examples.

## Known Issues

- Derived fields are currently not supported
- `where` clauses are currently not supported
- All response fields are returned as strings, regardless of their underlying type. See the `converter` package for some utility functions for converting to `*big.Int` or `*big.Float`

## Client Options

```go
type ClientOptions struct {
  HttpClient *http.Client // option to pass in your own http client (http.DefaultClient by default)
  CloseReq   bool // option to close the request immediately
}

func NewClient(url string, opts *ClientOptions) *Client
```

## Request Options

There are two ways to specify the fields you want to be included in the query. `IncludeFields` can be used to "opt in" to the fields you want, and `"*"` is a valid option to include all fields. Alternatively, you can include all fields and then exclude certain fields ("opt out") with `ExcludeFields`.

You can query data at a particular block with the `Block` option. For `List*` queries, pagination is supported with the `First` and `Skip` options, and sorting is supported with the `OrderBy` and `OrderDir` options.

```go
type RequestOptions struct {
  IncludeFields []string // fields to include in the query. '*' is a valid option meaning 'include all fields'. if any fields are listed in IncludeFields besides '*', ExcludeFields must be empty.
  ExcludeFields []string // fields to exclude from the query. only valid when '*' is in IncludeFields.
  Block         int      // query for data at a specific block number.
  First         int      // number of results to retrieve. `100` is the default. only valid for List queries.
  Skip          int      // number of results to skip. `0` is the default. only valid for List queries.
  OrderBy       string   // field to order by. `id` is the default. only valid for List queries.
  OrderDir      string   // order direction. `asc` for ascending and `desc` for descending are the only valid options. `asc` is the default. only valid for List queries.
}
```

## Endpoints

When creating a new client, you can specify any subgraph endpoint that supports a Uniswap v3 schema:

```go
client := unigraphclient.NewClient("https://<my graphql host>", nil)
```

For convenience, you can also use one of the provided endpoints, which are the same as the endpoints used in the [Uniswap Info](https://info.uniswap.org/#/) site:

```
// e.g.:
endpoint := unigraphclient.Endpoints[unigraphclient.Ethereum] // https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3
// or
endpoint := unigraphclient.Endpoints[unigraphclient.Arbitrum] // https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-arbitrum-one
// or
endpoint := unigraphclient.Endpoints[unigraphclient.Optimism] // https://api.thegraph.com/subgraphs/name/ianlapham/optimism-post-regenesis
// or
endpoint := unigraphclient.Endpoints[unigraphclient.Polygon] // https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-polygon
// or
endpoint := unigraphclient.Endpoints[unigraphclient.Base] // https://api.studio.thegraph.com/query/48211/uniswap-v3-base/version/latest
// or
endpoint := unigraphclient.Endpoints[unigraphclient.Celo] // https://api.thegraph.com/subgraphs/name/jesse-sawa/uniswap-celo
// or
endpoint := unigraphclient.Endpoints[unigraphclient.Avalanche] //https://api.thegraph.com/subgraphs/name/lynnshaoyu/uniswap-v3-avax
// or
endpoint := unigraphclient.Endpoints[unigraphclient.Bnb] // https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-bsc

client := unigraphclient.NewClient(endpoint, nil)
```

## Converter utility functions

```
func StringToBigInt(s string) (*big.Int, error)
func StringToBigFloat(s string) (*big.Float, error)
func ModelToJsonBytes(model any) ([]byte, error)
func ModelToJsonString(model any) (string, error)
```

## Resources

- [Uniswap Subgraph Overview](https://docs.uniswap.org/api/subgraph/overview)
- [Uniswap Subgraph Query Examples](https://docs.uniswap.org/api/subgraph/guides/examples)
- [Uniswap v3 Subgraph Schemas](https://github.com/Uniswap/v3-subgraph/blob/main/schema.graphql)
