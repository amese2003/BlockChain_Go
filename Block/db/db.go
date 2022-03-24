package db

import (
	utils "blockchain/Utils"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

const (
	dbName       = "blockchain"
	dataBucket   = "data"
	blocksBucket = "blocks"

	checkpoint = "checkpoint"
)

var db *bolt.DB

type DB struct{}

func (DB) FindBlock(hash string) []byte {
	return findBlock(hash)
}
func (DB) LoadChain() []byte {
	return loadChain()
}
func (DB) SaveBlock(hash string, data []byte) {
	saveBlock(hash, data)
}
func (DB) SaveChain(data []byte) {
	saveChain(data)
}

func (DB) DeleteAllBlocks() {
	emptyBlocks()
}

func InitDB() {
	if db == nil {
		dbPointer, err := bolt.Open(getDBName(), 0600, nil)
		db = dbPointer
		utils.HandleError(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleError(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleError(err)
	}
}

func getDBName() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

func saveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleError(err)
}

func saveChain(data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})

	utils.HandleError(err)
}

func loadChain() []byte {
	var data []byte
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})

	return data
}

func findBlock(hash string) []byte {
	var data []byte
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})

	return data
}

func Close() {
	db.Close()
}

func emptyBlocks() {
	db.Update(func(t *bolt.Tx) error {
		utils.HandleError(t.DeleteBucket([]byte(blocksBucket)))
		_, err := t.CreateBucket([]byte(blocksBucket))
		utils.HandleError(err)
		return nil
	})
}
