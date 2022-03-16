package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	conn  *websocket.Conn
	inbox chan []byte
}

func (p *peer) read() {
	for {
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("%s", m)
	}
}

func (p *peer) write() {

}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	p := &peer{
		conn,
		make(chan []byte, 5),
	}

	go p.read()
	go p.write()
	key := fmt.Sprintf("%s:%s", address, port)
	Peers[key] = p
	return p
}
