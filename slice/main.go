package main

import (
	"fmt"
)

var (
	t   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lim = 200
)

func main() {
	set := []int{0, 0, 0, 0, 0, 0, 0}
	ans := 1
	for set[0] = 0; sum(set[:1]) <= lim; set[0]++ { // 1s
		for set[1] = 0; sum(set[:2]) <= lim; set[1] += 2 { // 2s
			for set[2] = 0; sum(set[:3]) <= lim; set[2] += 5 { // 5s
				for set[3] = 0; sum(set[:4]) <= lim; set[3] += 10 { // 10s
					for set[4] = 0; sum(set[:5]) <= lim; set[4] += 20 { // 20s
						for set[5] = 0; sum(set[:6]) <= lim; set[5] += 50 { // 50s
							for set[6] = 0; sum(set) <= lim; set[6] += 100 { // 100s
								//fmt.Println(set)
								if sum(set) == lim {
									ans++
									fmt.Printf("%-3v %v\n", set, ans)
								}
							}
						}
					}
				}
			}
		}
	}
}

func sum(set []int) (total int) {
	for _, n := range set {
		total += n
	}

	return
}
