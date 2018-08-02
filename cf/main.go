package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/nboughton/num"
)

var l = []float64{0.245, 0.045, 1.56, 0.75}

func main() {
	for d := 3.; d <= 1000; d++ {
		f := math.Sqrt(d)
		if _, frac := math.Modf(f); frac == 0 {
			continue
		}

		for n := range num.ContinuedFraction(new(big.Rat).SetFloat64(f)) {
			//fmt.Println(n.Frac.Num(), big.NewInt(int64(d)), n.Frac.Denom(), diophantine(n.Frac.Denom(), big.NewInt(int64(d)), n.Frac.Num()))
			fmt.Println(n)
		}
		break
	}
}

func diophantine(x, d, y *big.Int) *big.Int {
	x2 := new(big.Int).Mul(x, x)
	y2 := new(big.Int).Mul(y, y)
	Dy2 := new(big.Int).Mul(d, y2)

	return new(big.Int).Sub(x2, Dy2)
}
