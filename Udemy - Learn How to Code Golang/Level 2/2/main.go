package main

import (
	"fmt"
)

func main() {
	v1 := 24 < 25
	v2 := 26 <= 23
	v3 := 26 > 23
	v4 := 30 >= 23
	v5 := 26 != 26
	fmt.Println(v1, v2, v3, v4, v5)

}
