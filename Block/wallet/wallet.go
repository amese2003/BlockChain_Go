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
	address    string
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

func restoreKey() *ecdsa.PrivateKey {
	keyAsBytes, err := os.ReadFile(fileName)
	utils.HandleError(err)
	key, err := x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleError(err)
	return key
}

// go에선 이렇게도 return이 가능하다더라...
// func restoreKey2() (key *ecdsa.PrivateKey) {
// 	keyAsBytes, err := os.ReadFile(fileName)
// 	utils.HandleError(err)
// 	key, err = x509.ParseECPrivateKey(keyAsBytes)
// 	utils.HandleError(err)
// 	return
// }

func saveKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleError(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleError(err)
}

func aFromK(key *ecdsa.PrivateKey) string {

	return ""
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivateKey()
			saveKey(key)
			w.privateKey = key
		}

		w.address = aFromK(w.privateKey)
	}

	return w
}
