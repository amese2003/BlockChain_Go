package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("테스트 코인입니다!\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API (추천!))\n\n")
	os.Exit(0)
}

func main() {

	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portflag := rest.Int("port", 4000, "포트번호를 정합니다.")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portflag)
		fmt.Println("서버 스타트")
	}

}
