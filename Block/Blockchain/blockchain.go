package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type BlockChain struct {
	blocks []*block
}

var b *BlockChain
var once sync.Once

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	total := len(GetBlockChain().blocks)

	if total == 0 {
		return ""
	}

	return GetBlockChain().blocks[total-1].Hash
}

func (c *BlockChain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *BlockChain {
	if b == nil {
		once.Do(func() {
			b = &BlockChain{}
			b.AddBlock("test block")
		})
	}

	return b
}

func (b *BlockChain) AllBlock() []*block {
	return GetBlockChain().blocks
}
