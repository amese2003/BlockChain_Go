package wallet

import (
	utils "blockchain/Utils"
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testkey     = "30770201010420099826f110baedbf3a5ea31efdf9079085e7bad40aa9e4ed7891e2fece83fb70a00a06082a8648ce3d030107a14403420004fc9eefc6a24aa724fa1fb7ebff97f886b2fd194e21684bedd3dcd25af3db4380b5256b82d6f12cfc836ed8e39371cea23029b13524e3ad45a95c93e22f76e7ba"
	testPayload = "0035e1f04b7e6b41d24873d3df88e90a9258cd9e1bbdd8f769e23be22e2edda1"
	testSign    = "41df91a4c59144ccc0d7ad8bf7f013197d8fa87dea1cba9ea018a9df19f2aa10344277bf0e25323c9bf96c7856ef616131b11872c45c7e0911966ca14a4d77d9"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, err := hex.DecodeString(testkey)
	utils.HandleError(err)
	key, err := x509.ParseECPrivateKey(b)
	utils.HandleError(err)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}

	tests := []test{
		{testPayload, true},
		//{"0035e1f04b7e6b41d24873d3df88e90a9258cd9e1bbdd8f769e23be22e2", false},
	}

	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSign, tc.input, w.Address)

		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("가나다라")
	if err == nil {
		t.Error("restoreBigInts는 값이 hex가 아니라면 에러를 반환해야합니다.")
	}
}
