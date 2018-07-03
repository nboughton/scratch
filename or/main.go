package main

import (
	"fmt"
)

const (
	ZERO = iota
	ONE
)

func main() {
	if 2 == ZERO|ONE {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
