package domain

import (
	"github.com/osmosis-labs/osmosis/osmomath"
)

// Token represents the token's domain model
type Token struct {
	// Name
	Name string `json:"name"`
	// HumanDenom is the human readable denom, e.g. atom
	HumanDenom string `json:"symbol"`
	// Denom is the chain denom of the token, e.g. ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2
	CoinMinimalDenom string `json:"coinMinimalDenom"`
	// Precision is the precision of the token.
	Precision int `json:"decimals"`
	// IsUnlisted is true if the token is unlisted.
	IsUnlisted  bool   `json:"preview"`
	CoingeckoID string `json:"coingeckoId"`
}

// PoolDenomMetaData contains the metadata about the denoms collected from the pools.
type PoolDenomMetaData struct {
	// TotalLiquidity represents the total liquidity across all pools.
	// @Type string
	TotalLiquidity osmomath.Int `json:"total_liquidity"`
	// TotalLiquidityCap represents the total liquidity capitalization across all pools.
	// If it is set to zero, that there was a failure in fetching the price.
	// @Type string
	TotalLiquidityCap osmomath.Int `json:"total_liquidity_cap"`
	// Price represents the price of the token.
	// @Type string
	Price osmomath.BigDec `json:"price"`
}

// DenomPoolLiquidityMap is a map of denoms to their pool liquidity data.
type DenomPoolLiquidityMap map[string]DenomPoolLiquidityData

// DenomPoolLiquidityData contains the pool liquidity data for a denom
// It has the total liquidity for the denom as well as all the
// pools with their individual contributions to the total.
type DenomPoolLiquidityData struct {
	// Total liquidity for this denom
	TotalLiquidity osmomath.Int
	// Mapping from pool ID to denom liquidity
	// in that pool
	Pools map[uint64]osmomath.Int
}

// GAMMSharePrefix is the prefix for the GAMM share
const GAMMSharePrefix = "gamm/pool"

// TokenRegistryLoader is loader of tokens from the chain registry.
// Loaded tokens are used to update the token registry.
type TokenRegistryLoader interface {
	// FetchAndUpdateTokens fetches tokens from the chain registry and updates the token registry.
	FetchAndUpdateTokens() error
}

// SwapMethod is the type of token swap method.
type TokenSwapMethod int

// SwapMethod constants represent specific type of swap method.
const (
	// Token swap method where the exact amount of token in is swapped for token out.
	TokenSwapMethodExactIn TokenSwapMethod = iota

	// Token swap method where the exact amount of token out is swapped for token in.
	TokenSwapMethodExactOut

	// Unknown swap method, used for error handling.
	TokenSwapMethodInvalid
)
