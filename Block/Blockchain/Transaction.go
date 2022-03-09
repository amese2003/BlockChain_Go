package blockchain

import (
	utils "blockchain/Utils"
	"errors"
	"time"
)

const (
	minerReward int = 50
)

var ErrNotEnough = errors.New("Not enough Coin")

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}

	txOuts := []*TxOut{
		{address, minerReward},
	}

	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}

	tx.getId()
	return &tx
}

func makeTx(from, to string, amount int) (*Tx, error) {
	if BlockChain().BalanceByAddress(from) < amount {
		return nil, ErrNotEnough
	}

	var txIns []*TxIn
	var txOuts []*TxOut

	total := 0
	oldTxOuts := BlockChain().TxOutsByAddress(from)

	for _, txOut := range oldTxOuts {
		if total > amount {
			break
		}

		txIn := &TxIn{txOut.Owner, txOut.Amount}
		txIns = append(txIns, txIn)
		total += txOut.Amount
	}

	change := total - amount

	if change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}

	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}

	tx.getId()

	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("nero", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("nero")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
