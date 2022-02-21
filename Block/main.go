package main

import (
	blockchain "blockchain/Blockchain"
	"fmt"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("두번째!")
	chain.AddBlock("세번째!")
	chain.AddBlock("네번째!")

	for _, val := range chain.AllBlock() {
		//fmt.Println("데이터 : %d\n", idx)
		fmt.Printf("데이터 : %s\n", val.Data)
		fmt.Printf("해시 : %s\n", val.Hash)
		fmt.Printf("전 해시 : %s\n", val.PrevHash)
	}
}
