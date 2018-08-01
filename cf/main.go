package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/nboughton/num"
)

var l = []float64{0.245, 0.045, 1.56, 0.75}

func main() {
	for n := range num.ContinuedFraction(new(big.Rat).SetFloat64(math.Sqrt2)) {
		fmt.Println(n)
	}
}
