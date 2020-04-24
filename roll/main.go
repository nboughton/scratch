package main

import (
	"fmt"
	"log"

	"github.com/nboughton/go-roll"
)

func main() {
	s := "3d6Kh2"

	// Roll some dice
	var results roll.Results
	for i := 0; i < 10000; i++ {
		result, err := roll.FromString(s)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.Ints())

		results = append(results, result)
	}

	var (
		rs  = results[0]
		min = len(rs.Ints()) * rs.Die().Min().N
		max = len(rs.Ints()) * rs.Die().Max().N
	)

	fmt.Println(min, max)
}
