package p2p

import (
	blockchain "blockchain/Blockchain"
	utils "blockchain/Utils"
	"encoding/json"
	"fmt"
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
		Kind:    kind,
		Payload: utils.ToJson(payload),
	}

	return utils.ToJson(m)
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.BlockChain().NewestHash)
	utils.HandleError(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var payload blockchain.Block
		utils.HandleError(json.Unmarshal(m.Payload, &payload))
		fmt.Println(payload)
	}
}
