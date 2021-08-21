package wallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GeneratePrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		// log.Fatal(err)
		fmt.Println("error generating private key", err)
	}
	return privateKey
}

func GetPrivateKeyBytes(privateKey *ecdsa.PrivateKey) []byte {
	return crypto.FromECDSA(privateKey)
}

// Private key to sign transactions
func GetPrivateKeyHex(privateKey *ecdsa.PrivateKey) string {
	privateKeyBytes := GetPrivateKeyBytes(privateKey)
	return hexutil.Encode(privateKeyBytes)[2:]
}

func GetPublicKey(privateKey *ecdsa.PrivateKey) interface{} {
	return privateKey.Public()
}

func GetPublicKeyECDSA(publicKey interface{}) *ecdsa.PublicKey {
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        // log.Fatal("error casting public key to ECDSA")
        fmt.Println("error casting public key to ECDSA")
    }
	return publicKeyECDSA
}

func GetPublicKeyBytes(publicKeyECDSA *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(publicKeyECDSA)
}

func GetPublicKeyHex(publicKeyECDSA *ecdsa.PublicKey) string {
	publicKeyBytes := GetPublicKeyBytes(publicKeyECDSA)
	return hexutil.Encode(publicKeyBytes)[4:]
}

func GetPublicAddress(publicKeyECDSA *ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}

func GetPublicAddressLegacy(publicKeyECDSA *ecdsa.PublicKey) string {
	publicKeyBytes := GetPublicKeyBytes(publicKeyECDSA)
	hash := sha3.NewLegacyKeccak256()
    hash.Write(publicKeyBytes[1:])
	return hexutil.Encode(hash.Sum(nil)[12:])
}
