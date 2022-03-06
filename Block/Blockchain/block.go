package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/db"
	"crypto/sha256"
	"errors"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("블록이.. 없어요!")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockbyte := db.Block(hash)

	if blockbyte == nil {
		return nil, ErrNotFound
	}

	block := &Block{}
	block.restore(blockbyte)
	return block, nil
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
