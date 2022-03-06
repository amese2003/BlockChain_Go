package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/db"
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newesthash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.HandleError(gob.NewDecoder(bytes.NewReader(data)).Decode(b))
}

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
			fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
			checkpoint := db.CheckPoint()
			if checkpoint == nil {
				b.AddBlock("test block")
				b.AddBlock("second block")
				b.AddBlock("third block")
			} else {
				b.restore(checkpoint)
			}

			fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
		})
	}

	fmt.Println(b.NewestHash)
	return b
}
