package wallet

import (
	utils "blockchain/Utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Start() {
	privatekey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleError(err)

	message := "무엇?"

	hashedMessage := utils.Hash(message)

	hashAsBytes, err := hex.DecodeString(hashedMessage)

	utils.HandleError(err)

	r, s, err := ecdsa.Sign(rand.Reader, privatekey, hashAsBytes)

	utils.HandleError(err)

	fmt.Printf("R:%d\nS:%d", r, s)

}
