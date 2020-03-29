package main

import (
	"fmt"
)

func main() {
	c := make(chan int,1)
	c<-24
	
	v,ok:=<-c
	close(c)
	
	fmt.Println(v,ok)
}

