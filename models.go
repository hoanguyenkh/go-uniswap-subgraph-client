package unigraphclient

type FactoryResponse struct {
	Factory Factory
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

var FactoryFields []string = []string{
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
}

type PoolResponse struct {
	Pool Pool
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
	// other fields
}

var PoolFields []string = []string{
	"id",                   // ID!
	"createdAtTimestamp",   // BigInt!
	"createdAtBlockNumber", // BigInt!
	// "token0", // Token!
	// "token1", // Token!
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
	// "poolHourData", // [PoolHourData!]! @derivedFrom(field: "pool")
	// "poolDayData", // [PoolDayData!]! @derivedFrom(field: "pool")
	// "mints", // [Mint!]! @derivedFrom(field: "pool")
	// "burns", // [Burn!]! @derivedFrom(field: "pool")
	// "swaps", // [Swap!]! @derivedFrom(field: "pool")
	// "collects", // [Collect!]! @derivedFrom(field: "pool")
	// "ticks", // [Tick!]! @derivedFrom(field: "pool")
}

type TokenResponse struct {
	Token Token
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
	// other fields
}

var TokenFields []string = []string{
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
	// "whitelistPools", // [Pool!]!
	// "tokenDayData", // [TokenDayData!]! @derivedFrom(field: "token")
}
