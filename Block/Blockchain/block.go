package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/db"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) toBytes() []byte {
	var blockbuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockbuffer)
	utils.HandleError(encoder.Encode(b))
	return blockbuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
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
	newBlock.persist()
	return newBlock
}
