package sapient

import (
	"golang.org/x/crypto/ed25519"
)

// SigningPublicKey type
type SigningPublicKey [SigningPublicKeyBytesSize]byte

// NewSigningPublicKey is way to generate new key in one line
func NewSigningPublicKey(oldKey interface{}) *SigningPublicKey {

	key := &SigningPublicKey{}

	if oldKey == nil {

		key.Generate()

	} else {

		switch pk := oldKey.(type) {

		case string:

			key.Load(pk)

		case []byte:

			key.Copy(pk[:])

		default:

			panic("Falied to load the key")
		}
	}

	return key
}

// Generate this is to generate new Sealing Public Key
// [WARNING] this will overwrite old key!!!
func (key *SigningPublicKey) Generate() {
	RandomBytes(key[:])
}

// Verify data signature
func (key *SigningPublicKey) Verify(message, signature []byte) bool {

	return ed25519.Verify(key[:], message, signature)
}

// Copy other key so you can import any key here,
// but it should be equal to SigningPublicKeyBytesSize
func (key *SigningPublicKey) Copy(otherKey []byte) {

	if len(otherKey) != SigningPublicKeyBytesSize {
		panic("Bytes key is not the correct size")
	}

	copy(key[:], otherKey[:])
}

// Load other key so you can import any key here in encoded format,
// but it should be equal to SigningPublicKeyEncodedSize
func (key *SigningPublicKey) Load(otherKey string) {

	if len(otherKey) != SigningPublicKeyEncodedSize {
		panic("Encoded key is not the correct size")
	}

	decoded, err := Base64UrlDecode(otherKey)

	if err != nil {
		panic("Failed to decode the key:" + err.Error())
	}

	key.Copy(decoded[:])
}

// Bytes return raw bytes value for the key
func (key *SigningPublicKey) Bytes() []byte {
	return key[:]
}

// Value return value of the key
func (key *SigningPublicKey) String() string {
	return Base64UrlEncode(key[:]) + SigningPublicKeyPadding
}
