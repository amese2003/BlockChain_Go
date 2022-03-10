package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/db"
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockchain struct {
	NewestHash        string `json:"newesthash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.HandleError(gob.NewDecoder(bytes.NewReader(data)).Decode(b))
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {

	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) calculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastrecal := allBlocks[difficultyInterval-1]
	actualtime := (newestBlock.Timestamp / 60) - (lastrecal.Timestamp / 60)
	expectedtime := difficultyInterval * blockInterval

	if actualtime <= (expectedtime - allowedRange) {
		return b.CurrentDifficulty + 1

	} else if actualtime >= (expectedtime + allowedRange) {
		return b.CurrentDifficulty - 1
	}

	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.calculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				b.restore(checkpoint)
			}

			fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
		})
	}

	return b
}

func (b *blockchain) UTxOutsByAddress(address string) []*UTxOut {
	var uTxOuts []*UTxOut
	createTxs := make(map[string]bool)

	for _, block := range b.Blocks() {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Owner == address {
					createTxs[input.TxID] = true
				}
			}

			for idx, output := range tx.TxOuts {
				if output.Owner == address {
					if _, ok := createTxs[tx.Id]; !ok {
						uTxOuts = append(uTxOuts, &UTxOut{tx.Id, idx, output.Amount})
					}
				}
			}
		}

	}

	return uTxOuts
}

func (b *blockchain) BalanceByAddress(address string) int {
	txOuts := b.UTxOutsByAddress(address)
	var amount int

	for _, txout := range txOuts {
		amount += txout.Amount
	}

	return amount
}
