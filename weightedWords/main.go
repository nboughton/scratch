package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"

	"github.com/jmcvetta/randutil"
)

var (
	notLetter = regexp.MustCompile("[^a-zA-Z]")
	vowel     = regexp.MustCompile("(?i)([aeiouy]+)")
	consonant = regexp.MustCompile("(?i)([^aeiouy]+)")
)

type pho struct {
	chars  string
	count  int
	weight int
}

type byCount []pho

func (b byCount) Len() int           { return len(b) }
func (b byCount) Less(i, j int) bool { return b[i].count > b[j].count }
func (b byCount) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byCount) totalCount() int {
	t := 0

	for _, v := range b {
		t += v.count
	}

	return t
}

var (
	vwls = make(map[string]int)
	cons = make(map[string]int)
	vtot = 0.
	ctot = 0.
	keep = 41
	base = 10
)

func main() {
	b, err := ioutil.ReadFile("lovecraft.txt")
	if err != nil {
		panic(err)
	}

	text := strings.ToLower(notLetter.ReplaceAllString(string(b), " "))

	for _, word := range strings.Fields(text) {
		if len(word) > 3 {
			breakdown(word)
		}
	}

	vpho := byCount{}
	for chars, count := range vwls {
		vpho = append(vpho, pho{chars, count, int(float64(count) / vtot * 10000)})
	}
	fmt.Println("vpho: ", len(vpho))

	vcon := byCount{}
	for chars, count := range cons {
		vcon = append(vcon, pho{chars, count, int(float64(count) / ctot * 10000)})
	}
	fmt.Println("vcon: ", len(vcon))

	sort.Sort(vpho)
	sort.Sort(vcon)

	weightedV := []randutil.Choice{}
	for _, v := range vpho {
		fmt.Printf("{Weight: %d, Item: \"%s\"},\n", v.count, v.chars)
		weightedV = append(weightedV, randutil.Choice{Weight: v.count, Item: v.chars})
	}

	fmt.Println(randutil.WeightedChoice(weightedV))

	weightedC := []randutil.Choice{}
	for _, c := range vcon {
		fmt.Printf("{Weight: %d, Item: \"%s\"},\n", c.count, c.chars)
		weightedC = append(weightedC, randutil.Choice{Weight: c.count, Item: c.chars})
	}

	/*
		vlist := bellcurve(vpho[:keep])
		fmt.Println("vlist: ", len(vlist))

		clist := bellcurve(vcon[:keep])
		fmt.Println("clist: ", len(clist))

		for i, c := range vlist {
			fmt.Printf("{Match: []int{%d}, Text: \"%s\"},\n", i+base, c)
		}

		fmt.Println()

		for i, c := range clist {
			fmt.Printf("{Match: []int{%d}, Text: \"%s\"},\n", i+base, c)
		}
	*/
}

func breakdown(word string) {
	for _, con := range consonant.FindAllString(word, -1) {
		if len(con) < 3 {
			cons[con]++
			ctot++
		}
	}

	for _, vow := range vowel.FindAllString(word, -1) {
		if len(vow) < 3 {
			vwls[vow]++
			vtot++
		}
	}
}

func bellcurve(list byCount) []string {
	out := []string{}

	for i, v := range list {
		if i%2 == 0 {
			out = append(out, v.chars)
		} else {
			out = append([]string{v.chars}, out...)
		}
	}

	return out
}
