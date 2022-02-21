package main

import (
	blockchain "blockchain/Blockchain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func Home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Templates/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockChain().AllBlock()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", Home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
