package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "테스트 JSON",
		},
		{
			URL:         "/home",
			Method:      "GET",
			Description: "테스트 JSON22",
			Payload:     "지긋지긋한 payload",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

	// b, err := json.Marshal(data)
	// utils.HandleError(err)

	// fmt.Printf("%s\n", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
