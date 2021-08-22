package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClientChainID (client *ethclient.Client) *big.Int {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return chainID
}

func GetTransactionSenderAddress(tx *types.Transaction, chainID *big.Int) string {
	if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), big.NewInt(1)); err == nil {
		return msg.From().Hex()
	}
	return ""
} 

func GetTransactionReceiptDetailsFromHash(hash common.Hash, client *ethclient.Client) *types.Receipt {
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Fatal("error getting transaction receipt ",err)
	}
	return receipt
}

type Transaction struct {
	tx *types.Transaction
	chainID *big.Int
	lenTransactions string
	client *ethclient.Client
	withReceipt bool
	index int
}

func (t Transaction) FormatTransactionDetails() string {
	data := "#############\n"
	data += "transaction " + strconv.Itoa(t.index) + " of " + t.lenTransactions + "\n"
	data += "transaction chain id " + strconv.FormatUint(t.chainID.Uint64(), 10) + "\n"
	data += "transaction hash hex " + t.tx.Hash().Hex() + "\n"
	data += "transaction value " + t.tx.Value().String() + "\n"
	data += "transaction gas limit " + strconv.FormatUint(t.tx.Gas(), 10) + "\n"
	data += "transaction gas price " + strconv.FormatUint(t.tx.GasPrice().Uint64(), 10) + "\n"
	data += "transaction sender nonce " + strconv.FormatUint(t.tx.Nonce(), 10) + "\n"
	data += "transaction input data " + string(t.tx.Data()[:]) + "\n"
	data += "transaction recipient address (or nil for contract creation) " + t.tx.To().Hex() + "\n"
	data += "transaction sender address " + GetTransactionSenderAddress(t.tx, t.chainID) + "\n"

	if t.withReceipt {
		receipt := GetTransactionReceiptDetailsFromHash(t.tx.Hash(), t.client)
		data += "transaction status " + strconv.FormatUint(receipt.Status, 10) + "\n"
		// fmt.Println(receipt.Logs)
	}

	data += "#############\n"
	return data
}

func GetTransactionDetails(client *ethclient.Client, block *types.Block, withReceipt bool) []string {
	lenTransactions := strconv.Itoa(len(block.Transactions()))
	chainID := GetClientChainID(client)
	var transactions []string
	for i, tx := range block.Transactions() {
		t := Transaction{tx, chainID, lenTransactions, client, withReceipt, i}
		data := t.FormatTransactionDetails()
		transactions = append(transactions, data)
    }
	return transactions
}

func GetBlockHashFromHex(hex string) common.Hash {
	return common.HexToHash(hex)
}

func GetTransactionsFromBlockHash(hash common.Hash, client *ethclient.Client) []string {
	lenTransactions, err := client.TransactionCount(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
	}
	chainID := GetClientChainID(client)
	var transactions []string
	for idx := uint(0); idx < lenTransactions; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), hash, idx)
		if err != nil {
			log.Fatal("error reading transaction in block ", err)
		}
		withReceipt := false
		t := Transaction{tx, chainID, strconv.FormatUint(uint64(lenTransactions), 10), client, withReceipt, idx}
		data := t.FormatTransactionDetails()
		transactions = append(transactions, data)
    }
	return transactions
}

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/7073cc887d0449feaf3017cc7bc6090e")
    if err != nil {
        log.Fatal(err)
    }

    blockNumber := big.NewInt(5671744)
    block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }

	withReceipt := true
	transactions := GetTransactionDetails(client, block, withReceipt)
	fmt.Println(transactions)

    // blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
    // count, err := client.TransactionCount(context.Background(), blockHash)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // for idx := uint(0); idx < count; idx++ {
    //     tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
    //     if err != nil {
    //         log.Fatal(err)
    //     }

    //     fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
    // }

    // txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
    // tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
    // fmt.Println(isPending)       // false
}