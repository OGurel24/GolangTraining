package main

import (
	"fmt"
)

func foo(a int) int {
	return a
}

func bar(a int) (int, string) {
	return a, "OG24"
}

func main() {
	fmt.Println(foo(3))
	fmt.Println(bar(7))

}
