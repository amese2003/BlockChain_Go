package cli

import (
	rest "blockchain/Rest"
	"blockchain/explorer"
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("테스트 코인입니다!\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "포트번호를 정합니다.")
	mode := flag.String("mode", "rest", "'html'이랑 'rest' 둘 중 하나")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
