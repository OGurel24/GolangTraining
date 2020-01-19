package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		s:= i%4
		switch s{
			case 1:
				fmt.Println("Muslera")
			case 2:
				fmt.Println("Lemina")
			case s:
				fmt.Println("Falcao")					
		}
	}
}
