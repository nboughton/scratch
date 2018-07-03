package main

import (
	"fmt"
	"math"
	"math/big"
	"os"

	"github.com/nboughton/numbers/sequence"
	"github.com/nboughton/numbers/set"
	"github.com/nboughton/numbers/slice"
)

var num = set.Int64{1, 3, 7, 9}

func main() {
	f, err := os.Create("truncatable-primes.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sets := set.Int64s{
		{2},
		{3},
		{5},
		{7},
		{9},
	}
	for ln := int64(1); ln < math.MaxInt64; ln++ {
		for n := range sequence.Permutations(num, ln) {
			for _, s := range sets {
				
			}
		}
	}
	/*
		n := big.NewInt(3)
		for {
			fmt.Println("Testing: ", n)
			if n.ProbablyPrime(10) && truncatablePrime(n) {
				fmt.Fprintln(f, n)
			}

			n.Add(n, big.NewInt(2))
		}
	*/
}

func truncatablePrime(n *big.Int) bool {
	s := n.String() // We probably don't want to be converting to string a lot

	for i := range s {
		p, success := big.NewInt(0).SetString(s[:i+1], 10)
		if !success {
			fmt.Println("Could not create bigInt: ", s[:i+1])
		}

		if !p.ProbablyPrime(10) {
			return false
		}
	}

	return true
}
