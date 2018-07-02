package main

import (
	"fmt"

	"github.com/nboughton/numbers/sequence"
)

func main() {
	for n := range sequence.PellLucas(10) {
		fmt.Printf("%s/%s\n", n[0], n[1])
	}
}
