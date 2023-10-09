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
}

type FactoryResponse struct {
	Factory Factory
}

type ListFactoriesResponse struct {
	Factories []Factory
}

type Factory struct {
	ID                           string
	PoolCount                    string
	TxCount                      string
	TotalVolumeUSD               string
	TotalVolumeETH               string
	TotalFeesUSD                 string
	TotalFeesETH                 string
	UntrackedVolumeUSD           string
	TotalValueLockedUSD          string
	TotalValueLockedETH          string
	TotalValueLockedUSDUntracked string
	TotalValueLockedETHUntracked string
	Owner                        string
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
	ID                           string
	CreatedAtTimestamp           string
	CreatedAtBlockNumber         string
	Token0                       Token
	Token1                       Token
	FeeTier                      string
	Liquidity                    string
	SqrtPrice                    string
	FeeGrowthGlobal0X128         string
	FeeGrowthGlobal1X128         string
	Token0Price                  string
	Token1Price                  string
	Tick                         string
	ObservationIndex             string
	VolumeToken0                 string
	VolumeToken1                 string
	VolumeUSD                    string
	UntrackedVolumeUSD           string
	FeesUSD                      string
	TxCount                      string
	CollectedFeesToken0          string
	CollectedFeesToken1          string
	CollectedFeesUSD             string
	TotalValueLockedToken0       string
	TotalValueLockedToken1       string
	TotalValueLockedETH          string
	TotalValueLockedUSD          string
	TotalValueLockedUSDUntracked string
	LiquidityProviderCount       string
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
	ID                           string
	Symbol                       string
	Name                         string
	Decimals                     string
	TotalSupply                  string
	Volume                       string
	VolumeUSD                    string
	UntrackedVolumeUSD           string
	FeesUSD                      string
	TxCount                      string
	PoolCount                    string
	TotalValueLocked             string
	TotalValueLockedUSD          string
	TotalValueLockedUSDUntracked string
	DerivedETH                   string
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
	ID          string
	EthPriceUSD string
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
	ID                     string
	PoolAddress            string
	TickIdx                string
	Pool                   Pool
	LiquidityGross         string
	LiquidityNet           string
	Price0                 string
	Price1                 string
	VolumeToken0           string
	VolumeToken1           string
	VolumeUSD              string
	UntrackedVolumeUSD     string
	FeesUSD                string
	CollectedFeesToken0    string
	CollectedFeesToken1    string
	CollectedFeesUSD       string
	CreatedAtTimestamp     string
	CreatedAtBlockNumber   string
	LiquidityProviderCount string
	FeeGrowthOutside0X128  string
	FeeGrowthOutside1X128  string
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
	ID                       string
	Owner                    string
	Pool                     Pool
	Token0                   Token
	Token1                   Token
	TickLower                Tick
	TickUpper                Tick
	Liquidity                string
	DepositedToken0          string
	DepositedToken1          string
	WithdrawnToken0          string
	WithdrawnToken1          string
	CollectedFeesToken0      string
	CollectedFeesToken1      string
	Transaction              Transaction
	FeeGrowthInside0LastX128 string
	FeeGrowthInside1LastX128 string
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
	ID                       string
	Owner                    string
	Pool                     Pool
	Position                 Position
	BlockNumber              string
	Timestamp                string
	Liquidity                string
	DepositedToken0          string
	DepositedToken1          string
	WithdrawnToken0          string
	WithdrawnToken1          string
	CollectedFeesToken0      string
	CollectedFeesToken1      string
	Transaction              Transaction
	FeeGrowthInside0LastX128 string
	FeeGrowthInside1LastX128 string
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
	ID          string
	BlockNumber string
	Timestamp   string
	GasUsed     string
	GasPrice    string
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
	ID          string
	Transaction Transaction
	Timestamp   string
	Pool        Pool
	Token0      Token
	Token1      Token
	Owner       string
	Sender      string
	Origin      string
	Amount      string
	Amount0     string
	Amount1     string
	AmountUSD   string
	TickLower   string
	TickUpper   string
	LogIndex    string
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
	ID          string
	Transaction Transaction
	Pool        Pool
	Token0      Token
	Token1      Token
	Timestamp   string
	Owner       string
	Origin      string
	Amount      string
	Amount0     string
	Amount1     string
	AmountUSD   string
	TickLower   string
	TickUpper   string
	LogIndex    string
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
	ID           string
	Transaction  Transaction
	Timestamp    string
	Pool         Pool
	Token0       Token
	Token1       Token
	Sender       string
	Recipient    string
	Origin       string
	Amount0      string
	Amount1      string
	AmountUSD    string
	SqrtPriceX96 string
	Tick         string
	LogIndex     string
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
	ID          string
	Transaction Transaction
	Timestamp   string
	Pool        Pool
	Owner       string
	Amount0     string
	Amount1     string
	AmountUSD   string
	TickLower   string
	TickUpper   string
	LogIndex    string
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
	ID          string
	Transaction Transaction
	Timestamp   string
	Pool        Pool
	Sender      string
	Recipient   string
	Amount0     string
	Amount1     string
	AmountUSD   string
	Amount0Paid string
	Amount1Paid string
	LogIndex    string
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
	ID                 string
	Date               string
	VolumeETH          string
	VolumeUSD          string
	VolumeUSDUntracked string
	FeesUSD            string
	TxCount            string
	TvlUSD             string
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
	ID                   string
	Date                 string
	Pool                 Pool
	Liquidity            string
	SqrtPrice            string
	Token0Price          string
	Token1Price          string
	Tick                 string
	FeeGrowthGlobal0X128 string
	FeeGrowthGlobal1X128 string
	TvlUSD               string
	VolumeToken0         string
	VolumeToken1         string
	VolumeUSD            string
	FeesUSD              string
	TxCount              string
	Open                 string
	High                 string
	Low                  string
	Close                string
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
	ID                   string
	PeriodStartUnix      string
	Pool                 Pool
	Liquidity            string
	SqrtPrice            string
	Token0Price          string
	Token1Price          string
	Tick                 string
	FeeGrowthGlobal0X128 string
	FeeGrowthGlobal1X128 string
	TvlUSD               string
	VolumeToken0         string
	VolumeToken1         string
	VolumeUSD            string
	FeesUSD              string
	TxCount              string
	Open                 string
	High                 string
	Low                  string
	Close                string
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
	ID              string
	PeriodStartUnix string
	Pool            Pool
	Tick            Tick
	LiquidityGross  string
	LiquidityNet    string
	VolumeToken0    string
	VolumeToken1    string
	VolumeUSD       string
	FeesUSD         string
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
	ID                    string
	Date                  string
	Pool                  Pool
	Tick                  Tick
	LiquidityGross        string
	LiquidityNet          string
	VolumeToken0          string
	VolumeToken1          string
	VolumeUSD             string
	FeesUSD               string
	FeeGrowthOutside0X128 string
	FeeGrowthOutside1X128 string
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
	ID                  string
	Date                string
	Token               Token
	Volume              string
	VolumeUSD           string
	UntrackedVolumeUSD  string
	TotalValueLocked    string
	TotalValueLockedUSD string
	PriceUSD            string
	FeesUSD             string
	Open                string
	High                string
	Low                 string
	Close               string
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
	ID                  string
	PeriodStartUnix     string
	Token               Token
	Volume              string
	VolumeUSD           string
	UntrackedVolumeUSD  string
	TotalValueLocked    string
	TotalValueLockedUSD string
	PriceUSD            string
	FeesUSD             string
	Open                string
	High                string
	Low                 string
	Close               string
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
