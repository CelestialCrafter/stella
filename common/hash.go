package common

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

func Hash() []byte {
	b := make([]byte, 4)
	timestamp := time.Now().UnixNano()
	binary.BigEndian.PutUint32(b, uint32(timestamp))
	hash := sha256.Sum256(b)
	return hash[:]
}
