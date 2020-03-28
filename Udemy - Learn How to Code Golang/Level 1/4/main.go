package main

import (
	"fmt"
)

type onur int

var v1 onur = 24

func main() {
	fmt.Printf("%v\t%T", v1, v1)
}
