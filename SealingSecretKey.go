package sapient

// SealingSecretKey type
type SealingSecretKey [SealingSecretKeyBytesSize]byte

// NewSealingSecretKey is way to generate new key in one line
func NewSealingSecretKey(oldKey interface{}) *SealingSecretKey {

	key := &SealingSecretKey{}

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

// Generate this is to generate new Sealing Secret Key
// [WARNING] this will overwrite old key!!!
func (key *SealingSecretKey) Generate() {
	RandomBytes(key[:])
}

// Copy other key so you can import any key here,
// but it should be equal to SealingSecretKeyBytesSize
func (key *SealingSecretKey) Copy(otherKey []byte) {

	if len(otherKey) != SealingSecretKeyBytesSize {
		panic("Bytes key is not the correct size")
	}

	copy(key[:], otherKey[:])

}

// Load other key so you can import any key here in encoded format,
// but it should be equal to SealingSecretKeyEncodedSize
func (key *SealingSecretKey) Load(otherKey string) {

	if len(otherKey) != SealingSecretKeyEncodedSize {
		panic("Encoded key is not the correct size")
	}

	decoded, err := Base64UrlDecode(otherKey)

	if err != nil {
		panic("Failed to decode the key:" + err.Error())
	}

	key.Copy(decoded[:])
}

// Bytes return raw bytes value for the key
func (key *SealingSecretKey) Bytes() []byte {
	return key[:]
}

// Value return value of the key
func (key *SealingSecretKey) String() string {
	return Base64UrlEncode(key[:]) + SealingSecretKeyPadding
}
