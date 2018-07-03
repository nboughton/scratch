package main

import (
	"fmt"

	"github.com/nboughton/numbers/set"
)

func main() {
	n := newNumberSpiral(3)
	for _, row := range n {
		fmt.Println(row)
	}
}

func newNumberSpiral(size int64) set.Int64s {
	if size%2 == 0 {
		size++
	}

	grid := make(set.Int64s, size)
	for row := range grid {
		grid[row] = make(set.Int64, size)
	}
	grid = set.Int64s{
		{9, 2, 3},
		{8, 1, 4},
		{7, 6, 5},
	}

	// Starting from the center head up 1...
	crd := set.Coord{Row: size / 2, Col: size / 2}
	//max := size * size

	var (
		crds []set.Coord
		err  error
		res  set.Int64
	)

	i, inc := int64(1), int64(1)
	for {
		res, crds, err = grid.Vector(crd, i+inc, set.UP)
		fmt.Println(res)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		//i = grid[crd.Row][crd.Col]

		fmt.Println("1:", crd, i, inc)

		res, crds, err = grid.Vector(crd, i+inc, set.LTR)
		fmt.Println(res)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		//i = grid[crd.Row][crd.Col]
		inc++

		fmt.Println("2:", crd, i, inc)

		res, crds, err = grid.Vector(crd, i+inc, set.DOWN)
		fmt.Println(res)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		//i = grid[crd.Row][crd.Col]

		fmt.Println("3:", crd, i, inc)

		res, crds, err = grid.Vector(crd, i+inc, set.RTL)
		fmt.Println(res)
		if err != nil {
			break
		}
		crd = crds[len(crds)-1]
		//i = grid[crd.Row][crd.Col]
		inc++

		fmt.Println("4:", crd, i, inc)
	}

	return grid
}
