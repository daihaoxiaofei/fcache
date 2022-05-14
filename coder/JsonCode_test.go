package coder

import (
	"fmt"
	"testing"
)

func TestNewJsonCode(t *testing.T) {
	gc := NewJsonCode()
	by, err := gc.EnCode(`sasdffffffffff`)
	if err != nil {
		t.Fatal(`err: gc.EnCode: `, err)
	}

	fmt.Println(`by:`, string(by))

	var ss string
	gc.DeCode(by, &ss)

	fmt.Println(`ss:`, ss)
}
