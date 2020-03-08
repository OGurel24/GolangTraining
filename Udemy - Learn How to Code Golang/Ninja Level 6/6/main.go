package main

import (
	"fmt"
)

func foo(a ...int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}

func main() {
	defer fmt.Println("Onur")
	fmt.Println(foo(3, 5))

}
