package helpers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/ethclient"
	Account "github.com/kaikoh95/web3go/src/account"
	"github.com/tyler-smith/go-bip39"
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

func IsValidEthAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	isValid := re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	fmt.Printf("is valid ETH address: %v\n", isValid)
	return isValid
}

// Checks if address is valid, 
// and also checks if it is a smart contract or account.
func IsContract(client *ethclient.Client, address string) bool {
	isValid := IsValidEthAddress(address)
	accountAddress := Account.GetAccountAddressFromHex(address)
	// nil is latest block
	bytecode, err := client.CodeAt(context.Background(), accountAddress, nil) 
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract)
	return isContract && isValid
}

// Generate a mnemonic, user-friendly seed
func GenerateMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}