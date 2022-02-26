package main

import (
	rest "blockchain/Rest"
	"blockchain/explorer"
)

func main() {
	go explorer.Start(5000)
	rest.Start(4000)
}
