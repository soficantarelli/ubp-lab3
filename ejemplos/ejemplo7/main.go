package main

import (
	"arrays"
	"fmt"
)

func main() {
	x, e := arrays.Get(2)

	if e != nil {
		fmt.Println(e.Error())
		exit(1)
	}
}
