package fcache

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestRemember(t *testing.T) {
	var result User
	Remember(`ss`, &result, func() interface{} {
		fmt.Println(`重新获取`)
		return User{
			Name: `xiaofei`,
			Age:  13,
		}
	})

	fmt.Println(`result: `, result)
}
