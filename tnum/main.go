package main

import (
	"fmt"

	"github.com/nboughton/num"
)

func main() {
	i := num.Set{
		10,
		30,
		210,
		2310,
		30030,
		510510,
		9699690,
	}

	for n := range num.Seq(num.TRIANGLE) {
		if n > 9699690 {
			break
		}

		fmt.Println(i.Contains(n))
	}
}
