package wallet

import (
	utils "blockchain/Utils"
	"crypto/x509"
	"encoding/hex"
	"math/big"
)

const (
	signature     string = "957df8baf7130c962f406295afe8ba172f8e6ee5bc37ef8f35837ce9c8bd1c8bf8727f11349ca19df2c58fd012c8872857155c29cabb451b4031d833cb3ff992"
	privatekey    string = "307702010104208a17af7b363b5f9e73142f86750233a90a20c014626b5f29a4373161d4bb47b4a00a06082a8648ce3d030107a1440342000450627448618a2e26e63143741d567707043b90bb7f4d112c4d7191593ad9b509f64674bf4ded43ce200264246ef49de7448a0d89c3861819a3199590120a3dc4"
	hashedMessage string = "a79f577e868cb214ce016d6d459ad1a4ec75ddce099fe2cfb13835ec6be266a2"
)

func Start() {
	privBytes, err := hex.DecodeString(privatekey)

	utils.HandleError(err)

	private, err := x509.ParseECPrivateKey(privBytes)

	println(private)
	utils.HandleError(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
}
