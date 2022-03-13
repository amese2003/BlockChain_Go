package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/db"
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
	utils.FromBytes(b, data)
}

func persistBlockchain(b *blockchain) {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
}

func Txs(b *blockchain) []*Tx {
	var txs []*Tx
	for _, block := range Blocks(b) {
		txs = append(txs, block.Transactions...)
	}

	return txs
}

func FindTx(b *blockchain, targetID string) *Tx {
	for _, tx := range Txs(b) {
		if tx.Id == targetID {
			return tx
		}
	}
	return nil
}

func calculateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
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

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return calculateDifficulty(b)
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

func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var uTxOuts []*UTxOut
	createTxs := make(map[string]bool)

	for _, block := range Blocks(b) {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {

				if input.Signature == "COINBASE" {
					break
				}

				if FindTx(b, input.TxID).TxOuts[input.Index].Address == address {
					createTxs[input.TxID] = true
				}
			}

			for idx, output := range tx.TxOuts {
				if output.Address == address {
					if _, ok := createTxs[tx.Id]; !ok {
						uTxOut := &UTxOut{tx.Id, idx, output.Amount}

						if isOnMempool(uTxOut) == false {

							uTxOuts = append(uTxOuts, uTxOut)
						}

					}
				}
			}
		}

	}

	return uTxOuts
}

func BalanceByAddress(address string, b *blockchain) int {
	txOuts := UTxOutsByAddress(address, b)
	var amount int

	for _, txout := range txOuts {
		amount += txout.Amount
	}

	return amount
}
