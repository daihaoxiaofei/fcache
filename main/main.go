package main

import (
	"fcache"
	"fmt"
)

func main() {
	var result string
	fc := fcache.DefaultFC
	err := fc.Remember(`test23`, &result, func() interface{} {
		return `adfieadfafkl`
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}
