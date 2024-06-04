package unigraphclient

import (
	"context"
	"net/http"

	"github.com/emersonmacro/go-uniswap-subgraph-client/graphql"

	"github.com/mitchellh/mapstructure"
)

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

func (c *Client) ListFactories(ctx context.Context, opts *RequestOptions) (*ListFactoriesResponse, error) {
	req, err := constructListQuery(FactoryFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListFactoriesResponse{}, c)
}

func (c *Client) GetPoolById(ctx context.Context, id string, opts *RequestOptions) (*PoolResponse, error) {
	req, err := constructByIdQuery(id, PoolFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, PoolResponse{}, c)
}

func (c *Client) GetSwapHistoryByPoolId(ctx context.Context, poolId string, opts *RequestOptions) (*ListSwapsResponse, error) {
	req, err := constructListQueryWithId(poolId, SwapFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListSwapsResponse{}, c)
}

func (c *Client) GetSwapHistoryByMemeToken(ctx context.Context, memeToken string, opts *RequestOptions) (*MemeCoinExitsResponse, error) {
	req, err := constructListQueryWithMemeToken(memeToken, MemeFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, MemeCoinExitsResponse{}, c)
}

func (c *Client) ListPools(ctx context.Context, opts *RequestOptions) (*ListPoolsResponse, error) {
	req, err := constructListQuery(PoolFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListPoolsResponse{}, c)
}

func (c *Client) GetTokenById(ctx context.Context, id string, opts *RequestOptions) (*TokenResponse, error) {
	req, err := constructByIdQuery(id, TokenFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, TokenResponse{}, c)
}

func (c *Client) ListTokens(ctx context.Context, opts *RequestOptions) (*ListTokensResponse, error) {
	req, err := constructListQuery(TokenFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListTokensResponse{}, c)
}

func (c *Client) GetBundleById(ctx context.Context, id string, opts *RequestOptions) (*BundleResponse, error) {
	req, err := constructByIdQuery(id, BundleFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, BundleResponse{}, c)
}

func (c *Client) ListBundles(ctx context.Context, opts *RequestOptions) (*ListBundlesResponse, error) {
	req, err := constructListQuery(BundleFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListBundlesResponse{}, c)
}

func (c *Client) GetTickById(ctx context.Context, id string, opts *RequestOptions) (*TickResponse, error) {
	req, err := constructByIdQuery(id, TickFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, TickResponse{}, c)
}

func (c *Client) ListTicks(ctx context.Context, opts *RequestOptions) (*ListTicksResponse, error) {
	req, err := constructListQuery(TickFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListTicksResponse{}, c)
}

func (c *Client) GetPositionById(ctx context.Context, id string, opts *RequestOptions) (*PositionResponse, error) {
	req, err := constructByIdQuery(id, PositionFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, PositionResponse{}, c)
}

func (c *Client) ListPositions(ctx context.Context, opts *RequestOptions) (*ListPositionsResponse, error) {
	req, err := constructListQuery(PositionFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListPositionsResponse{}, c)
}

func (c *Client) GetTransactionById(ctx context.Context, id string, opts *RequestOptions) (*TransactionResponse, error) {
	req, err := constructByIdQuery(id, TransactionFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, TransactionResponse{}, c)
}

func (c *Client) ListTransactions(ctx context.Context, opts *RequestOptions) (*ListTransactionsResponse, error) {
	req, err := constructListQuery(TransactionFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListTransactionsResponse{}, c)
}

func (c *Client) GetMintById(ctx context.Context, id string, opts *RequestOptions) (*MintResponse, error) {
	req, err := constructByIdQuery(id, MintFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, MintResponse{}, c)
}

func (c *Client) ListMints(ctx context.Context, opts *RequestOptions) (*ListMintsResponse, error) {
	req, err := constructListQuery(MintFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListMintsResponse{}, c)
}

func (c *Client) GetBurnById(ctx context.Context, id string, opts *RequestOptions) (*BurnResponse, error) {
	req, err := constructByIdQuery(id, BurnFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, BurnResponse{}, c)
}

func (c *Client) ListBurns(ctx context.Context, opts *RequestOptions) (*ListBurnsResponse, error) {
	req, err := constructListQuery(BurnFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListBurnsResponse{}, c)
}

func (c *Client) GetSwapById(ctx context.Context, id string, opts *RequestOptions) (*SwapResponse, error) {
	req, err := constructByIdQuery(id, SwapFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, SwapResponse{}, c)
}

