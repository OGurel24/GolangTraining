package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 24, 55, 66, 45, 23}
	for i, v := range array {
		fmt.Printf("index:%v\tvalue:%v\n", i, v)
	}
}
