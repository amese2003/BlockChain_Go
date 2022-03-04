package main

import (
	blockchain "blockchain/Blockchain"
)

func main() {
	blockchain.BlockChain().AddBlock("First")
	blockchain.BlockChain().AddBlock("Second")
	blockchain.BlockChain().AddBlock("Third")
}
