// Users wallet

package wallet

import (
	utils "blockchain/Utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/fs"
	"math/big"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
}

type fileLayer interface {
	hasWalletFile() bool
	writeFile(name string, data []byte, perm fs.FileMode) error
	readFile(name string) ([]byte, error)
}

type layer struct{}

func (layer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (layer) readFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

var w *wallet
var files fileLayer = layer{}

const (
	fileName string = "nerocoin.wallet"
)

func (layer) hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleError(err)
	return privKey
}

func restoreKey() *ecdsa.PrivateKey {
	keyAsBytes, err := files.readFile(fileName)
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
	err = files.writeFile(fileName, bytes, 0644)
	utils.HandleError(err)
}

func encodeBigInts(a, b []byte) string {
	z := append(a, b...)
	return fmt.Sprintf("%x", z)
}

func aFromK(key *ecdsa.PrivateKey) string {
	x := key.X.Bytes()
	y := key.Y.Bytes()

	return encodeBigInts(x, y)
}

func Sign(payload string, w *wallet) string {
	bytes, err := hex.DecodeString(payload)
	utils.HandleError(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, bytes)
	utils.HandleError(err)
	return encodeBigInts(r.Bytes(), s.Bytes())
}

func restoreBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)

	if err != nil {
		return nil, nil, err
	}

	frontBytes := bytes[:len(bytes)/2]
	backBytes := bytes[len(bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(frontBytes)
	bigB.SetBytes(backBytes)
	return &bigA, &bigB, nil
}

func Verify(signature, payload, address string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleError(err)

	x, y, err := restoreBigInts(address)
	utils.HandleError(err)

	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	payloadBytes, err := hex.DecodeString(payload)
	utils.HandleError(err)
	ok := ecdsa.Verify(&publicKey, payloadBytes, r, s)
	return ok
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if files.hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivateKey()
			saveKey(key)
			w.privateKey = key
		}

		w.Address = aFromK(w.privateKey)
	}

	return w
}
