package blocks

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

type Block struct {
	Number uint64
	Difficulty uint64
	Time uint64
	Hash common.Hash
	Hex string
	LenTransactions int
}

func GetBlockHeader(client *ethclient.Client) *types.Header {
	blockHeader, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal("Unable to get block header ", err)
    }
	fmt.Println("blockHeader number", blockHeader.Number.String()) // 5671744
	return blockHeader
}

func (block Block) FormatBlockDetails() string {
	format := "#############\n"
	format += "block number " + strconv.FormatUint(block.Number, 10) + "\n"
	format += "block number of transactions: " + strconv.Itoa(block.LenTransactions) + "\n"
	format += "block time " + strconv.FormatUint(block.Time, 10) + "\n"
	format += "block difficulty " + strconv.FormatUint(block.Difficulty, 10) + "\n"
	format += "block hex " + block.Hex + "\n"
	format += "block hash " + block.Hash.String() + "\n"
	format += "#############\n"
	return format
}

func GetFullBlock(client *ethclient.Client, blockNumber *big.Int) *types.Block {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal("Unable to get block from block number ", err)
    }
	return block
}

func GetBlockTransactionCount(client *ethclient.Client, hash common.Hash) uint {
	count, err := client.TransactionCount(context.Background(), hash)
    if err != nil {
        log.Fatal("Unable to get transaction count of block ", err)
    }
	return count
}