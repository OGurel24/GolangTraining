package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	age   int64
}

func main() {
	p1 := person{
		First: "Onur",
		Last:  "Gurel",
		age:   25,
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
