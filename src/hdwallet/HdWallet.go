package hdwallet

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	Helpers "github.com/kaikoh95/web3go/src/helpers"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func GenerateHdWalletFromMnemonic(mnemonic string) *hdwallet.Wallet {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal("Unable to generate wallet with mnemonic phrase ", err)
	}
	return wallet
}

// Generates hd wallet
// returns hd wallet and mnemonic phrase
func GenerateHdWalletNew() (*hdwallet.Wallet, string) {
	mnemonic := Helpers.GenerateMnemonic()
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal("Unable to generate wallet with mnemonic phrase ", err)
	}
	return wallet, mnemonic
}

// "m / purpose' / coin_type' / account' / change / address_index"
// reference from: https://ethereum.stackexchange.com/questions/70017/can-someone-explain-the-meaning-of-derivation-path-in-wallet-in-plain-english-s
func ParsedPathFromDefaultDerivationPath() accounts.DerivationPath {
	dPath := "m/44'/60'/0'/0/0"
	return hdwallet.MustParseDerivationPath(dPath)
}

func ParsedPathFromDerivationPath(dPath string) accounts.DerivationPath {
	return hdwallet.MustParseDerivationPath(dPath)
}

func DeriveAccountFromParsedPath(wallet *hdwallet.Wallet, path accounts.DerivationPath) accounts.Account {
	// pin set to true to add to list of tracked accounts (to read up on side-effects)
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal("Unable to get account from wallet and path ", err)
	}
	return account
}