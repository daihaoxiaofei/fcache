package coder

import (
	"encoding/json"
)

type JsonCode struct {
}

var _ Coder = (*JsonCode)(nil)

func NewJsonCode() *JsonCode {
	return &JsonCode{}
}

func (j *JsonCode) EnCode(i interface{}) ([]byte, error) {
	result, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (j *JsonCode) DeCode(in []byte, out interface{}) (err error) {
	err = json.Unmarshal(in, out)
	if err != nil {
		return err
	}
	return
}
