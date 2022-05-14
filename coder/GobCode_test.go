package coder

import (
	"fmt"
	"testing"
)

func TestNewGobCode(t *testing.T) {
	data:=`sasdffffffffff`
	gc := NewGobCode()
	by, err := gc.EnCode(data)
	if err != nil {
		t.Fatal(`err: gc.EnCode: `, err)
	}

	fmt.Println(`by:`, string(by))

	var ss string
	gc.DeCode(by, &ss)

	fmt.Println(`ss:`, ss)
}
