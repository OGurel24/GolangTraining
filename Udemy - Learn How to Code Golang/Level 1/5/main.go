package main

import (
	"fmt"
)

type onur int

var v1 onur = 24
var v2 int

func main() {
	fmt.Printf("%v\t%T\n", v1, v1)
	v2 = int(v1)
	fmt.Printf("%v\t%T", v2, v2)

}
