package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/db"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
)

const difficulty int = 2

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("블록이.. 없어요!")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *Block) Mine() {
	target := strings.Repeat("0", b.Difficulty)

	for {
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsString, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}

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
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}

	newBlock.Mine()
	newBlock.persist()
	return newBlock
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	currentHash := b.NewestHash

	for {
		block, _ := FindBlock(currentHash)
		blocks = append(blocks, block)
		if block.PrevHash == "" {
			break
		} else {
			currentHash = block.PrevHash
		}
	}

	return blocks
}
