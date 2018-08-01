package main

import (
	"fmt"
	"math/big"

	"github.com/nboughton/num"
)

func main() {
	fmt.Println(num.DecToFrac(0.245))
	/*
		a := 649.
		b := 200.

		c := a / b

		i, frac := math.Modf(c)
		fmt.Println(a, b, c, i, frac)

		var (
			i2    int
			frac2 float64
		)
		fmt.Sscanf(fmt.Sprint(c), "%d%f", &i2, &frac2)
		fmt.Println(a, b, c, i2, frac2)

		//r := big.NewRat(649, 200)
		fmt.Println(big.NewRat(3, 3).Float64())
	*/
}

func continuedFractions(f *num.Frac) []int64 {
	var (
		res []int64
		cf  func(r *num.Frac)
	)

	cf = func(r *num.Frac) {
		// This is a hack because math.Modf is weird
		var (
			i    int
			frac float64
		)
		fmt.Sscanf(fmt.Sprint(n), "%d%f", &i, &frac)
		if frac == 0 {
			return
		}

		fmt.Printf("%d; %.3f: %v\n", int64(i), frac, simplify(r))

		res = append(res, int64(i))

		next := new(big.Rat).SetFloat64(frac) // SetFloat64(frac)
		cf(next.Inv(simplify(next)))
	}

	cf(f)

	return res
}
