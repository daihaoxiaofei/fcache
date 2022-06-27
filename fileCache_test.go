package fcache

import (
	"os"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestRemember(t *testing.T) {
	os.Remove(`cache/ss.db`)
	var result User
	err := Remember(`ss`, &result, func() interface{} {
		return User{
			Name: `xiaofei`,
			Age:  13,
		}
	})
	if err != nil || result.Name != `xiaofei` {
		t.Fatal(`err: Remember first`, err)
	}

	err = Remember(`ss`, &result, func() interface{} {
		return User{
			Name: `xiaofei`,
			Age:  31,
		}
	})
	if err != nil || result.Age != 13 {
		t.Fatal(`err: Remember too`, result, err)
	}

	Remove(`ss`)
	err = Remember(`ss`, &result, func() interface{} {
		return User{
			Name: `xiaofei`,
			Age:  13,
		}
	})
	if err != nil || result.Age != 13 {
		t.Fatal(`err: Remember three`, err)
	}

}
