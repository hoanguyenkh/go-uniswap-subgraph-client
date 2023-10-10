package unigraphclient

// graphql types transcribed from https://github.com/Uniswap/v3-subgraph/blob/main/schema.graphql

type modelFields struct {
	name      string            // the name of the model
	direct    []string          // basic scalar types directly on the model e.g. Int, String
	reference map[string]string // fields that reference other models e.g. Token, Pool
	// TODO: add support for derived fields
}

var modelMap map[string]modelFields = map[string]modelFields{
	"factory":          FactoryFields,
	"pool":             PoolFields,
	"token":            TokenFields,
	"bundle":           BundleFields,
	"tick":             TickFields,
	"position":         PositionFields,
	"positionSnapshot": PositionSnapshotFields,
	"transaction":      TransactionFields,
	"mint":             MintFields,
	"burn":             BurnFields,
	"swap":             SwapFields,
	"collect":          CollectFields,
	"flash":            FlashFields,
	"uniswapDayData":   UniswapDayDataFields,
	"poolDayData":      PoolDayDataFields,
	"poolHourData":     PoolHourDataFields,
	"tickHourData":     TickHourDataFields,
	"tickDayData":      TickDayDataFields,
	"tokenDayData":     TokenDayDataFields,
	"tokenHourData":    TokenHourDataFields,
}

type FactoryResponse struct {
	Factory Factory
}

type ListFactoriesResponse struct {
	Factories []Factory
}

type Factory struct {
	ID                           string `json:"id"`
	PoolCount                    string `json:"poolCount"`
	TxCount                      string `json:"txCount"`
	TotalVolumeUSD               string `json:"totalVolumeUSD"`
	TotalVolumeETH               string `json:"totalVolumeETH"`
	TotalFeesUSD                 string `json:"totalFeesUSD"`
	TotalFeesETH                 string `json:"totalFeesETH"`
	UntrackedVolumeUSD           string `json:"untrackedVolumeUSD"`
	TotalValueLockedUSD          string `json:"totalValueLockedUSD"`
	TotalValueLockedETH          string `json:"totalValueLockedETH"`
	TotalValueLockedUSDUntracked string `json:"totalValueLockedUSDUntracked"`
	TotalValueLockedETHUntracked string `json:"totalValueLockedETHUntracked"`
	Owner                        string `json:"owner"`
}

var FactoryFields modelFields = modelFields{
	name: "factory",
	direct: []string{
		"id",                           // ID!
		"poolCount",                    // BigInt!
		"txCount",                      // BigInt!
		"totalVolumeUSD",               // BigDecimal!
		"totalVolumeETH",               // BigDecimal!
		"totalFeesUSD",                 // BigDecimal!
		"totalFeesETH",                 // BigDecimal!
		"untrackedVolumeUSD",           // BigDecimal!
		"totalValueLockedUSD",          // BigDecimal!
		"totalValueLockedETH",          // BigDecimal!
		"totalValueLockedUSDUntracked", // BigDecimal!
		"totalValueLockedETHUntracked", // BigDecimal!
		"owner",                        // ID!
	},
}

type PoolResponse struct {
	Pool Pool
}

type ListPoolsResponse struct {
	Pools []Pool
}

type Pool struct {
	ID                           string `json:"id"`
	CreatedAtTimestamp           string `json:"createdAtTimestamp"`
	CreatedAtBlockNumber         string `json:"createdAtBlockNumber"`
	Token0                       Token  `json:"token0"`
	Token1                       Token  `json:"token1"`
	FeeTier                      string `json:"feeTier"`
	Liquidity                    string `json:"liquidity"`
	SqrtPrice                    string `json:"sqrtPrice"`
	FeeGrowthGlobal0X128         string `json:"feeGrowthGlobal0X128"`
	FeeGrowthGlobal1X128         string `json:"feeGrowthGlobal1X128"`
	Token0Price                  string `json:"token0Price"`
	Token1Price                  string `json:"token1Price"`
	Tick                         string `json:"tick"`
	ObservationIndex             string `json:"observationIndex"`
	VolumeToken0                 string `json:"volumeToken0"`
	VolumeToken1                 string `json:"volumeToken1"`
	VolumeUSD                    string `json:"volumeUSD"`
	UntrackedVolumeUSD           string `json:"untrackedVolumeUSD"`
	FeesUSD                      string `json:"feesUSD"`
	TxCount                      string `json:"txCount"`
	CollectedFeesToken0          string `json:"collectedFeesToken0"`
	CollectedFeesToken1          string `json:"collectedFeesToken1"`
	CollectedFeesUSD             string `json:"collectedFeesUSD"`
	TotalValueLockedToken0       string `json:"totalValueLockedToken0"`
	TotalValueLockedToken1       string `json:"totalValueLockedToken1"`
	TotalValueLockedETH          string `json:"totalValueLockedETH"`
	TotalValueLockedUSD          string `json:"totalValueLockedUSD"`
	TotalValueLockedUSDUntracked string `json:"totalValueLockedUSDUntracked"`
	LiquidityProviderCount       string `json:"liquidityProviderCount"`
}

