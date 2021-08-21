package main

import (
	"fmt"
	"math/big"

	Account "github.com/kaikoh95/web3go/src/account"
	Client "github.com/kaikoh95/web3go/src/client"
	Helpers "github.com/kaikoh95/web3go/src/helpers"
	Wallet "github.com/kaikoh95/web3go/src/wallet"
)

func main() {
	///// Client setup
	// network := "https://mainnet.infura.io"
	network := "http://localhost:8545"
	
	client := Client.InitClient(network)
	fmt.Println("we have a connection", client)
	
	///// Accounts setup
	address := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	blockNumber := big.NewInt(5532993)
	account := Account.GetAccount(address)
	fmt.Println("hex", account.Hex())
	fmt.Println("hash hex", account.Hash().Hex())
	fmt.Println("bytes", account.Bytes())

	balance := Account.GetAccountBalance(client, account)
	fmt.Println("balance", balance)

	balanceAtBlockTime := Account.GetAccountBalanceAtBlockTime(client, account, blockNumber)
	fmt.Println("balance at block time", balanceAtBlockTime)

	weiBalance := Helpers.ConvertToWei(balance)
	fmt.Println("wei balance", weiBalance)

	pendingBalance := Account.GetAccountPendingBalance(client, account)
	fmt.Println("pending balance", pendingBalance)

	///// Wallet setup
	privateKey := Wallet.GeneratePrivateKey()
    fmt.Println("privateKey", privateKey)
	privateKeyBytes := Wallet.GetPrivateKeyBytes(privateKey)
    fmt.Println("privateKeyBytes", privateKeyBytes)
	
	publicKey := Wallet.GetPublicKey(privateKey)
    fmt.Println("publicKey", publicKey)
	
	publicKeyECDSA := Wallet.GetPublicKeyECDSA(publicKey)
	fmt.Println("publicKeyECDSA", publicKeyECDSA)

	publicKeyHex := Wallet.GetPublicKeyHex(publicKeyECDSA)
	fmt.Println("publicKeyHex", publicKeyHex)
	
	publicAddress := Wallet.GetPublicAddress(publicKeyECDSA)
	fmt.Println("publicAddress", publicAddress)
}
