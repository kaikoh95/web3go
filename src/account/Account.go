package account

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountAddressFromHex(address string) common.Address {
	return common.HexToAddress(address)
}

func GetAccountBalance(client *ethclient.Client, account common.Address) *big.Int {
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("error at getAccountBalance", err)
	}
	return balance
}

func GetAccountBalanceAtBlockTime(client *ethclient.Client, account common.Address, blockNumber *big.Int) *big.Int {
	balanceAtBlockTime, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal("error at getAccountBalanceAtBlockTime", err)
	}
	return balanceAtBlockTime
}

func GetAccountPendingBalance(client *ethclient.Client, account common.Address) *big.Int {
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal("error at getAccountPendingBalance", err)
	}
	return pendingBalance
}

func GetAccountAddressHex (accountAddress common.Address) string {
	return accountAddress.Hex()
}

func GetAccountAddressHashHex (accountAddress common.Address) string {
	return accountAddress.Hash().Hex()
}

func GetAccountAddressBytes (accountAddress common.Address) []byte {
	return accountAddress.Bytes()
}

func GetAccountAddress(account accounts.Account) common.Address {
	return account.Address
}