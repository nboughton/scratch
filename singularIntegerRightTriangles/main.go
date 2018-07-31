package main

import (
	"fmt"
	//"math"
	"sort"
	"time"

	"github.com/nboughton/num"
)

var (
	lim = num.Int(1500000)
)

func main() {
	z := time.Now()

	var f num.Set
	for n := range factors(lim) {
		fmt.Println(n)
		f = append(f, n)
	}
	sort.Sort(f)

	fmt.Println(len(f))
	fmt.Println(time.Since(z).Round(time.Millisecond))
}

func factors(lim num.Int) chan num.Int {
	c := make(chan num.Int)

	go func() {
		defer close(c)

		m := make(map[num.Int]bool)
		for r := num.Int(2); (r*r)/2 < lim; r++ {
			f := num.Int((r * r) / 2).Divisors()
			f = f[1 : len(f)-1]

			for _, n := range f {
				if _, ok := m[n]; !ok {
					c <- n
					m[n] = true
				}
			}
		}
	}()

	return c
}
