package coder

import (
	"encoding/gob"
	"fmt"
	"testing"
)

func Test_gob(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	// 若user类型未注册 会报错
	gob.Register(User{})
	coder := NewGobCode()
	code, err := coder.EnCode(map[string]interface{}{
		`a`: User{
			`asdf`, 15,
		},
		`df`: 65,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(code)
	var a map[string]interface{}
	err = coder.DeCode(code, &a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func Test_gob2(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	coder := NewGobCode()
	code, err := coder.EnCode(User{
		`asdf`,
		15,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(code)
	var a User
	err = coder.DeCode(code, &a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
