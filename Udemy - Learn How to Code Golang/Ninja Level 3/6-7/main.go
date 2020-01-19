package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println("Muslera")
		} else if i%4 == 1 {
			fmt.Println("Falcao")
		} else {
			fmt.Println("Lemina")
		}

	}
}
