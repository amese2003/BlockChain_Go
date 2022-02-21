package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "안녕 안녕")
}

func main() {
	http.HandleFunc("/", Home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
