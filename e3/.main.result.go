package main

import (
	"fmt"
)

var pl = []string{"Bulbasaur","Ivysaur","Venusaur","Charmander","Charmeleon","Charizard","Squirtle","Wartortle","Blastoise"}

func main() {
	b := pl[0:3]
	fmt.Println(b)

	c := pl[3:6]
	fmt.Println(c)

	s := pl[6:9]
	fmt.Println(s)

	//as := make([]string, 3)
	//s = append(s, as...)
	//fmt.Println(len(s))
}