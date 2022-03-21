package main

import (
	"blockchain/cli"
	"blockchain/db"
)

// 블록체인 프로그램을 시작합니다.
func main() {
	defer db.Close()
	cli.Start()
}
