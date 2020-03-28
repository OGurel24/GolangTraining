package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name  string
	lname string
	id    int
	admin bool
}

func main() {
	o := user{
		Name:  "Onur",
		lname: "Gurel",
		id:    24,
		admin: true,
	}
	ex, _ := json.Marshal(o)
	fmt.Println(ex)
	fmt.Println(string(ex))

}

