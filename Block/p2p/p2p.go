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

	openport := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openport != "" && ip != ""
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleError(err)
	initPeer(conn, ip, openport)

	time.Sleep(20 * time.Second)
	conn.WriteMessage(websocket.TextMessage, []byte("3000번이다!"))
}

func AddPeer(address, port, openPort string) {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port[1:], openPort), nil)
	utils.HandleError(err)
	initPeer(conn, address, port)
	time.Sleep(10 * time.Second)
	conn.WriteMessage(websocket.TextMessage, []byte("4000번이다!"))
}
