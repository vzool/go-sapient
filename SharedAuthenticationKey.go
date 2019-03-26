package sapient

import (
	"crypto/hmac"
	"crypto/sha512"
)

// SharedAuthenticationKey type
type SharedAuthenticationKey [SharedAuthenticationKeyBytesSize]byte

// NewSharedAuthenticationKey is way to generate new key in one line
func NewSharedAuthenticationKey(oldKey interface{}) *SharedAuthenticationKey {

	key := &SharedAuthenticationKey{}

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

// Generate new Shared Authentication Key.
// [WARNING] this will overwrite old key!!!
func (key *SharedAuthenticationKey) Generate() {
	RandomBytes(key[:])
}

// Copy other key so you can import any key here,
// but it should be equal to SharedAuthenticationKeyBytesSize
func (key *SharedAuthenticationKey) Copy(otherKey []byte) {

	if len(otherKey) != SharedAuthenticationKeyBytesSize {
		panic("Bytes key is not the correct size")
	}

	copy(key[:], otherKey[:])
}

// Load other key so you can import any key here in encoded format,
// but it should be equal to SharedAuthenticationKeyEncodedSize
func (key *SharedAuthenticationKey) Load(otherKey string) {

	if len(otherKey) != SharedAuthenticationKeyEncodedSize {
		panic("Encoded key is not the correct size")
	}

	decoded, err := Base64UrlDecode(otherKey)

	if err != nil {
		panic("Failed to decode the key:" + err.Error())
	}

	key.Copy(decoded[:])
}

// Bytes return raw bytes value for the key
func (key *SharedAuthenticationKey) Bytes() []byte {
	return key[:]
}

// Value return value of the key
func (key *SharedAuthenticationKey) String() string {
	return Base64UrlEncode(key[:]) + SharedAuthenticationKeyPadding
}

// HmacSha512256 generates a hash of data using HMAC-SHA-512/256. The tag is intended to
// be a natural-language string describing the purpose of the hash, such as
// "hash file for lookup key" or "master secret to client secret".  It serves
// as an HMAC "key" and ensures that different purposes will have different
// hash output. This function is NOT suitable for hashing passwords.
func (key *SharedAuthenticationKey) HmacSha512256(data []byte) []byte {
	h := hmac.New(sha512.New512_256, key[:])
	h.Write(data)
	return h.Sum(nil)
}

// Authenticate an data with a pre-shared key.
func (key *SharedAuthenticationKey) Authenticate(body []byte) (string, string) {
	return HeaderAuthName, Base64UrlEncode(key.HmacSha512256(body))
}

// Verify that the Body-HMAC-SHA512256 header correctly authenticates the
// HTTP Request. Will either return the request given, or throw an
// InvalidMessageException if the signature is invalid. Will also throw a
// HeaderMissingException is there is no Body-HMAC-SHA512256 header.
func (key *SharedAuthenticationKey) Verify(clientAuth string, body []byte) bool {

	_, err := Base64UrlDecode(clientAuth)

	if err != nil {
		panic(err.Error())
	}
	_, serverAuth := key.Authenticate(body)

	return serverAuth == clientAuth && serverAuth != ""
}
