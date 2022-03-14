package p2p

import (
	utils "blockchain/Utils"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleError(err)

	for {
		_, p, err := conn.ReadMessage()
		utils.HandleError(err)
		fmt.Printf("%s\n\n", p)
	}
}
