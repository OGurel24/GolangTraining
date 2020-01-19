package main

import (
	"fmt"
)

type person struct {
	name     string
	surname  string
	icecream string
}

func main() {
	p1 := person{
		name:     "Onur",
		surname:  "Gurel",
		icecream: "Chocalate",
	}
	p2 := person{
		name:     "Ugur",
		surname:  "Gurel",
		icecream: "Lemon",
	}

	mymap := map[string]person{p1.name: p1, p2.name: p2}

	fmt.Println(mymap)

}
