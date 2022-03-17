package p2p

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type peers struct {
	v map[string]*peer
	m sync.Mutex
}

var Peers peers = peers{
	v: make(map[string]*peer),
}

type peer struct {
	conn    *websocket.Conn
	inbox   chan []byte
	key     string
	address string
	port    string
}

func Allpeer(p *peers) []string {
	p.m.Lock()
	defer p.m.Unlock()

	var keys []string
	for key := range p.v {
		keys = append(keys, key)
	}

	return keys
}

func (p *peer) Close() {
	Peers.m.Lock()
	defer func() {
		time.Sleep(20 * time.Second)
		Peers.m.Unlock()
	}()

	p.conn.Close()
	delete(Peers.v, p.key)
}

func (p *peer) read() {
	defer p.Close()
	for {
		m := Message{}
		err := p.conn.ReadJSON(&m)
		if err != nil {
			break
		}

		handleMsg(&m, p)
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
	Peers.m.Lock()
	defer Peers.m.Unlock()

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

	Peers.v[key] = p
	return p
}