var PoolFields modelFields = modelFields{
	name: "pool",
	direct: []string{
		"id",                           // ID!
		"createdAtTimestamp",           // BigInt!
		"createdAtBlockNumber",         // BigInt!
		"feeTier",                      // BigInt!
		"liquidity",                    // BigInt!
		"sqrtPrice",                    // BigInt!
		"feeGrowthGlobal0X128",         // BigInt!
		"feeGrowthGlobal1X128",         // BigInt!
		"token0Price",                  // BigDecimal!
		"token1Price",                  // BigDecimal!
		"tick",                         // BigInt!
		"observationIndex",             // BigInt!
		"volumeToken0",                 // BigDecimal!
		"volumeToken1",                 // BigDecimal!
		"volumeUSD",                    // BigDecimal!
		"untrackedVolumeUSD",           // BigDecimal!
		"feesUSD",                      // BigDecimal!
		"txCount",                      // BigInt!
		"collectedFeesToken0",          // BigDecimal!
		"collectedFeesToken1",          // BigDecimal!
		"collectedFeesUSD",             // BigDecimal!
		"totalValueLockedToken0",       // BigDecimal!
		"totalValueLockedToken1",       // BigDecimal!
		"totalValueLockedETH",          // BigDecimal!
		"totalValueLockedUSD",          // BigDecimal!
		"totalValueLockedUSDUntracked", // BigDecimal!
		"liquidityProviderCount",       // BigInt!
	},
	reference: map[string]string{
		"token0": "token", // Token!
		"token1": "token", // Token!
	},
}

type TokenResponse struct {
	Token Token
}

type ListTokensResponse struct {
	Tokens []Token
}

type Token struct {
	ID                           string `json:"id"`
	Symbol                       string `json:"symbol"`
	Name                         string `json:"name"`
	Decimals                     string `json:"decimals"`
	TotalSupply                  string `json:"totalSupply"`
	Volume                       string `json:"volume"`
	VolumeUSD                    string `json:"volumeUSD"`
	UntrackedVolumeUSD           string `json:"untrackedVolumeUSD"`
	FeesUSD                      string `json:"feesUSD"`
	TxCount                      string `json:"txCount"`
	PoolCount                    string `json:"poolCount"`
	TotalValueLocked             string `json:"totalValueLocked"`
	TotalValueLockedUSD          string `json:"totalValueLockedUSD"`
	TotalValueLockedUSDUntracked string `json:"totalValueLockedUSDUntracked"`
	DerivedETH                   string `json:"derivedETH"`
}

var TokenFields modelFields = modelFields{
	name: "token",
	direct: []string{
		"id",                           // ID!
		"symbol",                       // String!
		"name",                         // String!
		"decimals",                     // BigInt!
		"totalSupply",                  // BigInt!
		"volume",                       // BigDecimal!
		"volumeUSD",                    // BigDecimal!
		"untrackedVolumeUSD",           // BigDecimal!
		"feesUSD",                      // BigDecimal!
		"txCount",                      // BigInt!
		"poolCount",                    // BigInt!
		"totalValueLocked",             // BigDecimal!
		"totalValueLockedUSD",          // BigDecimal!
		"totalValueLockedUSDUntracked", // BigDecimal!
		"derivedETH",                   // BigDecimal!
	},
	reference: map[string]string{
		"whitelistPools": "pool", // [Pool!]!
	},
}

type BundleResponse struct {
	Bundle Bundle
}

type ListBundlesResponse struct {
	Bundles []Bundle
}

type Bundle struct {
	ID          string `json:"id"`
	EthPriceUSD string `json:"ethPriceUSD"`
}

