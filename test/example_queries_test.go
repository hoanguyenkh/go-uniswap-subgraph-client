// example queries from Uniswap subgraph docs https://docs.uniswap.org/api/subgraph/guides/examples

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	unigraphclient "github.com/emersonmacro/go-uniswap-subgraph-client"

	"github.com/stretchr/testify/assert"
)

func TestUniswapExampleQueries(t *testing.T) {
	var endpoint string = unigraphclient.Endpoints[unigraphclient.Ethereum]

	t.Run("current global data query (factory by id)", func(t *testing.T) {
		factoryId := "0x1F98431c8aD98523631AE4a59f267346ea31F984"
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"poolCount",
				"txCount",
				"totalVolumeUSD",
				"totalVolumeETH",
			},
		}
		resp, err := client.GetFactoryById(context.Background(), factoryId, reqOpts)

		assert.Nil(t, err)
		assert.NotEqual(t, "", resp.Factory.PoolCount)
		assert.NotEqual(t, "", resp.Factory.TxCount)
		assert.NotEqual(t, "", resp.Factory.TotalVolumeUSD)
		assert.NotEqual(t, "", resp.Factory.TotalVolumeETH)
	})

	t.Run("historical global data query (factory by id)", func(t *testing.T) {
		factoryId := "0x1F98431c8aD98523631AE4a59f267346ea31F984"
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"poolCount",
				"txCount",
				"totalVolumeUSD",
				"totalVolumeETH",
			},
			Block: 13380584,
		}
		resp, err := client.GetFactoryById(context.Background(), factoryId, reqOpts)

		assert.Nil(t, err)
		assert.Equal(t, "4530", resp.Factory.PoolCount)
		assert.Equal(t, "6420502", resp.Factory.TxCount)
		assert.Equal(t, "175360787979.8419091840895098379928", resp.Factory.TotalVolumeUSD)
		assert.Equal(t, "63933597.84932237533939399849568679", resp.Factory.TotalVolumeETH)
	})

	t.Run("general pool query (pool by id)", func(t *testing.T) {
		poolId := "0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8"
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"tick",
				"token0.symbol",
				"token0.id",
				"token0.decimals",
				"token1.symbol",
				"token1.id",
				"token1.decimals",
				"feeTier",
				"sqrtPrice",
				"liquidity",
			},
		}
		resp, err := client.GetPoolById(context.Background(), poolId, reqOpts)

		assert.Nil(t, err)
		assert.NotEqual(t, "", resp.Pool.Tick)
		assert.NotEqual(t, "", resp.Pool.SqrtPrice)
		assert.NotEqual(t, "", resp.Pool.Token0.ID)
		assert.NotEqual(t, "", resp.Pool.Token1.Symbol)
	})

	t.Run("all possible pools - skipping first 1000 pools (list pools)", func(t *testing.T) {
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"id",
				"token0.symbol",
				"token0.id",
				"token1.symbol",
				"token1.id",
			},
			First: 10,
			Skip:  1000,
		}
		resp, err := client.ListPools(context.Background(), reqOpts)

		assert.Nil(t, err)
		assert.True(t, len(resp.Pools) > 0)
		assert.NotEqual(t, "", resp.Pools[0].ID)
		assert.NotEqual(t, "", resp.Pools[0].Token0.ID)
		assert.NotEqual(t, "", resp.Pools[0].Token1.ID)
	})

	// see test/pagination_test.go for a `creating a skip variable` example

	t.Run("all possible pools - most liquid pools (list pools)", func(t *testing.T) {
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"id",
			},
			First: 1000,
		}
		resp, err := client.ListPools(context.Background(), reqOpts)

		assert.Nil(t, err)
		assert.Equal(t, 1000, len(resp.Pools))
	})

	// pool daily aggregated - TODO: add support for `where` clause in queries

	t.Run("general swap data query (swap by id)", func(t *testing.T) {
		swapId := "0x000007e1111cbd97f74cfc6eea2879a5b02020f26960ac06f4af0f9395372b64#66785"
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"sender",
				"recipient",
				"amount0",
				"amount1",
				"transaction.id",
				"transaction.blockNumber",
				"transaction.gasUsed",
				"transaction.gasPrice",
				"timestamp",
				"token0.id",
				"token0.symbol",
				"token1.id",
				"token1.symbol",
			},
		}
		resp, err := client.GetSwapById(context.Background(), swapId, reqOpts)

		assert.Nil(t, err)
		assert.NotEqual(t, "", resp.Swap.Sender)
		assert.NotEqual(t, "", resp.Swap.Amount0)
		assert.NotEqual(t, "", resp.Swap.Transaction.ID)
		assert.NotEqual(t, "", resp.Swap.Token0.Symbol)
	})

	// recent swaps within a pool - TODO: add support for `where` clause in queries

	t.Run("general token data query (token by id)", func(t *testing.T) {
		tokenId := "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"symbol",
				"name",
				"decimals",
				"volumeUSD",
				"poolCount",
			},
		}
		resp, err := client.GetTokenById(context.Background(), tokenId, reqOpts)

		assert.Nil(t, err)
		assert.NotEqual(t, "", resp.Token.Symbol)
		assert.NotEqual(t, "", resp.Token.Name)
	})

	// token daily aggregated - TODO: add support for `where` clause in queries

	t.Run("all tokens query (list tokens)", func(t *testing.T) {
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"id",
				"symbol",
				"name",
			},
			First: 1000,
			Skip:  100,
		}
		resp, err := client.ListTokens(context.Background(), reqOpts)

		assert.Nil(t, err)
		assert.Equal(t, 1000, len(resp.Tokens))
	})

	t.Run("general position data query (position by id)", func(t *testing.T) {
		positionId := "3"
		client := unigraphclient.NewClient(endpoint, nil)

		reqOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"id",
				"collectedFeesToken0",
				"collectedFeesToken1",
				"liquidity",
				"token0.id",
				"token0.symbol",
				"token1.id",
				"token1.symbol",
			},
		}
		resp, err := client.GetPositionById(context.Background(), positionId, reqOpts)

		assert.Nil(t, err)
		assert.NotEqual(t, "", resp.Position.ID)
		assert.NotEqual(t, "", resp.Position.CollectedFeesToken0)
		assert.NotEqual(t, "", resp.Position.Token0.ID)
		assert.NotEqual(t, "", resp.Position.Token1.Symbol)
	})

	t.Run("Get history swap", func(t *testing.T) {
		poolId := "0x0c5527e51d6be6bf3e93711e38feb2ee611c99cb"
		endpoint = unigraphclient.Endpoints[unigraphclient.Base]
		client := unigraphclient.NewClient(endpoint, nil)

		requestOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"id",
				"timestamp",
				"amount0",
				"amount1",
				"amountUSD",
				"sqrtPriceX96",
				"sender",
				"recipient",
				"pool.token0.id",
				"pool.token0.symbol",
				"pool.token1.id",
				"pool.token1.symbol",
				"transaction.id",
			},
		}
		response, err := client.GetSwapHistoryByPoolId(context.Background(), poolId, requestOpts)
		if err != nil {
			fmt.Println(err)
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshaling response:", err)
			return
		}

		fmt.Println(string(responseBytes))
	})

	t.Run("Get meme buy/sell history", func(t *testing.T) {
		memeAddress := "0x9839e570cbaeb1715d8191f33d2b58745142cf87"
		client := unigraphclient.NewClient("https://api.studio.thegraph.com/query/76502/membots-ai-memeception-mvp/version/latest", nil)

		requestOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"*",
			},
		}
		response, err := client.GetSwapHistoryByMemeToken(context.Background(), memeAddress, requestOpts)
		if err != nil {
			fmt.Println(err)
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshaling response:", err)
			return
		}

		fmt.Println(string(responseBytes))
	})

	t.Run("Get meme create tiers", func(t *testing.T) {
		memeAddress := "0xc511c410beb32c206a657f2dd0cd412862281d29"
		client := unigraphclient.NewClient("https://api.studio.thegraph.com/query/76502/membots-ai-memeception-mvp/version/latest", nil)

		requestOpts := &unigraphclient.RequestOptions{
			IncludeFields: []string{
				"*",
			},
		}
		response, err := client.GetMemeTiersByMemeToken(context.Background(), memeAddress, requestOpts)
		if err != nil {
			fmt.Println(err)
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshaling response:", err)
			return
		}

		fmt.Println(string(responseBytes))
	})
}
