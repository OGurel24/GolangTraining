package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name  string
	lname string
	Id    int64
}

func main() {

	p1 := person{"Onur", "Gurel", 24}
	p2 := person{"Ugur", "Gurel", 7}
	p3 := person{"Muslera", "Nestor Fernando", 1}

	team := []person{p1, p2, p3}
	fmt.Println(team)

	json.NewEncoder(os.Stdout).Encode(team)

}
