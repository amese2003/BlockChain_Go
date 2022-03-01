package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type BlockChain struct {
	blocks []*Block
}

var b *BlockChain
var once sync.Once

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockChain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *Block) calculateHash() {
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

func (b *BlockChain) AllBlock() []*Block {
	return GetBlockChain().blocks
}

var ErrNotFound = errors.New("block 없음")

func (b *BlockChain) GetBlock(height int) (*Block, error) {
	count := len(b.blocks)

	if height > count {
		return nil, ErrNotFound
	}

	return b.blocks[height-1], nil
}
