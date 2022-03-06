package main

import (
	blockchain "blockchain/Blockchain"
	"blockchain/cli"
)

func main() {
	blockchain.BlockChain()
	cli.Start()
}
