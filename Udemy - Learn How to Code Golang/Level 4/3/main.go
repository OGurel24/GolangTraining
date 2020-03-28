package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 24, 55, 66, 45, 23}
	for i1, _ := range array {
		for i2, _ := range array {
			if i1 <= i2 {
				fmt.Println(array[i1:i2])
			}
		}
	}
}
