package sapient

import (
	"crypto/rand"
	"io"
)

// RandomBytes is to generate random bytes according to available buffer
func RandomBytes(buf []byte) {
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err.Error())
	}
}

// RandomNonce is to generate random bytes according to NonceSizeX
func RandomNonce() []byte {
	nonce := make([]byte, NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		panic(err.Error())
	}
	return nonce
}
