package helpers

import (
	"io/ioutil"
	"log"
	"math"
	"math/big"
)

// 18 decimal points precision
func ConvertToWei(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

func GetBytesFromFile(file string) []byte {
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("error reading file", err)
	}
	return jsonBytes
}