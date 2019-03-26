package sapient

const (

	// HeaderAuthName header key for Authenticate with Shared Key
	HeaderAuthName = "Body-HMAC-SHA512256"

	// HeaderSignatureName header key for Authenticate with Public Key
	HeaderSignatureName = "Body-Signature-Ed25519"

	// SealingPublicKeyPadding is padding will be adding to the end of the key
	SealingPublicKeyPadding = "="

	// SealingSecretKeyPadding is padding will be adding to the end of the key
	SealingSecretKeyPadding = "="

	// SharedAuthenticationKeyPadding is padding will be adding to the end of the key
	SharedAuthenticationKeyPadding = "="

	// SharedEncryptionKeyPadding is padding will be adding to the end of the key
	SharedEncryptionKeyPadding = "="

	// SigningPublicKeyPadding is padding will be adding to the end of the key
	SigningPublicKeyPadding = "="

	// SigningSecretKeyPadding is padding will be adding to the end of the key
	SigningSecretKeyPadding = "=="

	// NonceSizeX is the size of the nonce used with the XChaCha20-Poly1305
	// variant of this AEAD, in bytes.
	NonceSizeX = 24

	// SealingPublicKeyBytesSize is a buffer size for sealing public key in bytes
	SealingPublicKeyBytesSize = 32

	// SealingPublicKeyEncodedSize is a buffer size for sealing public encoded key
	SealingPublicKeyEncodedSize = 44

	// SealingSecretKeyBytesSize is a buffer size for sealing secret key in bytes
	SealingSecretKeyBytesSize = 32

	// SealingSecretKeyEncodedSize is a buffer size for sealing secret encoded key
	SealingSecretKeyEncodedSize = 44

	// SharedAuthenticationKeyBytesSize is a buffer size for shared authentication key in bytes
	SharedAuthenticationKeyBytesSize = 32

	// SharedAuthenticationKeyEncodedSize is a buffer size for shared authentication encoded key
	SharedAuthenticationKeyEncodedSize = 44

	// SharedEncryptionKeyBytesSize is a buffer size for shared encryption key in bytes
	SharedEncryptionKeyBytesSize = 32

	// SharedEncryptionKeyEncodedSize is a buffer size for shared encryption encoded key
	SharedEncryptionKeyEncodedSize = 44

	// SigningPublicKeyBytesSize is a buffer size for siging public key in bytes
	SigningPublicKeyBytesSize = 32

	// SigningPublicKeyEncodedSize is a buffer size for siging public encoded key
	SigningPublicKeyEncodedSize = 44

	// SigningSecretKeyBytesSize is a buffer size for signing secret key in bytes
	SigningSecretKeyBytesSize = 64

	// SigningSecretKeyEncodedSize is a buffer size for signing secret encoded key
	SigningSecretKeyEncodedSize = 88
)
