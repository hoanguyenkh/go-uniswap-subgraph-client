package test

import (
	"context"
	"testing"

	unigraphclient "github.com/emersonmacro/go-uniswap-subgraph-client"

	"github.com/stretchr/testify/assert"
)

func TestEndpoints(t *testing.T) {
	tests := map[string]struct {
		endpoint string
	}{
		"ethereum": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Ethereum],
		},
		"arbitrum": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Arbitrum],
		},
		"optimism": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Optimism],
		},
		"polygon": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Polygon],
		},
		"celo": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Celo],
		},
		"bnb": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Bnb],
		},
		"base": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Base],
		},
		"avalanche": {
			endpoint: unigraphclient.Endpoints[unigraphclient.Avalanche],
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			client := unigraphclient.NewClient(test.endpoint, nil)
			reqOpts := &unigraphclient.RequestOptions{
				IncludeFields: []string{"*"},
				ExcludeFields: []string{"feeGrowthGlobal0X128", "feeGrowthGlobal1X128"},
				First:         5,
			}
			resp, err := client.ListPools(ctx, reqOpts)
			assert.Nil(t, err)
			assert.True(t, len(resp.Pools) > 0)
		})
	}
}
