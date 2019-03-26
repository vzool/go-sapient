package sapient

// Sapient interface
type Sapient interface {
	Generate()
	Copy([]byte)
	Load(string)
	Bytes() []byte
	String() string
}
