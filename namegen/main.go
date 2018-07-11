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
	Dice: "3d12",
	Items: []rollt.Item{
		{Match: []int{3}, Text: "y"},
		{Match: []int{4}, Text: "ya"},
		{Match: []int{5}, Text: "ue"},
		{Match: []int{6}, Text: "ey"},
		{Match: []int{7}, Text: "ei"},
		{Match: []int{8}, Text: "ao"},
		{Match: []int{9}, Text: "ae"},
		{Match: []int{10}, Text: "ai"},
		{Match: []int{11}, Text: "eo"},
		{Match: []int{12}, Text: "eu"},
		{Match: []int{13}, Text: "iu"},
		{Match: []int{14}, Text: "oa"},
		{Match: []int{15}, Text: "oo"},
		{Match: []int{16}, Text: "ie"},
		{Match: []int{17}, Text: "ea"},
		{Match: []int{18}, Text: "a"},
		{Match: []int{19}, Text: "e"},
		{Match: []int{20}, Text: "i"},
		{Match: []int{21}, Text: "o"},
		{Match: []int{22}, Text: "u"},
		{Match: []int{23}, Text: "ou"},
		{Match: []int{24}, Text: "io"},
		{Match: []int{25}, Text: "ia"},
		{Match: []int{26}, Text: "oi"},
		{Match: []int{27}, Text: "ay"},
		{Match: []int{28}, Text: "oe"},
		{Match: []int{29}, Text: "oy"},
		{Match: []int{30}, Text: "au"},
		{Match: []int{31}, Text: "ui"},
		{Match: []int{32}, Text: "uo"},
		{Match: []int{33}, Text: "ua"},
		{Match: []int{34}, Text: "yo"},
		{Match: []int{35}, Text: "yi"},
		{Match: []int{36}, Text: "yu"},
	},
}

/*
"y","ya","ue","ey","ei","ao",

"ae","ai","eo","eu","iu","oa","oo",

"ie","ea","a","e","i","o","u","ou",

"io","ia","oi","ay","oe","ui","au",

"uo","ua","oy","yo","yi","yu",
*/

var con = rollt.List{
	Name: "consonants",
	Items: []string{
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z",
		"bh", "bl", "br",
		"ch", "cl", "cr",
		"dh", "dr",
		"fl", "fr",
		"gh", "gl", "gr", "gn",
		"kh", "kl", "kr", "kn",
		"nk",
		"ph", "pl", "pr",
		"qu",
		"sh", "sl", "st",
		"th"},
}

var badSuffix = []string{
	"bl", "br", "cl", "cr", "dh", "dr", "fl", "fr", "gl", "gr", "gn", "iw", "kl", "kr", "kn",
	"pl", "pr", "qu", "sl",
}

// Generate creates a random name by combining alternating vowels and consonants
func Generate(ln int) string {
	name := ""
	for i := rand.Intn(2); i <= ln; i++ {
		if i%2 != 0 {
			name += con.Roll()
		} else {
			name += vl.Roll()
		}
	}

	for _, suf := range badSuffix {
		if string(name[len(name)-2:]) == suf {
			name = string(name[:len(name)-2])
		}
	}

	return strings.ToUpper(string(name[0])) + string(name[1:])
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		fmt.Println(Generate(rand.Intn(3) + 3))
	}
}
