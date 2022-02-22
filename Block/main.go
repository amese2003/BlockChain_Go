package main

import (
	blockchain "blockchain/Blockchain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	port        string = ":4000"
	templateDir string = "Templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func Home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockChain().AllBlock()}
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", Home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
