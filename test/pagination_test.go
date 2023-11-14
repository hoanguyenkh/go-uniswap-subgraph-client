package test

import (
	"context"
	"testing"

	unigraphclient "github.com/emersonmacro/go-uniswap-subgraph-client"

	"github.com/stretchr/testify/assert"
)

func TestPagination(t *testing.T) {
	var (
		endpoint string = unigraphclient.Endpoints[unigraphclient.Ethereum]
		pageSize int    = 100
	)

	client := unigraphclient.NewClient(endpoint, nil)

	// first page (skip 0)
	reqOpts1 := &unigraphclient.RequestOptions{
		IncludeFields: []string{
			"id",
		},
		First: pageSize,
		Skip:  0,
	}
	resp1, err := client.ListPools(context.Background(), reqOpts1)

	assert.Nil(t, err)
	assert.Equal(t, 100, len(resp1.Pools))

	// second page (skip 100)
	reqOpts2 := &unigraphclient.RequestOptions{
		IncludeFields: []string{
			"id",
		},
		First: pageSize,
		Skip:  pageSize,
	}
	resp2, err := client.ListPools(context.Background(), reqOpts2)

	assert.Nil(t, err)
	assert.Equal(t, 100, len(resp2.Pools))
	assert.NotEqual(t, resp1.Pools[0].ID, resp2.Pools[0].ID)

	// third page (skip 200)
	reqOpts3 := &unigraphclient.RequestOptions{
		IncludeFields: []string{
			"id",
		},
		First: pageSize,
		Skip:  pageSize * 2,
	}
	resp3, err := client.ListPools(context.Background(), reqOpts3)

	assert.Nil(t, err)
	assert.Equal(t, 100, len(resp3.Pools))
	assert.NotEqual(t, resp2.Pools[0].ID, resp3.Pools[0].ID)
}
