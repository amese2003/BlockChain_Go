package rest

import (
	blockchain "blockchain/Blockchain"
	utils "blockchain/Utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
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
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See A Block",
		},
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOut for an Address",
		},
	}

	utils.HandleError(json.NewEncoder(rw).Encode(data))
}

func BlockPage(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		utils.HandleError(json.NewEncoder(rw).Encode(blockchain.BlockChain().Blocks()))

	case "POST":
		blockchain.BlockChain().AddBlock()
		rw.WriteHeader(http.StatusCreated)
	}
}

func Block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["hash"]
	block, err := blockchain.FindBlock(id)
	encoder := json.NewEncoder(rw)

	if err == blockchain.ErrNotFound {
		utils.HandleError(encoder.Encode(errorResponse{fmt.Sprint(err)}))
	} else {
		utils.HandleError(encoder.Encode(block))
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func status(rw http.ResponseWriter, r *http.Request) {
	utils.HandleError(json.NewEncoder(rw).Encode(blockchain.BlockChain()))
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")

	switch total {
	case "true":
		amount := blockchain.BlockChain().BalanceByAddress(address)
		json.NewEncoder(rw).Encode(balanceResponse{address, amount})
		break

	default:
		utils.HandleError(json.NewEncoder(rw).Encode(blockchain.BlockChain().TxOutsByAddress(address)))
		break
	}

}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status)
	router.HandleFunc("/blocks", BlockPage).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", Block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
