package blockchain

import (
	utils "blockchain/Utils"
	"blockchain/wallet"
	"errors"
	"sync"
	"time"
)

const (
	minerReward int = 50
)

var ErrNotEnough = errors.New("Not enough Coin")
var ErrNotValid = errors.New("Tx Invalid")

var m *mempool = &mempool{}
var memOnce sync.Once

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	TxID      string `json:"txId"`
	Index     int    `json:"index"`
	Signature string `json:"signature"`
}

type TxOut struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

type UTxOut struct {
	TxID   string `json:"txId"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

type mempool struct {
	Txs map[string]*Tx
	m   sync.Mutex
}

func Mempool() *mempool {
	memOnce.Do(func() {
		m = &mempool{
			Txs: make(map[string]*Tx),
		}
	})

	return m
}

func (t *Tx) sign() {
	for _, txIn := range t.TxIns {
		txIn.Signature = wallet.Sign(t.Id, wallet.Wallet())
	}
}

func validate(tx *Tx) bool {
	valid := true

	for _, txIn := range tx.TxIns {
		prevTx := FindTx(BlockChain(), txIn.TxID)
		if prevTx == nil {
			valid = false
			break
		}

		address := prevTx.TxOuts[txIn.Index].Address
		valid = wallet.Verify(txIn.Signature, tx.Id, address)

		if valid == false {
			break
		}
	}

	return valid
}

func isOnMempool(UTxOut *UTxOut) bool {
	exists := false

Outloop:
	for _, tx := range Mempool().Txs {
		for _, input := range tx.TxIns {
			if input.TxID == UTxOut.TxID && input.Index == UTxOut.Index {
				exists = true
				break Outloop
			}
		}
	}

	return exists
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
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

	if BalanceByAddress(from, BlockChain()) < amount {
		return nil, ErrNotEnough
	}

	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	UTxOuts := UTxOutsByAddress(from, BlockChain())

	for _, uTxOut := range UTxOuts {
		if total >= amount {
			break
		}

		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}

	if change := total - amount; change != 0 {
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
	tx.sign()
	valid := validate(tx)

	if valid == false {
		return nil, ErrNotValid
	}

	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) (*Tx, error) {
	tx, err := makeTx(wallet.Wallet().Address, to, amount)
	if err != nil {
		return nil, err
	}
	m.Txs[tx.Id] = tx
	return tx, nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx(wallet.Wallet().Address)
	var txs []*Tx

	for _, tx := range m.Txs {
		txs = append(txs, tx)
	}
	txs = append(txs, coinbase)
	m.Txs = make(map[string]*Tx)
	return txs
}

func (m *mempool) AddPeerTx(tx *Tx) {
	m.m.Lock()
	defer m.m.Unlock()

	m.Txs[tx.Id] = tx
}
