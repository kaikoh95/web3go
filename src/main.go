package main

import (
	"fmt"

	// Account "github.com/kaikoh95/web3go/src/account"
	// Blocks "github.com/kaikoh95/web3go/src/blocks"
	Client "github.com/kaikoh95/web3go/src/client"
	// Transactions "github.com/kaikoh95/web3go/src/transactions"
	// Helpers "github.com/kaikoh95/web3go/src/helpers"
	// Keystores "github.com/kaikoh95/web3go/src/keystores"
	// Wallet "github.com/kaikoh95/web3go/src/wallet"
)

func main() {
	client := Client.InitWithDefaultNetwork()
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
	// blockHeader := Blocks.GetBlockHeader(client)
	// fmt.Println("blockHeader number", blockHeader.Number.String()) // 5671744
	// var blockNumber *big.Int
    // blockNumber = big.NewInt(5671744)
    // block := Blocks.GetFullBlock(client, blockNumber)
	// blockStruct := Blocks.Block{
	// 	Number: block.Number().Uint64(),
	// 	Difficulty: block.Difficulty().Uint64(),
	// 	Time: block.Time(),
	// 	Hash: block.Hash(),
	// 	Hex: block.Hash().Hex(),
	// 	LenTransactions: len(block.Transactions()),
	// }
	// fmt.Println(blockStruct.FormatBlockDetails())

	
	///// Transactions
	// withReceipt := true
    // blockNumber = big.NewInt(5671744)
    // block := Blocks.GetFullBlock(client, blockNumber)
	// transactions := Transactions.GetTransactionDetails(client, block, withReceipt)
	// fmt.Println(transactions)
	
	// blockHash := Transactions.GetHashFromHex(block.Hash().Hex())
	// fmt.Println("block hash ", blockHash)
	// transactionsFromBlockHash := Transactions.GetTransactionsFromBlockHash(blockHash, client, withReceipt)
	// fmt.Println(transactionsFromBlockHash)
	
	// transactionHex := "0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2"
	// fmt.Println("transactionHex ", transactionHex)
    // transactionHash := Transactions.GetHashFromHex(transactionHex)
	// singleTransactionFromTransactionHash := Transactions.GetSingleTransactionFromTransactionHash(transactionHash, client, withReceipt)
	// fmt.Println("singleTransactionFromTransactionHash ", singleTransactionFromTransactionHash)
}