var BundleFields modelFields = modelFields{
	name: "bundle",
	direct: []string{
		"id",          // ID!
		"ethPriceUSD", // BigDecimal!
	},
}

type TickResponse struct {
	Tick Tick
}

type ListTicksResponse struct {
	Ticks []Tick
}

type Tick struct {
	ID                     string `json:"id"`
	PoolAddress            string `json:"poolAddress"`
	TickIdx                string `json:"tickIdx"`
	Pool                   Pool   `json:"pool"`
	LiquidityGross         string `json:"liquidityGross"`
	LiquidityNet           string `json:"liquidityNet"`
	Price0                 string `json:"price0"`
	Price1                 string `json:"price1"`
	VolumeToken0           string `json:"volumeToken0"`
	VolumeToken1           string `json:"volumeToken1"`
	VolumeUSD              string `json:"volumeUSD"`
	UntrackedVolumeUSD     string `json:"untrackedVolumeUSD"`
	FeesUSD                string `json:"feesUSD"`
	CollectedFeesToken0    string `json:"collectedFeesToken0"`
	CollectedFeesToken1    string `json:"collectedFeesToken1"`
	CollectedFeesUSD       string `json:"collectedFeesUSD"`
	CreatedAtTimestamp     string `json:"createdAtTimestamp"`
	CreatedAtBlockNumber   string `json:"createdAtBlockNumber"`
	LiquidityProviderCount string `json:"liquidityProviderCount"`
	FeeGrowthOutside0X128  string `json:"feeGrowthOutside0X128"`
	FeeGrowthOutside1X128  string `json:"feeGrowthOutside1X128"`
}

var TickFields modelFields = modelFields{
	name: "tick",
	direct: []string{
		"id",                     // ID!
		"poolAddress",            // String
		"tickIdx",                // BigInt!
		"liquidityGross",         // BigInt!
		"liquidityNet",           // BigInt!
		"price0",                 // BigDecimal!
		"price1",                 // BigDecimal!
		"volumeToken0",           // BigDecimal!
		"volumeToken1",           // BigDecimal!
		"volumeUSD",              // BigDecimal!
		"untrackedVolumeUSD",     // BigDecimal!
		"feesUSD",                // BigDecimal!
		"collectedFeesToken0",    // BigDecimal!
		"collectedFeesToken1",    // BigDecimal!
		"collectedFeesUSD",       // BigDecimal!
		"createdAtTimestamp",     // BigInt!
		"createdAtBlockNumber",   // BigInt!
		"liquidityProviderCount", // BigInt!
		"feeGrowthOutside0X128",  // BigInt!
		"feeGrowthOutside1X128",  // BigInt!
	},
	reference: map[string]string{
		"pool": "pool", // Pool!
	},
}

type PositionResponse struct {
	Position Position
}

type ListPositionsResponse struct {
	Positions []Position
}

type Position struct {
	ID                       string      `json:"id"`
	Owner                    string      `json:"owner"`
	Pool                     Pool        `json:"pool"`
	Token0                   Token       `json:"token0"`
	Token1                   Token       `json:"token1"`
	TickLower                Tick        `json:"tickLower"`
	TickUpper                Tick        `json:"tickUpper"`
	Liquidity                string      `json:"liquidity"`
	DepositedToken0          string      `json:"depositedToken0"`
	DepositedToken1          string      `json:"depositedToken1"`
	WithdrawnToken0          string      `json:"withdrawnToken0"`
	WithdrawnToken1          string      `json:"withdrawnToken1"`
	CollectedFeesToken0      string      `json:"collectedFeesToken0"`
	CollectedFeesToken1      string      `json:"collectedFeesToken1"`
	Transaction              Transaction `json:"transaction"`
	FeeGrowthInside0LastX128 string      `json:"feeGrowthInside0LastX128"`
	FeeGrowthInside1LastX128 string      `json:"feeGrowthInside1LastX128"`
}

var PositionFields modelFields = modelFields{
	name: "position",
	direct: []string{
		"id",                       // ID!
		"owner",                    // Bytes!
		"liquidity",                // BigInt!
		"depositedToken0",          // BigDecimal!
		"depositedToken1",          // BigDecimal!
		"withdrawnToken0",          // BigDecimal!
		"withdrawnToken1",          // BigDecimal!
		"collectedFeesToken0",      // BigDecimal!
		"collectedFeesToken1",      // BigDecimal!
		"feeGrowthInside0LastX128", // BigInt!
		"feeGrowthInside1LastX128", // BigInt!
	},
	reference: map[string]string{
		"pool":        "pool",        // Pool!
		"token0":      "token",       // Token!
		"token1":      "token",       // Token!
		"tickLower":   "tick",        // Tick!
		"tickUpper":   "tick",        // Tick!
		"transaction": "transaction", // Transaction!
	},
}

