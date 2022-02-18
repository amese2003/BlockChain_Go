package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type BlockChain struct {
	blocks []block
}

func (b *BlockChain) GetLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}

	return ""
}

func (b *BlockChain) AddBlock(data string) {
	newBlock := block{data, "", b.GetLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}

func (b *BlockChain) ListBlock() {

	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Hash : %s\n", block.hash)
		fmt.Printf("PrevHash : %s\n\n", block.prevHash)
	}

}

func main() {
	var testchain BlockChain
	testchain.AddBlock("genesis block")
	testchain.AddBlock("a block")
	testchain.AddBlock("b block")

	testchain.ListBlock()
}
