package p2p

import (
	blockchain "blockchain/Blockchain"
	utils "blockchain/Utils"
	"fmt"
	"net/http"

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
}

func AddPeer(address, port, openPort string, broadcast bool) {

	fmt.Printf("ws://%s:%s/ws?openPort=%s\n", address, port[1:], openPort)

	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port[1:], openPort), nil)
	utils.HandleError(err)
	p := initPeer(conn, address, port)

	if broadcast == false {
		BroadcastNewPeer(p)
		return
	}

	sendNewestBlock(p)
}

func BroadcastNewBlock(b *blockchain.Block) {
	for _, p := range Peers.v {
		notifyNewBlock(b, p)
	}
}

func BroadcastNewTx(tx *blockchain.Tx) {
	for _, p := range Peers.v {
		notifyNewTx(tx, p)
	}
}

func BroadcastNewPeer(newPeer *peer) {
	for key, p := range Peers.v {
		if key != newPeer.key {
			payload := fmt.Sprintf("%s:%s", newPeer.key, p.port)
			notifyNewPeer(payload, p)
		}
	}
}