type PositionSnapshotResponse struct {
	PositionSnapshot PositionSnapshot
}

type ListPositionSnapshotsResponse struct {
	PositionSnapshots []PositionSnapshot
}

type PositionSnapshot struct {
	ID                       string      `json:"id"`
	Owner                    string      `json:"owner"`
	Pool                     Pool        `json:"pool"`
	Position                 Position    `json:"position"`
	BlockNumber              string      `json:"blockNumber"`
	Timestamp                string      `json:"timestamp"`
	Liquidity                string      `json:"liquidity"`
	DepositedToken0          string      `json:"depositedToken0"`
	DepositedToken1          string      `json:"depositedToken1"`
	WithdrawnToken0          string      `json:"withdrawnToken0"`
	WithdrawnToken1          string      `json:"withdrawnToken1"`
	CollectedFeesToken0      string      `json:"collectedFeesToken0"`
	CollectedFeesToken1      string      `json:"collectedFeesToken1"`
	Transaction              Transaction `json:"transaction"`
	FeeGrowthInside0LastX128 string      `json:"feeGrowthInside0LastX128"`
	FeeGrowthInside1LastX128 string      `json:"feeGrowthInside1LastX128"`
}

var PositionSnapshotFields modelFields = modelFields{
	name: "positionSnapshot",
	direct: []string{
		"id",                       // ID!
		"owner",                    // Bytes!
		"blockNumber",              // BigInt!
		"timestamp",                // BigInt!
		"liquidity",                // BigInt!
		"depositedToken0",          // BigDecimal!
		"depositedToken1",          // BigDecimal!
		"withdrawnToken0",          // BigDecimal!
		"withdrawnToken1",          // BigDecimal!
		"collectedFeesToken0",      // BigDecimal!
		"collectedFeesToken1",      // BigDecimal!
		"feeGrowthInside0LastX128", // BigInt!
		"feeGrowthInside1LastX128", // BigInt!
	},
	reference: map[string]string{
		"pool":        "pool",        // Pool!
		"position":    "position",    // Position!
		"transaction": "transaction", // Transaction!
	},
}

type TransactionResponse struct {
	Transaction Transaction
}

type ListTransactionsResponse struct {
	Transactions []Transaction
}

type Transaction struct {
	ID          string `json:"id"`
	BlockNumber string `json:"blockNumber"`
	Timestamp   string `json:"timestamp"`
	GasUsed     string `json:"gasUsed"`
	GasPrice    string `json:"gasPrice"`
}

var TransactionFields modelFields = modelFields{
	name: "transaction",
	direct: []string{
		"id",          // ID!
		"blockNumber", // BigInt!
		"timestamp",   // BigInt!
		"gasUsed",     // BigInt!
		"gasPrice",    // BigInt!
	},
}

type MintResponse struct {
	Mint Mint
}

type ListMintsResponse struct {
	Mints []Mint
}

type Mint struct {
	ID          string      `json:"id"`
	Transaction Transaction `json:"transaction"`
	Timestamp   string      `json:"timestamp"`
	Pool        Pool        `json:"pool"`
	Token0      Token       `json:"token0"`
	Token1      Token       `json:"token1"`
	Owner       string      `json:"owner"`
	Sender      string      `json:"sender"`
	Origin      string      `json:"origin"`
	Amount      string      `json:"amount"`
	Amount0     string      `json:"amount0"`
	Amount1     string      `json:"amount1"`
	AmountUSD   string      `json:"amountUSD"`
	TickLower   string      `json:"tickLower"`
	TickUpper   string      `json:"tickUpper"`
	LogIndex    string      `json:"logIndex"`
}

