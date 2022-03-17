package p2p

import (
	blockchain "blockchain/Blockchain"
	utils "blockchain/Utils"
	"encoding/json"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func (m *Message) addPayload(p interface{}) {
	jsondata, err := json.Marshal(p)
	utils.HandleError(err)
	m.Payload = jsondata
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
	}

	m.addPayload(payload)
	jsonData, err := json.Marshal(m)
	utils.HandleError(err)
	return jsonData
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.BlockChain().NewestHash)
	utils.HandleError(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}
