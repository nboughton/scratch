package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		f, err := os.Lstat(os.Args[1])
		if err != nil {
			panic(err)
		}

		if f.Mode()&os.ModeSymlink != 0 {
			fmt.Println("is symlink")
		} else {
			fmt.Println("Is not symlink")
		}
	}
}
