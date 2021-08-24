package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Init(network string) *ethclient.Client {
    client, err := ethclient.Dial(network)
    if err != nil {
        log.Fatal(err)
    }
    return client
}

func InitWithDefaultNetwork() *ethclient.Client {
    return Init("https://rinkeby.infura.io/v3/7073cc887d0449feaf3017cc7bc6090e")
}