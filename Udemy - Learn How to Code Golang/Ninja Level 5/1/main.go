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

	slice := []string{p1.icecream, p2.icecream}
	fmt.Println(slice)

}
