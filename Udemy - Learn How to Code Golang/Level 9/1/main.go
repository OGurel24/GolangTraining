package main

import (
	"fmt"
	"sync"
)

func onur(a *sync.WaitGroup){
	fmt.Println("Onur")
	a.Done()
	
}

func onur2(a *sync.WaitGroup){
	fmt.Println("Onur2")
	a.Done()
	
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	
	go onur(&wg)
	wg.Wait()
	
	wg.Add(1)
	go onur2(&wg)
	wg.Wait()
	
}
