package main

import (
	"fmt"
)

func main() {

	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice1 = append(slice1, slice1[3:6]...)
	fmt.Println(slice1)

}
