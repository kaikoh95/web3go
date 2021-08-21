package helpers

import (
	"math"
	"math/big"
)

// 18 decimal points precision
func ConvertToWei(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}