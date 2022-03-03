package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func createBlock(data string, prevHash string, height int) *Block {

	newBlock := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}

	payload := newBlock.Data + newBlock.PrevHash + fmt.Sprint(newBlock.Height)
	newBlock.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	return newBlock
}
