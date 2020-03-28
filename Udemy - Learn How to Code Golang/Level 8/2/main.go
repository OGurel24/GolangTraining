package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var familyJson = []byte(`[
	{"Name": "Onur", "ID": 24},
	{"Name": "UGur",    "ID": 7}
]`)
	type animal struct {
		Name string
		ID   int64
	}
	var family []animal
	err := json.Unmarshal(familyJson, &family)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", family)
}
