package main

import (
	"flag"
	"fmt"

	"github.com/nboughton/go-dice"
)

var (
	d8         = "2d8"
	damageType = map[int]string{
		1: "Acid",
		2: "Cold",
		3: "Fire",
		4: "Force",
		5: "Lightning",
		6: "Poison",
		7: "Psychic",
		8: "Thunder",
	}
)

func main() {
	level := flag.Int("level", 1, "Set level at which the spell is cast")
	flag.Parse()

	bag, err := dice.NewBag(d8, fmt.Sprintf("%dd6", *level))
	if err != nil {
		panic(err)
	}

	dmg, rolls := bag.Roll()
	d8s := rolls[d8]
	if d8s[0] == d8s[1] {
		fmt.Printf("%d %s damage, attack again.\n", dmg, damageType[d8s[0]])
	} else {
		fmt.Printf("%d %s or %s damage.\n", dmg, damageType[d8s[0]], damageType[d8s[1]])
	}
}
