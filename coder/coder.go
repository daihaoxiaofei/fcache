package coder

type Coder interface {
	EnCode(i interface{}) ([]byte, error)
	DeCode(in []byte, out interface{}) (err error)
}