func (c *Client) ListSwaps(ctx context.Context, opts *RequestOptions) (*ListSwapsResponse, error) {
	req, err := constructListQuery(SwapFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListSwapsResponse{}, c)
}

func (c *Client) GetCollectById(ctx context.Context, id string, opts *RequestOptions) (*CollectResponse, error) {
	req, err := constructByIdQuery(id, CollectFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, CollectResponse{}, c)
}

func (c *Client) ListCollects(ctx context.Context, opts *RequestOptions) (*ListCollectsResponse, error) {
	req, err := constructListQuery(CollectFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListCollectsResponse{}, c)
}

func (c *Client) GetFlashById(ctx context.Context, id string, opts *RequestOptions) (*FlashResponse, error) {
	req, err := constructByIdQuery(id, FlashFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, FlashResponse{}, c)
}

func (c *Client) ListFlashes(ctx context.Context, opts *RequestOptions) (*ListFlashesResponse, error) {
	req, err := constructListQuery(FlashFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListFlashesResponse{}, c)
}

func (c *Client) GetUniswapDayDataById(ctx context.Context, id string, opts *RequestOptions) (*UniswapDayDataResponse, error) {
	req, err := constructByIdQuery(id, UniswapDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, UniswapDayDataResponse{}, c)
}

func (c *Client) ListUniswapDayDatas(ctx context.Context, opts *RequestOptions) (*ListUniswapDayDatasResponse, error) {
	req, err := constructListQuery(UniswapDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListUniswapDayDatasResponse{}, c)
}

func (c *Client) GetPoolDayDataById(ctx context.Context, id string, opts *RequestOptions) (*PoolDayDataResponse, error) {
	req, err := constructByIdQuery(id, PoolDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, PoolDayDataResponse{}, c)
}

func (c *Client) ListPoolDayDatas(ctx context.Context, opts *RequestOptions) (*ListPoolDayDatasResponse, error) {
	req, err := constructListQuery(PoolDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListPoolDayDatasResponse{}, c)
}

func (c *Client) GetPoolHourDataById(ctx context.Context, id string, opts *RequestOptions) (*PoolHourDataResponse, error) {
	req, err := constructByIdQuery(id, PoolHourDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, PoolHourDataResponse{}, c)
}

func (c *Client) ListPoolHourDatas(ctx context.Context, opts *RequestOptions) (*ListPoolHourDatasResponse, error) {
	req, err := constructListQuery(PoolHourDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListPoolHourDatasResponse{}, c)
}

func (c *Client) GetTickHourDataById(ctx context.Context, id string, opts *RequestOptions) (*TickHourDataResponse, error) {
	req, err := constructByIdQuery(id, TickHourDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, TickHourDataResponse{}, c)
}

func (c *Client) ListTickHourDatas(ctx context.Context, opts *RequestOptions) (*ListTickHourDatasResponse, error) {
	req, err := constructListQuery(TickHourDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListTickHourDatasResponse{}, c)
}

func (c *Client) GetTickDayDataById(ctx context.Context, id string, opts *RequestOptions) (*TickDayDataResponse, error) {
	req, err := constructByIdQuery(id, TickDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, TickDayDataResponse{}, c)
}

func (c *Client) ListTickDayDatas(ctx context.Context, opts *RequestOptions) (*ListTickDayDatasResponse, error) {
	req, err := constructListQuery(TickDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListTickDayDatasResponse{}, c)
}

func (c *Client) GetTokenDayDataById(ctx context.Context, id string, opts *RequestOptions) (*TokenDayDataResponse, error) {
	req, err := constructByIdQuery(id, TokenDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, TokenDayDataResponse{}, c)
}

func (c *Client) ListTokenDayDatas(ctx context.Context, opts *RequestOptions) (*ListTokenDayDatasResponse, error) {
	req, err := constructListQuery(TokenDayDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListTokenDayDatasResponse{}, c)
}

func (c *Client) GetTokenHourDataById(ctx context.Context, id string, opts *RequestOptions) (*TokenHourDataResponse, error) {
	req, err := constructByIdQuery(id, TokenHourDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, TokenHourDataResponse{}, c)
}

func (c *Client) ListTokenHourDatas(ctx context.Context, opts *RequestOptions) (*ListTokenHourDatasResponse, error) {
	req, err := constructListQuery(TokenHourDataFields, opts)
	if err != nil {
		return nil, err
	}
	return executeRequestAndConvert(ctx, req, ListTokenHourDatasResponse{}, c)
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
