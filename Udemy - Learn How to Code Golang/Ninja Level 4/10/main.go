package main

import (
	"fmt"
)

func main() {

	my_map := map[string][]int{
		"GK": {1, 99},
		"CB": {3, 4, 6, 32, 51},
	}
	my_map["ST"] = []int{9, 11, 21, 99}
	fmt.Println(my_map)
	delete(my_map, "GK")
	fmt.Println(my_map)

}
