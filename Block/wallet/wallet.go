package wallet

import (
	utils "blockchain/Utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

const (
	fileName string = "nerocoin.wallet"
)

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleError(err)
	return privKey
}

func saveKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleError(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleError(err)
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {

		} else {
			key := createPrivateKey()
			saveKey(key)
			w.privateKey = key
		}
	}

	return w
}
