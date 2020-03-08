package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (a person) foo() (string, int) {
	return a.first, a.age
}

func main() {
	onur := person{
		first: "Onur",
		last:  "Gurel",
		age:   25,
	}
	fmt.Println(onur.foo())

}
