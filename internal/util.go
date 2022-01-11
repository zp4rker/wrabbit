package wrabbit

import (
	"crypto/rand"
	"fmt"
)

func RandToken() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
