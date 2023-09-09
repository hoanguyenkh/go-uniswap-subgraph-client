package unigraphclient

type PoolResult struct {
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
	"id",
	"createdAtTimestamp",
	"createdAtBlockNumber",
	// "token0",
	// "token1",
	"feeTier",
	"liquidity",
	"sqrtPrice",
	"feeGrowthGlobal0X128",
	"feeGrowthGlobal1X128",
	"token0Price",
	"token1Price",
	"tick",
	"observationIndex",
	"volumeToken0",
	"volumeToken1",
	"volumeUSD",
	"untrackedVolumeUSD",
	"feesUSD",
	"txCount",
	"collectedFeesToken0",
	"collectedFeesToken1",
	"collectedFeesUSD",
	"totalValueLockedToken0",
	"totalValueLockedToken1",
	"totalValueLockedETH",
	"totalValueLockedUSD",
	"totalValueLockedUSDUntracked",
	"liquidityProviderCount",
	// "poolHourData",
	// "poolDayData",
	// "mints",
	// "burns",
	// "swaps",
	// "collects",
	// "ticks",
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

var TokenFields []string = []string{
	"id",
	"symbol",
	"name",
	"decimals",
	"totalSupply",
	"volume",
	"volumeUSD",
	"untrackedVolumeUSD",
	"feesUSD",
	"txCount",
	"poolCount",
	"totalValueLocked",
	"totalValueLockedUSD",
	"totalValueLockedUSDUntracked",
	"derivedETH",
	// "whitelistPools",
	// "tokenDayData",
}

// type Pool struct {
// 	ID        any
// 	Tick      int
// 	FeeTier   int
// 	SqrtPrice int
// 	Liquidity int
// 	Token0    Token
// 	Token1    Token
// }

// type Token struct {
// 	ID       any
// 	Symbol   string
// 	Decimals int
// }