var MintFields modelFields = modelFields{
	name: "mint",
	direct: []string{
		"id",        // ID!
		"timestamp", // BigInt!
		"owner",     // Bytes!
		"sender",    // Bytes!
		"origin",    // Bytes!
		"amount",    // BigInt!
		"amount0",   // BigDecimal!
		"amount1",   // BigDecimal!
		"amountUSD", // BigDecimal!
		"tickLower", // BigInt!
		"tickUpper", // BigInt!
		"logIndex",  // BigInt
	},
	reference: map[string]string{
		"transaction": "transaction", // Transaction!
		"pool":        "pool",        // Pool!
		"token0":      "token",       // Token!
		"token1":      "token",       // Token!
	},
}

type BurnResponse struct {
	Burn Burn
}

type ListBurnsResponse struct {
	Burns []Burn
}

type Burn struct {
	ID          string      `json:"id"`
	Transaction Transaction `json:"transaction"`
	Pool        Pool        `json:"pool"`
	Token0      Token       `json:"token0"`
	Token1      Token       `json:"token1"`
	Timestamp   string      `json:"timestamp"`
	Owner       string      `json:"owner"`
	Origin      string      `json:"origin"`
	Amount      string      `json:"amount"`
	Amount0     string      `json:"amount0"`
	Amount1     string      `json:"amount1"`
	AmountUSD   string      `json:"amountUSD"`
	TickLower   string      `json:"tickLower"`
	TickUpper   string      `json:"tickUpper"`
	LogIndex    string      `json:"logIndex"`
}

var BurnFields modelFields = modelFields{
	name: "burn",
	direct: []string{
		"id",        // ID!
		"timestamp", // BigInt!
		"owner",     // Bytes!
		"origin",    // Bytes!
		"amount",    // BigInt!
		"amount0",   // BigDecimal!
		"amount1",   // BigDecimal!
		"amountUSD", // BigDecimal!
		"tickLower", // BigInt!
		"tickUpper", // BigInt!
		"logIndex",  // BigInt
	},
	reference: map[string]string{
		"transaction": "transaction", // Transaction!
		"pool":        "pool",        // Pool!
		"token0":      "token",       // Token!
		"token1":      "token",       // Token!
	},
}

type SwapResponse struct {
	Swap Swap
}

type ListSwapsResponse struct {
	Swaps []Swap
}

type Swap struct {
	ID           string      `json:"id"`
	Transaction  Transaction `json:"transaction"`
	Timestamp    string      `json:"timestamp"`
	Pool         Pool        `json:"pool"`
	Token0       Token       `json:"token0"`
	Token1       Token       `json:"token1"`
	Sender       string      `json:"sender"`
	Recipient    string      `json:"recipient"`
	Origin       string      `json:"origin"`
	Amount0      string      `json:"amount0"`
	Amount1      string      `json:"amount1"`
	AmountUSD    string      `json:"amountUSD"`
	SqrtPriceX96 string      `json:"sqrtPriceX96"`
	Tick         string      `json:"tick"`
	LogIndex     string      `json:"logIndex"`
}

var SwapFields modelFields = modelFields{
	name: "swap",
	direct: []string{
		"id",           // ID!
		"timestamp",    // BigInt!
		"sender",       // Bytes!
		"recipient",    // Bytes!
		"origin",       // Bytes!
		"amount0",      // BigDecimal!
		"amount1",      // BigDecimal!
		"amountUSD",    // BigDecimal!
		"sqrtPriceX96", // BigInt!
		"tick",         // BigInt!
		"logIndex",     // BigInt
	},
	reference: map[string]string{
		"transaction": "transaction", // Transaction!
		"pool":        "pool",        // Pool!
		"token0":      "token",       // Token!
		"token1":      "token",       // Token!
	},
}

type CollectResponse struct {
	Collect Collect
}

type ListCollectsResponse struct {
	Collects []Collect
}

type Collect struct {
	ID          string      `json:"id"`
	Transaction Transaction `json:"transaction"`
	Timestamp   string      `json:"timestamp"`
	Pool        Pool        `json:"pool"`
	Owner       string      `json:"owner"`
	Amount0     string      `json:"amount0"`
	Amount1     string      `json:"amount1"`
	AmountUSD   string      `json:"amountUSD"`
	TickLower   string      `json:"tickLower"`
	TickUpper   string      `json:"tickUpper"`
	LogIndex    string      `json:"logIndex"`
}

