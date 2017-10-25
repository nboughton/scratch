package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	names := os.Args[2:]

	for i := 0; i < num; i++ {
		n := len(names)
		p := rand.Intn(n)

		fmt.Println(names[p])
		names = append(names[:p], names[p+1:]...)
	}
}
