package coder

import (
	"bytes"
	"encoding/gob"
)

// 这部分未完成
type GobCode struct {
	buf *bytes.Buffer
	enc *gob.Encoder // Will write to network.
	dec *gob.Decoder // Will read from network.
}

// var _ Coder = (*GobCode)(nil)

func NewGobCode() *GobCode {
	GobC := &GobCode{
		buf: &bytes.Buffer{},
	}
	GobC.enc = gob.NewEncoder(GobC.buf)
	GobC.dec = gob.NewDecoder(GobC.buf)
	return GobC
}

func (g *GobCode) EnCode(i interface{}) ([]byte, error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(i)
	return network.Bytes(), nil
}

func (g *GobCode) DeCode(ign []byte, out interface{}) (err error) {
	var network bytes.Buffer
	dec := gob.NewDecoder(&network)
	dec.Decode(out)
	return
}
