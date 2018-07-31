package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/nboughton/rollt"
)

var vl = rollt.Table{
	Name: "vowels",
	Dice: "1d4,1d3",
	Items: []rollt.Item{
		{Match: []int{2}, Text: "y"},
		{Match: []int{3}, Text: "a"},
		{Match: []int{4}, Text: "e"},
		{Match: []int{5}, Text: "i"},
		{Match: []int{6}, Text: "o"},
		{Match: []int{7}, Text: "u"},
	},
}

var con = rollt.List{
	Name: "consonants",
	Items: []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x",
		"ch", "gh", "kh", "ph", "rh", "sh", "th", "qu", "ck"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	name, ln := "", rand.Intn(10)+1
	for i := 0; i <= ln; i++ {
		if i%2 != 0 {
			name += con.Roll()
		} else {
			name += vl.Roll()
		}
	}

	fmt.Println(strings.ToUpper(string(name[0])) + string(name[1:]))
}

func genName(ln int) string {
	name := ""
	for i := 0; i <= ln; i++ {
		if i%2 != 0 {
			name += con.Roll()
		} else {
			name += vl.Roll()
		}
	}

	return strings.ToUpper(string(name[0])) + string(name[1:])
}
