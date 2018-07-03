package main

import (
	"fmt"

	"github.com/nboughton/num"
)

//  1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

func main() {
	for i := range num.SeqFibonacci() {
		if i.Int64() > 100 {
			break
		}

		fmt.Println(i)
	}
}
