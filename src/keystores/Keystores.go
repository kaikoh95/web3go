package keystores

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	Helpers "github.com/kaikoh95/web3go/src/helpers"
)

func InitKeyStore(folder string) *keystore.KeyStore {
	return keystore.NewKeyStore(folder, keystore.StandardScryptN, keystore.StandardScryptP)
}

func CreateNewAccount(ks *keystore.KeyStore, password string) accounts.Account {
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal("error creating new wallet", err)
	}
	return account
}

func ImportAccountFromKeyStore(ks *keystore.KeyStore, file string, password string) accounts.Account {
	jsonBytes := Helpers.GetBytesFromFile(file)
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal("error importing wallet", err)
	}
	if err := os.Remove(file); err != nil {
		log.Fatal("error removing file", err)
	}
	return account
}