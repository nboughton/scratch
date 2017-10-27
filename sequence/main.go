package main

import (
	"fmt"

	"github.com/nboughton/numbers/sequence"
)

func main() {
	for n := range sequence.Triangles() {
		if n > 32 {
			break
		}
		fmt.Println(n)
	}
}
