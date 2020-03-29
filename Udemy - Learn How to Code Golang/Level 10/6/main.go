package main

import (
	"fmt"
)

func main() {

	cap := 100
	c := make(chan int)

	go func() {
		for i := 0; i < cap; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

}
