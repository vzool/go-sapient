package sapient

import (
	"golang.org/x/crypto/ed25519"
)

// SigningSecretKey type
type SigningSecretKey [SigningSecretKeyBytesSize]byte

// NewSigningSecretKey is way to generate new key in one line
func NewSigningSecretKey(oldKey interface{}) *SigningSecretKey {

	key := &SigningSecretKey{}

	if oldKey == nil {

		key.Generate()

	} else {

		switch sk := oldKey.(type) {

		case string:

			key.Load(sk)

		case []byte:

			key.Copy(sk[:])

		default:

			panic("Falied to load the key")
		}
	}

	return key
}

// Sign data and generate signature
func (key *SigningSecretKey) Sign(message []byte) []byte {

	return ed25519.Sign(key[:], message)
}

// PublicKey is to be generated from this key and it will return Signing Public Key
func (key *SigningSecretKey) PublicKey() *SigningPublicKey {

	publicKey := &SigningPublicKey{}

	sk := ed25519.NewKeyFromSeed(key[:SigningPublicKeyBytesSize])
	publicKey.Copy(sk.Public().(ed25519.PublicKey))

	return publicKey
}

// Generate this is to generate new Signing Secret Key
// [WARNING] this will overwrite old key!!!
func (key *SigningSecretKey) Generate() {
	RandomBytes(key[:])
}

// Copy other key so you can import any key here,
// but it should be equal to SigningSecretKeyBytesSize
func (key *SigningSecretKey) Copy(otherKey []byte) {

	if len(otherKey) != SigningSecretKeyBytesSize {
		panic("Bytes key is not the correct size")
	}

	copy(key[:], otherKey[:])
}

// Load other key so you can import any key here in encoded format,
// but it should be equal to SigningSecretKeyEncodedSize
func (key *SigningSecretKey) Load(otherKey string) {

	if len(otherKey) != SigningSecretKeyEncodedSize {
		panic("Encoded key is not the correct size")
	}

	decoded, err := Base64UrlDecode(otherKey)

	if err != nil {
		panic("Failed to decode the key:" + err.Error())
	}

	key.Copy(decoded[:])
}

// Bytes return raw bytes value for the key
func (key *SigningSecretKey) Bytes() []byte {
	return key[:]
}

// Value return value of the key
func (key *SigningSecretKey) String() string {
	return Base64UrlEncode(key[:]) + SigningSecretKeyPadding
}
