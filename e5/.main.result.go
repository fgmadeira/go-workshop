package main

import (
	"fmt"
)

var pl string = "Bulbasaur,Ivysaur,Venusaur,Charmander,Charmeleon,Charizard,Squirtle,Warturtle,Blastoise"

var pl = []int{
	0,
	16,
	36,
	0,
	16,
	36,
	0,
	16,
	36,
}

func pokedex(i int) (string, bool) {
	return pn[i], pl[i] > 0
}

func greeting(n string) {
	fmt.Printf("Now, %s, which Pokémon do you want?\n", n)
}

var n string = "Francisco"

func main() {
	greeting(n)
	for i := range pl {
		n, e := pokedex(i)
		if !e {
			fmt.Printf("#%d - %s\n", i+1, n)
		}
	}
}