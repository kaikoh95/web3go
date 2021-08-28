package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kaikoh95/web3go/src/Account"
	"github.com/kaikoh95/web3go/src/client"
	"github.com/kaikoh95/web3go/src/hdwallet"

	"github.com/kaikoh95/web3go/src/helpers"

	"github.com/kaikoh95/web3go/src/transactions"
)

func main() {
	// get eth-ropsten from faucet
	// https://faucet.dimensions.network/

	client := client.InitWithDefaultNetwork()
	fmt.Println("we have a connection", client)

	chainID := transactions.GetClientChainID(client)
	fmt.Println("chain ID", chainID)
	
	// networkUrl := "https://ropsten.etherscan.io"
	networkUrl := "https://rinkeby.etherscan.io"
	fmt.Println("network url", networkUrl)

	////////////// start test code //////////////


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
	// privateKey := wallet.GeneratePrivateKey()
    // fmt.Println("privateKey ", privateKey)
	
	// privateKeyBytes := wallet.GetPrivateKeyBytes(privateKey)
    // fmt.Println("privateKeyBytes ", privateKeyBytes)
	
	// privateKeyHex := wallet.GetPrivateKeyHex(privateKey)
    // fmt.Println("privateKeyHex ", privateKeyHex)
	
	// publicKey := wallet.GetPublicKey(privateKey)
    // fmt.Println("publicKey ", publicKey)
	
	// publicKeyECDSA := wallet.GetPublicKeyECDSA(publicKey)
	// fmt.Println("publicKeyECDSA ", publicKeyECDSA)

	// // publicKeyHex := wallet.GetPublicKeyHex(publicKeyECDSA)
	// // fmt.Println("publicKeyHex ", publicKeyHex)
	
	// publicAddress := wallet.GetPublicAddress(publicKeyECDSA)
	// fmt.Println("publicAddress ", publicAddress)

	// ///// Keystores setup
	// var account accounts.Account
	// ks := keystores.InitKeyStore("./wallets")
	// password := "password"
	
	///// Create new account 
	// account = keystores.CreateNewAccount(ks, password)
	
	// ///// Import account
	// file := "./wallets/UTC--2021-08-21T23-38-28.676160000Z--896562a998b4b819f23c05dc78c39e6f43f70b3d"
	// account =  keystores.ImportAccountFromKeyStore(ks, file, password)
	
	// accountAddressFromKeystores := Account.GetAccountAddress(account)
	// fmt.Println("address hex ", Account.GetAccountAddressHex(accountAddressFromKeystores))

	///// Blocks
	// blockHeader := blocks.GetBlockHeader(client)
	// fmt.Println("blockHeader number", blockHeader.Number.String()) // 5671744
	// var blockNumber *big.Int
    // blockNumber = big.NewInt(5671744)
    // block := blocks.GetFullBlock(client, blockNumber)
	// blockStruct := blocks.Block{
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
    // block := blocks.GetFullBlock(client, blockNumber)
	// transactions := transactions.GetTransactionDetails(client, block, withReceipt)
	// fmt.Println(transactions)
	
	// blockHash := transactions.GetHashFromHex(block.Hash().Hex())
	// fmt.Println("block hash ", blockHash)
	// transactionsFromBlockHash := transactions.GetTransactionsFromBlockHash(blockHash, client, withReceipt)
	// fmt.Println(transactionsFromBlockHash)
	
	// transactionHex := "0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2"
	// fmt.Println("transactionHex ", transactionHex)
    // transactionHash := transactions.GetHashFromHex(transactionHex)
	// singleTransactionFromTransactionHash := transactions.GetSingleTransactionFromTransactionHash(transactionHash, client, withReceipt)
	// fmt.Println("singleTransactionFromTransactionHash ", singleTransactionFromTransactionHash)

	///// HDWallets
	ryanAddr := Account.GetAddressFromHex("0x3AbbBD290D35ea69E62eE1b5D45da6934490DC34")
	fmt.Println("Ryan address ", ryanAddr)

	// wallet to act as receiver
	recMnemonic := "cruel twelve firm dignity huge such boost vault meadow monkey grace outer element cruise danger live benefit morning draft fury someone fall rich pride"
	recWallet := hdwallet.GenerateHdWalletFromMnemonic(recMnemonic)
	recPath := hdwallet.ParsedPathFromDefaultDerivationPath()
	recAccount := hdwallet.DeriveAccountFromParsedPath(recWallet, recPath)
	recPub, _ := recWallet.PublicKey(recAccount)
	recPubAddr := crypto.PubkeyToAddress(*recPub)
	fmt.Println("Receiver address ", recPubAddr)

	// wallet to act as sender
	mnemonic := "man drastic shed trip rug extra bar trophy sign floor vibrant step square hour clap file brown black cable seminar squirrel holiday negative brain"
	wallet1 := hdwallet.GenerateHdWalletFromMnemonic(mnemonic)
	path := hdwallet.ParsedPathFromDefaultDerivationPath()
	account1 := hdwallet.DeriveAccountFromParsedPath(wallet1, path)
	
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
	amountToSend := 0.0005
	weiAmountToSend := helpers.ToWei(amountToSend, 18)
	fmt.Println("amount to send in wei", weiAmountToSend)
	// gasLimit := helpers.EstimateGasLimit(client, &recPubAddr, nil)
	gasPrice := helpers.SuggestGasPrice(client)
	// gasCost := helpers.CalcGasCost(gasLimit, gasPrice)
	// fmt.Println("gas cost in wei ", gasCost)

	//// nil for just sending ETH 
	// Legacy Tx
	// tx := types.NewTransaction(nonce, recPubAddr, weiAmountToSend, gasLimit, gasPrice, nil)
	
	//// EIP-1559 Tx
	// ethPrepTx := types.DynamicFeeTx{
	// 	ChainID: chainID,
	// 	Nonce: nonce,
	// 	GasFeeCap: gasPrice,
	// 	GasTipCap: gasPrice,
	// 	Gas: gasLimit,
	// 	To: &recPubAddr,
	// 	Value: weiAmountToSend,
	// }

	//// EIP-2930 Tx
	// var storageKeys []common.Hash
	
	// storageKeys = append(storageKeys, accountAddress1.Hash())
	// accessTuple := types.AccessTuple{
	// 	Address: accountAddress1,
	// 	StorageKeys: storageKeys,
	// }
	// var accessList []types.AccessTuple
	// accessList = append(accessList, accessTuple)
	// ethPrepTx := types.AccessListTx{
	// 	ChainID: chainID,
	// 	Nonce: nonce,
	// 	GasPrice: gasPrice,
	// 	Gas: uint64(800000),
	// 	To: &recPubAddr,
	// 	Value: weiAmountToSend,
	// 	AccessList: accessList,
	// }
	// ethTx := types.NewTx(&ethPrepTx)

	// signedEthTx, err := types.SignTx(ethTx, types.NewEIP2930Signer(chainID), privateKey)
	// if err != nil {
	// 	log.Fatal("Failed to sign transaction ", err)
	// }

	// err = client.SendTransaction(context.Background(), signedEthTx)
	// if err != nil {
	// 	log.Fatal("Unable to send transaction ", err)
	// }

	// fmt.Printf("check tx here %s/tx/%s \n", networkUrl, signedEthTx.Hash().Hex())
	// ///// fmt.Println(spew.Dump(signedEthTx))

	// ///// send ERC-20 Token

	//// Ropsten
	//// SatorV2 := "0xaced0798cbA611f1613BaD8E6dC0Ac30C2C4Bb66"
	
	//// Rinkeby
	// SatorV2 := "0xb082ee966A31089b08bcE68Ec43C0aB580be2A8a"
	Sator777 := "0x5794E45E4E06C099C0873c633C41943610DBfBe2"
	tokenAddress := Account.GetAddressFromHex(Sator777)
	fmt.Println(tokenAddress)

	methodName := "transfer(address,uint256)"
	tokenAmountToSend := 1000
	data := transactions.PrepareTokenTransactionData(
		methodName, &recPubAddr, tokenAmountToSend)
	// gasPrice := helpers.SuggestGasPrice(client)
	gasLimit := helpers.EstimateGasLimit(client, &recPubAddr, data)
	fmt.Println("gas ", gasLimit) // 23256
	fmt.Println("gas price", gasPrice) // 23256
	fmt.Println("nonce ", nonce)
	dynamicTx := types.DynamicFeeTx{
		ChainID: chainID,
		Nonce: nonce,
		GasFeeCap: gasPrice.Mul(gasPrice, big.NewInt(50)),
		GasTipCap: gasPrice.Mul(gasPrice, big.NewInt(50)),
		Gas: gasLimit * 20,
		To: &tokenAddress,
		Value: big.NewInt(0),
		Data: data,
	}
	tx := types.NewTx(&dynamicTx)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Unable to send transaction ", err)
	}
	fmt.Printf("tx hex %s \n", signedTx.Hash().Hex())
	fmt.Printf("check tx here %s/tx/%s \n", networkUrl, signedTx.Hash().Hex())
	
}
