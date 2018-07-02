package main

import (
	"fmt"
	"time"

	"github.com/nboughton/numbers/sequence"
	"github.com/nboughton/numbers/set"
	"github.com/nboughton/numbers/slice"
)

func main() {
	n, l, m := set.Int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, make(map[int64]int)

	t1 := time.Now()
	for prm := range n.Permutations(l) {
		t := slice.ToInt64(prm)
		if _, ok := m[t]; !ok {
			m[t] = 1
		} else {
			m[t]++
		}
	}
	fmt.Println(time.Since(t1).Round(time.Millisecond))
	//m = nil

	t2 := time.Now()
	for prm := range sequence.Permutations(n, int64(l)) {
		t := slice.ToInt64(prm)
		if _, ok := m[t]; !ok {
			fmt.Printf("%d not found in n.Permutations()\n", t)
		}
	}
	fmt.Println(time.Since(t2).Round(time.Millisecond))
	//m = nil

	t3 := time.Now()
	for cmb := range n.Combinations(l) {
		for prm := range cmb.Prm() {
			t := slice.ToInt64(prm)
			if _, ok := m[t]; !ok {
				fmt.Printf("%d not found in n.Permutations()\n", t)
			}
		}
	}
	fmt.Println(time.Since(t3).Round(time.Millisecond))
}
