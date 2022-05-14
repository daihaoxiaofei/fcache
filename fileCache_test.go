package fcache

import (
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestRemember(t *testing.T) {
	var result User
	Remember(`ss`, &result, func() interface{} {
		return User{
			Name: `xiaofei`,
			Age:  13,
		}
	})
	if result.Name != `xiaofei` {
		t.Fatal(`err: Remember first`)
	}

	Remember(`ss`, &result, func() interface{} {
		return User{
			Name: `xiaofei`,
			Age:  31,
		}
	})
	if result.Age != 13 {
		t.Fatal(`err: Remember too`)
	}

	Remome(`ss`)
	Remember(`ss`, &result, func() interface{} {
		return User{
			Name: `xiaofei`,
			Age:  31,
		}
	})
	if result.Age != 31 {
		t.Fatal(`err: Remember three`)
	}


}
