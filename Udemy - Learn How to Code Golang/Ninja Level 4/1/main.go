package main

import (
	"fmt"
)

func main() {

	array := [5]int{1, 2, 3}
	for i, v := range array {
		fmt.Printf("index:%v\tvalue:%v\n", i, v)
	}
}
