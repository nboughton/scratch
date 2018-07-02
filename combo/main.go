package main

import (
	"fmt"

	"github.com/nboughton/numbers/sequence"
)

func main() {
	var t []int64
	for i := range sequence.Ints() {
		if i > 10 {
			break
		}

		t = append(t, i)
	}

	for c := range sequence.Permutations(t, 5) {
		fmt.Println(c)
	}
}
