package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(3)
	for {
		if n.ProbablyPrime(10) && sequentialPrime(n) {
			fmt.Println(n)
		}

		n.Add(n, big.NewInt(2))
	}
}

func sequentialPrime(n *big.Int) bool {
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
