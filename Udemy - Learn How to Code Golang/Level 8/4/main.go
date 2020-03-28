package main

import (
	"fmt"
	"sort"
)

func main() {

	p1 := []int{0, -55, 34, 23, 99, 101, 4}
	p2 := []string{"Onur", "Ugur", "Muslera", "Falcao", "Onyekuru"}

	fmt.Println(p1)
	sort.Ints(p1)
	fmt.Println(p1)

	fmt.Println()

	fmt.Println(p2)
	sort.Strings(p2)
	fmt.Println(p2)

}
