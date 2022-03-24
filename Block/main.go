package main

import (
	"blockchain/cli"
	"blockchain/db"
)

// 블록체인 프로그램을 시작합니다.
// 테스트의 경우 go test -v -coverprofile "cover.out" ./... 를 사용하세요.
// 추적된 테스트는 go tool cover -html="cover.out"
// godoc -http:6060
func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