var CollectFields modelFields = modelFields{
	name: "collect",
	direct: []string{
		"id",        // ID!
		"timestamp", // BigInt!
		"owner",     // Bytes
		"amount0",   // BigDecimal!
		"amount1",   // BigDecimal!
		"amountUSD", // BigDecimal
		"tickLower", // BigInt!
		"tickUpper", // BigInt!
		"logIndex",  // BigInt
	},
	reference: map[string]string{
		"transaction": "transaction", // Transaction!
		"pool":        "pool",        // Pool!
	},
}

type FlashResponse struct {
	Flash Flash
}

type ListFlashesResponse struct {
	Flashes []Flash
}

type Flash struct {
	ID          string      `json:"id"`
	Transaction Transaction `json:"transaction"`
	Timestamp   string      `json:"timestamp"`
	Pool        Pool        `json:"pool"`
	Sender      string      `json:"sender"`
	Recipient   string      `json:"recipient"`
	Amount0     string      `json:"amount0"`
	Amount1     string      `json:"amount1"`
	AmountUSD   string      `json:"amountUSD"`
	Amount0Paid string      `json:"amount0Paid"`
	Amount1Paid string      `json:"amount1Paid"`
	LogIndex    string      `json:"logIndex"`
}

var FlashFields modelFields = modelFields{
	name: "flash",
	direct: []string{
		"id",          // ID!
		"timestamp",   // BigInt!
		"sender",      // Bytes!
		"recipient",   // Bytes!
		"amount0",     // BigDecimal!
		"amount1",     // BigDecimal!
		"amountUSD",   // BigDecimal
		"amount0Paid", // BigDecimal!
		"amount1Paid", // BigDecimal!
		"logIndex",    // BigInt
	},
	reference: map[string]string{
		"transaction": "transaction", // Transaction!
		"pool":        "pool",        // Pool!
	},
}

type UniswapDayDataResponse struct {
	UniswapDayData UniswapDayData
}

type ListUniswapDayDatasResponse struct {
	UniswapDayDatas []UniswapDayData
}

type UniswapDayData struct {
	ID                 string `json:"id"`
	Date               string `json:"date"`
	VolumeETH          string `json:"volumeETH"`
	VolumeUSD          string `json:"volumeUSD"`
	VolumeUSDUntracked string `json:"volumeUSDUntracked"`
	FeesUSD            string `json:"feesUSD"`
	TxCount            string `json:"txCount"`
	TvlUSD             string `json:"tvlUSD"`
}

var UniswapDayDataFields modelFields = modelFields{
	name: "uniswapDayData",
	direct: []string{
		"id",                 // ID!
		"date",               // Int!
		"volumeETH",          // BigDecimal!
		"volumeUSD",          // BigDecimal!
		"volumeUSDUntracked", // BigDecimal!
		"feesUSD",            // BigDecimal!
		"txCount",            // BigInt!
		"tvlUSD",             // BigDecimal!
	},
}

type PoolDayDataResponse struct {
	PoolDayData PoolDayData
}

type ListPoolDayDatasResponse struct {
	PoolDayDatas []PoolDayData
}

type PoolDayData struct {
	ID                   string `json:"id"`
	Date                 string `json:"date"`
	Pool                 Pool   `json:"pool"`
	Liquidity            string `json:"liquidity"`
	SqrtPrice            string `json:"sqrtPrice"`
	Token0Price          string `json:"token0Price"`
	Token1Price          string `json:"token1Price"`
	Tick                 string `json:"tick"`
	FeeGrowthGlobal0X128 string `json:"feeGrowthGlobal0X128"`
	FeeGrowthGlobal1X128 string `json:"feeGrowthGlobal1X128"`
	TvlUSD               string `json:"tvlUSD"`
	VolumeToken0         string `json:"volumeToken0"`
	VolumeToken1         string `json:"volumeToken1"`
	VolumeUSD            string `json:"volumeUSD"`
	FeesUSD              string `json:"feesUSD"`
	TxCount              string `json:"txCount"`
	Open                 string `json:"open"`
	High                 string `json:"high"`
	Low                  string `json:"low"`
	Close                string `json:"close"`
}

