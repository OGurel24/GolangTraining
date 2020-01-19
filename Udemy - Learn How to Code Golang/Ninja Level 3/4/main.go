package main

import (
	"fmt"
)

func main() {
	birthday := 1994
	year := 2020
	for {
		fmt.Println(birthday)
		birthday++
		if birthday > year {
			break
		}

	}
}
