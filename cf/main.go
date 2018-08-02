package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/nboughton/num"
)

var l = []float64{0.245, 0.045, 1.56, 0.75}

func main() {
	for d := 2.; d <= 10; d++ {
		f := math.Sqrt(d)
		if _, frac := math.Modf(f); frac == 0 {
			continue
		}

		for n := range num.ContinuedFraction(new(big.Rat).SetFloat64(f)) {
			x, y := new(big.Int).Set(n.Frac.Num()), new(big.Int).Set(n.Frac.Denom())
			//if diophantine(x, big.NewInt(int64(d)), y).Cmp(big.NewInt(1)) == 0 {
			fmt.Println(d, n, x, y, diophantine(x, big.NewInt(int64(d)), y))
			//}
		}
	}
}

func diophantine(y, d, x *big.Int) *big.Int {
	x2 := new(big.Int).Mul(x, x)
	y2 := new(big.Int).Mul(y, y)
	Dy2 := new(big.Int).Mul(d, y2)

	return new(big.Int).Sub(x2, Dy2)
}
