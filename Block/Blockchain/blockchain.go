package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
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
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	total := len(GetBlockChain().blocks)

	if total == 0 {
		return ""
	}

	return GetBlockChain().blocks[total-1].hash
}

func GetBlockChain() *BlockChain {
	if b == nil {
		once.Do(func() {
			b = &BlockChain{}
			b.blocks = append(b.blocks, createBlock("test block"))
		})
	}

	return b
}
