package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "테스트 JSON",
		},
		{
			URL:         URL("/home"),
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
