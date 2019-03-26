package sapient

import (
	"encoding/base64"

	"golang.org/x/crypto/chacha20poly1305"
)

// SharedEncryptionKey type
type SharedEncryptionKey [SharedEncryptionKeyBytesSize]byte

// NewSharedEncryptionKey is way to generate new key in one line
func NewSharedEncryptionKey(oldKey interface{}) *SharedEncryptionKey {

	key := &SharedEncryptionKey{}

	if oldKey == nil {

		key.Generate()

	} else {

		switch ek := oldKey.(type) {

		case string:

			key.Load(ek)

		case []byte:

			key.Copy(ek[:])

		default:

			panic("Falied to load the key")
		}
	}

	return key
}

// Generate a new Shared Encryption Key.
// [WARNING] this will overwrite old key!!!
func (key *SharedEncryptionKey) Generate() {
	RandomBytes(key[:])
}

// Copy other key so you can import any key here,
// but it should be equal to SharedEncryptionKeyBytesSize
func (key *SharedEncryptionKey) Copy(otherKey []byte) {

	if len(otherKey) != SharedEncryptionKeyBytesSize {
		panic("Key is not the correct size")
	}

	copy(key[:], otherKey[:])
}

// Load other key so you can import any key here in encoded format,
// but it should be equal to SharedEncryptionKeyEncodedSize
func (key *SharedEncryptionKey) Load(otherKey string) {

	if len(otherKey) != SharedEncryptionKeyEncodedSize {
		panic("Encoded key is not the correct size")
	}

	decoded, err := Base64UrlDecode(otherKey)

	if err != nil {
		panic("Failed to decode the key:" + err.Error())
	}

	key.Copy(decoded[:])
}

// Bytes return raw bytes value for the key
func (key *SharedEncryptionKey) Bytes() []byte {
	return key[:]
}

// Value return value of the key
func (key *SharedEncryptionKey) String() string {
	return Base64UrlEncode(key[:]) + SharedEncryptionKeyPadding
}

// Encrypt a data with a pre-shared key.
func (key *SharedEncryptionKey) Encrypt(body []byte) string {

	nonce := RandomNonce()

	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		panic(err.Error())
	}

	return Base64UrlEncode(append(
		nonce[:],
		aead.Seal(nil, nonce, body, nil)...,
	))
}

// Decrypt a data with a pre-shared key.
func (key *SharedEncryptionKey) Decrypt(body string) []byte {

	decoded, err := base64.RawURLEncoding.DecodeString(body)

	if err != nil {
		panic(err.Error())
	}

	nonce, cipher := decoded[:NonceSizeX], decoded[NonceSizeX:]

	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		panic(err.Error())
	}

	// Decryption.
	plaintext, err := aead.Open(nil, nonce, cipher, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}
