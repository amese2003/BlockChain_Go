package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type BlockChain struct {
	blocks []block
}

var b *BlockChain

func GetBlockChain() *BlockChain {
	if b == nil {
		b = &BlockChain{}
	}

	return b
}