var PoolDayDataFields modelFields = modelFields{
	name: "poolDayData",
	direct: []string{
		"id",                   // ID!
		"date",                 // Int!
		"liquidity",            // BigInt!
		"sqrtPrice",            // BigInt!
		"token0Price",          // BigDecimal!
		"token1Price",          // BigDecimal!
		"tick",                 // BigInt
		"feeGrowthGlobal0X128", // BigInt!
		"feeGrowthGlobal1X128", // BigInt!
		"tvlUSD",               // BigDecimal!
		"volumeToken0",         // BigDecimal!
		"volumeToken1",         // BigDecimal!
		"volumeUSD",            // BigDecimal!
		"feesUSD",              // BigDecimal!
		"txCount",              // BigInt!
		"open",                 // BigDecimal!
		"high",                 // BigDecimal!
		"low",                  // BigDecimal!
		"close",                // BigDecimal!
	},
	reference: map[string]string{
		"pool": "pool", // Pool!
	},
}

type PoolHourDataResponse struct {
	PoolHourData PoolHourData
}

type ListPoolHourDatasResponse struct {
	PoolHourDatas []PoolHourData
}

type PoolHourData struct {
	ID                   string `json:"id"`
	PeriodStartUnix      string `json:"periodStartUnix"`
	Pool                 Pool   `json:"pool"`
	Liquidity            string `json:"liquidity"`
	SqrtPrice            string `json:"sqrtPrice"`
	Token0Price          string `json:"token0Price"`
	Token1Price          string `json:"token1Price"`
	Tick                 string `json:"tick"`
	FeeGrowthGlobal0X128 string `json:"feeGrowthGlobal0X128"`
	FeeGrowthGlobal1X128 string `json:"feeGrowthGlobal1X128"`
	TvlUSD               string `json:"tvlUSD"`
	VolumeToken0         string `json:"volumeToken0"`
	VolumeToken1         string `json:"volumeToken1"`
	VolumeUSD            string `json:"volumeUSD"`
	FeesUSD              string `json:"feesUSD"`
	TxCount              string `json:"txCount"`
	Open                 string `json:"open"`
	High                 string `json:"high"`
	Low                  string `json:"low"`
	Close                string `json:"close"`
}

var PoolHourDataFields modelFields = modelFields{
	name: "poolHourData",
	direct: []string{
		"id",                   // ID!
		"periodStartUnix",      // Int!
		"liquidity",            // BigInt!
		"sqrtPrice",            // BigInt!
		"token0Price",          // BigDecimal!
		"token1Price",          // BigDecimal!
		"tick",                 // BigInt
		"feeGrowthGlobal0X128", // BigInt!
		"feeGrowthGlobal1X128", // BigInt!
		"tvlUSD",               // BigDecimal!
		"volumeToken0",         // BigDecimal!
		"volumeToken1",         // BigDecimal!
		"volumeUSD",            // BigDecimal!
		"feesUSD",              // BigDecimal!
		"txCount",              // BigInt!
		"open",                 // BigDecimal!
		"high",                 // BigDecimal!
		"low",                  // BigDecimal!
		"close",                // BigDecimal!
	},
	reference: map[string]string{
		"pool": "pool", // Pool!
	},
}

type TickHourDataResponse struct {
	TickHourData TickHourData
}

type ListTickHourDatasResponse struct {
	TickHourDatas []TickHourData
}

type TickHourData struct {
	ID              string `json:"id"`
	PeriodStartUnix string `json:"periodStartUnix"`
	Pool            Pool   `json:"pool"`
	Tick            Tick   `json:"tick"`
	LiquidityGross  string `json:"liquidityGross"`
	LiquidityNet    string `json:"liquidityNet"`
	VolumeToken0    string `json:"volumeToken0"`
	VolumeToken1    string `json:"volumeToken1"`
	VolumeUSD       string `json:"volumeUSD"`
	FeesUSD         string `json:"feesUSD"`
}

var TickHourDataFields modelFields = modelFields{
	name: "tickHourData",
	direct: []string{
		"id",              // ID!
		"periodStartUnix", // Int!
		"liquidityGross",  // BigInt!
		"liquidityNet",    // BigInt!
		"volumeToken0",    // BigDecimal!
		"volumeToken1",    // BigDecimal!
		"volumeUSD",       // BigDecimal!
		"feesUSD",         // BigDecimal!
	},
	reference: map[string]string{
		"pool": "pool", // Pool!
		"tick": "tick", // Tick!
	},
}

type TickDayDataResponse struct {
	TickDayData TickDayData
}

type ListTickDayDatasResponse struct {
	TickDayDatas []TickDayData
}

