package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	Account "github.com/kaikoh95/web3go/src/account"
	Client "github.com/kaikoh95/web3go/src/client"
	HdWallet "github.com/kaikoh95/web3go/src/hdwallet"

	// Helpers "github.com/kaikoh95/web3go/src/helpers"

	// Blocks "github.com/kaikoh95/web3go/src/blocks"
	Transactions "github.com/kaikoh95/web3go/src/transactions"
	// Keystores "github.com/kaikoh95/web3go/src/keystores"
	// Wallet "github.com/kaikoh95/web3go/src/wallet"
)

func main() {
	// get eth-ropsten from faucet
	// https://faucet.dimensions.network/

	client := Client.InitWithDefaultNetwork()
	fmt.Println("we have a connection", client)

	chainID := Transactions.GetClientChainID(client)
	fmt.Println("chain ID", chainID)

	////////////// end test code //////////////

	// ///// Accounts setup
	// address := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	// accountAddress := Account.GetAddressFromHex(address)
	// fmt.Println("hex", Account.GetAccountAddressHex(accountAddress))
	// fmt.Println("hash hex", Account.GetAccountAddressHashHex(accountAddress))
	// fmt.Println("bytes", Account.GetAccountAddressBytes(accountAddress))
	
	// balance := Account.GetAccountBalance(client, accountAddress)
	// fmt.Println("balance", balance)
	
	// blockNumber = big.NewInt(5532993)
	// // balanceAtBlockTime := Account.GetAccountBalanceAtBlockTime(client, accountAddress, blockNumber)
	// // fmt.Println("balance at block time", balanceAtBlockTime)

	// realBalance := Helpers.ConvertFromWei(balance)
	// fmt.Println("wei balance", realBalance)

	// pendingBalance := Account.GetAccountPendingBalance(client, accountAddress)
	// fmt.Println("pending balance", pendingBalance)

	///// Wallet setup
	// privateKey := Wallet.GeneratePrivateKey()
    // fmt.Println("privateKey ", privateKey)
	
	// privateKeyBytes := Wallet.GetPrivateKeyBytes(privateKey)
    // fmt.Println("privateKeyBytes ", privateKeyBytes)
	
	// privateKeyHex := Wallet.GetPrivateKeyHex(privateKey)
    // fmt.Println("privateKeyHex ", privateKeyHex)
	
	// publicKey := Wallet.GetPublicKey(privateKey)
    // fmt.Println("publicKey ", publicKey)
	
	// publicKeyECDSA := Wallet.GetPublicKeyECDSA(publicKey)
	// fmt.Println("publicKeyECDSA ", publicKeyECDSA)

	// // publicKeyHex := Wallet.GetPublicKeyHex(publicKeyECDSA)
	// // fmt.Println("publicKeyHex ", publicKeyHex)
	
	// publicAddress := Wallet.GetPublicAddress(publicKeyECDSA)
	// fmt.Println("publicAddress ", publicAddress)

	// ///// Keystores setup
	// var account accounts.Account
	// ks := Keystores.InitKeyStore("./wallets")
	// password := "password"
	
	///// Create new account 
	// account = Keystores.CreateNewAccount(ks, password)
	
	// ///// Import account
	// file := "./wallets/UTC--2021-08-21T23-38-28.676160000Z--896562a998b4b819f23c05dc78c39e6f43f70b3d"
	// account =  Keystores.ImportAccountFromKeyStore(ks, file, password)
	
	// accountAddressFromKeystores := Account.GetAccountAddress(account)
	// fmt.Println("address hex ", Account.GetAccountAddressHex(accountAddressFromKeystores))

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

	///// HDWallets
	// ryanAddr := Account.GetAddressFromHex("0x3AbbBD290D35ea69E62eE1b5D45da6934490DC34")

	// wallet to act as receiver
	recMnemonic := "cruel twelve firm dignity huge such boost vault meadow monkey grace outer element cruise danger live benefit morning draft fury someone fall rich pride"
	recWallet := HdWallet.GenerateHdWalletFromMnemonic(recMnemonic)
	recPath := HdWallet.ParsedPathFromDefaultDerivationPath()
	recAccount := HdWallet.DeriveAccountFromParsedPath(recWallet, recPath)
	recPub, _ := recWallet.PublicKey(recAccount)
	recPubAddr := crypto.PubkeyToAddress(*recPub)
	fmt.Println("Receiver address ", recPubAddr)

	// wallet to act as sender
	mnemonic := "man drastic shed trip rug extra bar trophy sign floor vibrant step square hour clap file brown black cable seminar squirrel holiday negative brain"
	wallet1 := HdWallet.GenerateHdWalletFromMnemonic(mnemonic)
	path := HdWallet.ParsedPathFromDefaultDerivationPath()
	account1 := HdWallet.DeriveAccountFromParsedPath(wallet1, path)
	
	privateKey, _ := wallet1.PrivateKey(account1)
	fmt.Println("pk ", privateKey)
	accountAddress1 := Account.GetAccountAddress(account1)
	
	nonce := Account.GetAccountPendingNonce(client, accountAddress1)
	fmt.Println("sender nonce ", nonce)

	// weiBalance := Account.GetAccountBalance(client, accountAddress1)
	// fmt.Println("balance in wei", weiBalance)
	// decBalance := Helpers.ToDecimal(weiBalance, 18)
	// fmt.Println("balance in decimal", decBalance)
	// weiBal := Helpers.ToWei(decBalance, 18)
	// fmt.Println("balance in wei", weiBal)
	
	// ///// Send ETH
	// amountToSend := 0.005
	// weiAmountToSend := Helpers.ToWei(amountToSend, 18)
	// fmt.Println("amount to send in wei", weiAmountToSend)
	// gasLimit := Helpers.EstimateGasLimit(client, &recPubAddr, nil)
	// gasPrice := Helpers.SuggestGasPrice(client)
	// gasCost := Helpers.CalcGasCost(gasLimit, gasPrice)
	// fmt.Println("gas cost in wei ", gasCost)

	// /// nil for just sending ETH 
	// tx := types.NewTransaction(nonce, recPubAddr, weiAmountToSend, gasLimit, gasPrice, nil)

	// signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	// if err != nil {
	// 	log.Fatal("Failed to sign transaction ", err)
	// }

	// err = client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	log.Fatal("Unable to send transaction ", err)
	// }

	// fmt.Printf("tx sent: %s \n", signedTx.Hash().Hex())
	// ///// fmt.Println(spew.Dump(signedTx))

	// ///// send ERC-20 Token
	// SatorV2 := "0xaced0798cbA611f1613BaD8E6dC0Ac30C2C4Bb66"
	// tokenAddress := Account.GetAddressFromHex(SatorV2)
	// fmt.Println(tokenAddress)

	// methodName := "transfer(address,uint256)"
	// tokenAmountToSend := 1000
	// data := Transactions.PrepareTokenTransactionData(methodName, &recPubAddr, tokenAmountToSend)
	// gasPrice := Helpers.SuggestGasPrice(client)
	// gasLimit := Helpers.EstimateGasLimit(client, &recPubAddr, data)
	// fmt.Println("gas ", gasLimit) // 23256
	// fmt.Println("gas price", gasPrice) // 23256
	// fmt.Println("nonce ", nonce)
	
	// dynamicTx := types.DynamicFeeTx{
	// 	ChainID: chainID,
	// 	Nonce: nonce,
	// 	GasFeeCap: gasPrice,
	// 	GasTipCap: gasPrice,
	// 	Gas: gasLimit,
	// 	To: &tokenAddress,
	// 	Value: big.NewInt(0),
	// 	Data: data,
	// }
	// tx := types.NewTx(&dynamicTx)

	// signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
    // if err != nil {
    //     log.Fatal(err)
    // }
	// err = client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	log.Fatal("Unable to send transaction ", err)
	// }
	// fmt.Printf("tx hex %s \n", signedTx.Hash().Hex())
	// networkUrl := "https://ropsten.etherscan.io"
	// fmt.Printf("check tx here %s/tx/%s \n", networkUrl, signedTx.Hash().Hex())
	
}
