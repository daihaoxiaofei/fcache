package coder

import (
	"encoding/json"
)

type JsonCoder struct {
}

var _ Coder = (*JsonCoder)(nil)

func NewJsonCoder() *JsonCoder {
	return &JsonCoder{}
}

func (j *JsonCoder) EnCode(i interface{}) ([]byte, error) {
	result, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (j *JsonCoder) DeCode(in []byte, out interface{}) (err error) {
	err = json.Unmarshal(in, out)
	if err != nil {
		return err
	}
	return
}
