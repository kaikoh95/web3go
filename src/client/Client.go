package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func InitClient(network string) *ethclient.Client {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }
    return client
}