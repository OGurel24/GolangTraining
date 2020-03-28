package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 122; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c\t", i)
		}
		fmt.Println("")
	}
}
