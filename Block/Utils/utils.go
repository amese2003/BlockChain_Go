package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func ToBytes(data interface{}) []byte {
	var blockbuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockbuffer)
	HandleError(encoder.Encode(data))
	return blockbuffer.Bytes()
}

func FromBytes(out interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleError(encoder.Decode(out))
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
