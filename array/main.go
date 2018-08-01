package main

import (
	"fmt"
	"time"
)

var (
	tot = 0
	lim = 100
)

func main() {
	z := time.Now()

	cache := make([]int, lim+1)
	for i := 0; i <= lim; i++ {
		for j := 0; j < i; j++ {
			cache[i] += j
		}
	}

	t := 0
	for _, n := range cache {
		if t == 0 {
			t = n
		}

		t *= n
	}
	fmt.Println(t)

	fmt.Printf("%-3v: %v\n", lim, tot)
	fmt.Println(time.Since(z).Round(time.Millisecond))
}
