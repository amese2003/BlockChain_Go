package p2p

import (
	utils "blockchain/Utils"
	"fmt"
	"net/http"
	"time"

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
		if err != nil {
			break
		}
		fmt.Printf("Just got:%s\n\n", p)
		time.Sleep(5 * time.Second)
		message := fmt.Sprintf("New message: %s", p)
		utils.HandleError(conn.WriteMessage(websocket.TextMessage, []byte(message)))
	}
}
