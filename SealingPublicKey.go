package sapient

// SealingPublicKey type
type SealingPublicKey [SealingPublicKeyBytesSize]byte

// NewSealingPublicKey is way to generate new key in one line
func NewSealingPublicKey(oldKey interface{}) *SealingPublicKey {

	key := &SealingPublicKey{}

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

// Generate this is to generate new Sealing Public Key
// [WARNING] this will overwrite old key!!!
func (key *SealingPublicKey) Generate() {
	RandomBytes(key[:])
}

// Copy other key so you can import any key here,
// but it should be equal to SealingPublicKeyBytesSize
func (key *SealingPublicKey) Copy(otherKey []byte) {

	if len(otherKey) != SealingPublicKeyBytesSize {
		panic("Bytes key is not the correct size")
	}

	copy(key[:], otherKey[:])
}

// Load other key so you can import any key here in encoded format,
// but it should be equal to SealingPublicKeyEncodedSize
func (key *SealingPublicKey) Load(otherKey string) {

	if len(otherKey) != SealingPublicKeyEncodedSize {
		panic("Encoded key is not the correct size")
	}

	decoded, err := Base64UrlDecode(otherKey)

	if err != nil {
		panic("Failed to decode the key:" + err.Error())
	}

	key.Copy(decoded[:])
}

// Bytes return raw bytes value for the key
func (key *SealingPublicKey) Bytes() []byte {
	return key[:]
}

// Value return value of the key
func (key *SealingPublicKey) String() string {
	return Base64UrlEncode(key[:]) + SealingPublicKeyPadding
}
