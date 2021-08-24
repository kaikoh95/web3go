package transactions

import (
	"context"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	Blocks "github.com/kaikoh95/web3go/src/blocks"
)

type Transaction struct {
	Tx *types.Transaction
	ChainID *big.Int
	LenTransactions string
	Index int
}

func GetClientChainID (client *ethclient.Client) *big.Int {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Unable to get chain ID from client ", err)
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

func (t Transaction) FormatTransactionDetails(client *ethclient.Client, withReceipt bool) string {
	data := "#############\n"
	data += "transaction " + strconv.Itoa(t.Index) + " of " + t.LenTransactions + "\n"
	data += "transaction chain id " + strconv.FormatUint(t.ChainID.Uint64(), 10) + "\n"
	data += "transaction hash hex " + t.Tx.Hash().Hex() + "\n"
	data += "transaction value " + t.Tx.Value().String() + "\n"
	data += "transaction gas limit " + strconv.FormatUint(t.Tx.Gas(), 10) + "\n"
	data += "transaction gas price " + strconv.FormatUint(t.Tx.GasPrice().Uint64(), 10) + "\n"
	data += "transaction sender nonce " + strconv.FormatUint(t.Tx.Nonce(), 10) + "\n"
	data += "transaction input data " + string(t.Tx.Data()[:]) + "\n"
	data += "transaction recipient address (or nil for contract creation) " + t.Tx.To().Hex() + "\n"
	data += "transaction sender address " + GetTransactionSenderAddress(t.Tx, t.ChainID) + "\n"

	if withReceipt {
		receipt := GetTransactionReceiptDetailsFromHash(t.Tx.Hash(), client)
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
		t := Transaction{Tx: tx, ChainID: chainID, LenTransactions: lenTransactions, Index: i}
		data := t.FormatTransactionDetails(client, withReceipt)
		transactions = append(transactions, data)
    }
	return transactions
}

func GetHashFromHex(hex string) common.Hash {
	return common.HexToHash(hex)
}

func GetTransactionsFromBlockHash(hash common.Hash, client *ethclient.Client, withReceipt bool) []string {
	lenTransactions := Blocks.GetBlockTransactionCount(client, hash)
	chainID := GetClientChainID(client)
	var transactions []string
	for idx := uint(0); idx < lenTransactions; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), hash, idx)
		if err != nil {
			log.Fatal("error reading transaction in block ", err)
		}
		t := Transaction{
			Tx: tx, 
			ChainID: chainID, 
			LenTransactions: strconv.FormatUint(uint64(lenTransactions), 10), 
			Index: int(idx),
		}
		data := t.FormatTransactionDetails(client, withReceipt)
		transactions = append(transactions, data)
    }
	return transactions
}

func GetSingleTransactionFromTransactionHash(hash common.Hash, client *ethclient.Client, withReceipt bool) string {
	tx, isPending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Fatal("No transactions found ", err)
	}
	if !isPending {
		lenTransactions := Blocks.GetBlockTransactionCount(client, hash)
		chainID := GetClientChainID(client)
		t := Transaction{
			Tx: tx, ChainID: chainID, 
			LenTransactions: strconv.FormatUint(uint64(lenTransactions), 10), 
			Index: 0,
		}
		return t.FormatTransactionDetails(client, withReceipt)
	}
	return "Transaction " + tx.Hash().Hex() + " is pending"
}