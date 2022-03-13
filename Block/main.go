package main

import (
	"blockchain/cli"
	"blockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
