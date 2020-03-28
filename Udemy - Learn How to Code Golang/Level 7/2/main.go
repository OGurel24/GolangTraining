package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func change(a *person, first string, last string, age int) {
	(*a).first = first
	(*a).last = last
	(*a).age = age
}

func main() {
	og := person{"onur", "gurel", 30}
	fmt.Println(og)
	change(&og, "Onur", "Gurel", 25)
	fmt.Println(og)

}
