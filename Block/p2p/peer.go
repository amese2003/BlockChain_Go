package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	conn    *websocket.Conn
	inbox   chan []byte
	key     string
	address string
	port    string
}

func (p *peer) Close() {
	p.conn.Close()
	delete(Peers, p.key)
}

func (p *peer) read() {
	defer p.Close()
	for {
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("%s", m)
	}
}

func (p *peer) write() {
	defer p.Close()
	for {
		m, ok := <-p.inbox
		if ok == false {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}

}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)

	p := &peer{
		conn:    conn,
		inbox:   make(chan []byte),
		address: address,
		key:     key,
		port:    port,
	}

	go p.read()
	go p.write()

	Peers[key] = p
	return p
}
