package rest

import (
	blockchain "blockchain/Blockchain"
	utils "blockchain/Utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "테스트 JSON",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func BlockPage(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlock())
		break
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		fmt.Println(addBlockBody)
		blockchain.GetBlockChain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
		break
	}
}

func Block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleError(err)

	block, err := blockchain.GetBlockChain().GetBlock(id)
	encoder := json.NewEncoder(rw)

	if err != nil {
		if err == blockchain.ErrNotFound {
			encoder.Encode(errorResponse{fmt.Sprint(err)})
		}
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", BlockPage).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", Block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}