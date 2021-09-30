package main

import (
	"fmt"
	"strings"
)

var pl string = "Bulbasaur,Ivysaur,Venusaur,Charmander,Charmeleon,Charizard,Squirtle,Warturtle,Blastoise"

func main() {
	ps := strings.Split(pl, ",")

	//for i := 0; i < 9; i++ {
	for i, p := range ps {
		fmt.Printf("#%d - %s\n", i+1, p)
	}
}