type TickDayData struct {
	ID                    string `json:"id"`
	Date                  string `json:"date"`
	Pool                  Pool   `json:"pool"`
	Tick                  Tick   `json:"tick"`
	LiquidityGross        string `json:"liquidityGross"`
	LiquidityNet          string `json:"liquidityNet"`
	VolumeToken0          string `json:"volumeToken0"`
	VolumeToken1          string `json:"volumeToken1"`
	VolumeUSD             string `json:"volumeUSD"`
	FeesUSD               string `json:"feesUSD"`
	FeeGrowthOutside0X128 string `json:"feeGrowthOutside0X128"`
	FeeGrowthOutside1X128 string `json:"feeGrowthOutside1X128"`
}

var TickDayDataFields modelFields = modelFields{
	name: "tickDayData",
	direct: []string{
		"id",                    // ID!
		"date",                  // Int!
		"liquidityGross",        // BigInt!
		"liquidityNet",          // BigInt!
		"volumeToken0",          // BigDecimal!
		"volumeToken1",          // BigDecimal!
		"volumeUSD",             // BigDecimal!
		"feesUSD",               // BigDecimal!
		"feeGrowthOutside0X128", // BigInt!
		"feeGrowthOutside1X128", // BigInt!
	},
	reference: map[string]string{
		"pool": "pool", // Pool!
		"tick": "tick", // Tick!
	},
}

type TokenDayDataResponse struct {
	TokenDayData TokenDayData
}

type ListTokenDayDatasResponse struct {
	TokenDayDatas []TokenDayData
}

type TokenDayData struct {
	ID                  string `json:"id"`
	Date                string `json:"date"`
	Token               Token  `json:"token"`
	Volume              string `json:"volume"`
	VolumeUSD           string `json:"volumeUSD"`
	UntrackedVolumeUSD  string `json:"untrackedVolumeUSD"`
	TotalValueLocked    string `json:"totalValueLocked"`
	TotalValueLockedUSD string `json:"totalValueLockedUSD"`
	PriceUSD            string `json:"priceUSD"`
	FeesUSD             string `json:"feesUSD"`
	Open                string `json:"open"`
	High                string `json:"high"`
	Low                 string `json:"low"`
	Close               string `json:"close"`
}

var TokenDayDataFields modelFields = modelFields{
	name: "tokenDayData",
	direct: []string{
		"id",                  // ID!
		"date",                // Int!
		"volume",              // BigDecimal!
		"volumeUSD",           // BigDecimal!
		"untrackedVolumeUSD",  // BigDecimal!
		"totalValueLocked",    // BigDecimal!
		"totalValueLockedUSD", // BigDecimal!
		"priceUSD",            // BigDecimal!
		"feesUSD",             // BigDecimal!
		"open",                // BigDecimal!
		"high",                // BigDecimal!
		"low",                 // BigDecimal!
		"close",               // BigDecimal!
	},
	reference: map[string]string{
		"token": "token", // Token!
	},
}

type TokenHourDataResponse struct {
	TokenHourData TokenHourData
}

type ListTokenHourDatasResponse struct {
	TokenHourDatas []TokenHourData
}

type TokenHourData struct {
	ID                  string `json:"id"`
	PeriodStartUnix     string `json:"periodStartUnix"`
	Token               Token  `json:"token"`
	Volume              string `json:"volume"`
	VolumeUSD           string `json:"volumeUSD"`
	UntrackedVolumeUSD  string `json:"untrackedVolumeUSD"`
	TotalValueLocked    string `json:"totalValueLocked"`
	TotalValueLockedUSD string `json:"totalValueLockedUSD"`
	PriceUSD            string `json:"priceUSD"`
	FeesUSD             string `json:"feesUSD"`
	Open                string `json:"open"`
	High                string `json:"high"`
	Low                 string `json:"low"`
	Close               string `json:"close"`
}

var TokenHourDataFields modelFields = modelFields{
	name: "tokenHourData",
	direct: []string{
		"id",                  // ID!
		"periodStartUnix",     // Int!
		"volume",              // BigDecimal!
		"volumeUSD",           // BigDecimal!
		"untrackedVolumeUSD",  // BigDecimal!
		"totalValueLocked",    // BigDecimal!
		"totalValueLockedUSD", // BigDecimal!
		"priceUSD",            // BigDecimal!
		"feesUSD",             // BigDecimal!
		"open",                // BigDecimal!
		"high",                // BigDecimal!
		"low",                 // BigDecimal!
		"close",               // BigDecimal!
	},
	reference: map[string]string{
		"token": "token", // Token!
	},
}
