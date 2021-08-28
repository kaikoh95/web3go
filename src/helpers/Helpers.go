package helpers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kaikoh95/web3go/src/Account"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/shopspring/decimal"
)

// 18 decimal points precision
// unused
// func ConvertFromWei(balance *big.Int) *big.Float {
// 	fbalance := new(big.Float)
// 	fbalance.SetString(balance.String())
// 	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
// }

func SuggestGasPrice(client *ethclient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Unable to suggest gas price ", err)
	}
	return gasPrice
}

func EstimateGasLimit(client *ethclient.Client, toAddress *common.Address, data []byte) uint64 {
    gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
    return gasLimit
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
    value := new(big.Int)
    switch v := ivalue.(type) {
    case string:
        value.SetString(v, 10)
    case *big.Int:
        value = v
    }

    mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
    num, _ := decimal.NewFromString(value.String())
    result := num.Div(mul)

    return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
    amount := decimal.NewFromFloat(0)
    switch v := iamount.(type) {
    case string:
        amount, _ = decimal.NewFromString(v)
    case int:
        amount = decimal.NewFromInt(int64(v))
    case float64:
        amount = decimal.NewFromFloat(v)
    case int64:
        amount = decimal.NewFromFloat(float64(v))
    case decimal.Decimal:
        amount = v
    case *decimal.Decimal:
        amount = *v
    }

    mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
    result := amount.Mul(mul)

    wei := new(big.Int)
    wei.SetString(result.String(), 10)

    return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
// example
// gasLimit := uint64(21000)
// gasPrice := new(big.Int)
// gasPrice.SetString("2000000000", 10)
// gasCost := util.CalcGasCost(gasLimit, gasPrice)
// fmt.Println(gasCost) // 42000000000000
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
    gasLimitBig := big.NewInt(int64(gasLimit))
    return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
// sig := "0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301"
// r, s, v := util.SigRSV(sig)
// fmt.Println(hexutil.Encode(r[:])[2:]) // 789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c6
// fmt.Println(hexutil.Encode(s[:])[2:]) // 2621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde023
// fmt.Println(v)                        // 28
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
    var sig []byte
    switch v := isig.(type) {
    case []byte:
        sig = v
    case string:
        sig, _ = hexutil.Decode(v)
    }

    sigstr := common.Bytes2Hex(sig)
    rS := sigstr[0:64]
    sS := sigstr[64:128]
    R := [32]byte{}
    S := [32]byte{}
    copy(R[:], common.FromHex(rS))
    copy(S[:], common.FromHex(sS))
    vStr := sigstr[128:130]
    vI, _ := strconv.Atoi(vStr)
    V := uint8(vI + 27)

    return R, S, V
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
	accountAddress := Account.GetAddressFromHex(address)
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
	entropy, _ := hdwallet.NewEntropy(256)
	mnemonic, _ := hdwallet.NewMnemonicFromEntropy(entropy)
	return mnemonic
}