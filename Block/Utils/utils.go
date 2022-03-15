package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
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

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}
