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

func DB() *bolt.DB {
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

	return db
}

func getDBName() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleError(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})

	utils.HandleError(err)
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})

	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})

	return data
}

func Close() {
	DB().Close()
}
