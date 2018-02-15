package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8}
	m := make(map[int]int)
	for _, v := range i {
		m[v]++
	}

	fmt.Println(m)
}
