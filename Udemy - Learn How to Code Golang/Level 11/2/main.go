package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "Onur",
		Last:    "Gurel",
		Sayings: []string{"QMECHANICS FOREVER!"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln("CRASHED!")
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("Error caught by Onur Gurel")
	}

	return bs, nil
}
