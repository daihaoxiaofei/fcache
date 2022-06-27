package coder

import (
	"bytes"
	"encoding/gob"
)

// GobCode 这部分..
// 如果选择使用gob序列化则须遵守其基础规则 如未知类型须提前注册 兼容性会差
type GobCode struct {
	buf *bytes.Buffer
	enc *gob.Encoder
	dec *gob.Decoder
}

var _ Coder = (*GobCode)(nil)

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
	err := enc.Encode(i)
	return network.Bytes(), err
}

func (g *GobCode) DeCode(ign []byte, out interface{}) (err error) {
	network := bytes.NewBuffer(ign)
	dec := gob.NewDecoder(network)
	dec.Decode(out)
	return
}
