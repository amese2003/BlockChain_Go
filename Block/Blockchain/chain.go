package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newesthash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (c *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
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
