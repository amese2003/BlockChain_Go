package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/db"
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newesthash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (c *blockchain) AddBlock(data string) {

	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("test block")
		})
	}

	return b
}
