package main

import (
	"fmt"
)

func main() {
	v1 := struct {
		doors int
		price int
		lux   bool
	}{
		doors: 4,
		price: 100000,
		lux:   false,
	}

	fmt.Println(v1)
}
