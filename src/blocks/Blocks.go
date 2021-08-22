package blocks

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockHeader(client *ethclient.Client) *types.Header {
	blockHeader, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("blockHeader number", blockHeader.Number.String()) // 5671744
	return blockHeader
}

func FormatBlockDetails(block *types.Block) string {
	format := "#############\n"
	format += "block number " + strconv.FormatUint(block.Number().Uint64(), 10) + "\n"
	format += "block number of transactions: " + strconv.Itoa(len(block.Transactions())) + "\n"
	format += "block time " + strconv.FormatUint(block.Time(), 10) + "\n"
	format += "block difficulty " + strconv.FormatUint(block.Difficulty().Uint64(), 10) + "\n"
	format += "block hash hex " + block.Hash().Hex() + "\n"
	format += "#############\n"
	return format
}

func GetFullBlock(client *ethclient.Client, blockNumber *big.Int) *types.Block {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }
	return block
}

func GetBlockTransactionCount(client *ethclient.Client, block *types.Block) uint {
	count, err := client.TransactionCount(context.Background(), block.Hash())
    if err != nil {
        log.Fatal(err)
    }
	return count
}