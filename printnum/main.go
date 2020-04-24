package main

import (
	"fmt"
	"strings"
)

var oneToSix = []string{"1", "2", "3", "4", "5", "6"}

func main() {
	o := []string{}

	for _, h := range oneToSix {
		for _, t := range oneToSix {
			for _, d := range oneToSix {
				o = append(o, h+t+d)
			}
		}
	}

	fmt.Println(strings.Join(o, ", "))
}