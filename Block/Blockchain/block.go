package blockchain

import (
	utils "blockchain/Utils"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

const difficulty int = 2

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

func persistBlock(b *Block) {
	dbStorage.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("블록이.. 없어요!")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *Block) Mine() {
	target := strings.Repeat("0", b.Difficulty)

	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		//fmt.Printf("\n\n\nTarget:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}

}

func FindBlock(hash string) (*Block, error) {
	blockbyte := dbStorage.FindBlock(hash)

	if blockbyte == nil {
		return nil, ErrNotFound
	}

	block := &Block{}
	block.restore(blockbyte)
	return block, nil
}

func createBlock(prevHash string, height int, difficulty int) *Block {
	newBlock := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}

	newBlock.Transactions = Mempool().TxToConfirm()
	newBlock.Mine()
	persistBlock(newBlock)
	return newBlock
}

func Blocks(b *blockchain) []*Block {
	b.m.Lock()
	defer b.m.Unlock()

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

func Status(b *blockchain, rw http.ResponseWriter) {
	b.m.Lock()
	defer b.m.Unlock()
	utils.HandleError(json.NewEncoder(rw).Encode(b))
}
