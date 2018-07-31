package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

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
}

func continuedFractions(f *big.Rat) []int64 {
	var (
		res []int64
		cf  func(r *big.Rat)
	)

	cf = func(r *big.Rat) {
		n, _ := r.Float64()

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

func simplify(r *big.Rat) *big.Rat {
	gcd := new(big.Int).GCD(nil, nil, r.Num(), r.Denom())

	a := new(big.Int).Div(r.Num(), gcd)
	b := new(big.Int).Div(r.Denom(), gcd)

	return new(big.Rat).SetFrac(a, b)
}
