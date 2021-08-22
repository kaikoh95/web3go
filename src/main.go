package main

import (
	"fmt"
	"math/big"

	// Account "github.com/kaikoh95/web3go/src/account"
	Blocks "github.com/kaikoh95/web3go/src/blocks"
	Client "github.com/kaikoh95/web3go/src/client"
	// Helpers "github.com/kaikoh95/web3go/src/helpers"
	// Keystores "github.com/kaikoh95/web3go/src/keystores"
	// Wallet "github.com/kaikoh95/web3go/src/wallet"
)

func main() {
	var blockNumber *big.Int

	///// Client setup
	//// refactor into dotenv
	network := "https://mainnet.infura.io/v3/7073cc887d0449feaf3017cc7bc6090e"
	// network := "http://localhost:8545"
	
	client := Client.InitClient(network)
	fmt.Println("we have a connection", client)
	
	// ///// Accounts setup
	// address := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	// accountAddress := Account.GetAccountAddress(address)
	// fmt.Println("hex", Account.GetAccountAddressHash(accountAddress))
	// fmt.Println("hash hex", Account.GetAccountAddressHexHash(accountAddress))
	// fmt.Println("bytes", Account.GetAccountAddressBytes(accountAddress))
	
	// balance := Account.GetAccountBalance(client, accountAddress)
	// fmt.Println("balance", balance)
	
	// blockNumber = big.NewInt(5532993)
	// // balanceAtBlockTime := Account.GetAccountBalanceAtBlockTime(client, accountAddress, blockNumber)
	// // fmt.Println("balance at block time", balanceAtBlockTime)

	// weiBalance := Helpers.ConvertToWei(balance)
	// fmt.Println("wei balance", weiBalance)

	// pendingBalance := Account.GetAccountPendingBalance(client, accountAddress)
	// fmt.Println("pending balance", pendingBalance)

	// ///// Wallet setup
	// privateKey := Wallet.GeneratePrivateKey()
    // fmt.Println("privateKey", privateKey)
	// privateKeyBytes := Wallet.GetPrivateKeyBytes(privateKey)
    // fmt.Println("privateKeyBytes", privateKeyBytes)
	
	// publicKey := Wallet.GetPublicKey(privateKey)
    // fmt.Println("publicKey", publicKey)
	
	// publicKeyECDSA := Wallet.GetPublicKeyECDSA(publicKey)
	// fmt.Println("publicKeyECDSA", publicKeyECDSA)

	// publicKeyHex := Wallet.GetPublicKeyHex(publicKeyECDSA)
	// fmt.Println("publicKeyHex", publicKeyHex)
	
	// publicAddress := Wallet.GetPublicAddress(publicKeyECDSA)
	// fmt.Println("publicAddress", publicAddress)

	// ///// Keystores setup
	// ks := Keystores.InitKeyStore("./wallets")
	// password := "secret"
	
	// var account accounts.Account
	// ///// Create new account 
	// // account = Keystores.CreateNewAccount(ks, password)

	// ///// Import account
	// file := "./wallets/UTC--2021-08-21T23-38-28.676160000Z--896562a998b4b819f23c05dc78c39e6f43f70b3d"
	// account =  Keystores.ImportAccountFromKeyStore(ks, file, password)
	
	// accountAddressFromKeystores := Keystores.GetAccountAddress(account)
	// fmt.Println(Account.GetAccountAddressHash(accountAddressFromKeystores))

	///// Blocks
	blockHeader := Blocks.GetBlockHeader(client)
	fmt.Println("blockHeader number", blockHeader.Number.String()) // 5671744
	
    blockNumber = big.NewInt(5671744)
    block := Blocks.GetFullBlock(client, blockNumber)
	fmt.Println(Blocks.FormatBlockDetails(block))
}
